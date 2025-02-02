package configs

type (
	Config struct {
		Service  Service  `mapsstructure:"service"`
		Database Database `mapsstructure:"database"`
	}

	Service struct {
		Port      string `mapstructure:"port"`
		JWTSecret string `mapstructure:"jwt_secret"`
	}

	Database struct {
		Datasource string `mapstructure:"datasource"`
	}
)
