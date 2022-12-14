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

// KalturaVendorCaptionsCatalogItem struct for KalturaVendorCaptionsCatalogItem
type KalturaVendorCaptionsCatalogItem struct {
	KalturaVendorCatalogItem
}

// NewKalturaVendorCaptionsCatalogItem instantiates a new KalturaVendorCaptionsCatalogItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaVendorCaptionsCatalogItem() *KalturaVendorCaptionsCatalogItem {
	this := KalturaVendorCaptionsCatalogItem{}
	return &this
}

// NewKalturaVendorCaptionsCatalogItemWithDefaults instantiates a new KalturaVendorCaptionsCatalogItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaVendorCaptionsCatalogItemWithDefaults() *KalturaVendorCaptionsCatalogItem {
	this := KalturaVendorCaptionsCatalogItem{}
	return &this
}

func (o KalturaVendorCaptionsCatalogItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedKalturaVendorCatalogItem, errKalturaVendorCatalogItem := json.Marshal(o.KalturaVendorCatalogItem)
	if errKalturaVendorCatalogItem != nil {
		return []byte{}, errKalturaVendorCatalogItem
	}
	errKalturaVendorCatalogItem = json.Unmarshal([]byte(serializedKalturaVendorCatalogItem), &toSerialize)
	if errKalturaVendorCatalogItem != nil {
		return []byte{}, errKalturaVendorCatalogItem
	}
	return json.Marshal(toSerialize)
}

type NullableKalturaVendorCaptionsCatalogItem struct {
	value *KalturaVendorCaptionsCatalogItem
	isSet bool
}

func (v NullableKalturaVendorCaptionsCatalogItem) Get() *KalturaVendorCaptionsCatalogItem {
	return v.value
}

func (v *NullableKalturaVendorCaptionsCatalogItem) Set(val *KalturaVendorCaptionsCatalogItem) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaVendorCaptionsCatalogItem) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaVendorCaptionsCatalogItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaVendorCaptionsCatalogItem(val *KalturaVendorCaptionsCatalogItem) *NullableKalturaVendorCaptionsCatalogItem {
	return &NullableKalturaVendorCaptionsCatalogItem{value: val, isSet: true}
}

func (v NullableKalturaVendorCaptionsCatalogItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaVendorCaptionsCatalogItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


