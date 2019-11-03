package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/speps/go-hashids"
	"gopkg.in/couchbase/gocb.v1"
)

type MyUrl struct {
	ID       string `json:"id,omitempty"`
	LongUrl  string `json:"longUrl,omitempty"`
	ShortUrl string `json:"shortUrl,omitempty"`
}

var bucket *gocb.Bucket
var bucketName string

func CreateIndex(response http.ResponseWriter, request *http.Request) {
	response.Write([]byte(`{"message": "my go docker demo project is working demo 2"}`))
}

func ExpandEndpoint(w http.ResponseWriter, req *http.Request) {
	var n1qlParams []interface{}
	query := gocb.NewN1qlQuery("SELECT `" + bucketName + "`.* FROM `" + bucketName + "` WHERE shortUrl = $1")
	params := req.URL.Query()
	n1qlParams = append(n1qlParams, params.Get("shortUrl"))
	rows, _ := bucket.ExecuteN1qlQuery(query, n1qlParams)
	var row MyUrl
	rows.One(&row)
	json.NewEncoder(w).Encode(row)
}

func CreateEndpoint(w http.ResponseWriter, req *http.Request) {
	fmt.Println("bucket1")
	var url MyUrl
	_ = json.NewDecoder(req.Body).Decode(&url)
	var n1qlParams []interface{}
	n1qlParams = append(n1qlParams, url.LongUrl)
	fmt.Println("bucket :: ", bucket)
	fmt.Println("bucket name :: ", bucketName)
	query := gocb.NewN1qlQuery("SELECT `" + bucketName + "`.* FROM `" + bucketName + "` WHERE longUrl = $1")
	rows, err := bucket.ExecuteN1qlQuery(query, n1qlParams)
	if err != nil {
		w.WriteHeader(401)
		w.Write([]byte(err.Error()))
		return
	}
	var row MyUrl
	rows.One(&row)
	if row == (MyUrl{}) {
		hd := hashids.NewData()
		h, _ := hashids.NewWithData(hd)
		now := time.Now()
		url.ID, _ = h.Encode([]int{int(now.Unix())})
		url.ShortUrl = "http://localhost:12345/" + url.ID
		bucket.Insert(url.ID, url, 0)
	} else {
		url = row
	}
	json.NewEncoder(w).Encode(url)
}

func RootEndpoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	var url MyUrl
	bucket.Get(params["id"], &url)
	http.Redirect(w, req, url.LongUrl, 301)
}

func main() {
	fmt.Println("starting app")
	router := mux.NewRouter()
	cluster, _ := gocb.Connect("couchbase://" + os.Getenv("COUCHBASE_HOST"))
	cluster.Authenticate(gocb.PasswordAuthenticator{
		Username: "Administrator",
		Password: "password",
	})
	bucketName = os.Getenv("COUCHBASE_BUCKET")
	bucket, _ = cluster.OpenBucket(bucketName, "")
	router.HandleFunc("/", CreateIndex).Methods("GET")
	router.HandleFunc("/{id}", RootEndpoint).Methods("GET")
	router.HandleFunc("/expand/", ExpandEndpoint).Methods("GET")
	router.HandleFunc("/create", CreateEndpoint).Methods("PUT")
	log.Fatal(http.ListenAndServe(":12345", router))
}
