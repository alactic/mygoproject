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
	fmt.Println("Starting go application ...")

	bucket = connection.Connection()
	router.Router()
	fmt.Println("host name :: ", os.Getenv("COUCHBASE_HOST"))
	fmt.Println("bucket name :: ", bucket.Name())

}
