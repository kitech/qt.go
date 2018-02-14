package qtquick

// /usr/include/qt/QtQuick/qsgninepatchnode.h
// #include <qsgninepatchnode.h>
// #include <QtQuick>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 18
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "log"
import "github.com/kitech/qt.go/qtrt"
import "github.com/kitech/qt.go/qtcore"
import "github.com/kitech/qt.go/qtnetwork"
import "github.com/kitech/qt.go/qtgui"
import "github.com/kitech/qt.go/qtqml"

//  ext block end

//  body block begin

type QSGNinePatchNode struct {
	*QSGGeometryNode
}
type QSGNinePatchNode_ITF interface {
	QSGGeometryNode_ITF
	QSGNinePatchNode_PTR() *QSGNinePatchNode
}

func (ptr *QSGNinePatchNode) QSGNinePatchNode_PTR() *QSGNinePatchNode { return ptr }

func (this *QSGNinePatchNode) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QSGGeometryNode.GetCthis()
	}
}
func (this *QSGNinePatchNode) SetCthis(cthis unsafe.Pointer) {
	this.QSGGeometryNode = NewQSGGeometryNodeFromPointer(cthis)
}
func NewQSGNinePatchNodeFromPointer(cthis unsafe.Pointer) *QSGNinePatchNode {
	bcthis0 := NewQSGGeometryNodeFromPointer(cthis)
	return &QSGNinePatchNode{bcthis0}
}
func (*QSGNinePatchNode) NewFromPointer(cthis unsafe.Pointer) *QSGNinePatchNode {
	return NewQSGNinePatchNodeFromPointer(cthis)
}

// /usr/include/qt/QtQuick/qsgninepatchnode.h:51
// index:0
// Public inline virtual Visibility=Default Availability=Available
// [-2] void ~QSGNinePatchNode()
func DeleteQSGNinePatchNode(this *QSGNinePatchNode) {
	rv, err := qtrt.InvokeQtFunc6("_ZN16QSGNinePatchNodeD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 144)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtQuick/qsgninepatchnode.h:53
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [-2] void setTexture(QSGTexture *)
func (this *QSGNinePatchNode) SetTexture(texture QSGTexture_ITF /*777 QSGTexture **/) {
	var convArg0 unsafe.Pointer
	if texture != nil && texture.QSGTexture_PTR() != nil {
		convArg0 = texture.QSGTexture_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN16QSGNinePatchNode10setTextureEP10QSGTexture", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qsgninepatchnode.h:54
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [-2] void setBounds(const QRectF &)
func (this *QSGNinePatchNode) SetBounds(bounds qtcore.QRectF_ITF) {
	var convArg0 unsafe.Pointer
	if bounds != nil && bounds.QRectF_PTR() != nil {
		convArg0 = bounds.QRectF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN16QSGNinePatchNode9setBoundsERK6QRectF", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qsgninepatchnode.h:55
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [-2] void setDevicePixelRatio(qreal)
func (this *QSGNinePatchNode) SetDevicePixelRatio(ratio float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN16QSGNinePatchNode19setDevicePixelRatioEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), ratio)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qsgninepatchnode.h:56
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [-2] void setPadding(qreal, qreal, qreal, qreal)
func (this *QSGNinePatchNode) SetPadding(left float64, top float64, right float64, bottom float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN16QSGNinePatchNode10setPaddingEdddd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), left, top, right, bottom)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qsgninepatchnode.h:57
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [-2] void update()
func (this *QSGNinePatchNode) Update() {
	rv, err := qtrt.InvokeQtFunc6("_ZN16QSGNinePatchNode6updateEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qsgninepatchnode.h:59
// index:0
// Public static Visibility=Default Availability=Available
// [-2] void rebuildGeometry(QSGTexture *, QSGGeometry *, const QVector4D &, const QRectF &, qreal)
func (this *QSGNinePatchNode) RebuildGeometry(texture QSGTexture_ITF /*777 QSGTexture **/, geometry QSGGeometry_ITF /*777 QSGGeometry **/, padding qtgui.QVector4D_ITF, bounds qtcore.QRectF_ITF, dpr float64) {
	var convArg0 unsafe.Pointer
	if texture != nil && texture.QSGTexture_PTR() != nil {
		convArg0 = texture.QSGTexture_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if geometry != nil && geometry.QSGGeometry_PTR() != nil {
		convArg1 = geometry.QSGGeometry_PTR().GetCthis()
	}
	var convArg2 unsafe.Pointer
	if padding != nil && padding.QVector4D_PTR() != nil {
		convArg2 = padding.QVector4D_PTR().GetCthis()
	}
	var convArg3 unsafe.Pointer
	if bounds != nil && bounds.QRectF_PTR() != nil {
		convArg3 = bounds.QRectF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN16QSGNinePatchNode15rebuildGeometryEP10QSGTextureP11QSGGeometryRK9QVector4DRK6QRectFd", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, convArg2, convArg3, dpr)
	qtrt.ErrPrint(err, rv)
}
func QSGNinePatchNode_RebuildGeometry(texture QSGTexture_ITF /*777 QSGTexture **/, geometry QSGGeometry_ITF /*777 QSGGeometry **/, padding qtgui.QVector4D_ITF, bounds qtcore.QRectF_ITF, dpr float64) {
	var nilthis *QSGNinePatchNode
	nilthis.RebuildGeometry(texture, geometry, padding, bounds, dpr)
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
		qtnetwork.KeepMe()
	}
	if false {
		qtgui.KeepMe()
	}
	if false {
		qtqml.KeepMe()
	}
}

//  keep block end
