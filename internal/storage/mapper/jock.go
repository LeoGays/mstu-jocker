package mapper

import (
	"jocer/internal/model"
	"jocer/internal/storage/ent"
)

func CreateJockList(jocks []*ent.Jock) []*model.Jock {
	items := make([]*model.Jock, 0, len(jocks))
	for _, jock := range jocks {
		items = append(items, &model.Jock{
			ID:      jock.ID,
			Name:    jock.Name,
			Content: jock.Content,
		})
	}

	return items
}
