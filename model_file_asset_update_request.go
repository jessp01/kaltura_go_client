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

// FileAssetUpdateRequest struct for FileAssetUpdateRequest
type FileAssetUpdateRequest struct {
	FileAsset KalturaFileAsset `json:"fileAsset"`
	Id int32 `json:"id"`
}

// NewFileAssetUpdateRequest instantiates a new FileAssetUpdateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFileAssetUpdateRequest(fileAsset KalturaFileAsset, id int32) *FileAssetUpdateRequest {
	this := FileAssetUpdateRequest{}
	this.FileAsset = fileAsset
	this.Id = id
	return &this
}

// NewFileAssetUpdateRequestWithDefaults instantiates a new FileAssetUpdateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFileAssetUpdateRequestWithDefaults() *FileAssetUpdateRequest {
	this := FileAssetUpdateRequest{}
	return &this
}

// GetFileAsset returns the FileAsset field value
func (o *FileAssetUpdateRequest) GetFileAsset() KalturaFileAsset {
	if o == nil {
		var ret KalturaFileAsset
		return ret
	}

	return o.FileAsset
}

// GetFileAssetOk returns a tuple with the FileAsset field value
// and a boolean to check if the value has been set.
func (o *FileAssetUpdateRequest) GetFileAssetOk() (*KalturaFileAsset, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FileAsset, true
}

// SetFileAsset sets field value
func (o *FileAssetUpdateRequest) SetFileAsset(v KalturaFileAsset) {
	o.FileAsset = v
}

// GetId returns the Id field value
func (o *FileAssetUpdateRequest) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *FileAssetUpdateRequest) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *FileAssetUpdateRequest) SetId(v int32) {
	o.Id = v
}

func (o FileAssetUpdateRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["fileAsset"] = o.FileAsset
	}
	if true {
		toSerialize["id"] = o.Id
	}
	return json.Marshal(toSerialize)
}

type NullableFileAssetUpdateRequest struct {
	value *FileAssetUpdateRequest
	isSet bool
}

func (v NullableFileAssetUpdateRequest) Get() *FileAssetUpdateRequest {
	return v.value
}

func (v *NullableFileAssetUpdateRequest) Set(val *FileAssetUpdateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableFileAssetUpdateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableFileAssetUpdateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFileAssetUpdateRequest(val *FileAssetUpdateRequest) *NullableFileAssetUpdateRequest {
	return &NullableFileAssetUpdateRequest{value: val, isSet: true}
}

func (v NullableFileAssetUpdateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFileAssetUpdateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

