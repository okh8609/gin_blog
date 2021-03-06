package service

type CreateAuthParam struct {
	UUID     string `form:"uuid" binding:"max=255"`
	Password string `form:"password" binding:"max=255"`
}

type VerifyAuthParam struct {
	UUID     string `form:"uuid" binding:"max=255"`
	Password string `form:"password" binding:"max=255"`
}

type UpdateAuthParam struct {
	UUID        string `form:"uuid" binding:"max=255"`
	Password    string `form:"password" binding:"max=255"`
	NewPassword string `form:"new_password" binding:"max=255"`
}

type DeleteAuthParam struct {
	UUID     string `form:"uuid" binding:"max=255"`
	Password string `form:"password" binding:"max=255"`
}
