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
	_, err := ic.db.PgSQL.Exec(query, itemCategory.ItemCategory.ItemCategoryID, itemCategory.ItemCategory.Name, itemCategory.ItemCategory.CreatedAt, itemCategory.ItemCategory.CompanyID)
	if err != nil {
		return err
	}

	return nil
}

// UpdateCategory ...
func (ic *ItemCommand) UpdateCategory(companyID uuid.UUID, id uuid.UUID, itemCategory *model.ItemCategoryUpdate) error {
	query := ic.q.UpdateWhere("item_categories", *itemCategory, []*qbmodel.Condition{&qbmodel.Condition{
		Key:      "id",
		Operator: "=",
		NextCond: "AND",
		Value:    id.String(),
	}, &qbmodel.Condition{
		Key:      "company_id",
		Operator: "=",
		NextCond: "",
		Value:    companyID.String(),
	}})
	_, err := ic.db.PgSQL.Exec(query, itemCategory.Name)
	if err != nil {
		return err
	}

	return nil
}

// DeleteCategory ...
func (ic *ItemCommand) DeleteCategory(companyID uuid.UUID, id uuid.UUID) error {
	status := struct {
		Active bool `db:"active"`
	}{
		Active: false,
	}

	query := ic.q.UpdateWhere("item_categories", status, []*qbmodel.Condition{&qbmodel.Condition{
		Key:      "id",
		Operator: "=",
		NextCond: "AND",
		Value:    id.String(),
	}, &qbmodel.Condition{
		Key:      "company_id",
		Operator: "=",
		NextCond: "",
		Value:    companyID.String(),
	}})
	_, err := ic.db.PgSQL.Exec(query, status.Active)
	if err != nil {
		return err
	}

	return nil
}

// CreateItem ...
func (ic *ItemCommand) CreateItem(item *model.ItemCreate) error {
	query := ic.q.Create("items", (*item).Item)
	_, err := ic.db.PgSQL.Exec(query, item.Item.ItemID, item.Item.ItemCategoryID, item.Item.Name, item.Item.UnitID, item.Item.CreatedAt, item.CompanyID)
	if err != nil {
		return err
	}

	return nil
}

// UpdateItem ...
func (ic *ItemCommand) UpdateItem(companyID uuid.UUID, id uuid.UUID, item *model.ItemUpdate) error {
	query := ic.q.UpdateWhere("items", *item, []*qbmodel.Condition{&qbmodel.Condition{
		Key:      "id",
		Operator: "=",
		NextCond: "AND",
		Value:    id.String(),
	}, &qbmodel.Condition{
		Key:      "company_id",
		Operator: "=",
		NextCond: "",
		Value:    companyID.String(),
	}})
	_, err := ic.db.PgSQL.Exec(query, item.Name, item.ItemCategoryID, item.UnitID)
	if err != nil {
		return err
	}

	return nil
}

// DeleteItem ...
func (ic *ItemCommand) DeleteItem(companyID uuid.UUID, id uuid.UUID) error {
	status := struct {
		Active bool `db:"active"`
	}{
		Active: false,
	}

	query := ic.q.UpdateWhere("items", status, []*qbmodel.Condition{&qbmodel.Condition{
		Key:      "id",
		Operator: "=",
		NextCond: "AND",
		Value:    id.String(),
	}, &qbmodel.Condition{
		Key:      "company_id",
		Operator: "=",
		NextCond: "",
		Value:    companyID.String(),
	}})
	_, err := ic.db.PgSQL.Exec(query, status.Active)
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
