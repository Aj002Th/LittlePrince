package db

import (
	"github.com/Aj002Th/LittlePrince/data"
	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB

func init() {
	db = data.GetMySQLConn()
	UserR.init()
}
