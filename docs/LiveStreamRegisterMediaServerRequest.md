# LiveStreamRegisterMediaServerRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApplicationName** | Pointer to **string** |  | [optional] 
**EntryId** | **string** |  | 
**Hostname** | **string** |  | 
**LiveEntryStatus** | Pointer to **int32** |  | [optional] 
**MediaServerIndex** | **string** |  | 
**ShouldCreateRecordedEntry** | Pointer to **bool** |  | [optional] [default to true]

## Methods

### NewLiveStreamRegisterMediaServerRequest

`func NewLiveStreamRegisterMediaServerRequest(entryId string, hostname string, mediaServerIndex string, ) *LiveStreamRegisterMediaServerRequest`

NewLiveStreamRegisterMediaServerRequest instantiates a new LiveStreamRegisterMediaServerRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLiveStreamRegisterMediaServerRequestWithDefaults

`func NewLiveStreamRegisterMediaServerRequestWithDefaults() *LiveStreamRegisterMediaServerRequest`

NewLiveStreamRegisterMediaServerRequestWithDefaults instantiates a new LiveStreamRegisterMediaServerRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplicationName

`func (o *LiveStreamRegisterMediaServerRequest) GetApplicationName() string`

GetApplicationName returns the ApplicationName field if non-nil, zero value otherwise.

### GetApplicationNameOk

`func (o *LiveStreamRegisterMediaServerRequest) GetApplicationNameOk() (*string, bool)`

GetApplicationNameOk returns a tuple with the ApplicationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationName

`func (o *LiveStreamRegisterMediaServerRequest) SetApplicationName(v string)`

SetApplicationName sets ApplicationName field to given value.

### HasApplicationName

`func (o *LiveStreamRegisterMediaServerRequest) HasApplicationName() bool`

HasApplicationName returns a boolean if a field has been set.

### GetEntryId

`func (o *LiveStreamRegisterMediaServerRequest) GetEntryId() string`

GetEntryId returns the EntryId field if non-nil, zero value otherwise.

### GetEntryIdOk

`func (o *LiveStreamRegisterMediaServerRequest) GetEntryIdOk() (*string, bool)`

GetEntryIdOk returns a tuple with the EntryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntryId

`func (o *LiveStreamRegisterMediaServerRequest) SetEntryId(v string)`

SetEntryId sets EntryId field to given value.


### GetHostname

`func (o *LiveStreamRegisterMediaServerRequest) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *LiveStreamRegisterMediaServerRequest) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *LiveStreamRegisterMediaServerRequest) SetHostname(v string)`

SetHostname sets Hostname field to given value.


### GetLiveEntryStatus

`func (o *LiveStreamRegisterMediaServerRequest) GetLiveEntryStatus() int32`

GetLiveEntryStatus returns the LiveEntryStatus field if non-nil, zero value otherwise.

### GetLiveEntryStatusOk

`func (o *LiveStreamRegisterMediaServerRequest) GetLiveEntryStatusOk() (*int32, bool)`

GetLiveEntryStatusOk returns a tuple with the LiveEntryStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLiveEntryStatus

`func (o *LiveStreamRegisterMediaServerRequest) SetLiveEntryStatus(v int32)`

SetLiveEntryStatus sets LiveEntryStatus field to given value.

### HasLiveEntryStatus

`func (o *LiveStreamRegisterMediaServerRequest) HasLiveEntryStatus() bool`

HasLiveEntryStatus returns a boolean if a field has been set.

### GetMediaServerIndex

`func (o *LiveStreamRegisterMediaServerRequest) GetMediaServerIndex() string`

GetMediaServerIndex returns the MediaServerIndex field if non-nil, zero value otherwise.

### GetMediaServerIndexOk

`func (o *LiveStreamRegisterMediaServerRequest) GetMediaServerIndexOk() (*string, bool)`

GetMediaServerIndexOk returns a tuple with the MediaServerIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaServerIndex

`func (o *LiveStreamRegisterMediaServerRequest) SetMediaServerIndex(v string)`

SetMediaServerIndex sets MediaServerIndex field to given value.


### GetShouldCreateRecordedEntry

`func (o *LiveStreamRegisterMediaServerRequest) GetShouldCreateRecordedEntry() bool`

GetShouldCreateRecordedEntry returns the ShouldCreateRecordedEntry field if non-nil, zero value otherwise.

### GetShouldCreateRecordedEntryOk

`func (o *LiveStreamRegisterMediaServerRequest) GetShouldCreateRecordedEntryOk() (*bool, bool)`

GetShouldCreateRecordedEntryOk returns a tuple with the ShouldCreateRecordedEntry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShouldCreateRecordedEntry

`func (o *LiveStreamRegisterMediaServerRequest) SetShouldCreateRecordedEntry(v bool)`

SetShouldCreateRecordedEntry sets ShouldCreateRecordedEntry field to given value.

### HasShouldCreateRecordedEntry

`func (o *LiveStreamRegisterMediaServerRequest) HasShouldCreateRecordedEntry() bool`

HasShouldCreateRecordedEntry returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


