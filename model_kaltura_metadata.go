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

// KalturaMetadata struct for KalturaMetadata
type KalturaMetadata struct {
	// `readOnly`
	CreatedAt *int32 `json:"createdAt,omitempty"`
	// `readOnly`
	Id *int32 `json:"id,omitempty"`
	// `readOnly`  Enum Type: `KalturaMetadataObjectType`
	MetadataObjectType *string `json:"metadataObjectType,omitempty"`
	// `readOnly`
	MetadataProfileId *int32 `json:"metadataProfileId,omitempty"`
	// `readOnly`
	MetadataProfileVersion *int32 `json:"metadataProfileVersion,omitempty"`
	// `readOnly`
	ObjectId *string `json:"objectId,omitempty"`
	// `readOnly`
	PartnerId *int32 `json:"partnerId,omitempty"`
	// `readOnly`  Enum Type: `KalturaMetadataStatus`
	Status *int32 `json:"status,omitempty"`
	// `readOnly`
	UpdatedAt *int32 `json:"updatedAt,omitempty"`
	// `readOnly`
	Version *int32 `json:"version,omitempty"`
	// `readOnly`
	Xml *string `json:"xml,omitempty"`
}

// NewKalturaMetadata instantiates a new KalturaMetadata object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaMetadata() *KalturaMetadata {
	this := KalturaMetadata{}
	return &this
}

// NewKalturaMetadataWithDefaults instantiates a new KalturaMetadata object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaMetadataWithDefaults() *KalturaMetadata {
	this := KalturaMetadata{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *KalturaMetadata) GetCreatedAt() int32 {
	if o == nil || o.CreatedAt == nil {
		var ret int32
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaMetadata) GetCreatedAtOk() (*int32, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *KalturaMetadata) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given int32 and assigns it to the CreatedAt field.
func (o *KalturaMetadata) SetCreatedAt(v int32) {
	o.CreatedAt = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *KalturaMetadata) GetId() int32 {
	if o == nil || o.Id == nil {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaMetadata) GetIdOk() (*int32, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *KalturaMetadata) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *KalturaMetadata) SetId(v int32) {
	o.Id = &v
}

// GetMetadataObjectType returns the MetadataObjectType field value if set, zero value otherwise.
func (o *KalturaMetadata) GetMetadataObjectType() string {
	if o == nil || o.MetadataObjectType == nil {
		var ret string
		return ret
	}
	return *o.MetadataObjectType
}

// GetMetadataObjectTypeOk returns a tuple with the MetadataObjectType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaMetadata) GetMetadataObjectTypeOk() (*string, bool) {
	if o == nil || o.MetadataObjectType == nil {
		return nil, false
	}
	return o.MetadataObjectType, true
}

// HasMetadataObjectType returns a boolean if a field has been set.
func (o *KalturaMetadata) HasMetadataObjectType() bool {
	if o != nil && o.MetadataObjectType != nil {
		return true
	}

	return false
}

// SetMetadataObjectType gets a reference to the given string and assigns it to the MetadataObjectType field.
func (o *KalturaMetadata) SetMetadataObjectType(v string) {
	o.MetadataObjectType = &v
}

// GetMetadataProfileId returns the MetadataProfileId field value if set, zero value otherwise.
func (o *KalturaMetadata) GetMetadataProfileId() int32 {
	if o == nil || o.MetadataProfileId == nil {
		var ret int32
		return ret
	}
	return *o.MetadataProfileId
}

// GetMetadataProfileIdOk returns a tuple with the MetadataProfileId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaMetadata) GetMetadataProfileIdOk() (*int32, bool) {
	if o == nil || o.MetadataProfileId == nil {
		return nil, false
	}
	return o.MetadataProfileId, true
}

// HasMetadataProfileId returns a boolean if a field has been set.
func (o *KalturaMetadata) HasMetadataProfileId() bool {
	if o != nil && o.MetadataProfileId != nil {
		return true
	}

	return false
}

// SetMetadataProfileId gets a reference to the given int32 and assigns it to the MetadataProfileId field.
func (o *KalturaMetadata) SetMetadataProfileId(v int32) {
	o.MetadataProfileId = &v
}

// GetMetadataProfileVersion returns the MetadataProfileVersion field value if set, zero value otherwise.
func (o *KalturaMetadata) GetMetadataProfileVersion() int32 {
	if o == nil || o.MetadataProfileVersion == nil {
		var ret int32
		return ret
	}
	return *o.MetadataProfileVersion
}

// GetMetadataProfileVersionOk returns a tuple with the MetadataProfileVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaMetadata) GetMetadataProfileVersionOk() (*int32, bool) {
	if o == nil || o.MetadataProfileVersion == nil {
		return nil, false
	}
	return o.MetadataProfileVersion, true
}

// HasMetadataProfileVersion returns a boolean if a field has been set.
func (o *KalturaMetadata) HasMetadataProfileVersion() bool {
	if o != nil && o.MetadataProfileVersion != nil {
		return true
	}

	return false
}

// SetMetadataProfileVersion gets a reference to the given int32 and assigns it to the MetadataProfileVersion field.
func (o *KalturaMetadata) SetMetadataProfileVersion(v int32) {
	o.MetadataProfileVersion = &v
}

// GetObjectId returns the ObjectId field value if set, zero value otherwise.
func (o *KalturaMetadata) GetObjectId() string {
	if o == nil || o.ObjectId == nil {
		var ret string
		return ret
	}
	return *o.ObjectId
}

// GetObjectIdOk returns a tuple with the ObjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaMetadata) GetObjectIdOk() (*string, bool) {
	if o == nil || o.ObjectId == nil {
		return nil, false
	}
	return o.ObjectId, true
}

// HasObjectId returns a boolean if a field has been set.
func (o *KalturaMetadata) HasObjectId() bool {
	if o != nil && o.ObjectId != nil {
		return true
	}

	return false
}

// SetObjectId gets a reference to the given string and assigns it to the ObjectId field.
func (o *KalturaMetadata) SetObjectId(v string) {
	o.ObjectId = &v
}

// GetPartnerId returns the PartnerId field value if set, zero value otherwise.
func (o *KalturaMetadata) GetPartnerId() int32 {
	if o == nil || o.PartnerId == nil {
		var ret int32
		return ret
	}
	return *o.PartnerId
}

// GetPartnerIdOk returns a tuple with the PartnerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaMetadata) GetPartnerIdOk() (*int32, bool) {
	if o == nil || o.PartnerId == nil {
		return nil, false
	}
	return o.PartnerId, true
}

// HasPartnerId returns a boolean if a field has been set.
func (o *KalturaMetadata) HasPartnerId() bool {
	if o != nil && o.PartnerId != nil {
		return true
	}

	return false
}

// SetPartnerId gets a reference to the given int32 and assigns it to the PartnerId field.
func (o *KalturaMetadata) SetPartnerId(v int32) {
	o.PartnerId = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *KalturaMetadata) GetStatus() int32 {
	if o == nil || o.Status == nil {
		var ret int32
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaMetadata) GetStatusOk() (*int32, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *KalturaMetadata) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given int32 and assigns it to the Status field.
func (o *KalturaMetadata) SetStatus(v int32) {
	o.Status = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *KalturaMetadata) GetUpdatedAt() int32 {
	if o == nil || o.UpdatedAt == nil {
		var ret int32
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaMetadata) GetUpdatedAtOk() (*int32, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *KalturaMetadata) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt != nil {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given int32 and assigns it to the UpdatedAt field.
func (o *KalturaMetadata) SetUpdatedAt(v int32) {
	o.UpdatedAt = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *KalturaMetadata) GetVersion() int32 {
	if o == nil || o.Version == nil {
		var ret int32
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaMetadata) GetVersionOk() (*int32, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *KalturaMetadata) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given int32 and assigns it to the Version field.
func (o *KalturaMetadata) SetVersion(v int32) {
	o.Version = &v
}

// GetXml returns the Xml field value if set, zero value otherwise.
func (o *KalturaMetadata) GetXml() string {
	if o == nil || o.Xml == nil {
		var ret string
		return ret
	}
	return *o.Xml
}

// GetXmlOk returns a tuple with the Xml field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaMetadata) GetXmlOk() (*string, bool) {
	if o == nil || o.Xml == nil {
		return nil, false
	}
	return o.Xml, true
}

// HasXml returns a boolean if a field has been set.
func (o *KalturaMetadata) HasXml() bool {
	if o != nil && o.Xml != nil {
		return true
	}

	return false
}

// SetXml gets a reference to the given string and assigns it to the Xml field.
func (o *KalturaMetadata) SetXml(v string) {
	o.Xml = &v
}

func (o KalturaMetadata) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CreatedAt != nil {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.MetadataObjectType != nil {
		toSerialize["metadataObjectType"] = o.MetadataObjectType
	}
	if o.MetadataProfileId != nil {
		toSerialize["metadataProfileId"] = o.MetadataProfileId
	}
	if o.MetadataProfileVersion != nil {
		toSerialize["metadataProfileVersion"] = o.MetadataProfileVersion
	}
	if o.ObjectId != nil {
		toSerialize["objectId"] = o.ObjectId
	}
	if o.PartnerId != nil {
		toSerialize["partnerId"] = o.PartnerId
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.UpdatedAt != nil {
		toSerialize["updatedAt"] = o.UpdatedAt
	}
	if o.Version != nil {
		toSerialize["version"] = o.Version
	}
	if o.Xml != nil {
		toSerialize["xml"] = o.Xml
	}
	return json.Marshal(toSerialize)
}

type NullableKalturaMetadata struct {
	value *KalturaMetadata
	isSet bool
}

func (v NullableKalturaMetadata) Get() *KalturaMetadata {
	return v.value
}

func (v *NullableKalturaMetadata) Set(val *KalturaMetadata) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaMetadata) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaMetadata) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaMetadata(val *KalturaMetadata) *NullableKalturaMetadata {
	return &NullableKalturaMetadata{value: val, isSet: true}
}

func (v NullableKalturaMetadata) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaMetadata) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


