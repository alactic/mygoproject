package connection

import (
	"fmt"
	"os"

	"gopkg.in/couchbase/gocb.v1"
)

func Connection() *gocb.Bucket {
	var bucket *gocb.Bucket
	// cluster, _ := gocb.Connect("couchbase://localhost")
	fmt.Println("host :: ", os.Getenv("COUCHBASEHOST"))
	cluster, _ := gocb.Connect("couchbase://" + os.Getenv("COUCHBASEHOST"))
	cluster.Authenticate(gocb.PasswordAuthenticator{
		Username: "user",
		Password: "password",
	})
	bucket, _ = cluster.OpenBucket("mygodb", "")
	return bucket
}
