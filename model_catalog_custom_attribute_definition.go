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

// Contains information defining a custom attribute. Custom attributes are intended to store additional information about a catalog object or to associate a catalog object with an entity in another system. Do not use custom attributes to store any sensitive information (personally identifiable information, card details, etc.). [Read more about custom attributes](https://developer.squareup.com/docs/catalog-api/add-custom-attributes)
type CatalogCustomAttributeDefinition struct {
	Type_ *CatalogCustomAttributeDefinitionType `json:"type"`
	//  The name of this definition for API and seller-facing UI purposes. The name must be unique within the (merchant, application) pair. Required. May not be empty and may not exceed 255 characters. Can be modified after creation.
	Name string `json:"name"`
	// Seller-oriented description of the meaning of this Custom Attribute, any constraints that the seller should observe, etc. May be displayed as a tooltip in Square UIs.
	Description string `json:"description,omitempty"`
	SourceApplication *SourceApplication `json:"source_application,omitempty"`
	// The set of `CatalogObject` types that this custom atttribute may be applied to. Currently, only `ITEM`, `ITEM_VARIATION`, `MODIFIER`, `MODIFIER_LIST`, and `CATEGORY` are allowed. At least one type must be included. See [CatalogObjectType](#type-catalogobjecttype) for possible values
	AllowedObjectTypes []CatalogObjectType `json:"allowed_object_types"`
	SellerVisibility *CatalogCustomAttributeDefinitionSellerVisibility `json:"seller_visibility,omitempty"`
	AppVisibility *CatalogCustomAttributeDefinitionAppVisibility `json:"app_visibility,omitempty"`
	StringConfig *CatalogCustomAttributeDefinitionStringConfig `json:"string_config,omitempty"`
	NumberConfig *CatalogCustomAttributeDefinitionNumberConfig `json:"number_config,omitempty"`
	SelectionConfig *CatalogCustomAttributeDefinitionSelectionConfig `json:"selection_config,omitempty"`
	// The number of custom attributes that reference this custom attribute definition. Set by the server in response to a ListCatalog request with `include_counts` set to `true`.  If the actual count is greater than 100, `custom_attribute_usage_count` will be set to `100`.
	CustomAttributeUsageCount int32 `json:"custom_attribute_usage_count,omitempty"`
	// The name of the desired custom attribute key that can be used to access the custom attribute value on catalog objects. Cannot be modified after the custom attribute definition has been created. Must be between 1 and 60 characters, and may only contain the characters `[a-zA-Z0-9_-]`.
	Key string `json:"key,omitempty"`
}
