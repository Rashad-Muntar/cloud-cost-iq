package billing

import (
	"github.com/cloud-cost-iq/internals/db"
)

type Repository struct {
	db *db.Database
}

type Service struct {
	repo *Repository
}




