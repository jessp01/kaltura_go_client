# MediaAddFromUrlRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MediaEntry** | [**KalturaMediaEntry**](KalturaMediaEntry.md) |  | 
**Url** | **string** |  | 

## Methods

### NewMediaAddFromUrlRequest

`func NewMediaAddFromUrlRequest(mediaEntry KalturaMediaEntry, url string, ) *MediaAddFromUrlRequest`

NewMediaAddFromUrlRequest instantiates a new MediaAddFromUrlRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMediaAddFromUrlRequestWithDefaults

`func NewMediaAddFromUrlRequestWithDefaults() *MediaAddFromUrlRequest`

NewMediaAddFromUrlRequestWithDefaults instantiates a new MediaAddFromUrlRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMediaEntry

`func (o *MediaAddFromUrlRequest) GetMediaEntry() KalturaMediaEntry`

GetMediaEntry returns the MediaEntry field if non-nil, zero value otherwise.

### GetMediaEntryOk

`func (o *MediaAddFromUrlRequest) GetMediaEntryOk() (*KalturaMediaEntry, bool)`

GetMediaEntryOk returns a tuple with the MediaEntry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaEntry

`func (o *MediaAddFromUrlRequest) SetMediaEntry(v KalturaMediaEntry)`

SetMediaEntry sets MediaEntry field to given value.


### GetUrl

`func (o *MediaAddFromUrlRequest) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *MediaAddFromUrlRequest) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *MediaAddFromUrlRequest) SetUrl(v string)`

SetUrl sets Url field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


