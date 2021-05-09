package main

import (
	"DBsummer/appDir"
	"context"
	"log"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	appNew, err1 := appDir.New()
	if err1 != nil {
		log.Fatal(err1)
	}

	//err1 = appNew.Run()
	//if err1 != nil {
	//	log.Fatal(err1)
	//}
	id, err1 := appNew.Repository.TableNew().Create(context.Background())
	if err1 != nil {
		log.Fatal(err1)
	}

	log.Println(id)
}
