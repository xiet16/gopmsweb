package conf

type DbConfig struct {
	DriverName   string
	Dsn          string
	ShowSql      bool
	ShowExecTime bool
	MaxIdle      int
	MaxOpen      int
}

var Db = map[string]DbConfig{
	"db": {
		DriverName:   "mysql",
		Dsn:          "",
		ShowSql:      true,
		ShowExecTime: false,
		MaxIdle:      10,
		MaxOpen:      200,
	},
}
