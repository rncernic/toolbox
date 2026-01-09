package tbmath

import (
	"math"
)

const (
	// Degrees to radians
	DegToRad = math.Pi / 180
	// Radians to degrees
	RadToDeg = 180 / math.Pi
)

// Signum returns the sign of number
func Signum[X Signed | Float](x X) X {
	if x > 0 {
		return X(1)
	} else if x == 0 {
		return X(0)
	} else {
		return X(-1)
	}
}

// Floor returns the greatest integer value ≤ x.
// For integer inputs, it returns x unchanged.
func Floor[T Number](x T) T {
	switch v := any(x).(type) {
	case float32:
		return T(math.Floor(float64(v)))
	case float64:
		return T(math.Floor(v))
	default:
		return x
	}
}

// IntFloor returns Floor as integer
func IntFloor[T Number](x T) int {
	return int(Floor(x))
}

// Ceil returns the smallest integer value ≥ x.
// For integer inputs, it returns x unchanged.
func Ceil[T Number](x T) T {
	switch v := any(x).(type) {
	case float32:
		return T(math.Ceil(float64(v)))
	case float64:
		return T(math.Ceil(v))
	default:
		return x
	}
}

// Round rounds to the nearest integer, halves away from zero.
// Integer inputs are returned unchanged.
func Round[T Number](x T) T {
	switch v := any(x).(type) {
	case float32:
		return T(math.Round(float64(v)))
	case float64:
		return T(math.Round(v))
	default:
		return x
	}
}

// Trunc removes the fractional part of x.
func Trunc[T Number](x T) T {
	switch v := any(x).(type) {
	case float32:
		return T(math.Trunc(float64(v)))
	case float64:
		return T(math.Trunc(v))
	default:
		return x
	}
}

// Abs returns the absolute value of x.
// Unsigned integers are returned unchanged.
func Abs[T Number](x T) T {
	switch v := any(x).(type) {
	case float32:
		return T(math.Abs(float64(v)))
	case float64:
		return T(math.Abs(v))
	default:
		if x < 0 {
			return -x
		}
		return x
	}
}

// Clamp restricts x to the range [min, max]
func Clamp[T Number](x, min, max T) T {
	if x < min {
		return min
	}
	if x > max {
		return max
	}
	return x
}

// Int converts a numeric type to int.
// Floating values are rounded.
func Int[T Number](x T) int {
	switch v := any(x).(type) {
	case float32:
		return int(math.Round(float64(v)))
	case float64:
		return int(math.Round(v))
	default:
		return int(x)
	}
}

// Mod computes the Euclidean modulo of x mod m.
//
// mod(x,m)=x−m⋅⌊x/m⌋
//
// Mixed numeric types are allowed.
// The return type is the type of x.
//
// Panics if m == 0.
func Mod[X Number, M Number](x X, m M) X {
	// Convert modulus to type X
	mx := X(m)

	if mx == 0 {
		panic("tbmath.Mod: modulo by zero")
	}

	switch xv := any(x).(type) {

	case float32:
		mv := float64(mx)
		r := float64(xv) - mv*math.Floor(float64(xv)/mv)
		return X(r)

	case float64:
		mv := float64(mx)
		r := xv - mv*math.Floor(xv/mv)
		return X(r)

	case int:
		mv := int(mx)
		r := xv % mv
		if r < 0 {
			if mv < 0 {
				r -= mv // Subtracting a negative is adding a positive
			} else {
				r += mv
			}
		}
		return X(r)

	case int8:
		r := xv % int8(mx)
		if r < 0 {
			r += int8(mx)
		}
		return X(r)

	case int16:
		r := xv % int16(mx)
		if r < 0 {
			r += int16(mx)
		}
		return X(r)

	case int32:
		r := xv % int32(mx)
		if r < 0 {
			r += int32(mx)
		}
		return X(r)

	case int64:
		r := xv % int64(mx)
		if r < 0 {
			r += int64(mx)
		}
		return X(r)

	case uint:
		return X(xv % uint(mx))
	case uint8:
		return X(xv % uint8(mx))
	case uint16:
		return X(xv % uint16(mx))
	case uint32:
		return X(xv % uint32(mx))
	case uint64:
		return X(xv % uint64(mx))
	case uintptr:
		return X(xv % uintptr(mx))

	default:
		panic("tbmath.Mod: unsupported type")
	}
}

func Atan2Deg(y, x float64) float64 {
	return math.Atan2(y*DegToRad, x*DegToRad) * RadToDeg
}
