package db

import (
	"database/sql"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/driver/pgdriver"
	_ "github.com/uptrace/bun/driver/pgdriver"
)


 var DB *bun.DB

 func InitDB(dsn string){
	sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))
	DB = bun.NewDB(sqldb, pgdialect.New())
 }