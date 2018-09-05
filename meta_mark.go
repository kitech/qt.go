package qtmeta

import (
	"fmt"
	"log"
	"reflect"
	"strings"
	"unsafe"

	"github.com/kitech/qt.go/qtrt"
)

// syntax parser and mark relate types
type Q_SIGNALS_BEGIN struct{}
type Q_SIGNALS_END struct{}
type Q_SLOTS_BEGIN struct{}
type Q_SLOTS_END struct{}
type Q_META_SECTION_END struct{}

const qt_tag = "qt"
const newfn_name = "NewForInherit"

/////
// Derive(this)
// pargs, args for base1, base2...
func Derive(obj interface{}, bargs ...[]interface{}) {
	// fill obj's inherit struct
	// parse obj's signal and fill function body
	// 所有的public方法都是slot
	_DeriveImpl(obj, bargs...)
}
func Underive(obj interface{}) {}

func _DeriveImpl(obj interface{}, bargs ...[]interface{}) {
	objval := reflect.ValueOf(obj)
	objty := objval.Type().Elem()

	mdo := NewQtMetaData()
	mdo.SetClassName(objty.Name()) // TODO need regular name

	fcnt := objty.NumField()
	for i := 0; i < fcnt; i++ {
		fldo := objty.Field(i)
		tago := fldo.Tag
		switch fldo.Tag.Get("qt") {
		case "inherit":
		case "classinfo":
			key, value := tago.Get("key"), tago.Get("value")
			mdo.AddClassInfo(key, value)
		case "property":
			// TODO set property default value
		case "signal":
			rettys, argtys := functyproto(fldo.Type)
			mdo.AddSignal(fldo.Name, rettys, argtys)
			// TODO gen func value for this signal
		case "slot":
			rettys, argtys := functyproto(fldo.Type)
			mdo.AddSlot(fldo.Name[1:], rettys, argtys)
			// TODO check real slot func exists
		}
	}

	mdo.FinalPass()
	log.Println(mdo.Dump())
	log.Println(mdo.MetaStrDat.Dump())

	for i := 0; i < fcnt; i++ {
		fldo := objty.Field(i)
		switch fldo.Tag.Get("qt") {
		case "inherit":
			nilthiso := reflect.NewAt(fldo.Type.Elem(), nil)
			var bargs2 = []interface{}{}
			if len(bargs) > 0 {
				bargs2 = bargs[0]
			}
			newmth := findMethod(nilthiso.Type(), bargs2)
			if newmth == nil {
				log.Println("Base class have no `NewForInherit` method:", fldo.Type)
				break
			}

			rets := newmth.Func.Call([]reflect.Value{nilthiso})
			var retx reflect.Value = rets[0]
			objval.Elem().Field(i).Set(retx)
			cobjx := retx.MethodByName("GetCthis").Call([]reflect.Value{})
			symname := fmt.Sprintf("C_%s_init_staticMetaObject", fldo.Name)
			qtrt.InvokeQtFunc6(symname, qtrt.FFI_TYPE_POINTER,
				cobjx[0].Interface().(unsafe.Pointer),
				mdo.MetaStrDat.ToC(), mdo.ToC(), staticMetaCallFn(), metaCastFn(), metaCallFn())
		}
	}
}

// simple version
// argc == && argtype ==
func findMethod(clsty reflect.Type, args []interface{}) *reflect.Method {
	for i := 0; i < clsty.NumMethod(); i++ {
		mth := clsty.Method(i)
		// log.Println(mth.Name, mth.Func, mth.Func.Type())
		if strings.HasPrefix(mth.Name, newfn_name) {
			if nameResolve(mth.Func.Type(), args) {
				return &mth
			}
		}
	}

	return nil
}

func nameResolve(fnty reflect.Type, args []interface{}) bool {
	if fnty.NumIn()-1 != len(args) {
		log.Println("argc not match", fnty.NumIn(), len(args))
		return false
	}
	for i := 0; i < len(args); i++ {
		argty := reflect.TypeOf(args[i])
		if !argty.AssignableTo(fnty.In(i + 1)) {
			log.Printf("arg%d type not match: want %s, give %s\n", i, fnty.In(i+1), argty)
			return false
		}
	}
	return true
}

type QMetaObjectData struct {
	// per class members
	Superdata          unsafe.Pointer
	Stringdata         unsafe.Pointer
	Data               unsafe.Pointer
	Static_metacall    unsafe.Pointer
	relatedMetaObjects unsafe.Pointer
	extradata          unsafe.Pointer

	// per instance members
	// Metacall unsafe.Pointer
	// Metacast unsafe.Pointer
}

func NewQMetaObjectData() {

}

func (this *QMetaObjectData) GetCthis() unsafe.Pointer {
	return unsafe.Pointer(this)
}

func (this *QMetaObjectData) SetCthis(unsafe.Pointer) {}
