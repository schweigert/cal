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

func Exp(x,y,m *big.Int) *big.Int {
  r := NewNum(0)
  r.Exp(x,y,m)
  return r
}

func Sqrt(n, unity *big.Int) *big.Int {
	// Initial guess = 2^(log_2(n)/2)
	guess := NewNum(2)
	guess.Exp(guess, NewNum(int64(n.BitLen()/2)), nil)

	n_big := NewNum(0).Mul(n, unity)
	// Now refine using Newton's method
	prev_guess := NewNum(0)
	for {
		prev_guess.Set(guess)
		guess.Add(guess, NewNum(0).Div(n_big, guess))
		guess.Div(guess, NewNum(2))
		if guess.Cmp(prev_guess) == 0  {
			break
		}
	}
	return guess
}
