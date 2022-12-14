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

// ServerNodeDeleteRequest struct for ServerNodeDeleteRequest
type ServerNodeDeleteRequest struct {
	ServerNodeId string `json:"serverNodeId"`
}

// NewServerNodeDeleteRequest instantiates a new ServerNodeDeleteRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServerNodeDeleteRequest(serverNodeId string) *ServerNodeDeleteRequest {
	this := ServerNodeDeleteRequest{}
	this.ServerNodeId = serverNodeId
	return &this
}

// NewServerNodeDeleteRequestWithDefaults instantiates a new ServerNodeDeleteRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServerNodeDeleteRequestWithDefaults() *ServerNodeDeleteRequest {
	this := ServerNodeDeleteRequest{}
	return &this
}

// GetServerNodeId returns the ServerNodeId field value
func (o *ServerNodeDeleteRequest) GetServerNodeId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ServerNodeId
}

// GetServerNodeIdOk returns a tuple with the ServerNodeId field value
// and a boolean to check if the value has been set.
func (o *ServerNodeDeleteRequest) GetServerNodeIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServerNodeId, true
}

// SetServerNodeId sets field value
func (o *ServerNodeDeleteRequest) SetServerNodeId(v string) {
	o.ServerNodeId = v
}

func (o ServerNodeDeleteRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["serverNodeId"] = o.ServerNodeId
	}
	return json.Marshal(toSerialize)
}

type NullableServerNodeDeleteRequest struct {
	value *ServerNodeDeleteRequest
	isSet bool
}

func (v NullableServerNodeDeleteRequest) Get() *ServerNodeDeleteRequest {
	return v.value
}

func (v *NullableServerNodeDeleteRequest) Set(val *ServerNodeDeleteRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableServerNodeDeleteRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableServerNodeDeleteRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServerNodeDeleteRequest(val *ServerNodeDeleteRequest) *NullableServerNodeDeleteRequest {
	return &NullableServerNodeDeleteRequest{value: val, isSet: true}
}

func (v NullableServerNodeDeleteRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServerNodeDeleteRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


