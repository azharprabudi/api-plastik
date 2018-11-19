package route

import (
	itemPresentation "github.com/api-plastik/presentations/item"
	itemCategoryPresentation "github.com/api-plastik/presentations/item-category"
	sellerPresentation "github.com/api-plastik/presentations/seller"
	supplierPresentation "github.com/api-plastik/presentations/supplier"

	"github.com/go-chi/chi"
)

// NewRoutesV1 ...
func NewRoutesV1(route *Route) {
	// initialize presentations
	item := itemPresentation.NewItemPresentation(route.db)
	itemCategory := itemCategoryPresentation.NewPresentationItemCategory(route.db)
	supplier := supplierPresentation.NewPresentationSupplier(route.db)
	seller := sellerPresentation.NewPresentationSeller(route.db)

	// route
	route.r.Route("/v1", func(r chi.Router) {
		/* item */
		r.Get("/item", item.Find)
		r.Get("/item/:id", item.FindByID)
		r.Post("/item", item.Create)
		r.Patch("/item", item.Update)
		r.Delete("/item", item.Delete)

		/* itemCategory */
		r.Get("/item-category", itemCategory.Find)
		r.Get("/item-category/:id", itemCategory.FindByID)
		r.Post("/item-category", itemCategory.Create)
		r.Patch("/item-category", itemCategory.Update)
		r.Delete("/item-category", itemCategory.Delete)

		/* item */
		r.Get("/supplier", supplier.Find)
		r.Get("/supplier/:id", supplier.FindByID)
		r.Post("/supplier", supplier.Create)
		r.Patch("/supplier", supplier.Update)
		r.Delete("/supplier", supplier.Delete)

		/* seller */
		r.Get("/seller", seller.Find)
		r.Get("/seller/:id", seller.FindByID)
		r.Post("/seller", seller.Create)
		r.Patch("/seller", seller.Update)
		r.Delete("/seller", seller.Delete)
	})
}
