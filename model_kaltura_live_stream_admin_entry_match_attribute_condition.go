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

// KalturaLiveStreamAdminEntryMatchAttributeCondition Auto-generated class.  Used to search KalturaLiveStreamAdminEntry attributes. Use KalturaLiveStreamAdminEntryMatchAttribute enum to provide attribute name. /
type KalturaLiveStreamAdminEntryMatchAttributeCondition struct {
	KalturaSearchMatchAttributeCondition
}

// NewKalturaLiveStreamAdminEntryMatchAttributeCondition instantiates a new KalturaLiveStreamAdminEntryMatchAttributeCondition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaLiveStreamAdminEntryMatchAttributeCondition() *KalturaLiveStreamAdminEntryMatchAttributeCondition {
	this := KalturaLiveStreamAdminEntryMatchAttributeCondition{}
	return &this
}

// NewKalturaLiveStreamAdminEntryMatchAttributeConditionWithDefaults instantiates a new KalturaLiveStreamAdminEntryMatchAttributeCondition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaLiveStreamAdminEntryMatchAttributeConditionWithDefaults() *KalturaLiveStreamAdminEntryMatchAttributeCondition {
	this := KalturaLiveStreamAdminEntryMatchAttributeCondition{}
	return &this
}

func (o KalturaLiveStreamAdminEntryMatchAttributeCondition) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedKalturaSearchMatchAttributeCondition, errKalturaSearchMatchAttributeCondition := json.Marshal(o.KalturaSearchMatchAttributeCondition)
	if errKalturaSearchMatchAttributeCondition != nil {
		return []byte{}, errKalturaSearchMatchAttributeCondition
	}
	errKalturaSearchMatchAttributeCondition = json.Unmarshal([]byte(serializedKalturaSearchMatchAttributeCondition), &toSerialize)
	if errKalturaSearchMatchAttributeCondition != nil {
		return []byte{}, errKalturaSearchMatchAttributeCondition
	}
	return json.Marshal(toSerialize)
}

type NullableKalturaLiveStreamAdminEntryMatchAttributeCondition struct {
	value *KalturaLiveStreamAdminEntryMatchAttributeCondition
	isSet bool
}

func (v NullableKalturaLiveStreamAdminEntryMatchAttributeCondition) Get() *KalturaLiveStreamAdminEntryMatchAttributeCondition {
	return v.value
}

func (v *NullableKalturaLiveStreamAdminEntryMatchAttributeCondition) Set(val *KalturaLiveStreamAdminEntryMatchAttributeCondition) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaLiveStreamAdminEntryMatchAttributeCondition) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaLiveStreamAdminEntryMatchAttributeCondition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaLiveStreamAdminEntryMatchAttributeCondition(val *KalturaLiveStreamAdminEntryMatchAttributeCondition) *NullableKalturaLiveStreamAdminEntryMatchAttributeCondition {
	return &NullableKalturaLiveStreamAdminEntryMatchAttributeCondition{value: val, isSet: true}
}

func (v NullableKalturaLiveStreamAdminEntryMatchAttributeCondition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaLiveStreamAdminEntryMatchAttributeCondition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


