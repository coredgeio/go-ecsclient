package bucket

import (
	"encoding/json"
	"log"
	"net/url"
	"strconv"

	client "github.com/coredgeio/goecsclient"
)

type BucketClient interface {
	GetList(param *BucketListParameters) (*BucketListResp, error)
	Create(req *BucketCreateReq) (*BucketCreateResp, error)
	Delete(name, namespace string) error
}

type bucketClient struct {
	apiClient client.EcsClient
}

func (c *bucketClient) GetList(param *BucketListParameters) (*BucketListResp, error) {
	var query url.Values
	if param != nil && (param.Namespace != "" || param.Marker != "" || param.Limit != 0 || param.Name != "") {
		query = url.Values{}
		if param.Namespace != "" {
			query.Add("namespace", param.Namespace)
		}
		if param.Marker != "" {
			query.Add("marker", param.Marker)
		}
		if param.Limit != 0 {
			query.Add("limit", strconv.Itoa(param.Limit))
		}
		if param.Name != "" {
			query.Add("name", param.Name)
		}
	}

	bytes, err := c.apiClient.Get("/object/bucket", query)
	if err != nil {
		return nil, err
	}

	resp := &BucketListResp{}
	if err = json.Unmarshal(bytes, resp); err != nil {
		log.Println("failed to decode response for get bucket list", err)
	}
	return resp, err
}

func (c *bucketClient) Create(req *BucketCreateReq) (*BucketCreateResp, error) {
	data, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	log.Println("sending data", string(data))

	bytes, err := c.apiClient.Post("/object/bucket", data, nil)
	if err != nil {
		return nil, err
	}

	resp := &BucketCreateResp{}
	if err = json.Unmarshal(bytes, resp); err != nil {
		log.Println("failed to decode response for create bucket", err)
	}
	return resp, err
}

func (c *bucketClient) Delete(name, namespace string) error {
	var query url.Values
	if namespace != "" {
		query = url.Values{}
		query.Add("namespace", namespace)
	}
	bytes, err := c.apiClient.Post("/object/bucket/"+name+"/deactivate", nil, query)
	if err != nil {
		log.Println("got error while deleting bucket", err)
		return err
	}

	log.Println("got response for delete:", bytes)
	return nil
}

// provides EcsBucketClient for give handler to EcsClient
func GetEcsBucketClient(apiClient client.EcsClient) BucketClient {
	return &bucketClient{
		apiClient: apiClient,
	}
}
