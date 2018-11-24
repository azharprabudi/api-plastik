package command

import (
	"github.com/api-plastik/db"
	"github.com/api-plastik/helper/querybuilder"
	"github.com/api-plastik/internal/item/model"
)

// CreateCategory ...
func (itemCommand *ItemCommand) CreateCategory(itemCat *model.ItemCategoryCreate) error {
	query := itemCommand.q.Create("itemCategory", *itemCat)
	_, err := itemCommand.db.PgSQL.Exec(query, itemCat.Name, itemCat.CreatedAt)
	if err != nil {
		return err
	}
	return nil
}

// NewItemCommand ...
func NewItemCommand(db *db.DB) ItemCommandInterface {
	q := qb.NewQueryBuilder()
	return &ItemCommand{
		q:  q,
		db: db,
	}
}
