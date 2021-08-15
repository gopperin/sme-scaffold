package config

import ()

// Database Database
type Database struct {
	Dialect      string `mapstructure:"dialect" json:"dialect" yaml:"dialect"`
	Database     string `mapstructure:"database" json:"database" yaml:"database"`
	User         string `mapstructure:"user" json:"user" yaml:"user"`
	Password     string `mapstructure:"password" json:"password" yaml:"password"`
	Host         string `mapstructure:"host" json:"host" yaml:"host"`
	Port         int    `mapstructure:"port" json:"port" yaml:"port"`
	Charset      string `mapstructure:"charset" json:"charset" yaml:"charset"`
	URL          string `mapstructure:"url" json:"url" yaml:"url"`
	MaxIdleConns int    `mapstructure:"maxIdleConns" json:"maxIdleConns" yaml:"maxIdleConns"`
	MaxOpenConns int    `mapstructure:"maxOpenConns" json:"maxOpenConns" yaml:"maxOpenConns"`
	LogMode      string `mapstructure:"logMode" json:"logMode" yaml:"logMode"`
}
