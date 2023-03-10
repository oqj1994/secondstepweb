package dbrepo

import (
	"database/sql"
	"github.com/vitaLemoTea/secondstepweb/internal/config"
	"github.com/vitaLemoTea/secondstepweb/internal/repository"
)

type postgresDBRepo struct {
	App *config.Config
	DB  *sql.DB
}

type testDBRepo struct {
	App *config.Config
	DB  *sql.DB
}

func NewPostgresRepo(a *config.Config, conn *sql.DB) repository.DatabaseRepo {
	p := postgresDBRepo{
		App: a,
		DB:  conn,
	}
	return &p
}

func NewTestRepo(a *config.Config) repository.DatabaseRepo {
	p := testDBRepo{
		App: a,
	}
	return &p
}
