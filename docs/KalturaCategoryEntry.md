# KalturaCategoryEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CategoryFullIds** | Pointer to **string** | &#x60;readOnly&#x60;  The full ids of the Category | [optional] [readonly] 
**CategoryId** | Pointer to **int32** |  | [optional] 
**CreatedAt** | Pointer to **int32** | &#x60;readOnly&#x60;  Creation date as Unix timestamp (In seconds) | [optional] [readonly] 
**CreatorUserId** | Pointer to **string** | &#x60;readOnly&#x60;  CategroyEntry creator puser ID | [optional] [readonly] 
**EntryId** | Pointer to **string** | entry id | [optional] 
**Status** | Pointer to **int32** | &#x60;readOnly&#x60;  Enum Type: &#x60;KalturaCategoryEntryStatus&#x60;  CategroyEntry status | [optional] [readonly] 

## Methods

### NewKalturaCategoryEntry

`func NewKalturaCategoryEntry() *KalturaCategoryEntry`

NewKalturaCategoryEntry instantiates a new KalturaCategoryEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKalturaCategoryEntryWithDefaults

`func NewKalturaCategoryEntryWithDefaults() *KalturaCategoryEntry`

NewKalturaCategoryEntryWithDefaults instantiates a new KalturaCategoryEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCategoryFullIds

`func (o *KalturaCategoryEntry) GetCategoryFullIds() string`

GetCategoryFullIds returns the CategoryFullIds field if non-nil, zero value otherwise.

### GetCategoryFullIdsOk

`func (o *KalturaCategoryEntry) GetCategoryFullIdsOk() (*string, bool)`

GetCategoryFullIdsOk returns a tuple with the CategoryFullIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryFullIds

`func (o *KalturaCategoryEntry) SetCategoryFullIds(v string)`

SetCategoryFullIds sets CategoryFullIds field to given value.

### HasCategoryFullIds

`func (o *KalturaCategoryEntry) HasCategoryFullIds() bool`

HasCategoryFullIds returns a boolean if a field has been set.

### GetCategoryId

`func (o *KalturaCategoryEntry) GetCategoryId() int32`

GetCategoryId returns the CategoryId field if non-nil, zero value otherwise.

### GetCategoryIdOk

`func (o *KalturaCategoryEntry) GetCategoryIdOk() (*int32, bool)`

GetCategoryIdOk returns a tuple with the CategoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryId

`func (o *KalturaCategoryEntry) SetCategoryId(v int32)`

SetCategoryId sets CategoryId field to given value.

### HasCategoryId

`func (o *KalturaCategoryEntry) HasCategoryId() bool`

HasCategoryId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *KalturaCategoryEntry) GetCreatedAt() int32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *KalturaCategoryEntry) GetCreatedAtOk() (*int32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *KalturaCategoryEntry) SetCreatedAt(v int32)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *KalturaCategoryEntry) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCreatorUserId

`func (o *KalturaCategoryEntry) GetCreatorUserId() string`

GetCreatorUserId returns the CreatorUserId field if non-nil, zero value otherwise.

### GetCreatorUserIdOk

`func (o *KalturaCategoryEntry) GetCreatorUserIdOk() (*string, bool)`

GetCreatorUserIdOk returns a tuple with the CreatorUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatorUserId

`func (o *KalturaCategoryEntry) SetCreatorUserId(v string)`

SetCreatorUserId sets CreatorUserId field to given value.

### HasCreatorUserId

`func (o *KalturaCategoryEntry) HasCreatorUserId() bool`

HasCreatorUserId returns a boolean if a field has been set.

### GetEntryId

`func (o *KalturaCategoryEntry) GetEntryId() string`

GetEntryId returns the EntryId field if non-nil, zero value otherwise.

### GetEntryIdOk

`func (o *KalturaCategoryEntry) GetEntryIdOk() (*string, bool)`

GetEntryIdOk returns a tuple with the EntryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntryId

`func (o *KalturaCategoryEntry) SetEntryId(v string)`

SetEntryId sets EntryId field to given value.

### HasEntryId

`func (o *KalturaCategoryEntry) HasEntryId() bool`

HasEntryId returns a boolean if a field has been set.

### GetStatus

`func (o *KalturaCategoryEntry) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *KalturaCategoryEntry) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *KalturaCategoryEntry) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *KalturaCategoryEntry) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


