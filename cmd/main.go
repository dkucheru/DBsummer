package main

import (
	"DBsummer/appDir"
	"github.com/unidoc/unipdf/v3/common/license"
	"log"
	"math/rand"
	"time"
)

func main() {
	err := license.SetMeteredKey(`10bc0bc50f7829dffd2a7e8b87d38a8ab09775e890aa1a92db4a8f2d70c695a8`)
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
