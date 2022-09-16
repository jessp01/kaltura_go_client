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

// VendorIntegrationDeleteRequest struct for VendorIntegrationDeleteRequest
type VendorIntegrationDeleteRequest struct {
	IntegrationId int32 `json:"integrationId"`
}

// NewVendorIntegrationDeleteRequest instantiates a new VendorIntegrationDeleteRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVendorIntegrationDeleteRequest(integrationId int32) *VendorIntegrationDeleteRequest {
	this := VendorIntegrationDeleteRequest{}
	this.IntegrationId = integrationId
	return &this
}

// NewVendorIntegrationDeleteRequestWithDefaults instantiates a new VendorIntegrationDeleteRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVendorIntegrationDeleteRequestWithDefaults() *VendorIntegrationDeleteRequest {
	this := VendorIntegrationDeleteRequest{}
	return &this
}

// GetIntegrationId returns the IntegrationId field value
func (o *VendorIntegrationDeleteRequest) GetIntegrationId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.IntegrationId
}

// GetIntegrationIdOk returns a tuple with the IntegrationId field value
// and a boolean to check if the value has been set.
func (o *VendorIntegrationDeleteRequest) GetIntegrationIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IntegrationId, true
}

// SetIntegrationId sets field value
func (o *VendorIntegrationDeleteRequest) SetIntegrationId(v int32) {
	o.IntegrationId = v
}

func (o VendorIntegrationDeleteRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["integrationId"] = o.IntegrationId
	}
	return json.Marshal(toSerialize)
}

type NullableVendorIntegrationDeleteRequest struct {
	value *VendorIntegrationDeleteRequest
	isSet bool
}

func (v NullableVendorIntegrationDeleteRequest) Get() *VendorIntegrationDeleteRequest {
	return v.value
}

func (v *NullableVendorIntegrationDeleteRequest) Set(val *VendorIntegrationDeleteRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableVendorIntegrationDeleteRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableVendorIntegrationDeleteRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVendorIntegrationDeleteRequest(val *VendorIntegrationDeleteRequest) *NullableVendorIntegrationDeleteRequest {
	return &NullableVendorIntegrationDeleteRequest{value: val, isSet: true}
}

func (v NullableVendorIntegrationDeleteRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVendorIntegrationDeleteRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


