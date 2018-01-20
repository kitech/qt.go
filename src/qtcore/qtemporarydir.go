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
	*qtrt.CObject
}

func (this *QTemporaryDir) GetCthis() unsafe.Pointer {
	return this.Cthis
}
func NewQTemporaryDirFromPointer(cthis unsafe.Pointer) *QTemporaryDir {
	return &QTemporaryDir{&qtrt.CObject{cthis}}
}

// /usr/include/qt/QtCore/qtemporarydir.h:56
// index:0
// Public
// void QTemporaryDir()
func NewQTemporaryDir() *QTemporaryDir {
	cthis := qtrt.Calloc(1, 256) // 8
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QTemporaryDirC2Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	gothis := NewQTemporaryDirFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qtemporarydir.h:57
// index:1
// Public
// void QTemporaryDir(const class QString &)
func NewQTemporaryDir_1(templateName *QString) *QTemporaryDir {
	cthis := qtrt.Calloc(1, 256) // 8
	var convArg0 = templateName.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QTemporaryDirC2ERK7QString", ffiqt.FFI_TYPE_VOID, cthis, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQTemporaryDirFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qtemporarydir.h:58
// index:0
// Public
// void ~QTemporaryDir()
func DeleteQTemporaryDir(*QTemporaryDir) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QTemporaryDirD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qtemporarydir.h:60
// index:0
// Public
// bool isValid()
func (this *QTemporaryDir) IsValid() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QTemporaryDir7isValidEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qtemporarydir.h:61
// index:0
// Public
// QString errorString()
func (this *QTemporaryDir) ErrorString() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QTemporaryDir11errorStringEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qtemporarydir.h:63
// index:0
// Public
// bool autoRemove()
func (this *QTemporaryDir) AutoRemove() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QTemporaryDir10autoRemoveEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qtemporarydir.h:64
// index:0
// Public
// void setAutoRemove(_Bool)
func (this *QTemporaryDir) SetAutoRemove(b bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QTemporaryDir13setAutoRemoveEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &b)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qtemporarydir.h:65
// index:0
// Public
// bool remove()
func (this *QTemporaryDir) Remove() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QTemporaryDir6removeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qtemporarydir.h:67
// index:0
// Public
// QString path()
func (this *QTemporaryDir) Path() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QTemporaryDir4pathEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qtemporarydir.h:68
// index:0
// Public
// QString filePath(const class QString &)
func (this *QTemporaryDir) FilePath(fileName *QString) interface{} {
	var convArg0 = fileName.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QTemporaryDir8filePathERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	return rv
}

//  body block end
