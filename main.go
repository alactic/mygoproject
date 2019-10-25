package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/alactic/mygoproject/utils/connection"
	"github.com/alactic/mygoproject/utils/router"
	"gopkg.in/couchbase/gocb.v1"
)

var bucket *gocb.Bucket

func main() {
	fmt.Println("Starting application ...")

	bucket = connection.Connection()
	router.Router()
	
}
