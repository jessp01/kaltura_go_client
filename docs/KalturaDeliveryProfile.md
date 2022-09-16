# KalturaDeliveryProfile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **int32** | &#x60;readOnly&#x60;  Creation time as Unix timestamp (In seconds) | [optional] [readonly] 
**Description** | Pointer to **string** | The description of the Delivery | [optional] 
**ExtraParams** | Pointer to **string** | Extra query string parameters that should be added to the url | [optional] 
**HostName** | Pointer to **string** | &#x60;readOnly&#x60;  the host part of the url | [optional] [readonly] 
**Id** | Pointer to **int32** | &#x60;readOnly&#x60;  The id of the Delivery | [optional] [readonly] 
**IsDefault** | Pointer to **int32** | &#x60;readOnly&#x60;  Enum Type: &#x60;KalturaNullableBoolean&#x60;  True if this is the systemwide default for the protocol | [optional] [readonly] 
**MediaProtocols** | Pointer to **string** | Comma separated list of supported media protocols. f.i. rtmpe | [optional] 
**Name** | Pointer to **string** | The name of the Delivery | [optional] 
**ObjectType** | Pointer to **string** |  | [optional] 
**ParentId** | Pointer to **int32** | &#x60;readOnly&#x60;  the object from which this object was cloned (or 0) | [optional] [readonly] 
**PartnerId** | Pointer to **int32** | &#x60;readOnly&#x60; | [optional] [readonly] 
**Priority** | Pointer to **int32** | priority used for ordering similar delivery profiles | [optional] 
**Recognizer** | Pointer to [**KalturaUrlRecognizer**](KalturaUrlRecognizer.md) |  | [optional] 
**Status** | Pointer to **int32** | Enum Type: &#x60;KalturaDeliveryStatus&#x60; | [optional] 
**StreamerType** | Pointer to **string** | Enum Type: &#x60;KalturaPlaybackProtocol&#x60; | [optional] 
**SupplementaryAssetsFilter** | Pointer to [**KalturaAssetFilter**](KalturaAssetFilter.md) |  | [optional] 
**SystemName** | Pointer to **string** | System name of the delivery | [optional] 
**Tokenizer** | Pointer to [**KalturaUrlTokenizer**](KalturaUrlTokenizer.md) |  | [optional] 
**Type** | Pointer to **string** | Enum Type: &#x60;KalturaDeliveryProfileType&#x60;  Delivery type | [optional] 
**UpdatedAt** | Pointer to **int32** | &#x60;readOnly&#x60;  Update time as Unix timestamp (In seconds) | [optional] [readonly] 
**Url** | Pointer to **string** |  | [optional] 

## Methods

### NewKalturaDeliveryProfile

`func NewKalturaDeliveryProfile() *KalturaDeliveryProfile`

NewKalturaDeliveryProfile instantiates a new KalturaDeliveryProfile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKalturaDeliveryProfileWithDefaults

`func NewKalturaDeliveryProfileWithDefaults() *KalturaDeliveryProfile`

NewKalturaDeliveryProfileWithDefaults instantiates a new KalturaDeliveryProfile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *KalturaDeliveryProfile) GetCreatedAt() int32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *KalturaDeliveryProfile) GetCreatedAtOk() (*int32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *KalturaDeliveryProfile) SetCreatedAt(v int32)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *KalturaDeliveryProfile) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDescription

`func (o *KalturaDeliveryProfile) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *KalturaDeliveryProfile) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *KalturaDeliveryProfile) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *KalturaDeliveryProfile) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetExtraParams

`func (o *KalturaDeliveryProfile) GetExtraParams() string`

GetExtraParams returns the ExtraParams field if non-nil, zero value otherwise.

### GetExtraParamsOk

`func (o *KalturaDeliveryProfile) GetExtraParamsOk() (*string, bool)`

GetExtraParamsOk returns a tuple with the ExtraParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraParams

`func (o *KalturaDeliveryProfile) SetExtraParams(v string)`

SetExtraParams sets ExtraParams field to given value.

### HasExtraParams

`func (o *KalturaDeliveryProfile) HasExtraParams() bool`

HasExtraParams returns a boolean if a field has been set.

### GetHostName

`func (o *KalturaDeliveryProfile) GetHostName() string`

GetHostName returns the HostName field if non-nil, zero value otherwise.

### GetHostNameOk

`func (o *KalturaDeliveryProfile) GetHostNameOk() (*string, bool)`

GetHostNameOk returns a tuple with the HostName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostName

`func (o *KalturaDeliveryProfile) SetHostName(v string)`

SetHostName sets HostName field to given value.

### HasHostName

`func (o *KalturaDeliveryProfile) HasHostName() bool`

HasHostName returns a boolean if a field has been set.

### GetId

`func (o *KalturaDeliveryProfile) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *KalturaDeliveryProfile) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *KalturaDeliveryProfile) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *KalturaDeliveryProfile) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIsDefault

`func (o *KalturaDeliveryProfile) GetIsDefault() int32`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *KalturaDeliveryProfile) GetIsDefaultOk() (*int32, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *KalturaDeliveryProfile) SetIsDefault(v int32)`

SetIsDefault sets IsDefault field to given value.

### HasIsDefault

`func (o *KalturaDeliveryProfile) HasIsDefault() bool`

HasIsDefault returns a boolean if a field has been set.

### GetMediaProtocols

`func (o *KalturaDeliveryProfile) GetMediaProtocols() string`

GetMediaProtocols returns the MediaProtocols field if non-nil, zero value otherwise.

### GetMediaProtocolsOk

`func (o *KalturaDeliveryProfile) GetMediaProtocolsOk() (*string, bool)`

GetMediaProtocolsOk returns a tuple with the MediaProtocols field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaProtocols

`func (o *KalturaDeliveryProfile) SetMediaProtocols(v string)`

SetMediaProtocols sets MediaProtocols field to given value.

### HasMediaProtocols

`func (o *KalturaDeliveryProfile) HasMediaProtocols() bool`

HasMediaProtocols returns a boolean if a field has been set.

### GetName

`func (o *KalturaDeliveryProfile) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *KalturaDeliveryProfile) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *KalturaDeliveryProfile) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *KalturaDeliveryProfile) HasName() bool`

HasName returns a boolean if a field has been set.

### GetObjectType

`func (o *KalturaDeliveryProfile) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *KalturaDeliveryProfile) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *KalturaDeliveryProfile) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.

### HasObjectType

`func (o *KalturaDeliveryProfile) HasObjectType() bool`

HasObjectType returns a boolean if a field has been set.

### GetParentId

`func (o *KalturaDeliveryProfile) GetParentId() int32`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *KalturaDeliveryProfile) GetParentIdOk() (*int32, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *KalturaDeliveryProfile) SetParentId(v int32)`

SetParentId sets ParentId field to given value.

### HasParentId

`func (o *KalturaDeliveryProfile) HasParentId() bool`

HasParentId returns a boolean if a field has been set.

### GetPartnerId

`func (o *KalturaDeliveryProfile) GetPartnerId() int32`

GetPartnerId returns the PartnerId field if non-nil, zero value otherwise.

### GetPartnerIdOk

`func (o *KalturaDeliveryProfile) GetPartnerIdOk() (*int32, bool)`

GetPartnerIdOk returns a tuple with the PartnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerId

`func (o *KalturaDeliveryProfile) SetPartnerId(v int32)`

SetPartnerId sets PartnerId field to given value.

### HasPartnerId

`func (o *KalturaDeliveryProfile) HasPartnerId() bool`

HasPartnerId returns a boolean if a field has been set.

### GetPriority

`func (o *KalturaDeliveryProfile) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *KalturaDeliveryProfile) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *KalturaDeliveryProfile) SetPriority(v int32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *KalturaDeliveryProfile) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetRecognizer

`func (o *KalturaDeliveryProfile) GetRecognizer() KalturaUrlRecognizer`

GetRecognizer returns the Recognizer field if non-nil, zero value otherwise.

### GetRecognizerOk

`func (o *KalturaDeliveryProfile) GetRecognizerOk() (*KalturaUrlRecognizer, bool)`

GetRecognizerOk returns a tuple with the Recognizer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecognizer

`func (o *KalturaDeliveryProfile) SetRecognizer(v KalturaUrlRecognizer)`

SetRecognizer sets Recognizer field to given value.

### HasRecognizer

`func (o *KalturaDeliveryProfile) HasRecognizer() bool`

HasRecognizer returns a boolean if a field has been set.

### GetStatus

`func (o *KalturaDeliveryProfile) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *KalturaDeliveryProfile) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *KalturaDeliveryProfile) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *KalturaDeliveryProfile) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStreamerType

`func (o *KalturaDeliveryProfile) GetStreamerType() string`

GetStreamerType returns the StreamerType field if non-nil, zero value otherwise.

### GetStreamerTypeOk

`func (o *KalturaDeliveryProfile) GetStreamerTypeOk() (*string, bool)`

GetStreamerTypeOk returns a tuple with the StreamerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreamerType

`func (o *KalturaDeliveryProfile) SetStreamerType(v string)`

SetStreamerType sets StreamerType field to given value.

### HasStreamerType

`func (o *KalturaDeliveryProfile) HasStreamerType() bool`

HasStreamerType returns a boolean if a field has been set.

### GetSupplementaryAssetsFilter

`func (o *KalturaDeliveryProfile) GetSupplementaryAssetsFilter() KalturaAssetFilter`

GetSupplementaryAssetsFilter returns the SupplementaryAssetsFilter field if non-nil, zero value otherwise.

### GetSupplementaryAssetsFilterOk

`func (o *KalturaDeliveryProfile) GetSupplementaryAssetsFilterOk() (*KalturaAssetFilter, bool)`

GetSupplementaryAssetsFilterOk returns a tuple with the SupplementaryAssetsFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupplementaryAssetsFilter

`func (o *KalturaDeliveryProfile) SetSupplementaryAssetsFilter(v KalturaAssetFilter)`

SetSupplementaryAssetsFilter sets SupplementaryAssetsFilter field to given value.

### HasSupplementaryAssetsFilter

`func (o *KalturaDeliveryProfile) HasSupplementaryAssetsFilter() bool`

HasSupplementaryAssetsFilter returns a boolean if a field has been set.

### GetSystemName

`func (o *KalturaDeliveryProfile) GetSystemName() string`

GetSystemName returns the SystemName field if non-nil, zero value otherwise.

### GetSystemNameOk

`func (o *KalturaDeliveryProfile) GetSystemNameOk() (*string, bool)`

GetSystemNameOk returns a tuple with the SystemName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemName

`func (o *KalturaDeliveryProfile) SetSystemName(v string)`

SetSystemName sets SystemName field to given value.

### HasSystemName

`func (o *KalturaDeliveryProfile) HasSystemName() bool`

HasSystemName returns a boolean if a field has been set.

### GetTokenizer

`func (o *KalturaDeliveryProfile) GetTokenizer() KalturaUrlTokenizer`

GetTokenizer returns the Tokenizer field if non-nil, zero value otherwise.

### GetTokenizerOk

`func (o *KalturaDeliveryProfile) GetTokenizerOk() (*KalturaUrlTokenizer, bool)`

GetTokenizerOk returns a tuple with the Tokenizer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenizer

`func (o *KalturaDeliveryProfile) SetTokenizer(v KalturaUrlTokenizer)`

SetTokenizer sets Tokenizer field to given value.

### HasTokenizer

`func (o *KalturaDeliveryProfile) HasTokenizer() bool`

HasTokenizer returns a boolean if a field has been set.

### GetType

`func (o *KalturaDeliveryProfile) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *KalturaDeliveryProfile) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *KalturaDeliveryProfile) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *KalturaDeliveryProfile) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *KalturaDeliveryProfile) GetUpdatedAt() int32`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *KalturaDeliveryProfile) GetUpdatedAtOk() (*int32, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *KalturaDeliveryProfile) SetUpdatedAt(v int32)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *KalturaDeliveryProfile) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUrl

`func (o *KalturaDeliveryProfile) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *KalturaDeliveryProfile) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *KalturaDeliveryProfile) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *KalturaDeliveryProfile) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


