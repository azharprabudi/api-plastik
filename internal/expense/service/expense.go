package service

import (
	"github.com/api-plastik/internal/expense/transform"
	uuid "github.com/satori/go.uuid"

	"github.com/api-plastik/db"
	"github.com/api-plastik/internal/expense/command"
	"github.com/api-plastik/internal/expense/dto"
	"github.com/api-plastik/internal/expense/query"
)

// GetExpenseType ...
func (expenseService *ExpenseService) GetExpenseType() ([]*dto.ExpenseTypeRes, error) {
	// add data to db
	expenseTypeModel, err := expenseService.query.GetExpenseType()
	if err != nil {
		return nil, err
	}

	// transform data from model
	expenseTypeDTO := expenseService.transform.TransformGetExpenseType(expenseTypeModel)
	return expenseTypeDTO, nil
}

// GetExpenseTypeByID ...
func (expenseService *ExpenseService) GetExpenseTypeByID(expenseTypeID int) *dto.ExpenseTypeRes {
	expenseTypeModel := expenseService.query.GetExpenseTypeByID(expenseTypeID)
	if expenseTypeModel == nil {
		return nil
	}

	// transform data from model
	expenseTypeDTO := expenseService.transform.TransformGetExpenseTypeByID(expenseTypeModel)
	return expenseTypeDTO
}

// CreateExpenseType ...
func (expenseService *ExpenseService) CreateExpenseType(itemCategory *dto.ExpenseTypeReq) (int64, error) {
	// transform dto to model
	itemCategoryCreate := expenseService.transform.TransformCreateCategory(itemCategory)

	// add data to db
	id, err := expenseService.command.CreateCategory(itemCategoryCreate)
	if err != nil {
		return 0, err
	}
	return id, nil
}

// UpdateExpenseType ...
func (expenseService *ExpenseService) UpdateExpenseType(categoryID int, itemCategory *dto.ExpenseTypeReq) error {
	// transform dto to model
	itemCategoryUpdate := expenseService.transform.TransformUpdateCategory(itemCategory)

	// update to db
	err := expenseService.command.UpdateCategory(categoryID, itemCategoryUpdate)
	if err != nil {
		return err
	}
	return nil
}

// DeleteExpenseType ...
func (expenseService *ExpenseService) DeleteExpenseType(categoryID int) error {
	// delete data from db
	err := expenseService.command.DeleteCategory(categoryID)
	if err != nil {
		return err
	}
	return nil
}

// GetItem ...
func (expenseService *ExpenseService) GetItem() ([]*dto.ItemRes, error) {
	// add data to db
	itemModel, err := expenseService.query.GetItem()
	if err != nil {
		return nil, err
	}

	// transform data from model
	itemDTO := expenseService.transform.TransformGetItem(itemModel)
	return itemDTO, nil
}

// GetItemByID ...
func (expenseService *ExpenseService) GetItemByID(itemID string) *dto.ItemRes {
	itemModel := expenseService.query.GetItemByID(itemID)
	if itemModel == nil {
		return nil
	}

	// transform data from model
	itemDTO := expenseService.transform.TransformGetItemByID(itemModel)
	return itemDTO
}

// CreateItem ...
func (expenseService *ExpenseService) CreateItem(item *dto.ItemReq) (uuid.UUID, error) {
	// transform dto to model
	itemCreate := expenseService.transform.TransformCreateItem(item)

	// add data to db
	err := expenseService.command.CreateItem(itemCreate)
	if err != nil {
		return uuid.Nil, err
	}
	return itemCreate.ItemID, nil
}

// UpdateItem ...
func (expenseService *ExpenseService) UpdateItem(itemID string, item *dto.ItemReq) error {
	// transform dto to model
	itemUpdate := expenseService.transform.TransformUpdateItem(item)

	// update to db
	err := expenseService.command.UpdateItem(itemID, itemUpdate)
	if err != nil {
		return err
	}
	return nil
}

// DeleteItem ...
func (expenseService *ExpenseService) DeleteItem(itemID string) error {
	// delete data from db
	err := expenseService.command.DeleteItem(itemID)
	if err != nil {
		return err
	}
	return nil
}

// NewExpenseService ...
func NewExpenseService(db *db.DB) ExpenseServiceInterface {
	return &ExpenseService{
		query:     query.NewExpenseQuery(db),
		command:   command.NewExpenseCommand(db),
		transform: transform.NewExpenseTransform(),
	}
}
