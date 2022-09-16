# MediaGetRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EntryId** | **string** |  | 
**Version** | Pointer to **int32** |  | [optional] 

## Methods

### NewMediaGetRequest

`func NewMediaGetRequest(entryId string, ) *MediaGetRequest`

NewMediaGetRequest instantiates a new MediaGetRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMediaGetRequestWithDefaults

`func NewMediaGetRequestWithDefaults() *MediaGetRequest`

NewMediaGetRequestWithDefaults instantiates a new MediaGetRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntryId

`func (o *MediaGetRequest) GetEntryId() string`

GetEntryId returns the EntryId field if non-nil, zero value otherwise.

### GetEntryIdOk

`func (o *MediaGetRequest) GetEntryIdOk() (*string, bool)`

GetEntryIdOk returns a tuple with the EntryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntryId

`func (o *MediaGetRequest) SetEntryId(v string)`

SetEntryId sets EntryId field to given value.


### GetVersion

`func (o *MediaGetRequest) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *MediaGetRequest) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *MediaGetRequest) SetVersion(v int32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *MediaGetRequest) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


