package qtqml

// /usr/include/qt/QtQml/qqmllist.h
// #include <qqmllist.h>
// #include <QtQml>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 3
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "github.com/kitech/qt.go/qtrt"
import "github.com/kitech/qt.go/qtcore"
import "github.com/kitech/qt.go/qtnetwork"

//  ext block end

//  body block begin

type QQmlListReference struct {
	*qtrt.CObject
}

func (this *QQmlListReference) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QQmlListReference) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQQmlListReferenceFromPointer(cthis unsafe.Pointer) *QQmlListReference {
	return &QQmlListReference{&qtrt.CObject{cthis}}
}
func (*QQmlListReference) NewFromPointer(cthis unsafe.Pointer) *QQmlListReference {
	return NewQQmlListReferenceFromPointer(cthis)
}

// /usr/include/qt/QtQml/qqmllist.h:142
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QQmlListReference()
func NewQQmlListReference() *QQmlListReference {
	rv, err := qtrt.InvokeQtFunc6("_ZN17QQmlListReferenceC2Ev", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	gothis := NewQQmlListReferenceFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQQmlListReference)
	return gothis
}

// /usr/include/qt/QtQml/qqmllist.h:143
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QQmlListReference(QObject *, const char *, QQmlEngine *)
func NewQQmlListReference_1(arg0 *qtcore.QObject /*777 QObject **/, property string, arg2 *QQmlEngine /*777 QQmlEngine **/) *QQmlListReference {
	var convArg0 = arg0.GetCthis()
	var convArg1 = qtrt.CString(property)
	defer qtrt.FreeMem(convArg1)
	var convArg2 = arg2.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN17QQmlListReferenceC2EP7QObjectPKcP10QQmlEngine", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, convArg2)
	qtrt.ErrPrint(err, rv)
	gothis := NewQQmlListReferenceFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQQmlListReference)
	return gothis
}

// /usr/include/qt/QtQml/qqmllist.h:146
// index:0
// Public Visibility=Default Availability=Available
// [-2] void ~QQmlListReference()
func DeleteQQmlListReference(this *QQmlListReference) {
	rv, err := qtrt.InvokeQtFunc6("_ZN17QQmlListReferenceD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 8)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtQml/qqmllist.h:148
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isValid()
func (this *QQmlListReference) IsValid() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QQmlListReference7isValidEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtQml/qqmllist.h:150
// index:0
// Public Visibility=Default Availability=Available
// [8] QObject * object()
func (this *QQmlListReference) Object() *qtcore.QObject /*777 QObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QQmlListReference6objectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtQml/qqmllist.h:151
// index:0
// Public Visibility=Default Availability=Available
// [8] const QMetaObject * listElementType()
func (this *QQmlListReference) ListElementType() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QQmlListReference15listElementTypeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtQml/qqmllist.h:153
// index:0
// Public Visibility=Default Availability=Available
// [1] bool canAppend()
func (this *QQmlListReference) CanAppend() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QQmlListReference9canAppendEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtQml/qqmllist.h:154
// index:0
// Public Visibility=Default Availability=Available
// [1] bool canAt()
func (this *QQmlListReference) CanAt() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QQmlListReference5canAtEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtQml/qqmllist.h:155
// index:0
// Public Visibility=Default Availability=Available
// [1] bool canClear()
func (this *QQmlListReference) CanClear() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QQmlListReference8canClearEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtQml/qqmllist.h:156
// index:0
// Public Visibility=Default Availability=Available
// [1] bool canCount()
func (this *QQmlListReference) CanCount() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QQmlListReference8canCountEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtQml/qqmllist.h:158
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isManipulable()
func (this *QQmlListReference) IsManipulable() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QQmlListReference13isManipulableEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtQml/qqmllist.h:159
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isReadable()
func (this *QQmlListReference) IsReadable() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QQmlListReference10isReadableEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtQml/qqmllist.h:162
// index:0
// Public Visibility=Default Availability=Available
// [8] QObject * at(int)
func (this *QQmlListReference) At(arg0 int) *qtcore.QObject /*777 QObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QQmlListReference2atEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtQml/qqmllist.h:163
// index:0
// Public Visibility=Default Availability=Available
// [1] bool clear()
func (this *QQmlListReference) Clear() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QQmlListReference5clearEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtQml/qqmllist.h:164
// index:0
// Public Visibility=Default Availability=Available
// [4] int count()
func (this *QQmlListReference) Count() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QQmlListReference5countEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
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
	if false {
		qtcore.KeepMe()
	}
	if false {
		qtnetwork.KeepMe()
	}
}

//  keep block end
