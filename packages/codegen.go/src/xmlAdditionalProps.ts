/*---------------------------------------------------------------------------------------------
 *  Copyright (c) Microsoft Corporation. All rights reserved.
 *  Licensed under the MIT License. See License.txt in the project root for license information.
 *--------------------------------------------------------------------------------------------*/

import * as go from '../../codemodel.go/src/index.js';
import { contentPreamble } from './helpers.js';
import { ImportManager } from './imports.js';

// Creates the content for required additional properties XML marshalling helpers.
// Will be empty if no helpers are required.
export async function generateXMLAdditionalPropsHelpers(codeModel: go.CodeModel): Promise<string> {
  // check if any models need this helper
  let required = false;
  for (const model of codeModel.models) {
    if (model.format !== 'xml') {
      continue;
    }
    for (const field of model.fields) {
      if (go.isMapType(field.type)) {
        required = true;
        break;
      }
    }
    if (required) {
      break;
    }
  }

  if (!required) {
    return '';
  }

  let text = contentPreamble(codeModel);
  // add standard imports
  const imports = new ImportManager();
  imports.add('encoding/xml');
  imports.add('errors');
  imports.add('github.com/Azure/azure-sdk-for-go/sdk/azcore/to');
  imports.add('io');
  imports.add('strings');
  text += imports.text();
  text += `
type additionalProperties map[string]*string

// UnmarshalXML implements the xml.Unmarshaler interface for additionalProperties.
func (ap *additionalProperties) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	tokName := ""
	tokValue := ""
	for {
		t, err := d.Token()
		if errors.Is(err, io.EOF) {
			break
		} else if err != nil {
			return err
		}
		switch tt := t.(type) {
		case xml.StartElement:
			tokName = strings.ToLower(tt.Name.Local)
			tokValue = ""
		case xml.CharData:
			if tokName == "" {
				continue
			}
			tokValue = string(tt)
		case xml.EndElement:
			if tokName == "" {
				continue
			}
			if *ap == nil {
				*ap = additionalProperties{}
			}
			(*ap)[tokName] = to.Ptr(tokValue)
			tokName = ""
		}
	}
	return nil
}
`;
  return text;
}
