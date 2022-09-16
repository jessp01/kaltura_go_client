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

// AttachmentAssetSetContentRequest struct for AttachmentAssetSetContentRequest
type AttachmentAssetSetContentRequest struct {
	ContentResource KalturaContentResource `json:"contentResource"`
	Id string `json:"id"`
}

// NewAttachmentAssetSetContentRequest instantiates a new AttachmentAssetSetContentRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAttachmentAssetSetContentRequest(contentResource KalturaContentResource, id string) *AttachmentAssetSetContentRequest {
	this := AttachmentAssetSetContentRequest{}
	this.ContentResource = contentResource
	this.Id = id
	return &this
}

// NewAttachmentAssetSetContentRequestWithDefaults instantiates a new AttachmentAssetSetContentRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAttachmentAssetSetContentRequestWithDefaults() *AttachmentAssetSetContentRequest {
	this := AttachmentAssetSetContentRequest{}
	return &this
}

// GetContentResource returns the ContentResource field value
func (o *AttachmentAssetSetContentRequest) GetContentResource() KalturaContentResource {
	if o == nil {
		var ret KalturaContentResource
		return ret
	}

	return o.ContentResource
}

// GetContentResourceOk returns a tuple with the ContentResource field value
// and a boolean to check if the value has been set.
func (o *AttachmentAssetSetContentRequest) GetContentResourceOk() (*KalturaContentResource, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ContentResource, true
}

// SetContentResource sets field value
func (o *AttachmentAssetSetContentRequest) SetContentResource(v KalturaContentResource) {
	o.ContentResource = v
}

// GetId returns the Id field value
func (o *AttachmentAssetSetContentRequest) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AttachmentAssetSetContentRequest) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *AttachmentAssetSetContentRequest) SetId(v string) {
	o.Id = v
}

func (o AttachmentAssetSetContentRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["contentResource"] = o.ContentResource
	}
	if true {
		toSerialize["id"] = o.Id
	}
	return json.Marshal(toSerialize)
}

type NullableAttachmentAssetSetContentRequest struct {
	value *AttachmentAssetSetContentRequest
	isSet bool
}

func (v NullableAttachmentAssetSetContentRequest) Get() *AttachmentAssetSetContentRequest {
	return v.value
}

func (v *NullableAttachmentAssetSetContentRequest) Set(val *AttachmentAssetSetContentRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAttachmentAssetSetContentRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAttachmentAssetSetContentRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAttachmentAssetSetContentRequest(val *AttachmentAssetSetContentRequest) *NullableAttachmentAssetSetContentRequest {
	return &NullableAttachmentAssetSetContentRequest{value: val, isSet: true}
}

func (v NullableAttachmentAssetSetContentRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAttachmentAssetSetContentRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

