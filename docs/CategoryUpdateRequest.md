# CategoryUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Category** | [**KalturaCategory**](KalturaCategory.md) |  | 
**Id** | **int32** |  | 

## Methods

### NewCategoryUpdateRequest

`func NewCategoryUpdateRequest(category KalturaCategory, id int32, ) *CategoryUpdateRequest`

NewCategoryUpdateRequest instantiates a new CategoryUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCategoryUpdateRequestWithDefaults

`func NewCategoryUpdateRequestWithDefaults() *CategoryUpdateRequest`

NewCategoryUpdateRequestWithDefaults instantiates a new CategoryUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCategory

`func (o *CategoryUpdateRequest) GetCategory() KalturaCategory`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *CategoryUpdateRequest) GetCategoryOk() (*KalturaCategory, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *CategoryUpdateRequest) SetCategory(v KalturaCategory)`

SetCategory sets Category field to given value.


### GetId

`func (o *CategoryUpdateRequest) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CategoryUpdateRequest) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CategoryUpdateRequest) SetId(v int32)`

SetId sets Id field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


