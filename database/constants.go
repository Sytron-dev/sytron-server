package database

import "os"

var DATABASE_NAME = os.Getenv("PRIMARY_DATABASE")
var CONFIGS_DATABASE_NAME = os.Getenv("CONFIGS_DATABASE")
