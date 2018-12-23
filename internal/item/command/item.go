package command

import (
	"github.com/azharprabudi/api-plastik/db"
	qb "github.com/azharprabudi/api-plastik/helper/querybuilder"
	qbmodel "github.com/azharprabudi/api-plastik/helper/querybuilder/model"
	"github.com/azharprabudi/api-plastik/internal/item/model"
	"github.com/jmoiron/sqlx"
	uuid "github.com/satori/go.uuid"
)

// CreateCategory ...
func (ic *ItemCommand) CreateCategory(itemCategory *model.ItemCategoryCreate) error {
	query := ic.q.Create("item_categories", (*itemCategory).ItemCategory)
	_, err := ic.db.PgSQL.Exec(query, itemCategory.ItemCategory.ItemCategoryID, itemCategory.ItemCategory.Name, itemCategory.ItemCategory.CreatedAt)
	if err != nil {
		return err
	}

	return nil
}

// UpdateCategory ...
func (ic *ItemCommand) UpdateCategory(id uuid.UUID, itemCategory *model.ItemCategoryUpdate) error {
	query := ic.q.UpdateWhere("item_categories", *itemCategory, []*qbmodel.Condition{&qbmodel.Condition{
		Key:      "id",
		Operator: "=",
		NextCond: "",
		Value:    id.String(),
	}})
	_, err := ic.db.PgSQL.Exec(query, itemCategory.Name)
	if err != nil {
		return err
	}

	return nil
}

// DeleteCategory ...
func (ic *ItemCommand) DeleteCategory(id uuid.UUID) error {
	query := ic.q.Delete("item_categories", []*qbmodel.Condition{&qbmodel.Condition{
		Key:      "id",
		Operator: "=",
		NextCond: "",
		Value:    id.String(),
	}})
	_, err := ic.db.PgSQL.Exec(query)
	if err != nil {
		return err
	}

	return nil
}

// CreateItem ...
func (ic *ItemCommand) CreateItem(item *model.ItemCreate) error {
	query := ic.q.Create("items", (*item).Item)
	_, err := ic.db.PgSQL.Exec(query, item.Item.ItemID, item.Item.ItemCategoryID, item.Item.Name, item.Item.UnitID, item.CreatedAt)
	if err != nil {
		return err
	}

	return nil
}

// UpdateItem ...
func (ic *ItemCommand) UpdateItem(id uuid.UUID, item *model.ItemUpdate) error {
	query := ic.q.UpdateWhere("items", *item, []*qbmodel.Condition{&qbmodel.Condition{
		Key:      "id",
		Operator: "=",
		NextCond: "",
		Value:    id.String(),
	}})
	_, err := ic.db.PgSQL.Exec(query, item.Name, item.ItemCategoryID)
	if err != nil {
		return err
	}

	return nil
}

// DeleteItem ...
func (ic *ItemCommand) DeleteItem(id uuid.UUID) error {
	query := ic.q.Delete("items", []*qbmodel.Condition{&qbmodel.Condition{
		Key:      "id",
		Operator: "=",
		NextCond: "",
		Value:    id.String(),
	}})
	_, err := ic.db.PgSQL.Exec(query)
	if err != nil {
		return err
	}

	return nil
}

// CreateItemStockLog ...
func (ic *ItemCommand) CreateItemStockLog(tx *sqlx.Tx, item *model.ItemStockLogCreate) error {
	query := ic.q.Create("item_stock_logs", *item)
	_, err := tx.Exec(query, item.ID, *(*item).ItemName, item.ItemID, item.TransactionID, item.Qty, item.CreatedAt)
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
