package lib

import (
	"errors"
	"fmt"
	"math"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type ConfigManager interface {
	NewConfig() Config
	//@description: server port getter/setters.
	SetPort(int) error
	GetPort() uint16
	//@description: DB_Uri Getter and Setters.
	SetDBUri(string)
	GetDBUri() string
}

type Config struct {
	Port    uint16
	DBUri   string
	Proxies []string
	APIKeys map[string]string
}

func NewConfig() Config {
	godotenv.Load()
	sENV_Port := os.Getenv("PORT")
	iENV_Port := 5000

	if sENV_Port != "" {
		port, err := strconv.Atoi(sENV_Port)
		Logger(err)

		fmt.Println("Port are setted on", port)
		iENV_Port = port
	}
	db_uri := os.Getenv("DB_URI")
	return Config{
		Port:    uint16(iENV_Port),
		DBUri:   db_uri,
		Proxies: []string{"0.0.0.0"},
		APIKeys: map[string]string{},
	}
}
func (c *Config) SetPort(port int) error {
	if (port > 0) && (port < int(math.Pow(2, 16))) {
		c.Port = uint16(port)
	} else {
		return errors.New("You cannot assign this value for port.")
	}
	return nil
}

func (c *Config) GetPort() uint16 {
	return c.Port
}

func (c *Config) SetDBUri(uri string) {
	c.DBUri = uri
}

func (c *Config) SetApiKey(key, value string) {
	c.APIKeys[key] = value
}

func (c *Config) GetApiKey(key string) string {
	return c.APIKeys[key]
}
