package i18n

import (
	"fmt"
	"log"
	"os"
)

func main() {

	dir, _ := os.Getwd()

	err := Start(dir + "/examples/", PtBR)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(Message(ResourceWithBodyEmpty))

}
