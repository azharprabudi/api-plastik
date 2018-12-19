package presentations

import (
	"net/http"

	"github.com/azharprabudi/api-plastik/db"
	newError "github.com/azharprabudi/api-plastik/httpserver/error"
	"github.com/azharprabudi/api-plastik/httpserver/response"
	"github.com/azharprabudi/api-plastik/internal/item/service"
)

// Find ...
func (iu *ItemUnit) Find(w http.ResponseWriter, r *http.Request) {
	results, err := iu.service.GetItemUnit()
	if err != nil {
		response.Send(w, http.StatusInternalServerError, nil, newError.NewErrorReponse(newError.InternalServerError, err.Error(), "", nil))
		return
	}
	response.Send(w, http.StatusOK, nil, results)
	return
}

// NewPresentationItemUnit ...
func NewPresentationItemUnit(db *db.DB) ItemUnitAbstract {
	return &ItemUnit{
		service: service.NewItemService(db),
	}
}
