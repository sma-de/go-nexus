# CondaProxyRepositoryApiRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | A unique identifier for this repository | 
**Online** | **bool** | Whether this repository accepts incoming requests | 
**Storage** | [**StorageAttributes**](StorageAttributes.md) |  | 
**Cleanup** | Pointer to [**CleanupPolicyAttributes**](CleanupPolicyAttributes.md) |  | [optional] 
**Proxy** | [**ProxyAttributes**](ProxyAttributes.md) |  | 
**NegativeCache** | [**NegativeCacheAttributes**](NegativeCacheAttributes.md) |  | 
**HttpClient** | [**HttpClientAttributes**](HttpClientAttributes.md) |  | 
**RoutingRule** | Pointer to **string** |  | [optional] 
**Replication** | Pointer to [**ReplicationAttributes**](ReplicationAttributes.md) |  | [optional] 

## Methods

### NewCondaProxyRepositoryApiRequest

`func NewCondaProxyRepositoryApiRequest(name string, online bool, storage StorageAttributes, proxy ProxyAttributes, negativeCache NegativeCacheAttributes, httpClient HttpClientAttributes, ) *CondaProxyRepositoryApiRequest`

NewCondaProxyRepositoryApiRequest instantiates a new CondaProxyRepositoryApiRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCondaProxyRepositoryApiRequestWithDefaults

`func NewCondaProxyRepositoryApiRequestWithDefaults() *CondaProxyRepositoryApiRequest`

NewCondaProxyRepositoryApiRequestWithDefaults instantiates a new CondaProxyRepositoryApiRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CondaProxyRepositoryApiRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CondaProxyRepositoryApiRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CondaProxyRepositoryApiRequest) SetName(v string)`

SetName sets Name field to given value.


### GetOnline

`func (o *CondaProxyRepositoryApiRequest) GetOnline() bool`

GetOnline returns the Online field if non-nil, zero value otherwise.

### GetOnlineOk

`func (o *CondaProxyRepositoryApiRequest) GetOnlineOk() (*bool, bool)`

GetOnlineOk returns a tuple with the Online field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnline

`func (o *CondaProxyRepositoryApiRequest) SetOnline(v bool)`

SetOnline sets Online field to given value.


### GetStorage

`func (o *CondaProxyRepositoryApiRequest) GetStorage() StorageAttributes`

GetStorage returns the Storage field if non-nil, zero value otherwise.

### GetStorageOk

`func (o *CondaProxyRepositoryApiRequest) GetStorageOk() (*StorageAttributes, bool)`

GetStorageOk returns a tuple with the Storage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorage

`func (o *CondaProxyRepositoryApiRequest) SetStorage(v StorageAttributes)`

SetStorage sets Storage field to given value.


### GetCleanup

`func (o *CondaProxyRepositoryApiRequest) GetCleanup() CleanupPolicyAttributes`

GetCleanup returns the Cleanup field if non-nil, zero value otherwise.

### GetCleanupOk

`func (o *CondaProxyRepositoryApiRequest) GetCleanupOk() (*CleanupPolicyAttributes, bool)`

GetCleanupOk returns a tuple with the Cleanup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCleanup

`func (o *CondaProxyRepositoryApiRequest) SetCleanup(v CleanupPolicyAttributes)`

SetCleanup sets Cleanup field to given value.

### HasCleanup

`func (o *CondaProxyRepositoryApiRequest) HasCleanup() bool`

HasCleanup returns a boolean if a field has been set.

### GetProxy

`func (o *CondaProxyRepositoryApiRequest) GetProxy() ProxyAttributes`

GetProxy returns the Proxy field if non-nil, zero value otherwise.

### GetProxyOk

`func (o *CondaProxyRepositoryApiRequest) GetProxyOk() (*ProxyAttributes, bool)`

GetProxyOk returns a tuple with the Proxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxy

`func (o *CondaProxyRepositoryApiRequest) SetProxy(v ProxyAttributes)`

SetProxy sets Proxy field to given value.


### GetNegativeCache

`func (o *CondaProxyRepositoryApiRequest) GetNegativeCache() NegativeCacheAttributes`

GetNegativeCache returns the NegativeCache field if non-nil, zero value otherwise.

### GetNegativeCacheOk

`func (o *CondaProxyRepositoryApiRequest) GetNegativeCacheOk() (*NegativeCacheAttributes, bool)`

GetNegativeCacheOk returns a tuple with the NegativeCache field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNegativeCache

`func (o *CondaProxyRepositoryApiRequest) SetNegativeCache(v NegativeCacheAttributes)`

SetNegativeCache sets NegativeCache field to given value.


### GetHttpClient

`func (o *CondaProxyRepositoryApiRequest) GetHttpClient() HttpClientAttributes`

GetHttpClient returns the HttpClient field if non-nil, zero value otherwise.

### GetHttpClientOk

`func (o *CondaProxyRepositoryApiRequest) GetHttpClientOk() (*HttpClientAttributes, bool)`

GetHttpClientOk returns a tuple with the HttpClient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpClient

`func (o *CondaProxyRepositoryApiRequest) SetHttpClient(v HttpClientAttributes)`

SetHttpClient sets HttpClient field to given value.


### GetRoutingRule

`func (o *CondaProxyRepositoryApiRequest) GetRoutingRule() string`

GetRoutingRule returns the RoutingRule field if non-nil, zero value otherwise.

### GetRoutingRuleOk

`func (o *CondaProxyRepositoryApiRequest) GetRoutingRuleOk() (*string, bool)`

GetRoutingRuleOk returns a tuple with the RoutingRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoutingRule

`func (o *CondaProxyRepositoryApiRequest) SetRoutingRule(v string)`

SetRoutingRule sets RoutingRule field to given value.

### HasRoutingRule

`func (o *CondaProxyRepositoryApiRequest) HasRoutingRule() bool`

HasRoutingRule returns a boolean if a field has been set.

### GetReplication

`func (o *CondaProxyRepositoryApiRequest) GetReplication() ReplicationAttributes`

GetReplication returns the Replication field if non-nil, zero value otherwise.

### GetReplicationOk

`func (o *CondaProxyRepositoryApiRequest) GetReplicationOk() (*ReplicationAttributes, bool)`

GetReplicationOk returns a tuple with the Replication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplication

`func (o *CondaProxyRepositoryApiRequest) SetReplication(v ReplicationAttributes)`

SetReplication sets Replication field to given value.

### HasReplication

`func (o *CondaProxyRepositoryApiRequest) HasReplication() bool`

HasReplication returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


