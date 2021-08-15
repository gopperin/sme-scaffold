package config

import ()

// Case 全局Setup config
var (
	Case Configs
)

// Configs Configs
type Configs struct {
	Application Application `mapstructure:"application" json:"application" yaml:"application"`

	CORS CORS `mapstructure:"cors" json:"cors" yaml:"cors"`

	Database Database `mapstructure:"database" json:"database" yaml:"database"`
}
