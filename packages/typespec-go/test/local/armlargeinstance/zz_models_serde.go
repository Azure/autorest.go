// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package armlargeinstance

import (
	"encoding/json"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"reflect"
)

// MarshalJSON implements the json.Marshaller interface for type AzureLargeInstance.
func (a AzureLargeInstance) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "id", a.ID)
	populate(objectMap, "location", a.Location)
	populate(objectMap, "name", a.Name)
	populate(objectMap, "properties", a.Properties)
	populate(objectMap, "systemData", a.SystemData)
	populate(objectMap, "tags", a.Tags)
	populate(objectMap, "type", a.Type)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type AzureLargeInstance.
func (a *AzureLargeInstance) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", a, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "id":
			err = unpopulate(val, "ID", &a.ID)
			delete(rawMsg, key)
		case "location":
			err = unpopulate(val, "Location", &a.Location)
			delete(rawMsg, key)
		case "name":
			err = unpopulate(val, "Name", &a.Name)
			delete(rawMsg, key)
		case "properties":
			err = unpopulate(val, "Properties", &a.Properties)
			delete(rawMsg, key)
		case "systemData":
			err = unpopulate(val, "SystemData", &a.SystemData)
			delete(rawMsg, key)
		case "tags":
			err = unpopulate(val, "Tags", &a.Tags)
			delete(rawMsg, key)
		case "type":
			err = unpopulate(val, "Type", &a.Type)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", a, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type AzureLargeStorageInstance.
func (a AzureLargeStorageInstance) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "id", a.ID)
	populate(objectMap, "identity", a.Identity)
	populate(objectMap, "location", a.Location)
	populate(objectMap, "name", a.Name)
	populate(objectMap, "properties", a.Properties)
	populate(objectMap, "systemData", a.SystemData)
	populate(objectMap, "tags", a.Tags)
	populate(objectMap, "type", a.Type)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type AzureLargeStorageInstance.
func (a *AzureLargeStorageInstance) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", a, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "id":
			err = unpopulate(val, "ID", &a.ID)
			delete(rawMsg, key)
		case "identity":
			err = unpopulate(val, "Identity", &a.Identity)
			delete(rawMsg, key)
		case "location":
			err = unpopulate(val, "Location", &a.Location)
			delete(rawMsg, key)
		case "name":
			err = unpopulate(val, "Name", &a.Name)
			delete(rawMsg, key)
		case "properties":
			err = unpopulate(val, "Properties", &a.Properties)
			delete(rawMsg, key)
		case "systemData":
			err = unpopulate(val, "SystemData", &a.SystemData)
			delete(rawMsg, key)
		case "tags":
			err = unpopulate(val, "Tags", &a.Tags)
			delete(rawMsg, key)
		case "type":
			err = unpopulate(val, "Type", &a.Type)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", a, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type AzureLargeStorageInstanceListResult.
func (a AzureLargeStorageInstanceListResult) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "nextLink", a.NextLink)
	populate(objectMap, "value", a.Value)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type AzureLargeStorageInstanceListResult.
func (a *AzureLargeStorageInstanceListResult) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", a, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "nextLink":
			err = unpopulate(val, "NextLink", &a.NextLink)
			delete(rawMsg, key)
		case "value":
			err = unpopulate(val, "Value", &a.Value)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", a, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type AzureLargeStorageInstanceProperties.
func (a AzureLargeStorageInstanceProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "azureLargeStorageInstanceUniqueIdentifier", a.AzureLargeStorageInstanceUniqueIdentifier)
	populate(objectMap, "storageProperties", a.StorageProperties)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type AzureLargeStorageInstanceProperties.
func (a *AzureLargeStorageInstanceProperties) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", a, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "azureLargeStorageInstanceUniqueIdentifier":
			err = unpopulate(val, "AzureLargeStorageInstanceUniqueIdentifier", &a.AzureLargeStorageInstanceUniqueIdentifier)
			delete(rawMsg, key)
		case "storageProperties":
			err = unpopulate(val, "StorageProperties", &a.StorageProperties)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", a, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type AzureLargeStorageInstanceTagsUpdate.
func (a AzureLargeStorageInstanceTagsUpdate) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "identity", a.Identity)
	populate(objectMap, "tags", a.Tags)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type AzureLargeStorageInstanceTagsUpdate.
func (a *AzureLargeStorageInstanceTagsUpdate) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", a, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "identity":
			err = unpopulate(val, "Identity", &a.Identity)
			delete(rawMsg, key)
		case "tags":
			err = unpopulate(val, "Tags", &a.Tags)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", a, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type Disk.
func (d Disk) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "diskSizeGB", d.DiskSizeGB)
	populate(objectMap, "lun", d.Lun)
	populate(objectMap, "name", d.Name)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type Disk.
func (d *Disk) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", d, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "diskSizeGB":
			err = unpopulate(val, "DiskSizeGB", &d.DiskSizeGB)
			delete(rawMsg, key)
		case "lun":
			err = unpopulate(val, "Lun", &d.Lun)
			delete(rawMsg, key)
		case "name":
			err = unpopulate(val, "Name", &d.Name)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", d, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type ErrorAdditionalInfo.
func (e ErrorAdditionalInfo) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populateAny(objectMap, "info", e.Info)
	populate(objectMap, "type", e.Type)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type ErrorAdditionalInfo.
func (e *ErrorAdditionalInfo) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", e, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "info":
			err = unpopulate(val, "Info", &e.Info)
			delete(rawMsg, key)
		case "type":
			err = unpopulate(val, "Type", &e.Type)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", e, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type ErrorDetail.
func (e ErrorDetail) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "additionalInfo", e.AdditionalInfo)
	populate(objectMap, "code", e.Code)
	populate(objectMap, "details", e.Details)
	populate(objectMap, "message", e.Message)
	populate(objectMap, "target", e.Target)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type ErrorDetail.
func (e *ErrorDetail) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", e, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "additionalInfo":
			err = unpopulate(val, "AdditionalInfo", &e.AdditionalInfo)
			delete(rawMsg, key)
		case "code":
			err = unpopulate(val, "Code", &e.Code)
			delete(rawMsg, key)
		case "details":
			err = unpopulate(val, "Details", &e.Details)
			delete(rawMsg, key)
		case "message":
			err = unpopulate(val, "Message", &e.Message)
			delete(rawMsg, key)
		case "target":
			err = unpopulate(val, "Target", &e.Target)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", e, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type ForceState.
func (f ForceState) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "forceState", f.ForceState)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type ForceState.
func (f *ForceState) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", f, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "forceState":
			err = unpopulate(val, "ForceState", &f.ForceState)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", f, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type HardwareProfile.
func (h HardwareProfile) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "azureLargeInstanceSize", h.AzureLargeInstanceSize)
	populate(objectMap, "hardwareType", h.HardwareType)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type HardwareProfile.
func (h *HardwareProfile) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", h, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "azureLargeInstanceSize":
			err = unpopulate(val, "AzureLargeInstanceSize", &h.AzureLargeInstanceSize)
			delete(rawMsg, key)
		case "hardwareType":
			err = unpopulate(val, "HardwareType", &h.HardwareType)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", h, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type IPAddress.
func (i IPAddress) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "ipAddress", i.IPAddress)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type IPAddress.
func (i *IPAddress) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", i, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "ipAddress":
			err = unpopulate(val, "IPAddress", &i.IPAddress)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", i, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type ListResult.
func (l ListResult) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "nextLink", l.NextLink)
	populate(objectMap, "value", l.Value)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type ListResult.
func (l *ListResult) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", l, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "nextLink":
			err = unpopulate(val, "NextLink", &l.NextLink)
			delete(rawMsg, key)
		case "value":
			err = unpopulate(val, "Value", &l.Value)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", l, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type ManagedServiceIdentity.
func (m ManagedServiceIdentity) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "principalId", m.PrincipalID)
	populate(objectMap, "tenantId", m.TenantID)
	populate(objectMap, "type", m.Type)
	populate(objectMap, "userAssignedIdentities", m.UserAssignedIdentities)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type ManagedServiceIdentity.
func (m *ManagedServiceIdentity) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", m, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "principalId":
			err = unpopulate(val, "PrincipalID", &m.PrincipalID)
			delete(rawMsg, key)
		case "tenantId":
			err = unpopulate(val, "TenantID", &m.TenantID)
			delete(rawMsg, key)
		case "type":
			err = unpopulate(val, "Type", &m.Type)
			delete(rawMsg, key)
		case "userAssignedIdentities":
			err = unpopulate(val, "UserAssignedIdentities", &m.UserAssignedIdentities)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", m, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type NetworkProfile.
func (n NetworkProfile) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "circuitId", n.CircuitID)
	populate(objectMap, "networkInterfaces", n.NetworkInterfaces)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type NetworkProfile.
func (n *NetworkProfile) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", n, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "circuitId":
			err = unpopulate(val, "CircuitID", &n.CircuitID)
			delete(rawMsg, key)
		case "networkInterfaces":
			err = unpopulate(val, "NetworkInterfaces", &n.NetworkInterfaces)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", n, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type Operation.
func (o Operation) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "actionType", o.ActionType)
	populate(objectMap, "display", o.Display)
	populate(objectMap, "isDataAction", o.IsDataAction)
	populate(objectMap, "name", o.Name)
	populate(objectMap, "origin", o.Origin)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type Operation.
func (o *Operation) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", o, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "actionType":
			err = unpopulate(val, "ActionType", &o.ActionType)
			delete(rawMsg, key)
		case "display":
			err = unpopulate(val, "Display", &o.Display)
			delete(rawMsg, key)
		case "isDataAction":
			err = unpopulate(val, "IsDataAction", &o.IsDataAction)
			delete(rawMsg, key)
		case "name":
			err = unpopulate(val, "Name", &o.Name)
			delete(rawMsg, key)
		case "origin":
			err = unpopulate(val, "Origin", &o.Origin)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", o, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type OperationDisplay.
func (o OperationDisplay) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "description", o.Description)
	populate(objectMap, "operation", o.Operation)
	populate(objectMap, "provider", o.Provider)
	populate(objectMap, "resource", o.Resource)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type OperationDisplay.
func (o *OperationDisplay) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", o, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "description":
			err = unpopulate(val, "Description", &o.Description)
			delete(rawMsg, key)
		case "operation":
			err = unpopulate(val, "Operation", &o.Operation)
			delete(rawMsg, key)
		case "provider":
			err = unpopulate(val, "Provider", &o.Provider)
			delete(rawMsg, key)
		case "resource":
			err = unpopulate(val, "Resource", &o.Resource)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", o, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type OperationListResult.
func (o OperationListResult) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "nextLink", o.NextLink)
	populate(objectMap, "value", o.Value)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type OperationListResult.
func (o *OperationListResult) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", o, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "nextLink":
			err = unpopulate(val, "NextLink", &o.NextLink)
			delete(rawMsg, key)
		case "value":
			err = unpopulate(val, "Value", &o.Value)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", o, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type OperationStatusResult.
func (o OperationStatusResult) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populateDateTimeRFC3339(objectMap, "endTime", o.EndTime)
	populate(objectMap, "error", o.Error)
	populate(objectMap, "id", o.ID)
	populate(objectMap, "name", o.Name)
	populate(objectMap, "operations", o.Operations)
	populate(objectMap, "percentComplete", o.PercentComplete)
	populate(objectMap, "resourceId", o.ResourceID)
	populateDateTimeRFC3339(objectMap, "startTime", o.StartTime)
	populate(objectMap, "status", o.Status)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type OperationStatusResult.
func (o *OperationStatusResult) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", o, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "endTime":
			err = unpopulateDateTimeRFC3339(val, "EndTime", &o.EndTime)
			delete(rawMsg, key)
		case "error":
			err = unpopulate(val, "Error", &o.Error)
			delete(rawMsg, key)
		case "id":
			err = unpopulate(val, "ID", &o.ID)
			delete(rawMsg, key)
		case "name":
			err = unpopulate(val, "Name", &o.Name)
			delete(rawMsg, key)
		case "operations":
			err = unpopulate(val, "Operations", &o.Operations)
			delete(rawMsg, key)
		case "percentComplete":
			err = unpopulate(val, "PercentComplete", &o.PercentComplete)
			delete(rawMsg, key)
		case "resourceId":
			err = unpopulate(val, "ResourceID", &o.ResourceID)
			delete(rawMsg, key)
		case "startTime":
			err = unpopulateDateTimeRFC3339(val, "StartTime", &o.StartTime)
			delete(rawMsg, key)
		case "status":
			err = unpopulate(val, "Status", &o.Status)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", o, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type OsProfile.
func (o OsProfile) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "computerName", o.ComputerName)
	populate(objectMap, "osType", o.OSType)
	populate(objectMap, "sshPublicKey", o.SSHPublicKey)
	populate(objectMap, "version", o.Version)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type OsProfile.
func (o *OsProfile) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", o, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "computerName":
			err = unpopulate(val, "ComputerName", &o.ComputerName)
			delete(rawMsg, key)
		case "osType":
			err = unpopulate(val, "OSType", &o.OSType)
			delete(rawMsg, key)
		case "sshPublicKey":
			err = unpopulate(val, "SSHPublicKey", &o.SSHPublicKey)
			delete(rawMsg, key)
		case "version":
			err = unpopulate(val, "Version", &o.Version)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", o, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type Properties.
func (p Properties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "azureLargeInstanceId", p.AzureLargeInstanceID)
	populate(objectMap, "hardwareProfile", p.HardwareProfile)
	populate(objectMap, "hwRevision", p.HwRevision)
	populate(objectMap, "networkProfile", p.NetworkProfile)
	populate(objectMap, "osProfile", p.OSProfile)
	populate(objectMap, "powerState", p.PowerState)
	populate(objectMap, "provisioningState", p.ProvisioningState)
	populate(objectMap, "proximityPlacementGroup", p.ProximityPlacementGroup)
	populate(objectMap, "storageProfile", p.StorageProfile)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type Properties.
func (p *Properties) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", p, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "azureLargeInstanceId":
			err = unpopulate(val, "AzureLargeInstanceID", &p.AzureLargeInstanceID)
			delete(rawMsg, key)
		case "hardwareProfile":
			err = unpopulate(val, "HardwareProfile", &p.HardwareProfile)
			delete(rawMsg, key)
		case "hwRevision":
			err = unpopulate(val, "HwRevision", &p.HwRevision)
			delete(rawMsg, key)
		case "networkProfile":
			err = unpopulate(val, "NetworkProfile", &p.NetworkProfile)
			delete(rawMsg, key)
		case "osProfile":
			err = unpopulate(val, "OSProfile", &p.OSProfile)
			delete(rawMsg, key)
		case "powerState":
			err = unpopulate(val, "PowerState", &p.PowerState)
			delete(rawMsg, key)
		case "provisioningState":
			err = unpopulate(val, "ProvisioningState", &p.ProvisioningState)
			delete(rawMsg, key)
		case "proximityPlacementGroup":
			err = unpopulate(val, "ProximityPlacementGroup", &p.ProximityPlacementGroup)
			delete(rawMsg, key)
		case "storageProfile":
			err = unpopulate(val, "StorageProfile", &p.StorageProfile)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", p, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type StorageBillingProperties.
func (s StorageBillingProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "billingMode", s.BillingMode)
	populate(objectMap, "sku", s.SKU)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type StorageBillingProperties.
func (s *StorageBillingProperties) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", s, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "billingMode":
			err = unpopulate(val, "BillingMode", &s.BillingMode)
			delete(rawMsg, key)
		case "sku":
			err = unpopulate(val, "SKU", &s.SKU)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", s, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type StorageProfile.
func (s StorageProfile) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "nfsIpAddress", s.NfsIPAddress)
	populate(objectMap, "osDisks", s.OSDisks)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type StorageProfile.
func (s *StorageProfile) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", s, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "nfsIpAddress":
			err = unpopulate(val, "NfsIPAddress", &s.NfsIPAddress)
			delete(rawMsg, key)
		case "osDisks":
			err = unpopulate(val, "OSDisks", &s.OSDisks)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", s, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type StorageProperties.
func (s StorageProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "generation", s.Generation)
	populate(objectMap, "hardwareType", s.HardwareType)
	populate(objectMap, "offeringType", s.OfferingType)
	populate(objectMap, "provisioningState", s.ProvisioningState)
	populate(objectMap, "storageBillingProperties", s.StorageBillingProperties)
	populate(objectMap, "storageType", s.StorageType)
	populate(objectMap, "workloadType", s.WorkloadType)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type StorageProperties.
func (s *StorageProperties) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", s, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "generation":
			err = unpopulate(val, "Generation", &s.Generation)
			delete(rawMsg, key)
		case "hardwareType":
			err = unpopulate(val, "HardwareType", &s.HardwareType)
			delete(rawMsg, key)
		case "offeringType":
			err = unpopulate(val, "OfferingType", &s.OfferingType)
			delete(rawMsg, key)
		case "provisioningState":
			err = unpopulate(val, "ProvisioningState", &s.ProvisioningState)
			delete(rawMsg, key)
		case "storageBillingProperties":
			err = unpopulate(val, "StorageBillingProperties", &s.StorageBillingProperties)
			delete(rawMsg, key)
		case "storageType":
			err = unpopulate(val, "StorageType", &s.StorageType)
			delete(rawMsg, key)
		case "workloadType":
			err = unpopulate(val, "WorkloadType", &s.WorkloadType)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", s, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type SystemData.
func (s SystemData) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populateDateTimeRFC3339(objectMap, "createdAt", s.CreatedAt)
	populate(objectMap, "createdBy", s.CreatedBy)
	populate(objectMap, "createdByType", s.CreatedByType)
	populateDateTimeRFC3339(objectMap, "lastModifiedAt", s.LastModifiedAt)
	populate(objectMap, "lastModifiedBy", s.LastModifiedBy)
	populate(objectMap, "lastModifiedByType", s.LastModifiedByType)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type SystemData.
func (s *SystemData) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", s, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "createdAt":
			err = unpopulateDateTimeRFC3339(val, "CreatedAt", &s.CreatedAt)
			delete(rawMsg, key)
		case "createdBy":
			err = unpopulate(val, "CreatedBy", &s.CreatedBy)
			delete(rawMsg, key)
		case "createdByType":
			err = unpopulate(val, "CreatedByType", &s.CreatedByType)
			delete(rawMsg, key)
		case "lastModifiedAt":
			err = unpopulateDateTimeRFC3339(val, "LastModifiedAt", &s.LastModifiedAt)
			delete(rawMsg, key)
		case "lastModifiedBy":
			err = unpopulate(val, "LastModifiedBy", &s.LastModifiedBy)
			delete(rawMsg, key)
		case "lastModifiedByType":
			err = unpopulate(val, "LastModifiedByType", &s.LastModifiedByType)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", s, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type TagsUpdate.
func (t TagsUpdate) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "tags", t.Tags)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type TagsUpdate.
func (t *TagsUpdate) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", t, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "tags":
			err = unpopulate(val, "Tags", &t.Tags)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", t, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type UserAssignedIdentity.
func (u UserAssignedIdentity) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "clientId", u.ClientID)
	populate(objectMap, "principalId", u.PrincipalID)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type UserAssignedIdentity.
func (u *UserAssignedIdentity) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", u, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "clientId":
			err = unpopulate(val, "ClientID", &u.ClientID)
			delete(rawMsg, key)
		case "principalId":
			err = unpopulate(val, "PrincipalID", &u.PrincipalID)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", u, err)
		}
	}
	return nil
}

func populate(m map[string]any, k string, v any) {
	if v == nil {
		return
	} else if azcore.IsNullValue(v) {
		m[k] = nil
	} else if !reflect.ValueOf(v).IsNil() {
		m[k] = v
	}
}

func populateAny(m map[string]any, k string, v any) {
	if v == nil {
		return
	} else if azcore.IsNullValue(v) {
		m[k] = nil
	} else {
		m[k] = v
	}
}

func unpopulate(data json.RawMessage, fn string, v any) error {
	if data == nil || string(data) == "null" {
		return nil
	}
	if err := json.Unmarshal(data, v); err != nil {
		return fmt.Errorf("struct field %s: %v", fn, err)
	}
	return nil
}
