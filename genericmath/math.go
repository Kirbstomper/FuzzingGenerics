package genericmath

type Number interface {
	int | int32 | int64 |float32 | float64
}

//Adds two numbers of the same type
func Add[T Number](a T, b T) T {
	return (a + b)
}
