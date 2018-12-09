package qtmeta

import (
	"fmt"
	"log"
	"os"
	"reflect"
	"runtime"
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

const qt_mark_tag = "qt"
const newfn_prefix = "NewForInherit"

/////
// Derive(this)
// pargs, args for base1, base2...
func Derive(obj interface{}, bargs ...[]interface{}) {
	// check exists
	objval := reflect.ValueOf(obj)
	objty := objval.Type().Elem()
	if mdcache.Has(objty.Name()) {
		return
	}

	// fill obj's inherit struct
	// parse obj's signal and fill function body
	// 所有的public方法都是slot
	_DeriveImpl(obj, bargs...)
}
func GetCthis(obj interface{}) unsafe.Pointer {
	objval := reflect.ValueOf(obj).Elem()
	retx := objval.MethodByName(qtrt.GetCthisName + "_").Call(nil)
	return retx[0].Interface().(unsafe.Pointer)
}

func Underive(obj interface{}) {
	mdcache.ClearSignal(obj) // decrease obj's refcnt
	mdcache.ClearSlot(obj)   // decrease obj's refcnt

	objval := reflect.ValueOf(obj)
	bclsfld := objval.Elem().Field(0).Interface()
	mdcache.RemoveC2go(GetCthis(bclsfld))

	objty := objval.Type().Elem()
	fldcnt := objty.NumField()
	for i := 0; i < fldcnt; i++ {
		fldo := objty.Field(i)
		switch fldo.Tag.Get("qt") {
		case "inherit":
			fldv := objval.Elem().Field(i)
			// fldv.Set(reflect.New(fldv.Type()).Elem())
			// 不能再此清0,finalizer的时候析构
			// 但是可以把其信号全部停掉吧
			cthis := GetCthis(fldv.Interface())
			qtrt.Assert(cthis != nil, "")

			var zeroptr unsafe.Pointer
			// QObject::disconnect(QObject const*, char const*, QObject const*, char const*)
			qtrt.InvokeQtFunc6("C_ZN7QObject10disconnectEPKS_PKcS1_S3_", qtrt.FFI_TYPE_POINTER, cthis, zeroptr, zeroptr, zeroptr)
		case "classinfo":
		case "property":
			// TODO clear property default value
		case "signal":
			fldv := objval.Elem().Field(i)
			fldv.Set(reflect.New(fldv.Type()).Elem())
		case "slot":
		default:
		}

		break //only process first one
	}
}

func _DeriveImpl(obj interface{}, bargs ...[]interface{}) bool {
	dctx := newDeriveContext(obj, bargs...)

	if !derive_precheck(dctx) {
		return false
	}

	derive_collect_metadata(dctx)
	derive_create_cobject(dctx)
	derive_fill_signals(dctx)
	derive_finalize(dctx)
	return true
}

type deriveContext struct {
	// in data
	obj    interface{}
	objval reflect.Value
	objty  reflect.Type
	bargs  [][]interface{}

	// out data
	mdo     *QtMetaData
	fldcnt  int
	cobjptr unsafe.Pointer
	cmdo    unsafe.Pointer
}

func newDeriveContext(obj interface{}, bargs ...[]interface{}) *deriveContext {
	this := &deriveContext{}
	this.obj = obj
	this.bargs = bargs
	this.objval = reflect.ValueOf(obj)
	this.objty = this.objval.Type().Elem()

	this.mdo = NewQtMetaData()
	this.mdo.SetClassName(this.objty.Name())
	this.fldcnt = this.objty.NumField()

	var dctx_dtored = false
	runtime.SetFinalizer(this, func(obj interface{}) { dctx_dtored = true })
	// we sure it will be GCed. But maybe need long to GCed
	// time.AfterFunc(5*time.Second, func() { qtrt.Assert(dctx_dtored, "dctx not success dtored") })
	return this
}

// check base class count
// check signal/slot param type supported
// check slot real exist.
func derive_precheck(dctx *deriveContext) bool {
	objty, fcnt := dctx.objty, dctx.fldcnt
	bclscnt := 0
	for i := 0; i < fcnt; i++ {
		fldo := objty.Field(i)
		switch fldo.Tag.Get("qt") {
		case "inherit":
			bclscnt += 1
		}
	}
	if bclscnt > 1 {
		log.Println("not support multiple base classes.", bclscnt)
		// return false
	}
	return true
}

func derive_collect_metadata(dctx *deriveContext) {
	obj, objval, objty, fcnt, mdo := dctx.obj, dctx.objval, dctx.objty, dctx.fldcnt, dctx.mdo

	// collect meta info
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
		case "slot":
			rettys, argtys := functyproto(fldo.Type)
			mdo.AddSlot(fldo.Name[1:], rettys, argtys)
			mtho := objval.MethodByName(fldo.Name[1:])
			if mtho.IsValid() {
				// TODO hold main object's reference, can not GCed
				slotfn := mtho.Interface()
				mdcache.AddSlot(obj, slotfn) // the only line cause obj not finalize
				_ = obj
				if !matchSignature(mtho.Type(), fldo.Type) {
					log.Panicln("coresponding slot func impl signature not match:", fldo.Type, mtho.Type())
				}
			} else {
				log.Panicln("can not found coresponding slot func impl:", fldo.Name[1:], objty)
			}
		}
	}

	mdo.FinalPass()
	// log.Println(mdo.Dump())
	// log.Println(mdo.MetaStrDat.Dump())
	mdcache.Add(objty.Name(), mdo)
}

func derive_create_cobject(dctx *deriveContext) {
	obj, bargs, objval, objty, fcnt, mdo :=
		dctx.obj, dctx.bargs, dctx.objval, dctx.objty, dctx.fldcnt, dctx.mdo

	// new foreign object
	var cobjptr unsafe.Pointer
	var cmdo unsafe.Pointer
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
			cobjx := retx.MethodByName(qtrt.GetCthisName).Call([]reflect.Value{})
			cobjptr = cobjx[0].Interface().(unsafe.Pointer)
			symname := fmt.Sprintf(qtrt.SetInitStObjName, fldo.Name)
			rv, _ := qtrt.InvokeQtFunc6(symname, qtrt.FFI_TYPE_POINTER, cobjptr,
				mdo.MetaStrDat.ToC(), mdo.ToC(), staticMetaCallFn(), metaCastFn(), metaCallFn())
			cmdo = unsafe.Pointer(uintptr(rv))
		}
		break //only process first one
	}
	dctx.cobjptr = cobjptr
	dctx.cmdo = cmdo
	mdcache.AddC2go(cobjptr, obj)
	_ = obj
	// TODO need think about destroyed() handler. see qdynslotobject.go:258 wtf
}

func derive_fill_signals(dctx *deriveContext) {
	obj, _, objval, objty, fcnt, _, cobjptr, cmdo :=
		dctx.obj, dctx.bargs, dctx.objval, dctx.objty, dctx.fldcnt, dctx.mdo, dctx.cobjptr, dctx.cmdo

	// collect meta info
	signal_index := 0
	for i := 0; i < fcnt; i++ {
		fldo := objty.Field(i)
		switch fldo.Tag.Get("qt") {
		case "signal":
			// TODO hold main object's reference, so obj can not GCed
			signalfn := reflect.MakeFunc(fldo.Type,
				mksignalfn(objty.String(), cobjptr, cmdo, signal_index, fldo.Name))
			objval.Elem().Field(i).Set(signalfn)
			signal_index += 1
			mdcache.AddSignal(obj, signalfn)
			_ = obj
		}
	}
}

func derive_finalize(dctx *deriveContext) {
	obj := dctx.obj
	runtime.SetFinalizer(obj, dtor_derived_obj)
}
func dtor_derived_obj(dobj interface{}) {
	objval := reflect.ValueOf(dobj)
	objty := objval.Type().Elem()
	bty := objty.Field(0).Type
	cthisx := objval.Elem().Field(0).MethodByName(qtrt.GetCthisName).Call(nil)
	cobjptr := cthisx[0].Interface().(unsafe.Pointer)
	clsname := mdcache.canName(bty.String())
	symname := fmt.Sprintf("C_ZN%d%sD2Ev", len(clsname), clsname)
	log.Printf("dtor~: %v, %s(%v)\n", objty, clsname, cobjptr)
	qtrt.InvokeQtFunc6(symname, qtrt.FFI_TYPE_POINTER, cobjptr)
}

// simple version
// argc == && argtype ==
func findMethod(clsty reflect.Type, args []interface{}) *reflect.Method {
	for i := 0; i < clsty.NumMethod(); i++ {
		mth := clsty.Method(i)
		// log.Println(mth.Name, mth.Func, mth.Func.Type())
		if strings.HasPrefix(mth.Name, newfn_prefix) {
			if nameResolve(mth.Func.Type(), args) {
				return &mth
			}
		}
	}
	// log.Printf("can not resolve method %s(%v) in %v\n", mthname, args,  clsty)
	return nil
}

func nameResolve(fnty reflect.Type, args []interface{}) bool {
	if fnty.NumIn()-1 != len(args) {
		// log.Println("argc not match", fnty.NumIn(), len(args))
		return false
	}
	for i := 0; i < len(args); i++ {
		argty := reflect.TypeOf(args[i])
		if !argty.AssignableTo(fnty.In(i + 1)) {
			// log.Printf("arg%d type not match: want %s, give %s\n", i, fnty.In(i+1), argty)
			return false
		}
	}
	return true
}

func matchSignature(fnty reflect.Type, yaty reflect.Type) bool {
	if fnty.NumIn() != yaty.NumIn() {
		return false
	}
	for i := 0; i < yaty.NumIn(); i++ {
		if !yaty.In(i).AssignableTo(fnty.In(i)) {
			return false
		}
	}
	return true
}

func mksignalfn(thisty string, cobjptr unsafe.Pointer, cmdo unsafe.Pointer, signal_index int, signal_name string) func(args []reflect.Value) (results []reflect.Value) {
	return func(args []reflect.Value) (results []reflect.Value) {
		var fullargs []uint64
		var argrefs []interface{}
		for idx, arg := range args {
			argty := arg.Type()
			switch argty.Kind() {
			case reflect.Ptr:
				if qtrt.IsQtclass(argty.String()) {
					cobjx := arg.MethodByName(qtrt.GetCthisName).Call(nil)
					fullargs = append(fullargs, uint64(uintptr(cobjx[0].UnsafeAddr())))
				} else {
					log.Panicln("not supported type:", argty)
				}
			default:
				argrefval := reflect.New(argty)
				argrefval.Elem().Set(arg)
				argrefs = append(argrefs, argrefval)
				log.Println(idx, arg, argty.Kind(), argrefval.Elem().UnsafeAddr(), cobjptr, cmdo, signal_index)
				fullargs = append(fullargs, uint64(uintptr(argrefval.Elem().UnsafeAddr())))
			}
		}

		var args_ unsafe.Pointer
		if len(fullargs) > 0 {
			args_ = unsafe.Pointer(&fullargs[0])
		}
		if cobjptr == nil || cmdo == nil {
			log.Panicln("wtf", cobjptr, cmdo, args)
		}
		qtrt.InvokeQtFunc6("C_ZN11QMetaObject8activateEP7QObjectPKS_iPPv", qtrt.FFI_TYPE_POINTER, cobjptr, cmdo, signal_index, args_)
		if debugQtMeta {
			log.Printf("emited signal: %s(%p)::%s:%d(%v|%v)\n", thisty, cobjptr, signal_name, signal_index, args, fullargs)
		}
		// void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(&_t1)) };
		// QMetaObject::activate(this, &staticMetaObject, 0, _a);
		return nil
	}
}

var debugQtMeta = false

func init() {
	debugQtMeta = os.Getenv("QTGO_DEBUG_QTMETA") == "1" ||
		strings.ToLower(os.Getenv("QTGO_DEBUG_QTMETA")) == "true"
}
func SetDebugQtMeta(on bool) { debugQtMeta = on }

/////
var mdcache = &_QtMetaDatas{
	mdos:      RwMap(),
	signalfns: RwMap(),
	slotfns:   RwMap(),
	c2gomap:   RwMap(),
}

// sync.Map有坑，即使delete也可能还是持有引用！！！// 换用这个了
type syncMap = Map

// Derive后，会在该类持有obj的3个引用，需要解引用
type _QtMetaDatas struct {
	mdos      syncMap // go struct name (class name) => *QtMetaData
	signalfns syncMap // go struct object (class object) = > []reflect.Value.Interface()
	slotfns   syncMap // go struct object (class object) = > []reflect.Value.Interface()
	c2gomap   syncMap // c object pointer => interface{}, go object pointer interface{}
}

func (this *_QtMetaDatas) Add(name string, mdo *QtMetaData) {
	name = this.canName(name)
	if !this.Has(name) {
		this.mdos.Store(name, mdo)
	} else {
		log.Println("already exists:", name)
	}
}
func (this *_QtMetaDatas) Has(name string) bool {
	name = this.canName(name)
	_, ok := this.mdos.Load(name)
	return ok
}
func (this *_QtMetaDatas) Get(name string) *QtMetaData {
	name = this.canName(name)
	mdox, ok := this.mdos.Load(name)
	if ok {
		return mdox.(*QtMetaData)
	}
	return nil
}
func (this *_QtMetaDatas) GetByObj(objx interface{}) *QtMetaData {
	objval := reflect.ValueOf(objx)
	objty := objval.Type().Elem()
	tycanname := mdcache.canName(objty.Name())
	tymeta := mdcache.Get(tycanname)
	return tymeta
}
func (this *_QtMetaDatas) canName(name string) string {
	if strings.Index(name, ".") >= 0 {
		segs := strings.Split(name, ".")
		return segs[len(segs)-1]
	}
	return name
}
func (this *_QtMetaDatas) AddSignal(obj interface{}, fnval interface{}) {
	var fns []interface{}
	fnsx, ok := this.signalfns.Load(obj)
	if ok {
		fns = fnsx.([]interface{})
	}
	fns = append(fns, fnval)
	this.signalfns.Store(obj, fns)
}
func (this *_QtMetaDatas) GetSignal(obj interface{}, idx int) interface{} {
	fnsx, ok := this.signalfns.Load(obj)
	if ok {
		fns := fnsx.([]interface{})
		if idx >= 0 && idx < len(fns) {
			return fns[idx]
		}
	}
	return nil
}
func (this *_QtMetaDatas) ClearSignal(obj interface{}) {
	fnsx, ok := this.signalfns.Load(obj)
	if ok {
		fns := fnsx.([]interface{})
		for i := 0; i < len(fns); i++ {
			fns[i] = nil
		}
		this.signalfns.Store(obj, fns)
		this.signalfns.Delete(obj)
	}
}
func (this *_QtMetaDatas) AddSlot(obj interface{}, fnval interface{}) {
	var fns []interface{}
	fnsx, ok := this.slotfns.Load(obj)
	if ok {
		fns = fnsx.([]interface{})
	}
	fns = append(fns, fnval)
	this.slotfns.Store(obj, fns)
}
func (this *_QtMetaDatas) GetSlot(obj interface{}, idx int) interface{} {
	fnsx, ok := this.slotfns.Load(obj)
	if ok {
		fns := fnsx.([]interface{})
		if idx >= 0 && idx < len(fns) {
			return fns[idx]
		}
	}
	return nil
}
func (this *_QtMetaDatas) ClearSlot(obj interface{}) {
	fnsx, ok := this.slotfns.Load(obj)
	if ok {
		fns := fnsx.([]interface{})
		for i := 0; i < len(fns); i++ {
			fns[i] = nil
		}
		this.slotfns.Store(obj, fns)
		this.slotfns.Delete(obj)
	}
}
func (this *_QtMetaDatas) AddC2go(cobj unsafe.Pointer, goobj interface{}) {
	this.c2gomap.Store(cobj, goobj)
}
func (this *_QtMetaDatas) HasC2go(cobj unsafe.Pointer) bool {
	_, ok := this.c2gomap.Load(cobj)
	return ok
}
func (this *_QtMetaDatas) RemoveC2go(cobj unsafe.Pointer) {
	this.c2gomap.Delete(cobj)
}
func (this *_QtMetaDatas) GetC2go(cobj unsafe.Pointer) (interface{}, bool) {
	return this.c2gomap.Load(cobj)
}
func countSyncMap(m syncMap) int { return m.Len() }

///
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
