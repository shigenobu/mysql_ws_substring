/*
 substring
*/
package _func

// 開始位置から長さ分の文字列を切り取る
func Substring(str string, start int, length int) string {
	// https://increment.hatenablog.com/entry/2016/04/17/115250
	if start < 0 || length <= 0 {
		return str
	}
	r := []rune(str)
	if start + length > len(r) {
		return string(r[start:])
	} else {
		return string(r[start:start + length])
	}
}
