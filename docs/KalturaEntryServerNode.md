# KalturaEntryServerNode

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **int32** | &#x60;readOnly&#x60; | [optional] [readonly] 
**EntryId** | Pointer to **string** | &#x60;readOnly&#x60; | [optional] [readonly] 
**Id** | Pointer to **int32** | &#x60;readOnly&#x60;  unique auto-generated identifier | [optional] [readonly] 
**ObjectType** | Pointer to **string** |  | [optional] 
**PartnerId** | Pointer to **int32** | &#x60;readOnly&#x60; | [optional] [readonly] 
**ServerNodeId** | Pointer to **int32** | &#x60;readOnly&#x60; | [optional] [readonly] 
**ServerType** | Pointer to **string** | &#x60;readOnly&#x60;  Enum Type: &#x60;KalturaEntryServerNodeType&#x60; | [optional] [readonly] 
**Status** | Pointer to **int32** | &#x60;readOnly&#x60;  Enum Type: &#x60;KalturaEntryServerNodeStatus&#x60; | [optional] [readonly] 
**UpdatedAt** | Pointer to **int32** | &#x60;readOnly&#x60; | [optional] [readonly] 

## Methods

### NewKalturaEntryServerNode

`func NewKalturaEntryServerNode() *KalturaEntryServerNode`

NewKalturaEntryServerNode instantiates a new KalturaEntryServerNode object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKalturaEntryServerNodeWithDefaults

`func NewKalturaEntryServerNodeWithDefaults() *KalturaEntryServerNode`

NewKalturaEntryServerNodeWithDefaults instantiates a new KalturaEntryServerNode object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *KalturaEntryServerNode) GetCreatedAt() int32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *KalturaEntryServerNode) GetCreatedAtOk() (*int32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *KalturaEntryServerNode) SetCreatedAt(v int32)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *KalturaEntryServerNode) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetEntryId

`func (o *KalturaEntryServerNode) GetEntryId() string`

GetEntryId returns the EntryId field if non-nil, zero value otherwise.

### GetEntryIdOk

`func (o *KalturaEntryServerNode) GetEntryIdOk() (*string, bool)`

GetEntryIdOk returns a tuple with the EntryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntryId

`func (o *KalturaEntryServerNode) SetEntryId(v string)`

SetEntryId sets EntryId field to given value.

### HasEntryId

`func (o *KalturaEntryServerNode) HasEntryId() bool`

HasEntryId returns a boolean if a field has been set.

### GetId

`func (o *KalturaEntryServerNode) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *KalturaEntryServerNode) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *KalturaEntryServerNode) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *KalturaEntryServerNode) HasId() bool`

HasId returns a boolean if a field has been set.

### GetObjectType

`func (o *KalturaEntryServerNode) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *KalturaEntryServerNode) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *KalturaEntryServerNode) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.

### HasObjectType

`func (o *KalturaEntryServerNode) HasObjectType() bool`

HasObjectType returns a boolean if a field has been set.

### GetPartnerId

`func (o *KalturaEntryServerNode) GetPartnerId() int32`

GetPartnerId returns the PartnerId field if non-nil, zero value otherwise.

### GetPartnerIdOk

`func (o *KalturaEntryServerNode) GetPartnerIdOk() (*int32, bool)`

GetPartnerIdOk returns a tuple with the PartnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerId

`func (o *KalturaEntryServerNode) SetPartnerId(v int32)`

SetPartnerId sets PartnerId field to given value.

### HasPartnerId

`func (o *KalturaEntryServerNode) HasPartnerId() bool`

HasPartnerId returns a boolean if a field has been set.

### GetServerNodeId

`func (o *KalturaEntryServerNode) GetServerNodeId() int32`

GetServerNodeId returns the ServerNodeId field if non-nil, zero value otherwise.

### GetServerNodeIdOk

`func (o *KalturaEntryServerNode) GetServerNodeIdOk() (*int32, bool)`

GetServerNodeIdOk returns a tuple with the ServerNodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerNodeId

`func (o *KalturaEntryServerNode) SetServerNodeId(v int32)`

SetServerNodeId sets ServerNodeId field to given value.

### HasServerNodeId

`func (o *KalturaEntryServerNode) HasServerNodeId() bool`

HasServerNodeId returns a boolean if a field has been set.

### GetServerType

`func (o *KalturaEntryServerNode) GetServerType() string`

GetServerType returns the ServerType field if non-nil, zero value otherwise.

### GetServerTypeOk

`func (o *KalturaEntryServerNode) GetServerTypeOk() (*string, bool)`

GetServerTypeOk returns a tuple with the ServerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerType

`func (o *KalturaEntryServerNode) SetServerType(v string)`

SetServerType sets ServerType field to given value.

### HasServerType

`func (o *KalturaEntryServerNode) HasServerType() bool`

HasServerType returns a boolean if a field has been set.

### GetStatus

`func (o *KalturaEntryServerNode) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *KalturaEntryServerNode) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *KalturaEntryServerNode) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *KalturaEntryServerNode) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *KalturaEntryServerNode) GetUpdatedAt() int32`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *KalturaEntryServerNode) GetUpdatedAtOk() (*int32, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *KalturaEntryServerNode) SetUpdatedAt(v int32)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *KalturaEntryServerNode) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


