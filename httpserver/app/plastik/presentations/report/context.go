package presentations

import (
	itemService "github.com/azharprabudi/api-plastik/internal/item/service"
	"github.com/azharprabudi/api-plastik/internal/transaction/service"
)

// Report ...
type Report struct {
	itemService        itemService.ItemServiceInterface
	transactionService service.TransactionServiceInterface
}
