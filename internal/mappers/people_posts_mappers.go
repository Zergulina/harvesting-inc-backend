package mappers

import (
	"backend/internal/dto"
	"backend/internal/models"
)

func FromPeoplePostsToDto(people *models.People, posts []string) *dto.PeoplePostsDto {
	peoplePostsDto := new(dto.PeoplePostsDto)
	peoplePostsDto.Id = people.Id
	peoplePostsDto.LastName = people.LastName
	peoplePostsDto.FirstName = people.FirstName
	if people.MiddleName.Valid {
		peoplePostsDto.MiddleName = &people.MiddleName.String
	}
	peoplePostsDto.BirthDate = people.BirthDate
	peoplePostsDto.Posts = posts
	return peoplePostsDto
}
