package qtqml

// /usr/include/qt/QtQml/qqmllist.h
// #include <qqmllist.h>
// #include <QtQml>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 3
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "gopp"
import "qt.go/cffiqt"
import "qt.go/qtrt"
import "qt.go/qtcore"
import "qt.go/qtnetwork"

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
		qtnetwork.KeepMe()
	}
}

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
	this.CObject = &qtrt.CObject{cthis}
}
func NewQQmlListReferenceFromPointer(cthis unsafe.Pointer) *QQmlListReference {
	return &QQmlListReference{&qtrt.CObject{cthis}}
}
func (*QQmlListReference) NewFromPointer(cthis unsafe.Pointer) *QQmlListReference {
	return NewQQmlListReferenceFromPointer(cthis)
}

// /usr/include/qt/QtQml/qqmllist.h:142
// index:0
// Public
// void QQmlListReference()
func NewQQmlListReference() *QQmlListReference {
	cthis := qtrt.Calloc(1, 256) // 8
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QQmlListReferenceC2Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	gothis := NewQQmlListReferenceFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtQml/qqmllist.h:143
// index:1
// Public
// void QQmlListReference(QObject *, const char *, QQmlEngine *)
func NewQQmlListReference_1(arg0 *qtcore.QObject /*777 QObject **/, property string, arg2 *QQmlEngine /*777 QQmlEngine **/) *QQmlListReference {
	cthis := qtrt.Calloc(1, 256) // 8
	var convArg0 = arg0.GetCthis()
	var convArg1 = qtrt.CString(property)
	defer qtrt.FreeMem(convArg1)
	var convArg2 = arg2.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QQmlListReferenceC2EP7QObjectPKcP10QQmlEngine", ffiqt.FFI_TYPE_VOID, cthis, convArg0, convArg1, convArg2)
	gopp.ErrPrint(err, rv)
	gothis := NewQQmlListReferenceFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtQml/qqmllist.h:146
// index:0
// Public
// void ~QQmlListReference()
func DeleteQQmlListReference(*QQmlListReference) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QQmlListReferenceD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQml/qqmllist.h:148
// index:0
// Public
// bool isValid()
func (this *QQmlListReference) IsValid() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QQmlListReference7isValidEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtQml/qqmllist.h:150
// index:0
// Public
// QObject * object()
func (this *QQmlListReference) Object() *qtcore.QObject /*777 QObject **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QQmlListReference6objectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtQml/qqmllist.h:151
// index:0
// Public
// const QMetaObject * listElementType()
func (this *QQmlListReference) ListElementType() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QQmlListReference15listElementTypeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtQml/qqmllist.h:153
// index:0
// Public
// bool canAppend()
func (this *QQmlListReference) CanAppend() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QQmlListReference9canAppendEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtQml/qqmllist.h:154
// index:0
// Public
// bool canAt()
func (this *QQmlListReference) CanAt() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QQmlListReference5canAtEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtQml/qqmllist.h:155
// index:0
// Public
// bool canClear()
func (this *QQmlListReference) CanClear() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QQmlListReference8canClearEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtQml/qqmllist.h:156
// index:0
// Public
// bool canCount()
func (this *QQmlListReference) CanCount() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QQmlListReference8canCountEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtQml/qqmllist.h:158
// index:0
// Public
// bool isManipulable()
func (this *QQmlListReference) IsManipulable() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QQmlListReference13isManipulableEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtQml/qqmllist.h:159
// index:0
// Public
// bool isReadable()
func (this *QQmlListReference) IsReadable() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QQmlListReference10isReadableEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtQml/qqmllist.h:162
// index:0
// Public
// QObject * at(int)
func (this *QQmlListReference) At(arg0 int) *qtcore.QObject /*777 QObject **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QQmlListReference2atEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtQml/qqmllist.h:163
// index:0
// Public
// bool clear()
func (this *QQmlListReference) Clear() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QQmlListReference5clearEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtQml/qqmllist.h:164
// index:0
// Public
// int count()
func (this *QQmlListReference) Count() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QQmlListReference5countEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

//  body block end
