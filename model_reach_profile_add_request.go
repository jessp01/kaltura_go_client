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

// ReachProfileAddRequest struct for ReachProfileAddRequest
type ReachProfileAddRequest struct {
	ReachProfile KalturaReachProfile `json:"reachProfile"`
}

// NewReachProfileAddRequest instantiates a new ReachProfileAddRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReachProfileAddRequest(reachProfile KalturaReachProfile) *ReachProfileAddRequest {
	this := ReachProfileAddRequest{}
	this.ReachProfile = reachProfile
	return &this
}

// NewReachProfileAddRequestWithDefaults instantiates a new ReachProfileAddRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReachProfileAddRequestWithDefaults() *ReachProfileAddRequest {
	this := ReachProfileAddRequest{}
	return &this
}

// GetReachProfile returns the ReachProfile field value
func (o *ReachProfileAddRequest) GetReachProfile() KalturaReachProfile {
	if o == nil {
		var ret KalturaReachProfile
		return ret
	}

	return o.ReachProfile
}

// GetReachProfileOk returns a tuple with the ReachProfile field value
// and a boolean to check if the value has been set.
func (o *ReachProfileAddRequest) GetReachProfileOk() (*KalturaReachProfile, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReachProfile, true
}

// SetReachProfile sets field value
func (o *ReachProfileAddRequest) SetReachProfile(v KalturaReachProfile) {
	o.ReachProfile = v
}

func (o ReachProfileAddRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["reachProfile"] = o.ReachProfile
	}
	return json.Marshal(toSerialize)
}

type NullableReachProfileAddRequest struct {
	value *ReachProfileAddRequest
	isSet bool
}

func (v NullableReachProfileAddRequest) Get() *ReachProfileAddRequest {
	return v.value
}

func (v *NullableReachProfileAddRequest) Set(val *ReachProfileAddRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableReachProfileAddRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableReachProfileAddRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReachProfileAddRequest(val *ReachProfileAddRequest) *NullableReachProfileAddRequest {
	return &NullableReachProfileAddRequest{value: val, isSet: true}
}

func (v NullableReachProfileAddRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReachProfileAddRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


