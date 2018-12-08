package qtmultimedia

// /usr/include/qt/QtMultimedia/qabstractvideofilter.h
// #include <qabstractvideofilter.h>
// #include <QtMultimedia>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 2
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "log"
import "github.com/kitech/qt.go/qtrt"
import "github.com/kitech/qt.go/qtcore"
import "github.com/kitech/qt.go/qtgui"
import "github.com/kitech/qt.go/qtnetwork"

//  ext block end

//  body block begin

/*

 */
type QAbstractVideoFilter struct {
	*qtcore.QObject
}
type QAbstractVideoFilter_ITF interface {
	qtcore.QObject_ITF
	QAbstractVideoFilter_PTR() *QAbstractVideoFilter
}

func (ptr *QAbstractVideoFilter) QAbstractVideoFilter_PTR() *QAbstractVideoFilter { return ptr }

func (this *QAbstractVideoFilter) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QObject.GetCthis()
	}
}
func (this *QAbstractVideoFilter) SetCthis(cthis unsafe.Pointer) {
	this.QObject = qtcore.NewQObjectFromPointer(cthis)
}
func NewQAbstractVideoFilterFromPointer(cthis unsafe.Pointer) *QAbstractVideoFilter {
	bcthis0 := qtcore.NewQObjectFromPointer(cthis)
	return &QAbstractVideoFilter{bcthis0}
}
func (*QAbstractVideoFilter) NewFromPointer(cthis unsafe.Pointer) *QAbstractVideoFilter {
	return NewQAbstractVideoFilterFromPointer(cthis)
}

// /usr/include/qt/QtMultimedia/qabstractvideofilter.h:67
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject() const

/*

 */
func (this *QAbstractVideoFilter) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK20QAbstractVideoFilter10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtMultimedia/qabstractvideofilter.h:71
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QAbstractVideoFilter(QObject *)

/*
Constructs a new QAbstractVideoFilter instance with parent object parent.
*/
func (*QAbstractVideoFilter) NewForInherit(parent qtcore.QObject_ITF /*777 QObject **/) *QAbstractVideoFilter {
	return NewQAbstractVideoFilter(parent)
}
func NewQAbstractVideoFilter(parent qtcore.QObject_ITF /*777 QObject **/) *QAbstractVideoFilter {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QObject_PTR() != nil {
		convArg0 = parent.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN20QAbstractVideoFilterC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQAbstractVideoFilterFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QAbstractVideoFilter")
	return gothis
}

// /usr/include/qt/QtMultimedia/qabstractvideofilter.h:71
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QAbstractVideoFilter(QObject *)

/*
Constructs a new QAbstractVideoFilter instance with parent object parent.
*/
func (*QAbstractVideoFilter) NewForInheritp() *QAbstractVideoFilter {
	return NewQAbstractVideoFilterp()
}
func NewQAbstractVideoFilterp() *QAbstractVideoFilter {
	// arg: 0, QObject *=Pointer, QObject=Record, , Invalid
	var convArg0 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN20QAbstractVideoFilterC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQAbstractVideoFilterFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QAbstractVideoFilter")
	return gothis
}

// /usr/include/qt/QtMultimedia/qabstractvideofilter.h:72
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QAbstractVideoFilter()

/*

 */
func DeleteQAbstractVideoFilter(this *QAbstractVideoFilter) {
	rv, err := qtrt.InvokeQtFunc6("_ZN20QAbstractVideoFilterD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 24)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtMultimedia/qabstractvideofilter.h:74
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isActive() const

/*

 */
func (this *QAbstractVideoFilter) IsActive() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK20QAbstractVideoFilter8isActiveEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtMultimedia/qabstractvideofilter.h:75
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setActive(bool)

/*

 */
func (this *QAbstractVideoFilter) SetActive(v bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN20QAbstractVideoFilter9setActiveEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), v)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qabstractvideofilter.h:77
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [8] QVideoFilterRunnable * createFilterRunnable()

/*
Factory function to create a new instance of a QVideoFilterRunnable subclass corresponding to this filter.

This function is called on the thread on which the Qt Quick scene graph performs rendering, with the OpenGL context bound. Ownership of the returned instance is transferred: the returned instance will live on the render thread and will be destroyed automatically when necessary.

Typically, implementations of the function will simply construct a new QVideoFilterRunnable instance, passing this to the constructor as the filter runnables must know their associated QAbstractVideoFilter instance to access dynamic properties and optionally emit signals.
*/
func (this *QAbstractVideoFilter) CreateFilterRunnable() *QVideoFilterRunnable /*777 QVideoFilterRunnable **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN20QAbstractVideoFilter20createFilterRunnableEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQVideoFilterRunnableFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtMultimedia/qabstractvideofilter.h:80
// index:0
// Public Visibility=Default Availability=Available
// [-2] void activeChanged()

/*

 */
func (this *QAbstractVideoFilter) ActiveChanged() {
	rv, err := qtrt.InvokeQtFunc6("_ZN20QAbstractVideoFilter13activeChangedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
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
	if false {
		qtcore.KeepMe()
	}
	if false {
		qtgui.KeepMe()
	}
	if false {
		qtnetwork.KeepMe()
	}
}

//  keep block end
