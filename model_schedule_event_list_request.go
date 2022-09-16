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

// ScheduleEventListRequest struct for ScheduleEventListRequest
type ScheduleEventListRequest struct {
	Filter *KalturaScheduleEventFilter `json:"filter,omitempty"`
	Pager *KalturaFilterPager `json:"pager,omitempty"`
}

// NewScheduleEventListRequest instantiates a new ScheduleEventListRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewScheduleEventListRequest() *ScheduleEventListRequest {
	this := ScheduleEventListRequest{}
	return &this
}

// NewScheduleEventListRequestWithDefaults instantiates a new ScheduleEventListRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewScheduleEventListRequestWithDefaults() *ScheduleEventListRequest {
	this := ScheduleEventListRequest{}
	return &this
}

// GetFilter returns the Filter field value if set, zero value otherwise.
func (o *ScheduleEventListRequest) GetFilter() KalturaScheduleEventFilter {
	if o == nil || o.Filter == nil {
		var ret KalturaScheduleEventFilter
		return ret
	}
	return *o.Filter
}

// GetFilterOk returns a tuple with the Filter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScheduleEventListRequest) GetFilterOk() (*KalturaScheduleEventFilter, bool) {
	if o == nil || o.Filter == nil {
		return nil, false
	}
	return o.Filter, true
}

// HasFilter returns a boolean if a field has been set.
func (o *ScheduleEventListRequest) HasFilter() bool {
	if o != nil && o.Filter != nil {
		return true
	}

	return false
}

// SetFilter gets a reference to the given KalturaScheduleEventFilter and assigns it to the Filter field.
func (o *ScheduleEventListRequest) SetFilter(v KalturaScheduleEventFilter) {
	o.Filter = &v
}

// GetPager returns the Pager field value if set, zero value otherwise.
func (o *ScheduleEventListRequest) GetPager() KalturaFilterPager {
	if o == nil || o.Pager == nil {
		var ret KalturaFilterPager
		return ret
	}
	return *o.Pager
}

// GetPagerOk returns a tuple with the Pager field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScheduleEventListRequest) GetPagerOk() (*KalturaFilterPager, bool) {
	if o == nil || o.Pager == nil {
		return nil, false
	}
	return o.Pager, true
}

// HasPager returns a boolean if a field has been set.
func (o *ScheduleEventListRequest) HasPager() bool {
	if o != nil && o.Pager != nil {
		return true
	}

	return false
}

// SetPager gets a reference to the given KalturaFilterPager and assigns it to the Pager field.
func (o *ScheduleEventListRequest) SetPager(v KalturaFilterPager) {
	o.Pager = &v
}

func (o ScheduleEventListRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Filter != nil {
		toSerialize["filter"] = o.Filter
	}
	if o.Pager != nil {
		toSerialize["pager"] = o.Pager
	}
	return json.Marshal(toSerialize)
}

type NullableScheduleEventListRequest struct {
	value *ScheduleEventListRequest
	isSet bool
}

func (v NullableScheduleEventListRequest) Get() *ScheduleEventListRequest {
	return v.value
}

func (v *NullableScheduleEventListRequest) Set(val *ScheduleEventListRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableScheduleEventListRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableScheduleEventListRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableScheduleEventListRequest(val *ScheduleEventListRequest) *NullableScheduleEventListRequest {
	return &NullableScheduleEventListRequest{value: val, isSet: true}
}

func (v NullableScheduleEventListRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableScheduleEventListRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


