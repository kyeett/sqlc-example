package main

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"math/rand"
	"time"

	"github.com/kyeett/sqlc-example/data"
	"log"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = ""
)

var sourceName = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, "")

func main() {
	// Seed randomizer
	rand.Seed(time.Now().UTC().UnixNano())

	db, err := sql.Open("postgres", sourceName)
	if err != nil {
		log.Fatalf("failed to open connection to DB: %v", err)
	}
	database := data.New(db)

	// Create animal
	err = database.CreateAnimal(context.Background(), data.CreateAnimalParams{
		Name: "Fido",
		Type: "dog",
	})
	if err != nil {
		log.Fatal("failed to create animal", err)
	}
	fmt.Println("* new animal with created")

	// List animals
	animals, err := database.ListAnimals(context.Background())
	if err != nil {
		log.Fatal("failed to list animals", err)
	}

	fmt.Println("* complete list of animals")
	for _, animal := range animals {
		fmt.Printf("-- %d\t%s\t%s\n", animal.ID, animal.Name, animal.Type)
	}
}

//func main() {
//	// Seed randomizer
//	rand.Seed(time.Now().UTC().UnixNano())
//
//	db, err := sql.Open("postgres", sourceName)
//	if err != nil {
//		log.Fatalf("failed to open connection to DB: %v", err)
//	}
//	database := data.New(db)
//
//	// Create animal
//	err = database.CreateAnimal(context.Background(), data.CreateAnimalParams{
//		Name: randomAnimalName(),
//		Type: randomAnimalType(),
//	})
//	if err != nil {
//		log.Fatal("failed to create animal", err)
//	}
//
//	// List animals
//	animals, err := database.ListAnimals(context.Background())
//	if err != nil {
//		log.Fatal("failed to list animals", err)
//	}
//
//	fmt.Println("* complete list of animals")
//	for _, animal := range animals {
//		fmt.Printf("-- %d\t%s\t%s\n", animal.ID, animal.Name, animal.Type)
//	}
//
//	// Delete random animal
//	n := rand.Intn(len(animals))
//	selectedAnimal := animals[n]
//	da, err := database.DeleteAnimal(context.Background(), selectedAnimal.ID)
//	if err != nil {
//		log.Fatal("failed to delete animal", err)
//	}
//	fmt.Printf("* animal with id %d deleted\n", da.ID)
//}
