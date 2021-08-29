package config

var (
	DB     = &db{}
	App    = &app{}
	Server = &server{}
	Logger = &logger{}
)

type db struct {
	User     string `required:"true" envconfig:"POSTGRES_USER"`
	Password string `required:"true" envconfig:"POSTGRES_PASSWORD"`
	Host     string `required:"true" envconfig:"POSTGRES_HOST"`
	Port     uint   `required:"true" envconfig:"POSTGRES_PORT"`
	Name     string `required:"true" envconfig:"POSTGRES_NAME"`
	Schema   string `required:"true" envconfig:"POSTGRES_SCHEMA"`
}

type app struct {
	Tag string `required:"true" envconfig:"PROJECT_TAG"`
	Env string `required:"true" envconfig:"PROJECT_ENV"`
}

type server struct {
	Port uint `required:"true" envconfig:"SERVER_PORT"`
}

type logger struct {
	Level string `required:"true" envconfig:"LOGGER_LEVEL"`
}
