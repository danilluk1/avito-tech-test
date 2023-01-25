package impl

import (
	"fmt"
	"github.com/danilluk1/avito-tech/internal/services/logger"
	"log"
	"strings"
)

type clogger struct{}

func NewLogger() logger.Logger {
	return &clogger{}
}

func (c clogger) Info(msg string, args ...string) {
	fmt.Printf("%s %s\n", msg, strings.Join(args, " "))
}

func (c clogger) Error(args ...any) {
	log.Println(args)
}
