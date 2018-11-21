package presentations

import (
	"log"
	"net/http"

	"github.com/api-plastik/db"
	"github.com/api-plastik/dto"
	helpers "github.com/api-plastik/helpers/json"
	"github.com/api-plastik/internal/item/service"
	"github.com/api-plastik/presentations"
)

// Find ...
func (item *ItemCategory) Find(w http.ResponseWriter, r *http.Request) {

}

// FindByID ...
func (item *ItemCategory) FindByID(w http.ResponseWriter, r *http.Request) {

}

// Create ...
func (item *ItemCategory) Create(w http.ResponseWriter, r *http.Request) {
	itemCatIncReq := new(dto.ItemCategoryIncReq)
	err := helpers.JSONDecode(r.Body, itemCatIncReq)
	if err != nil {
		log.Println("error")
	}
}

// Update ...
func (item *ItemCategory) Update(w http.ResponseWriter, r *http.Request) {

}

// Delete ...
func (item *ItemCategory) Delete(w http.ResponseWriter, r *http.Request) {

}

// NewPresentationItemCategory ...
func NewPresentationItemCategory(db *db.DB) presentations.BaseAbstract {
	return &ItemCategory{
		itemService: service.NewItemService(db),
	}
}
