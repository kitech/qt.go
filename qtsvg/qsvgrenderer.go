package qtsvg

// /usr/include/qt/QtSvg/qsvgrenderer.h
// #include <qsvgrenderer.h>
// #include <QtSvg>

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
import "github.com/kitech/qt.go/qtcore"
import "github.com/kitech/qt.go/qtgui"
import "github.com/kitech/qt.go/qtwidgets"

//  ext block end

//  body block begin

/*

 */
type QSvgRenderer struct {
	*qtcore.QObject
}
type QSvgRenderer_ITF interface {
	qtcore.QObject_ITF
	QSvgRenderer_PTR() *QSvgRenderer
}

func (ptr *QSvgRenderer) QSvgRenderer_PTR() *QSvgRenderer { return ptr }

func (this *QSvgRenderer) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QObject.GetCthis()
	}
}
func (this *QSvgRenderer) SetCthis(cthis unsafe.Pointer) {
	this.QObject = qtcore.NewQObjectFromPointer(cthis)
}
func NewQSvgRendererFromPointer(cthis unsafe.Pointer) *QSvgRenderer {
	bcthis0 := qtcore.NewQObjectFromPointer(cthis)
	return &QSvgRenderer{bcthis0}
}
func (*QSvgRenderer) NewFromPointer(cthis unsafe.Pointer) *QSvgRenderer {
	return NewQSvgRendererFromPointer(cthis)
}

// /usr/include/qt/QtSvg/qsvgrenderer.h:62
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject() const

/*

 */
func (this *QSvgRenderer) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QSvgRenderer10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtSvg/qsvgrenderer.h:68
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QSvgRenderer(QObject *)

/*
Constructs a new renderer with the given parent.
*/
func (*QSvgRenderer) NewForInherit(parent qtcore.QObject_ITF /*777 QObject **/) *QSvgRenderer {
	return NewQSvgRenderer(parent)
}
func NewQSvgRenderer(parent qtcore.QObject_ITF /*777 QObject **/) *QSvgRenderer {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QObject_PTR() != nil {
		convArg0 = parent.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN12QSvgRendererC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQSvgRendererFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QSvgRenderer")
	return gothis
}

// /usr/include/qt/QtSvg/qsvgrenderer.h:68
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QSvgRenderer(QObject *)

/*
Constructs a new renderer with the given parent.
*/
func (*QSvgRenderer) NewForInheritp() *QSvgRenderer {
	return NewQSvgRendererp()
}
func NewQSvgRendererp() *QSvgRenderer {
	// arg: 0, QObject *=Pointer, QObject=Record, , Invalid
	var convArg0 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN12QSvgRendererC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQSvgRendererFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QSvgRenderer")
	return gothis
}

// /usr/include/qt/QtSvg/qsvgrenderer.h:69
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QSvgRenderer(const QString &, QObject *)

/*
Constructs a new renderer with the given parent.
*/
func (*QSvgRenderer) NewForInherit1(filename string, parent qtcore.QObject_ITF /*777 QObject **/) *QSvgRenderer {
	return NewQSvgRenderer1(filename, parent)
}
func NewQSvgRenderer1(filename string, parent qtcore.QObject_ITF /*777 QObject **/) *QSvgRenderer {
	var tmpArg0 = qtcore.NewQString5(filename)
	var convArg0 = tmpArg0.GetCthis()
	var convArg1 unsafe.Pointer
	if parent != nil && parent.QObject_PTR() != nil {
		convArg1 = parent.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN12QSvgRendererC2ERK7QStringP7QObject", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	gothis := NewQSvgRendererFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QSvgRenderer")
	return gothis
}

// /usr/include/qt/QtSvg/qsvgrenderer.h:69
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QSvgRenderer(const QString &, QObject *)

/*
Constructs a new renderer with the given parent.
*/
func (*QSvgRenderer) NewForInherit1p(filename string) *QSvgRenderer {
	return NewQSvgRenderer1p(filename)
}
func NewQSvgRenderer1p(filename string) *QSvgRenderer {
	var tmpArg0 = qtcore.NewQString5(filename)
	var convArg0 = tmpArg0.GetCthis()
	// arg: 1, QObject *=Pointer, QObject=Record, , Invalid
	var convArg1 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN12QSvgRendererC2ERK7QStringP7QObject", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	gothis := NewQSvgRendererFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QSvgRenderer")
	return gothis
}

// /usr/include/qt/QtSvg/qsvgrenderer.h:70
// index:2
// Public Visibility=Default Availability=Available
// [-2] void QSvgRenderer(const QByteArray &, QObject *)

/*
Constructs a new renderer with the given parent.
*/
func (*QSvgRenderer) NewForInherit2(contents qtcore.QByteArray_ITF, parent qtcore.QObject_ITF /*777 QObject **/) *QSvgRenderer {
	return NewQSvgRenderer2(contents, parent)
}
func NewQSvgRenderer2(contents qtcore.QByteArray_ITF, parent qtcore.QObject_ITF /*777 QObject **/) *QSvgRenderer {
	var convArg0 unsafe.Pointer
	if contents != nil && contents.QByteArray_PTR() != nil {
		convArg0 = contents.QByteArray_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if parent != nil && parent.QObject_PTR() != nil {
		convArg1 = parent.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN12QSvgRendererC2ERK10QByteArrayP7QObject", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	gothis := NewQSvgRendererFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QSvgRenderer")
	return gothis
}

// /usr/include/qt/QtSvg/qsvgrenderer.h:70
// index:2
// Public Visibility=Default Availability=Available
// [-2] void QSvgRenderer(const QByteArray &, QObject *)

/*
Constructs a new renderer with the given parent.
*/
func (*QSvgRenderer) NewForInherit2p(contents qtcore.QByteArray_ITF) *QSvgRenderer {
	return NewQSvgRenderer2p(contents)
}
func NewQSvgRenderer2p(contents qtcore.QByteArray_ITF) *QSvgRenderer {
	var convArg0 unsafe.Pointer
	if contents != nil && contents.QByteArray_PTR() != nil {
		convArg0 = contents.QByteArray_PTR().GetCthis()
	}
	// arg: 1, QObject *=Pointer, QObject=Record, , Invalid
	var convArg1 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN12QSvgRendererC2ERK10QByteArrayP7QObject", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	gothis := NewQSvgRendererFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QSvgRenderer")
	return gothis
}

// /usr/include/qt/QtSvg/qsvgrenderer.h:71
// index:3
// Public Visibility=Default Availability=Available
// [-2] void QSvgRenderer(QXmlStreamReader *, QObject *)

/*
Constructs a new renderer with the given parent.
*/
func (*QSvgRenderer) NewForInherit3(contents qtcore.QXmlStreamReader_ITF /*777 QXmlStreamReader **/, parent qtcore.QObject_ITF /*777 QObject **/) *QSvgRenderer {
	return NewQSvgRenderer3(contents, parent)
}
func NewQSvgRenderer3(contents qtcore.QXmlStreamReader_ITF /*777 QXmlStreamReader **/, parent qtcore.QObject_ITF /*777 QObject **/) *QSvgRenderer {
	var convArg0 unsafe.Pointer
	if contents != nil && contents.QXmlStreamReader_PTR() != nil {
		convArg0 = contents.QXmlStreamReader_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if parent != nil && parent.QObject_PTR() != nil {
		convArg1 = parent.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN12QSvgRendererC2EP16QXmlStreamReaderP7QObject", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	gothis := NewQSvgRendererFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QSvgRenderer")
	return gothis
}

// /usr/include/qt/QtSvg/qsvgrenderer.h:71
// index:3
// Public Visibility=Default Availability=Available
// [-2] void QSvgRenderer(QXmlStreamReader *, QObject *)

/*
Constructs a new renderer with the given parent.
*/
func (*QSvgRenderer) NewForInherit3p(contents qtcore.QXmlStreamReader_ITF /*777 QXmlStreamReader **/) *QSvgRenderer {
	return NewQSvgRenderer3p(contents)
}
func NewQSvgRenderer3p(contents qtcore.QXmlStreamReader_ITF /*777 QXmlStreamReader **/) *QSvgRenderer {
	var convArg0 unsafe.Pointer
	if contents != nil && contents.QXmlStreamReader_PTR() != nil {
		convArg0 = contents.QXmlStreamReader_PTR().GetCthis()
	}
	// arg: 1, QObject *=Pointer, QObject=Record, , Invalid
	var convArg1 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN12QSvgRendererC2EP16QXmlStreamReaderP7QObject", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	gothis := NewQSvgRendererFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QSvgRenderer")
	return gothis
}

// /usr/include/qt/QtSvg/qsvgrenderer.h:72
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QSvgRenderer()

/*

 */
func DeleteQSvgRenderer(this *QSvgRenderer) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QSvgRendererD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 16)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtSvg/qsvgrenderer.h:74
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isValid() const

/*
Returns true if there is a valid current document; otherwise returns false.
*/
func (this *QSvgRenderer) IsValid() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QSvgRenderer7isValidEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtSvg/qsvgrenderer.h:76
// index:0
// Public Visibility=Default Availability=Available
// [8] QSize defaultSize() const

/*
Returns the default size of the document contents.
*/
func (this *QSvgRenderer) DefaultSize() *qtcore.QSize /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QSvgRenderer11defaultSizeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSize)
	return rv2
}

// /usr/include/qt/QtSvg/qsvgrenderer.h:78
// index:0
// Public Visibility=Default Availability=Available
// [16] QRect viewBox() const

/*
Returns viewBoxF().toRect().

See also setViewBox() and viewBoxF().
*/
func (this *QSvgRenderer) ViewBox() *qtcore.QRect /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QSvgRenderer7viewBoxEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQRectFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQRect)
	return rv2
}

// /usr/include/qt/QtSvg/qsvgrenderer.h:79
// index:0
// Public Visibility=Default Availability=Available
// [32] QRectF viewBoxF() const

/*

 */
func (this *QSvgRenderer) ViewBoxF() *qtcore.QRectF /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QSvgRenderer8viewBoxFEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQRectFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQRectF)
	return rv2
}

// /usr/include/qt/QtSvg/qsvgrenderer.h:80
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setViewBox(const QRect &)

/*

 */
func (this *QSvgRenderer) SetViewBox(viewbox qtcore.QRect_ITF) {
	var convArg0 unsafe.Pointer
	if viewbox != nil && viewbox.QRect_PTR() != nil {
		convArg0 = viewbox.QRect_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN12QSvgRenderer10setViewBoxERK5QRect", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtSvg/qsvgrenderer.h:81
// index:1
// Public Visibility=Default Availability=Available
// [-2] void setViewBox(const QRectF &)

/*

 */
func (this *QSvgRenderer) SetViewBox1(viewbox qtcore.QRectF_ITF) {
	var convArg0 unsafe.Pointer
	if viewbox != nil && viewbox.QRectF_PTR() != nil {
		convArg0 = viewbox.QRectF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN12QSvgRenderer10setViewBoxERK6QRectF", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtSvg/qsvgrenderer.h:83
// index:0
// Public Visibility=Default Availability=Available
// [1] bool animated() const

/*
Returns true if the current document contains animated elements; otherwise returns false.

See also framesPerSecond().
*/
func (this *QSvgRenderer) Animated() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QSvgRenderer8animatedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtSvg/qsvgrenderer.h:84
// index:0
// Public Visibility=Default Availability=Available
// [4] int framesPerSecond() const

/*

 */
func (this *QSvgRenderer) FramesPerSecond() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QSvgRenderer15framesPerSecondEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtSvg/qsvgrenderer.h:85
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setFramesPerSecond(int)

/*

 */
func (this *QSvgRenderer) SetFramesPerSecond(num int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QSvgRenderer18setFramesPerSecondEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), num)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtSvg/qsvgrenderer.h:86
// index:0
// Public Visibility=Default Availability=Available
// [4] int currentFrame() const

/*

 */
func (this *QSvgRenderer) CurrentFrame() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QSvgRenderer12currentFrameEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtSvg/qsvgrenderer.h:87
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setCurrentFrame(int)

/*

 */
func (this *QSvgRenderer) SetCurrentFrame(arg0 int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QSvgRenderer15setCurrentFrameEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtSvg/qsvgrenderer.h:88
// index:0
// Public Visibility=Default Availability=Available
// [4] int animationDuration() const

/*

 */
func (this *QSvgRenderer) AnimationDuration() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QSvgRenderer17animationDurationEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtSvg/qsvgrenderer.h:90
// index:0
// Public Visibility=Default Availability=Available
// [32] QRectF boundsOnElement(const QString &) const

/*
Returns bounding rectangle of the item with the given id. The transformation matrix of parent elements is not affecting the bounds of the element.

This function was introduced in  Qt 4.2.

See also matrixForElement().
*/
func (this *QSvgRenderer) BoundsOnElement(id string) *qtcore.QRectF /*123*/ {
	var tmpArg0 = qtcore.NewQString5(id)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QSvgRenderer15boundsOnElementERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQRectFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQRectF)
	return rv2
}

// /usr/include/qt/QtSvg/qsvgrenderer.h:91
// index:0
// Public Visibility=Default Availability=Available
// [1] bool elementExists(const QString &) const

/*
Returns true if the element with the given id exists in the currently parsed SVG file and is a renderable element.

Note: this method returns true only for elements that can be rendered. Which implies that elements that are considered part of the fill/stroke style properties, e.g. radialGradients even tough marked with "id" attributes will not be found by this method.

This function was introduced in  Qt 4.2.
*/
func (this *QSvgRenderer) ElementExists(id string) bool {
	var tmpArg0 = qtcore.NewQString5(id)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QSvgRenderer13elementExistsERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtSvg/qsvgrenderer.h:92
// index:0
// Public Visibility=Default Availability=Available
// [48] QMatrix matrixForElement(const QString &) const

/*
Returns the transformation matrix for the element with the given id. The matrix is a product of the transformation of the element's parents. The transformation of the element itself is not included.

To find the bounding rectangle of the element in logical coordinates, you can apply the matrix on the rectangle returned from boundsOnElement().

This function was introduced in  Qt 4.2.

See also boundsOnElement().
*/
func (this *QSvgRenderer) MatrixForElement(id string) *qtgui.QMatrix /*123*/ {
	var tmpArg0 = qtcore.NewQString5(id)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QSvgRenderer16matrixForElementERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := qtgui.NewQMatrixFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtgui.DeleteQMatrix)
	return rv2
}

// /usr/include/qt/QtSvg/qsvgrenderer.h:95
// index:0
// Public Visibility=Default Availability=Available
// [1] bool load(const QString &)

/*
Loads the SVG file specified by filename, returning true if the content was successfully parsed; otherwise returns false.
*/
func (this *QSvgRenderer) Load(filename string) bool {
	var tmpArg0 = qtcore.NewQString5(filename)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN12QSvgRenderer4loadERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtSvg/qsvgrenderer.h:96
// index:1
// Public Visibility=Default Availability=Available
// [1] bool load(const QByteArray &)

/*
Loads the SVG file specified by filename, returning true if the content was successfully parsed; otherwise returns false.
*/
func (this *QSvgRenderer) Load1(contents qtcore.QByteArray_ITF) bool {
	var convArg0 unsafe.Pointer
	if contents != nil && contents.QByteArray_PTR() != nil {
		convArg0 = contents.QByteArray_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN12QSvgRenderer4loadERK10QByteArray", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtSvg/qsvgrenderer.h:97
// index:2
// Public Visibility=Default Availability=Available
// [1] bool load(QXmlStreamReader *)

/*
Loads the SVG file specified by filename, returning true if the content was successfully parsed; otherwise returns false.
*/
func (this *QSvgRenderer) Load2(contents qtcore.QXmlStreamReader_ITF /*777 QXmlStreamReader **/) bool {
	var convArg0 unsafe.Pointer
	if contents != nil && contents.QXmlStreamReader_PTR() != nil {
		convArg0 = contents.QXmlStreamReader_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN12QSvgRenderer4loadEP16QXmlStreamReader", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtSvg/qsvgrenderer.h:98
// index:0
// Public Visibility=Default Availability=Available
// [-2] void render(QPainter *)

/*
Renders the current document, or the current frame of an animated document, using the given painter.
*/
func (this *QSvgRenderer) Render(p qtgui.QPainter_ITF /*777 QPainter **/) {
	var convArg0 unsafe.Pointer
	if p != nil && p.QPainter_PTR() != nil {
		convArg0 = p.QPainter_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN12QSvgRenderer6renderEP8QPainter", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtSvg/qsvgrenderer.h:99
// index:1
// Public Visibility=Default Availability=Available
// [-2] void render(QPainter *, const QRectF &)

/*
Renders the current document, or the current frame of an animated document, using the given painter.
*/
func (this *QSvgRenderer) Render1(p qtgui.QPainter_ITF /*777 QPainter **/, bounds qtcore.QRectF_ITF) {
	var convArg0 unsafe.Pointer
	if p != nil && p.QPainter_PTR() != nil {
		convArg0 = p.QPainter_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if bounds != nil && bounds.QRectF_PTR() != nil {
		convArg1 = bounds.QRectF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN12QSvgRenderer6renderEP8QPainterRK6QRectF", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtSvg/qsvgrenderer.h:101
// index:2
// Public Visibility=Default Availability=Available
// [-2] void render(QPainter *, const QString &, const QRectF &)

/*
Renders the current document, or the current frame of an animated document, using the given painter.
*/
func (this *QSvgRenderer) Render2(p qtgui.QPainter_ITF /*777 QPainter **/, elementId string, bounds qtcore.QRectF_ITF) {
	var convArg0 unsafe.Pointer
	if p != nil && p.QPainter_PTR() != nil {
		convArg0 = p.QPainter_PTR().GetCthis()
	}
	var tmpArg1 = qtcore.NewQString5(elementId)
	var convArg1 = tmpArg1.GetCthis()
	var convArg2 unsafe.Pointer
	if bounds != nil && bounds.QRectF_PTR() != nil {
		convArg2 = bounds.QRectF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN12QSvgRenderer6renderEP8QPainterRK7QStringRK6QRectF", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtSvg/qsvgrenderer.h:101
// index:2
// Public Visibility=Default Availability=Available
// [-2] void render(QPainter *, const QString &, const QRectF &)

/*
Renders the current document, or the current frame of an animated document, using the given painter.
*/
func (this *QSvgRenderer) Render2p(p qtgui.QPainter_ITF /*777 QPainter **/, elementId string) {
	var convArg0 unsafe.Pointer
	if p != nil && p.QPainter_PTR() != nil {
		convArg0 = p.QPainter_PTR().GetCthis()
	}
	var tmpArg1 = qtcore.NewQString5(elementId)
	var convArg1 = tmpArg1.GetCthis()
	// arg: 2, const QRectF &=LValueReference, QRectF=Record, , Invalid
	var convArg2 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN12QSvgRenderer6renderEP8QPainterRK7QStringRK6QRectF", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtSvg/qsvgrenderer.h:105
// index:0
// Public Visibility=Default Availability=Available
// [-2] void repaintNeeded()

/*
This signal is emitted whenever the rendering of the document needs to be updated, usually for the purposes of animation.
*/
func (this *QSvgRenderer) RepaintNeeded() {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QSvgRenderer13repaintNeededEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

//  body block end

//  keep block begin

func init_unused_11923() {
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
		qtwidgets.KeepMe()
	}
}

//  keep block end
