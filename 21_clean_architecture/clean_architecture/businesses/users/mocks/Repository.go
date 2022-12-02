package mocks

import(
  users "Go_Aldimas/21_clean_architecture/clean_architecture/businesses/users"
  mock "github.com/stretchr/testify/mock"

)

type Repository struct {
	mock.Mock
}


func (_m *Repository) GetByEmail(userDomain *users.Domain) users.Domain {
  ret := _m.Called(userDomain)

	var r0 users.Domain
	if rf, ok := ret.Get(0).(func(*users.Domain) users.Domain); ok {
		r0 = rf(userDomain)
	} else {
		r0 = ret.Get(0).(users.Domain)
	}

	return r0
}

func (_m *Repository) Register(userDomain *users.Domain) users.Domain {
	ret := _m.Called(userDomain)

	var r0 users.Domain
	if rf, ok := ret.Get(0).(func(*users.Domain) users.Domain); ok {
		r0 = rf(userDomain)
	} else {
		r0 = ret.Get(0).(users.Domain)
	}

	return r0
}


func (_m *Repository) GetAll() []users.Domain {
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



type mockConstructorTestingTNewRepository interface {
	mock.TestingT
	Cleanup(func())
}

func NewRepository(t mockConstructorTestingTNewRepository) *Repository {
	mock := &Repository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}