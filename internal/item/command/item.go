package command

import (
	"github.com/api-plastik/db"
	"github.com/api-plastik/helper/querybuilder"
	qbModel "github.com/api-plastik/helper/querybuilder/model"
	"github.com/api-plastik/internal/item/model"
)

// CreateCategory ...
func (itemCommand *ItemCommand) CreateCategory(itemCat *model.ItemCategoryCreate) (int64, error) {
	// temp id for returned
	var id int64

	query := itemCommand.q.Create("item_categories", *itemCat)
	err := itemCommand.db.PgSQL.QueryRowx(query, itemCat.Name, itemCat.CreatedAt).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

// UpdateCategory ...
func (itemCommand *ItemCommand) UpdateCategory(id int, itemCat *model.ItemCategoryUpdate) error {
	// create condition
	where := &qbModel.Condition{
		Key:      "id",
		Value:    id,
		Operator: "=",
		NextCond: "",
	}

	// create query
	query := itemCommand.q.UpdateWhere("item_categories", *itemCat, []*qbModel.Condition{where})

	// exec query
	_, err := itemCommand.db.PgSQL.Exec(query, itemCat.Name)
	if err != nil {
		return err
	}
	return nil
}

// DeleteCategory ...
func (itemCommand *ItemCommand) DeleteCategory(id int) error {
	// create condition
	where := &qbModel.Condition{
		Key:      "id",
		Value:    id,
		Operator: "=",
		NextCond: "",
	}

	// create query
	query := itemCommand.q.Delete("item_categories", []*qbModel.Condition{where})

	// exec query
	_, err := itemCommand.db.PgSQL.Exec(query)
	if err != nil {
		return err
	}
	return nil
}

// CreateItem ...
func (itemCommand *ItemCommand) CreateItem(item *model.ItemCreate) error {
	query := itemCommand.q.Create("items", *item)
	_, err := itemCommand.db.PgSQL.Exec(query, item.ItemID, item.CategoryItemID, item.Name, item.CreatedAt)
	if err != nil {
		return err
	}

	return nil
}

// UpdateItem ...
func (itemCommand *ItemCommand) UpdateItem(id string, item *model.ItemUpdate) error {
	// create condition
	where := &qbModel.Condition{
		Key:      "id",
		Value:    id,
		Operator: "=",
		NextCond: "",
	}

	// create query
	query := itemCommand.q.UpdateWhere("items", *item, []*qbModel.Condition{where})

	// exec query
	_, err := itemCommand.db.PgSQL.Exec(query, item.Name, item.CategoryItemID)
	if err != nil {
		return err
	}
	return nil
}

// DeleteItem ...
func (itemCommand *ItemCommand) DeleteItem(id string) error {
	// create condition
	where := &qbModel.Condition{
		Key:      "id",
		Value:    id,
		Operator: "=",
		NextCond: "",
	}

	// create query
	query := itemCommand.q.Delete("items", []*qbModel.Condition{where})

	// exec query
	_, err := itemCommand.db.PgSQL.Exec(query)
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
