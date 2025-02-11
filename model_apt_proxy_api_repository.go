/*
Nexus Repository Manager REST API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 3.76.0-03
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package nexus

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the AptProxyApiRepository type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AptProxyApiRepository{}

// AptProxyApiRepository struct for AptProxyApiRepository
type AptProxyApiRepository struct {
	// A unique identifier for this repository
	Name *string `json:"name,omitempty" validate:"regexp=^[a-zA-Z0-9\\\\-]{1}[a-zA-Z0-9_\\\\-\\\\.]*$"`
	// Whether this repository accepts incoming requests
	Online bool `json:"online"`
	Storage StorageAttributes `json:"storage"`
	Cleanup *CleanupPolicyAttributes `json:"cleanup,omitempty"`
	Proxy ProxyAttributes `json:"proxy"`
	NegativeCache NegativeCacheAttributes `json:"negativeCache"`
	HttpClient HttpClientAttributes `json:"httpClient"`
	// The name of the routing rule assigned to this repository
	RoutingRuleName *string `json:"routingRuleName,omitempty"`
	Replication *ReplicationAttributes `json:"replication,omitempty"`
	Apt AptProxyRepositoriesAttributes `json:"apt"`
}

type _AptProxyApiRepository AptProxyApiRepository

// NewAptProxyApiRepository instantiates a new AptProxyApiRepository object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAptProxyApiRepository(online bool, storage StorageAttributes, proxy ProxyAttributes, negativeCache NegativeCacheAttributes, httpClient HttpClientAttributes, apt AptProxyRepositoriesAttributes) *AptProxyApiRepository {
	this := AptProxyApiRepository{}
	this.Online = online
	this.Storage = storage
	this.Proxy = proxy
	this.NegativeCache = negativeCache
	this.HttpClient = httpClient
	this.Apt = apt
	return &this
}

// NewAptProxyApiRepositoryWithDefaults instantiates a new AptProxyApiRepository object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAptProxyApiRepositoryWithDefaults() *AptProxyApiRepository {
	this := AptProxyApiRepository{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *AptProxyApiRepository) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AptProxyApiRepository) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *AptProxyApiRepository) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *AptProxyApiRepository) SetName(v string) {
	o.Name = &v
}

// GetOnline returns the Online field value
func (o *AptProxyApiRepository) GetOnline() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Online
}

// GetOnlineOk returns a tuple with the Online field value
// and a boolean to check if the value has been set.
func (o *AptProxyApiRepository) GetOnlineOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Online, true
}

// SetOnline sets field value
func (o *AptProxyApiRepository) SetOnline(v bool) {
	o.Online = v
}

// GetStorage returns the Storage field value
func (o *AptProxyApiRepository) GetStorage() StorageAttributes {
	if o == nil {
		var ret StorageAttributes
		return ret
	}

	return o.Storage
}

// GetStorageOk returns a tuple with the Storage field value
// and a boolean to check if the value has been set.
func (o *AptProxyApiRepository) GetStorageOk() (*StorageAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Storage, true
}

// SetStorage sets field value
func (o *AptProxyApiRepository) SetStorage(v StorageAttributes) {
	o.Storage = v
}

// GetCleanup returns the Cleanup field value if set, zero value otherwise.
func (o *AptProxyApiRepository) GetCleanup() CleanupPolicyAttributes {
	if o == nil || IsNil(o.Cleanup) {
		var ret CleanupPolicyAttributes
		return ret
	}
	return *o.Cleanup
}

// GetCleanupOk returns a tuple with the Cleanup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AptProxyApiRepository) GetCleanupOk() (*CleanupPolicyAttributes, bool) {
	if o == nil || IsNil(o.Cleanup) {
		return nil, false
	}
	return o.Cleanup, true
}

// HasCleanup returns a boolean if a field has been set.
func (o *AptProxyApiRepository) HasCleanup() bool {
	if o != nil && !IsNil(o.Cleanup) {
		return true
	}

	return false
}

// SetCleanup gets a reference to the given CleanupPolicyAttributes and assigns it to the Cleanup field.
func (o *AptProxyApiRepository) SetCleanup(v CleanupPolicyAttributes) {
	o.Cleanup = &v
}

// GetProxy returns the Proxy field value
func (o *AptProxyApiRepository) GetProxy() ProxyAttributes {
	if o == nil {
		var ret ProxyAttributes
		return ret
	}

	return o.Proxy
}

// GetProxyOk returns a tuple with the Proxy field value
// and a boolean to check if the value has been set.
func (o *AptProxyApiRepository) GetProxyOk() (*ProxyAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Proxy, true
}

// SetProxy sets field value
func (o *AptProxyApiRepository) SetProxy(v ProxyAttributes) {
	o.Proxy = v
}

// GetNegativeCache returns the NegativeCache field value
func (o *AptProxyApiRepository) GetNegativeCache() NegativeCacheAttributes {
	if o == nil {
		var ret NegativeCacheAttributes
		return ret
	}

	return o.NegativeCache
}

// GetNegativeCacheOk returns a tuple with the NegativeCache field value
// and a boolean to check if the value has been set.
func (o *AptProxyApiRepository) GetNegativeCacheOk() (*NegativeCacheAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NegativeCache, true
}

// SetNegativeCache sets field value
func (o *AptProxyApiRepository) SetNegativeCache(v NegativeCacheAttributes) {
	o.NegativeCache = v
}

// GetHttpClient returns the HttpClient field value
func (o *AptProxyApiRepository) GetHttpClient() HttpClientAttributes {
	if o == nil {
		var ret HttpClientAttributes
		return ret
	}

	return o.HttpClient
}

// GetHttpClientOk returns a tuple with the HttpClient field value
// and a boolean to check if the value has been set.
func (o *AptProxyApiRepository) GetHttpClientOk() (*HttpClientAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HttpClient, true
}

// SetHttpClient sets field value
func (o *AptProxyApiRepository) SetHttpClient(v HttpClientAttributes) {
	o.HttpClient = v
}

// GetRoutingRuleName returns the RoutingRuleName field value if set, zero value otherwise.
func (o *AptProxyApiRepository) GetRoutingRuleName() string {
	if o == nil || IsNil(o.RoutingRuleName) {
		var ret string
		return ret
	}
	return *o.RoutingRuleName
}

// GetRoutingRuleNameOk returns a tuple with the RoutingRuleName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AptProxyApiRepository) GetRoutingRuleNameOk() (*string, bool) {
	if o == nil || IsNil(o.RoutingRuleName) {
		return nil, false
	}
	return o.RoutingRuleName, true
}

// HasRoutingRuleName returns a boolean if a field has been set.
func (o *AptProxyApiRepository) HasRoutingRuleName() bool {
	if o != nil && !IsNil(o.RoutingRuleName) {
		return true
	}

	return false
}

// SetRoutingRuleName gets a reference to the given string and assigns it to the RoutingRuleName field.
func (o *AptProxyApiRepository) SetRoutingRuleName(v string) {
	o.RoutingRuleName = &v
}

// GetReplication returns the Replication field value if set, zero value otherwise.
func (o *AptProxyApiRepository) GetReplication() ReplicationAttributes {
	if o == nil || IsNil(o.Replication) {
		var ret ReplicationAttributes
		return ret
	}
	return *o.Replication
}

// GetReplicationOk returns a tuple with the Replication field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AptProxyApiRepository) GetReplicationOk() (*ReplicationAttributes, bool) {
	if o == nil || IsNil(o.Replication) {
		return nil, false
	}
	return o.Replication, true
}

// HasReplication returns a boolean if a field has been set.
func (o *AptProxyApiRepository) HasReplication() bool {
	if o != nil && !IsNil(o.Replication) {
		return true
	}

	return false
}

// SetReplication gets a reference to the given ReplicationAttributes and assigns it to the Replication field.
func (o *AptProxyApiRepository) SetReplication(v ReplicationAttributes) {
	o.Replication = &v
}

// GetApt returns the Apt field value
func (o *AptProxyApiRepository) GetApt() AptProxyRepositoriesAttributes {
	if o == nil {
		var ret AptProxyRepositoriesAttributes
		return ret
	}

	return o.Apt
}

// GetAptOk returns a tuple with the Apt field value
// and a boolean to check if the value has been set.
func (o *AptProxyApiRepository) GetAptOk() (*AptProxyRepositoriesAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Apt, true
}

// SetApt sets field value
func (o *AptProxyApiRepository) SetApt(v AptProxyRepositoriesAttributes) {
	o.Apt = v
}

func (o AptProxyApiRepository) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AptProxyApiRepository) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	toSerialize["online"] = o.Online
	toSerialize["storage"] = o.Storage
	if !IsNil(o.Cleanup) {
		toSerialize["cleanup"] = o.Cleanup
	}
	toSerialize["proxy"] = o.Proxy
	toSerialize["negativeCache"] = o.NegativeCache
	toSerialize["httpClient"] = o.HttpClient
	if !IsNil(o.RoutingRuleName) {
		toSerialize["routingRuleName"] = o.RoutingRuleName
	}
	if !IsNil(o.Replication) {
		toSerialize["replication"] = o.Replication
	}
	toSerialize["apt"] = o.Apt
	return toSerialize, nil
}

func (o *AptProxyApiRepository) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"online",
		"storage",
		"proxy",
		"negativeCache",
		"httpClient",
		"apt",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varAptProxyApiRepository := _AptProxyApiRepository{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAptProxyApiRepository)

	if err != nil {
		return err
	}

	*o = AptProxyApiRepository(varAptProxyApiRepository)

	return err
}

type NullableAptProxyApiRepository struct {
	value *AptProxyApiRepository
	isSet bool
}

func (v NullableAptProxyApiRepository) Get() *AptProxyApiRepository {
	return v.value
}

func (v *NullableAptProxyApiRepository) Set(val *AptProxyApiRepository) {
	v.value = val
	v.isSet = true
}

func (v NullableAptProxyApiRepository) IsSet() bool {
	return v.isSet
}

func (v *NullableAptProxyApiRepository) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAptProxyApiRepository(val *AptProxyApiRepository) *NullableAptProxyApiRepository {
	return &NullableAptProxyApiRepository{value: val, isSet: true}
}

func (v NullableAptProxyApiRepository) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAptProxyApiRepository) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


