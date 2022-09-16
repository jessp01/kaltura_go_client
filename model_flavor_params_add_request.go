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

// FlavorParamsAddRequest struct for FlavorParamsAddRequest
type FlavorParamsAddRequest struct {
	FlavorParams KalturaFlavorParams `json:"flavorParams"`
}

// NewFlavorParamsAddRequest instantiates a new FlavorParamsAddRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFlavorParamsAddRequest(flavorParams KalturaFlavorParams) *FlavorParamsAddRequest {
	this := FlavorParamsAddRequest{}
	this.FlavorParams = flavorParams
	return &this
}

// NewFlavorParamsAddRequestWithDefaults instantiates a new FlavorParamsAddRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFlavorParamsAddRequestWithDefaults() *FlavorParamsAddRequest {
	this := FlavorParamsAddRequest{}
	return &this
}

// GetFlavorParams returns the FlavorParams field value
func (o *FlavorParamsAddRequest) GetFlavorParams() KalturaFlavorParams {
	if o == nil {
		var ret KalturaFlavorParams
		return ret
	}

	return o.FlavorParams
}

// GetFlavorParamsOk returns a tuple with the FlavorParams field value
// and a boolean to check if the value has been set.
func (o *FlavorParamsAddRequest) GetFlavorParamsOk() (*KalturaFlavorParams, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FlavorParams, true
}

// SetFlavorParams sets field value
func (o *FlavorParamsAddRequest) SetFlavorParams(v KalturaFlavorParams) {
	o.FlavorParams = v
}

func (o FlavorParamsAddRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["flavorParams"] = o.FlavorParams
	}
	return json.Marshal(toSerialize)
}

type NullableFlavorParamsAddRequest struct {
	value *FlavorParamsAddRequest
	isSet bool
}

func (v NullableFlavorParamsAddRequest) Get() *FlavorParamsAddRequest {
	return v.value
}

func (v *NullableFlavorParamsAddRequest) Set(val *FlavorParamsAddRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableFlavorParamsAddRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableFlavorParamsAddRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFlavorParamsAddRequest(val *FlavorParamsAddRequest) *NullableFlavorParamsAddRequest {
	return &NullableFlavorParamsAddRequest{value: val, isSet: true}
}

func (v NullableFlavorParamsAddRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFlavorParamsAddRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

