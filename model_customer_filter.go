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

// Represents the filtering criteria in a [search query]($m/CustomerQuery) that defines how to filter customer profiles returned in [SearchCustomers]($e/Customers/SearchCustomers) results.
type CustomerFilter struct {
	CreationSource *CustomerCreationSourceFilter `json:"creation_source,omitempty"`
	CreatedAt *TimeRange `json:"created_at,omitempty"`
	UpdatedAt *TimeRange `json:"updated_at,omitempty"`
	EmailAddress *CustomerTextFilter `json:"email_address,omitempty"`
	PhoneNumber *CustomerTextFilter `json:"phone_number,omitempty"`
	ReferenceId *CustomerTextFilter `json:"reference_id,omitempty"`
	GroupIds *FilterValue `json:"group_ids,omitempty"`
	CustomAttribute *CustomerCustomAttributeFilters `json:"custom_attribute,omitempty"`
	SegmentIds *FilterValue `json:"segment_ids,omitempty"`
}
