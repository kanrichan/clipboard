package entity

type NewPostRequest struct {
	Content string `json:"content" binding:"required"`
}

type NewPostResponse struct {
	ID int64 `json:"id"`
}

type GetPostByIDRequest struct {
}

type GetPostByIDResponse struct {
	ID      int64  `json:"id"`
	Content string `json:"content"`
	Author  struct {
		ID       int64  `json:"id"`
		Username string `json:"username"`
	} `json:"author"`
}

type UpdPostByIDRequest struct {
	Content string `json:"content"`
}

type UpdPostByIDResponse struct {
}

type DelPostByIDRequest struct {
}

type DelPostByIDResponse struct {
}
