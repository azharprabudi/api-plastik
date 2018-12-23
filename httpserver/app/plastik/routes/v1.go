package routes

import (
	"github.com/azharprabudi/api-plastik/db"
	itemPresentation "github.com/azharprabudi/api-plastik/httpserver/app/plastik/presentations/item"

	itemCategoryPresentation "github.com/azharprabudi/api-plastik/httpserver/app/plastik/presentations/item-category"
	itemUnitPresentation "github.com/azharprabudi/api-plastik/httpserver/app/plastik/presentations/item-unit"
	sellerPresentation "github.com/azharprabudi/api-plastik/httpserver/app/plastik/presentations/seller"
	supplierPresentation "github.com/azharprabudi/api-plastik/httpserver/app/plastik/presentations/supplier"

	transactionPresentation "github.com/azharprabudi/api-plastik/httpserver/app/plastik/presentations/transaction"

	"github.com/go-chi/chi"
)

// NewRoutesV1Plastik ...
func NewRoutesV1Plastik(newR *chi.Router, db *db.DB) {
	// initialize presentations
	item := itemPresentation.NewItemPresentation(db)
	itemUnit := itemUnitPresentation.NewItemUnitPresentation(db)
	itemCategory := itemCategoryPresentation.NewPresentationItemCategory(db)
	supplier := supplierPresentation.NewSupplierPresentation(db)
	seller := sellerPresentation.NewSellerPresentation(db)
	transaction := transactionPresentation.NewTransactionPresentation(db)

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

		/* transaction */
		r.Get("/transaction", transaction.Find)
		r.Get("/transaction/{id}", transaction.FindByID)
		r.Post("/transaction/in", transaction.CreateTransactionIn)
		r.Post("/transaction/out", transaction.CreateTransactionOut)
		r.Post("/transaction/etc", transaction.CreateTransactionEtc)
		r.Post("/transaction/etc/type", transaction.CreateTransactionEtcType)
		r.Get("/transaction/etc/type", transaction.FindTransactionEtcTypes)
		r.Get("/transaction/etc/type/{id}", transaction.FindTransactionEtcTypeByID)
		r.Patch("/transaction/etc/type/{id}", transaction.UpdateTransactionEtcType)
		r.Delete("/transaction/etc/type/{id}", transaction.DeleteTransactionEtcType)
	})
}
