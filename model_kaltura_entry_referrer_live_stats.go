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

// KalturaEntryReferrerLiveStats struct for KalturaEntryReferrerLiveStats
type KalturaEntryReferrerLiveStats struct {
	KalturaEntryLiveStats
}

// NewKalturaEntryReferrerLiveStats instantiates a new KalturaEntryReferrerLiveStats object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaEntryReferrerLiveStats() *KalturaEntryReferrerLiveStats {
	this := KalturaEntryReferrerLiveStats{}
	return &this
}

// NewKalturaEntryReferrerLiveStatsWithDefaults instantiates a new KalturaEntryReferrerLiveStats object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaEntryReferrerLiveStatsWithDefaults() *KalturaEntryReferrerLiveStats {
	this := KalturaEntryReferrerLiveStats{}
	return &this
}

func (o KalturaEntryReferrerLiveStats) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedKalturaEntryLiveStats, errKalturaEntryLiveStats := json.Marshal(o.KalturaEntryLiveStats)
	if errKalturaEntryLiveStats != nil {
		return []byte{}, errKalturaEntryLiveStats
	}
	errKalturaEntryLiveStats = json.Unmarshal([]byte(serializedKalturaEntryLiveStats), &toSerialize)
	if errKalturaEntryLiveStats != nil {
		return []byte{}, errKalturaEntryLiveStats
	}
	return json.Marshal(toSerialize)
}

type NullableKalturaEntryReferrerLiveStats struct {
	value *KalturaEntryReferrerLiveStats
	isSet bool
}

func (v NullableKalturaEntryReferrerLiveStats) Get() *KalturaEntryReferrerLiveStats {
	return v.value
}

func (v *NullableKalturaEntryReferrerLiveStats) Set(val *KalturaEntryReferrerLiveStats) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaEntryReferrerLiveStats) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaEntryReferrerLiveStats) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaEntryReferrerLiveStats(val *KalturaEntryReferrerLiveStats) *NullableKalturaEntryReferrerLiveStats {
	return &NullableKalturaEntryReferrerLiveStats{value: val, isSet: true}
}

func (v NullableKalturaEntryReferrerLiveStats) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaEntryReferrerLiveStats) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


