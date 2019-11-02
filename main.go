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
	fmt.Println("bucket name 2 ELVIS ")
	fmt.Println("host name 1:: ", os.Getenv("COUCHBASEHOST"))
	// fmt.Println("Starting go1 lang application ...")

	// fmt.Println("host name :: ", os.Getenv("COUCHBASEHOST"))
	// fmt.Println("bucket name :: ", bucket.Name())

	bucket = connection.Connection()
	router.Router()
	// fmt.Println("host name 1:: ", os.Getenv("COUCHBASE_HOST"))
	// fmt.Println("Starting go1 lang application ...")

	// fmt.Println("host name :: ", os.Getenv("COUCHBASE_HOST"))
	// fmt.Println("bucket name :: ", bucket.Name())

}
