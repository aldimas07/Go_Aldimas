package mocks

import(
  users "Go_Aldimas/21_clean_architecture/clean_architecture/businesses/users"
  mock "github.com/stretchr/testify/mock"
  
)

type Usecase struct {
	mock.Mock
}

func (_m *Usecase) Login(userDomain *users.Domain) string {
	ret := _m.Called(userDomain)

	var r0 string
	if rf, ok := ret.Get(0).(func(*users.Domain) string); ok {
		r0 = rf(userDomain)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}


func (_m *Usecase) Register(userDomain *users.Domain) users.Domain {
	ret := _m.Called(userDomain)

	var r0 users.Domain
	if rf, ok := ret.Get(0).(func(*users.Domain) users.Domain); ok {
		r0 = rf(userDomain)
	} else {
		r0 = ret.Get(0).(users.Domain)
	}

	return r0
}

func (_m *Usecase) GetAll() []users.Domain {
	ret := _m.Called()

	var r0 []users.Domain
	if rf, ok := ret.Get(0).(func() []users.Domain); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]users.Domain)
		}
	}

	return r0
}

type mockConstructorTestingTNewUsecase interface {
	mock.TestingT
	Cleanup(func())
}

func NewUsecase(t mockConstructorTestingTNewUsecase) *Usecase {
	mock := &Usecase{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}