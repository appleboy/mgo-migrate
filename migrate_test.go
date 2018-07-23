package migrate

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/mgo.v2"
)

var (
	dbHost = "mongodb"
	dhName = "test_migrate"
)

func TestMigrate(t *testing.T) {
	session, err := mgo.Dial(dbHost)
	assert.Nil(t, err)
	defer session.Close()
}
