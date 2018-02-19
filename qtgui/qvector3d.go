package qtgui

// /usr/include/qt/QtGui/qvector3d.h
// #include <qvector3d.h>
// #include <QtGui>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 32
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

type QVector3D struct {
	*qtrt.CObject
}
type QVector3D_ITF interface {
	QVector3D_PTR() *QVector3D
}

func (ptr *QVector3D) QVector3D_PTR() *QVector3D { return ptr }

func (this *QVector3D) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QVector3D) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQVector3DFromPointer(cthis unsafe.Pointer) *QVector3D {
	return &QVector3D{&qtrt.CObject{cthis}}
}
func (*QVector3D) NewFromPointer(cthis unsafe.Pointer) *QVector3D {
	return NewQVector3DFromPointer(cthis)
}

// /usr/include/qt/QtGui/qvector3d.h:60
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void QVector3D()
func NewQVector3D() *QVector3D {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QVector3DC2Ev", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	gothis := NewQVector3DFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQVector3D)
	return gothis
}

// /usr/include/qt/QtGui/qvector3d.h:61
// index:1
// Public inline Visibility=Default Availability=Available
// [-2] void QVector3D(Qt::Initialization)
func NewQVector3D_1(arg0 int) *QVector3D {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QVector3DC2EN2Qt14InitializationE", qtrt.FFI_TYPE_POINTER, arg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQVector3DFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQVector3D)
	return gothis
}

// /usr/include/qt/QtGui/qvector3d.h:62
// index:2
// Public inline Visibility=Default Availability=Available
// [-2] void QVector3D(float, float, float)
func NewQVector3D_2(xpos float32, ypos float32, zpos float32) *QVector3D {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QVector3DC2Efff", qtrt.FFI_TYPE_POINTER, xpos, ypos, zpos)
	qtrt.ErrPrint(err, rv)
	gothis := NewQVector3DFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQVector3D)
	return gothis
}

// /usr/include/qt/QtGui/qvector3d.h:64
// index:3
// Public inline Visibility=Default Availability=Available
// [-2] void QVector3D(const QPoint &)
func NewQVector3D_3(point qtcore.QPoint_ITF) *QVector3D {
	var convArg0 unsafe.Pointer
	if point != nil && point.QPoint_PTR() != nil {
		convArg0 = point.QPoint_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QVector3DC2ERK6QPoint", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQVector3DFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQVector3D)
	return gothis
}

// /usr/include/qt/QtGui/qvector3d.h:65
// index:4
// Public inline Visibility=Default Availability=Available
// [-2] void QVector3D(const QPointF &)
func NewQVector3D_4(point qtcore.QPointF_ITF) *QVector3D {
	var convArg0 unsafe.Pointer
	if point != nil && point.QPointF_PTR() != nil {
		convArg0 = point.QPointF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QVector3DC2ERK7QPointF", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQVector3DFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQVector3D)
	return gothis
}

// /usr/include/qt/QtGui/qvector3d.h:67
// index:5
// Public Visibility=Default Availability=Available
// [-2] void QVector3D(const QVector2D &)
func NewQVector3D_5(vector QVector2D_ITF) *QVector3D {
	var convArg0 unsafe.Pointer
	if vector != nil && vector.QVector2D_PTR() != nil {
		convArg0 = vector.QVector2D_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QVector3DC2ERK9QVector2D", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQVector3DFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQVector3D)
	return gothis
}

// /usr/include/qt/QtGui/qvector3d.h:68
// index:6
// Public Visibility=Default Availability=Available
// [-2] void QVector3D(const QVector2D &, float)
func NewQVector3D_6(vector QVector2D_ITF, zpos float32) *QVector3D {
	var convArg0 unsafe.Pointer
	if vector != nil && vector.QVector2D_PTR() != nil {
		convArg0 = vector.QVector2D_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QVector3DC2ERK9QVector2Df", qtrt.FFI_TYPE_POINTER, convArg0, zpos)
	qtrt.ErrPrint(err, rv)
	gothis := NewQVector3DFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQVector3D)
	return gothis
}

// /usr/include/qt/QtGui/qvector3d.h:71
// index:7
// Public Visibility=Default Availability=Available
// [-2] void QVector3D(const QVector4D &)
func NewQVector3D_7(vector QVector4D_ITF) *QVector3D {
	var convArg0 unsafe.Pointer
	if vector != nil && vector.QVector4D_PTR() != nil {
		convArg0 = vector.QVector4D_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QVector3DC2ERK9QVector4D", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQVector3DFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQVector3D)
	return gothis
}

// /usr/include/qt/QtGui/qvector3d.h:74
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isNull() const
func (this *QVector3D) IsNull() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QVector3D6isNullEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qvector3d.h:76
// index:0
// Public inline Visibility=Default Availability=Available
// [4] float x() const
func (this *QVector3D) X() float32 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QVector3D1xEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float32", rv).(float32) // 1111
}

// /usr/include/qt/QtGui/qvector3d.h:77
// index:0
// Public inline Visibility=Default Availability=Available
// [4] float y() const
func (this *QVector3D) Y() float32 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QVector3D1yEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float32", rv).(float32) // 1111
}

// /usr/include/qt/QtGui/qvector3d.h:78
// index:0
// Public inline Visibility=Default Availability=Available
// [4] float z() const
func (this *QVector3D) Z() float32 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QVector3D1zEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float32", rv).(float32) // 1111
}

// /usr/include/qt/QtGui/qvector3d.h:80
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setX(float)
func (this *QVector3D) SetX(x float32) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QVector3D4setXEf", qtrt.FFI_TYPE_POINTER, this.GetCthis(), x)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qvector3d.h:81
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setY(float)
func (this *QVector3D) SetY(y float32) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QVector3D4setYEf", qtrt.FFI_TYPE_POINTER, this.GetCthis(), y)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qvector3d.h:82
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setZ(float)
func (this *QVector3D) SetZ(z float32) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QVector3D4setZEf", qtrt.FFI_TYPE_POINTER, this.GetCthis(), z)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qvector3d.h:84
// index:0
// Public Visibility=Default Availability=Available
// [4] float & operator[](int)
func (this *QVector3D) Operator_get_index(i int) float32 {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QVector3DixEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), i)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cpretval2go("float32", rv).(float32) // 3331
}

// /usr/include/qt/QtGui/qvector3d.h:85
// index:1
// Public Visibility=Default Availability=Available
// [4] float operator[](int) const
func (this *QVector3D) Operator_get_index_1(i int) float32 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QVector3DixEi", qtrt.FFI_TYPE_DOUBLE, this.GetCthis(), i)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float32", rv).(float32) // 1111
}

// /usr/include/qt/QtGui/qvector3d.h:87
// index:0
// Public Visibility=Default Availability=Available
// [4] float length() const
func (this *QVector3D) Length() float32 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QVector3D6lengthEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float32", rv).(float32) // 1111
}

// /usr/include/qt/QtGui/qvector3d.h:88
// index:0
// Public Visibility=Default Availability=Available
// [4] float lengthSquared() const
func (this *QVector3D) LengthSquared() float32 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QVector3D13lengthSquaredEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float32", rv).(float32) // 1111
}

// /usr/include/qt/QtGui/qvector3d.h:90
// index:0
// Public Visibility=Default Availability=Available
// [12] QVector3D normalized() const
func (this *QVector3D) Normalized() *QVector3D /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QVector3D10normalizedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQVector3DFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQVector3D)
	return rv2
}

// /usr/include/qt/QtGui/qvector3d.h:91
// index:0
// Public Visibility=Default Availability=Available
// [-2] void normalize()
func (this *QVector3D) Normalize() {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QVector3D9normalizeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qvector3d.h:93
// index:0
// Public Visibility=Default Availability=Available
// [12] QVector3D & operator+=(const QVector3D &)
func (this *QVector3D) Operator_add_equal(vector QVector3D_ITF) *QVector3D {
	var convArg0 unsafe.Pointer
	if vector != nil && vector.QVector3D_PTR() != nil {
		convArg0 = vector.QVector3D_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QVector3DpLERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQVector3DFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQVector3D)
	return rv2
}

// /usr/include/qt/QtGui/qvector3d.h:94
// index:0
// Public Visibility=Default Availability=Available
// [12] QVector3D & operator-=(const QVector3D &)
func (this *QVector3D) Operator_minus_equal(vector QVector3D_ITF) *QVector3D {
	var convArg0 unsafe.Pointer
	if vector != nil && vector.QVector3D_PTR() != nil {
		convArg0 = vector.QVector3D_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QVector3DmIERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQVector3DFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQVector3D)
	return rv2
}

// /usr/include/qt/QtGui/qvector3d.h:95
// index:0
// Public Visibility=Default Availability=Available
// [12] QVector3D & operator*=(float)
func (this *QVector3D) Operator_mul_equal(factor float32) *QVector3D {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QVector3DmLEf", qtrt.FFI_TYPE_POINTER, this.GetCthis(), factor)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQVector3DFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQVector3D)
	return rv2
}

// /usr/include/qt/QtGui/qvector3d.h:96
// index:1
// Public Visibility=Default Availability=Available
// [12] QVector3D & operator*=(const QVector3D &)
func (this *QVector3D) Operator_mul_equal_1(vector QVector3D_ITF) *QVector3D {
	var convArg0 unsafe.Pointer
	if vector != nil && vector.QVector3D_PTR() != nil {
		convArg0 = vector.QVector3D_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QVector3DmLERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQVector3DFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQVector3D)
	return rv2
}

// /usr/include/qt/QtGui/qvector3d.h:97
// index:0
// Public Visibility=Default Availability=Available
// [12] QVector3D & operator/=(float)
func (this *QVector3D) Operator_div_equal(divisor float32) *QVector3D {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QVector3DdVEf", qtrt.FFI_TYPE_POINTER, this.GetCthis(), divisor)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQVector3DFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQVector3D)
	return rv2
}

// /usr/include/qt/QtGui/qvector3d.h:98
// index:1
// Public inline Visibility=Default Availability=Available
// [12] QVector3D & operator/=(const QVector3D &)
func (this *QVector3D) Operator_div_equal_1(vector QVector3D_ITF) *QVector3D {
	var convArg0 unsafe.Pointer
	if vector != nil && vector.QVector3D_PTR() != nil {
		convArg0 = vector.QVector3D_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QVector3DdVERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQVector3DFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQVector3D)
	return rv2
}

// /usr/include/qt/QtGui/qvector3d.h:100
// index:0
// Public static Visibility=Default Availability=Available
// [4] float dotProduct(const QVector3D &, const QVector3D &)
func (this *QVector3D) DotProduct(v1 QVector3D_ITF, v2 QVector3D_ITF) float32 {
	var convArg0 unsafe.Pointer
	if v1 != nil && v1.QVector3D_PTR() != nil {
		convArg0 = v1.QVector3D_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if v2 != nil && v2.QVector3D_PTR() != nil {
		convArg1 = v2.QVector3D_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QVector3D10dotProductERKS_S1_", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float32", rv).(float32) // 1111
}
func QVector3D_DotProduct(v1 QVector3D_ITF, v2 QVector3D_ITF) float32 {
	var nilthis *QVector3D
	rv := nilthis.DotProduct(v1, v2)
	return rv
}

// /usr/include/qt/QtGui/qvector3d.h:101
// index:0
// Public static Visibility=Default Availability=Available
// [12] QVector3D crossProduct(const QVector3D &, const QVector3D &)
func (this *QVector3D) CrossProduct(v1 QVector3D_ITF, v2 QVector3D_ITF) *QVector3D /*123*/ {
	var convArg0 unsafe.Pointer
	if v1 != nil && v1.QVector3D_PTR() != nil {
		convArg0 = v1.QVector3D_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if v2 != nil && v2.QVector3D_PTR() != nil {
		convArg1 = v2.QVector3D_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QVector3D12crossProductERKS_S1_", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQVector3DFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQVector3D)
	return rv2
}
func QVector3D_CrossProduct(v1 QVector3D_ITF, v2 QVector3D_ITF) *QVector3D /*123*/ {
	var nilthis *QVector3D
	rv := nilthis.CrossProduct(v1, v2)
	return rv
}

// /usr/include/qt/QtGui/qvector3d.h:103
// index:0
// Public static Visibility=Default Availability=Available
// [12] QVector3D normal(const QVector3D &, const QVector3D &)
func (this *QVector3D) Normal(v1 QVector3D_ITF, v2 QVector3D_ITF) *QVector3D /*123*/ {
	var convArg0 unsafe.Pointer
	if v1 != nil && v1.QVector3D_PTR() != nil {
		convArg0 = v1.QVector3D_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if v2 != nil && v2.QVector3D_PTR() != nil {
		convArg1 = v2.QVector3D_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QVector3D6normalERKS_S1_", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQVector3DFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQVector3D)
	return rv2
}
func QVector3D_Normal(v1 QVector3D_ITF, v2 QVector3D_ITF) *QVector3D /*123*/ {
	var nilthis *QVector3D
	rv := nilthis.Normal(v1, v2)
	return rv
}

// /usr/include/qt/QtGui/qvector3d.h:104
// index:1
// Public static Visibility=Default Availability=Available
// [12] QVector3D normal(const QVector3D &, const QVector3D &, const QVector3D &)
func (this *QVector3D) Normal_1(v1 QVector3D_ITF, v2 QVector3D_ITF, v3 QVector3D_ITF) *QVector3D /*123*/ {
	var convArg0 unsafe.Pointer
	if v1 != nil && v1.QVector3D_PTR() != nil {
		convArg0 = v1.QVector3D_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if v2 != nil && v2.QVector3D_PTR() != nil {
		convArg1 = v2.QVector3D_PTR().GetCthis()
	}
	var convArg2 unsafe.Pointer
	if v3 != nil && v3.QVector3D_PTR() != nil {
		convArg2 = v3.QVector3D_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QVector3D6normalERKS_S1_S1_", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, convArg2)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQVector3DFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQVector3D)
	return rv2
}
func QVector3D_Normal_1(v1 QVector3D_ITF, v2 QVector3D_ITF, v3 QVector3D_ITF) *QVector3D /*123*/ {
	var nilthis *QVector3D
	rv := nilthis.Normal_1(v1, v2, v3)
	return rv
}

// /usr/include/qt/QtGui/qvector3d.h:107
// index:0
// Public Visibility=Default Availability=Available
// [12] QVector3D project(const QMatrix4x4 &, const QMatrix4x4 &, const QRect &) const
func (this *QVector3D) Project(modelView QMatrix4x4_ITF, projection QMatrix4x4_ITF, viewport qtcore.QRect_ITF) *QVector3D /*123*/ {
	var convArg0 unsafe.Pointer
	if modelView != nil && modelView.QMatrix4x4_PTR() != nil {
		convArg0 = modelView.QMatrix4x4_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if projection != nil && projection.QMatrix4x4_PTR() != nil {
		convArg1 = projection.QMatrix4x4_PTR().GetCthis()
	}
	var convArg2 unsafe.Pointer
	if viewport != nil && viewport.QRect_PTR() != nil {
		convArg2 = viewport.QRect_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QVector3D7projectERK10QMatrix4x4S2_RK5QRect", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQVector3DFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQVector3D)
	return rv2
}

// /usr/include/qt/QtGui/qvector3d.h:108
// index:0
// Public Visibility=Default Availability=Available
// [12] QVector3D unproject(const QMatrix4x4 &, const QMatrix4x4 &, const QRect &) const
func (this *QVector3D) Unproject(modelView QMatrix4x4_ITF, projection QMatrix4x4_ITF, viewport qtcore.QRect_ITF) *QVector3D /*123*/ {
	var convArg0 unsafe.Pointer
	if modelView != nil && modelView.QMatrix4x4_PTR() != nil {
		convArg0 = modelView.QMatrix4x4_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if projection != nil && projection.QMatrix4x4_PTR() != nil {
		convArg1 = projection.QMatrix4x4_PTR().GetCthis()
	}
	var convArg2 unsafe.Pointer
	if viewport != nil && viewport.QRect_PTR() != nil {
		convArg2 = viewport.QRect_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QVector3D9unprojectERK10QMatrix4x4S2_RK5QRect", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQVector3DFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQVector3D)
	return rv2
}

// /usr/include/qt/QtGui/qvector3d.h:110
// index:0
// Public Visibility=Default Availability=Available
// [4] float distanceToPoint(const QVector3D &) const
func (this *QVector3D) DistanceToPoint(point QVector3D_ITF) float32 {
	var convArg0 unsafe.Pointer
	if point != nil && point.QVector3D_PTR() != nil {
		convArg0 = point.QVector3D_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QVector3D15distanceToPointERKS_", qtrt.FFI_TYPE_DOUBLE, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float32", rv).(float32) // 1111
}

// /usr/include/qt/QtGui/qvector3d.h:111
// index:0
// Public Visibility=Default Availability=Available
// [4] float distanceToPlane(const QVector3D &, const QVector3D &) const
func (this *QVector3D) DistanceToPlane(plane QVector3D_ITF, normal QVector3D_ITF) float32 {
	var convArg0 unsafe.Pointer
	if plane != nil && plane.QVector3D_PTR() != nil {
		convArg0 = plane.QVector3D_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if normal != nil && normal.QVector3D_PTR() != nil {
		convArg1 = normal.QVector3D_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QVector3D15distanceToPlaneERKS_S1_", qtrt.FFI_TYPE_DOUBLE, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float32", rv).(float32) // 1111
}

// /usr/include/qt/QtGui/qvector3d.h:112
// index:1
// Public Visibility=Default Availability=Available
// [4] float distanceToPlane(const QVector3D &, const QVector3D &, const QVector3D &) const
func (this *QVector3D) DistanceToPlane_1(plane1 QVector3D_ITF, plane2 QVector3D_ITF, plane3 QVector3D_ITF) float32 {
	var convArg0 unsafe.Pointer
	if plane1 != nil && plane1.QVector3D_PTR() != nil {
		convArg0 = plane1.QVector3D_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if plane2 != nil && plane2.QVector3D_PTR() != nil {
		convArg1 = plane2.QVector3D_PTR().GetCthis()
	}
	var convArg2 unsafe.Pointer
	if plane3 != nil && plane3.QVector3D_PTR() != nil {
		convArg2 = plane3.QVector3D_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QVector3D15distanceToPlaneERKS_S1_S1_", qtrt.FFI_TYPE_DOUBLE, this.GetCthis(), convArg0, convArg1, convArg2)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float32", rv).(float32) // 1111
}

// /usr/include/qt/QtGui/qvector3d.h:113
// index:0
// Public Visibility=Default Availability=Available
// [4] float distanceToLine(const QVector3D &, const QVector3D &) const
func (this *QVector3D) DistanceToLine(point QVector3D_ITF, direction QVector3D_ITF) float32 {
	var convArg0 unsafe.Pointer
	if point != nil && point.QVector3D_PTR() != nil {
		convArg0 = point.QVector3D_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if direction != nil && direction.QVector3D_PTR() != nil {
		convArg1 = direction.QVector3D_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QVector3D14distanceToLineERKS_S1_", qtrt.FFI_TYPE_DOUBLE, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float32", rv).(float32) // 1111
}

// /usr/include/qt/QtGui/qvector3d.h:129
// index:0
// Public Visibility=Default Availability=Available
// [8] QVector2D toVector2D() const
func (this *QVector3D) ToVector2D() *QVector2D /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QVector3D10toVector2DEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQVector2DFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQVector2D)
	return rv2
}

// /usr/include/qt/QtGui/qvector3d.h:132
// index:0
// Public Visibility=Default Availability=Available
// [16] QVector4D toVector4D() const
func (this *QVector3D) ToVector4D() *QVector4D /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QVector3D10toVector4DEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQVector4DFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQVector4D)
	return rv2
}

// /usr/include/qt/QtGui/qvector3d.h:135
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QPoint toPoint() const
func (this *QVector3D) ToPoint() *qtcore.QPoint /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QVector3D7toPointEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQPointFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQPoint)
	return rv2
}

// /usr/include/qt/QtGui/qvector3d.h:136
// index:0
// Public inline Visibility=Default Availability=Available
// [16] QPointF toPointF() const
func (this *QVector3D) ToPointF() *qtcore.QPointF /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QVector3D8toPointFEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQPointFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQPointF)
	return rv2
}

func DeleteQVector3D(this *QVector3D) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QVector3DD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
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
