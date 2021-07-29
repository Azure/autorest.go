/*---------------------------------------------------------------------------------------------
 *  Copyright (c) Microsoft Corporation. All rights reserved.
 *  Licensed under the MIT License. See License.txt in the project root for license information.
 *--------------------------------------------------------------------------------------------*/
import { SchemaType, Parameter, Schema } from '@autorest/codemodel';
import { sortParametersByRequired } from '../generator/helpers';

describe('parameter ordering', () => {
    test('required are before optional', () => {
        // Create an array of parameters with required intermixed 
        var params = Array(4).fill(0).map((_, i) => {
            var p = new Parameter (String(i), "",  new Schema("Test", "", SchemaType.String));
            p.required = i % 2 == 0;
            return p;
        });
        
        params.sort(sortParametersByRequired);
        expect(params[0].required).toBe(false);
        expect(params[1].required).toBe(true);
        expect(params[2].required).toBe(false);
        expect(params[3].required).toBe(false);
    });    
});

