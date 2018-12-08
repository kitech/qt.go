package qtsvg

// /usr/include/qt/QtSvg/qgraphicssvgitem.h
// #include <qgraphicssvgitem.h>
// #include <QtSvg>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 27
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
type QGraphicsSvgItem struct {
	*qtrt.CObject
}
type QGraphicsSvgItem_ITF interface {
	QGraphicsSvgItem_PTR() *QGraphicsSvgItem
}

func (ptr *QGraphicsSvgItem) QGraphicsSvgItem_PTR() *QGraphicsSvgItem { return ptr }

func (this *QGraphicsSvgItem) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QGraphicsSvgItem) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQGraphicsSvgItemFromPointer(cthis unsafe.Pointer) *QGraphicsSvgItem {
	return &QGraphicsSvgItem{&qtrt.CObject{cthis}}
}
func (*QGraphicsSvgItem) NewFromPointer(cthis unsafe.Pointer) *QGraphicsSvgItem {
	return NewQGraphicsSvgItemFromPointer(cthis)
}

// /usr/include/qt/QtSvg/qgraphicssvgitem.h:58
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject() const

/*

 */
func (this *QGraphicsSvgItem) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QGraphicsSvgItem10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtSvg/qgraphicssvgitem.h:64
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QGraphicsSvgItem(QGraphicsItem *)

/*
Constructs a new SVG item with the given parent.
*/
func (*QGraphicsSvgItem) NewForInherit(parentItem qtwidgets.QGraphicsItem_ITF /*777 QGraphicsItem **/) *QGraphicsSvgItem {
	return NewQGraphicsSvgItem(parentItem)
}
func NewQGraphicsSvgItem(parentItem qtwidgets.QGraphicsItem_ITF /*777 QGraphicsItem **/) *QGraphicsSvgItem {
	var convArg0 unsafe.Pointer
	if parentItem != nil && parentItem.QGraphicsItem_PTR() != nil {
		convArg0 = parentItem.QGraphicsItem_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN16QGraphicsSvgItemC2EP13QGraphicsItem", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQGraphicsSvgItemFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QGraphicsSvgItem")
	return gothis
}

// /usr/include/qt/QtSvg/qgraphicssvgitem.h:64
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QGraphicsSvgItem(QGraphicsItem *)

/*
Constructs a new SVG item with the given parent.
*/
func (*QGraphicsSvgItem) NewForInheritp() *QGraphicsSvgItem {
	return NewQGraphicsSvgItemp()
}
func NewQGraphicsSvgItemp() *QGraphicsSvgItem {
	// arg: 0, QGraphicsItem *=Pointer, QGraphicsItem=Record, , Invalid
	var convArg0 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN16QGraphicsSvgItemC2EP13QGraphicsItem", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQGraphicsSvgItemFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QGraphicsSvgItem")
	return gothis
}

// /usr/include/qt/QtSvg/qgraphicssvgitem.h:65
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QGraphicsSvgItem(const QString &, QGraphicsItem *)

/*
Constructs a new SVG item with the given parent.
*/
func (*QGraphicsSvgItem) NewForInherit1(fileName string, parentItem qtwidgets.QGraphicsItem_ITF /*777 QGraphicsItem **/) *QGraphicsSvgItem {
	return NewQGraphicsSvgItem1(fileName, parentItem)
}
func NewQGraphicsSvgItem1(fileName string, parentItem qtwidgets.QGraphicsItem_ITF /*777 QGraphicsItem **/) *QGraphicsSvgItem {
	var tmpArg0 = qtcore.NewQString5(fileName)
	var convArg0 = tmpArg0.GetCthis()
	var convArg1 unsafe.Pointer
	if parentItem != nil && parentItem.QGraphicsItem_PTR() != nil {
		convArg1 = parentItem.QGraphicsItem_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN16QGraphicsSvgItemC2ERK7QStringP13QGraphicsItem", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	gothis := NewQGraphicsSvgItemFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QGraphicsSvgItem")
	return gothis
}

// /usr/include/qt/QtSvg/qgraphicssvgitem.h:65
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QGraphicsSvgItem(const QString &, QGraphicsItem *)

/*
Constructs a new SVG item with the given parent.
*/
func (*QGraphicsSvgItem) NewForInherit1p(fileName string) *QGraphicsSvgItem {
	return NewQGraphicsSvgItem1p(fileName)
}
func NewQGraphicsSvgItem1p(fileName string) *QGraphicsSvgItem {
	var tmpArg0 = qtcore.NewQString5(fileName)
	var convArg0 = tmpArg0.GetCthis()
	// arg: 1, QGraphicsItem *=Pointer, QGraphicsItem=Record, , Invalid
	var convArg1 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN16QGraphicsSvgItemC2ERK7QStringP13QGraphicsItem", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	gothis := NewQGraphicsSvgItemFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QGraphicsSvgItem")
	return gothis
}

// /usr/include/qt/QtSvg/qgraphicssvgitem.h:67
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setSharedRenderer(QSvgRenderer *)

/*
Sets renderer to be a shared QSvgRenderer on the item. By using this method one can share the same QSvgRenderer on a number of items. This means that the SVG file will be parsed only once. QSvgRenderer passed to this method has to exist for as long as this item is used.
*/
func (this *QGraphicsSvgItem) SetSharedRenderer(renderer QSvgRenderer_ITF /*777 QSvgRenderer **/) {
	var convArg0 unsafe.Pointer
	if renderer != nil && renderer.QSvgRenderer_PTR() != nil {
		convArg0 = renderer.QSvgRenderer_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN16QGraphicsSvgItem17setSharedRendererEP12QSvgRenderer", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtSvg/qgraphicssvgitem.h:68
// index:0
// Public Visibility=Default Availability=Available
// [8] QSvgRenderer * renderer() const

/*
Returns the currently use QSvgRenderer.
*/
func (this *QGraphicsSvgItem) Renderer() *QSvgRenderer /*777 QSvgRenderer **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QGraphicsSvgItem8rendererEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQSvgRendererFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtSvg/qgraphicssvgitem.h:70
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setElementId(const QString &)

/*
Sets the XML ID of the element to id.

Note: Setter function for property elementId.

See also elementId().
*/
func (this *QGraphicsSvgItem) SetElementId(id string) {
	var tmpArg0 = qtcore.NewQString5(id)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN16QGraphicsSvgItem12setElementIdERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtSvg/qgraphicssvgitem.h:71
// index:0
// Public Visibility=Default Availability=Available
// [8] QString elementId() const

/*
Returns the XML ID the element that is currently being rendered. Returns an empty string if the whole file is being rendered.

Note: Getter function for property elementId.

See also setElementId().
*/
func (this *QGraphicsSvgItem) ElementId() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QGraphicsSvgItem9elementIdEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtSvg/qgraphicssvgitem.h:73
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setCachingEnabled(bool)

/*

 */
func (this *QGraphicsSvgItem) SetCachingEnabled(arg0 bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN16QGraphicsSvgItem17setCachingEnabledEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtSvg/qgraphicssvgitem.h:74
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isCachingEnabled() const

/*

 */
func (this *QGraphicsSvgItem) IsCachingEnabled() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QGraphicsSvgItem16isCachingEnabledEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtSvg/qgraphicssvgitem.h:76
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setMaximumCacheSize(const QSize &)

/*
Sets the maximum device coordinate cache size of the item to size. If the item is cached using QGraphicsItem::DeviceCoordinateCache mode, caching is bypassed if the extension of the item in device coordinates is larger than size.

The cache corresponds to the QPixmap which is used to cache the results of the rendering. Use QPixmapCache::setCacheLimit() to set limitations on the whole cache and use setMaximumCacheSize() when setting cache size for individual items.

Note: Setter function for property maximumCacheSize.

See also maximumCacheSize() and QGraphicsItem::cacheMode().
*/
func (this *QGraphicsSvgItem) SetMaximumCacheSize(size qtcore.QSize_ITF) {
	var convArg0 unsafe.Pointer
	if size != nil && size.QSize_PTR() != nil {
		convArg0 = size.QSize_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN16QGraphicsSvgItem19setMaximumCacheSizeERK5QSize", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtSvg/qgraphicssvgitem.h:77
// index:0
// Public Visibility=Default Availability=Available
// [8] QSize maximumCacheSize() const

/*
Returns the current maximum size of the device coordinate cache for this item. If the item is cached using QGraphicsItem::DeviceCoordinateCache mode, caching is bypassed if the extension of the item in device coordinates is larger than the maximum size.

The default maximum cache size is 1024x768. QPixmapCache::cacheLimit() gives the cumulative bounds of the whole cache, whereas maximumCacheSize() refers to a maximum cache size for this particular item.

Note: Getter function for property maximumCacheSize.

See also setMaximumCacheSize() and QGraphicsItem::cacheMode().
*/
func (this *QGraphicsSvgItem) MaximumCacheSize() *qtcore.QSize /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QGraphicsSvgItem16maximumCacheSizeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSize)
	return rv2
}

// /usr/include/qt/QtSvg/qgraphicssvgitem.h:79
// index:0
// Public virtual Visibility=Default Availability=Available
// [32] QRectF boundingRect() const

/*
Reimplemented from QGraphicsItem::boundingRect().

Returns the bounding rectangle of this item.
*/
func (this *QGraphicsSvgItem) BoundingRect() *qtcore.QRectF /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QGraphicsSvgItem12boundingRectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQRectFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQRectF)
	return rv2
}

// /usr/include/qt/QtSvg/qgraphicssvgitem.h:81
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void paint(QPainter *, const QStyleOptionGraphicsItem *, QWidget *)

/*
Reimplemented from QGraphicsItem::paint().
*/
func (this *QGraphicsSvgItem) Paint(painter qtgui.QPainter_ITF /*777 QPainter **/, option qtwidgets.QStyleOptionGraphicsItem_ITF /*777 const QStyleOptionGraphicsItem **/, widget qtwidgets.QWidget_ITF /*777 QWidget **/) {
	var convArg0 unsafe.Pointer
	if painter != nil && painter.QPainter_PTR() != nil {
		convArg0 = painter.QPainter_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if option != nil && option.QStyleOptionGraphicsItem_PTR() != nil {
		convArg1 = option.QStyleOptionGraphicsItem_PTR().GetCthis()
	}
	var convArg2 unsafe.Pointer
	if widget != nil && widget.QWidget_PTR() != nil {
		convArg2 = widget.QWidget_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN16QGraphicsSvgItem5paintEP8QPainterPK24QStyleOptionGraphicsItemP7QWidget", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtSvg/qgraphicssvgitem.h:81
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void paint(QPainter *, const QStyleOptionGraphicsItem *, QWidget *)

/*
Reimplemented from QGraphicsItem::paint().
*/
func (this *QGraphicsSvgItem) Paintp(painter qtgui.QPainter_ITF /*777 QPainter **/, option qtwidgets.QStyleOptionGraphicsItem_ITF /*777 const QStyleOptionGraphicsItem **/) {
	var convArg0 unsafe.Pointer
	if painter != nil && painter.QPainter_PTR() != nil {
		convArg0 = painter.QPainter_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if option != nil && option.QStyleOptionGraphicsItem_PTR() != nil {
		convArg1 = option.QStyleOptionGraphicsItem_PTR().GetCthis()
	}
	// arg: 2, QWidget *=Pointer, QWidget=Record, , Invalid
	var convArg2 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN16QGraphicsSvgItem5paintEP8QPainterPK24QStyleOptionGraphicsItemP7QWidget", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtSvg/qgraphicssvgitem.h:86
// index:0
// Public virtual Visibility=Default Availability=Available
// [4] int type() const

/*
Reimplemented from QGraphicsItem::type().
*/
func (this *QGraphicsSvgItem) Type() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QGraphicsSvgItem4typeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

func DeleteQGraphicsSvgItem(this *QGraphicsSvgItem) {
	rv, err := qtrt.InvokeQtFunc6("_ZN16QGraphicsSvgItemD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

/*


 */
type QGraphicsSvgItem__ = int

//
const QGraphicsSvgItem__Type QGraphicsSvgItem__ = 13

func (this *QGraphicsSvgItem) ItemName(val int) string {
	return qtrt.GetClassEnumItemName(this, val)
}
func QGraphicsSvgItem_ItemName(val int) string {
	var nilthis *QGraphicsSvgItem
	return nilthis.ItemName(val)
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
		qtwidgets.KeepMe()
	}
}

//  keep block end
