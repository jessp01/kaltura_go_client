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

// CaptionAssetUpdateRequest struct for CaptionAssetUpdateRequest
type CaptionAssetUpdateRequest struct {
	CaptionAsset KalturaCaptionAsset `json:"captionAsset"`
	Id string `json:"id"`
}

// NewCaptionAssetUpdateRequest instantiates a new CaptionAssetUpdateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCaptionAssetUpdateRequest(captionAsset KalturaCaptionAsset, id string) *CaptionAssetUpdateRequest {
	this := CaptionAssetUpdateRequest{}
	this.CaptionAsset = captionAsset
	this.Id = id
	return &this
}

// NewCaptionAssetUpdateRequestWithDefaults instantiates a new CaptionAssetUpdateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCaptionAssetUpdateRequestWithDefaults() *CaptionAssetUpdateRequest {
	this := CaptionAssetUpdateRequest{}
	return &this
}

// GetCaptionAsset returns the CaptionAsset field value
func (o *CaptionAssetUpdateRequest) GetCaptionAsset() KalturaCaptionAsset {
	if o == nil {
		var ret KalturaCaptionAsset
		return ret
	}

	return o.CaptionAsset
}

// GetCaptionAssetOk returns a tuple with the CaptionAsset field value
// and a boolean to check if the value has been set.
func (o *CaptionAssetUpdateRequest) GetCaptionAssetOk() (*KalturaCaptionAsset, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CaptionAsset, true
}

// SetCaptionAsset sets field value
func (o *CaptionAssetUpdateRequest) SetCaptionAsset(v KalturaCaptionAsset) {
	o.CaptionAsset = v
}

// GetId returns the Id field value
func (o *CaptionAssetUpdateRequest) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *CaptionAssetUpdateRequest) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *CaptionAssetUpdateRequest) SetId(v string) {
	o.Id = v
}

func (o CaptionAssetUpdateRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["captionAsset"] = o.CaptionAsset
	}
	if true {
		toSerialize["id"] = o.Id
	}
	return json.Marshal(toSerialize)
}

type NullableCaptionAssetUpdateRequest struct {
	value *CaptionAssetUpdateRequest
	isSet bool
}

func (v NullableCaptionAssetUpdateRequest) Get() *CaptionAssetUpdateRequest {
	return v.value
}

func (v *NullableCaptionAssetUpdateRequest) Set(val *CaptionAssetUpdateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCaptionAssetUpdateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCaptionAssetUpdateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCaptionAssetUpdateRequest(val *CaptionAssetUpdateRequest) *NullableCaptionAssetUpdateRequest {
	return &NullableCaptionAssetUpdateRequest{value: val, isSet: true}
}

func (v NullableCaptionAssetUpdateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCaptionAssetUpdateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

