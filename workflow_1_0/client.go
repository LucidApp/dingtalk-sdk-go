// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package workflow_1_0

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
)

type QueryFormInstanceHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryFormInstanceHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryFormInstanceHeaders) GoString() string {
	return s.String()
}

func (s *QueryFormInstanceHeaders) SetCommonHeaders(v map[string]*string) *QueryFormInstanceHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryFormInstanceHeaders) SetXAcsDingtalkAccessToken(v string) *QueryFormInstanceHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryFormInstanceRequest struct {
	// 表单实例id
	FormInstanceId *string `json:"formInstanceId,omitempty" xml:"formInstanceId,omitempty"`
	// 表单模板Code
	FormCode *string `json:"formCode,omitempty" xml:"formCode,omitempty"`
	// 应用搭建id
	AppUuid *string `json:"appUuid,omitempty" xml:"appUuid,omitempty"`
}

func (s QueryFormInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryFormInstanceRequest) GoString() string {
	return s.String()
}

func (s *QueryFormInstanceRequest) SetFormInstanceId(v string) *QueryFormInstanceRequest {
	s.FormInstanceId = &v
	return s
}

func (s *QueryFormInstanceRequest) SetFormCode(v string) *QueryFormInstanceRequest {
	s.FormCode = &v
	return s
}

func (s *QueryFormInstanceRequest) SetAppUuid(v string) *QueryFormInstanceRequest {
	s.AppUuid = &v
	return s
}

type QueryFormInstanceResponseBody struct {
	// 实例id
	FormInstanceId *string `json:"formInstanceId,omitempty" xml:"formInstanceId,omitempty"`
	// 表单数据
	FormInstDataList []*QueryFormInstanceResponseBodyFormInstDataList `json:"formInstDataList,omitempty" xml:"formInstDataList,omitempty" type:"Repeated"`
	// 应用搭建id
	AppUuid *string `json:"appUuid,omitempty" xml:"appUuid,omitempty"`
	// 表单模板id
	FormCode *string `json:"formCode,omitempty" xml:"formCode,omitempty"`
	// 表单标题
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
	// 创建人
	Creator *string `json:"creator,omitempty" xml:"creator,omitempty"`
	// 修改人
	Modifier *string `json:"modifier,omitempty" xml:"modifier,omitempty"`
	// 实例创建时间戳
	CreateTimestamp *int64 `json:"createTimestamp,omitempty" xml:"createTimestamp,omitempty"`
	// 实例最近修改时间戳
	ModifyTimestamp *int64 `json:"modifyTimestamp,omitempty" xml:"modifyTimestamp,omitempty"`
	// 外联业务实例id
	OutInstanceId *string `json:"outInstanceId,omitempty" xml:"outInstanceId,omitempty"`
	// 外联业务code
	OutBizCode *string `json:"outBizCode,omitempty" xml:"outBizCode,omitempty"`
	// 扩展信息
	Attributes map[string]interface{} `json:"attributes,omitempty" xml:"attributes,omitempty"`
}

func (s QueryFormInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryFormInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *QueryFormInstanceResponseBody) SetFormInstanceId(v string) *QueryFormInstanceResponseBody {
	s.FormInstanceId = &v
	return s
}

func (s *QueryFormInstanceResponseBody) SetFormInstDataList(v []*QueryFormInstanceResponseBodyFormInstDataList) *QueryFormInstanceResponseBody {
	s.FormInstDataList = v
	return s
}

func (s *QueryFormInstanceResponseBody) SetAppUuid(v string) *QueryFormInstanceResponseBody {
	s.AppUuid = &v
	return s
}

func (s *QueryFormInstanceResponseBody) SetFormCode(v string) *QueryFormInstanceResponseBody {
	s.FormCode = &v
	return s
}

func (s *QueryFormInstanceResponseBody) SetTitle(v string) *QueryFormInstanceResponseBody {
	s.Title = &v
	return s
}

func (s *QueryFormInstanceResponseBody) SetCreator(v string) *QueryFormInstanceResponseBody {
	s.Creator = &v
	return s
}

func (s *QueryFormInstanceResponseBody) SetModifier(v string) *QueryFormInstanceResponseBody {
	s.Modifier = &v
	return s
}

func (s *QueryFormInstanceResponseBody) SetCreateTimestamp(v int64) *QueryFormInstanceResponseBody {
	s.CreateTimestamp = &v
	return s
}

func (s *QueryFormInstanceResponseBody) SetModifyTimestamp(v int64) *QueryFormInstanceResponseBody {
	s.ModifyTimestamp = &v
	return s
}

func (s *QueryFormInstanceResponseBody) SetOutInstanceId(v string) *QueryFormInstanceResponseBody {
	s.OutInstanceId = &v
	return s
}

func (s *QueryFormInstanceResponseBody) SetOutBizCode(v string) *QueryFormInstanceResponseBody {
	s.OutBizCode = &v
	return s
}

func (s *QueryFormInstanceResponseBody) SetAttributes(v map[string]interface{}) *QueryFormInstanceResponseBody {
	s.Attributes = v
	return s
}

type QueryFormInstanceResponseBodyFormInstDataList struct {
	ComponentType *string `json:"componentType,omitempty" xml:"componentType,omitempty"`
	BizAlias      *string `json:"bizAlias,omitempty" xml:"bizAlias,omitempty"`
	ExtendValue   *string `json:"extendValue,omitempty" xml:"extendValue,omitempty"`
	Label         *string `json:"label,omitempty" xml:"label,omitempty"`
	Value         *string `json:"value,omitempty" xml:"value,omitempty"`
	Key           *string `json:"key,omitempty" xml:"key,omitempty"`
}

func (s QueryFormInstanceResponseBodyFormInstDataList) String() string {
	return tea.Prettify(s)
}

func (s QueryFormInstanceResponseBodyFormInstDataList) GoString() string {
	return s.String()
}

func (s *QueryFormInstanceResponseBodyFormInstDataList) SetComponentType(v string) *QueryFormInstanceResponseBodyFormInstDataList {
	s.ComponentType = &v
	return s
}

func (s *QueryFormInstanceResponseBodyFormInstDataList) SetBizAlias(v string) *QueryFormInstanceResponseBodyFormInstDataList {
	s.BizAlias = &v
	return s
}

func (s *QueryFormInstanceResponseBodyFormInstDataList) SetExtendValue(v string) *QueryFormInstanceResponseBodyFormInstDataList {
	s.ExtendValue = &v
	return s
}

func (s *QueryFormInstanceResponseBodyFormInstDataList) SetLabel(v string) *QueryFormInstanceResponseBodyFormInstDataList {
	s.Label = &v
	return s
}

func (s *QueryFormInstanceResponseBodyFormInstDataList) SetValue(v string) *QueryFormInstanceResponseBodyFormInstDataList {
	s.Value = &v
	return s
}

func (s *QueryFormInstanceResponseBodyFormInstDataList) SetKey(v string) *QueryFormInstanceResponseBodyFormInstDataList {
	s.Key = &v
	return s
}

type QueryFormInstanceResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryFormInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryFormInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryFormInstanceResponse) GoString() string {
	return s.String()
}

func (s *QueryFormInstanceResponse) SetHeaders(v map[string]*string) *QueryFormInstanceResponse {
	s.Headers = v
	return s
}

func (s *QueryFormInstanceResponse) SetBody(v *QueryFormInstanceResponseBody) *QueryFormInstanceResponse {
	s.Body = v
	return s
}

type ProcessForecastHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s ProcessForecastHeaders) String() string {
	return tea.Prettify(s)
}

func (s ProcessForecastHeaders) GoString() string {
	return s.String()
}

func (s *ProcessForecastHeaders) SetCommonHeaders(v map[string]*string) *ProcessForecastHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ProcessForecastHeaders) SetXAcsDingtalkAccessToken(v string) *ProcessForecastHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type ProcessForecastRequest struct {
	DingCorpId         *string `json:"dingCorpId,omitempty" xml:"dingCorpId,omitempty"`
	DingOrgId          *int64  `json:"dingOrgId,omitempty" xml:"dingOrgId,omitempty"`
	DingIsvOrgId       *int64  `json:"dingIsvOrgId,omitempty" xml:"dingIsvOrgId,omitempty"`
	DingSuiteKey       *string `json:"dingSuiteKey,omitempty" xml:"dingSuiteKey,omitempty"`
	DingTokenGrantType *int64  `json:"dingTokenGrantType,omitempty" xml:"dingTokenGrantType,omitempty"`
	RequestId          *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 审批流的唯一码
	ProcessCode *string `json:"processCode,omitempty" xml:"processCode,omitempty"`
	// 部门ID
	DeptId *int32 `json:"deptId,omitempty" xml:"deptId,omitempty"`
	// 审批发起人的userId
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
	// 表单数据内容，控件列表
	FormComponentValues []*ProcessForecastRequestFormComponentValues `json:"formComponentValues,omitempty" xml:"formComponentValues,omitempty" type:"Repeated"`
}

func (s ProcessForecastRequest) String() string {
	return tea.Prettify(s)
}

func (s ProcessForecastRequest) GoString() string {
	return s.String()
}

func (s *ProcessForecastRequest) SetDingCorpId(v string) *ProcessForecastRequest {
	s.DingCorpId = &v
	return s
}

func (s *ProcessForecastRequest) SetDingOrgId(v int64) *ProcessForecastRequest {
	s.DingOrgId = &v
	return s
}

func (s *ProcessForecastRequest) SetDingIsvOrgId(v int64) *ProcessForecastRequest {
	s.DingIsvOrgId = &v
	return s
}

func (s *ProcessForecastRequest) SetDingSuiteKey(v string) *ProcessForecastRequest {
	s.DingSuiteKey = &v
	return s
}

func (s *ProcessForecastRequest) SetDingTokenGrantType(v int64) *ProcessForecastRequest {
	s.DingTokenGrantType = &v
	return s
}

func (s *ProcessForecastRequest) SetRequestId(v string) *ProcessForecastRequest {
	s.RequestId = &v
	return s
}

func (s *ProcessForecastRequest) SetProcessCode(v string) *ProcessForecastRequest {
	s.ProcessCode = &v
	return s
}

func (s *ProcessForecastRequest) SetDeptId(v int32) *ProcessForecastRequest {
	s.DeptId = &v
	return s
}

func (s *ProcessForecastRequest) SetUserId(v string) *ProcessForecastRequest {
	s.UserId = &v
	return s
}

func (s *ProcessForecastRequest) SetFormComponentValues(v []*ProcessForecastRequestFormComponentValues) *ProcessForecastRequest {
	s.FormComponentValues = v
	return s
}

type ProcessForecastRequestFormComponentValues struct {
	// 控件id
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// 控件别名
	BizAlias *string `json:"bizAlias,omitempty" xml:"bizAlias,omitempty"`
	// 控件名称
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 控件值
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
	// 控件扩展值
	ExtValue *string `json:"extValue,omitempty" xml:"extValue,omitempty"`
	// 控件类型
	ComponentType *string                                             `json:"componentType,omitempty" xml:"componentType,omitempty"`
	Details       []*ProcessForecastRequestFormComponentValuesDetails `json:"details,omitempty" xml:"details,omitempty" type:"Repeated"`
}

func (s ProcessForecastRequestFormComponentValues) String() string {
	return tea.Prettify(s)
}

func (s ProcessForecastRequestFormComponentValues) GoString() string {
	return s.String()
}

func (s *ProcessForecastRequestFormComponentValues) SetId(v string) *ProcessForecastRequestFormComponentValues {
	s.Id = &v
	return s
}

func (s *ProcessForecastRequestFormComponentValues) SetBizAlias(v string) *ProcessForecastRequestFormComponentValues {
	s.BizAlias = &v
	return s
}

func (s *ProcessForecastRequestFormComponentValues) SetName(v string) *ProcessForecastRequestFormComponentValues {
	s.Name = &v
	return s
}

func (s *ProcessForecastRequestFormComponentValues) SetValue(v string) *ProcessForecastRequestFormComponentValues {
	s.Value = &v
	return s
}

func (s *ProcessForecastRequestFormComponentValues) SetExtValue(v string) *ProcessForecastRequestFormComponentValues {
	s.ExtValue = &v
	return s
}

func (s *ProcessForecastRequestFormComponentValues) SetComponentType(v string) *ProcessForecastRequestFormComponentValues {
	s.ComponentType = &v
	return s
}

func (s *ProcessForecastRequestFormComponentValues) SetDetails(v []*ProcessForecastRequestFormComponentValuesDetails) *ProcessForecastRequestFormComponentValues {
	s.Details = v
	return s
}

type ProcessForecastRequestFormComponentValuesDetails struct {
	// 控件id
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// 控件别名
	BizAlias *string `json:"bizAlias,omitempty" xml:"bizAlias,omitempty"`
	// 控件名称
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 控件值
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
	// 控件扩展值
	ExtValue *string                                                    `json:"extValue,omitempty" xml:"extValue,omitempty"`
	Details  []*ProcessForecastRequestFormComponentValuesDetailsDetails `json:"details,omitempty" xml:"details,omitempty" type:"Repeated"`
}

func (s ProcessForecastRequestFormComponentValuesDetails) String() string {
	return tea.Prettify(s)
}

func (s ProcessForecastRequestFormComponentValuesDetails) GoString() string {
	return s.String()
}

func (s *ProcessForecastRequestFormComponentValuesDetails) SetId(v string) *ProcessForecastRequestFormComponentValuesDetails {
	s.Id = &v
	return s
}

func (s *ProcessForecastRequestFormComponentValuesDetails) SetBizAlias(v string) *ProcessForecastRequestFormComponentValuesDetails {
	s.BizAlias = &v
	return s
}

func (s *ProcessForecastRequestFormComponentValuesDetails) SetName(v string) *ProcessForecastRequestFormComponentValuesDetails {
	s.Name = &v
	return s
}

func (s *ProcessForecastRequestFormComponentValuesDetails) SetValue(v string) *ProcessForecastRequestFormComponentValuesDetails {
	s.Value = &v
	return s
}

func (s *ProcessForecastRequestFormComponentValuesDetails) SetExtValue(v string) *ProcessForecastRequestFormComponentValuesDetails {
	s.ExtValue = &v
	return s
}

func (s *ProcessForecastRequestFormComponentValuesDetails) SetDetails(v []*ProcessForecastRequestFormComponentValuesDetailsDetails) *ProcessForecastRequestFormComponentValuesDetails {
	s.Details = v
	return s
}

type ProcessForecastRequestFormComponentValuesDetailsDetails struct {
	// 控件id
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// 控件别名
	BizAlias *string `json:"bizAlias,omitempty" xml:"bizAlias,omitempty"`
	// 控件名称
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 控件值
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
	// 控件扩展值
	ExtValue      *string `json:"extValue,omitempty" xml:"extValue,omitempty"`
	ComponentType *string `json:"componentType,omitempty" xml:"componentType,omitempty"`
}

func (s ProcessForecastRequestFormComponentValuesDetailsDetails) String() string {
	return tea.Prettify(s)
}

func (s ProcessForecastRequestFormComponentValuesDetailsDetails) GoString() string {
	return s.String()
}

func (s *ProcessForecastRequestFormComponentValuesDetailsDetails) SetId(v string) *ProcessForecastRequestFormComponentValuesDetailsDetails {
	s.Id = &v
	return s
}

func (s *ProcessForecastRequestFormComponentValuesDetailsDetails) SetBizAlias(v string) *ProcessForecastRequestFormComponentValuesDetailsDetails {
	s.BizAlias = &v
	return s
}

func (s *ProcessForecastRequestFormComponentValuesDetailsDetails) SetName(v string) *ProcessForecastRequestFormComponentValuesDetailsDetails {
	s.Name = &v
	return s
}

func (s *ProcessForecastRequestFormComponentValuesDetailsDetails) SetValue(v string) *ProcessForecastRequestFormComponentValuesDetailsDetails {
	s.Value = &v
	return s
}

func (s *ProcessForecastRequestFormComponentValuesDetailsDetails) SetExtValue(v string) *ProcessForecastRequestFormComponentValuesDetailsDetails {
	s.ExtValue = &v
	return s
}

func (s *ProcessForecastRequestFormComponentValuesDetailsDetails) SetComponentType(v string) *ProcessForecastRequestFormComponentValuesDetailsDetails {
	s.ComponentType = &v
	return s
}

type ProcessForecastResponseBody struct {
	// 返回结果
	Result *ProcessForecastResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s ProcessForecastResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ProcessForecastResponseBody) GoString() string {
	return s.String()
}

func (s *ProcessForecastResponseBody) SetResult(v *ProcessForecastResponseBodyResult) *ProcessForecastResponseBody {
	s.Result = v
	return s
}

type ProcessForecastResponseBodyResult struct {
	// 是否预测成功
	IsForecastSuccess *bool `json:"isForecastSuccess,omitempty" xml:"isForecastSuccess,omitempty"`
	// 流程 code
	ProcessCode *string `json:"processCode,omitempty" xml:"processCode,omitempty"`
	// 用户 id
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
	// 流程 id
	ProcessId *int64 `json:"processId,omitempty" xml:"processId,omitempty"`
	// 是否静态流程
	IsStaticWorkflow      *bool                                                     `json:"isStaticWorkflow,omitempty" xml:"isStaticWorkflow,omitempty"`
	WorkflowActivityRules []*ProcessForecastResponseBodyResultWorkflowActivityRules `json:"workflowActivityRules,omitempty" xml:"workflowActivityRules,omitempty" type:"Repeated"`
	WorkflowForecastNodes []*ProcessForecastResponseBodyResultWorkflowForecastNodes `json:"workflowForecastNodes,omitempty" xml:"workflowForecastNodes,omitempty" type:"Repeated"`
}

func (s ProcessForecastResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ProcessForecastResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ProcessForecastResponseBodyResult) SetIsForecastSuccess(v bool) *ProcessForecastResponseBodyResult {
	s.IsForecastSuccess = &v
	return s
}

func (s *ProcessForecastResponseBodyResult) SetProcessCode(v string) *ProcessForecastResponseBodyResult {
	s.ProcessCode = &v
	return s
}

func (s *ProcessForecastResponseBodyResult) SetUserId(v string) *ProcessForecastResponseBodyResult {
	s.UserId = &v
	return s
}

func (s *ProcessForecastResponseBodyResult) SetProcessId(v int64) *ProcessForecastResponseBodyResult {
	s.ProcessId = &v
	return s
}

func (s *ProcessForecastResponseBodyResult) SetIsStaticWorkflow(v bool) *ProcessForecastResponseBodyResult {
	s.IsStaticWorkflow = &v
	return s
}

func (s *ProcessForecastResponseBodyResult) SetWorkflowActivityRules(v []*ProcessForecastResponseBodyResultWorkflowActivityRules) *ProcessForecastResponseBodyResult {
	s.WorkflowActivityRules = v
	return s
}

func (s *ProcessForecastResponseBodyResult) SetWorkflowForecastNodes(v []*ProcessForecastResponseBodyResultWorkflowForecastNodes) *ProcessForecastResponseBodyResult {
	s.WorkflowForecastNodes = v
	return s
}

type ProcessForecastResponseBodyResultWorkflowActivityRules struct {
	// 节点 id
	ActivityId *string `json:"activityId,omitempty" xml:"activityId,omitempty"`
	// 流程中前一个节点的 id
	PrevActivityId *string `json:"prevActivityId,omitempty" xml:"prevActivityId,omitempty"`
	// 节点名称
	ActivityName *string `json:"activityName,omitempty" xml:"activityName,omitempty"`
	// 规则类型
	ActivityType *string `json:"activityType,omitempty" xml:"activityType,omitempty"`
	// 是否自选审批节点
	IsTargetSelect *bool `json:"isTargetSelect,omitempty" xml:"isTargetSelect,omitempty"`
	// 节点操作人信息
	WorkflowActor *ProcessForecastResponseBodyResultWorkflowActivityRulesWorkflowActor `json:"workflowActor,omitempty" xml:"workflowActor,omitempty" type:"Struct"`
}

func (s ProcessForecastResponseBodyResultWorkflowActivityRules) String() string {
	return tea.Prettify(s)
}

func (s ProcessForecastResponseBodyResultWorkflowActivityRules) GoString() string {
	return s.String()
}

func (s *ProcessForecastResponseBodyResultWorkflowActivityRules) SetActivityId(v string) *ProcessForecastResponseBodyResultWorkflowActivityRules {
	s.ActivityId = &v
	return s
}

func (s *ProcessForecastResponseBodyResultWorkflowActivityRules) SetPrevActivityId(v string) *ProcessForecastResponseBodyResultWorkflowActivityRules {
	s.PrevActivityId = &v
	return s
}

func (s *ProcessForecastResponseBodyResultWorkflowActivityRules) SetActivityName(v string) *ProcessForecastResponseBodyResultWorkflowActivityRules {
	s.ActivityName = &v
	return s
}

func (s *ProcessForecastResponseBodyResultWorkflowActivityRules) SetActivityType(v string) *ProcessForecastResponseBodyResultWorkflowActivityRules {
	s.ActivityType = &v
	return s
}

func (s *ProcessForecastResponseBodyResultWorkflowActivityRules) SetIsTargetSelect(v bool) *ProcessForecastResponseBodyResultWorkflowActivityRules {
	s.IsTargetSelect = &v
	return s
}

func (s *ProcessForecastResponseBodyResultWorkflowActivityRules) SetWorkflowActor(v *ProcessForecastResponseBodyResultWorkflowActivityRulesWorkflowActor) *ProcessForecastResponseBodyResultWorkflowActivityRules {
	s.WorkflowActor = v
	return s
}

type ProcessForecastResponseBodyResultWorkflowActivityRulesWorkflowActor struct {
	// 节点操作人 key
	ActorKey *string `json:"actorKey,omitempty" xml:"actorKey,omitempty"`
	// 节点操作人类型
	ActorType *string `json:"actorType,omitempty" xml:"actorType,omitempty"`
	// 节点操作人选择范围类型
	ActorSelectionType *string `json:"actorSelectionType,omitempty" xml:"actorSelectionType,omitempty"`
	// 节点操作人选择范围
	ActorSelectionRange *ProcessForecastResponseBodyResultWorkflowActivityRulesWorkflowActorActorSelectionRange `json:"actorSelectionRange,omitempty" xml:"actorSelectionRange,omitempty" type:"Struct"`
	// 是否允许多选，还是仅允许选一人
	AllowedMulti *bool `json:"allowedMulti,omitempty" xml:"allowedMulti,omitempty"`
	// 节点审批类型
	ApprovalType *string `json:"approvalType,omitempty" xml:"approvalType,omitempty"`
	// 节点审批方式
	ApprovalMethod *string `json:"approvalMethod,omitempty" xml:"approvalMethod,omitempty"`
	// 节点激活类型
	ActorActivateType *string `json:"actorActivateType,omitempty" xml:"actorActivateType,omitempty"`
	// 该审批人节点在发起审批时是否必填
	Required *bool `json:"required,omitempty" xml:"required,omitempty"`
}

func (s ProcessForecastResponseBodyResultWorkflowActivityRulesWorkflowActor) String() string {
	return tea.Prettify(s)
}

func (s ProcessForecastResponseBodyResultWorkflowActivityRulesWorkflowActor) GoString() string {
	return s.String()
}

func (s *ProcessForecastResponseBodyResultWorkflowActivityRulesWorkflowActor) SetActorKey(v string) *ProcessForecastResponseBodyResultWorkflowActivityRulesWorkflowActor {
	s.ActorKey = &v
	return s
}

func (s *ProcessForecastResponseBodyResultWorkflowActivityRulesWorkflowActor) SetActorType(v string) *ProcessForecastResponseBodyResultWorkflowActivityRulesWorkflowActor {
	s.ActorType = &v
	return s
}

func (s *ProcessForecastResponseBodyResultWorkflowActivityRulesWorkflowActor) SetActorSelectionType(v string) *ProcessForecastResponseBodyResultWorkflowActivityRulesWorkflowActor {
	s.ActorSelectionType = &v
	return s
}

func (s *ProcessForecastResponseBodyResultWorkflowActivityRulesWorkflowActor) SetActorSelectionRange(v *ProcessForecastResponseBodyResultWorkflowActivityRulesWorkflowActorActorSelectionRange) *ProcessForecastResponseBodyResultWorkflowActivityRulesWorkflowActor {
	s.ActorSelectionRange = v
	return s
}

func (s *ProcessForecastResponseBodyResultWorkflowActivityRulesWorkflowActor) SetAllowedMulti(v bool) *ProcessForecastResponseBodyResultWorkflowActivityRulesWorkflowActor {
	s.AllowedMulti = &v
	return s
}

func (s *ProcessForecastResponseBodyResultWorkflowActivityRulesWorkflowActor) SetApprovalType(v string) *ProcessForecastResponseBodyResultWorkflowActivityRulesWorkflowActor {
	s.ApprovalType = &v
	return s
}

func (s *ProcessForecastResponseBodyResultWorkflowActivityRulesWorkflowActor) SetApprovalMethod(v string) *ProcessForecastResponseBodyResultWorkflowActivityRulesWorkflowActor {
	s.ApprovalMethod = &v
	return s
}

func (s *ProcessForecastResponseBodyResultWorkflowActivityRulesWorkflowActor) SetActorActivateType(v string) *ProcessForecastResponseBodyResultWorkflowActivityRulesWorkflowActor {
	s.ActorActivateType = &v
	return s
}

func (s *ProcessForecastResponseBodyResultWorkflowActivityRulesWorkflowActor) SetRequired(v bool) *ProcessForecastResponseBodyResultWorkflowActivityRulesWorkflowActor {
	s.Required = &v
	return s
}

type ProcessForecastResponseBodyResultWorkflowActivityRulesWorkflowActorActorSelectionRange struct {
	// 审批指定成员
	Approvals []*ProcessForecastResponseBodyResultWorkflowActivityRulesWorkflowActorActorSelectionRangeApprovals `json:"approvals,omitempty" xml:"approvals,omitempty" type:"Repeated"`
	// 审批指定角色
	Labels []*ProcessForecastResponseBodyResultWorkflowActivityRulesWorkflowActorActorSelectionRangeLabels `json:"labels,omitempty" xml:"labels,omitempty" type:"Repeated"`
}

func (s ProcessForecastResponseBodyResultWorkflowActivityRulesWorkflowActorActorSelectionRange) String() string {
	return tea.Prettify(s)
}

func (s ProcessForecastResponseBodyResultWorkflowActivityRulesWorkflowActorActorSelectionRange) GoString() string {
	return s.String()
}

func (s *ProcessForecastResponseBodyResultWorkflowActivityRulesWorkflowActorActorSelectionRange) SetApprovals(v []*ProcessForecastResponseBodyResultWorkflowActivityRulesWorkflowActorActorSelectionRangeApprovals) *ProcessForecastResponseBodyResultWorkflowActivityRulesWorkflowActorActorSelectionRange {
	s.Approvals = v
	return s
}

func (s *ProcessForecastResponseBodyResultWorkflowActivityRulesWorkflowActorActorSelectionRange) SetLabels(v []*ProcessForecastResponseBodyResultWorkflowActivityRulesWorkflowActorActorSelectionRangeLabels) *ProcessForecastResponseBodyResultWorkflowActivityRulesWorkflowActorActorSelectionRange {
	s.Labels = v
	return s
}

type ProcessForecastResponseBodyResultWorkflowActivityRulesWorkflowActorActorSelectionRangeApprovals struct {
	// 员工 userId
	WorkNo *string `json:"workNo,omitempty" xml:"workNo,omitempty"`
	// 员工姓名
	UserName *string `json:"userName,omitempty" xml:"userName,omitempty"`
}

func (s ProcessForecastResponseBodyResultWorkflowActivityRulesWorkflowActorActorSelectionRangeApprovals) String() string {
	return tea.Prettify(s)
}

func (s ProcessForecastResponseBodyResultWorkflowActivityRulesWorkflowActorActorSelectionRangeApprovals) GoString() string {
	return s.String()
}

func (s *ProcessForecastResponseBodyResultWorkflowActivityRulesWorkflowActorActorSelectionRangeApprovals) SetWorkNo(v string) *ProcessForecastResponseBodyResultWorkflowActivityRulesWorkflowActorActorSelectionRangeApprovals {
	s.WorkNo = &v
	return s
}

func (s *ProcessForecastResponseBodyResultWorkflowActivityRulesWorkflowActorActorSelectionRangeApprovals) SetUserName(v string) *ProcessForecastResponseBodyResultWorkflowActivityRulesWorkflowActorActorSelectionRangeApprovals {
	s.UserName = &v
	return s
}

type ProcessForecastResponseBodyResultWorkflowActivityRulesWorkflowActorActorSelectionRangeLabels struct {
	// 角色 id
	Labels *string `json:"labels,omitempty" xml:"labels,omitempty"`
	// 角色名字
	LabelNames *string `json:"labelNames,omitempty" xml:"labelNames,omitempty"`
}

func (s ProcessForecastResponseBodyResultWorkflowActivityRulesWorkflowActorActorSelectionRangeLabels) String() string {
	return tea.Prettify(s)
}

func (s ProcessForecastResponseBodyResultWorkflowActivityRulesWorkflowActorActorSelectionRangeLabels) GoString() string {
	return s.String()
}

func (s *ProcessForecastResponseBodyResultWorkflowActivityRulesWorkflowActorActorSelectionRangeLabels) SetLabels(v string) *ProcessForecastResponseBodyResultWorkflowActivityRulesWorkflowActorActorSelectionRangeLabels {
	s.Labels = &v
	return s
}

func (s *ProcessForecastResponseBodyResultWorkflowActivityRulesWorkflowActorActorSelectionRangeLabels) SetLabelNames(v string) *ProcessForecastResponseBodyResultWorkflowActivityRulesWorkflowActorActorSelectionRangeLabels {
	s.LabelNames = &v
	return s
}

type ProcessForecastResponseBodyResultWorkflowForecastNodes struct {
	// 节点 id
	ActivityId *string `json:"activityId,omitempty" xml:"activityId,omitempty"`
	// 节点出线 id
	OutId *string `json:"outId,omitempty" xml:"outId,omitempty"`
}

func (s ProcessForecastResponseBodyResultWorkflowForecastNodes) String() string {
	return tea.Prettify(s)
}

func (s ProcessForecastResponseBodyResultWorkflowForecastNodes) GoString() string {
	return s.String()
}

func (s *ProcessForecastResponseBodyResultWorkflowForecastNodes) SetActivityId(v string) *ProcessForecastResponseBodyResultWorkflowForecastNodes {
	s.ActivityId = &v
	return s
}

func (s *ProcessForecastResponseBodyResultWorkflowForecastNodes) SetOutId(v string) *ProcessForecastResponseBodyResultWorkflowForecastNodes {
	s.OutId = &v
	return s
}

type ProcessForecastResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ProcessForecastResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ProcessForecastResponse) String() string {
	return tea.Prettify(s)
}

func (s ProcessForecastResponse) GoString() string {
	return s.String()
}

func (s *ProcessForecastResponse) SetHeaders(v map[string]*string) *ProcessForecastResponse {
	s.Headers = v
	return s
}

func (s *ProcessForecastResponse) SetBody(v *ProcessForecastResponseBody) *ProcessForecastResponse {
	s.Body = v
	return s
}

type QueryAllProcessInstancesHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryAllProcessInstancesHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryAllProcessInstancesHeaders) GoString() string {
	return s.String()
}

func (s *QueryAllProcessInstancesHeaders) SetCommonHeaders(v map[string]*string) *QueryAllProcessInstancesHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryAllProcessInstancesHeaders) SetXAcsDingtalkAccessToken(v string) *QueryAllProcessInstancesHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryAllProcessInstancesRequest struct {
	// 分页起始值
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// 分页大小
	MaxResults *int64 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// 开始时间
	StartTimeInMills *int64 `json:"startTimeInMills,omitempty" xml:"startTimeInMills,omitempty"`
	// 结束时间
	EndTimeInMills *int64 `json:"endTimeInMills,omitempty" xml:"endTimeInMills,omitempty"`
	// 模板编码
	ProcessCode *string `json:"processCode,omitempty" xml:"processCode,omitempty"`
	// 应用编码
	AppUuid *string `json:"appUuid,omitempty" xml:"appUuid,omitempty"`
}

func (s QueryAllProcessInstancesRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryAllProcessInstancesRequest) GoString() string {
	return s.String()
}

func (s *QueryAllProcessInstancesRequest) SetNextToken(v string) *QueryAllProcessInstancesRequest {
	s.NextToken = &v
	return s
}

func (s *QueryAllProcessInstancesRequest) SetMaxResults(v int64) *QueryAllProcessInstancesRequest {
	s.MaxResults = &v
	return s
}

func (s *QueryAllProcessInstancesRequest) SetStartTimeInMills(v int64) *QueryAllProcessInstancesRequest {
	s.StartTimeInMills = &v
	return s
}

func (s *QueryAllProcessInstancesRequest) SetEndTimeInMills(v int64) *QueryAllProcessInstancesRequest {
	s.EndTimeInMills = &v
	return s
}

func (s *QueryAllProcessInstancesRequest) SetProcessCode(v string) *QueryAllProcessInstancesRequest {
	s.ProcessCode = &v
	return s
}

func (s *QueryAllProcessInstancesRequest) SetAppUuid(v string) *QueryAllProcessInstancesRequest {
	s.AppUuid = &v
	return s
}

type QueryAllProcessInstancesResponseBody struct {
	// result
	Result *QueryAllProcessInstancesResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s QueryAllProcessInstancesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryAllProcessInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *QueryAllProcessInstancesResponseBody) SetResult(v *QueryAllProcessInstancesResponseBodyResult) *QueryAllProcessInstancesResponseBody {
	s.Result = v
	return s
}

type QueryAllProcessInstancesResponseBodyResult struct {
	// 下次获取数据的游标
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// 是否有更多数据
	HasMore *bool `json:"hasMore,omitempty" xml:"hasMore,omitempty"`
	// 总数
	MaxResults *int64                                            `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	List       []*QueryAllProcessInstancesResponseBodyResultList `json:"list,omitempty" xml:"list,omitempty" type:"Repeated"`
}

func (s QueryAllProcessInstancesResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s QueryAllProcessInstancesResponseBodyResult) GoString() string {
	return s.String()
}

func (s *QueryAllProcessInstancesResponseBodyResult) SetNextToken(v string) *QueryAllProcessInstancesResponseBodyResult {
	s.NextToken = &v
	return s
}

func (s *QueryAllProcessInstancesResponseBodyResult) SetHasMore(v bool) *QueryAllProcessInstancesResponseBodyResult {
	s.HasMore = &v
	return s
}

func (s *QueryAllProcessInstancesResponseBodyResult) SetMaxResults(v int64) *QueryAllProcessInstancesResponseBodyResult {
	s.MaxResults = &v
	return s
}

func (s *QueryAllProcessInstancesResponseBodyResult) SetList(v []*QueryAllProcessInstancesResponseBodyResultList) *QueryAllProcessInstancesResponseBodyResult {
	s.List = v
	return s
}

type QueryAllProcessInstancesResponseBodyResultList struct {
	// 流程实例ID
	ProcessInstanceId *string `json:"processInstanceId,omitempty" xml:"processInstanceId,omitempty"`
	// 主单实例Id
	MainProcessInstanceId *string `json:"mainProcessInstanceId,omitempty" xml:"mainProcessInstanceId,omitempty"`
	// 审批结束时间
	FinishTime *int64 `json:"finishTime,omitempty" xml:"finishTime,omitempty"`
	// 附属单信息
	AttachedProcessInstanceIds *string `json:"attachedProcessInstanceIds,omitempty" xml:"attachedProcessInstanceIds,omitempty"`
	// 审批单编号
	BusinessId *string `json:"businessId,omitempty" xml:"businessId,omitempty"`
	// 审批单标题
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
	// 发起人部门id
	OriginatorDeptId *string `json:"originatorDeptId,omitempty" xml:"originatorDeptId,omitempty"`
	// 审批结果
	Result *string `json:"result,omitempty" xml:"result,omitempty"`
	// 审批单创建时间
	CreateTime *int64 `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// 发起者userId
	OriginatorUserid *string `json:"originatorUserid,omitempty" xml:"originatorUserid,omitempty"`
	// 审批单状态
	Status              *string                                                              `json:"status,omitempty" xml:"status,omitempty"`
	FormComponentValues []*QueryAllProcessInstancesResponseBodyResultListFormComponentValues `json:"formComponentValues,omitempty" xml:"formComponentValues,omitempty" type:"Repeated"`
}

func (s QueryAllProcessInstancesResponseBodyResultList) String() string {
	return tea.Prettify(s)
}

func (s QueryAllProcessInstancesResponseBodyResultList) GoString() string {
	return s.String()
}

func (s *QueryAllProcessInstancesResponseBodyResultList) SetProcessInstanceId(v string) *QueryAllProcessInstancesResponseBodyResultList {
	s.ProcessInstanceId = &v
	return s
}

func (s *QueryAllProcessInstancesResponseBodyResultList) SetMainProcessInstanceId(v string) *QueryAllProcessInstancesResponseBodyResultList {
	s.MainProcessInstanceId = &v
	return s
}

func (s *QueryAllProcessInstancesResponseBodyResultList) SetFinishTime(v int64) *QueryAllProcessInstancesResponseBodyResultList {
	s.FinishTime = &v
	return s
}

func (s *QueryAllProcessInstancesResponseBodyResultList) SetAttachedProcessInstanceIds(v string) *QueryAllProcessInstancesResponseBodyResultList {
	s.AttachedProcessInstanceIds = &v
	return s
}

func (s *QueryAllProcessInstancesResponseBodyResultList) SetBusinessId(v string) *QueryAllProcessInstancesResponseBodyResultList {
	s.BusinessId = &v
	return s
}

func (s *QueryAllProcessInstancesResponseBodyResultList) SetTitle(v string) *QueryAllProcessInstancesResponseBodyResultList {
	s.Title = &v
	return s
}

func (s *QueryAllProcessInstancesResponseBodyResultList) SetOriginatorDeptId(v string) *QueryAllProcessInstancesResponseBodyResultList {
	s.OriginatorDeptId = &v
	return s
}

func (s *QueryAllProcessInstancesResponseBodyResultList) SetResult(v string) *QueryAllProcessInstancesResponseBodyResultList {
	s.Result = &v
	return s
}

func (s *QueryAllProcessInstancesResponseBodyResultList) SetCreateTime(v int64) *QueryAllProcessInstancesResponseBodyResultList {
	s.CreateTime = &v
	return s
}

func (s *QueryAllProcessInstancesResponseBodyResultList) SetOriginatorUserid(v string) *QueryAllProcessInstancesResponseBodyResultList {
	s.OriginatorUserid = &v
	return s
}

func (s *QueryAllProcessInstancesResponseBodyResultList) SetStatus(v string) *QueryAllProcessInstancesResponseBodyResultList {
	s.Status = &v
	return s
}

func (s *QueryAllProcessInstancesResponseBodyResultList) SetFormComponentValues(v []*QueryAllProcessInstancesResponseBodyResultListFormComponentValues) *QueryAllProcessInstancesResponseBodyResultList {
	s.FormComponentValues = v
	return s
}

type QueryAllProcessInstancesResponseBodyResultListFormComponentValues struct {
	// 控件名称
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 控件id
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// 控件数据
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
	// 控件扩展数据
	ExtValue *string `json:"extValue,omitempty" xml:"extValue,omitempty"`
}

func (s QueryAllProcessInstancesResponseBodyResultListFormComponentValues) String() string {
	return tea.Prettify(s)
}

func (s QueryAllProcessInstancesResponseBodyResultListFormComponentValues) GoString() string {
	return s.String()
}

func (s *QueryAllProcessInstancesResponseBodyResultListFormComponentValues) SetName(v string) *QueryAllProcessInstancesResponseBodyResultListFormComponentValues {
	s.Name = &v
	return s
}

func (s *QueryAllProcessInstancesResponseBodyResultListFormComponentValues) SetId(v string) *QueryAllProcessInstancesResponseBodyResultListFormComponentValues {
	s.Id = &v
	return s
}

func (s *QueryAllProcessInstancesResponseBodyResultListFormComponentValues) SetValue(v string) *QueryAllProcessInstancesResponseBodyResultListFormComponentValues {
	s.Value = &v
	return s
}

func (s *QueryAllProcessInstancesResponseBodyResultListFormComponentValues) SetExtValue(v string) *QueryAllProcessInstancesResponseBodyResultListFormComponentValues {
	s.ExtValue = &v
	return s
}

type QueryAllProcessInstancesResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryAllProcessInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryAllProcessInstancesResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryAllProcessInstancesResponse) GoString() string {
	return s.String()
}

func (s *QueryAllProcessInstancesResponse) SetHeaders(v map[string]*string) *QueryAllProcessInstancesResponse {
	s.Headers = v
	return s
}

func (s *QueryAllProcessInstancesResponse) SetBody(v *QueryAllProcessInstancesResponseBody) *QueryAllProcessInstancesResponse {
	s.Body = v
	return s
}

type QueryAllFormInstancesHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryAllFormInstancesHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryAllFormInstancesHeaders) GoString() string {
	return s.String()
}

func (s *QueryAllFormInstancesHeaders) SetCommonHeaders(v map[string]*string) *QueryAllFormInstancesHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryAllFormInstancesHeaders) SetXAcsDingtalkAccessToken(v string) *QueryAllFormInstancesHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryAllFormInstancesRequest struct {
	// 分页游标，第一次调用传空或者null
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// 翻页size
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// 应用搭建id
	AppUuid *string `json:"appUuid,omitempty" xml:"appUuid,omitempty"`
	// 表单模板id
	FormCode *string `json:"formCode,omitempty" xml:"formCode,omitempty"`
}

func (s QueryAllFormInstancesRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryAllFormInstancesRequest) GoString() string {
	return s.String()
}

func (s *QueryAllFormInstancesRequest) SetNextToken(v string) *QueryAllFormInstancesRequest {
	s.NextToken = &v
	return s
}

func (s *QueryAllFormInstancesRequest) SetMaxResults(v int32) *QueryAllFormInstancesRequest {
	s.MaxResults = &v
	return s
}

func (s *QueryAllFormInstancesRequest) SetAppUuid(v string) *QueryAllFormInstancesRequest {
	s.AppUuid = &v
	return s
}

func (s *QueryAllFormInstancesRequest) SetFormCode(v string) *QueryAllFormInstancesRequest {
	s.FormCode = &v
	return s
}

type QueryAllFormInstancesResponseBody struct {
	// 分页结果
	Result *QueryAllFormInstancesResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s QueryAllFormInstancesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryAllFormInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *QueryAllFormInstancesResponseBody) SetResult(v *QueryAllFormInstancesResponseBodyResult) *QueryAllFormInstancesResponseBody {
	s.Result = v
	return s
}

type QueryAllFormInstancesResponseBodyResult struct {
	// 下一页的游标
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// 是否有更多数据
	HasMore *bool `json:"hasMore,omitempty" xml:"hasMore,omitempty"`
	// 分页大小
	MaxResults *int64 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// 表单列表
	Values []*QueryAllFormInstancesResponseBodyResultValues `json:"values,omitempty" xml:"values,omitempty" type:"Repeated"`
}

func (s QueryAllFormInstancesResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s QueryAllFormInstancesResponseBodyResult) GoString() string {
	return s.String()
}

func (s *QueryAllFormInstancesResponseBodyResult) SetNextToken(v string) *QueryAllFormInstancesResponseBodyResult {
	s.NextToken = &v
	return s
}

func (s *QueryAllFormInstancesResponseBodyResult) SetHasMore(v bool) *QueryAllFormInstancesResponseBodyResult {
	s.HasMore = &v
	return s
}

func (s *QueryAllFormInstancesResponseBodyResult) SetMaxResults(v int64) *QueryAllFormInstancesResponseBodyResult {
	s.MaxResults = &v
	return s
}

func (s *QueryAllFormInstancesResponseBodyResult) SetValues(v []*QueryAllFormInstancesResponseBodyResultValues) *QueryAllFormInstancesResponseBodyResult {
	s.Values = v
	return s
}

type QueryAllFormInstancesResponseBodyResultValues struct {
	// 表单实例id
	FormInstanceId *string `json:"formInstanceId,omitempty" xml:"formInstanceId,omitempty"`
	// 应用搭建id
	AppUuid *string `json:"appUuid,omitempty" xml:"appUuid,omitempty"`
	// 表单模板code
	FormCode *string `json:"formCode,omitempty" xml:"formCode,omitempty"`
	// 标题
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
	// 创建人
	Creator *string `json:"creator,omitempty" xml:"creator,omitempty"`
	// 修改人
	Modifier *string `json:"modifier,omitempty" xml:"modifier,omitempty"`
	// 创建时间
	CreateTimestamp *int64 `json:"createTimestamp,omitempty" xml:"createTimestamp,omitempty"`
	// 修改时间
	ModifyTimestamp *int64 `json:"modifyTimestamp,omitempty" xml:"modifyTimestamp,omitempty"`
	// 外部实例编码
	OutInstanceId *string `json:"outInstanceId,omitempty" xml:"outInstanceId,omitempty"`
	// 外部业务编码
	OutBizCode *string `json:"outBizCode,omitempty" xml:"outBizCode,omitempty"`
	// 扩展信息
	Attributes map[string]interface{} `json:"attributes,omitempty" xml:"attributes,omitempty"`
	// 表单实例数据
	FormInstDataList []*QueryAllFormInstancesResponseBodyResultValuesFormInstDataList `json:"formInstDataList,omitempty" xml:"formInstDataList,omitempty" type:"Repeated"`
}

func (s QueryAllFormInstancesResponseBodyResultValues) String() string {
	return tea.Prettify(s)
}

func (s QueryAllFormInstancesResponseBodyResultValues) GoString() string {
	return s.String()
}

func (s *QueryAllFormInstancesResponseBodyResultValues) SetFormInstanceId(v string) *QueryAllFormInstancesResponseBodyResultValues {
	s.FormInstanceId = &v
	return s
}

func (s *QueryAllFormInstancesResponseBodyResultValues) SetAppUuid(v string) *QueryAllFormInstancesResponseBodyResultValues {
	s.AppUuid = &v
	return s
}

func (s *QueryAllFormInstancesResponseBodyResultValues) SetFormCode(v string) *QueryAllFormInstancesResponseBodyResultValues {
	s.FormCode = &v
	return s
}

func (s *QueryAllFormInstancesResponseBodyResultValues) SetTitle(v string) *QueryAllFormInstancesResponseBodyResultValues {
	s.Title = &v
	return s
}

func (s *QueryAllFormInstancesResponseBodyResultValues) SetCreator(v string) *QueryAllFormInstancesResponseBodyResultValues {
	s.Creator = &v
	return s
}

func (s *QueryAllFormInstancesResponseBodyResultValues) SetModifier(v string) *QueryAllFormInstancesResponseBodyResultValues {
	s.Modifier = &v
	return s
}

func (s *QueryAllFormInstancesResponseBodyResultValues) SetCreateTimestamp(v int64) *QueryAllFormInstancesResponseBodyResultValues {
	s.CreateTimestamp = &v
	return s
}

func (s *QueryAllFormInstancesResponseBodyResultValues) SetModifyTimestamp(v int64) *QueryAllFormInstancesResponseBodyResultValues {
	s.ModifyTimestamp = &v
	return s
}

func (s *QueryAllFormInstancesResponseBodyResultValues) SetOutInstanceId(v string) *QueryAllFormInstancesResponseBodyResultValues {
	s.OutInstanceId = &v
	return s
}

func (s *QueryAllFormInstancesResponseBodyResultValues) SetOutBizCode(v string) *QueryAllFormInstancesResponseBodyResultValues {
	s.OutBizCode = &v
	return s
}

func (s *QueryAllFormInstancesResponseBodyResultValues) SetAttributes(v map[string]interface{}) *QueryAllFormInstancesResponseBodyResultValues {
	s.Attributes = v
	return s
}

func (s *QueryAllFormInstancesResponseBodyResultValues) SetFormInstDataList(v []*QueryAllFormInstancesResponseBodyResultValuesFormInstDataList) *QueryAllFormInstancesResponseBodyResultValues {
	s.FormInstDataList = v
	return s
}

type QueryAllFormInstancesResponseBodyResultValuesFormInstDataList struct {
	// 控件类型
	ComponentType *string `json:"componentType,omitempty" xml:"componentType,omitempty"`
	// 控件别名
	BizAlias *string `json:"bizAlias,omitempty" xml:"bizAlias,omitempty"`
	// 表单控件扩展数据
	ExtendValue *string `json:"extendValue,omitempty" xml:"extendValue,omitempty"`
	// 控件名称
	Label *string `json:"label,omitempty" xml:"label,omitempty"`
	// 控件填写的数据
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
	// 控件唯一id
	Key *string `json:"key,omitempty" xml:"key,omitempty"`
}

func (s QueryAllFormInstancesResponseBodyResultValuesFormInstDataList) String() string {
	return tea.Prettify(s)
}

func (s QueryAllFormInstancesResponseBodyResultValuesFormInstDataList) GoString() string {
	return s.String()
}

func (s *QueryAllFormInstancesResponseBodyResultValuesFormInstDataList) SetComponentType(v string) *QueryAllFormInstancesResponseBodyResultValuesFormInstDataList {
	s.ComponentType = &v
	return s
}

func (s *QueryAllFormInstancesResponseBodyResultValuesFormInstDataList) SetBizAlias(v string) *QueryAllFormInstancesResponseBodyResultValuesFormInstDataList {
	s.BizAlias = &v
	return s
}

func (s *QueryAllFormInstancesResponseBodyResultValuesFormInstDataList) SetExtendValue(v string) *QueryAllFormInstancesResponseBodyResultValuesFormInstDataList {
	s.ExtendValue = &v
	return s
}

func (s *QueryAllFormInstancesResponseBodyResultValuesFormInstDataList) SetLabel(v string) *QueryAllFormInstancesResponseBodyResultValuesFormInstDataList {
	s.Label = &v
	return s
}

func (s *QueryAllFormInstancesResponseBodyResultValuesFormInstDataList) SetValue(v string) *QueryAllFormInstancesResponseBodyResultValuesFormInstDataList {
	s.Value = &v
	return s
}

func (s *QueryAllFormInstancesResponseBodyResultValuesFormInstDataList) SetKey(v string) *QueryAllFormInstancesResponseBodyResultValuesFormInstDataList {
	s.Key = &v
	return s
}

type QueryAllFormInstancesResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryAllFormInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryAllFormInstancesResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryAllFormInstancesResponse) GoString() string {
	return s.String()
}

func (s *QueryAllFormInstancesResponse) SetHeaders(v map[string]*string) *QueryAllFormInstancesResponse {
	s.Headers = v
	return s
}

func (s *QueryAllFormInstancesResponse) SetBody(v *QueryAllFormInstancesResponseBody) *QueryAllFormInstancesResponse {
	s.Body = v
	return s
}

type QueryFormByBizTypeHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryFormByBizTypeHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryFormByBizTypeHeaders) GoString() string {
	return s.String()
}

func (s *QueryFormByBizTypeHeaders) SetCommonHeaders(v map[string]*string) *QueryFormByBizTypeHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryFormByBizTypeHeaders) SetXAcsDingtalkAccessToken(v string) *QueryFormByBizTypeHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryFormByBizTypeRequest struct {
	// 应用搭建id
	AppUuid *string `json:"appUuid,omitempty" xml:"appUuid,omitempty"`
	// 表单业务标识
	BizTypes []*string `json:"bizTypes,omitempty" xml:"bizTypes,omitempty" type:"Repeated"`
}

func (s QueryFormByBizTypeRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryFormByBizTypeRequest) GoString() string {
	return s.String()
}

func (s *QueryFormByBizTypeRequest) SetAppUuid(v string) *QueryFormByBizTypeRequest {
	s.AppUuid = &v
	return s
}

func (s *QueryFormByBizTypeRequest) SetBizTypes(v []*string) *QueryFormByBizTypeRequest {
	s.BizTypes = v
	return s
}

type QueryFormByBizTypeResponseBody struct {
	// 模板列表
	Result []*QueryFormByBizTypeResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
}

func (s QueryFormByBizTypeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryFormByBizTypeResponseBody) GoString() string {
	return s.String()
}

func (s *QueryFormByBizTypeResponseBody) SetResult(v []*QueryFormByBizTypeResponseBodyResult) *QueryFormByBizTypeResponseBody {
	s.Result = v
	return s
}

type QueryFormByBizTypeResponseBodyResult struct {
	// 创建人
	Creator *string `json:"creator,omitempty" xml:"creator,omitempty"`
	// 应用搭建id
	AppUuid *string `json:"appUuid,omitempty" xml:"appUuid,omitempty"`
	// 模板code
	FormCode *string `json:"formCode,omitempty" xml:"formCode,omitempty"`
	// 表单uuid
	FormUuid *string `json:"formUuid,omitempty" xml:"formUuid,omitempty"`
	// 模板名称
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 模板描述
	Memo *string `json:"memo,omitempty" xml:"memo,omitempty"`
	// 数据归属id
	OwnerId *string `json:"ownerId,omitempty" xml:"ownerId,omitempty"`
	// 表单类型，0为流程表单，1为数据表单
	AppType *int32 `json:"appType,omitempty" xml:"appType,omitempty"`
	// 业务标识
	BizType *string `json:"bizType,omitempty" xml:"bizType,omitempty"`
	// 模板状态
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// 创建时间
	CreateTime *int64 `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// 修改时间
	ModifedTime *int64 `json:"modifedTime,omitempty" xml:"modifedTime,omitempty"`
	// 表单控件描述
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
}

func (s QueryFormByBizTypeResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s QueryFormByBizTypeResponseBodyResult) GoString() string {
	return s.String()
}

func (s *QueryFormByBizTypeResponseBodyResult) SetCreator(v string) *QueryFormByBizTypeResponseBodyResult {
	s.Creator = &v
	return s
}

func (s *QueryFormByBizTypeResponseBodyResult) SetAppUuid(v string) *QueryFormByBizTypeResponseBodyResult {
	s.AppUuid = &v
	return s
}

func (s *QueryFormByBizTypeResponseBodyResult) SetFormCode(v string) *QueryFormByBizTypeResponseBodyResult {
	s.FormCode = &v
	return s
}

func (s *QueryFormByBizTypeResponseBodyResult) SetFormUuid(v string) *QueryFormByBizTypeResponseBodyResult {
	s.FormUuid = &v
	return s
}

func (s *QueryFormByBizTypeResponseBodyResult) SetName(v string) *QueryFormByBizTypeResponseBodyResult {
	s.Name = &v
	return s
}

func (s *QueryFormByBizTypeResponseBodyResult) SetMemo(v string) *QueryFormByBizTypeResponseBodyResult {
	s.Memo = &v
	return s
}

func (s *QueryFormByBizTypeResponseBodyResult) SetOwnerId(v string) *QueryFormByBizTypeResponseBodyResult {
	s.OwnerId = &v
	return s
}

func (s *QueryFormByBizTypeResponseBodyResult) SetAppType(v int32) *QueryFormByBizTypeResponseBodyResult {
	s.AppType = &v
	return s
}

func (s *QueryFormByBizTypeResponseBodyResult) SetBizType(v string) *QueryFormByBizTypeResponseBodyResult {
	s.BizType = &v
	return s
}

func (s *QueryFormByBizTypeResponseBodyResult) SetStatus(v string) *QueryFormByBizTypeResponseBodyResult {
	s.Status = &v
	return s
}

func (s *QueryFormByBizTypeResponseBodyResult) SetCreateTime(v int64) *QueryFormByBizTypeResponseBodyResult {
	s.CreateTime = &v
	return s
}

func (s *QueryFormByBizTypeResponseBodyResult) SetModifedTime(v int64) *QueryFormByBizTypeResponseBodyResult {
	s.ModifedTime = &v
	return s
}

func (s *QueryFormByBizTypeResponseBodyResult) SetContent(v string) *QueryFormByBizTypeResponseBodyResult {
	s.Content = &v
	return s
}

type QueryFormByBizTypeResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryFormByBizTypeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryFormByBizTypeResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryFormByBizTypeResponse) GoString() string {
	return s.String()
}

func (s *QueryFormByBizTypeResponse) SetHeaders(v map[string]*string) *QueryFormByBizTypeResponse {
	s.Headers = v
	return s
}

func (s *QueryFormByBizTypeResponse) SetBody(v *QueryFormByBizTypeResponseBody) *QueryFormByBizTypeResponse {
	s.Body = v
	return s
}

type FormCreateHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s FormCreateHeaders) String() string {
	return tea.Prettify(s)
}

func (s FormCreateHeaders) GoString() string {
	return s.String()
}

func (s *FormCreateHeaders) SetCommonHeaders(v map[string]*string) *FormCreateHeaders {
	s.CommonHeaders = v
	return s
}

func (s *FormCreateHeaders) SetXAcsDingtalkAccessToken(v string) *FormCreateHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type FormCreateRequest struct {
	DingCorpId         *string `json:"dingCorpId,omitempty" xml:"dingCorpId,omitempty"`
	DingOrgId          *int64  `json:"dingOrgId,omitempty" xml:"dingOrgId,omitempty"`
	DingIsvOrgId       *int64  `json:"dingIsvOrgId,omitempty" xml:"dingIsvOrgId,omitempty"`
	DingSuiteKey       *string `json:"dingSuiteKey,omitempty" xml:"dingSuiteKey,omitempty"`
	DingTokenGrantType *int64  `json:"dingTokenGrantType,omitempty" xml:"dingTokenGrantType,omitempty"`
	RequestId          *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ProcessCode        *string `json:"processCode,omitempty" xml:"processCode,omitempty"`
	// 表单模板名称
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 表单模板描述
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// 表单控件列表
	FormComponents []*FormCreateRequestFormComponents `json:"formComponents,omitempty" xml:"formComponents,omitempty" type:"Repeated"`
}

func (s FormCreateRequest) String() string {
	return tea.Prettify(s)
}

func (s FormCreateRequest) GoString() string {
	return s.String()
}

func (s *FormCreateRequest) SetDingCorpId(v string) *FormCreateRequest {
	s.DingCorpId = &v
	return s
}

func (s *FormCreateRequest) SetDingOrgId(v int64) *FormCreateRequest {
	s.DingOrgId = &v
	return s
}

func (s *FormCreateRequest) SetDingIsvOrgId(v int64) *FormCreateRequest {
	s.DingIsvOrgId = &v
	return s
}

func (s *FormCreateRequest) SetDingSuiteKey(v string) *FormCreateRequest {
	s.DingSuiteKey = &v
	return s
}

func (s *FormCreateRequest) SetDingTokenGrantType(v int64) *FormCreateRequest {
	s.DingTokenGrantType = &v
	return s
}

func (s *FormCreateRequest) SetRequestId(v string) *FormCreateRequest {
	s.RequestId = &v
	return s
}

func (s *FormCreateRequest) SetProcessCode(v string) *FormCreateRequest {
	s.ProcessCode = &v
	return s
}

func (s *FormCreateRequest) SetName(v string) *FormCreateRequest {
	s.Name = &v
	return s
}

func (s *FormCreateRequest) SetDescription(v string) *FormCreateRequest {
	s.Description = &v
	return s
}

func (s *FormCreateRequest) SetFormComponents(v []*FormCreateRequestFormComponents) *FormCreateRequest {
	s.FormComponents = v
	return s
}

type FormCreateRequestFormComponents struct {
	// 控件类型
	ComponentType *string `json:"componentType,omitempty" xml:"componentType,omitempty"`
	// 控件属性
	Props *FormCreateRequestFormComponentsProps `json:"props,omitempty" xml:"props,omitempty" type:"Struct"`
	// 子控件列表
	Children []*FormCreateRequestFormComponentsChildren `json:"children,omitempty" xml:"children,omitempty" type:"Repeated"`
}

func (s FormCreateRequestFormComponents) String() string {
	return tea.Prettify(s)
}

func (s FormCreateRequestFormComponents) GoString() string {
	return s.String()
}

func (s *FormCreateRequestFormComponents) SetComponentType(v string) *FormCreateRequestFormComponents {
	s.ComponentType = &v
	return s
}

func (s *FormCreateRequestFormComponents) SetProps(v *FormCreateRequestFormComponentsProps) *FormCreateRequestFormComponents {
	s.Props = v
	return s
}

func (s *FormCreateRequestFormComponents) SetChildren(v []*FormCreateRequestFormComponentsChildren) *FormCreateRequestFormComponents {
	s.Children = v
	return s
}

type FormCreateRequestFormComponentsProps struct {
	// 控件表单内唯一id
	ComponentId *string `json:"componentId,omitempty" xml:"componentId,omitempty"`
	// 控件标题
	Label *string `json:"label,omitempty" xml:"label,omitempty"`
	// 套件中控件是否可设置为分条件字段
	AsyncCondition *bool `json:"asyncCondition,omitempty" xml:"asyncCondition,omitempty"`
	// 是否必填
	Required *bool `json:"required,omitempty" xml:"required,omitempty"`
	// 说明文字控件内容
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// 时间格式
	Format *string `json:"format,omitempty" xml:"format,omitempty"`
	// 金额控件是否需要大写，1不需要大写，其他需要大写
	Upper *string `json:"upper,omitempty" xml:"upper,omitempty"`
	// 时间单位（天、小时）
	Unit *string `json:"unit,omitempty" xml:"unit,omitempty"`
	// 输入提示
	Placeholder *string `json:"placeholder,omitempty" xml:"placeholder,omitempty"`
	// 业务别名
	BizAlias *string `json:"bizAlias,omitempty" xml:"bizAlias,omitempty"`
	// 套件的业务标识
	BizType *string `json:"bizType,omitempty" xml:"bizType,omitempty"`
	// 是否自动计算时长
	Duration *bool `json:"duration,omitempty" xml:"duration,omitempty"`
	// 联系人控件是否支持多选，1多选，0单选
	Choice *string `json:"choice,omitempty" xml:"choice,omitempty"`
	// 是否不可编辑
	Disabled *bool `json:"disabled,omitempty" xml:"disabled,omitempty"`
	// 文字提示控件显示方式:top|middle|bottom
	Align *string `json:"align,omitempty" xml:"align,omitempty"`
	// 是否隐藏字段
	Invisible *bool `json:"invisible,omitempty" xml:"invisible,omitempty"`
	// 说明文字控件链接地址
	Link *string `json:"link,omitempty" xml:"link,omitempty"`
	// 明细打印方式false横向，true纵向
	VerticalPrint *bool `json:"verticalPrint,omitempty" xml:"verticalPrint,omitempty"`
	// 公式
	Formula *string `json:"formula,omitempty" xml:"formula,omitempty"`
	// 自定义控件渲染标识
	CommonBizType *string `json:"commonBizType,omitempty" xml:"commonBizType,omitempty"`
	// 单选多选控件选项列表
	Options []*FormCreateRequestFormComponentsPropsOptions `json:"options,omitempty" xml:"options,omitempty" type:"Repeated"`
	// 字段是否可打印，1打印，0不打印，默认打印
	Print *string `json:"print,omitempty" xml:"print,omitempty"`
	// 明细控件数据汇总统计
	StatField []*FormCreateRequestFormComponentsPropsStatField `json:"statField,omitempty" xml:"statField,omitempty" type:"Repeated"`
	// 关联审批单、关联表单数据源配置
	DataSource *FormCreateRequestFormComponentsPropsDataSource `json:"dataSource,omitempty" xml:"dataSource,omitempty" type:"Struct"`
	// 关联显示字段
	Fields []*FormCreateRequestFormComponentsPropsFields `json:"fields,omitempty" xml:"fields,omitempty" type:"Repeated"`
	// 地址控件模式city省市,district省市区,street省市区街道
	AddressModel *string `json:"addressModel,omitempty" xml:"addressModel,omitempty"`
	// 部门控件是否可多选
	Multiple *bool `json:"multiple,omitempty" xml:"multiple,omitempty"`
	// 评分控件上限
	Limit *int32 `json:"limit,omitempty" xml:"limit,omitempty"`
	// 关联审批单控件限定模板列表
	AvailableTemplates []*FormCreateRequestFormComponentsPropsAvailableTemplates `json:"availableTemplates,omitempty" xml:"availableTemplates,omitempty" type:"Repeated"`
	// 明细填写方式，table（表格）、list（列表）
	TableViewMode *string `json:"tableViewMode,omitempty" xml:"tableViewMode,omitempty"`
}

func (s FormCreateRequestFormComponentsProps) String() string {
	return tea.Prettify(s)
}

func (s FormCreateRequestFormComponentsProps) GoString() string {
	return s.String()
}

func (s *FormCreateRequestFormComponentsProps) SetComponentId(v string) *FormCreateRequestFormComponentsProps {
	s.ComponentId = &v
	return s
}

func (s *FormCreateRequestFormComponentsProps) SetLabel(v string) *FormCreateRequestFormComponentsProps {
	s.Label = &v
	return s
}

func (s *FormCreateRequestFormComponentsProps) SetAsyncCondition(v bool) *FormCreateRequestFormComponentsProps {
	s.AsyncCondition = &v
	return s
}

func (s *FormCreateRequestFormComponentsProps) SetRequired(v bool) *FormCreateRequestFormComponentsProps {
	s.Required = &v
	return s
}

func (s *FormCreateRequestFormComponentsProps) SetContent(v string) *FormCreateRequestFormComponentsProps {
	s.Content = &v
	return s
}

func (s *FormCreateRequestFormComponentsProps) SetFormat(v string) *FormCreateRequestFormComponentsProps {
	s.Format = &v
	return s
}

func (s *FormCreateRequestFormComponentsProps) SetUpper(v string) *FormCreateRequestFormComponentsProps {
	s.Upper = &v
	return s
}

func (s *FormCreateRequestFormComponentsProps) SetUnit(v string) *FormCreateRequestFormComponentsProps {
	s.Unit = &v
	return s
}

func (s *FormCreateRequestFormComponentsProps) SetPlaceholder(v string) *FormCreateRequestFormComponentsProps {
	s.Placeholder = &v
	return s
}

func (s *FormCreateRequestFormComponentsProps) SetBizAlias(v string) *FormCreateRequestFormComponentsProps {
	s.BizAlias = &v
	return s
}

func (s *FormCreateRequestFormComponentsProps) SetBizType(v string) *FormCreateRequestFormComponentsProps {
	s.BizType = &v
	return s
}

func (s *FormCreateRequestFormComponentsProps) SetDuration(v bool) *FormCreateRequestFormComponentsProps {
	s.Duration = &v
	return s
}

func (s *FormCreateRequestFormComponentsProps) SetChoice(v string) *FormCreateRequestFormComponentsProps {
	s.Choice = &v
	return s
}

func (s *FormCreateRequestFormComponentsProps) SetDisabled(v bool) *FormCreateRequestFormComponentsProps {
	s.Disabled = &v
	return s
}

func (s *FormCreateRequestFormComponentsProps) SetAlign(v string) *FormCreateRequestFormComponentsProps {
	s.Align = &v
	return s
}

func (s *FormCreateRequestFormComponentsProps) SetInvisible(v bool) *FormCreateRequestFormComponentsProps {
	s.Invisible = &v
	return s
}

func (s *FormCreateRequestFormComponentsProps) SetLink(v string) *FormCreateRequestFormComponentsProps {
	s.Link = &v
	return s
}

func (s *FormCreateRequestFormComponentsProps) SetVerticalPrint(v bool) *FormCreateRequestFormComponentsProps {
	s.VerticalPrint = &v
	return s
}

func (s *FormCreateRequestFormComponentsProps) SetFormula(v string) *FormCreateRequestFormComponentsProps {
	s.Formula = &v
	return s
}

func (s *FormCreateRequestFormComponentsProps) SetCommonBizType(v string) *FormCreateRequestFormComponentsProps {
	s.CommonBizType = &v
	return s
}

func (s *FormCreateRequestFormComponentsProps) SetOptions(v []*FormCreateRequestFormComponentsPropsOptions) *FormCreateRequestFormComponentsProps {
	s.Options = v
	return s
}

func (s *FormCreateRequestFormComponentsProps) SetPrint(v string) *FormCreateRequestFormComponentsProps {
	s.Print = &v
	return s
}

func (s *FormCreateRequestFormComponentsProps) SetStatField(v []*FormCreateRequestFormComponentsPropsStatField) *FormCreateRequestFormComponentsProps {
	s.StatField = v
	return s
}

func (s *FormCreateRequestFormComponentsProps) SetDataSource(v *FormCreateRequestFormComponentsPropsDataSource) *FormCreateRequestFormComponentsProps {
	s.DataSource = v
	return s
}

func (s *FormCreateRequestFormComponentsProps) SetFields(v []*FormCreateRequestFormComponentsPropsFields) *FormCreateRequestFormComponentsProps {
	s.Fields = v
	return s
}

func (s *FormCreateRequestFormComponentsProps) SetAddressModel(v string) *FormCreateRequestFormComponentsProps {
	s.AddressModel = &v
	return s
}

func (s *FormCreateRequestFormComponentsProps) SetMultiple(v bool) *FormCreateRequestFormComponentsProps {
	s.Multiple = &v
	return s
}

func (s *FormCreateRequestFormComponentsProps) SetLimit(v int32) *FormCreateRequestFormComponentsProps {
	s.Limit = &v
	return s
}

func (s *FormCreateRequestFormComponentsProps) SetAvailableTemplates(v []*FormCreateRequestFormComponentsPropsAvailableTemplates) *FormCreateRequestFormComponentsProps {
	s.AvailableTemplates = v
	return s
}

func (s *FormCreateRequestFormComponentsProps) SetTableViewMode(v string) *FormCreateRequestFormComponentsProps {
	s.TableViewMode = &v
	return s
}

type FormCreateRequestFormComponentsPropsOptions struct {
	// 选项的显示内容
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
	// 选项的唯一主键
	Key *string `json:"key,omitempty" xml:"key,omitempty"`
}

func (s FormCreateRequestFormComponentsPropsOptions) String() string {
	return tea.Prettify(s)
}

func (s FormCreateRequestFormComponentsPropsOptions) GoString() string {
	return s.String()
}

func (s *FormCreateRequestFormComponentsPropsOptions) SetValue(v string) *FormCreateRequestFormComponentsPropsOptions {
	s.Value = &v
	return s
}

func (s *FormCreateRequestFormComponentsPropsOptions) SetKey(v string) *FormCreateRequestFormComponentsPropsOptions {
	s.Key = &v
	return s
}

type FormCreateRequestFormComponentsPropsStatField struct {
	// 需要统计的明细控件内子控件id
	ComponentId *string `json:"componentId,omitempty" xml:"componentId,omitempty"`
	// 子控件标题
	Label *string `json:"label,omitempty" xml:"label,omitempty"`
	// 金额控件是否需要大写
	Upper     *bool   `json:"upper,omitempty" xml:"upper,omitempty"`
	PayEnable *string `json:"payEnable,omitempty" xml:"payEnable,omitempty"`
}

func (s FormCreateRequestFormComponentsPropsStatField) String() string {
	return tea.Prettify(s)
}

func (s FormCreateRequestFormComponentsPropsStatField) GoString() string {
	return s.String()
}

func (s *FormCreateRequestFormComponentsPropsStatField) SetComponentId(v string) *FormCreateRequestFormComponentsPropsStatField {
	s.ComponentId = &v
	return s
}

func (s *FormCreateRequestFormComponentsPropsStatField) SetLabel(v string) *FormCreateRequestFormComponentsPropsStatField {
	s.Label = &v
	return s
}

func (s *FormCreateRequestFormComponentsPropsStatField) SetUpper(v bool) *FormCreateRequestFormComponentsPropsStatField {
	s.Upper = &v
	return s
}

func (s *FormCreateRequestFormComponentsPropsStatField) SetPayEnable(v string) *FormCreateRequestFormComponentsPropsStatField {
	s.PayEnable = &v
	return s
}

type FormCreateRequestFormComponentsPropsDataSource struct {
	// 关联类型，form关联表单
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// 关联表单信息
	Target *FormCreateRequestFormComponentsPropsDataSourceTarget `json:"target,omitempty" xml:"target,omitempty" type:"Struct"`
}

func (s FormCreateRequestFormComponentsPropsDataSource) String() string {
	return tea.Prettify(s)
}

func (s FormCreateRequestFormComponentsPropsDataSource) GoString() string {
	return s.String()
}

func (s *FormCreateRequestFormComponentsPropsDataSource) SetType(v string) *FormCreateRequestFormComponentsPropsDataSource {
	s.Type = &v
	return s
}

func (s *FormCreateRequestFormComponentsPropsDataSource) SetTarget(v *FormCreateRequestFormComponentsPropsDataSourceTarget) *FormCreateRequestFormComponentsPropsDataSource {
	s.Target = v
	return s
}

type FormCreateRequestFormComponentsPropsDataSourceTarget struct {
	// 应用appUuid
	AppUuid *string `json:"appUuid,omitempty" xml:"appUuid,omitempty"`
	// 表单类型，0流程表单
	AppType *int32 `json:"appType,omitempty" xml:"appType,omitempty"`
	// 关联表单业务标识
	BizType *string `json:"bizType,omitempty" xml:"bizType,omitempty"`
	// 关联表单的formCode
	FormCode *string `json:"formCode,omitempty" xml:"formCode,omitempty"`
}

func (s FormCreateRequestFormComponentsPropsDataSourceTarget) String() string {
	return tea.Prettify(s)
}

func (s FormCreateRequestFormComponentsPropsDataSourceTarget) GoString() string {
	return s.String()
}

func (s *FormCreateRequestFormComponentsPropsDataSourceTarget) SetAppUuid(v string) *FormCreateRequestFormComponentsPropsDataSourceTarget {
	s.AppUuid = &v
	return s
}

func (s *FormCreateRequestFormComponentsPropsDataSourceTarget) SetAppType(v int32) *FormCreateRequestFormComponentsPropsDataSourceTarget {
	s.AppType = &v
	return s
}

func (s *FormCreateRequestFormComponentsPropsDataSourceTarget) SetBizType(v string) *FormCreateRequestFormComponentsPropsDataSourceTarget {
	s.BizType = &v
	return s
}

func (s *FormCreateRequestFormComponentsPropsDataSourceTarget) SetFormCode(v string) *FormCreateRequestFormComponentsPropsDataSourceTarget {
	s.FormCode = &v
	return s
}

type FormCreateRequestFormComponentsPropsFields struct {
	// 关联显示字段类型
	ComponentType *string `json:"componentType,omitempty" xml:"componentType,omitempty"`
	// 关联显示字段属性
	Props *FormCreateRequestFormComponentsPropsFieldsProps `json:"props,omitempty" xml:"props,omitempty" type:"Struct"`
}

func (s FormCreateRequestFormComponentsPropsFields) String() string {
	return tea.Prettify(s)
}

func (s FormCreateRequestFormComponentsPropsFields) GoString() string {
	return s.String()
}

func (s *FormCreateRequestFormComponentsPropsFields) SetComponentType(v string) *FormCreateRequestFormComponentsPropsFields {
	s.ComponentType = &v
	return s
}

func (s *FormCreateRequestFormComponentsPropsFields) SetProps(v *FormCreateRequestFormComponentsPropsFieldsProps) *FormCreateRequestFormComponentsPropsFields {
	s.Props = v
	return s
}

type FormCreateRequestFormComponentsPropsFieldsProps struct {
	// 表单中控件的唯一id
	ComponentId *string `json:"componentId,omitempty" xml:"componentId,omitempty"`
	// 控件标题
	Label *string `json:"label,omitempty" xml:"label,omitempty"`
	// label是否禁用修改
	LabelEditableFreeze *bool `json:"labelEditableFreeze,omitempty" xml:"labelEditableFreeze,omitempty"`
	// 必填
	Required *bool `json:"required,omitempty" xml:"required,omitempty"`
	// 必填是否可修改
	RequiredEditableFreeze *bool `json:"requiredEditableFreeze,omitempty" xml:"requiredEditableFreeze,omitempty"`
	// 是否可打印
	Print *string `json:"print,omitempty" xml:"print,omitempty"`
	// 说明文字控件内容
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// 时间格式
	Format *string `json:"format,omitempty" xml:"format,omitempty"`
	// 选项内容
	Options []*FormCreateRequestFormComponentsPropsFieldsPropsOptions `json:"options,omitempty" xml:"options,omitempty" type:"Repeated"`
	// 是否需要大写，1需要大写，0不需要
	Upper *string `json:"upper,omitempty" xml:"upper,omitempty"`
	// 时间单位（天、小时）
	Unit *string `json:"unit,omitempty" xml:"unit,omitempty"`
	// 输入提示
	Placeholder *string `json:"placeholder,omitempty" xml:"placeholder,omitempty"`
	// 业务别名
	BizAlias *string `json:"bizAlias,omitempty" xml:"bizAlias,omitempty"`
	// 套件的业务标识
	BizType *string `json:"bizType,omitempty" xml:"bizType,omitempty"`
	// 是否自动计算时长
	Duration *bool `json:"duration,omitempty" xml:"duration,omitempty"`
	// 联系人控件是否支持多选，1多选，0单选
	Choice *string `json:"choice,omitempty" xml:"choice,omitempty"`
	// 是否不可编辑
	Disabled *bool `json:"disabled,omitempty" xml:"disabled,omitempty"`
	// 文字提示控件显示方式（top|middle|bottom）
	Align *string `json:"align,omitempty" xml:"align,omitempty"`
}

func (s FormCreateRequestFormComponentsPropsFieldsProps) String() string {
	return tea.Prettify(s)
}

func (s FormCreateRequestFormComponentsPropsFieldsProps) GoString() string {
	return s.String()
}

func (s *FormCreateRequestFormComponentsPropsFieldsProps) SetComponentId(v string) *FormCreateRequestFormComponentsPropsFieldsProps {
	s.ComponentId = &v
	return s
}

func (s *FormCreateRequestFormComponentsPropsFieldsProps) SetLabel(v string) *FormCreateRequestFormComponentsPropsFieldsProps {
	s.Label = &v
	return s
}

func (s *FormCreateRequestFormComponentsPropsFieldsProps) SetLabelEditableFreeze(v bool) *FormCreateRequestFormComponentsPropsFieldsProps {
	s.LabelEditableFreeze = &v
	return s
}

func (s *FormCreateRequestFormComponentsPropsFieldsProps) SetRequired(v bool) *FormCreateRequestFormComponentsPropsFieldsProps {
	s.Required = &v
	return s
}

func (s *FormCreateRequestFormComponentsPropsFieldsProps) SetRequiredEditableFreeze(v bool) *FormCreateRequestFormComponentsPropsFieldsProps {
	s.RequiredEditableFreeze = &v
	return s
}

func (s *FormCreateRequestFormComponentsPropsFieldsProps) SetPrint(v string) *FormCreateRequestFormComponentsPropsFieldsProps {
	s.Print = &v
	return s
}

func (s *FormCreateRequestFormComponentsPropsFieldsProps) SetContent(v string) *FormCreateRequestFormComponentsPropsFieldsProps {
	s.Content = &v
	return s
}

func (s *FormCreateRequestFormComponentsPropsFieldsProps) SetFormat(v string) *FormCreateRequestFormComponentsPropsFieldsProps {
	s.Format = &v
	return s
}

func (s *FormCreateRequestFormComponentsPropsFieldsProps) SetOptions(v []*FormCreateRequestFormComponentsPropsFieldsPropsOptions) *FormCreateRequestFormComponentsPropsFieldsProps {
	s.Options = v
	return s
}

func (s *FormCreateRequestFormComponentsPropsFieldsProps) SetUpper(v string) *FormCreateRequestFormComponentsPropsFieldsProps {
	s.Upper = &v
	return s
}

func (s *FormCreateRequestFormComponentsPropsFieldsProps) SetUnit(v string) *FormCreateRequestFormComponentsPropsFieldsProps {
	s.Unit = &v
	return s
}

func (s *FormCreateRequestFormComponentsPropsFieldsProps) SetPlaceholder(v string) *FormCreateRequestFormComponentsPropsFieldsProps {
	s.Placeholder = &v
	return s
}

func (s *FormCreateRequestFormComponentsPropsFieldsProps) SetBizAlias(v string) *FormCreateRequestFormComponentsPropsFieldsProps {
	s.BizAlias = &v
	return s
}

func (s *FormCreateRequestFormComponentsPropsFieldsProps) SetBizType(v string) *FormCreateRequestFormComponentsPropsFieldsProps {
	s.BizType = &v
	return s
}

func (s *FormCreateRequestFormComponentsPropsFieldsProps) SetDuration(v bool) *FormCreateRequestFormComponentsPropsFieldsProps {
	s.Duration = &v
	return s
}

func (s *FormCreateRequestFormComponentsPropsFieldsProps) SetChoice(v string) *FormCreateRequestFormComponentsPropsFieldsProps {
	s.Choice = &v
	return s
}

func (s *FormCreateRequestFormComponentsPropsFieldsProps) SetDisabled(v bool) *FormCreateRequestFormComponentsPropsFieldsProps {
	s.Disabled = &v
	return s
}

func (s *FormCreateRequestFormComponentsPropsFieldsProps) SetAlign(v string) *FormCreateRequestFormComponentsPropsFieldsProps {
	s.Align = &v
	return s
}

type FormCreateRequestFormComponentsPropsFieldsPropsOptions struct {
	// 每一个选项的唯一键
	Key *string `json:"key,omitempty" xml:"key,omitempty"`
	// 每一个选项的值
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s FormCreateRequestFormComponentsPropsFieldsPropsOptions) String() string {
	return tea.Prettify(s)
}

func (s FormCreateRequestFormComponentsPropsFieldsPropsOptions) GoString() string {
	return s.String()
}

func (s *FormCreateRequestFormComponentsPropsFieldsPropsOptions) SetKey(v string) *FormCreateRequestFormComponentsPropsFieldsPropsOptions {
	s.Key = &v
	return s
}

func (s *FormCreateRequestFormComponentsPropsFieldsPropsOptions) SetValue(v string) *FormCreateRequestFormComponentsPropsFieldsPropsOptions {
	s.Value = &v
	return s
}

type FormCreateRequestFormComponentsPropsAvailableTemplates struct {
	// 表单名称
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 表单模板processCode
	ProcessCode *string `json:"processCode,omitempty" xml:"processCode,omitempty"`
}

func (s FormCreateRequestFormComponentsPropsAvailableTemplates) String() string {
	return tea.Prettify(s)
}

func (s FormCreateRequestFormComponentsPropsAvailableTemplates) GoString() string {
	return s.String()
}

func (s *FormCreateRequestFormComponentsPropsAvailableTemplates) SetName(v string) *FormCreateRequestFormComponentsPropsAvailableTemplates {
	s.Name = &v
	return s
}

func (s *FormCreateRequestFormComponentsPropsAvailableTemplates) SetProcessCode(v string) *FormCreateRequestFormComponentsPropsAvailableTemplates {
	s.ProcessCode = &v
	return s
}

type FormCreateRequestFormComponentsChildren struct {
	// 控件类型
	ComponentType *string `json:"componentType,omitempty" xml:"componentType,omitempty"`
	// 控件属性
	Props *FormCreateRequestFormComponentsChildrenProps `json:"props,omitempty" xml:"props,omitempty" type:"Struct"`
	// 子控件列表
	Children []*FormCreateRequestFormComponentsChildrenChildren `json:"children,omitempty" xml:"children,omitempty" type:"Repeated"`
}

func (s FormCreateRequestFormComponentsChildren) String() string {
	return tea.Prettify(s)
}

func (s FormCreateRequestFormComponentsChildren) GoString() string {
	return s.String()
}

func (s *FormCreateRequestFormComponentsChildren) SetComponentType(v string) *FormCreateRequestFormComponentsChildren {
	s.ComponentType = &v
	return s
}

func (s *FormCreateRequestFormComponentsChildren) SetProps(v *FormCreateRequestFormComponentsChildrenProps) *FormCreateRequestFormComponentsChildren {
	s.Props = v
	return s
}

func (s *FormCreateRequestFormComponentsChildren) SetChildren(v []*FormCreateRequestFormComponentsChildrenChildren) *FormCreateRequestFormComponentsChildren {
	s.Children = v
	return s
}

type FormCreateRequestFormComponentsChildrenProps struct {
	ComponentId *string `json:"componentId,omitempty" xml:"componentId,omitempty"`
	// 控件标题
	Label *string `json:"label,omitempty" xml:"label,omitempty"`
	// 套件中控件是否可设置为分条件字段
	AsyncCondition *bool `json:"asyncCondition,omitempty" xml:"asyncCondition,omitempty"`
	// 必填
	Required *bool `json:"required,omitempty" xml:"required,omitempty"`
	// 说明文字控件内容
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// 时间格式
	Format *string `json:"format,omitempty" xml:"format,omitempty"`
	// 金额是否需要大写，1不需要大写，其他需要大写
	Upper *string `json:"upper,omitempty" xml:"upper,omitempty"`
	// 时间单位（天、小时）
	Unit *string `json:"unit,omitempty" xml:"unit,omitempty"`
	// 输入提示
	Placeholder *string `json:"placeholder,omitempty" xml:"placeholder,omitempty"`
	// 业务别名
	BizAlias *string `json:"bizAlias,omitempty" xml:"bizAlias,omitempty"`
	// 套件的业务标识
	BizType *string `json:"bizType,omitempty" xml:"bizType,omitempty"`
	// 是否自动计算时长
	Duration *bool `json:"duration,omitempty" xml:"duration,omitempty"`
	// 联系人控件是否支持多选，1多选，0单选
	Choice *string `json:"choice,omitempty" xml:"choice,omitempty"`
	// 是否不可编辑
	Disabled *bool `json:"disabled,omitempty" xml:"disabled,omitempty"`
	// 文字提示控件显示方式:top|middle|bottom
	Align *string `json:"align,omitempty" xml:"align,omitempty"`
	// 是否隐藏字段
	Invisible *bool `json:"invisible,omitempty" xml:"invisible,omitempty"`
	// 说明文字控件链接地址
	Link *string `json:"link,omitempty" xml:"link,omitempty"`
	// 明细打印方式false横向，true纵向
	VerticalPrint *bool `json:"verticalPrint,omitempty" xml:"verticalPrint,omitempty"`
	// 公式
	Formula *string `json:"formula,omitempty" xml:"formula,omitempty"`
	// 自定义控件渲染标识
	CommonBizType *string `json:"commonBizType,omitempty" xml:"commonBizType,omitempty"`
	// 单选多选控件选项列表
	Options []*FormCreateRequestFormComponentsChildrenPropsOptions `json:"options,omitempty" xml:"options,omitempty" type:"Repeated"`
	// 是否可被打印
	Print *string `json:"print,omitempty" xml:"print,omitempty"`
	// 明细汇总统计设置
	StatField []*FormCreateRequestFormComponentsChildrenPropsStatField `json:"statField,omitempty" xml:"statField,omitempty" type:"Repeated"`
	// 数据源配置
	DataSource *FormCreateRequestFormComponentsChildrenPropsDataSource `json:"dataSource,omitempty" xml:"dataSource,omitempty" type:"Struct"`
	// 关联显示字段配置
	Fields []*FormCreateRequestFormComponentsChildrenPropsFields `json:"fields,omitempty" xml:"fields,omitempty" type:"Repeated"`
	// 地址控件模式
	AddressModel *string `json:"addressModel,omitempty" xml:"addressModel,omitempty"`
	// 部门控件是否支持多选
	Multiple *bool `json:"multiple,omitempty" xml:"multiple,omitempty"`
	// 评分控件上限
	Limit *int32 `json:"limit,omitempty" xml:"limit,omitempty"`
	// 关联审批单表单筛选列表
	AvailableTemplates []*FormCreateRequestFormComponentsChildrenPropsAvailableTemplates `json:"availableTemplates,omitempty" xml:"availableTemplates,omitempty" type:"Repeated"`
	// 明细填写方式，table（表格）、list（列表）
	TableViewMode *string `json:"tableViewMode,omitempty" xml:"tableViewMode,omitempty"`
}

func (s FormCreateRequestFormComponentsChildrenProps) String() string {
	return tea.Prettify(s)
}

func (s FormCreateRequestFormComponentsChildrenProps) GoString() string {
	return s.String()
}

func (s *FormCreateRequestFormComponentsChildrenProps) SetComponentId(v string) *FormCreateRequestFormComponentsChildrenProps {
	s.ComponentId = &v
	return s
}

func (s *FormCreateRequestFormComponentsChildrenProps) SetLabel(v string) *FormCreateRequestFormComponentsChildrenProps {
	s.Label = &v
	return s
}

func (s *FormCreateRequestFormComponentsChildrenProps) SetAsyncCondition(v bool) *FormCreateRequestFormComponentsChildrenProps {
	s.AsyncCondition = &v
	return s
}

func (s *FormCreateRequestFormComponentsChildrenProps) SetRequired(v bool) *FormCreateRequestFormComponentsChildrenProps {
	s.Required = &v
	return s
}

func (s *FormCreateRequestFormComponentsChildrenProps) SetContent(v string) *FormCreateRequestFormComponentsChildrenProps {
	s.Content = &v
	return s
}

func (s *FormCreateRequestFormComponentsChildrenProps) SetFormat(v string) *FormCreateRequestFormComponentsChildrenProps {
	s.Format = &v
	return s
}

func (s *FormCreateRequestFormComponentsChildrenProps) SetUpper(v string) *FormCreateRequestFormComponentsChildrenProps {
	s.Upper = &v
	return s
}

func (s *FormCreateRequestFormComponentsChildrenProps) SetUnit(v string) *FormCreateRequestFormComponentsChildrenProps {
	s.Unit = &v
	return s
}

func (s *FormCreateRequestFormComponentsChildrenProps) SetPlaceholder(v string) *FormCreateRequestFormComponentsChildrenProps {
	s.Placeholder = &v
	return s
}

func (s *FormCreateRequestFormComponentsChildrenProps) SetBizAlias(v string) *FormCreateRequestFormComponentsChildrenProps {
	s.BizAlias = &v
	return s
}

func (s *FormCreateRequestFormComponentsChildrenProps) SetBizType(v string) *FormCreateRequestFormComponentsChildrenProps {
	s.BizType = &v
	return s
}

func (s *FormCreateRequestFormComponentsChildrenProps) SetDuration(v bool) *FormCreateRequestFormComponentsChildrenProps {
	s.Duration = &v
	return s
}

func (s *FormCreateRequestFormComponentsChildrenProps) SetChoice(v string) *FormCreateRequestFormComponentsChildrenProps {
	s.Choice = &v
	return s
}

func (s *FormCreateRequestFormComponentsChildrenProps) SetDisabled(v bool) *FormCreateRequestFormComponentsChildrenProps {
	s.Disabled = &v
	return s
}

func (s *FormCreateRequestFormComponentsChildrenProps) SetAlign(v string) *FormCreateRequestFormComponentsChildrenProps {
	s.Align = &v
	return s
}

func (s *FormCreateRequestFormComponentsChildrenProps) SetInvisible(v bool) *FormCreateRequestFormComponentsChildrenProps {
	s.Invisible = &v
	return s
}

func (s *FormCreateRequestFormComponentsChildrenProps) SetLink(v string) *FormCreateRequestFormComponentsChildrenProps {
	s.Link = &v
	return s
}

func (s *FormCreateRequestFormComponentsChildrenProps) SetVerticalPrint(v bool) *FormCreateRequestFormComponentsChildrenProps {
	s.VerticalPrint = &v
	return s
}

func (s *FormCreateRequestFormComponentsChildrenProps) SetFormula(v string) *FormCreateRequestFormComponentsChildrenProps {
	s.Formula = &v
	return s
}

func (s *FormCreateRequestFormComponentsChildrenProps) SetCommonBizType(v string) *FormCreateRequestFormComponentsChildrenProps {
	s.CommonBizType = &v
	return s
}

func (s *FormCreateRequestFormComponentsChildrenProps) SetOptions(v []*FormCreateRequestFormComponentsChildrenPropsOptions) *FormCreateRequestFormComponentsChildrenProps {
	s.Options = v
	return s
}

func (s *FormCreateRequestFormComponentsChildrenProps) SetPrint(v string) *FormCreateRequestFormComponentsChildrenProps {
	s.Print = &v
	return s
}

func (s *FormCreateRequestFormComponentsChildrenProps) SetStatField(v []*FormCreateRequestFormComponentsChildrenPropsStatField) *FormCreateRequestFormComponentsChildrenProps {
	s.StatField = v
	return s
}

func (s *FormCreateRequestFormComponentsChildrenProps) SetDataSource(v *FormCreateRequestFormComponentsChildrenPropsDataSource) *FormCreateRequestFormComponentsChildrenProps {
	s.DataSource = v
	return s
}

func (s *FormCreateRequestFormComponentsChildrenProps) SetFields(v []*FormCreateRequestFormComponentsChildrenPropsFields) *FormCreateRequestFormComponentsChildrenProps {
	s.Fields = v
	return s
}

func (s *FormCreateRequestFormComponentsChildrenProps) SetAddressModel(v string) *FormCreateRequestFormComponentsChildrenProps {
	s.AddressModel = &v
	return s
}

func (s *FormCreateRequestFormComponentsChildrenProps) SetMultiple(v bool) *FormCreateRequestFormComponentsChildrenProps {
	s.Multiple = &v
	return s
}

func (s *FormCreateRequestFormComponentsChildrenProps) SetLimit(v int32) *FormCreateRequestFormComponentsChildrenProps {
	s.Limit = &v
	return s
}

func (s *FormCreateRequestFormComponentsChildrenProps) SetAvailableTemplates(v []*FormCreateRequestFormComponentsChildrenPropsAvailableTemplates) *FormCreateRequestFormComponentsChildrenProps {
	s.AvailableTemplates = v
	return s
}

func (s *FormCreateRequestFormComponentsChildrenProps) SetTableViewMode(v string) *FormCreateRequestFormComponentsChildrenProps {
	s.TableViewMode = &v
	return s
}

type FormCreateRequestFormComponentsChildrenPropsOptions struct {
	// 选项的显示内容
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
	// 选项的唯一主键
	Key *string `json:"key,omitempty" xml:"key,omitempty"`
}

func (s FormCreateRequestFormComponentsChildrenPropsOptions) String() string {
	return tea.Prettify(s)
}

func (s FormCreateRequestFormComponentsChildrenPropsOptions) GoString() string {
	return s.String()
}

func (s *FormCreateRequestFormComponentsChildrenPropsOptions) SetValue(v string) *FormCreateRequestFormComponentsChildrenPropsOptions {
	s.Value = &v
	return s
}

func (s *FormCreateRequestFormComponentsChildrenPropsOptions) SetKey(v string) *FormCreateRequestFormComponentsChildrenPropsOptions {
	s.Key = &v
	return s
}

type FormCreateRequestFormComponentsChildrenPropsStatField struct {
	ComponentId *string `json:"componentId,omitempty" xml:"componentId,omitempty"`
	Label       *string `json:"label,omitempty" xml:"label,omitempty"`
	Upper       *bool   `json:"upper,omitempty" xml:"upper,omitempty"`
	PayEnable   *string `json:"payEnable,omitempty" xml:"payEnable,omitempty"`
}

func (s FormCreateRequestFormComponentsChildrenPropsStatField) String() string {
	return tea.Prettify(s)
}

func (s FormCreateRequestFormComponentsChildrenPropsStatField) GoString() string {
	return s.String()
}

func (s *FormCreateRequestFormComponentsChildrenPropsStatField) SetComponentId(v string) *FormCreateRequestFormComponentsChildrenPropsStatField {
	s.ComponentId = &v
	return s
}

func (s *FormCreateRequestFormComponentsChildrenPropsStatField) SetLabel(v string) *FormCreateRequestFormComponentsChildrenPropsStatField {
	s.Label = &v
	return s
}

func (s *FormCreateRequestFormComponentsChildrenPropsStatField) SetUpper(v bool) *FormCreateRequestFormComponentsChildrenPropsStatField {
	s.Upper = &v
	return s
}

func (s *FormCreateRequestFormComponentsChildrenPropsStatField) SetPayEnable(v string) *FormCreateRequestFormComponentsChildrenPropsStatField {
	s.PayEnable = &v
	return s
}

type FormCreateRequestFormComponentsChildrenPropsDataSource struct {
	Type   *string                                                       `json:"type,omitempty" xml:"type,omitempty"`
	Target *FormCreateRequestFormComponentsChildrenPropsDataSourceTarget `json:"target,omitempty" xml:"target,omitempty" type:"Struct"`
}

func (s FormCreateRequestFormComponentsChildrenPropsDataSource) String() string {
	return tea.Prettify(s)
}

func (s FormCreateRequestFormComponentsChildrenPropsDataSource) GoString() string {
	return s.String()
}

func (s *FormCreateRequestFormComponentsChildrenPropsDataSource) SetType(v string) *FormCreateRequestFormComponentsChildrenPropsDataSource {
	s.Type = &v
	return s
}

func (s *FormCreateRequestFormComponentsChildrenPropsDataSource) SetTarget(v *FormCreateRequestFormComponentsChildrenPropsDataSourceTarget) *FormCreateRequestFormComponentsChildrenPropsDataSource {
	s.Target = v
	return s
}

type FormCreateRequestFormComponentsChildrenPropsDataSourceTarget struct {
	AppUuid  *string `json:"appUuid,omitempty" xml:"appUuid,omitempty"`
	AppType  *int64  `json:"appType,omitempty" xml:"appType,omitempty"`
	BizType  *string `json:"bizType,omitempty" xml:"bizType,omitempty"`
	FormCode *string `json:"formCode,omitempty" xml:"formCode,omitempty"`
}

func (s FormCreateRequestFormComponentsChildrenPropsDataSourceTarget) String() string {
	return tea.Prettify(s)
}

func (s FormCreateRequestFormComponentsChildrenPropsDataSourceTarget) GoString() string {
	return s.String()
}

func (s *FormCreateRequestFormComponentsChildrenPropsDataSourceTarget) SetAppUuid(v string) *FormCreateRequestFormComponentsChildrenPropsDataSourceTarget {
	s.AppUuid = &v
	return s
}

func (s *FormCreateRequestFormComponentsChildrenPropsDataSourceTarget) SetAppType(v int64) *FormCreateRequestFormComponentsChildrenPropsDataSourceTarget {
	s.AppType = &v
	return s
}

func (s *FormCreateRequestFormComponentsChildrenPropsDataSourceTarget) SetBizType(v string) *FormCreateRequestFormComponentsChildrenPropsDataSourceTarget {
	s.BizType = &v
	return s
}

func (s *FormCreateRequestFormComponentsChildrenPropsDataSourceTarget) SetFormCode(v string) *FormCreateRequestFormComponentsChildrenPropsDataSourceTarget {
	s.FormCode = &v
	return s
}

type FormCreateRequestFormComponentsChildrenPropsFields struct {
	ComponentType *string                                                  `json:"componentType,omitempty" xml:"componentType,omitempty"`
	Props         *FormCreateRequestFormComponentsChildrenPropsFieldsProps `json:"props,omitempty" xml:"props,omitempty" type:"Struct"`
}

func (s FormCreateRequestFormComponentsChildrenPropsFields) String() string {
	return tea.Prettify(s)
}

func (s FormCreateRequestFormComponentsChildrenPropsFields) GoString() string {
	return s.String()
}

func (s *FormCreateRequestFormComponentsChildrenPropsFields) SetComponentType(v string) *FormCreateRequestFormComponentsChildrenPropsFields {
	s.ComponentType = &v
	return s
}

func (s *FormCreateRequestFormComponentsChildrenPropsFields) SetProps(v *FormCreateRequestFormComponentsChildrenPropsFieldsProps) *FormCreateRequestFormComponentsChildrenPropsFields {
	s.Props = v
	return s
}

type FormCreateRequestFormComponentsChildrenPropsFieldsProps struct {
	// 表单中控件的唯一id
	ComponentId *string `json:"componentId,omitempty" xml:"componentId,omitempty"`
	// 控件标题
	Label *string `json:"label,omitempty" xml:"label,omitempty"`
	// label是否禁用修改
	LabelEditableFreeze *bool `json:"labelEditableFreeze,omitempty" xml:"labelEditableFreeze,omitempty"`
	// 必填
	Required *bool `json:"required,omitempty" xml:"required,omitempty"`
	// 必填是否可修改
	RequiredEditableFreeze *bool   `json:"requiredEditableFreeze,omitempty" xml:"requiredEditableFreeze,omitempty"`
	Print                  *string `json:"print,omitempty" xml:"print,omitempty"`
	// 说明文字控件内容
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// 时间格式
	Format *string `json:"format,omitempty" xml:"format,omitempty"`
	// 选项内容
	Options []*FormCreateRequestFormComponentsChildrenPropsFieldsPropsOptions `json:"options,omitempty" xml:"options,omitempty" type:"Repeated"`
	// 是否需要大写，1不需要大写，其他需要大写
	NotUpper *string `json:"notUpper,omitempty" xml:"notUpper,omitempty"`
	// 时间单位（天、小时）
	Unit *string `json:"unit,omitempty" xml:"unit,omitempty"`
	// 输入提示
	Placeholder *string `json:"placeholder,omitempty" xml:"placeholder,omitempty"`
	// 业务别名
	BizAlias *string `json:"bizAlias,omitempty" xml:"bizAlias,omitempty"`
	// 套件的业务标识
	BizType *string `json:"bizType,omitempty" xml:"bizType,omitempty"`
	// 是否自动计算时长
	Duration *bool `json:"duration,omitempty" xml:"duration,omitempty"`
	// 联系人控件是否支持多选，1多选，0单选
	Choice *string `json:"choice,omitempty" xml:"choice,omitempty"`
	// 是否不可编辑
	Disabled *bool `json:"disabled,omitempty" xml:"disabled,omitempty"`
	// 文字提示控件显示方式（top|middle|bottom）
	Align *string `json:"align,omitempty" xml:"align,omitempty"`
}

func (s FormCreateRequestFormComponentsChildrenPropsFieldsProps) String() string {
	return tea.Prettify(s)
}

func (s FormCreateRequestFormComponentsChildrenPropsFieldsProps) GoString() string {
	return s.String()
}

func (s *FormCreateRequestFormComponentsChildrenPropsFieldsProps) SetComponentId(v string) *FormCreateRequestFormComponentsChildrenPropsFieldsProps {
	s.ComponentId = &v
	return s
}

func (s *FormCreateRequestFormComponentsChildrenPropsFieldsProps) SetLabel(v string) *FormCreateRequestFormComponentsChildrenPropsFieldsProps {
	s.Label = &v
	return s
}

func (s *FormCreateRequestFormComponentsChildrenPropsFieldsProps) SetLabelEditableFreeze(v bool) *FormCreateRequestFormComponentsChildrenPropsFieldsProps {
	s.LabelEditableFreeze = &v
	return s
}

func (s *FormCreateRequestFormComponentsChildrenPropsFieldsProps) SetRequired(v bool) *FormCreateRequestFormComponentsChildrenPropsFieldsProps {
	s.Required = &v
	return s
}

func (s *FormCreateRequestFormComponentsChildrenPropsFieldsProps) SetRequiredEditableFreeze(v bool) *FormCreateRequestFormComponentsChildrenPropsFieldsProps {
	s.RequiredEditableFreeze = &v
	return s
}

func (s *FormCreateRequestFormComponentsChildrenPropsFieldsProps) SetPrint(v string) *FormCreateRequestFormComponentsChildrenPropsFieldsProps {
	s.Print = &v
	return s
}

func (s *FormCreateRequestFormComponentsChildrenPropsFieldsProps) SetContent(v string) *FormCreateRequestFormComponentsChildrenPropsFieldsProps {
	s.Content = &v
	return s
}

func (s *FormCreateRequestFormComponentsChildrenPropsFieldsProps) SetFormat(v string) *FormCreateRequestFormComponentsChildrenPropsFieldsProps {
	s.Format = &v
	return s
}

func (s *FormCreateRequestFormComponentsChildrenPropsFieldsProps) SetOptions(v []*FormCreateRequestFormComponentsChildrenPropsFieldsPropsOptions) *FormCreateRequestFormComponentsChildrenPropsFieldsProps {
	s.Options = v
	return s
}

func (s *FormCreateRequestFormComponentsChildrenPropsFieldsProps) SetNotUpper(v string) *FormCreateRequestFormComponentsChildrenPropsFieldsProps {
	s.NotUpper = &v
	return s
}

func (s *FormCreateRequestFormComponentsChildrenPropsFieldsProps) SetUnit(v string) *FormCreateRequestFormComponentsChildrenPropsFieldsProps {
	s.Unit = &v
	return s
}

func (s *FormCreateRequestFormComponentsChildrenPropsFieldsProps) SetPlaceholder(v string) *FormCreateRequestFormComponentsChildrenPropsFieldsProps {
	s.Placeholder = &v
	return s
}

func (s *FormCreateRequestFormComponentsChildrenPropsFieldsProps) SetBizAlias(v string) *FormCreateRequestFormComponentsChildrenPropsFieldsProps {
	s.BizAlias = &v
	return s
}

func (s *FormCreateRequestFormComponentsChildrenPropsFieldsProps) SetBizType(v string) *FormCreateRequestFormComponentsChildrenPropsFieldsProps {
	s.BizType = &v
	return s
}

func (s *FormCreateRequestFormComponentsChildrenPropsFieldsProps) SetDuration(v bool) *FormCreateRequestFormComponentsChildrenPropsFieldsProps {
	s.Duration = &v
	return s
}

func (s *FormCreateRequestFormComponentsChildrenPropsFieldsProps) SetChoice(v string) *FormCreateRequestFormComponentsChildrenPropsFieldsProps {
	s.Choice = &v
	return s
}

func (s *FormCreateRequestFormComponentsChildrenPropsFieldsProps) SetDisabled(v bool) *FormCreateRequestFormComponentsChildrenPropsFieldsProps {
	s.Disabled = &v
	return s
}

func (s *FormCreateRequestFormComponentsChildrenPropsFieldsProps) SetAlign(v string) *FormCreateRequestFormComponentsChildrenPropsFieldsProps {
	s.Align = &v
	return s
}

type FormCreateRequestFormComponentsChildrenPropsFieldsPropsOptions struct {
	// 每一个选项的唯一键
	Key *string `json:"key,omitempty" xml:"key,omitempty"`
	// 每一个选项的值
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s FormCreateRequestFormComponentsChildrenPropsFieldsPropsOptions) String() string {
	return tea.Prettify(s)
}

func (s FormCreateRequestFormComponentsChildrenPropsFieldsPropsOptions) GoString() string {
	return s.String()
}

func (s *FormCreateRequestFormComponentsChildrenPropsFieldsPropsOptions) SetKey(v string) *FormCreateRequestFormComponentsChildrenPropsFieldsPropsOptions {
	s.Key = &v
	return s
}

func (s *FormCreateRequestFormComponentsChildrenPropsFieldsPropsOptions) SetValue(v string) *FormCreateRequestFormComponentsChildrenPropsFieldsPropsOptions {
	s.Value = &v
	return s
}

type FormCreateRequestFormComponentsChildrenPropsAvailableTemplates struct {
	// 模板名称
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 模板processCode
	ProcessCode *string `json:"processCode,omitempty" xml:"processCode,omitempty"`
}

func (s FormCreateRequestFormComponentsChildrenPropsAvailableTemplates) String() string {
	return tea.Prettify(s)
}

func (s FormCreateRequestFormComponentsChildrenPropsAvailableTemplates) GoString() string {
	return s.String()
}

func (s *FormCreateRequestFormComponentsChildrenPropsAvailableTemplates) SetName(v string) *FormCreateRequestFormComponentsChildrenPropsAvailableTemplates {
	s.Name = &v
	return s
}

func (s *FormCreateRequestFormComponentsChildrenPropsAvailableTemplates) SetProcessCode(v string) *FormCreateRequestFormComponentsChildrenPropsAvailableTemplates {
	s.ProcessCode = &v
	return s
}

type FormCreateRequestFormComponentsChildrenChildren struct {
	ComponentType *string                                               `json:"componentType,omitempty" xml:"componentType,omitempty"`
	Props         *FormCreateRequestFormComponentsChildrenChildrenProps `json:"props,omitempty" xml:"props,omitempty" type:"Struct"`
}

func (s FormCreateRequestFormComponentsChildrenChildren) String() string {
	return tea.Prettify(s)
}

func (s FormCreateRequestFormComponentsChildrenChildren) GoString() string {
	return s.String()
}

func (s *FormCreateRequestFormComponentsChildrenChildren) SetComponentType(v string) *FormCreateRequestFormComponentsChildrenChildren {
	s.ComponentType = &v
	return s
}

func (s *FormCreateRequestFormComponentsChildrenChildren) SetProps(v *FormCreateRequestFormComponentsChildrenChildrenProps) *FormCreateRequestFormComponentsChildrenChildren {
	s.Props = v
	return s
}

type FormCreateRequestFormComponentsChildrenChildrenProps struct {
	ComponentId *string `json:"componentId,omitempty" xml:"componentId,omitempty"`
	// 控件标题
	Label *string `json:"label,omitempty" xml:"label,omitempty"`
	// 套件中控件是否可设置为分条件字段
	AsyncCondition *bool `json:"asyncCondition,omitempty" xml:"asyncCondition,omitempty"`
	// 必填
	Required *bool `json:"required,omitempty" xml:"required,omitempty"`
	// 说明文字控件内容
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// 时间格式
	Format *string `json:"format,omitempty" xml:"format,omitempty"`
	// 是否需要大写，1需要大写，0不需要大写
	Upper *string `json:"upper,omitempty" xml:"upper,omitempty"`
	// 时间单位（天、小时）
	Unit *string `json:"unit,omitempty" xml:"unit,omitempty"`
	// 输入提示
	Placeholder *string `json:"placeholder,omitempty" xml:"placeholder,omitempty"`
	// 业务别名
	BizAlias *string `json:"bizAlias,omitempty" xml:"bizAlias,omitempty"`
	// 套件的业务标识
	BizType *string `json:"bizType,omitempty" xml:"bizType,omitempty"`
	// 是否自动计算时长
	Duration *bool `json:"duration,omitempty" xml:"duration,omitempty"`
	// 联系人控件是否支持多选，1多选，0单选
	Choice *string `json:"choice,omitempty" xml:"choice,omitempty"`
	// 是否不可编辑
	Disabled *bool `json:"disabled,omitempty" xml:"disabled,omitempty"`
	// 文字提示控件显示方式:top|middle|bottom
	Align *string `json:"align,omitempty" xml:"align,omitempty"`
	// 是否隐藏字段
	Invisible *bool `json:"invisible,omitempty" xml:"invisible,omitempty"`
	// 说明文字控件链接地址
	Link *string `json:"link,omitempty" xml:"link,omitempty"`
	// 明细排版方式false横向，true纵向
	VerticalPrint *bool `json:"verticalPrint,omitempty" xml:"verticalPrint,omitempty"`
	// 公式
	Formula *string `json:"formula,omitempty" xml:"formula,omitempty"`
	// 自定义控件渲染标识
	CommonBizType *string                                                          `json:"commonBizType,omitempty" xml:"commonBizType,omitempty"`
	Options       []*FormCreateRequestFormComponentsChildrenChildrenPropsOptions   `json:"options,omitempty" xml:"options,omitempty" type:"Repeated"`
	Print         *string                                                          `json:"print,omitempty" xml:"print,omitempty"`
	StatField     []*FormCreateRequestFormComponentsChildrenChildrenPropsStatField `json:"statField,omitempty" xml:"statField,omitempty" type:"Repeated"`
	DataSource    *FormCreateRequestFormComponentsChildrenChildrenPropsDataSource  `json:"dataSource,omitempty" xml:"dataSource,omitempty" type:"Struct"`
	Fields        []*FormCreateRequestFormComponentsChildrenChildrenPropsFields    `json:"fields,omitempty" xml:"fields,omitempty" type:"Repeated"`
}

func (s FormCreateRequestFormComponentsChildrenChildrenProps) String() string {
	return tea.Prettify(s)
}

func (s FormCreateRequestFormComponentsChildrenChildrenProps) GoString() string {
	return s.String()
}

func (s *FormCreateRequestFormComponentsChildrenChildrenProps) SetComponentId(v string) *FormCreateRequestFormComponentsChildrenChildrenProps {
	s.ComponentId = &v
	return s
}

func (s *FormCreateRequestFormComponentsChildrenChildrenProps) SetLabel(v string) *FormCreateRequestFormComponentsChildrenChildrenProps {
	s.Label = &v
	return s
}

func (s *FormCreateRequestFormComponentsChildrenChildrenProps) SetAsyncCondition(v bool) *FormCreateRequestFormComponentsChildrenChildrenProps {
	s.AsyncCondition = &v
	return s
}

func (s *FormCreateRequestFormComponentsChildrenChildrenProps) SetRequired(v bool) *FormCreateRequestFormComponentsChildrenChildrenProps {
	s.Required = &v
	return s
}

func (s *FormCreateRequestFormComponentsChildrenChildrenProps) SetContent(v string) *FormCreateRequestFormComponentsChildrenChildrenProps {
	s.Content = &v
	return s
}

func (s *FormCreateRequestFormComponentsChildrenChildrenProps) SetFormat(v string) *FormCreateRequestFormComponentsChildrenChildrenProps {
	s.Format = &v
	return s
}

func (s *FormCreateRequestFormComponentsChildrenChildrenProps) SetUpper(v string) *FormCreateRequestFormComponentsChildrenChildrenProps {
	s.Upper = &v
	return s
}

func (s *FormCreateRequestFormComponentsChildrenChildrenProps) SetUnit(v string) *FormCreateRequestFormComponentsChildrenChildrenProps {
	s.Unit = &v
	return s
}

func (s *FormCreateRequestFormComponentsChildrenChildrenProps) SetPlaceholder(v string) *FormCreateRequestFormComponentsChildrenChildrenProps {
	s.Placeholder = &v
	return s
}

func (s *FormCreateRequestFormComponentsChildrenChildrenProps) SetBizAlias(v string) *FormCreateRequestFormComponentsChildrenChildrenProps {
	s.BizAlias = &v
	return s
}

func (s *FormCreateRequestFormComponentsChildrenChildrenProps) SetBizType(v string) *FormCreateRequestFormComponentsChildrenChildrenProps {
	s.BizType = &v
	return s
}

func (s *FormCreateRequestFormComponentsChildrenChildrenProps) SetDuration(v bool) *FormCreateRequestFormComponentsChildrenChildrenProps {
	s.Duration = &v
	return s
}

func (s *FormCreateRequestFormComponentsChildrenChildrenProps) SetChoice(v string) *FormCreateRequestFormComponentsChildrenChildrenProps {
	s.Choice = &v
	return s
}

func (s *FormCreateRequestFormComponentsChildrenChildrenProps) SetDisabled(v bool) *FormCreateRequestFormComponentsChildrenChildrenProps {
	s.Disabled = &v
	return s
}

func (s *FormCreateRequestFormComponentsChildrenChildrenProps) SetAlign(v string) *FormCreateRequestFormComponentsChildrenChildrenProps {
	s.Align = &v
	return s
}

func (s *FormCreateRequestFormComponentsChildrenChildrenProps) SetInvisible(v bool) *FormCreateRequestFormComponentsChildrenChildrenProps {
	s.Invisible = &v
	return s
}

func (s *FormCreateRequestFormComponentsChildrenChildrenProps) SetLink(v string) *FormCreateRequestFormComponentsChildrenChildrenProps {
	s.Link = &v
	return s
}

func (s *FormCreateRequestFormComponentsChildrenChildrenProps) SetVerticalPrint(v bool) *FormCreateRequestFormComponentsChildrenChildrenProps {
	s.VerticalPrint = &v
	return s
}

func (s *FormCreateRequestFormComponentsChildrenChildrenProps) SetFormula(v string) *FormCreateRequestFormComponentsChildrenChildrenProps {
	s.Formula = &v
	return s
}

func (s *FormCreateRequestFormComponentsChildrenChildrenProps) SetCommonBizType(v string) *FormCreateRequestFormComponentsChildrenChildrenProps {
	s.CommonBizType = &v
	return s
}

func (s *FormCreateRequestFormComponentsChildrenChildrenProps) SetOptions(v []*FormCreateRequestFormComponentsChildrenChildrenPropsOptions) *FormCreateRequestFormComponentsChildrenChildrenProps {
	s.Options = v
	return s
}

func (s *FormCreateRequestFormComponentsChildrenChildrenProps) SetPrint(v string) *FormCreateRequestFormComponentsChildrenChildrenProps {
	s.Print = &v
	return s
}

func (s *FormCreateRequestFormComponentsChildrenChildrenProps) SetStatField(v []*FormCreateRequestFormComponentsChildrenChildrenPropsStatField) *FormCreateRequestFormComponentsChildrenChildrenProps {
	s.StatField = v
	return s
}

func (s *FormCreateRequestFormComponentsChildrenChildrenProps) SetDataSource(v *FormCreateRequestFormComponentsChildrenChildrenPropsDataSource) *FormCreateRequestFormComponentsChildrenChildrenProps {
	s.DataSource = v
	return s
}

func (s *FormCreateRequestFormComponentsChildrenChildrenProps) SetFields(v []*FormCreateRequestFormComponentsChildrenChildrenPropsFields) *FormCreateRequestFormComponentsChildrenChildrenProps {
	s.Fields = v
	return s
}

type FormCreateRequestFormComponentsChildrenChildrenPropsOptions struct {
	// 选项的显示内容
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
	// 选项的唯一主键
	Key *string `json:"key,omitempty" xml:"key,omitempty"`
}

func (s FormCreateRequestFormComponentsChildrenChildrenPropsOptions) String() string {
	return tea.Prettify(s)
}

func (s FormCreateRequestFormComponentsChildrenChildrenPropsOptions) GoString() string {
	return s.String()
}

func (s *FormCreateRequestFormComponentsChildrenChildrenPropsOptions) SetValue(v string) *FormCreateRequestFormComponentsChildrenChildrenPropsOptions {
	s.Value = &v
	return s
}

func (s *FormCreateRequestFormComponentsChildrenChildrenPropsOptions) SetKey(v string) *FormCreateRequestFormComponentsChildrenChildrenPropsOptions {
	s.Key = &v
	return s
}

type FormCreateRequestFormComponentsChildrenChildrenPropsStatField struct {
	ComponentId *string `json:"componentId,omitempty" xml:"componentId,omitempty"`
	Label       *string `json:"label,omitempty" xml:"label,omitempty"`
	Upper       *bool   `json:"upper,omitempty" xml:"upper,omitempty"`
	PayEnable   *string `json:"payEnable,omitempty" xml:"payEnable,omitempty"`
}

func (s FormCreateRequestFormComponentsChildrenChildrenPropsStatField) String() string {
	return tea.Prettify(s)
}

func (s FormCreateRequestFormComponentsChildrenChildrenPropsStatField) GoString() string {
	return s.String()
}

func (s *FormCreateRequestFormComponentsChildrenChildrenPropsStatField) SetComponentId(v string) *FormCreateRequestFormComponentsChildrenChildrenPropsStatField {
	s.ComponentId = &v
	return s
}

func (s *FormCreateRequestFormComponentsChildrenChildrenPropsStatField) SetLabel(v string) *FormCreateRequestFormComponentsChildrenChildrenPropsStatField {
	s.Label = &v
	return s
}

func (s *FormCreateRequestFormComponentsChildrenChildrenPropsStatField) SetUpper(v bool) *FormCreateRequestFormComponentsChildrenChildrenPropsStatField {
	s.Upper = &v
	return s
}

func (s *FormCreateRequestFormComponentsChildrenChildrenPropsStatField) SetPayEnable(v string) *FormCreateRequestFormComponentsChildrenChildrenPropsStatField {
	s.PayEnable = &v
	return s
}

type FormCreateRequestFormComponentsChildrenChildrenPropsDataSource struct {
	Type   *string                                                               `json:"type,omitempty" xml:"type,omitempty"`
	Target *FormCreateRequestFormComponentsChildrenChildrenPropsDataSourceTarget `json:"target,omitempty" xml:"target,omitempty" type:"Struct"`
}

func (s FormCreateRequestFormComponentsChildrenChildrenPropsDataSource) String() string {
	return tea.Prettify(s)
}

func (s FormCreateRequestFormComponentsChildrenChildrenPropsDataSource) GoString() string {
	return s.String()
}

func (s *FormCreateRequestFormComponentsChildrenChildrenPropsDataSource) SetType(v string) *FormCreateRequestFormComponentsChildrenChildrenPropsDataSource {
	s.Type = &v
	return s
}

func (s *FormCreateRequestFormComponentsChildrenChildrenPropsDataSource) SetTarget(v *FormCreateRequestFormComponentsChildrenChildrenPropsDataSourceTarget) *FormCreateRequestFormComponentsChildrenChildrenPropsDataSource {
	s.Target = v
	return s
}

type FormCreateRequestFormComponentsChildrenChildrenPropsDataSourceTarget struct {
	AppUuid  *string `json:"appUuid,omitempty" xml:"appUuid,omitempty"`
	AppType  *int64  `json:"appType,omitempty" xml:"appType,omitempty"`
	BizType  *string `json:"bizType,omitempty" xml:"bizType,omitempty"`
	FormCode *string `json:"formCode,omitempty" xml:"formCode,omitempty"`
}

func (s FormCreateRequestFormComponentsChildrenChildrenPropsDataSourceTarget) String() string {
	return tea.Prettify(s)
}

func (s FormCreateRequestFormComponentsChildrenChildrenPropsDataSourceTarget) GoString() string {
	return s.String()
}

func (s *FormCreateRequestFormComponentsChildrenChildrenPropsDataSourceTarget) SetAppUuid(v string) *FormCreateRequestFormComponentsChildrenChildrenPropsDataSourceTarget {
	s.AppUuid = &v
	return s
}

func (s *FormCreateRequestFormComponentsChildrenChildrenPropsDataSourceTarget) SetAppType(v int64) *FormCreateRequestFormComponentsChildrenChildrenPropsDataSourceTarget {
	s.AppType = &v
	return s
}

func (s *FormCreateRequestFormComponentsChildrenChildrenPropsDataSourceTarget) SetBizType(v string) *FormCreateRequestFormComponentsChildrenChildrenPropsDataSourceTarget {
	s.BizType = &v
	return s
}

func (s *FormCreateRequestFormComponentsChildrenChildrenPropsDataSourceTarget) SetFormCode(v string) *FormCreateRequestFormComponentsChildrenChildrenPropsDataSourceTarget {
	s.FormCode = &v
	return s
}

type FormCreateRequestFormComponentsChildrenChildrenPropsFields struct {
	ComponentType *string                                                          `json:"componentType,omitempty" xml:"componentType,omitempty"`
	Props         *FormCreateRequestFormComponentsChildrenChildrenPropsFieldsProps `json:"props,omitempty" xml:"props,omitempty" type:"Struct"`
}

func (s FormCreateRequestFormComponentsChildrenChildrenPropsFields) String() string {
	return tea.Prettify(s)
}

func (s FormCreateRequestFormComponentsChildrenChildrenPropsFields) GoString() string {
	return s.String()
}

func (s *FormCreateRequestFormComponentsChildrenChildrenPropsFields) SetComponentType(v string) *FormCreateRequestFormComponentsChildrenChildrenPropsFields {
	s.ComponentType = &v
	return s
}

func (s *FormCreateRequestFormComponentsChildrenChildrenPropsFields) SetProps(v *FormCreateRequestFormComponentsChildrenChildrenPropsFieldsProps) *FormCreateRequestFormComponentsChildrenChildrenPropsFields {
	s.Props = v
	return s
}

type FormCreateRequestFormComponentsChildrenChildrenPropsFieldsProps struct {
	// 表单中控件的唯一id
	ComponentId *string `json:"componentId,omitempty" xml:"componentId,omitempty"`
	// 控件标题
	Label *string `json:"label,omitempty" xml:"label,omitempty"`
	// 必填
	Required *bool `json:"required,omitempty" xml:"required,omitempty"`
	// 字段是否可被打印，1表示打印, 0表示打印，默认打印
	Print *string `json:"print,omitempty" xml:"print,omitempty"`
	// 说明文字控件内容
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// 时间格式
	Format *string `json:"format,omitempty" xml:"format,omitempty"`
	// 选项内容
	Options []*FormCreateRequestFormComponentsChildrenChildrenPropsFieldsPropsOptions `json:"options,omitempty" xml:"options,omitempty" type:"Repeated"`
	// 是否需要大写，1需要大写，0不需要，默认1
	Upper *string `json:"upper,omitempty" xml:"upper,omitempty"`
	// 时间单位（天、小时）
	Unit *string `json:"unit,omitempty" xml:"unit,omitempty"`
	// 输入提示
	Placeholder *string `json:"placeholder,omitempty" xml:"placeholder,omitempty"`
	// 业务别名
	BizAlias *string `json:"bizAlias,omitempty" xml:"bizAlias,omitempty"`
	// 套件的业务标识
	BizType *string `json:"bizType,omitempty" xml:"bizType,omitempty"`
	// 是否自动计算时长
	Duration *bool `json:"duration,omitempty" xml:"duration,omitempty"`
	// 联系人控件是否支持多选，1多选，0单选
	Choice *string `json:"choice,omitempty" xml:"choice,omitempty"`
	// 是否不可编辑
	Disabled *bool `json:"disabled,omitempty" xml:"disabled,omitempty"`
	// 文字提示控件显示方式（top|middle|bottom）
	Align *string `json:"align,omitempty" xml:"align,omitempty"`
}

func (s FormCreateRequestFormComponentsChildrenChildrenPropsFieldsProps) String() string {
	return tea.Prettify(s)
}

func (s FormCreateRequestFormComponentsChildrenChildrenPropsFieldsProps) GoString() string {
	return s.String()
}

func (s *FormCreateRequestFormComponentsChildrenChildrenPropsFieldsProps) SetComponentId(v string) *FormCreateRequestFormComponentsChildrenChildrenPropsFieldsProps {
	s.ComponentId = &v
	return s
}

func (s *FormCreateRequestFormComponentsChildrenChildrenPropsFieldsProps) SetLabel(v string) *FormCreateRequestFormComponentsChildrenChildrenPropsFieldsProps {
	s.Label = &v
	return s
}

func (s *FormCreateRequestFormComponentsChildrenChildrenPropsFieldsProps) SetRequired(v bool) *FormCreateRequestFormComponentsChildrenChildrenPropsFieldsProps {
	s.Required = &v
	return s
}

func (s *FormCreateRequestFormComponentsChildrenChildrenPropsFieldsProps) SetPrint(v string) *FormCreateRequestFormComponentsChildrenChildrenPropsFieldsProps {
	s.Print = &v
	return s
}

func (s *FormCreateRequestFormComponentsChildrenChildrenPropsFieldsProps) SetContent(v string) *FormCreateRequestFormComponentsChildrenChildrenPropsFieldsProps {
	s.Content = &v
	return s
}

func (s *FormCreateRequestFormComponentsChildrenChildrenPropsFieldsProps) SetFormat(v string) *FormCreateRequestFormComponentsChildrenChildrenPropsFieldsProps {
	s.Format = &v
	return s
}

func (s *FormCreateRequestFormComponentsChildrenChildrenPropsFieldsProps) SetOptions(v []*FormCreateRequestFormComponentsChildrenChildrenPropsFieldsPropsOptions) *FormCreateRequestFormComponentsChildrenChildrenPropsFieldsProps {
	s.Options = v
	return s
}

func (s *FormCreateRequestFormComponentsChildrenChildrenPropsFieldsProps) SetUpper(v string) *FormCreateRequestFormComponentsChildrenChildrenPropsFieldsProps {
	s.Upper = &v
	return s
}

func (s *FormCreateRequestFormComponentsChildrenChildrenPropsFieldsProps) SetUnit(v string) *FormCreateRequestFormComponentsChildrenChildrenPropsFieldsProps {
	s.Unit = &v
	return s
}

func (s *FormCreateRequestFormComponentsChildrenChildrenPropsFieldsProps) SetPlaceholder(v string) *FormCreateRequestFormComponentsChildrenChildrenPropsFieldsProps {
	s.Placeholder = &v
	return s
}

func (s *FormCreateRequestFormComponentsChildrenChildrenPropsFieldsProps) SetBizAlias(v string) *FormCreateRequestFormComponentsChildrenChildrenPropsFieldsProps {
	s.BizAlias = &v
	return s
}

func (s *FormCreateRequestFormComponentsChildrenChildrenPropsFieldsProps) SetBizType(v string) *FormCreateRequestFormComponentsChildrenChildrenPropsFieldsProps {
	s.BizType = &v
	return s
}

func (s *FormCreateRequestFormComponentsChildrenChildrenPropsFieldsProps) SetDuration(v bool) *FormCreateRequestFormComponentsChildrenChildrenPropsFieldsProps {
	s.Duration = &v
	return s
}

func (s *FormCreateRequestFormComponentsChildrenChildrenPropsFieldsProps) SetChoice(v string) *FormCreateRequestFormComponentsChildrenChildrenPropsFieldsProps {
	s.Choice = &v
	return s
}

func (s *FormCreateRequestFormComponentsChildrenChildrenPropsFieldsProps) SetDisabled(v bool) *FormCreateRequestFormComponentsChildrenChildrenPropsFieldsProps {
	s.Disabled = &v
	return s
}

func (s *FormCreateRequestFormComponentsChildrenChildrenPropsFieldsProps) SetAlign(v string) *FormCreateRequestFormComponentsChildrenChildrenPropsFieldsProps {
	s.Align = &v
	return s
}

type FormCreateRequestFormComponentsChildrenChildrenPropsFieldsPropsOptions struct {
	// 每一个选项的唯一键
	Key *string `json:"key,omitempty" xml:"key,omitempty"`
	// 每一个选项的值
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s FormCreateRequestFormComponentsChildrenChildrenPropsFieldsPropsOptions) String() string {
	return tea.Prettify(s)
}

func (s FormCreateRequestFormComponentsChildrenChildrenPropsFieldsPropsOptions) GoString() string {
	return s.String()
}

func (s *FormCreateRequestFormComponentsChildrenChildrenPropsFieldsPropsOptions) SetKey(v string) *FormCreateRequestFormComponentsChildrenChildrenPropsFieldsPropsOptions {
	s.Key = &v
	return s
}

func (s *FormCreateRequestFormComponentsChildrenChildrenPropsFieldsPropsOptions) SetValue(v string) *FormCreateRequestFormComponentsChildrenChildrenPropsFieldsPropsOptions {
	s.Value = &v
	return s
}

type FormCreateResponseBody struct {
	// 表单模板信息
	Result *FormCreateResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s FormCreateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s FormCreateResponseBody) GoString() string {
	return s.String()
}

func (s *FormCreateResponseBody) SetResult(v *FormCreateResponseBodyResult) *FormCreateResponseBody {
	s.Result = v
	return s
}

type FormCreateResponseBodyResult struct {
	// 保存或更新的表单code
	ProcessCode *string `json:"processCode,omitempty" xml:"processCode,omitempty"`
}

func (s FormCreateResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s FormCreateResponseBodyResult) GoString() string {
	return s.String()
}

func (s *FormCreateResponseBodyResult) SetProcessCode(v string) *FormCreateResponseBodyResult {
	s.ProcessCode = &v
	return s
}

type FormCreateResponse struct {
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *FormCreateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s FormCreateResponse) String() string {
	return tea.Prettify(s)
}

func (s FormCreateResponse) GoString() string {
	return s.String()
}

func (s *FormCreateResponse) SetHeaders(v map[string]*string) *FormCreateResponse {
	s.Headers = v
	return s
}

func (s *FormCreateResponse) SetBody(v *FormCreateResponseBody) *FormCreateResponse {
	s.Body = v
	return s
}

type StartProcessInstanceHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s StartProcessInstanceHeaders) String() string {
	return tea.Prettify(s)
}

func (s StartProcessInstanceHeaders) GoString() string {
	return s.String()
}

func (s *StartProcessInstanceHeaders) SetCommonHeaders(v map[string]*string) *StartProcessInstanceHeaders {
	s.CommonHeaders = v
	return s
}

func (s *StartProcessInstanceHeaders) SetXAcsDingtalkAccessToken(v string) *StartProcessInstanceHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type StartProcessInstanceRequest struct {
	// 审批发起人的userId
	OriginatorUserId *string `json:"originatorUserId,omitempty" xml:"originatorUserId,omitempty"`
	// 审批流的唯一码
	ProcessCode *string `json:"processCode,omitempty" xml:"processCode,omitempty"`
	// 部门ID
	DeptId *int64 `json:"deptId,omitempty" xml:"deptId,omitempty"`
	// 企业微应用标识
	MicroappAgentId *int64 `json:"microappAgentId,omitempty" xml:"microappAgentId,omitempty"`
	// 不使用审批流模板时，直接指定审批人列表
	Approvers []*StartProcessInstanceRequestApprovers `json:"approvers,omitempty" xml:"approvers,omitempty" type:"Repeated"`
	// 抄送人userId列表
	CcList []*string `json:"ccList,omitempty" xml:"ccList,omitempty" type:"Repeated"`
	// 抄送时间
	CcPosition *string `json:"ccPosition,omitempty" xml:"ccPosition,omitempty"`
	// 使用审批流模板时，模板上的自选操作人列表
	TargetSelectActioners []*StartProcessInstanceRequestTargetSelectActioners `json:"targetSelectActioners,omitempty" xml:"targetSelectActioners,omitempty" type:"Repeated"`
	// 表单数据内容，控件列表
	FormComponentValues []*StartProcessInstanceRequestFormComponentValues `json:"formComponentValues,omitempty" xml:"formComponentValues,omitempty" type:"Repeated"`
	RequestId           *string                                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	DingCorpId          *string                                           `json:"dingCorpId,omitempty" xml:"dingCorpId,omitempty"`
	DingOrgId           *int64                                            `json:"dingOrgId,omitempty" xml:"dingOrgId,omitempty"`
	DingIsvOrgId        *int64                                            `json:"dingIsvOrgId,omitempty" xml:"dingIsvOrgId,omitempty"`
	DingSuiteKey        *string                                           `json:"dingSuiteKey,omitempty" xml:"dingSuiteKey,omitempty"`
	DingTokenGrantType  *int64                                            `json:"dingTokenGrantType,omitempty" xml:"dingTokenGrantType,omitempty"`
}

func (s StartProcessInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s StartProcessInstanceRequest) GoString() string {
	return s.String()
}

func (s *StartProcessInstanceRequest) SetOriginatorUserId(v string) *StartProcessInstanceRequest {
	s.OriginatorUserId = &v
	return s
}

func (s *StartProcessInstanceRequest) SetProcessCode(v string) *StartProcessInstanceRequest {
	s.ProcessCode = &v
	return s
}

func (s *StartProcessInstanceRequest) SetDeptId(v int64) *StartProcessInstanceRequest {
	s.DeptId = &v
	return s
}

func (s *StartProcessInstanceRequest) SetMicroappAgentId(v int64) *StartProcessInstanceRequest {
	s.MicroappAgentId = &v
	return s
}

func (s *StartProcessInstanceRequest) SetApprovers(v []*StartProcessInstanceRequestApprovers) *StartProcessInstanceRequest {
	s.Approvers = v
	return s
}

func (s *StartProcessInstanceRequest) SetCcList(v []*string) *StartProcessInstanceRequest {
	s.CcList = v
	return s
}

func (s *StartProcessInstanceRequest) SetCcPosition(v string) *StartProcessInstanceRequest {
	s.CcPosition = &v
	return s
}

func (s *StartProcessInstanceRequest) SetTargetSelectActioners(v []*StartProcessInstanceRequestTargetSelectActioners) *StartProcessInstanceRequest {
	s.TargetSelectActioners = v
	return s
}

func (s *StartProcessInstanceRequest) SetFormComponentValues(v []*StartProcessInstanceRequestFormComponentValues) *StartProcessInstanceRequest {
	s.FormComponentValues = v
	return s
}

func (s *StartProcessInstanceRequest) SetRequestId(v string) *StartProcessInstanceRequest {
	s.RequestId = &v
	return s
}

func (s *StartProcessInstanceRequest) SetDingCorpId(v string) *StartProcessInstanceRequest {
	s.DingCorpId = &v
	return s
}

func (s *StartProcessInstanceRequest) SetDingOrgId(v int64) *StartProcessInstanceRequest {
	s.DingOrgId = &v
	return s
}

func (s *StartProcessInstanceRequest) SetDingIsvOrgId(v int64) *StartProcessInstanceRequest {
	s.DingIsvOrgId = &v
	return s
}

func (s *StartProcessInstanceRequest) SetDingSuiteKey(v string) *StartProcessInstanceRequest {
	s.DingSuiteKey = &v
	return s
}

func (s *StartProcessInstanceRequest) SetDingTokenGrantType(v int64) *StartProcessInstanceRequest {
	s.DingTokenGrantType = &v
	return s
}

type StartProcessInstanceRequestApprovers struct {
	// 审批类型
	ActionType *string `json:"actionType,omitempty" xml:"actionType,omitempty"`
	// 审批人列表
	UserIds []*string `json:"userIds,omitempty" xml:"userIds,omitempty" type:"Repeated"`
}

func (s StartProcessInstanceRequestApprovers) String() string {
	return tea.Prettify(s)
}

func (s StartProcessInstanceRequestApprovers) GoString() string {
	return s.String()
}

func (s *StartProcessInstanceRequestApprovers) SetActionType(v string) *StartProcessInstanceRequestApprovers {
	s.ActionType = &v
	return s
}

func (s *StartProcessInstanceRequestApprovers) SetUserIds(v []*string) *StartProcessInstanceRequestApprovers {
	s.UserIds = v
	return s
}

type StartProcessInstanceRequestTargetSelectActioners struct {
	// 自选节点的规则key
	ActionerKey *string `json:"actionerKey,omitempty" xml:"actionerKey,omitempty"`
	// 操作人userId列表
	ActionerUserIds []*string `json:"actionerUserIds,omitempty" xml:"actionerUserIds,omitempty" type:"Repeated"`
}

func (s StartProcessInstanceRequestTargetSelectActioners) String() string {
	return tea.Prettify(s)
}

func (s StartProcessInstanceRequestTargetSelectActioners) GoString() string {
	return s.String()
}

func (s *StartProcessInstanceRequestTargetSelectActioners) SetActionerKey(v string) *StartProcessInstanceRequestTargetSelectActioners {
	s.ActionerKey = &v
	return s
}

func (s *StartProcessInstanceRequestTargetSelectActioners) SetActionerUserIds(v []*string) *StartProcessInstanceRequestTargetSelectActioners {
	s.ActionerUserIds = v
	return s
}

type StartProcessInstanceRequestFormComponentValues struct {
	// 控件id
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// 控件别名
	BizAlias *string `json:"bizAlias,omitempty" xml:"bizAlias,omitempty"`
	// 控件名称
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 控件值
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
	// 控件扩展值
	ExtValue *string `json:"extValue,omitempty" xml:"extValue,omitempty"`
	// 控件类型
	ComponentType *string                                                  `json:"componentType,omitempty" xml:"componentType,omitempty"`
	Details       []*StartProcessInstanceRequestFormComponentValuesDetails `json:"details,omitempty" xml:"details,omitempty" type:"Repeated"`
}

func (s StartProcessInstanceRequestFormComponentValues) String() string {
	return tea.Prettify(s)
}

func (s StartProcessInstanceRequestFormComponentValues) GoString() string {
	return s.String()
}

func (s *StartProcessInstanceRequestFormComponentValues) SetId(v string) *StartProcessInstanceRequestFormComponentValues {
	s.Id = &v
	return s
}

func (s *StartProcessInstanceRequestFormComponentValues) SetBizAlias(v string) *StartProcessInstanceRequestFormComponentValues {
	s.BizAlias = &v
	return s
}

func (s *StartProcessInstanceRequestFormComponentValues) SetName(v string) *StartProcessInstanceRequestFormComponentValues {
	s.Name = &v
	return s
}

func (s *StartProcessInstanceRequestFormComponentValues) SetValue(v string) *StartProcessInstanceRequestFormComponentValues {
	s.Value = &v
	return s
}

func (s *StartProcessInstanceRequestFormComponentValues) SetExtValue(v string) *StartProcessInstanceRequestFormComponentValues {
	s.ExtValue = &v
	return s
}

func (s *StartProcessInstanceRequestFormComponentValues) SetComponentType(v string) *StartProcessInstanceRequestFormComponentValues {
	s.ComponentType = &v
	return s
}

func (s *StartProcessInstanceRequestFormComponentValues) SetDetails(v []*StartProcessInstanceRequestFormComponentValuesDetails) *StartProcessInstanceRequestFormComponentValues {
	s.Details = v
	return s
}

type StartProcessInstanceRequestFormComponentValuesDetails struct {
	// 控件id
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// 控件别名
	BizAlias *string `json:"bizAlias,omitempty" xml:"bizAlias,omitempty"`
	// 控件名称
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 控件值
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
	// 控件扩展值
	ExtValue *string                                                         `json:"extValue,omitempty" xml:"extValue,omitempty"`
	Details  []*StartProcessInstanceRequestFormComponentValuesDetailsDetails `json:"details,omitempty" xml:"details,omitempty" type:"Repeated"`
}

func (s StartProcessInstanceRequestFormComponentValuesDetails) String() string {
	return tea.Prettify(s)
}

func (s StartProcessInstanceRequestFormComponentValuesDetails) GoString() string {
	return s.String()
}

func (s *StartProcessInstanceRequestFormComponentValuesDetails) SetId(v string) *StartProcessInstanceRequestFormComponentValuesDetails {
	s.Id = &v
	return s
}

func (s *StartProcessInstanceRequestFormComponentValuesDetails) SetBizAlias(v string) *StartProcessInstanceRequestFormComponentValuesDetails {
	s.BizAlias = &v
	return s
}

func (s *StartProcessInstanceRequestFormComponentValuesDetails) SetName(v string) *StartProcessInstanceRequestFormComponentValuesDetails {
	s.Name = &v
	return s
}

func (s *StartProcessInstanceRequestFormComponentValuesDetails) SetValue(v string) *StartProcessInstanceRequestFormComponentValuesDetails {
	s.Value = &v
	return s
}

func (s *StartProcessInstanceRequestFormComponentValuesDetails) SetExtValue(v string) *StartProcessInstanceRequestFormComponentValuesDetails {
	s.ExtValue = &v
	return s
}

func (s *StartProcessInstanceRequestFormComponentValuesDetails) SetDetails(v []*StartProcessInstanceRequestFormComponentValuesDetailsDetails) *StartProcessInstanceRequestFormComponentValuesDetails {
	s.Details = v
	return s
}

type StartProcessInstanceRequestFormComponentValuesDetailsDetails struct {
	// 控件id
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// 控件别名
	BizAlias *string `json:"bizAlias,omitempty" xml:"bizAlias,omitempty"`
	// 控件名称
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 控件值
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
	// 控件扩展值
	ExtValue *string `json:"extValue,omitempty" xml:"extValue,omitempty"`
	// 控件类型
	ComponentType *string `json:"componentType,omitempty" xml:"componentType,omitempty"`
}

func (s StartProcessInstanceRequestFormComponentValuesDetailsDetails) String() string {
	return tea.Prettify(s)
}

func (s StartProcessInstanceRequestFormComponentValuesDetailsDetails) GoString() string {
	return s.String()
}

func (s *StartProcessInstanceRequestFormComponentValuesDetailsDetails) SetId(v string) *StartProcessInstanceRequestFormComponentValuesDetailsDetails {
	s.Id = &v
	return s
}

func (s *StartProcessInstanceRequestFormComponentValuesDetailsDetails) SetBizAlias(v string) *StartProcessInstanceRequestFormComponentValuesDetailsDetails {
	s.BizAlias = &v
	return s
}

func (s *StartProcessInstanceRequestFormComponentValuesDetailsDetails) SetName(v string) *StartProcessInstanceRequestFormComponentValuesDetailsDetails {
	s.Name = &v
	return s
}

func (s *StartProcessInstanceRequestFormComponentValuesDetailsDetails) SetValue(v string) *StartProcessInstanceRequestFormComponentValuesDetailsDetails {
	s.Value = &v
	return s
}

func (s *StartProcessInstanceRequestFormComponentValuesDetailsDetails) SetExtValue(v string) *StartProcessInstanceRequestFormComponentValuesDetailsDetails {
	s.ExtValue = &v
	return s
}

func (s *StartProcessInstanceRequestFormComponentValuesDetailsDetails) SetComponentType(v string) *StartProcessInstanceRequestFormComponentValuesDetailsDetails {
	s.ComponentType = &v
	return s
}

type StartProcessInstanceResponseBody struct {
	// 审批实例id
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
}

func (s StartProcessInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StartProcessInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *StartProcessInstanceResponseBody) SetInstanceId(v string) *StartProcessInstanceResponseBody {
	s.InstanceId = &v
	return s
}

type StartProcessInstanceResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *StartProcessInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s StartProcessInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s StartProcessInstanceResponse) GoString() string {
	return s.String()
}

func (s *StartProcessInstanceResponse) SetHeaders(v map[string]*string) *StartProcessInstanceResponse {
	s.Headers = v
	return s
}

func (s *StartProcessInstanceResponse) SetBody(v *StartProcessInstanceResponseBody) *StartProcessInstanceResponse {
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

func (client *Client) QueryFormInstance(request *QueryFormInstanceRequest) (_result *QueryFormInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryFormInstanceHeaders{}
	_result = &QueryFormInstanceResponse{}
	_body, _err := client.QueryFormInstanceWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryFormInstanceWithOptions(request *QueryFormInstanceRequest, headers *QueryFormInstanceHeaders, runtime *util.RuntimeOptions) (_result *QueryFormInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FormInstanceId)) {
		query["formInstanceId"] = request.FormInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.FormCode)) {
		query["formCode"] = request.FormCode
	}

	if !tea.BoolValue(util.IsUnset(request.AppUuid)) {
		query["appUuid"] = request.AppUuid
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsDingtalkAccessToken)) {
		realHeaders["x-acs-dingtalk-access-token"] = headers.XAcsDingtalkAccessToken
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	_result = &QueryFormInstanceResponse{}
	_body, _err := client.DoROARequest(tea.String("QueryFormInstance"), tea.String("workflow_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/workflow/forms/instances"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ProcessForecast(request *ProcessForecastRequest) (_result *ProcessForecastResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ProcessForecastHeaders{}
	_result = &ProcessForecastResponse{}
	_body, _err := client.ProcessForecastWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ProcessForecastWithOptions(request *ProcessForecastRequest, headers *ProcessForecastHeaders, runtime *util.RuntimeOptions) (_result *ProcessForecastResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DingCorpId)) {
		body["dingCorpId"] = request.DingCorpId
	}

	if !tea.BoolValue(util.IsUnset(request.DingOrgId)) {
		body["dingOrgId"] = request.DingOrgId
	}

	if !tea.BoolValue(util.IsUnset(request.DingIsvOrgId)) {
		body["dingIsvOrgId"] = request.DingIsvOrgId
	}

	if !tea.BoolValue(util.IsUnset(request.DingSuiteKey)) {
		body["dingSuiteKey"] = request.DingSuiteKey
	}

	if !tea.BoolValue(util.IsUnset(request.DingTokenGrantType)) {
		body["dingTokenGrantType"] = request.DingTokenGrantType
	}

	if !tea.BoolValue(util.IsUnset(request.RequestId)) {
		body["RequestId"] = request.RequestId
	}

	if !tea.BoolValue(util.IsUnset(request.ProcessCode)) {
		body["processCode"] = request.ProcessCode
	}

	if !tea.BoolValue(util.IsUnset(request.DeptId)) {
		body["deptId"] = request.DeptId
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		body["userId"] = request.UserId
	}

	if !tea.BoolValue(util.IsUnset(request.FormComponentValues)) {
		body["formComponentValues"] = request.FormComponentValues
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsDingtalkAccessToken)) {
		realHeaders["x-acs-dingtalk-access-token"] = headers.XAcsDingtalkAccessToken
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	_result = &ProcessForecastResponse{}
	_body, _err := client.DoROARequest(tea.String("ProcessForecast"), tea.String("workflow_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/workflow/processes/forecast"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryAllProcessInstances(request *QueryAllProcessInstancesRequest) (_result *QueryAllProcessInstancesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryAllProcessInstancesHeaders{}
	_result = &QueryAllProcessInstancesResponse{}
	_body, _err := client.QueryAllProcessInstancesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryAllProcessInstancesWithOptions(request *QueryAllProcessInstancesRequest, headers *QueryAllProcessInstancesHeaders, runtime *util.RuntimeOptions) (_result *QueryAllProcessInstancesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["nextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["maxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.StartTimeInMills)) {
		query["startTimeInMills"] = request.StartTimeInMills
	}

	if !tea.BoolValue(util.IsUnset(request.EndTimeInMills)) {
		query["endTimeInMills"] = request.EndTimeInMills
	}

	if !tea.BoolValue(util.IsUnset(request.ProcessCode)) {
		query["processCode"] = request.ProcessCode
	}

	if !tea.BoolValue(util.IsUnset(request.AppUuid)) {
		query["appUuid"] = request.AppUuid
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsDingtalkAccessToken)) {
		realHeaders["x-acs-dingtalk-access-token"] = headers.XAcsDingtalkAccessToken
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	_result = &QueryAllProcessInstancesResponse{}
	_body, _err := client.DoROARequest(tea.String("QueryAllProcessInstances"), tea.String("workflow_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/workflow/processes/pages/instances"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryAllFormInstances(request *QueryAllFormInstancesRequest) (_result *QueryAllFormInstancesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryAllFormInstancesHeaders{}
	_result = &QueryAllFormInstancesResponse{}
	_body, _err := client.QueryAllFormInstancesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryAllFormInstancesWithOptions(request *QueryAllFormInstancesRequest, headers *QueryAllFormInstancesHeaders, runtime *util.RuntimeOptions) (_result *QueryAllFormInstancesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["nextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["maxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.AppUuid)) {
		query["appUuid"] = request.AppUuid
	}

	if !tea.BoolValue(util.IsUnset(request.FormCode)) {
		query["formCode"] = request.FormCode
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsDingtalkAccessToken)) {
		realHeaders["x-acs-dingtalk-access-token"] = headers.XAcsDingtalkAccessToken
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	_result = &QueryAllFormInstancesResponse{}
	_body, _err := client.DoROARequest(tea.String("QueryAllFormInstances"), tea.String("workflow_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/workflow/forms/pages/instances"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryFormByBizType(request *QueryFormByBizTypeRequest) (_result *QueryFormByBizTypeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryFormByBizTypeHeaders{}
	_result = &QueryFormByBizTypeResponse{}
	_body, _err := client.QueryFormByBizTypeWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryFormByBizTypeWithOptions(request *QueryFormByBizTypeRequest, headers *QueryFormByBizTypeHeaders, runtime *util.RuntimeOptions) (_result *QueryFormByBizTypeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppUuid)) {
		body["appUuid"] = request.AppUuid
	}

	if !tea.BoolValue(util.IsUnset(request.BizTypes)) {
		body["bizTypes"] = request.BizTypes
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsDingtalkAccessToken)) {
		realHeaders["x-acs-dingtalk-access-token"] = headers.XAcsDingtalkAccessToken
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	_result = &QueryFormByBizTypeResponse{}
	_body, _err := client.DoROARequest(tea.String("QueryFormByBizType"), tea.String("workflow_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/workflow/forms/forminfos/query"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) FormCreate(request *FormCreateRequest) (_result *FormCreateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &FormCreateHeaders{}
	_result = &FormCreateResponse{}
	_body, _err := client.FormCreateWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) FormCreateWithOptions(request *FormCreateRequest, headers *FormCreateHeaders, runtime *util.RuntimeOptions) (_result *FormCreateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DingCorpId)) {
		body["dingCorpId"] = request.DingCorpId
	}

	if !tea.BoolValue(util.IsUnset(request.DingOrgId)) {
		body["dingOrgId"] = request.DingOrgId
	}

	if !tea.BoolValue(util.IsUnset(request.DingIsvOrgId)) {
		body["dingIsvOrgId"] = request.DingIsvOrgId
	}

	if !tea.BoolValue(util.IsUnset(request.DingSuiteKey)) {
		body["dingSuiteKey"] = request.DingSuiteKey
	}

	if !tea.BoolValue(util.IsUnset(request.DingTokenGrantType)) {
		body["dingTokenGrantType"] = request.DingTokenGrantType
	}

	if !tea.BoolValue(util.IsUnset(request.RequestId)) {
		body["RequestId"] = request.RequestId
	}

	if !tea.BoolValue(util.IsUnset(request.ProcessCode)) {
		body["processCode"] = request.ProcessCode
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.FormComponents)) {
		body["formComponents"] = request.FormComponents
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsDingtalkAccessToken)) {
		realHeaders["x-acs-dingtalk-access-token"] = headers.XAcsDingtalkAccessToken
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	_result = &FormCreateResponse{}
	_body, _err := client.DoROARequest(tea.String("FormCreate"), tea.String("workflow_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/workflow/forms"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) StartProcessInstance(request *StartProcessInstanceRequest) (_result *StartProcessInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &StartProcessInstanceHeaders{}
	_result = &StartProcessInstanceResponse{}
	_body, _err := client.StartProcessInstanceWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) StartProcessInstanceWithOptions(request *StartProcessInstanceRequest, headers *StartProcessInstanceHeaders, runtime *util.RuntimeOptions) (_result *StartProcessInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OriginatorUserId)) {
		body["originatorUserId"] = request.OriginatorUserId
	}

	if !tea.BoolValue(util.IsUnset(request.ProcessCode)) {
		body["processCode"] = request.ProcessCode
	}

	if !tea.BoolValue(util.IsUnset(request.DeptId)) {
		body["deptId"] = request.DeptId
	}

	if !tea.BoolValue(util.IsUnset(request.MicroappAgentId)) {
		body["microappAgentId"] = request.MicroappAgentId
	}

	if !tea.BoolValue(util.IsUnset(request.Approvers)) {
		body["approvers"] = request.Approvers
	}

	if !tea.BoolValue(util.IsUnset(request.CcList)) {
		body["ccList"] = request.CcList
	}

	if !tea.BoolValue(util.IsUnset(request.CcPosition)) {
		body["ccPosition"] = request.CcPosition
	}

	if !tea.BoolValue(util.IsUnset(request.TargetSelectActioners)) {
		body["targetSelectActioners"] = request.TargetSelectActioners
	}

	if !tea.BoolValue(util.IsUnset(request.FormComponentValues)) {
		body["formComponentValues"] = request.FormComponentValues
	}

	if !tea.BoolValue(util.IsUnset(request.RequestId)) {
		body["RequestId"] = request.RequestId
	}

	if !tea.BoolValue(util.IsUnset(request.DingCorpId)) {
		body["dingCorpId"] = request.DingCorpId
	}

	if !tea.BoolValue(util.IsUnset(request.DingOrgId)) {
		body["dingOrgId"] = request.DingOrgId
	}

	if !tea.BoolValue(util.IsUnset(request.DingIsvOrgId)) {
		body["dingIsvOrgId"] = request.DingIsvOrgId
	}

	if !tea.BoolValue(util.IsUnset(request.DingSuiteKey)) {
		body["dingSuiteKey"] = request.DingSuiteKey
	}

	if !tea.BoolValue(util.IsUnset(request.DingTokenGrantType)) {
		body["dingTokenGrantType"] = request.DingTokenGrantType
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsDingtalkAccessToken)) {
		realHeaders["x-acs-dingtalk-access-token"] = headers.XAcsDingtalkAccessToken
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	_result = &StartProcessInstanceResponse{}
	_body, _err := client.DoROARequest(tea.String("StartProcessInstance"), tea.String("workflow_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/workflow/processInstances"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}
