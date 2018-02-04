package qtcore

// /usr/include/qt/QtCore/qtemporarydir.h
// #include <qtemporarydir.h>
// #include <QtCore>

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
import "gopp"
import "qt.go/qtrt"

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
}

//  ext block end

//  body block begin

type QTemporaryDir struct {
	*qtrt.CObject
}

func (this *QTemporaryDir) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QTemporaryDir) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQTemporaryDirFromPointer(cthis unsafe.Pointer) *QTemporaryDir {
	return &QTemporaryDir{&qtrt.CObject{cthis}}
}
func (*QTemporaryDir) NewFromPointer(cthis unsafe.Pointer) *QTemporaryDir {
	return NewQTemporaryDirFromPointer(cthis)
}

// /usr/include/qt/QtCore/qtemporarydir.h:56
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QTemporaryDir()
func NewQTemporaryDir() *QTemporaryDir {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QTemporaryDirC2Ev", qtrt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	gothis := NewQTemporaryDirFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQTemporaryDir)
	return gothis
}

// /usr/include/qt/QtCore/qtemporarydir.h:57
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QTemporaryDir(const QString &)
func NewQTemporaryDir_1(templateName *QString) *QTemporaryDir {
	var convArg0 = templateName.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN13QTemporaryDirC2ERK7QString", qtrt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQTemporaryDirFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQTemporaryDir)
	return gothis
}

// /usr/include/qt/QtCore/qtemporarydir.h:58
// index:0
// Public Visibility=Default Availability=Available
// [-2] void ~QTemporaryDir()
func DeleteQTemporaryDir(this *QTemporaryDir) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QTemporaryDirD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 8)
	gopp.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtCore/qtemporarydir.h:60
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isValid()
func (this *QTemporaryDir) IsValid() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QTemporaryDir7isValidEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qtemporarydir.h:61
// index:0
// Public Visibility=Default Availability=Available
// [8] QString errorString()
func (this *QTemporaryDir) ErrorString() *QString /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QTemporaryDir11errorStringEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQString)
	return rv2
}

// /usr/include/qt/QtCore/qtemporarydir.h:63
// index:0
// Public Visibility=Default Availability=Available
// [1] bool autoRemove()
func (this *QTemporaryDir) AutoRemove() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QTemporaryDir10autoRemoveEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qtemporarydir.h:64
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setAutoRemove(_Bool)
func (this *QTemporaryDir) SetAutoRemove(b bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QTemporaryDir13setAutoRemoveEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), b)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qtemporarydir.h:65
// index:0
// Public Visibility=Default Availability=Available
// [1] bool remove()
func (this *QTemporaryDir) Remove() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QTemporaryDir6removeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qtemporarydir.h:67
// index:0
// Public Visibility=Default Availability=Available
// [8] QString path()
func (this *QTemporaryDir) Path() *QString /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QTemporaryDir4pathEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQString)
	return rv2
}

// /usr/include/qt/QtCore/qtemporarydir.h:68
// index:0
// Public Visibility=Default Availability=Available
// [8] QString filePath(const QString &)
func (this *QTemporaryDir) FilePath(fileName *QString) *QString /*123*/ {
	var convArg0 = fileName.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QTemporaryDir8filePathERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQString)
	return rv2
}

//  body block end
