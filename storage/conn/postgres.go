package conn

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/nedpals/supabase-go"

	"sytron-server/constants"
)

func initSupabaseDB() *supabase.Client {
	supabaseUrl := constants.SUPABASE_PROJECT_URL
	supabaseKey := constants.SUPABASE_API_KEY
	return supabase.CreateClient(supabaseUrl, supabaseKey)
}

func initPGX() *pgxpool.Pool {
	conn, _ := pgxpool.New(context.Background(), constants.POSTGRES_CONNECTION_STRING)
	log.Println("Connected to Postgress")
	return conn
}

func Close() {
	PgxConn.Close()
}

var PgxConn *pgxpool.Pool = initPGX()

// Supa              = initSupabaseDB()
