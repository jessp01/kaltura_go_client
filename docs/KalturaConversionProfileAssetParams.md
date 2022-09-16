# KalturaConversionProfileAssetParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AssetParamsId** | Pointer to **int32** | &#x60;readOnly&#x60;  The id of the asset params | [optional] [readonly] 
**ChunkedEncodeMode** | Pointer to **int32** |  | [optional] 
**ContentAwareness** | Pointer to **float32** |  | [optional] 
**ConversionProfileId** | Pointer to **int32** | &#x60;readOnly&#x60;  The id of the conversion profile | [optional] [readonly] 
**DeletePolicy** | Pointer to **int32** | Enum Type: &#x60;KalturaAssetParamsDeletePolicy&#x60;  Specifies how to treat the flavor after conversion is finished | [optional] 
**ForceNoneComplied** | Pointer to **int32** | Enum Type: &#x60;KalturaNullableBoolean&#x60;  Starts conversion even if the decision layer reduced the configuration to comply with the source | [optional] 
**IsEncrypted** | Pointer to **int32** | Enum Type: &#x60;KalturaNullableBoolean&#x60; | [optional] 
**Origin** | Pointer to **int32** | Enum Type: &#x60;KalturaAssetParamsOrigin&#x60;  The ingestion origin of the asset params | [optional] 
**OverloadParams** | Pointer to **string** | JSON string containing an array of flavotParams field-value pairs. | [optional] 
**ReadyBehavior** | Pointer to **int32** | Enum Type: &#x60;KalturaFlavorReadyBehaviorType&#x60;  The ingestion origin of the asset params | [optional] 
**SystemName** | Pointer to **string** | Asset params system name | [optional] 
**Tags** | Pointer to **string** |  | [optional] 
**TwoPass** | Pointer to **int32** | Enum Type: &#x60;KalturaNullableBoolean&#x60; | [optional] 

## Methods

### NewKalturaConversionProfileAssetParams

`func NewKalturaConversionProfileAssetParams() *KalturaConversionProfileAssetParams`

NewKalturaConversionProfileAssetParams instantiates a new KalturaConversionProfileAssetParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKalturaConversionProfileAssetParamsWithDefaults

`func NewKalturaConversionProfileAssetParamsWithDefaults() *KalturaConversionProfileAssetParams`

NewKalturaConversionProfileAssetParamsWithDefaults instantiates a new KalturaConversionProfileAssetParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssetParamsId

`func (o *KalturaConversionProfileAssetParams) GetAssetParamsId() int32`

GetAssetParamsId returns the AssetParamsId field if non-nil, zero value otherwise.

### GetAssetParamsIdOk

`func (o *KalturaConversionProfileAssetParams) GetAssetParamsIdOk() (*int32, bool)`

GetAssetParamsIdOk returns a tuple with the AssetParamsId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetParamsId

`func (o *KalturaConversionProfileAssetParams) SetAssetParamsId(v int32)`

SetAssetParamsId sets AssetParamsId field to given value.

### HasAssetParamsId

`func (o *KalturaConversionProfileAssetParams) HasAssetParamsId() bool`

HasAssetParamsId returns a boolean if a field has been set.

### GetChunkedEncodeMode

`func (o *KalturaConversionProfileAssetParams) GetChunkedEncodeMode() int32`

GetChunkedEncodeMode returns the ChunkedEncodeMode field if non-nil, zero value otherwise.

### GetChunkedEncodeModeOk

`func (o *KalturaConversionProfileAssetParams) GetChunkedEncodeModeOk() (*int32, bool)`

GetChunkedEncodeModeOk returns a tuple with the ChunkedEncodeMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChunkedEncodeMode

`func (o *KalturaConversionProfileAssetParams) SetChunkedEncodeMode(v int32)`

SetChunkedEncodeMode sets ChunkedEncodeMode field to given value.

### HasChunkedEncodeMode

`func (o *KalturaConversionProfileAssetParams) HasChunkedEncodeMode() bool`

HasChunkedEncodeMode returns a boolean if a field has been set.

### GetContentAwareness

`func (o *KalturaConversionProfileAssetParams) GetContentAwareness() float32`

GetContentAwareness returns the ContentAwareness field if non-nil, zero value otherwise.

### GetContentAwarenessOk

`func (o *KalturaConversionProfileAssetParams) GetContentAwarenessOk() (*float32, bool)`

GetContentAwarenessOk returns a tuple with the ContentAwareness field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentAwareness

`func (o *KalturaConversionProfileAssetParams) SetContentAwareness(v float32)`

SetContentAwareness sets ContentAwareness field to given value.

### HasContentAwareness

`func (o *KalturaConversionProfileAssetParams) HasContentAwareness() bool`

HasContentAwareness returns a boolean if a field has been set.

### GetConversionProfileId

`func (o *KalturaConversionProfileAssetParams) GetConversionProfileId() int32`

GetConversionProfileId returns the ConversionProfileId field if non-nil, zero value otherwise.

### GetConversionProfileIdOk

`func (o *KalturaConversionProfileAssetParams) GetConversionProfileIdOk() (*int32, bool)`

GetConversionProfileIdOk returns a tuple with the ConversionProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConversionProfileId

`func (o *KalturaConversionProfileAssetParams) SetConversionProfileId(v int32)`

SetConversionProfileId sets ConversionProfileId field to given value.

### HasConversionProfileId

`func (o *KalturaConversionProfileAssetParams) HasConversionProfileId() bool`

HasConversionProfileId returns a boolean if a field has been set.

### GetDeletePolicy

`func (o *KalturaConversionProfileAssetParams) GetDeletePolicy() int32`

GetDeletePolicy returns the DeletePolicy field if non-nil, zero value otherwise.

### GetDeletePolicyOk

`func (o *KalturaConversionProfileAssetParams) GetDeletePolicyOk() (*int32, bool)`

GetDeletePolicyOk returns a tuple with the DeletePolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletePolicy

`func (o *KalturaConversionProfileAssetParams) SetDeletePolicy(v int32)`

SetDeletePolicy sets DeletePolicy field to given value.

### HasDeletePolicy

`func (o *KalturaConversionProfileAssetParams) HasDeletePolicy() bool`

HasDeletePolicy returns a boolean if a field has been set.

### GetForceNoneComplied

`func (o *KalturaConversionProfileAssetParams) GetForceNoneComplied() int32`

GetForceNoneComplied returns the ForceNoneComplied field if non-nil, zero value otherwise.

### GetForceNoneCompliedOk

`func (o *KalturaConversionProfileAssetParams) GetForceNoneCompliedOk() (*int32, bool)`

GetForceNoneCompliedOk returns a tuple with the ForceNoneComplied field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForceNoneComplied

`func (o *KalturaConversionProfileAssetParams) SetForceNoneComplied(v int32)`

SetForceNoneComplied sets ForceNoneComplied field to given value.

### HasForceNoneComplied

`func (o *KalturaConversionProfileAssetParams) HasForceNoneComplied() bool`

HasForceNoneComplied returns a boolean if a field has been set.

### GetIsEncrypted

`func (o *KalturaConversionProfileAssetParams) GetIsEncrypted() int32`

GetIsEncrypted returns the IsEncrypted field if non-nil, zero value otherwise.

### GetIsEncryptedOk

`func (o *KalturaConversionProfileAssetParams) GetIsEncryptedOk() (*int32, bool)`

GetIsEncryptedOk returns a tuple with the IsEncrypted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEncrypted

`func (o *KalturaConversionProfileAssetParams) SetIsEncrypted(v int32)`

SetIsEncrypted sets IsEncrypted field to given value.

### HasIsEncrypted

`func (o *KalturaConversionProfileAssetParams) HasIsEncrypted() bool`

HasIsEncrypted returns a boolean if a field has been set.

### GetOrigin

`func (o *KalturaConversionProfileAssetParams) GetOrigin() int32`

GetOrigin returns the Origin field if non-nil, zero value otherwise.

### GetOriginOk

`func (o *KalturaConversionProfileAssetParams) GetOriginOk() (*int32, bool)`

GetOriginOk returns a tuple with the Origin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigin

`func (o *KalturaConversionProfileAssetParams) SetOrigin(v int32)`

SetOrigin sets Origin field to given value.

### HasOrigin

`func (o *KalturaConversionProfileAssetParams) HasOrigin() bool`

HasOrigin returns a boolean if a field has been set.

### GetOverloadParams

`func (o *KalturaConversionProfileAssetParams) GetOverloadParams() string`

GetOverloadParams returns the OverloadParams field if non-nil, zero value otherwise.

### GetOverloadParamsOk

`func (o *KalturaConversionProfileAssetParams) GetOverloadParamsOk() (*string, bool)`

GetOverloadParamsOk returns a tuple with the OverloadParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverloadParams

`func (o *KalturaConversionProfileAssetParams) SetOverloadParams(v string)`

SetOverloadParams sets OverloadParams field to given value.

### HasOverloadParams

`func (o *KalturaConversionProfileAssetParams) HasOverloadParams() bool`

HasOverloadParams returns a boolean if a field has been set.

### GetReadyBehavior

`func (o *KalturaConversionProfileAssetParams) GetReadyBehavior() int32`

GetReadyBehavior returns the ReadyBehavior field if non-nil, zero value otherwise.

### GetReadyBehaviorOk

`func (o *KalturaConversionProfileAssetParams) GetReadyBehaviorOk() (*int32, bool)`

GetReadyBehaviorOk returns a tuple with the ReadyBehavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadyBehavior

`func (o *KalturaConversionProfileAssetParams) SetReadyBehavior(v int32)`

SetReadyBehavior sets ReadyBehavior field to given value.

### HasReadyBehavior

`func (o *KalturaConversionProfileAssetParams) HasReadyBehavior() bool`

HasReadyBehavior returns a boolean if a field has been set.

### GetSystemName

`func (o *KalturaConversionProfileAssetParams) GetSystemName() string`

GetSystemName returns the SystemName field if non-nil, zero value otherwise.

### GetSystemNameOk

`func (o *KalturaConversionProfileAssetParams) GetSystemNameOk() (*string, bool)`

GetSystemNameOk returns a tuple with the SystemName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemName

`func (o *KalturaConversionProfileAssetParams) SetSystemName(v string)`

SetSystemName sets SystemName field to given value.

### HasSystemName

`func (o *KalturaConversionProfileAssetParams) HasSystemName() bool`

HasSystemName returns a boolean if a field has been set.

### GetTags

`func (o *KalturaConversionProfileAssetParams) GetTags() string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *KalturaConversionProfileAssetParams) GetTagsOk() (*string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *KalturaConversionProfileAssetParams) SetTags(v string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *KalturaConversionProfileAssetParams) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTwoPass

`func (o *KalturaConversionProfileAssetParams) GetTwoPass() int32`

GetTwoPass returns the TwoPass field if non-nil, zero value otherwise.

### GetTwoPassOk

`func (o *KalturaConversionProfileAssetParams) GetTwoPassOk() (*int32, bool)`

GetTwoPassOk returns a tuple with the TwoPass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTwoPass

`func (o *KalturaConversionProfileAssetParams) SetTwoPass(v int32)`

SetTwoPass sets TwoPass field to given value.

### HasTwoPass

`func (o *KalturaConversionProfileAssetParams) HasTwoPass() bool`

HasTwoPass returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


