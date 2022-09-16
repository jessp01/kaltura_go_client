# KalturaConfMaps

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChangeDescription** | Pointer to **string** |  | [optional] 
**Content** | Pointer to **string** | Ini file content | [optional] 
**CreatedAt** | Pointer to **int32** | &#x60;readOnly&#x60;  Time of the last update | [optional] [readonly] 
**IsEditable** | Pointer to **bool** | &#x60;readOnly&#x60;  IsEditable - true / false | [optional] [readonly] 
**Name** | Pointer to **string** | &#x60;insertOnly&#x60;  Name of the map | [optional] 
**RawData** | Pointer to **string** |  | [optional] 
**RelatedHost** | Pointer to **string** | Regex that represent the host/s that this map affect | [optional] 
**Remarks** | Pointer to **string** | &#x60;insertOnly&#x60; | [optional] 
**SourceLocation** | Pointer to **string** | &#x60;insertOnly&#x60;  Enum Type: &#x60;KalturaConfMapsSourceLocation&#x60; | [optional] 
**Status** | Pointer to **int32** | Enum Type: &#x60;KalturaConfMapsStatus&#x60;  map status | [optional] 
**UserId** | Pointer to **string** |  | [optional] 
**Version** | Pointer to **int32** | &#x60;readOnly&#x60; | [optional] [readonly] 

## Methods

### NewKalturaConfMaps

`func NewKalturaConfMaps() *KalturaConfMaps`

NewKalturaConfMaps instantiates a new KalturaConfMaps object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKalturaConfMapsWithDefaults

`func NewKalturaConfMapsWithDefaults() *KalturaConfMaps`

NewKalturaConfMapsWithDefaults instantiates a new KalturaConfMaps object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChangeDescription

`func (o *KalturaConfMaps) GetChangeDescription() string`

GetChangeDescription returns the ChangeDescription field if non-nil, zero value otherwise.

### GetChangeDescriptionOk

`func (o *KalturaConfMaps) GetChangeDescriptionOk() (*string, bool)`

GetChangeDescriptionOk returns a tuple with the ChangeDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangeDescription

`func (o *KalturaConfMaps) SetChangeDescription(v string)`

SetChangeDescription sets ChangeDescription field to given value.

### HasChangeDescription

`func (o *KalturaConfMaps) HasChangeDescription() bool`

HasChangeDescription returns a boolean if a field has been set.

### GetContent

`func (o *KalturaConfMaps) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *KalturaConfMaps) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *KalturaConfMaps) SetContent(v string)`

SetContent sets Content field to given value.

### HasContent

`func (o *KalturaConfMaps) HasContent() bool`

HasContent returns a boolean if a field has been set.

### GetCreatedAt

`func (o *KalturaConfMaps) GetCreatedAt() int32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *KalturaConfMaps) GetCreatedAtOk() (*int32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *KalturaConfMaps) SetCreatedAt(v int32)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *KalturaConfMaps) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetIsEditable

`func (o *KalturaConfMaps) GetIsEditable() bool`

GetIsEditable returns the IsEditable field if non-nil, zero value otherwise.

### GetIsEditableOk

`func (o *KalturaConfMaps) GetIsEditableOk() (*bool, bool)`

GetIsEditableOk returns a tuple with the IsEditable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEditable

`func (o *KalturaConfMaps) SetIsEditable(v bool)`

SetIsEditable sets IsEditable field to given value.

### HasIsEditable

`func (o *KalturaConfMaps) HasIsEditable() bool`

HasIsEditable returns a boolean if a field has been set.

### GetName

`func (o *KalturaConfMaps) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *KalturaConfMaps) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *KalturaConfMaps) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *KalturaConfMaps) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRawData

`func (o *KalturaConfMaps) GetRawData() string`

GetRawData returns the RawData field if non-nil, zero value otherwise.

### GetRawDataOk

`func (o *KalturaConfMaps) GetRawDataOk() (*string, bool)`

GetRawDataOk returns a tuple with the RawData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRawData

`func (o *KalturaConfMaps) SetRawData(v string)`

SetRawData sets RawData field to given value.

### HasRawData

`func (o *KalturaConfMaps) HasRawData() bool`

HasRawData returns a boolean if a field has been set.

### GetRelatedHost

`func (o *KalturaConfMaps) GetRelatedHost() string`

GetRelatedHost returns the RelatedHost field if non-nil, zero value otherwise.

### GetRelatedHostOk

`func (o *KalturaConfMaps) GetRelatedHostOk() (*string, bool)`

GetRelatedHostOk returns a tuple with the RelatedHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelatedHost

`func (o *KalturaConfMaps) SetRelatedHost(v string)`

SetRelatedHost sets RelatedHost field to given value.

### HasRelatedHost

`func (o *KalturaConfMaps) HasRelatedHost() bool`

HasRelatedHost returns a boolean if a field has been set.

### GetRemarks

`func (o *KalturaConfMaps) GetRemarks() string`

GetRemarks returns the Remarks field if non-nil, zero value otherwise.

### GetRemarksOk

`func (o *KalturaConfMaps) GetRemarksOk() (*string, bool)`

GetRemarksOk returns a tuple with the Remarks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemarks

`func (o *KalturaConfMaps) SetRemarks(v string)`

SetRemarks sets Remarks field to given value.

### HasRemarks

`func (o *KalturaConfMaps) HasRemarks() bool`

HasRemarks returns a boolean if a field has been set.

### GetSourceLocation

`func (o *KalturaConfMaps) GetSourceLocation() string`

GetSourceLocation returns the SourceLocation field if non-nil, zero value otherwise.

### GetSourceLocationOk

`func (o *KalturaConfMaps) GetSourceLocationOk() (*string, bool)`

GetSourceLocationOk returns a tuple with the SourceLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceLocation

`func (o *KalturaConfMaps) SetSourceLocation(v string)`

SetSourceLocation sets SourceLocation field to given value.

### HasSourceLocation

`func (o *KalturaConfMaps) HasSourceLocation() bool`

HasSourceLocation returns a boolean if a field has been set.

### GetStatus

`func (o *KalturaConfMaps) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *KalturaConfMaps) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *KalturaConfMaps) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *KalturaConfMaps) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUserId

`func (o *KalturaConfMaps) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *KalturaConfMaps) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *KalturaConfMaps) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *KalturaConfMaps) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetVersion

`func (o *KalturaConfMaps) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *KalturaConfMaps) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *KalturaConfMaps) SetVersion(v int32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *KalturaConfMaps) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


