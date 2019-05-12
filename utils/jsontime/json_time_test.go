package jsontime

import (
	"github.com/labstack/gommon/log"
	"testing"
)

func TestTime(t *testing.T) {
	time := Time()
	log.Info(time)
}
