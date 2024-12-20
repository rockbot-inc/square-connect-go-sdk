/*
 * Square
 *
 * Use Square APIs to manage and run business including payment, customer, product, inventory, and employee management.
 *
 * API version: 2.0
 * Contact: developers@squareup.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package square

// Represents the snippet that is added to a Square Online site. The snippet code is injected into the `head` element of all pages on the site, except for checkout pages.
type Snippet struct {
	// The Square-assigned ID for the snippet.
	Id string `json:"id,omitempty"`
	// The ID of the site that contains the snippet.
	SiteId string `json:"site_id,omitempty"`
	// The snippet code, which can contain valid HTML, JavaScript, or both.
	Content string `json:"content"`
	// The timestamp of when the snippet was initially added to the site, in RFC 3339 format.
	CreatedAt string `json:"created_at,omitempty"`
	// The timestamp of when the snippet was last updated on the site, in RFC 3339 format.
	UpdatedAt string `json:"updated_at,omitempty"`
}
