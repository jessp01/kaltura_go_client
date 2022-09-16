# InteractivityGetRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DataFilter** | Pointer to [**KalturaInteractivityDataFilter**](KalturaInteractivityDataFilter.md) |  | [optional] 
**EntryId** | **string** |  | 

## Methods

### NewInteractivityGetRequest

`func NewInteractivityGetRequest(entryId string, ) *InteractivityGetRequest`

NewInteractivityGetRequest instantiates a new InteractivityGetRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInteractivityGetRequestWithDefaults

`func NewInteractivityGetRequestWithDefaults() *InteractivityGetRequest`

NewInteractivityGetRequestWithDefaults instantiates a new InteractivityGetRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDataFilter

`func (o *InteractivityGetRequest) GetDataFilter() KalturaInteractivityDataFilter`

GetDataFilter returns the DataFilter field if non-nil, zero value otherwise.

### GetDataFilterOk

`func (o *InteractivityGetRequest) GetDataFilterOk() (*KalturaInteractivityDataFilter, bool)`

GetDataFilterOk returns a tuple with the DataFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataFilter

`func (o *InteractivityGetRequest) SetDataFilter(v KalturaInteractivityDataFilter)`

SetDataFilter sets DataFilter field to given value.

### HasDataFilter

`func (o *InteractivityGetRequest) HasDataFilter() bool`

HasDataFilter returns a boolean if a field has been set.

### GetEntryId

`func (o *InteractivityGetRequest) GetEntryId() string`

GetEntryId returns the EntryId field if non-nil, zero value otherwise.

### GetEntryIdOk

`func (o *InteractivityGetRequest) GetEntryIdOk() (*string, bool)`

GetEntryIdOk returns a tuple with the EntryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntryId

`func (o *InteractivityGetRequest) SetEntryId(v string)`

SetEntryId sets EntryId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


