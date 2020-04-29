package assert

import (
	"testing"

	"github.com/ppapapetrou76/go-testing/internal/pkg/types"
)

// AssertableBool is the assertable structure for bool values
type AssertableBool struct {
	t      *testing.T
	actual types.BoolValue
}

// ThatBool returns an AssertableBool structure initialized with the test reference and the actual bool value to assert
func ThatBool(t *testing.T, actual bool) AssertableBool {
	return AssertableBool{
		t:      t,
		actual: types.NewBoolValue(actual),
	}
}

// IsEqualTo asserts if the expected bool is equal to the assertable bool value
// It errors the tests if the compared values (actual VS expected) are not equal
func (a AssertableBool) IsEqualTo(expected interface{}) AssertableBool {
	if !a.actual.IsEqualTo(expected) {
		a.t.Error(ShouldBeEqual(a.actual, expected))
	}
	return a
}

// IsEqualTo asserts if the expected bool is not equal to the assertable bool value
// It errors the tests if the compared values (actual VS expected) are equal
func (a AssertableBool) IsNotEqualTo(expected interface{}) AssertableBool {
	if a.actual.IsEqualTo(expected) {
		a.t.Error(ShouldNotBeEqual(a.actual, expected))
	}
	return a
}

// IsTrue asserts if the expected bool value is true
func (a AssertableBool) IsTrue() AssertableBool {
	return a.IsEqualTo(true)
}

// IsFalse asserts if the expected bool value is false
func (a AssertableBool) IsFalse() AssertableBool {
	return a.IsEqualTo(false)
}
