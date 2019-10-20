package connection

import "gopkg.in/couchbase/gocb.v1"

var bucket *gocb.Bucket

func Connection() *gocb.Bucket {
	cluster, _ := gocb.Connect("couchbase://localhost")
	cluster.Authenticate(gocb.PasswordAuthenticator{
		Username: "elvis",
		Password: "password",
	})
	bucket, _ = cluster.OpenBucket("demoproject", "")
	return bucket
}
