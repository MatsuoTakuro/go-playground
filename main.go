// 関数handleFunc型の周りの簡略化

// 公式ドキュメント https://pkg.go.dev/net/http#pkg-index
// 参考記事 https://qiita.com/tenntenn/items/eac962a49c56b2b15ee8

package main

import "fmt"

// (http.Handlerの簡略版)
// Doメソッドを持つインタフェース型A
type A interface {
	Do()
}

// (http.HandlerFuncの簡略版)
// func()という関数型B
type B func()

// B型の関数bのDoメソッド
// つまり、B型の関数("b")は、インタフェースA型を実装している
// さらに、メソッド内で自身の関数("b()")を呼び出している
func (b B) Do() {
	b()
}

// (http.Handleの簡略版)
// 関数C
// 引数にA型のインタフェース("a")を持つ
func C(s string, a A) {
	a.Do()
	fmt.Println("C func is called")
}
func main() {
	// B型の関数はA型のインタフェースを実装しているため、
	// A型のインタフェースにB型の関数を代入できる
	// 右辺では、"fmt.Println("B func is called")"という値を持つ無名関数を、関数型Bの関数リテラルとしてキャストしている
	// キャストされた関数型BをA型のインタフェースaに代入している
	var a A = B(func() {
		fmt.Println("B func is called")
	})
	C("/", a)
	// このmain関数での処理を簡略化した関数が、http.HandleFunc

}
