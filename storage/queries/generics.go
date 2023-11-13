package queries

import "sytron-server/storage/conn"

var (
	db      = conn.Supa.DB
	pgxConn = conn.PgxConn
)
