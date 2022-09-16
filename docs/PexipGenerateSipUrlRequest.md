# PexipGenerateSipUrlRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EntryId** | **string** |  | 
**Regenerate** | Pointer to **bool** |  | [optional] [default to false]
**SourceType** | Pointer to **int32** |  | [optional] 

## Methods

### NewPexipGenerateSipUrlRequest

`func NewPexipGenerateSipUrlRequest(entryId string, ) *PexipGenerateSipUrlRequest`

NewPexipGenerateSipUrlRequest instantiates a new PexipGenerateSipUrlRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPexipGenerateSipUrlRequestWithDefaults

`func NewPexipGenerateSipUrlRequestWithDefaults() *PexipGenerateSipUrlRequest`

NewPexipGenerateSipUrlRequestWithDefaults instantiates a new PexipGenerateSipUrlRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntryId

`func (o *PexipGenerateSipUrlRequest) GetEntryId() string`

GetEntryId returns the EntryId field if non-nil, zero value otherwise.

### GetEntryIdOk

`func (o *PexipGenerateSipUrlRequest) GetEntryIdOk() (*string, bool)`

GetEntryIdOk returns a tuple with the EntryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntryId

`func (o *PexipGenerateSipUrlRequest) SetEntryId(v string)`

SetEntryId sets EntryId field to given value.


### GetRegenerate

`func (o *PexipGenerateSipUrlRequest) GetRegenerate() bool`

GetRegenerate returns the Regenerate field if non-nil, zero value otherwise.

### GetRegenerateOk

`func (o *PexipGenerateSipUrlRequest) GetRegenerateOk() (*bool, bool)`

GetRegenerateOk returns a tuple with the Regenerate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegenerate

`func (o *PexipGenerateSipUrlRequest) SetRegenerate(v bool)`

SetRegenerate sets Regenerate field to given value.

### HasRegenerate

`func (o *PexipGenerateSipUrlRequest) HasRegenerate() bool`

HasRegenerate returns a boolean if a field has been set.

### GetSourceType

`func (o *PexipGenerateSipUrlRequest) GetSourceType() int32`

GetSourceType returns the SourceType field if non-nil, zero value otherwise.

### GetSourceTypeOk

`func (o *PexipGenerateSipUrlRequest) GetSourceTypeOk() (*int32, bool)`

GetSourceTypeOk returns a tuple with the SourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceType

`func (o *PexipGenerateSipUrlRequest) SetSourceType(v int32)`

SetSourceType sets SourceType field to given value.

### HasSourceType

`func (o *PexipGenerateSipUrlRequest) HasSourceType() bool`

HasSourceType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


