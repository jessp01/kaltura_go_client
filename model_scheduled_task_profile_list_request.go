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

// ScheduledTaskProfileListRequest struct for ScheduledTaskProfileListRequest
type ScheduledTaskProfileListRequest struct {
	Filter *KalturaScheduledTaskProfileFilter `json:"filter,omitempty"`
	Pager *KalturaFilterPager `json:"pager,omitempty"`
}

// NewScheduledTaskProfileListRequest instantiates a new ScheduledTaskProfileListRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewScheduledTaskProfileListRequest() *ScheduledTaskProfileListRequest {
	this := ScheduledTaskProfileListRequest{}
	return &this
}

// NewScheduledTaskProfileListRequestWithDefaults instantiates a new ScheduledTaskProfileListRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewScheduledTaskProfileListRequestWithDefaults() *ScheduledTaskProfileListRequest {
	this := ScheduledTaskProfileListRequest{}
	return &this
}

// GetFilter returns the Filter field value if set, zero value otherwise.
func (o *ScheduledTaskProfileListRequest) GetFilter() KalturaScheduledTaskProfileFilter {
	if o == nil || o.Filter == nil {
		var ret KalturaScheduledTaskProfileFilter
		return ret
	}
	return *o.Filter
}

// GetFilterOk returns a tuple with the Filter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScheduledTaskProfileListRequest) GetFilterOk() (*KalturaScheduledTaskProfileFilter, bool) {
	if o == nil || o.Filter == nil {
		return nil, false
	}
	return o.Filter, true
}

// HasFilter returns a boolean if a field has been set.
func (o *ScheduledTaskProfileListRequest) HasFilter() bool {
	if o != nil && o.Filter != nil {
		return true
	}

	return false
}

// SetFilter gets a reference to the given KalturaScheduledTaskProfileFilter and assigns it to the Filter field.
func (o *ScheduledTaskProfileListRequest) SetFilter(v KalturaScheduledTaskProfileFilter) {
	o.Filter = &v
}

// GetPager returns the Pager field value if set, zero value otherwise.
func (o *ScheduledTaskProfileListRequest) GetPager() KalturaFilterPager {
	if o == nil || o.Pager == nil {
		var ret KalturaFilterPager
		return ret
	}
	return *o.Pager
}

// GetPagerOk returns a tuple with the Pager field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScheduledTaskProfileListRequest) GetPagerOk() (*KalturaFilterPager, bool) {
	if o == nil || o.Pager == nil {
		return nil, false
	}
	return o.Pager, true
}

// HasPager returns a boolean if a field has been set.
func (o *ScheduledTaskProfileListRequest) HasPager() bool {
	if o != nil && o.Pager != nil {
		return true
	}

	return false
}

// SetPager gets a reference to the given KalturaFilterPager and assigns it to the Pager field.
func (o *ScheduledTaskProfileListRequest) SetPager(v KalturaFilterPager) {
	o.Pager = &v
}

func (o ScheduledTaskProfileListRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Filter != nil {
		toSerialize["filter"] = o.Filter
	}
	if o.Pager != nil {
		toSerialize["pager"] = o.Pager
	}
	return json.Marshal(toSerialize)
}

type NullableScheduledTaskProfileListRequest struct {
	value *ScheduledTaskProfileListRequest
	isSet bool
}

func (v NullableScheduledTaskProfileListRequest) Get() *ScheduledTaskProfileListRequest {
	return v.value
}

func (v *NullableScheduledTaskProfileListRequest) Set(val *ScheduledTaskProfileListRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableScheduledTaskProfileListRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableScheduledTaskProfileListRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableScheduledTaskProfileListRequest(val *ScheduledTaskProfileListRequest) *NullableScheduledTaskProfileListRequest {
	return &NullableScheduledTaskProfileListRequest{value: val, isSet: true}
}

func (v NullableScheduledTaskProfileListRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableScheduledTaskProfileListRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


