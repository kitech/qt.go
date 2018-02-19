package qtgui

// /usr/include/qt/QtGui/qpolygon.h
// #include <qpolygon.h>
// #include <QtGui>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 25
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "log"
import "github.com/kitech/qt.go/qtrt"
import "github.com/kitech/qt.go/qtcore"

//  ext block end

//  body block begin

type QPolygonF struct {
	*qtrt.CObject
}
type QPolygonF_ITF interface {
	QPolygonF_PTR() *QPolygonF
}

func (ptr *QPolygonF) QPolygonF_PTR() *QPolygonF { return ptr }

func (this *QPolygonF) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QPolygonF) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQPolygonFFromPointer(cthis unsafe.Pointer) *QPolygonF {
	return &QPolygonF{&qtrt.CObject{cthis}}
}
func (*QPolygonF) NewFromPointer(cthis unsafe.Pointer) *QPolygonF {
	return NewQPolygonFFromPointer(cthis)
}

// /usr/include/qt/QtGui/qpolygon.h:144
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void QPolygonF()
func NewQPolygonF() *QPolygonF {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QPolygonFC2Ev", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	gothis := NewQPolygonFFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQPolygonF)
	return gothis
}

// /usr/include/qt/QtGui/qpolygon.h:146
// index:1
// Public inline Visibility=Default Availability=Available
// [-2] void QPolygonF(int)
func NewQPolygonF_1(size int) *QPolygonF {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QPolygonFC2Ei", qtrt.FFI_TYPE_POINTER, size)
	qtrt.ErrPrint(err, rv)
	gothis := NewQPolygonFFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQPolygonF)
	return gothis
}

// /usr/include/qt/QtGui/qpolygon.h:151
// index:2
// Public Visibility=Default Availability=Available
// [-2] void QPolygonF(const QRectF &)
func NewQPolygonF_2(r qtcore.QRectF_ITF) *QPolygonF {
	var convArg0 unsafe.Pointer
	if r != nil && r.QRectF_PTR() != nil {
		convArg0 = r.QRectF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QPolygonFC2ERK6QRectF", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQPolygonFFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQPolygonF)
	return gothis
}

// /usr/include/qt/QtGui/qpolygon.h:152
// index:3
// Public Visibility=Default Availability=Available
// [-2] void QPolygonF(const QPolygon &)
func NewQPolygonF_3(a QPolygon_ITF) *QPolygonF {
	var convArg0 unsafe.Pointer
	if a != nil && a.QPolygon_PTR() != nil {
		convArg0 = a.QPolygon_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QPolygonFC2ERK8QPolygon", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQPolygonFFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQPolygonF)
	return gothis
}

// /usr/include/qt/QtGui/qpolygon.h:145
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void ~QPolygonF()
func DeleteQPolygonF(this *QPolygonF) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QPolygonFD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 8)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtGui/qpolygon.h:156
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QPolygonF & operator=(QPolygonF &&)
func (this *QPolygonF) Operator_equal(other unsafe.Pointer /*333*/) *QPolygonF {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QPolygonFaSEOS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), other)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQPolygonFFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQPolygonF)
	return rv2
}

// /usr/include/qt/QtGui/qpolygon.h:158
// index:1
// Public inline Visibility=Default Availability=Available
// [8] QPolygonF & operator=(const QPolygonF &)
func (this *QPolygonF) Operator_equal_1(other QPolygonF_ITF) *QPolygonF {
	var convArg0 unsafe.Pointer
	if other != nil && other.QPolygonF_PTR() != nil {
		convArg0 = other.QPolygonF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QPolygonFaSERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQPolygonFFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQPolygonF)
	return rv2
}

// /usr/include/qt/QtGui/qpolygon.h:159
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void swap(QPolygonF &)
func (this *QPolygonF) Swap(other QPolygonF_ITF) {
	var convArg0 unsafe.Pointer
	if other != nil && other.QPolygonF_PTR() != nil {
		convArg0 = other.QPolygonF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QPolygonF4swapERS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpolygon.h:163
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void translate(qreal, qreal)
func (this *QPolygonF) Translate(dx float64, dy float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QPolygonF9translateEdd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), dx, dy)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpolygon.h:164
// index:1
// Public Visibility=Default Availability=Available
// [-2] void translate(const QPointF &)
func (this *QPolygonF) Translate_1(offset qtcore.QPointF_ITF) {
	var convArg0 unsafe.Pointer
	if offset != nil && offset.QPointF_PTR() != nil {
		convArg0 = offset.QPointF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QPolygonF9translateERK7QPointF", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpolygon.h:166
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QPolygonF translated(qreal, qreal) const
func (this *QPolygonF) Translated(dx float64, dy float64) *QPolygonF /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QPolygonF10translatedEdd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), dx, dy)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQPolygonFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQPolygonF)
	return rv2
}

// /usr/include/qt/QtGui/qpolygon.h:167
// index:1
// Public Visibility=Default Availability=Available
// [8] QPolygonF translated(const QPointF &) const
func (this *QPolygonF) Translated_1(offset qtcore.QPointF_ITF) *QPolygonF /*123*/ {
	var convArg0 unsafe.Pointer
	if offset != nil && offset.QPointF_PTR() != nil {
		convArg0 = offset.QPointF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QPolygonF10translatedERK7QPointF", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQPolygonFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQPolygonF)
	return rv2
}

// /usr/include/qt/QtGui/qpolygon.h:169
// index:0
// Public Visibility=Default Availability=Available
// [8] QPolygon toPolygon() const
func (this *QPolygonF) ToPolygon() *QPolygon /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QPolygonF9toPolygonEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQPolygonFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQPolygon)
	return rv2
}

// /usr/include/qt/QtGui/qpolygon.h:171
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isClosed() const
func (this *QPolygonF) IsClosed() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QPolygonF8isClosedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qpolygon.h:173
// index:0
// Public Visibility=Default Availability=Available
// [32] QRectF boundingRect() const
func (this *QPolygonF) BoundingRect() *qtcore.QRectF /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QPolygonF12boundingRectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQRectFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQRectF)
	return rv2
}

// /usr/include/qt/QtGui/qpolygon.h:175
// index:0
// Public Visibility=Default Availability=Available
// [1] bool containsPoint(const QPointF &, Qt::FillRule) const
func (this *QPolygonF) ContainsPoint(pt qtcore.QPointF_ITF, fillRule int) bool {
	var convArg0 unsafe.Pointer
	if pt != nil && pt.QPointF_PTR() != nil {
		convArg0 = pt.QPointF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QPolygonF13containsPointERK7QPointFN2Qt8FillRuleE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, fillRule)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qpolygon.h:177
// index:0
// Public Visibility=Default Availability=Available
// [8] QPolygonF united(const QPolygonF &) const
func (this *QPolygonF) United(r QPolygonF_ITF) *QPolygonF /*123*/ {
	var convArg0 unsafe.Pointer
	if r != nil && r.QPolygonF_PTR() != nil {
		convArg0 = r.QPolygonF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QPolygonF6unitedERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQPolygonFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQPolygonF)
	return rv2
}

// /usr/include/qt/QtGui/qpolygon.h:178
// index:0
// Public Visibility=Default Availability=Available
// [8] QPolygonF intersected(const QPolygonF &) const
func (this *QPolygonF) Intersected(r QPolygonF_ITF) *QPolygonF /*123*/ {
	var convArg0 unsafe.Pointer
	if r != nil && r.QPolygonF_PTR() != nil {
		convArg0 = r.QPolygonF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QPolygonF11intersectedERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQPolygonFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQPolygonF)
	return rv2
}

// /usr/include/qt/QtGui/qpolygon.h:179
// index:0
// Public Visibility=Default Availability=Available
// [8] QPolygonF subtracted(const QPolygonF &) const
func (this *QPolygonF) Subtracted(r QPolygonF_ITF) *QPolygonF /*123*/ {
	var convArg0 unsafe.Pointer
	if r != nil && r.QPolygonF_PTR() != nil {
		convArg0 = r.QPolygonF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QPolygonF10subtractedERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQPolygonFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQPolygonF)
	return rv2
}

// /usr/include/qt/QtGui/qpolygon.h:181
// index:0
// Public Visibility=Default Availability=Available
// [1] bool intersects(const QPolygonF &) const
func (this *QPolygonF) Intersects(r QPolygonF_ITF) bool {
	var convArg0 unsafe.Pointer
	if r != nil && r.QPolygonF_PTR() != nil {
		convArg0 = r.QPolygonF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QPolygonF10intersectsERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
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
}

//  keep block end
