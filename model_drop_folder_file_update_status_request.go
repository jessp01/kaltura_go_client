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

// DropFolderFileUpdateStatusRequest struct for DropFolderFileUpdateStatusRequest
type DropFolderFileUpdateStatusRequest struct {
	DropFolderFileId int32 `json:"dropFolderFileId"`
	Status int32 `json:"status"`
}

// NewDropFolderFileUpdateStatusRequest instantiates a new DropFolderFileUpdateStatusRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDropFolderFileUpdateStatusRequest(dropFolderFileId int32, status int32) *DropFolderFileUpdateStatusRequest {
	this := DropFolderFileUpdateStatusRequest{}
	this.DropFolderFileId = dropFolderFileId
	this.Status = status
	return &this
}

// NewDropFolderFileUpdateStatusRequestWithDefaults instantiates a new DropFolderFileUpdateStatusRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDropFolderFileUpdateStatusRequestWithDefaults() *DropFolderFileUpdateStatusRequest {
	this := DropFolderFileUpdateStatusRequest{}
	return &this
}

// GetDropFolderFileId returns the DropFolderFileId field value
func (o *DropFolderFileUpdateStatusRequest) GetDropFolderFileId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.DropFolderFileId
}

// GetDropFolderFileIdOk returns a tuple with the DropFolderFileId field value
// and a boolean to check if the value has been set.
func (o *DropFolderFileUpdateStatusRequest) GetDropFolderFileIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DropFolderFileId, true
}

// SetDropFolderFileId sets field value
func (o *DropFolderFileUpdateStatusRequest) SetDropFolderFileId(v int32) {
	o.DropFolderFileId = v
}

// GetStatus returns the Status field value
func (o *DropFolderFileUpdateStatusRequest) GetStatus() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *DropFolderFileUpdateStatusRequest) GetStatusOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *DropFolderFileUpdateStatusRequest) SetStatus(v int32) {
	o.Status = v
}

func (o DropFolderFileUpdateStatusRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["dropFolderFileId"] = o.DropFolderFileId
	}
	if true {
		toSerialize["status"] = o.Status
	}
	return json.Marshal(toSerialize)
}

type NullableDropFolderFileUpdateStatusRequest struct {
	value *DropFolderFileUpdateStatusRequest
	isSet bool
}

func (v NullableDropFolderFileUpdateStatusRequest) Get() *DropFolderFileUpdateStatusRequest {
	return v.value
}

func (v *NullableDropFolderFileUpdateStatusRequest) Set(val *DropFolderFileUpdateStatusRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableDropFolderFileUpdateStatusRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableDropFolderFileUpdateStatusRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDropFolderFileUpdateStatusRequest(val *DropFolderFileUpdateStatusRequest) *NullableDropFolderFileUpdateStatusRequest {
	return &NullableDropFolderFileUpdateStatusRequest{value: val, isSet: true}
}

func (v NullableDropFolderFileUpdateStatusRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDropFolderFileUpdateStatusRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


