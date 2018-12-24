package command

import (
	"github.com/azharprabudi/api-plastik/db"
	qb "github.com/azharprabudi/api-plastik/helper/querybuilder"
	qbModel "github.com/azharprabudi/api-plastik/helper/querybuilder/model"
	"github.com/azharprabudi/api-plastik/internal/seller/model"
	uuid "github.com/satori/go.uuid"
)

// CreateSeller ...
func (sc *SellerCommand) CreateSeller(s *model.SellerCreate) error {
	query := sc.q.Create("sellers", (*s).Seller)
	_, err := sc.db.PgSQL.Exec(query, s.Seller.SellerID, s.Seller.Name, s.Seller.Phone, s.Seller.Address, s.Seller.CreatedAt, s.Seller.CompanyID)
	if err != nil {
		return err
	}

	return nil
}

// UpdateSeller ...
func (sc *SellerCommand) UpdateSeller(companyID uuid.UUID, id uuid.UUID, s *model.SellerUpdate) error {
	query := sc.q.UpdateWhere("sellers", *s, []*qbModel.Condition{&qbModel.Condition{
		Key:      "id",
		Operator: "=",
		NextCond: "AND",
		Value:    id.String(),
	}, &qbModel.Condition{
		Key:      "company_id",
		Operator: "=",
		NextCond: "",
		Value:    companyID.String(),
	}})
	_, err := sc.db.PgSQL.Exec(query, s.Name, s.Phone, s.Address)
	if err != nil {
		return err
	}

	return nil
}

// DeleteSeller ...
func (sc *SellerCommand) DeleteSeller(companyID uuid.UUID, id uuid.UUID) error {
	query := sc.q.Delete("sellers", []*qbModel.Condition{&qbModel.Condition{
		Key:      "id",
		Operator: "=",
		NextCond: "AND",
		Value:    id.String(),
	}, &qbModel.Condition{
		Key:      "company_id",
		Operator: "=",
		NextCond: "",
		Value:    companyID.String(),
	}})
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
