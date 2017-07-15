// These are all variants of an unsigned 64x64->128 multiplier written for the
// purpose of demonstrating Go's code inlining behavior.
//
// For a general implementation, see mulWW in math/big/arith.go
package main

// This has an AST visitor cost of 99.
func multiply1(a, b uint64) [2]uint64 {
	al := a & 0xFFFFFFFF
	ah := a >> 32
	bl := b & 0xFFFFFFFF
	bh := b >> 32

	c0 := (al * bl) >> 32
	t1 := ah*bl + c0
	t1_lo := t1 & 0xFFFFFFFF
	c1 := t1 >> 32

	t2 := al*bh + t1_lo
	c2 := t2 >> 32

	hi := ah*bh + c1 + c2
	lo := (a * b)

	return [2]uint64{lo, hi}
}

// This costs 97
func multiply2(a, b uint64) (uint64, uint64) {
	al := a & 0xFFFFFFFF
	ah := a >> 32
	bl := b & 0xFFFFFFFF
	bh := b >> 32

	c0 := (al * bl) >> 32
	t1 := ah*bl + c0
	t1_lo := t1 & 0xFFFFFFFF
	c1 := t1 >> 32

	t2 := al*bh + t1_lo
	c2 := t2 >> 32

	hi := ah*bh + c1 + c2
	lo := (a * b)

	return lo, hi
}

// This costs 77 - it inlines! But we can do "better"
func multiply3(a, b uint64) (uint64, uint64) {
	al := a & 0xFFFFFFFF
	ah := a >> 32
	bl := b & 0xFFFFFFFF
	bh := b >> 32

	t1 := ah*bl + ((al * bl) >> 32)
	t2 := al*bh + (t1 & 0xFFFFFFFF)
	hi := ah*bh + t1>>32 + t2>>32
	lo := (a * b)

	return lo, hi
}

// This costs 71, but is still pretty readable.
func multiply4(a, b uint64) (lo uint64, hi uint64) {
	al := a & 0xFFFFFFFF
	ah := a >> 32
	bl := b & 0xFFFFFFFF
	bh := b >> 32

	t1 := ah*bl + ((al * bl) >> 32)
	t2 := al*bh + (t1 & 0xFFFFFFFF)
	hi = ah*bh + t1>>32 + t2>>32
	lo = (a * b)

	return
}

// Now we're at 59.
func multiply5(a, b uint64) (lo uint64, hi uint64) {
	t1 := (a>>32)*(b&0xFFFFFFFF) + ((a & 0xFFFFFFFF) * (b & 0xFFFFFFFF) >> 32)
	t2 := (a&0xFFFFFFFF)*(b>>32) + (t1 & 0xFFFFFFFF)
	hi = (a>>32)*(b>>32) + t1>>32 + t2>>32
	lo = (a * b)
	return
}

// 55. Don't actually write code like this.
func multiply6(a, b uint64) (uint64, uint64) {
	t1 := (a>>32)*(b&0xFFFFFFFF) + ((a & 0xFFFFFFFF) * (b & 0xFFFFFFFF) >> 32)
	t2 := (a&0xFFFFFFFF)*(b>>32) + (t1 & 0xFFFFFFFF)
	return (a * b), (a>>32)*(b>>32) + t1>>32 + t2>>32
}

//go:noinline
func multiplyNoInline(a, b uint64) (uint64, uint64) {
	t1 := (a>>32)*(b&0xFFFFFFFF) + ((a & 0xFFFFFFFF) * (b & 0xFFFFFFFF) >> 32)
	t2 := (a&0xFFFFFFFF)*(b>>32) + (t1 & 0xFFFFFFFF)
	return (a * b), (a>>32)*(b>>32) + t1>>32 + t2>>32
}
