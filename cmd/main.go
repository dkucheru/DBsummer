package main

import (
	"DBsummer/appDir"
	"github.com/unidoc/unipdf/v3/common/license"
	"log"
	"math/rand"
	"time"
)

func main() {
	err := license.SetMeteredKey(`57ea372ede16a7f5e5cec5b38036aa87adae616fb05784f6e77dba5389870483`)
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
