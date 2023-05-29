package lib

import (
	"fmt"
	"log"
)

func Logger(err error) {
	if err != nil {
		log.Fatal(fmt.Sprintf("Error:\n%s", err))
		panic("ERROR!!")
	} else {
		log.Printf("Successfuly continued")
	}
}
