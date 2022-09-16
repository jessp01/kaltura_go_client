# CategoryCloneRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CategoryId** | **int32** |  | 
**FromPartnerId** | **int32** |  | 
**ParentCategoryId** | Pointer to **int32** |  | [optional] 

## Methods

### NewCategoryCloneRequest

`func NewCategoryCloneRequest(categoryId int32, fromPartnerId int32, ) *CategoryCloneRequest`

NewCategoryCloneRequest instantiates a new CategoryCloneRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCategoryCloneRequestWithDefaults

`func NewCategoryCloneRequestWithDefaults() *CategoryCloneRequest`

NewCategoryCloneRequestWithDefaults instantiates a new CategoryCloneRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCategoryId

`func (o *CategoryCloneRequest) GetCategoryId() int32`

GetCategoryId returns the CategoryId field if non-nil, zero value otherwise.

### GetCategoryIdOk

`func (o *CategoryCloneRequest) GetCategoryIdOk() (*int32, bool)`

GetCategoryIdOk returns a tuple with the CategoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryId

`func (o *CategoryCloneRequest) SetCategoryId(v int32)`

SetCategoryId sets CategoryId field to given value.


### GetFromPartnerId

`func (o *CategoryCloneRequest) GetFromPartnerId() int32`

GetFromPartnerId returns the FromPartnerId field if non-nil, zero value otherwise.

### GetFromPartnerIdOk

`func (o *CategoryCloneRequest) GetFromPartnerIdOk() (*int32, bool)`

GetFromPartnerIdOk returns a tuple with the FromPartnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromPartnerId

`func (o *CategoryCloneRequest) SetFromPartnerId(v int32)`

SetFromPartnerId sets FromPartnerId field to given value.


### GetParentCategoryId

`func (o *CategoryCloneRequest) GetParentCategoryId() int32`

GetParentCategoryId returns the ParentCategoryId field if non-nil, zero value otherwise.

### GetParentCategoryIdOk

`func (o *CategoryCloneRequest) GetParentCategoryIdOk() (*int32, bool)`

GetParentCategoryIdOk returns a tuple with the ParentCategoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentCategoryId

`func (o *CategoryCloneRequest) SetParentCategoryId(v int32)`

SetParentCategoryId sets ParentCategoryId field to given value.

### HasParentCategoryId

`func (o *CategoryCloneRequest) HasParentCategoryId() bool`

HasParentCategoryId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


