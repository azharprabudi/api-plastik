package command

import (
	"github.com/api-plastik/db"
	"github.com/api-plastik/helper/querybuilder"
	qbmodel "github.com/api-plastik/helper/querybuilder/model"
	"github.com/api-plastik/internal/item/model"
	"github.com/satori/go.uuid"
)

// CreateCategory ...
func (ic *ItemCommand) CreateCategory(category *model.ItemCategoryCreate) error {
	query := ic.q.Create("item_categories", (*category).ItemCategory)
	_, err := ic.db.PgSQL.Exec(query, category.ItemCategoryID, category.ItemCategory.Name, category.ItemCategory.CreatedAt)

	// return the id has been created
	if err != nil {
		return err
	}

	return nil
}

// UpdateCategory ...
func (ic *ItemCommand) UpdateCategory(id uuid.UUID, category *model.ItemCategoryUpdate) error {
	// create condition
	where := &qbmodel.Condition{
		Key:      "id",
		Operator: "=",
		NextCond: "",
		Value:    id.String(),
	}

	// create query
	query := ic.q.UpdateWhere("item_categories", *category, []*qbmodel.Condition{where})

	// exec query
	_, err := ic.db.PgSQL.Exec(query, category.Name)
	if err != nil {
		return err
	}
	return nil
}

// DeleteCategory ...
func (ic *ItemCommand) DeleteCategory(id uuid.UUID) error {
	// create condition
	where := &qbmodel.Condition{
		Key:      "id",
		Operator: "=",
		NextCond: "",
		Value:    id.String(),
	}

	// create query
	query := ic.q.Delete("item_categories", []*qbmodel.Condition{where})

	// exec query
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
	// create condition
	where := &qbmodel.Condition{
		Key:      "id",
		Operator: "=",
		NextCond: "",
		Value:    id.String(),
	}

	// create query
	query := ic.q.UpdateWhere("items", *item, []*qbmodel.Condition{where})

	// exec query
	_, err := ic.db.PgSQL.Exec(query, item.Name, item.ItemCategoryID)
	if err != nil {
		return err
	}
	return nil
}

// DeleteItem ...
func (ic *ItemCommand) DeleteItem(id uuid.UUID) error {
	// create condition
	where := &qbmodel.Condition{
		Key:      "id",
		Operator: "=",
		NextCond: "",
		Value:    id.String(),
	}

	// create query
	query := ic.q.Delete("items", []*qbmodel.Condition{where})

	// exec query
	_, err := ic.db.PgSQL.Exec(query)
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
