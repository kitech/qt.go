package qtcore

// /usr/include/qt/QtCore/qfileselector.h
// #include <qfileselector.h>
// #include <QtCore>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 1
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
func NewQFileSelectorFromPointer(cthis unsafe.Pointer) *QFileSelector {
	bcthis0 := NewQObjectFromPointer(cthis)
	return &QFileSelector{bcthis0}
}

// /usr/include/qt/QtCore/qfileselector.h:51
// index:0
// Public virtual
// const QMetaObject * metaObject()
func (this *QFileSelector) MetaObject() *QMetaObject /*444 const QMetaObject **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QFileSelector10metaObjectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtCore/qfileselector.h:53
// index:0
// Public
// void QFileSelector(class QObject *)
func NewQFileSelector(parent *QObject /*444 QObject **/) *QFileSelector {
	cthis := qtrt.Calloc(1, 256) // 16
	var convArg0 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QFileSelectorC2EP7QObject", ffiqt.FFI_TYPE_VOID, cthis, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQFileSelectorFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qfileselector.h:54
// index:0
// Public virtual
// void ~QFileSelector()
func DeleteQFileSelector(*QFileSelector) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QFileSelectorD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qfileselector.h:56
// index:0
// Public
// QString select(const class QString &)
func (this *QFileSelector) Select(filePath *QString) *QString /*123*/ {
	var convArg0 = filePath.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QFileSelector6selectERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qfileselector.h:57
// index:1
// Public
// QUrl select(const class QUrl &)
func (this *QFileSelector) Select_1(filePath *QUrl) *QUrl /*123*/ {
	var convArg0 = filePath.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QFileSelector6selectERK4QUrl", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQUrlFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qfileselector.h:60
// index:0
// Public
// void setExtraSelectors(const class QStringList &)
func (this *QFileSelector) SetExtraSelectors(list *QStringList) {
	var convArg0 = list.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QFileSelector17setExtraSelectorsERK11QStringList", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

//  body block end
