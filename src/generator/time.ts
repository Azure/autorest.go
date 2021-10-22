/*---------------------------------------------------------------------------------------------
 *  Copyright (c) Microsoft Corporation. All rights reserved.
 *  Licensed under the MIT License. See License.txt in the project root for license information.
 *--------------------------------------------------------------------------------------------*/

import { Session } from '@autorest/extension-base';
import { CodeModel } from '@autorest/codemodel';
import { contentPreamble } from './helpers'

// represents the generated content for an operation group
export class Content {
  readonly name: string;
  readonly content: string;

  constructor(name: string, content: string) {
    this.name = name;
    this.content = content;
  }
}

// Creates the content for required time marshalling helpers.
// Will be empty if no helpers are required.
export async function generateTimeHelpers(session: Session<CodeModel>): Promise<Content[]> {
  const preamble = await contentPreamble(session);
  const content = new Array<Content>();
  if (session.model.language.go!.hasTimeRFC1123) {
    content.push(new Content('time_rfc1123', generateRFC1123Helper(preamble)));
  }
  if (session.model.language.go!.hasTimeRFC3339) {
    content.push(new Content('time_rfc3339', generateRFC3339Helper(preamble)));
  }
  if (session.model.language.go!.hasUnixTime) {
    content.push(new Content('time_unix', generateUnixTimeHelper(preamble)));
  }
  if (session.model.language.go!.hasDate) {
    content.push(new Content('date_type', generateDateHelper(preamble)));
  }
  return content;
}

function generateRFC1123Helper(preamble: string): string {
  return `${preamble}

import (
	"encoding/json"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"reflect"
	"strings"
	"time"
)

const (
	rfc1123JSON = \`"\` + time.RFC1123 + \`"\`
)

type timeRFC1123 time.Time

func (t timeRFC1123) MarshalJSON() ([]byte, error) {
	b := []byte(time.Time(t).Format(rfc1123JSON))
	return b, nil
}

func (t timeRFC1123) MarshalText() ([]byte, error) {
	b := []byte(time.Time(t).Format(time.RFC1123))
	return b, nil
}

func (t *timeRFC1123) UnmarshalJSON(data []byte) error {
	p, err := time.Parse(rfc1123JSON, strings.ToUpper(string(data)))
	*t = timeRFC1123(p)
	return err
}

func (t *timeRFC1123) UnmarshalText(data []byte) error {
	p, err := time.Parse(time.RFC1123, string(data))
	*t = timeRFC1123(p)
	return err
}

func populateTimeRFC1123(m map[string]interface{}, k string, t *time.Time) {
	if t == nil {
		return
	} else if azcore.IsNullValue(t) {
		m[k] = nil
		return
	} else if reflect.ValueOf(t).IsNil() {
		return
	}
	m[k] = (*timeRFC1123)(t)
}

func unpopulateTimeRFC1123(data json.RawMessage, t **time.Time) error {
	if data == nil || strings.EqualFold(string(data), "null") {
		return nil
	}
	var aux timeRFC1123
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	*t = (*time.Time)(&aux)
	return nil
}
`;
}

function generateRFC3339Helper(preamble: string): string {
  return `${preamble}

import (
	"encoding/json"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"reflect"
	"regexp"
	"strings"
	"time"
)

const (
	utcLayoutJSON = \`"2006-01-02T15:04:05.999999999"\`
	utcLayout     = "2006-01-02T15:04:05.999999999"
	rfc3339JSON   = \`"\` + time.RFC3339Nano + \`"\`
)

// Azure reports time in UTC but it doesn't include the 'Z' time zone suffix in some cases.
var tzOffsetRegex = regexp.MustCompile(\`(Z|z|\\+|-)(\\d+:\\d+)*"*$\`)

type timeRFC3339 time.Time

func (t timeRFC3339) MarshalJSON() (json []byte, err error) {
	tt := time.Time(t)
	return tt.MarshalJSON()
}

func (t timeRFC3339) MarshalText() (text []byte, err error) {
	tt := time.Time(t)
	return tt.MarshalText()
}

func (t *timeRFC3339) UnmarshalJSON(data []byte) error {
	layout := utcLayoutJSON
	if tzOffsetRegex.Match(data) {
		layout = rfc3339JSON
	}
	return t.Parse(layout, string(data))
}

func (t *timeRFC3339) UnmarshalText(data []byte) (err error) {
	layout := utcLayout
	if tzOffsetRegex.Match(data) {
		layout = time.RFC3339Nano
	}
	return t.Parse(layout, string(data))
}

func (t *timeRFC3339) Parse(layout, value string) error {
	p, err := time.Parse(layout, strings.ToUpper(value))
	*t = timeRFC3339(p)
	return err
}

func populateTimeRFC3339(m map[string]interface{}, k string, t *time.Time) {
	if t == nil {
		return
	} else if azcore.IsNullValue(t) {
		m[k] = nil
		return
	} else if reflect.ValueOf(t).IsNil() {
		return
	}
	m[k] = (*timeRFC3339)(t)
}

func unpopulateTimeRFC3339(data json.RawMessage, t **time.Time) error {
	if data == nil || strings.EqualFold(string(data), "null") {
		return nil
	}
	var aux timeRFC3339
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	*t = (*time.Time)(&aux)
	return nil
}
`;
}

function generateUnixTimeHelper(preamble: string): string {
  return `${preamble}

import (
	"encoding/json"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"reflect"
	"strings"
	"time"
)

type timeUnix time.Time

func (t timeUnix) MarshalJSON() ([]byte, error) {
	return json.Marshal(time.Time(t).Unix())
}

func (t *timeUnix) UnmarshalJSON(data []byte) error {
	var seconds int64
	if err := json.Unmarshal(data, &seconds); err != nil {
		return err
	}
	*t = timeUnix(time.Unix(seconds, 0))
	return nil
}

func (t timeUnix) String() string {
	return fmt.Sprintf("%d", time.Time(t).Unix())
}

func populateTimeUnix(m map[string]interface{}, k string, t *time.Time) {
	if t == nil {
		return
	} else if azcore.IsNullValue(t) {
		m[k] = nil
		return
	} else if reflect.ValueOf(t).IsNil() {
		return
	}
	m[k] = (*timeUnix)(t)
}

func unpopulateTimeUnix(data json.RawMessage, t **time.Time) error {
	if data == nil || strings.EqualFold(string(data), "null") {
		return nil
	}
	var aux timeUnix
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	*t = (*time.Time)(&aux)
	return nil
}
`;
}

function generateDateHelper(preamble: string): string {
  return `${preamble}

import (
	"encoding/json"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"reflect"
	"strings"
	"time"
)

const (
	fullDateJSON = \`"2006-01-02"\`
	jsonFormat   = \`"%04d-%02d-%02d"\`
)

type dateType time.Time

func (t dateType) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(jsonFormat, time.Time(t).Year(), time.Time(t).Month(), time.Time(t).Day())), nil
}

func (d *dateType) UnmarshalJSON(data []byte) (err error) {
	t, err := time.Parse(fullDateJSON, string(data))
	*d = (dateType)(t)
	return err
}

func populateDateType(m map[string]interface{}, k string, t *time.Time) {
	if t == nil {
		return
	} else if azcore.IsNullValue(t) {
		m[k] = nil
		return
	} else if reflect.ValueOf(t).IsNil() {
		return
	}
	m[k] = (*dateType)(t)
}

func unpopulateDateType(data json.RawMessage, t **time.Time) error {
	if data == nil || strings.EqualFold(string(data), "null") {
		return nil
	}
	var aux dateType
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	*t = (*time.Time)(&aux)
	return nil
}
`;
}
