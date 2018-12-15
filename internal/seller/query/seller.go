package query

import (
	"github.com/api-plastik/db"
	qb "github.com/api-plastik/helper/querybuilder"
	"github.com/api-plastik/helper/querybuilder/model"
	"github.com/api-plastik/internal/seller/model"
	uuid "github.com/satori/go.uuid"
)

// Get ...
func (iq *SellerQuery) Get() ([]*model.SellerRead, error) {
	// init variable
	var results = []*model.SellerRead{}

	// orders
	orders := []*qbmodel.Order{
		&qbmodel.Order{
			Key:   "created_at",
			Value: "desc",
		},
	}

	// get query
	query := iq.qb.Query("sellers", 0, 0, orders)
	rows, err := iq.db.PgSQL.Queryx(query)
	if err != nil {
		return nil, err
	}

	// get struct
	for rows.Next() {
		tmp := new(model.SellerRead)
		rows.StructScan(tmp)
		results = append(results, tmp)
	}

	return results, nil
}

// GetByID ...
func (iq *SellerQuery) GetByID(id uuid.UUID) *model.SellerRead {
	// init variable
	result := new(model.SellerRead)

	// create conditional
	where := &qbmodel.Condition{
		Key:      "id",
		NextCond: "",
		Operator: "=",
		Value:    id.String(),
	}

	// get query and execute
	query := iq.qb.QueryWhere("sellers", []*qbmodel.Condition{where}, nil)
	err := iq.db.PgSQL.QueryRowx(query).StructScan(result)
	if err != nil {
		return nil
	}
	return result
}

// NewSellerQuery ...
func NewSellerQuery(db *db.DB) SellerQueryInterface {
	q := qb.NewQueryBuilder()
	return &SellerQuery{
		qb: q,
		db: db,
	}
}
