package main

import (
	"crypto/rand"
	"fmt"
	"math/big"

	"time"
)

func NewPrime(c int) *big.Int {
	p1, err := rand.Prime(rand.Reader,c)
	if err != nil {
		fmt.Println("error:", err)
		return nil
	}
	return p1

}
func BigPow(c *big.Int,m *big.Int)  {
	n := m.Int64()-1
	m1 := big.NewInt(c.Int64())
	for n>0 {
		c.Mul(c,m1)
		n-=1
	}

}
func RandExponent(q *big.Int,k int,c int64) *big.Int{
	c = 1
	b := make([]byte, c)
	rand.Read(b)
	e := new(big.Int).SetBytes(b)
	r := new(big.Int).GCD(nil, nil, e, q)
	t := time.Now();
	for true {
		if r.Int64() == 1 && e.Int64() != 1{
			mod :=  new(big.Int).SetBytes(q.Bytes())
			mod.Mul(mod,big.NewInt(int64(k)) )
			mod.Add(mod,big.NewInt((1)))
			mod.Mod(mod,e)
			if mod.Int64() == 0 {
				fmt.Printf("e = %v  speed %v \n",e,time.Now().Sub(t))
				return e
			}
		}
		b = make([]byte, c)
		rand.Read(b)
		e = new(big.Int).SetBytes(b)
		r = new(big.Int).GCD(nil, nil, e, q)

	}
    return nil
}
func main() {


	co := 9
	//生产2个质数
	p1 := NewPrime(co)
	p2 := NewPrime(co)
	for p1.Int64() == p2.Int64() {
		p2 = NewPrime(co)
	}
	//n = p1 * p2
	n := big.NewInt(0).Mul(p1,p2)
	//q = φ(n) = (p1-1) * (p2-1)
	q := big.NewInt(0).Mul(big.NewInt(0).Add(p1,big.NewInt(-1)),big.NewInt(0).Add(p2,big.NewInt(-1)))
	//k 随机数
	k := 3
	fmt.Printf("p1: %v \n", p1)
	fmt.Printf("p2: %v \n", p2)
	fmt.Printf("n: %v \n", n)
	fmt.Printf("q: %v \n", q)
	fmt.Printf("k: %v \n", k)

    //e 随机数 必须为奇数 不同与φ(n)的公因数 并满足 (k * q + 1) mod e = 0
    e := RandExponent(q, k,int64(co))
    //d = (k * q + 1)/e
    d := big.NewInt(1).Mul(big.NewInt(int64(k)),q)
    d.Add(d,big.NewInt(1))
    d.Div(d,e)

    //m 内容
    m := "hi"
    //c = m^e mod n 加密
    c := big.NewInt(0).SetBytes([]byte(m))
    BigPow(c,e)
	c.Mod(c,n)

	// m = c^d mod n 解密
	s := big.NewInt(0).SetBytes(c.Bytes())
	BigPow(s,d)
	s.Mod(s,n)


	fmt.Printf("e: %v \n", e)
	fmt.Printf("m: %v \n", m)
	fmt.Printf("d: %v \n", d)
	fmt.Printf("c: %v \n", c)
	fmt.Printf("s: %v \n", string(s.Bytes()))




}