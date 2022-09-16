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

// KalturaCielo24JobProviderData struct for KalturaCielo24JobProviderData
type KalturaCielo24JobProviderData struct {
	KalturaIntegrationJobProviderData
}

// NewKalturaCielo24JobProviderData instantiates a new KalturaCielo24JobProviderData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaCielo24JobProviderData() *KalturaCielo24JobProviderData {
	this := KalturaCielo24JobProviderData{}
	return &this
}

// NewKalturaCielo24JobProviderDataWithDefaults instantiates a new KalturaCielo24JobProviderData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaCielo24JobProviderDataWithDefaults() *KalturaCielo24JobProviderData {
	this := KalturaCielo24JobProviderData{}
	return &this
}

func (o KalturaCielo24JobProviderData) MarshalJSON() ([]byte, error) {
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

type NullableKalturaCielo24JobProviderData struct {
	value *KalturaCielo24JobProviderData
	isSet bool
}

func (v NullableKalturaCielo24JobProviderData) Get() *KalturaCielo24JobProviderData {
	return v.value
}

func (v *NullableKalturaCielo24JobProviderData) Set(val *KalturaCielo24JobProviderData) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaCielo24JobProviderData) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaCielo24JobProviderData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaCielo24JobProviderData(val *KalturaCielo24JobProviderData) *NullableKalturaCielo24JobProviderData {
	return &NullableKalturaCielo24JobProviderData{value: val, isSet: true}
}

func (v NullableKalturaCielo24JobProviderData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaCielo24JobProviderData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

