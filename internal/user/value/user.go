package value

import (
	uuid "github.com/satori/go.uuid"
)

var (
	COMPANY_ID, _ = uuid.FromString("8f17a762-ddf8-472d-8a4a-c0e79924e269")
	USER_ID, _    = uuid.FromString("0f78335b-f693-42af-8a67-19503daecd4e")
	GROUP_ID, _   = uuid.FromString("bbb814b8-8ad8-4301-85b5-050ed74ca809")
	RETAIL        = "RETAIL"
	NON_RETAIL    = "NON_RETAIL"
)
