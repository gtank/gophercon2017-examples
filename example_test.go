package main

import (
	"testing"
)

// Arbitrary value: 29273397577908224
// aLargeNumber^2 is 110 bits
const aLargeNumber uint64 = 3*(1<<52) + 7*(1<<51)

func BenchmarkMultiply1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = multiply1(aLargeNumber, aLargeNumber)
	}
}

func BenchmarkMultiply2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = multiply2(aLargeNumber, aLargeNumber)
	}
}

func BenchmarkMultiplyAsm(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = mulq(aLargeNumber, aLargeNumber)
	}
}

func BenchmarkMultiply3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = multiply3(aLargeNumber, aLargeNumber)
	}
}

func BenchmarkMultiply4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = multiply4(aLargeNumber, aLargeNumber)
	}
}

func BenchmarkMultiply5(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = multiply5(aLargeNumber, aLargeNumber)
	}
}

func BenchmarkMultiply6(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = multiply6(aLargeNumber, aLargeNumber)
	}
}

// #7 is #6 with a noinline directive.
// Code that looks like #7 is not actually better!
func BenchmarkMultiplyNoInline(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = multiplyNoInline(aLargeNumber, aLargeNumber)
	}
}
