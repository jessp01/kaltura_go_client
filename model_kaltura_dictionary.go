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

// KalturaDictionary struct for KalturaDictionary
type KalturaDictionary struct {
	Data *string `json:"data,omitempty"`
	// Enum Type: `KalturaCatalogItemLanguage`
	Language *string `json:"language,omitempty"`
}

// NewKalturaDictionary instantiates a new KalturaDictionary object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaDictionary() *KalturaDictionary {
	this := KalturaDictionary{}
	return &this
}

// NewKalturaDictionaryWithDefaults instantiates a new KalturaDictionary object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaDictionaryWithDefaults() *KalturaDictionary {
	this := KalturaDictionary{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *KalturaDictionary) GetData() string {
	if o == nil || o.Data == nil {
		var ret string
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaDictionary) GetDataOk() (*string, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *KalturaDictionary) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given string and assigns it to the Data field.
func (o *KalturaDictionary) SetData(v string) {
	o.Data = &v
}

// GetLanguage returns the Language field value if set, zero value otherwise.
func (o *KalturaDictionary) GetLanguage() string {
	if o == nil || o.Language == nil {
		var ret string
		return ret
	}
	return *o.Language
}

// GetLanguageOk returns a tuple with the Language field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaDictionary) GetLanguageOk() (*string, bool) {
	if o == nil || o.Language == nil {
		return nil, false
	}
	return o.Language, true
}

// HasLanguage returns a boolean if a field has been set.
func (o *KalturaDictionary) HasLanguage() bool {
	if o != nil && o.Language != nil {
		return true
	}

	return false
}

// SetLanguage gets a reference to the given string and assigns it to the Language field.
func (o *KalturaDictionary) SetLanguage(v string) {
	o.Language = &v
}

func (o KalturaDictionary) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	if o.Language != nil {
		toSerialize["language"] = o.Language
	}
	return json.Marshal(toSerialize)
}

type NullableKalturaDictionary struct {
	value *KalturaDictionary
	isSet bool
}

func (v NullableKalturaDictionary) Get() *KalturaDictionary {
	return v.value
}

func (v *NullableKalturaDictionary) Set(val *KalturaDictionary) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaDictionary) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaDictionary) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaDictionary(val *KalturaDictionary) *NullableKalturaDictionary {
	return &NullableKalturaDictionary{value: val, isSet: true}
}

func (v NullableKalturaDictionary) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaDictionary) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

