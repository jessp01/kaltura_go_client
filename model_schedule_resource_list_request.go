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

// ScheduleResourceListRequest struct for ScheduleResourceListRequest
type ScheduleResourceListRequest struct {
	Filter *KalturaScheduleResourceFilter `json:"filter,omitempty"`
	Pager *KalturaFilterPager `json:"pager,omitempty"`
}

// NewScheduleResourceListRequest instantiates a new ScheduleResourceListRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewScheduleResourceListRequest() *ScheduleResourceListRequest {
	this := ScheduleResourceListRequest{}
	return &this
}

// NewScheduleResourceListRequestWithDefaults instantiates a new ScheduleResourceListRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewScheduleResourceListRequestWithDefaults() *ScheduleResourceListRequest {
	this := ScheduleResourceListRequest{}
	return &this
}

// GetFilter returns the Filter field value if set, zero value otherwise.
func (o *ScheduleResourceListRequest) GetFilter() KalturaScheduleResourceFilter {
	if o == nil || o.Filter == nil {
		var ret KalturaScheduleResourceFilter
		return ret
	}
	return *o.Filter
}

// GetFilterOk returns a tuple with the Filter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScheduleResourceListRequest) GetFilterOk() (*KalturaScheduleResourceFilter, bool) {
	if o == nil || o.Filter == nil {
		return nil, false
	}
	return o.Filter, true
}

// HasFilter returns a boolean if a field has been set.
func (o *ScheduleResourceListRequest) HasFilter() bool {
	if o != nil && o.Filter != nil {
		return true
	}

	return false
}

// SetFilter gets a reference to the given KalturaScheduleResourceFilter and assigns it to the Filter field.
func (o *ScheduleResourceListRequest) SetFilter(v KalturaScheduleResourceFilter) {
	o.Filter = &v
}

// GetPager returns the Pager field value if set, zero value otherwise.
func (o *ScheduleResourceListRequest) GetPager() KalturaFilterPager {
	if o == nil || o.Pager == nil {
		var ret KalturaFilterPager
		return ret
	}
	return *o.Pager
}

// GetPagerOk returns a tuple with the Pager field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScheduleResourceListRequest) GetPagerOk() (*KalturaFilterPager, bool) {
	if o == nil || o.Pager == nil {
		return nil, false
	}
	return o.Pager, true
}

// HasPager returns a boolean if a field has been set.
func (o *ScheduleResourceListRequest) HasPager() bool {
	if o != nil && o.Pager != nil {
		return true
	}

	return false
}

// SetPager gets a reference to the given KalturaFilterPager and assigns it to the Pager field.
func (o *ScheduleResourceListRequest) SetPager(v KalturaFilterPager) {
	o.Pager = &v
}

func (o ScheduleResourceListRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Filter != nil {
		toSerialize["filter"] = o.Filter
	}
	if o.Pager != nil {
		toSerialize["pager"] = o.Pager
	}
	return json.Marshal(toSerialize)
}

type NullableScheduleResourceListRequest struct {
	value *ScheduleResourceListRequest
	isSet bool
}

func (v NullableScheduleResourceListRequest) Get() *ScheduleResourceListRequest {
	return v.value
}

func (v *NullableScheduleResourceListRequest) Set(val *ScheduleResourceListRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableScheduleResourceListRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableScheduleResourceListRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableScheduleResourceListRequest(val *ScheduleResourceListRequest) *NullableScheduleResourceListRequest {
	return &NullableScheduleResourceListRequest{value: val, isSet: true}
}

func (v NullableScheduleResourceListRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableScheduleResourceListRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


