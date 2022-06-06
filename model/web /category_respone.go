package web

type CategoryRespone struct {
	Id int `validate:"required"`
	Name string `validate:"required,min=1,max=100"`
}
