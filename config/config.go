package config

type dbConstants struct {
	PostgresURI string
	MysqlURI string


}

var DBConstants = dbConstants {
	PostgresURI: "http://www.google.com",
	MysqlURI: "",

}


var _ = DBConstants
