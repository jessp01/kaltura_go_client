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

// KalturaVendorAlignmentCatalogItemFilter struct for KalturaVendorAlignmentCatalogItemFilter
type KalturaVendorAlignmentCatalogItemFilter struct {
	KalturaVendorCaptionsCatalogItemBaseFilter
}

// NewKalturaVendorAlignmentCatalogItemFilter instantiates a new KalturaVendorAlignmentCatalogItemFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaVendorAlignmentCatalogItemFilter() *KalturaVendorAlignmentCatalogItemFilter {
	this := KalturaVendorAlignmentCatalogItemFilter{}
	return &this
}

// NewKalturaVendorAlignmentCatalogItemFilterWithDefaults instantiates a new KalturaVendorAlignmentCatalogItemFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaVendorAlignmentCatalogItemFilterWithDefaults() *KalturaVendorAlignmentCatalogItemFilter {
	this := KalturaVendorAlignmentCatalogItemFilter{}
	return &this
}

func (o KalturaVendorAlignmentCatalogItemFilter) MarshalJSON() ([]byte, error) {
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

type NullableKalturaVendorAlignmentCatalogItemFilter struct {
	value *KalturaVendorAlignmentCatalogItemFilter
	isSet bool
}

func (v NullableKalturaVendorAlignmentCatalogItemFilter) Get() *KalturaVendorAlignmentCatalogItemFilter {
	return v.value
}

func (v *NullableKalturaVendorAlignmentCatalogItemFilter) Set(val *KalturaVendorAlignmentCatalogItemFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaVendorAlignmentCatalogItemFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaVendorAlignmentCatalogItemFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaVendorAlignmentCatalogItemFilter(val *KalturaVendorAlignmentCatalogItemFilter) *NullableKalturaVendorAlignmentCatalogItemFilter {
	return &NullableKalturaVendorAlignmentCatalogItemFilter{value: val, isSet: true}
}

func (v NullableKalturaVendorAlignmentCatalogItemFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaVendorAlignmentCatalogItemFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


