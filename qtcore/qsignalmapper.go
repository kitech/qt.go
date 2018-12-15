package qtcore

// /usr/include/qt/QtCore/qsignalmapper.h
// #include <qsignalmapper.h>
// #include <QtCore>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 20
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "log"
import "github.com/kitech/qt.go/qtrt"

//  ext block end

//  body block begin

/*

 */
type QSignalMapper struct {
	*QObject
}
type QSignalMapper_ITF interface {
	QObject_ITF
	QSignalMapper_PTR() *QSignalMapper
}

func (ptr *QSignalMapper) QSignalMapper_PTR() *QSignalMapper { return ptr }

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
// [8] const QMetaObject * metaObject() const

/*

 */
func (this *QSignalMapper) MetaObject() *QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QSignalMapper10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtCore/qsignalmapper.h:56
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QSignalMapper(QObject *)

/*
Constructs a QSignalMapper with parent parent.
*/
func (*QSignalMapper) NewForInherit(parent QObject_ITF /*777 QObject **/) *QSignalMapper {
	return NewQSignalMapper(parent)
}
func NewQSignalMapper(parent QObject_ITF /*777 QObject **/) *QSignalMapper {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QObject_PTR() != nil {
		convArg0 = parent.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QSignalMapperC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQSignalMapperFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QSignalMapper")
	return gothis
}

// /usr/include/qt/QtCore/qsignalmapper.h:56
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QSignalMapper(QObject *)

/*
Constructs a QSignalMapper with parent parent.
*/
func (*QSignalMapper) NewForInheritp() *QSignalMapper {
	return NewQSignalMapperp()
}
func NewQSignalMapperp() *QSignalMapper {
	// arg: 0, QObject *=Pointer, QObject=Record, , Invalid
	var convArg0 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN13QSignalMapperC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQSignalMapperFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QSignalMapper")
	return gothis
}

// /usr/include/qt/QtCore/qsignalmapper.h:57
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QSignalMapper()

/*

 */
func DeleteQSignalMapper(this *QSignalMapper) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QSignalMapperD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 16)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtCore/qsignalmapper.h:59
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setMapping(QObject *, int)

/*
Adds a mapping so that when map() is signalled from the given sender, the signal mapped(id) is emitted.

There may be at most one integer ID for each sender.

See also mapping().
*/
func (this *QSignalMapper) SetMapping(sender QObject_ITF /*777 QObject **/, id int) {
	var convArg0 unsafe.Pointer
	if sender != nil && sender.QObject_PTR() != nil {
		convArg0 = sender.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QSignalMapper10setMappingEP7QObjecti", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, id)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsignalmapper.h:60
// index:1
// Public Visibility=Default Availability=Available
// [-2] void setMapping(QObject *, const QString &)

/*
Adds a mapping so that when map() is signalled from the given sender, the signal mapped(id) is emitted.

There may be at most one integer ID for each sender.

See also mapping().
*/
func (this *QSignalMapper) SetMapping1(sender QObject_ITF /*777 QObject **/, text string) {
	var convArg0 unsafe.Pointer
	if sender != nil && sender.QObject_PTR() != nil {
		convArg0 = sender.QObject_PTR().GetCthis()
	}
	var tmpArg1 = NewQString5(text)
	var convArg1 = tmpArg1.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN13QSignalMapper10setMappingEP7QObjectRK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsignalmapper.h:61
// index:2
// Public Visibility=Default Availability=Available
// [-2] void setMapping(QObject *, QWidget *)

/*
Adds a mapping so that when map() is signalled from the given sender, the signal mapped(id) is emitted.

There may be at most one integer ID for each sender.

See also mapping().
*/
func (this *QSignalMapper) SetMapping2(sender QObject_ITF /*777 QObject **/, widget unsafe.Pointer /*666*/) {
	var convArg0 unsafe.Pointer
	if sender != nil && sender.QObject_PTR() != nil {
		convArg0 = sender.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QSignalMapper10setMappingEP7QObjectP7QWidget", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, widget)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsignalmapper.h:62
// index:3
// Public Visibility=Default Availability=Available
// [-2] void setMapping(QObject *, QObject *)

/*
Adds a mapping so that when map() is signalled from the given sender, the signal mapped(id) is emitted.

There may be at most one integer ID for each sender.

See also mapping().
*/
func (this *QSignalMapper) SetMapping3(sender QObject_ITF /*777 QObject **/, object QObject_ITF /*777 QObject **/) {
	var convArg0 unsafe.Pointer
	if sender != nil && sender.QObject_PTR() != nil {
		convArg0 = sender.QObject_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if object != nil && object.QObject_PTR() != nil {
		convArg1 = object.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QSignalMapper10setMappingEP7QObjectS1_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsignalmapper.h:63
// index:0
// Public Visibility=Default Availability=Available
// [-2] void removeMappings(QObject *)

/*
Removes all mappings for sender.

This is done automatically when mapped objects are destroyed.

Note: This does not disconnect any signals. If sender is not destroyed then this will need to be done explicitly if required.
*/
func (this *QSignalMapper) RemoveMappings(sender QObject_ITF /*777 QObject **/) {
	var convArg0 unsafe.Pointer
	if sender != nil && sender.QObject_PTR() != nil {
		convArg0 = sender.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QSignalMapper14removeMappingsEP7QObject", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsignalmapper.h:65
// index:0
// Public Visibility=Default Availability=Available
// [8] QObject * mapping(int) const

/*
Returns the sender QObject that is associated with the id.

See also setMapping().
*/
func (this *QSignalMapper) Mapping(id int) *QObject /*777 QObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QSignalMapper7mappingEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), id)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtCore/qsignalmapper.h:66
// index:1
// Public Visibility=Default Availability=Available
// [8] QObject * mapping(const QString &) const

/*
Returns the sender QObject that is associated with the id.

See also setMapping().
*/
func (this *QSignalMapper) Mapping1(text string) *QObject /*777 QObject **/ {
	var tmpArg0 = NewQString5(text)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QSignalMapper7mappingERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtCore/qsignalmapper.h:67
// index:2
// Public Visibility=Default Availability=Available
// [8] QObject * mapping(QWidget *) const

/*
Returns the sender QObject that is associated with the id.

See also setMapping().
*/
func (this *QSignalMapper) Mapping2(widget unsafe.Pointer /*666*/) *QObject /*777 QObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QSignalMapper7mappingEP7QWidget", qtrt.FFI_TYPE_POINTER, this.GetCthis(), widget)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtCore/qsignalmapper.h:68
// index:3
// Public Visibility=Default Availability=Available
// [8] QObject * mapping(QObject *) const

/*
Returns the sender QObject that is associated with the id.

See also setMapping().
*/
func (this *QSignalMapper) Mapping3(object QObject_ITF /*777 QObject **/) *QObject /*777 QObject **/ {
	var convArg0 unsafe.Pointer
	if object != nil && object.QObject_PTR() != nil {
		convArg0 = object.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QSignalMapper7mappingEP7QObject", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtCore/qsignalmapper.h:71
// index:0
// Public Visibility=Default Availability=Available
// [-2] void mapped(int)

/*
This signal is emitted when map() is signalled from an object that has an integer mapping set. The object's mapped integer is passed in i.

Note: Signal mapped is overloaded in this class. To connect to this signal by using the function pointer syntax, Qt provides a convenient helper for obtaining the function pointer as shown in this example:


  connect(signalMapper, QOverload<int>::of(&QSignalMapper::mapped),
      [=](int i){ /-* ... *-/ });



See also setMapping().
*/
func (this *QSignalMapper) Mapped(arg0 int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QSignalMapper6mappedEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsignalmapper.h:72
// index:1
// Public Visibility=Default Availability=Available
// [-2] void mapped(const QString &)

/*
This signal is emitted when map() is signalled from an object that has an integer mapping set. The object's mapped integer is passed in i.

Note: Signal mapped is overloaded in this class. To connect to this signal by using the function pointer syntax, Qt provides a convenient helper for obtaining the function pointer as shown in this example:


  connect(signalMapper, QOverload<int>::of(&QSignalMapper::mapped),
      [=](int i){ /-* ... *-/ });



See also setMapping().
*/
func (this *QSignalMapper) Mapped1(arg0 string) {
	var tmpArg0 = NewQString5(arg0)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN13QSignalMapper6mappedERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsignalmapper.h:73
// index:2
// Public Visibility=Default Availability=Available
// [-2] void mapped(QWidget *)

/*
This signal is emitted when map() is signalled from an object that has an integer mapping set. The object's mapped integer is passed in i.

Note: Signal mapped is overloaded in this class. To connect to this signal by using the function pointer syntax, Qt provides a convenient helper for obtaining the function pointer as shown in this example:


  connect(signalMapper, QOverload<int>::of(&QSignalMapper::mapped),
      [=](int i){ /-* ... *-/ });



See also setMapping().
*/
func (this *QSignalMapper) Mapped2(arg0 unsafe.Pointer /*666*/) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QSignalMapper6mappedEP7QWidget", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsignalmapper.h:74
// index:3
// Public Visibility=Default Availability=Available
// [-2] void mapped(QObject *)

/*
This signal is emitted when map() is signalled from an object that has an integer mapping set. The object's mapped integer is passed in i.

Note: Signal mapped is overloaded in this class. To connect to this signal by using the function pointer syntax, Qt provides a convenient helper for obtaining the function pointer as shown in this example:


  connect(signalMapper, QOverload<int>::of(&QSignalMapper::mapped),
      [=](int i){ /-* ... *-/ });



See also setMapping().
*/
func (this *QSignalMapper) Mapped3(arg0 QObject_ITF /*777 QObject **/) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QObject_PTR() != nil {
		convArg0 = arg0.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QSignalMapper6mappedEP7QObject", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsignalmapper.h:77
// index:0
// Public Visibility=Default Availability=Available
// [-2] void map()

/*
This slot emits signals based on which object sends signals to it.
*/
func (this *QSignalMapper) Map() {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QSignalMapper3mapEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsignalmapper.h:78
// index:1
// Public Visibility=Default Availability=Available
// [-2] void map(QObject *)

/*
This slot emits signals based on which object sends signals to it.
*/
func (this *QSignalMapper) Map1(sender QObject_ITF /*777 QObject **/) {
	var convArg0 unsafe.Pointer
	if sender != nil && sender.QObject_PTR() != nil {
		convArg0 = sender.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QSignalMapper3mapEP7QObject", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

//  body block end

//  keep block begin

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
		log.Println(123)
	}
	if false {
		qtrt.KeepMe()
	}
}

//  keep block end
