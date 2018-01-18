package main

import (
	"fmt"
	"gopp"
	"io/ioutil"
	"log"
	"mkuse/cffiqt/qtsymbols"
	"os"
	"strings"

	simplejson "github.com/bitly/go-simplejson"
	"github.com/clbanning/mxj"
	"github.com/kitech/godsts/maps/hashmap"
)

func main_a() {
	log.Println("begin ...")
	fp, err := os.Open("./docs/5.9.0/qtcore.index")
	gopp.ErrPrint(err)

	m, err := mxj.NewMapXmlReader(fp)
	gopp.ErrPrint(err, m)

	log.Println(m.Root())

	jbcc, err := m.Json()
	gopp.ErrPrint(err, len(jbcc))

	jso, err := simplejson.NewJson(jbcc)
	gopp.ErrPrint(err)
	_ = jso

	cp := NewCodePager()
	cp.AddPointer("import")
	cp.AddPointer("body")

	cp.Append("import", "package gentmp")

	foundsigns := []string{}
	foundsymcnt := 0
	totalsymcnt := 0
	classcnt := 0
	// jsop := gopp.NewJsonFromObject(jso)
	log.Println(len(jso.Get("INDEX").Get("namespace").MustMap()))
	classeso := jso.Get("INDEX").Get("namespace").Get("class")
	for i, v := range classeso.MustArray() {
		_, _ = i, v
		ci := UnmarshalClass(classeso.GetIndex(i))
		if ci.Access == "private" {
			continue
		}
		classcnt += 1
		// log.Println(i, v == nil, ci)
		methodso := classeso.GetIndex(i).GetPath("function")
		for j, v := range methodso.MustArray() {
			_, _ = j, v
			// log.Println(j, v == nil, ci.Name, methodso.GetIndex(j).Get("-name").MustString())
			mi := UnmarshalMethod(methodso.GetIndex(j), ci)
			if mi.Access != "public" {
				continue
			}
			if FilterMethod(mi) {
				// log.Println("filtered:", mi.Signature)
				// continue
			}
			symname, chksym := qtsymbols.Qtsymbolsbi.GetKey(mi.SymSignature)
			_, _ = symname, chksym
			foundsymcnt += gopp.IfElseInt(chksym, 1, 0)
			totalsymcnt += 1
			if chksym {
				// log.Println(ci.Name, chksym, symname, mi.Access, mi.Status, mi.SymSignature)
				foundsigns = append(foundsigns, mi.Signature)
			}

			cp.Appendv("body", "// ", chksym, mi.FullName, mi.Signature)
			{
				overloadSuffix := gopp.IfElseStr(mi.OverloadNumber == 0, "", fmt.Sprintf("_%d", mi.OverloadNumber))
				if strings.HasPrefix(mi.Name, "operator") {
					valiname := rewriteOperatorMethodName(mi.Name)
					// log.Println(mi.Name, "=>", valiname)
					cp.Append("body", fmt.Sprintf("func %s_%s%s(){}",
						ci.Name, valiname, overloadSuffix))
				} else if strings.HasPrefix(mi.Name, "~") {
					cp.Append("body", fmt.Sprintf("func %s_%s%s(){}",
						ci.Name, "Dtor_"+mi.Name[1:], overloadSuffix))
				} else {
					cp.Append("body", fmt.Sprintf("func %s_%s%s(){}",
						ci.Name, mi.Name, overloadSuffix))
				}
			}

			argumento := methodso.GetIndex(j).Get("parameter")
			for k := 0; k < len(argumento.MustArray()); k++ {
				log.Println(mi.FullName, argumento.GetIndex(k))
			}
		}
	}
	log.Println("foundsymcnt:", foundsymcnt, "totalsymcnt:", totalsymcnt)
	log.Println("classcnt:", classcnt)
	ioutil.WriteFile("foundsigns.txt", []byte(strings.Join(foundsigns, "\n")), 0644)
	ioutil.WriteFile("tmp/gened.go", []byte(cp.ExportAll()), 0644)
}

var overloads = hashmap.New()

func rewriteOperatorMethodName(name string) string {
	replaces := []string{
		"&=", "_and_equal",
		"^=", "_jian_equal",
		"|=", "_or_equal",
		"+=", "_add_equal",
		"-=", "_minus_equal",
		"==", "_equal_equal",
		"!=", "_not_equal",
		"!", "_not", "=", "_equal",
		"<<", "_left_shift",
		">>", "_right_shift",
		"[]", "_get_index",
		"()", "_fncall",
		"->", "_pointer_selector",
		"<", "_less_than", ">", "_greater_than",
		"&", "_and", "^", "_jian", "|", "_or", "~", "_pozhehao",
		"/", "_div", "*", "_mul", "-", "_minus", "+", "_add",
		" ", "_"}
	valiname := name
	for i := 0; i < len(replaces)/2; i += 1 {
		valiname = strings.Replace(valiname, replaces[i*2], replaces[i*2+1], -1)
	}
	return valiname
}

func FilterMethod(mi *MethodInfo) bool {
	switch {
	case mi.Name == "setInterval" && strings.Contains(mi.Signature, "std::chrono::milliseconds"):
		return true
	}
	return false
}

func UnmarshalClass(clsjo *simplejson.Json) *ClassInfo {
	ci := &ClassInfo{}
	ci.Name = clsjo.Get("-name").MustString()
	ci.Access = clsjo.Get("-access").MustString()
	ci.Status = clsjo.Get("-status").MustString()
	ci.Module = clsjo.Get("-module").MustString()
	ci.Bases = strings.Split(clsjo.Get("-bases").MustString(), ",")
	return ci
}

type ClassInfo struct {
	Name   string
	Access string
	Status string
	Bases  []string
	Module string
}

func UnmarshalMethod(jso *simplejson.Json, ci *ClassInfo) *MethodInfo {
	mi := &MethodInfo{}
	mi.Name = jso.Get("-name").MustString()
	mi.FullName = jso.Get("-fullname").MustString()
	mi.Access = jso.Get("-access").MustString()
	mi.Status = jso.Get("-status").MustString()
	mi.Virtual = jso.Get("-virtual").MustString() == "true"
	mi.Type = jso.Get("-type").MustString()
	mi.Overload = jso.Get("-overload").MustString() == "true"
	mi.OverloadNumber = gopp.MustInt(jso.Get("-overload-number").MustString())
	mi.Signature = jso.Get("-signature").MustString()
	mi.Signature = strings.Replace(mi.Signature, "&amp;", "&", -1)
	mi.Signature = strings.Replace(mi.Signature, " )", ")", -1)

	if mi.Type == "" {
		mi.SymSignature = ci.Name + "::" + mi.Signature
	} else {
		mi.SymSignature = mi.Signature[len(mi.Type):]
		mi.SymSignature = ci.Name + "::" + strings.TrimSpace(mi.SymSignature)
	}
	mi.SymSignature = strings.Replace(mi.SymSignature, " )", ")", -1)

	// fix overload value error
	index := 0
	if index_, found := overloads.Get(mi.FullName); !found {
		overloads.Put(mi.FullName, int(0))
	} else {
		index = index_.(int) + 1
		overloads.Put(mi.FullName, index)
	}
	mi.Overload = true
	mi.OverloadNumber = index

	return mi
}

type MethodInfo struct {
	Name           string
	FullName       string
	Access         string
	Status         string
	Virtual        bool
	Const          bool
	Static         bool
	Overload       bool
	OverloadNumber int
	Delete         bool
	Default        bool
	Final          bool
	Override       bool
	Type           string
	Signature      string
	SymSignature   string
}
