// Code generated by mockery v1.0.0
package automock

import mock "github.com/stretchr/testify/mock"
import unstructured "k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"

// GenericUsageBindingAnnotationTracer is an autogenerated mock type for the GenericUsageBindingAnnotationTracer type
type GenericUsageBindingAnnotationTracer struct {
	mock.Mock
}

// DeleteAnnotationAboutBindingUsage provides a mock function with given fields: res, usageName
func (_m *GenericUsageBindingAnnotationTracer) DeleteAnnotationAboutBindingUsage(res *unstructured.Unstructured, usageName string) error {
	ret := _m.Called(res, usageName)

	var r0 error
	if rf, ok := ret.Get(0).(func(*unstructured.Unstructured, string) error); ok {
		r0 = rf(res, usageName)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetInjectedLabels provides a mock function with given fields: res, usageName
func (_m *GenericUsageBindingAnnotationTracer) GetInjectedLabels(res *unstructured.Unstructured, usageName string) (map[string]string, error) {
	ret := _m.Called(res, usageName)

	var r0 map[string]string
	if rf, ok := ret.Get(0).(func(*unstructured.Unstructured, string) map[string]string); ok {
		r0 = rf(res, usageName)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]string)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*unstructured.Unstructured, string) error); ok {
		r1 = rf(res, usageName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SetAnnotationAboutBindingUsage provides a mock function with given fields: res, usageName, labels
func (_m *GenericUsageBindingAnnotationTracer) SetAnnotationAboutBindingUsage(res *unstructured.Unstructured, usageName string, labels map[string]string) error {
	ret := _m.Called(res, usageName, labels)

	var r0 error
	if rf, ok := ret.Get(0).(func(*unstructured.Unstructured, string, map[string]string) error); ok {
		r0 = rf(res, usageName, labels)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
