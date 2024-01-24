package entities

import (
	"github.com/chrisbward/templstrap.io/pkg/components/breadcrumb/breadcrumbitem"
	"github.com/chrisbward/templstrap.io/pkg/components/navbar/navitem"
)

type IPageProps interface {
	GetBreadcrumbs() []breadcrumbitem.BreadcrumbItemProps
}

type PageProps struct {
	Title                     string
	ComponentsNavItems        []navitem.NavItemProps
	ComponentsNavItemsSidebar []navitem.NavItemProps
	BreadcrumbItems           []breadcrumbitem.BreadcrumbItemProps
}

func NewPageProps(
	title string,
	componentsNavItems []navitem.NavItemProps,
	componentsNavItemsSidebar []navitem.NavItemProps,
	breadcrumbItems []breadcrumbitem.BreadcrumbItemProps,
) IPageProps {
	return &PageProps{
		Title:                     title,
		ComponentsNavItems:        componentsNavItems,
		ComponentsNavItemsSidebar: componentsNavItemsSidebar,
		BreadcrumbItems:           breadcrumbItems,
	}
}

func (pp PageProps) GetBreadcrumbs() []breadcrumbitem.BreadcrumbItemProps {
	return pp.BreadcrumbItems
}
