package qtqml

// /usr/include/qt/QtQml/qqmlfileselector.h
// #include <qqmlfileselector.h>
// #include <QtQml>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 28
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "mkuse/cffiqt"
import "gopp"
import "qt.go/qtrt"
import "qt.go/qtcore"
import "qt.go/qtnetwork"

func init() {
	if false {
		reflect.TypeOf(123)
	}
	if false {
		reflect.TypeOf(unsafe.Sizeof(0))
	}
	if false {
		fmt.Println(123)
	}
	if false {
		qtrt.KeepMe()
	}
	if false {
		ffiqt.KeepMe()
	}
	if false {
		gopp.KeepMe()
	}
	if false {
		qtcore.KeepMe()
	}
	if false {
		qtnetwork.KeepMe()
	}
}

//  ext block end

//  body block begin
type QQmlFileSelector struct {
	*qtcore.QObject
}

func (this *QQmlFileSelector) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QObject.GetCthis()
	}
}
func (this *QQmlFileSelector) SetCthis(cthis unsafe.Pointer) {
	this.QObject = qtcore.NewQObjectFromPointer(cthis)
}
func NewQQmlFileSelectorFromPointer(cthis unsafe.Pointer) *QQmlFileSelector {
	bcthis0 := qtcore.NewQObjectFromPointer(cthis)
	return &QQmlFileSelector{bcthis0}
}
func (*QQmlFileSelector) NewFromPointer(cthis unsafe.Pointer) *QQmlFileSelector {
	return NewQQmlFileSelectorFromPointer(cthis)
}

// /usr/include/qt/QtQml/qqmlfileselector.h:54
// index:0
// Public virtual
// const QMetaObject * metaObject()
func (this *QQmlFileSelector) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QQmlFileSelector10metaObjectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtQml/qqmlfileselector.h:57
// index:0
// Public
// void QQmlFileSelector(class QQmlEngine *, class QObject *)
func NewQQmlFileSelector(engine *QQmlEngine /*777 QQmlEngine **/, parent *qtcore.QObject /*777 QObject **/) *QQmlFileSelector {
	cthis := qtrt.Calloc(1, 256) // 16
	var convArg0 = engine.GetCthis()
	var convArg1 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QQmlFileSelectorC2EP10QQmlEngineP7QObject", ffiqt.FFI_TYPE_VOID, cthis, convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	gothis := NewQQmlFileSelectorFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtQml/qqmlfileselector.h:58
// index:0
// Public virtual
// void ~QQmlFileSelector()
func DeleteQQmlFileSelector(*QQmlFileSelector) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QQmlFileSelectorD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQml/qqmlfileselector.h:59
// index:0
// Public
// QFileSelector * selector()
func (this *QQmlFileSelector) Selector() *qtcore.QFileSelector /*777 QFileSelector **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QQmlFileSelector8selectorEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQFileSelectorFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtQml/qqmlfileselector.h:60
// index:0
// Public
// void setSelector(class QFileSelector *)
func (this *QQmlFileSelector) SetSelector(selector *qtcore.QFileSelector /*777 QFileSelector **/) {
	var convArg0 = selector.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QQmlFileSelector11setSelectorEP13QFileSelector", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQml/qqmlfileselector.h:61
// index:0
// Public
// void setExtraSelectors(class QStringList &)
func (this *QQmlFileSelector) SetExtraSelectors(strings *qtcore.QStringList) {
	var convArg0 = strings.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QQmlFileSelector17setExtraSelectorsER11QStringList", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQml/qqmlfileselector.h:62
// index:1
// Public
// void setExtraSelectors(const class QStringList &)
func (this *QQmlFileSelector) SetExtraSelectors_1(strings *qtcore.QStringList) {
	var convArg0 = strings.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QQmlFileSelector17setExtraSelectorsERK11QStringList", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQml/qqmlfileselector.h:63
// index:0
// Public static
// QQmlFileSelector * get(class QQmlEngine *)
func (this *QQmlFileSelector) Get(arg0 *QQmlEngine /*777 QQmlEngine **/) *QQmlFileSelector /*777 QQmlFileSelector **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QQmlFileSelector3getEP10QQmlEngine", ffiqt.FFI_TYPE_POINTER, arg0)
	gopp.ErrPrint(err, rv)
	// return rv
	rv2 := /*==*/ NewQQmlFileSelectorFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}
func QQmlFileSelector_Get(arg0 *QQmlEngine /*777 QQmlEngine **/) *QQmlFileSelector /*777 QQmlFileSelector **/ {
	var nilthis *QQmlFileSelector
	rv := nilthis.Get(arg0)
	return rv
}

//  body block end
