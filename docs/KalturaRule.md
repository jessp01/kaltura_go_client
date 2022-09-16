# KalturaRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Actions** | Pointer to [**[]KalturaRuleAction**](KalturaRuleAction.md) |  | [optional] 
**Code** | Pointer to **string** | Code to be thrown to the player in case the rule is fulfilled | [optional] 
**Conditions** | Pointer to [**[]KalturaCondition**](KalturaCondition.md) |  | [optional] 
**Contexts** | Pointer to [**[]KalturaContextTypeHolder**](KalturaContextTypeHolder.md) |  | [optional] 
**Description** | Pointer to **string** | Short Rule Description | [optional] 
**ForceAdminValidation** | Pointer to **int32** | Enum Type: &#x60;KalturaNullableBoolean&#x60;  Indicates if we should force ks validation for admin ks users as well | [optional] 
**Message** | Pointer to **string** | Message to be thrown to the player in case the rule is fulfilled | [optional] 
**RuleData** | Pointer to **string** | Rule Custom Data to allow saving rule specific information | [optional] 
**StopProcessing** | Pointer to **bool** | Indicates that this rule is enough and no need to continue checking the rest of the rules | [optional] 

## Methods

### NewKalturaRule

`func NewKalturaRule() *KalturaRule`

NewKalturaRule instantiates a new KalturaRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKalturaRuleWithDefaults

`func NewKalturaRuleWithDefaults() *KalturaRule`

NewKalturaRuleWithDefaults instantiates a new KalturaRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActions

`func (o *KalturaRule) GetActions() []KalturaRuleAction`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *KalturaRule) GetActionsOk() (*[]KalturaRuleAction, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *KalturaRule) SetActions(v []KalturaRuleAction)`

SetActions sets Actions field to given value.

### HasActions

`func (o *KalturaRule) HasActions() bool`

HasActions returns a boolean if a field has been set.

### GetCode

`func (o *KalturaRule) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *KalturaRule) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *KalturaRule) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *KalturaRule) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetConditions

`func (o *KalturaRule) GetConditions() []KalturaCondition`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *KalturaRule) GetConditionsOk() (*[]KalturaCondition, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *KalturaRule) SetConditions(v []KalturaCondition)`

SetConditions sets Conditions field to given value.

### HasConditions

`func (o *KalturaRule) HasConditions() bool`

HasConditions returns a boolean if a field has been set.

### GetContexts

`func (o *KalturaRule) GetContexts() []KalturaContextTypeHolder`

GetContexts returns the Contexts field if non-nil, zero value otherwise.

### GetContextsOk

`func (o *KalturaRule) GetContextsOk() (*[]KalturaContextTypeHolder, bool)`

GetContextsOk returns a tuple with the Contexts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContexts

`func (o *KalturaRule) SetContexts(v []KalturaContextTypeHolder)`

SetContexts sets Contexts field to given value.

### HasContexts

`func (o *KalturaRule) HasContexts() bool`

HasContexts returns a boolean if a field has been set.

### GetDescription

`func (o *KalturaRule) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *KalturaRule) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *KalturaRule) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *KalturaRule) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetForceAdminValidation

`func (o *KalturaRule) GetForceAdminValidation() int32`

GetForceAdminValidation returns the ForceAdminValidation field if non-nil, zero value otherwise.

### GetForceAdminValidationOk

`func (o *KalturaRule) GetForceAdminValidationOk() (*int32, bool)`

GetForceAdminValidationOk returns a tuple with the ForceAdminValidation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForceAdminValidation

`func (o *KalturaRule) SetForceAdminValidation(v int32)`

SetForceAdminValidation sets ForceAdminValidation field to given value.

### HasForceAdminValidation

`func (o *KalturaRule) HasForceAdminValidation() bool`

HasForceAdminValidation returns a boolean if a field has been set.

### GetMessage

`func (o *KalturaRule) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *KalturaRule) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *KalturaRule) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *KalturaRule) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetRuleData

`func (o *KalturaRule) GetRuleData() string`

GetRuleData returns the RuleData field if non-nil, zero value otherwise.

### GetRuleDataOk

`func (o *KalturaRule) GetRuleDataOk() (*string, bool)`

GetRuleDataOk returns a tuple with the RuleData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleData

`func (o *KalturaRule) SetRuleData(v string)`

SetRuleData sets RuleData field to given value.

### HasRuleData

`func (o *KalturaRule) HasRuleData() bool`

HasRuleData returns a boolean if a field has been set.

### GetStopProcessing

`func (o *KalturaRule) GetStopProcessing() bool`

GetStopProcessing returns the StopProcessing field if non-nil, zero value otherwise.

### GetStopProcessingOk

`func (o *KalturaRule) GetStopProcessingOk() (*bool, bool)`

GetStopProcessingOk returns a tuple with the StopProcessing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStopProcessing

`func (o *KalturaRule) SetStopProcessing(v bool)`

SetStopProcessing sets StopProcessing field to given value.

### HasStopProcessing

`func (o *KalturaRule) HasStopProcessing() bool`

HasStopProcessing returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


