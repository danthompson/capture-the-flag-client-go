/*
 * Capture the Flag
 *
 * A friendly game of Capture the Flag
 *
 * API version: 1.0.0
 * Contact: hello@testdouble.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package game-client
// Flag Opponent flag information
type Flag struct {
	Revealed bool `json:"revealed"`
	Held bool `json:"held"`
	X int32 `json:"x"`
	Y int32 `json:"y"`
}
