package dto

type Category struct {
	ID   int64  `db:"id" form:"id"`
	Name string `db:"name" form:"name"`
}

type ReqCreateCategory struct {
	Name string `json:"name" form:"name" validate:"required"`
}

type ReqUpdateCategory struct {
	ID   int64  `json:"id" form:"id" validate:"required,numeric"`
	Name string `json:"name" form:"name" validate:"required"`
}

type ReqDeleteCategory struct {
	ID int64 `json:"id" form:"id" validate:"required,numeric"`
}

type ResponseCategory struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

func (c *Category) ToResponseCategory() ResponseCategory {
	return ResponseCategory{
		ID:   c.ID,
		Name: c.Name,
	}
}

func (c *Category) SetName(name *string) {
	c.Name = *name
}

func (c *Category) SetID(id *int64) {
	c.ID = *id
}

func (c *Category) GetID() *int64 {
	return &c.ID
}

func (c *Category) GetName() *string {
	return &c.Name
}
