package security

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

// CertificateID an internal certificate ID value.
type CertificateID int64

// Int64 returns the CertificateID as int64 value.
func (t CertificateID) Int64() int64 {
	return int64(t)
}

// SecurityState the security level of a page or resource.
type SecurityState string

// String returns the SecurityState as string value.
func (t SecurityState) String() string {
	return string(t)
}

// SecurityState values.
const (
	SecurityStateUnknown  SecurityState = "unknown"
	SecurityStateNeutral  SecurityState = "neutral"
	SecurityStateInsecure SecurityState = "insecure"
	SecurityStateWarning  SecurityState = "warning"
	SecurityStateSecure   SecurityState = "secure"
	SecurityStateInfo     SecurityState = "info"
)

// MarshalEasyJSON satisfies easyjson.Marshaler.
func (t SecurityState) MarshalEasyJSON(out *jwriter.Writer) {
	out.String(string(t))
}

// MarshalJSON satisfies json.Marshaler.
func (t SecurityState) MarshalJSON() ([]byte, error) {
	return easyjson.Marshal(t)
}

// UnmarshalEasyJSON satisfies easyjson.Unmarshaler.
func (t *SecurityState) UnmarshalEasyJSON(in *jlexer.Lexer) {
	switch SecurityState(in.String()) {
	case SecurityStateUnknown:
		*t = SecurityStateUnknown
	case SecurityStateNeutral:
		*t = SecurityStateNeutral
	case SecurityStateInsecure:
		*t = SecurityStateInsecure
	case SecurityStateWarning:
		*t = SecurityStateWarning
	case SecurityStateSecure:
		*t = SecurityStateSecure
	case SecurityStateInfo:
		*t = SecurityStateInfo

	default:
		in.AddError(errors.New("unknown SecurityState value"))
	}
}

// UnmarshalJSON satisfies json.Unmarshaler.
func (t *SecurityState) UnmarshalJSON(buf []byte) error {
	return easyjson.Unmarshal(buf, t)
}

// SecurityStateExplanation an explanation of an factor contributing to the
// security state.
type SecurityStateExplanation struct {
	SecurityState  SecurityState `json:"securityState,omitempty"`  // Security state representing the severity of the factor being explained.
	Summary        string        `json:"summary,omitempty"`        // Short phrase describing the type of factor.
	Description    string        `json:"description,omitempty"`    // Full text explanation of the factor.
	HasCertificate bool          `json:"hasCertificate,omitempty"` // True if the page has a certificate.
}

// InsecureContentStatus information about insecure content on the page.
type InsecureContentStatus struct {
	RanMixedContent                bool          `json:"ranMixedContent,omitempty"`                // True if the page was loaded over HTTPS and ran mixed (HTTP) content such as scripts.
	DisplayedMixedContent          bool          `json:"displayedMixedContent,omitempty"`          // True if the page was loaded over HTTPS and displayed mixed (HTTP) content such as images.
	RanContentWithCertErrors       bool          `json:"ranContentWithCertErrors,omitempty"`       // True if the page was loaded over HTTPS without certificate errors, and ran content such as scripts that were loaded with certificate errors.
	DisplayedContentWithCertErrors bool          `json:"displayedContentWithCertErrors,omitempty"` // True if the page was loaded over HTTPS without certificate errors, and displayed content such as images that were loaded with certificate errors.
	RanInsecureContentStyle        SecurityState `json:"ranInsecureContentStyle,omitempty"`        // Security state representing a page that ran insecure content.
	DisplayedInsecureContentStyle  SecurityState `json:"displayedInsecureContentStyle,omitempty"`  // Security state representing a page that displayed insecure content.
}