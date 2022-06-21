package config

import (
	"flag"
	"log"
	"os"
	"sync"

	"github.com/ilyakaznacheev/cleanenv"
	tld "github.com/jpillora/go-tld"
)

type (
	// Config -.
	Config struct {
		App      `yaml:"app"`
		HTTP     `yaml:"http"`
		Log      `yaml:"logger"`
		PG       `yaml:"postgres"`
		Security `yaml:"security"`
		Workers  `yaml:"workers"`
		API      `yaml:"api"`
	}

	App struct {
		Name    string `env-required:"true" yaml:"name"    env:"APP_NAME"`
		Version string `env-required:"true" yaml:"version" env:"APP_VERSION"`
	}

	HTTP struct {
		RunAddress string `env-required:"true" yaml:"run_address" env:"RUN_ADDRESS"`
		DomainName string
	}

	Log struct {
		Level string `env-required:"true" yaml:"log_level"   env:"LOG_LEVEL"`
	}

	PG struct {
		PoolMax int    `env-required:"true" yaml:"pool_max" env:"PG_POOL_MAX"`
		URL     string `env-required:"true" yaml:"pg_url" env:"DATABASE_URI"`
	}

	Security struct {
		SecretKey         string `env-required:"true" yaml:"secret_key" env:"SECRET_KEY"`
		TokenHourLifespan int    `env-required:"true" yaml:"token_hour_lifespan" env:"TOKEN_HOUR_LIFESPAN"`
		CookieTokenName   string `env-required:"true" yaml:"cookie_token_name" env:"COOKIE_TOKEN_NAME"`
	}

	Workers struct {
		WorkersNumber int `env-required:"true" yaml:"workers_number" env:"WORKERS_NUMBER"`
		PoolLength    int `env-required:"true" yaml:"pool_length" env:"POOL_LENGTH"`
	}
	API struct {
		AccrualSystemAddress string `env-required:"true" yaml:"accrual_system_address" env:"ACCRUAL_SYSTEM_ADDRESS"`
	}
)

func (c *Config) checkFlags() {
	flagAddress := flag.String("a", "", "RUN_ADDRESS")
	flagDSN := flag.String("d", "", "DATABASE_URI")
	flagASA := flag.String("r", "", "ACCRUAL_SYSTEM_ADDRESS")
	flag.Parse()
	if *flagAddress != "" {
		c.HTTP.RunAddress = *flagAddress
	}
	if *flagDSN != "" {
		c.PG.URL = *flagDSN
	}
	if *flagASA != "" {
		c.API.AccrualSystemAddress = *flagASA
	}
}

var (
	cfg  *Config
	once sync.Once
)

// NewConfig returns app config.
func NewConfig() (*Config, error) {
	var err error

	once.Do(func() {
		cfg = &Config{}

		tempYML := "config.yml"
		localYML := "../../config/config.yml"
		_, err := os.Stat(localYML)
		currentConfig := localYML
		if os.IsNotExist(err) {
			log.Println(localYML + " doesn't exist")
			yamlToFile(tempYML)
			currentConfig = tempYML
		}

		err = cleanenv.ReadConfig(currentConfig, cfg)
		if err != nil {
			log.Fatalf("config error: %v", err)
			return
		}

		err = cleanenv.ReadEnv(cfg)
		if err != nil {
			log.Fatalf("readenv error: %v", err)
			return
		}
		cfg.checkFlags()
		url, err := tld.Parse(cfg.HTTP.RunAddress)
		if err != nil {
			log.Fatalf("domain parsing error: %v", err)
			return
		}
		cfg.HTTP.DomainName = url.Domain
	})
	if err != nil {
		return nil, err
	}
	return cfg, nil
}
