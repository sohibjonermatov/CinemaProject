package models

type Config struct {
	App App `json:"app"`
	DB  DB  `json:"db"`
}

type App struct {
	Port int `json:"port"`
}

type DB struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	DBName   string `json:"dbname"`
}
