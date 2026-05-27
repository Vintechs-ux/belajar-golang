package database

var connection string

func init() {
	connection = "postgresql"
}

func GetDatabase() string {
	return connection
}
