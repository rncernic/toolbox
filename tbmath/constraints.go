package tbmath

// Signed is a constraint that allows any signed integer type.
type Signed interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

// Unsigned is a constraint that allows any unsigned integer type.
type Unsigned interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
}

// Integer is a constraint that allows any integer type.
type Integer interface {
	Signed | Unsigned
}

// Float is a constraint that allows any floating-point type.
type Float interface {
	~float32 | ~float64
}

// Number allows any number type
type Number interface {
	Integer | Float
}
