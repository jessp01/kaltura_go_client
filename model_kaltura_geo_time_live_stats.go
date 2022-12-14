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

// KalturaGeoTimeLiveStats struct for KalturaGeoTimeLiveStats
type KalturaGeoTimeLiveStats struct {
	KalturaEntryLiveStats
}

// NewKalturaGeoTimeLiveStats instantiates a new KalturaGeoTimeLiveStats object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaGeoTimeLiveStats() *KalturaGeoTimeLiveStats {
	this := KalturaGeoTimeLiveStats{}
	return &this
}

// NewKalturaGeoTimeLiveStatsWithDefaults instantiates a new KalturaGeoTimeLiveStats object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaGeoTimeLiveStatsWithDefaults() *KalturaGeoTimeLiveStats {
	this := KalturaGeoTimeLiveStats{}
	return &this
}

func (o KalturaGeoTimeLiveStats) MarshalJSON() ([]byte, error) {
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

type NullableKalturaGeoTimeLiveStats struct {
	value *KalturaGeoTimeLiveStats
	isSet bool
}

func (v NullableKalturaGeoTimeLiveStats) Get() *KalturaGeoTimeLiveStats {
	return v.value
}

func (v *NullableKalturaGeoTimeLiveStats) Set(val *KalturaGeoTimeLiveStats) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaGeoTimeLiveStats) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaGeoTimeLiveStats) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaGeoTimeLiveStats(val *KalturaGeoTimeLiveStats) *NullableKalturaGeoTimeLiveStats {
	return &NullableKalturaGeoTimeLiveStats{value: val, isSet: true}
}

func (v NullableKalturaGeoTimeLiveStats) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaGeoTimeLiveStats) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


