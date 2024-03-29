// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armnetwork

import "encoding/json"

func unmarshalActiveBaseSecurityAdminRuleClassification(rawMsg json.RawMessage) (ActiveBaseSecurityAdminRuleClassification, error) {
	if rawMsg == nil || string(rawMsg) == "null" {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b ActiveBaseSecurityAdminRuleClassification
	switch m["kind"] {
	case string(EffectiveAdminRuleKindCustom):
		b = &ActiveSecurityAdminRule{}
	case string(EffectiveAdminRuleKindDefault):
		b = &ActiveDefaultSecurityAdminRule{}
	default:
		b = &ActiveBaseSecurityAdminRule{}
	}
	if err := json.Unmarshal(rawMsg, b); err != nil {
		return nil, err
	}
	return b, nil
}

func unmarshalActiveBaseSecurityAdminRuleClassificationArray(rawMsg json.RawMessage) ([]ActiveBaseSecurityAdminRuleClassification, error) {
	if rawMsg == nil || string(rawMsg) == "null" {
		return nil, nil
	}
	var rawMessages []json.RawMessage
	if err := json.Unmarshal(rawMsg, &rawMessages); err != nil {
		return nil, err
	}
	fArray := make([]ActiveBaseSecurityAdminRuleClassification, len(rawMessages))
	for index, rawMessage := range rawMessages {
		f, err := unmarshalActiveBaseSecurityAdminRuleClassification(rawMessage)
		if err != nil {
			return nil, err
		}
		fArray[index] = f
	}
	return fArray, nil
}

func unmarshalBaseAdminRuleClassification(rawMsg json.RawMessage) (BaseAdminRuleClassification, error) {
	if rawMsg == nil || string(rawMsg) == "null" {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b BaseAdminRuleClassification
	switch m["kind"] {
	case string(AdminRuleKindCustom):
		b = &AdminRule{}
	case string(AdminRuleKindDefault):
		b = &DefaultAdminRule{}
	default:
		b = &BaseAdminRule{}
	}
	if err := json.Unmarshal(rawMsg, b); err != nil {
		return nil, err
	}
	return b, nil
}

func unmarshalBaseAdminRuleClassificationArray(rawMsg json.RawMessage) ([]BaseAdminRuleClassification, error) {
	if rawMsg == nil || string(rawMsg) == "null" {
		return nil, nil
	}
	var rawMessages []json.RawMessage
	if err := json.Unmarshal(rawMsg, &rawMessages); err != nil {
		return nil, err
	}
	fArray := make([]BaseAdminRuleClassification, len(rawMessages))
	for index, rawMessage := range rawMessages {
		f, err := unmarshalBaseAdminRuleClassification(rawMessage)
		if err != nil {
			return nil, err
		}
		fArray[index] = f
	}
	return fArray, nil
}

func unmarshalEffectiveBaseSecurityAdminRuleClassification(rawMsg json.RawMessage) (EffectiveBaseSecurityAdminRuleClassification, error) {
	if rawMsg == nil || string(rawMsg) == "null" {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b EffectiveBaseSecurityAdminRuleClassification
	switch m["kind"] {
	case string(EffectiveAdminRuleKindCustom):
		b = &EffectiveSecurityAdminRule{}
	case string(EffectiveAdminRuleKindDefault):
		b = &EffectiveDefaultSecurityAdminRule{}
	default:
		b = &EffectiveBaseSecurityAdminRule{}
	}
	if err := json.Unmarshal(rawMsg, b); err != nil {
		return nil, err
	}
	return b, nil
}

func unmarshalEffectiveBaseSecurityAdminRuleClassificationArray(rawMsg json.RawMessage) ([]EffectiveBaseSecurityAdminRuleClassification, error) {
	if rawMsg == nil || string(rawMsg) == "null" {
		return nil, nil
	}
	var rawMessages []json.RawMessage
	if err := json.Unmarshal(rawMsg, &rawMessages); err != nil {
		return nil, err
	}
	fArray := make([]EffectiveBaseSecurityAdminRuleClassification, len(rawMessages))
	for index, rawMessage := range rawMessages {
		f, err := unmarshalEffectiveBaseSecurityAdminRuleClassification(rawMessage)
		if err != nil {
			return nil, err
		}
		fArray[index] = f
	}
	return fArray, nil
}

func unmarshalFirewallPolicyRuleClassification(rawMsg json.RawMessage) (FirewallPolicyRuleClassification, error) {
	if rawMsg == nil || string(rawMsg) == "null" {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b FirewallPolicyRuleClassification
	switch m["ruleType"] {
	case string(FirewallPolicyRuleTypeApplicationRule):
		b = &ApplicationRule{}
	case string(FirewallPolicyRuleTypeNatRule):
		b = &NatRule{}
	case string(FirewallPolicyRuleTypeNetworkRule):
		b = &Rule{}
	default:
		b = &FirewallPolicyRule{}
	}
	if err := json.Unmarshal(rawMsg, b); err != nil {
		return nil, err
	}
	return b, nil
}

func unmarshalFirewallPolicyRuleClassificationArray(rawMsg json.RawMessage) ([]FirewallPolicyRuleClassification, error) {
	if rawMsg == nil || string(rawMsg) == "null" {
		return nil, nil
	}
	var rawMessages []json.RawMessage
	if err := json.Unmarshal(rawMsg, &rawMessages); err != nil {
		return nil, err
	}
	fArray := make([]FirewallPolicyRuleClassification, len(rawMessages))
	for index, rawMessage := range rawMessages {
		f, err := unmarshalFirewallPolicyRuleClassification(rawMessage)
		if err != nil {
			return nil, err
		}
		fArray[index] = f
	}
	return fArray, nil
}

func unmarshalFirewallPolicyRuleCollectionClassification(rawMsg json.RawMessage) (FirewallPolicyRuleCollectionClassification, error) {
	if rawMsg == nil || string(rawMsg) == "null" {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b FirewallPolicyRuleCollectionClassification
	switch m["ruleCollectionType"] {
	case string(FirewallPolicyRuleCollectionTypeFirewallPolicyFilterRuleCollection):
		b = &FirewallPolicyFilterRuleCollection{}
	case string(FirewallPolicyRuleCollectionTypeFirewallPolicyNatRuleCollection):
		b = &FirewallPolicyNatRuleCollection{}
	default:
		b = &FirewallPolicyRuleCollection{}
	}
	if err := json.Unmarshal(rawMsg, b); err != nil {
		return nil, err
	}
	return b, nil
}

func unmarshalFirewallPolicyRuleCollectionClassificationArray(rawMsg json.RawMessage) ([]FirewallPolicyRuleCollectionClassification, error) {
	if rawMsg == nil || string(rawMsg) == "null" {
		return nil, nil
	}
	var rawMessages []json.RawMessage
	if err := json.Unmarshal(rawMsg, &rawMessages); err != nil {
		return nil, err
	}
	fArray := make([]FirewallPolicyRuleCollectionClassification, len(rawMessages))
	for index, rawMessage := range rawMessages {
		f, err := unmarshalFirewallPolicyRuleCollectionClassification(rawMessage)
		if err != nil {
			return nil, err
		}
		fArray[index] = f
	}
	return fArray, nil
}
