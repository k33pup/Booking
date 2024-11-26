package logger

import (
	"log"
)

//mock logger -> Grafana (future)

func WriteLog(str string) error {
	log.Fatalln(str)
	return nil
}
