package qtcore

// /usr/include/qt/QtCore/qsignalmapper.h
// #include <qsignalmapper.h>
// #include <QtCore>

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
type QSignalMapper struct {
	*QObject
}

func (this *QSignalMapper) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QObject.GetCthis()
	}
}
func (this *QSignalMapper) SetCthis(cthis unsafe.Pointer) {
	this.QObject = NewQObjectFromPointer(cthis)
}
func NewQSignalMapperFromPointer(cthis unsafe.Pointer) *QSignalMapper {
	bcthis0 := NewQObjectFromPointer(cthis)
	return &QSignalMapper{bcthis0}
}
func (*QSignalMapper) NewFromPointer(cthis unsafe.Pointer) *QSignalMapper {
	return NewQSignalMapperFromPointer(cthis)
}

// /usr/include/qt/QtCore/qsignalmapper.h:53
// index:0
// Public virtual
// const QMetaObject * metaObject()
func (this *QSignalMapper) MetaObject() *QMetaObject /*777 const QMetaObject **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QSignalMapper10metaObjectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtCore/qsignalmapper.h:56
// index:0
// Public
// void QSignalMapper(class QObject *)
func NewQSignalMapper(parent *QObject /*777 QObject **/) *QSignalMapper {
	cthis := qtrt.Calloc(1, 256) // 16
	var convArg0 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QSignalMapperC2EP7QObject", ffiqt.FFI_TYPE_VOID, cthis, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQSignalMapperFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qsignalmapper.h:57
// index:0
// Public virtual
// void ~QSignalMapper()
func DeleteQSignalMapper(*QSignalMapper) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QSignalMapperD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsignalmapper.h:59
// index:0
// Public
// void setMapping(class QObject *, int)
func (this *QSignalMapper) SetMapping(sender *QObject /*777 QObject **/, id int) {
	var convArg0 = sender.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QSignalMapper10setMappingEP7QObjecti", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, id)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsignalmapper.h:60
// index:1
// Public
// void setMapping(class QObject *, const class QString &)
func (this *QSignalMapper) SetMapping_1(sender *QObject /*777 QObject **/, text *QString) {
	var convArg0 = sender.GetCthis()
	var convArg1 = text.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QSignalMapper10setMappingEP7QObjectRK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsignalmapper.h:61
// index:2
// Public
// void setMapping(class QObject *, class QWidget *)
func (this *QSignalMapper) SetMapping_2(sender *QObject /*777 QObject **/, widget unsafe.Pointer /*666*/) {
	var convArg0 = sender.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QSignalMapper10setMappingEP7QObjectP7QWidget", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, widget)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsignalmapper.h:62
// index:3
// Public
// void setMapping(class QObject *, class QObject *)
func (this *QSignalMapper) SetMapping_3(sender *QObject /*777 QObject **/, object *QObject /*777 QObject **/) {
	var convArg0 = sender.GetCthis()
	var convArg1 = object.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QSignalMapper10setMappingEP7QObjectS1_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsignalmapper.h:63
// index:0
// Public
// void removeMappings(class QObject *)
func (this *QSignalMapper) RemoveMappings(sender *QObject /*777 QObject **/) {
	var convArg0 = sender.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QSignalMapper14removeMappingsEP7QObject", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsignalmapper.h:65
// index:0
// Public
// QObject * mapping(int)
func (this *QSignalMapper) Mapping(id int) *QObject /*777 QObject **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QSignalMapper7mappingEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), id)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtCore/qsignalmapper.h:66
// index:1
// Public
// QObject * mapping(const class QString &)
func (this *QSignalMapper) Mapping_1(text *QString) *QObject /*777 QObject **/ {
	var convArg0 = text.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QSignalMapper7mappingERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtCore/qsignalmapper.h:67
// index:2
// Public
// QObject * mapping(class QWidget *)
func (this *QSignalMapper) Mapping_2(widget unsafe.Pointer /*666*/) *QObject /*777 QObject **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QSignalMapper7mappingEP7QWidget", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), widget)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtCore/qsignalmapper.h:68
// index:3
// Public
// QObject * mapping(class QObject *)
func (this *QSignalMapper) Mapping_3(object *QObject /*777 QObject **/) *QObject /*777 QObject **/ {
	var convArg0 = object.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QSignalMapper7mappingEP7QObject", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtCore/qsignalmapper.h:71
// index:0
// Public
// void mapped(int)
func (this *QSignalMapper) Mapped(arg0 int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QSignalMapper6mappedEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsignalmapper.h:72
// index:1
// Public
// void mapped(const class QString &)
func (this *QSignalMapper) Mapped_1(arg0 *QString) {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QSignalMapper6mappedERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsignalmapper.h:73
// index:2
// Public
// void mapped(class QWidget *)
func (this *QSignalMapper) Mapped_2(arg0 unsafe.Pointer /*666*/) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QSignalMapper6mappedEP7QWidget", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsignalmapper.h:74
// index:3
// Public
// void mapped(class QObject *)
func (this *QSignalMapper) Mapped_3(arg0 *QObject /*777 QObject **/) {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QSignalMapper6mappedEP7QObject", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsignalmapper.h:77
// index:0
// Public
// void map()
func (this *QSignalMapper) Map() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QSignalMapper3mapEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsignalmapper.h:78
// index:1
// Public
// void map(class QObject *)
func (this *QSignalMapper) Map_1(sender *QObject /*777 QObject **/) {
	var convArg0 = sender.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QSignalMapper3mapEP7QObject", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

//  body block end
