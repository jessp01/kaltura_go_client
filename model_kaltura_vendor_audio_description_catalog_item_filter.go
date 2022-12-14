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

// KalturaVendorAudioDescriptionCatalogItemFilter struct for KalturaVendorAudioDescriptionCatalogItemFilter
type KalturaVendorAudioDescriptionCatalogItemFilter struct {
	KalturaVendorCaptionsCatalogItemBaseFilter
}

// NewKalturaVendorAudioDescriptionCatalogItemFilter instantiates a new KalturaVendorAudioDescriptionCatalogItemFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaVendorAudioDescriptionCatalogItemFilter() *KalturaVendorAudioDescriptionCatalogItemFilter {
	this := KalturaVendorAudioDescriptionCatalogItemFilter{}
	return &this
}

// NewKalturaVendorAudioDescriptionCatalogItemFilterWithDefaults instantiates a new KalturaVendorAudioDescriptionCatalogItemFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaVendorAudioDescriptionCatalogItemFilterWithDefaults() *KalturaVendorAudioDescriptionCatalogItemFilter {
	this := KalturaVendorAudioDescriptionCatalogItemFilter{}
	return &this
}

func (o KalturaVendorAudioDescriptionCatalogItemFilter) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedKalturaVendorCaptionsCatalogItemBaseFilter, errKalturaVendorCaptionsCatalogItemBaseFilter := json.Marshal(o.KalturaVendorCaptionsCatalogItemBaseFilter)
	if errKalturaVendorCaptionsCatalogItemBaseFilter != nil {
		return []byte{}, errKalturaVendorCaptionsCatalogItemBaseFilter
	}
	errKalturaVendorCaptionsCatalogItemBaseFilter = json.Unmarshal([]byte(serializedKalturaVendorCaptionsCatalogItemBaseFilter), &toSerialize)
	if errKalturaVendorCaptionsCatalogItemBaseFilter != nil {
		return []byte{}, errKalturaVendorCaptionsCatalogItemBaseFilter
	}
	return json.Marshal(toSerialize)
}

type NullableKalturaVendorAudioDescriptionCatalogItemFilter struct {
	value *KalturaVendorAudioDescriptionCatalogItemFilter
	isSet bool
}

func (v NullableKalturaVendorAudioDescriptionCatalogItemFilter) Get() *KalturaVendorAudioDescriptionCatalogItemFilter {
	return v.value
}

func (v *NullableKalturaVendorAudioDescriptionCatalogItemFilter) Set(val *KalturaVendorAudioDescriptionCatalogItemFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaVendorAudioDescriptionCatalogItemFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaVendorAudioDescriptionCatalogItemFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaVendorAudioDescriptionCatalogItemFilter(val *KalturaVendorAudioDescriptionCatalogItemFilter) *NullableKalturaVendorAudioDescriptionCatalogItemFilter {
	return &NullableKalturaVendorAudioDescriptionCatalogItemFilter{value: val, isSet: true}
}

func (v NullableKalturaVendorAudioDescriptionCatalogItemFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaVendorAudioDescriptionCatalogItemFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


