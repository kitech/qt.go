package qtwidgets

// /usr/include/qt/QtWidgets/qstyleplugin.h
// #include <qstyleplugin.h>
// #include <QtWidgets>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 11
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "log"
import "github.com/kitech/qt.go/qtrt"
import "github.com/kitech/qt.go/qtcore"
import "github.com/kitech/qt.go/qtgui"

//  ext block end

//  body block begin

/*

 */
type QStylePlugin struct {
	*qtcore.QObject
}
type QStylePlugin_ITF interface {
	qtcore.QObject_ITF
	QStylePlugin_PTR() *QStylePlugin
}

func (ptr *QStylePlugin) QStylePlugin_PTR() *QStylePlugin { return ptr }

func (this *QStylePlugin) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QObject.GetCthis()
	}
}
func (this *QStylePlugin) SetCthis(cthis unsafe.Pointer) {
	this.QObject = qtcore.NewQObjectFromPointer(cthis)
}
func NewQStylePluginFromPointer(cthis unsafe.Pointer) *QStylePlugin {
	bcthis0 := qtcore.NewQObjectFromPointer(cthis)
	return &QStylePlugin{bcthis0}
}
func (*QStylePlugin) NewFromPointer(cthis unsafe.Pointer) *QStylePlugin {
	return NewQStylePluginFromPointer(cthis)
}

// /usr/include/qt/QtWidgets/qstyleplugin.h:56
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject() const

/*

 */
func (this *QStylePlugin) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QStylePlugin10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qstyleplugin.h:58
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QStylePlugin(QObject *)

/*
Constructs a style plugin with the given parent.

Note that this constructor is invoked automatically by the moc generated code that exports the plugin, so there is no need for calling it explicitly.
*/
func NewQStylePlugin(parent qtcore.QObject_ITF /*777 QObject **/) *QStylePlugin {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QObject_PTR() != nil {
		convArg0 = parent.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN12QStylePluginC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQStylePluginFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QStylePlugin")
	return gothis
}

// /usr/include/qt/QtWidgets/qstyleplugin.h:58
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QStylePlugin(QObject *)

/*
Constructs a style plugin with the given parent.

Note that this constructor is invoked automatically by the moc generated code that exports the plugin, so there is no need for calling it explicitly.
*/
func NewQStylePlugin__() *QStylePlugin {
	// arg: 0, QObject *=Pointer, QObject=Record, , Invalid
	var convArg0 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN12QStylePluginC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQStylePluginFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QStylePlugin")
	return gothis
}

// /usr/include/qt/QtWidgets/qstyleplugin.h:59
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QStylePlugin()

/*

 */
func DeleteQStylePlugin(this *QStylePlugin) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QStylePluginD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 16)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtWidgets/qstyleplugin.h:61
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [8] QStyle * create(const QString &)

/*
Creates and returns a QStyle object for the given style key. If a plugin cannot create a style, it should return 0 instead.

The style key is usually the class name of the required style. Note that the keys are case insensitive. For example:


  QStyle *MyStylePlugin::create(const QString &key)
  {
      QString lcKey = key.toLower();
      if (lcKey == "rocket") {
          return new RocketStyle;
      } else if (lcKey == "starbuster") {
          return new StarBusterStyle;
      }
      return 0;
  }
*/
func (this *QStylePlugin) Create(key string) *QStyle /*777 QStyle **/ {
	var tmpArg0 = qtcore.NewQString_5(key)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN12QStylePlugin6createERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQStyleFromPointer(unsafe.Pointer(uintptr(rv))) // 444
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
	if false {
		qtgui.KeepMe()
	}
}

//  keep block end
