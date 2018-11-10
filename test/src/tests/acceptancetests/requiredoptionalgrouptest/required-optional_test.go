package requiredoptionalgrouptest

import (
	"context"
	"fmt"
	"net/http"
	"testing"
	"tests/acceptancetests/utils"
	. "tests/generated/optionalgroup"

	chk "gopkg.in/check.v1"
)

func Test(t *testing.T) { chk.TestingT(t) }

type RequiredOptionalSuite struct{}

var _ = chk.Suite(&RequiredOptionalSuite{})

var explicitClient = getRequiredExplicitTestClient()
var implicitClient = getRequiredImplicitTestClient()

func getRequiredExplicitTestClient() ExplicitClient {
	c := NewExplicitClient("", "", nil)
	c.RetryDuration = 1
	c.BaseURI = utils.GetBaseURI()
	return c
}

func getRequiredImplicitTestClient() ImplicitClient {
	c := NewImplicitClient("", "", nil)
	c.RetryDuration = 1
	c.BaseURI = utils.GetBaseURI()
	return c
}

//Explicit tests

func (s *RequiredOptionalSuite) TestPostRequiredArrayHeader(c *chk.C) {
	_, err := explicitClient.PostRequiredArrayHeader(context.Background(), nil)
	c.Assert(err, chk.NotNil)
	expected := fmt.Errorf("autorest/validation: validation failed: parameter=%s constraint=%s value=%#v details: %s",
		"headerParameter", "Null", []string(nil), "value can not be null; required parameter")
	c.Assert(err.Error(), chk.Equals,
		fmt.Sprintf("optionalgroup.ExplicitClient#PostRequiredArrayHeader: Invalid input: %v", expected))
}

func (s *RequiredOptionalSuite) TestPostRequiredArrayParameter(c *chk.C) {
	_, err := explicitClient.PostRequiredArrayParameter(context.Background(), nil)
	c.Assert(err, chk.NotNil)
	expected := fmt.Errorf("autorest/validation: validation failed: parameter=%s constraint=%s value=%#v details: %s",
		"bodyParameter", "Null", []string(nil), "value can not be null; required parameter")
	c.Assert(err.Error(), chk.Equals,
		fmt.Sprintf("optionalgroup.ExplicitClient#PostRequiredArrayParameter: Invalid input: %v", expected))
}

func (s *RequiredOptionalSuite) TestPostRequiredArrayProperty(c *chk.C) {
	_, err := explicitClient.PostRequiredArrayProperty(context.Background(), ArrayWrapper{})
	c.Assert(err, chk.NotNil)
	expected := fmt.Errorf("autorest/validation: validation failed: parameter=%s constraint=%s value=%#v details: %s",
		"bodyParameter.Value", "Null", (*[]string)(nil), "value can not be null; required parameter")
	c.Assert(err.Error(), chk.Equals,
		fmt.Sprintf("optionalgroup.ExplicitClient#PostRequiredArrayProperty: Invalid input: %v", expected))
}

func (s *RequiredOptionalSuite) TestPostRequiredClassParameter(c *chk.C) {
	_, err := explicitClient.PostRequiredClassParameter(context.Background(), Product{})
	c.Assert(err, chk.NotNil)
	expected := fmt.Errorf("autorest/validation: validation failed: parameter=%s constraint=%s value=%#v details: %s",
		"bodyParameter.ID", "Null", (*int32)(nil), "value can not be null; required parameter")
	c.Assert(err.Error(), chk.Equals,
		fmt.Sprintf("optionalgroup.ExplicitClient#PostRequiredClassParameter: Invalid input: %v", expected))
}

func (s *RequiredOptionalSuite) TestPostRequiredClassProperty(c *chk.C) {
	_, err := explicitClient.PostRequiredClassProperty(context.Background(), ClassWrapper{})
	c.Assert(err, chk.NotNil)
	expected := fmt.Errorf("autorest/validation: validation failed: parameter=%s constraint=%s value=%#v details: %s",
		"bodyParameter.Value", "Null", (*Product)(nil), "value can not be null; required parameter")
	c.Assert(err.Error(), chk.Equals,
		fmt.Sprintf("optionalgroup.ExplicitClient#PostRequiredClassProperty: Invalid input: %v", expected))
}

func (s *RequiredOptionalSuite) TestPostRequiredIntegerProperty(c *chk.C) {
	_, err := explicitClient.PostRequiredIntegerProperty(context.Background(), IntWrapper{})
	c.Assert(err, chk.NotNil)
	expected := fmt.Errorf("autorest/validation: validation failed: parameter=%s constraint=%s value=%#v details: %s",
		"bodyParameter.Value", "Null", (*int32)(nil), "value can not be null; required parameter")
	c.Assert(err.Error(), chk.Equals,
		fmt.Sprintf("optionalgroup.ExplicitClient#PostRequiredIntegerProperty: Invalid input: %v", expected))
}

// Integer can't be null
// func (s *RequiredOptionalSuite) TestPostRequiredIntegerParameter(c *chk.C) {
// 	_, err := explicitClient.PostRequiredIntegerParameter(context.Background(), nil)
// 	c.Assert(err, chk.NotNil)
// }

// func (s *RequiredOptionalSuite) TestPostRequiredIntegerHeader(c *chk.C) {
// 	_, err := explicitClient.PostRequiredIntegerHeader(context.Background(), nil)
// 	c.Assert(err, chk.NotNil)
// }

func (s *RequiredOptionalSuite) TestPostOptionalArrayHeader(c *chk.C) {
	_, err := explicitClient.PostOptionalArrayHeader(context.Background(), nil)
	c.Assert(err, chk.IsNil)
}

func (s *RequiredOptionalSuite) TestPostOptionalArrayParameter(c *chk.C) {
	_, err := explicitClient.PostOptionalArrayParameter(context.Background(), nil)
	c.Assert(err, chk.IsNil)
}

func (s *RequiredOptionalSuite) TestPostOptionalArrayProperty(c *chk.C) {
	_, err := explicitClient.PostOptionalArrayProperty(context.Background(), &ArrayOptionalWrapper{nil})
	c.Assert(err, chk.IsNil)
}

func (s *RequiredOptionalSuite) TestPostOptionalClassParameter(c *chk.C) {
	_, err := explicitClient.PostOptionalClassParameter(context.Background(), nil)
	c.Assert(err, chk.IsNil)
}

func (s *RequiredOptionalSuite) TestPostOptionalClassProperty(c *chk.C) {
	_, err := explicitClient.PostOptionalClassProperty(context.Background(), nil)
	c.Assert(err, chk.IsNil)
}

func (s *RequiredOptionalSuite) TestPostOptionalIntegerHeader(c *chk.C) {
	_, err := explicitClient.PostOptionalIntegerHeader(context.Background(), nil)
	c.Assert(err, chk.IsNil)
}

func (s *RequiredOptionalSuite) TestPostOptionalIntegerParameter(c *chk.C) {
	_, err := explicitClient.PostOptionalIntegerParameter(context.Background(), nil)
	c.Assert(err, chk.IsNil)
}

func (s *RequiredOptionalSuite) TestPostOptionalIntegerProperty(c *chk.C) {
	_, err := explicitClient.PostOptionalIntegerProperty(context.Background(), nil)
	c.Assert(err, chk.IsNil)
}

func (s *RequiredOptionalSuite) TestPostOptionalStringHeader(c *chk.C) {
	_, err := explicitClient.PostOptionalStringHeader(context.Background(), "")
	c.Assert(err, chk.IsNil)
}

func (s *RequiredOptionalSuite) TestPostOptionalStringParameter(c *chk.C) {
	_, err := explicitClient.PostOptionalStringParameter(context.Background(), "")
	c.Assert(err, chk.IsNil)
}

func (s *RequiredOptionalSuite) TestPostOptionalStringProperty(c *chk.C) {
	_, err := explicitClient.PostOptionalStringProperty(context.Background(), nil)
	c.Assert(err, chk.IsNil)
}

//Implicit tests

// GlobalPath string parameter can't be set to null so test invalid for go.
// func (s *RequiredOptionalSuite) TestGetRequiredGlobalPath(c *chk.C) {
// 	_, err := implicitClient.GetRequiredGlobalPath(context.Background())
// 	c.Assert(err, chk.NotNil)
// }

// GlobalQuery string parameter can't be set to null so test invalid for go.
// func (s *RequiredOptionalSuite) TestGetRequiredGlobalQuery(c *chk.C) {
// 	_, err := implicitClient.GetRequiredGlobalQuery(context.Background())
// 	c.Assert(err, chk.NotNil)
// }

// String parameter can't be set to null so test invalid for go.
// func (s *RequiredOptionalSuite) TestGetRequiredPath(c *chk.C) {
// 	_, err := implicitClient.GetRequiredPath(context.Background(), nil) // compile time error
// 	c.Assert(err, chk.NotNil)
// }

func (s *RequiredOptionalSuite) TestPutOptionalBody(c *chk.C) {
	res, err := implicitClient.PutOptionalBody(context.Background(), "")
	c.Assert(err, chk.IsNil)
	c.Assert(res.StatusCode, chk.Equals, http.StatusOK)
}

func (s *RequiredOptionalSuite) TestPutOptionalHeader(c *chk.C) {
	res, err := implicitClient.PutOptionalHeader(context.Background(), "")
	c.Assert(err, chk.IsNil)
	c.Assert(res.StatusCode, chk.Equals, http.StatusOK)
}

func (s *RequiredOptionalSuite) TestPutOptionalQuery(c *chk.C) {
	res, err := implicitClient.PutOptionalQuery(context.Background(), "")
	c.Assert(err, chk.IsNil)
	c.Assert(res.StatusCode, chk.Equals, http.StatusOK)
}

func (s *RequiredOptionalSuite) TestGetOptionalGlobalQuery(c *chk.C) {
	res, err := implicitClient.GetOptionalGlobalQuery(context.Background())
	c.Assert(err, chk.IsNil)
	c.Assert(res.StatusCode, chk.Equals, http.StatusOK)
}
