package main

import (
	"fmt"

	"github.com/alactic/mygoproject/utils/router"
)

// var bucket *gocb.Bucket

func main() {
	fmt.Println("Starting application ...")

	// bucket = connection.Connection()
	router.Router()
}
