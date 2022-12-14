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

// FlavorAssetGetUrlRequest struct for FlavorAssetGetUrlRequest
type FlavorAssetGetUrlRequest struct {
	ForceProxy *bool `json:"forceProxy,omitempty"`
	Id string `json:"id"`
	Options *KalturaFlavorAssetUrlOptions `json:"options,omitempty"`
	StorageId *int32 `json:"storageId,omitempty"`
}

// NewFlavorAssetGetUrlRequest instantiates a new FlavorAssetGetUrlRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFlavorAssetGetUrlRequest(id string) *FlavorAssetGetUrlRequest {
	this := FlavorAssetGetUrlRequest{}
	var forceProxy bool = false
	this.ForceProxy = &forceProxy
	this.Id = id
	return &this
}

// NewFlavorAssetGetUrlRequestWithDefaults instantiates a new FlavorAssetGetUrlRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFlavorAssetGetUrlRequestWithDefaults() *FlavorAssetGetUrlRequest {
	this := FlavorAssetGetUrlRequest{}
	var forceProxy bool = false
	this.ForceProxy = &forceProxy
	return &this
}

// GetForceProxy returns the ForceProxy field value if set, zero value otherwise.
func (o *FlavorAssetGetUrlRequest) GetForceProxy() bool {
	if o == nil || o.ForceProxy == nil {
		var ret bool
		return ret
	}
	return *o.ForceProxy
}

// GetForceProxyOk returns a tuple with the ForceProxy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlavorAssetGetUrlRequest) GetForceProxyOk() (*bool, bool) {
	if o == nil || o.ForceProxy == nil {
		return nil, false
	}
	return o.ForceProxy, true
}

// HasForceProxy returns a boolean if a field has been set.
func (o *FlavorAssetGetUrlRequest) HasForceProxy() bool {
	if o != nil && o.ForceProxy != nil {
		return true
	}

	return false
}

// SetForceProxy gets a reference to the given bool and assigns it to the ForceProxy field.
func (o *FlavorAssetGetUrlRequest) SetForceProxy(v bool) {
	o.ForceProxy = &v
}

// GetId returns the Id field value
func (o *FlavorAssetGetUrlRequest) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *FlavorAssetGetUrlRequest) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *FlavorAssetGetUrlRequest) SetId(v string) {
	o.Id = v
}

// GetOptions returns the Options field value if set, zero value otherwise.
func (o *FlavorAssetGetUrlRequest) GetOptions() KalturaFlavorAssetUrlOptions {
	if o == nil || o.Options == nil {
		var ret KalturaFlavorAssetUrlOptions
		return ret
	}
	return *o.Options
}

// GetOptionsOk returns a tuple with the Options field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlavorAssetGetUrlRequest) GetOptionsOk() (*KalturaFlavorAssetUrlOptions, bool) {
	if o == nil || o.Options == nil {
		return nil, false
	}
	return o.Options, true
}

// HasOptions returns a boolean if a field has been set.
func (o *FlavorAssetGetUrlRequest) HasOptions() bool {
	if o != nil && o.Options != nil {
		return true
	}

	return false
}

// SetOptions gets a reference to the given KalturaFlavorAssetUrlOptions and assigns it to the Options field.
func (o *FlavorAssetGetUrlRequest) SetOptions(v KalturaFlavorAssetUrlOptions) {
	o.Options = &v
}

// GetStorageId returns the StorageId field value if set, zero value otherwise.
func (o *FlavorAssetGetUrlRequest) GetStorageId() int32 {
	if o == nil || o.StorageId == nil {
		var ret int32
		return ret
	}
	return *o.StorageId
}

// GetStorageIdOk returns a tuple with the StorageId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlavorAssetGetUrlRequest) GetStorageIdOk() (*int32, bool) {
	if o == nil || o.StorageId == nil {
		return nil, false
	}
	return o.StorageId, true
}

// HasStorageId returns a boolean if a field has been set.
func (o *FlavorAssetGetUrlRequest) HasStorageId() bool {
	if o != nil && o.StorageId != nil {
		return true
	}

	return false
}

// SetStorageId gets a reference to the given int32 and assigns it to the StorageId field.
func (o *FlavorAssetGetUrlRequest) SetStorageId(v int32) {
	o.StorageId = &v
}

func (o FlavorAssetGetUrlRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ForceProxy != nil {
		toSerialize["forceProxy"] = o.ForceProxy
	}
	if true {
		toSerialize["id"] = o.Id
	}
	if o.Options != nil {
		toSerialize["options"] = o.Options
	}
	if o.StorageId != nil {
		toSerialize["storageId"] = o.StorageId
	}
	return json.Marshal(toSerialize)
}

type NullableFlavorAssetGetUrlRequest struct {
	value *FlavorAssetGetUrlRequest
	isSet bool
}

func (v NullableFlavorAssetGetUrlRequest) Get() *FlavorAssetGetUrlRequest {
	return v.value
}

func (v *NullableFlavorAssetGetUrlRequest) Set(val *FlavorAssetGetUrlRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableFlavorAssetGetUrlRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableFlavorAssetGetUrlRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFlavorAssetGetUrlRequest(val *FlavorAssetGetUrlRequest) *NullableFlavorAssetGetUrlRequest {
	return &NullableFlavorAssetGetUrlRequest{value: val, isSet: true}
}

func (v NullableFlavorAssetGetUrlRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFlavorAssetGetUrlRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


