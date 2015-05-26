package config

type AppConfig struct {
	 SoftwareStatement SoftwareStatementConfig
}

type SoftwareStatementConfig struct {
	SigningMethod string `section:"softwarestatement" key:"signing_method"`
}