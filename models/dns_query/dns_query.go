package dns_query

import (
	"main/models"
	"main/models/record_type"
)

type DNSQuery struct {
	RecordType record_type.RecordType
	Domain     models.Domain
}
