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
// extern C begin: 12
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
type QWidgetItem struct {
	*QLayoutItem
}
type QWidgetItem_ITF interface {
	QLayoutItem_ITF
	QWidgetItem_PTR() *QWidgetItem
}

func (ptr *QWidgetItem) QWidgetItem_PTR() *QWidgetItem { return ptr }

func (this *QWidgetItem) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QLayoutItem.GetCthis()
	}
}
func (this *QWidgetItem) SetCthis(cthis unsafe.Pointer) {
	this.QLayoutItem = NewQLayoutItemFromPointer(cthis)
}
func NewQWidgetItemFromPointer(cthis unsafe.Pointer) *QWidgetItem {
	bcthis0 := NewQLayoutItemFromPointer(cthis)
	return &QWidgetItem{bcthis0}
}
func (*QWidgetItem) NewFromPointer(cthis unsafe.Pointer) *QWidgetItem {
	return NewQWidgetItemFromPointer(cthis)
}

// /usr/include/qt/QtWidgets/qlayoutitem.h:126
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void QWidgetItem(QWidget *)

/*

 */
func (*QWidgetItem) NewForInherit(w QWidget_ITF /*777 QWidget **/) *QWidgetItem {
	return NewQWidgetItem(w)
}
func NewQWidgetItem(w QWidget_ITF /*777 QWidget **/) *QWidgetItem {
	var convArg0 unsafe.Pointer
	if w != nil && w.QWidget_PTR() != nil {
		convArg0 = w.QWidget_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QWidgetItemC2EP7QWidget", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQWidgetItemFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQWidgetItem)
	return gothis
}

// /usr/include/qt/QtWidgets/qlayoutitem.h:127
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QWidgetItem()

/*

 */
func DeleteQWidgetItem(this *QWidgetItem) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QWidgetItemD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 24)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtWidgets/qlayoutitem.h:129
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QSize sizeHint() const

/*
Implemented in subclasses to return the preferred size of this item.
*/
func (this *QWidgetItem) SizeHint() *qtcore.QSize /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QWidgetItem8sizeHintEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSize)
	return rv2
}

// /usr/include/qt/QtWidgets/qlayoutitem.h:130
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QSize minimumSize() const

/*
Implemented in subclasses to return the minimum size of this item.
*/
func (this *QWidgetItem) MinimumSize() *qtcore.QSize /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QWidgetItem11minimumSizeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSize)
	return rv2
}

// /usr/include/qt/QtWidgets/qlayoutitem.h:131
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QSize maximumSize() const

/*
Implemented in subclasses to return the maximum size of this item.
*/
func (this *QWidgetItem) MaximumSize() *qtcore.QSize /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QWidgetItem11maximumSizeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSize)
	return rv2
}

// /usr/include/qt/QtWidgets/qlayoutitem.h:132
// index:0
// Public virtual Visibility=Default Availability=Available
// [4] Qt::Orientations expandingDirections() const

/*
Returns whether this layout item can make use of more space than sizeHint(). A value of Qt::Vertical or Qt::Horizontal means that it wants to grow in only one dimension, whereas Qt::Vertical | Qt::Horizontal means that it wants to grow in both dimensions.
*/
func (this *QWidgetItem) ExpandingDirections() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QWidgetItem19expandingDirectionsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qlayoutitem.h:133
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool isEmpty() const

/*
Implemented in subclasses to return whether this item is empty, i.e. whether it contains any widgets.
*/
func (this *QWidgetItem) IsEmpty() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QWidgetItem7isEmptyEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qlayoutitem.h:134
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void setGeometry(const QRect &)

/*
Implemented in subclasses to set this item's geometry to r.

See also geometry().
*/
func (this *QWidgetItem) SetGeometry(arg0 qtcore.QRect_ITF) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QRect_PTR() != nil {
		convArg0 = arg0.QRect_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QWidgetItem11setGeometryERK5QRect", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlayoutitem.h:135
// index:0
// Public virtual Visibility=Default Availability=Available
// [16] QRect geometry() const

/*
Returns the rectangle covered by this layout item.

See also setGeometry().
*/
func (this *QWidgetItem) Geometry() *qtcore.QRect /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QWidgetItem8geometryEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQRectFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQRect)
	return rv2
}

// /usr/include/qt/QtWidgets/qlayoutitem.h:136
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QWidget * widget()

/*
If this item manages a QWidget, returns that widget. Otherwise, nullptr is returned.

Note: While the functions layout() and spacerItem() perform casts, this function returns another object: QLayout and QSpacerItem inherit QLayoutItem, while QWidget does not.

See also layout() and spacerItem().
*/
func (this *QWidgetItem) Widget() *QWidget /*777 QWidget **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QWidgetItem6widgetEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQWidgetFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qlayoutitem.h:138
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool hasHeightForWidth() const

/*
Returns true if this layout's preferred height depends on its width; otherwise returns false. The default implementation returns false.

Reimplement this function in layout managers that support height for width.

See also heightForWidth() and QWidget::heightForWidth().
*/
func (this *QWidgetItem) HasHeightForWidth() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QWidgetItem17hasHeightForWidthEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qlayoutitem.h:139
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
func (this *QWidgetItem) HeightForWidth(arg0 int) int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QWidgetItem14heightForWidthEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qlayoutitem.h:140
// index:0
// Public virtual Visibility=Default Availability=Available
// [4] QSizePolicy::ControlTypes controlTypes() const

/*
Returns the control type(s) for the layout item. For a QWidgetItem, the control type comes from the widget's size policy; for a QLayoutItem, the control types is derived from the layout's contents.

See also QSizePolicy::controlType().
*/
func (this *QWidgetItem) ControlTypes() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QWidgetItem12controlTypesEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

//  body block end

//  keep block begin

func init_unused_11075() {
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
