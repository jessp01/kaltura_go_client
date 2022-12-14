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

// CaptionParamsUpdateRequest struct for CaptionParamsUpdateRequest
type CaptionParamsUpdateRequest struct {
	CaptionParams KalturaCaptionParams `json:"captionParams"`
	Id int32 `json:"id"`
}

// NewCaptionParamsUpdateRequest instantiates a new CaptionParamsUpdateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCaptionParamsUpdateRequest(captionParams KalturaCaptionParams, id int32) *CaptionParamsUpdateRequest {
	this := CaptionParamsUpdateRequest{}
	this.CaptionParams = captionParams
	this.Id = id
	return &this
}

// NewCaptionParamsUpdateRequestWithDefaults instantiates a new CaptionParamsUpdateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCaptionParamsUpdateRequestWithDefaults() *CaptionParamsUpdateRequest {
	this := CaptionParamsUpdateRequest{}
	return &this
}

// GetCaptionParams returns the CaptionParams field value
func (o *CaptionParamsUpdateRequest) GetCaptionParams() KalturaCaptionParams {
	if o == nil {
		var ret KalturaCaptionParams
		return ret
	}

	return o.CaptionParams
}

// GetCaptionParamsOk returns a tuple with the CaptionParams field value
// and a boolean to check if the value has been set.
func (o *CaptionParamsUpdateRequest) GetCaptionParamsOk() (*KalturaCaptionParams, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CaptionParams, true
}

// SetCaptionParams sets field value
func (o *CaptionParamsUpdateRequest) SetCaptionParams(v KalturaCaptionParams) {
	o.CaptionParams = v
}

// GetId returns the Id field value
func (o *CaptionParamsUpdateRequest) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *CaptionParamsUpdateRequest) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *CaptionParamsUpdateRequest) SetId(v int32) {
	o.Id = v
}

func (o CaptionParamsUpdateRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["captionParams"] = o.CaptionParams
	}
	if true {
		toSerialize["id"] = o.Id
	}
	return json.Marshal(toSerialize)
}

type NullableCaptionParamsUpdateRequest struct {
	value *CaptionParamsUpdateRequest
	isSet bool
}

func (v NullableCaptionParamsUpdateRequest) Get() *CaptionParamsUpdateRequest {
	return v.value
}

func (v *NullableCaptionParamsUpdateRequest) Set(val *CaptionParamsUpdateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCaptionParamsUpdateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCaptionParamsUpdateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCaptionParamsUpdateRequest(val *CaptionParamsUpdateRequest) *NullableCaptionParamsUpdateRequest {
	return &NullableCaptionParamsUpdateRequest{value: val, isSet: true}
}

func (v NullableCaptionParamsUpdateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCaptionParamsUpdateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


