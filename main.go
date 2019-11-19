package main

import (
	"fmt"
	"github.com/alactic/mygoproject/utils/connection"
	"github.com/alactic/mygoproject/utils/router"
	"gopkg.in/couchbase/gocb.v1"
	"os"
)

var bucket *gocb.Bucket

func main() {
	fmt.Println("starting application")
	os.Setenv("SO_TIMEOUT", "8000")
	os.Setenv("CONNECTION_TIMEOUT", "12000")
	bucket = connection.Connection()
	router.Router()
}
