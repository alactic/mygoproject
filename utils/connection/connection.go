package connection

import "gopkg.in/couchbase/gocb.v1"

func Connection() *gocb.Bucket {
	var bucket *gocb.Bucket
	cluster, _ := gocb.Connect("couchbase://10.4.3.42")
	cluster.Authenticate(gocb.PasswordAuthenticator{
		Username: "elvis",
		Password: "password",
	})
	bucket, _ = cluster.OpenBucket("demoproject", "")
	return bucket
}
