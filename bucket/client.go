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

// provides EcsBucketClient for give handler to EcsClient
func GetEcsBucketClient(apiClient client.EcsClient) BucketClient {
	return &bucketClient{
		apiClient: apiClient,
	}
}
