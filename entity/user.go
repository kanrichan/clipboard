package entity

type NewUserRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type NewUserResponse struct {
	ID int64 `json:"id"`
}

type GetAllUsersRequest struct {
}

type GetAllUsersResponse struct {
	Users []struct {
		ID       int64  `json:"id"`
		Username string `json:"username"`
	} `json:"users"`
}
