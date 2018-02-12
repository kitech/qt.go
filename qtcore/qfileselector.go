package qtcore

// /usr/include/qt/QtCore/qfileselector.h
// #include <qfileselector.h>
// #include <QtCore>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 2
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "github.com/kitech/qt.go/qtrt"

//  ext block end

//  body block begin

type QFileSelector struct {
	*QObject
}

func (this *QFileSelector) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QObject.GetCthis()
	}
}
func (this *QFileSelector) SetCthis(cthis unsafe.Pointer) {
	this.QObject = NewQObjectFromPointer(cthis)
}
func NewQFileSelectorFromPointer(cthis unsafe.Pointer) *QFileSelector {
	bcthis0 := NewQObjectFromPointer(cthis)
	return &QFileSelector{bcthis0}
}
func (*QFileSelector) NewFromPointer(cthis unsafe.Pointer) *QFileSelector {
	return NewQFileSelectorFromPointer(cthis)
}

// /usr/include/qt/QtCore/qfileselector.h:51
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject()
func (this *QFileSelector) MetaObject() *QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QFileSelector10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtCore/qfileselector.h:53
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QFileSelector(QObject *)
func NewQFileSelector(parent *QObject /*777 QObject **/) *QFileSelector {
	var convArg0 = parent.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN13QFileSelectorC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQFileSelectorFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtCore/qfileselector.h:54
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QFileSelector()
func DeleteQFileSelector(this *QFileSelector) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QFileSelectorD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 16)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtCore/qfileselector.h:56
// index:0
// Public Visibility=Default Availability=Available
// [8] QString select(const QString &)
func (this *QFileSelector) Select(filePath string) string {
	var tmpArg0 = NewQString_5(filePath)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QFileSelector6selectERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qfileselector.h:57
// index:1
// Public Visibility=Default Availability=Available
// [8] QUrl select(const QUrl &)
func (this *QFileSelector) Select_1(filePath *QUrl) *QUrl /*123*/ {
	var convArg0 = filePath.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QFileSelector6selectERK4QUrl", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQUrlFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQUrl)
	return rv2
}

// /usr/include/qt/QtCore/qfileselector.h:59
// index:0
// Public Visibility=Default Availability=Available
// [8] QStringList extraSelectors()
func (this *QFileSelector) ExtraSelectors() *QStringList /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QFileSelector14extraSelectorsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringListFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQStringList)
	return rv2
}

// /usr/include/qt/QtCore/qfileselector.h:60
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setExtraSelectors(const QStringList &)
func (this *QFileSelector) SetExtraSelectors(list *QStringList) {
	var convArg0 = list.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN13QFileSelector17setExtraSelectorsERK11QStringList", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qfileselector.h:62
// index:0
// Public Visibility=Default Availability=Available
// [8] QStringList allSelectors()
func (this *QFileSelector) AllSelectors() *QStringList /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QFileSelector12allSelectorsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringListFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQStringList)
	return rv2
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
		qtrt.KeepMe()
	}
}

//  keep block end
