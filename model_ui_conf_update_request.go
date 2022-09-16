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

// UiConfUpdateRequest struct for UiConfUpdateRequest
type UiConfUpdateRequest struct {
	Id int32 `json:"id"`
	UiConf KalturaUiConf `json:"uiConf"`
}

// NewUiConfUpdateRequest instantiates a new UiConfUpdateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUiConfUpdateRequest(id int32, uiConf KalturaUiConf) *UiConfUpdateRequest {
	this := UiConfUpdateRequest{}
	this.Id = id
	this.UiConf = uiConf
	return &this
}

// NewUiConfUpdateRequestWithDefaults instantiates a new UiConfUpdateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUiConfUpdateRequestWithDefaults() *UiConfUpdateRequest {
	this := UiConfUpdateRequest{}
	return &this
}

// GetId returns the Id field value
func (o *UiConfUpdateRequest) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *UiConfUpdateRequest) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *UiConfUpdateRequest) SetId(v int32) {
	o.Id = v
}

// GetUiConf returns the UiConf field value
func (o *UiConfUpdateRequest) GetUiConf() KalturaUiConf {
	if o == nil {
		var ret KalturaUiConf
		return ret
	}

	return o.UiConf
}

// GetUiConfOk returns a tuple with the UiConf field value
// and a boolean to check if the value has been set.
func (o *UiConfUpdateRequest) GetUiConfOk() (*KalturaUiConf, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UiConf, true
}

// SetUiConf sets field value
func (o *UiConfUpdateRequest) SetUiConf(v KalturaUiConf) {
	o.UiConf = v
}

func (o UiConfUpdateRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["uiConf"] = o.UiConf
	}
	return json.Marshal(toSerialize)
}

type NullableUiConfUpdateRequest struct {
	value *UiConfUpdateRequest
	isSet bool
}

func (v NullableUiConfUpdateRequest) Get() *UiConfUpdateRequest {
	return v.value
}

func (v *NullableUiConfUpdateRequest) Set(val *UiConfUpdateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUiConfUpdateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUiConfUpdateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUiConfUpdateRequest(val *UiConfUpdateRequest) *NullableUiConfUpdateRequest {
	return &NullableUiConfUpdateRequest{value: val, isSet: true}
}

func (v NullableUiConfUpdateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUiConfUpdateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

