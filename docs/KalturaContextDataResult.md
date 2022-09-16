# KalturaContextDataResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Actions** | Pointer to [**[]KalturaRuleAction**](KalturaRuleAction.md) |  | [optional] 
**Messages** | Pointer to [**[]KalturaString**](KalturaString.md) |  | [optional] 
**ObjectType** | Pointer to **string** |  | [optional] 

## Methods

### NewKalturaContextDataResult

`func NewKalturaContextDataResult() *KalturaContextDataResult`

NewKalturaContextDataResult instantiates a new KalturaContextDataResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKalturaContextDataResultWithDefaults

`func NewKalturaContextDataResultWithDefaults() *KalturaContextDataResult`

NewKalturaContextDataResultWithDefaults instantiates a new KalturaContextDataResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActions

`func (o *KalturaContextDataResult) GetActions() []KalturaRuleAction`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *KalturaContextDataResult) GetActionsOk() (*[]KalturaRuleAction, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *KalturaContextDataResult) SetActions(v []KalturaRuleAction)`

SetActions sets Actions field to given value.

### HasActions

`func (o *KalturaContextDataResult) HasActions() bool`

HasActions returns a boolean if a field has been set.

### GetMessages

`func (o *KalturaContextDataResult) GetMessages() []KalturaString`

GetMessages returns the Messages field if non-nil, zero value otherwise.

### GetMessagesOk

`func (o *KalturaContextDataResult) GetMessagesOk() (*[]KalturaString, bool)`

GetMessagesOk returns a tuple with the Messages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessages

`func (o *KalturaContextDataResult) SetMessages(v []KalturaString)`

SetMessages sets Messages field to given value.

### HasMessages

`func (o *KalturaContextDataResult) HasMessages() bool`

HasMessages returns a boolean if a field has been set.

### GetObjectType

`func (o *KalturaContextDataResult) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *KalturaContextDataResult) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *KalturaContextDataResult) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.

### HasObjectType

`func (o *KalturaContextDataResult) HasObjectType() bool`

HasObjectType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


