package main

import (
	"log"

	client "github.com/coredgeio/goecsclient"
	"github.com/coredgeio/goecsclient/bucket"
)

func main() {
	c, _ := client.CreateEcsClientWithUserCred("user", "password", "https://127.0.0.1:4443")
	b := bucket.GetEcsBucketClient(c)

	// create new bucket
	req := &bucket.BucketCreateReq{
		Name:      "conreg-test-domain",
		Namespace: "container-registry",
		HeadType:  "s3",
		//Vpool: "YOTTAWEKARG01",
		Owner:            "orbiter",
		BlockSize:        25,
		NotificationSize: -1,
	}
	res1, err := b.Create(req)
	if err != nil {
		log.Println("got error while creating bucket", err)
	} else {
		log.Printf("bucket created %v", res1)
	}

	// list buckets
	param := &bucket.BucketListParameters{Namespace: "container-registry"}
	resp, err := b.GetList(param)
	if err != nil {
		log.Println("got error", err)
	}

	log.Println(resp)
	for i, b := range resp.Buckets {
		log.Println("bucket", i, ":", *b)
	}

	b.Delete("conreg-test-domain", "container-registry")
}
