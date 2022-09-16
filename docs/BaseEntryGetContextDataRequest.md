# BaseEntryGetContextDataRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContextDataParams** | [**KalturaEntryContextDataParams**](KalturaEntryContextDataParams.md) |  | 
**EntryId** | **string** |  | 

## Methods

### NewBaseEntryGetContextDataRequest

`func NewBaseEntryGetContextDataRequest(contextDataParams KalturaEntryContextDataParams, entryId string, ) *BaseEntryGetContextDataRequest`

NewBaseEntryGetContextDataRequest instantiates a new BaseEntryGetContextDataRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBaseEntryGetContextDataRequestWithDefaults

`func NewBaseEntryGetContextDataRequestWithDefaults() *BaseEntryGetContextDataRequest`

NewBaseEntryGetContextDataRequestWithDefaults instantiates a new BaseEntryGetContextDataRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContextDataParams

`func (o *BaseEntryGetContextDataRequest) GetContextDataParams() KalturaEntryContextDataParams`

GetContextDataParams returns the ContextDataParams field if non-nil, zero value otherwise.

### GetContextDataParamsOk

`func (o *BaseEntryGetContextDataRequest) GetContextDataParamsOk() (*KalturaEntryContextDataParams, bool)`

GetContextDataParamsOk returns a tuple with the ContextDataParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContextDataParams

`func (o *BaseEntryGetContextDataRequest) SetContextDataParams(v KalturaEntryContextDataParams)`

SetContextDataParams sets ContextDataParams field to given value.


### GetEntryId

`func (o *BaseEntryGetContextDataRequest) GetEntryId() string`

GetEntryId returns the EntryId field if non-nil, zero value otherwise.

### GetEntryIdOk

`func (o *BaseEntryGetContextDataRequest) GetEntryIdOk() (*string, bool)`

GetEntryIdOk returns a tuple with the EntryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntryId

`func (o *BaseEntryGetContextDataRequest) SetEntryId(v string)`

SetEntryId sets EntryId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


