// +build !minimal

package qtwidgets

import (
	"unsafe"

	"github.com/kitech/qt.go/qtrt"
)

// func NewQGraphicsItemListFromPointer(cthis unsafe.Pointer) *QGraphicsItemList {
// 	return &QGraphicsItemList{&qtrt.CObject{cthis}}
// }

func (this *QGraphicsItemListx) At(i int) *QGraphicsItem {
	rv, err := qtrt.InvokeQtFunc6("C_QGraphicsItemList_at_0", qtrt.FFI_TYPE_POINTER, this.GetCthis(), i)
	qtrt.ErrPrint(err, rv)
	return QGraphicsItemFromptr(unsafe.Pointer(uintptr(rv)))
}

func (this *QGraphicsItemListx) Count(that *QGraphicsItem) int {
	rv, err := qtrt.InvokeQtFunc6("C_QGraphicsItemList_count_0", qtrt.FFI_TYPE_POINTER, this.GetCthis(), that.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int)
}

func (this *QGraphicsItemListx) ConvertToSlice() (lst []*QGraphicsItem) {
	for i := 0; i < this.Count1(); i++ {
		lst = append(lst, this.At(i))
	}
	return
}

func (*QGraphicsItemListx) ConvertFromSlice(lst []*QGraphicsItem) *QGraphicsItemListx {
	return nil
}
