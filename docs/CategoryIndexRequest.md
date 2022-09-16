# CategoryIndexRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**ShouldUpdate** | Pointer to **bool** |  | [optional] [default to true]

## Methods

### NewCategoryIndexRequest

`func NewCategoryIndexRequest(id int32, ) *CategoryIndexRequest`

NewCategoryIndexRequest instantiates a new CategoryIndexRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCategoryIndexRequestWithDefaults

`func NewCategoryIndexRequestWithDefaults() *CategoryIndexRequest`

NewCategoryIndexRequestWithDefaults instantiates a new CategoryIndexRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CategoryIndexRequest) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CategoryIndexRequest) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CategoryIndexRequest) SetId(v int32)`

SetId sets Id field to given value.


### GetShouldUpdate

`func (o *CategoryIndexRequest) GetShouldUpdate() bool`

GetShouldUpdate returns the ShouldUpdate field if non-nil, zero value otherwise.

### GetShouldUpdateOk

`func (o *CategoryIndexRequest) GetShouldUpdateOk() (*bool, bool)`

GetShouldUpdateOk returns a tuple with the ShouldUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShouldUpdate

`func (o *CategoryIndexRequest) SetShouldUpdate(v bool)`

SetShouldUpdate sets ShouldUpdate field to given value.

### HasShouldUpdate

`func (o *CategoryIndexRequest) HasShouldUpdate() bool`

HasShouldUpdate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


