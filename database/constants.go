package database

import (
	"os"

	"github.com/joho/godotenv"
)

// Load the .env file first before fe==getting environment variables
var _ = godotenv.Load(".env")

// Database constants from the .env file
var (
	MONGODB_CONNECTION_STRING = os.Getenv("MONGODB_CONNECTION_STRING")
	DATABASE_NAME             = os.Getenv("PRIMARY_DATABASE")
	CONFIGS_DATABASE_NAME     = os.Getenv("CONFIGS_DATABASE")
)

// ---- Collection names -----------------------------------------------------------------
const (
	// Users
	BACKOFFICERS_COLLECTION = "backofficers"
	USERS_COLLECTION        = "users"
	MERCHANTS_COLLECTION    = "merchant"

	// CMS
	LOCATIONS_COLLECTION    = "locations"
	DESTINATIONS_COLLECTION = "destinations"
	ACTIVITIES_COLLECTION   = "activities"

	// Provisions (Booking and Recreation)
	ACCOMMODATIONS_COLLECTION = "accommodations"
	EXPERIENCES_COLLECTION    = "experiences"

	// Events
	EVENTS_COLLECTION        = "events"
	EVENT_TICKETS_COLLECTION = "eventTickets"

	// Configs
	CITIES_COLLECTION    = "cities"
	COUNTRIES_COLLECTION = "countries"
)
