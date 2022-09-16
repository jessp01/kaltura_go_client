# BaseEntryIndexRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**ShouldUpdate** | Pointer to **bool** |  | [optional] [default to true]

## Methods

### NewBaseEntryIndexRequest

`func NewBaseEntryIndexRequest(id string, ) *BaseEntryIndexRequest`

NewBaseEntryIndexRequest instantiates a new BaseEntryIndexRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBaseEntryIndexRequestWithDefaults

`func NewBaseEntryIndexRequestWithDefaults() *BaseEntryIndexRequest`

NewBaseEntryIndexRequestWithDefaults instantiates a new BaseEntryIndexRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BaseEntryIndexRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BaseEntryIndexRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BaseEntryIndexRequest) SetId(v string)`

SetId sets Id field to given value.


### GetShouldUpdate

`func (o *BaseEntryIndexRequest) GetShouldUpdate() bool`

GetShouldUpdate returns the ShouldUpdate field if non-nil, zero value otherwise.

### GetShouldUpdateOk

`func (o *BaseEntryIndexRequest) GetShouldUpdateOk() (*bool, bool)`

GetShouldUpdateOk returns a tuple with the ShouldUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShouldUpdate

`func (o *BaseEntryIndexRequest) SetShouldUpdate(v bool)`

SetShouldUpdate sets ShouldUpdate field to given value.

### HasShouldUpdate

`func (o *BaseEntryIndexRequest) HasShouldUpdate() bool`

HasShouldUpdate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


