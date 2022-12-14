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

// EntryServerNodeListRequest struct for EntryServerNodeListRequest
type EntryServerNodeListRequest struct {
	Filter *KalturaEntryServerNodeFilter `json:"filter,omitempty"`
	Pager *KalturaFilterPager `json:"pager,omitempty"`
}

// NewEntryServerNodeListRequest instantiates a new EntryServerNodeListRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEntryServerNodeListRequest() *EntryServerNodeListRequest {
	this := EntryServerNodeListRequest{}
	return &this
}

// NewEntryServerNodeListRequestWithDefaults instantiates a new EntryServerNodeListRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEntryServerNodeListRequestWithDefaults() *EntryServerNodeListRequest {
	this := EntryServerNodeListRequest{}
	return &this
}

// GetFilter returns the Filter field value if set, zero value otherwise.
func (o *EntryServerNodeListRequest) GetFilter() KalturaEntryServerNodeFilter {
	if o == nil || o.Filter == nil {
		var ret KalturaEntryServerNodeFilter
		return ret
	}
	return *o.Filter
}

// GetFilterOk returns a tuple with the Filter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntryServerNodeListRequest) GetFilterOk() (*KalturaEntryServerNodeFilter, bool) {
	if o == nil || o.Filter == nil {
		return nil, false
	}
	return o.Filter, true
}

// HasFilter returns a boolean if a field has been set.
func (o *EntryServerNodeListRequest) HasFilter() bool {
	if o != nil && o.Filter != nil {
		return true
	}

	return false
}

// SetFilter gets a reference to the given KalturaEntryServerNodeFilter and assigns it to the Filter field.
func (o *EntryServerNodeListRequest) SetFilter(v KalturaEntryServerNodeFilter) {
	o.Filter = &v
}

// GetPager returns the Pager field value if set, zero value otherwise.
func (o *EntryServerNodeListRequest) GetPager() KalturaFilterPager {
	if o == nil || o.Pager == nil {
		var ret KalturaFilterPager
		return ret
	}
	return *o.Pager
}

// GetPagerOk returns a tuple with the Pager field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntryServerNodeListRequest) GetPagerOk() (*KalturaFilterPager, bool) {
	if o == nil || o.Pager == nil {
		return nil, false
	}
	return o.Pager, true
}

// HasPager returns a boolean if a field has been set.
func (o *EntryServerNodeListRequest) HasPager() bool {
	if o != nil && o.Pager != nil {
		return true
	}

	return false
}

// SetPager gets a reference to the given KalturaFilterPager and assigns it to the Pager field.
func (o *EntryServerNodeListRequest) SetPager(v KalturaFilterPager) {
	o.Pager = &v
}

func (o EntryServerNodeListRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Filter != nil {
		toSerialize["filter"] = o.Filter
	}
	if o.Pager != nil {
		toSerialize["pager"] = o.Pager
	}
	return json.Marshal(toSerialize)
}

type NullableEntryServerNodeListRequest struct {
	value *EntryServerNodeListRequest
	isSet bool
}

func (v NullableEntryServerNodeListRequest) Get() *EntryServerNodeListRequest {
	return v.value
}

func (v *NullableEntryServerNodeListRequest) Set(val *EntryServerNodeListRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableEntryServerNodeListRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableEntryServerNodeListRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEntryServerNodeListRequest(val *EntryServerNodeListRequest) *NullableEntryServerNodeListRequest {
	return &NullableEntryServerNodeListRequest{value: val, isSet: true}
}

func (v NullableEntryServerNodeListRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEntryServerNodeListRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


