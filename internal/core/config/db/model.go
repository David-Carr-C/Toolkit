package db

type DbConfig struct {
	BackupDbs []string `yaml:"backup_dbs"`
}

type DatabaseConfig struct {
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
}

type Config struct {
	DatabaseConfig DatabaseConfig
	Databases      map[string]DbConfig `yaml:"databases"`
}
