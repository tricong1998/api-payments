package models

import (
	"api-payments/database"
)

// Database name
var databaseName = "api-payments"

// Create a connection
var dbConnect = database.NewDatastore(databaseName)
