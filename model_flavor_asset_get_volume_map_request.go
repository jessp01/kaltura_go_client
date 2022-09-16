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

// FlavorAssetGetVolumeMapRequest struct for FlavorAssetGetVolumeMapRequest
type FlavorAssetGetVolumeMapRequest struct {
	FlavorId string `json:"flavorId"`
}

// NewFlavorAssetGetVolumeMapRequest instantiates a new FlavorAssetGetVolumeMapRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFlavorAssetGetVolumeMapRequest(flavorId string) *FlavorAssetGetVolumeMapRequest {
	this := FlavorAssetGetVolumeMapRequest{}
	this.FlavorId = flavorId
	return &this
}

// NewFlavorAssetGetVolumeMapRequestWithDefaults instantiates a new FlavorAssetGetVolumeMapRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFlavorAssetGetVolumeMapRequestWithDefaults() *FlavorAssetGetVolumeMapRequest {
	this := FlavorAssetGetVolumeMapRequest{}
	return &this
}

// GetFlavorId returns the FlavorId field value
func (o *FlavorAssetGetVolumeMapRequest) GetFlavorId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FlavorId
}

// GetFlavorIdOk returns a tuple with the FlavorId field value
// and a boolean to check if the value has been set.
func (o *FlavorAssetGetVolumeMapRequest) GetFlavorIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FlavorId, true
}

// SetFlavorId sets field value
func (o *FlavorAssetGetVolumeMapRequest) SetFlavorId(v string) {
	o.FlavorId = v
}

func (o FlavorAssetGetVolumeMapRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["flavorId"] = o.FlavorId
	}
	return json.Marshal(toSerialize)
}

type NullableFlavorAssetGetVolumeMapRequest struct {
	value *FlavorAssetGetVolumeMapRequest
	isSet bool
}

func (v NullableFlavorAssetGetVolumeMapRequest) Get() *FlavorAssetGetVolumeMapRequest {
	return v.value
}

func (v *NullableFlavorAssetGetVolumeMapRequest) Set(val *FlavorAssetGetVolumeMapRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableFlavorAssetGetVolumeMapRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableFlavorAssetGetVolumeMapRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFlavorAssetGetVolumeMapRequest(val *FlavorAssetGetVolumeMapRequest) *NullableFlavorAssetGetVolumeMapRequest {
	return &NullableFlavorAssetGetVolumeMapRequest{value: val, isSet: true}
}

func (v NullableFlavorAssetGetVolumeMapRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFlavorAssetGetVolumeMapRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

