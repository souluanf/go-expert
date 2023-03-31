package configs

var cfg *conf

type conf struct {
	DBDriver      string
	DBHost        string
	DBPort        string
	DBName        string
	DBUser        string
	DBPassword    string
	WebServerPort string
	JWTSecret     string
	JWTExpiresIn  int
}

func loadConfig(path string) (*conf, error) {
	// ...
}
