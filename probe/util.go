package probe

import (
	"log"
	"os"
)

func checkError(err error) {
	if err == nil {
		return
	}
	log.Println(err)
	os.Exit(1)
}
