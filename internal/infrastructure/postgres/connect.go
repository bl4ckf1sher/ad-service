package postgres

import (
	//"database/sql"
	"github.com/bl4ckf1sher/ad-service/internal/config"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Connect struct {
	Db *sqlx.DB //ideally manage via a method (not a public field)
}

func NewConnect(cfg *config.DB) *Connect {
	db, err := sqlx.Connect("postgres", cfg.GetConnectionString())
	if err != nil {
		panic(err)
	}

	return &Connect{db}
}
