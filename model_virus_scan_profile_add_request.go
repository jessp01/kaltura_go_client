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

// VirusScanProfileAddRequest struct for VirusScanProfileAddRequest
type VirusScanProfileAddRequest struct {
	VirusScanProfile KalturaVirusScanProfile `json:"virusScanProfile"`
}

// NewVirusScanProfileAddRequest instantiates a new VirusScanProfileAddRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVirusScanProfileAddRequest(virusScanProfile KalturaVirusScanProfile) *VirusScanProfileAddRequest {
	this := VirusScanProfileAddRequest{}
	this.VirusScanProfile = virusScanProfile
	return &this
}

// NewVirusScanProfileAddRequestWithDefaults instantiates a new VirusScanProfileAddRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVirusScanProfileAddRequestWithDefaults() *VirusScanProfileAddRequest {
	this := VirusScanProfileAddRequest{}
	return &this
}

// GetVirusScanProfile returns the VirusScanProfile field value
func (o *VirusScanProfileAddRequest) GetVirusScanProfile() KalturaVirusScanProfile {
	if o == nil {
		var ret KalturaVirusScanProfile
		return ret
	}

	return o.VirusScanProfile
}

// GetVirusScanProfileOk returns a tuple with the VirusScanProfile field value
// and a boolean to check if the value has been set.
func (o *VirusScanProfileAddRequest) GetVirusScanProfileOk() (*KalturaVirusScanProfile, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VirusScanProfile, true
}

// SetVirusScanProfile sets field value
func (o *VirusScanProfileAddRequest) SetVirusScanProfile(v KalturaVirusScanProfile) {
	o.VirusScanProfile = v
}

func (o VirusScanProfileAddRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["virusScanProfile"] = o.VirusScanProfile
	}
	return json.Marshal(toSerialize)
}

type NullableVirusScanProfileAddRequest struct {
	value *VirusScanProfileAddRequest
	isSet bool
}

func (v NullableVirusScanProfileAddRequest) Get() *VirusScanProfileAddRequest {
	return v.value
}

func (v *NullableVirusScanProfileAddRequest) Set(val *VirusScanProfileAddRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableVirusScanProfileAddRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableVirusScanProfileAddRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVirusScanProfileAddRequest(val *VirusScanProfileAddRequest) *NullableVirusScanProfileAddRequest {
	return &NullableVirusScanProfileAddRequest{value: val, isSet: true}
}

func (v NullableVirusScanProfileAddRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVirusScanProfileAddRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

