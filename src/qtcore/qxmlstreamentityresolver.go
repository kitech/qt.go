package qtcore

// /usr/include/qt/QtCore/qxmlstream.h
// #include <qxmlstream.h>
// #include <QtCore>

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
	*qtrt.CObject
}

func (this *QXmlStreamEntityResolver) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QXmlStreamEntityResolver) SetCthis(cthis unsafe.Pointer) {
	this.CObject = &qtrt.CObject{cthis}
}
func NewQXmlStreamEntityResolverFromPointer(cthis unsafe.Pointer) *QXmlStreamEntityResolver {
	return &QXmlStreamEntityResolver{&qtrt.CObject{cthis}}
}
func (*QXmlStreamEntityResolver) NewFromPointer(cthis unsafe.Pointer) *QXmlStreamEntityResolver {
	return NewQXmlStreamEntityResolverFromPointer(cthis)
}

// /usr/include/qt/QtCore/qxmlstream.h:336
// index:0
// Public virtual
// void ~QXmlStreamEntityResolver()
func DeleteQXmlStreamEntityResolver(*QXmlStreamEntityResolver) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN24QXmlStreamEntityResolverD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qxmlstream.h:337
// index:0
// Public virtual
// QString resolveEntity(const class QString &, const class QString &)
func (this *QXmlStreamEntityResolver) ResolveEntity(publicId *QString, systemId *QString) *QString /*123*/ {
	mv := qtrt.Calloc(1, 256)
	var convArg0 = publicId.GetCthis()
	var convArg1 = systemId.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN24QXmlStreamEntityResolver13resolveEntityERK7QStringS2_", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qxmlstream.h:338
// index:0
// Public virtual
// QString resolveUndeclaredEntity(const class QString &)
func (this *QXmlStreamEntityResolver) ResolveUndeclaredEntity(name *QString) *QString /*123*/ {
	mv := qtrt.Calloc(1, 256)
	var convArg0 = name.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN24QXmlStreamEntityResolver23resolveUndeclaredEntityERK7QString", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

//  body block end
