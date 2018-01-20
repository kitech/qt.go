//  header block begin
// /usr/include/qt/QtGui/qpictureformatplugin.h
// #include <qpictureformatplugin.h>
// #include <QtGui>
package qtgui

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 29
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "qtrt"
import "mkuse/cffiqt"
import "gopp"
import "qtcore"

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
}

//  ext block end

//  body block begin
type QPictureFormatPlugin struct {
	*qtcore.QObject
}

func (this *QPictureFormatPlugin) GetCthis() unsafe.Pointer {
	return this.QObject.GetCthis()
}

// /usr/include/qt/QtGui/qpictureformatplugin.h:61
// index:0
// virtual
// const QMetaObject * metaObject()
func (this *QPictureFormatPlugin) MetaObject() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK20QPictureFormatPlugin10metaObjectEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpictureformatplugin.h:63
// index:0
// void QPictureFormatPlugin(class QObject *)
func NewQPictureFormatPlugin(parent unsafe.Pointer) *QPictureFormatPlugin {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN20QPictureFormatPluginC1EP7QObject", ffiqt.FFI_TYPE_VOID, cthis, parent)
	gopp.ErrPrint(err, rv)
	gothis := NewQPictureFormatPluginFromPointer(cthis)
	return gothis
}
func NewQPictureFormatPluginFromPointer(cthis unsafe.Pointer) *QPictureFormatPlugin {
	bcthis0 := qtcore.NewQObjectFromPointer(cthis)
	return &QPictureFormatPlugin{bcthis0}
}

// /usr/include/qt/QtGui/qpictureformatplugin.h:64
// index:0
// virtual
// void ~QPictureFormatPlugin()
func DeleteQPictureFormatPlugin(*QPictureFormatPlugin) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN20QPictureFormatPluginD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpictureformatplugin.h:66
// index:0
// virtual
// bool loadPicture(const class QString &, const class QString &, class QPicture *)
func (this *QPictureFormatPlugin) LoadPicture(format unsafe.Pointer, filename unsafe.Pointer, pic unsafe.Pointer) {
	// 0: (, format const QString &, filename const QString &, pic QPicture *), (format, filename, pic)
	rv, err := ffiqt.InvokeQtFunc6("_ZN20QPictureFormatPlugin11loadPictureERK7QStringS2_P8QPicture", ffiqt.FFI_TYPE_VOID, this.GetCthis(), format, filename, pic)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpictureformatplugin.h:67
// index:0
// virtual
// bool savePicture(const class QString &, const class QString &, const class QPicture &)
func (this *QPictureFormatPlugin) SavePicture(format unsafe.Pointer, filename unsafe.Pointer, pic unsafe.Pointer) {
	// 0: (, format const QString &, filename const QString &, pic const QPicture &), (format, filename, pic)
	rv, err := ffiqt.InvokeQtFunc6("_ZN20QPictureFormatPlugin11savePictureERK7QStringS2_RK8QPicture", ffiqt.FFI_TYPE_VOID, this.GetCthis(), format, filename, pic)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpictureformatplugin.h:68
// index:0
// pure virtual
// bool installIOHandler(const class QString &)
func (this *QPictureFormatPlugin) InstallIOHandler(format unsafe.Pointer) {
	// 0: (, format const QString &), (format)
	rv, err := ffiqt.InvokeQtFunc6("_ZN20QPictureFormatPlugin16installIOHandlerERK7QString", ffiqt.FFI_TYPE_VOID, this.GetCthis(), format)
	gopp.ErrPrint(err, rv)
}

//  body block end
