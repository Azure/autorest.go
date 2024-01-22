/*---------------------------------------------------------------------------------------------
 *  Copyright (c) Microsoft Corporation. All rights reserved.
 *  Licensed under the MIT License. See License.txt in the project root for license information.
 *--------------------------------------------------------------------------------------------*/

import { values } from '@azure-tools/linq';
import * as go from '../../codemodel.go/src/gocodemodel.js';
import { contentPreamble } from './helpers.js';
import { ImportManager } from './imports.js';

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
export async function generateTimeHelpers(codeModel: go.CodeModel, packageName?: string): Promise<Array<Content>> {
  const content = new Array<Content>();
  if (!codeModel.marshallingRequirements.generateDateTimeRFC1123Helper &&
    !codeModel.marshallingRequirements.generateDateTimeRFC3339Helper &&
	!codeModel.marshallingRequirements.generateTimeRFC3339Helper &&
    !codeModel.marshallingRequirements.generateUnixTimeHelper &&
    !codeModel.marshallingRequirements.generateDateHelper) {
    return content;
  }
  let needsPopulate = false;
  for (const model of codeModel.models) {
    if (model.format !== 'json') {
      // population helpers are for JSON only
      continue;
    }
    for (const field of values(model.fields)) {
      if (go.isTimeType(field.type)) {
        needsPopulate = true;
        break;
      }
    }
    if (needsPopulate) {
      break;
    }
  }
  const preamble = contentPreamble(codeModel, packageName);
  if (codeModel.marshallingRequirements.generateDateTimeRFC1123Helper) {
    content.push(new Content('time_rfc1123', generateRFC1123Helper(preamble, needsPopulate)));
  }
  if (codeModel.marshallingRequirements.generateDateTimeRFC3339Helper || codeModel.marshallingRequirements.generateTimeRFC3339Helper) {
    content.push(new Content('time_rfc3339', generateRFC3339Helper(preamble, codeModel.marshallingRequirements.generateDateTimeRFC3339Helper, codeModel.marshallingRequirements.generateTimeRFC3339Helper, needsPopulate)));
  }
  if (codeModel.marshallingRequirements.generateUnixTimeHelper) {
    content.push(new Content('time_unix', generateUnixTimeHelper(preamble, needsPopulate)));
  }
  if (codeModel.marshallingRequirements.generateDateHelper) {
    content.push(new Content('date_type', generateDateHelper(preamble, needsPopulate)));
  }
  return content;
}

function generateRFC1123Helper(preamble: string, needsPopulate: boolean): string {
  const imports = new ImportManager();
  imports.add('strings');
  imports.add('time');
  if (needsPopulate) {
    imports.add('encoding/json');
    imports.add('fmt');
    imports.add('github.com/Azure/azure-sdk-for-go/sdk/azcore');
    imports.add('reflect');
  }
  let text = `${preamble}

${imports.text()}

const (
	dateTimeRFC1123JSON = \`"\` + time.RFC1123 + \`"\`
)

type dateTimeRFC1123 time.Time

func (t dateTimeRFC1123) MarshalJSON() ([]byte, error) {
	b := []byte(time.Time(t).Format(dateTimeRFC1123JSON))
	return b, nil
}

func (t dateTimeRFC1123) MarshalText() ([]byte, error) {
	b := []byte(time.Time(t).Format(time.RFC1123))
	return b, nil
}

func (t *dateTimeRFC1123) UnmarshalJSON(data []byte) error {
	p, err := time.Parse(dateTimeRFC1123JSON, strings.ToUpper(string(data)))
	*t = dateTimeRFC1123(p)
	return err
}

func (t *dateTimeRFC1123) UnmarshalText(data []byte) error {
	p, err := time.Parse(time.RFC1123, string(data))
	*t = dateTimeRFC1123(p)
	return err
}

func (t dateTimeRFC1123) String() string {
	return time.Time(t).Format(time.RFC1123)
}
`;
  if (needsPopulate) {
    text +=
`

func populateDateTimeRFC1123(m map[string]any, k string, t *time.Time) {
	if t == nil {
		return
	} else if azcore.IsNullValue(t) {
		m[k] = nil
		return
	} else if reflect.ValueOf(t).IsNil() {
		return
	}
	m[k] = (*dateTimeRFC1123)(t)
}

func unpopulateDateTimeRFC1123(data json.RawMessage, fn string, t **time.Time) error {
	if data == nil || string(data) == "null" {
		return nil
	}
	var aux dateTimeRFC1123
	if err := json.Unmarshal(data, &aux); err != nil {
		return fmt.Errorf("struct field %s: %v", fn, err)
	}
	*t = (*time.Time)(&aux)
	return nil
}
`;
  }
  return text;
}

function generateRFC3339Helper(preamble: string, dateTime: boolean, time: boolean, needsPopulate: boolean): string {
  const imports = new ImportManager();
  imports.add('regexp');
  imports.add('strings');
  imports.add('time');
  if (needsPopulate) {
    imports.add('encoding/json');
    imports.add('fmt');
    imports.add('github.com/Azure/azure-sdk-for-go/sdk/azcore');
    imports.add('reflect');
  }
  let text = `${preamble}

${imports.text()}

// Azure reports time in UTC but it doesn't include the 'Z' time zone suffix in some cases.
var tzOffsetRegex = regexp.MustCompile(\`(Z|z|\\+|-)(\\d+:\\d+)*"*$\`)
`;

  if (dateTime) {
    text +=
`
const (
	utcDateTimeJSON = \`"2006-01-02T15:04:05.999999999"\`
	utcDateTime     = "2006-01-02T15:04:05.999999999"
	dateTimeJSON   = \`"\` + time.RFC3339Nano + \`"\`
)

type dateTimeRFC3339 time.Time

func (t dateTimeRFC3339) MarshalJSON() ([]byte, error) {
	tt := time.Time(t)
	return tt.MarshalJSON()
}

func (t dateTimeRFC3339) MarshalText() ([]byte, error) {
	tt := time.Time(t)
	return tt.MarshalText()
}

func (t *dateTimeRFC3339) UnmarshalJSON(data []byte) error {
	layout := utcDateTimeJSON
	if tzOffsetRegex.Match(data) {
		layout = dateTimeJSON
	}
	return t.Parse(layout, string(data))
}

func (t *dateTimeRFC3339) UnmarshalText(data []byte) (error) {
	layout := utcDateTime
	if tzOffsetRegex.Match(data) {
		layout = time.RFC3339Nano
	}
	return t.Parse(layout, string(data))
}

func (t *dateTimeRFC3339) Parse(layout, value string) error {
	p, err := time.Parse(layout, strings.ToUpper(value))
	*t = dateTimeRFC3339(p)
	return err
}

func (t dateTimeRFC3339) String() string {
	return time.Time(t).Format(time.RFC3339Nano)
}
`;
    if (needsPopulate) {
      text +=
`

func populateDateTimeRFC3339(m map[string]any, k string, t *time.Time) {
	if t == nil {
		return
	} else if azcore.IsNullValue(t) {
		m[k] = nil
		return
	} else if reflect.ValueOf(t).IsNil() {
		return
	}
	m[k] = (*dateTimeRFC3339)(t)
}

func unpopulateDateTimeRFC3339(data json.RawMessage, fn string, t **time.Time) error {
	if data == nil || string(data) == "null" {
		return nil
	}
	var aux dateTimeRFC3339
	if err := json.Unmarshal(data, &aux); err != nil {
		return fmt.Errorf("struct field %s: %v", fn, err)
	}
	*t = (*time.Time)(&aux)
	return nil
}
`;
    }
  }

  if (time) {
    text += `
const (
	utcTimeJSON = \`"15:04:05.999999999"\`
	utcTime     = "15:04:05.999999999"
	timeFormat  = "15:04:05.999999999Z07:00"
)

type timeRFC3339 time.Time

func (t timeRFC3339) MarshalJSON() ([]byte, error) {
	s, _ := t.MarshalText()
	return []byte(fmt.Sprintf("\\"%s\\"", s)), nil
}

func (t timeRFC3339) MarshalText() ([]byte, error) {
	tt := time.Time(t)
	return []byte(tt.Format(timeFormat)), nil
}

func (t *timeRFC3339) UnmarshalJSON(data []byte) error {
	layout := utcTimeJSON
	if tzOffsetRegex.Match(data) {
		layout = timeFormat
	}
	return t.Parse(layout, string(data))
}

func (t *timeRFC3339) UnmarshalText(data []byte) error {
	layout := utcTime
	if tzOffsetRegex.Match(data) {
		layout = timeFormat
	}
	return t.Parse(layout, string(data))
}

func (t *timeRFC3339) Parse(layout, value string) error {
	p, err := time.Parse(layout, strings.ToUpper(value))
	*t = timeRFC3339(p)
	return err
}

func (t timeRFC3339) String() string {
	tt := time.Time(t)
	return tt.Format(timeFormat)
}
`;
    if (needsPopulate) {
      text += `
func populateTimeRFC3339(m map[string]any, k string, t *time.Time) {
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

func unpopulateTimeRFC3339(data json.RawMessage, fn string, t **time.Time) error {
	if data == nil || string(data) == "null" {
		return nil
	}
	var aux timeRFC3339
	if err := json.Unmarshal(data, &aux); err != nil {
		return fmt.Errorf("struct field %s: %v", fn, err)
	}
	*t = (*time.Time)(&aux)
	return nil
}
`;
    }
  }
  return text;
}

function generateUnixTimeHelper(preamble: string, needsPopulate: boolean): string {
  const imports = new ImportManager();
  imports.add('encoding/json');
  imports.add('fmt');
  imports.add('time');
  if (needsPopulate) {
    imports.add('github.com/Azure/azure-sdk-for-go/sdk/azcore');
    imports.add('reflect');
  }
  let text = `${preamble}

${imports.text()}

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
`;
  if (needsPopulate) {
    text +=
`
func populateTimeUnix(m map[string]any, k string, t *time.Time) {
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

func unpopulateTimeUnix(data json.RawMessage, fn string, t **time.Time) error {
	if data == nil || string(data) == "null" {
		return nil
	}
	var aux timeUnix
	if err := json.Unmarshal(data, &aux); err != nil {
		return fmt.Errorf("struct field %s: %v", fn, err)
	}
	*t = (*time.Time)(&aux)
	return nil
}
`;
  }
  return text;
}

function generateDateHelper(preamble: string, needsPopulate: boolean): string {
  const imports = new ImportManager();
  imports.add('fmt');
  imports.add('time');
  if (needsPopulate) {
    imports.add('encoding/json');
    imports.add('github.com/Azure/azure-sdk-for-go/sdk/azcore');
    imports.add('reflect');
  }
  let text = `${preamble}

${imports.text()}

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
`;
  if (needsPopulate) {
    text +=
`
func populateDateType(m map[string]any, k string, t *time.Time) {
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

func unpopulateDateType(data json.RawMessage, fn string, t **time.Time) error {
	if data == nil || string(data) == "null" {
		return nil
	}
	var aux dateType
	if err := json.Unmarshal(data, &aux); err != nil {
		return fmt.Errorf("struct field %s: %v", fn, err)
	}
	*t = (*time.Time)(&aux)
	return nil
}
`;
  }
  return text;
}
