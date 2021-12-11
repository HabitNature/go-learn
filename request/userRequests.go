package request

type UserCreate struct {
	Name string `json:"name" form:"name" comment:"姓名" validate:"required"`
}
