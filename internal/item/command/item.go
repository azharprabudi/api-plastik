package command

import (
	"fmt"

	"github.com/api-plastik/db"
	"github.com/api-plastik/helper/querybuilder"
	"github.com/api-plastik/internal/item/model"
)

// CreateCategory ...
func (itemCommand *ItemCommand) CreateCategory(itemCat *model.ItemCategoryCreate) error {
	q := qb.NewQueryBuilder()
	query := q.Create("itemCategory", *itemCat)
	fmt.Println(query)
	_, err := itemCommand.db.PgSQL.Exec(query)
	if err != nil {
		return err
	}
	return nil
}

// NewItemCommand ...
func NewItemCommand(db *db.DB) ItemCommandInterface {
	return &ItemCommand{
		db: db,
	}
}
