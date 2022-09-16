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

// KalturaSipServerNodeBaseFilter `abstract`
type KalturaSipServerNodeBaseFilter struct {
	KalturaServerNodeFilter
}

// NewKalturaSipServerNodeBaseFilter instantiates a new KalturaSipServerNodeBaseFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaSipServerNodeBaseFilter() *KalturaSipServerNodeBaseFilter {
	this := KalturaSipServerNodeBaseFilter{}
	return &this
}

// NewKalturaSipServerNodeBaseFilterWithDefaults instantiates a new KalturaSipServerNodeBaseFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaSipServerNodeBaseFilterWithDefaults() *KalturaSipServerNodeBaseFilter {
	this := KalturaSipServerNodeBaseFilter{}
	return &this
}

func (o KalturaSipServerNodeBaseFilter) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedKalturaServerNodeFilter, errKalturaServerNodeFilter := json.Marshal(o.KalturaServerNodeFilter)
	if errKalturaServerNodeFilter != nil {
		return []byte{}, errKalturaServerNodeFilter
	}
	errKalturaServerNodeFilter = json.Unmarshal([]byte(serializedKalturaServerNodeFilter), &toSerialize)
	if errKalturaServerNodeFilter != nil {
		return []byte{}, errKalturaServerNodeFilter
	}
	return json.Marshal(toSerialize)
}

type NullableKalturaSipServerNodeBaseFilter struct {
	value *KalturaSipServerNodeBaseFilter
	isSet bool
}

func (v NullableKalturaSipServerNodeBaseFilter) Get() *KalturaSipServerNodeBaseFilter {
	return v.value
}

func (v *NullableKalturaSipServerNodeBaseFilter) Set(val *KalturaSipServerNodeBaseFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaSipServerNodeBaseFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaSipServerNodeBaseFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaSipServerNodeBaseFilter(val *KalturaSipServerNodeBaseFilter) *NullableKalturaSipServerNodeBaseFilter {
	return &NullableKalturaSipServerNodeBaseFilter{value: val, isSet: true}
}

func (v NullableKalturaSipServerNodeBaseFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaSipServerNodeBaseFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


