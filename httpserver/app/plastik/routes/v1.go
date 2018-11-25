package routes

import (
	"github.com/api-plastik/db"
	itemPresentation "github.com/api-plastik/httpserver/app/plastik/presentations/item"
	itemCategoryPresentation "github.com/api-plastik/httpserver/app/plastik/presentations/item-category"
	sellerPresentation "github.com/api-plastik/httpserver/app/plastik/presentations/seller"
	supplierPresentation "github.com/api-plastik/httpserver/app/plastik/presentations/supplier"

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
		r.Patch("/item", item.Update)
		r.Delete("/item", item.Delete)

		/* itemCategory */
		r.Get("/itemcategory", itemCategory.Find)
		r.Get("/itemcategory/{id}", itemCategory.FindByID)
		r.Post("/itemcategory", itemCategory.Create)
		r.Patch("/itemcategory", itemCategory.Update)
		r.Delete("/itemcategory", itemCategory.Delete)

		/* item */
		r.Get("/supplier", supplier.Find)
		r.Get("/supplier/{id}", supplier.FindByID)
		r.Post("/supplier", supplier.Create)
		r.Patch("/supplier", supplier.Update)
		r.Delete("/supplier", supplier.Delete)

		/* seller */
		r.Get("/seller", seller.Find)
		r.Get("/seller/{id}", seller.FindByID)
		r.Post("/seller", seller.Create)
		r.Patch("/seller", seller.Update)
		r.Delete("/seller", seller.Delete)
	})
}
