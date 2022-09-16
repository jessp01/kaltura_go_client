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

// KalturaConferenceServerNode struct for KalturaConferenceServerNode
type KalturaConferenceServerNode struct {
	KalturaServerNode
}

// NewKalturaConferenceServerNode instantiates a new KalturaConferenceServerNode object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaConferenceServerNode() *KalturaConferenceServerNode {
	this := KalturaConferenceServerNode{}
	return &this
}

// NewKalturaConferenceServerNodeWithDefaults instantiates a new KalturaConferenceServerNode object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaConferenceServerNodeWithDefaults() *KalturaConferenceServerNode {
	this := KalturaConferenceServerNode{}
	return &this
}

func (o KalturaConferenceServerNode) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedKalturaServerNode, errKalturaServerNode := json.Marshal(o.KalturaServerNode)
	if errKalturaServerNode != nil {
		return []byte{}, errKalturaServerNode
	}
	errKalturaServerNode = json.Unmarshal([]byte(serializedKalturaServerNode), &toSerialize)
	if errKalturaServerNode != nil {
		return []byte{}, errKalturaServerNode
	}
	return json.Marshal(toSerialize)
}

type NullableKalturaConferenceServerNode struct {
	value *KalturaConferenceServerNode
	isSet bool
}

func (v NullableKalturaConferenceServerNode) Get() *KalturaConferenceServerNode {
	return v.value
}

func (v *NullableKalturaConferenceServerNode) Set(val *KalturaConferenceServerNode) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaConferenceServerNode) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaConferenceServerNode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaConferenceServerNode(val *KalturaConferenceServerNode) *NullableKalturaConferenceServerNode {
	return &NullableKalturaConferenceServerNode{value: val, isSet: true}
}

func (v NullableKalturaConferenceServerNode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaConferenceServerNode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

