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

// EntryServerNodeUpdateRequest struct for EntryServerNodeUpdateRequest
type EntryServerNodeUpdateRequest struct {
	EntryServerNode KalturaEntryServerNode `json:"entryServerNode"`
	Id int32 `json:"id"`
}

// NewEntryServerNodeUpdateRequest instantiates a new EntryServerNodeUpdateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEntryServerNodeUpdateRequest(entryServerNode KalturaEntryServerNode, id int32) *EntryServerNodeUpdateRequest {
	this := EntryServerNodeUpdateRequest{}
	this.EntryServerNode = entryServerNode
	this.Id = id
	return &this
}

// NewEntryServerNodeUpdateRequestWithDefaults instantiates a new EntryServerNodeUpdateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEntryServerNodeUpdateRequestWithDefaults() *EntryServerNodeUpdateRequest {
	this := EntryServerNodeUpdateRequest{}
	return &this
}

// GetEntryServerNode returns the EntryServerNode field value
func (o *EntryServerNodeUpdateRequest) GetEntryServerNode() KalturaEntryServerNode {
	if o == nil {
		var ret KalturaEntryServerNode
		return ret
	}

	return o.EntryServerNode
}

// GetEntryServerNodeOk returns a tuple with the EntryServerNode field value
// and a boolean to check if the value has been set.
func (o *EntryServerNodeUpdateRequest) GetEntryServerNodeOk() (*KalturaEntryServerNode, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EntryServerNode, true
}

// SetEntryServerNode sets field value
func (o *EntryServerNodeUpdateRequest) SetEntryServerNode(v KalturaEntryServerNode) {
	o.EntryServerNode = v
}

// GetId returns the Id field value
func (o *EntryServerNodeUpdateRequest) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *EntryServerNodeUpdateRequest) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *EntryServerNodeUpdateRequest) SetId(v int32) {
	o.Id = v
}

func (o EntryServerNodeUpdateRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["entryServerNode"] = o.EntryServerNode
	}
	if true {
		toSerialize["id"] = o.Id
	}
	return json.Marshal(toSerialize)
}

type NullableEntryServerNodeUpdateRequest struct {
	value *EntryServerNodeUpdateRequest
	isSet bool
}

func (v NullableEntryServerNodeUpdateRequest) Get() *EntryServerNodeUpdateRequest {
	return v.value
}

func (v *NullableEntryServerNodeUpdateRequest) Set(val *EntryServerNodeUpdateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableEntryServerNodeUpdateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableEntryServerNodeUpdateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEntryServerNodeUpdateRequest(val *EntryServerNodeUpdateRequest) *NullableEntryServerNodeUpdateRequest {
	return &NullableEntryServerNodeUpdateRequest{value: val, isSet: true}
}

func (v NullableEntryServerNodeUpdateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEntryServerNodeUpdateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

