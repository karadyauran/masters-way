/*
Masters way general API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the SchemasCreateJobDonePayload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SchemasCreateJobDonePayload{}

// SchemasCreateJobDonePayload struct for SchemasCreateJobDonePayload
type SchemasCreateJobDonePayload struct {
	DayReportUuid string
	Description string
	JobTagUuids []string
	OwnerUuid string
	Time int32
}

type _SchemasCreateJobDonePayload SchemasCreateJobDonePayload

// NewSchemasCreateJobDonePayload instantiates a new SchemasCreateJobDonePayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSchemasCreateJobDonePayload(dayReportUuid string, description string, jobTagUuids []string, ownerUuid string, time int32) *SchemasCreateJobDonePayload {
	this := SchemasCreateJobDonePayload{}
	this.DayReportUuid = dayReportUuid
	this.Description = description
	this.JobTagUuids = jobTagUuids
	this.OwnerUuid = ownerUuid
	this.Time = time
	return &this
}

// NewSchemasCreateJobDonePayloadWithDefaults instantiates a new SchemasCreateJobDonePayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSchemasCreateJobDonePayloadWithDefaults() *SchemasCreateJobDonePayload {
	this := SchemasCreateJobDonePayload{}
	return &this
}

// GetDayReportUuid returns the DayReportUuid field value
func (o *SchemasCreateJobDonePayload) GetDayReportUuid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DayReportUuid
}

// GetDayReportUuidOk returns a tuple with the DayReportUuid field value
// and a boolean to check if the value has been set.
func (o *SchemasCreateJobDonePayload) GetDayReportUuidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DayReportUuid, true
}

// SetDayReportUuid sets field value
func (o *SchemasCreateJobDonePayload) SetDayReportUuid(v string) {
	o.DayReportUuid = v
}

// GetDescription returns the Description field value
func (o *SchemasCreateJobDonePayload) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *SchemasCreateJobDonePayload) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *SchemasCreateJobDonePayload) SetDescription(v string) {
	o.Description = v
}

// GetJobTagUuids returns the JobTagUuids field value
func (o *SchemasCreateJobDonePayload) GetJobTagUuids() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.JobTagUuids
}

// GetJobTagUuidsOk returns a tuple with the JobTagUuids field value
// and a boolean to check if the value has been set.
func (o *SchemasCreateJobDonePayload) GetJobTagUuidsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.JobTagUuids, true
}

// SetJobTagUuids sets field value
func (o *SchemasCreateJobDonePayload) SetJobTagUuids(v []string) {
	o.JobTagUuids = v
}

// GetOwnerUuid returns the OwnerUuid field value
func (o *SchemasCreateJobDonePayload) GetOwnerUuid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OwnerUuid
}

// GetOwnerUuidOk returns a tuple with the OwnerUuid field value
// and a boolean to check if the value has been set.
func (o *SchemasCreateJobDonePayload) GetOwnerUuidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OwnerUuid, true
}

// SetOwnerUuid sets field value
func (o *SchemasCreateJobDonePayload) SetOwnerUuid(v string) {
	o.OwnerUuid = v
}

// GetTime returns the Time field value
func (o *SchemasCreateJobDonePayload) GetTime() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Time
}

// GetTimeOk returns a tuple with the Time field value
// and a boolean to check if the value has been set.
func (o *SchemasCreateJobDonePayload) GetTimeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Time, true
}

// SetTime sets field value
func (o *SchemasCreateJobDonePayload) SetTime(v int32) {
	o.Time = v
}

func (o SchemasCreateJobDonePayload) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SchemasCreateJobDonePayload) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["dayReportUuid"] = o.DayReportUuid
	toSerialize["description"] = o.Description
	toSerialize["jobTagUuids"] = o.JobTagUuids
	toSerialize["ownerUuid"] = o.OwnerUuid
	toSerialize["time"] = o.Time
	return toSerialize, nil
}

func (o *SchemasCreateJobDonePayload) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"dayReportUuid",
		"description",
		"jobTagUuids",
		"ownerUuid",
		"time",
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

	varSchemasCreateJobDonePayload := _SchemasCreateJobDonePayload{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSchemasCreateJobDonePayload)

	if err != nil {
		return err
	}

	*o = SchemasCreateJobDonePayload(varSchemasCreateJobDonePayload)

	return err
}

type NullableSchemasCreateJobDonePayload struct {
	value *SchemasCreateJobDonePayload
	isSet bool
}

func (v NullableSchemasCreateJobDonePayload) Get() *SchemasCreateJobDonePayload {
	return v.value
}

func (v *NullableSchemasCreateJobDonePayload) Set(val *SchemasCreateJobDonePayload) {
	v.value = val
	v.isSet = true
}

func (v NullableSchemasCreateJobDonePayload) IsSet() bool {
	return v.isSet
}

func (v *NullableSchemasCreateJobDonePayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSchemasCreateJobDonePayload(val *SchemasCreateJobDonePayload) *NullableSchemasCreateJobDonePayload {
	return &NullableSchemasCreateJobDonePayload{value: val, isSet: true}
}

func (v NullableSchemasCreateJobDonePayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSchemasCreateJobDonePayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


