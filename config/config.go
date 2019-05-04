package utils

type Config struct {
	version         string
	PrefixUrl       string
	JwtSecret       string
	PageSize        int
	RuntimeRootPath string
	LogSavePath     string
	LogSaveName     string
	LogFileExt      string
	TimeFormat      string
}

type Server struct {
	RunMode      string
	HttpPort     int
	ReadTimeout  int
	WriteTimeout int
}

type database struct {
	Type        string
	User        string
	PassWord    string
	Host        string
	Port        int
	Name        string
	TablePrefix string
}
