package tags

import "github.com/gophercloud/gophercloud"

// List all tags on a server.
func List(client *gophercloud.ServiceClient, serverID string) (r ListResult) {
	url := listURL(client, serverID)
	_, r.Err = client.Get(url, &r.Body, &gophercloud.RequestOpts{
		OkCodes: []int{200},
	})
	return
}

// Check if a tag exists on a server.
func Check(client *gophercloud.ServiceClient, serverID, tag string) (r CheckResult) {
	url := checkURL(client, serverID, tag)
	_, r.Err = client.Get(url, nil, &gophercloud.RequestOpts{
		OkCodes: []int{204},
	})
	return
}

// ReplaceAllOptsBuilder allows to add additional parameters to the ReplaceAll request.
type ReplaceAllOptsBuilder interface {
	ToTagsReplaceAllMap() (map[string]interface{}, error)
}

// ReplaceAllOpts provides options used to replace Tags on a server.
type ReplaceAllOpts struct {
	Tags []string `json:"tags" required:"true"`
}

// ToTagsReplaceAllMap formats a ReplaceALlOpts into the body of the ReplaceAll request.
func (opts ReplaceAllOpts) ToTagsReplaceAllMap() (map[string]interface{}, error) {
	return gophercloud.BuildRequestBody(opts, "")
}

// ReplaceAll replaces all tags on a server.
func ReplaceAll(client *gophercloud.ServiceClient, serverID string, opts ReplaceAllOptsBuilder) (r ReplaceAllResult) {
	b, err := opts.ToTagsReplaceAllMap()
	url := replaceAllURL(client, serverID)
	if err != nil {
		r.Err = err
		return
	}
	_, r.Err = client.Put(url, &b, &r.Body, &gophercloud.RequestOpts{
		OkCodes: []int{200},
	})
	return
}
