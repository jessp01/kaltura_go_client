# KalturaLike

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **int32** | The date of the like&#39;s creation | [optional] 
**EntryId** | Pointer to **string** | The id of the entry that the like belongs to | [optional] 
**UserId** | Pointer to **string** | The id of user that the like belongs to | [optional] 

## Methods

### NewKalturaLike

`func NewKalturaLike() *KalturaLike`

NewKalturaLike instantiates a new KalturaLike object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKalturaLikeWithDefaults

`func NewKalturaLikeWithDefaults() *KalturaLike`

NewKalturaLikeWithDefaults instantiates a new KalturaLike object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *KalturaLike) GetCreatedAt() int32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *KalturaLike) GetCreatedAtOk() (*int32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *KalturaLike) SetCreatedAt(v int32)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *KalturaLike) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetEntryId

`func (o *KalturaLike) GetEntryId() string`

GetEntryId returns the EntryId field if non-nil, zero value otherwise.

### GetEntryIdOk

`func (o *KalturaLike) GetEntryIdOk() (*string, bool)`

GetEntryIdOk returns a tuple with the EntryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntryId

`func (o *KalturaLike) SetEntryId(v string)`

SetEntryId sets EntryId field to given value.

### HasEntryId

`func (o *KalturaLike) HasEntryId() bool`

HasEntryId returns a boolean if a field has been set.

### GetUserId

`func (o *KalturaLike) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *KalturaLike) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *KalturaLike) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *KalturaLike) HasUserId() bool`

HasUserId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


