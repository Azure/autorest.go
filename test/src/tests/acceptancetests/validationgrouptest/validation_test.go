package validationgrouptest

import (
	"context"
	"testing"
	"tests/acceptancetests/utils"
	. "tests/generated/validationgroup"

	chk "gopkg.in/check.v1"
)

func Test(t *testing.T) { chk.TestingT(t) }

type ValidationSuite struct{}

var _ = chk.Suite(&ValidationSuite{})

var validationClient = getValidationClient()

func getValidationClient() BaseClient {
	c := New("abc123")
	c.RetryDuration = 1
	c.BaseURI = utils.GetBaseURI()
	return c
}

func (s *ValidationSuite) TestGetWithConstantInPath(c *chk.C) {
	_, err := validationClient.GetWithConstantInPath(context.Background())
	c.Assert(err, chk.IsNil)
}

func (s *ValidationSuite) TestPostWithConstantInBody(c *chk.C) {
	cInt, cString, constProperty2 := int32(0), "constant", "constant2"
	p := Product{
		ConstInt:    &cInt,
		ConstString: &cString,
		Child: &ChildProduct{
			ConstProperty: &cString,
		},
		ConstChild: &ConstantProduct{
			ConstProperty:  &cString,
			ConstProperty2: &constProperty2,
		},
	}
	res, err := validationClient.PostWithConstantInBody(context.Background(), &p)
	p.Response = res.Response
	c.Assert(err, chk.IsNil)
	c.Assert(res, chk.DeepEquals, p)
}

func (s *ValidationSuite) TestValidationOfBody(c *chk.C) {
	_, err := validationClient.ValidationOfBody(context.Background(), "123", 150, &Product{
		DisplayNames: &[]string{
			"displayname1",
			"displayname2",
			"displayname3",
			"displayname4",
			"displayname5",
			"displayname6",
			"displayname7"},
	})
	c.Assert(err, chk.NotNil)
}
