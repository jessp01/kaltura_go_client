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

// DropFolderDeleteRequest struct for DropFolderDeleteRequest
type DropFolderDeleteRequest struct {
	DropFolderId int32 `json:"dropFolderId"`
}

// NewDropFolderDeleteRequest instantiates a new DropFolderDeleteRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDropFolderDeleteRequest(dropFolderId int32) *DropFolderDeleteRequest {
	this := DropFolderDeleteRequest{}
	this.DropFolderId = dropFolderId
	return &this
}

// NewDropFolderDeleteRequestWithDefaults instantiates a new DropFolderDeleteRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDropFolderDeleteRequestWithDefaults() *DropFolderDeleteRequest {
	this := DropFolderDeleteRequest{}
	return &this
}

// GetDropFolderId returns the DropFolderId field value
func (o *DropFolderDeleteRequest) GetDropFolderId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.DropFolderId
}

// GetDropFolderIdOk returns a tuple with the DropFolderId field value
// and a boolean to check if the value has been set.
func (o *DropFolderDeleteRequest) GetDropFolderIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DropFolderId, true
}

// SetDropFolderId sets field value
func (o *DropFolderDeleteRequest) SetDropFolderId(v int32) {
	o.DropFolderId = v
}

func (o DropFolderDeleteRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["dropFolderId"] = o.DropFolderId
	}
	return json.Marshal(toSerialize)
}

type NullableDropFolderDeleteRequest struct {
	value *DropFolderDeleteRequest
	isSet bool
}

func (v NullableDropFolderDeleteRequest) Get() *DropFolderDeleteRequest {
	return v.value
}

func (v *NullableDropFolderDeleteRequest) Set(val *DropFolderDeleteRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableDropFolderDeleteRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableDropFolderDeleteRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDropFolderDeleteRequest(val *DropFolderDeleteRequest) *NullableDropFolderDeleteRequest {
	return &NullableDropFolderDeleteRequest{value: val, isSet: true}
}

func (v NullableDropFolderDeleteRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDropFolderDeleteRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

