package event

import "github.com/azharprabudi/api-plastik/internal/item/service"

// TransactionEvent ...
type TransactionEvent struct {
	itemService service.ItemServiceInterface
}
