package main

import (
	"fmt"
	"os"

	"github.com/alactic/mygoproject/utils/connection"
	"github.com/alactic/mygoproject/utils/router"
	"gopkg.in/couchbase/gocb.v1"
)

var bucket *gocb.Bucket

func main() {

	bucket = connection.Connection()
	router.Router()
	fmt.Println("bucket name1 :: ", bucket.Name())
	fmt.Println("host name 1:: ", os.Getenv("COUCHBASE_HOST"))
	fmt.Println("Starting go1 lang application ...")

	fmt.Println("host name :: ", os.Getenv("COUCHBASE_HOST"))
	fmt.Println("bucket name :: ", bucket.Name())

}
