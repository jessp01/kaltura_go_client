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

// GroupListRequest struct for GroupListRequest
type GroupListRequest struct {
	Filter *KalturaGroupFilter `json:"filter,omitempty"`
	Pager *KalturaFilterPager `json:"pager,omitempty"`
}

// NewGroupListRequest instantiates a new GroupListRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGroupListRequest() *GroupListRequest {
	this := GroupListRequest{}
	return &this
}

// NewGroupListRequestWithDefaults instantiates a new GroupListRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGroupListRequestWithDefaults() *GroupListRequest {
	this := GroupListRequest{}
	return &this
}

// GetFilter returns the Filter field value if set, zero value otherwise.
func (o *GroupListRequest) GetFilter() KalturaGroupFilter {
	if o == nil || o.Filter == nil {
		var ret KalturaGroupFilter
		return ret
	}
	return *o.Filter
}

// GetFilterOk returns a tuple with the Filter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupListRequest) GetFilterOk() (*KalturaGroupFilter, bool) {
	if o == nil || o.Filter == nil {
		return nil, false
	}
	return o.Filter, true
}

// HasFilter returns a boolean if a field has been set.
func (o *GroupListRequest) HasFilter() bool {
	if o != nil && o.Filter != nil {
		return true
	}

	return false
}

// SetFilter gets a reference to the given KalturaGroupFilter and assigns it to the Filter field.
func (o *GroupListRequest) SetFilter(v KalturaGroupFilter) {
	o.Filter = &v
}

// GetPager returns the Pager field value if set, zero value otherwise.
func (o *GroupListRequest) GetPager() KalturaFilterPager {
	if o == nil || o.Pager == nil {
		var ret KalturaFilterPager
		return ret
	}
	return *o.Pager
}

// GetPagerOk returns a tuple with the Pager field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupListRequest) GetPagerOk() (*KalturaFilterPager, bool) {
	if o == nil || o.Pager == nil {
		return nil, false
	}
	return o.Pager, true
}

// HasPager returns a boolean if a field has been set.
func (o *GroupListRequest) HasPager() bool {
	if o != nil && o.Pager != nil {
		return true
	}

	return false
}

// SetPager gets a reference to the given KalturaFilterPager and assigns it to the Pager field.
func (o *GroupListRequest) SetPager(v KalturaFilterPager) {
	o.Pager = &v
}

func (o GroupListRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Filter != nil {
		toSerialize["filter"] = o.Filter
	}
	if o.Pager != nil {
		toSerialize["pager"] = o.Pager
	}
	return json.Marshal(toSerialize)
}

type NullableGroupListRequest struct {
	value *GroupListRequest
	isSet bool
}

func (v NullableGroupListRequest) Get() *GroupListRequest {
	return v.value
}

func (v *NullableGroupListRequest) Set(val *GroupListRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableGroupListRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableGroupListRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGroupListRequest(val *GroupListRequest) *NullableGroupListRequest {
	return &NullableGroupListRequest{value: val, isSet: true}
}

func (v NullableGroupListRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGroupListRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


