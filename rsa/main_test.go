package main

import (
  "./rsa"
  "testing"
)

func benchmarkBruteForce(s int, b *testing.B) {
  rsa.LEN_BIG_PRIME = s
  rsa.LEN_CRIVO = 10000
  _, pub := rsa.NewCert()
  for n := 0; n < b.N; n++ {
    bruteForce(pub)
  }
}

func BenchmarkBruteForce5(b *testing.B) { benchmarkBruteForce(5,b) }
func BenchmarkBruteForce6(b *testing.B) { benchmarkBruteForce(6,b) }
func BenchmarkBruteForce7(b *testing.B) { benchmarkBruteForce(7,b) }
func BenchmarkBruteForce8(b *testing.B) { benchmarkBruteForce(8,b) }
func BenchmarkBruteForce9(b *testing.B) { benchmarkBruteForce(9,b) }
func BenchmarkBruteForce10(b *testing.B) { benchmarkBruteForce(10,b) }
func BenchmarkBruteForce11(b *testing.B) { benchmarkBruteForce(11,b) }
func BenchmarkBruteForce12(b *testing.B) { benchmarkBruteForce(12,b) }
func BenchmarkBruteForce13(b *testing.B) { benchmarkBruteForce(13,b) }
func BenchmarkBruteForce14(b *testing.B) { benchmarkBruteForce(14,b) }
func BenchmarkBruteForce15(b *testing.B) { benchmarkBruteForce(15,b) }
func BenchmarkBruteForce16(b *testing.B) { benchmarkBruteForce(16,b) }

func TestNewInt(t *testing.T) {
  rsa.LEN_CRIVO = 10000
  rsa.LEN_BIG_PRIME = 32

  // O(n² + Klog(K)²)
  priv, pub := rsa.NewCert()
  fmt.Println(priv)
  fmt.Println(pub)
}
