package connection

import (
	"fmt"

	"gopkg.in/couchbase/gocb.v1"
)

func Connection() *gocb.Bucket {
	var bucket *gocb.Bucket
	 cluster, _ := gocb.Connect("http://localhost")
	// cluster, _ := gocb.Connect("http://192.168.0.107:31053")
	// cluster, _ := gocb.Connect("couchbase://192.168.0.100")
	cluster.Authenticate(gocb.PasswordAuthenticator{
		Username: "softloft",
		Password: "password",
	})
	bucket, _ = cluster.OpenBucket("default", "")

	fmt.Println("host bucket:: ", bucket)

	return bucket
}
