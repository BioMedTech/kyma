// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"
import v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
import v1alpha1 "github.com/kyma-project/kyma/components/application-operator/pkg/apis/applicationconnector/v1alpha1"

// ApplicationClient is an autogenerated mock type for the ApplicationClient type
type ApplicationClient struct {
	mock.Mock
}

// List provides a mock function with given fields: opts
func (_m *ApplicationClient) List(opts v1.ListOptions) (*v1alpha1.ApplicationList, error) {
	ret := _m.Called(opts)

	var r0 *v1alpha1.ApplicationList
	if rf, ok := ret.Get(0).(func(v1.ListOptions) *v1alpha1.ApplicationList); ok {
		r0 = rf(opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1alpha1.ApplicationList)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(v1.ListOptions) error); ok {
		r1 = rf(opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: _a0
func (_m *ApplicationClient) Update(_a0 *v1alpha1.Application) (*v1alpha1.Application, error) {
	ret := _m.Called(_a0)

	var r0 *v1alpha1.Application
	if rf, ok := ret.Get(0).(func(*v1alpha1.Application) *v1alpha1.Application); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1alpha1.Application)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*v1alpha1.Application) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
