/*---------------------------------------------------------------------------------------------
 *  Copyright (c) Microsoft Corporation. All rights reserved.
 *  Licensed under the MIT License. See License.txt in the project root for license information.
 *--------------------------------------------------------------------------------------------*/

import * as go from '../../../codemodel.go/src/index.js';
import * as helpers from './helpers.js';
import { ImportManager } from './imports.js';

/**
 * Creates the content for the required additional properties XML marshalling helpers.
 * 
 * @param pkg contains the package content
 * @returns the text for the file or the empty string
 */
export function generateXMLAdditionalPropsHelpers(pkg: go.PackageContent): string {
  // check if any models need this helper
  let required = false;
  for (const model of pkg.models) {
    if (helpers.getSerDeFormat(model, pkg) !== 'XML') {
      continue;
    }
    for (const field of model.fields) {
      if (field.type.kind === 'map') {
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

  let text = helpers.contentPreamble(pkg);
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

// MarshalXML implements the xml.Marshaler interface for additionalProperties.
func (ap additionalProperties) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if err := e.EncodeToken(start); err != nil {
		return err
	}
	for k, v := range ap {
		err := e.EncodeToken(xml.StartElement{
			Name: xml.Name{
				Local: k,
			},
		})
		if err != nil {
			return err
		}
		if v != nil {
			err = e.EncodeToken(xml.CharData(*v))
			if err != nil {
				return err
			}
		}
		err = e.EncodeToken(xml.EndElement{
			Name: xml.Name{
				Local: k,
			},
		})
		if err != nil {
			return err
		}
	}
	return e.EncodeToken(xml.EndElement{
		Name: start.Name,
	})
}

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
