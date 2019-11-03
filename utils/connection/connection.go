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
	cluster, _ := gocb.Connect("couchbase://" + os.Getenv("COUCHBASE_HOST"))
	// cluster.Authenticate(gocb.PasswordAuthenticator{
	// 	Username: "user",
	// 	Password: "password",
	// })
	bucket, _ = cluster.OpenBucket(os.Getenv("COUCHBASE_BUCKET"), "")
	return bucket
}
