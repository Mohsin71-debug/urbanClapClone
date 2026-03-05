package config

type Config struct {
        DB *DBConfig
}

type DBConfig struct {
        Dialect  string
        Host     string
        Username string
        Password string
        Name     string
        Charset  string
}

func GetConfig() *Config {
        return &Config{
                DB: &DBConfig{
                        Dialect:  "mysql",
                        Host:     "mysql",
                        Username: "root",
                        Password: "123456",
                        Name:     "urbanClap",
                        Charset:  "utf8",
                },
        }
}
