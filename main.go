// ヒント？https://qiita.com/tenntenn/items/eac962a49c56b2b15ee8#%E5%9F%8B%E3%82%81%E8%BE%BC%E3%81%BF%E3%82%92%E4%BD%BF%E3%81%A3%E3%81%9F%E3%82%A4%E3%83%B3%E3%82%BF%E3%83%95%E3%82%A7%E3%83%BC%E3%82%B9%E3%81%AE%E5%8B%95%E7%9A%84%E5%AE%9F%E8%A3%85
package main

import "fmt"

type Hoge interface {
	M()
	N()
}
type fuga struct{ Hoge } // インタフェースを埋め込む

func (f fuga) M() {
	fmt.Println("Hi") // Mの振る舞いを変える
	// f.Hoge.M()        // 元のメソッドを呼ぶ
}

func hiHoge(h Hoge) Hoge {
	return fuga{h} // 構造体を作る
}
func main() {
	// ???
	var h Hoge
	hs := hiHoge(h)
	hs.M()
	hs.N()
}
