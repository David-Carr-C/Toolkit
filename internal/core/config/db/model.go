package db

type DatabaseConfig struct {
	User      string   `yaml:"user"`
	Password  string   `yaml:"password"`
	Host      string   `yaml:"host"`
	Port      int      `yaml:"port"`
	Tunnel    bool     `yaml:"tunnel"`
	BackupDbs []string `yaml:"backup_dbs"`
}

type Config struct {
	Databases map[string]DatabaseConfig `yaml:"databases"`
}
