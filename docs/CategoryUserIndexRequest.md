# CategoryUserIndexRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CategoryId** | **int32** |  | 
**ShouldUpdate** | Pointer to **bool** |  | [optional] [default to true]
**UserId** | **string** |  | 

## Methods

### NewCategoryUserIndexRequest

`func NewCategoryUserIndexRequest(categoryId int32, userId string, ) *CategoryUserIndexRequest`

NewCategoryUserIndexRequest instantiates a new CategoryUserIndexRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCategoryUserIndexRequestWithDefaults

`func NewCategoryUserIndexRequestWithDefaults() *CategoryUserIndexRequest`

NewCategoryUserIndexRequestWithDefaults instantiates a new CategoryUserIndexRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCategoryId

`func (o *CategoryUserIndexRequest) GetCategoryId() int32`

GetCategoryId returns the CategoryId field if non-nil, zero value otherwise.

### GetCategoryIdOk

`func (o *CategoryUserIndexRequest) GetCategoryIdOk() (*int32, bool)`

GetCategoryIdOk returns a tuple with the CategoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryId

`func (o *CategoryUserIndexRequest) SetCategoryId(v int32)`

SetCategoryId sets CategoryId field to given value.


### GetShouldUpdate

`func (o *CategoryUserIndexRequest) GetShouldUpdate() bool`

GetShouldUpdate returns the ShouldUpdate field if non-nil, zero value otherwise.

### GetShouldUpdateOk

`func (o *CategoryUserIndexRequest) GetShouldUpdateOk() (*bool, bool)`

GetShouldUpdateOk returns a tuple with the ShouldUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShouldUpdate

`func (o *CategoryUserIndexRequest) SetShouldUpdate(v bool)`

SetShouldUpdate sets ShouldUpdate field to given value.

### HasShouldUpdate

`func (o *CategoryUserIndexRequest) HasShouldUpdate() bool`

HasShouldUpdate returns a boolean if a field has been set.

### GetUserId

`func (o *CategoryUserIndexRequest) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *CategoryUserIndexRequest) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *CategoryUserIndexRequest) SetUserId(v string)`

SetUserId sets UserId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


