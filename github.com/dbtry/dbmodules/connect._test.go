package dbmodules

import (
	"reflect"
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func TestExampleSuccess(t *testing.T) {
	result := GormConnect()

	if reflect.TypeOf(result) != *gorm.DB {
		t.Fatal("failed test")
	}
}
