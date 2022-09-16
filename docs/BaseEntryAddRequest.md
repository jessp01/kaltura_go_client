# BaseEntryAddRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Entry** | [**KalturaBaseEntry**](KalturaBaseEntry.md) |  | 
**Type** | Pointer to **string** |  | [optional] 

## Methods

### NewBaseEntryAddRequest

`func NewBaseEntryAddRequest(entry KalturaBaseEntry, ) *BaseEntryAddRequest`

NewBaseEntryAddRequest instantiates a new BaseEntryAddRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBaseEntryAddRequestWithDefaults

`func NewBaseEntryAddRequestWithDefaults() *BaseEntryAddRequest`

NewBaseEntryAddRequestWithDefaults instantiates a new BaseEntryAddRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntry

`func (o *BaseEntryAddRequest) GetEntry() KalturaBaseEntry`

GetEntry returns the Entry field if non-nil, zero value otherwise.

### GetEntryOk

`func (o *BaseEntryAddRequest) GetEntryOk() (*KalturaBaseEntry, bool)`

GetEntryOk returns a tuple with the Entry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntry

`func (o *BaseEntryAddRequest) SetEntry(v KalturaBaseEntry)`

SetEntry sets Entry field to given value.


### GetType

`func (o *BaseEntryAddRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BaseEntryAddRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BaseEntryAddRequest) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *BaseEntryAddRequest) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


