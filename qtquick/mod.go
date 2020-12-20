package qtquick

/*
#cgo CFLAGS: -std=c11
// #cgo LDFLAGS: -lQt5Inline
//  -lQt5Core -lQt5Gui -lQt5Widgets
*/
import "C"
import (
	"unsafe"

	"github.com/kitech/qt.go/qtrt"
)

func KeepMe() {}
func init()   { qtrt.RegisterSubPackage("Quick") }

type Voidptr = unsafe.Pointer

/////
type QQuickItemListx struct {
	*qtrt.CObject
}

func (this *QQuickItemListx) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QQuickItemListx) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQQuickItemListxFromPointer(cthis unsafe.Pointer) *QQuickItemListx {
	return &QQuickItemListx{&qtrt.CObject{cthis}}
}
func NewQQuickItemListFromPointer(cthis unsafe.Pointer) *QQuickItemList {
	return &QQuickItemList{&qtrt.CObject{cthis}}
}
func (*QQuickItemListx) NewFromPointer(cthis unsafe.Pointer) *QQuickItemListx {
	return NewQQuickItemListxFromPointer(cthis)
}

func (this *QQuickItemListx) At(i int) *QQuickItem {
	rv, err := qtrt.InvokeQtFunc6("C_QQuickItemList_at_0", qtrt.FFI_TYPE_POINTER, this.GetCthis(), i)
	qtrt.ErrPrint(err, rv)
	return NewQQuickItemFromPointer(unsafe.Pointer(uintptr(rv)))
}

func (this *QQuickItemListx) Count(that *QQuickItem) int {
	rv, err := qtrt.InvokeQtFunc6("C_QQuickItemList_count_0", qtrt.FFI_TYPE_POINTER, this.GetCthis(), that.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int)
}
func (this *QQuickItemListx) Count_1() int {
	rv, err := qtrt.InvokeQtFunc6("C_QQuickItemList_count_1", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int)
}

func (this *QQuickItemListx) ConvertToSlice() (lst []*QQuickItem) {
	for i := 0; i < this.Count_1(); i++ {
		lst = append(lst, this.At(i))
	}
	return
}

func (*QQuickItemListx) ConvertFromSlice(lst []*QQuickItem) *QQuickItemListx {
	return nil
}
