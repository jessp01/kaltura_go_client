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

// GroupAddRequest struct for GroupAddRequest
type GroupAddRequest struct {
	Group KalturaGroup `json:"group"`
}

// NewGroupAddRequest instantiates a new GroupAddRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGroupAddRequest(group KalturaGroup) *GroupAddRequest {
	this := GroupAddRequest{}
	this.Group = group
	return &this
}

// NewGroupAddRequestWithDefaults instantiates a new GroupAddRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGroupAddRequestWithDefaults() *GroupAddRequest {
	this := GroupAddRequest{}
	return &this
}

// GetGroup returns the Group field value
func (o *GroupAddRequest) GetGroup() KalturaGroup {
	if o == nil {
		var ret KalturaGroup
		return ret
	}

	return o.Group
}

// GetGroupOk returns a tuple with the Group field value
// and a boolean to check if the value has been set.
func (o *GroupAddRequest) GetGroupOk() (*KalturaGroup, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Group, true
}

// SetGroup sets field value
func (o *GroupAddRequest) SetGroup(v KalturaGroup) {
	o.Group = v
}

func (o GroupAddRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["group"] = o.Group
	}
	return json.Marshal(toSerialize)
}

type NullableGroupAddRequest struct {
	value *GroupAddRequest
	isSet bool
}

func (v NullableGroupAddRequest) Get() *GroupAddRequest {
	return v.value
}

func (v *NullableGroupAddRequest) Set(val *GroupAddRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableGroupAddRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableGroupAddRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGroupAddRequest(val *GroupAddRequest) *NullableGroupAddRequest {
	return &NullableGroupAddRequest{value: val, isSet: true}
}

func (v NullableGroupAddRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGroupAddRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


