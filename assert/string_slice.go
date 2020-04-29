package assert

import (
	"testing"

	"github.com/ppapapetrou76/go-testing/internal/pkg/types"
)

// SliceOpt is a configuration option to initialize an Assertable Slice
type SliceOpt func(*AssertableStringSlice)

// AssertableSlice is the implementation of CommonAssertable for string types
type AssertableStringSlice struct {
	t             *testing.T
	actual        types.Containable
	customMessage string
}

// WithCustomMessage provides a custom message to be added before the assertion error message
func WithCustomMessage(customMessage string) SliceOpt {
	return func(c *AssertableStringSlice) {
		c.customMessage = customMessage
	}
}

// ThatSlice returns a proper assertable structure based on the slice type
func ThatSlice(t *testing.T, actual interface{}, opts ...SliceOpt) AssertableSlice {
	switch v := actual.(type) {
	case []string:
		assertable := &AssertableStringSlice{
			t:      t,
			actual: types.NewStringSliceValue(actual),
		}
		for _, opt := range opts {
			opt(assertable)
		}
		return assertable
	default:
		t.Fatalf("only string slice type is supported but got %T type", v)
	}
	return nil
}

// IsEqualTo asserts if the expected slice is equal to the assertable slice value
// It errors the tests if the compared values (actual VS expected) are not equal
func (a AssertableStringSlice) IsEqualTo(expected interface{}) AssertableSlice {
	if a.actual.IsNotEqualTo(expected) {
		a.t.Error(ShouldBeEqual(a.actual, expected))
	}
	return a
}

// IsNotEqualTo asserts if the expected slice is not equal to the assertable slice value
// It errors the tests if the compared values (actual VS expected) are equal
func (a AssertableStringSlice) IsNotEqualTo(expected interface{}) AssertableSlice {
	if a.actual.IsEqualTo(expected) {
		a.t.Error(ShouldNotBeEqual(a.actual, expected))
	}
	return a
}

// HasSize asserts if the assertable string slice has the expected length size
// It errors the test if it doesn't have the expected size
func (a AssertableStringSlice) HasSize(size int) AssertableSlice {
	if !a.actual.HasSize(size) {
		a.t.Error(ShouldHaveSize(a.actual, size))
	}
	return a
}

// HasSize asserts if the assertable string slice has the expected length size
// It errors the test if it doesn't have the expected size
func (a AssertableStringSlice) IsEmpty() AssertableSlice {
	if a.actual.IsEmpty() {
		a.t.Error(ShouldBeEmpty(a.actual))
	}
	return a
}

func (a AssertableStringSlice) IsNotEmpty() AssertableSlice {
	if a.actual.IsNotEmpty() {
		a.t.Error(ShouldNotBeEmpty(a.actual))
	}
	return a
}

// Contains asserts if the assertable string slice contains the given element(s)
// It errors the test if it does not contain it/them
func (a AssertableStringSlice) Contains(elements interface{}) AssertableSlice {
	if a.actual.DoesNotContain(elements) {
		a.t.Error(ShouldContain(a.actual, elements))
	}
	return a
}

// Contains asserts if the assertable string slice contains only the given element(s)
// It errors the test if it does not contain it/them
func (a AssertableStringSlice) ContainsOnly(elements interface{}) AssertableSlice {
	if !a.actual.ContainsOnly(elements) {
		a.t.Error(ShouldContainOnly(a.actual, elements))
	}
	return a
}

// Contains asserts if the assertable string slice does not contain the given element
// It errors the test if it contains it/them
func (a AssertableStringSlice) DoesNotContain(elements interface{}) AssertableSlice {
	if a.actual.Contains(elements) {
		a.t.Error(ShouldNotContain(a.actual, elements))
	}
	return a
}
