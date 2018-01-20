//  header block begin
// /usr/include/qt/QtCore/qvariant.h
// #include <qvariant.h>
// #include <QtCore>
package qtcore

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 5
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
type QAssociativeIterable struct {
	*qtrt.CObject
}

func (this *QAssociativeIterable) GetCthis() unsafe.Pointer {
	return this.Cthis
}

// /usr/include/qt/QtCore/qvariant.h:680
// index:0
// QAssociativeIterable::const_iterator begin()
func (this *QAssociativeIterable) Begin() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK20QAssociativeIterable5beginEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qvariant.h:681
// index:0
// QAssociativeIterable::const_iterator end()
func (this *QAssociativeIterable) End() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK20QAssociativeIterable3endEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qvariant.h:682
// index:0
// QAssociativeIterable::const_iterator find(const class QVariant &)
func (this *QAssociativeIterable) Find(key unsafe.Pointer) {
	// 0: (, key const QVariant &), (key)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK20QAssociativeIterable4findERK8QVariant", ffiqt.FFI_TYPE_VOID, this.GetCthis(), key)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qvariant.h:684
// index:0
// QVariant value(const class QVariant &)
func (this *QAssociativeIterable) Value(key unsafe.Pointer) {
	// 0: (, key const QVariant &), (key)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK20QAssociativeIterable5valueERK8QVariant", ffiqt.FFI_TYPE_VOID, this.GetCthis(), key)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qvariant.h:686
// index:0
// int size()
func (this *QAssociativeIterable) Size() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK20QAssociativeIterable4sizeEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

//  body block end
