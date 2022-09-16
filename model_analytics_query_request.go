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

// AnalyticsQueryRequest struct for AnalyticsQueryRequest
type AnalyticsQueryRequest struct {
	Filter KalturaAnalyticsFilter `json:"filter"`
	Pager *KalturaFilterPager `json:"pager,omitempty"`
}

// NewAnalyticsQueryRequest instantiates a new AnalyticsQueryRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAnalyticsQueryRequest(filter KalturaAnalyticsFilter) *AnalyticsQueryRequest {
	this := AnalyticsQueryRequest{}
	this.Filter = filter
	return &this
}

// NewAnalyticsQueryRequestWithDefaults instantiates a new AnalyticsQueryRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAnalyticsQueryRequestWithDefaults() *AnalyticsQueryRequest {
	this := AnalyticsQueryRequest{}
	return &this
}

// GetFilter returns the Filter field value
func (o *AnalyticsQueryRequest) GetFilter() KalturaAnalyticsFilter {
	if o == nil {
		var ret KalturaAnalyticsFilter
		return ret
	}

	return o.Filter
}

// GetFilterOk returns a tuple with the Filter field value
// and a boolean to check if the value has been set.
func (o *AnalyticsQueryRequest) GetFilterOk() (*KalturaAnalyticsFilter, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Filter, true
}

// SetFilter sets field value
func (o *AnalyticsQueryRequest) SetFilter(v KalturaAnalyticsFilter) {
	o.Filter = v
}

// GetPager returns the Pager field value if set, zero value otherwise.
func (o *AnalyticsQueryRequest) GetPager() KalturaFilterPager {
	if o == nil || o.Pager == nil {
		var ret KalturaFilterPager
		return ret
	}
	return *o.Pager
}

// GetPagerOk returns a tuple with the Pager field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnalyticsQueryRequest) GetPagerOk() (*KalturaFilterPager, bool) {
	if o == nil || o.Pager == nil {
		return nil, false
	}
	return o.Pager, true
}

// HasPager returns a boolean if a field has been set.
func (o *AnalyticsQueryRequest) HasPager() bool {
	if o != nil && o.Pager != nil {
		return true
	}

	return false
}

// SetPager gets a reference to the given KalturaFilterPager and assigns it to the Pager field.
func (o *AnalyticsQueryRequest) SetPager(v KalturaFilterPager) {
	o.Pager = &v
}

func (o AnalyticsQueryRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["filter"] = o.Filter
	}
	if o.Pager != nil {
		toSerialize["pager"] = o.Pager
	}
	return json.Marshal(toSerialize)
}

type NullableAnalyticsQueryRequest struct {
	value *AnalyticsQueryRequest
	isSet bool
}

func (v NullableAnalyticsQueryRequest) Get() *AnalyticsQueryRequest {
	return v.value
}

func (v *NullableAnalyticsQueryRequest) Set(val *AnalyticsQueryRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAnalyticsQueryRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAnalyticsQueryRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAnalyticsQueryRequest(val *AnalyticsQueryRequest) *NullableAnalyticsQueryRequest {
	return &NullableAnalyticsQueryRequest{value: val, isSet: true}
}

func (v NullableAnalyticsQueryRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAnalyticsQueryRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


