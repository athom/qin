package qin

import (
	"fmt"
	"testing"

	//. "github.com/smartystreets/goconvey/convey"
)

type tc struct {
	in  string
	out []string
}

func TestSplit(t *testing.T) {
	var testcases = []tc{
		tc{
			" 江苏省",
			[]string{
				"江苏",
			},
		},
		tc{
			" 江苏",
			[]string{
				"江苏",
			},
		},
		tc{
			" 江苏\t ",
			[]string{
				"江苏",
			},
		},
		tc{
			"\t江苏省\t",
			[]string{
				"江苏",
			},
		},
		tc{
			" 江苏省 ",
			[]string{
				"江苏",
			},
		},
		tc{
			"江苏省,南京市",
			[]string{
				"江苏",
				"南京",
			},
		},
		tc{
			"江苏省 ,南京市",
			[]string{
				"江苏",
				"南京",
			},
		},
		tc{
			"江苏省, 南京市",
			[]string{
				"江苏",
				"南京",
			},
		},
		tc{
			"江苏省 , 南京市",
			[]string{
				"江苏",
				"南京",
			},
		},
		tc{
			" 江苏省 , 南京市",
			[]string{
				"江苏",
				"南京",
			},
		},
		tc{
			" 江苏省, 南京市 ",
			[]string{
				"江苏",
				"南京",
			},
		},
		tc{
			" 江苏省\t , \t南京市 ",
			[]string{
				"江苏",
				"南京",
			},
		},
		tc{
			"广东省深圳市福田区车公庙",
			[]string{
				"广东",
				"深圳",
			},
		},
		tc{
			"广东深圳市福田区车公庙",
			[]string{
				"广东",
				"深圳",
			},
		},
		tc{
			"广东省深圳市福田区车公庙",
			[]string{
				"广东",
				"深圳",
			},
		},
		tc{
			"广东深圳福田区车公庙",
			[]string{
				"广东",
				"深圳",
			},
		},
		tc{
			" 广东 深圳 福田区车公庙",
			[]string{
				"广东",
				"深圳",
			},
		},
		tc{
			"广东 深圳 福田区车公庙",
			[]string{
				"广东",
				"深圳",
			},
		},
		tc{
			"广东\t 深圳 福田区车公庙",
			[]string{
				"广东",
				"深圳",
			},
		},
	}

	for _, tc := range testcases {
		r := Parse(tc.in)
		for i, out := range tc.out {
			if out != r[i] {
				t.Errorf("Parse(%v) want (%v), but got (%v)\n", tc.in, out, r[i])
			}
		}
	}
}

func ExampleParse() {
	in := "江苏省南京市三牌楼校区"
	out := Parse(in)
	fmt.Println(out)
	// Output:
	// [江苏 南京]
}

func ExampleParseComma() {
	in := "江苏省,南京市, 三牌楼校区"
	out := Parse(in)
	fmt.Println(out)
	// Output:
	// [江苏 南京]
}

func ExampleParseProvinceOnly() {
	in := "上海徐汇区番禺路655号7#619"
	out := Parse(in)
	fmt.Println(out)
	// Output:
	// [上海]
}
