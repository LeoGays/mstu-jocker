package mapper

import (
	"jocer/internal/model"
	"jocer/internal/server/generated"
)

func CreateJockListResponse(jocks []*model.Jock) generated.JockList {
	items := make([]generated.Jock, 0, len(jocks))
	for _, jock := range jocks {
		items = append(items, generated.Jock{
			Content: jock.Content,
			Name:    jock.Name,
			Id:      jock.ID.String(),
		})
	}

	return generated.JockList{Items: items}
}

func CreateJockResponse(jock *model.Jock) generated.Jock {
	return generated.Jock{
		Content: jock.Content,
		Id:      jock.ID.String(),
		Name:    jock.Name,
	}
}

func CreateJock(jock generated.JockRequestBody) *model.Jock {
	return &model.Jock{
		Name:    jock.Name,
		Content: jock.Content,
	}
}
