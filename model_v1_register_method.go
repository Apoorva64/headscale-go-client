/*
 * headscale/v1/headscale.proto
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: version not set
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package headscale_api_client

type V1RegisterMethod string

// List of v1RegisterMethod
const (
	UNSPECIFIED_V1RegisterMethod V1RegisterMethod = "REGISTER_METHOD_UNSPECIFIED"
	AUTH_KEY_V1RegisterMethod    V1RegisterMethod = "REGISTER_METHOD_AUTH_KEY"
	CLI_V1RegisterMethod         V1RegisterMethod = "REGISTER_METHOD_CLI"
	OIDC_V1RegisterMethod        V1RegisterMethod = "REGISTER_METHOD_OIDC"
)
