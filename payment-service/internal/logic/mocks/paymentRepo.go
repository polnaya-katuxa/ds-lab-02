// Code generated by mockery v2.45.1. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	models "github.com/polnaya-katuxa/ds-lab-02/payment-service/internal/models"

	uuid "github.com/google/uuid"
)

// PaymentRepo is an autogenerated mock type for the paymentRepo type
type PaymentRepo struct {
	mock.Mock
}

type PaymentRepo_Expecter struct {
	mock *mock.Mock
}

func (_m *PaymentRepo) EXPECT() *PaymentRepo_Expecter {
	return &PaymentRepo_Expecter{mock: &_m.Mock}
}

// ChangeStatus provides a mock function with given fields: ctx, uid, status
func (_m *PaymentRepo) ChangeStatus(ctx context.Context, uid uuid.UUID, status models.PaymentStatus) error {
	ret := _m.Called(ctx, uid, status)

	if len(ret) == 0 {
		panic("no return value specified for ChangeStatus")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID, models.PaymentStatus) error); ok {
		r0 = rf(ctx, uid, status)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// PaymentRepo_ChangeStatus_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ChangeStatus'
type PaymentRepo_ChangeStatus_Call struct {
	*mock.Call
}

// ChangeStatus is a helper method to define mock.On call
//   - ctx context.Context
//   - uid uuid.UUID
//   - status models.PaymentStatus
func (_e *PaymentRepo_Expecter) ChangeStatus(ctx interface{}, uid interface{}, status interface{}) *PaymentRepo_ChangeStatus_Call {
	return &PaymentRepo_ChangeStatus_Call{Call: _e.mock.On("ChangeStatus", ctx, uid, status)}
}

func (_c *PaymentRepo_ChangeStatus_Call) Run(run func(ctx context.Context, uid uuid.UUID, status models.PaymentStatus)) *PaymentRepo_ChangeStatus_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uuid.UUID), args[2].(models.PaymentStatus))
	})
	return _c
}

func (_c *PaymentRepo_ChangeStatus_Call) Return(_a0 error) *PaymentRepo_ChangeStatus_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *PaymentRepo_ChangeStatus_Call) RunAndReturn(run func(context.Context, uuid.UUID, models.PaymentStatus) error) *PaymentRepo_ChangeStatus_Call {
	_c.Call.Return(run)
	return _c
}

// Create provides a mock function with given fields: ctx, payment
func (_m *PaymentRepo) Create(ctx context.Context, payment models.Payment) (*models.Payment, error) {
	ret := _m.Called(ctx, payment)

	if len(ret) == 0 {
		panic("no return value specified for Create")
	}

	var r0 *models.Payment
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, models.Payment) (*models.Payment, error)); ok {
		return rf(ctx, payment)
	}
	if rf, ok := ret.Get(0).(func(context.Context, models.Payment) *models.Payment); ok {
		r0 = rf(ctx, payment)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Payment)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, models.Payment) error); ok {
		r1 = rf(ctx, payment)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PaymentRepo_Create_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Create'
type PaymentRepo_Create_Call struct {
	*mock.Call
}

// Create is a helper method to define mock.On call
//   - ctx context.Context
//   - payment models.Payment
func (_e *PaymentRepo_Expecter) Create(ctx interface{}, payment interface{}) *PaymentRepo_Create_Call {
	return &PaymentRepo_Create_Call{Call: _e.mock.On("Create", ctx, payment)}
}

func (_c *PaymentRepo_Create_Call) Run(run func(ctx context.Context, payment models.Payment)) *PaymentRepo_Create_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(models.Payment))
	})
	return _c
}

func (_c *PaymentRepo_Create_Call) Return(_a0 *models.Payment, _a1 error) *PaymentRepo_Create_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *PaymentRepo_Create_Call) RunAndReturn(run func(context.Context, models.Payment) (*models.Payment, error)) *PaymentRepo_Create_Call {
	_c.Call.Return(run)
	return _c
}

// Get provides a mock function with given fields: ctx, uid
func (_m *PaymentRepo) Get(ctx context.Context, uid uuid.UUID) (*models.Payment, error) {
	ret := _m.Called(ctx, uid)

	if len(ret) == 0 {
		panic("no return value specified for Get")
	}

	var r0 *models.Payment
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) (*models.Payment, error)); ok {
		return rf(ctx, uid)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) *models.Payment); ok {
		r0 = rf(ctx, uid)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Payment)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, uuid.UUID) error); ok {
		r1 = rf(ctx, uid)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PaymentRepo_Get_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Get'
type PaymentRepo_Get_Call struct {
	*mock.Call
}

// Get is a helper method to define mock.On call
//   - ctx context.Context
//   - uid uuid.UUID
func (_e *PaymentRepo_Expecter) Get(ctx interface{}, uid interface{}) *PaymentRepo_Get_Call {
	return &PaymentRepo_Get_Call{Call: _e.mock.On("Get", ctx, uid)}
}

func (_c *PaymentRepo_Get_Call) Run(run func(ctx context.Context, uid uuid.UUID)) *PaymentRepo_Get_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uuid.UUID))
	})
	return _c
}

func (_c *PaymentRepo_Get_Call) Return(_a0 *models.Payment, _a1 error) *PaymentRepo_Get_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *PaymentRepo_Get_Call) RunAndReturn(run func(context.Context, uuid.UUID) (*models.Payment, error)) *PaymentRepo_Get_Call {
	_c.Call.Return(run)
	return _c
}

// NewPaymentRepo creates a new instance of PaymentRepo. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewPaymentRepo(t interface {
	mock.TestingT
	Cleanup(func())
}) *PaymentRepo {
	mock := &PaymentRepo{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
