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

// KalturaWowzaMediaServerNode struct for KalturaWowzaMediaServerNode
type KalturaWowzaMediaServerNode struct {
	KalturaMediaServerNode
}

// NewKalturaWowzaMediaServerNode instantiates a new KalturaWowzaMediaServerNode object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaWowzaMediaServerNode() *KalturaWowzaMediaServerNode {
	this := KalturaWowzaMediaServerNode{}
	return &this
}

// NewKalturaWowzaMediaServerNodeWithDefaults instantiates a new KalturaWowzaMediaServerNode object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaWowzaMediaServerNodeWithDefaults() *KalturaWowzaMediaServerNode {
	this := KalturaWowzaMediaServerNode{}
	return &this
}

func (o KalturaWowzaMediaServerNode) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedKalturaMediaServerNode, errKalturaMediaServerNode := json.Marshal(o.KalturaMediaServerNode)
	if errKalturaMediaServerNode != nil {
		return []byte{}, errKalturaMediaServerNode
	}
	errKalturaMediaServerNode = json.Unmarshal([]byte(serializedKalturaMediaServerNode), &toSerialize)
	if errKalturaMediaServerNode != nil {
		return []byte{}, errKalturaMediaServerNode
	}
	return json.Marshal(toSerialize)
}

type NullableKalturaWowzaMediaServerNode struct {
	value *KalturaWowzaMediaServerNode
	isSet bool
}

func (v NullableKalturaWowzaMediaServerNode) Get() *KalturaWowzaMediaServerNode {
	return v.value
}

func (v *NullableKalturaWowzaMediaServerNode) Set(val *KalturaWowzaMediaServerNode) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaWowzaMediaServerNode) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaWowzaMediaServerNode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaWowzaMediaServerNode(val *KalturaWowzaMediaServerNode) *NullableKalturaWowzaMediaServerNode {
	return &NullableKalturaWowzaMediaServerNode{value: val, isSet: true}
}

func (v NullableKalturaWowzaMediaServerNode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaWowzaMediaServerNode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


