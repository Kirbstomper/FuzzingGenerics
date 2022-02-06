package genericmath

type Number interface {
	int | int32 | int64 | float32 | float64
}

func Add[T Number](a T, b T) T {
	return (a + b)
}
