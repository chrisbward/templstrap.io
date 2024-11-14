package base

import "github.com/a-h/templ"

type ScrollSpyProps struct {
}

type HTMXSwapType string

const OuterSwapType HTMXSwapType = "outerHTML"
const InnerSwapType HTMXSwapType = "innerHTML"

type VerbType string

const PostVerbType VerbType = "post"
const GetVerbType VerbType = "get"

type HTMXProps struct {
	IsBoosted          bool
	Verb               VerbType
	PostDestination    string
	Encoding           string
	TargetSelector     string
	IndicatorSelector  string
	Swap               HTMXSwapType
	SwapOutOfBand      bool
	Include            string
	Trigger            string
	PushURL            bool
	On                 string
	HasHTML5Validation bool
}

func (hp *HTMXProps) GetAttributes() (htmxAttributes map[string]any) {
	htmxAttributes = make(map[string]any)
	if hp.IsBoosted {
		htmxAttributes["hx-boosted"] = "true"
	}
	if hp.Encoding != "" {
		htmxAttributes["hx-encoding"] = hp.Encoding
	}
	if hp.Verb == PostVerbType {
		htmxAttributes["hx-post"] = hp.PostDestination + "?requesttype=htmx"
	}
	if hp.Verb == GetVerbType {
		htmxAttributes["hx-get"] = "?requesttype=htmx"
		if hp.PushURL {
			htmxAttributes["hx-push-url"] = "true"
		}
	}
	if hp.TargetSelector != "" {
		htmxAttributes["hx-target"] = hp.TargetSelector
	}
	if hp.IndicatorSelector != "" {
		htmxAttributes["hx-indicator"] = hp.IndicatorSelector
	}
	if hp.Swap != "" {
		htmxAttributes["hx-swap"] = string(hp.Swap)
	}
	if hp.Include != "" {
		htmxAttributes["hx-include"] = hp.Include
	}
	if hp.Trigger != "" {
		htmxAttributes["hx-trigger"] = hp.Trigger
	}
	if hp.On != "" {
		htmxAttributes["hx-on"] = hp.On
	}
	if hp.HasHTML5Validation {
		htmxAttributes["hx-validate"] = "true"
	}

	return
}

type ElementProps struct {
	Id                string
	AdditionalClasses []string
	DataAttributes    templ.Attributes
	Style             string
	IsDisabled        bool
	IsVisible         bool
	AriaLabel         string
	AriaDescribedById string
	HTMX              HTMXProps
	ExtraAttributes   map[string]any
}

type FormElementProps struct {
	ElementProps
	IsReadOnly bool
	Value      string
}

type FormListProps struct {
	Text  string
	Value string
}
