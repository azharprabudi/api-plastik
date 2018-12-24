package query

import (
	"fmt"
	"time"

	"github.com/azharprabudi/api-plastik/db"
	"github.com/azharprabudi/api-plastik/internal/transaction/model"
	uuid "github.com/satori/go.uuid"

	qb "github.com/azharprabudi/api-plastik/helper/querybuilder"
	qbModel "github.com/azharprabudi/api-plastik/helper/querybuilder/model"
)

// GetTransactions ...
func (tq *TransactionQuery) GetTransactions(companyID uuid.UUID, limit int, start int, startAt string, endAt string, orderBy string) ([]*model.TransactionRead, error) {
	var queryLimit string
	queryFilter := fmt.Sprintf("where company_id='%s'", companyID.String())
	queryOrder := "order by transactions.created_at desc"
	results := []*model.TransactionRead{}

	if limit > 0 && start > 0 {
		queryLimit = fmt.Sprintf("limit %d offset %d", limit, start)
	}

	if startAt != "" && endAt != "" {
		queryFilter = fmt.Sprintf("where transactions.created_at::timestamp between '%s'::timestamp AND '%s'::timestamp AND company_id='%s'", startAt, endAt, companyID.String())
	}

	if orderBy != "" {
		queryOrder = fmt.Sprintf("order by transactions.%s", orderBy)
	}

	query := fmt.Sprintf(`
	select * , 
		case 
			when transactions.type = 'TRANSACTION_IN' 
				then 'Transaksi Masuk' 
			when transactions.type = 'TRANSACTION_OUT' 
				then 'Transaksi Keluar' 
			else 'Transaksi Lainnya' 
		end as type_name 
	from transactions %s %s %s`, queryFilter, queryOrder, queryLimit,
	)
	rows, err := tq.db.PgSQL.Queryx(query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		tmp := new(model.TransactionRead)
		rows.StructScan(tmp)
		results = append(results, tmp)
	}

	return results, nil
}

// GetTransactionByID ...
func (tq *TransactionQuery) GetTransactionByID(companyID uuid.UUID, id uuid.UUID) (*model.TransactionReadDetail, error) {
	query := fmt.Sprintf(`
	select 
		transactions.*,
		case 
			when transactions.type = 'TRANSACTION_IN' 
				then 'Transaksi Masuk' 
			when transactions.type = 'TRANSACTION_OUT' 
				then 'Transaksi Keluar' 
			else 'Transaksi Lainnya' 
		end as type_name ,
		transactions_in.id as transaction_in_id,
		transactions_in.supplier_id as supplier_id,
		suppliers."name" as supplier_name,
		transactions_out.id as transaction_out_id,
		transactions_out.seller_id,
		sellers."name" as seller_name,
		transactions_etc.id as transaction_etc_id,
		transaction_etc_types."name" as transaction_etc_type_name,
		transaction_details.id as transaction_detail_id,
		transaction_details.item_id as item_id,
		transaction_details.item_name as item_name,
		transaction_details.qty as qty,
		transaction_details.amount as amount,
		transaction_images.id as transaction_image_id,
		transaction_images.image as image
	from transactions
	left join transactions_in on transactions.id = transactions_in.transaction_id
	left join transactions_out on transactions.id = transactions_out.transaction_id
	left join transactions_etc on transactions.id = transactions_etc.transaction_id
	left join transaction_etc_types on transactions_etc.transaction_etc_type = transaction_etc_types.id
	left join suppliers on transactions_in.supplier_id = suppliers.id
	left join sellers on transactions_out.seller_id = sellers.id
	left join transaction_details on transactions.id = transaction_details.transaction_id
	left join transaction_images on transactions.id = transaction_images.transaction_id
	where transactions.id = '%s' and transactions.company_id = '%s' order by transaction_details.id asc, transaction_images.id asc
	`, id.String(), companyID.String())

	var result *model.TransactionReadDetail
	var images []*model.TransactionImageRead
	var details []*model.TransactionDetailRead
	var tmpImageID, tmpDetailID uuid.UUID
	rows, err := tq.db.PgSQL.Queryx(query)
	if err != nil {
		return nil, err
	}

	i := 0
	for rows.Next() {
		var id uuid.UUID
		var note string
		var userID uuid.UUID
		var companyID uuid.UUID
		var transactionType string
		var typeName string
		var amount float64
		var createdAt time.Time
		var transactionInID *uuid.UUID
		var supplierID *uuid.UUID
		var supplierName *string
		var transactionOutID *uuid.UUID
		var sellerID *uuid.UUID
		var sellerName *string
		var transactionEtcID *uuid.UUID
		var transactionEtcTypeName *string
		var transactionDetailID *uuid.UUID
		var itemID *uuid.UUID
		var itemName *string
		var qty *int
		var amountDetail *float64
		var transactionImageID *uuid.UUID
		var image *string

		_err := rows.Scan(&id, &note, &userID, &companyID, &transactionType, &amount, &createdAt, &typeName, &transactionInID, &supplierID, &supplierName, &transactionOutID, &sellerID, &sellerName, &transactionEtcID, &transactionEtcTypeName, &transactionDetailID, &itemID, &itemName, &qty, &amountDetail, &transactionImageID, &image)
		if _err != nil {
			err = _err
			break
		}

		if i == 0 {
			result = &model.TransactionReadDetail{
				TransactionRead: model.TransactionRead{
					Transaction: model.Transaction{
						ID:        id,
						Note:      note,
						UserID:    userID,
						Amount:    amount,
						CreatedAt: createdAt,
						Type:      transactionType,
						CompanyID: companyID,
					},
					TypeName: typeName,
				},
				TransactionOutID:       transactionOutID,
				SellerID:               sellerID,
				SellerName:             sellerName,
				TransactionInID:        transactionInID,
				SupplierID:             supplierID,
				SupplierName:           supplierName,
				TransactionEtcID:       transactionEtcID,
				TransactionEtcTypeName: transactionEtcTypeName,
				Images:                 []*model.TransactionImageRead{&model.TransactionImageRead{}},
				Details:                []*model.TransactionDetailRead{&model.TransactionDetailRead{}},
			}

			if tmpImageID != *transactionImageID && transactionImageID != nil {
				images = append(images, &model.TransactionImageRead{
					ID:    transactionImageID,
					Image: image,
				})
			}

			if tmpDetailID != *transactionDetailID && transactionDetailID != nil {
				details = append(details, &model.TransactionDetailRead{
					ID:       transactionDetailID,
					ItemID:   itemID,
					ItemName: itemName,
					Qty:      qty,
					Amount:   amountDetail,
				})
			}

			tmpImageID = *transactionImageID
			tmpDetailID = *transactionDetailID
		}
	}

	if len(images) > 0 {
		result.Images = images
	}

	if len(details) > 0 {
		result.Details = details
	}

	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetTransactionEtcTypes ...
func (tq *TransactionQuery) GetTransactionEtcTypes(companyID uuid.UUID) ([]*model.TransactionEtcTypeRead, error) {
	results := []*model.TransactionEtcTypeRead{}
	query := tq.qb.Query("transaction_etc_types", 0, 0, []*qbModel.Condition{
		&qbModel.Condition{
			Key:      "company_id",
			NextCond: "",
			Operator: "=",
			Value:    companyID.String(),
		},
	}, []*qbModel.Order{
		&qbModel.Order{
			Key:   "created_at",
			Value: "desc",
		},
	})
	rows, err := tq.db.PgSQL.Queryx(query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		tmp := new(model.TransactionEtcTypeRead)
		rows.StructScan(tmp)
		results = append(results, tmp)
	}

	return results, nil
}

// GetTransactionEtcTypeByID ...
func (tq *TransactionQuery) GetTransactionEtcTypeByID(companyID uuid.UUID, id uuid.UUID) (*model.TransactionEtcTypeRead, error) {
	result := new(model.TransactionEtcTypeRead)
	query := tq.qb.QueryWhere("transaction_etc_types", []*qbModel.Condition{&qbModel.Condition{
		Key:      "id",
		NextCond: "AND",
		Operator: "=",
		Value:    id.String(),
	}, &qbModel.Condition{
		Key:      "company_id",
		NextCond: "",
		Operator: "=",
		Value:    companyID.String(),
	}}, nil)
	err := tq.db.PgSQL.QueryRowx(query).StructScan(result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// NewTransactionQuery ...
func NewTransactionQuery(db *db.DB) TransactionQueryInterface {
	q := qb.NewQueryBuilder()
	return &TransactionQuery{
		qb: q,
		db: db,
	}

}
