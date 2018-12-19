package presentations

import "net/http"

// ItemUnitAbstract ...
type ItemUnitAbstract interface {
	Find(w http.ResponseWriter, r *http.Request)
}
