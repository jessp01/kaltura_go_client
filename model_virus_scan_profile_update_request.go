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

// VirusScanProfileUpdateRequest struct for VirusScanProfileUpdateRequest
type VirusScanProfileUpdateRequest struct {
	VirusScanProfile KalturaVirusScanProfile `json:"virusScanProfile"`
	VirusScanProfileId int32 `json:"virusScanProfileId"`
}

// NewVirusScanProfileUpdateRequest instantiates a new VirusScanProfileUpdateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVirusScanProfileUpdateRequest(virusScanProfile KalturaVirusScanProfile, virusScanProfileId int32) *VirusScanProfileUpdateRequest {
	this := VirusScanProfileUpdateRequest{}
	this.VirusScanProfile = virusScanProfile
	this.VirusScanProfileId = virusScanProfileId
	return &this
}

// NewVirusScanProfileUpdateRequestWithDefaults instantiates a new VirusScanProfileUpdateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVirusScanProfileUpdateRequestWithDefaults() *VirusScanProfileUpdateRequest {
	this := VirusScanProfileUpdateRequest{}
	return &this
}

// GetVirusScanProfile returns the VirusScanProfile field value
func (o *VirusScanProfileUpdateRequest) GetVirusScanProfile() KalturaVirusScanProfile {
	if o == nil {
		var ret KalturaVirusScanProfile
		return ret
	}

	return o.VirusScanProfile
}

// GetVirusScanProfileOk returns a tuple with the VirusScanProfile field value
// and a boolean to check if the value has been set.
func (o *VirusScanProfileUpdateRequest) GetVirusScanProfileOk() (*KalturaVirusScanProfile, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VirusScanProfile, true
}

// SetVirusScanProfile sets field value
func (o *VirusScanProfileUpdateRequest) SetVirusScanProfile(v KalturaVirusScanProfile) {
	o.VirusScanProfile = v
}

// GetVirusScanProfileId returns the VirusScanProfileId field value
func (o *VirusScanProfileUpdateRequest) GetVirusScanProfileId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.VirusScanProfileId
}

// GetVirusScanProfileIdOk returns a tuple with the VirusScanProfileId field value
// and a boolean to check if the value has been set.
func (o *VirusScanProfileUpdateRequest) GetVirusScanProfileIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VirusScanProfileId, true
}

// SetVirusScanProfileId sets field value
func (o *VirusScanProfileUpdateRequest) SetVirusScanProfileId(v int32) {
	o.VirusScanProfileId = v
}

func (o VirusScanProfileUpdateRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["virusScanProfile"] = o.VirusScanProfile
	}
	if true {
		toSerialize["virusScanProfileId"] = o.VirusScanProfileId
	}
	return json.Marshal(toSerialize)
}

type NullableVirusScanProfileUpdateRequest struct {
	value *VirusScanProfileUpdateRequest
	isSet bool
}

func (v NullableVirusScanProfileUpdateRequest) Get() *VirusScanProfileUpdateRequest {
	return v.value
}

func (v *NullableVirusScanProfileUpdateRequest) Set(val *VirusScanProfileUpdateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableVirusScanProfileUpdateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableVirusScanProfileUpdateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVirusScanProfileUpdateRequest(val *VirusScanProfileUpdateRequest) *NullableVirusScanProfileUpdateRequest {
	return &NullableVirusScanProfileUpdateRequest{value: val, isSet: true}
}

func (v NullableVirusScanProfileUpdateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVirusScanProfileUpdateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


