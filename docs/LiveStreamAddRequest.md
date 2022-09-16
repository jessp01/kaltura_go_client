# LiveStreamAddRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LiveStreamEntry** | [**KalturaLiveStreamEntry**](KalturaLiveStreamEntry.md) |  | 
**SourceType** | Pointer to **string** |  | [optional] 

## Methods

### NewLiveStreamAddRequest

`func NewLiveStreamAddRequest(liveStreamEntry KalturaLiveStreamEntry, ) *LiveStreamAddRequest`

NewLiveStreamAddRequest instantiates a new LiveStreamAddRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLiveStreamAddRequestWithDefaults

`func NewLiveStreamAddRequestWithDefaults() *LiveStreamAddRequest`

NewLiveStreamAddRequestWithDefaults instantiates a new LiveStreamAddRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLiveStreamEntry

`func (o *LiveStreamAddRequest) GetLiveStreamEntry() KalturaLiveStreamEntry`

GetLiveStreamEntry returns the LiveStreamEntry field if non-nil, zero value otherwise.

### GetLiveStreamEntryOk

`func (o *LiveStreamAddRequest) GetLiveStreamEntryOk() (*KalturaLiveStreamEntry, bool)`

GetLiveStreamEntryOk returns a tuple with the LiveStreamEntry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLiveStreamEntry

`func (o *LiveStreamAddRequest) SetLiveStreamEntry(v KalturaLiveStreamEntry)`

SetLiveStreamEntry sets LiveStreamEntry field to given value.


### GetSourceType

`func (o *LiveStreamAddRequest) GetSourceType() string`

GetSourceType returns the SourceType field if non-nil, zero value otherwise.

### GetSourceTypeOk

`func (o *LiveStreamAddRequest) GetSourceTypeOk() (*string, bool)`

GetSourceTypeOk returns a tuple with the SourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceType

`func (o *LiveStreamAddRequest) SetSourceType(v string)`

SetSourceType sets SourceType field to given value.

### HasSourceType

`func (o *LiveStreamAddRequest) HasSourceType() bool`

HasSourceType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


