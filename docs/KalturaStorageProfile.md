# KalturaStorageProfile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowAutoDelete** | Pointer to **int32** | Flag sugnifying that the storage exported content should be deleted when soure entry is deleted | [optional] 
**CreateFileLink** | Pointer to **bool** | Indicates to the local file transfer manager to create a link to the file instead of copying it | [optional] 
**CreatedAt** | Pointer to **int32** | &#x60;readOnly&#x60; | [optional] [readonly] 
**DeliveryPriority** | Pointer to **int32** | Delivery Priority | [optional] 
**DeliveryProfileIds** | Pointer to [**[]KalturaKeyValue**](KalturaKeyValue.md) |  | [optional] 
**DeliveryStatus** | Pointer to **int32** | Enum Type: &#x60;KalturaStorageProfileDeliveryStatus&#x60; | [optional] 
**Desciption** | Pointer to **string** |  | [optional] 
**ExcludedEntryTypes** | Pointer to **string** |  | [optional] 
**ExcludedFlavorParamsIds** | Pointer to **string** |  | [optional] 
**ExportPeriodically** | Pointer to **bool** |  | [optional] 
**FlavorParamsIds** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **int32** | &#x60;readOnly&#x60; | [optional] [readonly] 
**MaxConcurrentConnections** | Pointer to **int32** |  | [optional] 
**MaxFileSize** | Pointer to **int32** |  | [optional] 
**MinFileSize** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**ObjectType** | Pointer to **string** |  | [optional] 
**PackagerUrl** | Pointer to **string** |  | [optional] 
**PartnerId** | Pointer to **int32** | &#x60;readOnly&#x60; | [optional] [readonly] 
**PassPhrase** | Pointer to **string** |  | [optional] 
**PathManagerClass** | Pointer to **string** |  | [optional] 
**PathManagerParams** | Pointer to [**[]KalturaKeyValue**](KalturaKeyValue.md) |  | [optional] 
**PathPrefix** | Pointer to **string** |  | [optional] 
**Port** | Pointer to **int32** |  | [optional] 
**PrivateKey** | Pointer to **string** |  | [optional] 
**Protocol** | Pointer to **string** | Enum Type: &#x60;KalturaStorageProfileProtocol&#x60; | [optional] 
**PublicKey** | Pointer to **string** |  | [optional] 
**ReadyBehavior** | Pointer to **int32** | Enum Type: &#x60;KalturaStorageProfileReadyBehavior&#x60; | [optional] 
**Rules** | Pointer to [**[]KalturaRule**](KalturaRule.md) |  | [optional] 
**ShouldExportCaptions** | Pointer to **bool** |  | [optional] 
**ShouldExportThumbs** | Pointer to **bool** |  | [optional] 
**Status** | Pointer to **int32** | Enum Type: &#x60;KalturaStorageProfileStatus&#x60; | [optional] 
**StorageBaseDir** | Pointer to **string** |  | [optional] 
**StorageFtpPassiveMode** | Pointer to **bool** |  | [optional] 
**StoragePassword** | Pointer to **string** |  | [optional] 
**StorageUrl** | Pointer to **string** |  | [optional] 
**StorageUsername** | Pointer to **string** |  | [optional] 
**SystemName** | Pointer to **string** |  | [optional] 
**Trigger** | Pointer to **int32** | No need to create enum for temp field | [optional] 
**UpdatedAt** | Pointer to **int32** | &#x60;readOnly&#x60; | [optional] [readonly] 

## Methods

### NewKalturaStorageProfile

`func NewKalturaStorageProfile() *KalturaStorageProfile`

NewKalturaStorageProfile instantiates a new KalturaStorageProfile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKalturaStorageProfileWithDefaults

`func NewKalturaStorageProfileWithDefaults() *KalturaStorageProfile`

NewKalturaStorageProfileWithDefaults instantiates a new KalturaStorageProfile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowAutoDelete

`func (o *KalturaStorageProfile) GetAllowAutoDelete() int32`

GetAllowAutoDelete returns the AllowAutoDelete field if non-nil, zero value otherwise.

### GetAllowAutoDeleteOk

`func (o *KalturaStorageProfile) GetAllowAutoDeleteOk() (*int32, bool)`

GetAllowAutoDeleteOk returns a tuple with the AllowAutoDelete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowAutoDelete

`func (o *KalturaStorageProfile) SetAllowAutoDelete(v int32)`

SetAllowAutoDelete sets AllowAutoDelete field to given value.

### HasAllowAutoDelete

`func (o *KalturaStorageProfile) HasAllowAutoDelete() bool`

HasAllowAutoDelete returns a boolean if a field has been set.

### GetCreateFileLink

`func (o *KalturaStorageProfile) GetCreateFileLink() bool`

GetCreateFileLink returns the CreateFileLink field if non-nil, zero value otherwise.

### GetCreateFileLinkOk

`func (o *KalturaStorageProfile) GetCreateFileLinkOk() (*bool, bool)`

GetCreateFileLinkOk returns a tuple with the CreateFileLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateFileLink

`func (o *KalturaStorageProfile) SetCreateFileLink(v bool)`

SetCreateFileLink sets CreateFileLink field to given value.

### HasCreateFileLink

`func (o *KalturaStorageProfile) HasCreateFileLink() bool`

HasCreateFileLink returns a boolean if a field has been set.

### GetCreatedAt

`func (o *KalturaStorageProfile) GetCreatedAt() int32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *KalturaStorageProfile) GetCreatedAtOk() (*int32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *KalturaStorageProfile) SetCreatedAt(v int32)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *KalturaStorageProfile) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDeliveryPriority

`func (o *KalturaStorageProfile) GetDeliveryPriority() int32`

GetDeliveryPriority returns the DeliveryPriority field if non-nil, zero value otherwise.

### GetDeliveryPriorityOk

`func (o *KalturaStorageProfile) GetDeliveryPriorityOk() (*int32, bool)`

GetDeliveryPriorityOk returns a tuple with the DeliveryPriority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryPriority

`func (o *KalturaStorageProfile) SetDeliveryPriority(v int32)`

SetDeliveryPriority sets DeliveryPriority field to given value.

### HasDeliveryPriority

`func (o *KalturaStorageProfile) HasDeliveryPriority() bool`

HasDeliveryPriority returns a boolean if a field has been set.

### GetDeliveryProfileIds

`func (o *KalturaStorageProfile) GetDeliveryProfileIds() []KalturaKeyValue`

GetDeliveryProfileIds returns the DeliveryProfileIds field if non-nil, zero value otherwise.

### GetDeliveryProfileIdsOk

`func (o *KalturaStorageProfile) GetDeliveryProfileIdsOk() (*[]KalturaKeyValue, bool)`

GetDeliveryProfileIdsOk returns a tuple with the DeliveryProfileIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryProfileIds

`func (o *KalturaStorageProfile) SetDeliveryProfileIds(v []KalturaKeyValue)`

SetDeliveryProfileIds sets DeliveryProfileIds field to given value.

### HasDeliveryProfileIds

`func (o *KalturaStorageProfile) HasDeliveryProfileIds() bool`

HasDeliveryProfileIds returns a boolean if a field has been set.

### GetDeliveryStatus

`func (o *KalturaStorageProfile) GetDeliveryStatus() int32`

GetDeliveryStatus returns the DeliveryStatus field if non-nil, zero value otherwise.

### GetDeliveryStatusOk

`func (o *KalturaStorageProfile) GetDeliveryStatusOk() (*int32, bool)`

GetDeliveryStatusOk returns a tuple with the DeliveryStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryStatus

`func (o *KalturaStorageProfile) SetDeliveryStatus(v int32)`

SetDeliveryStatus sets DeliveryStatus field to given value.

### HasDeliveryStatus

`func (o *KalturaStorageProfile) HasDeliveryStatus() bool`

HasDeliveryStatus returns a boolean if a field has been set.

### GetDesciption

`func (o *KalturaStorageProfile) GetDesciption() string`

GetDesciption returns the Desciption field if non-nil, zero value otherwise.

### GetDesciptionOk

`func (o *KalturaStorageProfile) GetDesciptionOk() (*string, bool)`

GetDesciptionOk returns a tuple with the Desciption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesciption

`func (o *KalturaStorageProfile) SetDesciption(v string)`

SetDesciption sets Desciption field to given value.

### HasDesciption

`func (o *KalturaStorageProfile) HasDesciption() bool`

HasDesciption returns a boolean if a field has been set.

### GetExcludedEntryTypes

`func (o *KalturaStorageProfile) GetExcludedEntryTypes() string`

GetExcludedEntryTypes returns the ExcludedEntryTypes field if non-nil, zero value otherwise.

### GetExcludedEntryTypesOk

`func (o *KalturaStorageProfile) GetExcludedEntryTypesOk() (*string, bool)`

GetExcludedEntryTypesOk returns a tuple with the ExcludedEntryTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludedEntryTypes

`func (o *KalturaStorageProfile) SetExcludedEntryTypes(v string)`

SetExcludedEntryTypes sets ExcludedEntryTypes field to given value.

### HasExcludedEntryTypes

`func (o *KalturaStorageProfile) HasExcludedEntryTypes() bool`

HasExcludedEntryTypes returns a boolean if a field has been set.

### GetExcludedFlavorParamsIds

`func (o *KalturaStorageProfile) GetExcludedFlavorParamsIds() string`

GetExcludedFlavorParamsIds returns the ExcludedFlavorParamsIds field if non-nil, zero value otherwise.

### GetExcludedFlavorParamsIdsOk

`func (o *KalturaStorageProfile) GetExcludedFlavorParamsIdsOk() (*string, bool)`

GetExcludedFlavorParamsIdsOk returns a tuple with the ExcludedFlavorParamsIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludedFlavorParamsIds

`func (o *KalturaStorageProfile) SetExcludedFlavorParamsIds(v string)`

SetExcludedFlavorParamsIds sets ExcludedFlavorParamsIds field to given value.

### HasExcludedFlavorParamsIds

`func (o *KalturaStorageProfile) HasExcludedFlavorParamsIds() bool`

HasExcludedFlavorParamsIds returns a boolean if a field has been set.

### GetExportPeriodically

`func (o *KalturaStorageProfile) GetExportPeriodically() bool`

GetExportPeriodically returns the ExportPeriodically field if non-nil, zero value otherwise.

### GetExportPeriodicallyOk

`func (o *KalturaStorageProfile) GetExportPeriodicallyOk() (*bool, bool)`

GetExportPeriodicallyOk returns a tuple with the ExportPeriodically field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExportPeriodically

`func (o *KalturaStorageProfile) SetExportPeriodically(v bool)`

SetExportPeriodically sets ExportPeriodically field to given value.

### HasExportPeriodically

`func (o *KalturaStorageProfile) HasExportPeriodically() bool`

HasExportPeriodically returns a boolean if a field has been set.

### GetFlavorParamsIds

`func (o *KalturaStorageProfile) GetFlavorParamsIds() string`

GetFlavorParamsIds returns the FlavorParamsIds field if non-nil, zero value otherwise.

### GetFlavorParamsIdsOk

`func (o *KalturaStorageProfile) GetFlavorParamsIdsOk() (*string, bool)`

GetFlavorParamsIdsOk returns a tuple with the FlavorParamsIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlavorParamsIds

`func (o *KalturaStorageProfile) SetFlavorParamsIds(v string)`

SetFlavorParamsIds sets FlavorParamsIds field to given value.

### HasFlavorParamsIds

`func (o *KalturaStorageProfile) HasFlavorParamsIds() bool`

HasFlavorParamsIds returns a boolean if a field has been set.

### GetId

`func (o *KalturaStorageProfile) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *KalturaStorageProfile) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *KalturaStorageProfile) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *KalturaStorageProfile) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMaxConcurrentConnections

`func (o *KalturaStorageProfile) GetMaxConcurrentConnections() int32`

GetMaxConcurrentConnections returns the MaxConcurrentConnections field if non-nil, zero value otherwise.

### GetMaxConcurrentConnectionsOk

`func (o *KalturaStorageProfile) GetMaxConcurrentConnectionsOk() (*int32, bool)`

GetMaxConcurrentConnectionsOk returns a tuple with the MaxConcurrentConnections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxConcurrentConnections

`func (o *KalturaStorageProfile) SetMaxConcurrentConnections(v int32)`

SetMaxConcurrentConnections sets MaxConcurrentConnections field to given value.

### HasMaxConcurrentConnections

`func (o *KalturaStorageProfile) HasMaxConcurrentConnections() bool`

HasMaxConcurrentConnections returns a boolean if a field has been set.

### GetMaxFileSize

`func (o *KalturaStorageProfile) GetMaxFileSize() int32`

GetMaxFileSize returns the MaxFileSize field if non-nil, zero value otherwise.

### GetMaxFileSizeOk

`func (o *KalturaStorageProfile) GetMaxFileSizeOk() (*int32, bool)`

GetMaxFileSizeOk returns a tuple with the MaxFileSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxFileSize

`func (o *KalturaStorageProfile) SetMaxFileSize(v int32)`

SetMaxFileSize sets MaxFileSize field to given value.

### HasMaxFileSize

`func (o *KalturaStorageProfile) HasMaxFileSize() bool`

HasMaxFileSize returns a boolean if a field has been set.

### GetMinFileSize

`func (o *KalturaStorageProfile) GetMinFileSize() int32`

GetMinFileSize returns the MinFileSize field if non-nil, zero value otherwise.

### GetMinFileSizeOk

`func (o *KalturaStorageProfile) GetMinFileSizeOk() (*int32, bool)`

GetMinFileSizeOk returns a tuple with the MinFileSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinFileSize

`func (o *KalturaStorageProfile) SetMinFileSize(v int32)`

SetMinFileSize sets MinFileSize field to given value.

### HasMinFileSize

`func (o *KalturaStorageProfile) HasMinFileSize() bool`

HasMinFileSize returns a boolean if a field has been set.

### GetName

`func (o *KalturaStorageProfile) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *KalturaStorageProfile) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *KalturaStorageProfile) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *KalturaStorageProfile) HasName() bool`

HasName returns a boolean if a field has been set.

### GetObjectType

`func (o *KalturaStorageProfile) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *KalturaStorageProfile) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *KalturaStorageProfile) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.

### HasObjectType

`func (o *KalturaStorageProfile) HasObjectType() bool`

HasObjectType returns a boolean if a field has been set.

### GetPackagerUrl

`func (o *KalturaStorageProfile) GetPackagerUrl() string`

GetPackagerUrl returns the PackagerUrl field if non-nil, zero value otherwise.

### GetPackagerUrlOk

`func (o *KalturaStorageProfile) GetPackagerUrlOk() (*string, bool)`

GetPackagerUrlOk returns a tuple with the PackagerUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackagerUrl

`func (o *KalturaStorageProfile) SetPackagerUrl(v string)`

SetPackagerUrl sets PackagerUrl field to given value.

### HasPackagerUrl

`func (o *KalturaStorageProfile) HasPackagerUrl() bool`

HasPackagerUrl returns a boolean if a field has been set.

### GetPartnerId

`func (o *KalturaStorageProfile) GetPartnerId() int32`

GetPartnerId returns the PartnerId field if non-nil, zero value otherwise.

### GetPartnerIdOk

`func (o *KalturaStorageProfile) GetPartnerIdOk() (*int32, bool)`

GetPartnerIdOk returns a tuple with the PartnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerId

`func (o *KalturaStorageProfile) SetPartnerId(v int32)`

SetPartnerId sets PartnerId field to given value.

### HasPartnerId

`func (o *KalturaStorageProfile) HasPartnerId() bool`

HasPartnerId returns a boolean if a field has been set.

### GetPassPhrase

`func (o *KalturaStorageProfile) GetPassPhrase() string`

GetPassPhrase returns the PassPhrase field if non-nil, zero value otherwise.

### GetPassPhraseOk

`func (o *KalturaStorageProfile) GetPassPhraseOk() (*string, bool)`

GetPassPhraseOk returns a tuple with the PassPhrase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassPhrase

`func (o *KalturaStorageProfile) SetPassPhrase(v string)`

SetPassPhrase sets PassPhrase field to given value.

### HasPassPhrase

`func (o *KalturaStorageProfile) HasPassPhrase() bool`

HasPassPhrase returns a boolean if a field has been set.

### GetPathManagerClass

`func (o *KalturaStorageProfile) GetPathManagerClass() string`

GetPathManagerClass returns the PathManagerClass field if non-nil, zero value otherwise.

### GetPathManagerClassOk

`func (o *KalturaStorageProfile) GetPathManagerClassOk() (*string, bool)`

GetPathManagerClassOk returns a tuple with the PathManagerClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPathManagerClass

`func (o *KalturaStorageProfile) SetPathManagerClass(v string)`

SetPathManagerClass sets PathManagerClass field to given value.

### HasPathManagerClass

`func (o *KalturaStorageProfile) HasPathManagerClass() bool`

HasPathManagerClass returns a boolean if a field has been set.

### GetPathManagerParams

`func (o *KalturaStorageProfile) GetPathManagerParams() []KalturaKeyValue`

GetPathManagerParams returns the PathManagerParams field if non-nil, zero value otherwise.

### GetPathManagerParamsOk

`func (o *KalturaStorageProfile) GetPathManagerParamsOk() (*[]KalturaKeyValue, bool)`

GetPathManagerParamsOk returns a tuple with the PathManagerParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPathManagerParams

`func (o *KalturaStorageProfile) SetPathManagerParams(v []KalturaKeyValue)`

SetPathManagerParams sets PathManagerParams field to given value.

### HasPathManagerParams

`func (o *KalturaStorageProfile) HasPathManagerParams() bool`

HasPathManagerParams returns a boolean if a field has been set.

### GetPathPrefix

`func (o *KalturaStorageProfile) GetPathPrefix() string`

GetPathPrefix returns the PathPrefix field if non-nil, zero value otherwise.

### GetPathPrefixOk

`func (o *KalturaStorageProfile) GetPathPrefixOk() (*string, bool)`

GetPathPrefixOk returns a tuple with the PathPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPathPrefix

`func (o *KalturaStorageProfile) SetPathPrefix(v string)`

SetPathPrefix sets PathPrefix field to given value.

### HasPathPrefix

`func (o *KalturaStorageProfile) HasPathPrefix() bool`

HasPathPrefix returns a boolean if a field has been set.

### GetPort

`func (o *KalturaStorageProfile) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *KalturaStorageProfile) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *KalturaStorageProfile) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *KalturaStorageProfile) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetPrivateKey

`func (o *KalturaStorageProfile) GetPrivateKey() string`

GetPrivateKey returns the PrivateKey field if non-nil, zero value otherwise.

### GetPrivateKeyOk

`func (o *KalturaStorageProfile) GetPrivateKeyOk() (*string, bool)`

GetPrivateKeyOk returns a tuple with the PrivateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateKey

`func (o *KalturaStorageProfile) SetPrivateKey(v string)`

SetPrivateKey sets PrivateKey field to given value.

### HasPrivateKey

`func (o *KalturaStorageProfile) HasPrivateKey() bool`

HasPrivateKey returns a boolean if a field has been set.

### GetProtocol

`func (o *KalturaStorageProfile) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *KalturaStorageProfile) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *KalturaStorageProfile) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *KalturaStorageProfile) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetPublicKey

`func (o *KalturaStorageProfile) GetPublicKey() string`

GetPublicKey returns the PublicKey field if non-nil, zero value otherwise.

### GetPublicKeyOk

`func (o *KalturaStorageProfile) GetPublicKeyOk() (*string, bool)`

GetPublicKeyOk returns a tuple with the PublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicKey

`func (o *KalturaStorageProfile) SetPublicKey(v string)`

SetPublicKey sets PublicKey field to given value.

### HasPublicKey

`func (o *KalturaStorageProfile) HasPublicKey() bool`

HasPublicKey returns a boolean if a field has been set.

### GetReadyBehavior

`func (o *KalturaStorageProfile) GetReadyBehavior() int32`

GetReadyBehavior returns the ReadyBehavior field if non-nil, zero value otherwise.

### GetReadyBehaviorOk

`func (o *KalturaStorageProfile) GetReadyBehaviorOk() (*int32, bool)`

GetReadyBehaviorOk returns a tuple with the ReadyBehavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadyBehavior

`func (o *KalturaStorageProfile) SetReadyBehavior(v int32)`

SetReadyBehavior sets ReadyBehavior field to given value.

### HasReadyBehavior

`func (o *KalturaStorageProfile) HasReadyBehavior() bool`

HasReadyBehavior returns a boolean if a field has been set.

### GetRules

`func (o *KalturaStorageProfile) GetRules() []KalturaRule`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *KalturaStorageProfile) GetRulesOk() (*[]KalturaRule, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *KalturaStorageProfile) SetRules(v []KalturaRule)`

SetRules sets Rules field to given value.

### HasRules

`func (o *KalturaStorageProfile) HasRules() bool`

HasRules returns a boolean if a field has been set.

### GetShouldExportCaptions

`func (o *KalturaStorageProfile) GetShouldExportCaptions() bool`

GetShouldExportCaptions returns the ShouldExportCaptions field if non-nil, zero value otherwise.

### GetShouldExportCaptionsOk

`func (o *KalturaStorageProfile) GetShouldExportCaptionsOk() (*bool, bool)`

GetShouldExportCaptionsOk returns a tuple with the ShouldExportCaptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShouldExportCaptions

`func (o *KalturaStorageProfile) SetShouldExportCaptions(v bool)`

SetShouldExportCaptions sets ShouldExportCaptions field to given value.

### HasShouldExportCaptions

`func (o *KalturaStorageProfile) HasShouldExportCaptions() bool`

HasShouldExportCaptions returns a boolean if a field has been set.

### GetShouldExportThumbs

`func (o *KalturaStorageProfile) GetShouldExportThumbs() bool`

GetShouldExportThumbs returns the ShouldExportThumbs field if non-nil, zero value otherwise.

### GetShouldExportThumbsOk

`func (o *KalturaStorageProfile) GetShouldExportThumbsOk() (*bool, bool)`

GetShouldExportThumbsOk returns a tuple with the ShouldExportThumbs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShouldExportThumbs

`func (o *KalturaStorageProfile) SetShouldExportThumbs(v bool)`

SetShouldExportThumbs sets ShouldExportThumbs field to given value.

### HasShouldExportThumbs

`func (o *KalturaStorageProfile) HasShouldExportThumbs() bool`

HasShouldExportThumbs returns a boolean if a field has been set.

### GetStatus

`func (o *KalturaStorageProfile) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *KalturaStorageProfile) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *KalturaStorageProfile) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *KalturaStorageProfile) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStorageBaseDir

`func (o *KalturaStorageProfile) GetStorageBaseDir() string`

GetStorageBaseDir returns the StorageBaseDir field if non-nil, zero value otherwise.

### GetStorageBaseDirOk

`func (o *KalturaStorageProfile) GetStorageBaseDirOk() (*string, bool)`

GetStorageBaseDirOk returns a tuple with the StorageBaseDir field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageBaseDir

`func (o *KalturaStorageProfile) SetStorageBaseDir(v string)`

SetStorageBaseDir sets StorageBaseDir field to given value.

### HasStorageBaseDir

`func (o *KalturaStorageProfile) HasStorageBaseDir() bool`

HasStorageBaseDir returns a boolean if a field has been set.

### GetStorageFtpPassiveMode

`func (o *KalturaStorageProfile) GetStorageFtpPassiveMode() bool`

GetStorageFtpPassiveMode returns the StorageFtpPassiveMode field if non-nil, zero value otherwise.

### GetStorageFtpPassiveModeOk

`func (o *KalturaStorageProfile) GetStorageFtpPassiveModeOk() (*bool, bool)`

GetStorageFtpPassiveModeOk returns a tuple with the StorageFtpPassiveMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageFtpPassiveMode

`func (o *KalturaStorageProfile) SetStorageFtpPassiveMode(v bool)`

SetStorageFtpPassiveMode sets StorageFtpPassiveMode field to given value.

### HasStorageFtpPassiveMode

`func (o *KalturaStorageProfile) HasStorageFtpPassiveMode() bool`

HasStorageFtpPassiveMode returns a boolean if a field has been set.

### GetStoragePassword

`func (o *KalturaStorageProfile) GetStoragePassword() string`

GetStoragePassword returns the StoragePassword field if non-nil, zero value otherwise.

### GetStoragePasswordOk

`func (o *KalturaStorageProfile) GetStoragePasswordOk() (*string, bool)`

GetStoragePasswordOk returns a tuple with the StoragePassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoragePassword

`func (o *KalturaStorageProfile) SetStoragePassword(v string)`

SetStoragePassword sets StoragePassword field to given value.

### HasStoragePassword

`func (o *KalturaStorageProfile) HasStoragePassword() bool`

HasStoragePassword returns a boolean if a field has been set.

### GetStorageUrl

`func (o *KalturaStorageProfile) GetStorageUrl() string`

GetStorageUrl returns the StorageUrl field if non-nil, zero value otherwise.

### GetStorageUrlOk

`func (o *KalturaStorageProfile) GetStorageUrlOk() (*string, bool)`

GetStorageUrlOk returns a tuple with the StorageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageUrl

`func (o *KalturaStorageProfile) SetStorageUrl(v string)`

SetStorageUrl sets StorageUrl field to given value.

### HasStorageUrl

`func (o *KalturaStorageProfile) HasStorageUrl() bool`

HasStorageUrl returns a boolean if a field has been set.

### GetStorageUsername

`func (o *KalturaStorageProfile) GetStorageUsername() string`

GetStorageUsername returns the StorageUsername field if non-nil, zero value otherwise.

### GetStorageUsernameOk

`func (o *KalturaStorageProfile) GetStorageUsernameOk() (*string, bool)`

GetStorageUsernameOk returns a tuple with the StorageUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageUsername

`func (o *KalturaStorageProfile) SetStorageUsername(v string)`

SetStorageUsername sets StorageUsername field to given value.

### HasStorageUsername

`func (o *KalturaStorageProfile) HasStorageUsername() bool`

HasStorageUsername returns a boolean if a field has been set.

### GetSystemName

`func (o *KalturaStorageProfile) GetSystemName() string`

GetSystemName returns the SystemName field if non-nil, zero value otherwise.

### GetSystemNameOk

`func (o *KalturaStorageProfile) GetSystemNameOk() (*string, bool)`

GetSystemNameOk returns a tuple with the SystemName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemName

`func (o *KalturaStorageProfile) SetSystemName(v string)`

SetSystemName sets SystemName field to given value.

### HasSystemName

`func (o *KalturaStorageProfile) HasSystemName() bool`

HasSystemName returns a boolean if a field has been set.

### GetTrigger

`func (o *KalturaStorageProfile) GetTrigger() int32`

GetTrigger returns the Trigger field if non-nil, zero value otherwise.

### GetTriggerOk

`func (o *KalturaStorageProfile) GetTriggerOk() (*int32, bool)`

GetTriggerOk returns a tuple with the Trigger field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrigger

`func (o *KalturaStorageProfile) SetTrigger(v int32)`

SetTrigger sets Trigger field to given value.

### HasTrigger

`func (o *KalturaStorageProfile) HasTrigger() bool`

HasTrigger returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *KalturaStorageProfile) GetUpdatedAt() int32`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *KalturaStorageProfile) GetUpdatedAtOk() (*int32, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *KalturaStorageProfile) SetUpdatedAt(v int32)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *KalturaStorageProfile) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


