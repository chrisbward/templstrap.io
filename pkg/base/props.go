package base

type ScrollSpyProps struct {
}

type ElementProps struct {
	Id                string
	AdditionalClasses []string
	Style             string
	IsDisabled        bool
	IsVisible         bool
	AriaLabel         string
	AriaDescribedById string
}

type FormElementProps struct {
	ElementProps
	IsReadOnly bool
}
