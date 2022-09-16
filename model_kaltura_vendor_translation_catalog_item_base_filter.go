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

// KalturaVendorTranslationCatalogItemBaseFilter `abstract`
type KalturaVendorTranslationCatalogItemBaseFilter struct {
	KalturaVendorCaptionsCatalogItemFilter
}

// NewKalturaVendorTranslationCatalogItemBaseFilter instantiates a new KalturaVendorTranslationCatalogItemBaseFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaVendorTranslationCatalogItemBaseFilter() *KalturaVendorTranslationCatalogItemBaseFilter {
	this := KalturaVendorTranslationCatalogItemBaseFilter{}
	return &this
}

// NewKalturaVendorTranslationCatalogItemBaseFilterWithDefaults instantiates a new KalturaVendorTranslationCatalogItemBaseFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaVendorTranslationCatalogItemBaseFilterWithDefaults() *KalturaVendorTranslationCatalogItemBaseFilter {
	this := KalturaVendorTranslationCatalogItemBaseFilter{}
	return &this
}

func (o KalturaVendorTranslationCatalogItemBaseFilter) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedKalturaVendorCaptionsCatalogItemFilter, errKalturaVendorCaptionsCatalogItemFilter := json.Marshal(o.KalturaVendorCaptionsCatalogItemFilter)
	if errKalturaVendorCaptionsCatalogItemFilter != nil {
		return []byte{}, errKalturaVendorCaptionsCatalogItemFilter
	}
	errKalturaVendorCaptionsCatalogItemFilter = json.Unmarshal([]byte(serializedKalturaVendorCaptionsCatalogItemFilter), &toSerialize)
	if errKalturaVendorCaptionsCatalogItemFilter != nil {
		return []byte{}, errKalturaVendorCaptionsCatalogItemFilter
	}
	return json.Marshal(toSerialize)
}

type NullableKalturaVendorTranslationCatalogItemBaseFilter struct {
	value *KalturaVendorTranslationCatalogItemBaseFilter
	isSet bool
}

func (v NullableKalturaVendorTranslationCatalogItemBaseFilter) Get() *KalturaVendorTranslationCatalogItemBaseFilter {
	return v.value
}

func (v *NullableKalturaVendorTranslationCatalogItemBaseFilter) Set(val *KalturaVendorTranslationCatalogItemBaseFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaVendorTranslationCatalogItemBaseFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaVendorTranslationCatalogItemBaseFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaVendorTranslationCatalogItemBaseFilter(val *KalturaVendorTranslationCatalogItemBaseFilter) *NullableKalturaVendorTranslationCatalogItemBaseFilter {
	return &NullableKalturaVendorTranslationCatalogItemBaseFilter{value: val, isSet: true}
}

func (v NullableKalturaVendorTranslationCatalogItemBaseFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaVendorTranslationCatalogItemBaseFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

