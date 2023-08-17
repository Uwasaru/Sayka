package json

import (
	"github.com/Uwasaru/Sayka/domain/entity"
)

type TagJson struct {
	ID      int    `json:"id"`
	SaykaID string `json:"sayka_id"`
	Name    string `json:"name"`
}

func TagJsonToEntity(tag *entity.Tag) *TagJson {
	return &TagJson{
		ID:      tag.ID,
		SaykaID: tag.SaykaID,
		Name:    tag.Name,
	}
}

func TagEntityToJson(tagJson *TagJson) *entity.Tag {
	return &entity.Tag{
		ID:      tagJson.ID,
		SaykaID: tagJson.SaykaID,
		Name:    tagJson.Name,
	}
}
