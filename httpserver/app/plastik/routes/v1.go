package routes

import (
	"github.com/azharprabudi/api-plastik/db"
	itemPresentation "github.com/azharprabudi/api-plastik/httpserver/app/plastik/presentations/item"
	itemCategoryPresentation "github.com/azharprabudi/api-plastik/httpserver/app/plastik/presentations/item-category"
	sellerPresentation "github.com/azharprabudi/api-plastik/httpserver/app/plastik/presentations/seller"
	supplierPresentation "github.com/azharprabudi/api-plastik/httpserver/app/plastik/presentations/supplier"

	"github.com/go-chi/chi"
)

// NewRoutesV1Plastik ...
func NewRoutesV1Plastik(newR *chi.Router, db *db.DB) {
	// initialize presentations
	item := itemPresentation.NewItemPresentation(db)
	itemCategory := itemCategoryPresentation.NewPresentationItemCategory(db)
	supplier := supplierPresentation.NewPresentationSupplier(db)
	seller := sellerPresentation.NewPresentationSeller(db)

	// route
	(*newR).Route("/v1", func(r chi.Router) {
		/* item */
		r.Get("/item", item.Find)
		r.Get("/item/{id}", item.FindByID)
		r.Post("/item", item.Create)
		r.Patch("/item/{id}", item.Update)
		r.Delete("/item/{id}", item.Delete)

		/* itemCategory */
		r.Get("/itemcategory", itemCategory.Find)
		r.Get("/itemcategory/{id}", itemCategory.FindByID)
		r.Post("/itemcategory", itemCategory.Create)
		r.Patch("/itemcategory/{id}", itemCategory.Update)
		r.Delete("/itemcategory/{id}", itemCategory.Delete)

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
	})
}
