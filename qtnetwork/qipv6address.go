package qtnetwork

// /usr/include/qt/QtNetwork/qhostaddress.h
// #include <qhostaddress.h>
// #include <QtNetwork>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 20
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "gopp"
import "qt.go/cffiqt"
import "qt.go/qtrt"
import "qt.go/qtcore"

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
}

//  ext block end

//  body block begin
type QIPv6Address struct {
	*qtrt.CObject
}

func (this *QIPv6Address) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QIPv6Address) SetCthis(cthis unsafe.Pointer) {
	this.CObject = &qtrt.CObject{cthis}
}
func NewQIPv6AddressFromPointer(cthis unsafe.Pointer) *QIPv6Address {
	return &QIPv6Address{&qtrt.CObject{cthis}}
}
func (*QIPv6Address) NewFromPointer(cthis unsafe.Pointer) *QIPv6Address {
	return NewQIPv6AddressFromPointer(cthis)
}

//  body block end
