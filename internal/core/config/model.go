package config

type DirConfig struct {
	BackupDirs []string `yaml:"backup_dirs"`
}

type Config struct {
	Projects map[string]DirConfig `yaml:"projects"`
}
