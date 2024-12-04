/*
 * Square
 *
 * Use Square APIs to manage and run business including payment, customer, product, inventory, and employee management.
 *
 * API version: 2.0
 * Contact: developers@squareup.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// A category to which a `CatalogItem` instance belongs.
type CatalogCategory struct {
	// The category name. This is a searchable attribute for use in applicable query filters, and its value length is of Unicode code points.
	Name string `json:"name,omitempty"`
	// The IDs of images associated with this `CatalogCategory` instance. Currently these images are not displayed by Square, but are free to be displayed in 3rd party applications.
	ImageIds []string `json:"image_ids,omitempty"`
	CategoryType *CatalogCategoryType `json:"category_type,omitempty"`
	ParentCategory *CatalogObjectCategory `json:"parent_category,omitempty"`
	// Indicates whether a category is a top level category, which does not have any parent_category.
	IsTopLevel bool `json:"is_top_level,omitempty"`
	// A list of IDs representing channels, such as a Square Online site, where the category can be made visible.
	Channels []string `json:"channels,omitempty"`
	// The IDs of the `CatalogAvailabilityPeriod` objects associated with the category.
	AvailabilityPeriodIds []string `json:"availability_period_ids,omitempty"`
	// Indicates whether the category is visible (`true`) or hidden (`false`) on all of the seller's Square Online sites.
	OnlineVisibility bool `json:"online_visibility,omitempty"`
	// The top-level category in a category hierarchy.
	RootCategory string `json:"root_category,omitempty"`
	EcomSeoData *CatalogEcomSeoData `json:"ecom_seo_data,omitempty"`
	// The path from the category to its root category. The first node of the path is the parent of the category and the last is the root category. The path is empty if the category is a root category.
	PathToRoot []CategoryPathToRootNode `json:"path_to_root,omitempty"`
}
