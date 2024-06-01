package main

import (
	"log"

	client "github.com/coredgeio/goecsclient"
	"github.com/coredgeio/goecsclient/bucket"
)

func main() {
	c, _ := client.CreateEcsClientWithUserCred("user", "password", "https://127.0.0.1:4443")
	b := bucket.GetEcsBucketClient(c)
	param := &bucket.BucketListParameters{Namespace: "container-registry"}
	resp, err := b.GetList(param)
	if err != nil {
		log.Println("got error", err)
	}

	log.Println(resp)
	for i, b := range resp.Buckets {
		log.Println("bucket", i, ":", *b)
	}
}
