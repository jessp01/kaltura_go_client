# LiveStreamFinishConfRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EntryId** | **string** |  | 
**ServerNodeId** | Pointer to **int32** |  | [optional] 

## Methods

### NewLiveStreamFinishConfRequest

`func NewLiveStreamFinishConfRequest(entryId string, ) *LiveStreamFinishConfRequest`

NewLiveStreamFinishConfRequest instantiates a new LiveStreamFinishConfRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLiveStreamFinishConfRequestWithDefaults

`func NewLiveStreamFinishConfRequestWithDefaults() *LiveStreamFinishConfRequest`

NewLiveStreamFinishConfRequestWithDefaults instantiates a new LiveStreamFinishConfRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntryId

`func (o *LiveStreamFinishConfRequest) GetEntryId() string`

GetEntryId returns the EntryId field if non-nil, zero value otherwise.

### GetEntryIdOk

`func (o *LiveStreamFinishConfRequest) GetEntryIdOk() (*string, bool)`

GetEntryIdOk returns a tuple with the EntryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntryId

`func (o *LiveStreamFinishConfRequest) SetEntryId(v string)`

SetEntryId sets EntryId field to given value.


### GetServerNodeId

`func (o *LiveStreamFinishConfRequest) GetServerNodeId() int32`

GetServerNodeId returns the ServerNodeId field if non-nil, zero value otherwise.

### GetServerNodeIdOk

`func (o *LiveStreamFinishConfRequest) GetServerNodeIdOk() (*int32, bool)`

GetServerNodeIdOk returns a tuple with the ServerNodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerNodeId

`func (o *LiveStreamFinishConfRequest) SetServerNodeId(v int32)`

SetServerNodeId sets ServerNodeId field to given value.

### HasServerNodeId

`func (o *LiveStreamFinishConfRequest) HasServerNodeId() bool`

HasServerNodeId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


