package mappers

import (
	"backend/internal/dto"
	"backend/internal/models"
)

func FromPostToDto(post *models.Post) *dto.PostDto {
	postDto := new(dto.PostDto)
	postDto.Id = post.Id
	postDto.Name = post.Name
	return postDto
}

func FromCreateRequestDtoToPost(createRequestDto *dto.CreatePostRequestDto) *models.Post {
	post := new(models.Post)
	post.Name = createRequestDto.Name
	return post
}

func FromUpdateReqeustDtoToPost(updateRequestDto *dto.UpdatePostRequestDto) *models.Post {
	post := new(models.Post)
	post.Name = updateRequestDto.Name
	return post
}
