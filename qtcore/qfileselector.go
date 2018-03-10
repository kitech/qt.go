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
import "log"
import "github.com/kitech/qt.go/qtrt"

//  ext block end

//  body block begin

/*

 */
type QFileSelector struct {
	*QObject
}
type QFileSelector_ITF interface {
	QObject_ITF
	QFileSelector_PTR() *QFileSelector
}

func (ptr *QFileSelector) QFileSelector_PTR() *QFileSelector { return ptr }

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
// [8] const QMetaObject * metaObject() const

/*

 */
func (this *QFileSelector) MetaObject() *QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QFileSelector10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtCore/qfileselector.h:53
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QFileSelector(QObject *)

/*
Create a QFileSelector instance. This instance will have the same static selectors as other QFileSelector instances, but its own set of extra selectors.

If supplied, it will have the given QObject parent.
*/
func NewQFileSelector(parent QObject_ITF /*777 QObject **/) *QFileSelector {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QObject_PTR() != nil {
		convArg0 = parent.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QFileSelectorC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQFileSelectorFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QFileSelector")
	return gothis
}

// /usr/include/qt/QtCore/qfileselector.h:53
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QFileSelector(QObject *)

/*
Create a QFileSelector instance. This instance will have the same static selectors as other QFileSelector instances, but its own set of extra selectors.

If supplied, it will have the given QObject parent.
*/
func NewQFileSelector__() *QFileSelector {
	// arg: 0, QObject *=Pointer, QObject=Record,
	var convArg0 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN13QFileSelectorC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQFileSelectorFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QFileSelector")
	return gothis
}

// /usr/include/qt/QtCore/qfileselector.h:54
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QFileSelector()

/*

 */
func DeleteQFileSelector(this *QFileSelector) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QFileSelectorD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 16)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtCore/qfileselector.h:56
// index:0
// Public Visibility=Default Availability=Available
// [8] QString select(const QString &) const

/*
This function returns the selected version of the path, based on the conditions at runtime. If no selectable files are present, returns the original filePath.

If the original file does not exist, the original filePath is returned. This means that you must have a base file to fall back on, you cannot have only files in selectable sub-directories.

See the class overview for the selection algorithm.
*/
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
// [8] QUrl select(const QUrl &) const

/*
This function returns the selected version of the path, based on the conditions at runtime. If no selectable files are present, returns the original filePath.

If the original file does not exist, the original filePath is returned. This means that you must have a base file to fall back on, you cannot have only files in selectable sub-directories.

See the class overview for the selection algorithm.
*/
func (this *QFileSelector) Select_1(filePath QUrl_ITF) *QUrl /*123*/ {
	var convArg0 unsafe.Pointer
	if filePath != nil && filePath.QUrl_PTR() != nil {
		convArg0 = filePath.QUrl_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QFileSelector6selectERK4QUrl", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQUrlFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQUrl)
	return rv2
}

// /usr/include/qt/QtCore/qfileselector.h:59
// index:0
// Public Visibility=Default Availability=Available
// [8] QStringList extraSelectors() const

/*
Returns the list of extra selectors which have been added programmatically to this instance.

See also setExtraSelectors().
*/
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

/*
Sets the list of extra selectors which have been added programmatically to this instance.

These selectors have priority over any which have been automatically picked up.

See also extraSelectors().
*/
func (this *QFileSelector) SetExtraSelectors(list QStringList_ITF) {
	var convArg0 unsafe.Pointer
	if list != nil && list.QStringList_PTR() != nil {
		convArg0 = list.QStringList_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QFileSelector17setExtraSelectorsERK11QStringList", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qfileselector.h:62
// index:0
// Public Visibility=Default Availability=Available
// [8] QStringList allSelectors() const

/*
Returns the complete, ordered list of selectors used by this instance
*/
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
		log.Println(123)
	}
	if false {
		qtrt.KeepMe()
	}
}

//  keep block end
