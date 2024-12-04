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

type TerminalRefundQueryFilter struct {
	// `TerminalRefund` objects associated with a specific device. If no device is specified, then all `TerminalRefund` objects for the signed-in account are displayed.
	DeviceId string `json:"device_id,omitempty"`
	CreatedAt *TimeRange `json:"created_at,omitempty"`
	// Filtered results with the desired status of the `TerminalRefund`. Options: `PENDING`, `IN_PROGRESS`, `CANCEL_REQUESTED`, `CANCELED`, or `COMPLETED`.
	Status string `json:"status,omitempty"`
}
