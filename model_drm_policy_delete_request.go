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

// DrmPolicyDeleteRequest struct for DrmPolicyDeleteRequest
type DrmPolicyDeleteRequest struct {
	DrmPolicyId int32 `json:"drmPolicyId"`
}

// NewDrmPolicyDeleteRequest instantiates a new DrmPolicyDeleteRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDrmPolicyDeleteRequest(drmPolicyId int32) *DrmPolicyDeleteRequest {
	this := DrmPolicyDeleteRequest{}
	this.DrmPolicyId = drmPolicyId
	return &this
}

// NewDrmPolicyDeleteRequestWithDefaults instantiates a new DrmPolicyDeleteRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDrmPolicyDeleteRequestWithDefaults() *DrmPolicyDeleteRequest {
	this := DrmPolicyDeleteRequest{}
	return &this
}

// GetDrmPolicyId returns the DrmPolicyId field value
func (o *DrmPolicyDeleteRequest) GetDrmPolicyId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.DrmPolicyId
}

// GetDrmPolicyIdOk returns a tuple with the DrmPolicyId field value
// and a boolean to check if the value has been set.
func (o *DrmPolicyDeleteRequest) GetDrmPolicyIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DrmPolicyId, true
}

// SetDrmPolicyId sets field value
func (o *DrmPolicyDeleteRequest) SetDrmPolicyId(v int32) {
	o.DrmPolicyId = v
}

func (o DrmPolicyDeleteRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["drmPolicyId"] = o.DrmPolicyId
	}
	return json.Marshal(toSerialize)
}

type NullableDrmPolicyDeleteRequest struct {
	value *DrmPolicyDeleteRequest
	isSet bool
}

func (v NullableDrmPolicyDeleteRequest) Get() *DrmPolicyDeleteRequest {
	return v.value
}

func (v *NullableDrmPolicyDeleteRequest) Set(val *DrmPolicyDeleteRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableDrmPolicyDeleteRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableDrmPolicyDeleteRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDrmPolicyDeleteRequest(val *DrmPolicyDeleteRequest) *NullableDrmPolicyDeleteRequest {
	return &NullableDrmPolicyDeleteRequest{value: val, isSet: true}
}

func (v NullableDrmPolicyDeleteRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDrmPolicyDeleteRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


