# KalturaServerNode

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **int32** | &#x60;readOnly&#x60; | [optional] [readonly] 
**Dc** | Pointer to **int32** | &#x60;readOnly&#x60;  DC where the serverNode is located | [optional] [readonly] 
**Description** | Pointer to **string** |  | [optional] 
**Environment** | Pointer to **string** | Environment | [optional] 
**HeartbeatTime** | Pointer to **int32** | &#x60;readOnly&#x60; | [optional] [readonly] 
**HostName** | Pointer to **string** | serverNode hostName | [optional] 
**Id** | Pointer to **int32** | &#x60;readOnly&#x60; | [optional] [readonly] 
**Name** | Pointer to **string** | serverNode name | [optional] 
**ObjectType** | Pointer to **string** |  | [optional] 
**ParentId** | Pointer to **string** | Id of the parent serverNode | [optional] 
**PartnerId** | Pointer to **int32** | &#x60;readOnly&#x60; | [optional] [readonly] 
**Status** | Pointer to **int32** | &#x60;readOnly&#x60;  Enum Type: &#x60;KalturaServerNodeStatus&#x60; | [optional] [readonly] 
**SystemName** | Pointer to **string** | serverNode uniqe system name | [optional] 
**Tags** | Pointer to **string** | serverNode tags | [optional] 
**Type** | Pointer to **string** | &#x60;readOnly&#x60;  Enum Type: &#x60;KalturaServerNodeType&#x60; | [optional] [readonly] 
**UpdatedAt** | Pointer to **int32** | &#x60;readOnly&#x60; | [optional] [readonly] 

## Methods

### NewKalturaServerNode

`func NewKalturaServerNode() *KalturaServerNode`

NewKalturaServerNode instantiates a new KalturaServerNode object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKalturaServerNodeWithDefaults

`func NewKalturaServerNodeWithDefaults() *KalturaServerNode`

NewKalturaServerNodeWithDefaults instantiates a new KalturaServerNode object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *KalturaServerNode) GetCreatedAt() int32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *KalturaServerNode) GetCreatedAtOk() (*int32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *KalturaServerNode) SetCreatedAt(v int32)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *KalturaServerNode) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDc

`func (o *KalturaServerNode) GetDc() int32`

GetDc returns the Dc field if non-nil, zero value otherwise.

### GetDcOk

`func (o *KalturaServerNode) GetDcOk() (*int32, bool)`

GetDcOk returns a tuple with the Dc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDc

`func (o *KalturaServerNode) SetDc(v int32)`

SetDc sets Dc field to given value.

### HasDc

`func (o *KalturaServerNode) HasDc() bool`

HasDc returns a boolean if a field has been set.

### GetDescription

`func (o *KalturaServerNode) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *KalturaServerNode) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *KalturaServerNode) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *KalturaServerNode) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnvironment

`func (o *KalturaServerNode) GetEnvironment() string`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *KalturaServerNode) GetEnvironmentOk() (*string, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *KalturaServerNode) SetEnvironment(v string)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *KalturaServerNode) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetHeartbeatTime

`func (o *KalturaServerNode) GetHeartbeatTime() int32`

GetHeartbeatTime returns the HeartbeatTime field if non-nil, zero value otherwise.

### GetHeartbeatTimeOk

`func (o *KalturaServerNode) GetHeartbeatTimeOk() (*int32, bool)`

GetHeartbeatTimeOk returns a tuple with the HeartbeatTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeartbeatTime

`func (o *KalturaServerNode) SetHeartbeatTime(v int32)`

SetHeartbeatTime sets HeartbeatTime field to given value.

### HasHeartbeatTime

`func (o *KalturaServerNode) HasHeartbeatTime() bool`

HasHeartbeatTime returns a boolean if a field has been set.

### GetHostName

`func (o *KalturaServerNode) GetHostName() string`

GetHostName returns the HostName field if non-nil, zero value otherwise.

### GetHostNameOk

`func (o *KalturaServerNode) GetHostNameOk() (*string, bool)`

GetHostNameOk returns a tuple with the HostName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostName

`func (o *KalturaServerNode) SetHostName(v string)`

SetHostName sets HostName field to given value.

### HasHostName

`func (o *KalturaServerNode) HasHostName() bool`

HasHostName returns a boolean if a field has been set.

### GetId

`func (o *KalturaServerNode) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *KalturaServerNode) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *KalturaServerNode) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *KalturaServerNode) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *KalturaServerNode) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *KalturaServerNode) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *KalturaServerNode) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *KalturaServerNode) HasName() bool`

HasName returns a boolean if a field has been set.

### GetObjectType

`func (o *KalturaServerNode) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *KalturaServerNode) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *KalturaServerNode) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.

### HasObjectType

`func (o *KalturaServerNode) HasObjectType() bool`

HasObjectType returns a boolean if a field has been set.

### GetParentId

`func (o *KalturaServerNode) GetParentId() string`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *KalturaServerNode) GetParentIdOk() (*string, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *KalturaServerNode) SetParentId(v string)`

SetParentId sets ParentId field to given value.

### HasParentId

`func (o *KalturaServerNode) HasParentId() bool`

HasParentId returns a boolean if a field has been set.

### GetPartnerId

`func (o *KalturaServerNode) GetPartnerId() int32`

GetPartnerId returns the PartnerId field if non-nil, zero value otherwise.

### GetPartnerIdOk

`func (o *KalturaServerNode) GetPartnerIdOk() (*int32, bool)`

GetPartnerIdOk returns a tuple with the PartnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerId

`func (o *KalturaServerNode) SetPartnerId(v int32)`

SetPartnerId sets PartnerId field to given value.

### HasPartnerId

`func (o *KalturaServerNode) HasPartnerId() bool`

HasPartnerId returns a boolean if a field has been set.

### GetStatus

`func (o *KalturaServerNode) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *KalturaServerNode) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *KalturaServerNode) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *KalturaServerNode) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSystemName

`func (o *KalturaServerNode) GetSystemName() string`

GetSystemName returns the SystemName field if non-nil, zero value otherwise.

### GetSystemNameOk

`func (o *KalturaServerNode) GetSystemNameOk() (*string, bool)`

GetSystemNameOk returns a tuple with the SystemName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemName

`func (o *KalturaServerNode) SetSystemName(v string)`

SetSystemName sets SystemName field to given value.

### HasSystemName

`func (o *KalturaServerNode) HasSystemName() bool`

HasSystemName returns a boolean if a field has been set.

### GetTags

`func (o *KalturaServerNode) GetTags() string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *KalturaServerNode) GetTagsOk() (*string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *KalturaServerNode) SetTags(v string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *KalturaServerNode) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetType

`func (o *KalturaServerNode) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *KalturaServerNode) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *KalturaServerNode) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *KalturaServerNode) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *KalturaServerNode) GetUpdatedAt() int32`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *KalturaServerNode) GetUpdatedAtOk() (*int32, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *KalturaServerNode) SetUpdatedAt(v int32)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *KalturaServerNode) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


