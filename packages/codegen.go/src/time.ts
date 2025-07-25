/*---------------------------------------------------------------------------------------------
 *  Copyright (c) Microsoft Corporation. All rights reserved.
 *  Licensed under the MIT License. See License.txt in the project root for license information.
 *--------------------------------------------------------------------------------------------*/

import { values } from '@azure-tools/linq';
import * as go from '../../codemodel.go/src/index.js';
import { contentPreamble, getSerDeFormat, recursiveUnwrapMapSlice } from './helpers.js';
import { ImportManager } from './imports.js';
import { CodegenError } from './errors.js';

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
  let needsDateHelper = false;
  let needsDateTimeRFC1123Helper = false;
  let needsDateTimeRFC3339Helper = false;
  let needsTimeRFC3339Helper = false;
  let needsUnixTimeHelper = false;

  const setHelper = function(dateTimeFormat: go.TimeFormat): void {
    switch (dateTimeFormat) {
      case 'dateTimeRFC1123':
        needsDateTimeRFC1123Helper = true;
        break;
      case 'dateTimeRFC3339':
        needsDateTimeRFC3339Helper = true;
        break;
      case 'dateType':
        needsDateHelper = true;
        break;
      case 'timeRFC3339':
        needsTimeRFC3339Helper = true;
        break;
      case 'timeUnix':
        needsUnixTimeHelper = true;
        break;
      default:
        throw new CodegenError('InternalError', `unhandled date-time format ${dateTimeFormat}`);
    }
  };

  // needsSerDeHelpers is only required when time.Time is a struct field
  let needsSerDeHelpers = false;

  // find the required helpers.
  // for most packages, we must check params, response envelopes, and models
  // for fakes, we check a subset
  if (packageName !== 'fake') {
    for (const client of codeModel.clients) {
      for (const method of client.methods) {
        for (const param of method.parameters) {
          const unwrappedParam = recursiveUnwrapMapSlice(param.type);
          if (unwrappedParam.kind !== 'time') {
            continue;
          }
          // for body params, the helpers are always required.
          // for header/path/query params, the conversion happens in place. the only
          // exceptions are for timeRFC3339 and timeUnix
          // TODO: clean this up when moving to DateTime type in azcore
          if (param.kind === 'bodyParam' || unwrappedParam.format === 'timeRFC3339' || unwrappedParam.format === 'timeUnix') {
            setHelper(unwrappedParam.format);
          }
        }
      }
    }

    for (const model of codeModel.models) {
      for (const field of values(model.fields)) {
        const unwrappedField = recursiveUnwrapMapSlice(field.type);
        if (unwrappedField.kind !== 'time') {
          continue;
        }
        if (getSerDeFormat(model, codeModel) === 'JSON') {
          // needsSerDeHelpers helpers are for JSON only
          needsSerDeHelpers = true;
        }
        setHelper(unwrappedField.format);
      }
    }

    for (const respEnv of codeModel.responseEnvelopes) {
      if (!respEnv.result || respEnv.result.kind !== 'monomorphicResult' || respEnv.result.format !== 'JSON') {
        continue;
      }
      const unwrappedResult = recursiveUnwrapMapSlice(respEnv.result.monomorphicType);
      if (unwrappedResult.kind !== 'time') {
        continue;
      }
      setHelper(unwrappedResult.format);
    }
  } else {
	// for fakes, only need to check the if the body params are of type time.Time.
	// otherwise, the conversion happens in place
    for (const client of codeModel.clients) {
      for (const method of client.methods) {
        for (const param of method.parameters) {
          if (param.kind === 'bodyParam' && param.type.kind === 'time') {
            setHelper(param.type.format);
          }
        }
      }
    }

    for (const respEnv of codeModel.responseEnvelopes) {
      for (const header of respEnv.headers) {
        // for header/path/query params, the conversion happens in place. the only
        // exceptions are for timeRFC3339 and timeUnix
        if (header.type.kind === 'time' && (header.type.format === 'timeRFC3339' || header.type.format === 'timeUnix')) {
          setHelper(header.type.format);
        }
      }
      if (respEnv.result?.kind === 'monomorphicResult' && respEnv.result.monomorphicType.kind === 'time') {
        setHelper(respEnv.result.monomorphicType.format);
      }
    }
  }

  const content = new Array<Content>();
  if (!needsDateHelper &&
    !needsDateTimeRFC1123Helper &&
    !needsDateTimeRFC3339Helper &&
    !needsTimeRFC3339Helper &&
    !needsUnixTimeHelper) {
    return content;
  }

  const preamble = contentPreamble(codeModel, true, packageName);
  if (needsDateTimeRFC1123Helper) {
    content.push(new Content('time_rfc1123', generateRFC1123Helper(preamble, needsSerDeHelpers)));
  }
  if (needsDateTimeRFC3339Helper || needsTimeRFC3339Helper) {
    content.push(new Content('time_rfc3339', generateRFC3339Helper(preamble, needsDateTimeRFC3339Helper, needsTimeRFC3339Helper, needsSerDeHelpers)));
  }
  if (needsUnixTimeHelper) {
    content.push(new Content('time_unix', generateUnixTimeHelper(preamble, needsSerDeHelpers)));
  }
  if (needsDateHelper) {
    content.push(new Content('date_type', generateDateHelper(preamble, needsSerDeHelpers)));
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
	if len(data) == 0 {
		return nil
	}
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
var tzOffsetRegex = regexp.MustCompile(\`(?:Z|z|\\+|-)(?:\\d+:\\d+)*"*$\`)
`;

  if (dateTime) {
    text +=
`
const (
	utcDateTime        = "2006-01-02T15:04:05.999999999"
	utcDateTimeJSON    = \`"\` + utcDateTime + \`"\`
	utcDateTimeNoT     = "2006-01-02 15:04:05.999999999"
	utcDateTimeJSONNoT = \`"\` + utcDateTimeNoT + \`"\`
	dateTimeNoT        = \`2006-01-02 15:04:05.999999999Z07:00\`
	dateTimeJSON       = \`"\` + time.RFC3339Nano + \`"\`
	dateTimeJSONNoT    = \`"\` + dateTimeNoT + \`"\`
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
	tzOffset := tzOffsetRegex.Match(data)
	hasT := strings.Contains(string(data), "T") || strings.Contains(string(data), "t")
	var layout string
	if tzOffset && hasT {
		layout = dateTimeJSON
	} else if tzOffset {
		layout = dateTimeJSONNoT
	} else if hasT {
		layout = utcDateTimeJSON
	} else {
		layout = utcDateTimeJSONNoT
	}
	return t.Parse(layout, string(data))
}

func (t *dateTimeRFC3339) UnmarshalText(data []byte) error {
	if len(data) == 0 {
		return nil
	}
	tzOffset := tzOffsetRegex.Match(data)
	hasT := strings.Contains(string(data), "T") || strings.Contains(string(data), "t")
	var layout string
	if tzOffset && hasT {
		layout = time.RFC3339Nano
	} else if tzOffset {
		layout = dateTimeNoT
	} else if hasT {
		layout = utcDateTime
	} else {
		layout = utcDateTimeNoT
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
	if len(data) == 0 {
		return nil
	}
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
