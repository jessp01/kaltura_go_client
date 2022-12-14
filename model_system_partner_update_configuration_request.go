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

// SystemPartnerUpdateConfigurationRequest struct for SystemPartnerUpdateConfigurationRequest
type SystemPartnerUpdateConfigurationRequest struct {
	Configuration KalturaSystemPartnerConfiguration `json:"configuration"`
	PId int32 `json:"pId"`
}

// NewSystemPartnerUpdateConfigurationRequest instantiates a new SystemPartnerUpdateConfigurationRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSystemPartnerUpdateConfigurationRequest(configuration KalturaSystemPartnerConfiguration, pId int32) *SystemPartnerUpdateConfigurationRequest {
	this := SystemPartnerUpdateConfigurationRequest{}
	this.Configuration = configuration
	this.PId = pId
	return &this
}

// NewSystemPartnerUpdateConfigurationRequestWithDefaults instantiates a new SystemPartnerUpdateConfigurationRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSystemPartnerUpdateConfigurationRequestWithDefaults() *SystemPartnerUpdateConfigurationRequest {
	this := SystemPartnerUpdateConfigurationRequest{}
	return &this
}

// GetConfiguration returns the Configuration field value
func (o *SystemPartnerUpdateConfigurationRequest) GetConfiguration() KalturaSystemPartnerConfiguration {
	if o == nil {
		var ret KalturaSystemPartnerConfiguration
		return ret
	}

	return o.Configuration
}

// GetConfigurationOk returns a tuple with the Configuration field value
// and a boolean to check if the value has been set.
func (o *SystemPartnerUpdateConfigurationRequest) GetConfigurationOk() (*KalturaSystemPartnerConfiguration, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Configuration, true
}

// SetConfiguration sets field value
func (o *SystemPartnerUpdateConfigurationRequest) SetConfiguration(v KalturaSystemPartnerConfiguration) {
	o.Configuration = v
}

// GetPId returns the PId field value
func (o *SystemPartnerUpdateConfigurationRequest) GetPId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PId
}

// GetPIdOk returns a tuple with the PId field value
// and a boolean to check if the value has been set.
func (o *SystemPartnerUpdateConfigurationRequest) GetPIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PId, true
}

// SetPId sets field value
func (o *SystemPartnerUpdateConfigurationRequest) SetPId(v int32) {
	o.PId = v
}

func (o SystemPartnerUpdateConfigurationRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["configuration"] = o.Configuration
	}
	if true {
		toSerialize["pId"] = o.PId
	}
	return json.Marshal(toSerialize)
}

type NullableSystemPartnerUpdateConfigurationRequest struct {
	value *SystemPartnerUpdateConfigurationRequest
	isSet bool
}

func (v NullableSystemPartnerUpdateConfigurationRequest) Get() *SystemPartnerUpdateConfigurationRequest {
	return v.value
}

func (v *NullableSystemPartnerUpdateConfigurationRequest) Set(val *SystemPartnerUpdateConfigurationRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSystemPartnerUpdateConfigurationRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSystemPartnerUpdateConfigurationRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSystemPartnerUpdateConfigurationRequest(val *SystemPartnerUpdateConfigurationRequest) *NullableSystemPartnerUpdateConfigurationRequest {
	return &NullableSystemPartnerUpdateConfigurationRequest{value: val, isSet: true}
}

func (v NullableSystemPartnerUpdateConfigurationRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSystemPartnerUpdateConfigurationRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


