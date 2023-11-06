package conn

import (
	"database/sql"

	_ "github.com/lib/pq"
	"github.com/nedpals/supabase-go"

	"sytron-server/constants"
)

func loadPostgresDB() *sql.DB {
	_db, err := sql.Open("postgres", constants.POSTGRES_CONNECTION_STRING)
	if err != nil {
		print("Failed to connect to DB")
		print(err.Error())
	} else {
		print("Connected to Postgres DB")
	}
	return _db
}

func initSupabaseDB() *supabase.Client {
	supabaseUrl := constants.SUPABASE_PROJECT_URL
	supabaseKey := constants.SUPABASE_API_KEY
	return supabase.CreateClient(supabaseUrl, supabaseKey)
}

var (
	db   *sql.DB = loadPostgresDB()
	Supa         = initSupabaseDB()
)
