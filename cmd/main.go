package main

import (
	"DBsummer/appDir"
	"github.com/unidoc/unipdf/v3/common/license"
	"log"
	"math/rand"
	"time"
)

func main() {
	err := license.SetMeteredKey(`92336d25a4a8ea959e7f560bec6cb054eec41af24552c8da56354090cb2552a2`)
	if err != nil {
		return
	}

	rand.Seed(time.Now().Unix())
	appNew, err1 := appDir.New()
	if err1 != nil {
		log.Fatal(err1)
	}

	err1 = appNew.Run()
	if err1 != nil {
		log.Fatal(err1)
	}
}

//id, err1 := appNew.repository.TableNew().Create(context.Background())
//if err1 != nil {
//	log.Fatal(err1)
//}
//
//log.Println(id)
