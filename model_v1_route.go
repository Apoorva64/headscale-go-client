/*
 * headscale/v1/headscale.proto
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: version not set
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package headscale_api_client

import (
	"time"
)

type V1Route struct {
	Id         string     `json:"id,omitempty"`
	Machine    *V1Machine `json:"machine,omitempty"`
	Prefix     string     `json:"prefix,omitempty"`
	Advertised bool       `json:"advertised,omitempty"`
	Enabled    bool       `json:"enabled,omitempty"`
	IsPrimary  bool       `json:"isPrimary,omitempty"`
	CreatedAt  time.Time  `json:"createdAt,omitempty"`
	UpdatedAt  time.Time  `json:"updatedAt,omitempty"`
	DeletedAt  time.Time  `json:"deletedAt,omitempty"`
}
