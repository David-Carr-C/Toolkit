package constants

func GetArchives() []string {
	archives := []string{
		"APIs/Codigos/",
		"APIs/ComplementoDescarga/",
		"APIs/Complementos/",
		"APIs/CSD/",
		"APIs/DM/log_proceso/",
		"APIs/DM/zip_files/",
		"APIs/DM/zip_missing/",
		"APIs/DM/zip_processed/",
		"APIs/Expediente/Descargas/",
		"APIs/Facturas/",
		"APIs/FacturasCadenas/",
		"APIs/FacturasCanceladas/",
		"APIs/FacturasDescarga/",
		"APIs/libs/PHP8/log_proceso/",
		"APIs/log/",
		"APIs/logProcesos/",
		"APIs/pcbFiles/",
		"APIs/Prospecto/",
		"APIs/Proveedor/",
		"APIs/Reportes/",
	}

	return archives
}

/*
type PCB struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
}

arreglo := []PCB{
	{
		Host:     "localhost",
		Port:     3306,
		User:     "admin",
		Password: "root",
	}, {
		Host:     "localhost",
		Port:     3306,
		User:     "admin",
		Password: "root",
	},
}

*/
