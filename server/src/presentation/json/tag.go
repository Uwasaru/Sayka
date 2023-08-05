package json

import (
	"github.com/Uwasaru/Sayka/domain/entity"
)

type TagJson struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func TagJsonToEntity(tag *entity.Tag) *TagJson {
	return &TagJson{
		ID:   tag.ID,
		Name: tag.Name,
	}
}

func TagEntityToJson(tagJson *TagJson) *entity.Tag {
	return &entity.Tag{
		ID:   tagJson.ID,
		Name: tagJson.Name,
	}
}
