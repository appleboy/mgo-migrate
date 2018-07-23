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
	err = m.RollbackLast()
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

func TestMissingID(t *testing.T) {
	session, err := mgo.Dial(dbHost)
	assert.Nil(t, err)
	defer session.Close()

	newMigrations := append(migrations, &Migration{
		ID: "",
		Migrate: func(s *mgo.Session) error {
			return nil
		},
		Rollback: func(s *mgo.Session) error {
			return nil
		},
	})

	m := New(session, dbName, DefaultOptions, newMigrations)
	err = m.Migrate()
	assert.NotNil(t, err)
	assert.Equal(t, ErrMissingID, err)

	err = dropDB(session)
	assert.Nil(t, err)
}

func TestErrNoMigrationDefined(t *testing.T) {
	session, err := mgo.Dial(dbHost)
	assert.Nil(t, err)
	defer session.Close()

	// empty migration
	m := New(session, dbName, DefaultOptions, []*Migration{})
	err = m.Migrate()
	assert.Nil(t, err)
	err = m.RollbackLast()
	assert.NotNil(t, err)
	assert.Equal(t, ErrNoMigrationDefined, err)

	err = dropDB(session)
	assert.Nil(t, err)
}

func TestErrRollbackImpossible(t *testing.T) {
	session, err := mgo.Dial(dbHost)
	assert.Nil(t, err)
	defer session.Close()

	// empty migration
	m := New(session, dbName, DefaultOptions, []*Migration{{
		ID: "201709201400",
		Migrate: func(s *mgo.Session) error {
			return nil
		},
	}})
	err = m.Migrate()
	assert.Nil(t, err)
	err = m.RollbackLast()
	assert.NotNil(t, err)
	assert.Equal(t, ErrRollbackImpossible, err)

	err = dropDB(session)
	assert.Nil(t, err)
}
