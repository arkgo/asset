package asset

import (
	"fmt"
	"strings"
)

var (
	VarNil = Var{nil: true}
)

type (
	Any = interface{}
	Map = map[string]Any
	Env = int
	Res struct {
		Code int
		Text string
		Args []Any
	}

	Vars map[string]Var
	Var  struct {
		nil      bool
		Type     string `json:"type"`
		Required bool   `json:"require"`
		Unique   bool   `json:"unique"`
		Nullable bool   `json:"nullable"`
		Name     string `json:"name"`
		Desc     string `json:"desc"`
		Default  Any    `json:"default"`
		Setting  Map    `json:"setting"`
		Children Vars   `json:"children"`
		Option   Map    `json:"option"`

		Empty *Res `json:"-"`
		Error *Res `json:"-"`

		Encode string              `json:"-"`
		Decode string              `json:"-"`
		Valid  func(Any, Var) bool `json:"-"`
		Value  func(Any, Var) Any  `json:"-"`
	}
)

const (
	_ Env = iota
	Developing
	Testing
	Production
)
const (
	// REMOVED = "removed"

	DELIMS = `"` //字段以及表名边界符，自己实现数据驱动才需要处理这个，必须能启标识作用
	RANDBY = `$RANDBY$`

	COUNT = "COUNT"
	SUM   = "SUM"
	MAX   = "MAX"
	MIN   = "MIN"
	AVG   = "AVG"

	IS  = "="  //等于
	NOT = "!=" //不等于
	EQ  = "="  //等于
	NE  = "!=" //不等于
	NEQ = "!=" //不等于

	//约等于	正则等于
	AE   = "~*" //正则等于，约等于
	AEC  = "~"  //正则等于，区分大小写，
	RE   = "~*" //正则等于，约等于
	REC  = "~"  //正则等于，区分大小写，
	REQ  = "~*" //正则等于，约等于
	REQC = "~"  //正则等于，区分大小写，

	NAE   = "!~*" //正则不等于，
	NAEC  = "!~"  //正则不等于，区分大小写，
	NRE   = "!~*" //正则不等于，
	NREC  = "!~"  //正则不等于，区分大小写，
	NREQ  = "!~*" //正则不等于，
	NREQC = "!~"  //正则不等于，区分大小写，

	//换位约等于，值在前，字段在后，用于黑名单查询
	EA   = "$$~*$$" //正则等于，约等于
	EAC  = "$$~$$"  //正则等于，区分大小写，
	ER   = "$$~*$$" //正则等于，约等于
	ERC  = "$$~$$"  //正则等于，区分大小写，
	EQR  = "$$~*$$" //正则等于，约等于
	EQRC = "$$~$$"  //正则等于，区分大小写，

	NEA   = "$$!~*$$" //正则不等于，
	NEAC  = "$$!~$$"  //正则不等于，区分大小写，
	NER   = "$$!~*$$" //正则不等于，
	NERC  = "$$!~$$"  //正则不等于，区分大小写，
	NEQR  = "$$!~*$$" //正则不等于，
	NEQRC = "$$!~$$"  //正则不等于，区分大小写，

	GT  = ">"  //大于
	GE  = ">=" //大于等于
	GTE = ">=" //大于等于
	LT  = "<"  //小于
	LE  = "<=" //小于等于
	LTE = "<=" //小于等于

	IN  = "$$IN$$"    //支持  WHERE id IN (1,2,3)			//这条还没支持
	NI  = "$$NOTIN$$" //支持	WHERE id NOT IN(1,2,3)
	NIN = "$$NOTIN$$" //支持	WHERE id NOT IN(1,2,3)
	ANY = "$$ANY$$"   //支持数组字段的

	SEARCH    = "$$full$$"  //like搜索
	FULLLIKE  = "$$full$$"  //like搜索
	LEFTLIKE  = "$$left$$"  //like left搜索
	RIGHTLIKE = "$$right$$" //like right搜索

	INC = "$$inc$$" //累加，    UPDATE时用，解析成：views=views+value

	BYASC  = "asc"
	BYDESC = "desc"
)

type (
	dataNil  struct{}
	dataNol  struct{}
	dataRand struct{}
	dataAsc  struct{}
	dataDesc struct{}
)

var (
	NIL  dataNil //为空	IS NULL
	NOL  dataNol //不为空	IS NOT NULL
	NULL dataNil //为空	IS NULL
	NOLL dataNol //不为空	IS NOT NULL
	RAND dataRand
	ASC  dataAsc  //正序	asc
	DESC dataDesc //倒序	desc
)

const (
	GET     = "get"
	POST    = "post"
	PUT     = "put"
	DELETE  = "delete"
	HEAD    = "head"
	PATCH   = "patch"
	OPTIONS = "options"
	TRACE   = "trace"
	CONNECT = "connect"
)

func (res *Res) OK() bool {
	if res == nil {
		return true
	}
	return res.Code == 0
}
func (res *Res) Fail() bool {
	return res.OK() == false
}
func (res *Res) Affix(args ...Any) *Res {
	res.Args = args
	return res
}

func (res *Res) Error() string {
	if res == nil {
		return ""
	}
	if res.Args != nil && len(res.Args) > 0 {
		ccc := strings.Count(res.Text, "%") - strings.Count(res.Text, "%%")
		if ccc == len(res.Args) {
			return fmt.Sprintf(res.Text, res.Args...)
		}
	}
	return res.Text
}

func (vvv *Var) Nil() bool {
	return vvv.nil
}
