// Code generated by mockery v2.26.1. DO NOT EDIT.

package mocks

import (
	gin "github.com/gin-gonic/gin"

	mock "github.com/stretchr/testify/mock"
)

// IProductHandler is an autogenerated mock type for the IProductHandler type
type IProductHandler struct {
	mock.Mock
}

// AddProduct provides a mock function with given fields: c
func (_m *IProductHandler) AddProduct(c *gin.Context) {
	_m.Called(c)
}

// GetProducts provides a mock function with given fields: c
func (_m *IProductHandler) GetProducts(c *gin.Context) {
	_m.Called(c)
}

type mockConstructorTestingTNewIProductHandler interface {
	mock.TestingT
	Cleanup(func())
}

// NewIProductHandler creates a new instance of IProductHandler. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewIProductHandler(t mockConstructorTestingTNewIProductHandler) *IProductHandler {
	mock := &IProductHandler{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}