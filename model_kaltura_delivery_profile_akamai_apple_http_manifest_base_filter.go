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

// KalturaDeliveryProfileAkamaiAppleHttpManifestBaseFilter `abstract`
type KalturaDeliveryProfileAkamaiAppleHttpManifestBaseFilter struct {
	KalturaDeliveryProfileFilter
}

// NewKalturaDeliveryProfileAkamaiAppleHttpManifestBaseFilter instantiates a new KalturaDeliveryProfileAkamaiAppleHttpManifestBaseFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaDeliveryProfileAkamaiAppleHttpManifestBaseFilter() *KalturaDeliveryProfileAkamaiAppleHttpManifestBaseFilter {
	this := KalturaDeliveryProfileAkamaiAppleHttpManifestBaseFilter{}
	return &this
}

// NewKalturaDeliveryProfileAkamaiAppleHttpManifestBaseFilterWithDefaults instantiates a new KalturaDeliveryProfileAkamaiAppleHttpManifestBaseFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaDeliveryProfileAkamaiAppleHttpManifestBaseFilterWithDefaults() *KalturaDeliveryProfileAkamaiAppleHttpManifestBaseFilter {
	this := KalturaDeliveryProfileAkamaiAppleHttpManifestBaseFilter{}
	return &this
}

func (o KalturaDeliveryProfileAkamaiAppleHttpManifestBaseFilter) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedKalturaDeliveryProfileFilter, errKalturaDeliveryProfileFilter := json.Marshal(o.KalturaDeliveryProfileFilter)
	if errKalturaDeliveryProfileFilter != nil {
		return []byte{}, errKalturaDeliveryProfileFilter
	}
	errKalturaDeliveryProfileFilter = json.Unmarshal([]byte(serializedKalturaDeliveryProfileFilter), &toSerialize)
	if errKalturaDeliveryProfileFilter != nil {
		return []byte{}, errKalturaDeliveryProfileFilter
	}
	return json.Marshal(toSerialize)
}

type NullableKalturaDeliveryProfileAkamaiAppleHttpManifestBaseFilter struct {
	value *KalturaDeliveryProfileAkamaiAppleHttpManifestBaseFilter
	isSet bool
}

func (v NullableKalturaDeliveryProfileAkamaiAppleHttpManifestBaseFilter) Get() *KalturaDeliveryProfileAkamaiAppleHttpManifestBaseFilter {
	return v.value
}

func (v *NullableKalturaDeliveryProfileAkamaiAppleHttpManifestBaseFilter) Set(val *KalturaDeliveryProfileAkamaiAppleHttpManifestBaseFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaDeliveryProfileAkamaiAppleHttpManifestBaseFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaDeliveryProfileAkamaiAppleHttpManifestBaseFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaDeliveryProfileAkamaiAppleHttpManifestBaseFilter(val *KalturaDeliveryProfileAkamaiAppleHttpManifestBaseFilter) *NullableKalturaDeliveryProfileAkamaiAppleHttpManifestBaseFilter {
	return &NullableKalturaDeliveryProfileAkamaiAppleHttpManifestBaseFilter{value: val, isSet: true}
}

func (v NullableKalturaDeliveryProfileAkamaiAppleHttpManifestBaseFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaDeliveryProfileAkamaiAppleHttpManifestBaseFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


