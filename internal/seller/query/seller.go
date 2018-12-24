package query

import (
	"github.com/azharprabudi/api-plastik/db"
	qb "github.com/azharprabudi/api-plastik/helper/querybuilder"
	qbmodel "github.com/azharprabudi/api-plastik/helper/querybuilder/model"
	"github.com/azharprabudi/api-plastik/internal/seller/model"
	uuid "github.com/satori/go.uuid"
)

// GetSellers ...
func (iq *SellerQuery) GetSellers(companyID uuid.UUID) ([]*model.SellerRead, error) {
	results := []*model.SellerRead{}
	query := iq.qb.Query("sellers", 0, 0, []*qbmodel.Condition{
		&qbmodel.Condition{
			Key:      "company_id",
			NextCond: "",
			Operator: "=",
			Value:    companyID.String(),
		},
	}, []*qbmodel.Order{
		&qbmodel.Order{
			Key:   "created_at",
			Value: "desc",
		},
	})

	rows, err := iq.db.PgSQL.Queryx(query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		tmp := new(model.SellerRead)
		rows.StructScan(tmp)
		results = append(results, tmp)
	}
	return results, nil
}

// GetSellerByID ...
func (iq *SellerQuery) GetSellerByID(companyID uuid.UUID, id uuid.UUID) (*model.SellerRead, error) {
	result := new(model.SellerRead)
	query := iq.qb.QueryWhere("sellers", []*qbmodel.Condition{&qbmodel.Condition{
		Key:      "id",
		NextCond: "AND",
		Operator: "=",
		Value:    id.String(),
	}, &qbmodel.Condition{
		Key:      "company_id",
		NextCond: "",
		Operator: "=",
		Value:    companyID.String(),
	}}, nil)

	err := iq.db.PgSQL.QueryRowx(query).StructScan(result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// NewSellerQuery ...
func NewSellerQuery(db *db.DB) SellerQueryInterface {
	q := qb.NewQueryBuilder()
	return &SellerQuery{
		qb: q,
		db: db,
	}
}
