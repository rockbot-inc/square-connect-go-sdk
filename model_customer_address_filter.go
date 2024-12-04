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

// The customer address filter. This filter is used in a [CustomerCustomAttributeFilterValue]($m/CustomerCustomAttributeFilterValue) filter when searching by an `Address`-type custom attribute.
type CustomerAddressFilter struct {
	PostalCode *CustomerTextFilter `json:"postal_code,omitempty"`
	Country *Country `json:"country,omitempty"`
}
