package repository

import (
	"vandyahmad24/maxsol/app/model"

	"github.com/opentracing/opentracing-go"
	"github.com/stretchr/testify/mock"
)

type MockCakeRepository struct {
	mock.Mock
}

func (m *MockCakeRepository) InsertCake(span opentracing.Span, input *model.Cake) (interface{}, error) {
	call := m.Called(input)
	res := call.Get(0)
	if res == nil {
		return nil, call.Error(1)
	}
	return call.Get(0), call.Error(1)
}

func (m *MockCakeRepository) GetAll(span opentracing.Span) (interface{}, error) {
	call := m.Called()
	res := call.Get(0)
	if res == nil {
		return nil, call.Error(1)
	}
	return call.Get(0), call.Error(1)
}

func (m *MockCakeRepository) Get(span opentracing.Span, id int) (interface{}, error) {
	call := m.Called(id)
	res := call.Get(0)
	if res == nil {
		return nil, call.Error(1)
	}
	return call.Get(0), call.Error(1)
}

func (m *MockCakeRepository) Delete(span opentracing.Span, id int) error {
	call := m.Called(id)
	res := call.Get(0)
	if res == nil {
		return call.Error(1)
	}
	return nil
}

func (m *MockCakeRepository) Update(span opentracing.Span, id int, input *model.Cake) (interface{}, error) {
	call := m.Called(input)
	res := call.Get(0)
	if res == nil {
		return nil, call.Error(1)
	}
	return call.Get(0), call.Error(1)
}
