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
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject()
func (this *QSignalMapper) MetaObject() *QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QSignalMapper10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtCore/qsignalmapper.h:56
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QSignalMapper(QObject *)
func NewQSignalMapper(parent *QObject /*777 QObject **/) *QSignalMapper {
	var convArg0 = parent.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN13QSignalMapperC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQSignalMapperFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtCore/qsignalmapper.h:57
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QSignalMapper()
func DeleteQSignalMapper(this *QSignalMapper) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QSignalMapperD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 16)
	gopp.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtCore/qsignalmapper.h:59
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setMapping(QObject *, int)
func (this *QSignalMapper) SetMapping(sender *QObject /*777 QObject **/, id int) {
	var convArg0 = sender.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN13QSignalMapper10setMappingEP7QObjecti", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, id)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsignalmapper.h:60
// index:1
// Public Visibility=Default Availability=Available
// [-2] void setMapping(QObject *, const QString &)
func (this *QSignalMapper) SetMapping_1(sender *QObject /*777 QObject **/, text *QString) {
	var convArg0 = sender.GetCthis()
	var convArg1 = text.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN13QSignalMapper10setMappingEP7QObjectRK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsignalmapper.h:61
// index:2
// Public Visibility=Default Availability=Available
// [-2] void setMapping(QObject *, QWidget *)
func (this *QSignalMapper) SetMapping_2(sender *QObject /*777 QObject **/, widget unsafe.Pointer /*666*/) {
	var convArg0 = sender.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN13QSignalMapper10setMappingEP7QObjectP7QWidget", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, widget)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsignalmapper.h:62
// index:3
// Public Visibility=Default Availability=Available
// [-2] void setMapping(QObject *, QObject *)
func (this *QSignalMapper) SetMapping_3(sender *QObject /*777 QObject **/, object *QObject /*777 QObject **/) {
	var convArg0 = sender.GetCthis()
	var convArg1 = object.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN13QSignalMapper10setMappingEP7QObjectS1_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsignalmapper.h:63
// index:0
// Public Visibility=Default Availability=Available
// [-2] void removeMappings(QObject *)
func (this *QSignalMapper) RemoveMappings(sender *QObject /*777 QObject **/) {
	var convArg0 = sender.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN13QSignalMapper14removeMappingsEP7QObject", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsignalmapper.h:65
// index:0
// Public Visibility=Default Availability=Available
// [8] QObject * mapping(int)
func (this *QSignalMapper) Mapping(id int) *QObject /*777 QObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QSignalMapper7mappingEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), id)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtCore/qsignalmapper.h:66
// index:1
// Public Visibility=Default Availability=Available
// [8] QObject * mapping(const QString &)
func (this *QSignalMapper) Mapping_1(text *QString) *QObject /*777 QObject **/ {
	var convArg0 = text.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QSignalMapper7mappingERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtCore/qsignalmapper.h:67
// index:2
// Public Visibility=Default Availability=Available
// [8] QObject * mapping(QWidget *)
func (this *QSignalMapper) Mapping_2(widget unsafe.Pointer /*666*/) *QObject /*777 QObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QSignalMapper7mappingEP7QWidget", qtrt.FFI_TYPE_POINTER, this.GetCthis(), widget)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtCore/qsignalmapper.h:68
// index:3
// Public Visibility=Default Availability=Available
// [8] QObject * mapping(QObject *)
func (this *QSignalMapper) Mapping_3(object *QObject /*777 QObject **/) *QObject /*777 QObject **/ {
	var convArg0 = object.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QSignalMapper7mappingEP7QObject", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtCore/qsignalmapper.h:71
// index:0
// Public Visibility=Default Availability=Available
// [-2] void mapped(int)
func (this *QSignalMapper) Mapped(arg0 int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QSignalMapper6mappedEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsignalmapper.h:72
// index:1
// Public Visibility=Default Availability=Available
// [-2] void mapped(const QString &)
func (this *QSignalMapper) Mapped_1(arg0 *QString) {
	var convArg0 = arg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN13QSignalMapper6mappedERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsignalmapper.h:73
// index:2
// Public Visibility=Default Availability=Available
// [-2] void mapped(QWidget *)
func (this *QSignalMapper) Mapped_2(arg0 unsafe.Pointer /*666*/) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QSignalMapper6mappedEP7QWidget", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsignalmapper.h:74
// index:3
// Public Visibility=Default Availability=Available
// [-2] void mapped(QObject *)
func (this *QSignalMapper) Mapped_3(arg0 *QObject /*777 QObject **/) {
	var convArg0 = arg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN13QSignalMapper6mappedEP7QObject", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsignalmapper.h:77
// index:0
// Public Visibility=Default Availability=Available
// [-2] void map()
func (this *QSignalMapper) Map() {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QSignalMapper3mapEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsignalmapper.h:78
// index:1
// Public Visibility=Default Availability=Available
// [-2] void map(QObject *)
func (this *QSignalMapper) Map_1(sender *QObject /*777 QObject **/) {
	var convArg0 = sender.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN13QSignalMapper3mapEP7QObject", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

//  body block end
