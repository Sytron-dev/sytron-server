package conn

import (
	"context"
	"log"
	"sytron-server/constants"

	"github.com/jackc/pgx/v5"
	"github.com/nedpals/supabase-go"
)

func initSupabaseDB() *supabase.Client {
	supabaseUrl := constants.SUPABASE_PROJECT_URL
	supabaseKey := constants.SUPABASE_API_KEY
	return supabase.CreateClient(supabaseUrl, supabaseKey)
}

func initPGX() *pgx.Conn {
	conn, _ := pgx.Connect(context.Background(), constants.POSTGRES_CONNECTION_STRING)
	log.Println("Connected to Postgress")
	return conn
}

func Close() {
	PgxConn.Close(context.Background())
}

var PgxConn *pgx.Conn = initPGX()

// Supa              = initSupabaseDB()
