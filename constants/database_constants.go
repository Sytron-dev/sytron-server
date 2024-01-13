package constants

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

// ---- Table names ----------------------------------------------------------------------
const (
	// AUTH
	TABLE_AUTH_BACKOFFICERS = "auth_backofficers"

	// CONFIGS
	TABLE_CONFIGS_COUNTRIES  = "countries"
	TABLE_CONFIGS_CITIES     = "cities"
	TABLE_CONFIGS_CURRENCIES = "currencies"

	// CMS
	TABLE_CMS_TOURIST_DESTINATIONS = "tourist_destinations"

	// OTHERS
	TABLE_BNB = "bnb"
)
