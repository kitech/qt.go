package qtgui

// /usr/include/qt/QtGui/qpictureformatplugin.h
// #include <qpictureformatplugin.h>
// #include <QtGui>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 26
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
type QPictureFormatPlugin struct {
	*qtcore.QObject
}
type QPictureFormatPlugin_ITF interface {
	qtcore.QObject_ITF
	QPictureFormatPlugin_PTR() *QPictureFormatPlugin
}

func (ptr *QPictureFormatPlugin) QPictureFormatPlugin_PTR() *QPictureFormatPlugin { return ptr }

func (this *QPictureFormatPlugin) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QObject.GetCthis()
	}
}
func (this *QPictureFormatPlugin) SetCthis(cthis unsafe.Pointer) {
	this.QObject = qtcore.NewQObjectFromPointer(cthis)
}
func NewQPictureFormatPluginFromPointer(cthis unsafe.Pointer) *QPictureFormatPlugin {
	bcthis0 := qtcore.NewQObjectFromPointer(cthis)
	return &QPictureFormatPlugin{bcthis0}
}
func (*QPictureFormatPlugin) NewFromPointer(cthis unsafe.Pointer) *QPictureFormatPlugin {
	return NewQPictureFormatPluginFromPointer(cthis)
}

// /usr/include/qt/QtGui/qpictureformatplugin.h:61
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject() const

/*

 */
func (this *QPictureFormatPlugin) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK20QPictureFormatPlugin10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtGui/qpictureformatplugin.h:63
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QPictureFormatPlugin(QObject *)

/*
Constructs an picture format plugin with the given parent. This is invoked automatically by the moc generated code that exports the plugin.
*/
func NewQPictureFormatPlugin(parent qtcore.QObject_ITF /*777 QObject **/) *QPictureFormatPlugin {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QObject_PTR() != nil {
		convArg0 = parent.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN20QPictureFormatPluginC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQPictureFormatPluginFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QPictureFormatPlugin")
	return gothis
}

// /usr/include/qt/QtGui/qpictureformatplugin.h:63
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QPictureFormatPlugin(QObject *)

/*
Constructs an picture format plugin with the given parent. This is invoked automatically by the moc generated code that exports the plugin.
*/
func NewQPictureFormatPlugin__() *QPictureFormatPlugin {
	// arg: 0, QObject *=Pointer, QObject=Record, , Invalid
	var convArg0 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN20QPictureFormatPluginC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQPictureFormatPluginFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QPictureFormatPlugin")
	return gothis
}

// /usr/include/qt/QtGui/qpictureformatplugin.h:64
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QPictureFormatPlugin()

/*

 */
func DeleteQPictureFormatPlugin(this *QPictureFormatPlugin) {
	rv, err := qtrt.InvokeQtFunc6("_ZN20QPictureFormatPluginD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 16)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtGui/qpictureformatplugin.h:66
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool loadPicture(const QString &, const QString &, QPicture *)

/*
Loads the picture stored in the file called fileName, with the given format, into *picture. Returns true on success; otherwise returns false.

See also savePicture().
*/
func (this *QPictureFormatPlugin) LoadPicture(format string, filename string, pic QPicture_ITF /*777 QPicture **/) bool {
	var tmpArg0 = qtcore.NewQString_5(format)
	var convArg0 = tmpArg0.GetCthis()
	var tmpArg1 = qtcore.NewQString_5(filename)
	var convArg1 = tmpArg1.GetCthis()
	var convArg2 unsafe.Pointer
	if pic != nil && pic.QPicture_PTR() != nil {
		convArg2 = pic.QPicture_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN20QPictureFormatPlugin11loadPictureERK7QStringS2_P8QPicture", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qpictureformatplugin.h:67
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool savePicture(const QString &, const QString &, const QPicture &)

/*
Saves the given picture into the file called fileName, using the specified format. Returns true on success; otherwise returns false.

See also loadPicture().
*/
func (this *QPictureFormatPlugin) SavePicture(format string, filename string, pic QPicture_ITF) bool {
	var tmpArg0 = qtcore.NewQString_5(format)
	var convArg0 = tmpArg0.GetCthis()
	var tmpArg1 = qtcore.NewQString_5(filename)
	var convArg1 = tmpArg1.GetCthis()
	var convArg2 unsafe.Pointer
	if pic != nil && pic.QPicture_PTR() != nil {
		convArg2 = pic.QPicture_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN20QPictureFormatPlugin11savePictureERK7QStringS2_RK8QPicture", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qpictureformatplugin.h:68
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [1] bool installIOHandler(const QString &)

/*
Installs a QPictureIO picture I/O handler for the picture format format. Returns true on success.
*/
func (this *QPictureFormatPlugin) InstallIOHandler(format string) bool {
	var tmpArg0 = qtcore.NewQString_5(format)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN20QPictureFormatPlugin16installIOHandlerERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
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
