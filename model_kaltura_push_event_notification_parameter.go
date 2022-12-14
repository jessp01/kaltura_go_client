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

// KalturaPushEventNotificationParameter struct for KalturaPushEventNotificationParameter
type KalturaPushEventNotificationParameter struct {
	KalturaEventNotificationParameter
}

// NewKalturaPushEventNotificationParameter instantiates a new KalturaPushEventNotificationParameter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaPushEventNotificationParameter() *KalturaPushEventNotificationParameter {
	this := KalturaPushEventNotificationParameter{}
	return &this
}

// NewKalturaPushEventNotificationParameterWithDefaults instantiates a new KalturaPushEventNotificationParameter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaPushEventNotificationParameterWithDefaults() *KalturaPushEventNotificationParameter {
	this := KalturaPushEventNotificationParameter{}
	return &this
}

func (o KalturaPushEventNotificationParameter) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedKalturaEventNotificationParameter, errKalturaEventNotificationParameter := json.Marshal(o.KalturaEventNotificationParameter)
	if errKalturaEventNotificationParameter != nil {
		return []byte{}, errKalturaEventNotificationParameter
	}
	errKalturaEventNotificationParameter = json.Unmarshal([]byte(serializedKalturaEventNotificationParameter), &toSerialize)
	if errKalturaEventNotificationParameter != nil {
		return []byte{}, errKalturaEventNotificationParameter
	}
	return json.Marshal(toSerialize)
}

type NullableKalturaPushEventNotificationParameter struct {
	value *KalturaPushEventNotificationParameter
	isSet bool
}

func (v NullableKalturaPushEventNotificationParameter) Get() *KalturaPushEventNotificationParameter {
	return v.value
}

func (v *NullableKalturaPushEventNotificationParameter) Set(val *KalturaPushEventNotificationParameter) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaPushEventNotificationParameter) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaPushEventNotificationParameter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaPushEventNotificationParameter(val *KalturaPushEventNotificationParameter) *NullableKalturaPushEventNotificationParameter {
	return &NullableKalturaPushEventNotificationParameter{value: val, isSet: true}
}

func (v NullableKalturaPushEventNotificationParameter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaPushEventNotificationParameter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


