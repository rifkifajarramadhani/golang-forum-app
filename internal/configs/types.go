package configs

type (
	Config struct {
		Service Service `mapsstructure:"service"`
	}

	Service struct {
		Port string `mapstructure:"port"`
	}
)