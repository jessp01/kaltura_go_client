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

// SystemPartnerGetRequest struct for SystemPartnerGetRequest
type SystemPartnerGetRequest struct {
	PId int32 `json:"pId"`
}

// NewSystemPartnerGetRequest instantiates a new SystemPartnerGetRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSystemPartnerGetRequest(pId int32) *SystemPartnerGetRequest {
	this := SystemPartnerGetRequest{}
	this.PId = pId
	return &this
}

// NewSystemPartnerGetRequestWithDefaults instantiates a new SystemPartnerGetRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSystemPartnerGetRequestWithDefaults() *SystemPartnerGetRequest {
	this := SystemPartnerGetRequest{}
	return &this
}

// GetPId returns the PId field value
func (o *SystemPartnerGetRequest) GetPId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PId
}

// GetPIdOk returns a tuple with the PId field value
// and a boolean to check if the value has been set.
func (o *SystemPartnerGetRequest) GetPIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PId, true
}

// SetPId sets field value
func (o *SystemPartnerGetRequest) SetPId(v int32) {
	o.PId = v
}

func (o SystemPartnerGetRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["pId"] = o.PId
	}
	return json.Marshal(toSerialize)
}

type NullableSystemPartnerGetRequest struct {
	value *SystemPartnerGetRequest
	isSet bool
}

func (v NullableSystemPartnerGetRequest) Get() *SystemPartnerGetRequest {
	return v.value
}

func (v *NullableSystemPartnerGetRequest) Set(val *SystemPartnerGetRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSystemPartnerGetRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSystemPartnerGetRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSystemPartnerGetRequest(val *SystemPartnerGetRequest) *NullableSystemPartnerGetRequest {
	return &NullableSystemPartnerGetRequest{value: val, isSet: true}
}

func (v NullableSystemPartnerGetRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSystemPartnerGetRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


