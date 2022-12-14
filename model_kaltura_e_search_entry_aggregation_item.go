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

// KalturaESearchEntryAggregationItem struct for KalturaESearchEntryAggregationItem
type KalturaESearchEntryAggregationItem struct {
	KalturaESearchAggregationItem
}

// NewKalturaESearchEntryAggregationItem instantiates a new KalturaESearchEntryAggregationItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaESearchEntryAggregationItem() *KalturaESearchEntryAggregationItem {
	this := KalturaESearchEntryAggregationItem{}
	return &this
}

// NewKalturaESearchEntryAggregationItemWithDefaults instantiates a new KalturaESearchEntryAggregationItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaESearchEntryAggregationItemWithDefaults() *KalturaESearchEntryAggregationItem {
	this := KalturaESearchEntryAggregationItem{}
	return &this
}

func (o KalturaESearchEntryAggregationItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedKalturaESearchAggregationItem, errKalturaESearchAggregationItem := json.Marshal(o.KalturaESearchAggregationItem)
	if errKalturaESearchAggregationItem != nil {
		return []byte{}, errKalturaESearchAggregationItem
	}
	errKalturaESearchAggregationItem = json.Unmarshal([]byte(serializedKalturaESearchAggregationItem), &toSerialize)
	if errKalturaESearchAggregationItem != nil {
		return []byte{}, errKalturaESearchAggregationItem
	}
	return json.Marshal(toSerialize)
}

type NullableKalturaESearchEntryAggregationItem struct {
	value *KalturaESearchEntryAggregationItem
	isSet bool
}

func (v NullableKalturaESearchEntryAggregationItem) Get() *KalturaESearchEntryAggregationItem {
	return v.value
}

func (v *NullableKalturaESearchEntryAggregationItem) Set(val *KalturaESearchEntryAggregationItem) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaESearchEntryAggregationItem) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaESearchEntryAggregationItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaESearchEntryAggregationItem(val *KalturaESearchEntryAggregationItem) *NullableKalturaESearchEntryAggregationItem {
	return &NullableKalturaESearchEntryAggregationItem{value: val, isSet: true}
}

func (v NullableKalturaESearchEntryAggregationItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaESearchEntryAggregationItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


