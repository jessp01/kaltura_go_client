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

// KalturaVendorChapteringCatalogItem struct for KalturaVendorChapteringCatalogItem
type KalturaVendorChapteringCatalogItem struct {
	KalturaVendorCatalogItem
}

// NewKalturaVendorChapteringCatalogItem instantiates a new KalturaVendorChapteringCatalogItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaVendorChapteringCatalogItem() *KalturaVendorChapteringCatalogItem {
	this := KalturaVendorChapteringCatalogItem{}
	return &this
}

// NewKalturaVendorChapteringCatalogItemWithDefaults instantiates a new KalturaVendorChapteringCatalogItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaVendorChapteringCatalogItemWithDefaults() *KalturaVendorChapteringCatalogItem {
	this := KalturaVendorChapteringCatalogItem{}
	return &this
}

func (o KalturaVendorChapteringCatalogItem) MarshalJSON() ([]byte, error) {
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

type NullableKalturaVendorChapteringCatalogItem struct {
	value *KalturaVendorChapteringCatalogItem
	isSet bool
}

func (v NullableKalturaVendorChapteringCatalogItem) Get() *KalturaVendorChapteringCatalogItem {
	return v.value
}

func (v *NullableKalturaVendorChapteringCatalogItem) Set(val *KalturaVendorChapteringCatalogItem) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaVendorChapteringCatalogItem) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaVendorChapteringCatalogItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaVendorChapteringCatalogItem(val *KalturaVendorChapteringCatalogItem) *NullableKalturaVendorChapteringCatalogItem {
	return &NullableKalturaVendorChapteringCatalogItem{value: val, isSet: true}
}

func (v NullableKalturaVendorChapteringCatalogItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaVendorChapteringCatalogItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


