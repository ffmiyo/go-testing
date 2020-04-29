package assert

type AssertableSlice interface {
	IsEqualTo(expected interface{}) AssertableSlice
	IsNotEqualTo(expected interface{}) AssertableSlice
	IsEmpty() AssertableSlice
	IsNotEmpty() AssertableSlice
	HasSize(size int) AssertableSlice
	Contains(elements interface{}) AssertableSlice
	ContainsOnly(elements interface{}) AssertableSlice
	DoesNotContain(elements interface{}) AssertableSlice
}
