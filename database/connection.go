package database

var driver string

func init() {
	driver = "PostgreSQL"
}

func GetDBDriver() string {
	return driver
}
