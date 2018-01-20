//  header block begin
// /usr/include/qt/QtCore/qfileselector.h
// #include <qfileselector.h>
// #include <QtCore>
package qtcore

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 13
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
	return this.QObject.GetCthis()
}

// /usr/include/qt/QtCore/qfileselector.h:51
// index:0
// virtual
// const QMetaObject * metaObject()
func (this *QFileSelector) MetaObject() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QFileSelector10metaObjectEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qfileselector.h:53
// index:0
// void QFileSelector(class QObject *)
func NewQFileSelector(parent unsafe.Pointer) *QFileSelector {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QFileSelectorC2EP7QObject", ffiqt.FFI_TYPE_VOID, cthis, parent)
	gopp.ErrPrint(err, rv)
	gothis := NewQFileSelectorFromPointer(cthis)
	return gothis
}
func NewQFileSelectorFromPointer(cthis unsafe.Pointer) *QFileSelector {
	bcthis0 := NewQObjectFromPointer(cthis)
	return &QFileSelector{bcthis0}
}

// /usr/include/qt/QtCore/qfileselector.h:54
// index:0
// virtual
// void ~QFileSelector()
func DeleteQFileSelector(*QFileSelector) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QFileSelectorD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qfileselector.h:56
// index:0
// QString select(const class QString &)
func (this *QFileSelector) Select(filePath unsafe.Pointer) {
	// 0: (, filePath const QString &), (filePath)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QFileSelector6selectERK7QString", ffiqt.FFI_TYPE_VOID, this.GetCthis(), filePath)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qfileselector.h:57
// index:1
// QUrl select(const class QUrl &)
func (this *QFileSelector) Select_1(filePath unsafe.Pointer) {
	// 1: (, filePath const QUrl &), (filePath)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QFileSelector6selectERK4QUrl", ffiqt.FFI_TYPE_VOID, this.GetCthis(), filePath)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qfileselector.h:59
// index:0
// QStringList extraSelectors()
func (this *QFileSelector) ExtraSelectors() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QFileSelector14extraSelectorsEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qfileselector.h:60
// index:0
// void setExtraSelectors(const class QStringList &)
func (this *QFileSelector) SetExtraSelectors(list unsafe.Pointer) {
	// 0: (, list const QStringList &), (list)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QFileSelector17setExtraSelectorsERK11QStringList", ffiqt.FFI_TYPE_VOID, this.GetCthis(), list)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qfileselector.h:62
// index:0
// QStringList allSelectors()
func (this *QFileSelector) AllSelectors() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QFileSelector12allSelectorsEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

//  body block end
