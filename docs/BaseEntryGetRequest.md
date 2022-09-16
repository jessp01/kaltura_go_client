# BaseEntryGetRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EntryId** | **string** |  | 
**Version** | Pointer to **int32** |  | [optional] 

## Methods

### NewBaseEntryGetRequest

`func NewBaseEntryGetRequest(entryId string, ) *BaseEntryGetRequest`

NewBaseEntryGetRequest instantiates a new BaseEntryGetRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBaseEntryGetRequestWithDefaults

`func NewBaseEntryGetRequestWithDefaults() *BaseEntryGetRequest`

NewBaseEntryGetRequestWithDefaults instantiates a new BaseEntryGetRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntryId

`func (o *BaseEntryGetRequest) GetEntryId() string`

GetEntryId returns the EntryId field if non-nil, zero value otherwise.

### GetEntryIdOk

`func (o *BaseEntryGetRequest) GetEntryIdOk() (*string, bool)`

GetEntryIdOk returns a tuple with the EntryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntryId

`func (o *BaseEntryGetRequest) SetEntryId(v string)`

SetEntryId sets EntryId field to given value.


### GetVersion

`func (o *BaseEntryGetRequest) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *BaseEntryGetRequest) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *BaseEntryGetRequest) SetVersion(v int32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *BaseEntryGetRequest) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


