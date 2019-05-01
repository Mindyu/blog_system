package md5

import (
	"github.com/labstack/gommon/log"
	"testing"
)

func TestEncryptPassword(t *testing.T) {
	password := EncryptPassword("123456")
	log.Info(password)
}
