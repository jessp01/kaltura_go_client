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

// KalturaFilterPager The KalturaFilterPager object enables paging management to be applied upon service list actions.
type KalturaFilterPager struct {
	KalturaPager
}

// NewKalturaFilterPager instantiates a new KalturaFilterPager object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaFilterPager() *KalturaFilterPager {
	this := KalturaFilterPager{}
	return &this
}

// NewKalturaFilterPagerWithDefaults instantiates a new KalturaFilterPager object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaFilterPagerWithDefaults() *KalturaFilterPager {
	this := KalturaFilterPager{}
	return &this
}

func (o KalturaFilterPager) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedKalturaPager, errKalturaPager := json.Marshal(o.KalturaPager)
	if errKalturaPager != nil {
		return []byte{}, errKalturaPager
	}
	errKalturaPager = json.Unmarshal([]byte(serializedKalturaPager), &toSerialize)
	if errKalturaPager != nil {
		return []byte{}, errKalturaPager
	}
	return json.Marshal(toSerialize)
}

type NullableKalturaFilterPager struct {
	value *KalturaFilterPager
	isSet bool
}

func (v NullableKalturaFilterPager) Get() *KalturaFilterPager {
	return v.value
}

func (v *NullableKalturaFilterPager) Set(val *KalturaFilterPager) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaFilterPager) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaFilterPager) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaFilterPager(val *KalturaFilterPager) *NullableKalturaFilterPager {
	return &NullableKalturaFilterPager{value: val, isSet: true}
}

func (v NullableKalturaFilterPager) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaFilterPager) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


