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

// AttachmentAssetAddRequest struct for AttachmentAssetAddRequest
type AttachmentAssetAddRequest struct {
	AttachmentAsset KalturaAttachmentAsset `json:"attachmentAsset"`
	EntryId string `json:"entryId"`
}

// NewAttachmentAssetAddRequest instantiates a new AttachmentAssetAddRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAttachmentAssetAddRequest(attachmentAsset KalturaAttachmentAsset, entryId string) *AttachmentAssetAddRequest {
	this := AttachmentAssetAddRequest{}
	this.AttachmentAsset = attachmentAsset
	this.EntryId = entryId
	return &this
}

// NewAttachmentAssetAddRequestWithDefaults instantiates a new AttachmentAssetAddRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAttachmentAssetAddRequestWithDefaults() *AttachmentAssetAddRequest {
	this := AttachmentAssetAddRequest{}
	return &this
}

// GetAttachmentAsset returns the AttachmentAsset field value
func (o *AttachmentAssetAddRequest) GetAttachmentAsset() KalturaAttachmentAsset {
	if o == nil {
		var ret KalturaAttachmentAsset
		return ret
	}

	return o.AttachmentAsset
}

// GetAttachmentAssetOk returns a tuple with the AttachmentAsset field value
// and a boolean to check if the value has been set.
func (o *AttachmentAssetAddRequest) GetAttachmentAssetOk() (*KalturaAttachmentAsset, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AttachmentAsset, true
}

// SetAttachmentAsset sets field value
func (o *AttachmentAssetAddRequest) SetAttachmentAsset(v KalturaAttachmentAsset) {
	o.AttachmentAsset = v
}

// GetEntryId returns the EntryId field value
func (o *AttachmentAssetAddRequest) GetEntryId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EntryId
}

// GetEntryIdOk returns a tuple with the EntryId field value
// and a boolean to check if the value has been set.
func (o *AttachmentAssetAddRequest) GetEntryIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EntryId, true
}

// SetEntryId sets field value
func (o *AttachmentAssetAddRequest) SetEntryId(v string) {
	o.EntryId = v
}

func (o AttachmentAssetAddRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["attachmentAsset"] = o.AttachmentAsset
	}
	if true {
		toSerialize["entryId"] = o.EntryId
	}
	return json.Marshal(toSerialize)
}

type NullableAttachmentAssetAddRequest struct {
	value *AttachmentAssetAddRequest
	isSet bool
}

func (v NullableAttachmentAssetAddRequest) Get() *AttachmentAssetAddRequest {
	return v.value
}

func (v *NullableAttachmentAssetAddRequest) Set(val *AttachmentAssetAddRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAttachmentAssetAddRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAttachmentAssetAddRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAttachmentAssetAddRequest(val *AttachmentAssetAddRequest) *NullableAttachmentAssetAddRequest {
	return &NullableAttachmentAssetAddRequest{value: val, isSet: true}
}

func (v NullableAttachmentAssetAddRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAttachmentAssetAddRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

