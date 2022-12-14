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

// KalturaSearch struct for KalturaSearch
type KalturaSearch struct {
	AuthData *string `json:"authData,omitempty"`
	// Use this field to pass dynamic data for searching  For example - if you set this field to \"mymovies_$partner_id\"  The $partner_id will be automatically replcaed with your real partner Id
	ExtraData *string `json:"extraData,omitempty"`
	KeyWords *string `json:"keyWords,omitempty"`
	// Enum Type: `KalturaMediaType`
	MediaType *int32 `json:"mediaType,omitempty"`
	ObjectType *string `json:"objectType,omitempty"`
	// Enum Type: `KalturaSearchProviderType`
	SearchSource *int32 `json:"searchSource,omitempty"`
}

// NewKalturaSearch instantiates a new KalturaSearch object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaSearch() *KalturaSearch {
	this := KalturaSearch{}
	return &this
}

// NewKalturaSearchWithDefaults instantiates a new KalturaSearch object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaSearchWithDefaults() *KalturaSearch {
	this := KalturaSearch{}
	return &this
}

// GetAuthData returns the AuthData field value if set, zero value otherwise.
func (o *KalturaSearch) GetAuthData() string {
	if o == nil || o.AuthData == nil {
		var ret string
		return ret
	}
	return *o.AuthData
}

// GetAuthDataOk returns a tuple with the AuthData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaSearch) GetAuthDataOk() (*string, bool) {
	if o == nil || o.AuthData == nil {
		return nil, false
	}
	return o.AuthData, true
}

// HasAuthData returns a boolean if a field has been set.
func (o *KalturaSearch) HasAuthData() bool {
	if o != nil && o.AuthData != nil {
		return true
	}

	return false
}

// SetAuthData gets a reference to the given string and assigns it to the AuthData field.
func (o *KalturaSearch) SetAuthData(v string) {
	o.AuthData = &v
}

// GetExtraData returns the ExtraData field value if set, zero value otherwise.
func (o *KalturaSearch) GetExtraData() string {
	if o == nil || o.ExtraData == nil {
		var ret string
		return ret
	}
	return *o.ExtraData
}

// GetExtraDataOk returns a tuple with the ExtraData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaSearch) GetExtraDataOk() (*string, bool) {
	if o == nil || o.ExtraData == nil {
		return nil, false
	}
	return o.ExtraData, true
}

// HasExtraData returns a boolean if a field has been set.
func (o *KalturaSearch) HasExtraData() bool {
	if o != nil && o.ExtraData != nil {
		return true
	}

	return false
}

// SetExtraData gets a reference to the given string and assigns it to the ExtraData field.
func (o *KalturaSearch) SetExtraData(v string) {
	o.ExtraData = &v
}

// GetKeyWords returns the KeyWords field value if set, zero value otherwise.
func (o *KalturaSearch) GetKeyWords() string {
	if o == nil || o.KeyWords == nil {
		var ret string
		return ret
	}
	return *o.KeyWords
}

// GetKeyWordsOk returns a tuple with the KeyWords field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaSearch) GetKeyWordsOk() (*string, bool) {
	if o == nil || o.KeyWords == nil {
		return nil, false
	}
	return o.KeyWords, true
}

// HasKeyWords returns a boolean if a field has been set.
func (o *KalturaSearch) HasKeyWords() bool {
	if o != nil && o.KeyWords != nil {
		return true
	}

	return false
}

// SetKeyWords gets a reference to the given string and assigns it to the KeyWords field.
func (o *KalturaSearch) SetKeyWords(v string) {
	o.KeyWords = &v
}

// GetMediaType returns the MediaType field value if set, zero value otherwise.
func (o *KalturaSearch) GetMediaType() int32 {
	if o == nil || o.MediaType == nil {
		var ret int32
		return ret
	}
	return *o.MediaType
}

// GetMediaTypeOk returns a tuple with the MediaType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaSearch) GetMediaTypeOk() (*int32, bool) {
	if o == nil || o.MediaType == nil {
		return nil, false
	}
	return o.MediaType, true
}

// HasMediaType returns a boolean if a field has been set.
func (o *KalturaSearch) HasMediaType() bool {
	if o != nil && o.MediaType != nil {
		return true
	}

	return false
}

// SetMediaType gets a reference to the given int32 and assigns it to the MediaType field.
func (o *KalturaSearch) SetMediaType(v int32) {
	o.MediaType = &v
}

// GetObjectType returns the ObjectType field value if set, zero value otherwise.
func (o *KalturaSearch) GetObjectType() string {
	if o == nil || o.ObjectType == nil {
		var ret string
		return ret
	}
	return *o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaSearch) GetObjectTypeOk() (*string, bool) {
	if o == nil || o.ObjectType == nil {
		return nil, false
	}
	return o.ObjectType, true
}

// HasObjectType returns a boolean if a field has been set.
func (o *KalturaSearch) HasObjectType() bool {
	if o != nil && o.ObjectType != nil {
		return true
	}

	return false
}

// SetObjectType gets a reference to the given string and assigns it to the ObjectType field.
func (o *KalturaSearch) SetObjectType(v string) {
	o.ObjectType = &v
}

// GetSearchSource returns the SearchSource field value if set, zero value otherwise.
func (o *KalturaSearch) GetSearchSource() int32 {
	if o == nil || o.SearchSource == nil {
		var ret int32
		return ret
	}
	return *o.SearchSource
}

// GetSearchSourceOk returns a tuple with the SearchSource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaSearch) GetSearchSourceOk() (*int32, bool) {
	if o == nil || o.SearchSource == nil {
		return nil, false
	}
	return o.SearchSource, true
}

// HasSearchSource returns a boolean if a field has been set.
func (o *KalturaSearch) HasSearchSource() bool {
	if o != nil && o.SearchSource != nil {
		return true
	}

	return false
}

// SetSearchSource gets a reference to the given int32 and assigns it to the SearchSource field.
func (o *KalturaSearch) SetSearchSource(v int32) {
	o.SearchSource = &v
}

func (o KalturaSearch) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AuthData != nil {
		toSerialize["authData"] = o.AuthData
	}
	if o.ExtraData != nil {
		toSerialize["extraData"] = o.ExtraData
	}
	if o.KeyWords != nil {
		toSerialize["keyWords"] = o.KeyWords
	}
	if o.MediaType != nil {
		toSerialize["mediaType"] = o.MediaType
	}
	if o.ObjectType != nil {
		toSerialize["objectType"] = o.ObjectType
	}
	if o.SearchSource != nil {
		toSerialize["searchSource"] = o.SearchSource
	}
	return json.Marshal(toSerialize)
}

type NullableKalturaSearch struct {
	value *KalturaSearch
	isSet bool
}

func (v NullableKalturaSearch) Get() *KalturaSearch {
	return v.value
}

func (v *NullableKalturaSearch) Set(val *KalturaSearch) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaSearch) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaSearch) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaSearch(val *KalturaSearch) *NullableKalturaSearch {
	return &NullableKalturaSearch{value: val, isSet: true}
}

func (v NullableKalturaSearch) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaSearch) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


