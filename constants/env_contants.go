package constants

import (
	"os"

	"github.com/joho/godotenv"
)

var (
	ENVIRONMENT string = ""
	PORT        string = ""

	MONGODB_CONNECTION_STRING string = ""
	DATABASE_NAME             string = ""
	CONFIGS_DATABASE_NAME     string = ""

	POSTGRES_CONNECTION_STRING string = ""
	SUPABASE_HOST              string = ""
	SUPABASE_DATABASE_NAME     string = ""
	SUPABASE_PORT              string = ""
	SUPABASE_USER              string = ""
	SUPABASE_PASSWORD          string = ""
	SUPABASE_PROJECT_URL       string = ""
	SUPABASE_API_KEY           string = ""

	SECRET_KEY string = ""
)

func initEnv() (err error) {
	if err = godotenv.Load(".env"); err != nil {
		return
	}

	// Defaults
	ENVIRONMENT = os.Getenv("ENVIRONMENT")
	PORT = os.Getenv("PORT")

	// Mongodb
	MONGODB_CONNECTION_STRING = os.Getenv("MONGODB_CONNECTION_STRING")
	DATABASE_NAME = os.Getenv("PRIMARY_DATABASE")
	CONFIGS_DATABASE_NAME = os.Getenv("CONFIGS_DATABASE")

	// Postgres
	POSTGRES_CONNECTION_STRING = os.Getenv("POSTGRES_CONNECTION_STRING")
	SUPABASE_HOST = os.Getenv("SUPABASE_HOST")
	SUPABASE_DATABASE_NAME = os.Getenv("SUPABASE_DATABASE_NAME")
	SUPABASE_PORT = os.Getenv("SUPABASE_PORT")
	SUPABASE_USER = os.Getenv("SUPABASE_USER")
	SUPABASE_PASSWORD = os.Getenv("SUPABASE_PASSWORD")
	SUPABASE_PROJECT_URL = os.Getenv("SUPABASE_PROJECT_URL")
	SUPABASE_API_KEY = os.Getenv("SUPABASE_API_KEY")

	// JWT
	SECRET_KEY = os.Getenv("SECRET_KEY")

	return
}

var _ = initEnv()
