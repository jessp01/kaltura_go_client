# CategoryUserUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CategoryId** | **int32** |  | 
**CategoryUser** | [**KalturaCategoryUser**](KalturaCategoryUser.md) |  | 
**Override** | Pointer to **bool** |  | [optional] [default to false]
**UserId** | **string** |  | 

## Methods

### NewCategoryUserUpdateRequest

`func NewCategoryUserUpdateRequest(categoryId int32, categoryUser KalturaCategoryUser, userId string, ) *CategoryUserUpdateRequest`

NewCategoryUserUpdateRequest instantiates a new CategoryUserUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCategoryUserUpdateRequestWithDefaults

`func NewCategoryUserUpdateRequestWithDefaults() *CategoryUserUpdateRequest`

NewCategoryUserUpdateRequestWithDefaults instantiates a new CategoryUserUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCategoryId

`func (o *CategoryUserUpdateRequest) GetCategoryId() int32`

GetCategoryId returns the CategoryId field if non-nil, zero value otherwise.

### GetCategoryIdOk

`func (o *CategoryUserUpdateRequest) GetCategoryIdOk() (*int32, bool)`

GetCategoryIdOk returns a tuple with the CategoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryId

`func (o *CategoryUserUpdateRequest) SetCategoryId(v int32)`

SetCategoryId sets CategoryId field to given value.


### GetCategoryUser

`func (o *CategoryUserUpdateRequest) GetCategoryUser() KalturaCategoryUser`

GetCategoryUser returns the CategoryUser field if non-nil, zero value otherwise.

### GetCategoryUserOk

`func (o *CategoryUserUpdateRequest) GetCategoryUserOk() (*KalturaCategoryUser, bool)`

GetCategoryUserOk returns a tuple with the CategoryUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryUser

`func (o *CategoryUserUpdateRequest) SetCategoryUser(v KalturaCategoryUser)`

SetCategoryUser sets CategoryUser field to given value.


### GetOverride

`func (o *CategoryUserUpdateRequest) GetOverride() bool`

GetOverride returns the Override field if non-nil, zero value otherwise.

### GetOverrideOk

`func (o *CategoryUserUpdateRequest) GetOverrideOk() (*bool, bool)`

GetOverrideOk returns a tuple with the Override field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverride

`func (o *CategoryUserUpdateRequest) SetOverride(v bool)`

SetOverride sets Override field to given value.

### HasOverride

`func (o *CategoryUserUpdateRequest) HasOverride() bool`

HasOverride returns a boolean if a field has been set.

### GetUserId

`func (o *CategoryUserUpdateRequest) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *CategoryUserUpdateRequest) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *CategoryUserUpdateRequest) SetUserId(v string)`

SetUserId sets UserId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


