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

// KalturaPlaybackSource struct for KalturaPlaybackSource
type KalturaPlaybackSource struct {
	DeliveryProfileId *string `json:"deliveryProfileId,omitempty"`
	Drm []KalturaDrmPlaybackPluginData `json:"drm,omitempty"`
	// comma separated string of flavor ids
	FlavorIds *string `json:"flavorIds,omitempty"`
	// source format according to delivery profile streamer type (applehttp, mpegdash etc.)
	Format *string `json:"format,omitempty"`
	// comma separated string according to deliveryProfile media protocols ('http,https' etc.)
	Protocols *string `json:"protocols,omitempty"`
	Url *string `json:"url,omitempty"`
}

// NewKalturaPlaybackSource instantiates a new KalturaPlaybackSource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaPlaybackSource() *KalturaPlaybackSource {
	this := KalturaPlaybackSource{}
	return &this
}

// NewKalturaPlaybackSourceWithDefaults instantiates a new KalturaPlaybackSource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaPlaybackSourceWithDefaults() *KalturaPlaybackSource {
	this := KalturaPlaybackSource{}
	return &this
}

// GetDeliveryProfileId returns the DeliveryProfileId field value if set, zero value otherwise.
func (o *KalturaPlaybackSource) GetDeliveryProfileId() string {
	if o == nil || o.DeliveryProfileId == nil {
		var ret string
		return ret
	}
	return *o.DeliveryProfileId
}

// GetDeliveryProfileIdOk returns a tuple with the DeliveryProfileId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaPlaybackSource) GetDeliveryProfileIdOk() (*string, bool) {
	if o == nil || o.DeliveryProfileId == nil {
		return nil, false
	}
	return o.DeliveryProfileId, true
}

// HasDeliveryProfileId returns a boolean if a field has been set.
func (o *KalturaPlaybackSource) HasDeliveryProfileId() bool {
	if o != nil && o.DeliveryProfileId != nil {
		return true
	}

	return false
}

// SetDeliveryProfileId gets a reference to the given string and assigns it to the DeliveryProfileId field.
func (o *KalturaPlaybackSource) SetDeliveryProfileId(v string) {
	o.DeliveryProfileId = &v
}

// GetDrm returns the Drm field value if set, zero value otherwise.
func (o *KalturaPlaybackSource) GetDrm() []KalturaDrmPlaybackPluginData {
	if o == nil || o.Drm == nil {
		var ret []KalturaDrmPlaybackPluginData
		return ret
	}
	return o.Drm
}

// GetDrmOk returns a tuple with the Drm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaPlaybackSource) GetDrmOk() ([]KalturaDrmPlaybackPluginData, bool) {
	if o == nil || o.Drm == nil {
		return nil, false
	}
	return o.Drm, true
}

// HasDrm returns a boolean if a field has been set.
func (o *KalturaPlaybackSource) HasDrm() bool {
	if o != nil && o.Drm != nil {
		return true
	}

	return false
}

// SetDrm gets a reference to the given []KalturaDrmPlaybackPluginData and assigns it to the Drm field.
func (o *KalturaPlaybackSource) SetDrm(v []KalturaDrmPlaybackPluginData) {
	o.Drm = v
}

// GetFlavorIds returns the FlavorIds field value if set, zero value otherwise.
func (o *KalturaPlaybackSource) GetFlavorIds() string {
	if o == nil || o.FlavorIds == nil {
		var ret string
		return ret
	}
	return *o.FlavorIds
}

// GetFlavorIdsOk returns a tuple with the FlavorIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaPlaybackSource) GetFlavorIdsOk() (*string, bool) {
	if o == nil || o.FlavorIds == nil {
		return nil, false
	}
	return o.FlavorIds, true
}

// HasFlavorIds returns a boolean if a field has been set.
func (o *KalturaPlaybackSource) HasFlavorIds() bool {
	if o != nil && o.FlavorIds != nil {
		return true
	}

	return false
}

// SetFlavorIds gets a reference to the given string and assigns it to the FlavorIds field.
func (o *KalturaPlaybackSource) SetFlavorIds(v string) {
	o.FlavorIds = &v
}

// GetFormat returns the Format field value if set, zero value otherwise.
func (o *KalturaPlaybackSource) GetFormat() string {
	if o == nil || o.Format == nil {
		var ret string
		return ret
	}
	return *o.Format
}

// GetFormatOk returns a tuple with the Format field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaPlaybackSource) GetFormatOk() (*string, bool) {
	if o == nil || o.Format == nil {
		return nil, false
	}
	return o.Format, true
}

// HasFormat returns a boolean if a field has been set.
func (o *KalturaPlaybackSource) HasFormat() bool {
	if o != nil && o.Format != nil {
		return true
	}

	return false
}

// SetFormat gets a reference to the given string and assigns it to the Format field.
func (o *KalturaPlaybackSource) SetFormat(v string) {
	o.Format = &v
}

// GetProtocols returns the Protocols field value if set, zero value otherwise.
func (o *KalturaPlaybackSource) GetProtocols() string {
	if o == nil || o.Protocols == nil {
		var ret string
		return ret
	}
	return *o.Protocols
}

// GetProtocolsOk returns a tuple with the Protocols field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaPlaybackSource) GetProtocolsOk() (*string, bool) {
	if o == nil || o.Protocols == nil {
		return nil, false
	}
	return o.Protocols, true
}

// HasProtocols returns a boolean if a field has been set.
func (o *KalturaPlaybackSource) HasProtocols() bool {
	if o != nil && o.Protocols != nil {
		return true
	}

	return false
}

// SetProtocols gets a reference to the given string and assigns it to the Protocols field.
func (o *KalturaPlaybackSource) SetProtocols(v string) {
	o.Protocols = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *KalturaPlaybackSource) GetUrl() string {
	if o == nil || o.Url == nil {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaPlaybackSource) GetUrlOk() (*string, bool) {
	if o == nil || o.Url == nil {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *KalturaPlaybackSource) HasUrl() bool {
	if o != nil && o.Url != nil {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *KalturaPlaybackSource) SetUrl(v string) {
	o.Url = &v
}

func (o KalturaPlaybackSource) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DeliveryProfileId != nil {
		toSerialize["deliveryProfileId"] = o.DeliveryProfileId
	}
	if o.Drm != nil {
		toSerialize["drm"] = o.Drm
	}
	if o.FlavorIds != nil {
		toSerialize["flavorIds"] = o.FlavorIds
	}
	if o.Format != nil {
		toSerialize["format"] = o.Format
	}
	if o.Protocols != nil {
		toSerialize["protocols"] = o.Protocols
	}
	if o.Url != nil {
		toSerialize["url"] = o.Url
	}
	return json.Marshal(toSerialize)
}

type NullableKalturaPlaybackSource struct {
	value *KalturaPlaybackSource
	isSet bool
}

func (v NullableKalturaPlaybackSource) Get() *KalturaPlaybackSource {
	return v.value
}

func (v *NullableKalturaPlaybackSource) Set(val *KalturaPlaybackSource) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaPlaybackSource) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaPlaybackSource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaPlaybackSource(val *KalturaPlaybackSource) *NullableKalturaPlaybackSource {
	return &NullableKalturaPlaybackSource{value: val, isSet: true}
}

func (v NullableKalturaPlaybackSource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaPlaybackSource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


