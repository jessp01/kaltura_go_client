# ExternalMediaUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Entry** | [**KalturaExternalMediaEntry**](KalturaExternalMediaEntry.md) |  | 
**Id** | **string** |  | 

## Methods

### NewExternalMediaUpdateRequest

`func NewExternalMediaUpdateRequest(entry KalturaExternalMediaEntry, id string, ) *ExternalMediaUpdateRequest`

NewExternalMediaUpdateRequest instantiates a new ExternalMediaUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExternalMediaUpdateRequestWithDefaults

`func NewExternalMediaUpdateRequestWithDefaults() *ExternalMediaUpdateRequest`

NewExternalMediaUpdateRequestWithDefaults instantiates a new ExternalMediaUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntry

`func (o *ExternalMediaUpdateRequest) GetEntry() KalturaExternalMediaEntry`

GetEntry returns the Entry field if non-nil, zero value otherwise.

### GetEntryOk

`func (o *ExternalMediaUpdateRequest) GetEntryOk() (*KalturaExternalMediaEntry, bool)`

GetEntryOk returns a tuple with the Entry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntry

`func (o *ExternalMediaUpdateRequest) SetEntry(v KalturaExternalMediaEntry)`

SetEntry sets Entry field to given value.


### GetId

`func (o *ExternalMediaUpdateRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ExternalMediaUpdateRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ExternalMediaUpdateRequest) SetId(v string)`

SetId sets Id field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


