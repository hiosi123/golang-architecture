package main

import (
	"fmt"

	architecture "github.com/hiosi123/golang-architecture"
	"github.com/hiosi123/golang-architecture/storage/mongo"
	"github.com/hiosi123/golang-architecture/storage/postgres"
)

func main() {
	dbm := mongo.Db{}
	dbp := postgres.Db{}

	p1 := architecture.Person{
		First: "Jenny",
	}

	p2 := architecture.Person{
		First: "James",
	}

	ps := architecture.NewPersonService(dbm)

	// or store in some other data storage
	architecture.Put(dbm, 1, p1)
	architecture.Put(dbm, 2, p2)
	fmt.Println(architecture.Get(dbm, 1))
	fmt.Println(architecture.Get(dbm, 2))

	fmt.Println(ps.Get(1))

	architecture.Put(dbp, 1, p1)
	architecture.Put(dbp, 2, p2)
	fmt.Println(architecture.Get(dbp, 1))
	fmt.Println(architecture.Get(dbp, 2))

}
