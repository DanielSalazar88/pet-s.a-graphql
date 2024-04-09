package settings

type mysqlFields struct {
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Address  string `yaml:"address"`
	Port     string `yaml:"port"`
	Database string `yaml:"database"`
}

type handlersFields struct {
	Port string `yaml:"port"`
}
