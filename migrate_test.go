package migrate

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/mgo.v2"
)

var (
	dbHost     = "mongodb"
	dbName     = "test_migrate"
	migrations = []*Migration{
		{
			ID: "201709201400",
			Migrate: func(s *mgo.Session) error {
				return nil
			},
			Rollback: func(s *mgo.Session) error {
				return nil
			},
		},
	}

	initSchema = func(s *mgo.Session) error {
		return nil
	}
)

func dropDB(s *mgo.Session) error {
	db := s.DB(dbName)

	return db.DropDatabase()
}

func TestMigrate(t *testing.T) {
	session, err := mgo.Dial(dbHost)
	assert.Nil(t, err)
	defer session.Close()

	m := New(session, dbName, DefaultOptions, migrations)
	err = m.Migrate()
	assert.Nil(t, err)

	err = dropDB(session)
	assert.Nil(t, err)
}

func TestInitSchema(t *testing.T) {
	session, err := mgo.Dial(dbHost)
	assert.Nil(t, err)
	defer session.Close()

	m := New(session, dbName, DefaultOptions, migrations)
	m.InitSchema(initSchema)
	err = m.Migrate()
	assert.Nil(t, err)

	err = dropDB(session)
	assert.Nil(t, err)
}
