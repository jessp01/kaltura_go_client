# KalturaBulkUpload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BulkFileUrl** | Pointer to **string** |  | [optional] 
**BulkUploadObjectType** | Pointer to **string** | Enum Type: &#x60;KalturaBulkUploadObjectType&#x60; | [optional] 
**BulkUploadType** | Pointer to **string** | Enum Type: &#x60;KalturaBulkUploadType&#x60; | [optional] 
**CsvFileUrl** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Error** | Pointer to **string** |  | [optional] 
**ErrorNumber** | Pointer to **int32** |  | [optional] 
**ErrorType** | Pointer to **int32** | Enum Type: &#x60;KalturaBatchJobErrorTypes&#x60; | [optional] 
**FileName** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **int32** |  | [optional] 
**LogFileUrl** | Pointer to **string** |  | [optional] 
**NumOfEntries** | Pointer to **int32** |  | [optional] 
**NumOfObjects** | Pointer to **int32** |  | [optional] 
**Results** | Pointer to [**[]KalturaBulkUploadResult**](KalturaBulkUploadResult.md) |  | [optional] 
**Status** | Pointer to **int32** | Enum Type: &#x60;KalturaBatchJobStatus&#x60; | [optional] 
**UploadedBy** | Pointer to **string** |  | [optional] 
**UploadedByUserId** | Pointer to **string** |  | [optional] 
**UploadedOn** | Pointer to **int32** |  | [optional] 

## Methods

### NewKalturaBulkUpload

`func NewKalturaBulkUpload() *KalturaBulkUpload`

NewKalturaBulkUpload instantiates a new KalturaBulkUpload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKalturaBulkUploadWithDefaults

`func NewKalturaBulkUploadWithDefaults() *KalturaBulkUpload`

NewKalturaBulkUploadWithDefaults instantiates a new KalturaBulkUpload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBulkFileUrl

`func (o *KalturaBulkUpload) GetBulkFileUrl() string`

GetBulkFileUrl returns the BulkFileUrl field if non-nil, zero value otherwise.

### GetBulkFileUrlOk

`func (o *KalturaBulkUpload) GetBulkFileUrlOk() (*string, bool)`

GetBulkFileUrlOk returns a tuple with the BulkFileUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBulkFileUrl

`func (o *KalturaBulkUpload) SetBulkFileUrl(v string)`

SetBulkFileUrl sets BulkFileUrl field to given value.

### HasBulkFileUrl

`func (o *KalturaBulkUpload) HasBulkFileUrl() bool`

HasBulkFileUrl returns a boolean if a field has been set.

### GetBulkUploadObjectType

`func (o *KalturaBulkUpload) GetBulkUploadObjectType() string`

GetBulkUploadObjectType returns the BulkUploadObjectType field if non-nil, zero value otherwise.

### GetBulkUploadObjectTypeOk

`func (o *KalturaBulkUpload) GetBulkUploadObjectTypeOk() (*string, bool)`

GetBulkUploadObjectTypeOk returns a tuple with the BulkUploadObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBulkUploadObjectType

`func (o *KalturaBulkUpload) SetBulkUploadObjectType(v string)`

SetBulkUploadObjectType sets BulkUploadObjectType field to given value.

### HasBulkUploadObjectType

`func (o *KalturaBulkUpload) HasBulkUploadObjectType() bool`

HasBulkUploadObjectType returns a boolean if a field has been set.

### GetBulkUploadType

`func (o *KalturaBulkUpload) GetBulkUploadType() string`

GetBulkUploadType returns the BulkUploadType field if non-nil, zero value otherwise.

### GetBulkUploadTypeOk

`func (o *KalturaBulkUpload) GetBulkUploadTypeOk() (*string, bool)`

GetBulkUploadTypeOk returns a tuple with the BulkUploadType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBulkUploadType

`func (o *KalturaBulkUpload) SetBulkUploadType(v string)`

SetBulkUploadType sets BulkUploadType field to given value.

### HasBulkUploadType

`func (o *KalturaBulkUpload) HasBulkUploadType() bool`

HasBulkUploadType returns a boolean if a field has been set.

### GetCsvFileUrl

`func (o *KalturaBulkUpload) GetCsvFileUrl() string`

GetCsvFileUrl returns the CsvFileUrl field if non-nil, zero value otherwise.

### GetCsvFileUrlOk

`func (o *KalturaBulkUpload) GetCsvFileUrlOk() (*string, bool)`

GetCsvFileUrlOk returns a tuple with the CsvFileUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCsvFileUrl

`func (o *KalturaBulkUpload) SetCsvFileUrl(v string)`

SetCsvFileUrl sets CsvFileUrl field to given value.

### HasCsvFileUrl

`func (o *KalturaBulkUpload) HasCsvFileUrl() bool`

HasCsvFileUrl returns a boolean if a field has been set.

### GetDescription

`func (o *KalturaBulkUpload) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *KalturaBulkUpload) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *KalturaBulkUpload) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *KalturaBulkUpload) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetError

`func (o *KalturaBulkUpload) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *KalturaBulkUpload) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *KalturaBulkUpload) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *KalturaBulkUpload) HasError() bool`

HasError returns a boolean if a field has been set.

### GetErrorNumber

`func (o *KalturaBulkUpload) GetErrorNumber() int32`

GetErrorNumber returns the ErrorNumber field if non-nil, zero value otherwise.

### GetErrorNumberOk

`func (o *KalturaBulkUpload) GetErrorNumberOk() (*int32, bool)`

GetErrorNumberOk returns a tuple with the ErrorNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorNumber

`func (o *KalturaBulkUpload) SetErrorNumber(v int32)`

SetErrorNumber sets ErrorNumber field to given value.

### HasErrorNumber

`func (o *KalturaBulkUpload) HasErrorNumber() bool`

HasErrorNumber returns a boolean if a field has been set.

### GetErrorType

`func (o *KalturaBulkUpload) GetErrorType() int32`

GetErrorType returns the ErrorType field if non-nil, zero value otherwise.

### GetErrorTypeOk

`func (o *KalturaBulkUpload) GetErrorTypeOk() (*int32, bool)`

GetErrorTypeOk returns a tuple with the ErrorType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorType

`func (o *KalturaBulkUpload) SetErrorType(v int32)`

SetErrorType sets ErrorType field to given value.

### HasErrorType

`func (o *KalturaBulkUpload) HasErrorType() bool`

HasErrorType returns a boolean if a field has been set.

### GetFileName

`func (o *KalturaBulkUpload) GetFileName() string`

GetFileName returns the FileName field if non-nil, zero value otherwise.

### GetFileNameOk

`func (o *KalturaBulkUpload) GetFileNameOk() (*string, bool)`

GetFileNameOk returns a tuple with the FileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileName

`func (o *KalturaBulkUpload) SetFileName(v string)`

SetFileName sets FileName field to given value.

### HasFileName

`func (o *KalturaBulkUpload) HasFileName() bool`

HasFileName returns a boolean if a field has been set.

### GetId

`func (o *KalturaBulkUpload) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *KalturaBulkUpload) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *KalturaBulkUpload) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *KalturaBulkUpload) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLogFileUrl

`func (o *KalturaBulkUpload) GetLogFileUrl() string`

GetLogFileUrl returns the LogFileUrl field if non-nil, zero value otherwise.

### GetLogFileUrlOk

`func (o *KalturaBulkUpload) GetLogFileUrlOk() (*string, bool)`

GetLogFileUrlOk returns a tuple with the LogFileUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogFileUrl

`func (o *KalturaBulkUpload) SetLogFileUrl(v string)`

SetLogFileUrl sets LogFileUrl field to given value.

### HasLogFileUrl

`func (o *KalturaBulkUpload) HasLogFileUrl() bool`

HasLogFileUrl returns a boolean if a field has been set.

### GetNumOfEntries

`func (o *KalturaBulkUpload) GetNumOfEntries() int32`

GetNumOfEntries returns the NumOfEntries field if non-nil, zero value otherwise.

### GetNumOfEntriesOk

`func (o *KalturaBulkUpload) GetNumOfEntriesOk() (*int32, bool)`

GetNumOfEntriesOk returns a tuple with the NumOfEntries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumOfEntries

`func (o *KalturaBulkUpload) SetNumOfEntries(v int32)`

SetNumOfEntries sets NumOfEntries field to given value.

### HasNumOfEntries

`func (o *KalturaBulkUpload) HasNumOfEntries() bool`

HasNumOfEntries returns a boolean if a field has been set.

### GetNumOfObjects

`func (o *KalturaBulkUpload) GetNumOfObjects() int32`

GetNumOfObjects returns the NumOfObjects field if non-nil, zero value otherwise.

### GetNumOfObjectsOk

`func (o *KalturaBulkUpload) GetNumOfObjectsOk() (*int32, bool)`

GetNumOfObjectsOk returns a tuple with the NumOfObjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumOfObjects

`func (o *KalturaBulkUpload) SetNumOfObjects(v int32)`

SetNumOfObjects sets NumOfObjects field to given value.

### HasNumOfObjects

`func (o *KalturaBulkUpload) HasNumOfObjects() bool`

HasNumOfObjects returns a boolean if a field has been set.

### GetResults

`func (o *KalturaBulkUpload) GetResults() []KalturaBulkUploadResult`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *KalturaBulkUpload) GetResultsOk() (*[]KalturaBulkUploadResult, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *KalturaBulkUpload) SetResults(v []KalturaBulkUploadResult)`

SetResults sets Results field to given value.

### HasResults

`func (o *KalturaBulkUpload) HasResults() bool`

HasResults returns a boolean if a field has been set.

### GetStatus

`func (o *KalturaBulkUpload) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *KalturaBulkUpload) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *KalturaBulkUpload) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *KalturaBulkUpload) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUploadedBy

`func (o *KalturaBulkUpload) GetUploadedBy() string`

GetUploadedBy returns the UploadedBy field if non-nil, zero value otherwise.

### GetUploadedByOk

`func (o *KalturaBulkUpload) GetUploadedByOk() (*string, bool)`

GetUploadedByOk returns a tuple with the UploadedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUploadedBy

`func (o *KalturaBulkUpload) SetUploadedBy(v string)`

SetUploadedBy sets UploadedBy field to given value.

### HasUploadedBy

`func (o *KalturaBulkUpload) HasUploadedBy() bool`

HasUploadedBy returns a boolean if a field has been set.

### GetUploadedByUserId

`func (o *KalturaBulkUpload) GetUploadedByUserId() string`

GetUploadedByUserId returns the UploadedByUserId field if non-nil, zero value otherwise.

### GetUploadedByUserIdOk

`func (o *KalturaBulkUpload) GetUploadedByUserIdOk() (*string, bool)`

GetUploadedByUserIdOk returns a tuple with the UploadedByUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUploadedByUserId

`func (o *KalturaBulkUpload) SetUploadedByUserId(v string)`

SetUploadedByUserId sets UploadedByUserId field to given value.

### HasUploadedByUserId

`func (o *KalturaBulkUpload) HasUploadedByUserId() bool`

HasUploadedByUserId returns a boolean if a field has been set.

### GetUploadedOn

`func (o *KalturaBulkUpload) GetUploadedOn() int32`

GetUploadedOn returns the UploadedOn field if non-nil, zero value otherwise.

### GetUploadedOnOk

`func (o *KalturaBulkUpload) GetUploadedOnOk() (*int32, bool)`

GetUploadedOnOk returns a tuple with the UploadedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUploadedOn

`func (o *KalturaBulkUpload) SetUploadedOn(v int32)`

SetUploadedOn sets UploadedOn field to given value.

### HasUploadedOn

`func (o *KalturaBulkUpload) HasUploadedOn() bool`

HasUploadedOn returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


