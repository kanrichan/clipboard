package entity

type NewSessionRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type NewSessionResponse struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
	Token    string `json:"token"`
}

type GetSessionRequest struct {
}

type GetSessionResponse struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
}

type UpdSessionRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UpdSessionResponse struct {
}

type DelSessionRequest struct {
}

type DelSessionResponse struct {
}
