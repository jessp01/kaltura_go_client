# KalturaPartnerStatistics

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Bandwidth** | Pointer to **float32** | &#x60;readOnly&#x60;  Partner total bandwidth in GB | [optional] [readonly] 
**Hosting** | Pointer to **float32** | &#x60;readOnly&#x60;  Partner total hosting in GB on the disk | [optional] [readonly] 
**PackageBandwidthAndStorage** | Pointer to **int32** | &#x60;readOnly&#x60;  Package total allowed bandwidth and storage | [optional] [readonly] 
**ReachedLimitDate** | Pointer to **int32** | &#x60;readOnly&#x60;  date when partner reached the limit of his package (timestamp) | [optional] [readonly] 
**Usage** | Pointer to **int32** | &#x60;readOnly&#x60;  total usage in GB - including bandwidth and storage | [optional] [readonly] 
**UsagePercent** | Pointer to **float32** | &#x60;readOnly&#x60;  Percent of usage out of partner&#39;s package. if usage is 5GB and package is 10GB, this value will be 50 | [optional] [readonly] 

## Methods

### NewKalturaPartnerStatistics

`func NewKalturaPartnerStatistics() *KalturaPartnerStatistics`

NewKalturaPartnerStatistics instantiates a new KalturaPartnerStatistics object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKalturaPartnerStatisticsWithDefaults

`func NewKalturaPartnerStatisticsWithDefaults() *KalturaPartnerStatistics`

NewKalturaPartnerStatisticsWithDefaults instantiates a new KalturaPartnerStatistics object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBandwidth

`func (o *KalturaPartnerStatistics) GetBandwidth() float32`

GetBandwidth returns the Bandwidth field if non-nil, zero value otherwise.

### GetBandwidthOk

`func (o *KalturaPartnerStatistics) GetBandwidthOk() (*float32, bool)`

GetBandwidthOk returns a tuple with the Bandwidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBandwidth

`func (o *KalturaPartnerStatistics) SetBandwidth(v float32)`

SetBandwidth sets Bandwidth field to given value.

### HasBandwidth

`func (o *KalturaPartnerStatistics) HasBandwidth() bool`

HasBandwidth returns a boolean if a field has been set.

### GetHosting

`func (o *KalturaPartnerStatistics) GetHosting() float32`

GetHosting returns the Hosting field if non-nil, zero value otherwise.

### GetHostingOk

`func (o *KalturaPartnerStatistics) GetHostingOk() (*float32, bool)`

GetHostingOk returns a tuple with the Hosting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHosting

`func (o *KalturaPartnerStatistics) SetHosting(v float32)`

SetHosting sets Hosting field to given value.

### HasHosting

`func (o *KalturaPartnerStatistics) HasHosting() bool`

HasHosting returns a boolean if a field has been set.

### GetPackageBandwidthAndStorage

`func (o *KalturaPartnerStatistics) GetPackageBandwidthAndStorage() int32`

GetPackageBandwidthAndStorage returns the PackageBandwidthAndStorage field if non-nil, zero value otherwise.

### GetPackageBandwidthAndStorageOk

`func (o *KalturaPartnerStatistics) GetPackageBandwidthAndStorageOk() (*int32, bool)`

GetPackageBandwidthAndStorageOk returns a tuple with the PackageBandwidthAndStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageBandwidthAndStorage

`func (o *KalturaPartnerStatistics) SetPackageBandwidthAndStorage(v int32)`

SetPackageBandwidthAndStorage sets PackageBandwidthAndStorage field to given value.

### HasPackageBandwidthAndStorage

`func (o *KalturaPartnerStatistics) HasPackageBandwidthAndStorage() bool`

HasPackageBandwidthAndStorage returns a boolean if a field has been set.

### GetReachedLimitDate

`func (o *KalturaPartnerStatistics) GetReachedLimitDate() int32`

GetReachedLimitDate returns the ReachedLimitDate field if non-nil, zero value otherwise.

### GetReachedLimitDateOk

`func (o *KalturaPartnerStatistics) GetReachedLimitDateOk() (*int32, bool)`

GetReachedLimitDateOk returns a tuple with the ReachedLimitDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReachedLimitDate

`func (o *KalturaPartnerStatistics) SetReachedLimitDate(v int32)`

SetReachedLimitDate sets ReachedLimitDate field to given value.

### HasReachedLimitDate

`func (o *KalturaPartnerStatistics) HasReachedLimitDate() bool`

HasReachedLimitDate returns a boolean if a field has been set.

### GetUsage

`func (o *KalturaPartnerStatistics) GetUsage() int32`

GetUsage returns the Usage field if non-nil, zero value otherwise.

### GetUsageOk

`func (o *KalturaPartnerStatistics) GetUsageOk() (*int32, bool)`

GetUsageOk returns a tuple with the Usage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsage

`func (o *KalturaPartnerStatistics) SetUsage(v int32)`

SetUsage sets Usage field to given value.

### HasUsage

`func (o *KalturaPartnerStatistics) HasUsage() bool`

HasUsage returns a boolean if a field has been set.

### GetUsagePercent

`func (o *KalturaPartnerStatistics) GetUsagePercent() float32`

GetUsagePercent returns the UsagePercent field if non-nil, zero value otherwise.

### GetUsagePercentOk

`func (o *KalturaPartnerStatistics) GetUsagePercentOk() (*float32, bool)`

GetUsagePercentOk returns a tuple with the UsagePercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsagePercent

`func (o *KalturaPartnerStatistics) SetUsagePercent(v float32)`

SetUsagePercent sets UsagePercent field to given value.

### HasUsagePercent

`func (o *KalturaPartnerStatistics) HasUsagePercent() bool`

HasUsagePercent returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


