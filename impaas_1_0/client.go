// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package impaas_1_0

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
)

type AddGroupMembersHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	OperationSource         *string            `json:"operationSource,omitempty" xml:"operationSource,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s AddGroupMembersHeaders) String() string {
	return tea.Prettify(s)
}

func (s AddGroupMembersHeaders) GoString() string {
	return s.String()
}

func (s *AddGroupMembersHeaders) SetCommonHeaders(v map[string]*string) *AddGroupMembersHeaders {
	s.CommonHeaders = v
	return s
}

func (s *AddGroupMembersHeaders) SetOperationSource(v string) *AddGroupMembersHeaders {
	s.OperationSource = &v
	return s
}

func (s *AddGroupMembersHeaders) SetXAcsDingtalkAccessToken(v string) *AddGroupMembersHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type AddGroupMembersRequest struct {
	ConversationId *string                          `json:"conversationId,omitempty" xml:"conversationId,omitempty"`
	Members        []*AddGroupMembersRequestMembers `json:"members,omitempty" xml:"members,omitempty" type:"Repeated"`
	OperatorUid    *string                          `json:"operatorUid,omitempty" xml:"operatorUid,omitempty"`
}

func (s AddGroupMembersRequest) String() string {
	return tea.Prettify(s)
}

func (s AddGroupMembersRequest) GoString() string {
	return s.String()
}

func (s *AddGroupMembersRequest) SetConversationId(v string) *AddGroupMembersRequest {
	s.ConversationId = &v
	return s
}

func (s *AddGroupMembersRequest) SetMembers(v []*AddGroupMembersRequestMembers) *AddGroupMembersRequest {
	s.Members = v
	return s
}

func (s *AddGroupMembersRequest) SetOperatorUid(v string) *AddGroupMembersRequest {
	s.OperatorUid = &v
	return s
}

type AddGroupMembersRequestMembers struct {
	Nick *string `json:"nick,omitempty" xml:"nick,omitempty"`
	Uid  *string `json:"uid,omitempty" xml:"uid,omitempty"`
}

func (s AddGroupMembersRequestMembers) String() string {
	return tea.Prettify(s)
}

func (s AddGroupMembersRequestMembers) GoString() string {
	return s.String()
}

func (s *AddGroupMembersRequestMembers) SetNick(v string) *AddGroupMembersRequestMembers {
	s.Nick = &v
	return s
}

func (s *AddGroupMembersRequestMembers) SetUid(v string) *AddGroupMembersRequestMembers {
	s.Uid = &v
	return s
}

type AddGroupMembersResponseBody struct {
	MemberUids []*string `json:"memberUids,omitempty" xml:"memberUids,omitempty" type:"Repeated"`
}

func (s AddGroupMembersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddGroupMembersResponseBody) GoString() string {
	return s.String()
}

func (s *AddGroupMembersResponseBody) SetMemberUids(v []*string) *AddGroupMembersResponseBody {
	s.MemberUids = v
	return s
}

type AddGroupMembersResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AddGroupMembersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AddGroupMembersResponse) String() string {
	return tea.Prettify(s)
}

func (s AddGroupMembersResponse) GoString() string {
	return s.String()
}

func (s *AddGroupMembersResponse) SetHeaders(v map[string]*string) *AddGroupMembersResponse {
	s.Headers = v
	return s
}

func (s *AddGroupMembersResponse) SetBody(v *AddGroupMembersResponseBody) *AddGroupMembersResponse {
	s.Body = v
	return s
}

type AddProfileHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s AddProfileHeaders) String() string {
	return tea.Prettify(s)
}

func (s AddProfileHeaders) GoString() string {
	return s.String()
}

func (s *AddProfileHeaders) SetCommonHeaders(v map[string]*string) *AddProfileHeaders {
	s.CommonHeaders = v
	return s
}

func (s *AddProfileHeaders) SetXAcsDingtalkAccessToken(v string) *AddProfileHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type AddProfileRequest struct {
	// 外部app的账号，格式：xxx@channel格式
	AppUid *string `json:"appUid,omitempty" xml:"appUid,omitempty"`
	// 头像mediaId，调用Upload接口获得
	AvatarMediaId *string `json:"avatarMediaId,omitempty" xml:"avatarMediaId,omitempty"`
	// 手机号
	MobileNumber *string `json:"mobileNumber,omitempty" xml:"mobileNumber,omitempty"`
	// 昵称
	Nick *string `json:"nick,omitempty" xml:"nick,omitempty"`
}

func (s AddProfileRequest) String() string {
	return tea.Prettify(s)
}

func (s AddProfileRequest) GoString() string {
	return s.String()
}

func (s *AddProfileRequest) SetAppUid(v string) *AddProfileRequest {
	s.AppUid = &v
	return s
}

func (s *AddProfileRequest) SetAvatarMediaId(v string) *AddProfileRequest {
	s.AvatarMediaId = &v
	return s
}

func (s *AddProfileRequest) SetMobileNumber(v string) *AddProfileRequest {
	s.MobileNumber = &v
	return s
}

func (s *AddProfileRequest) SetNick(v string) *AddProfileRequest {
	s.Nick = &v
	return s
}

type AddProfileResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
}

func (s AddProfileResponse) String() string {
	return tea.Prettify(s)
}

func (s AddProfileResponse) GoString() string {
	return s.String()
}

func (s *AddProfileResponse) SetHeaders(v map[string]*string) *AddProfileResponse {
	s.Headers = v
	return s
}

type BatchSendHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s BatchSendHeaders) String() string {
	return tea.Prettify(s)
}

func (s BatchSendHeaders) GoString() string {
	return s.String()
}

func (s *BatchSendHeaders) SetCommonHeaders(v map[string]*string) *BatchSendHeaders {
	s.CommonHeaders = v
	return s
}

func (s *BatchSendHeaders) SetXAcsDingtalkAccessToken(v string) *BatchSendHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type BatchSendRequest struct {
	// 接受者列表，外部用户
	AppUids []*string `json:"appUids,omitempty" xml:"appUids,omitempty" type:"Repeated"`
	// 消息内容
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// 接收消息的群聊列表
	ConversationIds []*string `json:"conversationIds,omitempty" xml:"conversationIds,omitempty" type:"Repeated"`
	// 发送者，企业员工账号
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s BatchSendRequest) String() string {
	return tea.Prettify(s)
}

func (s BatchSendRequest) GoString() string {
	return s.String()
}

func (s *BatchSendRequest) SetAppUids(v []*string) *BatchSendRequest {
	s.AppUids = v
	return s
}

func (s *BatchSendRequest) SetContent(v string) *BatchSendRequest {
	s.Content = &v
	return s
}

func (s *BatchSendRequest) SetConversationIds(v []*string) *BatchSendRequest {
	s.ConversationIds = v
	return s
}

func (s *BatchSendRequest) SetUserId(v string) *BatchSendRequest {
	s.UserId = &v
	return s
}

type BatchSendResponseBody struct {
	// 任务Id
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
}

func (s BatchSendResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BatchSendResponseBody) GoString() string {
	return s.String()
}

func (s *BatchSendResponseBody) SetTaskId(v string) *BatchSendResponseBody {
	s.TaskId = &v
	return s
}

type BatchSendResponse struct {
	Headers map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *BatchSendResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s BatchSendResponse) String() string {
	return tea.Prettify(s)
}

func (s BatchSendResponse) GoString() string {
	return s.String()
}

func (s *BatchSendResponse) SetHeaders(v map[string]*string) *BatchSendResponse {
	s.Headers = v
	return s
}

func (s *BatchSendResponse) SetBody(v *BatchSendResponseBody) *BatchSendResponse {
	s.Body = v
	return s
}

type CreateGroupHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	OperationSource         *string            `json:"operationSource,omitempty" xml:"operationSource,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CreateGroupHeaders) String() string {
	return tea.Prettify(s)
}

func (s CreateGroupHeaders) GoString() string {
	return s.String()
}

func (s *CreateGroupHeaders) SetCommonHeaders(v map[string]*string) *CreateGroupHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreateGroupHeaders) SetOperationSource(v string) *CreateGroupHeaders {
	s.OperationSource = &v
	return s
}

func (s *CreateGroupHeaders) SetXAcsDingtalkAccessToken(v string) *CreateGroupHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CreateGroupRequest struct {
	Channel     *string            `json:"channel,omitempty" xml:"channel,omitempty"`
	CreatorUid  *string            `json:"creatorUid,omitempty" xml:"creatorUid,omitempty"`
	IconMediaId *string            `json:"iconMediaId,omitempty" xml:"iconMediaId,omitempty"`
	Name        *string            `json:"name,omitempty" xml:"name,omitempty"`
	Properties  map[string]*string `json:"properties,omitempty" xml:"properties,omitempty"`
	Uuid        *string            `json:"uuid,omitempty" xml:"uuid,omitempty"`
}

func (s CreateGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateGroupRequest) GoString() string {
	return s.String()
}

func (s *CreateGroupRequest) SetChannel(v string) *CreateGroupRequest {
	s.Channel = &v
	return s
}

func (s *CreateGroupRequest) SetCreatorUid(v string) *CreateGroupRequest {
	s.CreatorUid = &v
	return s
}

func (s *CreateGroupRequest) SetIconMediaId(v string) *CreateGroupRequest {
	s.IconMediaId = &v
	return s
}

func (s *CreateGroupRequest) SetName(v string) *CreateGroupRequest {
	s.Name = &v
	return s
}

func (s *CreateGroupRequest) SetProperties(v map[string]*string) *CreateGroupRequest {
	s.Properties = v
	return s
}

func (s *CreateGroupRequest) SetUuid(v string) *CreateGroupRequest {
	s.Uuid = &v
	return s
}

type CreateGroupResponseBody struct {
	ChatId         *string `json:"chatId,omitempty" xml:"chatId,omitempty"`
	ConversationId *string `json:"conversationId,omitempty" xml:"conversationId,omitempty"`
	CreateTime     *int64  `json:"createTime,omitempty" xml:"createTime,omitempty"`
}

func (s CreateGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateGroupResponseBody) GoString() string {
	return s.String()
}

func (s *CreateGroupResponseBody) SetChatId(v string) *CreateGroupResponseBody {
	s.ChatId = &v
	return s
}

func (s *CreateGroupResponseBody) SetConversationId(v string) *CreateGroupResponseBody {
	s.ConversationId = &v
	return s
}

func (s *CreateGroupResponseBody) SetCreateTime(v int64) *CreateGroupResponseBody {
	s.CreateTime = &v
	return s
}

type CreateGroupResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateGroupResponse) GoString() string {
	return s.String()
}

func (s *CreateGroupResponse) SetHeaders(v map[string]*string) *CreateGroupResponse {
	s.Headers = v
	return s
}

func (s *CreateGroupResponse) SetBody(v *CreateGroupResponseBody) *CreateGroupResponse {
	s.Body = v
	return s
}

type CreateTrustGroupHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	OperationSource         *string            `json:"operationSource,omitempty" xml:"operationSource,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CreateTrustGroupHeaders) String() string {
	return tea.Prettify(s)
}

func (s CreateTrustGroupHeaders) GoString() string {
	return s.String()
}

func (s *CreateTrustGroupHeaders) SetCommonHeaders(v map[string]*string) *CreateTrustGroupHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreateTrustGroupHeaders) SetOperationSource(v string) *CreateTrustGroupHeaders {
	s.OperationSource = &v
	return s
}

func (s *CreateTrustGroupHeaders) SetXAcsDingtalkAccessToken(v string) *CreateTrustGroupHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CreateTrustGroupRequest struct {
	// MPASS渠道编码
	Channel *string `json:"channel,omitempty" xml:"channel,omitempty"`
	// 素材ID
	IconMediaId *string `json:"iconMediaId,omitempty" xml:"iconMediaId,omitempty"`
	// 群名称
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 其他扩展参数
	Properties map[string]*string `json:"properties,omitempty" xml:"properties,omitempty"`
	// 外部系统映射唯一标识
	Uuid *string `json:"uuid,omitempty" xml:"uuid,omitempty"`
}

func (s CreateTrustGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateTrustGroupRequest) GoString() string {
	return s.String()
}

func (s *CreateTrustGroupRequest) SetChannel(v string) *CreateTrustGroupRequest {
	s.Channel = &v
	return s
}

func (s *CreateTrustGroupRequest) SetIconMediaId(v string) *CreateTrustGroupRequest {
	s.IconMediaId = &v
	return s
}

func (s *CreateTrustGroupRequest) SetName(v string) *CreateTrustGroupRequest {
	s.Name = &v
	return s
}

func (s *CreateTrustGroupRequest) SetProperties(v map[string]*string) *CreateTrustGroupRequest {
	s.Properties = v
	return s
}

func (s *CreateTrustGroupRequest) SetUuid(v string) *CreateTrustGroupRequest {
	s.Uuid = &v
	return s
}

type CreateTrustGroupResponseBody struct {
	ChatId             *string `json:"chatId,omitempty" xml:"chatId,omitempty"`
	CreateTime         *int64  `json:"createTime,omitempty" xml:"createTime,omitempty"`
	OpenConversationId *string `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
}

func (s CreateTrustGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateTrustGroupResponseBody) GoString() string {
	return s.String()
}

func (s *CreateTrustGroupResponseBody) SetChatId(v string) *CreateTrustGroupResponseBody {
	s.ChatId = &v
	return s
}

func (s *CreateTrustGroupResponseBody) SetCreateTime(v int64) *CreateTrustGroupResponseBody {
	s.CreateTime = &v
	return s
}

func (s *CreateTrustGroupResponseBody) SetOpenConversationId(v string) *CreateTrustGroupResponseBody {
	s.OpenConversationId = &v
	return s
}

type CreateTrustGroupResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateTrustGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateTrustGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateTrustGroupResponse) GoString() string {
	return s.String()
}

func (s *CreateTrustGroupResponse) SetHeaders(v map[string]*string) *CreateTrustGroupResponse {
	s.Headers = v
	return s
}

func (s *CreateTrustGroupResponse) SetBody(v *CreateTrustGroupResponseBody) *CreateTrustGroupResponse {
	s.Body = v
	return s
}

type DismissGroupHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	OperationSource         *string            `json:"operationSource,omitempty" xml:"operationSource,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s DismissGroupHeaders) String() string {
	return tea.Prettify(s)
}

func (s DismissGroupHeaders) GoString() string {
	return s.String()
}

func (s *DismissGroupHeaders) SetCommonHeaders(v map[string]*string) *DismissGroupHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DismissGroupHeaders) SetOperationSource(v string) *DismissGroupHeaders {
	s.OperationSource = &v
	return s
}

func (s *DismissGroupHeaders) SetXAcsDingtalkAccessToken(v string) *DismissGroupHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type DismissGroupRequest struct {
	ConversationId *string `json:"conversationId,omitempty" xml:"conversationId,omitempty"`
	OperatorUid    *string `json:"operatorUid,omitempty" xml:"operatorUid,omitempty"`
}

func (s DismissGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s DismissGroupRequest) GoString() string {
	return s.String()
}

func (s *DismissGroupRequest) SetConversationId(v string) *DismissGroupRequest {
	s.ConversationId = &v
	return s
}

func (s *DismissGroupRequest) SetOperatorUid(v string) *DismissGroupRequest {
	s.OperatorUid = &v
	return s
}

type DismissGroupResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
}

func (s DismissGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s DismissGroupResponse) GoString() string {
	return s.String()
}

func (s *DismissGroupResponse) SetHeaders(v map[string]*string) *DismissGroupResponse {
	s.Headers = v
	return s
}

type GetConversationIdHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetConversationIdHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetConversationIdHeaders) GoString() string {
	return s.String()
}

func (s *GetConversationIdHeaders) SetCommonHeaders(v map[string]*string) *GetConversationIdHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetConversationIdHeaders) SetXAcsDingtalkAccessToken(v string) *GetConversationIdHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetConversationIdRequest struct {
	// 外部用户账号：outerId@channel
	AppUid *string `json:"appUid,omitempty" xml:"appUid,omitempty"`
	// 员工企业账号：staffId#corpId@dingding
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s GetConversationIdRequest) String() string {
	return tea.Prettify(s)
}

func (s GetConversationIdRequest) GoString() string {
	return s.String()
}

func (s *GetConversationIdRequest) SetAppUid(v string) *GetConversationIdRequest {
	s.AppUid = &v
	return s
}

func (s *GetConversationIdRequest) SetUserId(v string) *GetConversationIdRequest {
	s.UserId = &v
	return s
}

type GetConversationIdResponseBody struct {
	ConversationId *string `json:"conversationId,omitempty" xml:"conversationId,omitempty"`
}

func (s GetConversationIdResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetConversationIdResponseBody) GoString() string {
	return s.String()
}

func (s *GetConversationIdResponseBody) SetConversationId(v string) *GetConversationIdResponseBody {
	s.ConversationId = &v
	return s
}

type GetConversationIdResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetConversationIdResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetConversationIdResponse) String() string {
	return tea.Prettify(s)
}

func (s GetConversationIdResponse) GoString() string {
	return s.String()
}

func (s *GetConversationIdResponse) SetHeaders(v map[string]*string) *GetConversationIdResponse {
	s.Headers = v
	return s
}

func (s *GetConversationIdResponse) SetBody(v *GetConversationIdResponseBody) *GetConversationIdResponse {
	s.Body = v
	return s
}

type GetMediaUrlHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetMediaUrlHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetMediaUrlHeaders) GoString() string {
	return s.String()
}

func (s *GetMediaUrlHeaders) SetCommonHeaders(v map[string]*string) *GetMediaUrlHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetMediaUrlHeaders) SetXAcsDingtalkAccessToken(v string) *GetMediaUrlHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetMediaUrlRequest struct {
	// 多媒体id
	MediaId *string `json:"mediaId,omitempty" xml:"mediaId,omitempty"`
	// 过期时间
	UrlExpireTime *int32 `json:"urlExpireTime,omitempty" xml:"urlExpireTime,omitempty"`
}

func (s GetMediaUrlRequest) String() string {
	return tea.Prettify(s)
}

func (s GetMediaUrlRequest) GoString() string {
	return s.String()
}

func (s *GetMediaUrlRequest) SetMediaId(v string) *GetMediaUrlRequest {
	s.MediaId = &v
	return s
}

func (s *GetMediaUrlRequest) SetUrlExpireTime(v int32) *GetMediaUrlRequest {
	s.UrlExpireTime = &v
	return s
}

type GetMediaUrlResponseBody struct {
	Url *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s GetMediaUrlResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetMediaUrlResponseBody) GoString() string {
	return s.String()
}

func (s *GetMediaUrlResponseBody) SetUrl(v string) *GetMediaUrlResponseBody {
	s.Url = &v
	return s
}

type GetMediaUrlResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetMediaUrlResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetMediaUrlResponse) String() string {
	return tea.Prettify(s)
}

func (s GetMediaUrlResponse) GoString() string {
	return s.String()
}

func (s *GetMediaUrlResponse) SetHeaders(v map[string]*string) *GetMediaUrlResponse {
	s.Headers = v
	return s
}

func (s *GetMediaUrlResponse) SetBody(v *GetMediaUrlResponseBody) *GetMediaUrlResponse {
	s.Body = v
	return s
}

type ListGroupStaffMembersHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s ListGroupStaffMembersHeaders) String() string {
	return tea.Prettify(s)
}

func (s ListGroupStaffMembersHeaders) GoString() string {
	return s.String()
}

func (s *ListGroupStaffMembersHeaders) SetCommonHeaders(v map[string]*string) *ListGroupStaffMembersHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListGroupStaffMembersHeaders) SetXAcsDingtalkAccessToken(v string) *ListGroupStaffMembersHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type ListGroupStaffMembersRequest struct {
	ConversationId *string `json:"conversationId,omitempty" xml:"conversationId,omitempty"`
}

func (s ListGroupStaffMembersRequest) String() string {
	return tea.Prettify(s)
}

func (s ListGroupStaffMembersRequest) GoString() string {
	return s.String()
}

func (s *ListGroupStaffMembersRequest) SetConversationId(v string) *ListGroupStaffMembersRequest {
	s.ConversationId = &v
	return s
}

type ListGroupStaffMembersResponseBody struct {
	Members []*ListGroupStaffMembersResponseBodyMembers `json:"members,omitempty" xml:"members,omitempty" type:"Repeated"`
}

func (s ListGroupStaffMembersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListGroupStaffMembersResponseBody) GoString() string {
	return s.String()
}

func (s *ListGroupStaffMembersResponseBody) SetMembers(v []*ListGroupStaffMembersResponseBodyMembers) *ListGroupStaffMembersResponseBody {
	s.Members = v
	return s
}

type ListGroupStaffMembersResponseBodyMembers struct {
	Nick *string `json:"nick,omitempty" xml:"nick,omitempty"`
	Uid  *string `json:"uid,omitempty" xml:"uid,omitempty"`
}

func (s ListGroupStaffMembersResponseBodyMembers) String() string {
	return tea.Prettify(s)
}

func (s ListGroupStaffMembersResponseBodyMembers) GoString() string {
	return s.String()
}

func (s *ListGroupStaffMembersResponseBodyMembers) SetNick(v string) *ListGroupStaffMembersResponseBodyMembers {
	s.Nick = &v
	return s
}

func (s *ListGroupStaffMembersResponseBodyMembers) SetUid(v string) *ListGroupStaffMembersResponseBodyMembers {
	s.Uid = &v
	return s
}

type ListGroupStaffMembersResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListGroupStaffMembersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListGroupStaffMembersResponse) String() string {
	return tea.Prettify(s)
}

func (s ListGroupStaffMembersResponse) GoString() string {
	return s.String()
}

func (s *ListGroupStaffMembersResponse) SetHeaders(v map[string]*string) *ListGroupStaffMembersResponse {
	s.Headers = v
	return s
}

func (s *ListGroupStaffMembersResponse) SetBody(v *ListGroupStaffMembersResponseBody) *ListGroupStaffMembersResponse {
	s.Body = v
	return s
}

type QueryBatchSendResultHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryBatchSendResultHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryBatchSendResultHeaders) GoString() string {
	return s.String()
}

func (s *QueryBatchSendResultHeaders) SetCommonHeaders(v map[string]*string) *QueryBatchSendResultHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryBatchSendResultHeaders) SetXAcsDingtalkAccessToken(v string) *QueryBatchSendResultHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryBatchSendResultRequest struct {
	// 发送者，必须是B端用户
	SenderUserId *string `json:"senderUserId,omitempty" xml:"senderUserId,omitempty"`
	// batchSend返回的taskId
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
}

func (s QueryBatchSendResultRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryBatchSendResultRequest) GoString() string {
	return s.String()
}

func (s *QueryBatchSendResultRequest) SetSenderUserId(v string) *QueryBatchSendResultRequest {
	s.SenderUserId = &v
	return s
}

func (s *QueryBatchSendResultRequest) SetTaskId(v string) *QueryBatchSendResultRequest {
	s.TaskId = &v
	return s
}

type QueryBatchSendResultResponseBody struct {
	Results []*QueryBatchSendResultResponseBodyResults `json:"results,omitempty" xml:"results,omitempty" type:"Repeated"`
	// status
	Status *int32 `json:"status,omitempty" xml:"status,omitempty"`
}

func (s QueryBatchSendResultResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryBatchSendResultResponseBody) GoString() string {
	return s.String()
}

func (s *QueryBatchSendResultResponseBody) SetResults(v []*QueryBatchSendResultResponseBodyResults) *QueryBatchSendResultResponseBody {
	s.Results = v
	return s
}

func (s *QueryBatchSendResultResponseBody) SetStatus(v int32) *QueryBatchSendResultResponseBody {
	s.Status = &v
	return s
}

type QueryBatchSendResultResponseBodyResults struct {
	AppUid         *string `json:"appUid,omitempty" xml:"appUid,omitempty"`
	ConversationId *string `json:"conversationId,omitempty" xml:"conversationId,omitempty"`
	ErrorCode      *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	ErrorMessage   *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	MsgId          *string `json:"msgId,omitempty" xml:"msgId,omitempty"`
}

func (s QueryBatchSendResultResponseBodyResults) String() string {
	return tea.Prettify(s)
}

func (s QueryBatchSendResultResponseBodyResults) GoString() string {
	return s.String()
}

func (s *QueryBatchSendResultResponseBodyResults) SetAppUid(v string) *QueryBatchSendResultResponseBodyResults {
	s.AppUid = &v
	return s
}

func (s *QueryBatchSendResultResponseBodyResults) SetConversationId(v string) *QueryBatchSendResultResponseBodyResults {
	s.ConversationId = &v
	return s
}

func (s *QueryBatchSendResultResponseBodyResults) SetErrorCode(v string) *QueryBatchSendResultResponseBodyResults {
	s.ErrorCode = &v
	return s
}

func (s *QueryBatchSendResultResponseBodyResults) SetErrorMessage(v string) *QueryBatchSendResultResponseBodyResults {
	s.ErrorMessage = &v
	return s
}

func (s *QueryBatchSendResultResponseBodyResults) SetMsgId(v string) *QueryBatchSendResultResponseBodyResults {
	s.MsgId = &v
	return s
}

type QueryBatchSendResultResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryBatchSendResultResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryBatchSendResultResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryBatchSendResultResponse) GoString() string {
	return s.String()
}

func (s *QueryBatchSendResultResponse) SetHeaders(v map[string]*string) *QueryBatchSendResultResponse {
	s.Headers = v
	return s
}

func (s *QueryBatchSendResultResponse) SetBody(v *QueryBatchSendResultResponseBody) *QueryBatchSendResultResponse {
	s.Body = v
	return s
}

type ReadMessageHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	OperationSource         *string            `json:"operationSource,omitempty" xml:"operationSource,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s ReadMessageHeaders) String() string {
	return tea.Prettify(s)
}

func (s ReadMessageHeaders) GoString() string {
	return s.String()
}

func (s *ReadMessageHeaders) SetCommonHeaders(v map[string]*string) *ReadMessageHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ReadMessageHeaders) SetOperationSource(v string) *ReadMessageHeaders {
	s.OperationSource = &v
	return s
}

func (s *ReadMessageHeaders) SetXAcsDingtalkAccessToken(v string) *ReadMessageHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type ReadMessageRequest struct {
	MessageId   *string `json:"messageId,omitempty" xml:"messageId,omitempty"`
	OperatorUid *string `json:"operatorUid,omitempty" xml:"operatorUid,omitempty"`
}

func (s ReadMessageRequest) String() string {
	return tea.Prettify(s)
}

func (s ReadMessageRequest) GoString() string {
	return s.String()
}

func (s *ReadMessageRequest) SetMessageId(v string) *ReadMessageRequest {
	s.MessageId = &v
	return s
}

func (s *ReadMessageRequest) SetOperatorUid(v string) *ReadMessageRequest {
	s.OperatorUid = &v
	return s
}

type ReadMessageResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
}

func (s ReadMessageResponse) String() string {
	return tea.Prettify(s)
}

func (s ReadMessageResponse) GoString() string {
	return s.String()
}

func (s *ReadMessageResponse) SetHeaders(v map[string]*string) *ReadMessageResponse {
	s.Headers = v
	return s
}

type RecallMessageHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	OperationSource         *string            `json:"operationSource,omitempty" xml:"operationSource,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s RecallMessageHeaders) String() string {
	return tea.Prettify(s)
}

func (s RecallMessageHeaders) GoString() string {
	return s.String()
}

func (s *RecallMessageHeaders) SetCommonHeaders(v map[string]*string) *RecallMessageHeaders {
	s.CommonHeaders = v
	return s
}

func (s *RecallMessageHeaders) SetOperationSource(v string) *RecallMessageHeaders {
	s.OperationSource = &v
	return s
}

func (s *RecallMessageHeaders) SetXAcsDingtalkAccessToken(v string) *RecallMessageHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type RecallMessageRequest struct {
	MessageId   *string `json:"messageId,omitempty" xml:"messageId,omitempty"`
	OperatorUid *string `json:"operatorUid,omitempty" xml:"operatorUid,omitempty"`
	Type        *int32  `json:"type,omitempty" xml:"type,omitempty"`
}

func (s RecallMessageRequest) String() string {
	return tea.Prettify(s)
}

func (s RecallMessageRequest) GoString() string {
	return s.String()
}

func (s *RecallMessageRequest) SetMessageId(v string) *RecallMessageRequest {
	s.MessageId = &v
	return s
}

func (s *RecallMessageRequest) SetOperatorUid(v string) *RecallMessageRequest {
	s.OperatorUid = &v
	return s
}

func (s *RecallMessageRequest) SetType(v int32) *RecallMessageRequest {
	s.Type = &v
	return s
}

type RecallMessageResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
}

func (s RecallMessageResponse) String() string {
	return tea.Prettify(s)
}

func (s RecallMessageResponse) GoString() string {
	return s.String()
}

func (s *RecallMessageResponse) SetHeaders(v map[string]*string) *RecallMessageResponse {
	s.Headers = v
	return s
}

type RemoveGroupMembersHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	OperationSource         *string            `json:"operationSource,omitempty" xml:"operationSource,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s RemoveGroupMembersHeaders) String() string {
	return tea.Prettify(s)
}

func (s RemoveGroupMembersHeaders) GoString() string {
	return s.String()
}

func (s *RemoveGroupMembersHeaders) SetCommonHeaders(v map[string]*string) *RemoveGroupMembersHeaders {
	s.CommonHeaders = v
	return s
}

func (s *RemoveGroupMembersHeaders) SetOperationSource(v string) *RemoveGroupMembersHeaders {
	s.OperationSource = &v
	return s
}

func (s *RemoveGroupMembersHeaders) SetXAcsDingtalkAccessToken(v string) *RemoveGroupMembersHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type RemoveGroupMembersRequest struct {
	ConversationId *string   `json:"conversationId,omitempty" xml:"conversationId,omitempty"`
	MemberUids     []*string `json:"memberUids,omitempty" xml:"memberUids,omitempty" type:"Repeated"`
	OperatorUid    *string   `json:"operatorUid,omitempty" xml:"operatorUid,omitempty"`
}

func (s RemoveGroupMembersRequest) String() string {
	return tea.Prettify(s)
}

func (s RemoveGroupMembersRequest) GoString() string {
	return s.String()
}

func (s *RemoveGroupMembersRequest) SetConversationId(v string) *RemoveGroupMembersRequest {
	s.ConversationId = &v
	return s
}

func (s *RemoveGroupMembersRequest) SetMemberUids(v []*string) *RemoveGroupMembersRequest {
	s.MemberUids = v
	return s
}

func (s *RemoveGroupMembersRequest) SetOperatorUid(v string) *RemoveGroupMembersRequest {
	s.OperatorUid = &v
	return s
}

type RemoveGroupMembersResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
}

func (s RemoveGroupMembersResponse) String() string {
	return tea.Prettify(s)
}

func (s RemoveGroupMembersResponse) GoString() string {
	return s.String()
}

func (s *RemoveGroupMembersResponse) SetHeaders(v map[string]*string) *RemoveGroupMembersResponse {
	s.Headers = v
	return s
}

type SendMessageHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	OperationSource         *string            `json:"operationSource,omitempty" xml:"operationSource,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s SendMessageHeaders) String() string {
	return tea.Prettify(s)
}

func (s SendMessageHeaders) GoString() string {
	return s.String()
}

func (s *SendMessageHeaders) SetCommonHeaders(v map[string]*string) *SendMessageHeaders {
	s.CommonHeaders = v
	return s
}

func (s *SendMessageHeaders) SetOperationSource(v string) *SendMessageHeaders {
	s.OperationSource = &v
	return s
}

func (s *SendMessageHeaders) SetXAcsDingtalkAccessToken(v string) *SendMessageHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type SendMessageRequest struct {
	Content        *string `json:"content,omitempty" xml:"content,omitempty"`
	ConversationId *string `json:"conversationId,omitempty" xml:"conversationId,omitempty"`
	CreateTime     *int64  `json:"createTime,omitempty" xml:"createTime,omitempty"`
	ReceiverUid    *string `json:"receiverUid,omitempty" xml:"receiverUid,omitempty"`
	SenderUid      *string `json:"senderUid,omitempty" xml:"senderUid,omitempty"`
	Uuid           *string `json:"uuid,omitempty" xml:"uuid,omitempty"`
}

func (s SendMessageRequest) String() string {
	return tea.Prettify(s)
}

func (s SendMessageRequest) GoString() string {
	return s.String()
}

func (s *SendMessageRequest) SetContent(v string) *SendMessageRequest {
	s.Content = &v
	return s
}

func (s *SendMessageRequest) SetConversationId(v string) *SendMessageRequest {
	s.ConversationId = &v
	return s
}

func (s *SendMessageRequest) SetCreateTime(v int64) *SendMessageRequest {
	s.CreateTime = &v
	return s
}

func (s *SendMessageRequest) SetReceiverUid(v string) *SendMessageRequest {
	s.ReceiverUid = &v
	return s
}

func (s *SendMessageRequest) SetSenderUid(v string) *SendMessageRequest {
	s.SenderUid = &v
	return s
}

func (s *SendMessageRequest) SetUuid(v string) *SendMessageRequest {
	s.Uuid = &v
	return s
}

type SendMessageResponseBody struct {
	CreateTime *int64  `json:"createTime,omitempty" xml:"createTime,omitempty"`
	MessageId  *string `json:"messageId,omitempty" xml:"messageId,omitempty"`
	MsgId      *string `json:"msgId,omitempty" xml:"msgId,omitempty"`
}

func (s SendMessageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SendMessageResponseBody) GoString() string {
	return s.String()
}

func (s *SendMessageResponseBody) SetCreateTime(v int64) *SendMessageResponseBody {
	s.CreateTime = &v
	return s
}

func (s *SendMessageResponseBody) SetMessageId(v string) *SendMessageResponseBody {
	s.MessageId = &v
	return s
}

func (s *SendMessageResponseBody) SetMsgId(v string) *SendMessageResponseBody {
	s.MsgId = &v
	return s
}

type SendMessageResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SendMessageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SendMessageResponse) String() string {
	return tea.Prettify(s)
}

func (s SendMessageResponse) GoString() string {
	return s.String()
}

func (s *SendMessageResponse) SetHeaders(v map[string]*string) *SendMessageResponse {
	s.Headers = v
	return s
}

func (s *SendMessageResponse) SetBody(v *SendMessageResponseBody) *SendMessageResponse {
	s.Body = v
	return s
}

type UpdateGroupNameHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	OperationSource         *string            `json:"operationSource,omitempty" xml:"operationSource,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s UpdateGroupNameHeaders) String() string {
	return tea.Prettify(s)
}

func (s UpdateGroupNameHeaders) GoString() string {
	return s.String()
}

func (s *UpdateGroupNameHeaders) SetCommonHeaders(v map[string]*string) *UpdateGroupNameHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdateGroupNameHeaders) SetOperationSource(v string) *UpdateGroupNameHeaders {
	s.OperationSource = &v
	return s
}

func (s *UpdateGroupNameHeaders) SetXAcsDingtalkAccessToken(v string) *UpdateGroupNameHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type UpdateGroupNameRequest struct {
	ConversationId *string `json:"conversationId,omitempty" xml:"conversationId,omitempty"`
	Name           *string `json:"name,omitempty" xml:"name,omitempty"`
	OperatorUid    *string `json:"operatorUid,omitempty" xml:"operatorUid,omitempty"`
}

func (s UpdateGroupNameRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateGroupNameRequest) GoString() string {
	return s.String()
}

func (s *UpdateGroupNameRequest) SetConversationId(v string) *UpdateGroupNameRequest {
	s.ConversationId = &v
	return s
}

func (s *UpdateGroupNameRequest) SetName(v string) *UpdateGroupNameRequest {
	s.Name = &v
	return s
}

func (s *UpdateGroupNameRequest) SetOperatorUid(v string) *UpdateGroupNameRequest {
	s.OperatorUid = &v
	return s
}

type UpdateGroupNameResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
}

func (s UpdateGroupNameResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateGroupNameResponse) GoString() string {
	return s.String()
}

func (s *UpdateGroupNameResponse) SetHeaders(v map[string]*string) *UpdateGroupNameResponse {
	s.Headers = v
	return s
}

type UpdateGroupOwnerHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	OperationSource         *string            `json:"operationSource,omitempty" xml:"operationSource,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s UpdateGroupOwnerHeaders) String() string {
	return tea.Prettify(s)
}

func (s UpdateGroupOwnerHeaders) GoString() string {
	return s.String()
}

func (s *UpdateGroupOwnerHeaders) SetCommonHeaders(v map[string]*string) *UpdateGroupOwnerHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdateGroupOwnerHeaders) SetOperationSource(v string) *UpdateGroupOwnerHeaders {
	s.OperationSource = &v
	return s
}

func (s *UpdateGroupOwnerHeaders) SetXAcsDingtalkAccessToken(v string) *UpdateGroupOwnerHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type UpdateGroupOwnerRequest struct {
	ConversationId *string `json:"conversationId,omitempty" xml:"conversationId,omitempty"`
	OperatorUid    *string `json:"operatorUid,omitempty" xml:"operatorUid,omitempty"`
	OwnerUid       *string `json:"ownerUid,omitempty" xml:"ownerUid,omitempty"`
}

func (s UpdateGroupOwnerRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateGroupOwnerRequest) GoString() string {
	return s.String()
}

func (s *UpdateGroupOwnerRequest) SetConversationId(v string) *UpdateGroupOwnerRequest {
	s.ConversationId = &v
	return s
}

func (s *UpdateGroupOwnerRequest) SetOperatorUid(v string) *UpdateGroupOwnerRequest {
	s.OperatorUid = &v
	return s
}

func (s *UpdateGroupOwnerRequest) SetOwnerUid(v string) *UpdateGroupOwnerRequest {
	s.OwnerUid = &v
	return s
}

type UpdateGroupOwnerResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
}

func (s UpdateGroupOwnerResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateGroupOwnerResponse) GoString() string {
	return s.String()
}

func (s *UpdateGroupOwnerResponse) SetHeaders(v map[string]*string) *UpdateGroupOwnerResponse {
	s.Headers = v
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

func (client *Client) AddGroupMembers(request *AddGroupMembersRequest) (_result *AddGroupMembersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &AddGroupMembersHeaders{}
	_result = &AddGroupMembersResponse{}
	_body, _err := client.AddGroupMembersWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AddGroupMembersWithOptions(request *AddGroupMembersRequest, headers *AddGroupMembersHeaders, runtime *util.RuntimeOptions) (_result *AddGroupMembersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ConversationId)) {
		body["conversationId"] = request.ConversationId
	}

	if !tea.BoolValue(util.IsUnset(request.Members)) {
		body["members"] = request.Members
	}

	if !tea.BoolValue(util.IsUnset(request.OperatorUid)) {
		body["operatorUid"] = request.OperatorUid
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.OperationSource)) {
		realHeaders["operationSource"] = util.ToJSONString(headers.OperationSource)
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsDingtalkAccessToken)) {
		realHeaders["x-acs-dingtalk-access-token"] = util.ToJSONString(headers.XAcsDingtalkAccessToken)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	_result = &AddGroupMembersResponse{}
	_body, _err := client.DoROARequest(tea.String("AddGroupMembers"), tea.String("impaas_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/impaas/interconnections/groups/members/batchAdd"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddProfile(request *AddProfileRequest) (_result *AddProfileResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &AddProfileHeaders{}
	_result = &AddProfileResponse{}
	_body, _err := client.AddProfileWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AddProfileWithOptions(request *AddProfileRequest, headers *AddProfileHeaders, runtime *util.RuntimeOptions) (_result *AddProfileResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppUid)) {
		body["appUid"] = request.AppUid
	}

	if !tea.BoolValue(util.IsUnset(request.AvatarMediaId)) {
		body["avatarMediaId"] = request.AvatarMediaId
	}

	if !tea.BoolValue(util.IsUnset(request.MobileNumber)) {
		body["mobileNumber"] = request.MobileNumber
	}

	if !tea.BoolValue(util.IsUnset(request.Nick)) {
		body["nick"] = request.Nick
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
	_result = &AddProfileResponse{}
	_body, _err := client.DoROARequest(tea.String("AddProfile"), tea.String("impaas_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/impaas/interconnections/users/profiles"), tea.String("none"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) BatchSend(request *BatchSendRequest) (_result *BatchSendResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &BatchSendHeaders{}
	_result = &BatchSendResponse{}
	_body, _err := client.BatchSendWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) BatchSendWithOptions(request *BatchSendRequest, headers *BatchSendHeaders, runtime *util.RuntimeOptions) (_result *BatchSendResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppUids)) {
		body["appUids"] = request.AppUids
	}

	if !tea.BoolValue(util.IsUnset(request.Content)) {
		body["content"] = request.Content
	}

	if !tea.BoolValue(util.IsUnset(request.ConversationIds)) {
		body["conversationIds"] = request.ConversationIds
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		body["userId"] = request.UserId
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
	_result = &BatchSendResponse{}
	_body, _err := client.DoROARequest(tea.String("BatchSend"), tea.String("impaas_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/impaas/interconnections/messages/batchSend"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateGroup(request *CreateGroupRequest) (_result *CreateGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CreateGroupHeaders{}
	_result = &CreateGroupResponse{}
	_body, _err := client.CreateGroupWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateGroupWithOptions(request *CreateGroupRequest, headers *CreateGroupHeaders, runtime *util.RuntimeOptions) (_result *CreateGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Channel)) {
		body["channel"] = request.Channel
	}

	if !tea.BoolValue(util.IsUnset(request.CreatorUid)) {
		body["creatorUid"] = request.CreatorUid
	}

	if !tea.BoolValue(util.IsUnset(request.IconMediaId)) {
		body["iconMediaId"] = request.IconMediaId
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.Properties)) {
		body["properties"] = request.Properties
	}

	if !tea.BoolValue(util.IsUnset(request.Uuid)) {
		body["uuid"] = request.Uuid
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.OperationSource)) {
		realHeaders["operationSource"] = util.ToJSONString(headers.OperationSource)
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsDingtalkAccessToken)) {
		realHeaders["x-acs-dingtalk-access-token"] = util.ToJSONString(headers.XAcsDingtalkAccessToken)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	_result = &CreateGroupResponse{}
	_body, _err := client.DoROARequest(tea.String("CreateGroup"), tea.String("impaas_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/impaas/interconnections/groups"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateTrustGroup(request *CreateTrustGroupRequest) (_result *CreateTrustGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CreateTrustGroupHeaders{}
	_result = &CreateTrustGroupResponse{}
	_body, _err := client.CreateTrustGroupWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateTrustGroupWithOptions(request *CreateTrustGroupRequest, headers *CreateTrustGroupHeaders, runtime *util.RuntimeOptions) (_result *CreateTrustGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Channel)) {
		body["channel"] = request.Channel
	}

	if !tea.BoolValue(util.IsUnset(request.IconMediaId)) {
		body["iconMediaId"] = request.IconMediaId
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.Properties)) {
		body["properties"] = request.Properties
	}

	if !tea.BoolValue(util.IsUnset(request.Uuid)) {
		body["uuid"] = request.Uuid
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.OperationSource)) {
		realHeaders["operationSource"] = util.ToJSONString(headers.OperationSource)
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsDingtalkAccessToken)) {
		realHeaders["x-acs-dingtalk-access-token"] = util.ToJSONString(headers.XAcsDingtalkAccessToken)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	_result = &CreateTrustGroupResponse{}
	_body, _err := client.DoROARequest(tea.String("CreateTrustGroup"), tea.String("impaas_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/impaas/interconnections/groups/trusts"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DismissGroup(request *DismissGroupRequest) (_result *DismissGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &DismissGroupHeaders{}
	_result = &DismissGroupResponse{}
	_body, _err := client.DismissGroupWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DismissGroupWithOptions(request *DismissGroupRequest, headers *DismissGroupHeaders, runtime *util.RuntimeOptions) (_result *DismissGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ConversationId)) {
		body["conversationId"] = request.ConversationId
	}

	if !tea.BoolValue(util.IsUnset(request.OperatorUid)) {
		body["operatorUid"] = request.OperatorUid
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.OperationSource)) {
		realHeaders["operationSource"] = util.ToJSONString(headers.OperationSource)
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsDingtalkAccessToken)) {
		realHeaders["x-acs-dingtalk-access-token"] = util.ToJSONString(headers.XAcsDingtalkAccessToken)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	_result = &DismissGroupResponse{}
	_body, _err := client.DoROARequest(tea.String("DismissGroup"), tea.String("impaas_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/impaas/interconnections/groups/dismiss"), tea.String("none"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetConversationId(request *GetConversationIdRequest) (_result *GetConversationIdResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetConversationIdHeaders{}
	_result = &GetConversationIdResponse{}
	_body, _err := client.GetConversationIdWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetConversationIdWithOptions(request *GetConversationIdRequest, headers *GetConversationIdHeaders, runtime *util.RuntimeOptions) (_result *GetConversationIdResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppUid)) {
		body["appUid"] = request.AppUid
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		body["userId"] = request.UserId
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
	_result = &GetConversationIdResponse{}
	_body, _err := client.DoROARequest(tea.String("GetConversationId"), tea.String("impaas_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/impaas/interconnections/conversations"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetMediaUrl(request *GetMediaUrlRequest) (_result *GetMediaUrlResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetMediaUrlHeaders{}
	_result = &GetMediaUrlResponse{}
	_body, _err := client.GetMediaUrlWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetMediaUrlWithOptions(request *GetMediaUrlRequest, headers *GetMediaUrlHeaders, runtime *util.RuntimeOptions) (_result *GetMediaUrlResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MediaId)) {
		body["mediaId"] = request.MediaId
	}

	if !tea.BoolValue(util.IsUnset(request.UrlExpireTime)) {
		body["urlExpireTime"] = request.UrlExpireTime
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
	_result = &GetMediaUrlResponse{}
	_body, _err := client.DoROARequest(tea.String("GetMediaUrl"), tea.String("impaas_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/impaas/interconnections/medium/urls"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListGroupStaffMembers(request *ListGroupStaffMembersRequest) (_result *ListGroupStaffMembersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ListGroupStaffMembersHeaders{}
	_result = &ListGroupStaffMembersResponse{}
	_body, _err := client.ListGroupStaffMembersWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListGroupStaffMembersWithOptions(request *ListGroupStaffMembersRequest, headers *ListGroupStaffMembersHeaders, runtime *util.RuntimeOptions) (_result *ListGroupStaffMembersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ConversationId)) {
		body["conversationId"] = request.ConversationId
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
	_result = &ListGroupStaffMembersResponse{}
	_body, _err := client.DoROARequest(tea.String("ListGroupStaffMembers"), tea.String("impaas_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/impaas/interconnections/groups/staffMemers/query"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryBatchSendResult(request *QueryBatchSendResultRequest) (_result *QueryBatchSendResultResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryBatchSendResultHeaders{}
	_result = &QueryBatchSendResultResponse{}
	_body, _err := client.QueryBatchSendResultWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryBatchSendResultWithOptions(request *QueryBatchSendResultRequest, headers *QueryBatchSendResultHeaders, runtime *util.RuntimeOptions) (_result *QueryBatchSendResultResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.SenderUserId)) {
		query["senderUserId"] = request.SenderUserId
	}

	if !tea.BoolValue(util.IsUnset(request.TaskId)) {
		query["taskId"] = request.TaskId
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
	_result = &QueryBatchSendResultResponse{}
	_body, _err := client.DoROARequest(tea.String("QueryBatchSendResult"), tea.String("impaas_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/impaas/interconnections/messages/batchSendResults"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ReadMessage(request *ReadMessageRequest) (_result *ReadMessageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ReadMessageHeaders{}
	_result = &ReadMessageResponse{}
	_body, _err := client.ReadMessageWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ReadMessageWithOptions(request *ReadMessageRequest, headers *ReadMessageHeaders, runtime *util.RuntimeOptions) (_result *ReadMessageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MessageId)) {
		body["messageId"] = request.MessageId
	}

	if !tea.BoolValue(util.IsUnset(request.OperatorUid)) {
		body["operatorUid"] = request.OperatorUid
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.OperationSource)) {
		realHeaders["operationSource"] = util.ToJSONString(headers.OperationSource)
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsDingtalkAccessToken)) {
		realHeaders["x-acs-dingtalk-access-token"] = util.ToJSONString(headers.XAcsDingtalkAccessToken)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	_result = &ReadMessageResponse{}
	_body, _err := client.DoROARequest(tea.String("ReadMessage"), tea.String("impaas_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/impaas/interconnections/messages/read"), tea.String("none"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RecallMessage(request *RecallMessageRequest) (_result *RecallMessageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &RecallMessageHeaders{}
	_result = &RecallMessageResponse{}
	_body, _err := client.RecallMessageWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RecallMessageWithOptions(request *RecallMessageRequest, headers *RecallMessageHeaders, runtime *util.RuntimeOptions) (_result *RecallMessageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MessageId)) {
		body["messageId"] = request.MessageId
	}

	if !tea.BoolValue(util.IsUnset(request.OperatorUid)) {
		body["operatorUid"] = request.OperatorUid
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		body["type"] = request.Type
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.OperationSource)) {
		realHeaders["operationSource"] = util.ToJSONString(headers.OperationSource)
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsDingtalkAccessToken)) {
		realHeaders["x-acs-dingtalk-access-token"] = util.ToJSONString(headers.XAcsDingtalkAccessToken)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	_result = &RecallMessageResponse{}
	_body, _err := client.DoROARequest(tea.String("RecallMessage"), tea.String("impaas_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/impaas/interconnections/messages/recall"), tea.String("none"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RemoveGroupMembers(request *RemoveGroupMembersRequest) (_result *RemoveGroupMembersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &RemoveGroupMembersHeaders{}
	_result = &RemoveGroupMembersResponse{}
	_body, _err := client.RemoveGroupMembersWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RemoveGroupMembersWithOptions(request *RemoveGroupMembersRequest, headers *RemoveGroupMembersHeaders, runtime *util.RuntimeOptions) (_result *RemoveGroupMembersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ConversationId)) {
		body["conversationId"] = request.ConversationId
	}

	if !tea.BoolValue(util.IsUnset(request.MemberUids)) {
		body["memberUids"] = request.MemberUids
	}

	if !tea.BoolValue(util.IsUnset(request.OperatorUid)) {
		body["operatorUid"] = request.OperatorUid
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.OperationSource)) {
		realHeaders["operationSource"] = util.ToJSONString(headers.OperationSource)
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsDingtalkAccessToken)) {
		realHeaders["x-acs-dingtalk-access-token"] = util.ToJSONString(headers.XAcsDingtalkAccessToken)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	_result = &RemoveGroupMembersResponse{}
	_body, _err := client.DoROARequest(tea.String("RemoveGroupMembers"), tea.String("impaas_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/impaas/interconnections/groups/members/batchRemove"), tea.String("none"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SendMessage(request *SendMessageRequest) (_result *SendMessageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &SendMessageHeaders{}
	_result = &SendMessageResponse{}
	_body, _err := client.SendMessageWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SendMessageWithOptions(request *SendMessageRequest, headers *SendMessageHeaders, runtime *util.RuntimeOptions) (_result *SendMessageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Content)) {
		body["content"] = request.Content
	}

	if !tea.BoolValue(util.IsUnset(request.ConversationId)) {
		body["conversationId"] = request.ConversationId
	}

	if !tea.BoolValue(util.IsUnset(request.CreateTime)) {
		body["createTime"] = request.CreateTime
	}

	if !tea.BoolValue(util.IsUnset(request.ReceiverUid)) {
		body["receiverUid"] = request.ReceiverUid
	}

	if !tea.BoolValue(util.IsUnset(request.SenderUid)) {
		body["senderUid"] = request.SenderUid
	}

	if !tea.BoolValue(util.IsUnset(request.Uuid)) {
		body["uuid"] = request.Uuid
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.OperationSource)) {
		realHeaders["operationSource"] = util.ToJSONString(headers.OperationSource)
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsDingtalkAccessToken)) {
		realHeaders["x-acs-dingtalk-access-token"] = util.ToJSONString(headers.XAcsDingtalkAccessToken)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	_result = &SendMessageResponse{}
	_body, _err := client.DoROARequest(tea.String("SendMessage"), tea.String("impaas_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/impaas/interconnections/messages/send"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateGroupName(request *UpdateGroupNameRequest) (_result *UpdateGroupNameResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &UpdateGroupNameHeaders{}
	_result = &UpdateGroupNameResponse{}
	_body, _err := client.UpdateGroupNameWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateGroupNameWithOptions(request *UpdateGroupNameRequest, headers *UpdateGroupNameHeaders, runtime *util.RuntimeOptions) (_result *UpdateGroupNameResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ConversationId)) {
		body["conversationId"] = request.ConversationId
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.OperatorUid)) {
		body["operatorUid"] = request.OperatorUid
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.OperationSource)) {
		realHeaders["operationSource"] = util.ToJSONString(headers.OperationSource)
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsDingtalkAccessToken)) {
		realHeaders["x-acs-dingtalk-access-token"] = util.ToJSONString(headers.XAcsDingtalkAccessToken)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	_result = &UpdateGroupNameResponse{}
	_body, _err := client.DoROARequest(tea.String("UpdateGroupName"), tea.String("impaas_1.0"), tea.String("HTTP"), tea.String("PUT"), tea.String("AK"), tea.String("/v1.0/impaas/interconnections/groups/names"), tea.String("none"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateGroupOwner(request *UpdateGroupOwnerRequest) (_result *UpdateGroupOwnerResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &UpdateGroupOwnerHeaders{}
	_result = &UpdateGroupOwnerResponse{}
	_body, _err := client.UpdateGroupOwnerWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateGroupOwnerWithOptions(request *UpdateGroupOwnerRequest, headers *UpdateGroupOwnerHeaders, runtime *util.RuntimeOptions) (_result *UpdateGroupOwnerResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ConversationId)) {
		body["conversationId"] = request.ConversationId
	}

	if !tea.BoolValue(util.IsUnset(request.OperatorUid)) {
		body["operatorUid"] = request.OperatorUid
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerUid)) {
		body["ownerUid"] = request.OwnerUid
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.OperationSource)) {
		realHeaders["operationSource"] = util.ToJSONString(headers.OperationSource)
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsDingtalkAccessToken)) {
		realHeaders["x-acs-dingtalk-access-token"] = util.ToJSONString(headers.XAcsDingtalkAccessToken)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	_result = &UpdateGroupOwnerResponse{}
	_body, _err := client.DoROARequest(tea.String("UpdateGroupOwner"), tea.String("impaas_1.0"), tea.String("HTTP"), tea.String("PUT"), tea.String("AK"), tea.String("/v1.0/impaas/interconnections/groups/owners"), tea.String("none"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}
