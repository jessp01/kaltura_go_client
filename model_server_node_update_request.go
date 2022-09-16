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

// ServerNodeUpdateRequest struct for ServerNodeUpdateRequest
type ServerNodeUpdateRequest struct {
	ServerNode KalturaServerNode `json:"serverNode"`
	ServerNodeId int32 `json:"serverNodeId"`
}

// NewServerNodeUpdateRequest instantiates a new ServerNodeUpdateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServerNodeUpdateRequest(serverNode KalturaServerNode, serverNodeId int32) *ServerNodeUpdateRequest {
	this := ServerNodeUpdateRequest{}
	this.ServerNode = serverNode
	this.ServerNodeId = serverNodeId
	return &this
}

// NewServerNodeUpdateRequestWithDefaults instantiates a new ServerNodeUpdateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServerNodeUpdateRequestWithDefaults() *ServerNodeUpdateRequest {
	this := ServerNodeUpdateRequest{}
	return &this
}

// GetServerNode returns the ServerNode field value
func (o *ServerNodeUpdateRequest) GetServerNode() KalturaServerNode {
	if o == nil {
		var ret KalturaServerNode
		return ret
	}

	return o.ServerNode
}

// GetServerNodeOk returns a tuple with the ServerNode field value
// and a boolean to check if the value has been set.
func (o *ServerNodeUpdateRequest) GetServerNodeOk() (*KalturaServerNode, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServerNode, true
}

// SetServerNode sets field value
func (o *ServerNodeUpdateRequest) SetServerNode(v KalturaServerNode) {
	o.ServerNode = v
}

// GetServerNodeId returns the ServerNodeId field value
func (o *ServerNodeUpdateRequest) GetServerNodeId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ServerNodeId
}

// GetServerNodeIdOk returns a tuple with the ServerNodeId field value
// and a boolean to check if the value has been set.
func (o *ServerNodeUpdateRequest) GetServerNodeIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServerNodeId, true
}

// SetServerNodeId sets field value
func (o *ServerNodeUpdateRequest) SetServerNodeId(v int32) {
	o.ServerNodeId = v
}

func (o ServerNodeUpdateRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["serverNode"] = o.ServerNode
	}
	if true {
		toSerialize["serverNodeId"] = o.ServerNodeId
	}
	return json.Marshal(toSerialize)
}

type NullableServerNodeUpdateRequest struct {
	value *ServerNodeUpdateRequest
	isSet bool
}

func (v NullableServerNodeUpdateRequest) Get() *ServerNodeUpdateRequest {
	return v.value
}

func (v *NullableServerNodeUpdateRequest) Set(val *ServerNodeUpdateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableServerNodeUpdateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableServerNodeUpdateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServerNodeUpdateRequest(val *ServerNodeUpdateRequest) *NullableServerNodeUpdateRequest {
	return &NullableServerNodeUpdateRequest{value: val, isSet: true}
}

func (v NullableServerNodeUpdateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServerNodeUpdateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


