package testing

import (
	"fmt"
	"net/http"
	"testing"

	fake "github.com/gophercloud/gophercloud/openstack/networking/v2/common"
	"github.com/gophercloud/gophercloud/openstack/networking/v2/extensions/vlantransparent"
	"github.com/gophercloud/gophercloud/openstack/networking/v2/networks"
	th "github.com/gophercloud/gophercloud/testhelper"
)

func TestList(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/v2.0/networks", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		fmt.Fprintf(w, NetworksVLANTransparentListResult)
	})

	type networkVLANTransparentExt struct {
		networks.Network
		vlantransparent.TransparentExt
	}
	var actual []networkVLANTransparentExt

	allPages, err := networks.List(fake.ServiceClient(), networks.ListOpts{}).AllPages()
	th.AssertNoErr(t, err)

	err = networks.ExtractNetworksInto(allPages, &actual)
	th.AssertNoErr(t, err)

	th.AssertEquals(t, "db193ab3-96e3-4cb3-8fc5-05f4296d0324", actual[0].ID)
	th.AssertEquals(t, "private", actual[0].Name)
	th.AssertEquals(t, true, actual[0].AdminStateUp)
	th.AssertEquals(t, "ACTIVE", actual[0].Status)
	th.AssertDeepEquals(t, []string{"08eae331-0402-425a-923c-34f7cfe39c1b"}, actual[0].Subnets)
	th.AssertEquals(t, "26a7980765d0414dbc1fc1f88cdb7e6e", actual[0].TenantID)
	th.AssertEquals(t, false, actual[0].Shared)
	th.AssertEquals(t, true, actual[0].VLANTransparent)
}

func TestGet(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/v2.0/networks/db193ab3-96e3-4cb3-8fc5-05f4296d0324", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		fmt.Fprintf(w, NetworksVLANTransparentGetResult)
	})

	var s struct {
		networks.Network
		vlantransparent.TransparentExt
	}

	err := networks.Get(fake.ServiceClient(), "db193ab3-96e3-4cb3-8fc5-05f4296d0324").ExtractInto(&s)
	th.AssertNoErr(t, err)

	th.AssertEquals(t, "db193ab3-96e3-4cb3-8fc5-05f4296d0324", s.ID)
	th.AssertEquals(t, "private", s.Name)
	th.AssertEquals(t, true, s.AdminStateUp)
	th.AssertEquals(t, "ACTIVE", s.Status)
	th.AssertDeepEquals(t, []string{"08eae331-0402-425a-923c-34f7cfe39c1b"}, s.Subnets)
	th.AssertEquals(t, "26a7980765d0414dbc1fc1f88cdb7e6e", s.TenantID)
	th.AssertEquals(t, false, s.Shared)
	th.AssertEquals(t, true, s.VLANTransparent)
}
