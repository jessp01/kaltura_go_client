# CategoryDeleteRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**MoveEntriesToParentCategory** | Pointer to **int32** |  | [optional] 

## Methods

### NewCategoryDeleteRequest

`func NewCategoryDeleteRequest(id int32, ) *CategoryDeleteRequest`

NewCategoryDeleteRequest instantiates a new CategoryDeleteRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCategoryDeleteRequestWithDefaults

`func NewCategoryDeleteRequestWithDefaults() *CategoryDeleteRequest`

NewCategoryDeleteRequestWithDefaults instantiates a new CategoryDeleteRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CategoryDeleteRequest) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CategoryDeleteRequest) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CategoryDeleteRequest) SetId(v int32)`

SetId sets Id field to given value.


### GetMoveEntriesToParentCategory

`func (o *CategoryDeleteRequest) GetMoveEntriesToParentCategory() int32`

GetMoveEntriesToParentCategory returns the MoveEntriesToParentCategory field if non-nil, zero value otherwise.

### GetMoveEntriesToParentCategoryOk

`func (o *CategoryDeleteRequest) GetMoveEntriesToParentCategoryOk() (*int32, bool)`

GetMoveEntriesToParentCategoryOk returns a tuple with the MoveEntriesToParentCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoveEntriesToParentCategory

`func (o *CategoryDeleteRequest) SetMoveEntriesToParentCategory(v int32)`

SetMoveEntriesToParentCategory sets MoveEntriesToParentCategory field to given value.

### HasMoveEntriesToParentCategory

`func (o *CategoryDeleteRequest) HasMoveEntriesToParentCategory() bool`

HasMoveEntriesToParentCategory returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


