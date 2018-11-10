package urlgrouptest

import (
	"context"
	"encoding/base64"
	"testing"
	"tests/acceptancetests/utils"
	. "tests/generated/urlgroup"

	chk "gopkg.in/check.v1"
)

//Not in coverage, for now
//So swagger files are not changed, code for this tests won't be generated
//TestPathBase64URL
//TestPathStringUnicode
//TestPathGetUnixTimeUrl

func Test(t *testing.T) { chk.TestingT(t) }

type URLSuite struct{}

var _ = chk.Suite(&URLSuite{})

var pathClient = getPathClient()
var queryClient = getQueryClient()
var pathItemClient = getPathItemsClient()

func getPathItemsClient() PathItemsClient {
	c := NewPathItemsClient("globalStringPath", "globalStringQuery")
	c.RetryDuration = 1
	c.BaseURI = utils.GetBaseURI()
	return c
}

func getQueryClient() QueriesClient {
	c := NewQueriesClient("globalStringPath", "globalStringQuery")
	c.RetryDuration = 1
	c.BaseURI = utils.GetBaseURI()
	return c
}

func getPathClient() PathsClient {
	c := NewPathsClient("globalStringPath", "globalStringQuery")
	c.RetryDuration = 1
	c.BaseURI = utils.GetBaseURI()
	return c
}

//path tests

func (s *URLSuite) TestPathGetBooleanFalse(c *chk.C) {
	_, err := pathClient.GetBooleanFalse(context.Background())
	c.Assert(err, chk.IsNil)
}

func (s *URLSuite) TestPathGetBooleanTrue(c *chk.C) {
	_, err := pathClient.GetBooleanTrue(context.Background())
	c.Assert(err, chk.IsNil)
}

// Path parameter can't be empty or null.
func (s *URLSuite) TestPathByteEmpty(c *chk.C) {
	_, err := pathClient.ByteEmpty(context.Background())
	c.Assert(err, chk.IsNil)
}

func (s *URLSuite) TestPathByteMultiByte(c *chk.C) {
	encoded := base64.StdEncoding.EncodeToString([]byte("啊齄丂狛狜隣郎隣兀﨩"))
	_, err := pathClient.ByteMultiByte(context.Background(), []byte(encoded))
	c.Assert(err, chk.IsNil)
}

// func (s *URLSuite) TestPathGetUnixTimeUrl(c *chk.C) {
// 	_, err := pathClient.UnixTimeUrl(time.Date(2016, time.April, 13, 0, 0, 0, 0, time.UTC).Unix())
// 	c.Assert(err, chk.IsNil)
// }

func (s *URLSuite) TestPathDateTimeValid(c *chk.C) {
	_, err := pathClient.DateTimeValid(context.Background())
	c.Assert(err, chk.IsNil)
}

func (s *URLSuite) TestPathDateValid(c *chk.C) {
	_, err := pathClient.DateValid(context.Background())
	c.Assert(err, chk.IsNil)
}

func (s *URLSuite) TestPathDoubleDecimalNegative(c *chk.C) {
	_, err := pathClient.DoubleDecimalNegative(context.Background())
	c.Assert(err, chk.IsNil)
}

func (s *URLSuite) TestPathDoubleDecimalPositive(c *chk.C) {
	_, err := pathClient.DoubleDecimalPositive(context.Background())
	c.Assert(err, chk.IsNil)
}

func (s *URLSuite) TestPathEnumNull(c *chk.C) {
	_, err := pathClient.EnumNull(context.Background(), "")
	c.Assert(err, chk.NotNil)
}

func (s *URLSuite) TestPathEnumValid(c *chk.C) {
	_, err := pathClient.EnumValid(context.Background(), Greencolor)
	c.Assert(err, chk.IsNil)
}

func (s *URLSuite) TestPathFloatScientificNegative(c *chk.C) {
	_, err := pathClient.FloatScientificNegative(context.Background())
	c.Assert(err, chk.IsNil)
}

func (s *URLSuite) TestPathFloatScientificPositive(c *chk.C) {
	_, err := pathClient.FloatScientificPositive(context.Background())
	c.Assert(err, chk.IsNil)
}

func (s *URLSuite) TestPathGetIntNegativeOneMillion(c *chk.C) {
	_, err := pathClient.GetIntNegativeOneMillion(context.Background())
	c.Assert(err, chk.IsNil)
}

func (s *URLSuite) TestPathGetIntOneMillion(c *chk.C) {
	_, err := pathClient.GetIntOneMillion(context.Background())
	c.Assert(err, chk.IsNil)
}

func (s *URLSuite) TestPathGetNegativeTenBillion(c *chk.C) {
	_, err := pathClient.GetNegativeTenBillion(context.Background())
	c.Assert(err, chk.IsNil)
}

func (s *URLSuite) TestPathGetTenBillion(c *chk.C) {
	_, err := pathClient.GetTenBillion(context.Background())
	c.Assert(err, chk.IsNil)
}

// Path parameter can't be empty or null.
func (s *URLSuite) TestPathStringEmpty(c *chk.C) {
	_, err := pathClient.StringEmpty(context.Background())
	c.Assert(err, chk.IsNil)
}

func (s *URLSuite) TestPathStringNull(c *chk.C) {
	_, err := pathClient.StringNull(context.Background(), "")
	c.Assert(err, chk.NotNil)
}

func (s *URLSuite) TestPathStringURLEncoded(c *chk.C) {
	_, err := pathClient.StringURLEncoded(context.Background())
	c.Assert(err, chk.IsNil)
}

// not in coverage for now
// func (s *URLSuite) TestPathStringUnicode(c *chk.C) {
// 	_, err := pathClient.StringUnicode(context.Background(), `啊齄丂狛狜隣郎隣兀﨩`)
// 	c.Assert(err, chk.IsNil)
// }

// func (s *URLSuite) TestPathBase64URL(c *chk.C) {
// 	encoded := base64.URLEncoding.EncodeToString([]byte("lorem"))
// 	_, err := pathClient.Base64URL(context.Background(), encoded)
// 	c.Assert(err, chk.IsNil)
// }

// queries tests

// func (s *URLSuite) TestQueryArrayStringCsvEmpty(c *chk.C) {
// 	_, err := queryClient.ArrayStringCsvEmpty(context.Background(), []string{})
// 	c.Assert(err, chk.IsNil)
// }

func (s *URLSuite) TestQueryArrayStringCsvNull(c *chk.C) {
	_, err := queryClient.ArrayStringCsvNull(context.Background(), nil)
	c.Assert(err, chk.IsNil)
}

func (s *URLSuite) TestQueryArrayStringCsvValid(c *chk.C) {
	_, err := queryClient.ArrayStringCsvValid(context.Background(), []string{"ArrayQuery1", "begin!*'();:@ &=+$,/?#[]end", "", ""})
	c.Assert(err, chk.IsNil)
}

func (s *URLSuite) TestQueryArrayStringPipesValid(c *chk.C) {
	_, err := queryClient.ArrayStringPipesValid(context.Background(), []string{"ArrayQuery1", "begin!*'();:@ &=+$,/?#[]end", "", ""})
	c.Assert(err, chk.IsNil)
}

func (s *URLSuite) TestQueryArrayStringSsvValid(c *chk.C) {
	_, err := queryClient.ArrayStringSsvValid(context.Background(), []string{"ArrayQuery1", "begin!*'();:@ &=+$,/?#[]end", "", ""})
	c.Assert(err, chk.IsNil)
}

func (s *URLSuite) TestQueryArrayStringTsvValid(c *chk.C) {
	_, err := queryClient.ArrayStringTsvValid(context.Background(), []string{"ArrayQuery1", "begin!*'();:@ &=+$,/?#[]end", "", ""})
	c.Assert(err, chk.IsNil)
}

// Query parameter is required so can't be empty or null.
func (s *URLSuite) TestQueryByteEmpty(c *chk.C) {
	_, err := queryClient.ByteEmpty(context.Background())
	c.Assert(err, chk.IsNil)
}

func (s *URLSuite) TestQueryByteMultiByte(c *chk.C) {
	encoded := base64.StdEncoding.EncodeToString([]byte("啊齄丂狛狜隣郎隣兀﨩"))
	_, err := queryClient.ByteMultiByte(context.Background(), []byte(encoded))
	c.Assert(err, chk.IsNil)
}

func (s *URLSuite) TestQueryByteNull(c *chk.C) {
	_, err := queryClient.ByteNull(context.Background(), nil)
	c.Assert(err, chk.IsNil)
}

func (s *URLSuite) TestQueryDateNull(c *chk.C) {
	_, err := queryClient.DateNull(context.Background(), nil)
	c.Assert(err, chk.IsNil)
}

func (s *URLSuite) TestQueryDateTimeNull(c *chk.C) {
	_, err := queryClient.DateTimeNull(context.Background(), nil)
	c.Assert(err, chk.IsNil)
}

// dont why not working
func (s *URLSuite) TestQueryDateTimeValid(c *chk.C) {
	_, err := queryClient.DateTimeValid(context.Background())
	c.Assert(err, chk.IsNil)
}

func (s *URLSuite) TestQueryDateValid(c *chk.C) {
	_, err := queryClient.DateValid(context.Background())
	c.Assert(err, chk.IsNil)
}

func (s *URLSuite) TestQueryDoubleDecimalNegative(c *chk.C) {
	_, err := queryClient.DoubleDecimalNegative(context.Background())
	c.Assert(err, chk.IsNil)
}

func (s *URLSuite) TestQueryDoubleDecimalPositive(c *chk.C) {
	_, err := queryClient.DoubleDecimalPositive(context.Background())
	c.Assert(err, chk.IsNil)
}

func (s *URLSuite) TestQueryDoubleNull(c *chk.C) {
	_, err := queryClient.DoubleNull(context.Background(), nil)
	c.Assert(err, chk.IsNil)
}

func (s *URLSuite) TestQueryEnumNull(c *chk.C) {
	_, err := queryClient.EnumNull(context.Background(), "")
	c.Assert(err, chk.IsNil)
}

func (s *URLSuite) TestQueryEnumValid(c *chk.C) {
	_, err := queryClient.EnumValid(context.Background(), Greencolor)
	c.Assert(err, chk.IsNil)
}

func (s *URLSuite) TestQueryFloatNull(c *chk.C) {
	_, err := queryClient.FloatNull(context.Background(), nil)
	c.Assert(err, chk.IsNil)
}

func (s *URLSuite) TestQueryFloatScientificNegative(c *chk.C) {
	_, err := queryClient.FloatScientificNegative(context.Background())
	c.Assert(err, chk.IsNil)
}

func (s *URLSuite) TestQueryFloatScientificPositive(c *chk.C) {
	_, err := queryClient.FloatScientificPositive(context.Background())
	c.Assert(err, chk.IsNil)
}

func (s *URLSuite) TestQueryGetBooleanFalse(c *chk.C) {
	_, err := queryClient.GetBooleanFalse(context.Background())
	c.Assert(err, chk.IsNil)
}

func (s *URLSuite) TestQueryGetBooleanTrue(c *chk.C) {
	_, err := queryClient.GetBooleanTrue(context.Background())
	c.Assert(err, chk.IsNil)
}

func (s *URLSuite) TestQueryGetBooleanNull(c *chk.C) {
	_, err := queryClient.GetBooleanNull(context.Background(), nil)
	c.Assert(err, chk.IsNil)
}

func (s *URLSuite) TestQueryGetIntNegativeOneMillion(c *chk.C) {
	_, err := queryClient.GetIntNegativeOneMillion(context.Background())
	c.Assert(err, chk.IsNil)
}

func (s *URLSuite) TestQueryGetIntOneMillion(c *chk.C) {
	_, err := queryClient.GetIntOneMillion(context.Background())
	c.Assert(err, chk.IsNil)
}

func (s *URLSuite) TestQueryGetIntNull(c *chk.C) {
	_, err := queryClient.GetIntNull(context.Background(), nil)
	c.Assert(err, chk.IsNil)
}

func (s *URLSuite) TestQueryGetLongNull(c *chk.C) {
	_, err := queryClient.GetLongNull(context.Background(), nil)
	c.Assert(err, chk.IsNil)
}

func (s *URLSuite) TestQueryGetNegativeTenBillion(c *chk.C) {
	_, err := queryClient.GetNegativeTenBillion(context.Background())
	c.Assert(err, chk.IsNil)
}

func (s *URLSuite) TestQueryGetTenBillion(c *chk.C) {
	_, err := queryClient.GetTenBillion(context.Background())
	c.Assert(err, chk.IsNil)
}

// Query parameter is required so can't be empty or null.
func (s *URLSuite) TestQueryStringEmpty(c *chk.C) {
	_, err := queryClient.StringEmpty(context.Background())
	c.Assert(err, chk.IsNil)
}

///Can't send string as nil in Go
func (s *URLSuite) TestQueryStringNull(c *chk.C) {
	_, err := queryClient.StringNull(context.Background(), "")
	c.Assert(err, chk.IsNil)
}

func (s *URLSuite) TestQueryStringURLEncoded(c *chk.C) {
	_, err := queryClient.StringURLEncoded(context.Background())
	c.Assert(err, chk.IsNil)
}

//path items tests

func (s *URLSuite) TestPathItemGetAllWithValues(c *chk.C) {
	pathItemClient.GlobalStringPath = "globalStringPath"
	pathItemClient.GlobalStringQuery = "globalStringQuery"
	_, err := pathItemClient.GetAllWithValues(context.Background(), "localStringPath", "pathItemStringPath", "localStringQuery", "pathItemStringQuery")
	c.Assert(err, chk.IsNil)
}

func (s *URLSuite) TestPathItemGetGlobalAndLocalQueryNull(c *chk.C) {
	pathItemClient.GlobalStringQuery = ""
	_, err := pathItemClient.GetGlobalAndLocalQueryNull(context.Background(), "localStringPath", "pathItemStringPath", "", "pathItemStringQuery")
	c.Assert(err, chk.IsNil)
}

func (s *URLSuite) TestPathItemGetGlobalQueryNull(c *chk.C) {
	pathItemClient.GlobalStringQuery = ""
	_, err := pathItemClient.GetGlobalQueryNull(context.Background(), "localStringPath", "pathItemStringPath", "localStringQuery", "pathItemStringQuery")
	c.Assert(err, chk.IsNil)
}

func (s *URLSuite) TestGetLocalPathItemQueryNull(c *chk.C) {
	pathItemClient.GlobalStringQuery = "globalStringQuery"
	_, err := pathItemClient.GetLocalPathItemQueryNull(context.Background(), "localStringPath", "pathItemStringPath", "", "")
	c.Assert(err, chk.IsNil)
}
