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

// KalturaSystemPartnerOveragedLimit struct for KalturaSystemPartnerOveragedLimit
type KalturaSystemPartnerOveragedLimit struct {
	KalturaSystemPartnerLimit
}

// NewKalturaSystemPartnerOveragedLimit instantiates a new KalturaSystemPartnerOveragedLimit object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaSystemPartnerOveragedLimit() *KalturaSystemPartnerOveragedLimit {
	this := KalturaSystemPartnerOveragedLimit{}
	return &this
}

// NewKalturaSystemPartnerOveragedLimitWithDefaults instantiates a new KalturaSystemPartnerOveragedLimit object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaSystemPartnerOveragedLimitWithDefaults() *KalturaSystemPartnerOveragedLimit {
	this := KalturaSystemPartnerOveragedLimit{}
	return &this
}

func (o KalturaSystemPartnerOveragedLimit) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedKalturaSystemPartnerLimit, errKalturaSystemPartnerLimit := json.Marshal(o.KalturaSystemPartnerLimit)
	if errKalturaSystemPartnerLimit != nil {
		return []byte{}, errKalturaSystemPartnerLimit
	}
	errKalturaSystemPartnerLimit = json.Unmarshal([]byte(serializedKalturaSystemPartnerLimit), &toSerialize)
	if errKalturaSystemPartnerLimit != nil {
		return []byte{}, errKalturaSystemPartnerLimit
	}
	return json.Marshal(toSerialize)
}

type NullableKalturaSystemPartnerOveragedLimit struct {
	value *KalturaSystemPartnerOveragedLimit
	isSet bool
}

func (v NullableKalturaSystemPartnerOveragedLimit) Get() *KalturaSystemPartnerOveragedLimit {
	return v.value
}

func (v *NullableKalturaSystemPartnerOveragedLimit) Set(val *KalturaSystemPartnerOveragedLimit) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaSystemPartnerOveragedLimit) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaSystemPartnerOveragedLimit) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaSystemPartnerOveragedLimit(val *KalturaSystemPartnerOveragedLimit) *NullableKalturaSystemPartnerOveragedLimit {
	return &NullableKalturaSystemPartnerOveragedLimit{value: val, isSet: true}
}

func (v NullableKalturaSystemPartnerOveragedLimit) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaSystemPartnerOveragedLimit) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


