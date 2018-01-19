//  header block begin
// /usr/include/qt/QtCore/qxmlstream.h
// #include <qxmlstream.h>
// #include <QtCore>
package qtcore

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 7
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
type QXmlStreamEntityResolver struct {
	cthis unsafe.Pointer
}

// /usr/include/qt/QtCore/qxmlstream.h:336
// index:0
// virtual
// void ~QXmlStreamEntityResolver()
func DeleteQXmlStreamEntityResolver(*QXmlStreamEntityResolver) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN24QXmlStreamEntityResolverD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qxmlstream.h:337
// index:0
// virtual
// QString resolveEntity(const class QString &, const class QString &)
func (this *QXmlStreamEntityResolver) ResolveEntity(publicId unsafe.Pointer, systemId unsafe.Pointer) {
	// 0: (, const QString & publicId, const QString & systemId), (publicId, systemId)
	rv, err := ffiqt.InvokeQtFunc6("_ZN24QXmlStreamEntityResolver13resolveEntityERK7QStringS2_", ffiqt.FFI_TYPE_VOID, this.cthis, publicId, systemId)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qxmlstream.h:338
// index:0
// virtual
// QString resolveUndeclaredEntity(const class QString &)
func (this *QXmlStreamEntityResolver) ResolveUndeclaredEntity(name unsafe.Pointer) {
	// 0: (, const QString & name), (name)
	rv, err := ffiqt.InvokeQtFunc6("_ZN24QXmlStreamEntityResolver23resolveUndeclaredEntityERK7QString", ffiqt.FFI_TYPE_VOID, this.cthis, name)
	gopp.ErrPrint(err, rv)
}

//  body block end
