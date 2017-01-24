package runtime

// AUTOGENERATED. DO NOT EDIT.

import (
	"errors"

	. "github.com/knq/chromedp/cdp"
	"github.com/mailru/easyjson"
	"github.com/mailru/easyjson/jlexer"
	"github.com/mailru/easyjson/jwriter"
)

var (
	_ BackendNode
	_ BackendNodeID
	_ ComputedProperty
	_ ErrorType
	_ Frame
	_ FrameID
	_ LoaderID
	_ Message
	_ MessageError
	_ MethodType
	_ Node
	_ NodeID
	_ NodeType
	_ PseudoType
	_ RGBA
	_ ShadowRootType
	_ Timestamp
)

// ScriptID unique script identifier.
type ScriptID string

// String returns the ScriptID as string value.
func (t ScriptID) String() string {
	return string(t)
}

// RemoteObjectID unique object identifier.
type RemoteObjectID string

// String returns the RemoteObjectID as string value.
func (t RemoteObjectID) String() string {
	return string(t)
}

// UnserializableValue primitive value which cannot be JSON-stringified.
type UnserializableValue string

// String returns the UnserializableValue as string value.
func (t UnserializableValue) String() string {
	return string(t)
}

// UnserializableValue values.
const (
	UnserializableValueInfinity         UnserializableValue = "Infinity"
	UnserializableValueNaN              UnserializableValue = "NaN"
	UnserializableValueNegativeInfinity UnserializableValue = "-Infinity"
	UnserializableValueNegative         UnserializableValue = "-0"
)

// MarshalEasyJSON satisfies easyjson.Marshaler.
func (t UnserializableValue) MarshalEasyJSON(out *jwriter.Writer) {
	out.String(string(t))
}

// MarshalJSON satisfies json.Marshaler.
func (t UnserializableValue) MarshalJSON() ([]byte, error) {
	return easyjson.Marshal(t)
}

// UnmarshalEasyJSON satisfies easyjson.Unmarshaler.
func (t *UnserializableValue) UnmarshalEasyJSON(in *jlexer.Lexer) {
	switch UnserializableValue(in.String()) {
	case UnserializableValueInfinity:
		*t = UnserializableValueInfinity
	case UnserializableValueNaN:
		*t = UnserializableValueNaN
	case UnserializableValueNegativeInfinity:
		*t = UnserializableValueNegativeInfinity
	case UnserializableValueNegative:
		*t = UnserializableValueNegative

	default:
		in.AddError(errors.New("unknown UnserializableValue value"))
	}
}

// UnmarshalJSON satisfies json.Unmarshaler.
func (t *UnserializableValue) UnmarshalJSON(buf []byte) error {
	return easyjson.Unmarshal(buf, t)
}

// RemoteObject mirror object referencing original JavaScript object.
type RemoteObject struct {
	Type                Type                `json:"type,omitempty"`                // Object type.
	Subtype             Subtype             `json:"subtype,omitempty"`             // Object subtype hint. Specified for object type values only.
	ClassName           string              `json:"className,omitempty"`           // Object class (constructor) name. Specified for object type values only.
	Value               easyjson.RawMessage `json:"value,omitempty"`               // Remote object value in case of primitive values or JSON values (if it was requested).
	UnserializableValue UnserializableValue `json:"unserializableValue,omitempty"` // Primitive value which can not be JSON-stringified does not have value, but gets this property.
	Description         string              `json:"description,omitempty"`         // String representation of the object.
	ObjectID            RemoteObjectID      `json:"objectId,omitempty"`            // Unique object identifier (for non-primitive values).
	Preview             *ObjectPreview      `json:"preview,omitempty"`             // Preview containing abbreviated property values. Specified for object type values only.
	CustomPreview       *CustomPreview      `json:"customPreview,omitempty"`
}

type CustomPreview struct {
	Header                     string         `json:"header,omitempty"`
	HasBody                    bool           `json:"hasBody,omitempty"`
	FormatterObjectID          RemoteObjectID `json:"formatterObjectId,omitempty"`
	BindRemoteObjectFunctionID RemoteObjectID `json:"bindRemoteObjectFunctionId,omitempty"`
	ConfigObjectID             RemoteObjectID `json:"configObjectId,omitempty"`
}

// ObjectPreview object containing abbreviated remote object value.
type ObjectPreview struct {
	Type        Type               `json:"type,omitempty"`        // Object type.
	Subtype     Subtype            `json:"subtype,omitempty"`     // Object subtype hint. Specified for object type values only.
	Description string             `json:"description,omitempty"` // String representation of the object.
	Overflow    bool               `json:"overflow,omitempty"`    // True iff some of the properties or entries of the original object did not fit.
	Properties  []*PropertyPreview `json:"properties,omitempty"`  // List of the properties.
	Entries     []*EntryPreview    `json:"entries,omitempty"`     // List of the entries. Specified for map and set subtype values only.
}

type PropertyPreview struct {
	Name         string         `json:"name,omitempty"`         // Property name.
	Type         Type           `json:"type,omitempty"`         // Object type. Accessor means that the property itself is an accessor property.
	Value        string         `json:"value,omitempty"`        // User-friendly property value string.
	ValuePreview *ObjectPreview `json:"valuePreview,omitempty"` // Nested value preview.
	Subtype      Subtype        `json:"subtype,omitempty"`      // Object subtype hint. Specified for object type values only.
}

type EntryPreview struct {
	Key   *ObjectPreview `json:"key,omitempty"`   // Preview of the key. Specified for map-like collection entries.
	Value *ObjectPreview `json:"value,omitempty"` // Preview of the value.
}

// PropertyDescriptor object property descriptor.
type PropertyDescriptor struct {
	Name         string        `json:"name,omitempty"`         // Property name or symbol description.
	Value        *RemoteObject `json:"value,omitempty"`        // The value associated with the property.
	Writable     bool          `json:"writable,omitempty"`     // True if the value associated with the property may be changed (data descriptors only).
	Get          *RemoteObject `json:"get,omitempty"`          // A function which serves as a getter for the property, or undefined if there is no getter (accessor descriptors only).
	Set          *RemoteObject `json:"set,omitempty"`          // A function which serves as a setter for the property, or undefined if there is no setter (accessor descriptors only).
	Configurable bool          `json:"configurable,omitempty"` // True if the type of this property descriptor may be changed and if the property may be deleted from the corresponding object.
	Enumerable   bool          `json:"enumerable,omitempty"`   // True if this property shows up during enumeration of the properties on the corresponding object.
	WasThrown    bool          `json:"wasThrown,omitempty"`    // True if the result was thrown during the evaluation.
	IsOwn        bool          `json:"isOwn,omitempty"`        // True if the property is owned for the object.
	Symbol       *RemoteObject `json:"symbol,omitempty"`       // Property symbol object, if the property is of the symbol type.
}

// InternalPropertyDescriptor object internal property descriptor. This
// property isn't normally visible in JavaScript code.
type InternalPropertyDescriptor struct {
	Name  string        `json:"name,omitempty"`  // Conventional property name.
	Value *RemoteObject `json:"value,omitempty"` // The value associated with the property.
}

// CallArgument represents function call argument. Either remote object id
// objectId, primitive value, unserializable primitive value or neither of (for
// undefined) them should be specified.
type CallArgument struct {
	Value               easyjson.RawMessage `json:"value,omitempty"`               // Primitive value.
	UnserializableValue UnserializableValue `json:"unserializableValue,omitempty"` // Primitive value which can not be JSON-stringified.
	ObjectID            RemoteObjectID      `json:"objectId,omitempty"`            // Remote object handle.
}

// ExecutionContextID id of an execution context.
type ExecutionContextID int64

// Int64 returns the ExecutionContextID as int64 value.
func (t ExecutionContextID) Int64() int64 {
	return int64(t)
}

// ExecutionContextDescription description of an isolated world.
type ExecutionContextDescription struct {
	ID      ExecutionContextID  `json:"id,omitempty"`     // Unique id of the execution context. It can be used to specify in which execution context script evaluation should be performed.
	Origin  string              `json:"origin,omitempty"` // Execution context origin.
	Name    string              `json:"name,omitempty"`   // Human readable name describing given context.
	AuxData easyjson.RawMessage `json:"auxData,omitempty"`
}

// ExceptionDetails detailed information about exception (or error) that was
// thrown during script compilation or execution.
type ExceptionDetails struct {
	ExceptionID        int64              `json:"exceptionId,omitempty"`        // Exception id.
	Text               string             `json:"text,omitempty"`               // Exception text, which should be used together with exception object when available.
	LineNumber         int64              `json:"lineNumber,omitempty"`         // Line number of the exception location (0-based).
	ColumnNumber       int64              `json:"columnNumber,omitempty"`       // Column number of the exception location (0-based).
	ScriptID           ScriptID           `json:"scriptId,omitempty"`           // Script ID of the exception location.
	URL                string             `json:"url,omitempty"`                // URL of the exception location, to be used when the script was not reported.
	StackTrace         *StackTrace        `json:"stackTrace,omitempty"`         // JavaScript stack trace if available.
	Exception          *RemoteObject      `json:"exception,omitempty"`          // Exception object if available.
	ExecutionContextID ExecutionContextID `json:"executionContextId,omitempty"` // Identifier of the context where exception happened.
}

// CallFrame stack entry for runtime errors and assertions.
type CallFrame struct {
	FunctionName string   `json:"functionName,omitempty"` // JavaScript function name.
	ScriptID     ScriptID `json:"scriptId,omitempty"`     // JavaScript script id.
	URL          string   `json:"url,omitempty"`          // JavaScript script name or url.
	LineNumber   int64    `json:"lineNumber,omitempty"`   // JavaScript script line number (0-based).
	ColumnNumber int64    `json:"columnNumber,omitempty"` // JavaScript script column number (0-based).
}

// StackTrace call frames for assertions or error messages.
type StackTrace struct {
	Description string       `json:"description,omitempty"` // String label of this stack trace. For async traces this may be a name of the function that initiated the async call.
	CallFrames  []*CallFrame `json:"callFrames,omitempty"`  // JavaScript function name.
	Parent      *StackTrace  `json:"parent,omitempty"`      // Asynchronous JavaScript stack trace that preceded this stack, if available.
}

// Type object type.
type Type string

// String returns the Type as string value.
func (t Type) String() string {
	return string(t)
}

// Type values.
const (
	TypeObject    Type = "object"
	TypeFunction  Type = "function"
	TypeUndefined Type = "undefined"
	TypeString    Type = "string"
	TypeNumber    Type = "number"
	TypeBoolean   Type = "boolean"
	TypeSymbol    Type = "symbol"
	TypeAccessor  Type = "accessor"
)

// MarshalEasyJSON satisfies easyjson.Marshaler.
func (t Type) MarshalEasyJSON(out *jwriter.Writer) {
	out.String(string(t))
}

// MarshalJSON satisfies json.Marshaler.
func (t Type) MarshalJSON() ([]byte, error) {
	return easyjson.Marshal(t)
}

// UnmarshalEasyJSON satisfies easyjson.Unmarshaler.
func (t *Type) UnmarshalEasyJSON(in *jlexer.Lexer) {
	switch Type(in.String()) {
	case TypeObject:
		*t = TypeObject
	case TypeFunction:
		*t = TypeFunction
	case TypeUndefined:
		*t = TypeUndefined
	case TypeString:
		*t = TypeString
	case TypeNumber:
		*t = TypeNumber
	case TypeBoolean:
		*t = TypeBoolean
	case TypeSymbol:
		*t = TypeSymbol
	case TypeAccessor:
		*t = TypeAccessor

	default:
		in.AddError(errors.New("unknown Type value"))
	}
}

// UnmarshalJSON satisfies json.Unmarshaler.
func (t *Type) UnmarshalJSON(buf []byte) error {
	return easyjson.Unmarshal(buf, t)
}

// Subtype object subtype hint. Specified for object type values only.
type Subtype string

// String returns the Subtype as string value.
func (t Subtype) String() string {
	return string(t)
}

// Subtype values.
const (
	SubtypeArray      Subtype = "array"
	SubtypeNull       Subtype = "null"
	SubtypeNode       Subtype = "node"
	SubtypeRegexp     Subtype = "regexp"
	SubtypeDate       Subtype = "date"
	SubtypeMap        Subtype = "map"
	SubtypeSet        Subtype = "set"
	SubtypeIterator   Subtype = "iterator"
	SubtypeGenerator  Subtype = "generator"
	SubtypeError      Subtype = "error"
	SubtypeProxy      Subtype = "proxy"
	SubtypePromise    Subtype = "promise"
	SubtypeTypedarray Subtype = "typedarray"
)

// MarshalEasyJSON satisfies easyjson.Marshaler.
func (t Subtype) MarshalEasyJSON(out *jwriter.Writer) {
	out.String(string(t))
}

// MarshalJSON satisfies json.Marshaler.
func (t Subtype) MarshalJSON() ([]byte, error) {
	return easyjson.Marshal(t)
}

// UnmarshalEasyJSON satisfies easyjson.Unmarshaler.
func (t *Subtype) UnmarshalEasyJSON(in *jlexer.Lexer) {
	switch Subtype(in.String()) {
	case SubtypeArray:
		*t = SubtypeArray
	case SubtypeNull:
		*t = SubtypeNull
	case SubtypeNode:
		*t = SubtypeNode
	case SubtypeRegexp:
		*t = SubtypeRegexp
	case SubtypeDate:
		*t = SubtypeDate
	case SubtypeMap:
		*t = SubtypeMap
	case SubtypeSet:
		*t = SubtypeSet
	case SubtypeIterator:
		*t = SubtypeIterator
	case SubtypeGenerator:
		*t = SubtypeGenerator
	case SubtypeError:
		*t = SubtypeError
	case SubtypeProxy:
		*t = SubtypeProxy
	case SubtypePromise:
		*t = SubtypePromise
	case SubtypeTypedarray:
		*t = SubtypeTypedarray

	default:
		in.AddError(errors.New("unknown Subtype value"))
	}
}

// UnmarshalJSON satisfies json.Unmarshaler.
func (t *Subtype) UnmarshalJSON(buf []byte) error {
	return easyjson.Unmarshal(buf, t)
}

// APIType type of the call.
type APIType string

// String returns the APIType as string value.
func (t APIType) String() string {
	return string(t)
}

// APIType values.
const (
	APITypeLog                 APIType = "log"
	APITypeDebug               APIType = "debug"
	APITypeInfo                APIType = "info"
	APITypeError               APIType = "error"
	APITypeWarning             APIType = "warning"
	APITypeDir                 APIType = "dir"
	APITypeDirxml              APIType = "dirxml"
	APITypeTable               APIType = "table"
	APITypeTrace               APIType = "trace"
	APITypeClear               APIType = "clear"
	APITypeStartGroup          APIType = "startGroup"
	APITypeStartGroupCollapsed APIType = "startGroupCollapsed"
	APITypeEndGroup            APIType = "endGroup"
	APITypeAssert              APIType = "assert"
	APITypeProfile             APIType = "profile"
	APITypeProfileEnd          APIType = "profileEnd"
)

// MarshalEasyJSON satisfies easyjson.Marshaler.
func (t APIType) MarshalEasyJSON(out *jwriter.Writer) {
	out.String(string(t))
}

// MarshalJSON satisfies json.Marshaler.
func (t APIType) MarshalJSON() ([]byte, error) {
	return easyjson.Marshal(t)
}

// UnmarshalEasyJSON satisfies easyjson.Unmarshaler.
func (t *APIType) UnmarshalEasyJSON(in *jlexer.Lexer) {
	switch APIType(in.String()) {
	case APITypeLog:
		*t = APITypeLog
	case APITypeDebug:
		*t = APITypeDebug
	case APITypeInfo:
		*t = APITypeInfo
	case APITypeError:
		*t = APITypeError
	case APITypeWarning:
		*t = APITypeWarning
	case APITypeDir:
		*t = APITypeDir
	case APITypeDirxml:
		*t = APITypeDirxml
	case APITypeTable:
		*t = APITypeTable
	case APITypeTrace:
		*t = APITypeTrace
	case APITypeClear:
		*t = APITypeClear
	case APITypeStartGroup:
		*t = APITypeStartGroup
	case APITypeStartGroupCollapsed:
		*t = APITypeStartGroupCollapsed
	case APITypeEndGroup:
		*t = APITypeEndGroup
	case APITypeAssert:
		*t = APITypeAssert
	case APITypeProfile:
		*t = APITypeProfile
	case APITypeProfileEnd:
		*t = APITypeProfileEnd

	default:
		in.AddError(errors.New("unknown APIType value"))
	}
}

// UnmarshalJSON satisfies json.Unmarshaler.
func (t *APIType) UnmarshalJSON(buf []byte) error {
	return easyjson.Unmarshal(buf, t)
}