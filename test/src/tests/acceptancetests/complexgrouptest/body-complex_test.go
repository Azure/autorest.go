package complexgrouptest

import (
	"context"
	"encoding/json"
	"net/http"
	"testing"
	"time"

	"github.com/Azure/go-autorest/autorest/date"
	"github.com/Azure/go-autorest/autorest/to"
	chk "gopkg.in/check.v1"

	"tests/acceptancetests/utils"
	. "tests/generated/body-complex"
)

func Test(t *testing.T) { chk.TestingT(t) }

type ComplexGroupSuite struct{}

var _ = chk.Suite(&ComplexGroupSuite{})

var complexPrimitiveClient = getPrimitiveClient()
var complexArrayClient = getArrayComplexClient()
var complexDictionaryClient = getDictionaryComplexClient()
var complexBasicOperationsClient = getBasicOperationsClient()
var complexInheritanceClient = getInheritanceClient()
var complexReadOnlyClient = getReadOnlyClient()
var complexPolymorphicClient = getPolymorphismClient()
var complexPolymorphicRecursiveClient = getPolymorphismRecursiveClient()

func getArrayComplexClient() ArrayClient {
	c := NewArrayClient()
	c.RetryDuration = 1
	c.BaseURI = utils.GetBaseURI()
	return c
}

func getBasicOperationsClient() BasicClient {
	c := NewBasicClient()
	c.RetryDuration = 1
	c.BaseURI = utils.GetBaseURI()
	return c
}

func getDictionaryComplexClient() DictionaryClient {
	c := NewDictionaryClient()
	c.RetryDuration = 1
	c.BaseURI = utils.GetBaseURI()
	return c
}

func getPrimitiveClient() PrimitiveClient {
	c := NewPrimitiveClient()
	c.RetryDuration = 1
	c.BaseURI = utils.GetBaseURI()
	return c
}

func getInheritanceClient() InheritanceClient {
	c := NewInheritanceClient()
	c.RetryDuration = 1
	c.BaseURI = utils.GetBaseURI()
	return c
}

func getPolymorphismClient() PolymorphismClient {
	c := NewPolymorphismClient()
	c.RetryDuration = 1
	c.BaseURI = utils.GetBaseURI()
	return c
}

func getReadOnlyClient() ReadonlypropertyClient {
	c := NewReadonlypropertyClient()
	c.RetryDuration = 1
	c.BaseURI = utils.GetBaseURI()
	return c
}

func getPolymorphismRecursiveClient() PolymorphicrecursiveClient {
	c := NewPolymorphicrecursiveClient()
	c.RetryDuration = 1
	c.BaseURI = utils.GetBaseURI()
	return c
}

// Primitive tests
func (s *ComplexGroupSuite) TestGetBoolComplex(c *chk.C) {
	res, err := complexPrimitiveClient.GetBool(context.Background())
	c.Assert(err, chk.IsNil)

	c.Assert(*res.FieldTrue, chk.Equals, true)
	c.Assert(*res.FieldFalse, chk.Equals, false)
}

func (s *ComplexGroupSuite) TestGetByteComplex(c *chk.C) {
	res, err := complexPrimitiveClient.GetByte(context.Background())
	c.Assert(err, chk.IsNil)
	c.Assert(*res.Field, chk.DeepEquals, []byte{0x0FF, 0x0FE, 0x0FD, 0x0FC, 0x000, 0x0FA, 0x0F9, 0x0F8, 0x0F7, 0x0F6})
}

func (s *ComplexGroupSuite) TestGetDateComplex(c *chk.C) {
	res, err := complexPrimitiveClient.GetDate(context.Background())
	c.Assert(err, chk.IsNil)

	c.Assert(*res.Field, chk.DeepEquals, date.Date{time.Date(0001, time.January, 1, 0, 0, 0, 0, time.UTC)})
	c.Assert(*res.Leap, chk.DeepEquals, date.Date{time.Date(2016, time.February, 29, 0, 0, 0, 0, time.UTC)})
}

func (s *ComplexGroupSuite) TestGetDateTimeComplex(c *chk.C) {
	res, err := complexPrimitiveClient.GetDateTime(context.Background())
	c.Assert(err, chk.IsNil)

	c.Assert(*res.Field, chk.DeepEquals, utils.ToDateTime("0001-01-01T00:00:00Z"))
	c.Assert(*res.Now, chk.DeepEquals, utils.ToDateTime("2015-05-18T18:38:00Z"))
}

func (s *ComplexGroupSuite) TestGetDateTimeRfc1123Complex(c *chk.C) {
	res, err := complexPrimitiveClient.GetDateTimeRfc1123(context.Background())
	c.Assert(err, chk.IsNil)

	c.Assert(*res.Field, chk.DeepEquals, utils.ToDateTimeRFC1123("Mon, 01 Jan 0001 00:00:00 GMT"))
	c.Assert(*res.Now, chk.DeepEquals, utils.ToDateTimeRFC1123("Mon, 18 May 2015 11:38:00 GMT"))
}

func (s *ComplexGroupSuite) TestGetDoubleComplex(c *chk.C) {
	res, err := complexPrimitiveClient.GetDouble(context.Background())
	c.Assert(err, chk.IsNil)

	c.Assert(*res.Field1, chk.Equals, 3e-100)
	c.Assert(*res.Field56ZerosAfterTheDotAndNegativeZeroBeforeDotAndThisIsALongFieldNameOnPurpose, chk.Equals, -0.000000000000000000000000000000000000000000000000000000005)
}

func (s *ComplexGroupSuite) TestGetDurationComplex(c *chk.C) {
	res, err := complexPrimitiveClient.GetDuration(context.Background())
	c.Assert(err, chk.IsNil)
	c.Assert(*res.Field, chk.Equals, "P123DT22H14M12.011S")
}

func (s *ComplexGroupSuite) TestGetFloatComplex(c *chk.C) {
	res, err := complexPrimitiveClient.GetFloat(context.Background())
	c.Assert(err, chk.IsNil)

	c.Assert(*res.Field1, chk.Equals, 1.05)
	c.Assert(*res.Field2, chk.Equals, -0.003)
}

func (s *ComplexGroupSuite) TestGetIntComplex(c *chk.C) {
	res, err := complexPrimitiveClient.GetInt(context.Background())
	c.Assert(err, chk.IsNil)
	c.Assert(*res.Field1, chk.Equals, int32(-1))
	c.Assert(*res.Field2, chk.Equals, int32(2))
}

func (s *ComplexGroupSuite) TestGetLongComplex(c *chk.C) {
	res, err := complexPrimitiveClient.GetLong(context.Background())
	c.Assert(err, chk.IsNil)
	c.Assert(*res.Field1, chk.Equals, int64(1099511627775))
	c.Assert(*res.Field2, chk.Equals, int64(-999511627788))
}

func (s *ComplexGroupSuite) TestGetStringComplex(c *chk.C) {
	res, err := complexPrimitiveClient.GetString(context.Background())
	c.Assert(err, chk.IsNil)

	c.Assert(*res.Field, chk.Equals, "goodrequest")
	c.Assert(*res.Empty, chk.Equals, "")
	c.Assert(res.Null, chk.IsNil)
}

func (s *ComplexGroupSuite) TestPutBoolComplex(c *chk.C) {
	a, b := true, false
	_, err := complexPrimitiveClient.PutBool(context.Background(), BooleanWrapper{FieldTrue: &a, FieldFalse: &b})
	c.Assert(err, chk.IsNil)
}

func (s *ComplexGroupSuite) TestPutByteComplex(c *chk.C) {
	_, err := complexPrimitiveClient.PutByte(context.Background(),
		ByteWrapper{Field: &[]byte{0x0FF, 0x0FE, 0x0FD, 0x0FC, 0x000, 0x0FA, 0x0F9, 0x0F8, 0x0F7, 0x0F6}})
	c.Assert(err, chk.IsNil)
}

func (s *ComplexGroupSuite) TestPutDateComplex(c *chk.C) {
	_, err := complexPrimitiveClient.PutDate(context.Background(),
		DateWrapper{Field: &date.Date{time.Date(0001, time.January, 1, 0, 0, 0, 0, time.UTC)},
			Leap: &date.Date{time.Date(2016, time.February, 29, 0, 0, 0, 0, time.UTC)}})
	c.Assert(err, chk.IsNil)
}

func (s *ComplexGroupSuite) TestPutDateTimeComplex(c *chk.C) {
	dt1, dt2 := utils.ToDateTime("0001-01-01T00:00:00Z"), utils.ToDateTime("2015-05-18T18:38:00Z")
	_, err := complexPrimitiveClient.PutDateTime(context.Background(), DatetimeWrapper{Field: &dt1, Now: &dt2})
	c.Assert(err, chk.IsNil)
}

func (s *ComplexGroupSuite) TestPutDateTimeRFC1123Complex(c *chk.C) {
	dt1, dt2 := utils.ToDateTimeRFC1123("Mon, 01 Jan 0001 00:00:00 GMT"), utils.ToDateTimeRFC1123("Mon, 18 May 2015 11:38:00 GMT")
	_, err := complexPrimitiveClient.PutDateTimeRfc1123(context.Background(), Datetimerfc1123Wrapper{Field: &dt1, Now: &dt2})
	c.Assert(err, chk.IsNil)
}

func (s *ComplexGroupSuite) TestPutDoubleComplex(c *chk.C) {
	d1, d2 := 3e-100, -0.000000000000000000000000000000000000000000000000000000005
	_, err := complexPrimitiveClient.PutDouble(context.Background(), DoubleWrapper{Field1: &d1,
		Field56ZerosAfterTheDotAndNegativeZeroBeforeDotAndThisIsALongFieldNameOnPurpose: &d2})
	c.Assert(err, chk.IsNil)
}

func (s *ComplexGroupSuite) TestPutDurationComplex(c *chk.C) {
	duration := "P123DT22H14M12.011S"
	_, err := complexPrimitiveClient.PutDuration(context.Background(), DurationWrapper{Field: &duration})
	c.Assert(err, chk.IsNil)
}

func (s *ComplexGroupSuite) TestPutFloatComplex(c *chk.C) {
	f1, f2 := 1.05, -0.003
	_, err := complexPrimitiveClient.PutFloat(context.Background(), FloatWrapper{Field1: &f1, Field2: &f2})
	c.Assert(err, chk.IsNil)
}

func (s *ComplexGroupSuite) TestPutIntComplex(c *chk.C) {
	var i1, i2 int32 = -1, 2
	_, err := complexPrimitiveClient.PutInt(context.Background(), IntWrapper{Field1: &i1, Field2: &i2})
	c.Assert(err, chk.IsNil)
}

func (s *ComplexGroupSuite) TestPutLongComplex(c *chk.C) {
	var l1, l2 int64 = 1099511627775, -999511627788
	_, err := complexPrimitiveClient.PutLong(context.Background(), LongWrapper{Field1: &l1, Field2: &l2})
	c.Assert(err, chk.IsNil)
}

func (s *ComplexGroupSuite) TestPutStringComplex(c *chk.C) {
	s1, s2 := "goodrequest", ""
	_, err := complexPrimitiveClient.PutString(context.Background(), StringWrapper{Field: &s1, Empty: &s2, Null: nil})
	c.Assert(err, chk.IsNil)
}

//Array Complex tests
func (s *ComplexGroupSuite) TestGetEmptyArrayComplex(c *chk.C) {
	res, err := complexArrayClient.GetEmpty(context.Background())
	c.Assert(err, chk.IsNil)
	c.Assert(*res.Array, chk.DeepEquals, []string{})
}

func (s *ComplexGroupSuite) TestGetNotProvidedArrayComplex(c *chk.C) {
	res, err := complexArrayClient.GetNotProvided(context.Background())
	c.Assert(err, chk.IsNil)
	c.Assert(res.Array, chk.IsNil)
}

func (s *ComplexGroupSuite) TestGetValidArrayComplex(c *chk.C) {
	// string can't be null
	array := []string{"1, 2, 3, 4", "", "", "&S#$(*Y", "The quick brown fox jumps over the lazy dog"}
	res, err := complexArrayClient.GetValid(context.Background())
	c.Assert(err, chk.IsNil)
	c.Assert(*res.Array, chk.DeepEquals, array)
}

func (s *ComplexGroupSuite) TestPutEmptyArrayComplex(c *chk.C) {
	_, err := complexArrayClient.PutEmpty(context.Background(), ArrayWrapper{Array: &[]string{}})
	c.Assert(err, chk.IsNil)
}

// func (s *ComplexGroupSuite) TestPutComplexArrayValid(c *chk.C) {
// 	// string can't be null
// 	array := []string{"1, 2, 3, 4", "", "", "&S#$(*Y", "The quick brown fox jumps over the lazy dog"}
// 	_, err := complexArray.PutValid(context.Background(), ArrayWrapper{Array: &array})
// 	c.Assert(err, chk.IsNil)
// }

// Dictionary complex tests
func (s *ComplexGroupSuite) TestGetEmptyDictionaryComplex(c *chk.C) {
	res, err := complexDictionaryClient.GetEmpty(context.Background())
	c.Assert(err, chk.IsNil)
	c.Assert(res.DefaultProgram, chk.DeepEquals, map[string]*string{})
}

func (s *ComplexGroupSuite) TestGetNotProvidedDictionaryComplex(c *chk.C) {
	res, err := complexDictionaryClient.GetNotProvided(context.Background())
	c.Assert(err, chk.IsNil)
	c.Assert(res.DefaultProgram, chk.IsNil)
}

func (s *ComplexGroupSuite) TestGetNullDictionaryComplex(c *chk.C) {
	res, err := complexDictionaryClient.GetNull(context.Background())
	c.Assert(err, chk.IsNil)
	c.Assert(res.DefaultProgram, chk.IsNil)
}

func (s *ComplexGroupSuite) TestGetValidDictionaryComplex(c *chk.C) {
	s1, s2, s3, s4 := "notepad", "mspaint", "excel", ""
	dic := map[string]*string{"txt": &s1, "bmp": &s2, "xls": &s3, "exe": &s4, "": nil}
	res, err := complexDictionaryClient.GetValid(context.Background())
	c.Assert(err, chk.IsNil)
	c.Assert(res.DefaultProgram, chk.DeepEquals, dic)
}

func (s *ComplexGroupSuite) TestPutEmptyDictionaryComplex(c *chk.C) {
	_, err := complexDictionaryClient.PutEmpty(context.Background(), DictionaryWrapper{DefaultProgram: map[string]*string{}})
	c.Assert(err, chk.IsNil)
}

func (s *ComplexGroupSuite) TestPutValidDictionaryComplex(c *chk.C) {
	jsonBlob := `{"defaultProgram":{"txt":"notepad","bmp":"mspaint","xls":"excel","exe":"","":null}}`
	var dw DictionaryWrapper
	err := json.Unmarshal([]byte(jsonBlob), &dw)
	c.Assert(err, chk.IsNil)
	_, err = complexDictionaryClient.PutValid(context.Background(), dw)
	c.Assert(err, chk.IsNil)
}

// Basic operations tests
func (s *ComplexGroupSuite) TestGetEmptyBasicOperationsComplex(c *chk.C) {
	res, err := complexBasicOperationsClient.GetEmpty(context.Background())
	c.Assert(err, chk.IsNil)

	c.Assert(res.ID, chk.IsNil)
	c.Assert(res.Name, chk.IsNil)

	c.Assert(string(res.Color), chk.Equals, "")
}

func (s *ComplexGroupSuite) TestGetInvalidBasicOperationsComplex(c *chk.C) {
	_, err := complexBasicOperationsClient.GetInvalid(context.Background())
	c.Assert(err, chk.NotNil)
}

func (s *ComplexGroupSuite) TestGetNotProvidedBasicOperationsComplex(c *chk.C) {
	res, err := complexBasicOperationsClient.GetNotProvided(context.Background())
	c.Assert(err, chk.IsNil)

	c.Assert(res.ID, chk.IsNil)
	c.Assert(res.Name, chk.IsNil)

	c.Assert(string(res.Color), chk.Equals, "")
}

func (s *ComplexGroupSuite) TestGetValidBasicOperationsComplex(c *chk.C) {
	res, err := complexBasicOperationsClient.GetValid(context.Background())
	c.Assert(err, chk.IsNil)

	c.Assert(*res.ID, chk.Equals, int32(2))
	c.Assert(*res.Name, chk.Equals, "abc")

	c.Assert(string(res.Color), chk.Equals, "YELLOW")
}

func (s *ComplexGroupSuite) TestGetNullBasicOperationsComplex(c *chk.C) {
	res, err := complexBasicOperationsClient.GetNull(context.Background())
	c.Assert(err, chk.IsNil)

	c.Assert(res.ID, chk.IsNil)
	c.Assert(res.Name, chk.IsNil)

	c.Assert(string(res.Color), chk.Equals, "")
}

func (s *ComplexGroupSuite) TestPutValidBasicOperationsComplex(c *chk.C) {
	m, n := int32(2), "abc"
	expected := Basic{ID: &m, Name: &n, Color: "Magenta"}
	_, err := complexBasicOperationsClient.PutValid(context.Background(), expected)
	c.Assert(err, chk.IsNil)
}

// Inheritance tests
func (s *ComplexGroupSuite) TestGetValidInheritanceComplex(c *chk.C) {
	res, err := complexInheritanceClient.GetValid(context.Background())
	c.Assert(err, chk.IsNil)
	dogs := []Dog{getDogObject("tomato", 1, "Potato"), getDogObject("french fries", -1, "Tomato")}

	c.Assert(*res.ID, chk.Equals, int32(2))
	c.Assert(*res.Name, chk.Equals, "Siameeee")
	c.Assert(*res.Color, chk.Equals, "green")
	c.Assert(*res.Breed, chk.Equals, "persian")
	c.Assert(*res.Hates, chk.DeepEquals, dogs)
}

func (s *ComplexGroupSuite) TestPutValidInheritanceComplex(c *chk.C) {
	a, b, x, y := int32(2), "Siameeee", "persian", "green"
	cat := Siamese{
		ID:    &a,
		Name:  &b,
		Breed: &x,
		Color: &y,
		Hates: &[]Dog{
			getDogObject("tomato", 1, "Potato"),
			getDogObject("french fries", -1, "Tomato"),
		},
	}
	_, err := complexInheritanceClient.PutValid(context.Background(), cat)
	c.Assert(err, chk.IsNil)
}

func getDogObject(food string, id int32, name string) Dog {
	return Dog{Food: &food, ID: &id, Name: &name}
}

// Polymorphic tests
var f = Salmon{
	Length:   to.Float64Ptr(1),
	Iswild:   to.BoolPtr(true),
	Location: to.StringPtr("alaska"),
	Species:  to.StringPtr("king"),
	Siblings: &[]BasicFish{
		Shark{
			Length:   to.Float64Ptr(20),
			Birthday: &date.Time{time.Date(2012, time.January, 5, 1, 0, 0, 0, time.UTC)},
			Age:      to.Int32Ptr(6),
			Species:  to.StringPtr("predator"),
		},
		Sawshark{
			Length:   to.Float64Ptr(10),
			Birthday: &date.Time{time.Date(1900, time.January, 5, 1, 0, 0, 0, time.UTC)},
			Age:      to.Int32Ptr(105),
			Species:  to.StringPtr("dangerous"),
			Picture:  &[]byte{255, 255, 255, 255, 254},
		},
		Goblinshark{
			Length:   to.Float64Ptr(30),
			Birthday: &date.Time{time.Date(2015, time.August, 8, 0, 0, 0, 0, time.UTC)},
			Age:      to.Int32Ptr(1),
			Species:  to.StringPtr("scary"),
			Jawsize:  to.Int32Ptr(5),
			Color:    GoblinSharkColor("pinkish-gray"),
		},
	},
}

var ss = SmartSalmon{
	Length:   to.Float64Ptr(1),
	Iswild:   to.BoolPtr(true),
	Location: to.StringPtr("alaska"),
	Species:  to.StringPtr("king"),
	AdditionalProperties: map[string]interface{}{
		"additionalProperty1": float64(1),
		"additionalProperty2": false,
		"additionalProperty3": "hello",
		"additionalProperty4": map[string]interface{}{
			"a": float64(1),
			"b": float64(2),
		},
		"additionalProperty5": []interface{}{float64(1), float64(3)},
	},
	Siblings: &[]BasicFish{
		Shark{
			Length:   to.Float64Ptr(20),
			Birthday: &date.Time{time.Date(2012, time.January, 5, 1, 0, 0, 0, time.UTC)},
			Age:      to.Int32Ptr(6),
			Species:  to.StringPtr("predator"),
		},
		Sawshark{
			Length:   to.Float64Ptr(10),
			Birthday: &date.Time{time.Date(1900, time.January, 5, 1, 0, 0, 0, time.UTC)},
			Age:      to.Int32Ptr(105),
			Species:  to.StringPtr("dangerous"),
			Picture:  &[]byte{255, 255, 255, 255, 254},
		},
		Goblinshark{
			Length:   to.Float64Ptr(30),
			Birthday: &date.Time{time.Date(2015, time.August, 8, 0, 0, 0, 0, time.UTC)},
			Age:      to.Int32Ptr(1),
			Species:  to.StringPtr("scary"),
			Jawsize:  to.Int32Ptr(5),
			Color:    "pinkish-gray",
		},
	},
}

func (s *ComplexGroupSuite) TestGetComplexPolymorphicValid(c *chk.C) {
	resp, err := complexPolymorphicClient.GetValid(context.Background())
	c.Assert(err, chk.IsNil)

	salmon, b := resp.Value.AsSalmon()
	c.Assert(*salmon, chk.FitsTypeOf, f)
	c.Assert(b, chk.Equals, true)

	c.Assert(resp.Value, chk.FitsTypeOf, f)
	c.Assert(*salmon.Iswild, chk.Equals, *f.Iswild)
	c.Assert(*salmon.Location, chk.Equals, *f.Location)
	c.Assert(*salmon.Species, chk.Equals, *f.Species)
	c.Assert(*salmon.Length, chk.Equals, *f.Length)
	c.Assert(*salmon.Siblings, chk.HasLen, len(*f.Siblings))

	c.Assert((*salmon.Siblings)[0], chk.FitsTypeOf, (*f.Siblings)[0].(Shark))
	c.Assert(*(*salmon.Siblings)[0].(Shark).Length, chk.Equals, *(*f.Siblings)[0].(Shark).Length)
	c.Assert(*(*salmon.Siblings)[0].(Shark).Birthday, chk.Equals, *(*f.Siblings)[0].(Shark).Birthday)
	c.Assert(*(*salmon.Siblings)[0].(Shark).Age, chk.Equals, *(*f.Siblings)[0].(Shark).Age)
	c.Assert(*(*salmon.Siblings)[0].(Shark).Species, chk.Equals, *(*f.Siblings)[0].(Shark).Species)

	c.Assert((*salmon.Siblings)[1], chk.FitsTypeOf, (*f.Siblings)[1].(Sawshark))
	c.Assert(*(*salmon.Siblings)[1].(Sawshark).Length, chk.Equals, *(*f.Siblings)[1].(Sawshark).Length)
	c.Assert(*(*salmon.Siblings)[1].(Sawshark).Birthday, chk.Equals, *(*f.Siblings)[1].(Sawshark).Birthday)
	c.Assert(*(*salmon.Siblings)[1].(Sawshark).Age, chk.Equals, *(*f.Siblings)[1].(Sawshark).Age)
	c.Assert(*(*salmon.Siblings)[1].(Sawshark).Species, chk.Equals, *(*f.Siblings)[1].(Sawshark).Species)
	c.Assert(*(*salmon.Siblings)[1].(Sawshark).Picture, chk.DeepEquals, *(*f.Siblings)[1].(Sawshark).Picture)

	c.Assert((*salmon.Siblings)[2], chk.FitsTypeOf, (*f.Siblings)[2].(Goblinshark))
	c.Assert(*(*salmon.Siblings)[2].(Goblinshark).Length, chk.Equals, *(*f.Siblings)[2].(Goblinshark).Length)
	c.Assert(*(*salmon.Siblings)[2].(Goblinshark).Birthday, chk.Equals, *(*f.Siblings)[2].(Goblinshark).Birthday)
	c.Assert(*(*salmon.Siblings)[2].(Goblinshark).Age, chk.Equals, *(*f.Siblings)[2].(Goblinshark).Age)
	c.Assert(*(*salmon.Siblings)[2].(Goblinshark).Species, chk.Equals, *(*f.Siblings)[2].(Goblinshark).Species)
	c.Assert(*(*salmon.Siblings)[2].(Goblinshark).Jawsize, chk.Equals, *(*f.Siblings)[2].(Goblinshark).Jawsize)
}

func (s *ComplexGroupSuite) TestPutComplexPolymorphismValid(c *chk.C) {
	_, err := complexPolymorphicClient.PutValid(context.Background(), f)
	c.Assert(err, chk.IsNil)
}

func (s *ComplexGroupSuite) TestPutComplexPolymorphismValidMissingRequired(c *chk.C) {
	badF := Salmon{
		Length:   to.Float64Ptr(1),
		Iswild:   to.BoolPtr(true),
		Location: to.StringPtr("alaska"),
		Species:  to.StringPtr("king"),
		Siblings: &[]BasicFish{
			Shark{
				Length:   to.Float64Ptr(20),
				Birthday: &date.Time{time.Date(2012, time.January, 5, 1, 0, 0, 0, time.UTC)},
				Age:      to.Int32Ptr(6),
				Species:  to.StringPtr("predator"),
			},
			Sawshark{
				Length:  to.Float64Ptr(10),
				Age:     to.Int32Ptr(105),
				Species: to.StringPtr("dangerous"),
				Picture: &[]byte{255, 255, 255, 255, 254},
			},
		},
	}
	resp, err := complexPolymorphicClient.PutValidMissingRequired(context.Background(), badF)
	c.Assert(err, chk.NotNil)
	c.Assert(resp.StatusCode, chk.Equals, http.StatusBadRequest)
}

// Polymorphic recursive tests
func (s *ComplexGroupSuite) TestGetComplexPolymorphicRecursive(c *chk.C) {
	resp, err := complexPolymorphicRecursiveClient.GetValid(context.Background())
	c.Assert(err, chk.IsNil)

	c.Assert(resp.Value, chk.FitsTypeOf, Salmon{})
	c.Assert((*resp.Value.(Salmon).Siblings)[0], chk.FitsTypeOf, Shark{})
	c.Assert((*(*resp.Value.(Salmon).Siblings)[0].(Shark).Siblings)[0], chk.FitsTypeOf, Salmon{})
	c.Assert(*(*(*resp.Value.(Salmon).Siblings)[0].(Shark).Siblings)[0].(Salmon).Location, chk.FitsTypeOf, "atlantic")
}

func (s *ComplexGroupSuite) TestPutComplexPolymorphicRecursive(c *chk.C) {
	recF := Salmon{
		Iswild:   to.BoolPtr(true),
		Length:   to.Float64Ptr(1),
		Species:  to.StringPtr("king"),
		Location: to.StringPtr("alaska"),
		Siblings: &[]BasicFish{
			Shark{
				Age:      to.Int32Ptr(6),
				Length:   to.Float64Ptr(20),
				Species:  to.StringPtr("predator"),
				Birthday: &date.Time{time.Date(2012, time.January, 5, 1, 0, 0, 0, time.UTC)},
				Siblings: &[]BasicFish{
					Salmon{
						Iswild:   to.BoolPtr(true),
						Length:   to.Float64Ptr(2),
						Location: to.StringPtr("atlantic"),
						Species:  to.StringPtr("coho"),
						Siblings: &[]BasicFish{
							Shark{
								Age:      to.Int32Ptr(6),
								Length:   to.Float64Ptr(20),
								Species:  to.StringPtr("predator"),
								Birthday: &date.Time{time.Date(2012, time.January, 5, 1, 0, 0, 0, time.UTC)},
							},
							Sawshark{
								Age:      to.Int32Ptr(105),
								Length:   to.Float64Ptr(10),
								Birthday: &date.Time{time.Date(1900, time.January, 5, 1, 0, 0, 0, time.UTC)},
								Species:  to.StringPtr("dangerous"),
								Picture:  &[]byte{255, 255, 255, 255, 254},
							},
						},
					},
					Sawshark{
						Age:      to.Int32Ptr(105),
						Length:   to.Float64Ptr(10),
						Species:  to.StringPtr("dangerous"),
						Siblings: &[]BasicFish{},
						Birthday: &date.Time{time.Date(1900, time.January, 5, 1, 0, 0, 0, time.UTC)},
						Picture:  &[]byte{255, 255, 255, 255, 254},
					},
				},
			},
			Sawshark{
				Age:      to.Int32Ptr(105),
				Length:   to.Float64Ptr(10),
				Species:  to.StringPtr("dangerous"),
				Siblings: &[]BasicFish{},
				Birthday: &date.Time{time.Date(1900, time.January, 5, 1, 0, 0, 0, time.UTC)},
				Picture:  &[]byte{255, 255, 255, 255, 254},
			},
		},
	}
	_, err := complexPolymorphicRecursiveClient.PutValid(context.Background(), recF)
	c.Assert(err, chk.IsNil)
}

func (s *ComplexGroupSuite) TestGetComplexPolymorphicComplicated(c *chk.C) {
	resp, err := complexPolymorphicClient.GetComplicated(context.Background())
	c.Assert(err, chk.IsNil)

	salmon, b := resp.Value.AsSmartSalmon()
	c.Assert(b, chk.Equals, true)
	c.Assert(*salmon, chk.FitsTypeOf, ss)

	c.Assert(resp.Value, chk.FitsTypeOf, ss)
	c.Assert(*salmon.Iswild, chk.Equals, *ss.Iswild)
	c.Assert(*salmon.Location, chk.Equals, *ss.Location)
	c.Assert(*salmon.Species, chk.Equals, *ss.Species)
	c.Assert(*salmon.Length, chk.Equals, *ss.Length)
	c.Assert(*salmon.Siblings, chk.HasLen, len(*ss.Siblings))

	c.Assert((*salmon.Siblings)[0], chk.FitsTypeOf, (*ss.Siblings)[0].(Shark))
	c.Assert(*(*salmon.Siblings)[0].(Shark).Length, chk.Equals, *(*ss.Siblings)[0].(Shark).Length)
	c.Assert(*(*salmon.Siblings)[0].(Shark).Birthday, chk.Equals, *(*ss.Siblings)[0].(Shark).Birthday)
	c.Assert(*(*salmon.Siblings)[0].(Shark).Age, chk.Equals, *(*ss.Siblings)[0].(Shark).Age)
	c.Assert(*(*salmon.Siblings)[0].(Shark).Species, chk.Equals, *(*ss.Siblings)[0].(Shark).Species)

	c.Assert((*salmon.Siblings)[1], chk.FitsTypeOf, (*ss.Siblings)[1].(Sawshark))
	c.Assert(*(*salmon.Siblings)[1].(Sawshark).Length, chk.Equals, *(*ss.Siblings)[1].(Sawshark).Length)
	c.Assert(*(*salmon.Siblings)[1].(Sawshark).Birthday, chk.Equals, *(*ss.Siblings)[1].(Sawshark).Birthday)
	c.Assert(*(*salmon.Siblings)[1].(Sawshark).Age, chk.Equals, *(*ss.Siblings)[1].(Sawshark).Age)
	c.Assert(*(*salmon.Siblings)[1].(Sawshark).Species, chk.Equals, *(*ss.Siblings)[1].(Sawshark).Species)
	c.Assert(*(*salmon.Siblings)[1].(Sawshark).Picture, chk.DeepEquals, *(*ss.Siblings)[1].(Sawshark).Picture)

	c.Assert((*salmon.Siblings)[2], chk.FitsTypeOf, (*ss.Siblings)[2].(Goblinshark))
	c.Assert(*(*salmon.Siblings)[2].(Goblinshark).Length, chk.Equals, *(*ss.Siblings)[2].(Goblinshark).Length)
	c.Assert(*(*salmon.Siblings)[2].(Goblinshark).Birthday, chk.Equals, *(*ss.Siblings)[2].(Goblinshark).Birthday)
	c.Assert(*(*salmon.Siblings)[2].(Goblinshark).Age, chk.Equals, *(*ss.Siblings)[2].(Goblinshark).Age)
	c.Assert(*(*salmon.Siblings)[2].(Goblinshark).Species, chk.Equals, *(*ss.Siblings)[2].(Goblinshark).Species)
	c.Assert(*(*salmon.Siblings)[2].(Goblinshark).Jawsize, chk.Equals, *(*ss.Siblings)[2].(Goblinshark).Jawsize)

	c.Assert(len(salmon.AdditionalProperties), chk.Equals, len(ss.AdditionalProperties))
	c.Assert(salmon.AdditionalProperties["additionalProperty1"], chk.Equals, ss.AdditionalProperties["additionalProperty1"])
	c.Assert(salmon.AdditionalProperties["additionalProperty2"], chk.Equals, ss.AdditionalProperties["additionalProperty2"])
	c.Assert(salmon.AdditionalProperties["additionalProperty3"], chk.Equals, ss.AdditionalProperties["additionalProperty3"])
	c.Assert(salmon.AdditionalProperties["additionalProperty4"], chk.DeepEquals, ss.AdditionalProperties["additionalProperty4"])
	c.Assert(salmon.AdditionalProperties["additionalProperty5"], chk.DeepEquals, ss.AdditionalProperties["additionalProperty5"])
}

func (s *ComplexGroupSuite) TestPutComplexPolymorphicComplicated(c *chk.C) {
	_, err := complexPolymorphicClient.PutComplicated(context.Background(), ss)
	c.Assert(err, chk.IsNil)
}
