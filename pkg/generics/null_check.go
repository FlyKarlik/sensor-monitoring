package generics

func NullCheck[T any](value *T) *T {
	if value == nil {
		return nil
	}
	return value
}
