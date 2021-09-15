package configuration

type Configuration struct {
	Server struct {
		Enviroment string `yaml:"enviroment"`
		Port       string `yaml:"port"`
		Host       string `yaml:"host"`
	} `yaml:"server"`
	Database struct {
		ConnectionString string `yaml:"connectionString"`
		Database         string `yaml:"database"`
	} `yaml:"database"`
	Messaging struct {
		URL string `yaml:"url"`
	} `yaml:"messaging"`
}
