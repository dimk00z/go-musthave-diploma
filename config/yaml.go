package config

// TODO REMOVE
import (
	"log"
	"os"
)

const yaml = `
app:
  name: 'go-musthave-diploma'
  version: '1.0.0'

http:
  run_address: 'localhost:8080'

logger:
  log_level: 'debug'
  rollbar_env: 'go-musthave-diploma'

postgres:
  pool_max: 2
  pg_url: 'postgres://admin:password@localhost:5432/gophermart?sslmode=disable'

security:
# only for debug
  secret_key: 'SECRET_KEY'
  token_hour_lifespan: 96
  cookie_token_name: 'gophermart_token'

workers:
  workers_number: 5
  pool_length: 10

api:
  accrual_system_address: 'pass'`

func yamlToFile(filePath string) {
	err := os.WriteFile(filePath, []byte(yaml), 0644)
	if err != nil {
		log.Fatalln(err)
	}
}
