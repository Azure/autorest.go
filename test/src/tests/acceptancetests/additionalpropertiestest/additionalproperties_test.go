package additionalpropertiestest

import (
	"context"
	"encoding/base64"
	"testing"

	"github.com/Azure/go-autorest/autorest/to"
	chk "gopkg.in/check.v1"

	"tests/acceptancetests/utils"
	addlprops "tests/generated/additional-properties"
)

func Test(t *testing.T) { chk.TestingT(t) }

type AdditionalPropertiesSuite struct{}

var _ = chk.Suite(&AdditionalPropertiesSuite{})

func getPetsClient() addlprops.PetsClient {
	c := addlprops.NewPetsClient()
	c.RetryDuration = 1
	c.BaseURI = utils.GetBaseURI()
	return c
}

var petsClient = getPetsClient()

func (s *AdditionalPropertiesSuite) TestCreateAPInProperties(c *chk.C) {
	res, err := petsClient.CreateAPInProperties(context.Background(), addlprops.PetAPInProperties{
		ID:   to.Int32Ptr(4),
		Name: to.StringPtr("Bunny"),
		AdditionalProperties: map[string]*float64{
			"height":   to.Float64Ptr(5.61),
			"weight":   to.Float64Ptr(599),
			"footsize": to.Float64Ptr(11.5),
		},
	})
	c.Assert(err, chk.IsNil)
	c.Assert(res.AdditionalProperties, chk.HasLen, 3)
}

func (s *AdditionalPropertiesSuite) TestCreateAPInPropertiesWithAPString(c *chk.C) {
	res, err := petsClient.CreateAPInPropertiesWithAPString(context.Background(), addlprops.PetAPInPropertiesWithAPString{
		ID:            to.Int32Ptr(5),
		Name:          to.StringPtr("Funny"),
		OdataLocation: to.StringPtr("westus"),
		AdditionalProperties: map[string]*string{
			"color": to.StringPtr("red"),
			"city":  to.StringPtr("Seattle"),
			"food":  to.StringPtr("tikka masala"),
		},
		AdditionalProperties1: map[string]*float64{
			"height":   to.Float64Ptr(5.61),
			"weight":   to.Float64Ptr(599),
			"footsize": to.Float64Ptr(11.5),
		},
	})
	c.Assert(err, chk.IsNil)
	c.Assert(res.AdditionalProperties, chk.HasLen, 3)
	c.Assert(*res.AdditionalProperties["color"], chk.Equals, "red")
	c.Assert(*res.AdditionalProperties["city"], chk.Equals, "Seattle")
	c.Assert(*res.AdditionalProperties["food"], chk.Equals, "tikka masala")
	c.Assert(res.AdditionalProperties1, chk.HasLen, 3)
}

func (s *AdditionalPropertiesSuite) TestCreateAPObject(c *chk.C) {
	addlProps := map[string]interface{}{
		"siblings": []interface{}{
			map[string]interface{}{
				"id":        float64(1),
				"name":      "Puppy",
				"birthdate": "2017-12-13T02:29:51Z",
				"complexProperty": map[string]interface{}{
					"color": "Red",
				},
			},
		},
		"picture": base64.StdEncoding.EncodeToString([]byte{255, 255, 255, 255, 254}),
	}
	res, err := petsClient.CreateAPObject(context.Background(), addlprops.PetAPObject{
		ID:                   to.Int32Ptr(2),
		Name:                 to.StringPtr("Hira"),
		AdditionalProperties: addlProps,
	})
	c.Assert(err, chk.IsNil)
	c.Assert(res.AdditionalProperties, chk.HasLen, 2)
	c.Assert(res.AdditionalProperties, chk.DeepEquals, addlProps)
}

func (s *AdditionalPropertiesSuite) TestCreateAPString(c *chk.C) {
	res, err := petsClient.CreateAPString(context.Background(), addlprops.PetAPString{
		ID:   to.Int32Ptr(3),
		Name: to.StringPtr("Tommy"),
		AdditionalProperties: map[string]*string{
			"color":  to.StringPtr("red"),
			"weight": to.StringPtr("10 kg"),
			"city":   to.StringPtr("Bombay"),
		},
	})
	c.Assert(err, chk.IsNil)
	c.Assert(res.AdditionalProperties, chk.HasLen, 3)
	c.Assert(*res.AdditionalProperties["color"], chk.Equals, "red")
	c.Assert(*res.AdditionalProperties["weight"], chk.Equals, "10 kg")
	c.Assert(*res.AdditionalProperties["city"], chk.Equals, "Bombay")
}

func (s *AdditionalPropertiesSuite) TestCreateAPTrue(c *chk.C) {
	addlProps := map[string]interface{}{
		"birthdate": "2017-12-13T02:29:51Z",
		"complexProperty": map[string]interface{}{
			"color": "Red",
		},
	}
	res, err := petsClient.CreateAPTrue(context.Background(), addlprops.PetAPTrue{
		ID:                   to.Int32Ptr(1),
		Name:                 to.StringPtr("Puppy"),
		AdditionalProperties: addlProps,
	})
	c.Assert(err, chk.IsNil)
	c.Assert(res.AdditionalProperties, chk.HasLen, 2)
	c.Assert(res.AdditionalProperties, chk.DeepEquals, addlProps)
}

func (s *AdditionalPropertiesSuite) TestCreateCatAPTrue(c *chk.C) {
	addlProps := map[string]interface{}{
		"birthdate": "2017-12-13T02:29:51Z",
		"complexProperty": map[string]interface{}{
			"color": "Red",
		},
	}
	res, err := petsClient.CreateCatAPTrue(context.Background(), addlprops.CatAPTrue{
		ID:                   to.Int32Ptr(1),
		Name:                 to.StringPtr("Lisa"),
		Friendly:             to.BoolPtr(true),
		AdditionalProperties: addlProps,
	})
	c.Assert(err, chk.IsNil)
	c.Assert(res.AdditionalProperties, chk.HasLen, 2)
	c.Assert(res.AdditionalProperties, chk.DeepEquals, addlProps)
}
