# InteractivityUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EntryId** | **string** |  | 
**KalturaInteractivity** | [**KalturaInteractivity**](KalturaInteractivity.md) |  | 
**Version** | **int32** |  | 

## Methods

### NewInteractivityUpdateRequest

`func NewInteractivityUpdateRequest(entryId string, kalturaInteractivity KalturaInteractivity, version int32, ) *InteractivityUpdateRequest`

NewInteractivityUpdateRequest instantiates a new InteractivityUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInteractivityUpdateRequestWithDefaults

`func NewInteractivityUpdateRequestWithDefaults() *InteractivityUpdateRequest`

NewInteractivityUpdateRequestWithDefaults instantiates a new InteractivityUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntryId

`func (o *InteractivityUpdateRequest) GetEntryId() string`

GetEntryId returns the EntryId field if non-nil, zero value otherwise.

### GetEntryIdOk

`func (o *InteractivityUpdateRequest) GetEntryIdOk() (*string, bool)`

GetEntryIdOk returns a tuple with the EntryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntryId

`func (o *InteractivityUpdateRequest) SetEntryId(v string)`

SetEntryId sets EntryId field to given value.


### GetKalturaInteractivity

`func (o *InteractivityUpdateRequest) GetKalturaInteractivity() KalturaInteractivity`

GetKalturaInteractivity returns the KalturaInteractivity field if non-nil, zero value otherwise.

### GetKalturaInteractivityOk

`func (o *InteractivityUpdateRequest) GetKalturaInteractivityOk() (*KalturaInteractivity, bool)`

GetKalturaInteractivityOk returns a tuple with the KalturaInteractivity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKalturaInteractivity

`func (o *InteractivityUpdateRequest) SetKalturaInteractivity(v KalturaInteractivity)`

SetKalturaInteractivity sets KalturaInteractivity field to given value.


### GetVersion

`func (o *InteractivityUpdateRequest) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *InteractivityUpdateRequest) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *InteractivityUpdateRequest) SetVersion(v int32)`

SetVersion sets Version field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


