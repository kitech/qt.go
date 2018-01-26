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
	this.CObject = &qtrt.CObject{cthis}
}
func NewQSequentialIterableFromPointer(cthis unsafe.Pointer) *QSequentialIterable {
	return &QSequentialIterable{&qtrt.CObject{cthis}}
}
func (*QSequentialIterable) NewFromPointer(cthis unsafe.Pointer) *QSequentialIterable {
	return NewQSequentialIterableFromPointer(cthis)
}

// /usr/include/qt/QtCore/qvariant.h:623
// index:0
// Public
// QSequentialIterable::const_iterator begin()
func (this *QSequentialIterable) Begin() unsafe.Pointer /*444*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QSequentialIterable5beginEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	return unsafe.Pointer(uintptr(rv))
}

// /usr/include/qt/QtCore/qvariant.h:624
// index:0
// Public
// QSequentialIterable::const_iterator end()
func (this *QSequentialIterable) End() unsafe.Pointer /*444*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QSequentialIterable3endEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	return unsafe.Pointer(uintptr(rv))
}

// /usr/include/qt/QtCore/qvariant.h:626
// index:0
// Public
// QVariant at(int)
func (this *QSequentialIterable) At(idx int) *QVariant /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QSequentialIterable2atEi", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), idx)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQVariantFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qvariant.h:627
// index:0
// Public
// int size()
func (this *QSequentialIterable) Size() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QSequentialIterable4sizeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtCore/qvariant.h:629
// index:0
// Public
// bool canReverseIterate()
func (this *QSequentialIterable) CanReverseIterate() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QSequentialIterable17canReverseIterateEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

//  body block end
