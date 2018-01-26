package qtquick

// /usr/include/qt/QtQuick/qsgninepatchnode.h
// #include <qsgninepatchnode.h>
// #include <QtQuick>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 16
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "mkuse/cffiqt"
import "gopp"
import "qt.go/qtrt"
import "qt.go/qtcore"
import "qt.go/qtgui"
import "qt.go/qtnetwork"
import "qt.go/qtqml"

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
	if false {
		qtcore.KeepMe()
	}
	if false {
		qtgui.KeepMe()
	}
	if false {
		qtnetwork.KeepMe()
	}
	if false {
		qtqml.KeepMe()
	}
}

//  ext block end

//  body block begin
type QSGNinePatchNode struct {
	*QSGGeometryNode
}

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
// Public inline virtual
// void ~QSGNinePatchNode()
func DeleteQSGNinePatchNode(*QSGNinePatchNode) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QSGNinePatchNodeD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qsgninepatchnode.h:53
// index:0
// Public pure virtual
// void setTexture(class QSGTexture *)
func (this *QSGNinePatchNode) SetTexture(texture *QSGTexture /*777 QSGTexture **/) {
	var convArg0 = texture.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QSGNinePatchNode10setTextureEP10QSGTexture", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qsgninepatchnode.h:54
// index:0
// Public pure virtual
// void setBounds(const class QRectF &)
func (this *QSGNinePatchNode) SetBounds(bounds *qtcore.QRectF) {
	var convArg0 = bounds.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QSGNinePatchNode9setBoundsERK6QRectF", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qsgninepatchnode.h:55
// index:0
// Public pure virtual
// void setDevicePixelRatio(qreal)
func (this *QSGNinePatchNode) SetDevicePixelRatio(ratio float64) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QSGNinePatchNode19setDevicePixelRatioEd", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), ratio)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qsgninepatchnode.h:56
// index:0
// Public pure virtual
// void setPadding(qreal, qreal, qreal, qreal)
func (this *QSGNinePatchNode) SetPadding(left float64, top float64, right float64, bottom float64) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QSGNinePatchNode10setPaddingEdddd", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), left, top, right, bottom)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qsgninepatchnode.h:57
// index:0
// Public pure virtual
// void update()
func (this *QSGNinePatchNode) Update() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QSGNinePatchNode6updateEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qsgninepatchnode.h:59
// index:0
// Public static
// void rebuildGeometry(class QSGTexture *, class QSGGeometry *, const class QVector4D &, const class QRectF &, qreal)
func (this *QSGNinePatchNode) RebuildGeometry(texture *QSGTexture /*777 QSGTexture **/, geometry *QSGGeometry /*777 QSGGeometry **/, padding *qtgui.QVector4D, bounds *qtcore.QRectF, dpr float64) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QSGNinePatchNode15rebuildGeometryEP10QSGTextureP11QSGGeometryRK9QVector4DRK6QRectFd", ffiqt.FFI_TYPE_POINTER, texture, geometry, padding, bounds, dpr)
	gopp.ErrPrint(err, rv)
}
func QSGNinePatchNode_RebuildGeometry(texture *QSGTexture /*777 QSGTexture **/, geometry *QSGGeometry /*777 QSGGeometry **/, padding *qtgui.QVector4D, bounds *qtcore.QRectF, dpr float64) {
	var nilthis *QSGNinePatchNode
	nilthis.RebuildGeometry(texture, geometry, padding, bounds, dpr)
}

//  body block end
