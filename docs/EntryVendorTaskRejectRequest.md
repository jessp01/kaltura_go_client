# EntryVendorTaskRejectRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**RejectReason** | Pointer to **string** |  | [optional] 

## Methods

### NewEntryVendorTaskRejectRequest

`func NewEntryVendorTaskRejectRequest(id int32, ) *EntryVendorTaskRejectRequest`

NewEntryVendorTaskRejectRequest instantiates a new EntryVendorTaskRejectRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEntryVendorTaskRejectRequestWithDefaults

`func NewEntryVendorTaskRejectRequestWithDefaults() *EntryVendorTaskRejectRequest`

NewEntryVendorTaskRejectRequestWithDefaults instantiates a new EntryVendorTaskRejectRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *EntryVendorTaskRejectRequest) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EntryVendorTaskRejectRequest) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EntryVendorTaskRejectRequest) SetId(v int32)`

SetId sets Id field to given value.


### GetRejectReason

`func (o *EntryVendorTaskRejectRequest) GetRejectReason() string`

GetRejectReason returns the RejectReason field if non-nil, zero value otherwise.

### GetRejectReasonOk

`func (o *EntryVendorTaskRejectRequest) GetRejectReasonOk() (*string, bool)`

GetRejectReasonOk returns a tuple with the RejectReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRejectReason

`func (o *EntryVendorTaskRejectRequest) SetRejectReason(v string)`

SetRejectReason sets RejectReason field to given value.

### HasRejectReason

`func (o *EntryVendorTaskRejectRequest) HasRejectReason() bool`

HasRejectReason returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


