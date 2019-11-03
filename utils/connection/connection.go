package connection

import (
	"fmt"
	"os"

	"gopkg.in/couchbase/gocb.v1"
)

func Connection() *gocb.Bucket {
	var bucket *gocb.Bucket
	// cluster, _ := gocb.Connect("couchbase://localhost")
	fmt.Println("host :: ", os.Getenv("COUCHBASE_HOST"))
	fmt.Println("COUCHBASE_ADMINISTRATOR_USERNAME :: ", os.Getenv("COUCHBASE_ADMINISTRATOR_USERNAME"))
	fmt.Println("COUCHBASE_ADMINISTRATOR_PASSWORD :: ", os.Getenv("COUCHBASE_ADMINISTRATOR_PASSWORD"))
	cluster, _ := gocb.Connect("couchbase://" + os.Getenv("COUCHBASE_HOST"))
	cluster.Authenticate(gocb.PasswordAuthenticator{
		Username: os.Getenv("COUCHBASE_ADMINISTRATOR_USERNAME"),
		Password: os.Getenv("COUCHBASE_ADMINISTRATOR_PASSWORD"),
	})
	bucket, _ = cluster.OpenBucket(os.Getenv("COUCHBASE_BUCKET"), "")
	fmt.Println("host bucket:: ", bucket)

	return bucket
}
