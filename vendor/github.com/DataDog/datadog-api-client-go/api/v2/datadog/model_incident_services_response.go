/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 * This product includes software developed at Datadog (https://www.datadoghq.com/).
 * Copyright 2019-Present Datadog, Inc.
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package datadog

import (
	"encoding/json"
)

// IncidentServicesResponse Response with a list of incident service payloads.
type IncidentServicesResponse struct {
	// An array of incident services.
	Data []IncidentServiceResponseData `json:"data"`
	// Included related resources which the user requested.
	Included *[]IncidentServiceIncludedItems `json:"included,omitempty"`
	Meta     *IncidentServicesResponseMeta   `json:"meta,omitempty"`
}

// NewIncidentServicesResponse instantiates a new IncidentServicesResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIncidentServicesResponse(data []IncidentServiceResponseData) *IncidentServicesResponse {
	this := IncidentServicesResponse{}
	this.Data = data
	return &this
}

// NewIncidentServicesResponseWithDefaults instantiates a new IncidentServicesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIncidentServicesResponseWithDefaults() *IncidentServicesResponse {
	this := IncidentServicesResponse{}
	return &this
}

// GetData returns the Data field value
func (o *IncidentServicesResponse) GetData() []IncidentServiceResponseData {
	if o == nil {
		var ret []IncidentServiceResponseData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *IncidentServicesResponse) GetDataOk() (*[]IncidentServiceResponseData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *IncidentServicesResponse) SetData(v []IncidentServiceResponseData) {
	o.Data = v
}

// GetIncluded returns the Included field value if set, zero value otherwise.
func (o *IncidentServicesResponse) GetIncluded() []IncidentServiceIncludedItems {
	if o == nil || o.Included == nil {
		var ret []IncidentServiceIncludedItems
		return ret
	}
	return *o.Included
}

// GetIncludedOk returns a tuple with the Included field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentServicesResponse) GetIncludedOk() (*[]IncidentServiceIncludedItems, bool) {
	if o == nil || o.Included == nil {
		return nil, false
	}
	return o.Included, true
}

// HasIncluded returns a boolean if a field has been set.
func (o *IncidentServicesResponse) HasIncluded() bool {
	if o != nil && o.Included != nil {
		return true
	}

	return false
}

// SetIncluded gets a reference to the given []IncidentServiceIncludedItems and assigns it to the Included field.
func (o *IncidentServicesResponse) SetIncluded(v []IncidentServiceIncludedItems) {
	o.Included = &v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *IncidentServicesResponse) GetMeta() IncidentServicesResponseMeta {
	if o == nil || o.Meta == nil {
		var ret IncidentServicesResponseMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentServicesResponse) GetMetaOk() (*IncidentServicesResponseMeta, bool) {
	if o == nil || o.Meta == nil {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *IncidentServicesResponse) HasMeta() bool {
	if o != nil && o.Meta != nil {
		return true
	}

	return false
}

// SetMeta gets a reference to the given IncidentServicesResponseMeta and assigns it to the Meta field.
func (o *IncidentServicesResponse) SetMeta(v IncidentServicesResponseMeta) {
	o.Meta = &v
}

func (o IncidentServicesResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["data"] = o.Data
	}
	if o.Included != nil {
		toSerialize["included"] = o.Included
	}
	if o.Meta != nil {
		toSerialize["meta"] = o.Meta
	}
	return json.Marshal(toSerialize)
}

type NullableIncidentServicesResponse struct {
	value *IncidentServicesResponse
	isSet bool
}

func (v NullableIncidentServicesResponse) Get() *IncidentServicesResponse {
	return v.value
}

func (v *NullableIncidentServicesResponse) Set(val *IncidentServicesResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableIncidentServicesResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableIncidentServicesResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIncidentServicesResponse(val *IncidentServicesResponse) *NullableIncidentServicesResponse {
	return &NullableIncidentServicesResponse{value: val, isSet: true}
}

func (v NullableIncidentServicesResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIncidentServicesResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}