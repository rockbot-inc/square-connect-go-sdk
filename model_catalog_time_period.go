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

// Represents a time period - either a single period or a repeating period.
type CatalogTimePeriod struct {
	// An iCalendar (RFC 5545) [event](https://tools.ietf.org/html/rfc5545#section-3.6.1), which specifies the name, timing, duration and recurrence of this time period.  Example:  ``` DTSTART:20190707T180000 DURATION:P2H RRULE:FREQ=WEEKLY;BYDAY=MO,WE,FR ```  Only `SUMMARY`, `DTSTART`, `DURATION` and `RRULE` fields are supported. `DTSTART` must be in local (unzoned) time format. Note that while `BEGIN:VEVENT` and `END:VEVENT` is not required in the request. The response will always include them.
	Event string `json:"event,omitempty"`
}
