# MediaUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EntryId** | **string** |  | 
**MediaEntry** | [**KalturaMediaEntry**](KalturaMediaEntry.md) |  | 

## Methods

### NewMediaUpdateRequest

`func NewMediaUpdateRequest(entryId string, mediaEntry KalturaMediaEntry, ) *MediaUpdateRequest`

NewMediaUpdateRequest instantiates a new MediaUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMediaUpdateRequestWithDefaults

`func NewMediaUpdateRequestWithDefaults() *MediaUpdateRequest`

NewMediaUpdateRequestWithDefaults instantiates a new MediaUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntryId

`func (o *MediaUpdateRequest) GetEntryId() string`

GetEntryId returns the EntryId field if non-nil, zero value otherwise.

### GetEntryIdOk

`func (o *MediaUpdateRequest) GetEntryIdOk() (*string, bool)`

GetEntryIdOk returns a tuple with the EntryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntryId

`func (o *MediaUpdateRequest) SetEntryId(v string)`

SetEntryId sets EntryId field to given value.


### GetMediaEntry

`func (o *MediaUpdateRequest) GetMediaEntry() KalturaMediaEntry`

GetMediaEntry returns the MediaEntry field if non-nil, zero value otherwise.

### GetMediaEntryOk

`func (o *MediaUpdateRequest) GetMediaEntryOk() (*KalturaMediaEntry, bool)`

GetMediaEntryOk returns a tuple with the MediaEntry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaEntry

`func (o *MediaUpdateRequest) SetMediaEntry(v KalturaMediaEntry)`

SetMediaEntry sets MediaEntry field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


