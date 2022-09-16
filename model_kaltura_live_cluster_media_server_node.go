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

// KalturaLiveClusterMediaServerNode struct for KalturaLiveClusterMediaServerNode
type KalturaLiveClusterMediaServerNode struct {
	KalturaMediaServerNode
}

// NewKalturaLiveClusterMediaServerNode instantiates a new KalturaLiveClusterMediaServerNode object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaLiveClusterMediaServerNode() *KalturaLiveClusterMediaServerNode {
	this := KalturaLiveClusterMediaServerNode{}
	return &this
}

// NewKalturaLiveClusterMediaServerNodeWithDefaults instantiates a new KalturaLiveClusterMediaServerNode object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaLiveClusterMediaServerNodeWithDefaults() *KalturaLiveClusterMediaServerNode {
	this := KalturaLiveClusterMediaServerNode{}
	return &this
}

func (o KalturaLiveClusterMediaServerNode) MarshalJSON() ([]byte, error) {
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

type NullableKalturaLiveClusterMediaServerNode struct {
	value *KalturaLiveClusterMediaServerNode
	isSet bool
}

func (v NullableKalturaLiveClusterMediaServerNode) Get() *KalturaLiveClusterMediaServerNode {
	return v.value
}

func (v *NullableKalturaLiveClusterMediaServerNode) Set(val *KalturaLiveClusterMediaServerNode) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaLiveClusterMediaServerNode) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaLiveClusterMediaServerNode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaLiveClusterMediaServerNode(val *KalturaLiveClusterMediaServerNode) *NullableKalturaLiveClusterMediaServerNode {
	return &NullableKalturaLiveClusterMediaServerNode{value: val, isSet: true}
}

func (v NullableKalturaLiveClusterMediaServerNode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaLiveClusterMediaServerNode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

