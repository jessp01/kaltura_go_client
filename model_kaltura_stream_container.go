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

// KalturaStreamContainer struct for KalturaStreamContainer
type KalturaStreamContainer struct {
	ChannelIndex *int32 `json:"channelIndex,omitempty"`
	ChannelLayout *string `json:"channelLayout,omitempty"`
	Id *string `json:"id,omitempty"`
	Label *string `json:"label,omitempty"`
	Language *string `json:"language,omitempty"`
	TrackIndex *int32 `json:"trackIndex,omitempty"`
	Type *string `json:"type,omitempty"`
}

// NewKalturaStreamContainer instantiates a new KalturaStreamContainer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaStreamContainer() *KalturaStreamContainer {
	this := KalturaStreamContainer{}
	return &this
}

// NewKalturaStreamContainerWithDefaults instantiates a new KalturaStreamContainer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaStreamContainerWithDefaults() *KalturaStreamContainer {
	this := KalturaStreamContainer{}
	return &this
}

// GetChannelIndex returns the ChannelIndex field value if set, zero value otherwise.
func (o *KalturaStreamContainer) GetChannelIndex() int32 {
	if o == nil || o.ChannelIndex == nil {
		var ret int32
		return ret
	}
	return *o.ChannelIndex
}

// GetChannelIndexOk returns a tuple with the ChannelIndex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaStreamContainer) GetChannelIndexOk() (*int32, bool) {
	if o == nil || o.ChannelIndex == nil {
		return nil, false
	}
	return o.ChannelIndex, true
}

// HasChannelIndex returns a boolean if a field has been set.
func (o *KalturaStreamContainer) HasChannelIndex() bool {
	if o != nil && o.ChannelIndex != nil {
		return true
	}

	return false
}

// SetChannelIndex gets a reference to the given int32 and assigns it to the ChannelIndex field.
func (o *KalturaStreamContainer) SetChannelIndex(v int32) {
	o.ChannelIndex = &v
}

// GetChannelLayout returns the ChannelLayout field value if set, zero value otherwise.
func (o *KalturaStreamContainer) GetChannelLayout() string {
	if o == nil || o.ChannelLayout == nil {
		var ret string
		return ret
	}
	return *o.ChannelLayout
}

// GetChannelLayoutOk returns a tuple with the ChannelLayout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaStreamContainer) GetChannelLayoutOk() (*string, bool) {
	if o == nil || o.ChannelLayout == nil {
		return nil, false
	}
	return o.ChannelLayout, true
}

// HasChannelLayout returns a boolean if a field has been set.
func (o *KalturaStreamContainer) HasChannelLayout() bool {
	if o != nil && o.ChannelLayout != nil {
		return true
	}

	return false
}

// SetChannelLayout gets a reference to the given string and assigns it to the ChannelLayout field.
func (o *KalturaStreamContainer) SetChannelLayout(v string) {
	o.ChannelLayout = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *KalturaStreamContainer) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaStreamContainer) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *KalturaStreamContainer) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *KalturaStreamContainer) SetId(v string) {
	o.Id = &v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *KalturaStreamContainer) GetLabel() string {
	if o == nil || o.Label == nil {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaStreamContainer) GetLabelOk() (*string, bool) {
	if o == nil || o.Label == nil {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *KalturaStreamContainer) HasLabel() bool {
	if o != nil && o.Label != nil {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *KalturaStreamContainer) SetLabel(v string) {
	o.Label = &v
}

// GetLanguage returns the Language field value if set, zero value otherwise.
func (o *KalturaStreamContainer) GetLanguage() string {
	if o == nil || o.Language == nil {
		var ret string
		return ret
	}
	return *o.Language
}

// GetLanguageOk returns a tuple with the Language field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaStreamContainer) GetLanguageOk() (*string, bool) {
	if o == nil || o.Language == nil {
		return nil, false
	}
	return o.Language, true
}

// HasLanguage returns a boolean if a field has been set.
func (o *KalturaStreamContainer) HasLanguage() bool {
	if o != nil && o.Language != nil {
		return true
	}

	return false
}

// SetLanguage gets a reference to the given string and assigns it to the Language field.
func (o *KalturaStreamContainer) SetLanguage(v string) {
	o.Language = &v
}

// GetTrackIndex returns the TrackIndex field value if set, zero value otherwise.
func (o *KalturaStreamContainer) GetTrackIndex() int32 {
	if o == nil || o.TrackIndex == nil {
		var ret int32
		return ret
	}
	return *o.TrackIndex
}

// GetTrackIndexOk returns a tuple with the TrackIndex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaStreamContainer) GetTrackIndexOk() (*int32, bool) {
	if o == nil || o.TrackIndex == nil {
		return nil, false
	}
	return o.TrackIndex, true
}

// HasTrackIndex returns a boolean if a field has been set.
func (o *KalturaStreamContainer) HasTrackIndex() bool {
	if o != nil && o.TrackIndex != nil {
		return true
	}

	return false
}

// SetTrackIndex gets a reference to the given int32 and assigns it to the TrackIndex field.
func (o *KalturaStreamContainer) SetTrackIndex(v int32) {
	o.TrackIndex = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *KalturaStreamContainer) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaStreamContainer) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *KalturaStreamContainer) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *KalturaStreamContainer) SetType(v string) {
	o.Type = &v
}

func (o KalturaStreamContainer) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ChannelIndex != nil {
		toSerialize["channelIndex"] = o.ChannelIndex
	}
	if o.ChannelLayout != nil {
		toSerialize["channelLayout"] = o.ChannelLayout
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Label != nil {
		toSerialize["label"] = o.Label
	}
	if o.Language != nil {
		toSerialize["language"] = o.Language
	}
	if o.TrackIndex != nil {
		toSerialize["trackIndex"] = o.TrackIndex
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableKalturaStreamContainer struct {
	value *KalturaStreamContainer
	isSet bool
}

func (v NullableKalturaStreamContainer) Get() *KalturaStreamContainer {
	return v.value
}

func (v *NullableKalturaStreamContainer) Set(val *KalturaStreamContainer) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaStreamContainer) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaStreamContainer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaStreamContainer(val *KalturaStreamContainer) *NullableKalturaStreamContainer {
	return &NullableKalturaStreamContainer{value: val, isSet: true}
}

func (v NullableKalturaStreamContainer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaStreamContainer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


