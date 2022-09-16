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

// MediaInfoListRequest struct for MediaInfoListRequest
type MediaInfoListRequest struct {
	Filter *KalturaMediaInfoFilter `json:"filter,omitempty"`
	Pager *KalturaFilterPager `json:"pager,omitempty"`
}

// NewMediaInfoListRequest instantiates a new MediaInfoListRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMediaInfoListRequest() *MediaInfoListRequest {
	this := MediaInfoListRequest{}
	return &this
}

// NewMediaInfoListRequestWithDefaults instantiates a new MediaInfoListRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMediaInfoListRequestWithDefaults() *MediaInfoListRequest {
	this := MediaInfoListRequest{}
	return &this
}

// GetFilter returns the Filter field value if set, zero value otherwise.
func (o *MediaInfoListRequest) GetFilter() KalturaMediaInfoFilter {
	if o == nil || o.Filter == nil {
		var ret KalturaMediaInfoFilter
		return ret
	}
	return *o.Filter
}

// GetFilterOk returns a tuple with the Filter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MediaInfoListRequest) GetFilterOk() (*KalturaMediaInfoFilter, bool) {
	if o == nil || o.Filter == nil {
		return nil, false
	}
	return o.Filter, true
}

// HasFilter returns a boolean if a field has been set.
func (o *MediaInfoListRequest) HasFilter() bool {
	if o != nil && o.Filter != nil {
		return true
	}

	return false
}

// SetFilter gets a reference to the given KalturaMediaInfoFilter and assigns it to the Filter field.
func (o *MediaInfoListRequest) SetFilter(v KalturaMediaInfoFilter) {
	o.Filter = &v
}

// GetPager returns the Pager field value if set, zero value otherwise.
func (o *MediaInfoListRequest) GetPager() KalturaFilterPager {
	if o == nil || o.Pager == nil {
		var ret KalturaFilterPager
		return ret
	}
	return *o.Pager
}

// GetPagerOk returns a tuple with the Pager field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MediaInfoListRequest) GetPagerOk() (*KalturaFilterPager, bool) {
	if o == nil || o.Pager == nil {
		return nil, false
	}
	return o.Pager, true
}

// HasPager returns a boolean if a field has been set.
func (o *MediaInfoListRequest) HasPager() bool {
	if o != nil && o.Pager != nil {
		return true
	}

	return false
}

// SetPager gets a reference to the given KalturaFilterPager and assigns it to the Pager field.
func (o *MediaInfoListRequest) SetPager(v KalturaFilterPager) {
	o.Pager = &v
}

func (o MediaInfoListRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Filter != nil {
		toSerialize["filter"] = o.Filter
	}
	if o.Pager != nil {
		toSerialize["pager"] = o.Pager
	}
	return json.Marshal(toSerialize)
}

type NullableMediaInfoListRequest struct {
	value *MediaInfoListRequest
	isSet bool
}

func (v NullableMediaInfoListRequest) Get() *MediaInfoListRequest {
	return v.value
}

func (v *NullableMediaInfoListRequest) Set(val *MediaInfoListRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableMediaInfoListRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableMediaInfoListRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMediaInfoListRequest(val *MediaInfoListRequest) *NullableMediaInfoListRequest {
	return &NullableMediaInfoListRequest{value: val, isSet: true}
}

func (v NullableMediaInfoListRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMediaInfoListRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

