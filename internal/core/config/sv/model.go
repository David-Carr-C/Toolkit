package sv

type ServerConfig struct {
	Host string `yaml:"host"`
	Path string `yaml:"path"`
}

type Config struct {
	Servers map[string]ServerConfig `yaml:"servers"`
}
