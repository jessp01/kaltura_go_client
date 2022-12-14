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

// KalturaDexterIntegrationJobProviderData struct for KalturaDexterIntegrationJobProviderData
type KalturaDexterIntegrationJobProviderData struct {
	KalturaIntegrationJobProviderData
}

// NewKalturaDexterIntegrationJobProviderData instantiates a new KalturaDexterIntegrationJobProviderData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaDexterIntegrationJobProviderData() *KalturaDexterIntegrationJobProviderData {
	this := KalturaDexterIntegrationJobProviderData{}
	return &this
}

// NewKalturaDexterIntegrationJobProviderDataWithDefaults instantiates a new KalturaDexterIntegrationJobProviderData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaDexterIntegrationJobProviderDataWithDefaults() *KalturaDexterIntegrationJobProviderData {
	this := KalturaDexterIntegrationJobProviderData{}
	return &this
}

func (o KalturaDexterIntegrationJobProviderData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedKalturaIntegrationJobProviderData, errKalturaIntegrationJobProviderData := json.Marshal(o.KalturaIntegrationJobProviderData)
	if errKalturaIntegrationJobProviderData != nil {
		return []byte{}, errKalturaIntegrationJobProviderData
	}
	errKalturaIntegrationJobProviderData = json.Unmarshal([]byte(serializedKalturaIntegrationJobProviderData), &toSerialize)
	if errKalturaIntegrationJobProviderData != nil {
		return []byte{}, errKalturaIntegrationJobProviderData
	}
	return json.Marshal(toSerialize)
}

type NullableKalturaDexterIntegrationJobProviderData struct {
	value *KalturaDexterIntegrationJobProviderData
	isSet bool
}

func (v NullableKalturaDexterIntegrationJobProviderData) Get() *KalturaDexterIntegrationJobProviderData {
	return v.value
}

func (v *NullableKalturaDexterIntegrationJobProviderData) Set(val *KalturaDexterIntegrationJobProviderData) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaDexterIntegrationJobProviderData) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaDexterIntegrationJobProviderData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaDexterIntegrationJobProviderData(val *KalturaDexterIntegrationJobProviderData) *NullableKalturaDexterIntegrationJobProviderData {
	return &NullableKalturaDexterIntegrationJobProviderData{value: val, isSet: true}
}

func (v NullableKalturaDexterIntegrationJobProviderData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaDexterIntegrationJobProviderData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


