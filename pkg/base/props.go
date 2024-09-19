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
	IsBoosted         bool
	Verb              VerbType
	PostDestination   string
	Encoding          string
	TargetSelector    string
	IndicatorSelector string
	Swap              HTMXSwapType
	SwapOutOfBand     bool
	Include           string
	Trigger           string
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
}

type FormElementProps struct {
	ElementProps
	IsReadOnly bool
}
