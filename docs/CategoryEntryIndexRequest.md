# CategoryEntryIndexRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CategoryId** | **int32** |  | 
**EntryId** | **string** |  | 
**ShouldUpdate** | Pointer to **bool** |  | [optional] [default to true]

## Methods

### NewCategoryEntryIndexRequest

`func NewCategoryEntryIndexRequest(categoryId int32, entryId string, ) *CategoryEntryIndexRequest`

NewCategoryEntryIndexRequest instantiates a new CategoryEntryIndexRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCategoryEntryIndexRequestWithDefaults

`func NewCategoryEntryIndexRequestWithDefaults() *CategoryEntryIndexRequest`

NewCategoryEntryIndexRequestWithDefaults instantiates a new CategoryEntryIndexRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCategoryId

`func (o *CategoryEntryIndexRequest) GetCategoryId() int32`

GetCategoryId returns the CategoryId field if non-nil, zero value otherwise.

### GetCategoryIdOk

`func (o *CategoryEntryIndexRequest) GetCategoryIdOk() (*int32, bool)`

GetCategoryIdOk returns a tuple with the CategoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryId

`func (o *CategoryEntryIndexRequest) SetCategoryId(v int32)`

SetCategoryId sets CategoryId field to given value.


### GetEntryId

`func (o *CategoryEntryIndexRequest) GetEntryId() string`

GetEntryId returns the EntryId field if non-nil, zero value otherwise.

### GetEntryIdOk

`func (o *CategoryEntryIndexRequest) GetEntryIdOk() (*string, bool)`

GetEntryIdOk returns a tuple with the EntryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntryId

`func (o *CategoryEntryIndexRequest) SetEntryId(v string)`

SetEntryId sets EntryId field to given value.


### GetShouldUpdate

`func (o *CategoryEntryIndexRequest) GetShouldUpdate() bool`

GetShouldUpdate returns the ShouldUpdate field if non-nil, zero value otherwise.

### GetShouldUpdateOk

`func (o *CategoryEntryIndexRequest) GetShouldUpdateOk() (*bool, bool)`

GetShouldUpdateOk returns a tuple with the ShouldUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShouldUpdate

`func (o *CategoryEntryIndexRequest) SetShouldUpdate(v bool)`

SetShouldUpdate sets ShouldUpdate field to given value.

### HasShouldUpdate

`func (o *CategoryEntryIndexRequest) HasShouldUpdate() bool`

HasShouldUpdate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


