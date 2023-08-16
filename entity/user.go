package entity

type NewUserRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type NewUserResponse struct {
	ID int64 `json:"id"`
}

type GetUserByIDRequest struct {
}

type GetUserByIDResponse struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
}

type UpdUserByIDRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UpdUserByIDResponse struct {
}

type DelUserByIDRequest struct {
}

type DelUserByIDResponse struct {
}
