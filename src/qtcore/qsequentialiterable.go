//  header block begin
// /usr/include/qt/QtCore/qvariant.h
// #include <qvariant.h>
// #include <QtCore>
package qtcore

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
	return this.Cthis
}

// /usr/include/qt/QtCore/qvariant.h:623
// index:0
// QSequentialIterable::const_iterator begin()
func (this *QSequentialIterable) Begin() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QSequentialIterable5beginEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qvariant.h:624
// index:0
// QSequentialIterable::const_iterator end()
func (this *QSequentialIterable) End() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QSequentialIterable3endEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qvariant.h:626
// index:0
// QVariant at(int)
func (this *QSequentialIterable) At(idx int) {
	// 0: (, idx int), (&idx)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QSequentialIterable2atEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &idx)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qvariant.h:627
// index:0
// int size()
func (this *QSequentialIterable) Size() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QSequentialIterable4sizeEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qvariant.h:629
// index:0
// bool canReverseIterate()
func (this *QSequentialIterable) CanReverseIterate() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QSequentialIterable17canReverseIterateEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

//  body block end
