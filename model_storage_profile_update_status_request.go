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

// StorageProfileUpdateStatusRequest struct for StorageProfileUpdateStatusRequest
type StorageProfileUpdateStatusRequest struct {
	Status int32 `json:"status"`
	StorageId int32 `json:"storageId"`
}

// NewStorageProfileUpdateStatusRequest instantiates a new StorageProfileUpdateStatusRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStorageProfileUpdateStatusRequest(status int32, storageId int32) *StorageProfileUpdateStatusRequest {
	this := StorageProfileUpdateStatusRequest{}
	this.Status = status
	this.StorageId = storageId
	return &this
}

// NewStorageProfileUpdateStatusRequestWithDefaults instantiates a new StorageProfileUpdateStatusRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStorageProfileUpdateStatusRequestWithDefaults() *StorageProfileUpdateStatusRequest {
	this := StorageProfileUpdateStatusRequest{}
	return &this
}

// GetStatus returns the Status field value
func (o *StorageProfileUpdateStatusRequest) GetStatus() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *StorageProfileUpdateStatusRequest) GetStatusOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *StorageProfileUpdateStatusRequest) SetStatus(v int32) {
	o.Status = v
}

// GetStorageId returns the StorageId field value
func (o *StorageProfileUpdateStatusRequest) GetStorageId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.StorageId
}

// GetStorageIdOk returns a tuple with the StorageId field value
// and a boolean to check if the value has been set.
func (o *StorageProfileUpdateStatusRequest) GetStorageIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StorageId, true
}

// SetStorageId sets field value
func (o *StorageProfileUpdateStatusRequest) SetStorageId(v int32) {
	o.StorageId = v
}

func (o StorageProfileUpdateStatusRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["status"] = o.Status
	}
	if true {
		toSerialize["storageId"] = o.StorageId
	}
	return json.Marshal(toSerialize)
}

type NullableStorageProfileUpdateStatusRequest struct {
	value *StorageProfileUpdateStatusRequest
	isSet bool
}

func (v NullableStorageProfileUpdateStatusRequest) Get() *StorageProfileUpdateStatusRequest {
	return v.value
}

func (v *NullableStorageProfileUpdateStatusRequest) Set(val *StorageProfileUpdateStatusRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableStorageProfileUpdateStatusRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableStorageProfileUpdateStatusRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStorageProfileUpdateStatusRequest(val *StorageProfileUpdateStatusRequest) *NullableStorageProfileUpdateStatusRequest {
	return &NullableStorageProfileUpdateStatusRequest{value: val, isSet: true}
}

func (v NullableStorageProfileUpdateStatusRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStorageProfileUpdateStatusRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


