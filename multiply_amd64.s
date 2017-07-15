// +build amd64

// func mulq(x, y uint64) (lo uint64, hi uint64)
TEXT ·mulq(SB),4,$0
	MOVQ x+0(FP), AX
	MOVQ y+8(FP), CX
	MULQ CX
	MOVQ AX, ret+16(FP) // result low bits
	MOVQ DX, ret+24(FP) // result high bits
	RET
