package migrate_test

import (
	"github.com/appleboy/mgo-migrate"
	"gopkg.in/mgo.v2"
)

func Example() {
	session, err := mgo.Dial("127.0.0.1")
	if err != nil {
		panic(err)
	}

	m := migrate.New(session, "test_db", migrate.DefaultOptions, []*migrate.Migration{{
		ID: "201709201400",
		Migrate: func(s *mgo.Session) error {
			return nil
		},
		Rollback: func(s *mgo.Session) error {
			return nil
		},
	}})

	if err := m.Migrate(); err != nil {
		panic(err)
	}

	if err := m.RollbackLast(); err != nil {
		panic(err)
	}
}
