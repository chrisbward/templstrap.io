package base

type ScrollSpyProps struct {
}

type HTMXSwapType string

const OuterSwapType HTMXSwapType = "outerHTML"
const InnerSwapType HTMXSwapType = "innerHTML"

type VerbType string

const PostVerbType VerbType = "post"

type HTMXProps struct {
	Verb              VerbType
	PostDestination   string
	Encoding          string
	TargetSelector    string
	IndicatorSelector string
	Swap              HTMXSwapType
	SwapOutOfBand     bool
}

type ElementProps struct {
	Id                string
	AdditionalClasses []string
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
