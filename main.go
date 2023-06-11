package main

// import (
// 	"net/http"
// 	"fmt"
// 	"log"
// 	"github.com/nedpals/supabase-go"
// )
import (	
	"log"
	"github.com/joho/godotenv"
	supa "github.com/nedpals/supabase-go"
) 

type Country struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Capital string `json:"capital"`
  }

func main(){
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	supabaseUrl := "supabaseUrl"
	supabaseKey := "supabaseKey"
	supabase := supa.CreateClient(supabaseUrl, supabaseKey)
}
