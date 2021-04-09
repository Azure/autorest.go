/*---------------------------------------------------------------------------------------------
 *  Copyright (c) Microsoft Corporation. All rights reserved.
 *  Licensed under the MIT License. See License.txt in the project root for license information.
 *--------------------------------------------------------------------------------------------*/

import { Session } from '@autorest/extension-base';
import { CodeModel } from '@autorest/codemodel';
import { contentPreamble } from './helpers'
import { ImportManager } from './imports';

// Creates the content for required additional properties XML marshalling helpers.
// Will be empty if no helpers are required.
export async function generateXMLAdditionalPropsHelpers(session: Session<CodeModel>): Promise<string> {
  if (!session.model.language.go!.needsXMLDictionaryUnmarshalling) {
    return '';
  }
  let text = await contentPreamble(session);
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
	tokVal := ""
	inserted := false
	for t, err := d.Token(); err == nil; t, err = d.Token() {
		switch tt := t.(type) {
		case xml.StartElement:
			tokName = strings.ToLower(tt.Name.Local)
		case xml.CharData:
			temp := string(tt)
			if tokName != "" {
				if *ap == nil {
					*ap = additionalProperties{}
				}
				(*ap)[tokName] = &temp
				tokName = ""
				inserted = true
			}
			tokVal = temp
		case xml.EndElement:
			if !inserted && tokName == "" && tt.Name.Local != "" {
				if *ap == nil {
					*ap = additionalProperties{}
				}
				(*ap)[tt.Name.Local] = &tokVal
				inserted = false
			}
		}
	}
	return nil
}
`;
  return text;
}
