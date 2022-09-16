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

// PermissionItemDeleteRequest struct for PermissionItemDeleteRequest
type PermissionItemDeleteRequest struct {
	PermissionItemId int32 `json:"permissionItemId"`
}

// NewPermissionItemDeleteRequest instantiates a new PermissionItemDeleteRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPermissionItemDeleteRequest(permissionItemId int32) *PermissionItemDeleteRequest {
	this := PermissionItemDeleteRequest{}
	this.PermissionItemId = permissionItemId
	return &this
}

// NewPermissionItemDeleteRequestWithDefaults instantiates a new PermissionItemDeleteRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPermissionItemDeleteRequestWithDefaults() *PermissionItemDeleteRequest {
	this := PermissionItemDeleteRequest{}
	return &this
}

// GetPermissionItemId returns the PermissionItemId field value
func (o *PermissionItemDeleteRequest) GetPermissionItemId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PermissionItemId
}

// GetPermissionItemIdOk returns a tuple with the PermissionItemId field value
// and a boolean to check if the value has been set.
func (o *PermissionItemDeleteRequest) GetPermissionItemIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PermissionItemId, true
}

// SetPermissionItemId sets field value
func (o *PermissionItemDeleteRequest) SetPermissionItemId(v int32) {
	o.PermissionItemId = v
}

func (o PermissionItemDeleteRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["permissionItemId"] = o.PermissionItemId
	}
	return json.Marshal(toSerialize)
}

type NullablePermissionItemDeleteRequest struct {
	value *PermissionItemDeleteRequest
	isSet bool
}

func (v NullablePermissionItemDeleteRequest) Get() *PermissionItemDeleteRequest {
	return v.value
}

func (v *NullablePermissionItemDeleteRequest) Set(val *PermissionItemDeleteRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePermissionItemDeleteRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePermissionItemDeleteRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePermissionItemDeleteRequest(val *PermissionItemDeleteRequest) *NullablePermissionItemDeleteRequest {
	return &NullablePermissionItemDeleteRequest{value: val, isSet: true}
}

func (v NullablePermissionItemDeleteRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePermissionItemDeleteRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


