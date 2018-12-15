package qtwidgets

// /usr/include/qt/QtWidgets/qlayoutitem.h
// #include <qlayoutitem.h>
// #include <QtWidgets>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 58
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "log"
import "github.com/kitech/qt.go/qtrt"
import "github.com/kitech/qt.go/qtcore"
import "github.com/kitech/qt.go/qtgui"

//  ext block end

//  body block begin

/*

 */
type QLayoutItem struct {
	*qtrt.CObject
}
type QLayoutItem_ITF interface {
	QLayoutItem_PTR() *QLayoutItem
}

func (ptr *QLayoutItem) QLayoutItem_PTR() *QLayoutItem { return ptr }

func (this *QLayoutItem) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QLayoutItem) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQLayoutItemFromPointer(cthis unsafe.Pointer) *QLayoutItem {
	return &QLayoutItem{&qtrt.CObject{cthis}}
}
func (*QLayoutItem) NewFromPointer(cthis unsafe.Pointer) *QLayoutItem {
	return NewQLayoutItemFromPointer(cthis)
}

// /usr/include/qt/QtWidgets/qlayoutitem.h:63
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void QLayoutItem(Qt::Alignment)

/*
Constructs a layout item with an alignment. Not all subclasses support alignment.
*/
func (*QLayoutItem) NewForInherit(alignment int) *QLayoutItem {
	return NewQLayoutItem(alignment)
}
func NewQLayoutItem(alignment int) *QLayoutItem {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QLayoutItemC2E6QFlagsIN2Qt13AlignmentFlagEE", qtrt.FFI_TYPE_POINTER, alignment)
	qtrt.ErrPrint(err, rv)
	gothis := NewQLayoutItemFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQLayoutItem)
	return gothis
}

// /usr/include/qt/QtWidgets/qlayoutitem.h:63
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void QLayoutItem(Qt::Alignment)

/*
Constructs a layout item with an alignment. Not all subclasses support alignment.
*/
func (*QLayoutItem) NewForInheritp() *QLayoutItem {
	return NewQLayoutItemp()
}
func NewQLayoutItemp() *QLayoutItem {
	// arg: 0, Qt::Alignment=Elaborated, Qt::Alignment=Typedef, QFlags<Qt::AlignmentFlag>, Unexposed
	alignment := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN11QLayoutItemC2E6QFlagsIN2Qt13AlignmentFlagEE", qtrt.FFI_TYPE_POINTER, alignment)
	qtrt.ErrPrint(err, rv)
	gothis := NewQLayoutItemFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQLayoutItem)
	return gothis
}

// /usr/include/qt/QtWidgets/qlayoutitem.h:64
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QLayoutItem()

/*

 */
func DeleteQLayoutItem(this *QLayoutItem) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QLayoutItemD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 16)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtWidgets/qlayoutitem.h:65
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [8] QSize sizeHint() const

/*
Implemented in subclasses to return the preferred size of this item.
*/
func (this *QLayoutItem) SizeHint() *qtcore.QSize /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QLayoutItem8sizeHintEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSize)
	return rv2
}

// /usr/include/qt/QtWidgets/qlayoutitem.h:66
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [8] QSize minimumSize() const

/*
Implemented in subclasses to return the minimum size of this item.
*/
func (this *QLayoutItem) MinimumSize() *qtcore.QSize /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QLayoutItem11minimumSizeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSize)
	return rv2
}

// /usr/include/qt/QtWidgets/qlayoutitem.h:67
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [8] QSize maximumSize() const

/*
Implemented in subclasses to return the maximum size of this item.
*/
func (this *QLayoutItem) MaximumSize() *qtcore.QSize /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QLayoutItem11maximumSizeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSize)
	return rv2
}

// /usr/include/qt/QtWidgets/qlayoutitem.h:68
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [4] Qt::Orientations expandingDirections() const

/*
Returns whether this layout item can make use of more space than sizeHint(). A value of Qt::Vertical or Qt::Horizontal means that it wants to grow in only one dimension, whereas Qt::Vertical | Qt::Horizontal means that it wants to grow in both dimensions.
*/
func (this *QLayoutItem) ExpandingDirections() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QLayoutItem19expandingDirectionsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qlayoutitem.h:69
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [-2] void setGeometry(const QRect &)

/*
Implemented in subclasses to set this item's geometry to r.

See also geometry().
*/
func (this *QLayoutItem) SetGeometry(arg0 qtcore.QRect_ITF) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QRect_PTR() != nil {
		convArg0 = arg0.QRect_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QLayoutItem11setGeometryERK5QRect", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlayoutitem.h:70
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [16] QRect geometry() const

/*
Returns the rectangle covered by this layout item.

See also setGeometry().
*/
func (this *QLayoutItem) Geometry() *qtcore.QRect /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QLayoutItem8geometryEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQRectFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQRect)
	return rv2
}

// /usr/include/qt/QtWidgets/qlayoutitem.h:71
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [1] bool isEmpty() const

/*
Implemented in subclasses to return whether this item is empty, i.e. whether it contains any widgets.
*/
func (this *QLayoutItem) IsEmpty() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QLayoutItem7isEmptyEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qlayoutitem.h:72
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool hasHeightForWidth() const

/*
Returns true if this layout's preferred height depends on its width; otherwise returns false. The default implementation returns false.

Reimplement this function in layout managers that support height for width.

See also heightForWidth() and QWidget::heightForWidth().
*/
func (this *QLayoutItem) HasHeightForWidth() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QLayoutItem17hasHeightForWidthEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qlayoutitem.h:73
// index:0
// Public virtual Visibility=Default Availability=Available
// [4] int heightForWidth(int) const

/*
Returns the preferred height for this layout item, given the width, which is not used in this default implementation.

The default implementation returns -1, indicating that the preferred height is independent of the width of the item. Using the function hasHeightForWidth() will typically be much faster than calling this function and testing for -1.

Reimplement this function in layout managers that support height for width. A typical implementation will look like this:


  int MyLayout::heightForWidth(int w) const
  {
      if (cache_dirty || cached_width != w) {
          // not all C++ compilers support "mutable"
          MyLayout *that = (MyLayout*)this;
          int h = calculateHeightForWidth(w);
          that->cached_hfw = h;
          return h;
      }
      return cached_hfw;
  }



Caching is strongly recommended; without it layout will take exponential time.

See also hasHeightForWidth().
*/
func (this *QLayoutItem) HeightForWidth(arg0 int) int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QLayoutItem14heightForWidthEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qlayoutitem.h:74
// index:0
// Public virtual Visibility=Default Availability=Available
// [4] int minimumHeightForWidth(int) const

/*
Returns the minimum height this widget needs for the given width, w. The default implementation simply returns heightForWidth(w).
*/
func (this *QLayoutItem) MinimumHeightForWidth(arg0 int) int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QLayoutItem21minimumHeightForWidthEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qlayoutitem.h:75
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void invalidate()

/*
Invalidates any cached information in this layout item.
*/
func (this *QLayoutItem) Invalidate() {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QLayoutItem10invalidateEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlayoutitem.h:77
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QWidget * widget()

/*
If this item manages a QWidget, returns that widget. Otherwise, nullptr is returned.

Note: While the functions layout() and spacerItem() perform casts, this function returns another object: QLayout and QSpacerItem inherit QLayoutItem, while QWidget does not.

See also layout() and spacerItem().
*/
func (this *QLayoutItem) Widget() *QWidget /*777 QWidget **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QLayoutItem6widgetEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQWidgetFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qlayoutitem.h:78
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QLayout * layout()

/*
If this item is a QLayout, it is returned as a QLayout; otherwise 0 is returned. This function provides type-safe casting.

See also spacerItem() and widget().
*/
func (this *QLayoutItem) Layout() *QLayout /*777 QLayout **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QLayoutItem6layoutEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQLayoutFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qlayoutitem.h:79
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QSpacerItem * spacerItem()

/*
If this item is a QSpacerItem, it is returned as a QSpacerItem; otherwise 0 is returned. This function provides type-safe casting.

See also layout() and widget().
*/
func (this *QLayoutItem) SpacerItem() *QSpacerItem /*777 QSpacerItem **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QLayoutItem10spacerItemEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQSpacerItemFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qlayoutitem.h:81
// index:0
// Public inline Visibility=Default Availability=Available
// [4] Qt::Alignment alignment() const

/*
Returns the alignment of this item.

See also setAlignment().
*/
func (this *QLayoutItem) Alignment() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QLayoutItem9alignmentEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qlayoutitem.h:82
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setAlignment(Qt::Alignment)

/*
Sets the alignment of this item to alignment.

Note: Item alignment is only supported by QLayoutItem subclasses where it would have a visual effect. Except for QSpacerItem, which provides blank space for layouts, all public Qt classes that inherit QLayoutItem support item alignment.

See also alignment().
*/
func (this *QLayoutItem) SetAlignment(a int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QLayoutItem12setAlignmentE6QFlagsIN2Qt13AlignmentFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), a)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlayoutitem.h:83
// index:0
// Public virtual Visibility=Default Availability=Available
// [4] QSizePolicy::ControlTypes controlTypes() const

/*
Returns the control type(s) for the layout item. For a QWidgetItem, the control type comes from the widget's size policy; for a QLayoutItem, the control types is derived from the layout's contents.

See also QSizePolicy::controlType().
*/
func (this *QLayoutItem) ControlTypes() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QLayoutItem12controlTypesEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
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
}

//  keep block end
