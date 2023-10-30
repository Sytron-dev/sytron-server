package constants

import (
	"os"

	"github.com/joho/godotenv"
)

var (
	MONGODB_CONNECTION_STRING string = ""
	DATABASE_NAME             string = ""
	CONFIGS_DATABASE_NAME     string = ""
)

func initEnv() (err error) {
	error := godotenv.Load(".env")
	if error != nil {
		return
	}

	// Database constants from the .env file
	MONGODB_CONNECTION_STRING = os.Getenv("MONGODB_CONNECTION_STRING")
	DATABASE_NAME = os.Getenv("PRIMARY_DATABASE")
	CONFIGS_DATABASE_NAME = os.Getenv("CONFIGS_DATABASE")

	return
}

var initErr = initEnv()

// ---- Collection names -----------------------------------------------------------------
const (
	// Users
	PROFILES_COLLECTION_BACK_OFFICERS = "profiles_back_officers"
	PROFILES_COLLECTION_USERS         = "profiles_users"
	USERS_COLLECTION                  = "users" //! deprecated
	PROFILES_COLLECTION_MERCHANTS     = "profiles_merchants"

	// User Credentials
	CREDENTIALS_COLLECTION_BACK_OFFICERS = "credentials_back_officers"
	CREDENTIALS_COLLECTION_USERS         = "credentials_users"
	CREDENTIALS_COLLECTION_MERCHANTS     = "credentials_merchants"

	// CMS
	CMS_COLLECTION_LOCATIONS    = "cms_locations"
	CMS_COLLECTION_DESTINATIONS = "cms_destinations"
	CMS_COLLECTION_ACTIVITIES   = "cms_activities"

	// Company
	COMPANIES_COLLECTION = "companies"

	// Provisions (Booking and Recreation)
	ACCOMMODATIONS_COLLECTION = "accommodations"
	EXPERIENCES_COLLECTION    = "experiences"

	// Events
	EVENTS_COLLECTION        = "events"
	EVENT_TICKETS_COLLECTION = "event_tickets"

	// Configs
	CITIES_COLLECTION    = "cities"
	COUNTRIES_COLLECTION = "countries"
)
