// +build !minimal

package qtcore

// func NewQModelIndexListFromPointer(cthis unsafe.Pointer) *QModelIndexList {
// 	return &QModelIndexList{&qtrt.CObject{cthis}}
// }

// func (this *QModelIndexListx) At(i int) *QModelIndex {
// 	rv, err := qtrt.InvokeQtFunc6("C_QModelIndexList_at_0", qtrt.FFI_TYPE_POINTER, this.GetCthis(), i)
// 	qtrt.ErrPrint(err, rv)
// 	return NewQModelIndexFromPointer(unsafe.Pointer(uintptr(rv)))
// }

// func (this *QModelIndexListx) Count(that *QModelIndex) int {
// 	rv, err := qtrt.InvokeQtFunc6("C_QModelIndexList_count_0", qtrt.FFI_TYPE_POINTER, this.GetCthis(), that.GetCthis())
// 	qtrt.ErrPrint(err, rv)
// 	return qtrt.Cretval2go("int", rv).(int)
// }

// func (this *QModelIndexListx) ConvertToSlice() (lst []*QModelIndex) {
// 	for i := 0; i < this.Count1(); i++ {
// 		lst = append(lst, this.At(i))
// 	}
// 	return
// }

// func (*QModelIndexListx) ConvertFromSlice(lst []*QModelIndex) *QModelIndexListx {
// 	return nil
// }
