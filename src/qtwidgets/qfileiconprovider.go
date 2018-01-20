//  header block begin
// /usr/include/qt/QtWidgets/qfileiconprovider.h
// #include <qfileiconprovider.h>
// #include <QtWidgets>
package qtwidgets

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 24
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "qtrt"
import "mkuse/cffiqt"
import "gopp"
import "qtcore"
import "qtgui"

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
		qtgui.KeepMe()
	}
}

//  ext block end

//  body block begin
type QFileIconProvider struct {
	*qtrt.CObject
}

func (this *QFileIconProvider) GetCthis() unsafe.Pointer {
	return this.Cthis
}
func NewQFileIconProviderFromPointer(cthis unsafe.Pointer) *QFileIconProvider {
	return &QFileIconProvider{&qtrt.CObject{cthis}}
}

// /usr/include/qt/QtWidgets/qfileiconprovider.h:56
// index:0
// Public
// void QFileIconProvider()
func NewQFileIconProvider() *QFileIconProvider {
	cthis := qtrt.Calloc(1, 256) // 16
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QFileIconProviderC2Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	gothis := NewQFileIconProviderFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtWidgets/qfileiconprovider.h:57
// index:0
// Public virtual
// void ~QFileIconProvider()
func DeleteQFileIconProvider(*QFileIconProvider) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QFileIconProviderD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qfileiconprovider.h:65
// index:0
// Public virtual
// QIcon icon(enum QFileIconProvider::IconType)
func (this *QFileIconProvider) Icon(type_ int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QFileIconProvider4iconENS_8IconTypeE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &type_)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qfileiconprovider.h:66
// index:1
// Public virtual
// QIcon icon(const class QFileInfo &)
func (this *QFileIconProvider) Icon_1(info *qtcore.QFileInfo) interface{} {
	var convArg0 = info.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QFileIconProvider4iconERK9QFileInfo", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qfileiconprovider.h:67
// index:0
// Public virtual
// QString type(const class QFileInfo &)
func (this *QFileIconProvider) Type(info *qtcore.QFileInfo) interface{} {
	var convArg0 = info.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QFileIconProvider4typeERK9QFileInfo", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qfileiconprovider.h:70
// index:0
// Public
// QFileIconProvider::Options options()
func (this *QFileIconProvider) Options() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QFileIconProvider7optionsEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

//  body block end
