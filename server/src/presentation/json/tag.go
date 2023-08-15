package json

import (
	"github.com/Uwasaru/Sayka/domain/entity"
)

type TagJson struct {
	ID   int `json:"id"`
	PostID string `json:"post_id"`
	Name string `json:"name"`
}

func TagJsonToEntity(tag *entity.Tag) *TagJson {
	return &TagJson{
		ID:   tag.ID,
		PostID: tag.PostID,
		Name: tag.Name,
	}
}

func TagEntityToJson(tagJson *TagJson) *entity.Tag {
	return &entity.Tag{
		ID:   tagJson.ID,
		PostID: tagJson.PostID,
		Name: tagJson.Name,
	}
}
