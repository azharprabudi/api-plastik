package routes

import (
	"github.com/azharprabudi/api-plastik/db"
	expenseTypePresentation "github.com/azharprabudi/api-plastik/httpserver/app/plastik/presentations/expense-type"
	itemPresentation "github.com/azharprabudi/api-plastik/httpserver/app/plastik/presentations/item"

	itemCategoryPresentation "github.com/azharprabudi/api-plastik/httpserver/app/plastik/presentations/item-category"
	itemUnitPresentation "github.com/azharprabudi/api-plastik/httpserver/app/plastik/presentations/item-unit"
	sellerPresentation "github.com/azharprabudi/api-plastik/httpserver/app/plastik/presentations/seller"
	supplierPresentation "github.com/azharprabudi/api-plastik/httpserver/app/plastik/presentations/supplier"

	expensePresentation "github.com/azharprabudi/api-plastik/httpserver/app/plastik/presentations/expense"

	transactionPresentation "github.com/azharprabudi/api-plastik/httpserver/app/plastik/presentations/transaction"

	"github.com/go-chi/chi"
)

// NewRoutesV1Plastik ...
func NewRoutesV1Plastik(newR *chi.Router, db *db.DB) {
	// initialize presentations
	item := itemPresentation.NewItemPresentation(db)
	itemUnit := itemUnitPresentation.NewPresentationItemUnit(db)
	itemCategory := itemCategoryPresentation.NewPresentationItemCategory(db)
	supplier := supplierPresentation.NewSupplierPresentation(db)
	seller := sellerPresentation.NewSellerPresentation(db)
	expenseType := expenseTypePresentation.NewExpenseTypePresentation(db)
	expense := expensePresentation.NewExpensePresentation(db)
	transaction := transactionPresentation.NewPresentationTransaction(db)

	// route
	(*newR).Route("/v1", func(r chi.Router) {
		/* item */
		r.Get("/item", item.Find)
		r.Get("/item/{id}", item.FindByID)
		r.Post("/item", item.Create)
		r.Patch("/item/{id}", item.Update)
		r.Delete("/item/{id}", item.Delete)

		/* item unit */
		r.Get("/item-unit", itemUnit.Find)

		/* itemCategory */
		r.Get("/item-category", itemCategory.Find)
		r.Get("/item-category/{id}", itemCategory.FindByID)
		r.Post("/item-category", itemCategory.Create)
		r.Patch("/item-category/{id}", itemCategory.Update)
		r.Delete("/item-category/{id}", itemCategory.Delete)

		/* supplier */
		r.Get("/supplier", supplier.Find)
		r.Get("/supplier/{id}", supplier.FindByID)
		r.Post("/supplier", supplier.Create)
		r.Patch("/supplier/{id}", supplier.Update)
		r.Delete("/supplier/{id}", supplier.Delete)

		/* seller */
		r.Get("/seller", seller.Find)
		r.Get("/seller/{id}", seller.FindByID)
		r.Post("/seller", seller.Create)
		r.Patch("/seller/{id}", seller.Update)
		r.Delete("/seller/{id}", seller.Delete)

		/* expense type */
		r.Get("/expense-type", expenseType.Find)
		r.Get("/expense-type/{id}", expenseType.FindByID)
		r.Post("/expense-type", expenseType.Create)
		r.Patch("/expense-type/{id}", expenseType.Update)
		r.Delete("/expense-type/{id}", expenseType.Delete)

		/* expense */
		r.Get("/expense", expense.Find)
		r.Get("/expense/{id}", expense.FindByID)
		r.Post("/expense", expense.Create)

		/* transaction */
		r.Get("/transaction", transaction.Find)
		r.Get("/transaction/{id}", transaction.FindByID)
		r.Post("/transaction", transaction.Create)
	})
}
