package dto

import (
	"github.com/Uwasaru/Sayka/domain/entity"
)

type TagDto struct {
	ID   string `db:"id"`
	Name string `db:"name"`
}

type TagsDto []*TagDto

func TagDtoToEntity(dto *TagDto) *entity.Tag {
	return &entity.Tag{
		ID:   dto.ID,
		Name: dto.Name,
	}
}

func TagEntityToDto(t *entity.Tag) TagDto {
	return TagDto{
		ID:   t.ID,
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