package settings

type mySQL struct {
	UserName string `yaml:"username"`
	Password string `yaml:"username"`
	Address  string `yaml:"address"`
	Port     string `yaml:"port"`
	Database string `yaml:"database"`
}
