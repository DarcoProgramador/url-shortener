// Code generated by mockery v2.37.1. DO NOT EDIT.

package mocks

import (
	url "github.com/DarcoProgramador/url-shortener/internal/url"
	mock "github.com/stretchr/testify/mock"
)

// Service is an autogenerated mock type for the Service type
type Service struct {
	mock.Mock
}

// GetUrl provides a mock function with given fields: shortUrl
func (_m *Service) GetUrl(shortUrl string) (string, error) {
	ret := _m.Called(shortUrl)

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (string, error)); ok {
		return rf(shortUrl)
	}
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(shortUrl)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(shortUrl)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SaveUrl provides a mock function with given fields: req
func (_m *Service) SaveUrl(req *url.UrlCreateRequest) (*url.UrlCreateResponse, error) {
	ret := _m.Called(req)

	var r0 *url.UrlCreateResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(*url.UrlCreateRequest) (*url.UrlCreateResponse, error)); ok {
		return rf(req)
	}
	if rf, ok := ret.Get(0).(func(*url.UrlCreateRequest) *url.UrlCreateResponse); ok {
		r0 = rf(req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*url.UrlCreateResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(*url.UrlCreateRequest) error); ok {
		r1 = rf(req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewService creates a new instance of Service. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewService(t interface {
	mock.TestingT
	Cleanup(func())
}) *Service {
	mock := &Service{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}