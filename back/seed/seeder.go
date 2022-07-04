package seed

import (
	"fmt"
	"github.com/Dan-Yyyy/vendingPanel.git/pkg/message"
	"github.com/jmoiron/sqlx"
	"log"
	"reflect"
)

type Seed struct {
	db *sqlx.DB
}

func Execute(db *sqlx.DB, seedMethodNames ...string) {
	s := Seed{db}

	seedType := reflect.TypeOf(s)

	if len(seedMethodNames) == 0 {
		log.Println("Running all seeder...")
		for i := 0; i < seedType.NumMethod(); i++ {
			method := seedType.Method(i)
			seed(s, method.Name)
		}
	}

	for _, item := range seedMethodNames {
		seed(s, item)
	}
}

func seed(s Seed, seedMethodName string) {
	m := reflect.ValueOf(s).MethodByName(seedMethodName)
	if !m.IsValid() {
		log.Fatal("No method called ", seedMethodName)
	}
	log.Println("Seeding", seedMethodName, "...")
	m.Call(nil)
	log.Println("Seed", seedMethodName, "succeed")
}

func processDictionarySeed(s Seed, data []string, tableName string) {
	for _, value := range data {
		query := fmt.Sprintf("INSERT INTO %s (name) values ($1)", tableName)
		_, err := s.db.Exec(query, value)

		if err != nil {
			log.Fatalf("%s на значении %s: %s", message.SeedingError, value, err.Error())
			return
		}
	}
}
