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

// VendorIntegrationUpdateStatusRequest struct for VendorIntegrationUpdateStatusRequest
type VendorIntegrationUpdateStatusRequest struct {
	Id int32 `json:"id"`
	Status int32 `json:"status"`
}

// NewVendorIntegrationUpdateStatusRequest instantiates a new VendorIntegrationUpdateStatusRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVendorIntegrationUpdateStatusRequest(id int32, status int32) *VendorIntegrationUpdateStatusRequest {
	this := VendorIntegrationUpdateStatusRequest{}
	this.Id = id
	this.Status = status
	return &this
}

// NewVendorIntegrationUpdateStatusRequestWithDefaults instantiates a new VendorIntegrationUpdateStatusRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVendorIntegrationUpdateStatusRequestWithDefaults() *VendorIntegrationUpdateStatusRequest {
	this := VendorIntegrationUpdateStatusRequest{}
	return &this
}

// GetId returns the Id field value
func (o *VendorIntegrationUpdateStatusRequest) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *VendorIntegrationUpdateStatusRequest) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *VendorIntegrationUpdateStatusRequest) SetId(v int32) {
	o.Id = v
}

// GetStatus returns the Status field value
func (o *VendorIntegrationUpdateStatusRequest) GetStatus() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *VendorIntegrationUpdateStatusRequest) GetStatusOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *VendorIntegrationUpdateStatusRequest) SetStatus(v int32) {
	o.Status = v
}

func (o VendorIntegrationUpdateStatusRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["status"] = o.Status
	}
	return json.Marshal(toSerialize)
}

type NullableVendorIntegrationUpdateStatusRequest struct {
	value *VendorIntegrationUpdateStatusRequest
	isSet bool
}

func (v NullableVendorIntegrationUpdateStatusRequest) Get() *VendorIntegrationUpdateStatusRequest {
	return v.value
}

func (v *NullableVendorIntegrationUpdateStatusRequest) Set(val *VendorIntegrationUpdateStatusRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableVendorIntegrationUpdateStatusRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableVendorIntegrationUpdateStatusRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVendorIntegrationUpdateStatusRequest(val *VendorIntegrationUpdateStatusRequest) *NullableVendorIntegrationUpdateStatusRequest {
	return &NullableVendorIntegrationUpdateStatusRequest{value: val, isSet: true}
}

func (v NullableVendorIntegrationUpdateStatusRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVendorIntegrationUpdateStatusRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


