// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"fmt"
	"io"
	"strconv"
	"time"
)

type App struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	InitScript   string `json:"initScript"`
	KeployConfig string `json:"keployConfig"`
	CreatedBy    *User  `json:"createdBy"`
	TestsetCount int    `json:"testsetCount"`
}

type AppInput struct {
	Name         string `json:"name"`
	InitScripts  string `json:"initScripts"`
	KeployConfig string `json:"keployConfig"`
}

type BodyResult struct {
	Normal   bool     `json:"normal"`
	Type     BodyType `json:"type"`
	Expected string   `json:"expected"`
	Actual   string   `json:"actual"`
}

type BodyResultInput struct {
	Normal   bool     `json:"normal"`
	Type     BodyType `json:"type"`
	Expected string   `json:"expected"`
	Actual   string   `json:"actual"`
}

type CmdInput struct {
	AppID    string  `json:"appId"`
	JwtToken string  `json:"jwtToken"`
	JobID    *string `json:"jobId,omitempty"`
	Cmd      Cmd     `json:"cmd"`
}

type CmdOutput struct {
	Status   bool    `json:"status"`
	CmdID    *string `json:"cmdId,omitempty"`
	Logs     *string `json:"logs,omitempty"`
	ErrorMsg *string `json:"errorMsg,omitempty"`
}

type DeleteMockInput struct {
	TestSetID string `json:"testSetId"`
	Kind      string `json:"kind"`
}

type GuestInput struct {
	FullName *string `json:"fullName,omitempty"`
	Message  string  `json:"message"`
	Company  *string `json:"company,omitempty"`
	Email    string  `json:"email"`
}

type HeaderResult struct {
	Normal   bool   `json:"normal"`
	Key      string `json:"key"`
	Expected string `json:"expected"`
	Actual   string `json:"actual"`
}

type HeaderResultInput struct {
	Normal   bool   `json:"normal"`
	Key      string `json:"key"`
	Expected string `json:"expected"`
	Actual   string `json:"actual"`
}

type HTTPReq struct {
	ProtoMajor int        `json:"protoMajor"`
	ProtoMinor int        `json:"protoMinor"`
	URL        string     `json:"url"`
	URLParam   string     `json:"urlParam"`
	Header     string     `json:"header"`
	Method     Method     `json:"method"`
	Body       string     `json:"body"`
	BodyType   BodyType   `json:"bodyType"`
	Timestamp  *time.Time `json:"timestamp,omitempty"`
}

type HTTPReqInput struct {
	ProtoMajor *int       `json:"protoMajor,omitempty"`
	ProtoMinor *int       `json:"protoMinor,omitempty"`
	URL        *string    `json:"url,omitempty"`
	URLParam   *string    `json:"urlParam,omitempty"`
	Header     *string    `json:"header,omitempty"`
	Method     *Method    `json:"method,omitempty"`
	Body       *string    `json:"body,omitempty"`
	BodyType   *BodyType  `json:"bodyType,omitempty"`
	Binary     *string    `json:"binary,omitempty"`
	Timestamp  *time.Time `json:"timestamp,omitempty"`
	Host       *string    `json:"host,omitempty"`
}

type HTTPResp struct {
	StatusCode int        `json:"statusCode"`
	Header     string     `json:"header"`
	Body       string     `json:"body"`
	BodyType   *BodyType  `json:"bodyType,omitempty"`
	ProtoMajor int        `json:"protoMajor"`
	ProtoMinor int        `json:"protoMinor"`
	Timestamp  *time.Time `json:"timestamp,omitempty"`
}

type HTTPRespInput struct {
	StatusCode    *int       `json:"statusCode,omitempty"`
	Header        *string    `json:"header,omitempty"`
	Body          *string    `json:"body,omitempty"`
	BodyType      *BodyType  `json:"bodyType,omitempty"`
	ProtoMajor    *int       `json:"protoMajor,omitempty"`
	ProtoMinor    *int       `json:"protoMinor,omitempty"`
	StatusMessage *string    `json:"StatusMessage,omitempty"`
	Binary        *string    `json:"Binary,omitempty"`
	Timestamp     *time.Time `json:"Timestamp,omitempty"`
}

type IntResult struct {
	Normal   bool `json:"normal"`
	Expected int  `json:"expected"`
	Actual   int  `json:"actual"`
}

type IntResultInput struct {
	Normal   bool `json:"normal"`
	Expected int  `json:"expected"`
	Actual   int  `json:"actual"`
}

type MockInput struct {
	ID        string  `json:"id"`
	Name      *string `json:"name,omitempty"`
	Version   *string `json:"version,omitempty"`
	Kind      *string `json:"kind,omitempty"`
	TestSetID *string `json:"testSetId,omitempty"`
	Spec      string  `json:"Spec"`
}

type Mocks struct {
	ID         string  `json:"id"`
	Name       string  `json:"name"`
	CreatedBy  *User   `json:"createdBy"`
	Version    string  `json:"version"`
	Kind       string  `json:"kind"`
	TestSetID  string  `json:"testSetId"`
	TestcaseID *string `json:"testcaseId,omitempty"`
	Spec       string  `json:"Spec"`
}

type Mutation struct {
}

type MutationResult struct {
	ID       string  `json:"id"`
	Status   bool    `json:"status"`
	ErrorMsg *string `json:"errorMsg,omitempty"`
}

type NoiseInput struct {
	TestRunID   string         `json:"testRunId"`
	TestSetID   string         `json:"testSetId"`
	NoiseParams []*NoiseParams `json:"noiseParams"`
	AppID       string         `json:"appId"`
}

type NoiseOutput struct {
	Status   bool    `json:"status"`
	ErrorMsg *string `json:"errorMsg,omitempty"`
}

type NoiseParams struct {
	TestCaseIDs  string    `json:"testCaseIDs"`
	EditedBy     string    `json:"editedBy"`
	NewAssertion string    `json:"newAssertion"`
	NoiseType    NoiseType `json:"noiseType"`
	NoiseOps     NoiseOps  `json:"noiseOps"`
}

type NormaliseOutput struct {
	Status   bool    `json:"status"`
	ErrorMsg *string `json:"errorMsg,omitempty"`
}

type NormalizeInput struct {
	TestRunID   string                 `json:"testRunId"`
	TestSetID   string                 `json:"testSetId"`
	TestCaseIDs []string               `json:"TestCaseIDs"`
	AppID       string                 `json:"appId"`
	TcReport    []*TestCaseReportInput `json:"tcReport,omitempty"`
	EditedBy    string                 `json:"editedBy"`
}

type Query struct {
}

type Result struct {
	StatusResult  *IntResult      `json:"statusResult"`
	HeadersResult []*HeaderResult `json:"headersResult,omitempty"`
	BodyResult    *BodyResult     `json:"bodyResult"`
}

type ResultInput struct {
	StatusResult  *IntResultInput      `json:"statusResult"`
	HeadersResult []*HeaderResultInput `json:"headersResult,omitempty"`
	BodyResult    *BodyResultInput     `json:"bodyResult"`
}

type TestCase struct {
	ID         string     `json:"id"`
	CreatedAt  time.Time  `json:"createdAt"`
	UpdatedAt  *time.Time `json:"updatedAt,omitempty"`
	CapturedAt *time.Time `json:"capturedAt,omitempty"`
	CreatedBy  *User      `json:"createdBy"`
	Version    string     `json:"version"`
	TestSetID  string     `json:"testSetId"`
	Name       string     `json:"name"`
	Assertions string     `json:"assertions"`
	HTTPReq    *HTTPReq   `json:"httpReq"`
	HTTPResp   *HTTPResp  `json:"httpResp"`
	Kind       *string    `json:"kind,omitempty"`
}

type TestCaseInput struct {
	ID         string         `json:"id"`
	Version    *string        `json:"version,omitempty"`
	Name       *string        `json:"name,omitempty"`
	Assertions *string        `json:"assertions,omitempty"`
	HTTPReq    *HTTPReqInput  `json:"httpReq,omitempty"`
	HTTPResp   *HTTPRespInput `json:"httpResp,omitempty"`
	TestSetID  *string        `json:"testSetId,omitempty"`
	Kind       *string        `json:"Kind,omitempty"`
	Captured   *int           `json:"Captured,omitempty"`
}

type TestCaseReport struct {
	ID              string        `json:"id"`
	Name            string        `json:"name"`
	Status          TestRunStatus `json:"status"`
	Testcase        *TestCase     `json:"testcase"`
	TestResults     *Result       `json:"testResults,omitempty"`
	TestSetReportID string        `json:"testSetReportId"`
}

type TestCaseReportInput struct {
	ID              string         `json:"id"`
	Status          TestRunStatus  `json:"status"`
	Testcase        *TestCaseInput `json:"testcase"`
	TestResults     *ResultInput   `json:"testResults"`
	TestSetReportID string         `json:"testSetReportId"`
	StartedAt       time.Time      `json:"startedAt"`
	CompletedAt     time.Time      `json:"completedAt"`
}

type TestReport struct {
	ID             string        `json:"id"`
	Name           string        `json:"name"`
	TestSetsPassed int           `json:"testSetsPassed"`
	TestSetsFailed int           `json:"testSetsFailed"`
	TotalTestSets  int           `json:"totalTestSets"`
	Status         TestRunStatus `json:"status"`
	App            *App          `json:"app"`
	RanBy          *User         `json:"ranBy"`
}

type TestReportInput struct {
	ID            string        `json:"id"`
	Name          *string       `json:"name,omitempty"`
	TotalTestSets *int          `json:"totalTestSets,omitempty"`
	Status        TestRunStatus `json:"status"`
	AppID         string        `json:"appId"`
}

type TestRunInfo struct {
	AppID     int    `json:"appId"`
	TestRunID string `json:"testRunId"`
}

type TestSet struct {
	ID            string    `json:"id"`
	Name          string    `json:"name"`
	TestcaseCount int       `json:"testcaseCount"`
	CreatedBy     *User     `json:"createdBy"`
	CreatedAt     time.Time `json:"createdAt"`
	AppID         string    `json:"appId"`
}

type TestSetInput struct {
	ID    string  `json:"id"`
	Name  *string `json:"name,omitempty"`
	AppID *string `json:"appId,omitempty"`
}

type TestSetReport struct {
	ID                string        `json:"id"`
	Name              string        `json:"name"`
	NumberOfTcsPassed int           `json:"numberOfTcsPassed"`
	NumberOfTcsFailed int           `json:"numberOfTcsFailed"`
	TotalTcs          int           `json:"totalTcs"`
	Status            TestRunStatus `json:"status"`
	TestSet           *TestSet      `json:"testSet"`
	TestRunID         string        `json:"testRunId"`
}

type TestSetReportInput struct {
	ID                string        `json:"id"`
	NumberOfTcsPassed int           `json:"numberOfTcsPassed"`
	NumberOfTcsFailed int           `json:"numberOfTcsFailed"`
	TotalTcs          int           `json:"totalTcs"`
	Status            TestRunStatus `json:"status"`
	Version           string        `json:"Version"`
	TestSetID         string        `json:"testSetId"`
	TestRunID         string        `json:"testRunId"`
}

type TestSetStatus struct {
	Status string `json:"status"`
}

type User struct {
	ID        string     `json:"id"`
	Name      string     `json:"name"`
	Email     string     `json:"email"`
	Role      UserRole   `json:"role"`
	APIKey    string     `json:"apiKey"`
	CreatedBy string     `json:"createdBy"`
	CreatedAt string     `json:"createdAt"`
	Status    UserStatus `json:"status"`
	Cid       string     `json:"cid"`
}

type UserInput struct {
	Name  string   `json:"name"`
	Email string   `json:"email"`
	Role  UserRole `json:"role"`
}

type BodyType string

const (
	BodyTypePlain BodyType = "PLAIN"
	BodyTypeJSON  BodyType = "JSON"
)

var AllBodyType = []BodyType{
	BodyTypePlain,
	BodyTypeJSON,
}

func (e BodyType) IsValid() bool {
	switch e {
	case BodyTypePlain, BodyTypeJSON:
		return true
	}
	return false
}

func (e BodyType) String() string {
	return string(e)
}

func (e *BodyType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = BodyType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid BodyType", str)
	}
	return nil
}

func (e BodyType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type Cmd string

const (
	CmdRecordStart Cmd = "RecordStart"
	CmdRecordStop  Cmd = "RecordStop"
	CmdTestStart   Cmd = "TestStart"
	CmdTestStop    Cmd = "TestStop"
)

var AllCmd = []Cmd{
	CmdRecordStart,
	CmdRecordStop,
	CmdTestStart,
	CmdTestStop,
}

func (e Cmd) IsValid() bool {
	switch e {
	case CmdRecordStart, CmdRecordStop, CmdTestStart, CmdTestStop:
		return true
	}
	return false
}

func (e Cmd) String() string {
	return string(e)
}

func (e *Cmd) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = Cmd(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid Cmd", str)
	}
	return nil
}

func (e Cmd) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type Method string

const (
	MethodGet     Method = "GET"
	MethodPut     Method = "PUT"
	MethodHead    Method = "HEAD"
	MethodPost    Method = "POST"
	MethodPatch   Method = "PATCH"
	MethodDelete  Method = "DELETE"
	MethodOptions Method = "OPTIONS"
	MethodTrace   Method = "TRACE"
)

var AllMethod = []Method{
	MethodGet,
	MethodPut,
	MethodHead,
	MethodPost,
	MethodPatch,
	MethodDelete,
	MethodOptions,
	MethodTrace,
}

func (e Method) IsValid() bool {
	switch e {
	case MethodGet, MethodPut, MethodHead, MethodPost, MethodPatch, MethodDelete, MethodOptions, MethodTrace:
		return true
	}
	return false
}

func (e Method) String() string {
	return string(e)
}

func (e *Method) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = Method(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid Method", str)
	}
	return nil
}

func (e Method) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type NoiseOps string

const (
	NoiseOpsAdd    NoiseOps = "ADD"
	NoiseOpsRemove NoiseOps = "REMOVE"
)

var AllNoiseOps = []NoiseOps{
	NoiseOpsAdd,
	NoiseOpsRemove,
}

func (e NoiseOps) IsValid() bool {
	switch e {
	case NoiseOpsAdd, NoiseOpsRemove:
		return true
	}
	return false
}

func (e NoiseOps) String() string {
	return string(e)
}

func (e *NoiseOps) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = NoiseOps(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid NoiseOps", str)
	}
	return nil
}

func (e NoiseOps) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type NoiseType string

const (
	NoiseTypeStatus  NoiseType = "STATUS"
	NoiseTypeHeaders NoiseType = "HEADERS"
	NoiseTypeBody    NoiseType = "BODY"
	NoiseTypeAll     NoiseType = "ALL"
)

var AllNoiseType = []NoiseType{
	NoiseTypeStatus,
	NoiseTypeHeaders,
	NoiseTypeBody,
	NoiseTypeAll,
}

func (e NoiseType) IsValid() bool {
	switch e {
	case NoiseTypeStatus, NoiseTypeHeaders, NoiseTypeBody, NoiseTypeAll:
		return true
	}
	return false
}

func (e NoiseType) String() string {
	return string(e)
}

func (e *NoiseType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = NoiseType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid NoiseType", str)
	}
	return nil
}

func (e NoiseType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type TestRunStatus string

const (
	TestRunStatusRunning TestRunStatus = "RUNNING"
	TestRunStatusFailed  TestRunStatus = "FAILED"
	TestRunStatusPassed  TestRunStatus = "PASSED"
)

var AllTestRunStatus = []TestRunStatus{
	TestRunStatusRunning,
	TestRunStatusFailed,
	TestRunStatusPassed,
}

func (e TestRunStatus) IsValid() bool {
	switch e {
	case TestRunStatusRunning, TestRunStatusFailed, TestRunStatusPassed:
		return true
	}
	return false
}

func (e TestRunStatus) String() string {
	return string(e)
}

func (e *TestRunStatus) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = TestRunStatus(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid TestRunStatus", str)
	}
	return nil
}

func (e TestRunStatus) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type UserRole string

const (
	UserRoleAdmin UserRole = "ADMIN"
	UserRoleUser  UserRole = "USER"
)

var AllUserRole = []UserRole{
	UserRoleAdmin,
	UserRoleUser,
}

func (e UserRole) IsValid() bool {
	switch e {
	case UserRoleAdmin, UserRoleUser:
		return true
	}
	return false
}

func (e UserRole) String() string {
	return string(e)
}

func (e *UserRole) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = UserRole(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid UserRole", str)
	}
	return nil
}

func (e UserRole) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type UserStatus string

const (
	UserStatusActive   UserStatus = "ACTIVE"
	UserStatusInactive UserStatus = "INACTIVE"
)

var AllUserStatus = []UserStatus{
	UserStatusActive,
	UserStatusInactive,
}

func (e UserStatus) IsValid() bool {
	switch e {
	case UserStatusActive, UserStatusInactive:
		return true
	}
	return false
}

func (e UserStatus) String() string {
	return string(e)
}

func (e *UserStatus) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = UserStatus(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid UserStatus", str)
	}
	return nil
}

func (e UserStatus) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
