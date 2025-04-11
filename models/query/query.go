package query

import (
	"main/models"
	"main/models/record_type"
)

type Query struct {
	RecordType record_type.RecordType
	Domain     models.Domain
}
