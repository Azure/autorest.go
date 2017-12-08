package numbergrouptest

import (
	"context"
	"testing"

	_ "github.com/shopspring/decimal"
	chk "gopkg.in/check.v1"

	"tests/acceptancetests/utils"
	. "tests/generated/body-number"
)

// TODO: decimal tests

func Test(t *testing.T) { chk.TestingT(t) }

type NumberSuite struct{}

var _ = chk.Suite(&NumberSuite{})

var numberClient = getNumberClient()

func getNumberClient() NumberClient {
	c := NewNumberClient()
	c.RetryDuration = 1
	c.BaseURI = utils.GetBaseURI()
	return c
}

func (s *NumberSuite) TestGetBigDouble(c *chk.C) {
	res, err := numberClient.GetBigDouble(context.Background())
	c.Assert(err, chk.IsNil)
	c.Assert(*res.Value, chk.Equals, 2.5976931e+101)
}

func (s *NumberSuite) TestGetBigDoubleNegativeDecimal(c *chk.C) {
	res, err := numberClient.GetBigDoubleNegativeDecimal(context.Background())
	c.Assert(err, chk.IsNil)
	c.Assert(*res.Value, chk.Equals, -99999999.99)
}

func (s *NumberSuite) TestGetBigDoublePositiveDecimal(c *chk.C) {
	res, err := numberClient.GetBigDoublePositiveDecimal(context.Background())
	c.Assert(err, chk.IsNil)
	c.Assert(*res.Value, chk.Equals, 99999999.99)
}

func (s *NumberSuite) TestGetBigFloat(c *chk.C) {
	res, err := numberClient.GetBigFloat(context.Background())
	c.Assert(err, chk.IsNil)
	c.Assert(*res.Value, chk.Equals, 3.402823e+20)
}

func (s *NumberSuite) TestGetInvalidDouble(c *chk.C) {
	_, err := numberClient.GetInvalidDouble(context.Background())
	c.Assert(err, chk.NotNil)
}

func (s *NumberSuite) TestGetInvalidFloat(c *chk.C) {
	_, err := numberClient.GetInvalidFloat(context.Background())
	c.Assert(err, chk.NotNil)
}

func (s *NumberSuite) TestGetNullNumber(c *chk.C) {
	res, err := numberClient.GetNull(context.Background())
	c.Assert(err, chk.IsNil)
	c.Assert(res.Value, chk.IsNil)
}

func (s *NumberSuite) TestGetSmallDouble(c *chk.C) {
	res, err := numberClient.GetSmallDouble(context.Background())
	c.Assert(err, chk.IsNil)
	c.Assert(*res.Value, chk.Equals, 2.5976931e-101)
}

func (s *NumberSuite) TestGetSmallFloat(c *chk.C) {
	res, err := numberClient.GetSmallFloat(context.Background())
	c.Assert(err, chk.IsNil)
	c.Assert(*res.Value, chk.Equals, 3.402823e-20)
}

func (s *NumberSuite) TestPutBigDouble(c *chk.C) {
	_, err := numberClient.PutBigDouble(context.Background(), 2.5976931e+101)
	c.Assert(err, chk.IsNil)
}

func (s *NumberSuite) TestPutBigDoubleNegativeDecimal(c *chk.C) {
	_, err := numberClient.PutBigDoubleNegativeDecimal(context.Background(), -99999999.99)
	c.Assert(err, chk.IsNil)
}

func (s *NumberSuite) TestPutBigDoublePositiveDecimal(c *chk.C) {
	_, err := numberClient.PutBigDoublePositiveDecimal(context.Background(), 99999999.99)
	c.Assert(err, chk.IsNil)
}

func (s *NumberSuite) TestPutBigFloat(c *chk.C) {
	_, err := numberClient.PutBigFloat(context.Background(), 3.402823e+20)
	c.Assert(err, chk.IsNil)
}

func (s *NumberSuite) TestPutSmallDouble(c *chk.C) {
	_, err := numberClient.PutSmallDouble(context.Background(), 2.5976931e-101)
	c.Assert(err, chk.IsNil)
}

func (s *NumberSuite) TestPutSmallFloat(c *chk.C) {
	_, err := numberClient.PutSmallFloat(context.Background(), 3.402823e-20)
	c.Assert(err, chk.IsNil)
}

// func (s *NumberSuite) TestGetBigDecimal(c *chk.C) {
// 	res, err := numberClient.GetBigDecimal(context.Background())
// 	c.Assert(err, chk.IsNil)
// 	fmt.Println(*res.Value)
// 	fmt.Println(decimal.NewFromFloatWithExponent(2.5976931, 101))
// 	//c.Assert(*(res.Value).String(), chk.DeepEquals, decimal.NewFromFloatWithExponent(2.5976931, 101).String())
// }

// func (s *NumberSuite) TestGetSmallDecimal(c *chk.C) {
// 	res, err := numberClient.GetSmallDecimal(context.Background())
// 	c.Assert(err, chk.IsNil)
// 	c.Assert(*res.Value, chk.DeepEquals, decimal.NewFromFloatWithExponent(2.5976931, -101))
// }

// func (s *NumberSuite) TestGetBigDecimalPositiveDecimal(c *chk.C) {
// 	res, err := numberClient.GetBigDecimalPositiveDecimal(context.Background())
// 	c.Assert(err, chk.IsNil)
// 	c.Assert(*res.Value, chk.DeepEquals, decimal.NewFromFloat(99999999.99))
// }

// func (s *NumberSuite) TestGetBigDecimalNegativeDecimal(c *chk.C) {
// 	res, err := numberClient.GetBigDecimalNegativeDecimal(context.Background())
// 	c.Assert(err, chk.IsNil)
// 	c.Assert(*res.Value, chk.DeepEquals, decimal.NewFromFloat(-99999999.99))
// }

// func (s *NumberSuite) TestPutBigDecimal(c *chk.C) {
// 	_, err := numberClient.PutBigDecimal(context.Background(), decimal.NewFromFloatWithExponent(2.5976931, 101))
// 	c.Assert(err, chk.IsNil)
// }

// func (s *NumberSuite) TestPutSmallDecimal(c *chk.C) {
// 	_, err := numberClient.PutSmallDecimal(context.Background(), decimal.NewFromFloatWithExponent(2.5976931, -101))
// 	c.Assert(err, chk.IsNil)
// }

// func (s *NumberSuite) TestPutBigDecimalPositiveDecimal(c *chk.C) {
// 	_, err := numberClient.PutBigDecimalPositiveDecimal(context.Background(), decimal.NewFromFloat(99999999.99))
// 	c.Assert(err, chk.IsNil)
// }

// func (s *NumberSuite) TestPutBigDecimalNegativeDecimal(c *chk.C) {
// 	_, err := numberClient.PutBigDecimalNegativeDecimal(context.Background(), decimal.NewFromFloat(context.Background()-99999999.99))
// 	c.Assert(err, chk.IsNil)
// }
