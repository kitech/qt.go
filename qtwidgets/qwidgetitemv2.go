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
// extern C begin: 13
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
type QWidgetItemV2 struct {
	*QWidgetItem
}
type QWidgetItemV2_ITF interface {
	QWidgetItem_ITF
	QWidgetItemV2_PTR() *QWidgetItemV2
}

func (ptr *QWidgetItemV2) QWidgetItemV2_PTR() *QWidgetItemV2 { return ptr }

func (this *QWidgetItemV2) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QWidgetItem.GetCthis()
	}
}
func (this *QWidgetItemV2) SetCthis(cthis unsafe.Pointer) {
	this.QWidgetItem = NewQWidgetItemFromPointer(cthis)
}
func NewQWidgetItemV2FromPointer(cthis unsafe.Pointer) *QWidgetItemV2 {
	bcthis0 := NewQWidgetItemFromPointer(cthis)
	return &QWidgetItemV2{bcthis0}
}
func (*QWidgetItemV2) NewFromPointer(cthis unsafe.Pointer) *QWidgetItemV2 {
	return NewQWidgetItemV2FromPointer(cthis)
}

// /usr/include/qt/QtWidgets/qlayoutitem.h:148
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QWidgetItemV2(QWidget *)

/*

 */
func NewQWidgetItemV2(widget QWidget_ITF /*777 QWidget **/) *QWidgetItemV2 {
	var convArg0 unsafe.Pointer
	if widget != nil && widget.QWidget_PTR() != nil {
		convArg0 = widget.QWidget_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QWidgetItemV2C2EP7QWidget", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQWidgetItemV2FromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQWidgetItemV2)
	return gothis
}

// /usr/include/qt/QtWidgets/qlayoutitem.h:149
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QWidgetItemV2()

/*

 */
func DeleteQWidgetItemV2(this *QWidgetItemV2) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QWidgetItemV2D2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 88)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtWidgets/qlayoutitem.h:151
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QSize sizeHint() const

/*
Implemented in subclasses to return the preferred size of this item.
*/
func (this *QWidgetItemV2) SizeHint() *qtcore.QSize /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QWidgetItemV28sizeHintEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSize)
	return rv2
}

// /usr/include/qt/QtWidgets/qlayoutitem.h:152
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QSize minimumSize() const

/*
Implemented in subclasses to return the minimum size of this item.
*/
func (this *QWidgetItemV2) MinimumSize() *qtcore.QSize /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QWidgetItemV211minimumSizeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSize)
	return rv2
}

// /usr/include/qt/QtWidgets/qlayoutitem.h:153
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QSize maximumSize() const

/*
Implemented in subclasses to return the maximum size of this item.
*/
func (this *QWidgetItemV2) MaximumSize() *qtcore.QSize /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QWidgetItemV211maximumSizeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSize)
	return rv2
}

// /usr/include/qt/QtWidgets/qlayoutitem.h:154
// index:0
// Public virtual Visibility=Default Availability=Available
// [4] int heightForWidth(int) const

/*
Returns the preferred height for this layout item, given the width w.

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
func (this *QWidgetItemV2) HeightForWidth(width int) int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QWidgetItemV214heightForWidthEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), width)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

/*


 */
type QWidgetItemV2__ = int

//
const QWidgetItemV2__Dirty QWidgetItemV2__ = -123

//
const QWidgetItemV2__HfwCacheMaxSize QWidgetItemV2__ = 3

func (this *QWidgetItemV2) ItemName(val int) string {
	switch val {
	case QWidgetItemV2__Dirty: // -123
		return "Dirty"
	case QWidgetItemV2__HfwCacheMaxSize: // 3
		return "HfwCacheMaxSize"
	default:
		return fmt.Sprintf("%d", val)
	}
}
func QWidgetItemV2_ItemName(val int) string {
	var nilthis *QWidgetItemV2
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
}

//  keep block end
