package model

type Log struct {
	ID                int64
	Element           string
	Message           string
	Exception         string
	Level             string
	SourceServer      string
	DestinationServer string
	DestinationPath   string
	BackupSHA256      string
	DatabaseSHA256    string
	CreatedAt         string
	Synchronized      bool
}
