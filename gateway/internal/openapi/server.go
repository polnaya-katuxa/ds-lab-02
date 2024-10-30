package openapi

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	openapi_types "github.com/oapi-codegen/runtime/types"
	"github.com/polnaya-katuxa/ds-lab-02/gateway/internal/clients"
	"github.com/polnaya-katuxa/ds-lab-02/gateway/internal/generated/openapi"
	cars_service "github.com/polnaya-katuxa/ds-lab-02/gateway/internal/generated/openapi/clients/cars-service"
	payment_service "github.com/polnaya-katuxa/ds-lab-02/gateway/internal/generated/openapi/clients/payment-service"
	rental_service "github.com/polnaya-katuxa/ds-lab-02/gateway/internal/generated/openapi/clients/rental-service"
	"github.com/polnaya-katuxa/ds-lab-02/gateway/internal/models"
	"github.com/samber/lo"
)

type Server struct {
	cars    *clients.CarsServiceClient
	payment *clients.PaymentServiceClient
	rental  *clients.RentalServiceClient
}

func New(cars *clients.CarsServiceClient, payment *clients.PaymentServiceClient, rental *clients.RentalServiceClient) *Server {
	return &Server{
		cars:    cars,
		payment: payment,
		rental:  rental,
	}
}

func (s *Server) GetCars(c echo.Context, params openapi.GetCarsParams) error {
	cars, err := s.cars.List(c.Request().Context(), &cars_service.ListParams{
		Page:    lo.ToPtr(float32(lo.FromPtr(params.Page) - 1)),
		Size:    lo.ToPtr(float32(lo.FromPtr(params.Size))),
		ShowAll: params.ShowAll,
	})
	if err != nil {
		return processError(c, err, "list cars")
	}

	return c.JSON(http.StatusOK, cars)
}

func (s *Server) GetUserRentals(c echo.Context, params openapi.GetUserRentalsParams) error {
	rentals, err := s.rental.List(c.Request().Context(), params.XUserName)
	if err != nil {
		return processError(c, err, "list user rentals")
	}

	result := make([]openapi.RentalResponse, len(rentals))
	for i, rental := range rentals {
		car, err := s.cars.Get(c.Request().Context(), rental.CarUid)
		if err != nil {
			return processError(c, err, "get car info")
		}

		payment, err := s.payment.Get(c.Request().Context(), rental.PaymentUid)
		if err != nil {
			return processError(c, err, "get payment info")
		}

		result[i] = openapi.RentalResponse{
			Car: openapi.CarInfo{
				Brand:              car.Brand,
				CarUid:             car.CarUid,
				Model:              car.Model,
				RegistrationNumber: car.RegistrationNumber,
			},
			DateFrom: rental.DateFrom,
			DateTo:   rental.DateTo,
			Payment: openapi.PaymentInfo{
				PaymentUid: payment.PaymentUid,
				Price:      payment.Price,
				Status:     openapi.PaymentInfoStatus(payment.Status),
			},
			RentalUid: rental.RentalUid,
			Status:    openapi.RentalResponseStatus(rental.Status),
		}
	}

	return c.JSON(http.StatusOK, result)
}

func (s *Server) BookCar(c echo.Context, params openapi.BookCarParams) error {
	var req openapi.BookCarJSONRequestBody
	err := json.NewDecoder(c.Request().Body).Decode(&req)
	if err != nil {
		return processError(c, err, "cannot unmarshal request body")
	}

	dateFrom, err := time.Parse(time.DateOnly, req.DateFrom)
	if err != nil {
		return processError(c, models.ValidationError{Message: err.Error()}, "parse date from")
	}

	dateTo, err := time.Parse(time.DateOnly, req.DateTo)
	if err != nil {
		return processError(c, models.ValidationError{Message: err.Error()}, "parse date to")
	}

	numDays := int(dateTo.Sub(dateFrom).Hours()) / 24
	if numDays < 1 {
		return processError(c, models.ValidationError{Message: "should rent min to 1 day"}, "check rent dates")
	}

	car, err := s.cars.Get(c.Request().Context(), req.CarUid)
	if err != nil {
		return processError(c, err, "get car")
	}

	_, err = s.cars.Book(c.Request().Context(), car.CarUid)
	if err != nil {
		return processError(c, err, "book car")
	}

	totalPrice := car.Price * numDays
	payment, err := s.payment.Create(c.Request().Context(), payment_service.CreatePaymentRequest{
		Price: totalPrice,
	})
	if err != nil {
		return processError(c, err, "create payment")
	}

	rental, err := s.rental.Create(c.Request().Context(), params.XUserName, rental_service.CreateRentalRequest{
		CarUid:     car.CarUid,
		DateFrom:   req.DateFrom,
		DateTo:     req.DateTo,
		PaymentUid: payment.PaymentUid,
	})
	if err != nil {
		return processError(c, err, "create rental")
	}

	result := openapi.CreateRentalResponse{

		CarUid:   car.CarUid,
		DateFrom: rental.DateFrom,
		DateTo:   rental.DateTo,
		Payment: openapi.PaymentInfo{
			PaymentUid: payment.PaymentUid,
			Price:      payment.Price,
			Status:     openapi.PaymentInfoStatus(payment.Status),
		},
		RentalUid: rental.RentalUid,
		Status:    openapi.CreateRentalResponseStatus(rental.Status),
	}

	return c.JSON(http.StatusOK, result)
}

func (s *Server) CancelRental(c echo.Context, rentalUid openapi_types.UUID, params openapi.CancelRentalParams) error {
	rental, err := s.rental.Get(c.Request().Context(), params.XUserName, rentalUid)
	if err != nil {
		return processError(c, err, "get user rental")
	}

	err = s.cars.Unbook(c.Request().Context(), rental.CarUid)
	if err != nil {
		return processError(c, err, "make car available")
	}

	err = s.rental.Cancel(c.Request().Context(), params.XUserName, rentalUid)
	if err != nil {
		return processError(c, err, "cancel rental")
	}

	err = s.payment.Cancel(c.Request().Context(), rental.PaymentUid)
	if err != nil {
		return processError(c, err, "cancel payment")
	}

	return c.NoContent(http.StatusNoContent)
}

func (s *Server) GetUserRental(c echo.Context, rentalUid openapi_types.UUID, params openapi.GetUserRentalParams) error {
	rental, err := s.rental.Get(c.Request().Context(), params.XUserName, rentalUid)
	if err != nil {
		return processError(c, err, "get user rental")
	}

	car, err := s.cars.Get(c.Request().Context(), rental.CarUid)
	if err != nil {
		return processError(c, err, "get car info")
	}

	payment, err := s.payment.Get(c.Request().Context(), rental.PaymentUid)
	if err != nil {
		return processError(c, err, "get payment info")
	}

	result := openapi.RentalResponse{
		Car: openapi.CarInfo{
			Brand:              car.Brand,
			CarUid:             car.CarUid,
			Model:              car.Model,
			RegistrationNumber: car.RegistrationNumber,
		},
		DateFrom: rental.DateFrom,
		DateTo:   rental.DateTo,
		Payment: openapi.PaymentInfo{
			PaymentUid: payment.PaymentUid,
			Price:      payment.Price,
			Status:     openapi.PaymentInfoStatus(payment.Status),
		},
		RentalUid: rental.RentalUid,
		Status:    openapi.RentalResponseStatus(rental.Status),
	}

	return c.JSON(http.StatusOK, result)
}

func (s *Server) FinishRental(c echo.Context, rentalUid openapi_types.UUID, params openapi.FinishRentalParams) error {
	rental, err := s.rental.Get(c.Request().Context(), params.XUserName, rentalUid)
	if err != nil {
		return processError(c, err, "get user rental")
	}

	err = s.cars.Unbook(c.Request().Context(), rental.CarUid)
	if err != nil {
		return processError(c, err, "make car available")
	}

	err = s.rental.Finish(c.Request().Context(), params.XUserName, rentalUid)
	if err != nil {
		return processError(c, err, "finish rental")
	}

	return c.NoContent(http.StatusNoContent)
}

func (s *Server) Live(c echo.Context) error {
	return c.NoContent(http.StatusOK)
}