# MediaAddFromRecordedWebcamRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MediaEntry** | [**KalturaMediaEntry**](KalturaMediaEntry.md) |  | 
**WebcamTokenId** | **string** |  | 

## Methods

### NewMediaAddFromRecordedWebcamRequest

`func NewMediaAddFromRecordedWebcamRequest(mediaEntry KalturaMediaEntry, webcamTokenId string, ) *MediaAddFromRecordedWebcamRequest`

NewMediaAddFromRecordedWebcamRequest instantiates a new MediaAddFromRecordedWebcamRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMediaAddFromRecordedWebcamRequestWithDefaults

`func NewMediaAddFromRecordedWebcamRequestWithDefaults() *MediaAddFromRecordedWebcamRequest`

NewMediaAddFromRecordedWebcamRequestWithDefaults instantiates a new MediaAddFromRecordedWebcamRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMediaEntry

`func (o *MediaAddFromRecordedWebcamRequest) GetMediaEntry() KalturaMediaEntry`

GetMediaEntry returns the MediaEntry field if non-nil, zero value otherwise.

### GetMediaEntryOk

`func (o *MediaAddFromRecordedWebcamRequest) GetMediaEntryOk() (*KalturaMediaEntry, bool)`

GetMediaEntryOk returns a tuple with the MediaEntry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaEntry

`func (o *MediaAddFromRecordedWebcamRequest) SetMediaEntry(v KalturaMediaEntry)`

SetMediaEntry sets MediaEntry field to given value.


### GetWebcamTokenId

`func (o *MediaAddFromRecordedWebcamRequest) GetWebcamTokenId() string`

GetWebcamTokenId returns the WebcamTokenId field if non-nil, zero value otherwise.

### GetWebcamTokenIdOk

`func (o *MediaAddFromRecordedWebcamRequest) GetWebcamTokenIdOk() (*string, bool)`

GetWebcamTokenIdOk returns a tuple with the WebcamTokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebcamTokenId

`func (o *MediaAddFromRecordedWebcamRequest) SetWebcamTokenId(v string)`

SetWebcamTokenId sets WebcamTokenId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


