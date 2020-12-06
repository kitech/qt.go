package qtrel

import (
	"go/scanner"
	"go/token"
	"log"
	"reflect"
	"strings"
	"unsafe"

	"github.com/ianlancetaylor/demangle"
	"github.com/kitech/qt.go/qtrt"
)

const ulpfx = "_ZN"
const culpfx = "C_ZN"

/*
const QStringNew0 = "_ZN7QStringC2Ev"
//const QStringNew1 = "_ZN7QStringC2EPKc"
const QStringLength0 = "_ZNK7QString6lengthEv"
*/

// no this version, for ctor and static member and c function
func CC0(mgname string, args ...interface{}) ccret {
	return ccimpl(nil, mgname, args...)
}

// only return pointer
func CC0p(mgname string, args ...interface{}) CObject {
	rv := CC0(mgname, args...)
	return CObject(rv.Ptr())
}

func ccimpl(cthis unsafe.Pointer, mgname string, args ...interface{}) ccret {
	mgname = ulpfx + mgname
	cmo := cppfilt(mgname)
	argv := convgoargs2c(cmo, args)
	if cthis != nil {
		argv = append([]interface{}{cthis}, argv...)
		_ = argv
	}

	rv, err := qtrt.InvokeQtFunc6(mgname, qtrt.FFI_TYPE_UINT64, argv...)
	// log.Printf("rv=%+v err=%v %s\n", rv, err, mgname)
	qtrt.ErrPrint(err)
	return newccret(rv, &mgname, cmo)
}

func (this CObject) CC0(mgname string, args ...interface{}) ccret {
	return ccimpl(this.Ptr(), mgname, args...)
}

func (this CObject) CC0p(mgname string, args ...interface{}) CObject {
	rv := ccimpl(this.Ptr(), mgname, args...)
	return rv.Cobj()
}

func Inherit(cbobj qtrt.CObjectITF, mgname string, fn interface{}) {
	// cthis unsafe.Pointer
	mgname = ulpfx + mgname
	cmo := cppfilt(mgname)
	qtrt.SetAllInheritCallback(cbobj, cmo.mth, fn)
}
func (this CObject) Inherit(mgname string, fn interface{}) {
	// qtrt.SetAllInheritCallback(this, "", fn)
	Inherit(this, mgname, fn)
}

func (this CObject) Connect(signal string, slotfn interface{}) {
	qtrt.Connect(this, signal, slotfn)
}
func (this CObject) Disconnect(signal string) {
	qtrt.Disconnect(this, signal)
}

type CObject uintptr

func (this CObject) GetCthis() unsafe.Pointer      { return unsafe.Pointer(this) }
func (this CObject) SetCthis(cthis unsafe.Pointer) {}
func (CObject) FromCthis(cthis unsafe.Pointer) CObject {
	return CObject(uintptr(cthis))
}
func (this CObject) Ptr() unsafe.Pointer { return unsafe.Pointer(this) }

// type QObject uintptr
// type QWidget uintptr

var ccrety = reflect.TypeOf(ccret(0))

// assert sizeof(ccret) == sizeof(uintptr)
// TODO how about x86 pointer
type ccret uint64

/*
type ccret struct {
	rv uint64
	//typ int
}
*/

func newccret(rv uint64, fnsym *string, cmo *cppmethod) ccret {
	return ccret(rv)
	// return ccret{rv}
	//return ccret{rv, 0}
}

func (this ccret) Int() int            { return int(this) }
func (this ccret) Uint() uint          { return uint(this) }
func (this ccret) I32() int32          { return int32(this) }
func (this ccret) I64() int64          { return int64(this) }
func (this ccret) U32() uint32         { return uint32(this) }
func (this ccret) U64() uint64         { return uint64(this) }
func (this ccret) F32() float32        { return float32(this) }
func (this ccret) F64() float64        { return float64(this) }
func (this ccret) Ptr() unsafe.Pointer { return unsafe.Pointer(uintptr(this)) }
func (this ccret) Uptr() uintptr       { return uintptr(this) }
func (this ccret) Bool() bool          { return this != 0 }
func (this ccret) Strc() string {
	// char* => string
	return qtrt.GoString(this.Ptr())
}
func (this ccret) Strq() string {
	// QString => string
	ba := this.Cobj().CC0(QStringToUtf8)
	rv := ba.Cobj().CC0(QByteArrayData)
	return rv.Strc()
}
func (this ccret) Cobj() CObject { return CObject(this) }

// for Qt multiple inherit
func (this ccret) Cobj2() CObject {
	return CObject(this.Uptr() + unsafe.Sizeof(uintptr(0)))
}

func convgoargs2c(cmo *cppmethod, args []interface{}) []interface{} {
	var argv []interface{}
	if len(args) != cmo.argc() {
		log.Printf("WARN argc not match want %d have %d when call %s.%s\n",
			cmo.argc(), len(args), cmo.cls, cmo.mth)
	}
	for idx, arg := range args {
		val := convgoarg2c(cmo, idx, arg)
		argv = append(argv, val)
	}
	// log.Println("argc/want", len(argv), cmo.argc(), ":", argv)
	return argv
}

func convgoarg2c(cmo *cppmethod, idx int, argx interface{}) interface{} {
	var rv interface{} = argx

	argty := reflect.TypeOf(argx)
	argval := reflect.ValueOf(argx)
	if false {
		log.Println(argty, argval)
	}

	cppty := cmo.types[idx]
	if argty == nil {
		if argx == nil {
			rv = unsafe.Pointer(nil)
			return rv
		}
	}
	switch argty.Kind() {
	case reflect.Invalid:
		log.Println("TODO", cmo.cls, cmo.mth)
	case reflect.String:
		// which one? // => QString // => charptr
		tval := argx.(string)
		if cppty == "char *" {
			rv = unsafe.Pointer(nil)
			// val1 := C.CString(arg.(string))
			// val2 := unsafe.Pointer(val1)
			val1 := []byte(tval)
			if len(val1) > 0 {
				val2 := unsafe.Pointer(&val1[0]) // works
				rv = val2
			}
		} else if strings.HasPrefix(cppty, "QString") {
			val1 := CC0(QStringNew5, tval)
			rv = val1.Ptr()
		}
	case reflect.Slice:
		if argty.Elem().Kind() == reflect.String {
			rv = qtrt.StringSliceToCCharPP(argx.([]string))
			// log.Println("converted", argty, arg)
		}
	case reflect.Int:
		if cmo.argisref(idx) { // fix int&
			rvp := reflect.New(argty)
			rvp.Elem().Set(argval)
			rv = unsafe.Pointer(rvp.Pointer())
		}
	case reflect.Bool: // just ok
	case reflect.Uintptr:
		log.Println("wtt", idx, argx)
	case reflect.Struct:
		if argty == ccrety {
			arg := argx.(ccret)
			rv = arg.Ptr()
		}
	default:
		log.Println("wtt", idx, argty, argty.Kind(), argval)
	}

	return rv
}

type cppmethod struct {
	cls     string
	mth     string
	types   []string
	names   []string
	isconst bool
}

func (this *cppmethod) isctor() bool {
	return this.cls != "" && (this.cls == this.mth)
}

func (this *cppmethod) isdtor() bool {
	return strings.HasPrefix(this.mth, "~")
}

func (this *cppmethod) ismth() bool {
	return this.cls != ""
}

func (this *cppmethod) argc() int { return len(this.types) }
func (this *cppmethod) argisref(idx int) bool {
	return strings.Count(this.types[idx], "&") == 1
}
func (this *cppmethod) argisptr(idx int) bool {
	return strings.Count(this.types[idx], "*") == 1
}

func cppfilt(mgname string) *cppmethod {
	cmo := &cppmethod{}
	line, err := demangle.ToString(mgname)
	qtrt.ErrPrint(err)
	// log.Println(mgname, line)

	{
		// src := []byte("cos(x) + 1i*sin(x) // Euler")
		src := []byte(line)

		// Initialize the scanner.
		var s scanner.Scanner
		fset := token.NewFileSet()                      // positions are relative to fset
		file := fset.AddFile("", fset.Base(), len(src)) // register input "file"
		s.Init(file, src, nil /* no error handler */, scanner.ScanComments)

		var idts []string
		var colon = 0
		// Repeated calls to Scan yield the token sequence found in the input.
		for {
			pos, tok, lit := s.Scan()
			if tok == token.EOF {
				break
			}
			// fmt.Printf("%s\t%s\t%q\n", fset.Position(pos), tok, lit)
			switch tok {
			case token.CONST:
			case token.IDENT:
				idts = append(idts, lit)
			case token.COLON:
				colon += 1
				if colon == 2 {
					cmo.cls = idts[len(idts)-1]
				}
			case token.LPAREN:
				cmo.mth = idts[len(idts)-1]
				idts = []string{}
			case token.RPAREN:
				if len(idts) == 0 {
					break
				}
				tyfld := strings.Join(idts, " ")
				cmo.types = append(cmo.types, tyfld)
				idts = []string{}
			case token.MUL:
				idts = append(idts, tok.String())
			case token.AND:
				idts = append(idts, tok.String())
			case token.COMMA:
				if len(idts) == 0 {
					break
				}
				tyfld := strings.Join(idts, " ")
				cmo.types = append(cmo.types, tyfld)
				idts = []string{}
			case token.SEMICOLON:
			case token.LSS: // TODO template<>
				idts = append(idts, tok.String())
			case token.GTR:
				idts = append(idts, tok.String())
			default:
				log.Println("wtt", fset.Position(pos), tok, lit, mgname)
			}
		}
		// log.Printf("%#v\n", cmo)
	}

	return cmo
}
