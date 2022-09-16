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

// FlavorParamsGetByConversionProfileIdRequest struct for FlavorParamsGetByConversionProfileIdRequest
type FlavorParamsGetByConversionProfileIdRequest struct {
	ConversionProfileId int32 `json:"conversionProfileId"`
}

// NewFlavorParamsGetByConversionProfileIdRequest instantiates a new FlavorParamsGetByConversionProfileIdRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFlavorParamsGetByConversionProfileIdRequest(conversionProfileId int32) *FlavorParamsGetByConversionProfileIdRequest {
	this := FlavorParamsGetByConversionProfileIdRequest{}
	this.ConversionProfileId = conversionProfileId
	return &this
}

// NewFlavorParamsGetByConversionProfileIdRequestWithDefaults instantiates a new FlavorParamsGetByConversionProfileIdRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFlavorParamsGetByConversionProfileIdRequestWithDefaults() *FlavorParamsGetByConversionProfileIdRequest {
	this := FlavorParamsGetByConversionProfileIdRequest{}
	return &this
}

// GetConversionProfileId returns the ConversionProfileId field value
func (o *FlavorParamsGetByConversionProfileIdRequest) GetConversionProfileId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ConversionProfileId
}

// GetConversionProfileIdOk returns a tuple with the ConversionProfileId field value
// and a boolean to check if the value has been set.
func (o *FlavorParamsGetByConversionProfileIdRequest) GetConversionProfileIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConversionProfileId, true
}

// SetConversionProfileId sets field value
func (o *FlavorParamsGetByConversionProfileIdRequest) SetConversionProfileId(v int32) {
	o.ConversionProfileId = v
}

func (o FlavorParamsGetByConversionProfileIdRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["conversionProfileId"] = o.ConversionProfileId
	}
	return json.Marshal(toSerialize)
}

type NullableFlavorParamsGetByConversionProfileIdRequest struct {
	value *FlavorParamsGetByConversionProfileIdRequest
	isSet bool
}

func (v NullableFlavorParamsGetByConversionProfileIdRequest) Get() *FlavorParamsGetByConversionProfileIdRequest {
	return v.value
}

func (v *NullableFlavorParamsGetByConversionProfileIdRequest) Set(val *FlavorParamsGetByConversionProfileIdRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableFlavorParamsGetByConversionProfileIdRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableFlavorParamsGetByConversionProfileIdRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFlavorParamsGetByConversionProfileIdRequest(val *FlavorParamsGetByConversionProfileIdRequest) *NullableFlavorParamsGetByConversionProfileIdRequest {
	return &NullableFlavorParamsGetByConversionProfileIdRequest{value: val, isSet: true}
}

func (v NullableFlavorParamsGetByConversionProfileIdRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFlavorParamsGetByConversionProfileIdRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


