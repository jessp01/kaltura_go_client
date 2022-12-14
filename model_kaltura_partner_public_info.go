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

// KalturaPartnerPublicInfo struct for KalturaPartnerPublicInfo
type KalturaPartnerPublicInfo struct {
	AnalyticsPersistentSessionId *bool `json:"analyticsPersistentSessionId,omitempty"`
	AnalyticsUrl *string `json:"analyticsUrl,omitempty"`
	OttEnvironmentUrl *string `json:"ottEnvironmentUrl,omitempty"`
}

// NewKalturaPartnerPublicInfo instantiates a new KalturaPartnerPublicInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaPartnerPublicInfo() *KalturaPartnerPublicInfo {
	this := KalturaPartnerPublicInfo{}
	return &this
}

// NewKalturaPartnerPublicInfoWithDefaults instantiates a new KalturaPartnerPublicInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaPartnerPublicInfoWithDefaults() *KalturaPartnerPublicInfo {
	this := KalturaPartnerPublicInfo{}
	return &this
}

// GetAnalyticsPersistentSessionId returns the AnalyticsPersistentSessionId field value if set, zero value otherwise.
func (o *KalturaPartnerPublicInfo) GetAnalyticsPersistentSessionId() bool {
	if o == nil || o.AnalyticsPersistentSessionId == nil {
		var ret bool
		return ret
	}
	return *o.AnalyticsPersistentSessionId
}

// GetAnalyticsPersistentSessionIdOk returns a tuple with the AnalyticsPersistentSessionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaPartnerPublicInfo) GetAnalyticsPersistentSessionIdOk() (*bool, bool) {
	if o == nil || o.AnalyticsPersistentSessionId == nil {
		return nil, false
	}
	return o.AnalyticsPersistentSessionId, true
}

// HasAnalyticsPersistentSessionId returns a boolean if a field has been set.
func (o *KalturaPartnerPublicInfo) HasAnalyticsPersistentSessionId() bool {
	if o != nil && o.AnalyticsPersistentSessionId != nil {
		return true
	}

	return false
}

// SetAnalyticsPersistentSessionId gets a reference to the given bool and assigns it to the AnalyticsPersistentSessionId field.
func (o *KalturaPartnerPublicInfo) SetAnalyticsPersistentSessionId(v bool) {
	o.AnalyticsPersistentSessionId = &v
}

// GetAnalyticsUrl returns the AnalyticsUrl field value if set, zero value otherwise.
func (o *KalturaPartnerPublicInfo) GetAnalyticsUrl() string {
	if o == nil || o.AnalyticsUrl == nil {
		var ret string
		return ret
	}
	return *o.AnalyticsUrl
}

// GetAnalyticsUrlOk returns a tuple with the AnalyticsUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaPartnerPublicInfo) GetAnalyticsUrlOk() (*string, bool) {
	if o == nil || o.AnalyticsUrl == nil {
		return nil, false
	}
	return o.AnalyticsUrl, true
}

// HasAnalyticsUrl returns a boolean if a field has been set.
func (o *KalturaPartnerPublicInfo) HasAnalyticsUrl() bool {
	if o != nil && o.AnalyticsUrl != nil {
		return true
	}

	return false
}

// SetAnalyticsUrl gets a reference to the given string and assigns it to the AnalyticsUrl field.
func (o *KalturaPartnerPublicInfo) SetAnalyticsUrl(v string) {
	o.AnalyticsUrl = &v
}

// GetOttEnvironmentUrl returns the OttEnvironmentUrl field value if set, zero value otherwise.
func (o *KalturaPartnerPublicInfo) GetOttEnvironmentUrl() string {
	if o == nil || o.OttEnvironmentUrl == nil {
		var ret string
		return ret
	}
	return *o.OttEnvironmentUrl
}

// GetOttEnvironmentUrlOk returns a tuple with the OttEnvironmentUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaPartnerPublicInfo) GetOttEnvironmentUrlOk() (*string, bool) {
	if o == nil || o.OttEnvironmentUrl == nil {
		return nil, false
	}
	return o.OttEnvironmentUrl, true
}

// HasOttEnvironmentUrl returns a boolean if a field has been set.
func (o *KalturaPartnerPublicInfo) HasOttEnvironmentUrl() bool {
	if o != nil && o.OttEnvironmentUrl != nil {
		return true
	}

	return false
}

// SetOttEnvironmentUrl gets a reference to the given string and assigns it to the OttEnvironmentUrl field.
func (o *KalturaPartnerPublicInfo) SetOttEnvironmentUrl(v string) {
	o.OttEnvironmentUrl = &v
}

func (o KalturaPartnerPublicInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AnalyticsPersistentSessionId != nil {
		toSerialize["analyticsPersistentSessionId"] = o.AnalyticsPersistentSessionId
	}
	if o.AnalyticsUrl != nil {
		toSerialize["analyticsUrl"] = o.AnalyticsUrl
	}
	if o.OttEnvironmentUrl != nil {
		toSerialize["ottEnvironmentUrl"] = o.OttEnvironmentUrl
	}
	return json.Marshal(toSerialize)
}

type NullableKalturaPartnerPublicInfo struct {
	value *KalturaPartnerPublicInfo
	isSet bool
}

func (v NullableKalturaPartnerPublicInfo) Get() *KalturaPartnerPublicInfo {
	return v.value
}

func (v *NullableKalturaPartnerPublicInfo) Set(val *KalturaPartnerPublicInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaPartnerPublicInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaPartnerPublicInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaPartnerPublicInfo(val *KalturaPartnerPublicInfo) *NullableKalturaPartnerPublicInfo {
	return &NullableKalturaPartnerPublicInfo{value: val, isSet: true}
}

func (v NullableKalturaPartnerPublicInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaPartnerPublicInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


