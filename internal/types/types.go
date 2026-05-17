package types

type Student struct {
	Id    int
	Name  string `validate:"required"` // means this field is required
	Email string `validate:"required"`
	Age   int    `validate:"required"`
}
