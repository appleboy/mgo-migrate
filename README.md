# mgo-migrate

[![GoDoc](https://godoc.org/github.com/appleboy/mgo-migrate?status.svg)](https://godoc.org/github.com/appleboy/mgo-migrate)
[![Build Status](http://drone.wu-boy.com/api/badges/appleboy/mgo-migrate/status.svg)](http://drone.wu-boy.com/appleboy/mgo-migrate)
[![codecov](https://codecov.io/gh/appleboy/mgo-migrate/branch/master/graph/badge.svg)](https://codecov.io/gh/appleboy/mgo-migrate)
[![Go Report Card](https://goreportcard.com/badge/github.com/appleboy/mgo-migrate)](https://goreportcard.com/report/github.com/appleboy/mgo-migrate)

Migrate function of MongoDB driver for Go

## How to use

```go
package main

import (
	"github.com/appleboy/mgo-migrate"
	"gopkg.in/mgo.v2"
)

func main() {
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
```
