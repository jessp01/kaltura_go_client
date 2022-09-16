# BaseEntryCloneRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CloneOptions** | Pointer to [**[]KalturaBaseEntryCloneOptionItem**](KalturaBaseEntryCloneOptionItem.md) |  | [optional] 
**EntryId** | **string** |  | 

## Methods

### NewBaseEntryCloneRequest

`func NewBaseEntryCloneRequest(entryId string, ) *BaseEntryCloneRequest`

NewBaseEntryCloneRequest instantiates a new BaseEntryCloneRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBaseEntryCloneRequestWithDefaults

`func NewBaseEntryCloneRequestWithDefaults() *BaseEntryCloneRequest`

NewBaseEntryCloneRequestWithDefaults instantiates a new BaseEntryCloneRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCloneOptions

`func (o *BaseEntryCloneRequest) GetCloneOptions() []KalturaBaseEntryCloneOptionItem`

GetCloneOptions returns the CloneOptions field if non-nil, zero value otherwise.

### GetCloneOptionsOk

`func (o *BaseEntryCloneRequest) GetCloneOptionsOk() (*[]KalturaBaseEntryCloneOptionItem, bool)`

GetCloneOptionsOk returns a tuple with the CloneOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloneOptions

`func (o *BaseEntryCloneRequest) SetCloneOptions(v []KalturaBaseEntryCloneOptionItem)`

SetCloneOptions sets CloneOptions field to given value.

### HasCloneOptions

`func (o *BaseEntryCloneRequest) HasCloneOptions() bool`

HasCloneOptions returns a boolean if a field has been set.

### GetEntryId

`func (o *BaseEntryCloneRequest) GetEntryId() string`

GetEntryId returns the EntryId field if non-nil, zero value otherwise.

### GetEntryIdOk

`func (o *BaseEntryCloneRequest) GetEntryIdOk() (*string, bool)`

GetEntryIdOk returns a tuple with the EntryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntryId

`func (o *BaseEntryCloneRequest) SetEntryId(v string)`

SetEntryId sets EntryId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


