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

// KalturaBusinessProcessServer `abstract`
type KalturaBusinessProcessServer struct {
	// `readOnly`  Server creation date as Unix timestamp (In seconds)
	CreatedAt *int32 `json:"createdAt,omitempty"`
	// The dc of the server
	Dc *int32 `json:"dc,omitempty"`
	Description *string `json:"description,omitempty"`
	// `readOnly`  Auto generated identifier
	Id *int32 `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	ObjectType *string `json:"objectType,omitempty"`
	// `readOnly`
	PartnerId *int32 `json:"partnerId,omitempty"`
	// `readOnly`  Enum Type: `KalturaBusinessProcessServerStatus`
	Status *string `json:"status,omitempty"`
	SystemName *string `json:"systemName,omitempty"`
	// `readOnly`  Enum Type: `KalturaBusinessProcessProvider`  The type of the server, this is auto filled by the derived server object
	Type *string `json:"type,omitempty"`
	// `readOnly`  Server update date as Unix timestamp (In seconds)
	UpdatedAt *int32 `json:"updatedAt,omitempty"`
}

// NewKalturaBusinessProcessServer instantiates a new KalturaBusinessProcessServer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaBusinessProcessServer() *KalturaBusinessProcessServer {
	this := KalturaBusinessProcessServer{}
	return &this
}

// NewKalturaBusinessProcessServerWithDefaults instantiates a new KalturaBusinessProcessServer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaBusinessProcessServerWithDefaults() *KalturaBusinessProcessServer {
	this := KalturaBusinessProcessServer{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *KalturaBusinessProcessServer) GetCreatedAt() int32 {
	if o == nil || o.CreatedAt == nil {
		var ret int32
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaBusinessProcessServer) GetCreatedAtOk() (*int32, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *KalturaBusinessProcessServer) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given int32 and assigns it to the CreatedAt field.
func (o *KalturaBusinessProcessServer) SetCreatedAt(v int32) {
	o.CreatedAt = &v
}

// GetDc returns the Dc field value if set, zero value otherwise.
func (o *KalturaBusinessProcessServer) GetDc() int32 {
	if o == nil || o.Dc == nil {
		var ret int32
		return ret
	}
	return *o.Dc
}

// GetDcOk returns a tuple with the Dc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaBusinessProcessServer) GetDcOk() (*int32, bool) {
	if o == nil || o.Dc == nil {
		return nil, false
	}
	return o.Dc, true
}

// HasDc returns a boolean if a field has been set.
func (o *KalturaBusinessProcessServer) HasDc() bool {
	if o != nil && o.Dc != nil {
		return true
	}

	return false
}

// SetDc gets a reference to the given int32 and assigns it to the Dc field.
func (o *KalturaBusinessProcessServer) SetDc(v int32) {
	o.Dc = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *KalturaBusinessProcessServer) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaBusinessProcessServer) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *KalturaBusinessProcessServer) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *KalturaBusinessProcessServer) SetDescription(v string) {
	o.Description = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *KalturaBusinessProcessServer) GetId() int32 {
	if o == nil || o.Id == nil {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaBusinessProcessServer) GetIdOk() (*int32, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *KalturaBusinessProcessServer) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *KalturaBusinessProcessServer) SetId(v int32) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *KalturaBusinessProcessServer) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaBusinessProcessServer) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *KalturaBusinessProcessServer) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *KalturaBusinessProcessServer) SetName(v string) {
	o.Name = &v
}

// GetObjectType returns the ObjectType field value if set, zero value otherwise.
func (o *KalturaBusinessProcessServer) GetObjectType() string {
	if o == nil || o.ObjectType == nil {
		var ret string
		return ret
	}
	return *o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaBusinessProcessServer) GetObjectTypeOk() (*string, bool) {
	if o == nil || o.ObjectType == nil {
		return nil, false
	}
	return o.ObjectType, true
}

// HasObjectType returns a boolean if a field has been set.
func (o *KalturaBusinessProcessServer) HasObjectType() bool {
	if o != nil && o.ObjectType != nil {
		return true
	}

	return false
}

// SetObjectType gets a reference to the given string and assigns it to the ObjectType field.
func (o *KalturaBusinessProcessServer) SetObjectType(v string) {
	o.ObjectType = &v
}

// GetPartnerId returns the PartnerId field value if set, zero value otherwise.
func (o *KalturaBusinessProcessServer) GetPartnerId() int32 {
	if o == nil || o.PartnerId == nil {
		var ret int32
		return ret
	}
	return *o.PartnerId
}

// GetPartnerIdOk returns a tuple with the PartnerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaBusinessProcessServer) GetPartnerIdOk() (*int32, bool) {
	if o == nil || o.PartnerId == nil {
		return nil, false
	}
	return o.PartnerId, true
}

// HasPartnerId returns a boolean if a field has been set.
func (o *KalturaBusinessProcessServer) HasPartnerId() bool {
	if o != nil && o.PartnerId != nil {
		return true
	}

	return false
}

// SetPartnerId gets a reference to the given int32 and assigns it to the PartnerId field.
func (o *KalturaBusinessProcessServer) SetPartnerId(v int32) {
	o.PartnerId = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *KalturaBusinessProcessServer) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaBusinessProcessServer) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *KalturaBusinessProcessServer) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *KalturaBusinessProcessServer) SetStatus(v string) {
	o.Status = &v
}

// GetSystemName returns the SystemName field value if set, zero value otherwise.
func (o *KalturaBusinessProcessServer) GetSystemName() string {
	if o == nil || o.SystemName == nil {
		var ret string
		return ret
	}
	return *o.SystemName
}

// GetSystemNameOk returns a tuple with the SystemName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaBusinessProcessServer) GetSystemNameOk() (*string, bool) {
	if o == nil || o.SystemName == nil {
		return nil, false
	}
	return o.SystemName, true
}

// HasSystemName returns a boolean if a field has been set.
func (o *KalturaBusinessProcessServer) HasSystemName() bool {
	if o != nil && o.SystemName != nil {
		return true
	}

	return false
}

// SetSystemName gets a reference to the given string and assigns it to the SystemName field.
func (o *KalturaBusinessProcessServer) SetSystemName(v string) {
	o.SystemName = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *KalturaBusinessProcessServer) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaBusinessProcessServer) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *KalturaBusinessProcessServer) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *KalturaBusinessProcessServer) SetType(v string) {
	o.Type = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *KalturaBusinessProcessServer) GetUpdatedAt() int32 {
	if o == nil || o.UpdatedAt == nil {
		var ret int32
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaBusinessProcessServer) GetUpdatedAtOk() (*int32, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *KalturaBusinessProcessServer) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt != nil {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given int32 and assigns it to the UpdatedAt field.
func (o *KalturaBusinessProcessServer) SetUpdatedAt(v int32) {
	o.UpdatedAt = &v
}

func (o KalturaBusinessProcessServer) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CreatedAt != nil {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if o.Dc != nil {
		toSerialize["dc"] = o.Dc
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.ObjectType != nil {
		toSerialize["objectType"] = o.ObjectType
	}
	if o.PartnerId != nil {
		toSerialize["partnerId"] = o.PartnerId
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.SystemName != nil {
		toSerialize["systemName"] = o.SystemName
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.UpdatedAt != nil {
		toSerialize["updatedAt"] = o.UpdatedAt
	}
	return json.Marshal(toSerialize)
}

type NullableKalturaBusinessProcessServer struct {
	value *KalturaBusinessProcessServer
	isSet bool
}

func (v NullableKalturaBusinessProcessServer) Get() *KalturaBusinessProcessServer {
	return v.value
}

func (v *NullableKalturaBusinessProcessServer) Set(val *KalturaBusinessProcessServer) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaBusinessProcessServer) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaBusinessProcessServer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaBusinessProcessServer(val *KalturaBusinessProcessServer) *NullableKalturaBusinessProcessServer {
	return &NullableKalturaBusinessProcessServer{value: val, isSet: true}
}

func (v NullableKalturaBusinessProcessServer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaBusinessProcessServer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


