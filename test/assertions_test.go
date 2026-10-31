package xi_sdk_resellers

import (
	"reflect"
	"testing"
)

type testAssertions struct {
	fatal bool
}

var (
	assert  = testAssertions{}
	require = testAssertions{fatal: true}
)

func (a testAssertions) Equal(t *testing.T, expected, actual any) bool {
	t.Helper()
	if reflect.DeepEqual(expected, actual) {
		return true
	}

	a.fail(t, "expected %v, got %v", expected, actual)
	return false
}

func (a testAssertions) Nil(t *testing.T, actual any) bool {
	t.Helper()
	if isNil(actual) {
		return true
	}

	a.fail(t, "expected nil, got %v", actual)
	return false
}

func (a testAssertions) NotNil(t *testing.T, actual any) bool {
	t.Helper()
	if !isNil(actual) {
		return true
	}

	a.fail(t, "expected a non-nil value")
	return false
}

func (a testAssertions) fail(t *testing.T, format string, args ...any) {
	t.Helper()
	if a.fatal {
		t.Fatalf(format, args...)
		return
	}

	t.Errorf(format, args...)
}

func isNil(value any) bool {
	if value == nil {
		return true
	}

	v := reflect.ValueOf(value)
	switch v.Kind() {
	case reflect.Chan, reflect.Func, reflect.Interface, reflect.Map, reflect.Ptr, reflect.Slice:
		return v.IsNil()
	default:
		return false
	}
}
