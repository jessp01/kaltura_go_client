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

// KalturaVendorTranslationCatalogItemFilter struct for KalturaVendorTranslationCatalogItemFilter
type KalturaVendorTranslationCatalogItemFilter struct {
	KalturaVendorTranslationCatalogItemBaseFilter
}

// NewKalturaVendorTranslationCatalogItemFilter instantiates a new KalturaVendorTranslationCatalogItemFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaVendorTranslationCatalogItemFilter() *KalturaVendorTranslationCatalogItemFilter {
	this := KalturaVendorTranslationCatalogItemFilter{}
	return &this
}

// NewKalturaVendorTranslationCatalogItemFilterWithDefaults instantiates a new KalturaVendorTranslationCatalogItemFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaVendorTranslationCatalogItemFilterWithDefaults() *KalturaVendorTranslationCatalogItemFilter {
	this := KalturaVendorTranslationCatalogItemFilter{}
	return &this
}

func (o KalturaVendorTranslationCatalogItemFilter) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedKalturaVendorTranslationCatalogItemBaseFilter, errKalturaVendorTranslationCatalogItemBaseFilter := json.Marshal(o.KalturaVendorTranslationCatalogItemBaseFilter)
	if errKalturaVendorTranslationCatalogItemBaseFilter != nil {
		return []byte{}, errKalturaVendorTranslationCatalogItemBaseFilter
	}
	errKalturaVendorTranslationCatalogItemBaseFilter = json.Unmarshal([]byte(serializedKalturaVendorTranslationCatalogItemBaseFilter), &toSerialize)
	if errKalturaVendorTranslationCatalogItemBaseFilter != nil {
		return []byte{}, errKalturaVendorTranslationCatalogItemBaseFilter
	}
	return json.Marshal(toSerialize)
}

type NullableKalturaVendorTranslationCatalogItemFilter struct {
	value *KalturaVendorTranslationCatalogItemFilter
	isSet bool
}

func (v NullableKalturaVendorTranslationCatalogItemFilter) Get() *KalturaVendorTranslationCatalogItemFilter {
	return v.value
}

func (v *NullableKalturaVendorTranslationCatalogItemFilter) Set(val *KalturaVendorTranslationCatalogItemFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaVendorTranslationCatalogItemFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaVendorTranslationCatalogItemFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaVendorTranslationCatalogItemFilter(val *KalturaVendorTranslationCatalogItemFilter) *NullableKalturaVendorTranslationCatalogItemFilter {
	return &NullableKalturaVendorTranslationCatalogItemFilter{value: val, isSet: true}
}

func (v NullableKalturaVendorTranslationCatalogItemFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaVendorTranslationCatalogItemFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


