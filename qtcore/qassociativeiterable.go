package qtcore

// /usr/include/qt/QtCore/qvariant.h
// #include <qvariant.h>
// #include <QtCore>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 5
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
type QAssociativeIterable struct {
	*qtrt.CObject
}
type QAssociativeIterable_ITF interface {
	QAssociativeIterable_PTR() *QAssociativeIterable
}

func (ptr *QAssociativeIterable) QAssociativeIterable_PTR() *QAssociativeIterable { return ptr }

func (this *QAssociativeIterable) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QAssociativeIterable) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQAssociativeIterableFromPointer(cthis unsafe.Pointer) *QAssociativeIterable {
	return &QAssociativeIterable{&qtrt.CObject{cthis}}
}
func (*QAssociativeIterable) NewFromPointer(cthis unsafe.Pointer) *QAssociativeIterable {
	return NewQAssociativeIterableFromPointer(cthis)
}

// /usr/include/qt/QtCore/qvariant.h:705
// index:0
// Public Visibility=Default Availability=Available
// [120] QAssociativeIterable::const_iterator begin() const

/*

 */
func (this *QAssociativeIterable) Begin() unsafe.Pointer /*444*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK20QAssociativeIterable5beginEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return unsafe.Pointer(uintptr(rv))
}

// /usr/include/qt/QtCore/qvariant.h:706
// index:0
// Public Visibility=Default Availability=Available
// [120] QAssociativeIterable::const_iterator end() const

/*

 */
func (this *QAssociativeIterable) End() unsafe.Pointer /*444*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK20QAssociativeIterable3endEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return unsafe.Pointer(uintptr(rv))
}

// /usr/include/qt/QtCore/qvariant.h:707
// index:0
// Public Visibility=Default Availability=Available
// [120] QAssociativeIterable::const_iterator find(const QVariant &) const

/*

 */
func (this *QAssociativeIterable) Find(key QVariant_ITF) unsafe.Pointer /*444*/ {
	var convArg0 unsafe.Pointer
	if key != nil && key.QVariant_PTR() != nil {
		convArg0 = key.QVariant_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK20QAssociativeIterable4findERK8QVariant", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return unsafe.Pointer(uintptr(rv))
}

// /usr/include/qt/QtCore/qvariant.h:709
// index:0
// Public Visibility=Default Availability=Available
// [16] QVariant value(const QVariant &) const

/*
Returns the stored value converted to the template type T. Call canConvert() to find out whether a type can be converted. If the value cannot be converted, a default-constructed value will be returned.

If the type T is supported by QVariant, this function behaves exactly as toString(), toInt() etc.

Example:


  QVariant v;

  MyCustomStruct c;
  if (v.canConvert<MyCustomStruct>())
      c = v.value<MyCustomStruct>();

  v = 7;
  int i = v.value<int>();                        // same as v.toInt()
  QString s = v.value<QString>();                // same as v.toString(), s is now "7"
  MyCustomStruct c2 = v.value<MyCustomStruct>(); // conversion failed, c2 is empty



If the QVariant contains a pointer to a type derived from QObject then T may be any QObject type. If the pointer stored in the QVariant can be qobject_cast to T, then that result is returned. Otherwise a null pointer is returned. Note that this only works for QObject subclasses which use the Q_OBJECT macro.

If the QVariant contains a sequential container and T is QVariantList, the elements of the container will be converted into QVariants and returned as a QVariantList.


  QList<int> intList = {7, 11, 42};

  QVariant variant = QVariant::fromValue(intList);
  if (variant.canConvert<QVariantList>()) {
      QSequentialIterable iterable = variant.value<QSequentialIterable>();
      // Can use foreach:
      foreach (const QVariant &v, iterable) {
          qDebug() << v;
      }
      // Can use C++11 range-for:
      for (const QVariant &v : iterable) {
          qDebug() << v;
      }
      // Can use iterators:
      QSequentialIterable::const_iterator it = iterable.begin();
      const QSequentialIterable::const_iterator end = iterable.end();
      for ( ; it != end; ++it) {
          qDebug() << *it;
      }
  }



See also setValue(), fromValue(), canConvert(), and Q_DECLARE_SEQUENTIAL_CONTAINER_METATYPE().
*/
func (this *QAssociativeIterable) Value(key QVariant_ITF) *QVariant /*123*/ {
	var convArg0 unsafe.Pointer
	if key != nil && key.QVariant_PTR() != nil {
		convArg0 = key.QVariant_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK20QAssociativeIterable5valueERK8QVariant", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQVariantFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQVariant)
	return rv2
}

// /usr/include/qt/QtCore/qvariant.h:711
// index:0
// Public Visibility=Default Availability=Available
// [4] int size() const

/*

 */
func (this *QAssociativeIterable) Size() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK20QAssociativeIterable4sizeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

func DeleteQAssociativeIterable(this *QAssociativeIterable) {
	rv, err := qtrt.InvokeQtFunc6("_ZN20QAssociativeIterableD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

//  body block end

//  keep block begin

func init_unused_10281() {
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
