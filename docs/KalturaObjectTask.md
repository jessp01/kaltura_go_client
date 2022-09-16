# KalturaObjectTask

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ObjectType** | Pointer to **string** |  | [optional] 
**StopProcessingOnError** | Pointer to **bool** |  | [optional] 
**Type** | Pointer to **string** | &#x60;readOnly&#x60;  Enum Type: &#x60;KalturaObjectTaskType&#x60; | [optional] [readonly] 

## Methods

### NewKalturaObjectTask

`func NewKalturaObjectTask() *KalturaObjectTask`

NewKalturaObjectTask instantiates a new KalturaObjectTask object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKalturaObjectTaskWithDefaults

`func NewKalturaObjectTaskWithDefaults() *KalturaObjectTask`

NewKalturaObjectTaskWithDefaults instantiates a new KalturaObjectTask object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObjectType

`func (o *KalturaObjectTask) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *KalturaObjectTask) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *KalturaObjectTask) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.

### HasObjectType

`func (o *KalturaObjectTask) HasObjectType() bool`

HasObjectType returns a boolean if a field has been set.

### GetStopProcessingOnError

`func (o *KalturaObjectTask) GetStopProcessingOnError() bool`

GetStopProcessingOnError returns the StopProcessingOnError field if non-nil, zero value otherwise.

### GetStopProcessingOnErrorOk

`func (o *KalturaObjectTask) GetStopProcessingOnErrorOk() (*bool, bool)`

GetStopProcessingOnErrorOk returns a tuple with the StopProcessingOnError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStopProcessingOnError

`func (o *KalturaObjectTask) SetStopProcessingOnError(v bool)`

SetStopProcessingOnError sets StopProcessingOnError field to given value.

### HasStopProcessingOnError

`func (o *KalturaObjectTask) HasStopProcessingOnError() bool`

HasStopProcessingOnError returns a boolean if a field has been set.

### GetType

`func (o *KalturaObjectTask) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *KalturaObjectTask) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *KalturaObjectTask) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *KalturaObjectTask) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


