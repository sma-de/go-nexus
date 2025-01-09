# MaliciousRiskSummaryResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CountByEcosystem** | Pointer to [**[]MaliciousRiskCountByEcosystem**](MaliciousRiskCountByEcosystem.md) |  | [optional] 
**TotalMaliciousRiskCount** | Pointer to **int64** |  | [optional] 
**TotalProxyRepositoryCount** | Pointer to **int64** |  | [optional] 
**QuarantineEnabledRepositoryCount** | Pointer to **int64** |  | [optional] 
**HdsError** | Pointer to **bool** |  | [optional] 

## Methods

### NewMaliciousRiskSummaryResponse

`func NewMaliciousRiskSummaryResponse() *MaliciousRiskSummaryResponse`

NewMaliciousRiskSummaryResponse instantiates a new MaliciousRiskSummaryResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMaliciousRiskSummaryResponseWithDefaults

`func NewMaliciousRiskSummaryResponseWithDefaults() *MaliciousRiskSummaryResponse`

NewMaliciousRiskSummaryResponseWithDefaults instantiates a new MaliciousRiskSummaryResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCountByEcosystem

`func (o *MaliciousRiskSummaryResponse) GetCountByEcosystem() []MaliciousRiskCountByEcosystem`

GetCountByEcosystem returns the CountByEcosystem field if non-nil, zero value otherwise.

### GetCountByEcosystemOk

`func (o *MaliciousRiskSummaryResponse) GetCountByEcosystemOk() (*[]MaliciousRiskCountByEcosystem, bool)`

GetCountByEcosystemOk returns a tuple with the CountByEcosystem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountByEcosystem

`func (o *MaliciousRiskSummaryResponse) SetCountByEcosystem(v []MaliciousRiskCountByEcosystem)`

SetCountByEcosystem sets CountByEcosystem field to given value.

### HasCountByEcosystem

`func (o *MaliciousRiskSummaryResponse) HasCountByEcosystem() bool`

HasCountByEcosystem returns a boolean if a field has been set.

### GetTotalMaliciousRiskCount

`func (o *MaliciousRiskSummaryResponse) GetTotalMaliciousRiskCount() int64`

GetTotalMaliciousRiskCount returns the TotalMaliciousRiskCount field if non-nil, zero value otherwise.

### GetTotalMaliciousRiskCountOk

`func (o *MaliciousRiskSummaryResponse) GetTotalMaliciousRiskCountOk() (*int64, bool)`

GetTotalMaliciousRiskCountOk returns a tuple with the TotalMaliciousRiskCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalMaliciousRiskCount

`func (o *MaliciousRiskSummaryResponse) SetTotalMaliciousRiskCount(v int64)`

SetTotalMaliciousRiskCount sets TotalMaliciousRiskCount field to given value.

### HasTotalMaliciousRiskCount

`func (o *MaliciousRiskSummaryResponse) HasTotalMaliciousRiskCount() bool`

HasTotalMaliciousRiskCount returns a boolean if a field has been set.

### GetTotalProxyRepositoryCount

`func (o *MaliciousRiskSummaryResponse) GetTotalProxyRepositoryCount() int64`

GetTotalProxyRepositoryCount returns the TotalProxyRepositoryCount field if non-nil, zero value otherwise.

### GetTotalProxyRepositoryCountOk

`func (o *MaliciousRiskSummaryResponse) GetTotalProxyRepositoryCountOk() (*int64, bool)`

GetTotalProxyRepositoryCountOk returns a tuple with the TotalProxyRepositoryCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalProxyRepositoryCount

`func (o *MaliciousRiskSummaryResponse) SetTotalProxyRepositoryCount(v int64)`

SetTotalProxyRepositoryCount sets TotalProxyRepositoryCount field to given value.

### HasTotalProxyRepositoryCount

`func (o *MaliciousRiskSummaryResponse) HasTotalProxyRepositoryCount() bool`

HasTotalProxyRepositoryCount returns a boolean if a field has been set.

### GetQuarantineEnabledRepositoryCount

`func (o *MaliciousRiskSummaryResponse) GetQuarantineEnabledRepositoryCount() int64`

GetQuarantineEnabledRepositoryCount returns the QuarantineEnabledRepositoryCount field if non-nil, zero value otherwise.

### GetQuarantineEnabledRepositoryCountOk

`func (o *MaliciousRiskSummaryResponse) GetQuarantineEnabledRepositoryCountOk() (*int64, bool)`

GetQuarantineEnabledRepositoryCountOk returns a tuple with the QuarantineEnabledRepositoryCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuarantineEnabledRepositoryCount

`func (o *MaliciousRiskSummaryResponse) SetQuarantineEnabledRepositoryCount(v int64)`

SetQuarantineEnabledRepositoryCount sets QuarantineEnabledRepositoryCount field to given value.

### HasQuarantineEnabledRepositoryCount

`func (o *MaliciousRiskSummaryResponse) HasQuarantineEnabledRepositoryCount() bool`

HasQuarantineEnabledRepositoryCount returns a boolean if a field has been set.

### GetHdsError

`func (o *MaliciousRiskSummaryResponse) GetHdsError() bool`

GetHdsError returns the HdsError field if non-nil, zero value otherwise.

### GetHdsErrorOk

`func (o *MaliciousRiskSummaryResponse) GetHdsErrorOk() (*bool, bool)`

GetHdsErrorOk returns a tuple with the HdsError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHdsError

`func (o *MaliciousRiskSummaryResponse) SetHdsError(v bool)`

SetHdsError sets HdsError field to given value.

### HasHdsError

`func (o *MaliciousRiskSummaryResponse) HasHdsError() bool`

HasHdsError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


