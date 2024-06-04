package main

import (
	"log"

	client "github.com/coredgeio/goecsclient"
	"github.com/coredgeio/goecsclient/bucket"
)

func ListBuckets(b bucket.BucketClient, namespace string) {
	// list buckets
	param := &bucket.BucketListParameters{Namespace: namespace}
	resp, err := b.GetList(param)
	if err != nil {
		log.Println("got error", err)
	}

	log.Println(resp)
	for i, b := range resp.Buckets {
		log.Println("bucket", i, ":", *b)
	}
}

func main() {
	c, _ := client.CreateEcsClientWithUserCred("user", "password", "https://127.0.0.1:4443")
	b := bucket.GetEcsBucketClient(c)

	bucketName := "conreg-test-domain"
	namespace := "container-registry"
	// create new bucket
	req := &bucket.BucketCreateReq{
		Name:      bucketName,
		Namespace: namespace,
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
	ListBuckets(b, namespace)

	// set bucket quota
	quotaReq := &bucket.BucketQuotaUpdateReq{
		BlockSize:        5,
		NotificationSize: -1,
		Namespace:        namespace,
	}
	err = b.SetQuota(bucketName, quotaReq)
	if err != nil {
		log.Println("got error while setting bucket quota", err)
	} else {
		log.Println("quota successfully updated")
	}

	// list buckets
	ListBuckets(b, namespace)

	// reset bucket quota
	quotaReq.BlockSize = -1
	err = b.SetQuota(bucketName, quotaReq)
	if err != nil {
		log.Println("got error while resetting bucket quota", err)
	} else {
		log.Println("quota successfully removed")
	}

	// list buckets
	ListBuckets(b, namespace)

	// get bucket usage
	usageResp, err := b.GetBillingInfo(bucketName, namespace, "KB")
	if err != nil {
		log.Println("got error while getting bucket billing info", err)
	}
	log.Println("bucket usage in KBs: ", usageResp.TotalSize)

	// deleting bucket
	b.Delete(bucketName, namespace)
}
