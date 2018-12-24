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

	/* item unit */
	(*newR).Get("/v1/item-unit", itemUnit.Find)

	// route
	(*newR).Route("/v1/company", func(r chi.Router) {
		/* item */
		r.Get("/{companyId}/item", item.Find)
		r.Get("/{companyId}/item/{id}", item.FindByID)
		r.Post("/{companyId}/item", item.Create)
		r.Patch("/{companyId}/item/{id}", item.Update)
		r.Delete("/{companyId}/item/{id}", item.Delete)

		/* itemCategory */
		r.Get("/{companyId}/item-category", itemCategory.Find)
		r.Get("/{companyId}/item-category/{id}", itemCategory.FindByID)
		r.Post("/{companyId}/item-category", itemCategory.Create)
		r.Patch("/{companyId}/item-category/{id}", itemCategory.Update)
		r.Delete("/{companyId}/item-category/{id}", itemCategory.Delete)

		/* supplier */
		r.Get("/{companyId}/supplier", supplier.Find)
		r.Get("/{companyId}/supplier/{id}", supplier.FindByID)
		r.Post("/{companyId}/supplier", supplier.Create)
		r.Patch("/{companyId}/supplier/{id}", supplier.Update)
		r.Delete("/{companyId}/supplier/{id}", supplier.Delete)

		/* seller */
		r.Get("/{companyId}/seller", seller.Find)
		r.Get("/{companyId}/seller/{id}", seller.FindByID)
		r.Post("/{companyId}/seller", seller.Create)
		r.Patch("/{companyId}/seller/{id}", seller.Update)
		r.Delete("/{companyId}/seller/{id}", seller.Delete)

		/* transaction */
		r.Get("/{companyId}/transaction", transaction.Find)
		r.Get("/{companyId}/transaction/{id}", transaction.FindByID)
		r.Post("/{companyId}/transaction/in", transaction.CreateTransactionIn)
		r.Post("/{companyId}/transaction/out", transaction.CreateTransactionOut)
		r.Post("/{companyId}/transaction/etc", transaction.CreateTransactionEtc)
		r.Post("/{companyId}/transaction/etc/type", transaction.CreateTransactionEtcType)
		r.Get("/{companyId}/transaction/etc/type", transaction.FindTransactionEtcTypes)
		r.Get("/{companyId}/transaction/etc/type/{id}", transaction.FindTransactionEtcTypeByID)
		r.Patch("/{companyId}/transaction/etc/type/{id}", transaction.UpdateTransactionEtcType)
		r.Delete("/{companyId}/transaction/etc/type/{id}", transaction.DeleteTransactionEtcType)
	})
}
