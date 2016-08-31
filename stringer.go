package main


import (
	"fmt"
	"strings"
)

type Person struct {
	Name string
	Age int
}



// String() string
// fmtに定義されているStringerインターフェースを満たす
func (p Person) String() string {
	return fmt.Sprintf("%v (%v yers)", p.Name, p.Age)
}


func main() {
	a := Person{"Arther Dent", 42}
	z := Person{"Zaphod Beeblebrex", 9001}

	// String() string が定義されていてStringerを満たすので
	// fmt.Printlnがa.String()を使って文字列を取得して表示してるっぽい
	fmt.Println(a, z)


	main2()
}


/* IPaddr型の実装 */
type IPaddr [4]byte

func (addr IPaddr) String() string {
	// 確認:goの配列は固定長
	// 配列の長さを指定しない宣言はスライスになる

	//var str []string // sliceの宣言
	str := make([]string, 0, 4) // はじめからキャパを取っていてもいい、少し早くなるはず
	for _, x := range addr {
		str = append(str, fmt.Sprintf("%v", x))
		// byte(数値)をそのまま文字列にするもっといい方法がありそう
	}
	return strings.Join(str, ".")

	// はーどこーでぃんぐ 多分速度はこちらが早い
	//return fmt.Sprintf("%v.%v.%v.%v",addr[0],addr[1],addr[2],addr[3])
}


func main2() {
	hosts := map[string]IPaddr{
		"loopback": {127,0,0,1},
		"googleDNS": {8,8,8,8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}

}

