# EntryVendorTaskAbortRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AbortReason** | Pointer to **string** |  | [optional] 
**Id** | **int32** |  | 

## Methods

### NewEntryVendorTaskAbortRequest

`func NewEntryVendorTaskAbortRequest(id int32, ) *EntryVendorTaskAbortRequest`

NewEntryVendorTaskAbortRequest instantiates a new EntryVendorTaskAbortRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEntryVendorTaskAbortRequestWithDefaults

`func NewEntryVendorTaskAbortRequestWithDefaults() *EntryVendorTaskAbortRequest`

NewEntryVendorTaskAbortRequestWithDefaults instantiates a new EntryVendorTaskAbortRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAbortReason

`func (o *EntryVendorTaskAbortRequest) GetAbortReason() string`

GetAbortReason returns the AbortReason field if non-nil, zero value otherwise.

### GetAbortReasonOk

`func (o *EntryVendorTaskAbortRequest) GetAbortReasonOk() (*string, bool)`

GetAbortReasonOk returns a tuple with the AbortReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbortReason

`func (o *EntryVendorTaskAbortRequest) SetAbortReason(v string)`

SetAbortReason sets AbortReason field to given value.

### HasAbortReason

`func (o *EntryVendorTaskAbortRequest) HasAbortReason() bool`

HasAbortReason returns a boolean if a field has been set.

### GetId

`func (o *EntryVendorTaskAbortRequest) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EntryVendorTaskAbortRequest) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EntryVendorTaskAbortRequest) SetId(v int32)`

SetId sets Id field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


