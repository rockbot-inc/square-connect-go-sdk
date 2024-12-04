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

// Represents a tax that applies to one or more line item in the order.  Fixed-amount, order-scoped taxes are distributed across all non-zero line item totals. The amount distributed to each line item is relative to the amount the item contributes to the order subtotal.
type OrderLineItemTax struct {
	// A unique ID that identifies the tax only within this order.
	Uid string `json:"uid,omitempty"`
	// The catalog object ID referencing [CatalogTax](entity:CatalogTax).
	CatalogObjectId string `json:"catalog_object_id,omitempty"`
	// The version of the catalog object that this tax references.
	CatalogVersion int64 `json:"catalog_version,omitempty"`
	// The tax's name.
	Name string `json:"name,omitempty"`
	Type_ *OrderLineItemTaxType `json:"type,omitempty"`
	// The percentage of the tax, as a string representation of a decimal number. For example, a value of `\"7.25\"` corresponds to a percentage of 7.25%.
	Percentage string `json:"percentage,omitempty"`
	// Application-defined data attached to this tax. Metadata fields are intended to store descriptive references or associations with an entity in another system or store brief information about the object. Square does not process this field; it only stores and returns it in relevant API calls. Do not use metadata to store any sensitive information (such as personally identifiable information or card details).  Keys written by applications must be 60 characters or less and must be in the character set `[a-zA-Z0-9_-]`. Entries can also include metadata generated by Square. These keys are prefixed with a namespace, separated from the key with a ':' character.  Values have a maximum length of 255 characters.  An application can have up to 10 entries per metadata field.  Entries written by applications are private and can only be read or modified by the same application.  For more information, see [Metadata](https://developer.squareup.com/docs/build-basics/metadata).
	Metadata map[string]string `json:"metadata,omitempty"`
	AppliedMoney *Money `json:"applied_money,omitempty"`
	Scope *OrderLineItemTaxScope `json:"scope,omitempty"`
	// Determines whether the tax was automatically applied to the order based on the catalog configuration. For an example, see [Automatically Apply Taxes to an Order](https://developer.squareup.com/docs/orders-api/apply-taxes-and-discounts/auto-apply-taxes).
	AutoApplied bool `json:"auto_applied,omitempty"`
}