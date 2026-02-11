package xconfig_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/iota101/xconfig"
)

func TestFromYAML(t *testing.T) {
	t.Parallel()

	yaml := `
server:
  host: "0.0.0.0"
  port: 3000
database:
  ssl: true
  ratio: 1.5
`
	cfg := createTempConfig(t, yaml)

	tests := []struct {
		name string
		key  xconfig.K
		want any
		get  func(xconfig.Value) any
	}{
		{"nested string", "server.host", "0.0.0.0", func(v xconfig.Value) any { return v.String() }},
		{"nested int", "server.port", 3000, func(v xconfig.Value) any { return v.Int() }},
		{"nested bool", "database.ssl", true, func(v xconfig.Value) any { return v.Bool() }},
		{"nested float", "database.ratio", 1.5, func(v xconfig.Value) any { return v.Float64() }},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.get(cfg.Get(tt.key))
			if got != tt.want {
				t.Errorf("Get(%q) = %v, want %v", tt.key, got, tt.want)
			}
		})
	}
}

func TestFromYAML_Has(t *testing.T) {
	t.Parallel()

	cfg := createTempConfig(t, `key: value`)

	tests := []struct {
		key  xconfig.K
		want bool
	}{
		{"key", true},
		{"missing", false},
	}

	for _, tt := range tests {
		t.Run(string(tt.key), func(t *testing.T) {
			if got := cfg.Has(tt.key); got != tt.want {
				t.Errorf("Has(%q) = %v, want %v", tt.key, got, tt.want)
			}
		})
	}
}

func TestFromYAML_FileNotFound(t *testing.T) {
	t.Parallel()

	_, err := xconfig.FromYAML("nonexistent.yaml")
	if err == nil {
		t.Error("expected error for non-existent file")
	}
}

func TestValue_PanicOnMissing(t *testing.T) {
	t.Parallel()

	cfg := xconfig.FromMap(map[xconfig.K]any{})

	methods := []struct {
		name string
		fn   func()
	}{
		{"String", func() { _ = cfg.Get("missing").String() }},
		{"Int", func() { _ = cfg.Get("missing").Int() }},
		{"Int64", func() { _ = cfg.Get("missing").Int64() }},
		{"Float64", func() { _ = cfg.Get("missing").Float64() }},
		{"Bool", func() { _ = cfg.Get("missing").Bool() }},
	}

	for _, tt := range methods {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				if r := recover(); r == nil {
					t.Errorf("%s() should panic on missing key", tt.name)
				}
			}()
			tt.fn()
		})
	}
}

func TestValue_PanicOnTypeMismatch(t *testing.T) {
	t.Parallel()

	cfg := xconfig.FromMap(map[xconfig.K]any{
		"str": "hello",
	})

	methods := []struct {
		name string
		fn   func()
	}{
		{"Int", func() { cfg.Get("str").Int() }},
		{"Int64", func() { cfg.Get("str").Int64() }},
		{"Float64", func() { cfg.Get("str").Float64() }},
		{"Bool", func() { cfg.Get("str").Bool() }},
	}

	for _, tt := range methods {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				if r := recover(); r == nil {
					t.Errorf("%s() should panic on type mismatch", tt.name)
				}
			}()
			tt.fn()
		})
	}
}

func TestValue_Defaults(t *testing.T) {
	t.Parallel()

	cfg := xconfig.FromMap(map[xconfig.K]any{
		"str":   "value",
		"num":   42,
		"flag":  true,
		"ratio": 3.14,
	})

	t.Run("returns value when exists", func(t *testing.T) {
		tests := []struct {
			name string
			got  any
			want any
		}{
			{"StringOr", cfg.Get("str").StringOr("def"), "value"},
			{"IntOr", cfg.Get("num").IntOr(0), 42},
			{"Int64Or", cfg.Get("num").Int64Or(0), int64(42)},
			{"Float64Or", cfg.Get("ratio").Float64Or(0), 3.14},
			{"BoolOr", cfg.Get("flag").BoolOr(false), true},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				if tt.got != tt.want {
					t.Errorf("%s = %v, want %v", tt.name, tt.got, tt.want)
				}
			})
		}
	})

	t.Run("returns default when missing", func(t *testing.T) {
		tests := []struct {
			name string
			got  any
			want any
		}{
			{"StringOr", cfg.Get("x").StringOr("def"), "def"},
			{"IntOr", cfg.Get("x").IntOr(99), 99},
			{"Int64Or", cfg.Get("x").Int64Or(99), int64(99)},
			{"Float64Or", cfg.Get("x").Float64Or(9.9), 9.9},
			{"BoolOr", cfg.Get("x").BoolOr(true), true},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				if tt.got != tt.want {
					t.Errorf("%s = %v, want %v", tt.name, tt.got, tt.want)
				}
			})
		}
	})
}

func TestValue_IsEmpty(t *testing.T) {
	t.Parallel()

	cfg := xconfig.FromMap(map[xconfig.K]any{
		"empty":    "",
		"nonempty": "value",
		"zero":     0,
	})

	tests := []struct {
		key  xconfig.K
		want bool
	}{
		{"missing", true},
		{"empty", true},
		{"nonempty", false},
		{"zero", false},
	}

	for _, tt := range tests {
		t.Run(string(tt.key), func(t *testing.T) {
			if got := cfg.Get(tt.key).IsEmpty(); got != tt.want {
				t.Errorf("Get(%q).IsEmpty() = %v, want %v", tt.key, got, tt.want)
			}
		})
	}
}

func TestValue_TypeConversions(t *testing.T) {
	t.Parallel()

	cfg := xconfig.FromMap(map[xconfig.K]any{
		"int":     42,
		"int64":   int64(100),
		"float64": 3.14,
	})

	tests := []struct {
		name string
		got  any
		want any
	}{
		{"int to int64", cfg.Get("int").Int64(), int64(42)},
		{"int to float64", cfg.Get("int").Float64(), float64(42)},
		{"int64 to int", cfg.Get("int64").Int(), 100},
		{"float64 to int", cfg.Get("float64").Int(), 3},
		{"int to string", cfg.Get("int").String(), "42"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.got != tt.want {
				t.Errorf("%s: got %v, want %v", tt.name, tt.got, tt.want)
			}
		})
	}
}

func TestFromEnv(t *testing.T) {
	const key xconfig.E = "XCONFIG_TEST_VAR"

	t.Setenv(string(key), "secret123")
	secret := xconfig.FromEnv()

	t.Run("Get", func(t *testing.T) {
		if got := secret.Get(key).String(); got != "secret123" {
			t.Errorf("Get(%q) = %q, want %q", key, got, "secret123")
		}
	})

	t.Run("Has", func(t *testing.T) {
		if !secret.Has(key) {
			t.Errorf("Has(%q) = false, want true", key)
		}
		if secret.Has("XCONFIG_MISSING") {
			t.Error("Has(missing) = true, want false")
		}
	})

	t.Run("StringOr for missing", func(t *testing.T) {
		if got := secret.Get("XCONFIG_MISSING").StringOr("default"); got != "default" {
			t.Errorf("StringOr() = %q, want %q", got, "default")
		}
	})
}

func TestFromEnvMap(t *testing.T) {
	t.Parallel()

	secret := xconfig.FromEnvMap(map[xconfig.E]string{
		"DB_PASS": "password123",
		"API_KEY": "key456",
	})

	tests := []struct {
		key  xconfig.E
		want string
	}{
		{"DB_PASS", "password123"},
		{"API_KEY", "key456"},
	}

	for _, tt := range tests {
		t.Run(string(tt.key), func(t *testing.T) {
			if got := secret.Get(tt.key).String(); got != tt.want {
				t.Errorf("Get(%q) = %q, want %q", tt.key, got, tt.want)
			}
		})
	}

	t.Run("Has", func(t *testing.T) {
		if !secret.Has("DB_PASS") {
			t.Error("Has(DB_PASS) = false, want true")
		}
		if secret.Has("MISSING") {
			t.Error("Has(MISSING) = true, want false")
		}
	})
}

func TestFromMap(t *testing.T) {
	t.Parallel()

	cfg := xconfig.FromMap(map[xconfig.K]any{
		"server.host": "localhost",
		"server.port": 8080,
		"debug":       true,
	})

	tests := []struct {
		key  xconfig.K
		want any
		get  func(xconfig.Value) any
	}{
		{"server.host", "localhost", func(v xconfig.Value) any { return v.String() }},
		{"server.port", 8080, func(v xconfig.Value) any { return v.Int() }},
		{"debug", true, func(v xconfig.Value) any { return v.Bool() }},
	}

	for _, tt := range tests {
		t.Run(string(tt.key), func(t *testing.T) {
			if got := tt.get(cfg.Get(tt.key)); got != tt.want {
				t.Errorf("Get(%q) = %v, want %v", tt.key, got, tt.want)
			}
		})
	}
}

// Helper

func createTempConfig(t *testing.T, content string) xconfig.Config {
	t.Helper()

	path := filepath.Join(t.TempDir(), "config.yaml")
	if err := os.WriteFile(path, []byte(content), 0644); err != nil {
		t.Fatalf("write temp config: %v", err)
	}

	cfg, err := xconfig.FromYAML(path)
	if err != nil {
		t.Fatalf("FromYAML: %v", err)
	}

	return cfg
}
