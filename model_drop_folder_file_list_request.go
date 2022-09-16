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

// DropFolderFileListRequest struct for DropFolderFileListRequest
type DropFolderFileListRequest struct {
	Filter *KalturaDropFolderFileFilter `json:"filter,omitempty"`
	Pager *KalturaFilterPager `json:"pager,omitempty"`
}

// NewDropFolderFileListRequest instantiates a new DropFolderFileListRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDropFolderFileListRequest() *DropFolderFileListRequest {
	this := DropFolderFileListRequest{}
	return &this
}

// NewDropFolderFileListRequestWithDefaults instantiates a new DropFolderFileListRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDropFolderFileListRequestWithDefaults() *DropFolderFileListRequest {
	this := DropFolderFileListRequest{}
	return &this
}

// GetFilter returns the Filter field value if set, zero value otherwise.
func (o *DropFolderFileListRequest) GetFilter() KalturaDropFolderFileFilter {
	if o == nil || o.Filter == nil {
		var ret KalturaDropFolderFileFilter
		return ret
	}
	return *o.Filter
}

// GetFilterOk returns a tuple with the Filter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DropFolderFileListRequest) GetFilterOk() (*KalturaDropFolderFileFilter, bool) {
	if o == nil || o.Filter == nil {
		return nil, false
	}
	return o.Filter, true
}

// HasFilter returns a boolean if a field has been set.
func (o *DropFolderFileListRequest) HasFilter() bool {
	if o != nil && o.Filter != nil {
		return true
	}

	return false
}

// SetFilter gets a reference to the given KalturaDropFolderFileFilter and assigns it to the Filter field.
func (o *DropFolderFileListRequest) SetFilter(v KalturaDropFolderFileFilter) {
	o.Filter = &v
}

// GetPager returns the Pager field value if set, zero value otherwise.
func (o *DropFolderFileListRequest) GetPager() KalturaFilterPager {
	if o == nil || o.Pager == nil {
		var ret KalturaFilterPager
		return ret
	}
	return *o.Pager
}

// GetPagerOk returns a tuple with the Pager field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DropFolderFileListRequest) GetPagerOk() (*KalturaFilterPager, bool) {
	if o == nil || o.Pager == nil {
		return nil, false
	}
	return o.Pager, true
}

// HasPager returns a boolean if a field has been set.
func (o *DropFolderFileListRequest) HasPager() bool {
	if o != nil && o.Pager != nil {
		return true
	}

	return false
}

// SetPager gets a reference to the given KalturaFilterPager and assigns it to the Pager field.
func (o *DropFolderFileListRequest) SetPager(v KalturaFilterPager) {
	o.Pager = &v
}

func (o DropFolderFileListRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Filter != nil {
		toSerialize["filter"] = o.Filter
	}
	if o.Pager != nil {
		toSerialize["pager"] = o.Pager
	}
	return json.Marshal(toSerialize)
}

type NullableDropFolderFileListRequest struct {
	value *DropFolderFileListRequest
	isSet bool
}

func (v NullableDropFolderFileListRequest) Get() *DropFolderFileListRequest {
	return v.value
}

func (v *NullableDropFolderFileListRequest) Set(val *DropFolderFileListRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableDropFolderFileListRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableDropFolderFileListRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDropFolderFileListRequest(val *DropFolderFileListRequest) *NullableDropFolderFileListRequest {
	return &NullableDropFolderFileListRequest{value: val, isSet: true}
}

func (v NullableDropFolderFileListRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDropFolderFileListRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

