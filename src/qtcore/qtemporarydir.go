//  header block begin
// /usr/include/qt/QtCore/qtemporarydir.h
// #include <qtemporarydir.h>
// #include <QtCore>
package qtcore

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 8
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "qtrt"
import "mkuse/cffiqt"
import "gopp"

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
}

//  ext block end

//  body block begin
type QTemporaryDir struct {
	cthis unsafe.Pointer
}

// /usr/include/qt/QtCore/qtemporarydir.h:56
// index:0
// void QTemporaryDir()
func NewQTemporaryDir() *QTemporaryDir {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QTemporaryDirC2Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	return &QTemporaryDir{cthis}
}

// /usr/include/qt/QtCore/qtemporarydir.h:57
// index:1
// void QTemporaryDir(const class QString &)
func NewQTemporaryDir_1(templateName unsafe.Pointer) *QTemporaryDir {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QTemporaryDirC2ERK7QString", ffiqt.FFI_TYPE_VOID, cthis, templateName)
	gopp.ErrPrint(err, rv)
	return &QTemporaryDir{cthis}
}

// /usr/include/qt/QtCore/qtemporarydir.h:58
// index:0
// void ~QTemporaryDir()
func DeleteQTemporaryDir(*QTemporaryDir) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QTemporaryDirD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qtemporarydir.h:60
// index:0
// bool isValid()
func (this *QTemporaryDir) IsValid() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QTemporaryDir7isValidEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qtemporarydir.h:61
// index:0
// QString errorString()
func (this *QTemporaryDir) ErrorString() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QTemporaryDir11errorStringEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qtemporarydir.h:63
// index:0
// bool autoRemove()
func (this *QTemporaryDir) AutoRemove() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QTemporaryDir10autoRemoveEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qtemporarydir.h:64
// index:0
// void setAutoRemove(_Bool)
func (this *QTemporaryDir) SetAutoRemove(b bool) {
	// 0: (, bool b), (&b)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QTemporaryDir13setAutoRemoveEb", ffiqt.FFI_TYPE_VOID, this.cthis, &b)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qtemporarydir.h:65
// index:0
// bool remove()
func (this *QTemporaryDir) Remove() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QTemporaryDir6removeEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qtemporarydir.h:67
// index:0
// QString path()
func (this *QTemporaryDir) Path() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QTemporaryDir4pathEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qtemporarydir.h:68
// index:0
// QString filePath(const class QString &)
func (this *QTemporaryDir) FilePath(fileName unsafe.Pointer) {
	// 0: (, const QString & fileName), (fileName)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QTemporaryDir8filePathERK7QString", ffiqt.FFI_TYPE_VOID, this.cthis, fileName)
	gopp.ErrPrint(err, rv)
}

//  body block end
