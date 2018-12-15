package command

import (
	"github.com/api-plastik/db"
	"github.com/api-plastik/helper/querybuilder"
	qbModel "github.com/api-plastik/helper/querybuilder/model"
	"github.com/api-plastik/internal/seller/model"
	uuid "github.com/satori/go.uuid"
)

// Create ...
func (sc *SellerCommand) Create(s *model.SellerCreate) error {
	query := sc.q.Create("sellers", (*s).Seller)
	_, err := sc.db.PgSQL.Exec(query, s.Seller.SellerID, s.Seller.Name, s.Seller.Phone, s.Seller.Address, s.Seller.CreatedAt)
	if err != nil {
		return err
	}

	return nil
}

// Update ...
func (sc *SellerCommand) Update(id uuid.UUID, s *model.SellerUpdate) error {
	// create condition
	where := &qbModel.Condition{
		Key:      "id",
		Operator: "=",
		NextCond: "",
		Value:    id.String(),
	}

	// create query
	query := sc.q.UpdateWhere("sellers", *s, []*qbModel.Condition{where})

	// exec query
	_, err := sc.db.PgSQL.Exec(query, s.Name, s.Phone, s.Address)
	if err != nil {
		return err
	}
	return nil
}

// Delete ...
func (sc *SellerCommand) Delete(id uuid.UUID) error {
	// create condition
	where := &qbModel.Condition{
		Key:      "id",
		Operator: "=",
		NextCond: "",
		Value:    id.String(),
	}

	// create query
	query := sc.q.Delete("sellers", []*qbModel.Condition{where})

	// exec query
	_, err := sc.db.PgSQL.Exec(query)
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
