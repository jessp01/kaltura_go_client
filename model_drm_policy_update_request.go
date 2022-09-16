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

// DrmPolicyUpdateRequest struct for DrmPolicyUpdateRequest
type DrmPolicyUpdateRequest struct {
	DrmPolicy KalturaDrmPolicy `json:"drmPolicy"`
	DrmPolicyId int32 `json:"drmPolicyId"`
}

// NewDrmPolicyUpdateRequest instantiates a new DrmPolicyUpdateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDrmPolicyUpdateRequest(drmPolicy KalturaDrmPolicy, drmPolicyId int32) *DrmPolicyUpdateRequest {
	this := DrmPolicyUpdateRequest{}
	this.DrmPolicy = drmPolicy
	this.DrmPolicyId = drmPolicyId
	return &this
}

// NewDrmPolicyUpdateRequestWithDefaults instantiates a new DrmPolicyUpdateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDrmPolicyUpdateRequestWithDefaults() *DrmPolicyUpdateRequest {
	this := DrmPolicyUpdateRequest{}
	return &this
}

// GetDrmPolicy returns the DrmPolicy field value
func (o *DrmPolicyUpdateRequest) GetDrmPolicy() KalturaDrmPolicy {
	if o == nil {
		var ret KalturaDrmPolicy
		return ret
	}

	return o.DrmPolicy
}

// GetDrmPolicyOk returns a tuple with the DrmPolicy field value
// and a boolean to check if the value has been set.
func (o *DrmPolicyUpdateRequest) GetDrmPolicyOk() (*KalturaDrmPolicy, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DrmPolicy, true
}

// SetDrmPolicy sets field value
func (o *DrmPolicyUpdateRequest) SetDrmPolicy(v KalturaDrmPolicy) {
	o.DrmPolicy = v
}

// GetDrmPolicyId returns the DrmPolicyId field value
func (o *DrmPolicyUpdateRequest) GetDrmPolicyId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.DrmPolicyId
}

// GetDrmPolicyIdOk returns a tuple with the DrmPolicyId field value
// and a boolean to check if the value has been set.
func (o *DrmPolicyUpdateRequest) GetDrmPolicyIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DrmPolicyId, true
}

// SetDrmPolicyId sets field value
func (o *DrmPolicyUpdateRequest) SetDrmPolicyId(v int32) {
	o.DrmPolicyId = v
}

func (o DrmPolicyUpdateRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["drmPolicy"] = o.DrmPolicy
	}
	if true {
		toSerialize["drmPolicyId"] = o.DrmPolicyId
	}
	return json.Marshal(toSerialize)
}

type NullableDrmPolicyUpdateRequest struct {
	value *DrmPolicyUpdateRequest
	isSet bool
}

func (v NullableDrmPolicyUpdateRequest) Get() *DrmPolicyUpdateRequest {
	return v.value
}

func (v *NullableDrmPolicyUpdateRequest) Set(val *DrmPolicyUpdateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableDrmPolicyUpdateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableDrmPolicyUpdateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDrmPolicyUpdateRequest(val *DrmPolicyUpdateRequest) *NullableDrmPolicyUpdateRequest {
	return &NullableDrmPolicyUpdateRequest{value: val, isSet: true}
}

func (v NullableDrmPolicyUpdateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDrmPolicyUpdateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


