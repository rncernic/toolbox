package tbmath

import (
	"testing"
)

func TestFloor(t *testing.T) {
	if got := Floor(3.8); got != 3.0 {
		t.Errorf("Floor(3.8) = %v; want 3.0", got)
	}
	if got := Floor(5); got != 5 {
		t.Errorf("Floor(5) = %v; want 5", got)
	}
}

func TestCeil(t *testing.T) {
	if got := Ceil(3.1); got != 4.0 {
		t.Errorf("Ceil(3.1) = %v; want 4.0", got)
	}
	if got := Ceil(-3.1); got != -3.0 {
		t.Errorf("Ceil(-3.1) = %v; want -3.0", got)
	}
}

func TestRound(t *testing.T) {
	tests := []struct {
		name string
		val  float64
		want float64
	}{
		{"round up", 3.5, 4.0},
		{"round down", 3.4, 3.0},
		{"negative round", -3.5, -4.0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Round(tt.val); got != tt.want {
				t.Errorf("Round(%v) = %v; want %v", tt.val, got, tt.want)
			}
		})
	}
}

func TestAbs(t *testing.T) {
	// Test Signed Integer
	if got := Abs(-10); got != 10 {
		t.Errorf("Abs(-10) = %v; want 10", got)
	}
	// Test Float
	if got := Abs(-10.5); got != 10.5 {
		t.Errorf("Abs(-10.5) = %v; want 10.5", got)
	}
	// Test Unsigned (should remain unchanged)
	var u uint = 10
	if got := Abs(u); got != 10 {
		t.Errorf("Abs(uint 10) = %v; want 10", got)
	}
}

func TestClamp(t *testing.T) {
	if got := Clamp(15, 0, 10); got != 10 {
		t.Errorf("Clamp(15) too high: got %v; want 10", got)
	}
	if got := Clamp(-5, 0, 10); got != 0 {
		t.Errorf("Clamp(-5) too low: got %v; want 0", got)
	}
	if got := Clamp(5, 0, 10); got != 5 {
		t.Errorf("Clamp(5) in range: got %v; want 5", got)
	}
}

func TestMod(t *testing.T) {
	// Test Euclidean property: result should always be positive for positive divisor
	tests := []struct {
		name string
		x    int
		m    int
		want int
	}{
		{"standard", 10, 3, 1},
		{"negative dividend", -1, 3, 2}, // -1 mod 3 = 2 in Euclidean
		{"negative both", -1, -3, 2},    // Depending on implementation, but usually pos
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Mod(tt.x, tt.m); got != tt.want {
				t.Errorf("Mod(%v, %v) = %v; want %v", tt.x, tt.m, got, tt.want)
			}
		})
	}
}

func TestModPanic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Mod did not panic on zero divisor")
		}
	}()
	Mod(10, 0)
}

func TestModMixedTypes(t *testing.T) {
	t.Run("Float64 and Int", func(t *testing.T) {
		// x is float64, m is int. Result should be float64.
		var x float64 = 5.5
		var m int = 2
		var want float64 = 1.5

		got := Mod(x, m)
		if got != want {
			t.Errorf("Mod(%T(%v), %T(%v)) = %v; want %v", x, x, m, m, got, want)
		}
	})

	t.Run("Int and Float32", func(t *testing.T) {
		// x is int, m is float32. Result should be int (truncated/cast via X(m)).
		var x int = 10
		var m float32 = 3.5 // X(m) becomes int(3.5) = 3
		var want int = 1    // 10 % 3 = 1

		got := Mod(x, m)
		if got != want {
			t.Errorf("Mod(%T(%v), %T(%v)) = %v; want %v", x, x, m, m, got, want)
		}
	})

	t.Run("Uint8 and Int64", func(t *testing.T) {
		var x uint8 = 20
		var m int64 = 6
		var want uint8 = 2

		got := Mod(x, m)
		if got != want {
			t.Errorf("Mod(%T(%v), %T(%v)) = %v; want %v", x, x, m, m, got, want)
		}
	})
}
