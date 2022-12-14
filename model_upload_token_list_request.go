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

// UploadTokenListRequest struct for UploadTokenListRequest
type UploadTokenListRequest struct {
	Filter *KalturaUploadTokenFilter `json:"filter,omitempty"`
	Pager *KalturaFilterPager `json:"pager,omitempty"`
}

// NewUploadTokenListRequest instantiates a new UploadTokenListRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUploadTokenListRequest() *UploadTokenListRequest {
	this := UploadTokenListRequest{}
	return &this
}

// NewUploadTokenListRequestWithDefaults instantiates a new UploadTokenListRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUploadTokenListRequestWithDefaults() *UploadTokenListRequest {
	this := UploadTokenListRequest{}
	return &this
}

// GetFilter returns the Filter field value if set, zero value otherwise.
func (o *UploadTokenListRequest) GetFilter() KalturaUploadTokenFilter {
	if o == nil || o.Filter == nil {
		var ret KalturaUploadTokenFilter
		return ret
	}
	return *o.Filter
}

// GetFilterOk returns a tuple with the Filter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UploadTokenListRequest) GetFilterOk() (*KalturaUploadTokenFilter, bool) {
	if o == nil || o.Filter == nil {
		return nil, false
	}
	return o.Filter, true
}

// HasFilter returns a boolean if a field has been set.
func (o *UploadTokenListRequest) HasFilter() bool {
	if o != nil && o.Filter != nil {
		return true
	}

	return false
}

// SetFilter gets a reference to the given KalturaUploadTokenFilter and assigns it to the Filter field.
func (o *UploadTokenListRequest) SetFilter(v KalturaUploadTokenFilter) {
	o.Filter = &v
}

// GetPager returns the Pager field value if set, zero value otherwise.
func (o *UploadTokenListRequest) GetPager() KalturaFilterPager {
	if o == nil || o.Pager == nil {
		var ret KalturaFilterPager
		return ret
	}
	return *o.Pager
}

// GetPagerOk returns a tuple with the Pager field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UploadTokenListRequest) GetPagerOk() (*KalturaFilterPager, bool) {
	if o == nil || o.Pager == nil {
		return nil, false
	}
	return o.Pager, true
}

// HasPager returns a boolean if a field has been set.
func (o *UploadTokenListRequest) HasPager() bool {
	if o != nil && o.Pager != nil {
		return true
	}

	return false
}

// SetPager gets a reference to the given KalturaFilterPager and assigns it to the Pager field.
func (o *UploadTokenListRequest) SetPager(v KalturaFilterPager) {
	o.Pager = &v
}

func (o UploadTokenListRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Filter != nil {
		toSerialize["filter"] = o.Filter
	}
	if o.Pager != nil {
		toSerialize["pager"] = o.Pager
	}
	return json.Marshal(toSerialize)
}

type NullableUploadTokenListRequest struct {
	value *UploadTokenListRequest
	isSet bool
}

func (v NullableUploadTokenListRequest) Get() *UploadTokenListRequest {
	return v.value
}

func (v *NullableUploadTokenListRequest) Set(val *UploadTokenListRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUploadTokenListRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUploadTokenListRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUploadTokenListRequest(val *UploadTokenListRequest) *NullableUploadTokenListRequest {
	return &NullableUploadTokenListRequest{value: val, isSet: true}
}

func (v NullableUploadTokenListRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUploadTokenListRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


