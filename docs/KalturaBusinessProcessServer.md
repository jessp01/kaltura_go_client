# KalturaBusinessProcessServer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **int32** | &#x60;readOnly&#x60;  Server creation date as Unix timestamp (In seconds) | [optional] [readonly] 
**Dc** | Pointer to **int32** | The dc of the server | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **int32** | &#x60;readOnly&#x60;  Auto generated identifier | [optional] [readonly] 
**Name** | Pointer to **string** |  | [optional] 
**ObjectType** | Pointer to **string** |  | [optional] 
**PartnerId** | Pointer to **int32** | &#x60;readOnly&#x60; | [optional] [readonly] 
**Status** | Pointer to **string** | &#x60;readOnly&#x60;  Enum Type: &#x60;KalturaBusinessProcessServerStatus&#x60; | [optional] [readonly] 
**SystemName** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** | &#x60;readOnly&#x60;  Enum Type: &#x60;KalturaBusinessProcessProvider&#x60;  The type of the server, this is auto filled by the derived server object | [optional] [readonly] 
**UpdatedAt** | Pointer to **int32** | &#x60;readOnly&#x60;  Server update date as Unix timestamp (In seconds) | [optional] [readonly] 

## Methods

### NewKalturaBusinessProcessServer

`func NewKalturaBusinessProcessServer() *KalturaBusinessProcessServer`

NewKalturaBusinessProcessServer instantiates a new KalturaBusinessProcessServer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKalturaBusinessProcessServerWithDefaults

`func NewKalturaBusinessProcessServerWithDefaults() *KalturaBusinessProcessServer`

NewKalturaBusinessProcessServerWithDefaults instantiates a new KalturaBusinessProcessServer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *KalturaBusinessProcessServer) GetCreatedAt() int32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *KalturaBusinessProcessServer) GetCreatedAtOk() (*int32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *KalturaBusinessProcessServer) SetCreatedAt(v int32)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *KalturaBusinessProcessServer) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDc

`func (o *KalturaBusinessProcessServer) GetDc() int32`

GetDc returns the Dc field if non-nil, zero value otherwise.

### GetDcOk

`func (o *KalturaBusinessProcessServer) GetDcOk() (*int32, bool)`

GetDcOk returns a tuple with the Dc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDc

`func (o *KalturaBusinessProcessServer) SetDc(v int32)`

SetDc sets Dc field to given value.

### HasDc

`func (o *KalturaBusinessProcessServer) HasDc() bool`

HasDc returns a boolean if a field has been set.

### GetDescription

`func (o *KalturaBusinessProcessServer) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *KalturaBusinessProcessServer) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *KalturaBusinessProcessServer) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *KalturaBusinessProcessServer) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetId

`func (o *KalturaBusinessProcessServer) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *KalturaBusinessProcessServer) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *KalturaBusinessProcessServer) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *KalturaBusinessProcessServer) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *KalturaBusinessProcessServer) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *KalturaBusinessProcessServer) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *KalturaBusinessProcessServer) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *KalturaBusinessProcessServer) HasName() bool`

HasName returns a boolean if a field has been set.

### GetObjectType

`func (o *KalturaBusinessProcessServer) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *KalturaBusinessProcessServer) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *KalturaBusinessProcessServer) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.

### HasObjectType

`func (o *KalturaBusinessProcessServer) HasObjectType() bool`

HasObjectType returns a boolean if a field has been set.

### GetPartnerId

`func (o *KalturaBusinessProcessServer) GetPartnerId() int32`

GetPartnerId returns the PartnerId field if non-nil, zero value otherwise.

### GetPartnerIdOk

`func (o *KalturaBusinessProcessServer) GetPartnerIdOk() (*int32, bool)`

GetPartnerIdOk returns a tuple with the PartnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerId

`func (o *KalturaBusinessProcessServer) SetPartnerId(v int32)`

SetPartnerId sets PartnerId field to given value.

### HasPartnerId

`func (o *KalturaBusinessProcessServer) HasPartnerId() bool`

HasPartnerId returns a boolean if a field has been set.

### GetStatus

`func (o *KalturaBusinessProcessServer) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *KalturaBusinessProcessServer) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *KalturaBusinessProcessServer) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *KalturaBusinessProcessServer) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSystemName

`func (o *KalturaBusinessProcessServer) GetSystemName() string`

GetSystemName returns the SystemName field if non-nil, zero value otherwise.

### GetSystemNameOk

`func (o *KalturaBusinessProcessServer) GetSystemNameOk() (*string, bool)`

GetSystemNameOk returns a tuple with the SystemName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemName

`func (o *KalturaBusinessProcessServer) SetSystemName(v string)`

SetSystemName sets SystemName field to given value.

### HasSystemName

`func (o *KalturaBusinessProcessServer) HasSystemName() bool`

HasSystemName returns a boolean if a field has been set.

### GetType

`func (o *KalturaBusinessProcessServer) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *KalturaBusinessProcessServer) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *KalturaBusinessProcessServer) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *KalturaBusinessProcessServer) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *KalturaBusinessProcessServer) GetUpdatedAt() int32`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *KalturaBusinessProcessServer) GetUpdatedAtOk() (*int32, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *KalturaBusinessProcessServer) SetUpdatedAt(v int32)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *KalturaBusinessProcessServer) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


