# CategoryMoveRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CategoryIds** | **string** |  | 
**TargetCategoryParentId** | **int32** |  | 

## Methods

### NewCategoryMoveRequest

`func NewCategoryMoveRequest(categoryIds string, targetCategoryParentId int32, ) *CategoryMoveRequest`

NewCategoryMoveRequest instantiates a new CategoryMoveRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCategoryMoveRequestWithDefaults

`func NewCategoryMoveRequestWithDefaults() *CategoryMoveRequest`

NewCategoryMoveRequestWithDefaults instantiates a new CategoryMoveRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCategoryIds

`func (o *CategoryMoveRequest) GetCategoryIds() string`

GetCategoryIds returns the CategoryIds field if non-nil, zero value otherwise.

### GetCategoryIdsOk

`func (o *CategoryMoveRequest) GetCategoryIdsOk() (*string, bool)`

GetCategoryIdsOk returns a tuple with the CategoryIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryIds

`func (o *CategoryMoveRequest) SetCategoryIds(v string)`

SetCategoryIds sets CategoryIds field to given value.


### GetTargetCategoryParentId

`func (o *CategoryMoveRequest) GetTargetCategoryParentId() int32`

GetTargetCategoryParentId returns the TargetCategoryParentId field if non-nil, zero value otherwise.

### GetTargetCategoryParentIdOk

`func (o *CategoryMoveRequest) GetTargetCategoryParentIdOk() (*int32, bool)`

GetTargetCategoryParentIdOk returns a tuple with the TargetCategoryParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetCategoryParentId

`func (o *CategoryMoveRequest) SetTargetCategoryParentId(v int32)`

SetTargetCategoryParentId sets TargetCategoryParentId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


