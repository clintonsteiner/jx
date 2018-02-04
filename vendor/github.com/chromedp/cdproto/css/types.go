package css

// Code generated by chromedp-gen. DO NOT EDIT.

import (
	"errors"

	"github.com/chromedp/cdproto/cdp"
	"github.com/mailru/easyjson"
	"github.com/mailru/easyjson/jlexer"
	"github.com/mailru/easyjson/jwriter"
)

// StyleSheetID [no description].
type StyleSheetID string

// String returns the StyleSheetID as string value.
func (t StyleSheetID) String() string {
	return string(t)
}

// StyleSheetOrigin stylesheet type: "injected" for stylesheets injected via
// extension, "user-agent" for user-agent stylesheets, "inspector" for
// stylesheets created by the inspector (i.e. those holding the "via inspector"
// rules), "regular" for regular stylesheets.
type StyleSheetOrigin string

// String returns the StyleSheetOrigin as string value.
func (t StyleSheetOrigin) String() string {
	return string(t)
}

// StyleSheetOrigin values.
const (
	StyleSheetOriginInjected  StyleSheetOrigin = "injected"
	StyleSheetOriginUserAgent StyleSheetOrigin = "user-agent"
	StyleSheetOriginInspector StyleSheetOrigin = "inspector"
	StyleSheetOriginRegular   StyleSheetOrigin = "regular"
)

// MarshalEasyJSON satisfies easyjson.Marshaler.
func (t StyleSheetOrigin) MarshalEasyJSON(out *jwriter.Writer) {
	out.String(string(t))
}

// MarshalJSON satisfies json.Marshaler.
func (t StyleSheetOrigin) MarshalJSON() ([]byte, error) {
	return easyjson.Marshal(t)
}

// UnmarshalEasyJSON satisfies easyjson.Unmarshaler.
func (t *StyleSheetOrigin) UnmarshalEasyJSON(in *jlexer.Lexer) {
	switch StyleSheetOrigin(in.String()) {
	case StyleSheetOriginInjected:
		*t = StyleSheetOriginInjected
	case StyleSheetOriginUserAgent:
		*t = StyleSheetOriginUserAgent
	case StyleSheetOriginInspector:
		*t = StyleSheetOriginInspector
	case StyleSheetOriginRegular:
		*t = StyleSheetOriginRegular

	default:
		in.AddError(errors.New("unknown StyleSheetOrigin value"))
	}
}

// UnmarshalJSON satisfies json.Unmarshaler.
func (t *StyleSheetOrigin) UnmarshalJSON(buf []byte) error {
	return easyjson.Unmarshal(buf, t)
}

// PseudoElementMatches CSS rule collection for a single pseudo style.
type PseudoElementMatches struct {
	PseudoType cdp.PseudoType `json:"pseudoType"` // Pseudo element type.
	Matches    []*RuleMatch   `json:"matches"`    // Matches of CSS rules applicable to the pseudo style.
}

// InheritedStyleEntry inherited CSS rule collection from ancestor node.
type InheritedStyleEntry struct {
	InlineStyle     *Style       `json:"inlineStyle,omitempty"` // The ancestor node's inline style, if any, in the style inheritance chain.
	MatchedCSSRules []*RuleMatch `json:"matchedCSSRules"`       // Matches of CSS rules matching the ancestor node in the style inheritance chain.
}

// RuleMatch match data for a CSS rule.
type RuleMatch struct {
	Rule              *Rule   `json:"rule"`              // CSS rule in the match.
	MatchingSelectors []int64 `json:"matchingSelectors"` // Matching selector indices in the rule's selectorList selectors (0-based).
}

// Value data for a simple selector (these are delimited by commas in a
// selector list).
type Value struct {
	Text  string       `json:"text"`            // Value text.
	Range *SourceRange `json:"range,omitempty"` // Value range in the underlying resource (if available).
}

// SelectorList selector list data.
type SelectorList struct {
	Selectors []*Value `json:"selectors"` // Selectors in the list.
	Text      string   `json:"text"`      // Rule selector text.
}

// StyleSheetHeader CSS stylesheet metainformation.
type StyleSheetHeader struct {
	StyleSheetID StyleSheetID      `json:"styleSheetId"`           // The stylesheet identifier.
	FrameID      cdp.FrameID       `json:"frameId"`                // Owner frame identifier.
	SourceURL    string            `json:"sourceURL"`              // Stylesheet resource URL.
	SourceMapURL string            `json:"sourceMapURL,omitempty"` // URL of source map associated with the stylesheet (if any).
	Origin       StyleSheetOrigin  `json:"origin"`                 // Stylesheet origin.
	Title        string            `json:"title"`                  // Stylesheet title.
	OwnerNode    cdp.BackendNodeID `json:"ownerNode,omitempty"`    // The backend id for the owner node of the stylesheet.
	Disabled     bool              `json:"disabled"`               // Denotes whether the stylesheet is disabled.
	HasSourceURL bool              `json:"hasSourceURL,omitempty"` // Whether the sourceURL field value comes from the sourceURL comment.
	IsInline     bool              `json:"isInline"`               // Whether this stylesheet is created for STYLE tag by parser. This flag is not set for document.written STYLE tags.
	StartLine    float64           `json:"startLine"`              // Line offset of the stylesheet within the resource (zero based).
	StartColumn  float64           `json:"startColumn"`            // Column offset of the stylesheet within the resource (zero based).
	Length       float64           `json:"length"`                 // Size of the content (in characters).
}

// Rule CSS rule representation.
type Rule struct {
	StyleSheetID StyleSheetID     `json:"styleSheetId,omitempty"` // The css style sheet identifier (absent for user agent stylesheet and user-specified stylesheet rules) this rule came from.
	SelectorList *SelectorList    `json:"selectorList"`           // Rule selector data.
	Origin       StyleSheetOrigin `json:"origin"`                 // Parent stylesheet's origin.
	Style        *Style           `json:"style"`                  // Associated style declaration.
	Media        []*Media         `json:"media,omitempty"`        // Media list array (for rules involving media queries). The array enumerates media queries starting with the innermost one, going outwards.
}

// RuleUsage CSS coverage information.
type RuleUsage struct {
	StyleSheetID StyleSheetID `json:"styleSheetId"` // The css style sheet identifier (absent for user agent stylesheet and user-specified stylesheet rules) this rule came from.
	StartOffset  float64      `json:"startOffset"`  // Offset of the start of the rule (including selector) from the beginning of the stylesheet.
	EndOffset    float64      `json:"endOffset"`    // Offset of the end of the rule body from the beginning of the stylesheet.
	Used         bool         `json:"used"`         // Indicates whether the rule was actually used by some element in the page.
}

// SourceRange text range within a resource. All numbers are zero-based.
type SourceRange struct {
	StartLine   int64 `json:"startLine"`   // Start line of range.
	StartColumn int64 `json:"startColumn"` // Start column of range (inclusive).
	EndLine     int64 `json:"endLine"`     // End line of range
	EndColumn   int64 `json:"endColumn"`   // End column of range (exclusive).
}

// ShorthandEntry [no description].
type ShorthandEntry struct {
	Name      string `json:"name"`                // Shorthand name.
	Value     string `json:"value"`               // Shorthand value.
	Important bool   `json:"important,omitempty"` // Whether the property has "!important" annotation (implies false if absent).
}

// ComputedProperty [no description].
type ComputedProperty struct {
	Name  string `json:"name"`  // Computed style property name.
	Value string `json:"value"` // Computed style property value.
}

// Style CSS style representation.
type Style struct {
	StyleSheetID     StyleSheetID      `json:"styleSheetId,omitempty"` // The css style sheet identifier (absent for user agent stylesheet and user-specified stylesheet rules) this rule came from.
	CSSProperties    []*Property       `json:"cssProperties"`          // CSS properties in the style.
	ShorthandEntries []*ShorthandEntry `json:"shorthandEntries"`       // Computed values for all shorthands found in the style.
	CSSText          string            `json:"cssText,omitempty"`      // Style declaration text (if available).
	Range            *SourceRange      `json:"range,omitempty"`        // Style declaration range in the enclosing stylesheet (if available).
}

// Property CSS property declaration data.
type Property struct {
	Name      string       `json:"name"`                // The property name.
	Value     string       `json:"value"`               // The property value.
	Important bool         `json:"important,omitempty"` // Whether the property has "!important" annotation (implies false if absent).
	Implicit  bool         `json:"implicit,omitempty"`  // Whether the property is implicit (implies false if absent).
	Text      string       `json:"text,omitempty"`      // The full property text as specified in the style.
	ParsedOk  bool         `json:"parsedOk,omitempty"`  // Whether the property is understood by the browser (implies true if absent).
	Disabled  bool         `json:"disabled,omitempty"`  // Whether the property is disabled by the user (present for source-based properties only).
	Range     *SourceRange `json:"range,omitempty"`     // The entire property range in the enclosing style declaration (if available).
}

// Media CSS media rule descriptor.
type Media struct {
	Text         string        `json:"text"`                   // Media query text.
	Source       MediaSource   `json:"source"`                 // Source of the media query: "mediaRule" if specified by a @media rule, "importRule" if specified by an @import rule, "linkedSheet" if specified by a "media" attribute in a linked stylesheet's LINK tag, "inlineSheet" if specified by a "media" attribute in an inline stylesheet's STYLE tag.
	SourceURL    string        `json:"sourceURL,omitempty"`    // URL of the document containing the media query description.
	Range        *SourceRange  `json:"range,omitempty"`        // The associated rule (@media or @import) header range in the enclosing stylesheet (if available).
	StyleSheetID StyleSheetID  `json:"styleSheetId,omitempty"` // Identifier of the stylesheet containing this object (if exists).
	MediaList    []*MediaQuery `json:"mediaList,omitempty"`    // Array of media queries.
}

// MediaQuery media query descriptor.
type MediaQuery struct {
	Expressions []*MediaQueryExpression `json:"expressions"` // Array of media query expressions.
	Active      bool                    `json:"active"`      // Whether the media query condition is satisfied.
}

// MediaQueryExpression media query expression descriptor.
type MediaQueryExpression struct {
	Value          float64      `json:"value"`                    // Media query expression value.
	Unit           string       `json:"unit"`                     // Media query expression units.
	Feature        string       `json:"feature"`                  // Media query expression feature.
	ValueRange     *SourceRange `json:"valueRange,omitempty"`     // The associated range of the value text in the enclosing stylesheet (if available).
	ComputedLength float64      `json:"computedLength,omitempty"` // Computed length of media query expression (if applicable).
}

// PlatformFontUsage information about amount of glyphs that were rendered
// with given font.
type PlatformFontUsage struct {
	FamilyName   string  `json:"familyName"`   // Font's family name reported by platform.
	IsCustomFont bool    `json:"isCustomFont"` // Indicates if the font was downloaded or resolved locally.
	GlyphCount   float64 `json:"glyphCount"`   // Amount of glyphs that were rendered with this font.
}

// KeyframesRule CSS keyframes rule representation.
type KeyframesRule struct {
	AnimationName *Value          `json:"animationName"` // Animation name.
	Keyframes     []*KeyframeRule `json:"keyframes"`     // List of keyframes.
}

// KeyframeRule CSS keyframe rule representation.
type KeyframeRule struct {
	StyleSheetID StyleSheetID     `json:"styleSheetId,omitempty"` // The css style sheet identifier (absent for user agent stylesheet and user-specified stylesheet rules) this rule came from.
	Origin       StyleSheetOrigin `json:"origin"`                 // Parent stylesheet's origin.
	KeyText      *Value           `json:"keyText"`                // Associated key text.
	Style        *Style           `json:"style"`                  // Associated style declaration.
}

// StyleDeclarationEdit a descriptor of operation to mutate style declaration
// text.
type StyleDeclarationEdit struct {
	StyleSheetID StyleSheetID `json:"styleSheetId"` // The css style sheet identifier.
	Range        *SourceRange `json:"range"`        // The range of the style text in the enclosing stylesheet.
	Text         string       `json:"text"`         // New style text.
}

// MediaSource source of the media query: "mediaRule" if specified by a
// @media rule, "importRule" if specified by an @import rule, "linkedSheet" if
// specified by a "media" attribute in a linked stylesheet's LINK tag,
// "inlineSheet" if specified by a "media" attribute in an inline stylesheet's
// STYLE tag.
type MediaSource string

// String returns the MediaSource as string value.
func (t MediaSource) String() string {
	return string(t)
}

// MediaSource values.
const (
	MediaSourceMediaRule   MediaSource = "mediaRule"
	MediaSourceImportRule  MediaSource = "importRule"
	MediaSourceLinkedSheet MediaSource = "linkedSheet"
	MediaSourceInlineSheet MediaSource = "inlineSheet"
)

// MarshalEasyJSON satisfies easyjson.Marshaler.
func (t MediaSource) MarshalEasyJSON(out *jwriter.Writer) {
	out.String(string(t))
}

// MarshalJSON satisfies json.Marshaler.
func (t MediaSource) MarshalJSON() ([]byte, error) {
	return easyjson.Marshal(t)
}

// UnmarshalEasyJSON satisfies easyjson.Unmarshaler.
func (t *MediaSource) UnmarshalEasyJSON(in *jlexer.Lexer) {
	switch MediaSource(in.String()) {
	case MediaSourceMediaRule:
		*t = MediaSourceMediaRule
	case MediaSourceImportRule:
		*t = MediaSourceImportRule
	case MediaSourceLinkedSheet:
		*t = MediaSourceLinkedSheet
	case MediaSourceInlineSheet:
		*t = MediaSourceInlineSheet

	default:
		in.AddError(errors.New("unknown MediaSource value"))
	}
}

// UnmarshalJSON satisfies json.Unmarshaler.
func (t *MediaSource) UnmarshalJSON(buf []byte) error {
	return easyjson.Unmarshal(buf, t)
}
