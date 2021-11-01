package common

const (
	EMPLOYEE   string = "EMPLOYEE"
	CONTRACTOR string = "CONTRACTOR"
)

const (
	REST_PORT string = "8000"
	REST_URL  string = "http://localhost:" + REST_PORT
)

//TODO Arquivo de configuração
// MongoDB
const (
	MONGO_URL     = "mongodb://localhost:27017/"
	MONGO_DB      = "mongodb"
	MONGO_TIMEOUT = 60000
)

// Postgres
const (
	POSTGRES_URL     = "mongodb://localhost:27017/"
	POSTGRES_DB      = "mongodb"
	POSTGRES_TIMEOUT = 60000
)
