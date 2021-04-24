package service

type GetTagsParam struct {
	Name  string `form:"name" binding:"max=100"`
	State uint8  `form:"state,default=1" binding:"oneof=0 1"`
}

type CreateTagParam struct {
	Name      string `form:"name" binding:"required,min=2,max=100"`
	CreatedBy string `form:"created_by" binding:"required,min=2,max=100"`
	State     uint8  `form:"state,default=1" binding:"oneof=0 1"`
}

type UpdateTagParam struct {
	ID         uint32 `form:"id" binding:"required,gte=1"`
	Name       string `form:"name" binding:"max=100"`
	State      uint8  `form:"state" binding:"oneof=0 1"`
	ModifiedBy string `form:"modified_by" binding:"required,min=2,max=100"`
}

type DeleteTagParam struct {
	ID uint32 `form:"id" binding:"required,gte=1"`
}
