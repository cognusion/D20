package main

import (
	"testing"
)

func BenchmarkRandString2(b *testing.B) {

	var everything = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz_-!$%^&();:.,<>/?"

	for i := 0; i < b.N; i++ {
		randString(2, everything)
	}
}

func BenchmarkRandString20(b *testing.B) {

	var everything = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz_-!$%^&();:.,<>/?"

	for i := 0; i < b.N; i++ {
		randString(20, everything)
	}
}

func BenchmarkRandString200(b *testing.B) {

	var everything = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz_-!$%^&();:.,<>/?"

	for i := 0; i < b.N; i++ {
		randString(200, everything)
	}
}

func BenchmarkRandBytes2(b *testing.B) {

	for i := 0; i < b.N; i++ {
		randBytes(2)
	}
}

func BenchmarkRandBytes20(b *testing.B) {

	for i := 0; i < b.N; i++ {
		randBytes(20)
	}
}

func BenchmarkRandBytes200(b *testing.B) {

	for i := 0; i < b.N; i++ {
		randBytes(200)
	}
}

func BenchmarkGenerateString20(b *testing.B) {

	var everything = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz_-!$%^&();:.,<>/?"

	gc := generateConfig{
		chars:     everything,
		length:    20,
		blocksize: 0,
		b64:       false,
		mangle:    "",
	}
	for i := 0; i < b.N; i++ {
		generateString(gc)
	}
}
