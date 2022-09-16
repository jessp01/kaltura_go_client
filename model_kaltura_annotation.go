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

// KalturaAnnotation struct for KalturaAnnotation
type KalturaAnnotation struct {
	KalturaCuePoint
}

// NewKalturaAnnotation instantiates a new KalturaAnnotation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaAnnotation() *KalturaAnnotation {
	this := KalturaAnnotation{}
	return &this
}

// NewKalturaAnnotationWithDefaults instantiates a new KalturaAnnotation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaAnnotationWithDefaults() *KalturaAnnotation {
	this := KalturaAnnotation{}
	return &this
}

func (o KalturaAnnotation) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedKalturaCuePoint, errKalturaCuePoint := json.Marshal(o.KalturaCuePoint)
	if errKalturaCuePoint != nil {
		return []byte{}, errKalturaCuePoint
	}
	errKalturaCuePoint = json.Unmarshal([]byte(serializedKalturaCuePoint), &toSerialize)
	if errKalturaCuePoint != nil {
		return []byte{}, errKalturaCuePoint
	}
	return json.Marshal(toSerialize)
}

type NullableKalturaAnnotation struct {
	value *KalturaAnnotation
	isSet bool
}

func (v NullableKalturaAnnotation) Get() *KalturaAnnotation {
	return v.value
}

func (v *NullableKalturaAnnotation) Set(val *KalturaAnnotation) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaAnnotation) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaAnnotation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaAnnotation(val *KalturaAnnotation) *NullableKalturaAnnotation {
	return &NullableKalturaAnnotation{value: val, isSet: true}
}

func (v NullableKalturaAnnotation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaAnnotation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


