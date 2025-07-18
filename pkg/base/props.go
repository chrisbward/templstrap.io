package base

import (
	"github.com/a-h/templ"
)

type ScrollSpyProps struct {
}

type HTMXSwapType string

const OuterSwapType HTMXSwapType = "outerHTML"
const InnerSwapType HTMXSwapType = "innerHTML"

type VerbType string

const PostVerbType VerbType = "post"
const GetVerbType VerbType = "get"

type HTMXVersion string

var (
	HTMX_V1 HTMXVersion = "1"
	HTMX_V2 HTMXVersion = "2"
)

type HTMXProps struct {
	HTMXVersion        HTMXVersion
	IsBoosted          bool
	Verb               VerbType
	PostDestination    string
	DeleteDestination  string
	Encoding           string
	TargetSelector     string
	IndicatorSelector  string
	Swap               HTMXSwapType
	SwapOutOfBand      bool
	Include            string
	Trigger            string
	PushURL            bool
	ReplaceURL         string
	On                 string
	HasHTML5Validation bool
}

func (hp *HTMXProps) GetAttributes() (htmxAttributes templ.Attributes) {
	htmxAttributes = make(templ.Attributes)

	if hp.IsBoosted {
		htmxAttributes["hx-boost"] = "true"
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
	if hp.ReplaceURL != "" {
		htmxAttributes["hx-replace-url"] = hp.ReplaceURL
	}
	if hp.DeleteDestination != "" {
		htmxAttributes["hx-delete"] = hp.DeleteDestination
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
	ExtraAttributes   templ.Attributes
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
