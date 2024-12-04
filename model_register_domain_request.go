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

// Defines the parameters that can be included in the body of a request to the [RegisterDomain]($e/ApplePay/RegisterDomain) endpoint.
type RegisterDomainRequest struct {
	// A domain name as described in RFC-1034 that will be registered with ApplePay.
	DomainName string `json:"domain_name"`
}
