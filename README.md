# Qin

Qin is a library for parsing Chinese address string into meaning full array. The name qin is inspired by the Qin dynasty which established the prefectures system.

[![Build Status](https://api.travis-ci.org/athom/qin.png?branch=master)](https://travis-ci.org/athom/qin)
[![GoDoc](https://godoc.org/github.com/athom/qin?status.png)](http://godoc.org/github.com/athom/qin)


## Installation

```bash
	go get "github.com/athom/qin"
```

## Useage

```go
	in := "江苏省南京市三牌楼校区"
	out := Parse(in)
	fmt.Println(out)
	// Output:
	// [江苏 南京]
```


## License

Released under the [WTFPL License](http://www.wtfpl.net/txt/copying).


