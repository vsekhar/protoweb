// Code generated by protoc-gen-go. DO NOT EDIT.
// source: response.proto

package web

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Response_Headers_AccessControlAllowCredentialsValue int32

const (
	Response_Headers_UNUSED_ACCESS_CONTROL_ALLOW_CREDENTIALS_VALUE Response_Headers_AccessControlAllowCredentialsValue = 0
	Response_Headers_FALSE                                         Response_Headers_AccessControlAllowCredentialsValue = 1
	Response_Headers_TRUE                                          Response_Headers_AccessControlAllowCredentialsValue = 2
)

var Response_Headers_AccessControlAllowCredentialsValue_name = map[int32]string{
	0: "UNUSED_ACCESS_CONTROL_ALLOW_CREDENTIALS_VALUE",
	1: "FALSE",
	2: "TRUE",
}

var Response_Headers_AccessControlAllowCredentialsValue_value = map[string]int32{
	"UNUSED_ACCESS_CONTROL_ALLOW_CREDENTIALS_VALUE": 0,
	"FALSE": 1,
	"TRUE":  2,
}

func (x Response_Headers_AccessControlAllowCredentialsValue) String() string {
	return proto.EnumName(Response_Headers_AccessControlAllowCredentialsValue_name, int32(x))
}

func (Response_Headers_AccessControlAllowCredentialsValue) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_0fbc901015fa5021, []int{0, 0, 0}
}

type Response_Headers_XFrameOptionsValue int32

const (
	Response_Headers_UNUSED_X_FRAME_OPTIONS_VALUE Response_Headers_XFrameOptionsValue = 0
	Response_Headers_DENY                         Response_Headers_XFrameOptionsValue = 1
	Response_Headers_SAMEORIGIN                   Response_Headers_XFrameOptionsValue = 2
)

var Response_Headers_XFrameOptionsValue_name = map[int32]string{
	0: "UNUSED_X_FRAME_OPTIONS_VALUE",
	1: "DENY",
	2: "SAMEORIGIN",
}

var Response_Headers_XFrameOptionsValue_value = map[string]int32{
	"UNUSED_X_FRAME_OPTIONS_VALUE": 0,
	"DENY":                         1,
	"SAMEORIGIN":                   2,
}

func (x Response_Headers_XFrameOptionsValue) String() string {
	return proto.EnumName(Response_Headers_XFrameOptionsValue_name, int32(x))
}

func (Response_Headers_XFrameOptionsValue) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_0fbc901015fa5021, []int{0, 0, 1}
}

type Response_Headers_SetCookieMessage_SameSiteValue int32

const (
	Response_Headers_SetCookieMessage_UNUSED_SAMESITE_VALUE Response_Headers_SetCookieMessage_SameSiteValue = 0
	Response_Headers_SetCookieMessage_STRICT                Response_Headers_SetCookieMessage_SameSiteValue = 1
	Response_Headers_SetCookieMessage_LAX                   Response_Headers_SetCookieMessage_SameSiteValue = 2
	Response_Headers_SetCookieMessage_NONE                  Response_Headers_SetCookieMessage_SameSiteValue = 3
)

var Response_Headers_SetCookieMessage_SameSiteValue_name = map[int32]string{
	0: "UNUSED_SAMESITE_VALUE",
	1: "STRICT",
	2: "LAX",
	3: "NONE",
}

var Response_Headers_SetCookieMessage_SameSiteValue_value = map[string]int32{
	"UNUSED_SAMESITE_VALUE": 0,
	"STRICT":                1,
	"LAX":                   2,
	"NONE":                  3,
}

func (x Response_Headers_SetCookieMessage_SameSiteValue) String() string {
	return proto.EnumName(Response_Headers_SetCookieMessage_SameSiteValue_name, int32(x))
}

func (Response_Headers_SetCookieMessage_SameSiteValue) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_0fbc901015fa5021, []int{0, 0, 1, 0}
}

type Response struct {
	Status StatusCodes `protobuf:"varint,1,opt,name=status,proto3,enum=web.StatusCodes" json:"status,omitempty"`
	// message Headers
	Header               *Response_Headers `protobuf:"bytes,2,opt,name=header,proto3" json:"header,omitempty"`
	Body                 []byte            `protobuf:"bytes,3,opt,name=body,proto3" json:"body,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_0fbc901015fa5021, []int{0}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetStatus() StatusCodes {
	if m != nil {
		return m.Status
	}
	return StatusCodes_STATUS_CODE_UNUSED
}

func (m *Response) GetHeader() *Response_Headers {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *Response) GetBody() []byte {
	if m != nil {
		return m.Body
	}
	return nil
}

type Response_Headers struct {
	// Types that are valid to be assigned to AccessControlAllowOrigin:
	//	*Response_Headers_AccessControlAllowOriginAll
	//	*Response_Headers_AccessControlAllowOriginOrigins
	AccessControlAllowOrigin      isResponse_Headers_AccessControlAllowOrigin         `protobuf_oneof:"AccessControlAllowOrigin"`
	AccessControlAllowCredentials Response_Headers_AccessControlAllowCredentialsValue `protobuf:"varint,3,opt,name=access_control_allow_credentials,json=accessControlAllowCredentials,proto3,enum=web.Response_Headers_AccessControlAllowCredentialsValue" json:"access_control_allow_credentials,omitempty"`
	AccessControlExposeHeaders    []string                                            `protobuf:"bytes,4,rep,name=access_control_expose_headers,json=accessControlExposeHeaders,proto3" json:"access_control_expose_headers,omitempty"`
	// Types that are valid to be assigned to Alt_Svc:
	//	*Response_Headers_AltSvcClear
	//	*Response_Headers_Service
	Alt_Svc      isResponse_Headers_Alt_Svc `protobuf_oneof:"Alt_Svc"`
	CacheControl *CacheControlResponse      `protobuf:"bytes,18,opt,name=cache_control,json=cacheControl,proto3" json:"cache_control,omitempty"`
	ContentType  *MIMEType                  `protobuf:"bytes,7,opt,name=content_type,json=contentType,proto3" json:"content_type,omitempty"`
	Date         *timestamp.Timestamp       `protobuf:"bytes,8,opt,name=date,proto3" json:"date,omitempty"`
	// Types that are valid to be assigned to Expires:
	//	*Response_Headers_ExpiresAlready
	//	*Response_Headers_ExpiresDate
	Expires isResponse_Headers_Expires `protobuf_oneof:"Expires"`
	Server  string                     `protobuf:"bytes,11,opt,name=server,proto3" json:"server,omitempty"`
	// Types that are valid to be assigned to Vary:
	//	*Response_Headers_VaryAll
	//	*Response_Headers_VaryHeaders
	Vary                 isResponse_Headers_Vary              `protobuf_oneof:"Vary"`
	XFrameOptions        Response_Headers_XFrameOptionsValue  `protobuf:"varint,16,opt,name=x_frame_options,json=xFrameOptions,proto3,enum=web.Response_Headers_XFrameOptionsValue" json:"x_frame_options,omitempty"`
	XXssProtection       string                               `protobuf:"bytes,15,opt,name=x_xss_protection,json=xXssProtection,proto3" json:"x_xss_protection,omitempty"`
	SetCookie            []*Response_Headers_SetCookieMessage `protobuf:"bytes,14,rep,name=set_cookie,json=setCookie,proto3" json:"set_cookie,omitempty"`
	Other                []*KeyValue                          `protobuf:"bytes,17,rep,name=other,proto3" json:"other,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                             `json:"-"`
	XXX_unrecognized     []byte                               `json:"-"`
	XXX_sizecache        int32                                `json:"-"`
}

func (m *Response_Headers) Reset()         { *m = Response_Headers{} }
func (m *Response_Headers) String() string { return proto.CompactTextString(m) }
func (*Response_Headers) ProtoMessage()    {}
func (*Response_Headers) Descriptor() ([]byte, []int) {
	return fileDescriptor_0fbc901015fa5021, []int{0, 0}
}

func (m *Response_Headers) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response_Headers.Unmarshal(m, b)
}
func (m *Response_Headers) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response_Headers.Marshal(b, m, deterministic)
}
func (m *Response_Headers) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response_Headers.Merge(m, src)
}
func (m *Response_Headers) XXX_Size() int {
	return xxx_messageInfo_Response_Headers.Size(m)
}
func (m *Response_Headers) XXX_DiscardUnknown() {
	xxx_messageInfo_Response_Headers.DiscardUnknown(m)
}

var xxx_messageInfo_Response_Headers proto.InternalMessageInfo

type isResponse_Headers_AccessControlAllowOrigin interface {
	isResponse_Headers_AccessControlAllowOrigin()
}

type Response_Headers_AccessControlAllowOriginAll struct {
	AccessControlAllowOriginAll Wildcard `protobuf:"varint,1,opt,name=access_control_allow_origin_all,json=accessControlAllowOriginAll,proto3,enum=web.Wildcard,oneof"`
}

type Response_Headers_AccessControlAllowOriginOrigins struct {
	AccessControlAllowOriginOrigins *StringList `protobuf:"bytes,2,opt,name=access_control_allow_origin_origins,json=accessControlAllowOriginOrigins,proto3,oneof"`
}

func (*Response_Headers_AccessControlAllowOriginAll) isResponse_Headers_AccessControlAllowOrigin() {}

func (*Response_Headers_AccessControlAllowOriginOrigins) isResponse_Headers_AccessControlAllowOrigin() {
}

func (m *Response_Headers) GetAccessControlAllowOrigin() isResponse_Headers_AccessControlAllowOrigin {
	if m != nil {
		return m.AccessControlAllowOrigin
	}
	return nil
}

func (m *Response_Headers) GetAccessControlAllowOriginAll() Wildcard {
	if x, ok := m.GetAccessControlAllowOrigin().(*Response_Headers_AccessControlAllowOriginAll); ok {
		return x.AccessControlAllowOriginAll
	}
	return Wildcard_STAR
}

func (m *Response_Headers) GetAccessControlAllowOriginOrigins() *StringList {
	if x, ok := m.GetAccessControlAllowOrigin().(*Response_Headers_AccessControlAllowOriginOrigins); ok {
		return x.AccessControlAllowOriginOrigins
	}
	return nil
}

func (m *Response_Headers) GetAccessControlAllowCredentials() Response_Headers_AccessControlAllowCredentialsValue {
	if m != nil {
		return m.AccessControlAllowCredentials
	}
	return Response_Headers_UNUSED_ACCESS_CONTROL_ALLOW_CREDENTIALS_VALUE
}

func (m *Response_Headers) GetAccessControlExposeHeaders() []string {
	if m != nil {
		return m.AccessControlExposeHeaders
	}
	return nil
}

type isResponse_Headers_Alt_Svc interface {
	isResponse_Headers_Alt_Svc()
}

type Response_Headers_AltSvcClear struct {
	AltSvcClear Clear `protobuf:"varint,5,opt,name=alt_svc_clear,json=altSvcClear,proto3,enum=web.Clear,oneof"`
}

type Response_Headers_Service struct {
	Service *Response_Headers_AltSvcMessage `protobuf:"bytes,6,opt,name=service,proto3,oneof"`
}

func (*Response_Headers_AltSvcClear) isResponse_Headers_Alt_Svc() {}

func (*Response_Headers_Service) isResponse_Headers_Alt_Svc() {}

func (m *Response_Headers) GetAlt_Svc() isResponse_Headers_Alt_Svc {
	if m != nil {
		return m.Alt_Svc
	}
	return nil
}

func (m *Response_Headers) GetAltSvcClear() Clear {
	if x, ok := m.GetAlt_Svc().(*Response_Headers_AltSvcClear); ok {
		return x.AltSvcClear
	}
	return Clear_CLEAR
}

func (m *Response_Headers) GetService() *Response_Headers_AltSvcMessage {
	if x, ok := m.GetAlt_Svc().(*Response_Headers_Service); ok {
		return x.Service
	}
	return nil
}

func (m *Response_Headers) GetCacheControl() *CacheControlResponse {
	if m != nil {
		return m.CacheControl
	}
	return nil
}

func (m *Response_Headers) GetContentType() *MIMEType {
	if m != nil {
		return m.ContentType
	}
	return nil
}

func (m *Response_Headers) GetDate() *timestamp.Timestamp {
	if m != nil {
		return m.Date
	}
	return nil
}

type isResponse_Headers_Expires interface {
	isResponse_Headers_Expires()
}

type Response_Headers_ExpiresAlready struct {
	ExpiresAlready Already `protobuf:"varint,9,opt,name=expires_already,json=expiresAlready,proto3,enum=web.Already,oneof"`
}

type Response_Headers_ExpiresDate struct {
	ExpiresDate *timestamp.Timestamp `protobuf:"bytes,10,opt,name=expires_date,json=expiresDate,proto3,oneof"`
}

func (*Response_Headers_ExpiresAlready) isResponse_Headers_Expires() {}

func (*Response_Headers_ExpiresDate) isResponse_Headers_Expires() {}

func (m *Response_Headers) GetExpires() isResponse_Headers_Expires {
	if m != nil {
		return m.Expires
	}
	return nil
}

func (m *Response_Headers) GetExpiresAlready() Already {
	if x, ok := m.GetExpires().(*Response_Headers_ExpiresAlready); ok {
		return x.ExpiresAlready
	}
	return Already_ALREADY
}

func (m *Response_Headers) GetExpiresDate() *timestamp.Timestamp {
	if x, ok := m.GetExpires().(*Response_Headers_ExpiresDate); ok {
		return x.ExpiresDate
	}
	return nil
}

func (m *Response_Headers) GetServer() string {
	if m != nil {
		return m.Server
	}
	return ""
}

type isResponse_Headers_Vary interface {
	isResponse_Headers_Vary()
}

type Response_Headers_VaryAll struct {
	VaryAll Wildcard `protobuf:"varint,12,opt,name=vary_all,json=varyAll,proto3,enum=web.Wildcard,oneof"`
}

type Response_Headers_VaryHeaders struct {
	VaryHeaders *StringList `protobuf:"bytes,13,opt,name=vary_headers,json=varyHeaders,proto3,oneof"`
}

func (*Response_Headers_VaryAll) isResponse_Headers_Vary() {}

func (*Response_Headers_VaryHeaders) isResponse_Headers_Vary() {}

func (m *Response_Headers) GetVary() isResponse_Headers_Vary {
	if m != nil {
		return m.Vary
	}
	return nil
}

func (m *Response_Headers) GetVaryAll() Wildcard {
	if x, ok := m.GetVary().(*Response_Headers_VaryAll); ok {
		return x.VaryAll
	}
	return Wildcard_STAR
}

func (m *Response_Headers) GetVaryHeaders() *StringList {
	if x, ok := m.GetVary().(*Response_Headers_VaryHeaders); ok {
		return x.VaryHeaders
	}
	return nil
}

func (m *Response_Headers) GetXFrameOptions() Response_Headers_XFrameOptionsValue {
	if m != nil {
		return m.XFrameOptions
	}
	return Response_Headers_UNUSED_X_FRAME_OPTIONS_VALUE
}

func (m *Response_Headers) GetXXssProtection() string {
	if m != nil {
		return m.XXssProtection
	}
	return ""
}

func (m *Response_Headers) GetSetCookie() []*Response_Headers_SetCookieMessage {
	if m != nil {
		return m.SetCookie
	}
	return nil
}

func (m *Response_Headers) GetOther() []*KeyValue {
	if m != nil {
		return m.Other
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Response_Headers) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Response_Headers_AccessControlAllowOriginAll)(nil),
		(*Response_Headers_AccessControlAllowOriginOrigins)(nil),
		(*Response_Headers_AltSvcClear)(nil),
		(*Response_Headers_Service)(nil),
		(*Response_Headers_ExpiresAlready)(nil),
		(*Response_Headers_ExpiresDate)(nil),
		(*Response_Headers_VaryAll)(nil),
		(*Response_Headers_VaryHeaders)(nil),
	}
}

type Response_Headers_AltSvcMessage struct {
	Services             []*Response_Headers_AltSvcMessage_Service `protobuf:"bytes,1,rep,name=services,proto3" json:"services,omitempty"`
	MaxAge               uint64                                    `protobuf:"varint,2,opt,name=max_age,json=maxAge,proto3" json:"max_age,omitempty"`
	Persist              bool                                      `protobuf:"varint,3,opt,name=persist,proto3" json:"persist,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                  `json:"-"`
	XXX_unrecognized     []byte                                    `json:"-"`
	XXX_sizecache        int32                                     `json:"-"`
}

func (m *Response_Headers_AltSvcMessage) Reset()         { *m = Response_Headers_AltSvcMessage{} }
func (m *Response_Headers_AltSvcMessage) String() string { return proto.CompactTextString(m) }
func (*Response_Headers_AltSvcMessage) ProtoMessage()    {}
func (*Response_Headers_AltSvcMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_0fbc901015fa5021, []int{0, 0, 0}
}

func (m *Response_Headers_AltSvcMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response_Headers_AltSvcMessage.Unmarshal(m, b)
}
func (m *Response_Headers_AltSvcMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response_Headers_AltSvcMessage.Marshal(b, m, deterministic)
}
func (m *Response_Headers_AltSvcMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response_Headers_AltSvcMessage.Merge(m, src)
}
func (m *Response_Headers_AltSvcMessage) XXX_Size() int {
	return xxx_messageInfo_Response_Headers_AltSvcMessage.Size(m)
}
func (m *Response_Headers_AltSvcMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_Response_Headers_AltSvcMessage.DiscardUnknown(m)
}

var xxx_messageInfo_Response_Headers_AltSvcMessage proto.InternalMessageInfo

func (m *Response_Headers_AltSvcMessage) GetServices() []*Response_Headers_AltSvcMessage_Service {
	if m != nil {
		return m.Services
	}
	return nil
}

func (m *Response_Headers_AltSvcMessage) GetMaxAge() uint64 {
	if m != nil {
		return m.MaxAge
	}
	return 0
}

func (m *Response_Headers_AltSvcMessage) GetPersist() bool {
	if m != nil {
		return m.Persist
	}
	return false
}

type Response_Headers_AltSvcMessage_Service struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	HostName             string   `protobuf:"bytes,2,opt,name=host_name,json=hostName,proto3" json:"host_name,omitempty"`
	Port                 uint32   `protobuf:"varint,3,opt,name=port,proto3" json:"port,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response_Headers_AltSvcMessage_Service) Reset() {
	*m = Response_Headers_AltSvcMessage_Service{}
}
func (m *Response_Headers_AltSvcMessage_Service) String() string { return proto.CompactTextString(m) }
func (*Response_Headers_AltSvcMessage_Service) ProtoMessage()    {}
func (*Response_Headers_AltSvcMessage_Service) Descriptor() ([]byte, []int) {
	return fileDescriptor_0fbc901015fa5021, []int{0, 0, 0, 0}
}

func (m *Response_Headers_AltSvcMessage_Service) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response_Headers_AltSvcMessage_Service.Unmarshal(m, b)
}
func (m *Response_Headers_AltSvcMessage_Service) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response_Headers_AltSvcMessage_Service.Marshal(b, m, deterministic)
}
func (m *Response_Headers_AltSvcMessage_Service) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response_Headers_AltSvcMessage_Service.Merge(m, src)
}
func (m *Response_Headers_AltSvcMessage_Service) XXX_Size() int {
	return xxx_messageInfo_Response_Headers_AltSvcMessage_Service.Size(m)
}
func (m *Response_Headers_AltSvcMessage_Service) XXX_DiscardUnknown() {
	xxx_messageInfo_Response_Headers_AltSvcMessage_Service.DiscardUnknown(m)
}

var xxx_messageInfo_Response_Headers_AltSvcMessage_Service proto.InternalMessageInfo

func (m *Response_Headers_AltSvcMessage_Service) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Response_Headers_AltSvcMessage_Service) GetHostName() string {
	if m != nil {
		return m.HostName
	}
	return ""
}

func (m *Response_Headers_AltSvcMessage_Service) GetPort() uint32 {
	if m != nil {
		return m.Port
	}
	return 0
}

type Response_Headers_SetCookieMessage struct {
	Name                 string                                          `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Value                string                                          `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	Expires              *timestamp.Timestamp                            `protobuf:"bytes,3,opt,name=expires,proto3" json:"expires,omitempty"`
	MaxAge               int64                                           `protobuf:"varint,4,opt,name=max_age,json=maxAge,proto3" json:"max_age,omitempty"`
	Domain               string                                          `protobuf:"bytes,5,opt,name=domain,proto3" json:"domain,omitempty"`
	Path                 string                                          `protobuf:"bytes,6,opt,name=path,proto3" json:"path,omitempty"`
	Secure               bool                                            `protobuf:"varint,7,opt,name=secure,proto3" json:"secure,omitempty"`
	HttpOnly             bool                                            `protobuf:"varint,8,opt,name=http_only,json=httpOnly,proto3" json:"http_only,omitempty"`
	Samesite             Response_Headers_SetCookieMessage_SameSiteValue `protobuf:"varint,9,opt,name=samesite,proto3,enum=web.Response_Headers_SetCookieMessage_SameSiteValue" json:"samesite,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                        `json:"-"`
	XXX_unrecognized     []byte                                          `json:"-"`
	XXX_sizecache        int32                                           `json:"-"`
}

func (m *Response_Headers_SetCookieMessage) Reset()         { *m = Response_Headers_SetCookieMessage{} }
func (m *Response_Headers_SetCookieMessage) String() string { return proto.CompactTextString(m) }
func (*Response_Headers_SetCookieMessage) ProtoMessage()    {}
func (*Response_Headers_SetCookieMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_0fbc901015fa5021, []int{0, 0, 1}
}

func (m *Response_Headers_SetCookieMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response_Headers_SetCookieMessage.Unmarshal(m, b)
}
func (m *Response_Headers_SetCookieMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response_Headers_SetCookieMessage.Marshal(b, m, deterministic)
}
func (m *Response_Headers_SetCookieMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response_Headers_SetCookieMessage.Merge(m, src)
}
func (m *Response_Headers_SetCookieMessage) XXX_Size() int {
	return xxx_messageInfo_Response_Headers_SetCookieMessage.Size(m)
}
func (m *Response_Headers_SetCookieMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_Response_Headers_SetCookieMessage.DiscardUnknown(m)
}

var xxx_messageInfo_Response_Headers_SetCookieMessage proto.InternalMessageInfo

func (m *Response_Headers_SetCookieMessage) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Response_Headers_SetCookieMessage) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func (m *Response_Headers_SetCookieMessage) GetExpires() *timestamp.Timestamp {
	if m != nil {
		return m.Expires
	}
	return nil
}

func (m *Response_Headers_SetCookieMessage) GetMaxAge() int64 {
	if m != nil {
		return m.MaxAge
	}
	return 0
}

func (m *Response_Headers_SetCookieMessage) GetDomain() string {
	if m != nil {
		return m.Domain
	}
	return ""
}

func (m *Response_Headers_SetCookieMessage) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *Response_Headers_SetCookieMessage) GetSecure() bool {
	if m != nil {
		return m.Secure
	}
	return false
}

func (m *Response_Headers_SetCookieMessage) GetHttpOnly() bool {
	if m != nil {
		return m.HttpOnly
	}
	return false
}

func (m *Response_Headers_SetCookieMessage) GetSamesite() Response_Headers_SetCookieMessage_SameSiteValue {
	if m != nil {
		return m.Samesite
	}
	return Response_Headers_SetCookieMessage_UNUSED_SAMESITE_VALUE
}

func init() {
	proto.RegisterEnum("web.Response_Headers_AccessControlAllowCredentialsValue", Response_Headers_AccessControlAllowCredentialsValue_name, Response_Headers_AccessControlAllowCredentialsValue_value)
	proto.RegisterEnum("web.Response_Headers_XFrameOptionsValue", Response_Headers_XFrameOptionsValue_name, Response_Headers_XFrameOptionsValue_value)
	proto.RegisterEnum("web.Response_Headers_SetCookieMessage_SameSiteValue", Response_Headers_SetCookieMessage_SameSiteValue_name, Response_Headers_SetCookieMessage_SameSiteValue_value)
	proto.RegisterType((*Response)(nil), "web.Response")
	proto.RegisterType((*Response_Headers)(nil), "web.Response.Headers")
	proto.RegisterType((*Response_Headers_AltSvcMessage)(nil), "web.Response.Headers.AltSvcMessage")
	proto.RegisterType((*Response_Headers_AltSvcMessage_Service)(nil), "web.Response.Headers.AltSvcMessage.Service")
	proto.RegisterType((*Response_Headers_SetCookieMessage)(nil), "web.Response.Headers.SetCookieMessage")
}

func init() { proto.RegisterFile("response.proto", fileDescriptor_0fbc901015fa5021) }

var fileDescriptor_0fbc901015fa5021 = []byte{
	// 1085 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x55, 0xdf, 0x6f, 0x1a, 0xc7,
	0x13, 0xcf, 0x01, 0x86, 0x63, 0xf8, 0xe1, 0xfb, 0xae, 0xbe, 0x49, 0x2f, 0xa4, 0x91, 0x91, 0x2d,
	0x55, 0xa8, 0x55, 0x48, 0xea, 0x5a, 0x6a, 0x9f, 0x1a, 0x9d, 0xf1, 0x39, 0xa0, 0x62, 0xb0, 0xf6,
	0xb0, 0xe3, 0x3e, 0xad, 0xd6, 0xc7, 0x06, 0x9f, 0x7a, 0x70, 0xa7, 0xdb, 0x35, 0x86, 0xc7, 0x3e,
	0xf6, 0x2f, 0xec, 0x43, 0xd5, 0xff, 0xa5, 0xda, 0x1f, 0xe7, 0xd8, 0x09, 0x76, 0xf2, 0x02, 0xbb,
	0x33, 0x9f, 0xf9, 0xcc, 0xce, 0xf0, 0x99, 0x01, 0x9a, 0x19, 0xe3, 0x69, 0xb2, 0xe0, 0xac, 0x9b,
	0x66, 0x89, 0x48, 0x50, 0xf1, 0x86, 0x5d, 0xb6, 0x76, 0x66, 0x49, 0x32, 0x8b, 0xd9, 0x6b, 0x65,
	0xba, 0xbc, 0xfe, 0xf0, 0x5a, 0x44, 0x73, 0xc6, 0x05, 0x9d, 0xa7, 0x1a, 0xd5, 0x42, 0x21, 0x0d,
	0xaf, 0x58, 0x98, 0x2c, 0x44, 0x96, 0xc4, 0xc6, 0x06, 0xf3, 0x68, 0x6e, 0x58, 0x5a, 0x75, 0x2e,
	0xa8, 0xb8, 0xe6, 0xb9, 0xe7, 0x5a, 0x44, 0x06, 0xb5, 0xfb, 0xb7, 0x03, 0x36, 0x36, 0x29, 0x51,
	0x07, 0xca, 0x1a, 0xe8, 0x5a, 0x6d, 0xab, 0xd3, 0xdc, 0x77, 0xba, 0x37, 0xec, 0xb2, 0x1b, 0x28,
	0x53, 0x2f, 0x99, 0x32, 0x8e, 0x8d, 0x1f, 0xbd, 0x82, 0xf2, 0x15, 0xa3, 0x53, 0x96, 0xb9, 0x85,
	0xb6, 0xd5, 0xa9, 0xed, 0x3f, 0x55, 0xc8, 0x9c, 0xa8, 0xdb, 0x57, 0x3e, 0x8e, 0x0d, 0x08, 0x21,
	0x28, 0x5d, 0x26, 0xd3, 0xb5, 0x5b, 0x6c, 0x5b, 0x9d, 0x3a, 0x56, 0xe7, 0xd6, 0xbf, 0xdb, 0x50,
	0x31, 0x38, 0x74, 0x06, 0x3b, 0x34, 0x0c, 0x19, 0xe7, 0xc4, 0xd4, 0x40, 0x68, 0x1c, 0x27, 0x37,
	0x24, 0xc9, 0xa2, 0x59, 0xb4, 0x90, 0x17, 0xf3, 0xa2, 0x86, 0xca, 0xf3, 0x3e, 0x8a, 0xa7, 0x21,
	0xcd, 0xa6, 0xfd, 0x27, 0xf8, 0x85, 0x8e, 0xeb, 0xe9, 0x30, 0x4f, 0x46, 0x8d, 0x55, 0x90, 0x17,
	0xc7, 0x88, 0xc0, 0xde, 0x63, 0xb4, 0xfa, 0x8b, 0x9b, 0x12, 0xb6, 0x4d, 0xb1, 0x59, 0xb4, 0x98,
	0x0d, 0x23, 0x2e, 0xfa, 0x4f, 0xf0, 0xce, 0x43, 0xe4, 0xfa, 0x93, 0xa3, 0x3f, 0x2d, 0x68, 0x6f,
	0xcc, 0x10, 0x66, 0x6c, 0xca, 0x16, 0x22, 0xa2, 0x31, 0x57, 0x45, 0x37, 0xf7, 0x7f, 0xd9, 0xd8,
	0xa1, 0xae, 0xf7, 0x59, 0x86, 0xde, 0xc7, 0xd0, 0x73, 0x1a, 0x5f, 0x33, 0xfc, 0x92, 0x3e, 0x86,
	0x41, 0x1e, 0xbc, 0xfc, 0xe4, 0x09, 0x6c, 0x95, 0x26, 0x9c, 0x11, 0xdd, 0x7b, 0xee, 0x96, 0xda,
	0xc5, 0x4e, 0x15, 0xb7, 0xee, 0xb1, 0xf8, 0x0a, 0x92, 0xb7, 0xff, 0x0d, 0x34, 0x68, 0x2c, 0x08,
	0x5f, 0x86, 0x24, 0x8c, 0x19, 0xcd, 0xdc, 0x2d, 0xf5, 0x64, 0x50, 0x4f, 0xee, 0x49, 0x4b, 0xdf,
	0xc2, 0x35, 0x1a, 0x8b, 0x60, 0x19, 0xaa, 0x2b, 0x7a, 0x0b, 0x15, 0xce, 0xb2, 0x65, 0x14, 0x32,
	0xb7, 0xac, 0xba, 0xb7, 0xf7, 0x40, 0x79, 0x2a, 0xe6, 0x84, 0x71, 0x4e, 0x67, 0xac, 0x6f, 0xe1,
	0x3c, 0x0a, 0xfd, 0x0a, 0x0d, 0xa5, 0xd9, 0xfc, 0xd1, 0x2e, 0x52, 0x34, 0xcf, 0x75, 0x4a, 0xe9,
	0x31, 0x2f, 0xcd, 0x29, 0x71, 0x3d, 0xbc, 0x63, 0x45, 0x6f, 0xa0, 0x2e, 0x23, 0xd9, 0x42, 0x10,
	0xb1, 0x4e, 0x99, 0x5b, 0x51, 0xe1, 0x5a, 0x1e, 0x27, 0x83, 0x13, 0x7f, 0xb2, 0x4e, 0x19, 0xae,
	0x19, 0x88, 0xbc, 0xa0, 0x2e, 0x94, 0xa6, 0x54, 0x30, 0xd7, 0x56, 0xc8, 0x56, 0x57, 0xcf, 0x54,
	0x37, 0x9f, 0xa9, 0xee, 0x24, 0x9f, 0x29, 0xac, 0x70, 0xe8, 0x67, 0xd8, 0x66, 0xab, 0x34, 0xca,
	0x18, 0x27, 0x34, 0xce, 0x18, 0x9d, 0xae, 0xdd, 0xaa, 0x6a, 0x4b, 0x5d, 0x25, 0xf1, 0xb4, 0xad,
	0x5f, 0xc0, 0x4d, 0x03, 0x33, 0x16, 0xf4, 0x16, 0xea, 0x79, 0xa0, 0x4a, 0x08, 0x5f, 0x4a, 0xd8,
	0x2f, 0xe0, 0x9a, 0x89, 0x38, 0x92, 0x99, 0x9f, 0x41, 0x59, 0xb6, 0x89, 0x65, 0x6e, 0xad, 0x6d,
	0x75, 0xaa, 0xd8, 0xdc, 0xd0, 0xf7, 0x60, 0x2f, 0x69, 0xb6, 0x56, 0xe3, 0x50, 0xdf, 0x34, 0x0e,
	0x45, 0x5c, 0x91, 0x00, 0x29, 0xfd, 0x03, 0xa8, 0x2b, 0x6c, 0x2e, 0x82, 0xc6, 0x66, 0x8d, 0x17,
	0x71, 0x4d, 0xc2, 0x72, 0x21, 0x9c, 0xc2, 0xf6, 0x8a, 0x7c, 0xc8, 0xe8, 0x9c, 0x91, 0x24, 0x15,
	0x51, 0xb2, 0xe0, 0xae, 0xa3, 0x12, 0x75, 0x36, 0xff, 0xbc, 0x17, 0xc7, 0x12, 0x3b, 0xd6, 0x50,
	0xad, 0xd6, 0xc6, 0xea, 0xae, 0x0d, 0x75, 0xc0, 0x59, 0x91, 0x15, 0xe7, 0x44, 0x96, 0xcd, 0x42,
	0x69, 0x74, 0xb7, 0x55, 0x55, 0xcd, 0xd5, 0x05, 0xe7, 0xa7, 0xb7, 0x56, 0xe4, 0x03, 0x70, 0x26,
	0x48, 0x98, 0x24, 0x7f, 0x44, 0xcc, 0x6d, 0xb6, 0x8b, 0x9d, 0xda, 0xfe, 0x77, 0x9b, 0xd3, 0x06,
	0x4c, 0xf4, 0x14, 0xcc, 0x08, 0x0b, 0x57, 0x79, 0x6e, 0x41, 0x7b, 0xb0, 0x95, 0x88, 0x2b, 0x96,
	0xb9, 0xff, 0x53, 0x0c, 0xba, 0x43, 0xbf, 0xb1, 0xb5, 0x7e, 0x9d, 0xf6, 0xb5, 0xfe, 0xb1, 0xa0,
	0x71, 0x4f, 0x9a, 0xe8, 0x1d, 0xd8, 0x46, 0x9a, 0x72, 0xf9, 0xc9, 0xc8, 0x1f, 0xbe, 0x42, 0xd1,
	0xdd, 0x40, 0xc7, 0xe0, 0xdb, 0x60, 0xf4, 0x0d, 0x54, 0xe6, 0x74, 0x45, 0xe8, 0x8c, 0xa9, 0xbd,
	0x52, 0xc2, 0xe5, 0x39, 0x5d, 0x79, 0x33, 0x86, 0x5c, 0xa8, 0xa4, 0x2c, 0xe3, 0x11, 0x17, 0x6a,
	0x23, 0xd8, 0x38, 0xbf, 0xb6, 0x46, 0x50, 0x31, 0x3c, 0x72, 0x51, 0x2e, 0xe8, 0x9c, 0xa9, 0x6d,
	0x57, 0xc5, 0xea, 0x8c, 0x5e, 0x40, 0xf5, 0x2a, 0xe1, 0x82, 0x28, 0x47, 0x41, 0x39, 0x6c, 0x69,
	0x18, 0x49, 0x27, 0x82, 0x52, 0x9a, 0x64, 0x9a, 0xb2, 0x81, 0xd5, 0xb9, 0xf5, 0x57, 0x11, 0x9c,
	0x4f, 0x5b, 0xb4, 0x91, 0xf9, 0xff, 0xb0, 0xb5, 0x94, 0x6d, 0x31, 0xac, 0xfa, 0x82, 0x0e, 0xa0,
	0x62, 0xd4, 0xa8, 0x58, 0x1f, 0x9f, 0x95, 0x1c, 0x7a, 0xb7, 0xee, 0x52, 0xdb, 0xea, 0x14, 0x6f,
	0xeb, 0x7e, 0x06, 0xe5, 0x69, 0x32, 0xa7, 0xd1, 0x42, 0x6d, 0x95, 0x2a, 0x36, 0x37, 0xf5, 0x72,
	0x2a, 0xae, 0xd4, 0xfe, 0xa8, 0x62, 0x75, 0xd6, 0xca, 0x0f, 0xaf, 0x33, 0x3d, 0xcf, 0x36, 0x36,
	0x37, 0xd5, 0x02, 0x21, 0x52, 0x92, 0x2c, 0xe2, 0xb5, 0x1a, 0x60, 0x1b, 0xdb, 0xd2, 0x30, 0x5e,
	0xc4, 0x6b, 0x74, 0x0a, 0x36, 0xa7, 0x73, 0xc6, 0x23, 0xc1, 0xcc, 0x84, 0x1e, 0x7c, 0x9d, 0x6c,
	0xba, 0x01, 0x9d, 0xb3, 0x20, 0x12, 0x4c, 0x6b, 0xe3, 0x96, 0x65, 0x77, 0x00, 0x8d, 0x7b, 0x2e,
	0xf4, 0x1c, 0x9e, 0x9e, 0x8d, 0xce, 0x02, 0xff, 0x88, 0x04, 0xde, 0x89, 0x1f, 0x0c, 0x26, 0x3e,
	0x39, 0xf7, 0x86, 0x67, 0xbe, 0xf3, 0x04, 0x01, 0x94, 0x83, 0x09, 0x1e, 0xf4, 0x26, 0x8e, 0x85,
	0x2a, 0x50, 0x1c, 0x7a, 0x17, 0x4e, 0x01, 0xd9, 0x50, 0x1a, 0x8d, 0x47, 0xbe, 0x53, 0xdc, 0x8d,
	0x61, 0xf7, 0xcb, 0x2b, 0x1e, 0xfd, 0x08, 0xaf, 0x0c, 0xbf, 0xd7, 0xeb, 0xf9, 0x41, 0x40, 0x7a,
	0xe3, 0xd1, 0x04, 0x8f, 0x87, 0xc4, 0x1b, 0x0e, 0xc7, 0xef, 0x49, 0x0f, 0xfb, 0x47, 0xfe, 0x68,
	0x32, 0xf0, 0x86, 0xc1, 0x6d, 0xde, 0x2a, 0x6c, 0x1d, 0x7b, 0xc3, 0xc0, 0x77, 0x2c, 0x99, 0x6d,
	0x82, 0xcf, 0x7c, 0xa7, 0xb0, 0x7b, 0x0a, 0xe8, 0xf3, 0x91, 0x44, 0x6d, 0xf8, 0xd6, 0xb0, 0x5f,
	0x90, 0x63, 0xec, 0x9d, 0xf8, 0x64, 0x7c, 0x3a, 0x19, 0x8c, 0x47, 0x1f, 0xc9, 0x6c, 0x28, 0x1d,
	0xf9, 0xa3, 0xdf, 0x1d, 0x0b, 0x35, 0x01, 0x64, 0x89, 0x63, 0x3c, 0x78, 0x37, 0x18, 0x39, 0x85,
	0xc3, 0x16, 0xb8, 0xde, 0x03, 0x7f, 0x82, 0x87, 0x55, 0xa8, 0x78, 0xb1, 0x20, 0xc1, 0x32, 0x94,
	0x47, 0x5f, 0x0b, 0xe1, 0xb0, 0x0c, 0xa5, 0x73, 0x9a, 0xad, 0x2f, 0xcb, 0x4a, 0x2d, 0x3f, 0xfd,
	0x17, 0x00, 0x00, 0xff, 0xff, 0xcd, 0x70, 0x89, 0x68, 0xd2, 0x08, 0x00, 0x00,
}
