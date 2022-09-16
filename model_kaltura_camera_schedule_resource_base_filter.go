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

// KalturaCameraScheduleResourceBaseFilter `abstract`
type KalturaCameraScheduleResourceBaseFilter struct {
	KalturaScheduleResourceFilter
}

// NewKalturaCameraScheduleResourceBaseFilter instantiates a new KalturaCameraScheduleResourceBaseFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaCameraScheduleResourceBaseFilter() *KalturaCameraScheduleResourceBaseFilter {
	this := KalturaCameraScheduleResourceBaseFilter{}
	return &this
}

// NewKalturaCameraScheduleResourceBaseFilterWithDefaults instantiates a new KalturaCameraScheduleResourceBaseFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaCameraScheduleResourceBaseFilterWithDefaults() *KalturaCameraScheduleResourceBaseFilter {
	this := KalturaCameraScheduleResourceBaseFilter{}
	return &this
}

func (o KalturaCameraScheduleResourceBaseFilter) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedKalturaScheduleResourceFilter, errKalturaScheduleResourceFilter := json.Marshal(o.KalturaScheduleResourceFilter)
	if errKalturaScheduleResourceFilter != nil {
		return []byte{}, errKalturaScheduleResourceFilter
	}
	errKalturaScheduleResourceFilter = json.Unmarshal([]byte(serializedKalturaScheduleResourceFilter), &toSerialize)
	if errKalturaScheduleResourceFilter != nil {
		return []byte{}, errKalturaScheduleResourceFilter
	}
	return json.Marshal(toSerialize)
}

type NullableKalturaCameraScheduleResourceBaseFilter struct {
	value *KalturaCameraScheduleResourceBaseFilter
	isSet bool
}

func (v NullableKalturaCameraScheduleResourceBaseFilter) Get() *KalturaCameraScheduleResourceBaseFilter {
	return v.value
}

func (v *NullableKalturaCameraScheduleResourceBaseFilter) Set(val *KalturaCameraScheduleResourceBaseFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaCameraScheduleResourceBaseFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaCameraScheduleResourceBaseFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaCameraScheduleResourceBaseFilter(val *KalturaCameraScheduleResourceBaseFilter) *NullableKalturaCameraScheduleResourceBaseFilter {
	return &NullableKalturaCameraScheduleResourceBaseFilter{value: val, isSet: true}
}

func (v NullableKalturaCameraScheduleResourceBaseFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaCameraScheduleResourceBaseFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


