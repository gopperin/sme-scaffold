package config

import ()

// DatabaseStruct DatabaseStruct
type DatabaseStruct struct {
	Dialect      string
	Database     string
	User         string
	Password     string
	Host         string
	Port         int
	Charset      string
	URL          string
	MaxIdleConns int
	MaxOpenConns int
	LogMode      string
}

// Database Database
var Database = new(DatabaseStruct)
