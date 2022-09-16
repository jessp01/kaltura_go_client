# PlayReadyDrmGetEntryContentKeyRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreateIfMissing** | Pointer to **bool** |  | [optional] [default to false]
**EntryId** | **string** |  | 

## Methods

### NewPlayReadyDrmGetEntryContentKeyRequest

`func NewPlayReadyDrmGetEntryContentKeyRequest(entryId string, ) *PlayReadyDrmGetEntryContentKeyRequest`

NewPlayReadyDrmGetEntryContentKeyRequest instantiates a new PlayReadyDrmGetEntryContentKeyRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlayReadyDrmGetEntryContentKeyRequestWithDefaults

`func NewPlayReadyDrmGetEntryContentKeyRequestWithDefaults() *PlayReadyDrmGetEntryContentKeyRequest`

NewPlayReadyDrmGetEntryContentKeyRequestWithDefaults instantiates a new PlayReadyDrmGetEntryContentKeyRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreateIfMissing

`func (o *PlayReadyDrmGetEntryContentKeyRequest) GetCreateIfMissing() bool`

GetCreateIfMissing returns the CreateIfMissing field if non-nil, zero value otherwise.

### GetCreateIfMissingOk

`func (o *PlayReadyDrmGetEntryContentKeyRequest) GetCreateIfMissingOk() (*bool, bool)`

GetCreateIfMissingOk returns a tuple with the CreateIfMissing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateIfMissing

`func (o *PlayReadyDrmGetEntryContentKeyRequest) SetCreateIfMissing(v bool)`

SetCreateIfMissing sets CreateIfMissing field to given value.

### HasCreateIfMissing

`func (o *PlayReadyDrmGetEntryContentKeyRequest) HasCreateIfMissing() bool`

HasCreateIfMissing returns a boolean if a field has been set.

### GetEntryId

`func (o *PlayReadyDrmGetEntryContentKeyRequest) GetEntryId() string`

GetEntryId returns the EntryId field if non-nil, zero value otherwise.

### GetEntryIdOk

`func (o *PlayReadyDrmGetEntryContentKeyRequest) GetEntryIdOk() (*string, bool)`

GetEntryIdOk returns a tuple with the EntryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntryId

`func (o *PlayReadyDrmGetEntryContentKeyRequest) SetEntryId(v string)`

SetEntryId sets EntryId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


