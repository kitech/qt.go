package qtqml

// /usr/include/qt/QtQml/qqmlapplicationengine.h
// #include <qqmlapplicationengine.h>
// #include <QtQml>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 36
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
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

type QQmlApplicationEngine struct {
	*QQmlEngine
}

func (this *QQmlApplicationEngine) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QQmlEngine.GetCthis()
	}
}
func (this *QQmlApplicationEngine) SetCthis(cthis unsafe.Pointer) {
	this.QQmlEngine = NewQQmlEngineFromPointer(cthis)
}
func NewQQmlApplicationEngineFromPointer(cthis unsafe.Pointer) *QQmlApplicationEngine {
	bcthis0 := NewQQmlEngineFromPointer(cthis)
	return &QQmlApplicationEngine{bcthis0}
}
func (*QQmlApplicationEngine) NewFromPointer(cthis unsafe.Pointer) *QQmlApplicationEngine {
	return NewQQmlApplicationEngineFromPointer(cthis)
}

// /usr/include/qt/QtQml/qqmlapplicationengine.h:54
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject()
func (this *QQmlApplicationEngine) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK21QQmlApplicationEngine10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtQml/qqmlapplicationengine.h:56
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QQmlApplicationEngine(QObject *)
func NewQQmlApplicationEngine(parent *qtcore.QObject /*777 QObject **/) *QQmlApplicationEngine {
	var convArg0 = parent.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN21QQmlApplicationEngineC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQQmlApplicationEngineFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtQml/qqmlapplicationengine.h:57
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QQmlApplicationEngine(const QUrl &, QObject *)
func NewQQmlApplicationEngine_1(url *qtcore.QUrl, parent *qtcore.QObject /*777 QObject **/) *QQmlApplicationEngine {
	var convArg0 = url.GetCthis()
	var convArg1 = parent.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN21QQmlApplicationEngineC2ERK4QUrlP7QObject", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	gothis := NewQQmlApplicationEngineFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtQml/qqmlapplicationengine.h:58
// index:2
// Public Visibility=Default Availability=Available
// [-2] void QQmlApplicationEngine(const QString &, QObject *)
func NewQQmlApplicationEngine_2(filePath *qtcore.QString, parent *qtcore.QObject /*777 QObject **/) *QQmlApplicationEngine {
	var convArg0 = filePath.GetCthis()
	var convArg1 = parent.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN21QQmlApplicationEngineC2ERK7QStringP7QObject", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	gothis := NewQQmlApplicationEngineFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtQml/qqmlapplicationengine.h:59
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QQmlApplicationEngine()
func DeleteQQmlApplicationEngine(this *QQmlApplicationEngine) {
	rv, err := qtrt.InvokeQtFunc6("_ZN21QQmlApplicationEngineD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 24)
	gopp.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtQml/qqmlapplicationengine.h:67
// index:0
// Public Visibility=Default Availability=Available
// [-2] void load(const QUrl &)
func (this *QQmlApplicationEngine) Load(url *qtcore.QUrl) {
	var convArg0 = url.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN21QQmlApplicationEngine4loadERK4QUrl", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQml/qqmlapplicationengine.h:68
// index:1
// Public Visibility=Default Availability=Available
// [-2] void load(const QString &)
func (this *QQmlApplicationEngine) Load_1(filePath *qtcore.QString) {
	var convArg0 = filePath.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN21QQmlApplicationEngine4loadERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQml/qqmlapplicationengine.h:69
// index:0
// Public Visibility=Default Availability=Available
// [-2] void loadData(const QByteArray &, const QUrl &)
func (this *QQmlApplicationEngine) LoadData(data *qtcore.QByteArray, url *qtcore.QUrl) {
	var convArg0 = data.GetCthis()
	var convArg1 = url.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN21QQmlApplicationEngine8loadDataERK10QByteArrayRK4QUrl", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQml/qqmlapplicationengine.h:72
// index:0
// Public Visibility=Default Availability=Available
// [-2] void objectCreated(QObject *, const QUrl &)
func (this *QQmlApplicationEngine) ObjectCreated(object *qtcore.QObject /*777 QObject **/, url *qtcore.QUrl) {
	var convArg0 = object.GetCthis()
	var convArg1 = url.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN21QQmlApplicationEngine13objectCreatedEP7QObjectRK4QUrl", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
}

//  body block end
