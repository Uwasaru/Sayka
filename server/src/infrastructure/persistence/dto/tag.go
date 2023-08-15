package dto

import (
	"github.com/Uwasaru/Sayka/domain/entity"
)

type TagDto struct {
	ID   int `db:"id"`
	PostID string `db:"post_id"`
	Name string `db:"name"`
}

type TagsDto []*TagDto

func TagDtoToEntity(dto *TagDto) *entity.Tag {
	return &entity.Tag{
		ID:   dto.ID,
		PostID: dto.PostID,
		Name: dto.Name,
	}
}

func TagEntityToDto(t *entity.Tag) TagDto {
	return TagDto{
		ID:   t.ID,
		PostID: t.PostID,
		Name: t.Name,
	}
}

func TagsDtoToEntity(dtos *TagsDto) *entity.Tags {
	tags := &entity.Tags{}
	for _, dto := range *dtos {
		tag := TagDtoToEntity(dto)
		*tags = append(*tags, tag)
	}
	return tags
}
