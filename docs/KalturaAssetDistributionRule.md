# KalturaAssetDistributionRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AssetDistributionConditions** | Pointer to [**[]KalturaAssetDistributionCondition**](KalturaAssetDistributionCondition.md) |  | [optional] 
**ValidationError** | Pointer to **string** | The validation error description that will be set on the \&quot;data\&quot; property on KalturaDistributionValidationErrorMissingAsset if rule was not fulfilled | [optional] 

## Methods

### NewKalturaAssetDistributionRule

`func NewKalturaAssetDistributionRule() *KalturaAssetDistributionRule`

NewKalturaAssetDistributionRule instantiates a new KalturaAssetDistributionRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKalturaAssetDistributionRuleWithDefaults

`func NewKalturaAssetDistributionRuleWithDefaults() *KalturaAssetDistributionRule`

NewKalturaAssetDistributionRuleWithDefaults instantiates a new KalturaAssetDistributionRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssetDistributionConditions

`func (o *KalturaAssetDistributionRule) GetAssetDistributionConditions() []KalturaAssetDistributionCondition`

GetAssetDistributionConditions returns the AssetDistributionConditions field if non-nil, zero value otherwise.

### GetAssetDistributionConditionsOk

`func (o *KalturaAssetDistributionRule) GetAssetDistributionConditionsOk() (*[]KalturaAssetDistributionCondition, bool)`

GetAssetDistributionConditionsOk returns a tuple with the AssetDistributionConditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetDistributionConditions

`func (o *KalturaAssetDistributionRule) SetAssetDistributionConditions(v []KalturaAssetDistributionCondition)`

SetAssetDistributionConditions sets AssetDistributionConditions field to given value.

### HasAssetDistributionConditions

`func (o *KalturaAssetDistributionRule) HasAssetDistributionConditions() bool`

HasAssetDistributionConditions returns a boolean if a field has been set.

### GetValidationError

`func (o *KalturaAssetDistributionRule) GetValidationError() string`

GetValidationError returns the ValidationError field if non-nil, zero value otherwise.

### GetValidationErrorOk

`func (o *KalturaAssetDistributionRule) GetValidationErrorOk() (*string, bool)`

GetValidationErrorOk returns a tuple with the ValidationError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationError

`func (o *KalturaAssetDistributionRule) SetValidationError(v string)`

SetValidationError sets ValidationError field to given value.

### HasValidationError

`func (o *KalturaAssetDistributionRule) HasValidationError() bool`

HasValidationError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


