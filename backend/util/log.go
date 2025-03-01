package util

import "log"

func Log(service string, msg string) {
	log.Printf("[%s] %s", service, msg)
}
