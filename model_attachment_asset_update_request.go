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

// AttachmentAssetUpdateRequest struct for AttachmentAssetUpdateRequest
type AttachmentAssetUpdateRequest struct {
	AttachmentAsset KalturaAttachmentAsset `json:"attachmentAsset"`
	Id string `json:"id"`
}

// NewAttachmentAssetUpdateRequest instantiates a new AttachmentAssetUpdateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAttachmentAssetUpdateRequest(attachmentAsset KalturaAttachmentAsset, id string) *AttachmentAssetUpdateRequest {
	this := AttachmentAssetUpdateRequest{}
	this.AttachmentAsset = attachmentAsset
	this.Id = id
	return &this
}

// NewAttachmentAssetUpdateRequestWithDefaults instantiates a new AttachmentAssetUpdateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAttachmentAssetUpdateRequestWithDefaults() *AttachmentAssetUpdateRequest {
	this := AttachmentAssetUpdateRequest{}
	return &this
}

// GetAttachmentAsset returns the AttachmentAsset field value
func (o *AttachmentAssetUpdateRequest) GetAttachmentAsset() KalturaAttachmentAsset {
	if o == nil {
		var ret KalturaAttachmentAsset
		return ret
	}

	return o.AttachmentAsset
}

// GetAttachmentAssetOk returns a tuple with the AttachmentAsset field value
// and a boolean to check if the value has been set.
func (o *AttachmentAssetUpdateRequest) GetAttachmentAssetOk() (*KalturaAttachmentAsset, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AttachmentAsset, true
}

// SetAttachmentAsset sets field value
func (o *AttachmentAssetUpdateRequest) SetAttachmentAsset(v KalturaAttachmentAsset) {
	o.AttachmentAsset = v
}

// GetId returns the Id field value
func (o *AttachmentAssetUpdateRequest) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AttachmentAssetUpdateRequest) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *AttachmentAssetUpdateRequest) SetId(v string) {
	o.Id = v
}

func (o AttachmentAssetUpdateRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["attachmentAsset"] = o.AttachmentAsset
	}
	if true {
		toSerialize["id"] = o.Id
	}
	return json.Marshal(toSerialize)
}

type NullableAttachmentAssetUpdateRequest struct {
	value *AttachmentAssetUpdateRequest
	isSet bool
}

func (v NullableAttachmentAssetUpdateRequest) Get() *AttachmentAssetUpdateRequest {
	return v.value
}

func (v *NullableAttachmentAssetUpdateRequest) Set(val *AttachmentAssetUpdateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAttachmentAssetUpdateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAttachmentAssetUpdateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAttachmentAssetUpdateRequest(val *AttachmentAssetUpdateRequest) *NullableAttachmentAssetUpdateRequest {
	return &NullableAttachmentAssetUpdateRequest{value: val, isSet: true}
}

func (v NullableAttachmentAssetUpdateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAttachmentAssetUpdateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


