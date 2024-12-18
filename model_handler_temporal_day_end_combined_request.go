/*
Treasury Management API

A comprehensive API for managing addresses, offering endpoints for creation, update, deletion, and retrieval of address data.

API version: 1.0
Contact: support_cept@indiapost.gov.in
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the HandlerTemporalDayEndCombinedRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HandlerTemporalDayEndCombinedRequest{}

// HandlerTemporalDayEndCombinedRequest struct for HandlerTemporalDayEndCombinedRequest
type HandlerTemporalDayEndCombinedRequest struct {
	Counterno *int32 `json:"counterno,omitempty"`
	DayEndUserid *int32 `json:"day_end_userid,omitempty"`
	Employeeid *int32 `json:"employeeid,omitempty"`
	Employeename *string `json:"employeename,omitempty"`
	Id *int32 `json:"id,omitempty"`
	OfficeId *int32 `json:"office_id,omitempty"`
	Officeid *int32 `json:"officeid,omitempty"`
	Officename *string `json:"officename,omitempty"`
	Remarks *string `json:"remarks,omitempty"`
	Revertedby *string `json:"revertedby,omitempty"`
	Shiftbegin *bool `json:"shiftbegin,omitempty"`
	Shiftend *bool `json:"shiftend,omitempty"`
	Shiftendchanneltype *string `json:"shiftendchanneltype,omitempty"`
	Shiftenddoneby *int32 `json:"shiftenddoneby,omitempty"`
	Shiftendipaddress *string `json:"shiftendipaddress,omitempty"`
	Shiftendremarks *string `json:"shiftendremarks,omitempty"`
	Shiftendusertype *string `json:"shiftendusertype,omitempty"`
	Shiftno *int32 `json:"shiftno,omitempty"`
	TransDate *string `json:"trans_date,omitempty"`
	Transdate *string `json:"transdate,omitempty"`
}

// NewHandlerTemporalDayEndCombinedRequest instantiates a new HandlerTemporalDayEndCombinedRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHandlerTemporalDayEndCombinedRequest() *HandlerTemporalDayEndCombinedRequest {
	this := HandlerTemporalDayEndCombinedRequest{}
	return &this
}

// NewHandlerTemporalDayEndCombinedRequestWithDefaults instantiates a new HandlerTemporalDayEndCombinedRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHandlerTemporalDayEndCombinedRequestWithDefaults() *HandlerTemporalDayEndCombinedRequest {
	this := HandlerTemporalDayEndCombinedRequest{}
	return &this
}

// GetCounterno returns the Counterno field value if set, zero value otherwise.
func (o *HandlerTemporalDayEndCombinedRequest) GetCounterno() int32 {
	if o == nil || IsNil(o.Counterno) {
		var ret int32
		return ret
	}
	return *o.Counterno
}

// GetCounternoOk returns a tuple with the Counterno field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HandlerTemporalDayEndCombinedRequest) GetCounternoOk() (*int32, bool) {
	if o == nil || IsNil(o.Counterno) {
		return nil, false
	}
	return o.Counterno, true
}

// HasCounterno returns a boolean if a field has been set.
func (o *HandlerTemporalDayEndCombinedRequest) HasCounterno() bool {
	if o != nil && !IsNil(o.Counterno) {
		return true
	}

	return false
}

// SetCounterno gets a reference to the given int32 and assigns it to the Counterno field.
func (o *HandlerTemporalDayEndCombinedRequest) SetCounterno(v int32) {
	o.Counterno = &v
}

// GetDayEndUserid returns the DayEndUserid field value if set, zero value otherwise.
func (o *HandlerTemporalDayEndCombinedRequest) GetDayEndUserid() int32 {
	if o == nil || IsNil(o.DayEndUserid) {
		var ret int32
		return ret
	}
	return *o.DayEndUserid
}

// GetDayEndUseridOk returns a tuple with the DayEndUserid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HandlerTemporalDayEndCombinedRequest) GetDayEndUseridOk() (*int32, bool) {
	if o == nil || IsNil(o.DayEndUserid) {
		return nil, false
	}
	return o.DayEndUserid, true
}

// HasDayEndUserid returns a boolean if a field has been set.
func (o *HandlerTemporalDayEndCombinedRequest) HasDayEndUserid() bool {
	if o != nil && !IsNil(o.DayEndUserid) {
		return true
	}

	return false
}

// SetDayEndUserid gets a reference to the given int32 and assigns it to the DayEndUserid field.
func (o *HandlerTemporalDayEndCombinedRequest) SetDayEndUserid(v int32) {
	o.DayEndUserid = &v
}

// GetEmployeeid returns the Employeeid field value if set, zero value otherwise.
func (o *HandlerTemporalDayEndCombinedRequest) GetEmployeeid() int32 {
	if o == nil || IsNil(o.Employeeid) {
		var ret int32
		return ret
	}
	return *o.Employeeid
}

// GetEmployeeidOk returns a tuple with the Employeeid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HandlerTemporalDayEndCombinedRequest) GetEmployeeidOk() (*int32, bool) {
	if o == nil || IsNil(o.Employeeid) {
		return nil, false
	}
	return o.Employeeid, true
}

// HasEmployeeid returns a boolean if a field has been set.
func (o *HandlerTemporalDayEndCombinedRequest) HasEmployeeid() bool {
	if o != nil && !IsNil(o.Employeeid) {
		return true
	}

	return false
}

// SetEmployeeid gets a reference to the given int32 and assigns it to the Employeeid field.
func (o *HandlerTemporalDayEndCombinedRequest) SetEmployeeid(v int32) {
	o.Employeeid = &v
}

// GetEmployeename returns the Employeename field value if set, zero value otherwise.
func (o *HandlerTemporalDayEndCombinedRequest) GetEmployeename() string {
	if o == nil || IsNil(o.Employeename) {
		var ret string
		return ret
	}
	return *o.Employeename
}

// GetEmployeenameOk returns a tuple with the Employeename field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HandlerTemporalDayEndCombinedRequest) GetEmployeenameOk() (*string, bool) {
	if o == nil || IsNil(o.Employeename) {
		return nil, false
	}
	return o.Employeename, true
}

// HasEmployeename returns a boolean if a field has been set.
func (o *HandlerTemporalDayEndCombinedRequest) HasEmployeename() bool {
	if o != nil && !IsNil(o.Employeename) {
		return true
	}

	return false
}

// SetEmployeename gets a reference to the given string and assigns it to the Employeename field.
func (o *HandlerTemporalDayEndCombinedRequest) SetEmployeename(v string) {
	o.Employeename = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *HandlerTemporalDayEndCombinedRequest) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HandlerTemporalDayEndCombinedRequest) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *HandlerTemporalDayEndCombinedRequest) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *HandlerTemporalDayEndCombinedRequest) SetId(v int32) {
	o.Id = &v
}

// GetOfficeId returns the OfficeId field value if set, zero value otherwise.
func (o *HandlerTemporalDayEndCombinedRequest) GetOfficeId() int32 {
	if o == nil || IsNil(o.OfficeId) {
		var ret int32
		return ret
	}
	return *o.OfficeId
}

// GetOfficeIdOk returns a tuple with the OfficeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HandlerTemporalDayEndCombinedRequest) GetOfficeIdOk() (*int32, bool) {
	if o == nil || IsNil(o.OfficeId) {
		return nil, false
	}
	return o.OfficeId, true
}

// HasOfficeId returns a boolean if a field has been set.
func (o *HandlerTemporalDayEndCombinedRequest) HasOfficeId() bool {
	if o != nil && !IsNil(o.OfficeId) {
		return true
	}

	return false
}

// SetOfficeId gets a reference to the given int32 and assigns it to the OfficeId field.
func (o *HandlerTemporalDayEndCombinedRequest) SetOfficeId(v int32) {
	o.OfficeId = &v
}

// GetOfficeid returns the Officeid field value if set, zero value otherwise.
func (o *HandlerTemporalDayEndCombinedRequest) GetOfficeid() int32 {
	if o == nil || IsNil(o.Officeid) {
		var ret int32
		return ret
	}
	return *o.Officeid
}

// GetOfficeidOk returns a tuple with the Officeid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HandlerTemporalDayEndCombinedRequest) GetOfficeidOk() (*int32, bool) {
	if o == nil || IsNil(o.Officeid) {
		return nil, false
	}
	return o.Officeid, true
}

// HasOfficeid returns a boolean if a field has been set.
func (o *HandlerTemporalDayEndCombinedRequest) HasOfficeid() bool {
	if o != nil && !IsNil(o.Officeid) {
		return true
	}

	return false
}

// SetOfficeid gets a reference to the given int32 and assigns it to the Officeid field.
func (o *HandlerTemporalDayEndCombinedRequest) SetOfficeid(v int32) {
	o.Officeid = &v
}

// GetOfficename returns the Officename field value if set, zero value otherwise.
func (o *HandlerTemporalDayEndCombinedRequest) GetOfficename() string {
	if o == nil || IsNil(o.Officename) {
		var ret string
		return ret
	}
	return *o.Officename
}

// GetOfficenameOk returns a tuple with the Officename field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HandlerTemporalDayEndCombinedRequest) GetOfficenameOk() (*string, bool) {
	if o == nil || IsNil(o.Officename) {
		return nil, false
	}
	return o.Officename, true
}

// HasOfficename returns a boolean if a field has been set.
func (o *HandlerTemporalDayEndCombinedRequest) HasOfficename() bool {
	if o != nil && !IsNil(o.Officename) {
		return true
	}

	return false
}

// SetOfficename gets a reference to the given string and assigns it to the Officename field.
func (o *HandlerTemporalDayEndCombinedRequest) SetOfficename(v string) {
	o.Officename = &v
}

// GetRemarks returns the Remarks field value if set, zero value otherwise.
func (o *HandlerTemporalDayEndCombinedRequest) GetRemarks() string {
	if o == nil || IsNil(o.Remarks) {
		var ret string
		return ret
	}
	return *o.Remarks
}

// GetRemarksOk returns a tuple with the Remarks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HandlerTemporalDayEndCombinedRequest) GetRemarksOk() (*string, bool) {
	if o == nil || IsNil(o.Remarks) {
		return nil, false
	}
	return o.Remarks, true
}

// HasRemarks returns a boolean if a field has been set.
func (o *HandlerTemporalDayEndCombinedRequest) HasRemarks() bool {
	if o != nil && !IsNil(o.Remarks) {
		return true
	}

	return false
}

// SetRemarks gets a reference to the given string and assigns it to the Remarks field.
func (o *HandlerTemporalDayEndCombinedRequest) SetRemarks(v string) {
	o.Remarks = &v
}

// GetRevertedby returns the Revertedby field value if set, zero value otherwise.
func (o *HandlerTemporalDayEndCombinedRequest) GetRevertedby() string {
	if o == nil || IsNil(o.Revertedby) {
		var ret string
		return ret
	}
	return *o.Revertedby
}

// GetRevertedbyOk returns a tuple with the Revertedby field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HandlerTemporalDayEndCombinedRequest) GetRevertedbyOk() (*string, bool) {
	if o == nil || IsNil(o.Revertedby) {
		return nil, false
	}
	return o.Revertedby, true
}

// HasRevertedby returns a boolean if a field has been set.
func (o *HandlerTemporalDayEndCombinedRequest) HasRevertedby() bool {
	if o != nil && !IsNil(o.Revertedby) {
		return true
	}

	return false
}

// SetRevertedby gets a reference to the given string and assigns it to the Revertedby field.
func (o *HandlerTemporalDayEndCombinedRequest) SetRevertedby(v string) {
	o.Revertedby = &v
}

// GetShiftbegin returns the Shiftbegin field value if set, zero value otherwise.
func (o *HandlerTemporalDayEndCombinedRequest) GetShiftbegin() bool {
	if o == nil || IsNil(o.Shiftbegin) {
		var ret bool
		return ret
	}
	return *o.Shiftbegin
}

// GetShiftbeginOk returns a tuple with the Shiftbegin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HandlerTemporalDayEndCombinedRequest) GetShiftbeginOk() (*bool, bool) {
	if o == nil || IsNil(o.Shiftbegin) {
		return nil, false
	}
	return o.Shiftbegin, true
}

// HasShiftbegin returns a boolean if a field has been set.
func (o *HandlerTemporalDayEndCombinedRequest) HasShiftbegin() bool {
	if o != nil && !IsNil(o.Shiftbegin) {
		return true
	}

	return false
}

// SetShiftbegin gets a reference to the given bool and assigns it to the Shiftbegin field.
func (o *HandlerTemporalDayEndCombinedRequest) SetShiftbegin(v bool) {
	o.Shiftbegin = &v
}

// GetShiftend returns the Shiftend field value if set, zero value otherwise.
func (o *HandlerTemporalDayEndCombinedRequest) GetShiftend() bool {
	if o == nil || IsNil(o.Shiftend) {
		var ret bool
		return ret
	}
	return *o.Shiftend
}

// GetShiftendOk returns a tuple with the Shiftend field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HandlerTemporalDayEndCombinedRequest) GetShiftendOk() (*bool, bool) {
	if o == nil || IsNil(o.Shiftend) {
		return nil, false
	}
	return o.Shiftend, true
}

// HasShiftend returns a boolean if a field has been set.
func (o *HandlerTemporalDayEndCombinedRequest) HasShiftend() bool {
	if o != nil && !IsNil(o.Shiftend) {
		return true
	}

	return false
}

// SetShiftend gets a reference to the given bool and assigns it to the Shiftend field.
func (o *HandlerTemporalDayEndCombinedRequest) SetShiftend(v bool) {
	o.Shiftend = &v
}

// GetShiftendchanneltype returns the Shiftendchanneltype field value if set, zero value otherwise.
func (o *HandlerTemporalDayEndCombinedRequest) GetShiftendchanneltype() string {
	if o == nil || IsNil(o.Shiftendchanneltype) {
		var ret string
		return ret
	}
	return *o.Shiftendchanneltype
}

// GetShiftendchanneltypeOk returns a tuple with the Shiftendchanneltype field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HandlerTemporalDayEndCombinedRequest) GetShiftendchanneltypeOk() (*string, bool) {
	if o == nil || IsNil(o.Shiftendchanneltype) {
		return nil, false
	}
	return o.Shiftendchanneltype, true
}

// HasShiftendchanneltype returns a boolean if a field has been set.
func (o *HandlerTemporalDayEndCombinedRequest) HasShiftendchanneltype() bool {
	if o != nil && !IsNil(o.Shiftendchanneltype) {
		return true
	}

	return false
}

// SetShiftendchanneltype gets a reference to the given string and assigns it to the Shiftendchanneltype field.
func (o *HandlerTemporalDayEndCombinedRequest) SetShiftendchanneltype(v string) {
	o.Shiftendchanneltype = &v
}

// GetShiftenddoneby returns the Shiftenddoneby field value if set, zero value otherwise.
func (o *HandlerTemporalDayEndCombinedRequest) GetShiftenddoneby() int32 {
	if o == nil || IsNil(o.Shiftenddoneby) {
		var ret int32
		return ret
	}
	return *o.Shiftenddoneby
}

// GetShiftenddonebyOk returns a tuple with the Shiftenddoneby field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HandlerTemporalDayEndCombinedRequest) GetShiftenddonebyOk() (*int32, bool) {
	if o == nil || IsNil(o.Shiftenddoneby) {
		return nil, false
	}
	return o.Shiftenddoneby, true
}

// HasShiftenddoneby returns a boolean if a field has been set.
func (o *HandlerTemporalDayEndCombinedRequest) HasShiftenddoneby() bool {
	if o != nil && !IsNil(o.Shiftenddoneby) {
		return true
	}

	return false
}

// SetShiftenddoneby gets a reference to the given int32 and assigns it to the Shiftenddoneby field.
func (o *HandlerTemporalDayEndCombinedRequest) SetShiftenddoneby(v int32) {
	o.Shiftenddoneby = &v
}

// GetShiftendipaddress returns the Shiftendipaddress field value if set, zero value otherwise.
func (o *HandlerTemporalDayEndCombinedRequest) GetShiftendipaddress() string {
	if o == nil || IsNil(o.Shiftendipaddress) {
		var ret string
		return ret
	}
	return *o.Shiftendipaddress
}

// GetShiftendipaddressOk returns a tuple with the Shiftendipaddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HandlerTemporalDayEndCombinedRequest) GetShiftendipaddressOk() (*string, bool) {
	if o == nil || IsNil(o.Shiftendipaddress) {
		return nil, false
	}
	return o.Shiftendipaddress, true
}

// HasShiftendipaddress returns a boolean if a field has been set.
func (o *HandlerTemporalDayEndCombinedRequest) HasShiftendipaddress() bool {
	if o != nil && !IsNil(o.Shiftendipaddress) {
		return true
	}

	return false
}

// SetShiftendipaddress gets a reference to the given string and assigns it to the Shiftendipaddress field.
func (o *HandlerTemporalDayEndCombinedRequest) SetShiftendipaddress(v string) {
	o.Shiftendipaddress = &v
}

// GetShiftendremarks returns the Shiftendremarks field value if set, zero value otherwise.
func (o *HandlerTemporalDayEndCombinedRequest) GetShiftendremarks() string {
	if o == nil || IsNil(o.Shiftendremarks) {
		var ret string
		return ret
	}
	return *o.Shiftendremarks
}

// GetShiftendremarksOk returns a tuple with the Shiftendremarks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HandlerTemporalDayEndCombinedRequest) GetShiftendremarksOk() (*string, bool) {
	if o == nil || IsNil(o.Shiftendremarks) {
		return nil, false
	}
	return o.Shiftendremarks, true
}

// HasShiftendremarks returns a boolean if a field has been set.
func (o *HandlerTemporalDayEndCombinedRequest) HasShiftendremarks() bool {
	if o != nil && !IsNil(o.Shiftendremarks) {
		return true
	}

	return false
}

// SetShiftendremarks gets a reference to the given string and assigns it to the Shiftendremarks field.
func (o *HandlerTemporalDayEndCombinedRequest) SetShiftendremarks(v string) {
	o.Shiftendremarks = &v
}

// GetShiftendusertype returns the Shiftendusertype field value if set, zero value otherwise.
func (o *HandlerTemporalDayEndCombinedRequest) GetShiftendusertype() string {
	if o == nil || IsNil(o.Shiftendusertype) {
		var ret string
		return ret
	}
	return *o.Shiftendusertype
}

// GetShiftendusertypeOk returns a tuple with the Shiftendusertype field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HandlerTemporalDayEndCombinedRequest) GetShiftendusertypeOk() (*string, bool) {
	if o == nil || IsNil(o.Shiftendusertype) {
		return nil, false
	}
	return o.Shiftendusertype, true
}

// HasShiftendusertype returns a boolean if a field has been set.
func (o *HandlerTemporalDayEndCombinedRequest) HasShiftendusertype() bool {
	if o != nil && !IsNil(o.Shiftendusertype) {
		return true
	}

	return false
}

// SetShiftendusertype gets a reference to the given string and assigns it to the Shiftendusertype field.
func (o *HandlerTemporalDayEndCombinedRequest) SetShiftendusertype(v string) {
	o.Shiftendusertype = &v
}

// GetShiftno returns the Shiftno field value if set, zero value otherwise.
func (o *HandlerTemporalDayEndCombinedRequest) GetShiftno() int32 {
	if o == nil || IsNil(o.Shiftno) {
		var ret int32
		return ret
	}
	return *o.Shiftno
}

// GetShiftnoOk returns a tuple with the Shiftno field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HandlerTemporalDayEndCombinedRequest) GetShiftnoOk() (*int32, bool) {
	if o == nil || IsNil(o.Shiftno) {
		return nil, false
	}
	return o.Shiftno, true
}

// HasShiftno returns a boolean if a field has been set.
func (o *HandlerTemporalDayEndCombinedRequest) HasShiftno() bool {
	if o != nil && !IsNil(o.Shiftno) {
		return true
	}

	return false
}

// SetShiftno gets a reference to the given int32 and assigns it to the Shiftno field.
func (o *HandlerTemporalDayEndCombinedRequest) SetShiftno(v int32) {
	o.Shiftno = &v
}

// GetTransDate returns the TransDate field value if set, zero value otherwise.
func (o *HandlerTemporalDayEndCombinedRequest) GetTransDate() string {
	if o == nil || IsNil(o.TransDate) {
		var ret string
		return ret
	}
	return *o.TransDate
}

// GetTransDateOk returns a tuple with the TransDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HandlerTemporalDayEndCombinedRequest) GetTransDateOk() (*string, bool) {
	if o == nil || IsNil(o.TransDate) {
		return nil, false
	}
	return o.TransDate, true
}

// HasTransDate returns a boolean if a field has been set.
func (o *HandlerTemporalDayEndCombinedRequest) HasTransDate() bool {
	if o != nil && !IsNil(o.TransDate) {
		return true
	}

	return false
}

// SetTransDate gets a reference to the given string and assigns it to the TransDate field.
func (o *HandlerTemporalDayEndCombinedRequest) SetTransDate(v string) {
	o.TransDate = &v
}

// GetTransdate returns the Transdate field value if set, zero value otherwise.
func (o *HandlerTemporalDayEndCombinedRequest) GetTransdate() string {
	if o == nil || IsNil(o.Transdate) {
		var ret string
		return ret
	}
	return *o.Transdate
}

// GetTransdateOk returns a tuple with the Transdate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HandlerTemporalDayEndCombinedRequest) GetTransdateOk() (*string, bool) {
	if o == nil || IsNil(o.Transdate) {
		return nil, false
	}
	return o.Transdate, true
}

// HasTransdate returns a boolean if a field has been set.
func (o *HandlerTemporalDayEndCombinedRequest) HasTransdate() bool {
	if o != nil && !IsNil(o.Transdate) {
		return true
	}

	return false
}

// SetTransdate gets a reference to the given string and assigns it to the Transdate field.
func (o *HandlerTemporalDayEndCombinedRequest) SetTransdate(v string) {
	o.Transdate = &v
}

func (o HandlerTemporalDayEndCombinedRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HandlerTemporalDayEndCombinedRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Counterno) {
		toSerialize["counterno"] = o.Counterno
	}
	if !IsNil(o.DayEndUserid) {
		toSerialize["day_end_userid"] = o.DayEndUserid
	}
	if !IsNil(o.Employeeid) {
		toSerialize["employeeid"] = o.Employeeid
	}
	if !IsNil(o.Employeename) {
		toSerialize["employeename"] = o.Employeename
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.OfficeId) {
		toSerialize["office_id"] = o.OfficeId
	}
	if !IsNil(o.Officeid) {
		toSerialize["officeid"] = o.Officeid
	}
	if !IsNil(o.Officename) {
		toSerialize["officename"] = o.Officename
	}
	if !IsNil(o.Remarks) {
		toSerialize["remarks"] = o.Remarks
	}
	if !IsNil(o.Revertedby) {
		toSerialize["revertedby"] = o.Revertedby
	}
	if !IsNil(o.Shiftbegin) {
		toSerialize["shiftbegin"] = o.Shiftbegin
	}
	if !IsNil(o.Shiftend) {
		toSerialize["shiftend"] = o.Shiftend
	}
	if !IsNil(o.Shiftendchanneltype) {
		toSerialize["shiftendchanneltype"] = o.Shiftendchanneltype
	}
	if !IsNil(o.Shiftenddoneby) {
		toSerialize["shiftenddoneby"] = o.Shiftenddoneby
	}
	if !IsNil(o.Shiftendipaddress) {
		toSerialize["shiftendipaddress"] = o.Shiftendipaddress
	}
	if !IsNil(o.Shiftendremarks) {
		toSerialize["shiftendremarks"] = o.Shiftendremarks
	}
	if !IsNil(o.Shiftendusertype) {
		toSerialize["shiftendusertype"] = o.Shiftendusertype
	}
	if !IsNil(o.Shiftno) {
		toSerialize["shiftno"] = o.Shiftno
	}
	if !IsNil(o.TransDate) {
		toSerialize["trans_date"] = o.TransDate
	}
	if !IsNil(o.Transdate) {
		toSerialize["transdate"] = o.Transdate
	}
	return toSerialize, nil
}

type NullableHandlerTemporalDayEndCombinedRequest struct {
	value *HandlerTemporalDayEndCombinedRequest
	isSet bool
}

func (v NullableHandlerTemporalDayEndCombinedRequest) Get() *HandlerTemporalDayEndCombinedRequest {
	return v.value
}

func (v *NullableHandlerTemporalDayEndCombinedRequest) Set(val *HandlerTemporalDayEndCombinedRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableHandlerTemporalDayEndCombinedRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableHandlerTemporalDayEndCombinedRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHandlerTemporalDayEndCombinedRequest(val *HandlerTemporalDayEndCombinedRequest) *NullableHandlerTemporalDayEndCombinedRequest {
	return &NullableHandlerTemporalDayEndCombinedRequest{value: val, isSet: true}
}

func (v NullableHandlerTemporalDayEndCombinedRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHandlerTemporalDayEndCombinedRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


