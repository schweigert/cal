package rsa

import ("testing"
				"math/big"
			 )

func TestNewInt(t *testing.T) {
	t.Log("TestNewInt")
	a := NewNum(1)
	b := big.NewInt(1)

	if(a.Cmp(b) != 0){
		t.Fail()
	}
}

func TestMul(t *testing.T) {
	t.Log("TestMull")
	a := NewNum(1)
  b := NewNum(2)

  c := Mul(a,b)

  d := NewNum(2)

  if(c.Cmp(d) != 0){
    t.Fail()
  }
}

func TestAdd(t *testing.T) {
	t.Log("TestAdd")
	a := NewNum(1)
  b := NewNum(2)

  c := Add(a,b)

  d := NewNum(3)

  if(c.Cmp(d) != 0){
    t.Fail()
  }
}

func TestSub(t *testing.T) {
	t.Log("TestSub")
	a := NewNum(5)
	b := NewNum(8)

	c := Sub(a,b)

	d := NewNum(-3)

	if(c.Cmp(d) != 0){
		t.Fail()
	}
}

func TestDiv(t *testing.T) {
	t.Log("TestDiv")
	a := NewNum(8)
	b := NewNum(2)

	c := Div(a,b)

	d := NewNum(4)

	if(c.Cmp(d) != 0){
		t.Fail()
	}
}

func TestMod(t *testing.T) {
	t.Log("TestMod")
	a := NewNum(9)
	b := NewNum(2)

	c := Mod(a,b)

	d := NewNum(1)

	if(c.Cmp(d) != 0){
		t.Fail()
	}
}

func TestComp(t *testing.T) {
	t.Log("TestComp")
	a := NewNum(9)
	b := NewNum(2)

	if(Comp(a,b)){
		t.Fail()
	}
}
