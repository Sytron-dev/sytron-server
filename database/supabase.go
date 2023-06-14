package database

import (
	"os"

	supa "github.com/nedpals/supabase-go"
)

func GetClient() *supa.Client {
	supabaseUrl := os.Getenv("supabaseUrl")
	supabaseKey := os.Getenv("supabaseKey")
	// supabaseApiKey := os.Getenv("supabaseApiKey")
	return supa.CreateClient(supabaseUrl, supabaseKey, true)
}
