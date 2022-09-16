/*
Kaltura VPaaS

The Kaltura VPaaS API

API version: 18.12.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// PermissionItemAddRequest struct for PermissionItemAddRequest
type PermissionItemAddRequest struct {
	PermissionItem KalturaPermissionItem `json:"permissionItem"`
}

// NewPermissionItemAddRequest instantiates a new PermissionItemAddRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPermissionItemAddRequest(permissionItem KalturaPermissionItem) *PermissionItemAddRequest {
	this := PermissionItemAddRequest{}
	this.PermissionItem = permissionItem
	return &this
}

// NewPermissionItemAddRequestWithDefaults instantiates a new PermissionItemAddRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPermissionItemAddRequestWithDefaults() *PermissionItemAddRequest {
	this := PermissionItemAddRequest{}
	return &this
}

// GetPermissionItem returns the PermissionItem field value
func (o *PermissionItemAddRequest) GetPermissionItem() KalturaPermissionItem {
	if o == nil {
		var ret KalturaPermissionItem
		return ret
	}

	return o.PermissionItem
}

// GetPermissionItemOk returns a tuple with the PermissionItem field value
// and a boolean to check if the value has been set.
func (o *PermissionItemAddRequest) GetPermissionItemOk() (*KalturaPermissionItem, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PermissionItem, true
}

// SetPermissionItem sets field value
func (o *PermissionItemAddRequest) SetPermissionItem(v KalturaPermissionItem) {
	o.PermissionItem = v
}

func (o PermissionItemAddRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["permissionItem"] = o.PermissionItem
	}
	return json.Marshal(toSerialize)
}

type NullablePermissionItemAddRequest struct {
	value *PermissionItemAddRequest
	isSet bool
}

func (v NullablePermissionItemAddRequest) Get() *PermissionItemAddRequest {
	return v.value
}

func (v *NullablePermissionItemAddRequest) Set(val *PermissionItemAddRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePermissionItemAddRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePermissionItemAddRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePermissionItemAddRequest(val *PermissionItemAddRequest) *NullablePermissionItemAddRequest {
	return &NullablePermissionItemAddRequest{value: val, isSet: true}
}

func (v NullablePermissionItemAddRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePermissionItemAddRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


