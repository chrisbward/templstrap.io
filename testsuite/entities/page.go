package entities

import (
	"github.com/chrisbward/templstrap.io/pkg/components/breadcrumb/breadcrumbitem"
	"github.com/chrisbward/templstrap.io/pkg/components/navbar/navitem"
)

type IPageProps interface {
	GetNavItems(currentPageName string) []*navitem.NavItemProps
	GetBreadcrumbs() []*breadcrumbitem.BreadcrumbItemProps
	GetTitle() string
}

type PageProps struct {
	Title                     string
	CurrentPageName           string
	ComponentsNavItems        []*navitem.NavItemProps
	ComponentsNavItemsSidebar []navitem.NavItemProps
	BreadcrumbItems           []*breadcrumbitem.BreadcrumbItemProps
}

func NewPageProps(
	title string,
	currentPageName string,
	componentsNavItems []*navitem.NavItemProps,
	componentsNavItemsSidebar []navitem.NavItemProps,
	breadcrumbItems []*breadcrumbitem.BreadcrumbItemProps,
) IPageProps {

	thePage := PageProps{
		Title:                     title,
		CurrentPageName:           currentPageName,
		ComponentsNavItems:        componentsNavItems,
		ComponentsNavItemsSidebar: componentsNavItemsSidebar,
		BreadcrumbItems:           breadcrumbItems,
	}
	thePage.ComponentsNavItems = thePage.GetNavItems(currentPageName)

	return thePage
}

func (pp PageProps) GetNavItems(currentPageName string) []*navitem.NavItemProps {

	// will change this back from a pointer...
	// as http is stateless and I could leverage immutability here
	for _, navigationItem := range pp.ComponentsNavItems {
		navigationItem.IsActive = (navigationItem.Name == currentPageName)
	}

	return pp.ComponentsNavItems
}

func (pp PageProps) GetBreadcrumbs() []*breadcrumbitem.BreadcrumbItemProps {
	return pp.BreadcrumbItems
}

func (pp PageProps) GetTitle() string {
	return pp.Title
}
