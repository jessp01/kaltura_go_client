# DataServeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EntryId** | **string** |  | 
**ForceProxy** | Pointer to **bool** |  | [optional] [default to false]
**Version** | Pointer to **int32** |  | [optional] 

## Methods

### NewDataServeRequest

`func NewDataServeRequest(entryId string, ) *DataServeRequest`

NewDataServeRequest instantiates a new DataServeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataServeRequestWithDefaults

`func NewDataServeRequestWithDefaults() *DataServeRequest`

NewDataServeRequestWithDefaults instantiates a new DataServeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntryId

`func (o *DataServeRequest) GetEntryId() string`

GetEntryId returns the EntryId field if non-nil, zero value otherwise.

### GetEntryIdOk

`func (o *DataServeRequest) GetEntryIdOk() (*string, bool)`

GetEntryIdOk returns a tuple with the EntryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntryId

`func (o *DataServeRequest) SetEntryId(v string)`

SetEntryId sets EntryId field to given value.


### GetForceProxy

`func (o *DataServeRequest) GetForceProxy() bool`

GetForceProxy returns the ForceProxy field if non-nil, zero value otherwise.

### GetForceProxyOk

`func (o *DataServeRequest) GetForceProxyOk() (*bool, bool)`

GetForceProxyOk returns a tuple with the ForceProxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForceProxy

`func (o *DataServeRequest) SetForceProxy(v bool)`

SetForceProxy sets ForceProxy field to given value.

### HasForceProxy

`func (o *DataServeRequest) HasForceProxy() bool`

HasForceProxy returns a boolean if a field has been set.

### GetVersion

`func (o *DataServeRequest) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *DataServeRequest) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *DataServeRequest) SetVersion(v int32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *DataServeRequest) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


