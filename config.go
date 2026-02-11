// Package xconfig provides configuration loading from YAML files and environment variables.
package xconfig

type K string // Config key (YAML path with dot notation: "database.host")
type E string // Env key (environment variable name: "DATABASE_PASSWORD")

type Value interface {
	String() string
	Int() int
	Int64() int64
	Float64() float64
	Bool() bool

	StringOr(def string) string
	IntOr(def int) int
	Int64Or(def int64) int64
	Float64Or(def float64) float64
	BoolOr(def bool) bool

	IsEmpty() bool
}

type Config interface {
	Get(key K) Value
	Has(key K) bool
}

type Secret interface {
	Get(key E) Value
	Has(key E) bool
}
