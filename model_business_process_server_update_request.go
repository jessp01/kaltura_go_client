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

// BusinessProcessServerUpdateRequest struct for BusinessProcessServerUpdateRequest
type BusinessProcessServerUpdateRequest struct {
	BusinessProcessServer KalturaBusinessProcessServer `json:"businessProcessServer"`
	Id int32 `json:"id"`
}

// NewBusinessProcessServerUpdateRequest instantiates a new BusinessProcessServerUpdateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBusinessProcessServerUpdateRequest(businessProcessServer KalturaBusinessProcessServer, id int32) *BusinessProcessServerUpdateRequest {
	this := BusinessProcessServerUpdateRequest{}
	this.BusinessProcessServer = businessProcessServer
	this.Id = id
	return &this
}

// NewBusinessProcessServerUpdateRequestWithDefaults instantiates a new BusinessProcessServerUpdateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBusinessProcessServerUpdateRequestWithDefaults() *BusinessProcessServerUpdateRequest {
	this := BusinessProcessServerUpdateRequest{}
	return &this
}

// GetBusinessProcessServer returns the BusinessProcessServer field value
func (o *BusinessProcessServerUpdateRequest) GetBusinessProcessServer() KalturaBusinessProcessServer {
	if o == nil {
		var ret KalturaBusinessProcessServer
		return ret
	}

	return o.BusinessProcessServer
}

// GetBusinessProcessServerOk returns a tuple with the BusinessProcessServer field value
// and a boolean to check if the value has been set.
func (o *BusinessProcessServerUpdateRequest) GetBusinessProcessServerOk() (*KalturaBusinessProcessServer, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BusinessProcessServer, true
}

// SetBusinessProcessServer sets field value
func (o *BusinessProcessServerUpdateRequest) SetBusinessProcessServer(v KalturaBusinessProcessServer) {
	o.BusinessProcessServer = v
}

// GetId returns the Id field value
func (o *BusinessProcessServerUpdateRequest) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *BusinessProcessServerUpdateRequest) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *BusinessProcessServerUpdateRequest) SetId(v int32) {
	o.Id = v
}

func (o BusinessProcessServerUpdateRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["businessProcessServer"] = o.BusinessProcessServer
	}
	if true {
		toSerialize["id"] = o.Id
	}
	return json.Marshal(toSerialize)
}

type NullableBusinessProcessServerUpdateRequest struct {
	value *BusinessProcessServerUpdateRequest
	isSet bool
}

func (v NullableBusinessProcessServerUpdateRequest) Get() *BusinessProcessServerUpdateRequest {
	return v.value
}

func (v *NullableBusinessProcessServerUpdateRequest) Set(val *BusinessProcessServerUpdateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableBusinessProcessServerUpdateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableBusinessProcessServerUpdateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBusinessProcessServerUpdateRequest(val *BusinessProcessServerUpdateRequest) *NullableBusinessProcessServerUpdateRequest {
	return &NullableBusinessProcessServerUpdateRequest{value: val, isSet: true}
}

func (v NullableBusinessProcessServerUpdateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBusinessProcessServerUpdateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


