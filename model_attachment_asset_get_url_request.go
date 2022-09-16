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

// AttachmentAssetGetUrlRequest struct for AttachmentAssetGetUrlRequest
type AttachmentAssetGetUrlRequest struct {
	Id string `json:"id"`
	StorageId *int32 `json:"storageId,omitempty"`
}

// NewAttachmentAssetGetUrlRequest instantiates a new AttachmentAssetGetUrlRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAttachmentAssetGetUrlRequest(id string) *AttachmentAssetGetUrlRequest {
	this := AttachmentAssetGetUrlRequest{}
	this.Id = id
	return &this
}

// NewAttachmentAssetGetUrlRequestWithDefaults instantiates a new AttachmentAssetGetUrlRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAttachmentAssetGetUrlRequestWithDefaults() *AttachmentAssetGetUrlRequest {
	this := AttachmentAssetGetUrlRequest{}
	return &this
}

// GetId returns the Id field value
func (o *AttachmentAssetGetUrlRequest) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AttachmentAssetGetUrlRequest) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *AttachmentAssetGetUrlRequest) SetId(v string) {
	o.Id = v
}

// GetStorageId returns the StorageId field value if set, zero value otherwise.
func (o *AttachmentAssetGetUrlRequest) GetStorageId() int32 {
	if o == nil || o.StorageId == nil {
		var ret int32
		return ret
	}
	return *o.StorageId
}

// GetStorageIdOk returns a tuple with the StorageId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AttachmentAssetGetUrlRequest) GetStorageIdOk() (*int32, bool) {
	if o == nil || o.StorageId == nil {
		return nil, false
	}
	return o.StorageId, true
}

// HasStorageId returns a boolean if a field has been set.
func (o *AttachmentAssetGetUrlRequest) HasStorageId() bool {
	if o != nil && o.StorageId != nil {
		return true
	}

	return false
}

// SetStorageId gets a reference to the given int32 and assigns it to the StorageId field.
func (o *AttachmentAssetGetUrlRequest) SetStorageId(v int32) {
	o.StorageId = &v
}

func (o AttachmentAssetGetUrlRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if o.StorageId != nil {
		toSerialize["storageId"] = o.StorageId
	}
	return json.Marshal(toSerialize)
}

type NullableAttachmentAssetGetUrlRequest struct {
	value *AttachmentAssetGetUrlRequest
	isSet bool
}

func (v NullableAttachmentAssetGetUrlRequest) Get() *AttachmentAssetGetUrlRequest {
	return v.value
}

func (v *NullableAttachmentAssetGetUrlRequest) Set(val *AttachmentAssetGetUrlRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAttachmentAssetGetUrlRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAttachmentAssetGetUrlRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAttachmentAssetGetUrlRequest(val *AttachmentAssetGetUrlRequest) *NullableAttachmentAssetGetUrlRequest {
	return &NullableAttachmentAssetGetUrlRequest{value: val, isSet: true}
}

func (v NullableAttachmentAssetGetUrlRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAttachmentAssetGetUrlRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


