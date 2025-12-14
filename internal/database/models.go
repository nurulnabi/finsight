package database

type DBConfig interface {
	GetDriver() string
}

type SQLConfig struct {
	PORT     string `json:"port"`
	HOST     string `json:"host"`
	DB_URI   string `json:"db_uri"`
	USER     string `json:"user"`
	PASSWORD string `json:"password"`
	NAME     string `json:"name"`
}

func (sc SQLConfig) GetDriver() string {
	return "sql"
}
