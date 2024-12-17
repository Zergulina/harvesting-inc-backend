package dto

type CreatePostRequestDto struct {
	Name string `json:"name"`
}

type PostDto struct {
	Id   uint64 `json:"id"`
	Name string `json:"name"`
}

type UpdatePostRequestDto struct {
	Name string `json:"name"`
}
