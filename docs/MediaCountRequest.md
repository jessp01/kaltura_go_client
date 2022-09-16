# MediaCountRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Filter** | Pointer to [**KalturaMediaEntryFilter**](KalturaMediaEntryFilter.md) |  | [optional] 

## Methods

### NewMediaCountRequest

`func NewMediaCountRequest() *MediaCountRequest`

NewMediaCountRequest instantiates a new MediaCountRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMediaCountRequestWithDefaults

`func NewMediaCountRequestWithDefaults() *MediaCountRequest`

NewMediaCountRequestWithDefaults instantiates a new MediaCountRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilter

`func (o *MediaCountRequest) GetFilter() KalturaMediaEntryFilter`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *MediaCountRequest) GetFilterOk() (*KalturaMediaEntryFilter, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *MediaCountRequest) SetFilter(v KalturaMediaEntryFilter)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *MediaCountRequest) HasFilter() bool`

HasFilter returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


