package main

import (
	"fmt"
	"log"

	"github.com/samersawan/gator/internal/config"
)

func main() {

	c, err := config.Read()
	if err != nil {
		log.Fatal(err)
	}
	c.SetUser("samersawan")
	updatedCfg, err := config.Read()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(updatedCfg)

}
