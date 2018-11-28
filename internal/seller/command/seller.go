package command

import (
	"github.com/api-plastik/db"
	"github.com/api-plastik/helper/querybuilder"
	qbModel "github.com/api-plastik/helper/querybuilder/model"
	"github.com/api-plastik/internal/seller/model"
	uuid "github.com/satori/go.uuid"
)

// Create ...
func (sellerCommand *SellerCommand) Create(seller *model.SellerCreate) error {
	query := sellerCommand.q.Create("sellers", *seller)
	_, err := sellerCommand.db.PgSQL.Exec(query, seller.Name, seller.CreatedAt)
	if err != nil {
		return err
	}

	return nil
}

// Update ...
func (sellerCommand *SellerCommand) Update(id uuid.UUID, seller *model.SellerUpdate) error {
	// create condition
	where := &qbModel.Condition{
		Key:      "id",
		Value:    id,
		Operator: "=",
		NextCond: "",
	}

	// create query
	query := sellerCommand.q.UpdateWhere("sellers", *seller, []*qbModel.Condition{where})

	// exec query
	_, err := sellerCommand.db.PgSQL.Exec(query, seller.Name)
	if err != nil {
		return err
	}
	return nil
}

// Delete ...
func (sellerCommand *SellerCommand) Delete(id uuid.UUID) error {
	// create condition
	where := &qbModel.Condition{
		Key:      "id",
		Value:    id,
		Operator: "=",
		NextCond: "",
	}

	// create query
	query := sellerCommand.q.Delete("sellers", []*qbModel.Condition{where})

	// exec query
	_, err := sellerCommand.db.PgSQL.Exec(query)
	if err != nil {
		return err
	}
	return nil
}

// NewSellerCommand ...
func NewSellerCommand(db *db.DB) SellerCommandInterface {
	q := qb.NewQueryBuilder()
	return &SellerCommand{
		q:  q,
		db: db,
	}
}
