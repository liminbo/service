package mysql

type Config struct {
	Host     string
	Port     int
	User     string
	Password string
	Database string
	Charset  string

	TablePrefix string
	IdleConn    int
	MaxConn     int
}
