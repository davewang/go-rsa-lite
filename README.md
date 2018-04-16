<script type="text/javascript" src="http://cdn.mathjax.org/mathjax/latest/MathJax.js?config=default"></script>
# go-ras-lite
------

RSA分为一下几部

> * 生成公私腰
> * 加密
> * 解密

![cmd-markdown-logo](https://www.zybuluo.com/static/img/logo.png)

------

## 生产公私腰
公腰{e,n}用来加密使用,私钥{d,n}用来解密使用。

### 1. 生成一个质数
golang自带随机生成质数的一些工具方法。
```go
func NewPrime(c int) *big.Int {
	p1, err := rand.Prime(rand.Reader,c)
	if err != nil {
		fmt.Println("error:", err)
		return nil
	}
	return p1

}
```

### 2. 计算N，N是两个随机质数的乘积

$$N=P1*P2$$

### 3. 计算E，E是一个奇数，不同与φ(N)的公因数（K为随机数）
$$DE=1 mod N$$$$DE=K*φ(N)+1$$$$E=(K*φ(N)+1)/D$$所以需满足$$0=(K*φ(N)+1) mod E$$

```go
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
```

### 4. 计算D，E计算出来再算D很简单了
$$D=(K*φ(N)+1)/E$$
```go
 d := big.NewInt(1).Mul(k,q) //q为φ(N)
 d.Add(d,big.NewInt(1))
 d.Div(d,e)
```
由上面4步，公私腰生成完成
## 加密
$$c=m^E mod N$$
```go
 c := big.NewInt(0).SetBytes([]byte(m))
 BigPow(c,e)
 c.Mod(c,n)
```
## 解密
 $$m=c^D mod N$$
```go
s := big.NewInt(0).SetBytes(c.Bytes())
BigPow(s,d)
s.Mod(s,n)
```
