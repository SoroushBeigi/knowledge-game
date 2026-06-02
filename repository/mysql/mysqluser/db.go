package mysqluser

import "github.com/SoroushBeigi/knowledge-game/repository/mysql"

type db struct {
	conn *mysql.MySQLDB
}

func New(conn *mysql.MySQLDB) *db {
	return &db{
		conn: conn,
	}
}
