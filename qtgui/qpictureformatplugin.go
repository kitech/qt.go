package qtgui

// /usr/include/qt/QtGui/qpictureformatplugin.h
// #include <qpictureformatplugin.h>
// #include <QtGui>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 26
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "gopp"
import "qt.go/qtrt"
import "qt.go/qtcore"

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
}

//  ext block end

//  body block begin

type QPictureFormatPlugin struct {
	*qtcore.QObject
}

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
// [8] const QMetaObject * metaObject()
func (this *QPictureFormatPlugin) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK20QPictureFormatPlugin10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtGui/qpictureformatplugin.h:63
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QPictureFormatPlugin(QObject *)
func NewQPictureFormatPlugin(parent *qtcore.QObject /*777 QObject **/) *QPictureFormatPlugin {
	var convArg0 = parent.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN20QPictureFormatPluginC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQPictureFormatPluginFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtGui/qpictureformatplugin.h:64
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QPictureFormatPlugin()
func DeleteQPictureFormatPlugin(this *QPictureFormatPlugin) {
	rv, err := qtrt.InvokeQtFunc6("_ZN20QPictureFormatPluginD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 16)
	gopp.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtGui/qpictureformatplugin.h:66
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool loadPicture(const QString &, const QString &, QPicture *)
func (this *QPictureFormatPlugin) LoadPicture(format *qtcore.QString, filename *qtcore.QString, pic *QPicture /*777 QPicture **/) bool {
	var convArg0 = format.GetCthis()
	var convArg1 = filename.GetCthis()
	var convArg2 = pic.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN20QPictureFormatPlugin11loadPictureERK7QStringS2_P8QPicture", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtGui/qpictureformatplugin.h:67
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool savePicture(const QString &, const QString &, const QPicture &)
func (this *QPictureFormatPlugin) SavePicture(format *qtcore.QString, filename *qtcore.QString, pic *QPicture) bool {
	var convArg0 = format.GetCthis()
	var convArg1 = filename.GetCthis()
	var convArg2 = pic.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN20QPictureFormatPlugin11savePictureERK7QStringS2_RK8QPicture", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtGui/qpictureformatplugin.h:68
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [1] bool installIOHandler(const QString &)
func (this *QPictureFormatPlugin) InstallIOHandler(format *qtcore.QString) bool {
	var convArg0 = format.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN20QPictureFormatPlugin16installIOHandlerERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

//  body block end
