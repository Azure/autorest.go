/*---------------------------------------------------------------------------------------------
 *  Copyright (c) Microsoft Corporation. All rights reserved.
 *  Licensed under the MIT License. See License.txt in the project root for license information.
 *--------------------------------------------------------------------------------------------*/

import { GoCodeModel } from '../gocodemodel/gocodemodel';
import { contentPreamble } from './helpers';
import { ImportManager } from './imports';

// Creates the content for required additional properties XML marshalling helpers.
// Will be empty if no helpers are required.
export async function generateXMLAdditionalPropsHelpers(codeModel: GoCodeModel): Promise<string> {
  if (!codeModel.marshallingRequirements.generateXMLDictionaryUnmarshallingHelper) {
    return '';
  }
  let text = contentPreamble(codeModel);
  // add standard imports
  const imports = new ImportManager();
  imports.add('encoding/xml');
  imports.add('strings');
  text += imports.text();
  text += `
type additionalProperties map[string]*string

// UnmarshalXML implements the xml.Unmarshaler interface for additionalProperties.
func (ap *additionalProperties) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	tokName := ""
	for t, err := d.Token(); err == nil; t, err = d.Token() {
		switch tt := t.(type) {
		case xml.StartElement:
			tokName = strings.ToLower(tt.Name.Local)
			break
		case xml.CharData:
			if tokName == "" {
				continue
			}
			if *ap == nil {
				*ap = additionalProperties{}
			}
			s := string(tt)
			(*ap)[tokName] = &s
			tokName = ""
			break
		}
	}
	return nil
}
`;
  return text;
}
