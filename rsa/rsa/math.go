package rsa

import "math/big"

func NewNum(a int64) *big.Int {
  return big.NewInt(a)
}

func Mul(a,b *big.Int) *big.Int {
  r := big.NewInt(0)
  r.Mul(a,b)
  return r
}

func Add(a,b *big.Int) *big.Int {
  r := big.NewInt(0)
  r.Add(a,b)
  return r
}

func Sub(a,b *big.Int) *big.Int {
  r := big.NewInt(0)
  r.Sub(a,b)
  return r
}

func Div(a,b *big.Int) *big.Int {
  r := big.NewInt(0)
  r.Div(a,b)
  return r
}

func Mod(a,b *big.Int) *big.Int {
  r := big.NewInt(0)
  r.Mod(a,b)
  return r
}

func Comp(a,b *big.Int) bool {
  if(a.Cmp(b) == 0) {
    return true
  }
  return false
}

func CompEqualZero(a *big.Int) bool {
  b := NewNum(0);
  if(Comp(a,b)) {
    return true
  }
  return false
}

func CompNotEqualZero(a *big.Int) bool {
  b := NewNum(0);
  if(!Comp(a,b)) {
    return true
  }
  return false
}

func CompLess(a, b *big.Int) bool {
  if(a.Cmp(b) < 0) {
    return true
  }
  return false
}

func CompBigger(a, b *big.Int) bool {
  if(a.Cmp(b) > 0) {
    return true
  }
  return false
}

func CompBiggerZero(a *big.Int) bool {
  b := NewNum(0)

  if(CompBigger(a,b)){
    return true
  }
  return false
}

func CompLessZero(a *big.Int) bool {
  b := NewNum(0)

  if(CompLess(a,b)){
    return true
  }
  return false
}
