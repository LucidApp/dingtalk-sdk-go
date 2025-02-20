// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package calendar_1_0

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
)

type AddAttendeeHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s AddAttendeeHeaders) String() string {
	return tea.Prettify(s)
}

func (s AddAttendeeHeaders) GoString() string {
	return s.String()
}

func (s *AddAttendeeHeaders) SetCommonHeaders(v map[string]*string) *AddAttendeeHeaders {
	s.CommonHeaders = v
	return s
}

func (s *AddAttendeeHeaders) SetXAcsDingtalkAccessToken(v string) *AddAttendeeHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type AddAttendeeRequest struct {
	AttendeesToAdd []*AddAttendeeRequestAttendeesToAdd `json:"attendeesToAdd,omitempty" xml:"attendeesToAdd,omitempty" type:"Repeated"`
}

func (s AddAttendeeRequest) String() string {
	return tea.Prettify(s)
}

func (s AddAttendeeRequest) GoString() string {
	return s.String()
}

func (s *AddAttendeeRequest) SetAttendeesToAdd(v []*AddAttendeeRequestAttendeesToAdd) *AddAttendeeRequest {
	s.AttendeesToAdd = v
	return s
}

type AddAttendeeRequestAttendeesToAdd struct {
	Id         *string `json:"id,omitempty" xml:"id,omitempty"`
	IsOptional *bool   `json:"isOptional,omitempty" xml:"isOptional,omitempty"`
}

func (s AddAttendeeRequestAttendeesToAdd) String() string {
	return tea.Prettify(s)
}

func (s AddAttendeeRequestAttendeesToAdd) GoString() string {
	return s.String()
}

func (s *AddAttendeeRequestAttendeesToAdd) SetId(v string) *AddAttendeeRequestAttendeesToAdd {
	s.Id = &v
	return s
}

func (s *AddAttendeeRequestAttendeesToAdd) SetIsOptional(v bool) *AddAttendeeRequestAttendeesToAdd {
	s.IsOptional = &v
	return s
}

type AddAttendeeResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
}

func (s AddAttendeeResponse) String() string {
	return tea.Prettify(s)
}

func (s AddAttendeeResponse) GoString() string {
	return s.String()
}

func (s *AddAttendeeResponse) SetHeaders(v map[string]*string) *AddAttendeeResponse {
	s.Headers = v
	return s
}

type CheckInHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CheckInHeaders) String() string {
	return tea.Prettify(s)
}

func (s CheckInHeaders) GoString() string {
	return s.String()
}

func (s *CheckInHeaders) SetCommonHeaders(v map[string]*string) *CheckInHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CheckInHeaders) SetXAcsDingtalkAccessToken(v string) *CheckInHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CheckInResponseBody struct {
	// 签到时间戳
	CheckInTime *int64 `json:"checkInTime,omitempty" xml:"checkInTime,omitempty"`
}

func (s CheckInResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CheckInResponseBody) GoString() string {
	return s.String()
}

func (s *CheckInResponseBody) SetCheckInTime(v int64) *CheckInResponseBody {
	s.CheckInTime = &v
	return s
}

type CheckInResponse struct {
	Headers map[string]*string   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CheckInResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CheckInResponse) String() string {
	return tea.Prettify(s)
}

func (s CheckInResponse) GoString() string {
	return s.String()
}

func (s *CheckInResponse) SetHeaders(v map[string]*string) *CheckInResponse {
	s.Headers = v
	return s
}

func (s *CheckInResponse) SetBody(v *CheckInResponseBody) *CheckInResponse {
	s.Body = v
	return s
}

type ConvertLegacyEventIdHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s ConvertLegacyEventIdHeaders) String() string {
	return tea.Prettify(s)
}

func (s ConvertLegacyEventIdHeaders) GoString() string {
	return s.String()
}

func (s *ConvertLegacyEventIdHeaders) SetCommonHeaders(v map[string]*string) *ConvertLegacyEventIdHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ConvertLegacyEventIdHeaders) SetXAcsDingtalkAccessToken(v string) *ConvertLegacyEventIdHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type ConvertLegacyEventIdRequest struct {
	LegacyEventIds []*string `json:"legacyEventIds,omitempty" xml:"legacyEventIds,omitempty" type:"Repeated"`
}

func (s ConvertLegacyEventIdRequest) String() string {
	return tea.Prettify(s)
}

func (s ConvertLegacyEventIdRequest) GoString() string {
	return s.String()
}

func (s *ConvertLegacyEventIdRequest) SetLegacyEventIds(v []*string) *ConvertLegacyEventIdRequest {
	s.LegacyEventIds = v
	return s
}

type ConvertLegacyEventIdResponseBody struct {
	// legacyEventIdMap
	LegacyEventIdMap map[string]interface{} `json:"legacyEventIdMap,omitempty" xml:"legacyEventIdMap,omitempty"`
}

func (s ConvertLegacyEventIdResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ConvertLegacyEventIdResponseBody) GoString() string {
	return s.String()
}

func (s *ConvertLegacyEventIdResponseBody) SetLegacyEventIdMap(v map[string]interface{}) *ConvertLegacyEventIdResponseBody {
	s.LegacyEventIdMap = v
	return s
}

type ConvertLegacyEventIdResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ConvertLegacyEventIdResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ConvertLegacyEventIdResponse) String() string {
	return tea.Prettify(s)
}

func (s ConvertLegacyEventIdResponse) GoString() string {
	return s.String()
}

func (s *ConvertLegacyEventIdResponse) SetHeaders(v map[string]*string) *ConvertLegacyEventIdResponse {
	s.Headers = v
	return s
}

func (s *ConvertLegacyEventIdResponse) SetBody(v *ConvertLegacyEventIdResponseBody) *ConvertLegacyEventIdResponse {
	s.Body = v
	return s
}

type CreateAclsHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CreateAclsHeaders) String() string {
	return tea.Prettify(s)
}

func (s CreateAclsHeaders) GoString() string {
	return s.String()
}

func (s *CreateAclsHeaders) SetCommonHeaders(v map[string]*string) *CreateAclsHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreateAclsHeaders) SetXAcsDingtalkAccessToken(v string) *CreateAclsHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CreateAclsRequest struct {
	// 对日历的访问权限
	Privilege *string `json:"privilege,omitempty" xml:"privilege,omitempty"`
	// 权限范围
	Scope *CreateAclsRequestScope `json:"scope,omitempty" xml:"scope,omitempty" type:"Struct"`
	// 是否向授权人发消息
	SendMsg *bool `json:"sendMsg,omitempty" xml:"sendMsg,omitempty"`
}

func (s CreateAclsRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateAclsRequest) GoString() string {
	return s.String()
}

func (s *CreateAclsRequest) SetPrivilege(v string) *CreateAclsRequest {
	s.Privilege = &v
	return s
}

func (s *CreateAclsRequest) SetScope(v *CreateAclsRequestScope) *CreateAclsRequest {
	s.Scope = v
	return s
}

func (s *CreateAclsRequest) SetSendMsg(v bool) *CreateAclsRequest {
	s.SendMsg = &v
	return s
}

type CreateAclsRequestScope struct {
	// 权限类型
	ScopeType *string `json:"scopeType,omitempty" xml:"scopeType,omitempty"`
	// 用户id
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s CreateAclsRequestScope) String() string {
	return tea.Prettify(s)
}

func (s CreateAclsRequestScope) GoString() string {
	return s.String()
}

func (s *CreateAclsRequestScope) SetScopeType(v string) *CreateAclsRequestScope {
	s.ScopeType = &v
	return s
}

func (s *CreateAclsRequestScope) SetUserId(v string) *CreateAclsRequestScope {
	s.UserId = &v
	return s
}

type CreateAclsResponseBody struct {
	// acl资源ID
	AclId *string `json:"aclId,omitempty" xml:"aclId,omitempty"`
	// 对日历的访问权限
	Privilege *string `json:"privilege,omitempty" xml:"privilege,omitempty"`
	// 权限范围
	Scope *CreateAclsResponseBodyScope `json:"scope,omitempty" xml:"scope,omitempty" type:"Struct"`
}

func (s CreateAclsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateAclsResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAclsResponseBody) SetAclId(v string) *CreateAclsResponseBody {
	s.AclId = &v
	return s
}

func (s *CreateAclsResponseBody) SetPrivilege(v string) *CreateAclsResponseBody {
	s.Privilege = &v
	return s
}

func (s *CreateAclsResponseBody) SetScope(v *CreateAclsResponseBodyScope) *CreateAclsResponseBody {
	s.Scope = v
	return s
}

type CreateAclsResponseBodyScope struct {
	// 权限类型
	ScopeType *string `json:"scopeType,omitempty" xml:"scopeType,omitempty"`
	// 用户id
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s CreateAclsResponseBodyScope) String() string {
	return tea.Prettify(s)
}

func (s CreateAclsResponseBodyScope) GoString() string {
	return s.String()
}

func (s *CreateAclsResponseBodyScope) SetScopeType(v string) *CreateAclsResponseBodyScope {
	s.ScopeType = &v
	return s
}

func (s *CreateAclsResponseBodyScope) SetUserId(v string) *CreateAclsResponseBodyScope {
	s.UserId = &v
	return s
}

type CreateAclsResponse struct {
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateAclsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateAclsResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateAclsResponse) GoString() string {
	return s.String()
}

func (s *CreateAclsResponse) SetHeaders(v map[string]*string) *CreateAclsResponse {
	s.Headers = v
	return s
}

func (s *CreateAclsResponse) SetBody(v *CreateAclsResponseBody) *CreateAclsResponse {
	s.Body = v
	return s
}

type CreateEventHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CreateEventHeaders) String() string {
	return tea.Prettify(s)
}

func (s CreateEventHeaders) GoString() string {
	return s.String()
}

func (s *CreateEventHeaders) SetCommonHeaders(v map[string]*string) *CreateEventHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreateEventHeaders) SetXAcsDingtalkAccessToken(v string) *CreateEventHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CreateEventRequest struct {
	Attendees []*CreateEventRequestAttendees `json:"attendees,omitempty" xml:"attendees,omitempty" type:"Repeated"`
	// 日程描述
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// 日程结束时间
	End *CreateEventRequestEnd `json:"end,omitempty" xml:"end,omitempty" type:"Struct"`
	// 扩展信息
	Extra map[string]*string `json:"extra,omitempty" xml:"extra,omitempty"`
	// 是否为全天日程
	IsAllDay          *bool                                `json:"isAllDay,omitempty" xml:"isAllDay,omitempty"`
	Location          *CreateEventRequestLocation          `json:"location,omitempty" xml:"location,omitempty" type:"Struct"`
	OnlineMeetingInfo *CreateEventRequestOnlineMeetingInfo `json:"onlineMeetingInfo,omitempty" xml:"onlineMeetingInfo,omitempty" type:"Struct"`
	// 日程循环规则
	Recurrence *CreateEventRequestRecurrence  `json:"recurrence,omitempty" xml:"recurrence,omitempty" type:"Struct"`
	Reminders  []*CreateEventRequestReminders `json:"reminders,omitempty" xml:"reminders,omitempty" type:"Repeated"`
	// 日程开始时间
	Start *CreateEventRequestStart `json:"start,omitempty" xml:"start,omitempty" type:"Struct"`
	// 日程标题
	Summary *string `json:"summary,omitempty" xml:"summary,omitempty"`
}

func (s CreateEventRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateEventRequest) GoString() string {
	return s.String()
}

func (s *CreateEventRequest) SetAttendees(v []*CreateEventRequestAttendees) *CreateEventRequest {
	s.Attendees = v
	return s
}

func (s *CreateEventRequest) SetDescription(v string) *CreateEventRequest {
	s.Description = &v
	return s
}

func (s *CreateEventRequest) SetEnd(v *CreateEventRequestEnd) *CreateEventRequest {
	s.End = v
	return s
}

func (s *CreateEventRequest) SetExtra(v map[string]*string) *CreateEventRequest {
	s.Extra = v
	return s
}

func (s *CreateEventRequest) SetIsAllDay(v bool) *CreateEventRequest {
	s.IsAllDay = &v
	return s
}

func (s *CreateEventRequest) SetLocation(v *CreateEventRequestLocation) *CreateEventRequest {
	s.Location = v
	return s
}

func (s *CreateEventRequest) SetOnlineMeetingInfo(v *CreateEventRequestOnlineMeetingInfo) *CreateEventRequest {
	s.OnlineMeetingInfo = v
	return s
}

func (s *CreateEventRequest) SetRecurrence(v *CreateEventRequestRecurrence) *CreateEventRequest {
	s.Recurrence = v
	return s
}

func (s *CreateEventRequest) SetReminders(v []*CreateEventRequestReminders) *CreateEventRequest {
	s.Reminders = v
	return s
}

func (s *CreateEventRequest) SetStart(v *CreateEventRequestStart) *CreateEventRequest {
	s.Start = v
	return s
}

func (s *CreateEventRequest) SetSummary(v string) *CreateEventRequest {
	s.Summary = &v
	return s
}

type CreateEventRequestAttendees struct {
	Id         *string `json:"id,omitempty" xml:"id,omitempty"`
	IsOptional *bool   `json:"isOptional,omitempty" xml:"isOptional,omitempty"`
}

func (s CreateEventRequestAttendees) String() string {
	return tea.Prettify(s)
}

func (s CreateEventRequestAttendees) GoString() string {
	return s.String()
}

func (s *CreateEventRequestAttendees) SetId(v string) *CreateEventRequestAttendees {
	s.Id = &v
	return s
}

func (s *CreateEventRequestAttendees) SetIsOptional(v bool) *CreateEventRequestAttendees {
	s.IsOptional = &v
	return s
}

type CreateEventRequestEnd struct {
	// 日程结束日期，如果是全天日程必须有值，非全天日程必须留空，格式：yyyy-MM-dd
	Date *string `json:"date,omitempty" xml:"date,omitempty"`
	// 日程结束时间，非全天日程必须有值，全天日程必须留空，格式为ISO-8601的date-time格式
	DateTime *string `json:"dateTime,omitempty" xml:"dateTime,omitempty"`
	// 日程结束时间所属时区，非全天日程必须有值，全天日程必须留空，tz database name格式，参考：https://en.wikipedia.org/wiki/List_of_tz_database_time_zones
	TimeZone *string `json:"timeZone,omitempty" xml:"timeZone,omitempty"`
}

func (s CreateEventRequestEnd) String() string {
	return tea.Prettify(s)
}

func (s CreateEventRequestEnd) GoString() string {
	return s.String()
}

func (s *CreateEventRequestEnd) SetDate(v string) *CreateEventRequestEnd {
	s.Date = &v
	return s
}

func (s *CreateEventRequestEnd) SetDateTime(v string) *CreateEventRequestEnd {
	s.DateTime = &v
	return s
}

func (s *CreateEventRequestEnd) SetTimeZone(v string) *CreateEventRequestEnd {
	s.TimeZone = &v
	return s
}

type CreateEventRequestLocation struct {
	DisplayName *string `json:"displayName,omitempty" xml:"displayName,omitempty"`
}

func (s CreateEventRequestLocation) String() string {
	return tea.Prettify(s)
}

func (s CreateEventRequestLocation) GoString() string {
	return s.String()
}

func (s *CreateEventRequestLocation) SetDisplayName(v string) *CreateEventRequestLocation {
	s.DisplayName = &v
	return s
}

type CreateEventRequestOnlineMeetingInfo struct {
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s CreateEventRequestOnlineMeetingInfo) String() string {
	return tea.Prettify(s)
}

func (s CreateEventRequestOnlineMeetingInfo) GoString() string {
	return s.String()
}

func (s *CreateEventRequestOnlineMeetingInfo) SetType(v string) *CreateEventRequestOnlineMeetingInfo {
	s.Type = &v
	return s
}

type CreateEventRequestRecurrence struct {
	// 循环规则
	Pattern *CreateEventRequestRecurrencePattern `json:"pattern,omitempty" xml:"pattern,omitempty" type:"Struct"`
	Range   *CreateEventRequestRecurrenceRange   `json:"range,omitempty" xml:"range,omitempty" type:"Struct"`
}

func (s CreateEventRequestRecurrence) String() string {
	return tea.Prettify(s)
}

func (s CreateEventRequestRecurrence) GoString() string {
	return s.String()
}

func (s *CreateEventRequestRecurrence) SetPattern(v *CreateEventRequestRecurrencePattern) *CreateEventRequestRecurrence {
	s.Pattern = v
	return s
}

func (s *CreateEventRequestRecurrence) SetRange(v *CreateEventRequestRecurrenceRange) *CreateEventRequestRecurrence {
	s.Range = v
	return s
}

type CreateEventRequestRecurrencePattern struct {
	DayOfMonth *int32  `json:"dayOfMonth,omitempty" xml:"dayOfMonth,omitempty"`
	DaysOfWeek *string `json:"daysOfWeek,omitempty" xml:"daysOfWeek,omitempty"`
	Index      *string `json:"index,omitempty" xml:"index,omitempty"`
	Interval   *int32  `json:"interval,omitempty" xml:"interval,omitempty"`
	// 循环规则类型：  daily：每interval天 weekly：每interval周的第daysOfWeek天 absoluteMonthly：每interval月的第dayOfMonth天 relativeMonthly：每interval月的第index周的第daysOfWeek天 absoluteYearly：每interval年
	//
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s CreateEventRequestRecurrencePattern) String() string {
	return tea.Prettify(s)
}

func (s CreateEventRequestRecurrencePattern) GoString() string {
	return s.String()
}

func (s *CreateEventRequestRecurrencePattern) SetDayOfMonth(v int32) *CreateEventRequestRecurrencePattern {
	s.DayOfMonth = &v
	return s
}

func (s *CreateEventRequestRecurrencePattern) SetDaysOfWeek(v string) *CreateEventRequestRecurrencePattern {
	s.DaysOfWeek = &v
	return s
}

func (s *CreateEventRequestRecurrencePattern) SetIndex(v string) *CreateEventRequestRecurrencePattern {
	s.Index = &v
	return s
}

func (s *CreateEventRequestRecurrencePattern) SetInterval(v int32) *CreateEventRequestRecurrencePattern {
	s.Interval = &v
	return s
}

func (s *CreateEventRequestRecurrencePattern) SetType(v string) *CreateEventRequestRecurrencePattern {
	s.Type = &v
	return s
}

type CreateEventRequestRecurrenceRange struct {
	EndDate             *string `json:"endDate,omitempty" xml:"endDate,omitempty"`
	NumberOfOccurrences *int32  `json:"numberOfOccurrences,omitempty" xml:"numberOfOccurrences,omitempty"`
	Type                *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s CreateEventRequestRecurrenceRange) String() string {
	return tea.Prettify(s)
}

func (s CreateEventRequestRecurrenceRange) GoString() string {
	return s.String()
}

func (s *CreateEventRequestRecurrenceRange) SetEndDate(v string) *CreateEventRequestRecurrenceRange {
	s.EndDate = &v
	return s
}

func (s *CreateEventRequestRecurrenceRange) SetNumberOfOccurrences(v int32) *CreateEventRequestRecurrenceRange {
	s.NumberOfOccurrences = &v
	return s
}

func (s *CreateEventRequestRecurrenceRange) SetType(v string) *CreateEventRequestRecurrenceRange {
	s.Type = &v
	return s
}

type CreateEventRequestReminders struct {
	Method  *string `json:"method,omitempty" xml:"method,omitempty"`
	Minutes *int32  `json:"minutes,omitempty" xml:"minutes,omitempty"`
}

func (s CreateEventRequestReminders) String() string {
	return tea.Prettify(s)
}

func (s CreateEventRequestReminders) GoString() string {
	return s.String()
}

func (s *CreateEventRequestReminders) SetMethod(v string) *CreateEventRequestReminders {
	s.Method = &v
	return s
}

func (s *CreateEventRequestReminders) SetMinutes(v int32) *CreateEventRequestReminders {
	s.Minutes = &v
	return s
}

type CreateEventRequestStart struct {
	// 日程开始日期，如果是全天日程必须有值，非全天日程必须留空，格式：yyyy-MM-dd
	Date *string `json:"date,omitempty" xml:"date,omitempty"`
	// 日程开始时间，非全天日程必须有值，全天日程必须留空，格式为ISO-8601的date-time格式
	DateTime *string `json:"dateTime,omitempty" xml:"dateTime,omitempty"`
	// 日程开始时间所属时区，非全天日程必须有值，全天日程必须留空，tz database name格式，参考：https://en.wikipedia.org/wiki/List_of_tz_database_time_zones
	TimeZone *string `json:"timeZone,omitempty" xml:"timeZone,omitempty"`
}

func (s CreateEventRequestStart) String() string {
	return tea.Prettify(s)
}

func (s CreateEventRequestStart) GoString() string {
	return s.String()
}

func (s *CreateEventRequestStart) SetDate(v string) *CreateEventRequestStart {
	s.Date = &v
	return s
}

func (s *CreateEventRequestStart) SetDateTime(v string) *CreateEventRequestStart {
	s.DateTime = &v
	return s
}

func (s *CreateEventRequestStart) SetTimeZone(v string) *CreateEventRequestStart {
	s.TimeZone = &v
	return s
}

type CreateEventResponseBody struct {
	Attendees []*CreateEventResponseBodyAttendees `json:"attendees,omitempty" xml:"attendees,omitempty" type:"Repeated"`
	// 创建时间
	CreateTime        *string                                   `json:"createTime,omitempty" xml:"createTime,omitempty"`
	Description       *string                                   `json:"description,omitempty" xml:"description,omitempty"`
	End               *CreateEventResponseBodyEnd               `json:"end,omitempty" xml:"end,omitempty" type:"Struct"`
	Id                *string                                   `json:"id,omitempty" xml:"id,omitempty"`
	IsAllDay          *bool                                     `json:"isAllDay,omitempty" xml:"isAllDay,omitempty"`
	Location          *CreateEventResponseBodyLocation          `json:"location,omitempty" xml:"location,omitempty" type:"Struct"`
	OnlineMeetingInfo *CreateEventResponseBodyOnlineMeetingInfo `json:"onlineMeetingInfo,omitempty" xml:"onlineMeetingInfo,omitempty" type:"Struct"`
	Organizer         *CreateEventResponseBodyOrganizer         `json:"organizer,omitempty" xml:"organizer,omitempty" type:"Struct"`
	Recurrence        *CreateEventResponseBodyRecurrence        `json:"recurrence,omitempty" xml:"recurrence,omitempty" type:"Struct"`
	Reminders         []*CreateEventResponseBodyReminders       `json:"reminders,omitempty" xml:"reminders,omitempty" type:"Repeated"`
	// 日程开始时间
	Start   *CreateEventResponseBodyStart `json:"start,omitempty" xml:"start,omitempty" type:"Struct"`
	Summary *string                       `json:"summary,omitempty" xml:"summary,omitempty"`
	// 更新时间
	UpdateTime *string `json:"updateTime,omitempty" xml:"updateTime,omitempty"`
}

func (s CreateEventResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateEventResponseBody) GoString() string {
	return s.String()
}

func (s *CreateEventResponseBody) SetAttendees(v []*CreateEventResponseBodyAttendees) *CreateEventResponseBody {
	s.Attendees = v
	return s
}

func (s *CreateEventResponseBody) SetCreateTime(v string) *CreateEventResponseBody {
	s.CreateTime = &v
	return s
}

func (s *CreateEventResponseBody) SetDescription(v string) *CreateEventResponseBody {
	s.Description = &v
	return s
}

func (s *CreateEventResponseBody) SetEnd(v *CreateEventResponseBodyEnd) *CreateEventResponseBody {
	s.End = v
	return s
}

func (s *CreateEventResponseBody) SetId(v string) *CreateEventResponseBody {
	s.Id = &v
	return s
}

func (s *CreateEventResponseBody) SetIsAllDay(v bool) *CreateEventResponseBody {
	s.IsAllDay = &v
	return s
}

func (s *CreateEventResponseBody) SetLocation(v *CreateEventResponseBodyLocation) *CreateEventResponseBody {
	s.Location = v
	return s
}

func (s *CreateEventResponseBody) SetOnlineMeetingInfo(v *CreateEventResponseBodyOnlineMeetingInfo) *CreateEventResponseBody {
	s.OnlineMeetingInfo = v
	return s
}

func (s *CreateEventResponseBody) SetOrganizer(v *CreateEventResponseBodyOrganizer) *CreateEventResponseBody {
	s.Organizer = v
	return s
}

func (s *CreateEventResponseBody) SetRecurrence(v *CreateEventResponseBodyRecurrence) *CreateEventResponseBody {
	s.Recurrence = v
	return s
}

func (s *CreateEventResponseBody) SetReminders(v []*CreateEventResponseBodyReminders) *CreateEventResponseBody {
	s.Reminders = v
	return s
}

func (s *CreateEventResponseBody) SetStart(v *CreateEventResponseBodyStart) *CreateEventResponseBody {
	s.Start = v
	return s
}

func (s *CreateEventResponseBody) SetSummary(v string) *CreateEventResponseBody {
	s.Summary = &v
	return s
}

func (s *CreateEventResponseBody) SetUpdateTime(v string) *CreateEventResponseBody {
	s.UpdateTime = &v
	return s
}

type CreateEventResponseBodyAttendees struct {
	DisplayName *string `json:"displayName,omitempty" xml:"displayName,omitempty"`
	Id          *string `json:"id,omitempty" xml:"id,omitempty"`
	IsOptional  *bool   `json:"isOptional,omitempty" xml:"isOptional,omitempty"`
	// 回复状态
	ResponseStatus *string `json:"responseStatus,omitempty" xml:"responseStatus,omitempty"`
	Self           *bool   `json:"self,omitempty" xml:"self,omitempty"`
}

func (s CreateEventResponseBodyAttendees) String() string {
	return tea.Prettify(s)
}

func (s CreateEventResponseBodyAttendees) GoString() string {
	return s.String()
}

func (s *CreateEventResponseBodyAttendees) SetDisplayName(v string) *CreateEventResponseBodyAttendees {
	s.DisplayName = &v
	return s
}

func (s *CreateEventResponseBodyAttendees) SetId(v string) *CreateEventResponseBodyAttendees {
	s.Id = &v
	return s
}

func (s *CreateEventResponseBodyAttendees) SetIsOptional(v bool) *CreateEventResponseBodyAttendees {
	s.IsOptional = &v
	return s
}

func (s *CreateEventResponseBodyAttendees) SetResponseStatus(v string) *CreateEventResponseBodyAttendees {
	s.ResponseStatus = &v
	return s
}

func (s *CreateEventResponseBodyAttendees) SetSelf(v bool) *CreateEventResponseBodyAttendees {
	s.Self = &v
	return s
}

type CreateEventResponseBodyEnd struct {
	Date     *string `json:"date,omitempty" xml:"date,omitempty"`
	DateTime *string `json:"dateTime,omitempty" xml:"dateTime,omitempty"`
	TimeZone *string `json:"timeZone,omitempty" xml:"timeZone,omitempty"`
}

func (s CreateEventResponseBodyEnd) String() string {
	return tea.Prettify(s)
}

func (s CreateEventResponseBodyEnd) GoString() string {
	return s.String()
}

func (s *CreateEventResponseBodyEnd) SetDate(v string) *CreateEventResponseBodyEnd {
	s.Date = &v
	return s
}

func (s *CreateEventResponseBodyEnd) SetDateTime(v string) *CreateEventResponseBodyEnd {
	s.DateTime = &v
	return s
}

func (s *CreateEventResponseBodyEnd) SetTimeZone(v string) *CreateEventResponseBodyEnd {
	s.TimeZone = &v
	return s
}

type CreateEventResponseBodyLocation struct {
	DisplayName *string `json:"displayName,omitempty" xml:"displayName,omitempty"`
}

func (s CreateEventResponseBodyLocation) String() string {
	return tea.Prettify(s)
}

func (s CreateEventResponseBodyLocation) GoString() string {
	return s.String()
}

func (s *CreateEventResponseBodyLocation) SetDisplayName(v string) *CreateEventResponseBodyLocation {
	s.DisplayName = &v
	return s
}

type CreateEventResponseBodyOnlineMeetingInfo struct {
	ConferenceId *string                `json:"conferenceId,omitempty" xml:"conferenceId,omitempty"`
	ExtraInfo    map[string]interface{} `json:"extraInfo,omitempty" xml:"extraInfo,omitempty"`
	Type         *string                `json:"type,omitempty" xml:"type,omitempty"`
	Url          *string                `json:"url,omitempty" xml:"url,omitempty"`
}

func (s CreateEventResponseBodyOnlineMeetingInfo) String() string {
	return tea.Prettify(s)
}

func (s CreateEventResponseBodyOnlineMeetingInfo) GoString() string {
	return s.String()
}

func (s *CreateEventResponseBodyOnlineMeetingInfo) SetConferenceId(v string) *CreateEventResponseBodyOnlineMeetingInfo {
	s.ConferenceId = &v
	return s
}

func (s *CreateEventResponseBodyOnlineMeetingInfo) SetExtraInfo(v map[string]interface{}) *CreateEventResponseBodyOnlineMeetingInfo {
	s.ExtraInfo = v
	return s
}

func (s *CreateEventResponseBodyOnlineMeetingInfo) SetType(v string) *CreateEventResponseBodyOnlineMeetingInfo {
	s.Type = &v
	return s
}

func (s *CreateEventResponseBodyOnlineMeetingInfo) SetUrl(v string) *CreateEventResponseBodyOnlineMeetingInfo {
	s.Url = &v
	return s
}

type CreateEventResponseBodyOrganizer struct {
	// 用户名
	DisplayName *string `json:"displayName,omitempty" xml:"displayName,omitempty"`
	Id          *string `json:"id,omitempty" xml:"id,omitempty"`
	// 回复状态
	ResponseStatus *string `json:"responseStatus,omitempty" xml:"responseStatus,omitempty"`
	Self           *bool   `json:"self,omitempty" xml:"self,omitempty"`
}

func (s CreateEventResponseBodyOrganizer) String() string {
	return tea.Prettify(s)
}

func (s CreateEventResponseBodyOrganizer) GoString() string {
	return s.String()
}

func (s *CreateEventResponseBodyOrganizer) SetDisplayName(v string) *CreateEventResponseBodyOrganizer {
	s.DisplayName = &v
	return s
}

func (s *CreateEventResponseBodyOrganizer) SetId(v string) *CreateEventResponseBodyOrganizer {
	s.Id = &v
	return s
}

func (s *CreateEventResponseBodyOrganizer) SetResponseStatus(v string) *CreateEventResponseBodyOrganizer {
	s.ResponseStatus = &v
	return s
}

func (s *CreateEventResponseBodyOrganizer) SetSelf(v bool) *CreateEventResponseBodyOrganizer {
	s.Self = &v
	return s
}

type CreateEventResponseBodyRecurrence struct {
	Pattern *CreateEventResponseBodyRecurrencePattern `json:"pattern,omitempty" xml:"pattern,omitempty" type:"Struct"`
	Range   *CreateEventResponseBodyRecurrenceRange   `json:"range,omitempty" xml:"range,omitempty" type:"Struct"`
}

func (s CreateEventResponseBodyRecurrence) String() string {
	return tea.Prettify(s)
}

func (s CreateEventResponseBodyRecurrence) GoString() string {
	return s.String()
}

func (s *CreateEventResponseBodyRecurrence) SetPattern(v *CreateEventResponseBodyRecurrencePattern) *CreateEventResponseBodyRecurrence {
	s.Pattern = v
	return s
}

func (s *CreateEventResponseBodyRecurrence) SetRange(v *CreateEventResponseBodyRecurrenceRange) *CreateEventResponseBodyRecurrence {
	s.Range = v
	return s
}

type CreateEventResponseBodyRecurrencePattern struct {
	DayOfMonth *int32  `json:"dayOfMonth,omitempty" xml:"dayOfMonth,omitempty"`
	DaysOfWeek *string `json:"daysOfWeek,omitempty" xml:"daysOfWeek,omitempty"`
	Index      *string `json:"index,omitempty" xml:"index,omitempty"`
	Interval   *int32  `json:"interval,omitempty" xml:"interval,omitempty"`
	Type       *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s CreateEventResponseBodyRecurrencePattern) String() string {
	return tea.Prettify(s)
}

func (s CreateEventResponseBodyRecurrencePattern) GoString() string {
	return s.String()
}

func (s *CreateEventResponseBodyRecurrencePattern) SetDayOfMonth(v int32) *CreateEventResponseBodyRecurrencePattern {
	s.DayOfMonth = &v
	return s
}

func (s *CreateEventResponseBodyRecurrencePattern) SetDaysOfWeek(v string) *CreateEventResponseBodyRecurrencePattern {
	s.DaysOfWeek = &v
	return s
}

func (s *CreateEventResponseBodyRecurrencePattern) SetIndex(v string) *CreateEventResponseBodyRecurrencePattern {
	s.Index = &v
	return s
}

func (s *CreateEventResponseBodyRecurrencePattern) SetInterval(v int32) *CreateEventResponseBodyRecurrencePattern {
	s.Interval = &v
	return s
}

func (s *CreateEventResponseBodyRecurrencePattern) SetType(v string) *CreateEventResponseBodyRecurrencePattern {
	s.Type = &v
	return s
}

type CreateEventResponseBodyRecurrenceRange struct {
	EndDate             *string `json:"endDate,omitempty" xml:"endDate,omitempty"`
	NumberOfOccurrences *int32  `json:"numberOfOccurrences,omitempty" xml:"numberOfOccurrences,omitempty"`
	Type                *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s CreateEventResponseBodyRecurrenceRange) String() string {
	return tea.Prettify(s)
}

func (s CreateEventResponseBodyRecurrenceRange) GoString() string {
	return s.String()
}

func (s *CreateEventResponseBodyRecurrenceRange) SetEndDate(v string) *CreateEventResponseBodyRecurrenceRange {
	s.EndDate = &v
	return s
}

func (s *CreateEventResponseBodyRecurrenceRange) SetNumberOfOccurrences(v int32) *CreateEventResponseBodyRecurrenceRange {
	s.NumberOfOccurrences = &v
	return s
}

func (s *CreateEventResponseBodyRecurrenceRange) SetType(v string) *CreateEventResponseBodyRecurrenceRange {
	s.Type = &v
	return s
}

type CreateEventResponseBodyReminders struct {
	Method  *string `json:"method,omitempty" xml:"method,omitempty"`
	Minutes *string `json:"minutes,omitempty" xml:"minutes,omitempty"`
}

func (s CreateEventResponseBodyReminders) String() string {
	return tea.Prettify(s)
}

func (s CreateEventResponseBodyReminders) GoString() string {
	return s.String()
}

func (s *CreateEventResponseBodyReminders) SetMethod(v string) *CreateEventResponseBodyReminders {
	s.Method = &v
	return s
}

func (s *CreateEventResponseBodyReminders) SetMinutes(v string) *CreateEventResponseBodyReminders {
	s.Minutes = &v
	return s
}

type CreateEventResponseBodyStart struct {
	Date     *string `json:"date,omitempty" xml:"date,omitempty"`
	DateTime *string `json:"dateTime,omitempty" xml:"dateTime,omitempty"`
	TimeZone *string `json:"timeZone,omitempty" xml:"timeZone,omitempty"`
}

func (s CreateEventResponseBodyStart) String() string {
	return tea.Prettify(s)
}

func (s CreateEventResponseBodyStart) GoString() string {
	return s.String()
}

func (s *CreateEventResponseBodyStart) SetDate(v string) *CreateEventResponseBodyStart {
	s.Date = &v
	return s
}

func (s *CreateEventResponseBodyStart) SetDateTime(v string) *CreateEventResponseBodyStart {
	s.DateTime = &v
	return s
}

func (s *CreateEventResponseBodyStart) SetTimeZone(v string) *CreateEventResponseBodyStart {
	s.TimeZone = &v
	return s
}

type CreateEventResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateEventResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateEventResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateEventResponse) GoString() string {
	return s.String()
}

func (s *CreateEventResponse) SetHeaders(v map[string]*string) *CreateEventResponse {
	s.Headers = v
	return s
}

func (s *CreateEventResponse) SetBody(v *CreateEventResponseBody) *CreateEventResponse {
	s.Body = v
	return s
}

type CreateSubscribedCalendarHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CreateSubscribedCalendarHeaders) String() string {
	return tea.Prettify(s)
}

func (s CreateSubscribedCalendarHeaders) GoString() string {
	return s.String()
}

func (s *CreateSubscribedCalendarHeaders) SetCommonHeaders(v map[string]*string) *CreateSubscribedCalendarHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreateSubscribedCalendarHeaders) SetXAcsDingtalkAccessToken(v string) *CreateSubscribedCalendarHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CreateSubscribedCalendarRequest struct {
	// 日历介绍
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// 日历管理员列表
	Managers []*string `json:"managers,omitempty" xml:"managers,omitempty" type:"Repeated"`
	// 日历名
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 可订阅列表
	SubscribeScope *CreateSubscribedCalendarRequestSubscribeScope `json:"subscribeScope,omitempty" xml:"subscribeScope,omitempty" type:"Struct"`
}

func (s CreateSubscribedCalendarRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateSubscribedCalendarRequest) GoString() string {
	return s.String()
}

func (s *CreateSubscribedCalendarRequest) SetDescription(v string) *CreateSubscribedCalendarRequest {
	s.Description = &v
	return s
}

func (s *CreateSubscribedCalendarRequest) SetManagers(v []*string) *CreateSubscribedCalendarRequest {
	s.Managers = v
	return s
}

func (s *CreateSubscribedCalendarRequest) SetName(v string) *CreateSubscribedCalendarRequest {
	s.Name = &v
	return s
}

func (s *CreateSubscribedCalendarRequest) SetSubscribeScope(v *CreateSubscribedCalendarRequestSubscribeScope) *CreateSubscribedCalendarRequest {
	s.SubscribeScope = v
	return s
}

type CreateSubscribedCalendarRequestSubscribeScope struct {
	// 可订阅组织列表
	CorpIds []*string `json:"corpIds,omitempty" xml:"corpIds,omitempty" type:"Repeated"`
	// 可订阅群组列表
	OpenConversationIds []*string `json:"openConversationIds,omitempty" xml:"openConversationIds,omitempty" type:"Repeated"`
	// 可订阅人员列表
	UnionIds []*string `json:"unionIds,omitempty" xml:"unionIds,omitempty" type:"Repeated"`
}

func (s CreateSubscribedCalendarRequestSubscribeScope) String() string {
	return tea.Prettify(s)
}

func (s CreateSubscribedCalendarRequestSubscribeScope) GoString() string {
	return s.String()
}

func (s *CreateSubscribedCalendarRequestSubscribeScope) SetCorpIds(v []*string) *CreateSubscribedCalendarRequestSubscribeScope {
	s.CorpIds = v
	return s
}

func (s *CreateSubscribedCalendarRequestSubscribeScope) SetOpenConversationIds(v []*string) *CreateSubscribedCalendarRequestSubscribeScope {
	s.OpenConversationIds = v
	return s
}

func (s *CreateSubscribedCalendarRequestSubscribeScope) SetUnionIds(v []*string) *CreateSubscribedCalendarRequestSubscribeScope {
	s.UnionIds = v
	return s
}

type CreateSubscribedCalendarResponseBody struct {
	// 日历id
	CalendarId *string `json:"calendarId,omitempty" xml:"calendarId,omitempty"`
}

func (s CreateSubscribedCalendarResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateSubscribedCalendarResponseBody) GoString() string {
	return s.String()
}

func (s *CreateSubscribedCalendarResponseBody) SetCalendarId(v string) *CreateSubscribedCalendarResponseBody {
	s.CalendarId = &v
	return s
}

type CreateSubscribedCalendarResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateSubscribedCalendarResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateSubscribedCalendarResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateSubscribedCalendarResponse) GoString() string {
	return s.String()
}

func (s *CreateSubscribedCalendarResponse) SetHeaders(v map[string]*string) *CreateSubscribedCalendarResponse {
	s.Headers = v
	return s
}

func (s *CreateSubscribedCalendarResponse) SetBody(v *CreateSubscribedCalendarResponseBody) *CreateSubscribedCalendarResponse {
	s.Body = v
	return s
}

type DeleteAclHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s DeleteAclHeaders) String() string {
	return tea.Prettify(s)
}

func (s DeleteAclHeaders) GoString() string {
	return s.String()
}

func (s *DeleteAclHeaders) SetCommonHeaders(v map[string]*string) *DeleteAclHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DeleteAclHeaders) SetXAcsDingtalkAccessToken(v string) *DeleteAclHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type DeleteAclResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
}

func (s DeleteAclResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteAclResponse) GoString() string {
	return s.String()
}

func (s *DeleteAclResponse) SetHeaders(v map[string]*string) *DeleteAclResponse {
	s.Headers = v
	return s
}

type DeleteEventHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s DeleteEventHeaders) String() string {
	return tea.Prettify(s)
}

func (s DeleteEventHeaders) GoString() string {
	return s.String()
}

func (s *DeleteEventHeaders) SetCommonHeaders(v map[string]*string) *DeleteEventHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DeleteEventHeaders) SetXAcsDingtalkAccessToken(v string) *DeleteEventHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type DeleteEventResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
}

func (s DeleteEventResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteEventResponse) GoString() string {
	return s.String()
}

func (s *DeleteEventResponse) SetHeaders(v map[string]*string) *DeleteEventResponse {
	s.Headers = v
	return s
}

type DeleteSubscribedCalendarHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s DeleteSubscribedCalendarHeaders) String() string {
	return tea.Prettify(s)
}

func (s DeleteSubscribedCalendarHeaders) GoString() string {
	return s.String()
}

func (s *DeleteSubscribedCalendarHeaders) SetCommonHeaders(v map[string]*string) *DeleteSubscribedCalendarHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DeleteSubscribedCalendarHeaders) SetXAcsDingtalkAccessToken(v string) *DeleteSubscribedCalendarHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type DeleteSubscribedCalendarResponseBody struct {
	Result *bool `json:"result,omitempty" xml:"result,omitempty"`
}

func (s DeleteSubscribedCalendarResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteSubscribedCalendarResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteSubscribedCalendarResponseBody) SetResult(v bool) *DeleteSubscribedCalendarResponseBody {
	s.Result = &v
	return s
}

type DeleteSubscribedCalendarResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteSubscribedCalendarResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteSubscribedCalendarResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteSubscribedCalendarResponse) GoString() string {
	return s.String()
}

func (s *DeleteSubscribedCalendarResponse) SetHeaders(v map[string]*string) *DeleteSubscribedCalendarResponse {
	s.Headers = v
	return s
}

func (s *DeleteSubscribedCalendarResponse) SetBody(v *DeleteSubscribedCalendarResponseBody) *DeleteSubscribedCalendarResponse {
	s.Body = v
	return s
}

type GenerateCaldavAccountHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// 授权本次调用的用户id，该字段有值时认为本次调用已被授权访问该用户可以访问的所有数据
	DingUid                 *string `json:"dingUid,omitempty" xml:"dingUid,omitempty"`
	XAcsDingtalkAccessToken *string `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GenerateCaldavAccountHeaders) String() string {
	return tea.Prettify(s)
}

func (s GenerateCaldavAccountHeaders) GoString() string {
	return s.String()
}

func (s *GenerateCaldavAccountHeaders) SetCommonHeaders(v map[string]*string) *GenerateCaldavAccountHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GenerateCaldavAccountHeaders) SetDingUid(v string) *GenerateCaldavAccountHeaders {
	s.DingUid = &v
	return s
}

func (s *GenerateCaldavAccountHeaders) SetXAcsDingtalkAccessToken(v string) *GenerateCaldavAccountHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GenerateCaldavAccountRequest struct {
	// 设备名称
	Device *string `json:"device,omitempty" xml:"device,omitempty"`
}

func (s GenerateCaldavAccountRequest) String() string {
	return tea.Prettify(s)
}

func (s GenerateCaldavAccountRequest) GoString() string {
	return s.String()
}

func (s *GenerateCaldavAccountRequest) SetDevice(v string) *GenerateCaldavAccountRequest {
	s.Device = &v
	return s
}

type GenerateCaldavAccountResponseBody struct {
	Password      *string `json:"password,omitempty" xml:"password,omitempty"`
	ServerAddress *string `json:"serverAddress,omitempty" xml:"serverAddress,omitempty"`
	Username      *string `json:"username,omitempty" xml:"username,omitempty"`
}

func (s GenerateCaldavAccountResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GenerateCaldavAccountResponseBody) GoString() string {
	return s.String()
}

func (s *GenerateCaldavAccountResponseBody) SetPassword(v string) *GenerateCaldavAccountResponseBody {
	s.Password = &v
	return s
}

func (s *GenerateCaldavAccountResponseBody) SetServerAddress(v string) *GenerateCaldavAccountResponseBody {
	s.ServerAddress = &v
	return s
}

func (s *GenerateCaldavAccountResponseBody) SetUsername(v string) *GenerateCaldavAccountResponseBody {
	s.Username = &v
	return s
}

type GenerateCaldavAccountResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GenerateCaldavAccountResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GenerateCaldavAccountResponse) String() string {
	return tea.Prettify(s)
}

func (s GenerateCaldavAccountResponse) GoString() string {
	return s.String()
}

func (s *GenerateCaldavAccountResponse) SetHeaders(v map[string]*string) *GenerateCaldavAccountResponse {
	s.Headers = v
	return s
}

func (s *GenerateCaldavAccountResponse) SetBody(v *GenerateCaldavAccountResponseBody) *GenerateCaldavAccountResponse {
	s.Body = v
	return s
}

type GetEventHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetEventHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetEventHeaders) GoString() string {
	return s.String()
}

func (s *GetEventHeaders) SetCommonHeaders(v map[string]*string) *GetEventHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetEventHeaders) SetXAcsDingtalkAccessToken(v string) *GetEventHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetEventRequest struct {
	// 返回参与人，上限500人，默认为0
	MaxAttendees *int64 `json:"maxAttendees,omitempty" xml:"maxAttendees,omitempty"`
}

func (s GetEventRequest) String() string {
	return tea.Prettify(s)
}

func (s GetEventRequest) GoString() string {
	return s.String()
}

func (s *GetEventRequest) SetMaxAttendees(v int64) *GetEventRequest {
	s.MaxAttendees = &v
	return s
}

type GetEventResponseBody struct {
	Attendees []*GetEventResponseBodyAttendees `json:"attendees,omitempty" xml:"attendees,omitempty" type:"Repeated"`
	// 创建时间
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// 日程描述
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// 日程结束时间
	End                *GetEventResponseBodyEnd                `json:"end,omitempty" xml:"end,omitempty" type:"Struct"`
	ExtendedProperties *GetEventResponseBodyExtendedProperties `json:"extendedProperties,omitempty" xml:"extendedProperties,omitempty" type:"Struct"`
	Id                 *string                                 `json:"id,omitempty" xml:"id,omitempty"`
	// 是否为全天日程
	IsAllDay          *bool                                  `json:"isAllDay,omitempty" xml:"isAllDay,omitempty"`
	Location          *GetEventResponseBodyLocation          `json:"location,omitempty" xml:"location,omitempty" type:"Struct"`
	OnlineMeetingInfo *GetEventResponseBodyOnlineMeetingInfo `json:"onlineMeetingInfo,omitempty" xml:"onlineMeetingInfo,omitempty" type:"Struct"`
	Organizer         *GetEventResponseBodyOrganizer         `json:"organizer,omitempty" xml:"organizer,omitempty" type:"Struct"`
	Recurrence        *GetEventResponseBodyRecurrence        `json:"recurrence,omitempty" xml:"recurrence,omitempty" type:"Struct"`
	Reminders         []*GetEventResponseBodyReminders       `json:"reminders,omitempty" xml:"reminders,omitempty" type:"Repeated"`
	// 重复日程的主日程id，非重复日程为空
	SeriesMasterId *string `json:"seriesMasterId,omitempty" xml:"seriesMasterId,omitempty"`
	// 日程开始时间
	Start *GetEventResponseBodyStart `json:"start,omitempty" xml:"start,omitempty" type:"Struct"`
	// 日程状态
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// 日程标题
	Summary *string `json:"summary,omitempty" xml:"summary,omitempty"`
	// 更新时间
	UpdateTime *string `json:"updateTime,omitempty" xml:"updateTime,omitempty"`
}

func (s GetEventResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetEventResponseBody) GoString() string {
	return s.String()
}

func (s *GetEventResponseBody) SetAttendees(v []*GetEventResponseBodyAttendees) *GetEventResponseBody {
	s.Attendees = v
	return s
}

func (s *GetEventResponseBody) SetCreateTime(v string) *GetEventResponseBody {
	s.CreateTime = &v
	return s
}

func (s *GetEventResponseBody) SetDescription(v string) *GetEventResponseBody {
	s.Description = &v
	return s
}

func (s *GetEventResponseBody) SetEnd(v *GetEventResponseBodyEnd) *GetEventResponseBody {
	s.End = v
	return s
}

func (s *GetEventResponseBody) SetExtendedProperties(v *GetEventResponseBodyExtendedProperties) *GetEventResponseBody {
	s.ExtendedProperties = v
	return s
}

func (s *GetEventResponseBody) SetId(v string) *GetEventResponseBody {
	s.Id = &v
	return s
}

func (s *GetEventResponseBody) SetIsAllDay(v bool) *GetEventResponseBody {
	s.IsAllDay = &v
	return s
}

func (s *GetEventResponseBody) SetLocation(v *GetEventResponseBodyLocation) *GetEventResponseBody {
	s.Location = v
	return s
}

func (s *GetEventResponseBody) SetOnlineMeetingInfo(v *GetEventResponseBodyOnlineMeetingInfo) *GetEventResponseBody {
	s.OnlineMeetingInfo = v
	return s
}

func (s *GetEventResponseBody) SetOrganizer(v *GetEventResponseBodyOrganizer) *GetEventResponseBody {
	s.Organizer = v
	return s
}

func (s *GetEventResponseBody) SetRecurrence(v *GetEventResponseBodyRecurrence) *GetEventResponseBody {
	s.Recurrence = v
	return s
}

func (s *GetEventResponseBody) SetReminders(v []*GetEventResponseBodyReminders) *GetEventResponseBody {
	s.Reminders = v
	return s
}

func (s *GetEventResponseBody) SetSeriesMasterId(v string) *GetEventResponseBody {
	s.SeriesMasterId = &v
	return s
}

func (s *GetEventResponseBody) SetStart(v *GetEventResponseBodyStart) *GetEventResponseBody {
	s.Start = v
	return s
}

func (s *GetEventResponseBody) SetStatus(v string) *GetEventResponseBody {
	s.Status = &v
	return s
}

func (s *GetEventResponseBody) SetSummary(v string) *GetEventResponseBody {
	s.Summary = &v
	return s
}

func (s *GetEventResponseBody) SetUpdateTime(v string) *GetEventResponseBody {
	s.UpdateTime = &v
	return s
}

type GetEventResponseBodyAttendees struct {
	// 用户名
	DisplayName *string `json:"displayName,omitempty" xml:"displayName,omitempty"`
	Id          *string `json:"id,omitempty" xml:"id,omitempty"`
	IsOptional  *bool   `json:"isOptional,omitempty" xml:"isOptional,omitempty"`
	// 回复状态
	ResponseStatus *string `json:"responseStatus,omitempty" xml:"responseStatus,omitempty"`
	// 是否是当前登陆用户
	Self *bool `json:"self,omitempty" xml:"self,omitempty"`
}

func (s GetEventResponseBodyAttendees) String() string {
	return tea.Prettify(s)
}

func (s GetEventResponseBodyAttendees) GoString() string {
	return s.String()
}

func (s *GetEventResponseBodyAttendees) SetDisplayName(v string) *GetEventResponseBodyAttendees {
	s.DisplayName = &v
	return s
}

func (s *GetEventResponseBodyAttendees) SetId(v string) *GetEventResponseBodyAttendees {
	s.Id = &v
	return s
}

func (s *GetEventResponseBodyAttendees) SetIsOptional(v bool) *GetEventResponseBodyAttendees {
	s.IsOptional = &v
	return s
}

func (s *GetEventResponseBodyAttendees) SetResponseStatus(v string) *GetEventResponseBodyAttendees {
	s.ResponseStatus = &v
	return s
}

func (s *GetEventResponseBodyAttendees) SetSelf(v bool) *GetEventResponseBodyAttendees {
	s.Self = &v
	return s
}

type GetEventResponseBodyEnd struct {
	Date     *string `json:"date,omitempty" xml:"date,omitempty"`
	DateTime *string `json:"dateTime,omitempty" xml:"dateTime,omitempty"`
	TimeZone *string `json:"timeZone,omitempty" xml:"timeZone,omitempty"`
}

func (s GetEventResponseBodyEnd) String() string {
	return tea.Prettify(s)
}

func (s GetEventResponseBodyEnd) GoString() string {
	return s.String()
}

func (s *GetEventResponseBodyEnd) SetDate(v string) *GetEventResponseBodyEnd {
	s.Date = &v
	return s
}

func (s *GetEventResponseBodyEnd) SetDateTime(v string) *GetEventResponseBodyEnd {
	s.DateTime = &v
	return s
}

func (s *GetEventResponseBodyEnd) SetTimeZone(v string) *GetEventResponseBodyEnd {
	s.TimeZone = &v
	return s
}

type GetEventResponseBodyExtendedProperties struct {
	SharedProperties *GetEventResponseBodyExtendedPropertiesSharedProperties `json:"sharedProperties,omitempty" xml:"sharedProperties,omitempty" type:"Struct"`
}

func (s GetEventResponseBodyExtendedProperties) String() string {
	return tea.Prettify(s)
}

func (s GetEventResponseBodyExtendedProperties) GoString() string {
	return s.String()
}

func (s *GetEventResponseBodyExtendedProperties) SetSharedProperties(v *GetEventResponseBodyExtendedPropertiesSharedProperties) *GetEventResponseBodyExtendedProperties {
	s.SharedProperties = v
	return s
}

type GetEventResponseBodyExtendedPropertiesSharedProperties struct {
	SourceOpenCid *string `json:"sourceOpenCid,omitempty" xml:"sourceOpenCid,omitempty"`
}

func (s GetEventResponseBodyExtendedPropertiesSharedProperties) String() string {
	return tea.Prettify(s)
}

func (s GetEventResponseBodyExtendedPropertiesSharedProperties) GoString() string {
	return s.String()
}

func (s *GetEventResponseBodyExtendedPropertiesSharedProperties) SetSourceOpenCid(v string) *GetEventResponseBodyExtendedPropertiesSharedProperties {
	s.SourceOpenCid = &v
	return s
}

type GetEventResponseBodyLocation struct {
	DisplayName  *string   `json:"displayName,omitempty" xml:"displayName,omitempty"`
	MeetingRooms []*string `json:"meetingRooms,omitempty" xml:"meetingRooms,omitempty" type:"Repeated"`
}

func (s GetEventResponseBodyLocation) String() string {
	return tea.Prettify(s)
}

func (s GetEventResponseBodyLocation) GoString() string {
	return s.String()
}

func (s *GetEventResponseBodyLocation) SetDisplayName(v string) *GetEventResponseBodyLocation {
	s.DisplayName = &v
	return s
}

func (s *GetEventResponseBodyLocation) SetMeetingRooms(v []*string) *GetEventResponseBodyLocation {
	s.MeetingRooms = v
	return s
}

type GetEventResponseBodyOnlineMeetingInfo struct {
	ConferenceId *string                `json:"conferenceId,omitempty" xml:"conferenceId,omitempty"`
	ExtraInfo    map[string]interface{} `json:"extraInfo,omitempty" xml:"extraInfo,omitempty"`
	Type         *string                `json:"type,omitempty" xml:"type,omitempty"`
	Url          *string                `json:"url,omitempty" xml:"url,omitempty"`
}

func (s GetEventResponseBodyOnlineMeetingInfo) String() string {
	return tea.Prettify(s)
}

func (s GetEventResponseBodyOnlineMeetingInfo) GoString() string {
	return s.String()
}

func (s *GetEventResponseBodyOnlineMeetingInfo) SetConferenceId(v string) *GetEventResponseBodyOnlineMeetingInfo {
	s.ConferenceId = &v
	return s
}

func (s *GetEventResponseBodyOnlineMeetingInfo) SetExtraInfo(v map[string]interface{}) *GetEventResponseBodyOnlineMeetingInfo {
	s.ExtraInfo = v
	return s
}

func (s *GetEventResponseBodyOnlineMeetingInfo) SetType(v string) *GetEventResponseBodyOnlineMeetingInfo {
	s.Type = &v
	return s
}

func (s *GetEventResponseBodyOnlineMeetingInfo) SetUrl(v string) *GetEventResponseBodyOnlineMeetingInfo {
	s.Url = &v
	return s
}

type GetEventResponseBodyOrganizer struct {
	// 用户名
	DisplayName *string `json:"displayName,omitempty" xml:"displayName,omitempty"`
	Id          *string `json:"id,omitempty" xml:"id,omitempty"`
	// 回复状态
	ResponseStatus *string `json:"responseStatus,omitempty" xml:"responseStatus,omitempty"`
	// 是否是当前登陆用户
	Self *bool `json:"self,omitempty" xml:"self,omitempty"`
}

func (s GetEventResponseBodyOrganizer) String() string {
	return tea.Prettify(s)
}

func (s GetEventResponseBodyOrganizer) GoString() string {
	return s.String()
}

func (s *GetEventResponseBodyOrganizer) SetDisplayName(v string) *GetEventResponseBodyOrganizer {
	s.DisplayName = &v
	return s
}

func (s *GetEventResponseBodyOrganizer) SetId(v string) *GetEventResponseBodyOrganizer {
	s.Id = &v
	return s
}

func (s *GetEventResponseBodyOrganizer) SetResponseStatus(v string) *GetEventResponseBodyOrganizer {
	s.ResponseStatus = &v
	return s
}

func (s *GetEventResponseBodyOrganizer) SetSelf(v bool) *GetEventResponseBodyOrganizer {
	s.Self = &v
	return s
}

type GetEventResponseBodyRecurrence struct {
	// 重复模式
	Pattern *GetEventResponseBodyRecurrencePattern `json:"pattern,omitempty" xml:"pattern,omitempty" type:"Struct"`
	// 重复范围
	Range *GetEventResponseBodyRecurrenceRange `json:"range,omitempty" xml:"range,omitempty" type:"Struct"`
}

func (s GetEventResponseBodyRecurrence) String() string {
	return tea.Prettify(s)
}

func (s GetEventResponseBodyRecurrence) GoString() string {
	return s.String()
}

func (s *GetEventResponseBodyRecurrence) SetPattern(v *GetEventResponseBodyRecurrencePattern) *GetEventResponseBodyRecurrence {
	s.Pattern = v
	return s
}

func (s *GetEventResponseBodyRecurrence) SetRange(v *GetEventResponseBodyRecurrenceRange) *GetEventResponseBodyRecurrence {
	s.Range = v
	return s
}

type GetEventResponseBodyRecurrencePattern struct {
	DayOfMonth *int32  `json:"dayOfMonth,omitempty" xml:"dayOfMonth,omitempty"`
	DaysOfWeek *string `json:"daysOfWeek,omitempty" xml:"daysOfWeek,omitempty"`
	Index      *string `json:"index,omitempty" xml:"index,omitempty"`
	Interval   *int32  `json:"interval,omitempty" xml:"interval,omitempty"`
	// 循环模式类型(type: daily, weekly, absoluteMonthly, relativeMonthly, absoluteYearly, relativeYearly)
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s GetEventResponseBodyRecurrencePattern) String() string {
	return tea.Prettify(s)
}

func (s GetEventResponseBodyRecurrencePattern) GoString() string {
	return s.String()
}

func (s *GetEventResponseBodyRecurrencePattern) SetDayOfMonth(v int32) *GetEventResponseBodyRecurrencePattern {
	s.DayOfMonth = &v
	return s
}

func (s *GetEventResponseBodyRecurrencePattern) SetDaysOfWeek(v string) *GetEventResponseBodyRecurrencePattern {
	s.DaysOfWeek = &v
	return s
}

func (s *GetEventResponseBodyRecurrencePattern) SetIndex(v string) *GetEventResponseBodyRecurrencePattern {
	s.Index = &v
	return s
}

func (s *GetEventResponseBodyRecurrencePattern) SetInterval(v int32) *GetEventResponseBodyRecurrencePattern {
	s.Interval = &v
	return s
}

func (s *GetEventResponseBodyRecurrencePattern) SetType(v string) *GetEventResponseBodyRecurrencePattern {
	s.Type = &v
	return s
}

type GetEventResponseBodyRecurrenceRange struct {
	EndDate             *string `json:"endDate,omitempty" xml:"endDate,omitempty"`
	NumberOfOccurrences *int32  `json:"numberOfOccurrences,omitempty" xml:"numberOfOccurrences,omitempty"`
	// 范围类型(endDate, noEnd, numbered)
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s GetEventResponseBodyRecurrenceRange) String() string {
	return tea.Prettify(s)
}

func (s GetEventResponseBodyRecurrenceRange) GoString() string {
	return s.String()
}

func (s *GetEventResponseBodyRecurrenceRange) SetEndDate(v string) *GetEventResponseBodyRecurrenceRange {
	s.EndDate = &v
	return s
}

func (s *GetEventResponseBodyRecurrenceRange) SetNumberOfOccurrences(v int32) *GetEventResponseBodyRecurrenceRange {
	s.NumberOfOccurrences = &v
	return s
}

func (s *GetEventResponseBodyRecurrenceRange) SetType(v string) *GetEventResponseBodyRecurrenceRange {
	s.Type = &v
	return s
}

type GetEventResponseBodyReminders struct {
	Method  *string `json:"method,omitempty" xml:"method,omitempty"`
	Minutes *string `json:"minutes,omitempty" xml:"minutes,omitempty"`
}

func (s GetEventResponseBodyReminders) String() string {
	return tea.Prettify(s)
}

func (s GetEventResponseBodyReminders) GoString() string {
	return s.String()
}

func (s *GetEventResponseBodyReminders) SetMethod(v string) *GetEventResponseBodyReminders {
	s.Method = &v
	return s
}

func (s *GetEventResponseBodyReminders) SetMinutes(v string) *GetEventResponseBodyReminders {
	s.Minutes = &v
	return s
}

type GetEventResponseBodyStart struct {
	// 日期，格式：yyyyMMdd
	Date *string `json:"date,omitempty" xml:"date,omitempty"`
	// 时间戳，按照ISO 8601格式
	DateTime *string `json:"dateTime,omitempty" xml:"dateTime,omitempty"`
	// 时区
	TimeZone *string `json:"timeZone,omitempty" xml:"timeZone,omitempty"`
}

func (s GetEventResponseBodyStart) String() string {
	return tea.Prettify(s)
}

func (s GetEventResponseBodyStart) GoString() string {
	return s.String()
}

func (s *GetEventResponseBodyStart) SetDate(v string) *GetEventResponseBodyStart {
	s.Date = &v
	return s
}

func (s *GetEventResponseBodyStart) SetDateTime(v string) *GetEventResponseBodyStart {
	s.DateTime = &v
	return s
}

func (s *GetEventResponseBodyStart) SetTimeZone(v string) *GetEventResponseBodyStart {
	s.TimeZone = &v
	return s
}

type GetEventResponse struct {
	Headers map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetEventResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetEventResponse) String() string {
	return tea.Prettify(s)
}

func (s GetEventResponse) GoString() string {
	return s.String()
}

func (s *GetEventResponse) SetHeaders(v map[string]*string) *GetEventResponse {
	s.Headers = v
	return s
}

func (s *GetEventResponse) SetBody(v *GetEventResponseBody) *GetEventResponse {
	s.Body = v
	return s
}

type GetScheduleHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetScheduleHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetScheduleHeaders) GoString() string {
	return s.String()
}

func (s *GetScheduleHeaders) SetCommonHeaders(v map[string]*string) *GetScheduleHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetScheduleHeaders) SetXAcsDingtalkAccessToken(v string) *GetScheduleHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetScheduleRequest struct {
	// 查询结束时间
	EndTime *string `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// 查询开始时间
	StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty"`
	// 待查询的用户列表
	UserIds []*string `json:"userIds,omitempty" xml:"userIds,omitempty" type:"Repeated"`
}

func (s GetScheduleRequest) String() string {
	return tea.Prettify(s)
}

func (s GetScheduleRequest) GoString() string {
	return s.String()
}

func (s *GetScheduleRequest) SetEndTime(v string) *GetScheduleRequest {
	s.EndTime = &v
	return s
}

func (s *GetScheduleRequest) SetStartTime(v string) *GetScheduleRequest {
	s.StartTime = &v
	return s
}

func (s *GetScheduleRequest) SetUserIds(v []*string) *GetScheduleRequest {
	s.UserIds = v
	return s
}

type GetScheduleResponseBody struct {
	// 闲忙信息
	ScheduleInformation []*GetScheduleResponseBodyScheduleInformation `json:"scheduleInformation,omitempty" xml:"scheduleInformation,omitempty" type:"Repeated"`
}

func (s GetScheduleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetScheduleResponseBody) GoString() string {
	return s.String()
}

func (s *GetScheduleResponseBody) SetScheduleInformation(v []*GetScheduleResponseBodyScheduleInformation) *GetScheduleResponseBody {
	s.ScheduleInformation = v
	return s
}

type GetScheduleResponseBodyScheduleInformation struct {
	// 异常描述
	Error         *string                                                    `json:"error,omitempty" xml:"error,omitempty"`
	ScheduleItems []*GetScheduleResponseBodyScheduleInformationScheduleItems `json:"scheduleItems,omitempty" xml:"scheduleItems,omitempty" type:"Repeated"`
	// 用户userId
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s GetScheduleResponseBodyScheduleInformation) String() string {
	return tea.Prettify(s)
}

func (s GetScheduleResponseBodyScheduleInformation) GoString() string {
	return s.String()
}

func (s *GetScheduleResponseBodyScheduleInformation) SetError(v string) *GetScheduleResponseBodyScheduleInformation {
	s.Error = &v
	return s
}

func (s *GetScheduleResponseBodyScheduleInformation) SetScheduleItems(v []*GetScheduleResponseBodyScheduleInformationScheduleItems) *GetScheduleResponseBodyScheduleInformation {
	s.ScheduleItems = v
	return s
}

func (s *GetScheduleResponseBodyScheduleInformation) SetUserId(v string) *GetScheduleResponseBodyScheduleInformation {
	s.UserId = &v
	return s
}

type GetScheduleResponseBodyScheduleInformationScheduleItems struct {
	// 结束时间，表示一个日期，或者一个带时区的时间戳
	End *GetScheduleResponseBodyScheduleInformationScheduleItemsEnd `json:"end,omitempty" xml:"end,omitempty" type:"Struct"`
	// 开始时间，表示一个日期，或者一个带时区的时间戳
	Start *GetScheduleResponseBodyScheduleInformationScheduleItemsStart `json:"start,omitempty" xml:"start,omitempty" type:"Struct"`
	// 状态: - BUSY：繁忙, - TENTATIVE：暂定繁忙
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s GetScheduleResponseBodyScheduleInformationScheduleItems) String() string {
	return tea.Prettify(s)
}

func (s GetScheduleResponseBodyScheduleInformationScheduleItems) GoString() string {
	return s.String()
}

func (s *GetScheduleResponseBodyScheduleInformationScheduleItems) SetEnd(v *GetScheduleResponseBodyScheduleInformationScheduleItemsEnd) *GetScheduleResponseBodyScheduleInformationScheduleItems {
	s.End = v
	return s
}

func (s *GetScheduleResponseBodyScheduleInformationScheduleItems) SetStart(v *GetScheduleResponseBodyScheduleInformationScheduleItemsStart) *GetScheduleResponseBodyScheduleInformationScheduleItems {
	s.Start = v
	return s
}

func (s *GetScheduleResponseBodyScheduleInformationScheduleItems) SetStatus(v string) *GetScheduleResponseBodyScheduleInformationScheduleItems {
	s.Status = &v
	return s
}

type GetScheduleResponseBodyScheduleInformationScheduleItemsEnd struct {
	// 结束日期
	Date *string `json:"date,omitempty" xml:"date,omitempty"`
	// 结束时间戳，按照ISO 8601格式
	DateTime *string `json:"dateTime,omitempty" xml:"dateTime,omitempty"`
	// 时间戳所属时区
	TimeZone *string `json:"timeZone,omitempty" xml:"timeZone,omitempty"`
}

func (s GetScheduleResponseBodyScheduleInformationScheduleItemsEnd) String() string {
	return tea.Prettify(s)
}

func (s GetScheduleResponseBodyScheduleInformationScheduleItemsEnd) GoString() string {
	return s.String()
}

func (s *GetScheduleResponseBodyScheduleInformationScheduleItemsEnd) SetDate(v string) *GetScheduleResponseBodyScheduleInformationScheduleItemsEnd {
	s.Date = &v
	return s
}

func (s *GetScheduleResponseBodyScheduleInformationScheduleItemsEnd) SetDateTime(v string) *GetScheduleResponseBodyScheduleInformationScheduleItemsEnd {
	s.DateTime = &v
	return s
}

func (s *GetScheduleResponseBodyScheduleInformationScheduleItemsEnd) SetTimeZone(v string) *GetScheduleResponseBodyScheduleInformationScheduleItemsEnd {
	s.TimeZone = &v
	return s
}

type GetScheduleResponseBodyScheduleInformationScheduleItemsStart struct {
	// 开始日期
	Date *string `json:"date,omitempty" xml:"date,omitempty"`
	// 开始时间戳，按照ISO 8601格式
	DateTime *string `json:"dateTime,omitempty" xml:"dateTime,omitempty"`
	// 所属时区
	TimeZone *string `json:"timeZone,omitempty" xml:"timeZone,omitempty"`
}

func (s GetScheduleResponseBodyScheduleInformationScheduleItemsStart) String() string {
	return tea.Prettify(s)
}

func (s GetScheduleResponseBodyScheduleInformationScheduleItemsStart) GoString() string {
	return s.String()
}

func (s *GetScheduleResponseBodyScheduleInformationScheduleItemsStart) SetDate(v string) *GetScheduleResponseBodyScheduleInformationScheduleItemsStart {
	s.Date = &v
	return s
}

func (s *GetScheduleResponseBodyScheduleInformationScheduleItemsStart) SetDateTime(v string) *GetScheduleResponseBodyScheduleInformationScheduleItemsStart {
	s.DateTime = &v
	return s
}

func (s *GetScheduleResponseBodyScheduleInformationScheduleItemsStart) SetTimeZone(v string) *GetScheduleResponseBodyScheduleInformationScheduleItemsStart {
	s.TimeZone = &v
	return s
}

type GetScheduleResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetScheduleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetScheduleResponse) String() string {
	return tea.Prettify(s)
}

func (s GetScheduleResponse) GoString() string {
	return s.String()
}

func (s *GetScheduleResponse) SetHeaders(v map[string]*string) *GetScheduleResponse {
	s.Headers = v
	return s
}

func (s *GetScheduleResponse) SetBody(v *GetScheduleResponseBody) *GetScheduleResponse {
	s.Body = v
	return s
}

type GetSignInListHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetSignInListHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetSignInListHeaders) GoString() string {
	return s.String()
}

func (s *GetSignInListHeaders) SetCommonHeaders(v map[string]*string) *GetSignInListHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetSignInListHeaders) SetXAcsDingtalkAccessToken(v string) *GetSignInListHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetSignInListRequest struct {
	// 查询返回结果数（上限200）
	MaxResults *int32  `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	NextToken  *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// 签到信息类型（check_in，not_yet_check_in)
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s GetSignInListRequest) String() string {
	return tea.Prettify(s)
}

func (s GetSignInListRequest) GoString() string {
	return s.String()
}

func (s *GetSignInListRequest) SetMaxResults(v int32) *GetSignInListRequest {
	s.MaxResults = &v
	return s
}

func (s *GetSignInListRequest) SetNextToken(v string) *GetSignInListRequest {
	s.NextToken = &v
	return s
}

func (s *GetSignInListRequest) SetType(v string) *GetSignInListRequest {
	s.Type = &v
	return s
}

type GetSignInListResponseBody struct {
	// 翻页token
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// 签到信息
	Users []*GetSignInListResponseBodyUsers `json:"users,omitempty" xml:"users,omitempty" type:"Repeated"`
}

func (s GetSignInListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetSignInListResponseBody) GoString() string {
	return s.String()
}

func (s *GetSignInListResponseBody) SetNextToken(v string) *GetSignInListResponseBody {
	s.NextToken = &v
	return s
}

func (s *GetSignInListResponseBody) SetUsers(v []*GetSignInListResponseBodyUsers) *GetSignInListResponseBody {
	s.Users = v
	return s
}

type GetSignInListResponseBodyUsers struct {
	// 签到时间
	CheckInTime *int64 `json:"checkInTime,omitempty" xml:"checkInTime,omitempty"`
	// 用户名
	DisplayName *string `json:"displayName,omitempty" xml:"displayName,omitempty"`
	UserId      *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s GetSignInListResponseBodyUsers) String() string {
	return tea.Prettify(s)
}

func (s GetSignInListResponseBodyUsers) GoString() string {
	return s.String()
}

func (s *GetSignInListResponseBodyUsers) SetCheckInTime(v int64) *GetSignInListResponseBodyUsers {
	s.CheckInTime = &v
	return s
}

func (s *GetSignInListResponseBodyUsers) SetDisplayName(v string) *GetSignInListResponseBodyUsers {
	s.DisplayName = &v
	return s
}

func (s *GetSignInListResponseBodyUsers) SetUserId(v string) *GetSignInListResponseBodyUsers {
	s.UserId = &v
	return s
}

type GetSignInListResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetSignInListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetSignInListResponse) String() string {
	return tea.Prettify(s)
}

func (s GetSignInListResponse) GoString() string {
	return s.String()
}

func (s *GetSignInListResponse) SetHeaders(v map[string]*string) *GetSignInListResponse {
	s.Headers = v
	return s
}

func (s *GetSignInListResponse) SetBody(v *GetSignInListResponseBody) *GetSignInListResponse {
	s.Body = v
	return s
}

type GetSignOutListHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetSignOutListHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetSignOutListHeaders) GoString() string {
	return s.String()
}

func (s *GetSignOutListHeaders) SetCommonHeaders(v map[string]*string) *GetSignOutListHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetSignOutListHeaders) SetXAcsDingtalkAccessToken(v string) *GetSignOutListHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetSignOutListRequest struct {
	// 查询返回结果数（上限200）
	MaxResults *int32  `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	NextToken  *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// 签到信息类型（sign_out，not_yet_sign_out)
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s GetSignOutListRequest) String() string {
	return tea.Prettify(s)
}

func (s GetSignOutListRequest) GoString() string {
	return s.String()
}

func (s *GetSignOutListRequest) SetMaxResults(v int32) *GetSignOutListRequest {
	s.MaxResults = &v
	return s
}

func (s *GetSignOutListRequest) SetNextToken(v string) *GetSignOutListRequest {
	s.NextToken = &v
	return s
}

func (s *GetSignOutListRequest) SetType(v string) *GetSignOutListRequest {
	s.Type = &v
	return s
}

type GetSignOutListResponseBody struct {
	// 翻页token
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// 签退信息
	Users []*GetSignOutListResponseBodyUsers `json:"users,omitempty" xml:"users,omitempty" type:"Repeated"`
}

func (s GetSignOutListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetSignOutListResponseBody) GoString() string {
	return s.String()
}

func (s *GetSignOutListResponseBody) SetNextToken(v string) *GetSignOutListResponseBody {
	s.NextToken = &v
	return s
}

func (s *GetSignOutListResponseBody) SetUsers(v []*GetSignOutListResponseBodyUsers) *GetSignOutListResponseBody {
	s.Users = v
	return s
}

type GetSignOutListResponseBodyUsers struct {
	// 签退时间
	CheckOutTime *int64 `json:"checkOutTime,omitempty" xml:"checkOutTime,omitempty"`
	// 用户名
	DisplayName *string `json:"displayName,omitempty" xml:"displayName,omitempty"`
	UserId      *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s GetSignOutListResponseBodyUsers) String() string {
	return tea.Prettify(s)
}

func (s GetSignOutListResponseBodyUsers) GoString() string {
	return s.String()
}

func (s *GetSignOutListResponseBodyUsers) SetCheckOutTime(v int64) *GetSignOutListResponseBodyUsers {
	s.CheckOutTime = &v
	return s
}

func (s *GetSignOutListResponseBodyUsers) SetDisplayName(v string) *GetSignOutListResponseBodyUsers {
	s.DisplayName = &v
	return s
}

func (s *GetSignOutListResponseBodyUsers) SetUserId(v string) *GetSignOutListResponseBodyUsers {
	s.UserId = &v
	return s
}

type GetSignOutListResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetSignOutListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetSignOutListResponse) String() string {
	return tea.Prettify(s)
}

func (s GetSignOutListResponse) GoString() string {
	return s.String()
}

func (s *GetSignOutListResponse) SetHeaders(v map[string]*string) *GetSignOutListResponse {
	s.Headers = v
	return s
}

func (s *GetSignOutListResponse) SetBody(v *GetSignOutListResponseBody) *GetSignOutListResponse {
	s.Body = v
	return s
}

type GetSubscribedCalendarHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetSubscribedCalendarHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetSubscribedCalendarHeaders) GoString() string {
	return s.String()
}

func (s *GetSubscribedCalendarHeaders) SetCommonHeaders(v map[string]*string) *GetSubscribedCalendarHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetSubscribedCalendarHeaders) SetXAcsDingtalkAccessToken(v string) *GetSubscribedCalendarHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetSubscribedCalendarResponseBody struct {
	// 日历作者
	Author *string `json:"author,omitempty" xml:"author,omitempty"`
	// 订阅日历id
	CalendarId *string `json:"calendarId,omitempty" xml:"calendarId,omitempty"`
	// 日历描述
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// 可管理人群
	Managers []*string `json:"managers,omitempty" xml:"managers,omitempty" type:"Repeated"`
	// 日历名
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 可订阅范围
	SubscribeScope *GetSubscribedCalendarResponseBodySubscribeScope `json:"subscribeScope,omitempty" xml:"subscribeScope,omitempty" type:"Struct"`
}

func (s GetSubscribedCalendarResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetSubscribedCalendarResponseBody) GoString() string {
	return s.String()
}

func (s *GetSubscribedCalendarResponseBody) SetAuthor(v string) *GetSubscribedCalendarResponseBody {
	s.Author = &v
	return s
}

func (s *GetSubscribedCalendarResponseBody) SetCalendarId(v string) *GetSubscribedCalendarResponseBody {
	s.CalendarId = &v
	return s
}

func (s *GetSubscribedCalendarResponseBody) SetDescription(v string) *GetSubscribedCalendarResponseBody {
	s.Description = &v
	return s
}

func (s *GetSubscribedCalendarResponseBody) SetManagers(v []*string) *GetSubscribedCalendarResponseBody {
	s.Managers = v
	return s
}

func (s *GetSubscribedCalendarResponseBody) SetName(v string) *GetSubscribedCalendarResponseBody {
	s.Name = &v
	return s
}

func (s *GetSubscribedCalendarResponseBody) SetSubscribeScope(v *GetSubscribedCalendarResponseBodySubscribeScope) *GetSubscribedCalendarResponseBody {
	s.SubscribeScope = v
	return s
}

type GetSubscribedCalendarResponseBodySubscribeScope struct {
	// 可订阅组织
	CorpIds []*string `json:"corpIds,omitempty" xml:"corpIds,omitempty" type:"Repeated"`
	// 可订阅群组
	OpenConversationIds []*string `json:"openConversationIds,omitempty" xml:"openConversationIds,omitempty" type:"Repeated"`
	// 可订阅用户
	UnionIds []*string `json:"unionIds,omitempty" xml:"unionIds,omitempty" type:"Repeated"`
}

func (s GetSubscribedCalendarResponseBodySubscribeScope) String() string {
	return tea.Prettify(s)
}

func (s GetSubscribedCalendarResponseBodySubscribeScope) GoString() string {
	return s.String()
}

func (s *GetSubscribedCalendarResponseBodySubscribeScope) SetCorpIds(v []*string) *GetSubscribedCalendarResponseBodySubscribeScope {
	s.CorpIds = v
	return s
}

func (s *GetSubscribedCalendarResponseBodySubscribeScope) SetOpenConversationIds(v []*string) *GetSubscribedCalendarResponseBodySubscribeScope {
	s.OpenConversationIds = v
	return s
}

func (s *GetSubscribedCalendarResponseBodySubscribeScope) SetUnionIds(v []*string) *GetSubscribedCalendarResponseBodySubscribeScope {
	s.UnionIds = v
	return s
}

type GetSubscribedCalendarResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetSubscribedCalendarResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetSubscribedCalendarResponse) String() string {
	return tea.Prettify(s)
}

func (s GetSubscribedCalendarResponse) GoString() string {
	return s.String()
}

func (s *GetSubscribedCalendarResponse) SetHeaders(v map[string]*string) *GetSubscribedCalendarResponse {
	s.Headers = v
	return s
}

func (s *GetSubscribedCalendarResponse) SetBody(v *GetSubscribedCalendarResponseBody) *GetSubscribedCalendarResponse {
	s.Body = v
	return s
}

type ListAclsHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s ListAclsHeaders) String() string {
	return tea.Prettify(s)
}

func (s ListAclsHeaders) GoString() string {
	return s.String()
}

func (s *ListAclsHeaders) SetCommonHeaders(v map[string]*string) *ListAclsHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListAclsHeaders) SetXAcsDingtalkAccessToken(v string) *ListAclsHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type ListAclsResponseBody struct {
	// 访问控制列表
	Acls []*ListAclsResponseBodyAcls `json:"acls,omitempty" xml:"acls,omitempty" type:"Repeated"`
}

func (s ListAclsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListAclsResponseBody) GoString() string {
	return s.String()
}

func (s *ListAclsResponseBody) SetAcls(v []*ListAclsResponseBodyAcls) *ListAclsResponseBody {
	s.Acls = v
	return s
}

type ListAclsResponseBodyAcls struct {
	// acl资源ID
	AclId *string `json:"aclId,omitempty" xml:"aclId,omitempty"`
	// 权限信息
	Privilege *string `json:"privilege,omitempty" xml:"privilege,omitempty"`
	// 权限范围
	Scope *ListAclsResponseBodyAclsScope `json:"scope,omitempty" xml:"scope,omitempty" type:"Struct"`
}

func (s ListAclsResponseBodyAcls) String() string {
	return tea.Prettify(s)
}

func (s ListAclsResponseBodyAcls) GoString() string {
	return s.String()
}

func (s *ListAclsResponseBodyAcls) SetAclId(v string) *ListAclsResponseBodyAcls {
	s.AclId = &v
	return s
}

func (s *ListAclsResponseBodyAcls) SetPrivilege(v string) *ListAclsResponseBodyAcls {
	s.Privilege = &v
	return s
}

func (s *ListAclsResponseBodyAcls) SetScope(v *ListAclsResponseBodyAclsScope) *ListAclsResponseBodyAcls {
	s.Scope = v
	return s
}

type ListAclsResponseBodyAclsScope struct {
	// 权限类型
	ScopeType *string `json:"scopeType,omitempty" xml:"scopeType,omitempty"`
	// 用户id
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s ListAclsResponseBodyAclsScope) String() string {
	return tea.Prettify(s)
}

func (s ListAclsResponseBodyAclsScope) GoString() string {
	return s.String()
}

func (s *ListAclsResponseBodyAclsScope) SetScopeType(v string) *ListAclsResponseBodyAclsScope {
	s.ScopeType = &v
	return s
}

func (s *ListAclsResponseBodyAclsScope) SetUserId(v string) *ListAclsResponseBodyAclsScope {
	s.UserId = &v
	return s
}

type ListAclsResponse struct {
	Headers map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListAclsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListAclsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListAclsResponse) GoString() string {
	return s.String()
}

func (s *ListAclsResponse) SetHeaders(v map[string]*string) *ListAclsResponse {
	s.Headers = v
	return s
}

func (s *ListAclsResponse) SetBody(v *ListAclsResponseBody) *ListAclsResponse {
	s.Body = v
	return s
}

type ListAttendeesHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s ListAttendeesHeaders) String() string {
	return tea.Prettify(s)
}

func (s ListAttendeesHeaders) GoString() string {
	return s.String()
}

func (s *ListAttendeesHeaders) SetCommonHeaders(v map[string]*string) *ListAttendeesHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListAttendeesHeaders) SetXAcsDingtalkAccessToken(v string) *ListAttendeesHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type ListAttendeesRequest struct {
	MaxResults *int32  `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	NextToken  *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
}

func (s ListAttendeesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListAttendeesRequest) GoString() string {
	return s.String()
}

func (s *ListAttendeesRequest) SetMaxResults(v int32) *ListAttendeesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListAttendeesRequest) SetNextToken(v string) *ListAttendeesRequest {
	s.NextToken = &v
	return s
}

type ListAttendeesResponseBody struct {
	// 参与人
	Attendees []*ListAttendeesResponseBodyAttendees `json:"attendees,omitempty" xml:"attendees,omitempty" type:"Repeated"`
	// 翻页token
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
}

func (s ListAttendeesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListAttendeesResponseBody) GoString() string {
	return s.String()
}

func (s *ListAttendeesResponseBody) SetAttendees(v []*ListAttendeesResponseBodyAttendees) *ListAttendeesResponseBody {
	s.Attendees = v
	return s
}

func (s *ListAttendeesResponseBody) SetNextToken(v string) *ListAttendeesResponseBody {
	s.NextToken = &v
	return s
}

type ListAttendeesResponseBodyAttendees struct {
	// 用户名
	DisplayName *string `json:"displayName,omitempty" xml:"displayName,omitempty"`
	// 用户id
	Id         *string `json:"id,omitempty" xml:"id,omitempty"`
	IsOptional *bool   `json:"isOptional,omitempty" xml:"isOptional,omitempty"`
	// 回复状态
	ResponseStatus *string `json:"responseStatus,omitempty" xml:"responseStatus,omitempty"`
	// 是否当前用户
	Self *bool `json:"self,omitempty" xml:"self,omitempty"`
}

func (s ListAttendeesResponseBodyAttendees) String() string {
	return tea.Prettify(s)
}

func (s ListAttendeesResponseBodyAttendees) GoString() string {
	return s.String()
}

func (s *ListAttendeesResponseBodyAttendees) SetDisplayName(v string) *ListAttendeesResponseBodyAttendees {
	s.DisplayName = &v
	return s
}

func (s *ListAttendeesResponseBodyAttendees) SetId(v string) *ListAttendeesResponseBodyAttendees {
	s.Id = &v
	return s
}

func (s *ListAttendeesResponseBodyAttendees) SetIsOptional(v bool) *ListAttendeesResponseBodyAttendees {
	s.IsOptional = &v
	return s
}

func (s *ListAttendeesResponseBodyAttendees) SetResponseStatus(v string) *ListAttendeesResponseBodyAttendees {
	s.ResponseStatus = &v
	return s
}

func (s *ListAttendeesResponseBodyAttendees) SetSelf(v bool) *ListAttendeesResponseBodyAttendees {
	s.Self = &v
	return s
}

type ListAttendeesResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListAttendeesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListAttendeesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListAttendeesResponse) GoString() string {
	return s.String()
}

func (s *ListAttendeesResponse) SetHeaders(v map[string]*string) *ListAttendeesResponse {
	s.Headers = v
	return s
}

func (s *ListAttendeesResponse) SetBody(v *ListAttendeesResponseBody) *ListAttendeesResponse {
	s.Body = v
	return s
}

type ListCalendarsHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s ListCalendarsHeaders) String() string {
	return tea.Prettify(s)
}

func (s ListCalendarsHeaders) GoString() string {
	return s.String()
}

func (s *ListCalendarsHeaders) SetCommonHeaders(v map[string]*string) *ListCalendarsHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListCalendarsHeaders) SetXAcsDingtalkAccessToken(v string) *ListCalendarsHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type ListCalendarsResponseBody struct {
	// 日历信息
	Response *ListCalendarsResponseBodyResponse `json:"response,omitempty" xml:"response,omitempty" type:"Struct"`
}

func (s ListCalendarsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListCalendarsResponseBody) GoString() string {
	return s.String()
}

func (s *ListCalendarsResponseBody) SetResponse(v *ListCalendarsResponseBodyResponse) *ListCalendarsResponseBody {
	s.Response = v
	return s
}

type ListCalendarsResponseBodyResponse struct {
	Calendars []*ListCalendarsResponseBodyResponseCalendars `json:"calendars,omitempty" xml:"calendars,omitempty" type:"Repeated"`
}

func (s ListCalendarsResponseBodyResponse) String() string {
	return tea.Prettify(s)
}

func (s ListCalendarsResponseBodyResponse) GoString() string {
	return s.String()
}

func (s *ListCalendarsResponseBodyResponse) SetCalendars(v []*ListCalendarsResponseBodyResponseCalendars) *ListCalendarsResponseBodyResponse {
	s.Calendars = v
	return s
}

type ListCalendarsResponseBodyResponseCalendars struct {
	// 日历描述
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// Calendar资源的ETag，用于检测该Calendar以及内部的Event是否有被更新
	ETag *string `json:"eTag,omitempty" xml:"eTag,omitempty"`
	// 日历id
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// 权限信息
	Privilege *string `json:"privilege,omitempty" xml:"privilege,omitempty"`
	// 日历标题
	Summary *string `json:"summary,omitempty" xml:"summary,omitempty"`
	// 时区
	TimeZone *string `json:"timeZone,omitempty" xml:"timeZone,omitempty"`
	// 日历类型
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s ListCalendarsResponseBodyResponseCalendars) String() string {
	return tea.Prettify(s)
}

func (s ListCalendarsResponseBodyResponseCalendars) GoString() string {
	return s.String()
}

func (s *ListCalendarsResponseBodyResponseCalendars) SetDescription(v string) *ListCalendarsResponseBodyResponseCalendars {
	s.Description = &v
	return s
}

func (s *ListCalendarsResponseBodyResponseCalendars) SetETag(v string) *ListCalendarsResponseBodyResponseCalendars {
	s.ETag = &v
	return s
}

func (s *ListCalendarsResponseBodyResponseCalendars) SetId(v string) *ListCalendarsResponseBodyResponseCalendars {
	s.Id = &v
	return s
}

func (s *ListCalendarsResponseBodyResponseCalendars) SetPrivilege(v string) *ListCalendarsResponseBodyResponseCalendars {
	s.Privilege = &v
	return s
}

func (s *ListCalendarsResponseBodyResponseCalendars) SetSummary(v string) *ListCalendarsResponseBodyResponseCalendars {
	s.Summary = &v
	return s
}

func (s *ListCalendarsResponseBodyResponseCalendars) SetTimeZone(v string) *ListCalendarsResponseBodyResponseCalendars {
	s.TimeZone = &v
	return s
}

func (s *ListCalendarsResponseBodyResponseCalendars) SetType(v string) *ListCalendarsResponseBodyResponseCalendars {
	s.Type = &v
	return s
}

type ListCalendarsResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListCalendarsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListCalendarsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListCalendarsResponse) GoString() string {
	return s.String()
}

func (s *ListCalendarsResponse) SetHeaders(v map[string]*string) *ListCalendarsResponse {
	s.Headers = v
	return s
}

func (s *ListCalendarsResponse) SetBody(v *ListCalendarsResponseBody) *ListCalendarsResponse {
	s.Body = v
	return s
}

type ListEventsHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s ListEventsHeaders) String() string {
	return tea.Prettify(s)
}

func (s ListEventsHeaders) GoString() string {
	return s.String()
}

func (s *ListEventsHeaders) SetCommonHeaders(v map[string]*string) *ListEventsHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListEventsHeaders) SetXAcsDingtalkAccessToken(v string) *ListEventsHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type ListEventsRequest struct {
	// 每个日程的参与者查询个数，默认100，最大100
	MaxAttendees *int32 `json:"maxAttendees,omitempty" xml:"maxAttendees,omitempty"`
	// 返回的最大日程数，最大100个，默认100个
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// 查询翻页token
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// 是否返回删除事件
	ShowDeleted *bool `json:"showDeleted,omitempty" xml:"showDeleted,omitempty"`
	// 增量查询token
	SyncToken *string `json:"syncToken,omitempty" xml:"syncToken,omitempty"`
	// 查询截止时间
	TimeMax *string `json:"timeMax,omitempty" xml:"timeMax,omitempty"`
	// 查询开始时间
	TimeMin *string `json:"timeMin,omitempty" xml:"timeMin,omitempty"`
}

func (s ListEventsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListEventsRequest) GoString() string {
	return s.String()
}

func (s *ListEventsRequest) SetMaxAttendees(v int32) *ListEventsRequest {
	s.MaxAttendees = &v
	return s
}

func (s *ListEventsRequest) SetMaxResults(v int32) *ListEventsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListEventsRequest) SetNextToken(v string) *ListEventsRequest {
	s.NextToken = &v
	return s
}

func (s *ListEventsRequest) SetShowDeleted(v bool) *ListEventsRequest {
	s.ShowDeleted = &v
	return s
}

func (s *ListEventsRequest) SetSyncToken(v string) *ListEventsRequest {
	s.SyncToken = &v
	return s
}

func (s *ListEventsRequest) SetTimeMax(v string) *ListEventsRequest {
	s.TimeMax = &v
	return s
}

func (s *ListEventsRequest) SetTimeMin(v string) *ListEventsRequest {
	s.TimeMin = &v
	return s
}

type ListEventsResponseBody struct {
	// 日程
	Events []*ListEventsResponseBodyEvents `json:"events,omitempty" xml:"events,omitempty" type:"Repeated"`
	// 翻页token
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// 增量同步token
	SyncToken *string `json:"syncToken,omitempty" xml:"syncToken,omitempty"`
}

func (s ListEventsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListEventsResponseBody) GoString() string {
	return s.String()
}

func (s *ListEventsResponseBody) SetEvents(v []*ListEventsResponseBodyEvents) *ListEventsResponseBody {
	s.Events = v
	return s
}

func (s *ListEventsResponseBody) SetNextToken(v string) *ListEventsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListEventsResponseBody) SetSyncToken(v string) *ListEventsResponseBody {
	s.SyncToken = &v
	return s
}

type ListEventsResponseBodyEvents struct {
	// 日程参与人
	Attendees []*ListEventsResponseBodyEventsAttendees `json:"attendees,omitempty" xml:"attendees,omitempty" type:"Repeated"`
	// 创建时间
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// 日程描述
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// 日程结束时间
	End                *ListEventsResponseBodyEventsEnd                `json:"end,omitempty" xml:"end,omitempty" type:"Struct"`
	ExtendedProperties *ListEventsResponseBodyEventsExtendedProperties `json:"extendedProperties,omitempty" xml:"extendedProperties,omitempty" type:"Struct"`
	// 日程事件id
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// 是否为全天日程
	IsAllDay *bool `json:"isAllDay,omitempty" xml:"isAllDay,omitempty"`
	// 日程地点
	Location          *ListEventsResponseBodyEventsLocation          `json:"location,omitempty" xml:"location,omitempty" type:"Struct"`
	OnlineMeetingInfo *ListEventsResponseBodyEventsOnlineMeetingInfo `json:"onlineMeetingInfo,omitempty" xml:"onlineMeetingInfo,omitempty" type:"Struct"`
	// 日程组织人
	Organizer *ListEventsResponseBodyEventsOrganizer `json:"organizer,omitempty" xml:"organizer,omitempty" type:"Struct"`
	// 日程重复规则
	Recurrence *ListEventsResponseBodyEventsRecurrence  `json:"recurrence,omitempty" xml:"recurrence,omitempty" type:"Struct"`
	Reminders  []*ListEventsResponseBodyEventsReminders `json:"reminders,omitempty" xml:"reminders,omitempty" type:"Repeated"`
	// 重复日程的主日程id，非重复日程为空
	SeriesMasterId *string `json:"seriesMasterId,omitempty" xml:"seriesMasterId,omitempty"`
	// 日程开始时间
	Start *ListEventsResponseBodyEventsStart `json:"start,omitempty" xml:"start,omitempty" type:"Struct"`
	// 日程状态
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// 日程标题
	Summary *string `json:"summary,omitempty" xml:"summary,omitempty"`
	// 更新时间
	UpdateTime *string `json:"updateTime,omitempty" xml:"updateTime,omitempty"`
}

func (s ListEventsResponseBodyEvents) String() string {
	return tea.Prettify(s)
}

func (s ListEventsResponseBodyEvents) GoString() string {
	return s.String()
}

func (s *ListEventsResponseBodyEvents) SetAttendees(v []*ListEventsResponseBodyEventsAttendees) *ListEventsResponseBodyEvents {
	s.Attendees = v
	return s
}

func (s *ListEventsResponseBodyEvents) SetCreateTime(v string) *ListEventsResponseBodyEvents {
	s.CreateTime = &v
	return s
}

func (s *ListEventsResponseBodyEvents) SetDescription(v string) *ListEventsResponseBodyEvents {
	s.Description = &v
	return s
}

func (s *ListEventsResponseBodyEvents) SetEnd(v *ListEventsResponseBodyEventsEnd) *ListEventsResponseBodyEvents {
	s.End = v
	return s
}

func (s *ListEventsResponseBodyEvents) SetExtendedProperties(v *ListEventsResponseBodyEventsExtendedProperties) *ListEventsResponseBodyEvents {
	s.ExtendedProperties = v
	return s
}

func (s *ListEventsResponseBodyEvents) SetId(v string) *ListEventsResponseBodyEvents {
	s.Id = &v
	return s
}

func (s *ListEventsResponseBodyEvents) SetIsAllDay(v bool) *ListEventsResponseBodyEvents {
	s.IsAllDay = &v
	return s
}

func (s *ListEventsResponseBodyEvents) SetLocation(v *ListEventsResponseBodyEventsLocation) *ListEventsResponseBodyEvents {
	s.Location = v
	return s
}

func (s *ListEventsResponseBodyEvents) SetOnlineMeetingInfo(v *ListEventsResponseBodyEventsOnlineMeetingInfo) *ListEventsResponseBodyEvents {
	s.OnlineMeetingInfo = v
	return s
}

func (s *ListEventsResponseBodyEvents) SetOrganizer(v *ListEventsResponseBodyEventsOrganizer) *ListEventsResponseBodyEvents {
	s.Organizer = v
	return s
}

func (s *ListEventsResponseBodyEvents) SetRecurrence(v *ListEventsResponseBodyEventsRecurrence) *ListEventsResponseBodyEvents {
	s.Recurrence = v
	return s
}

func (s *ListEventsResponseBodyEvents) SetReminders(v []*ListEventsResponseBodyEventsReminders) *ListEventsResponseBodyEvents {
	s.Reminders = v
	return s
}

func (s *ListEventsResponseBodyEvents) SetSeriesMasterId(v string) *ListEventsResponseBodyEvents {
	s.SeriesMasterId = &v
	return s
}

func (s *ListEventsResponseBodyEvents) SetStart(v *ListEventsResponseBodyEventsStart) *ListEventsResponseBodyEvents {
	s.Start = v
	return s
}

func (s *ListEventsResponseBodyEvents) SetStatus(v string) *ListEventsResponseBodyEvents {
	s.Status = &v
	return s
}

func (s *ListEventsResponseBodyEvents) SetSummary(v string) *ListEventsResponseBodyEvents {
	s.Summary = &v
	return s
}

func (s *ListEventsResponseBodyEvents) SetUpdateTime(v string) *ListEventsResponseBodyEvents {
	s.UpdateTime = &v
	return s
}

type ListEventsResponseBodyEventsAttendees struct {
	// 用户名
	DisplayName *string `json:"displayName,omitempty" xml:"displayName,omitempty"`
	// 用户id
	Id         *string `json:"id,omitempty" xml:"id,omitempty"`
	IsOptional *bool   `json:"isOptional,omitempty" xml:"isOptional,omitempty"`
	// 回复状态
	ResponseStatus *string `json:"responseStatus,omitempty" xml:"responseStatus,omitempty"`
	// 是否是当前登陆用户
	Self *bool `json:"self,omitempty" xml:"self,omitempty"`
}

func (s ListEventsResponseBodyEventsAttendees) String() string {
	return tea.Prettify(s)
}

func (s ListEventsResponseBodyEventsAttendees) GoString() string {
	return s.String()
}

func (s *ListEventsResponseBodyEventsAttendees) SetDisplayName(v string) *ListEventsResponseBodyEventsAttendees {
	s.DisplayName = &v
	return s
}

func (s *ListEventsResponseBodyEventsAttendees) SetId(v string) *ListEventsResponseBodyEventsAttendees {
	s.Id = &v
	return s
}

func (s *ListEventsResponseBodyEventsAttendees) SetIsOptional(v bool) *ListEventsResponseBodyEventsAttendees {
	s.IsOptional = &v
	return s
}

func (s *ListEventsResponseBodyEventsAttendees) SetResponseStatus(v string) *ListEventsResponseBodyEventsAttendees {
	s.ResponseStatus = &v
	return s
}

func (s *ListEventsResponseBodyEventsAttendees) SetSelf(v bool) *ListEventsResponseBodyEventsAttendees {
	s.Self = &v
	return s
}

type ListEventsResponseBodyEventsEnd struct {
	Date     *string `json:"date,omitempty" xml:"date,omitempty"`
	DateTime *string `json:"dateTime,omitempty" xml:"dateTime,omitempty"`
	TimeZone *string `json:"timeZone,omitempty" xml:"timeZone,omitempty"`
}

func (s ListEventsResponseBodyEventsEnd) String() string {
	return tea.Prettify(s)
}

func (s ListEventsResponseBodyEventsEnd) GoString() string {
	return s.String()
}

func (s *ListEventsResponseBodyEventsEnd) SetDate(v string) *ListEventsResponseBodyEventsEnd {
	s.Date = &v
	return s
}

func (s *ListEventsResponseBodyEventsEnd) SetDateTime(v string) *ListEventsResponseBodyEventsEnd {
	s.DateTime = &v
	return s
}

func (s *ListEventsResponseBodyEventsEnd) SetTimeZone(v string) *ListEventsResponseBodyEventsEnd {
	s.TimeZone = &v
	return s
}

type ListEventsResponseBodyEventsExtendedProperties struct {
	SharedProperties *ListEventsResponseBodyEventsExtendedPropertiesSharedProperties `json:"sharedProperties,omitempty" xml:"sharedProperties,omitempty" type:"Struct"`
}

func (s ListEventsResponseBodyEventsExtendedProperties) String() string {
	return tea.Prettify(s)
}

func (s ListEventsResponseBodyEventsExtendedProperties) GoString() string {
	return s.String()
}

func (s *ListEventsResponseBodyEventsExtendedProperties) SetSharedProperties(v *ListEventsResponseBodyEventsExtendedPropertiesSharedProperties) *ListEventsResponseBodyEventsExtendedProperties {
	s.SharedProperties = v
	return s
}

type ListEventsResponseBodyEventsExtendedPropertiesSharedProperties struct {
	SourceOpenCid *string `json:"sourceOpenCid,omitempty" xml:"sourceOpenCid,omitempty"`
}

func (s ListEventsResponseBodyEventsExtendedPropertiesSharedProperties) String() string {
	return tea.Prettify(s)
}

func (s ListEventsResponseBodyEventsExtendedPropertiesSharedProperties) GoString() string {
	return s.String()
}

func (s *ListEventsResponseBodyEventsExtendedPropertiesSharedProperties) SetSourceOpenCid(v string) *ListEventsResponseBodyEventsExtendedPropertiesSharedProperties {
	s.SourceOpenCid = &v
	return s
}

type ListEventsResponseBodyEventsLocation struct {
	// 展示名称
	DisplayName  *string   `json:"displayName,omitempty" xml:"displayName,omitempty"`
	MeetingRooms []*string `json:"meetingRooms,omitempty" xml:"meetingRooms,omitempty" type:"Repeated"`
}

func (s ListEventsResponseBodyEventsLocation) String() string {
	return tea.Prettify(s)
}

func (s ListEventsResponseBodyEventsLocation) GoString() string {
	return s.String()
}

func (s *ListEventsResponseBodyEventsLocation) SetDisplayName(v string) *ListEventsResponseBodyEventsLocation {
	s.DisplayName = &v
	return s
}

func (s *ListEventsResponseBodyEventsLocation) SetMeetingRooms(v []*string) *ListEventsResponseBodyEventsLocation {
	s.MeetingRooms = v
	return s
}

type ListEventsResponseBodyEventsOnlineMeetingInfo struct {
	ConferenceId *string                `json:"conferenceId,omitempty" xml:"conferenceId,omitempty"`
	ExtraInfo    map[string]interface{} `json:"extraInfo,omitempty" xml:"extraInfo,omitempty"`
	Type         *string                `json:"type,omitempty" xml:"type,omitempty"`
	Url          *string                `json:"url,omitempty" xml:"url,omitempty"`
}

func (s ListEventsResponseBodyEventsOnlineMeetingInfo) String() string {
	return tea.Prettify(s)
}

func (s ListEventsResponseBodyEventsOnlineMeetingInfo) GoString() string {
	return s.String()
}

func (s *ListEventsResponseBodyEventsOnlineMeetingInfo) SetConferenceId(v string) *ListEventsResponseBodyEventsOnlineMeetingInfo {
	s.ConferenceId = &v
	return s
}

func (s *ListEventsResponseBodyEventsOnlineMeetingInfo) SetExtraInfo(v map[string]interface{}) *ListEventsResponseBodyEventsOnlineMeetingInfo {
	s.ExtraInfo = v
	return s
}

func (s *ListEventsResponseBodyEventsOnlineMeetingInfo) SetType(v string) *ListEventsResponseBodyEventsOnlineMeetingInfo {
	s.Type = &v
	return s
}

func (s *ListEventsResponseBodyEventsOnlineMeetingInfo) SetUrl(v string) *ListEventsResponseBodyEventsOnlineMeetingInfo {
	s.Url = &v
	return s
}

type ListEventsResponseBodyEventsOrganizer struct {
	// 用户名
	DisplayName *string `json:"displayName,omitempty" xml:"displayName,omitempty"`
	// 用户id
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// 回复状态
	ResponseStatus *string `json:"responseStatus,omitempty" xml:"responseStatus,omitempty"`
	// 是否是当前登陆用户
	Self *bool `json:"self,omitempty" xml:"self,omitempty"`
}

func (s ListEventsResponseBodyEventsOrganizer) String() string {
	return tea.Prettify(s)
}

func (s ListEventsResponseBodyEventsOrganizer) GoString() string {
	return s.String()
}

func (s *ListEventsResponseBodyEventsOrganizer) SetDisplayName(v string) *ListEventsResponseBodyEventsOrganizer {
	s.DisplayName = &v
	return s
}

func (s *ListEventsResponseBodyEventsOrganizer) SetId(v string) *ListEventsResponseBodyEventsOrganizer {
	s.Id = &v
	return s
}

func (s *ListEventsResponseBodyEventsOrganizer) SetResponseStatus(v string) *ListEventsResponseBodyEventsOrganizer {
	s.ResponseStatus = &v
	return s
}

func (s *ListEventsResponseBodyEventsOrganizer) SetSelf(v bool) *ListEventsResponseBodyEventsOrganizer {
	s.Self = &v
	return s
}

type ListEventsResponseBodyEventsRecurrence struct {
	// 重复模式
	Pattern *ListEventsResponseBodyEventsRecurrencePattern `json:"pattern,omitempty" xml:"pattern,omitempty" type:"Struct"`
	// 重复范围
	Range *ListEventsResponseBodyEventsRecurrenceRange `json:"range,omitempty" xml:"range,omitempty" type:"Struct"`
}

func (s ListEventsResponseBodyEventsRecurrence) String() string {
	return tea.Prettify(s)
}

func (s ListEventsResponseBodyEventsRecurrence) GoString() string {
	return s.String()
}

func (s *ListEventsResponseBodyEventsRecurrence) SetPattern(v *ListEventsResponseBodyEventsRecurrencePattern) *ListEventsResponseBodyEventsRecurrence {
	s.Pattern = v
	return s
}

func (s *ListEventsResponseBodyEventsRecurrence) SetRange(v *ListEventsResponseBodyEventsRecurrenceRange) *ListEventsResponseBodyEventsRecurrence {
	s.Range = v
	return s
}

type ListEventsResponseBodyEventsRecurrencePattern struct {
	DayOfMonth *int32  `json:"dayOfMonth,omitempty" xml:"dayOfMonth,omitempty"`
	DaysOfWeek *string `json:"daysOfWeek,omitempty" xml:"daysOfWeek,omitempty"`
	Index      *string `json:"index,omitempty" xml:"index,omitempty"`
	Interval   *int32  `json:"interval,omitempty" xml:"interval,omitempty"`
	// 循环模式类型(type: daily, weekly, absoluteMonthly, relativeMonthly, absoluteYearly, relativeYearly)
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s ListEventsResponseBodyEventsRecurrencePattern) String() string {
	return tea.Prettify(s)
}

func (s ListEventsResponseBodyEventsRecurrencePattern) GoString() string {
	return s.String()
}

func (s *ListEventsResponseBodyEventsRecurrencePattern) SetDayOfMonth(v int32) *ListEventsResponseBodyEventsRecurrencePattern {
	s.DayOfMonth = &v
	return s
}

func (s *ListEventsResponseBodyEventsRecurrencePattern) SetDaysOfWeek(v string) *ListEventsResponseBodyEventsRecurrencePattern {
	s.DaysOfWeek = &v
	return s
}

func (s *ListEventsResponseBodyEventsRecurrencePattern) SetIndex(v string) *ListEventsResponseBodyEventsRecurrencePattern {
	s.Index = &v
	return s
}

func (s *ListEventsResponseBodyEventsRecurrencePattern) SetInterval(v int32) *ListEventsResponseBodyEventsRecurrencePattern {
	s.Interval = &v
	return s
}

func (s *ListEventsResponseBodyEventsRecurrencePattern) SetType(v string) *ListEventsResponseBodyEventsRecurrencePattern {
	s.Type = &v
	return s
}

type ListEventsResponseBodyEventsRecurrenceRange struct {
	EndDate             *string `json:"endDate,omitempty" xml:"endDate,omitempty"`
	NumberOfOccurrences *int32  `json:"numberOfOccurrences,omitempty" xml:"numberOfOccurrences,omitempty"`
	// 范围类型(endDate, noEnd, numbered)
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s ListEventsResponseBodyEventsRecurrenceRange) String() string {
	return tea.Prettify(s)
}

func (s ListEventsResponseBodyEventsRecurrenceRange) GoString() string {
	return s.String()
}

func (s *ListEventsResponseBodyEventsRecurrenceRange) SetEndDate(v string) *ListEventsResponseBodyEventsRecurrenceRange {
	s.EndDate = &v
	return s
}

func (s *ListEventsResponseBodyEventsRecurrenceRange) SetNumberOfOccurrences(v int32) *ListEventsResponseBodyEventsRecurrenceRange {
	s.NumberOfOccurrences = &v
	return s
}

func (s *ListEventsResponseBodyEventsRecurrenceRange) SetType(v string) *ListEventsResponseBodyEventsRecurrenceRange {
	s.Type = &v
	return s
}

type ListEventsResponseBodyEventsReminders struct {
	Method  *string `json:"method,omitempty" xml:"method,omitempty"`
	Minutes *string `json:"minutes,omitempty" xml:"minutes,omitempty"`
}

func (s ListEventsResponseBodyEventsReminders) String() string {
	return tea.Prettify(s)
}

func (s ListEventsResponseBodyEventsReminders) GoString() string {
	return s.String()
}

func (s *ListEventsResponseBodyEventsReminders) SetMethod(v string) *ListEventsResponseBodyEventsReminders {
	s.Method = &v
	return s
}

func (s *ListEventsResponseBodyEventsReminders) SetMinutes(v string) *ListEventsResponseBodyEventsReminders {
	s.Minutes = &v
	return s
}

type ListEventsResponseBodyEventsStart struct {
	// 日期，格式：yyyyMMdd
	Date *string `json:"date,omitempty" xml:"date,omitempty"`
	// 时间戳，按照ISO 8601格式
	DateTime *string `json:"dateTime,omitempty" xml:"dateTime,omitempty"`
	// 时区
	TimeZone *string `json:"timeZone,omitempty" xml:"timeZone,omitempty"`
}

func (s ListEventsResponseBodyEventsStart) String() string {
	return tea.Prettify(s)
}

func (s ListEventsResponseBodyEventsStart) GoString() string {
	return s.String()
}

func (s *ListEventsResponseBodyEventsStart) SetDate(v string) *ListEventsResponseBodyEventsStart {
	s.Date = &v
	return s
}

func (s *ListEventsResponseBodyEventsStart) SetDateTime(v string) *ListEventsResponseBodyEventsStart {
	s.DateTime = &v
	return s
}

func (s *ListEventsResponseBodyEventsStart) SetTimeZone(v string) *ListEventsResponseBodyEventsStart {
	s.TimeZone = &v
	return s
}

type ListEventsResponse struct {
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListEventsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListEventsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListEventsResponse) GoString() string {
	return s.String()
}

func (s *ListEventsResponse) SetHeaders(v map[string]*string) *ListEventsResponse {
	s.Headers = v
	return s
}

func (s *ListEventsResponse) SetBody(v *ListEventsResponseBody) *ListEventsResponse {
	s.Body = v
	return s
}

type ListEventsInstancesHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s ListEventsInstancesHeaders) String() string {
	return tea.Prettify(s)
}

func (s ListEventsInstancesHeaders) GoString() string {
	return s.String()
}

func (s *ListEventsInstancesHeaders) SetCommonHeaders(v map[string]*string) *ListEventsInstancesHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListEventsInstancesHeaders) SetXAcsDingtalkAccessToken(v string) *ListEventsInstancesHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type ListEventsInstancesRequest struct {
	// listInstances每个日程的参与者查询个数，默认100，最大100。
	MaxAttendees *int32 `json:"maxAttendees,omitempty" xml:"maxAttendees,omitempty"`
	// listInstances返回的最大日程数，最大100个，默认100个。
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// 循环主日程id。
	SeriesMasterId *string `json:"seriesMasterId,omitempty" xml:"seriesMasterId,omitempty"`
	// 大于等于次序列id的所有实例
	StartRecurrenceId *string `json:"startRecurrenceId,omitempty" xml:"startRecurrenceId,omitempty"`
}

func (s ListEventsInstancesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListEventsInstancesRequest) GoString() string {
	return s.String()
}

func (s *ListEventsInstancesRequest) SetMaxAttendees(v int32) *ListEventsInstancesRequest {
	s.MaxAttendees = &v
	return s
}

func (s *ListEventsInstancesRequest) SetMaxResults(v int32) *ListEventsInstancesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListEventsInstancesRequest) SetSeriesMasterId(v string) *ListEventsInstancesRequest {
	s.SeriesMasterId = &v
	return s
}

func (s *ListEventsInstancesRequest) SetStartRecurrenceId(v string) *ListEventsInstancesRequest {
	s.StartRecurrenceId = &v
	return s
}

type ListEventsInstancesResponseBody struct {
	// 日程
	Events []*ListEventsInstancesResponseBodyEvents `json:"events,omitempty" xml:"events,omitempty" type:"Repeated"`
}

func (s ListEventsInstancesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListEventsInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *ListEventsInstancesResponseBody) SetEvents(v []*ListEventsInstancesResponseBodyEvents) *ListEventsInstancesResponseBody {
	s.Events = v
	return s
}

type ListEventsInstancesResponseBodyEvents struct {
	// 日程参与人
	Attendees []*ListEventsInstancesResponseBodyEventsAttendees `json:"attendees,omitempty" xml:"attendees,omitempty" type:"Repeated"`
	// 创建时间
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// 日程描述
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// 日程结束时间
	End                *ListEventsInstancesResponseBodyEventsEnd                `json:"end,omitempty" xml:"end,omitempty" type:"Struct"`
	ExtendedProperties *ListEventsInstancesResponseBodyEventsExtendedProperties `json:"extendedProperties,omitempty" xml:"extendedProperties,omitempty" type:"Struct"`
	// 日程事件id
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// 是否为全天日程
	IsAllDay *bool `json:"isAllDay,omitempty" xml:"isAllDay,omitempty"`
	// 日程地点
	Location *ListEventsInstancesResponseBodyEventsLocation `json:"location,omitempty" xml:"location,omitempty" type:"Struct"`
	// 线上会议
	OnlineMeetingInfo *ListEventsInstancesResponseBodyEventsOnlineMeetingInfo `json:"onlineMeetingInfo,omitempty" xml:"onlineMeetingInfo,omitempty" type:"Struct"`
	// 日程组织人
	Organizer *ListEventsInstancesResponseBodyEventsOrganizer `json:"organizer,omitempty" xml:"organizer,omitempty" type:"Struct"`
	// 日程重复规则
	Recurrence *ListEventsInstancesResponseBodyEventsRecurrence `json:"recurrence,omitempty" xml:"recurrence,omitempty" type:"Struct"`
	// 日程提醒
	Reminders []*ListEventsInstancesResponseBodyEventsReminders `json:"reminders,omitempty" xml:"reminders,omitempty" type:"Repeated"`
	// 重复日程的主日程id，非重复日程为空
	SeriesMasterId *string `json:"seriesMasterId,omitempty" xml:"seriesMasterId,omitempty"`
	// 日程开始时间
	Start *ListEventsInstancesResponseBodyEventsStart `json:"start,omitempty" xml:"start,omitempty" type:"Struct"`
	// 日程状态
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// 日程标题
	Summary *string `json:"summary,omitempty" xml:"summary,omitempty"`
	// 更新时间
	UpdateTime *string `json:"updateTime,omitempty" xml:"updateTime,omitempty"`
}

func (s ListEventsInstancesResponseBodyEvents) String() string {
	return tea.Prettify(s)
}

func (s ListEventsInstancesResponseBodyEvents) GoString() string {
	return s.String()
}

func (s *ListEventsInstancesResponseBodyEvents) SetAttendees(v []*ListEventsInstancesResponseBodyEventsAttendees) *ListEventsInstancesResponseBodyEvents {
	s.Attendees = v
	return s
}

func (s *ListEventsInstancesResponseBodyEvents) SetCreateTime(v string) *ListEventsInstancesResponseBodyEvents {
	s.CreateTime = &v
	return s
}

func (s *ListEventsInstancesResponseBodyEvents) SetDescription(v string) *ListEventsInstancesResponseBodyEvents {
	s.Description = &v
	return s
}

func (s *ListEventsInstancesResponseBodyEvents) SetEnd(v *ListEventsInstancesResponseBodyEventsEnd) *ListEventsInstancesResponseBodyEvents {
	s.End = v
	return s
}

func (s *ListEventsInstancesResponseBodyEvents) SetExtendedProperties(v *ListEventsInstancesResponseBodyEventsExtendedProperties) *ListEventsInstancesResponseBodyEvents {
	s.ExtendedProperties = v
	return s
}

func (s *ListEventsInstancesResponseBodyEvents) SetId(v string) *ListEventsInstancesResponseBodyEvents {
	s.Id = &v
	return s
}

func (s *ListEventsInstancesResponseBodyEvents) SetIsAllDay(v bool) *ListEventsInstancesResponseBodyEvents {
	s.IsAllDay = &v
	return s
}

func (s *ListEventsInstancesResponseBodyEvents) SetLocation(v *ListEventsInstancesResponseBodyEventsLocation) *ListEventsInstancesResponseBodyEvents {
	s.Location = v
	return s
}

func (s *ListEventsInstancesResponseBodyEvents) SetOnlineMeetingInfo(v *ListEventsInstancesResponseBodyEventsOnlineMeetingInfo) *ListEventsInstancesResponseBodyEvents {
	s.OnlineMeetingInfo = v
	return s
}

func (s *ListEventsInstancesResponseBodyEvents) SetOrganizer(v *ListEventsInstancesResponseBodyEventsOrganizer) *ListEventsInstancesResponseBodyEvents {
	s.Organizer = v
	return s
}

func (s *ListEventsInstancesResponseBodyEvents) SetRecurrence(v *ListEventsInstancesResponseBodyEventsRecurrence) *ListEventsInstancesResponseBodyEvents {
	s.Recurrence = v
	return s
}

func (s *ListEventsInstancesResponseBodyEvents) SetReminders(v []*ListEventsInstancesResponseBodyEventsReminders) *ListEventsInstancesResponseBodyEvents {
	s.Reminders = v
	return s
}

func (s *ListEventsInstancesResponseBodyEvents) SetSeriesMasterId(v string) *ListEventsInstancesResponseBodyEvents {
	s.SeriesMasterId = &v
	return s
}

func (s *ListEventsInstancesResponseBodyEvents) SetStart(v *ListEventsInstancesResponseBodyEventsStart) *ListEventsInstancesResponseBodyEvents {
	s.Start = v
	return s
}

func (s *ListEventsInstancesResponseBodyEvents) SetStatus(v string) *ListEventsInstancesResponseBodyEvents {
	s.Status = &v
	return s
}

func (s *ListEventsInstancesResponseBodyEvents) SetSummary(v string) *ListEventsInstancesResponseBodyEvents {
	s.Summary = &v
	return s
}

func (s *ListEventsInstancesResponseBodyEvents) SetUpdateTime(v string) *ListEventsInstancesResponseBodyEvents {
	s.UpdateTime = &v
	return s
}

type ListEventsInstancesResponseBodyEventsAttendees struct {
	// 用户名
	DisplayName *string `json:"displayName,omitempty" xml:"displayName,omitempty"`
	// 用户id
	Id         *string `json:"id,omitempty" xml:"id,omitempty"`
	IsOptional *bool   `json:"isOptional,omitempty" xml:"isOptional,omitempty"`
	// 回复状态
	ResponseStatus *string `json:"responseStatus,omitempty" xml:"responseStatus,omitempty"`
	// 是否是当前登陆用户
	Self *bool `json:"self,omitempty" xml:"self,omitempty"`
}

func (s ListEventsInstancesResponseBodyEventsAttendees) String() string {
	return tea.Prettify(s)
}

func (s ListEventsInstancesResponseBodyEventsAttendees) GoString() string {
	return s.String()
}

func (s *ListEventsInstancesResponseBodyEventsAttendees) SetDisplayName(v string) *ListEventsInstancesResponseBodyEventsAttendees {
	s.DisplayName = &v
	return s
}

func (s *ListEventsInstancesResponseBodyEventsAttendees) SetId(v string) *ListEventsInstancesResponseBodyEventsAttendees {
	s.Id = &v
	return s
}

func (s *ListEventsInstancesResponseBodyEventsAttendees) SetIsOptional(v bool) *ListEventsInstancesResponseBodyEventsAttendees {
	s.IsOptional = &v
	return s
}

func (s *ListEventsInstancesResponseBodyEventsAttendees) SetResponseStatus(v string) *ListEventsInstancesResponseBodyEventsAttendees {
	s.ResponseStatus = &v
	return s
}

func (s *ListEventsInstancesResponseBodyEventsAttendees) SetSelf(v bool) *ListEventsInstancesResponseBodyEventsAttendees {
	s.Self = &v
	return s
}

type ListEventsInstancesResponseBodyEventsEnd struct {
	// 日期，格式：yyyyMMdd
	Date *string `json:"date,omitempty" xml:"date,omitempty"`
	// 时间戳，按照ISO 8601格式
	DateTime *string `json:"dateTime,omitempty" xml:"dateTime,omitempty"`
	// 时区
	TimeZone *string `json:"timeZone,omitempty" xml:"timeZone,omitempty"`
}

func (s ListEventsInstancesResponseBodyEventsEnd) String() string {
	return tea.Prettify(s)
}

func (s ListEventsInstancesResponseBodyEventsEnd) GoString() string {
	return s.String()
}

func (s *ListEventsInstancesResponseBodyEventsEnd) SetDate(v string) *ListEventsInstancesResponseBodyEventsEnd {
	s.Date = &v
	return s
}

func (s *ListEventsInstancesResponseBodyEventsEnd) SetDateTime(v string) *ListEventsInstancesResponseBodyEventsEnd {
	s.DateTime = &v
	return s
}

func (s *ListEventsInstancesResponseBodyEventsEnd) SetTimeZone(v string) *ListEventsInstancesResponseBodyEventsEnd {
	s.TimeZone = &v
	return s
}

type ListEventsInstancesResponseBodyEventsExtendedProperties struct {
	SharedProperties *ListEventsInstancesResponseBodyEventsExtendedPropertiesSharedProperties `json:"sharedProperties,omitempty" xml:"sharedProperties,omitempty" type:"Struct"`
}

func (s ListEventsInstancesResponseBodyEventsExtendedProperties) String() string {
	return tea.Prettify(s)
}

func (s ListEventsInstancesResponseBodyEventsExtendedProperties) GoString() string {
	return s.String()
}

func (s *ListEventsInstancesResponseBodyEventsExtendedProperties) SetSharedProperties(v *ListEventsInstancesResponseBodyEventsExtendedPropertiesSharedProperties) *ListEventsInstancesResponseBodyEventsExtendedProperties {
	s.SharedProperties = v
	return s
}

type ListEventsInstancesResponseBodyEventsExtendedPropertiesSharedProperties struct {
	SourceOpenCid *string `json:"sourceOpenCid,omitempty" xml:"sourceOpenCid,omitempty"`
}

func (s ListEventsInstancesResponseBodyEventsExtendedPropertiesSharedProperties) String() string {
	return tea.Prettify(s)
}

func (s ListEventsInstancesResponseBodyEventsExtendedPropertiesSharedProperties) GoString() string {
	return s.String()
}

func (s *ListEventsInstancesResponseBodyEventsExtendedPropertiesSharedProperties) SetSourceOpenCid(v string) *ListEventsInstancesResponseBodyEventsExtendedPropertiesSharedProperties {
	s.SourceOpenCid = &v
	return s
}

type ListEventsInstancesResponseBodyEventsLocation struct {
	// 展示名称
	DisplayName  *string   `json:"displayName,omitempty" xml:"displayName,omitempty"`
	MeetingRooms []*string `json:"meetingRooms,omitempty" xml:"meetingRooms,omitempty" type:"Repeated"`
}

func (s ListEventsInstancesResponseBodyEventsLocation) String() string {
	return tea.Prettify(s)
}

func (s ListEventsInstancesResponseBodyEventsLocation) GoString() string {
	return s.String()
}

func (s *ListEventsInstancesResponseBodyEventsLocation) SetDisplayName(v string) *ListEventsInstancesResponseBodyEventsLocation {
	s.DisplayName = &v
	return s
}

func (s *ListEventsInstancesResponseBodyEventsLocation) SetMeetingRooms(v []*string) *ListEventsInstancesResponseBodyEventsLocation {
	s.MeetingRooms = v
	return s
}

type ListEventsInstancesResponseBodyEventsOnlineMeetingInfo struct {
	// 会议ID
	ConferenceId *string `json:"conferenceId,omitempty" xml:"conferenceId,omitempty"`
	// 线上会议类型，目前支持：  dingtalk：钉钉视频会议
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// 会议url
	Url *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s ListEventsInstancesResponseBodyEventsOnlineMeetingInfo) String() string {
	return tea.Prettify(s)
}

func (s ListEventsInstancesResponseBodyEventsOnlineMeetingInfo) GoString() string {
	return s.String()
}

func (s *ListEventsInstancesResponseBodyEventsOnlineMeetingInfo) SetConferenceId(v string) *ListEventsInstancesResponseBodyEventsOnlineMeetingInfo {
	s.ConferenceId = &v
	return s
}

func (s *ListEventsInstancesResponseBodyEventsOnlineMeetingInfo) SetType(v string) *ListEventsInstancesResponseBodyEventsOnlineMeetingInfo {
	s.Type = &v
	return s
}

func (s *ListEventsInstancesResponseBodyEventsOnlineMeetingInfo) SetUrl(v string) *ListEventsInstancesResponseBodyEventsOnlineMeetingInfo {
	s.Url = &v
	return s
}

type ListEventsInstancesResponseBodyEventsOrganizer struct {
	// 用户名
	DisplayName *string `json:"displayName,omitempty" xml:"displayName,omitempty"`
	// 用户id
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// 回复状态
	ResponseStatus *string `json:"responseStatus,omitempty" xml:"responseStatus,omitempty"`
	// 是否是当前登陆用户
	Self *bool `json:"self,omitempty" xml:"self,omitempty"`
}

func (s ListEventsInstancesResponseBodyEventsOrganizer) String() string {
	return tea.Prettify(s)
}

func (s ListEventsInstancesResponseBodyEventsOrganizer) GoString() string {
	return s.String()
}

func (s *ListEventsInstancesResponseBodyEventsOrganizer) SetDisplayName(v string) *ListEventsInstancesResponseBodyEventsOrganizer {
	s.DisplayName = &v
	return s
}

func (s *ListEventsInstancesResponseBodyEventsOrganizer) SetId(v string) *ListEventsInstancesResponseBodyEventsOrganizer {
	s.Id = &v
	return s
}

func (s *ListEventsInstancesResponseBodyEventsOrganizer) SetResponseStatus(v string) *ListEventsInstancesResponseBodyEventsOrganizer {
	s.ResponseStatus = &v
	return s
}

func (s *ListEventsInstancesResponseBodyEventsOrganizer) SetSelf(v bool) *ListEventsInstancesResponseBodyEventsOrganizer {
	s.Self = &v
	return s
}

type ListEventsInstancesResponseBodyEventsRecurrence struct {
	// 重复模式
	Pattern *ListEventsInstancesResponseBodyEventsRecurrencePattern `json:"pattern,omitempty" xml:"pattern,omitempty" type:"Struct"`
	// 重复范围
	Range *ListEventsInstancesResponseBodyEventsRecurrenceRange `json:"range,omitempty" xml:"range,omitempty" type:"Struct"`
}

func (s ListEventsInstancesResponseBodyEventsRecurrence) String() string {
	return tea.Prettify(s)
}

func (s ListEventsInstancesResponseBodyEventsRecurrence) GoString() string {
	return s.String()
}

func (s *ListEventsInstancesResponseBodyEventsRecurrence) SetPattern(v *ListEventsInstancesResponseBodyEventsRecurrencePattern) *ListEventsInstancesResponseBodyEventsRecurrence {
	s.Pattern = v
	return s
}

func (s *ListEventsInstancesResponseBodyEventsRecurrence) SetRange(v *ListEventsInstancesResponseBodyEventsRecurrenceRange) *ListEventsInstancesResponseBodyEventsRecurrence {
	s.Range = v
	return s
}

type ListEventsInstancesResponseBodyEventsRecurrencePattern struct {
	// 每月的第几天
	DayOfMonth *int32 `json:"dayOfMonth,omitempty" xml:"dayOfMonth,omitempty"`
	// 每周的第几天
	DaysOfWeek *string `json:"daysOfWeek,omitempty" xml:"daysOfWeek,omitempty"`
	// 指定事件发生在daysOfsWeek中指定的允许天数的哪个实例上，从该月的第一个实例开始计算。取值为:first, second, third, fourth, last。默认是first。如果类型是relativMonthly或relativeYear，则可选并使用
	Index *string `json:"index,omitempty" xml:"index,omitempty"`
	// 循环间隔
	Interval *int32 `json:"interval,omitempty" xml:"interval,omitempty"`
	// 循环模式类型(type: daily, weekly, absoluteMonthly, relativeMonthly, absoluteYearly, relativeYearly)
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s ListEventsInstancesResponseBodyEventsRecurrencePattern) String() string {
	return tea.Prettify(s)
}

func (s ListEventsInstancesResponseBodyEventsRecurrencePattern) GoString() string {
	return s.String()
}

func (s *ListEventsInstancesResponseBodyEventsRecurrencePattern) SetDayOfMonth(v int32) *ListEventsInstancesResponseBodyEventsRecurrencePattern {
	s.DayOfMonth = &v
	return s
}

func (s *ListEventsInstancesResponseBodyEventsRecurrencePattern) SetDaysOfWeek(v string) *ListEventsInstancesResponseBodyEventsRecurrencePattern {
	s.DaysOfWeek = &v
	return s
}

func (s *ListEventsInstancesResponseBodyEventsRecurrencePattern) SetIndex(v string) *ListEventsInstancesResponseBodyEventsRecurrencePattern {
	s.Index = &v
	return s
}

func (s *ListEventsInstancesResponseBodyEventsRecurrencePattern) SetInterval(v int32) *ListEventsInstancesResponseBodyEventsRecurrencePattern {
	s.Interval = &v
	return s
}

func (s *ListEventsInstancesResponseBodyEventsRecurrencePattern) SetType(v string) *ListEventsInstancesResponseBodyEventsRecurrencePattern {
	s.Type = &v
	return s
}

type ListEventsInstancesResponseBodyEventsRecurrenceRange struct {
	// 循环终止日期
	EndDate *string `json:"endDate,omitempty" xml:"endDate,omitempty"`
	// 循环出现次数
	NumberOfOccurrences *int32 `json:"numberOfOccurrences,omitempty" xml:"numberOfOccurrences,omitempty"`
	// 范围类型(endDate, noEnd, numbered)
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s ListEventsInstancesResponseBodyEventsRecurrenceRange) String() string {
	return tea.Prettify(s)
}

func (s ListEventsInstancesResponseBodyEventsRecurrenceRange) GoString() string {
	return s.String()
}

func (s *ListEventsInstancesResponseBodyEventsRecurrenceRange) SetEndDate(v string) *ListEventsInstancesResponseBodyEventsRecurrenceRange {
	s.EndDate = &v
	return s
}

func (s *ListEventsInstancesResponseBodyEventsRecurrenceRange) SetNumberOfOccurrences(v int32) *ListEventsInstancesResponseBodyEventsRecurrenceRange {
	s.NumberOfOccurrences = &v
	return s
}

func (s *ListEventsInstancesResponseBodyEventsRecurrenceRange) SetType(v string) *ListEventsInstancesResponseBodyEventsRecurrenceRange {
	s.Type = &v
	return s
}

type ListEventsInstancesResponseBodyEventsReminders struct {
	// 提醒方式
	Method *string `json:"method,omitempty" xml:"method,omitempty"`
	// 在日程开始前N分钟发出提醒
	Minutes *string `json:"minutes,omitempty" xml:"minutes,omitempty"`
}

func (s ListEventsInstancesResponseBodyEventsReminders) String() string {
	return tea.Prettify(s)
}

func (s ListEventsInstancesResponseBodyEventsReminders) GoString() string {
	return s.String()
}

func (s *ListEventsInstancesResponseBodyEventsReminders) SetMethod(v string) *ListEventsInstancesResponseBodyEventsReminders {
	s.Method = &v
	return s
}

func (s *ListEventsInstancesResponseBodyEventsReminders) SetMinutes(v string) *ListEventsInstancesResponseBodyEventsReminders {
	s.Minutes = &v
	return s
}

type ListEventsInstancesResponseBodyEventsStart struct {
	// 日期，格式：yyyyMMdd
	Date *string `json:"date,omitempty" xml:"date,omitempty"`
	// 时间戳，按照ISO 8601格式
	DateTime *string `json:"dateTime,omitempty" xml:"dateTime,omitempty"`
	// 时区
	TimeZone *string `json:"timeZone,omitempty" xml:"timeZone,omitempty"`
}

func (s ListEventsInstancesResponseBodyEventsStart) String() string {
	return tea.Prettify(s)
}

func (s ListEventsInstancesResponseBodyEventsStart) GoString() string {
	return s.String()
}

func (s *ListEventsInstancesResponseBodyEventsStart) SetDate(v string) *ListEventsInstancesResponseBodyEventsStart {
	s.Date = &v
	return s
}

func (s *ListEventsInstancesResponseBodyEventsStart) SetDateTime(v string) *ListEventsInstancesResponseBodyEventsStart {
	s.DateTime = &v
	return s
}

func (s *ListEventsInstancesResponseBodyEventsStart) SetTimeZone(v string) *ListEventsInstancesResponseBodyEventsStart {
	s.TimeZone = &v
	return s
}

type ListEventsInstancesResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListEventsInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListEventsInstancesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListEventsInstancesResponse) GoString() string {
	return s.String()
}

func (s *ListEventsInstancesResponse) SetHeaders(v map[string]*string) *ListEventsInstancesResponse {
	s.Headers = v
	return s
}

func (s *ListEventsInstancesResponse) SetBody(v *ListEventsInstancesResponseBody) *ListEventsInstancesResponse {
	s.Body = v
	return s
}

type ListEventsViewHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s ListEventsViewHeaders) String() string {
	return tea.Prettify(s)
}

func (s ListEventsViewHeaders) GoString() string {
	return s.String()
}

func (s *ListEventsViewHeaders) SetCommonHeaders(v map[string]*string) *ListEventsViewHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListEventsViewHeaders) SetXAcsDingtalkAccessToken(v string) *ListEventsViewHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type ListEventsViewRequest struct {
	// 每个日程的参与者查询个数，默认100，最大100。
	MaxAttendees *int32 `json:"maxAttendees,omitempty" xml:"maxAttendees,omitempty"`
	// 返回的最大日程数，最大100个，默认100个。
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// 查询翻页token
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// 查询截止时间
	TimeMax *string `json:"timeMax,omitempty" xml:"timeMax,omitempty"`
	// 查询开始时间
	TimeMin *string `json:"timeMin,omitempty" xml:"timeMin,omitempty"`
}

func (s ListEventsViewRequest) String() string {
	return tea.Prettify(s)
}

func (s ListEventsViewRequest) GoString() string {
	return s.String()
}

func (s *ListEventsViewRequest) SetMaxAttendees(v int32) *ListEventsViewRequest {
	s.MaxAttendees = &v
	return s
}

func (s *ListEventsViewRequest) SetMaxResults(v int32) *ListEventsViewRequest {
	s.MaxResults = &v
	return s
}

func (s *ListEventsViewRequest) SetNextToken(v string) *ListEventsViewRequest {
	s.NextToken = &v
	return s
}

func (s *ListEventsViewRequest) SetTimeMax(v string) *ListEventsViewRequest {
	s.TimeMax = &v
	return s
}

func (s *ListEventsViewRequest) SetTimeMin(v string) *ListEventsViewRequest {
	s.TimeMin = &v
	return s
}

type ListEventsViewResponseBody struct {
	// 日程
	Events []*ListEventsViewResponseBodyEvents `json:"events,omitempty" xml:"events,omitempty" type:"Repeated"`
	// 翻页token
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
}

func (s ListEventsViewResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListEventsViewResponseBody) GoString() string {
	return s.String()
}

func (s *ListEventsViewResponseBody) SetEvents(v []*ListEventsViewResponseBodyEvents) *ListEventsViewResponseBody {
	s.Events = v
	return s
}

func (s *ListEventsViewResponseBody) SetNextToken(v string) *ListEventsViewResponseBody {
	s.NextToken = &v
	return s
}

type ListEventsViewResponseBodyEvents struct {
	// 日程参与人
	Attendees []*ListEventsViewResponseBodyEventsAttendees `json:"attendees,omitempty" xml:"attendees,omitempty" type:"Repeated"`
	// 创建时间
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// 日程描述
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// 日程结束时间
	End                *ListEventsViewResponseBodyEventsEnd                `json:"end,omitempty" xml:"end,omitempty" type:"Struct"`
	ExtendedProperties *ListEventsViewResponseBodyEventsExtendedProperties `json:"extendedProperties,omitempty" xml:"extendedProperties,omitempty" type:"Struct"`
	// 日程事件id
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// 是否为全天日程
	IsAllDay *bool `json:"isAllDay,omitempty" xml:"isAllDay,omitempty"`
	// 日程地点
	Location          *ListEventsViewResponseBodyEventsLocation          `json:"location,omitempty" xml:"location,omitempty" type:"Struct"`
	OnlineMeetingInfo *ListEventsViewResponseBodyEventsOnlineMeetingInfo `json:"onlineMeetingInfo,omitempty" xml:"onlineMeetingInfo,omitempty" type:"Struct"`
	// 日程组织人
	Organizer *ListEventsViewResponseBodyEventsOrganizer `json:"organizer,omitempty" xml:"organizer,omitempty" type:"Struct"`
	// 日程重复规则
	Recurrence *ListEventsViewResponseBodyEventsRecurrence `json:"recurrence,omitempty" xml:"recurrence,omitempty" type:"Struct"`
	// 重复日程的主日程id，非重复日程为空
	SeriesMasterId *string `json:"seriesMasterId,omitempty" xml:"seriesMasterId,omitempty"`
	// 日程开始时间
	Start *ListEventsViewResponseBodyEventsStart `json:"start,omitempty" xml:"start,omitempty" type:"Struct"`
	// 日程状态
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// 日程标题
	Summary *string `json:"summary,omitempty" xml:"summary,omitempty"`
	// 更新时间
	UpdateTime *string `json:"updateTime,omitempty" xml:"updateTime,omitempty"`
}

func (s ListEventsViewResponseBodyEvents) String() string {
	return tea.Prettify(s)
}

func (s ListEventsViewResponseBodyEvents) GoString() string {
	return s.String()
}

func (s *ListEventsViewResponseBodyEvents) SetAttendees(v []*ListEventsViewResponseBodyEventsAttendees) *ListEventsViewResponseBodyEvents {
	s.Attendees = v
	return s
}

func (s *ListEventsViewResponseBodyEvents) SetCreateTime(v string) *ListEventsViewResponseBodyEvents {
	s.CreateTime = &v
	return s
}

func (s *ListEventsViewResponseBodyEvents) SetDescription(v string) *ListEventsViewResponseBodyEvents {
	s.Description = &v
	return s
}

func (s *ListEventsViewResponseBodyEvents) SetEnd(v *ListEventsViewResponseBodyEventsEnd) *ListEventsViewResponseBodyEvents {
	s.End = v
	return s
}

func (s *ListEventsViewResponseBodyEvents) SetExtendedProperties(v *ListEventsViewResponseBodyEventsExtendedProperties) *ListEventsViewResponseBodyEvents {
	s.ExtendedProperties = v
	return s
}

func (s *ListEventsViewResponseBodyEvents) SetId(v string) *ListEventsViewResponseBodyEvents {
	s.Id = &v
	return s
}

func (s *ListEventsViewResponseBodyEvents) SetIsAllDay(v bool) *ListEventsViewResponseBodyEvents {
	s.IsAllDay = &v
	return s
}

func (s *ListEventsViewResponseBodyEvents) SetLocation(v *ListEventsViewResponseBodyEventsLocation) *ListEventsViewResponseBodyEvents {
	s.Location = v
	return s
}

func (s *ListEventsViewResponseBodyEvents) SetOnlineMeetingInfo(v *ListEventsViewResponseBodyEventsOnlineMeetingInfo) *ListEventsViewResponseBodyEvents {
	s.OnlineMeetingInfo = v
	return s
}

func (s *ListEventsViewResponseBodyEvents) SetOrganizer(v *ListEventsViewResponseBodyEventsOrganizer) *ListEventsViewResponseBodyEvents {
	s.Organizer = v
	return s
}

func (s *ListEventsViewResponseBodyEvents) SetRecurrence(v *ListEventsViewResponseBodyEventsRecurrence) *ListEventsViewResponseBodyEvents {
	s.Recurrence = v
	return s
}

func (s *ListEventsViewResponseBodyEvents) SetSeriesMasterId(v string) *ListEventsViewResponseBodyEvents {
	s.SeriesMasterId = &v
	return s
}

func (s *ListEventsViewResponseBodyEvents) SetStart(v *ListEventsViewResponseBodyEventsStart) *ListEventsViewResponseBodyEvents {
	s.Start = v
	return s
}

func (s *ListEventsViewResponseBodyEvents) SetStatus(v string) *ListEventsViewResponseBodyEvents {
	s.Status = &v
	return s
}

func (s *ListEventsViewResponseBodyEvents) SetSummary(v string) *ListEventsViewResponseBodyEvents {
	s.Summary = &v
	return s
}

func (s *ListEventsViewResponseBodyEvents) SetUpdateTime(v string) *ListEventsViewResponseBodyEvents {
	s.UpdateTime = &v
	return s
}

type ListEventsViewResponseBodyEventsAttendees struct {
	// 用户名
	DisplayName *string `json:"displayName,omitempty" xml:"displayName,omitempty"`
	// 用户id
	Id         *string `json:"id,omitempty" xml:"id,omitempty"`
	IsOptional *bool   `json:"isOptional,omitempty" xml:"isOptional,omitempty"`
	// 回复状态
	ResponseStatus *string `json:"responseStatus,omitempty" xml:"responseStatus,omitempty"`
	// 是否是当前登陆用户
	Self *bool `json:"self,omitempty" xml:"self,omitempty"`
}

func (s ListEventsViewResponseBodyEventsAttendees) String() string {
	return tea.Prettify(s)
}

func (s ListEventsViewResponseBodyEventsAttendees) GoString() string {
	return s.String()
}

func (s *ListEventsViewResponseBodyEventsAttendees) SetDisplayName(v string) *ListEventsViewResponseBodyEventsAttendees {
	s.DisplayName = &v
	return s
}

func (s *ListEventsViewResponseBodyEventsAttendees) SetId(v string) *ListEventsViewResponseBodyEventsAttendees {
	s.Id = &v
	return s
}

func (s *ListEventsViewResponseBodyEventsAttendees) SetIsOptional(v bool) *ListEventsViewResponseBodyEventsAttendees {
	s.IsOptional = &v
	return s
}

func (s *ListEventsViewResponseBodyEventsAttendees) SetResponseStatus(v string) *ListEventsViewResponseBodyEventsAttendees {
	s.ResponseStatus = &v
	return s
}

func (s *ListEventsViewResponseBodyEventsAttendees) SetSelf(v bool) *ListEventsViewResponseBodyEventsAttendees {
	s.Self = &v
	return s
}

type ListEventsViewResponseBodyEventsEnd struct {
	Date     *string `json:"date,omitempty" xml:"date,omitempty"`
	DateTime *string `json:"dateTime,omitempty" xml:"dateTime,omitempty"`
	TimeZone *string `json:"timeZone,omitempty" xml:"timeZone,omitempty"`
}

func (s ListEventsViewResponseBodyEventsEnd) String() string {
	return tea.Prettify(s)
}

func (s ListEventsViewResponseBodyEventsEnd) GoString() string {
	return s.String()
}

func (s *ListEventsViewResponseBodyEventsEnd) SetDate(v string) *ListEventsViewResponseBodyEventsEnd {
	s.Date = &v
	return s
}

func (s *ListEventsViewResponseBodyEventsEnd) SetDateTime(v string) *ListEventsViewResponseBodyEventsEnd {
	s.DateTime = &v
	return s
}

func (s *ListEventsViewResponseBodyEventsEnd) SetTimeZone(v string) *ListEventsViewResponseBodyEventsEnd {
	s.TimeZone = &v
	return s
}

type ListEventsViewResponseBodyEventsExtendedProperties struct {
	SharedProperties *ListEventsViewResponseBodyEventsExtendedPropertiesSharedProperties `json:"sharedProperties,omitempty" xml:"sharedProperties,omitempty" type:"Struct"`
}

func (s ListEventsViewResponseBodyEventsExtendedProperties) String() string {
	return tea.Prettify(s)
}

func (s ListEventsViewResponseBodyEventsExtendedProperties) GoString() string {
	return s.String()
}

func (s *ListEventsViewResponseBodyEventsExtendedProperties) SetSharedProperties(v *ListEventsViewResponseBodyEventsExtendedPropertiesSharedProperties) *ListEventsViewResponseBodyEventsExtendedProperties {
	s.SharedProperties = v
	return s
}

type ListEventsViewResponseBodyEventsExtendedPropertiesSharedProperties struct {
	SourceOpenCid *string `json:"sourceOpenCid,omitempty" xml:"sourceOpenCid,omitempty"`
}

func (s ListEventsViewResponseBodyEventsExtendedPropertiesSharedProperties) String() string {
	return tea.Prettify(s)
}

func (s ListEventsViewResponseBodyEventsExtendedPropertiesSharedProperties) GoString() string {
	return s.String()
}

func (s *ListEventsViewResponseBodyEventsExtendedPropertiesSharedProperties) SetSourceOpenCid(v string) *ListEventsViewResponseBodyEventsExtendedPropertiesSharedProperties {
	s.SourceOpenCid = &v
	return s
}

type ListEventsViewResponseBodyEventsLocation struct {
	// 展示名称
	DisplayName  *string   `json:"displayName,omitempty" xml:"displayName,omitempty"`
	MeetingRooms []*string `json:"meetingRooms,omitempty" xml:"meetingRooms,omitempty" type:"Repeated"`
}

func (s ListEventsViewResponseBodyEventsLocation) String() string {
	return tea.Prettify(s)
}

func (s ListEventsViewResponseBodyEventsLocation) GoString() string {
	return s.String()
}

func (s *ListEventsViewResponseBodyEventsLocation) SetDisplayName(v string) *ListEventsViewResponseBodyEventsLocation {
	s.DisplayName = &v
	return s
}

func (s *ListEventsViewResponseBodyEventsLocation) SetMeetingRooms(v []*string) *ListEventsViewResponseBodyEventsLocation {
	s.MeetingRooms = v
	return s
}

type ListEventsViewResponseBodyEventsOnlineMeetingInfo struct {
	ConferenceId *string                `json:"conferenceId,omitempty" xml:"conferenceId,omitempty"`
	ExtraInfo    map[string]interface{} `json:"extraInfo,omitempty" xml:"extraInfo,omitempty"`
	Type         *string                `json:"type,omitempty" xml:"type,omitempty"`
	Url          *string                `json:"url,omitempty" xml:"url,omitempty"`
}

func (s ListEventsViewResponseBodyEventsOnlineMeetingInfo) String() string {
	return tea.Prettify(s)
}

func (s ListEventsViewResponseBodyEventsOnlineMeetingInfo) GoString() string {
	return s.String()
}

func (s *ListEventsViewResponseBodyEventsOnlineMeetingInfo) SetConferenceId(v string) *ListEventsViewResponseBodyEventsOnlineMeetingInfo {
	s.ConferenceId = &v
	return s
}

func (s *ListEventsViewResponseBodyEventsOnlineMeetingInfo) SetExtraInfo(v map[string]interface{}) *ListEventsViewResponseBodyEventsOnlineMeetingInfo {
	s.ExtraInfo = v
	return s
}

func (s *ListEventsViewResponseBodyEventsOnlineMeetingInfo) SetType(v string) *ListEventsViewResponseBodyEventsOnlineMeetingInfo {
	s.Type = &v
	return s
}

func (s *ListEventsViewResponseBodyEventsOnlineMeetingInfo) SetUrl(v string) *ListEventsViewResponseBodyEventsOnlineMeetingInfo {
	s.Url = &v
	return s
}

type ListEventsViewResponseBodyEventsOrganizer struct {
	// 用户名
	DisplayName *string `json:"displayName,omitempty" xml:"displayName,omitempty"`
	// 用户id
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// 回复状态
	ResponseStatus *string `json:"responseStatus,omitempty" xml:"responseStatus,omitempty"`
	// 是否是当前登陆用户
	Self *bool `json:"self,omitempty" xml:"self,omitempty"`
}

func (s ListEventsViewResponseBodyEventsOrganizer) String() string {
	return tea.Prettify(s)
}

func (s ListEventsViewResponseBodyEventsOrganizer) GoString() string {
	return s.String()
}

func (s *ListEventsViewResponseBodyEventsOrganizer) SetDisplayName(v string) *ListEventsViewResponseBodyEventsOrganizer {
	s.DisplayName = &v
	return s
}

func (s *ListEventsViewResponseBodyEventsOrganizer) SetId(v string) *ListEventsViewResponseBodyEventsOrganizer {
	s.Id = &v
	return s
}

func (s *ListEventsViewResponseBodyEventsOrganizer) SetResponseStatus(v string) *ListEventsViewResponseBodyEventsOrganizer {
	s.ResponseStatus = &v
	return s
}

func (s *ListEventsViewResponseBodyEventsOrganizer) SetSelf(v bool) *ListEventsViewResponseBodyEventsOrganizer {
	s.Self = &v
	return s
}

type ListEventsViewResponseBodyEventsRecurrence struct {
	// 重复模式
	Pattern *ListEventsViewResponseBodyEventsRecurrencePattern `json:"pattern,omitempty" xml:"pattern,omitempty" type:"Struct"`
	// 重复范围
	Range *ListEventsViewResponseBodyEventsRecurrenceRange `json:"range,omitempty" xml:"range,omitempty" type:"Struct"`
}

func (s ListEventsViewResponseBodyEventsRecurrence) String() string {
	return tea.Prettify(s)
}

func (s ListEventsViewResponseBodyEventsRecurrence) GoString() string {
	return s.String()
}

func (s *ListEventsViewResponseBodyEventsRecurrence) SetPattern(v *ListEventsViewResponseBodyEventsRecurrencePattern) *ListEventsViewResponseBodyEventsRecurrence {
	s.Pattern = v
	return s
}

func (s *ListEventsViewResponseBodyEventsRecurrence) SetRange(v *ListEventsViewResponseBodyEventsRecurrenceRange) *ListEventsViewResponseBodyEventsRecurrence {
	s.Range = v
	return s
}

type ListEventsViewResponseBodyEventsRecurrencePattern struct {
	DayOfMonth *int32  `json:"dayOfMonth,omitempty" xml:"dayOfMonth,omitempty"`
	DaysOfWeek *string `json:"daysOfWeek,omitempty" xml:"daysOfWeek,omitempty"`
	Index      *string `json:"index,omitempty" xml:"index,omitempty"`
	Interval   *int32  `json:"interval,omitempty" xml:"interval,omitempty"`
	// 循环模式类型(type: daily, weekly, absoluteMonthly, relativeMonthly, absoluteYearly, relativeYearly)
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s ListEventsViewResponseBodyEventsRecurrencePattern) String() string {
	return tea.Prettify(s)
}

func (s ListEventsViewResponseBodyEventsRecurrencePattern) GoString() string {
	return s.String()
}

func (s *ListEventsViewResponseBodyEventsRecurrencePattern) SetDayOfMonth(v int32) *ListEventsViewResponseBodyEventsRecurrencePattern {
	s.DayOfMonth = &v
	return s
}

func (s *ListEventsViewResponseBodyEventsRecurrencePattern) SetDaysOfWeek(v string) *ListEventsViewResponseBodyEventsRecurrencePattern {
	s.DaysOfWeek = &v
	return s
}

func (s *ListEventsViewResponseBodyEventsRecurrencePattern) SetIndex(v string) *ListEventsViewResponseBodyEventsRecurrencePattern {
	s.Index = &v
	return s
}

func (s *ListEventsViewResponseBodyEventsRecurrencePattern) SetInterval(v int32) *ListEventsViewResponseBodyEventsRecurrencePattern {
	s.Interval = &v
	return s
}

func (s *ListEventsViewResponseBodyEventsRecurrencePattern) SetType(v string) *ListEventsViewResponseBodyEventsRecurrencePattern {
	s.Type = &v
	return s
}

type ListEventsViewResponseBodyEventsRecurrenceRange struct {
	EndDate             *string `json:"endDate,omitempty" xml:"endDate,omitempty"`
	NumberOfOccurrences *int32  `json:"numberOfOccurrences,omitempty" xml:"numberOfOccurrences,omitempty"`
	// 范围类型(endDate, noEnd, numbered)
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s ListEventsViewResponseBodyEventsRecurrenceRange) String() string {
	return tea.Prettify(s)
}

func (s ListEventsViewResponseBodyEventsRecurrenceRange) GoString() string {
	return s.String()
}

func (s *ListEventsViewResponseBodyEventsRecurrenceRange) SetEndDate(v string) *ListEventsViewResponseBodyEventsRecurrenceRange {
	s.EndDate = &v
	return s
}

func (s *ListEventsViewResponseBodyEventsRecurrenceRange) SetNumberOfOccurrences(v int32) *ListEventsViewResponseBodyEventsRecurrenceRange {
	s.NumberOfOccurrences = &v
	return s
}

func (s *ListEventsViewResponseBodyEventsRecurrenceRange) SetType(v string) *ListEventsViewResponseBodyEventsRecurrenceRange {
	s.Type = &v
	return s
}

type ListEventsViewResponseBodyEventsStart struct {
	// 日期，格式：yyyyMMdd
	Date *string `json:"date,omitempty" xml:"date,omitempty"`
	// 时间戳，按照ISO 8601格式
	DateTime *string `json:"dateTime,omitempty" xml:"dateTime,omitempty"`
	// 时区
	TimeZone *string `json:"timeZone,omitempty" xml:"timeZone,omitempty"`
}

func (s ListEventsViewResponseBodyEventsStart) String() string {
	return tea.Prettify(s)
}

func (s ListEventsViewResponseBodyEventsStart) GoString() string {
	return s.String()
}

func (s *ListEventsViewResponseBodyEventsStart) SetDate(v string) *ListEventsViewResponseBodyEventsStart {
	s.Date = &v
	return s
}

func (s *ListEventsViewResponseBodyEventsStart) SetDateTime(v string) *ListEventsViewResponseBodyEventsStart {
	s.DateTime = &v
	return s
}

func (s *ListEventsViewResponseBodyEventsStart) SetTimeZone(v string) *ListEventsViewResponseBodyEventsStart {
	s.TimeZone = &v
	return s
}

type ListEventsViewResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListEventsViewResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListEventsViewResponse) String() string {
	return tea.Prettify(s)
}

func (s ListEventsViewResponse) GoString() string {
	return s.String()
}

func (s *ListEventsViewResponse) SetHeaders(v map[string]*string) *ListEventsViewResponse {
	s.Headers = v
	return s
}

func (s *ListEventsViewResponse) SetBody(v *ListEventsViewResponseBody) *ListEventsViewResponse {
	s.Body = v
	return s
}

type PatchEventHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s PatchEventHeaders) String() string {
	return tea.Prettify(s)
}

func (s PatchEventHeaders) GoString() string {
	return s.String()
}

func (s *PatchEventHeaders) SetCommonHeaders(v map[string]*string) *PatchEventHeaders {
	s.CommonHeaders = v
	return s
}

func (s *PatchEventHeaders) SetXAcsDingtalkAccessToken(v string) *PatchEventHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type PatchEventRequest struct {
	Attendees   []*PatchEventRequestAttendees `json:"attendees,omitempty" xml:"attendees,omitempty" type:"Repeated"`
	Description *string                       `json:"description,omitempty" xml:"description,omitempty"`
	End         *PatchEventRequestEnd         `json:"end,omitempty" xml:"end,omitempty" type:"Struct"`
	// 扩展信息
	Extra map[string]*string `json:"extra,omitempty" xml:"extra,omitempty"`
	// 日程id
	Id         *string                       `json:"id,omitempty" xml:"id,omitempty"`
	IsAllDay   *bool                         `json:"isAllDay,omitempty" xml:"isAllDay,omitempty"`
	Location   *PatchEventRequestLocation    `json:"location,omitempty" xml:"location,omitempty" type:"Struct"`
	Recurrence *PatchEventRequestRecurrence  `json:"recurrence,omitempty" xml:"recurrence,omitempty" type:"Struct"`
	Reminders  []*PatchEventRequestReminders `json:"reminders,omitempty" xml:"reminders,omitempty" type:"Repeated"`
	// 日程开始时间
	Start *PatchEventRequestStart `json:"start,omitempty" xml:"start,omitempty" type:"Struct"`
	// 日程标题
	Summary *string `json:"summary,omitempty" xml:"summary,omitempty"`
}

func (s PatchEventRequest) String() string {
	return tea.Prettify(s)
}

func (s PatchEventRequest) GoString() string {
	return s.String()
}

func (s *PatchEventRequest) SetAttendees(v []*PatchEventRequestAttendees) *PatchEventRequest {
	s.Attendees = v
	return s
}

func (s *PatchEventRequest) SetDescription(v string) *PatchEventRequest {
	s.Description = &v
	return s
}

func (s *PatchEventRequest) SetEnd(v *PatchEventRequestEnd) *PatchEventRequest {
	s.End = v
	return s
}

func (s *PatchEventRequest) SetExtra(v map[string]*string) *PatchEventRequest {
	s.Extra = v
	return s
}

func (s *PatchEventRequest) SetId(v string) *PatchEventRequest {
	s.Id = &v
	return s
}

func (s *PatchEventRequest) SetIsAllDay(v bool) *PatchEventRequest {
	s.IsAllDay = &v
	return s
}

func (s *PatchEventRequest) SetLocation(v *PatchEventRequestLocation) *PatchEventRequest {
	s.Location = v
	return s
}

func (s *PatchEventRequest) SetRecurrence(v *PatchEventRequestRecurrence) *PatchEventRequest {
	s.Recurrence = v
	return s
}

func (s *PatchEventRequest) SetReminders(v []*PatchEventRequestReminders) *PatchEventRequest {
	s.Reminders = v
	return s
}

func (s *PatchEventRequest) SetStart(v *PatchEventRequestStart) *PatchEventRequest {
	s.Start = v
	return s
}

func (s *PatchEventRequest) SetSummary(v string) *PatchEventRequest {
	s.Summary = &v
	return s
}

type PatchEventRequestAttendees struct {
	Id         *string `json:"id,omitempty" xml:"id,omitempty"`
	IsOptional *bool   `json:"isOptional,omitempty" xml:"isOptional,omitempty"`
}

func (s PatchEventRequestAttendees) String() string {
	return tea.Prettify(s)
}

func (s PatchEventRequestAttendees) GoString() string {
	return s.String()
}

func (s *PatchEventRequestAttendees) SetId(v string) *PatchEventRequestAttendees {
	s.Id = &v
	return s
}

func (s *PatchEventRequestAttendees) SetIsOptional(v bool) *PatchEventRequestAttendees {
	s.IsOptional = &v
	return s
}

type PatchEventRequestEnd struct {
	Date     *string `json:"date,omitempty" xml:"date,omitempty"`
	DateTime *string `json:"dateTime,omitempty" xml:"dateTime,omitempty"`
	TimeZone *string `json:"timeZone,omitempty" xml:"timeZone,omitempty"`
}

func (s PatchEventRequestEnd) String() string {
	return tea.Prettify(s)
}

func (s PatchEventRequestEnd) GoString() string {
	return s.String()
}

func (s *PatchEventRequestEnd) SetDate(v string) *PatchEventRequestEnd {
	s.Date = &v
	return s
}

func (s *PatchEventRequestEnd) SetDateTime(v string) *PatchEventRequestEnd {
	s.DateTime = &v
	return s
}

func (s *PatchEventRequestEnd) SetTimeZone(v string) *PatchEventRequestEnd {
	s.TimeZone = &v
	return s
}

type PatchEventRequestLocation struct {
	DisplayName *string `json:"displayName,omitempty" xml:"displayName,omitempty"`
}

func (s PatchEventRequestLocation) String() string {
	return tea.Prettify(s)
}

func (s PatchEventRequestLocation) GoString() string {
	return s.String()
}

func (s *PatchEventRequestLocation) SetDisplayName(v string) *PatchEventRequestLocation {
	s.DisplayName = &v
	return s
}

type PatchEventRequestRecurrence struct {
	Pattern *PatchEventRequestRecurrencePattern `json:"pattern,omitempty" xml:"pattern,omitempty" type:"Struct"`
	Range   *PatchEventRequestRecurrenceRange   `json:"range,omitempty" xml:"range,omitempty" type:"Struct"`
}

func (s PatchEventRequestRecurrence) String() string {
	return tea.Prettify(s)
}

func (s PatchEventRequestRecurrence) GoString() string {
	return s.String()
}

func (s *PatchEventRequestRecurrence) SetPattern(v *PatchEventRequestRecurrencePattern) *PatchEventRequestRecurrence {
	s.Pattern = v
	return s
}

func (s *PatchEventRequestRecurrence) SetRange(v *PatchEventRequestRecurrenceRange) *PatchEventRequestRecurrence {
	s.Range = v
	return s
}

type PatchEventRequestRecurrencePattern struct {
	DayOfMonth *int32  `json:"dayOfMonth,omitempty" xml:"dayOfMonth,omitempty"`
	DaysOfWeek *string `json:"daysOfWeek,omitempty" xml:"daysOfWeek,omitempty"`
	Index      *string `json:"index,omitempty" xml:"index,omitempty"`
	Interval   *int32  `json:"interval,omitempty" xml:"interval,omitempty"`
	Type       *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s PatchEventRequestRecurrencePattern) String() string {
	return tea.Prettify(s)
}

func (s PatchEventRequestRecurrencePattern) GoString() string {
	return s.String()
}

func (s *PatchEventRequestRecurrencePattern) SetDayOfMonth(v int32) *PatchEventRequestRecurrencePattern {
	s.DayOfMonth = &v
	return s
}

func (s *PatchEventRequestRecurrencePattern) SetDaysOfWeek(v string) *PatchEventRequestRecurrencePattern {
	s.DaysOfWeek = &v
	return s
}

func (s *PatchEventRequestRecurrencePattern) SetIndex(v string) *PatchEventRequestRecurrencePattern {
	s.Index = &v
	return s
}

func (s *PatchEventRequestRecurrencePattern) SetInterval(v int32) *PatchEventRequestRecurrencePattern {
	s.Interval = &v
	return s
}

func (s *PatchEventRequestRecurrencePattern) SetType(v string) *PatchEventRequestRecurrencePattern {
	s.Type = &v
	return s
}

type PatchEventRequestRecurrenceRange struct {
	EndDate             *string `json:"endDate,omitempty" xml:"endDate,omitempty"`
	NumberOfOccurrences *int32  `json:"numberOfOccurrences,omitempty" xml:"numberOfOccurrences,omitempty"`
	Type                *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s PatchEventRequestRecurrenceRange) String() string {
	return tea.Prettify(s)
}

func (s PatchEventRequestRecurrenceRange) GoString() string {
	return s.String()
}

func (s *PatchEventRequestRecurrenceRange) SetEndDate(v string) *PatchEventRequestRecurrenceRange {
	s.EndDate = &v
	return s
}

func (s *PatchEventRequestRecurrenceRange) SetNumberOfOccurrences(v int32) *PatchEventRequestRecurrenceRange {
	s.NumberOfOccurrences = &v
	return s
}

func (s *PatchEventRequestRecurrenceRange) SetType(v string) *PatchEventRequestRecurrenceRange {
	s.Type = &v
	return s
}

type PatchEventRequestReminders struct {
	Method  *string `json:"method,omitempty" xml:"method,omitempty"`
	Minutes *int32  `json:"minutes,omitempty" xml:"minutes,omitempty"`
}

func (s PatchEventRequestReminders) String() string {
	return tea.Prettify(s)
}

func (s PatchEventRequestReminders) GoString() string {
	return s.String()
}

func (s *PatchEventRequestReminders) SetMethod(v string) *PatchEventRequestReminders {
	s.Method = &v
	return s
}

func (s *PatchEventRequestReminders) SetMinutes(v int32) *PatchEventRequestReminders {
	s.Minutes = &v
	return s
}

type PatchEventRequestStart struct {
	Date     *string `json:"date,omitempty" xml:"date,omitempty"`
	DateTime *string `json:"dateTime,omitempty" xml:"dateTime,omitempty"`
	TimeZone *string `json:"timeZone,omitempty" xml:"timeZone,omitempty"`
}

func (s PatchEventRequestStart) String() string {
	return tea.Prettify(s)
}

func (s PatchEventRequestStart) GoString() string {
	return s.String()
}

func (s *PatchEventRequestStart) SetDate(v string) *PatchEventRequestStart {
	s.Date = &v
	return s
}

func (s *PatchEventRequestStart) SetDateTime(v string) *PatchEventRequestStart {
	s.DateTime = &v
	return s
}

func (s *PatchEventRequestStart) SetTimeZone(v string) *PatchEventRequestStart {
	s.TimeZone = &v
	return s
}

type PatchEventResponseBody struct {
	Attendees []*PatchEventResponseBodyAttendees `json:"attendees,omitempty" xml:"attendees,omitempty" type:"Repeated"`
	// 创建时间
	CreateTime  *string                            `json:"createTime,omitempty" xml:"createTime,omitempty"`
	Description *string                            `json:"description,omitempty" xml:"description,omitempty"`
	End         *PatchEventResponseBodyEnd         `json:"end,omitempty" xml:"end,omitempty" type:"Struct"`
	Id          *string                            `json:"id,omitempty" xml:"id,omitempty"`
	IsAllDay    *bool                              `json:"isAllDay,omitempty" xml:"isAllDay,omitempty"`
	Location    *PatchEventResponseBodyLocation    `json:"location,omitempty" xml:"location,omitempty" type:"Struct"`
	Organizer   *PatchEventResponseBodyOrganizer   `json:"organizer,omitempty" xml:"organizer,omitempty" type:"Struct"`
	Recurrence  *PatchEventResponseBodyRecurrence  `json:"recurrence,omitempty" xml:"recurrence,omitempty" type:"Struct"`
	Reminders   []*PatchEventResponseBodyReminders `json:"reminders,omitempty" xml:"reminders,omitempty" type:"Repeated"`
	// 日程开始时间
	Start   *PatchEventResponseBodyStart `json:"start,omitempty" xml:"start,omitempty" type:"Struct"`
	Summary *string                      `json:"summary,omitempty" xml:"summary,omitempty"`
	// 更新时间
	UpdateTime *string `json:"updateTime,omitempty" xml:"updateTime,omitempty"`
}

func (s PatchEventResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PatchEventResponseBody) GoString() string {
	return s.String()
}

func (s *PatchEventResponseBody) SetAttendees(v []*PatchEventResponseBodyAttendees) *PatchEventResponseBody {
	s.Attendees = v
	return s
}

func (s *PatchEventResponseBody) SetCreateTime(v string) *PatchEventResponseBody {
	s.CreateTime = &v
	return s
}

func (s *PatchEventResponseBody) SetDescription(v string) *PatchEventResponseBody {
	s.Description = &v
	return s
}

func (s *PatchEventResponseBody) SetEnd(v *PatchEventResponseBodyEnd) *PatchEventResponseBody {
	s.End = v
	return s
}

func (s *PatchEventResponseBody) SetId(v string) *PatchEventResponseBody {
	s.Id = &v
	return s
}

func (s *PatchEventResponseBody) SetIsAllDay(v bool) *PatchEventResponseBody {
	s.IsAllDay = &v
	return s
}

func (s *PatchEventResponseBody) SetLocation(v *PatchEventResponseBodyLocation) *PatchEventResponseBody {
	s.Location = v
	return s
}

func (s *PatchEventResponseBody) SetOrganizer(v *PatchEventResponseBodyOrganizer) *PatchEventResponseBody {
	s.Organizer = v
	return s
}

func (s *PatchEventResponseBody) SetRecurrence(v *PatchEventResponseBodyRecurrence) *PatchEventResponseBody {
	s.Recurrence = v
	return s
}

func (s *PatchEventResponseBody) SetReminders(v []*PatchEventResponseBodyReminders) *PatchEventResponseBody {
	s.Reminders = v
	return s
}

func (s *PatchEventResponseBody) SetStart(v *PatchEventResponseBodyStart) *PatchEventResponseBody {
	s.Start = v
	return s
}

func (s *PatchEventResponseBody) SetSummary(v string) *PatchEventResponseBody {
	s.Summary = &v
	return s
}

func (s *PatchEventResponseBody) SetUpdateTime(v string) *PatchEventResponseBody {
	s.UpdateTime = &v
	return s
}

type PatchEventResponseBodyAttendees struct {
	// 用户名
	DisplayName *string `json:"displayName,omitempty" xml:"displayName,omitempty"`
	Id          *string `json:"id,omitempty" xml:"id,omitempty"`
	IsOptional  *bool   `json:"isOptional,omitempty" xml:"isOptional,omitempty"`
	// 回复状态
	ResponseStatus *string `json:"responseStatus,omitempty" xml:"responseStatus,omitempty"`
	// 是否是当前登陆用户
	Self *bool `json:"self,omitempty" xml:"self,omitempty"`
}

func (s PatchEventResponseBodyAttendees) String() string {
	return tea.Prettify(s)
}

func (s PatchEventResponseBodyAttendees) GoString() string {
	return s.String()
}

func (s *PatchEventResponseBodyAttendees) SetDisplayName(v string) *PatchEventResponseBodyAttendees {
	s.DisplayName = &v
	return s
}

func (s *PatchEventResponseBodyAttendees) SetId(v string) *PatchEventResponseBodyAttendees {
	s.Id = &v
	return s
}

func (s *PatchEventResponseBodyAttendees) SetIsOptional(v bool) *PatchEventResponseBodyAttendees {
	s.IsOptional = &v
	return s
}

func (s *PatchEventResponseBodyAttendees) SetResponseStatus(v string) *PatchEventResponseBodyAttendees {
	s.ResponseStatus = &v
	return s
}

func (s *PatchEventResponseBodyAttendees) SetSelf(v bool) *PatchEventResponseBodyAttendees {
	s.Self = &v
	return s
}

type PatchEventResponseBodyEnd struct {
	Date     *string `json:"date,omitempty" xml:"date,omitempty"`
	DateTime *string `json:"dateTime,omitempty" xml:"dateTime,omitempty"`
	TimeZone *string `json:"timeZone,omitempty" xml:"timeZone,omitempty"`
}

func (s PatchEventResponseBodyEnd) String() string {
	return tea.Prettify(s)
}

func (s PatchEventResponseBodyEnd) GoString() string {
	return s.String()
}

func (s *PatchEventResponseBodyEnd) SetDate(v string) *PatchEventResponseBodyEnd {
	s.Date = &v
	return s
}

func (s *PatchEventResponseBodyEnd) SetDateTime(v string) *PatchEventResponseBodyEnd {
	s.DateTime = &v
	return s
}

func (s *PatchEventResponseBodyEnd) SetTimeZone(v string) *PatchEventResponseBodyEnd {
	s.TimeZone = &v
	return s
}

type PatchEventResponseBodyLocation struct {
	DisplayName  *string   `json:"displayName,omitempty" xml:"displayName,omitempty"`
	MeetingRooms []*string `json:"meetingRooms,omitempty" xml:"meetingRooms,omitempty" type:"Repeated"`
}

func (s PatchEventResponseBodyLocation) String() string {
	return tea.Prettify(s)
}

func (s PatchEventResponseBodyLocation) GoString() string {
	return s.String()
}

func (s *PatchEventResponseBodyLocation) SetDisplayName(v string) *PatchEventResponseBodyLocation {
	s.DisplayName = &v
	return s
}

func (s *PatchEventResponseBodyLocation) SetMeetingRooms(v []*string) *PatchEventResponseBodyLocation {
	s.MeetingRooms = v
	return s
}

type PatchEventResponseBodyOrganizer struct {
	// 用户名
	DisplayName *string `json:"displayName,omitempty" xml:"displayName,omitempty"`
	Id          *string `json:"id,omitempty" xml:"id,omitempty"`
	// 回复状态
	ResponseStatus *string `json:"responseStatus,omitempty" xml:"responseStatus,omitempty"`
	// 是否是当前登陆用户
	Self *bool `json:"self,omitempty" xml:"self,omitempty"`
}

func (s PatchEventResponseBodyOrganizer) String() string {
	return tea.Prettify(s)
}

func (s PatchEventResponseBodyOrganizer) GoString() string {
	return s.String()
}

func (s *PatchEventResponseBodyOrganizer) SetDisplayName(v string) *PatchEventResponseBodyOrganizer {
	s.DisplayName = &v
	return s
}

func (s *PatchEventResponseBodyOrganizer) SetId(v string) *PatchEventResponseBodyOrganizer {
	s.Id = &v
	return s
}

func (s *PatchEventResponseBodyOrganizer) SetResponseStatus(v string) *PatchEventResponseBodyOrganizer {
	s.ResponseStatus = &v
	return s
}

func (s *PatchEventResponseBodyOrganizer) SetSelf(v bool) *PatchEventResponseBodyOrganizer {
	s.Self = &v
	return s
}

type PatchEventResponseBodyRecurrence struct {
	Pattern *PatchEventResponseBodyRecurrencePattern `json:"pattern,omitempty" xml:"pattern,omitempty" type:"Struct"`
	Range   *PatchEventResponseBodyRecurrenceRange   `json:"range,omitempty" xml:"range,omitempty" type:"Struct"`
}

func (s PatchEventResponseBodyRecurrence) String() string {
	return tea.Prettify(s)
}

func (s PatchEventResponseBodyRecurrence) GoString() string {
	return s.String()
}

func (s *PatchEventResponseBodyRecurrence) SetPattern(v *PatchEventResponseBodyRecurrencePattern) *PatchEventResponseBodyRecurrence {
	s.Pattern = v
	return s
}

func (s *PatchEventResponseBodyRecurrence) SetRange(v *PatchEventResponseBodyRecurrenceRange) *PatchEventResponseBodyRecurrence {
	s.Range = v
	return s
}

type PatchEventResponseBodyRecurrencePattern struct {
	DayOfMonth *int32  `json:"dayOfMonth,omitempty" xml:"dayOfMonth,omitempty"`
	DaysOfWeek *string `json:"daysOfWeek,omitempty" xml:"daysOfWeek,omitempty"`
	Index      *string `json:"index,omitempty" xml:"index,omitempty"`
	Interval   *int32  `json:"interval,omitempty" xml:"interval,omitempty"`
	Type       *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s PatchEventResponseBodyRecurrencePattern) String() string {
	return tea.Prettify(s)
}

func (s PatchEventResponseBodyRecurrencePattern) GoString() string {
	return s.String()
}

func (s *PatchEventResponseBodyRecurrencePattern) SetDayOfMonth(v int32) *PatchEventResponseBodyRecurrencePattern {
	s.DayOfMonth = &v
	return s
}

func (s *PatchEventResponseBodyRecurrencePattern) SetDaysOfWeek(v string) *PatchEventResponseBodyRecurrencePattern {
	s.DaysOfWeek = &v
	return s
}

func (s *PatchEventResponseBodyRecurrencePattern) SetIndex(v string) *PatchEventResponseBodyRecurrencePattern {
	s.Index = &v
	return s
}

func (s *PatchEventResponseBodyRecurrencePattern) SetInterval(v int32) *PatchEventResponseBodyRecurrencePattern {
	s.Interval = &v
	return s
}

func (s *PatchEventResponseBodyRecurrencePattern) SetType(v string) *PatchEventResponseBodyRecurrencePattern {
	s.Type = &v
	return s
}

type PatchEventResponseBodyRecurrenceRange struct {
	EndDate             *string `json:"endDate,omitempty" xml:"endDate,omitempty"`
	NumberOfOccurrences *int32  `json:"numberOfOccurrences,omitempty" xml:"numberOfOccurrences,omitempty"`
	Type                *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s PatchEventResponseBodyRecurrenceRange) String() string {
	return tea.Prettify(s)
}

func (s PatchEventResponseBodyRecurrenceRange) GoString() string {
	return s.String()
}

func (s *PatchEventResponseBodyRecurrenceRange) SetEndDate(v string) *PatchEventResponseBodyRecurrenceRange {
	s.EndDate = &v
	return s
}

func (s *PatchEventResponseBodyRecurrenceRange) SetNumberOfOccurrences(v int32) *PatchEventResponseBodyRecurrenceRange {
	s.NumberOfOccurrences = &v
	return s
}

func (s *PatchEventResponseBodyRecurrenceRange) SetType(v string) *PatchEventResponseBodyRecurrenceRange {
	s.Type = &v
	return s
}

type PatchEventResponseBodyReminders struct {
	Method  *string `json:"method,omitempty" xml:"method,omitempty"`
	Minutes *string `json:"minutes,omitempty" xml:"minutes,omitempty"`
}

func (s PatchEventResponseBodyReminders) String() string {
	return tea.Prettify(s)
}

func (s PatchEventResponseBodyReminders) GoString() string {
	return s.String()
}

func (s *PatchEventResponseBodyReminders) SetMethod(v string) *PatchEventResponseBodyReminders {
	s.Method = &v
	return s
}

func (s *PatchEventResponseBodyReminders) SetMinutes(v string) *PatchEventResponseBodyReminders {
	s.Minutes = &v
	return s
}

type PatchEventResponseBodyStart struct {
	Date     *string `json:"date,omitempty" xml:"date,omitempty"`
	DateTime *string `json:"dateTime,omitempty" xml:"dateTime,omitempty"`
	TimeZone *string `json:"timeZone,omitempty" xml:"timeZone,omitempty"`
}

func (s PatchEventResponseBodyStart) String() string {
	return tea.Prettify(s)
}

func (s PatchEventResponseBodyStart) GoString() string {
	return s.String()
}

func (s *PatchEventResponseBodyStart) SetDate(v string) *PatchEventResponseBodyStart {
	s.Date = &v
	return s
}

func (s *PatchEventResponseBodyStart) SetDateTime(v string) *PatchEventResponseBodyStart {
	s.DateTime = &v
	return s
}

func (s *PatchEventResponseBodyStart) SetTimeZone(v string) *PatchEventResponseBodyStart {
	s.TimeZone = &v
	return s
}

type PatchEventResponse struct {
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *PatchEventResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s PatchEventResponse) String() string {
	return tea.Prettify(s)
}

func (s PatchEventResponse) GoString() string {
	return s.String()
}

func (s *PatchEventResponse) SetHeaders(v map[string]*string) *PatchEventResponse {
	s.Headers = v
	return s
}

func (s *PatchEventResponse) SetBody(v *PatchEventResponseBody) *PatchEventResponse {
	s.Body = v
	return s
}

type RemoveAttendeeHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s RemoveAttendeeHeaders) String() string {
	return tea.Prettify(s)
}

func (s RemoveAttendeeHeaders) GoString() string {
	return s.String()
}

func (s *RemoveAttendeeHeaders) SetCommonHeaders(v map[string]*string) *RemoveAttendeeHeaders {
	s.CommonHeaders = v
	return s
}

func (s *RemoveAttendeeHeaders) SetXAcsDingtalkAccessToken(v string) *RemoveAttendeeHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type RemoveAttendeeRequest struct {
	AttendeesToRemove []*RemoveAttendeeRequestAttendeesToRemove `json:"attendeesToRemove,omitempty" xml:"attendeesToRemove,omitempty" type:"Repeated"`
}

func (s RemoveAttendeeRequest) String() string {
	return tea.Prettify(s)
}

func (s RemoveAttendeeRequest) GoString() string {
	return s.String()
}

func (s *RemoveAttendeeRequest) SetAttendeesToRemove(v []*RemoveAttendeeRequestAttendeesToRemove) *RemoveAttendeeRequest {
	s.AttendeesToRemove = v
	return s
}

type RemoveAttendeeRequestAttendeesToRemove struct {
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
}

func (s RemoveAttendeeRequestAttendeesToRemove) String() string {
	return tea.Prettify(s)
}

func (s RemoveAttendeeRequestAttendeesToRemove) GoString() string {
	return s.String()
}

func (s *RemoveAttendeeRequestAttendeesToRemove) SetId(v string) *RemoveAttendeeRequestAttendeesToRemove {
	s.Id = &v
	return s
}

type RemoveAttendeeResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
}

func (s RemoveAttendeeResponse) String() string {
	return tea.Prettify(s)
}

func (s RemoveAttendeeResponse) GoString() string {
	return s.String()
}

func (s *RemoveAttendeeResponse) SetHeaders(v map[string]*string) *RemoveAttendeeResponse {
	s.Headers = v
	return s
}

type RespondEventHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s RespondEventHeaders) String() string {
	return tea.Prettify(s)
}

func (s RespondEventHeaders) GoString() string {
	return s.String()
}

func (s *RespondEventHeaders) SetCommonHeaders(v map[string]*string) *RespondEventHeaders {
	s.CommonHeaders = v
	return s
}

func (s *RespondEventHeaders) SetXAcsDingtalkAccessToken(v string) *RespondEventHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type RespondEventRequest struct {
	ResponseStatus *string `json:"responseStatus,omitempty" xml:"responseStatus,omitempty"`
}

func (s RespondEventRequest) String() string {
	return tea.Prettify(s)
}

func (s RespondEventRequest) GoString() string {
	return s.String()
}

func (s *RespondEventRequest) SetResponseStatus(v string) *RespondEventRequest {
	s.ResponseStatus = &v
	return s
}

type RespondEventResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
}

func (s RespondEventResponse) String() string {
	return tea.Prettify(s)
}

func (s RespondEventResponse) GoString() string {
	return s.String()
}

func (s *RespondEventResponse) SetHeaders(v map[string]*string) *RespondEventResponse {
	s.Headers = v
	return s
}

type SignInHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s SignInHeaders) String() string {
	return tea.Prettify(s)
}

func (s SignInHeaders) GoString() string {
	return s.String()
}

func (s *SignInHeaders) SetCommonHeaders(v map[string]*string) *SignInHeaders {
	s.CommonHeaders = v
	return s
}

func (s *SignInHeaders) SetXAcsDingtalkAccessToken(v string) *SignInHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type SignInResponseBody struct {
	// 签到时间戳
	CheckInTime *int64 `json:"checkInTime,omitempty" xml:"checkInTime,omitempty"`
}

func (s SignInResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SignInResponseBody) GoString() string {
	return s.String()
}

func (s *SignInResponseBody) SetCheckInTime(v int64) *SignInResponseBody {
	s.CheckInTime = &v
	return s
}

type SignInResponse struct {
	Headers map[string]*string  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SignInResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SignInResponse) String() string {
	return tea.Prettify(s)
}

func (s SignInResponse) GoString() string {
	return s.String()
}

func (s *SignInResponse) SetHeaders(v map[string]*string) *SignInResponse {
	s.Headers = v
	return s
}

func (s *SignInResponse) SetBody(v *SignInResponseBody) *SignInResponse {
	s.Body = v
	return s
}

type SignOutHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s SignOutHeaders) String() string {
	return tea.Prettify(s)
}

func (s SignOutHeaders) GoString() string {
	return s.String()
}

func (s *SignOutHeaders) SetCommonHeaders(v map[string]*string) *SignOutHeaders {
	s.CommonHeaders = v
	return s
}

func (s *SignOutHeaders) SetXAcsDingtalkAccessToken(v string) *SignOutHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type SignOutResponseBody struct {
	// 签退时间戳
	CheckOutTime *int64 `json:"checkOutTime,omitempty" xml:"checkOutTime,omitempty"`
}

func (s SignOutResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SignOutResponseBody) GoString() string {
	return s.String()
}

func (s *SignOutResponseBody) SetCheckOutTime(v int64) *SignOutResponseBody {
	s.CheckOutTime = &v
	return s
}

type SignOutResponse struct {
	Headers map[string]*string   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SignOutResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SignOutResponse) String() string {
	return tea.Prettify(s)
}

func (s SignOutResponse) GoString() string {
	return s.String()
}

func (s *SignOutResponse) SetHeaders(v map[string]*string) *SignOutResponse {
	s.Headers = v
	return s
}

func (s *SignOutResponse) SetBody(v *SignOutResponseBody) *SignOutResponse {
	s.Body = v
	return s
}

type SubscribeCalendarHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s SubscribeCalendarHeaders) String() string {
	return tea.Prettify(s)
}

func (s SubscribeCalendarHeaders) GoString() string {
	return s.String()
}

func (s *SubscribeCalendarHeaders) SetCommonHeaders(v map[string]*string) *SubscribeCalendarHeaders {
	s.CommonHeaders = v
	return s
}

func (s *SubscribeCalendarHeaders) SetXAcsDingtalkAccessToken(v string) *SubscribeCalendarHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type SubscribeCalendarResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
}

func (s SubscribeCalendarResponse) String() string {
	return tea.Prettify(s)
}

func (s SubscribeCalendarResponse) GoString() string {
	return s.String()
}

func (s *SubscribeCalendarResponse) SetHeaders(v map[string]*string) *SubscribeCalendarResponse {
	s.Headers = v
	return s
}

type UpdateSubscribedCalendarsHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s UpdateSubscribedCalendarsHeaders) String() string {
	return tea.Prettify(s)
}

func (s UpdateSubscribedCalendarsHeaders) GoString() string {
	return s.String()
}

func (s *UpdateSubscribedCalendarsHeaders) SetCommonHeaders(v map[string]*string) *UpdateSubscribedCalendarsHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdateSubscribedCalendarsHeaders) SetXAcsDingtalkAccessToken(v string) *UpdateSubscribedCalendarsHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type UpdateSubscribedCalendarsRequest struct {
	// 日历介绍
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// 日历管理员列表
	Managers []*string `json:"managers,omitempty" xml:"managers,omitempty" type:"Repeated"`
	// 日历名
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 可订阅列表
	SubscribeScope *UpdateSubscribedCalendarsRequestSubscribeScope `json:"subscribeScope,omitempty" xml:"subscribeScope,omitempty" type:"Struct"`
}

func (s UpdateSubscribedCalendarsRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateSubscribedCalendarsRequest) GoString() string {
	return s.String()
}

func (s *UpdateSubscribedCalendarsRequest) SetDescription(v string) *UpdateSubscribedCalendarsRequest {
	s.Description = &v
	return s
}

func (s *UpdateSubscribedCalendarsRequest) SetManagers(v []*string) *UpdateSubscribedCalendarsRequest {
	s.Managers = v
	return s
}

func (s *UpdateSubscribedCalendarsRequest) SetName(v string) *UpdateSubscribedCalendarsRequest {
	s.Name = &v
	return s
}

func (s *UpdateSubscribedCalendarsRequest) SetSubscribeScope(v *UpdateSubscribedCalendarsRequestSubscribeScope) *UpdateSubscribedCalendarsRequest {
	s.SubscribeScope = v
	return s
}

type UpdateSubscribedCalendarsRequestSubscribeScope struct {
	// 可订阅组织列表
	CorpIds []*string `json:"corpIds,omitempty" xml:"corpIds,omitempty" type:"Repeated"`
	// 可订阅群组列表
	OpenConversationIds []*string `json:"openConversationIds,omitempty" xml:"openConversationIds,omitempty" type:"Repeated"`
	// 可订阅人员列表
	UnionIds []*string `json:"unionIds,omitempty" xml:"unionIds,omitempty" type:"Repeated"`
}

func (s UpdateSubscribedCalendarsRequestSubscribeScope) String() string {
	return tea.Prettify(s)
}

func (s UpdateSubscribedCalendarsRequestSubscribeScope) GoString() string {
	return s.String()
}

func (s *UpdateSubscribedCalendarsRequestSubscribeScope) SetCorpIds(v []*string) *UpdateSubscribedCalendarsRequestSubscribeScope {
	s.CorpIds = v
	return s
}

func (s *UpdateSubscribedCalendarsRequestSubscribeScope) SetOpenConversationIds(v []*string) *UpdateSubscribedCalendarsRequestSubscribeScope {
	s.OpenConversationIds = v
	return s
}

func (s *UpdateSubscribedCalendarsRequestSubscribeScope) SetUnionIds(v []*string) *UpdateSubscribedCalendarsRequestSubscribeScope {
	s.UnionIds = v
	return s
}

type UpdateSubscribedCalendarsResponseBody struct {
	Result *bool `json:"result,omitempty" xml:"result,omitempty"`
}

func (s UpdateSubscribedCalendarsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateSubscribedCalendarsResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateSubscribedCalendarsResponseBody) SetResult(v bool) *UpdateSubscribedCalendarsResponseBody {
	s.Result = &v
	return s
}

type UpdateSubscribedCalendarsResponse struct {
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateSubscribedCalendarsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateSubscribedCalendarsResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateSubscribedCalendarsResponse) GoString() string {
	return s.String()
}

func (s *UpdateSubscribedCalendarsResponse) SetHeaders(v map[string]*string) *UpdateSubscribedCalendarsResponse {
	s.Headers = v
	return s
}

func (s *UpdateSubscribedCalendarsResponse) SetBody(v *UpdateSubscribedCalendarsResponseBody) *UpdateSubscribedCalendarsResponse {
	s.Body = v
	return s
}

type Client struct {
	openapi.Client
}

func NewClient(config *openapi.Config) (*Client, error) {
	client := new(Client)
	err := client.Init(config)
	return client, err
}

func (client *Client) Init(config *openapi.Config) (_err error) {
	_err = client.Client.Init(config)
	if _err != nil {
		return _err
	}
	client.EndpointRule = tea.String("")
	if tea.BoolValue(util.Empty(client.Endpoint)) {
		client.Endpoint = tea.String("api.dingtalk.com")
	}

	return nil
}

func (client *Client) AddAttendee(userId *string, calendarId *string, eventId *string, request *AddAttendeeRequest) (_result *AddAttendeeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &AddAttendeeHeaders{}
	_result = &AddAttendeeResponse{}
	_body, _err := client.AddAttendeeWithOptions(userId, calendarId, eventId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AddAttendeeWithOptions(userId *string, calendarId *string, eventId *string, request *AddAttendeeRequest, headers *AddAttendeeHeaders, runtime *util.RuntimeOptions) (_result *AddAttendeeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	userId = openapiutil.GetEncodeParam(userId)
	calendarId = openapiutil.GetEncodeParam(calendarId)
	eventId = openapiutil.GetEncodeParam(eventId)
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AttendeesToAdd)) {
		body["attendeesToAdd"] = request.AttendeesToAdd
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsDingtalkAccessToken)) {
		realHeaders["x-acs-dingtalk-access-token"] = util.ToJSONString(headers.XAcsDingtalkAccessToken)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	_result = &AddAttendeeResponse{}
	_body, _err := client.DoROARequest(tea.String("AddAttendee"), tea.String("calendar_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/calendar/users/"+tea.StringValue(userId)+"/calendars/"+tea.StringValue(calendarId)+"/events/"+tea.StringValue(eventId)+"/attendees"), tea.String("none"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CheckIn(userId *string, calendarId *string, eventId *string) (_result *CheckInResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CheckInHeaders{}
	_result = &CheckInResponse{}
	_body, _err := client.CheckInWithOptions(userId, calendarId, eventId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CheckInWithOptions(userId *string, calendarId *string, eventId *string, headers *CheckInHeaders, runtime *util.RuntimeOptions) (_result *CheckInResponse, _err error) {
	userId = openapiutil.GetEncodeParam(userId)
	calendarId = openapiutil.GetEncodeParam(calendarId)
	eventId = openapiutil.GetEncodeParam(eventId)
	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsDingtalkAccessToken)) {
		realHeaders["x-acs-dingtalk-access-token"] = util.ToJSONString(headers.XAcsDingtalkAccessToken)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
	}
	_result = &CheckInResponse{}
	_body, _err := client.DoROARequest(tea.String("CheckIn"), tea.String("calendar_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/calendar/users/"+tea.StringValue(userId)+"/calendars/"+tea.StringValue(calendarId)+"/events/"+tea.StringValue(eventId)+"/checkIn"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ConvertLegacyEventId(userId *string, request *ConvertLegacyEventIdRequest) (_result *ConvertLegacyEventIdResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ConvertLegacyEventIdHeaders{}
	_result = &ConvertLegacyEventIdResponse{}
	_body, _err := client.ConvertLegacyEventIdWithOptions(userId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ConvertLegacyEventIdWithOptions(userId *string, request *ConvertLegacyEventIdRequest, headers *ConvertLegacyEventIdHeaders, runtime *util.RuntimeOptions) (_result *ConvertLegacyEventIdResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	userId = openapiutil.GetEncodeParam(userId)
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.LegacyEventIds)) {
		body["legacyEventIds"] = request.LegacyEventIds
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsDingtalkAccessToken)) {
		realHeaders["x-acs-dingtalk-access-token"] = util.ToJSONString(headers.XAcsDingtalkAccessToken)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	_result = &ConvertLegacyEventIdResponse{}
	_body, _err := client.DoROARequest(tea.String("ConvertLegacyEventId"), tea.String("calendar_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/calendar/users/"+tea.StringValue(userId)+"/legacyEventIds/convert"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateAcls(userId *string, calendarId *string, request *CreateAclsRequest) (_result *CreateAclsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CreateAclsHeaders{}
	_result = &CreateAclsResponse{}
	_body, _err := client.CreateAclsWithOptions(userId, calendarId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateAclsWithOptions(userId *string, calendarId *string, request *CreateAclsRequest, headers *CreateAclsHeaders, runtime *util.RuntimeOptions) (_result *CreateAclsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	userId = openapiutil.GetEncodeParam(userId)
	calendarId = openapiutil.GetEncodeParam(calendarId)
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Privilege)) {
		body["privilege"] = request.Privilege
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(request.Scope))) {
		body["scope"] = request.Scope
	}

	if !tea.BoolValue(util.IsUnset(request.SendMsg)) {
		body["sendMsg"] = request.SendMsg
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsDingtalkAccessToken)) {
		realHeaders["x-acs-dingtalk-access-token"] = util.ToJSONString(headers.XAcsDingtalkAccessToken)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	_result = &CreateAclsResponse{}
	_body, _err := client.DoROARequest(tea.String("CreateAcls"), tea.String("calendar_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/calendar/users/"+tea.StringValue(userId)+"/calendars/"+tea.StringValue(calendarId)+"/acls"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateEvent(userId *string, calendarId *string, request *CreateEventRequest) (_result *CreateEventResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CreateEventHeaders{}
	_result = &CreateEventResponse{}
	_body, _err := client.CreateEventWithOptions(userId, calendarId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateEventWithOptions(userId *string, calendarId *string, request *CreateEventRequest, headers *CreateEventHeaders, runtime *util.RuntimeOptions) (_result *CreateEventResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	userId = openapiutil.GetEncodeParam(userId)
	calendarId = openapiutil.GetEncodeParam(calendarId)
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Attendees)) {
		body["attendees"] = request.Attendees
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(request.End))) {
		body["end"] = request.End
	}

	if !tea.BoolValue(util.IsUnset(request.Extra)) {
		body["extra"] = request.Extra
	}

	if !tea.BoolValue(util.IsUnset(request.IsAllDay)) {
		body["isAllDay"] = request.IsAllDay
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(request.Location))) {
		body["location"] = request.Location
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(request.OnlineMeetingInfo))) {
		body["onlineMeetingInfo"] = request.OnlineMeetingInfo
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(request.Recurrence))) {
		body["recurrence"] = request.Recurrence
	}

	if !tea.BoolValue(util.IsUnset(request.Reminders)) {
		body["reminders"] = request.Reminders
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(request.Start))) {
		body["start"] = request.Start
	}

	if !tea.BoolValue(util.IsUnset(request.Summary)) {
		body["summary"] = request.Summary
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsDingtalkAccessToken)) {
		realHeaders["x-acs-dingtalk-access-token"] = util.ToJSONString(headers.XAcsDingtalkAccessToken)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	_result = &CreateEventResponse{}
	_body, _err := client.DoROARequest(tea.String("CreateEvent"), tea.String("calendar_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/calendar/users/"+tea.StringValue(userId)+"/calendars/"+tea.StringValue(calendarId)+"/events"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateSubscribedCalendar(userId *string, request *CreateSubscribedCalendarRequest) (_result *CreateSubscribedCalendarResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CreateSubscribedCalendarHeaders{}
	_result = &CreateSubscribedCalendarResponse{}
	_body, _err := client.CreateSubscribedCalendarWithOptions(userId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateSubscribedCalendarWithOptions(userId *string, request *CreateSubscribedCalendarRequest, headers *CreateSubscribedCalendarHeaders, runtime *util.RuntimeOptions) (_result *CreateSubscribedCalendarResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	userId = openapiutil.GetEncodeParam(userId)
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.Managers)) {
		body["managers"] = request.Managers
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(request.SubscribeScope))) {
		body["subscribeScope"] = request.SubscribeScope
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsDingtalkAccessToken)) {
		realHeaders["x-acs-dingtalk-access-token"] = util.ToJSONString(headers.XAcsDingtalkAccessToken)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	_result = &CreateSubscribedCalendarResponse{}
	_body, _err := client.DoROARequest(tea.String("CreateSubscribedCalendar"), tea.String("calendar_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/calendar/users/"+tea.StringValue(userId)+"/subscribedCalendars"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteAcl(userId *string, calendarId *string, aclId *string) (_result *DeleteAclResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &DeleteAclHeaders{}
	_result = &DeleteAclResponse{}
	_body, _err := client.DeleteAclWithOptions(userId, calendarId, aclId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteAclWithOptions(userId *string, calendarId *string, aclId *string, headers *DeleteAclHeaders, runtime *util.RuntimeOptions) (_result *DeleteAclResponse, _err error) {
	userId = openapiutil.GetEncodeParam(userId)
	calendarId = openapiutil.GetEncodeParam(calendarId)
	aclId = openapiutil.GetEncodeParam(aclId)
	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsDingtalkAccessToken)) {
		realHeaders["x-acs-dingtalk-access-token"] = util.ToJSONString(headers.XAcsDingtalkAccessToken)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
	}
	_result = &DeleteAclResponse{}
	_body, _err := client.DoROARequest(tea.String("DeleteAcl"), tea.String("calendar_1.0"), tea.String("HTTP"), tea.String("DELETE"), tea.String("AK"), tea.String("/v1.0/calendar/users/"+tea.StringValue(userId)+"/calendars/"+tea.StringValue(calendarId)+"/acls/"+tea.StringValue(aclId)), tea.String("none"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteEvent(userId *string, calendarId *string, eventId *string) (_result *DeleteEventResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &DeleteEventHeaders{}
	_result = &DeleteEventResponse{}
	_body, _err := client.DeleteEventWithOptions(userId, calendarId, eventId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteEventWithOptions(userId *string, calendarId *string, eventId *string, headers *DeleteEventHeaders, runtime *util.RuntimeOptions) (_result *DeleteEventResponse, _err error) {
	userId = openapiutil.GetEncodeParam(userId)
	calendarId = openapiutil.GetEncodeParam(calendarId)
	eventId = openapiutil.GetEncodeParam(eventId)
	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsDingtalkAccessToken)) {
		realHeaders["x-acs-dingtalk-access-token"] = util.ToJSONString(headers.XAcsDingtalkAccessToken)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
	}
	_result = &DeleteEventResponse{}
	_body, _err := client.DoROARequest(tea.String("DeleteEvent"), tea.String("calendar_1.0"), tea.String("HTTP"), tea.String("DELETE"), tea.String("AK"), tea.String("/v1.0/calendar/users/"+tea.StringValue(userId)+"/calendars/"+tea.StringValue(calendarId)+"/events/"+tea.StringValue(eventId)), tea.String("none"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteSubscribedCalendar(userId *string, calendarId *string) (_result *DeleteSubscribedCalendarResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &DeleteSubscribedCalendarHeaders{}
	_result = &DeleteSubscribedCalendarResponse{}
	_body, _err := client.DeleteSubscribedCalendarWithOptions(userId, calendarId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteSubscribedCalendarWithOptions(userId *string, calendarId *string, headers *DeleteSubscribedCalendarHeaders, runtime *util.RuntimeOptions) (_result *DeleteSubscribedCalendarResponse, _err error) {
	userId = openapiutil.GetEncodeParam(userId)
	calendarId = openapiutil.GetEncodeParam(calendarId)
	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsDingtalkAccessToken)) {
		realHeaders["x-acs-dingtalk-access-token"] = util.ToJSONString(headers.XAcsDingtalkAccessToken)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
	}
	_result = &DeleteSubscribedCalendarResponse{}
	_body, _err := client.DoROARequest(tea.String("DeleteSubscribedCalendar"), tea.String("calendar_1.0"), tea.String("HTTP"), tea.String("DELETE"), tea.String("AK"), tea.String("/v1.0/calendar/users/"+tea.StringValue(userId)+"/subscribedCalendars/"+tea.StringValue(calendarId)), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GenerateCaldavAccount(userId *string, request *GenerateCaldavAccountRequest) (_result *GenerateCaldavAccountResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GenerateCaldavAccountHeaders{}
	_result = &GenerateCaldavAccountResponse{}
	_body, _err := client.GenerateCaldavAccountWithOptions(userId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GenerateCaldavAccountWithOptions(userId *string, request *GenerateCaldavAccountRequest, headers *GenerateCaldavAccountHeaders, runtime *util.RuntimeOptions) (_result *GenerateCaldavAccountResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	userId = openapiutil.GetEncodeParam(userId)
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Device)) {
		body["device"] = request.Device
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.DingUid)) {
		realHeaders["dingUid"] = util.ToJSONString(headers.DingUid)
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsDingtalkAccessToken)) {
		realHeaders["x-acs-dingtalk-access-token"] = util.ToJSONString(headers.XAcsDingtalkAccessToken)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	_result = &GenerateCaldavAccountResponse{}
	_body, _err := client.DoROARequest(tea.String("GenerateCaldavAccount"), tea.String("calendar_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/calendar/users/"+tea.StringValue(userId)+"/caldavAccounts"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetEvent(userId *string, calendarId *string, eventId *string, request *GetEventRequest) (_result *GetEventResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetEventHeaders{}
	_result = &GetEventResponse{}
	_body, _err := client.GetEventWithOptions(userId, calendarId, eventId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetEventWithOptions(userId *string, calendarId *string, eventId *string, request *GetEventRequest, headers *GetEventHeaders, runtime *util.RuntimeOptions) (_result *GetEventResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	userId = openapiutil.GetEncodeParam(userId)
	calendarId = openapiutil.GetEncodeParam(calendarId)
	eventId = openapiutil.GetEncodeParam(eventId)
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MaxAttendees)) {
		query["maxAttendees"] = request.MaxAttendees
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsDingtalkAccessToken)) {
		realHeaders["x-acs-dingtalk-access-token"] = util.ToJSONString(headers.XAcsDingtalkAccessToken)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	_result = &GetEventResponse{}
	_body, _err := client.DoROARequest(tea.String("GetEvent"), tea.String("calendar_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/calendar/users/"+tea.StringValue(userId)+"/calendars/"+tea.StringValue(calendarId)+"/events/"+tea.StringValue(eventId)), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetSchedule(userId *string, request *GetScheduleRequest) (_result *GetScheduleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetScheduleHeaders{}
	_result = &GetScheduleResponse{}
	_body, _err := client.GetScheduleWithOptions(userId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetScheduleWithOptions(userId *string, request *GetScheduleRequest, headers *GetScheduleHeaders, runtime *util.RuntimeOptions) (_result *GetScheduleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	userId = openapiutil.GetEncodeParam(userId)
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		body["endTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		body["startTime"] = request.StartTime
	}

	if !tea.BoolValue(util.IsUnset(request.UserIds)) {
		body["userIds"] = request.UserIds
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsDingtalkAccessToken)) {
		realHeaders["x-acs-dingtalk-access-token"] = util.ToJSONString(headers.XAcsDingtalkAccessToken)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	_result = &GetScheduleResponse{}
	_body, _err := client.DoROARequest(tea.String("GetSchedule"), tea.String("calendar_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/calendar/users/"+tea.StringValue(userId)+"/querySchedule"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetSignInList(userId *string, calendarId *string, eventId *string, request *GetSignInListRequest) (_result *GetSignInListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetSignInListHeaders{}
	_result = &GetSignInListResponse{}
	_body, _err := client.GetSignInListWithOptions(userId, calendarId, eventId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetSignInListWithOptions(userId *string, calendarId *string, eventId *string, request *GetSignInListRequest, headers *GetSignInListHeaders, runtime *util.RuntimeOptions) (_result *GetSignInListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	userId = openapiutil.GetEncodeParam(userId)
	calendarId = openapiutil.GetEncodeParam(calendarId)
	eventId = openapiutil.GetEncodeParam(eventId)
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["maxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["nextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		query["type"] = request.Type
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsDingtalkAccessToken)) {
		realHeaders["x-acs-dingtalk-access-token"] = util.ToJSONString(headers.XAcsDingtalkAccessToken)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	_result = &GetSignInListResponse{}
	_body, _err := client.DoROARequest(tea.String("GetSignInList"), tea.String("calendar_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/calendar/users/"+tea.StringValue(userId)+"/calendars/"+tea.StringValue(calendarId)+"/events/"+tea.StringValue(eventId)+"/signin"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetSignOutList(userId *string, calendarId *string, eventId *string, request *GetSignOutListRequest) (_result *GetSignOutListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetSignOutListHeaders{}
	_result = &GetSignOutListResponse{}
	_body, _err := client.GetSignOutListWithOptions(userId, calendarId, eventId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetSignOutListWithOptions(userId *string, calendarId *string, eventId *string, request *GetSignOutListRequest, headers *GetSignOutListHeaders, runtime *util.RuntimeOptions) (_result *GetSignOutListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	userId = openapiutil.GetEncodeParam(userId)
	calendarId = openapiutil.GetEncodeParam(calendarId)
	eventId = openapiutil.GetEncodeParam(eventId)
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["maxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["nextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		query["type"] = request.Type
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsDingtalkAccessToken)) {
		realHeaders["x-acs-dingtalk-access-token"] = util.ToJSONString(headers.XAcsDingtalkAccessToken)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	_result = &GetSignOutListResponse{}
	_body, _err := client.DoROARequest(tea.String("GetSignOutList"), tea.String("calendar_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/calendar/users/"+tea.StringValue(userId)+"/calendars/"+tea.StringValue(calendarId)+"/events/"+tea.StringValue(eventId)+"/signOut"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetSubscribedCalendar(userId *string, calendarId *string) (_result *GetSubscribedCalendarResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetSubscribedCalendarHeaders{}
	_result = &GetSubscribedCalendarResponse{}
	_body, _err := client.GetSubscribedCalendarWithOptions(userId, calendarId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetSubscribedCalendarWithOptions(userId *string, calendarId *string, headers *GetSubscribedCalendarHeaders, runtime *util.RuntimeOptions) (_result *GetSubscribedCalendarResponse, _err error) {
	userId = openapiutil.GetEncodeParam(userId)
	calendarId = openapiutil.GetEncodeParam(calendarId)
	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsDingtalkAccessToken)) {
		realHeaders["x-acs-dingtalk-access-token"] = util.ToJSONString(headers.XAcsDingtalkAccessToken)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
	}
	_result = &GetSubscribedCalendarResponse{}
	_body, _err := client.DoROARequest(tea.String("GetSubscribedCalendar"), tea.String("calendar_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/calendar/users/"+tea.StringValue(userId)+"/subscribedCalendars/"+tea.StringValue(calendarId)), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListAcls(userId *string, calendarId *string) (_result *ListAclsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ListAclsHeaders{}
	_result = &ListAclsResponse{}
	_body, _err := client.ListAclsWithOptions(userId, calendarId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListAclsWithOptions(userId *string, calendarId *string, headers *ListAclsHeaders, runtime *util.RuntimeOptions) (_result *ListAclsResponse, _err error) {
	userId = openapiutil.GetEncodeParam(userId)
	calendarId = openapiutil.GetEncodeParam(calendarId)
	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsDingtalkAccessToken)) {
		realHeaders["x-acs-dingtalk-access-token"] = util.ToJSONString(headers.XAcsDingtalkAccessToken)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
	}
	_result = &ListAclsResponse{}
	_body, _err := client.DoROARequest(tea.String("ListAcls"), tea.String("calendar_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/calendar/users/"+tea.StringValue(userId)+"/calendars/"+tea.StringValue(calendarId)+"/acls"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListAttendees(userId *string, calendarId *string, eventId *string, request *ListAttendeesRequest) (_result *ListAttendeesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ListAttendeesHeaders{}
	_result = &ListAttendeesResponse{}
	_body, _err := client.ListAttendeesWithOptions(userId, calendarId, eventId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListAttendeesWithOptions(userId *string, calendarId *string, eventId *string, request *ListAttendeesRequest, headers *ListAttendeesHeaders, runtime *util.RuntimeOptions) (_result *ListAttendeesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	userId = openapiutil.GetEncodeParam(userId)
	calendarId = openapiutil.GetEncodeParam(calendarId)
	eventId = openapiutil.GetEncodeParam(eventId)
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["maxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["nextToken"] = request.NextToken
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsDingtalkAccessToken)) {
		realHeaders["x-acs-dingtalk-access-token"] = util.ToJSONString(headers.XAcsDingtalkAccessToken)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	_result = &ListAttendeesResponse{}
	_body, _err := client.DoROARequest(tea.String("ListAttendees"), tea.String("calendar_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/calendar/users/"+tea.StringValue(userId)+"/calendars/"+tea.StringValue(calendarId)+"/events/"+tea.StringValue(eventId)+"/attendees"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListCalendars(userId *string) (_result *ListCalendarsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ListCalendarsHeaders{}
	_result = &ListCalendarsResponse{}
	_body, _err := client.ListCalendarsWithOptions(userId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListCalendarsWithOptions(userId *string, headers *ListCalendarsHeaders, runtime *util.RuntimeOptions) (_result *ListCalendarsResponse, _err error) {
	userId = openapiutil.GetEncodeParam(userId)
	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsDingtalkAccessToken)) {
		realHeaders["x-acs-dingtalk-access-token"] = util.ToJSONString(headers.XAcsDingtalkAccessToken)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
	}
	_result = &ListCalendarsResponse{}
	_body, _err := client.DoROARequest(tea.String("ListCalendars"), tea.String("calendar_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/calendar/users/"+tea.StringValue(userId)+"/calendars"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListEvents(userId *string, calendarId *string, request *ListEventsRequest) (_result *ListEventsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ListEventsHeaders{}
	_result = &ListEventsResponse{}
	_body, _err := client.ListEventsWithOptions(userId, calendarId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListEventsWithOptions(userId *string, calendarId *string, request *ListEventsRequest, headers *ListEventsHeaders, runtime *util.RuntimeOptions) (_result *ListEventsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	userId = openapiutil.GetEncodeParam(userId)
	calendarId = openapiutil.GetEncodeParam(calendarId)
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MaxAttendees)) {
		query["maxAttendees"] = request.MaxAttendees
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["maxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["nextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.ShowDeleted)) {
		query["showDeleted"] = request.ShowDeleted
	}

	if !tea.BoolValue(util.IsUnset(request.SyncToken)) {
		query["syncToken"] = request.SyncToken
	}

	if !tea.BoolValue(util.IsUnset(request.TimeMax)) {
		query["timeMax"] = request.TimeMax
	}

	if !tea.BoolValue(util.IsUnset(request.TimeMin)) {
		query["timeMin"] = request.TimeMin
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsDingtalkAccessToken)) {
		realHeaders["x-acs-dingtalk-access-token"] = util.ToJSONString(headers.XAcsDingtalkAccessToken)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	_result = &ListEventsResponse{}
	_body, _err := client.DoROARequest(tea.String("ListEvents"), tea.String("calendar_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/calendar/users/"+tea.StringValue(userId)+"/calendars/"+tea.StringValue(calendarId)+"/events"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListEventsInstances(userId *string, calendarId *string, request *ListEventsInstancesRequest) (_result *ListEventsInstancesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ListEventsInstancesHeaders{}
	_result = &ListEventsInstancesResponse{}
	_body, _err := client.ListEventsInstancesWithOptions(userId, calendarId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListEventsInstancesWithOptions(userId *string, calendarId *string, request *ListEventsInstancesRequest, headers *ListEventsInstancesHeaders, runtime *util.RuntimeOptions) (_result *ListEventsInstancesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	userId = openapiutil.GetEncodeParam(userId)
	calendarId = openapiutil.GetEncodeParam(calendarId)
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MaxAttendees)) {
		query["maxAttendees"] = request.MaxAttendees
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["maxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.SeriesMasterId)) {
		query["seriesMasterId"] = request.SeriesMasterId
	}

	if !tea.BoolValue(util.IsUnset(request.StartRecurrenceId)) {
		query["startRecurrenceId"] = request.StartRecurrenceId
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsDingtalkAccessToken)) {
		realHeaders["x-acs-dingtalk-access-token"] = util.ToJSONString(headers.XAcsDingtalkAccessToken)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	_result = &ListEventsInstancesResponse{}
	_body, _err := client.DoROARequest(tea.String("ListEventsInstances"), tea.String("calendar_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/calendar/users/"+tea.StringValue(userId)+"/calendars/"+tea.StringValue(calendarId)+"/instances"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListEventsView(userId *string, calendarId *string, request *ListEventsViewRequest) (_result *ListEventsViewResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ListEventsViewHeaders{}
	_result = &ListEventsViewResponse{}
	_body, _err := client.ListEventsViewWithOptions(userId, calendarId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListEventsViewWithOptions(userId *string, calendarId *string, request *ListEventsViewRequest, headers *ListEventsViewHeaders, runtime *util.RuntimeOptions) (_result *ListEventsViewResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	userId = openapiutil.GetEncodeParam(userId)
	calendarId = openapiutil.GetEncodeParam(calendarId)
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MaxAttendees)) {
		query["maxAttendees"] = request.MaxAttendees
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["maxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["nextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.TimeMax)) {
		query["timeMax"] = request.TimeMax
	}

	if !tea.BoolValue(util.IsUnset(request.TimeMin)) {
		query["timeMin"] = request.TimeMin
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsDingtalkAccessToken)) {
		realHeaders["x-acs-dingtalk-access-token"] = util.ToJSONString(headers.XAcsDingtalkAccessToken)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	_result = &ListEventsViewResponse{}
	_body, _err := client.DoROARequest(tea.String("ListEventsView"), tea.String("calendar_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/calendar/users/"+tea.StringValue(userId)+"/calendars/"+tea.StringValue(calendarId)+"/eventsview"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) PatchEvent(userId *string, calendarId *string, eventId *string, request *PatchEventRequest) (_result *PatchEventResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &PatchEventHeaders{}
	_result = &PatchEventResponse{}
	_body, _err := client.PatchEventWithOptions(userId, calendarId, eventId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) PatchEventWithOptions(userId *string, calendarId *string, eventId *string, request *PatchEventRequest, headers *PatchEventHeaders, runtime *util.RuntimeOptions) (_result *PatchEventResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	userId = openapiutil.GetEncodeParam(userId)
	calendarId = openapiutil.GetEncodeParam(calendarId)
	eventId = openapiutil.GetEncodeParam(eventId)
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Attendees)) {
		body["attendees"] = request.Attendees
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(request.End))) {
		body["end"] = request.End
	}

	if !tea.BoolValue(util.IsUnset(request.Extra)) {
		body["extra"] = request.Extra
	}

	if !tea.BoolValue(util.IsUnset(request.Id)) {
		body["id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.IsAllDay)) {
		body["isAllDay"] = request.IsAllDay
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(request.Location))) {
		body["location"] = request.Location
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(request.Recurrence))) {
		body["recurrence"] = request.Recurrence
	}

	if !tea.BoolValue(util.IsUnset(request.Reminders)) {
		body["reminders"] = request.Reminders
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(request.Start))) {
		body["start"] = request.Start
	}

	if !tea.BoolValue(util.IsUnset(request.Summary)) {
		body["summary"] = request.Summary
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsDingtalkAccessToken)) {
		realHeaders["x-acs-dingtalk-access-token"] = util.ToJSONString(headers.XAcsDingtalkAccessToken)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	_result = &PatchEventResponse{}
	_body, _err := client.DoROARequest(tea.String("PatchEvent"), tea.String("calendar_1.0"), tea.String("HTTP"), tea.String("PUT"), tea.String("AK"), tea.String("/v1.0/calendar/users/"+tea.StringValue(userId)+"/calendars/"+tea.StringValue(calendarId)+"/events/"+tea.StringValue(eventId)), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RemoveAttendee(userId *string, calendarId *string, eventId *string, request *RemoveAttendeeRequest) (_result *RemoveAttendeeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &RemoveAttendeeHeaders{}
	_result = &RemoveAttendeeResponse{}
	_body, _err := client.RemoveAttendeeWithOptions(userId, calendarId, eventId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RemoveAttendeeWithOptions(userId *string, calendarId *string, eventId *string, request *RemoveAttendeeRequest, headers *RemoveAttendeeHeaders, runtime *util.RuntimeOptions) (_result *RemoveAttendeeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	userId = openapiutil.GetEncodeParam(userId)
	calendarId = openapiutil.GetEncodeParam(calendarId)
	eventId = openapiutil.GetEncodeParam(eventId)
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AttendeesToRemove)) {
		body["attendeesToRemove"] = request.AttendeesToRemove
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsDingtalkAccessToken)) {
		realHeaders["x-acs-dingtalk-access-token"] = util.ToJSONString(headers.XAcsDingtalkAccessToken)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	_result = &RemoveAttendeeResponse{}
	_body, _err := client.DoROARequest(tea.String("RemoveAttendee"), tea.String("calendar_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/calendar/users/"+tea.StringValue(userId)+"/calendars/"+tea.StringValue(calendarId)+"/events/"+tea.StringValue(eventId)+"/attendees/batchRemove"), tea.String("none"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RespondEvent(userId *string, calendarId *string, eventId *string, request *RespondEventRequest) (_result *RespondEventResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &RespondEventHeaders{}
	_result = &RespondEventResponse{}
	_body, _err := client.RespondEventWithOptions(userId, calendarId, eventId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RespondEventWithOptions(userId *string, calendarId *string, eventId *string, request *RespondEventRequest, headers *RespondEventHeaders, runtime *util.RuntimeOptions) (_result *RespondEventResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	userId = openapiutil.GetEncodeParam(userId)
	calendarId = openapiutil.GetEncodeParam(calendarId)
	eventId = openapiutil.GetEncodeParam(eventId)
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ResponseStatus)) {
		body["responseStatus"] = request.ResponseStatus
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsDingtalkAccessToken)) {
		realHeaders["x-acs-dingtalk-access-token"] = util.ToJSONString(headers.XAcsDingtalkAccessToken)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	_result = &RespondEventResponse{}
	_body, _err := client.DoROARequest(tea.String("RespondEvent"), tea.String("calendar_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/calendar/users/"+tea.StringValue(userId)+"/calendars/"+tea.StringValue(calendarId)+"/events/"+tea.StringValue(eventId)+"/respond"), tea.String("none"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SignIn(userId *string, calendarId *string, eventId *string) (_result *SignInResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &SignInHeaders{}
	_result = &SignInResponse{}
	_body, _err := client.SignInWithOptions(userId, calendarId, eventId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SignInWithOptions(userId *string, calendarId *string, eventId *string, headers *SignInHeaders, runtime *util.RuntimeOptions) (_result *SignInResponse, _err error) {
	userId = openapiutil.GetEncodeParam(userId)
	calendarId = openapiutil.GetEncodeParam(calendarId)
	eventId = openapiutil.GetEncodeParam(eventId)
	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsDingtalkAccessToken)) {
		realHeaders["x-acs-dingtalk-access-token"] = util.ToJSONString(headers.XAcsDingtalkAccessToken)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
	}
	_result = &SignInResponse{}
	_body, _err := client.DoROARequest(tea.String("SignIn"), tea.String("calendar_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/calendar/users/"+tea.StringValue(userId)+"/calendars/"+tea.StringValue(calendarId)+"/events/"+tea.StringValue(eventId)+"/signin"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SignOut(userId *string, calendarId *string, eventId *string) (_result *SignOutResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &SignOutHeaders{}
	_result = &SignOutResponse{}
	_body, _err := client.SignOutWithOptions(userId, calendarId, eventId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SignOutWithOptions(userId *string, calendarId *string, eventId *string, headers *SignOutHeaders, runtime *util.RuntimeOptions) (_result *SignOutResponse, _err error) {
	userId = openapiutil.GetEncodeParam(userId)
	calendarId = openapiutil.GetEncodeParam(calendarId)
	eventId = openapiutil.GetEncodeParam(eventId)
	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsDingtalkAccessToken)) {
		realHeaders["x-acs-dingtalk-access-token"] = util.ToJSONString(headers.XAcsDingtalkAccessToken)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
	}
	_result = &SignOutResponse{}
	_body, _err := client.DoROARequest(tea.String("SignOut"), tea.String("calendar_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/calendar/users/"+tea.StringValue(userId)+"/calendars/"+tea.StringValue(calendarId)+"/events/"+tea.StringValue(eventId)+"/signOut"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SubscribeCalendar(userId *string, calendarId *string) (_result *SubscribeCalendarResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &SubscribeCalendarHeaders{}
	_result = &SubscribeCalendarResponse{}
	_body, _err := client.SubscribeCalendarWithOptions(userId, calendarId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SubscribeCalendarWithOptions(userId *string, calendarId *string, headers *SubscribeCalendarHeaders, runtime *util.RuntimeOptions) (_result *SubscribeCalendarResponse, _err error) {
	userId = openapiutil.GetEncodeParam(userId)
	calendarId = openapiutil.GetEncodeParam(calendarId)
	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsDingtalkAccessToken)) {
		realHeaders["x-acs-dingtalk-access-token"] = util.ToJSONString(headers.XAcsDingtalkAccessToken)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
	}
	_result = &SubscribeCalendarResponse{}
	_body, _err := client.DoROARequest(tea.String("SubscribeCalendar"), tea.String("calendar_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/calendar/users/"+tea.StringValue(userId)+"/calendars/"+tea.StringValue(calendarId)+"/subscribe"), tea.String("none"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateSubscribedCalendars(calendarId *string, userId *string, request *UpdateSubscribedCalendarsRequest) (_result *UpdateSubscribedCalendarsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &UpdateSubscribedCalendarsHeaders{}
	_result = &UpdateSubscribedCalendarsResponse{}
	_body, _err := client.UpdateSubscribedCalendarsWithOptions(calendarId, userId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateSubscribedCalendarsWithOptions(calendarId *string, userId *string, request *UpdateSubscribedCalendarsRequest, headers *UpdateSubscribedCalendarsHeaders, runtime *util.RuntimeOptions) (_result *UpdateSubscribedCalendarsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	calendarId = openapiutil.GetEncodeParam(calendarId)
	userId = openapiutil.GetEncodeParam(userId)
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.Managers)) {
		body["managers"] = request.Managers
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(request.SubscribeScope))) {
		body["subscribeScope"] = request.SubscribeScope
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsDingtalkAccessToken)) {
		realHeaders["x-acs-dingtalk-access-token"] = util.ToJSONString(headers.XAcsDingtalkAccessToken)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	_result = &UpdateSubscribedCalendarsResponse{}
	_body, _err := client.DoROARequest(tea.String("UpdateSubscribedCalendars"), tea.String("calendar_1.0"), tea.String("HTTP"), tea.String("PUT"), tea.String("AK"), tea.String("/v1.0/calendar/users/"+tea.StringValue(userId)+"/subscribedCalendars/"+tea.StringValue(calendarId)), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}
