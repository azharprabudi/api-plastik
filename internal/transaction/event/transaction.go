package event

import (
	"github.com/azharprabudi/api-plastik/internal/item/model"
	"github.com/azharprabudi/api-plastik/internal/item/service"
	"github.com/jmoiron/sqlx"
	uuid "github.com/satori/go.uuid"
)

// TriggerAfterCreateTransaction ...
func (te *TransactionEvent) TriggerAfterCreateTransaction(tx *sqlx.Tx, items []*model.ItemStockLogCreate, companyID uuid.UUID) error {
	var err error
	for _, item := range items {
		findItem, _err := te.itemService.GetItemByID(companyID, item.ItemID)
		if _err != nil {
			err = _err
			break
		}

		item.ItemName = &((*findItem).Item.Name)
		_err = te.itemService.CreateItemStockLog(tx, item)
		if _err != nil {
			err = _err
			break
		}
	}

	if err != nil {
		return err
	}
	return nil
}

// NewTransactionEvent ...
func NewTransactionEvent(itemService service.ItemServiceInterface) TransactionEventInterface {
	return &TransactionEvent{
		itemService: itemService,
	}
}
