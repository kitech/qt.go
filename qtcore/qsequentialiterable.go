package qtcore

// /usr/include/qt/QtCore/qvariant.h
// #include <qvariant.h>
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

type QSequentialIterable struct {
	*qtrt.CObject
}

func (this *QSequentialIterable) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QSequentialIterable) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQSequentialIterableFromPointer(cthis unsafe.Pointer) *QSequentialIterable {
	return &QSequentialIterable{&qtrt.CObject{cthis}}
}
func (*QSequentialIterable) NewFromPointer(cthis unsafe.Pointer) *QSequentialIterable {
	return NewQSequentialIterableFromPointer(cthis)
}

// /usr/include/qt/QtCore/qvariant.h:623
// index:0
// Public Visibility=Default Availability=Available
// [112] QSequentialIterable::const_iterator begin()
func (this *QSequentialIterable) Begin() unsafe.Pointer /*444*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QSequentialIterable5beginEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return unsafe.Pointer(uintptr(rv))
}

// /usr/include/qt/QtCore/qvariant.h:624
// index:0
// Public Visibility=Default Availability=Available
// [112] QSequentialIterable::const_iterator end()
func (this *QSequentialIterable) End() unsafe.Pointer /*444*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QSequentialIterable3endEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return unsafe.Pointer(uintptr(rv))
}

// /usr/include/qt/QtCore/qvariant.h:626
// index:0
// Public Visibility=Default Availability=Available
// [16] QVariant at(int)
func (this *QSequentialIterable) At(idx int) *QVariant /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QSequentialIterable2atEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), idx)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQVariantFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQVariant)
	return rv2
}

// /usr/include/qt/QtCore/qvariant.h:627
// index:0
// Public Visibility=Default Availability=Available
// [4] int size()
func (this *QSequentialIterable) Size() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QSequentialIterable4sizeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qvariant.h:629
// index:0
// Public Visibility=Default Availability=Available
// [1] bool canReverseIterate()
func (this *QSequentialIterable) CanReverseIterate() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QSequentialIterable17canReverseIterateEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

func DeleteQSequentialIterable(this *QSequentialIterable) {
	rv, err := qtrt.InvokeQtFunc6("_ZN19QSequentialIterableD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
	this.SetCthis(nil)
}

//  body block end
