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

// CategoryListRequest struct for CategoryListRequest
type CategoryListRequest struct {
	Filter *KalturaCategoryFilter `json:"filter,omitempty"`
	Pager *KalturaFilterPager `json:"pager,omitempty"`
}

// NewCategoryListRequest instantiates a new CategoryListRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCategoryListRequest() *CategoryListRequest {
	this := CategoryListRequest{}
	return &this
}

// NewCategoryListRequestWithDefaults instantiates a new CategoryListRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCategoryListRequestWithDefaults() *CategoryListRequest {
	this := CategoryListRequest{}
	return &this
}

// GetFilter returns the Filter field value if set, zero value otherwise.
func (o *CategoryListRequest) GetFilter() KalturaCategoryFilter {
	if o == nil || o.Filter == nil {
		var ret KalturaCategoryFilter
		return ret
	}
	return *o.Filter
}

// GetFilterOk returns a tuple with the Filter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CategoryListRequest) GetFilterOk() (*KalturaCategoryFilter, bool) {
	if o == nil || o.Filter == nil {
		return nil, false
	}
	return o.Filter, true
}

// HasFilter returns a boolean if a field has been set.
func (o *CategoryListRequest) HasFilter() bool {
	if o != nil && o.Filter != nil {
		return true
	}

	return false
}

// SetFilter gets a reference to the given KalturaCategoryFilter and assigns it to the Filter field.
func (o *CategoryListRequest) SetFilter(v KalturaCategoryFilter) {
	o.Filter = &v
}

// GetPager returns the Pager field value if set, zero value otherwise.
func (o *CategoryListRequest) GetPager() KalturaFilterPager {
	if o == nil || o.Pager == nil {
		var ret KalturaFilterPager
		return ret
	}
	return *o.Pager
}

// GetPagerOk returns a tuple with the Pager field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CategoryListRequest) GetPagerOk() (*KalturaFilterPager, bool) {
	if o == nil || o.Pager == nil {
		return nil, false
	}
	return o.Pager, true
}

// HasPager returns a boolean if a field has been set.
func (o *CategoryListRequest) HasPager() bool {
	if o != nil && o.Pager != nil {
		return true
	}

	return false
}

// SetPager gets a reference to the given KalturaFilterPager and assigns it to the Pager field.
func (o *CategoryListRequest) SetPager(v KalturaFilterPager) {
	o.Pager = &v
}

func (o CategoryListRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Filter != nil {
		toSerialize["filter"] = o.Filter
	}
	if o.Pager != nil {
		toSerialize["pager"] = o.Pager
	}
	return json.Marshal(toSerialize)
}

type NullableCategoryListRequest struct {
	value *CategoryListRequest
	isSet bool
}

func (v NullableCategoryListRequest) Get() *CategoryListRequest {
	return v.value
}

func (v *NullableCategoryListRequest) Set(val *CategoryListRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCategoryListRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCategoryListRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCategoryListRequest(val *CategoryListRequest) *NullableCategoryListRequest {
	return &NullableCategoryListRequest{value: val, isSet: true}
}

func (v NullableCategoryListRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCategoryListRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


