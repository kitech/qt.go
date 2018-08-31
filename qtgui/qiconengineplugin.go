package qtgui

// /usr/include/qt/QtGui/qiconengineplugin.h
// #include <qiconengineplugin.h>
// #include <QtGui>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 16
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "log"
import "github.com/kitech/qt.go/qtrt"
import "github.com/kitech/qt.go/qtcore"

//  ext block end

//  body block begin

/*

 */
type QIconEnginePlugin struct {
	*qtcore.QObject
}
type QIconEnginePlugin_ITF interface {
	qtcore.QObject_ITF
	QIconEnginePlugin_PTR() *QIconEnginePlugin
}

func (ptr *QIconEnginePlugin) QIconEnginePlugin_PTR() *QIconEnginePlugin { return ptr }

func (this *QIconEnginePlugin) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QObject.GetCthis()
	}
}
func (this *QIconEnginePlugin) SetCthis(cthis unsafe.Pointer) {
	this.QObject = qtcore.NewQObjectFromPointer(cthis)
}
func NewQIconEnginePluginFromPointer(cthis unsafe.Pointer) *QIconEnginePlugin {
	bcthis0 := qtcore.NewQObjectFromPointer(cthis)
	return &QIconEnginePlugin{bcthis0}
}
func (*QIconEnginePlugin) NewFromPointer(cthis unsafe.Pointer) *QIconEnginePlugin {
	return NewQIconEnginePluginFromPointer(cthis)
}

// /usr/include/qt/QtGui/qiconengineplugin.h:56
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject() const

/*

 */
func (this *QIconEnginePlugin) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QIconEnginePlugin10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtGui/qiconengineplugin.h:58
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QIconEnginePlugin(QObject *)

/*
Constructs a icon engine plugin with the given parent. This is invoked automatically by the plugin loader.
*/
func (*QIconEnginePlugin) NewForInherit(parent qtcore.QObject_ITF /*777 QObject **/) *QIconEnginePlugin {
	return NewQIconEnginePlugin(parent)
}
func NewQIconEnginePlugin(parent qtcore.QObject_ITF /*777 QObject **/) *QIconEnginePlugin {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QObject_PTR() != nil {
		convArg0 = parent.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN17QIconEnginePluginC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQIconEnginePluginFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QIconEnginePlugin")
	return gothis
}

// /usr/include/qt/QtGui/qiconengineplugin.h:58
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QIconEnginePlugin(QObject *)

/*
Constructs a icon engine plugin with the given parent. This is invoked automatically by the plugin loader.
*/
func (*QIconEnginePlugin) NewForInherit__() *QIconEnginePlugin {
	return NewQIconEnginePlugin__()
}
func NewQIconEnginePlugin__() *QIconEnginePlugin {
	// arg: 0, QObject *=Pointer, QObject=Record, , Invalid
	var convArg0 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN17QIconEnginePluginC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQIconEnginePluginFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QIconEnginePlugin")
	return gothis
}

// /usr/include/qt/QtGui/qiconengineplugin.h:59
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QIconEnginePlugin()

/*

 */
func DeleteQIconEnginePlugin(this *QIconEnginePlugin) {
	rv, err := qtrt.InvokeQtFunc6("_ZN17QIconEnginePluginD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 16)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtGui/qiconengineplugin.h:61
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [8] QIconEngine * create(const QString &)

/*
Creates and returns a QIconEngine object for the icon with the given filename.
*/
func (this *QIconEnginePlugin) Create(filename string) *QIconEngine /*777 QIconEngine **/ {
	var tmpArg0 = qtcore.NewQString_5(filename)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN17QIconEnginePlugin6createERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQIconEngineFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtGui/qiconengineplugin.h:61
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [8] QIconEngine * create(const QString &)

/*
Creates and returns a QIconEngine object for the icon with the given filename.
*/
func (this *QIconEnginePlugin) Create__() *QIconEngine /*777 QIconEngine **/ {
	// arg: 0, const QString &=LValueReference, QString=Record, , Invalid
	var convArg0 = qtcore.NewQString()
	rv, err := qtrt.InvokeQtFunc6("_ZN17QIconEnginePlugin6createERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQIconEngineFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

//  body block end

//  keep block begin

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
		log.Println(123)
	}
	if false {
		qtrt.KeepMe()
	}
	if false {
		qtcore.KeepMe()
	}
}

//  keep block end
