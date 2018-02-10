package qtcore

import (
	"gopp"
	"unsafe"

	"qt.go/qtrt"
)

/*
#cgo CFLAGS: -std=c11
// #cgo LDFLAGS: -lQt5Inline
//  -lQt5Core -lQt5Gui -lQt5Widgets
*/
// import "C"

func KeepMe() {}

/////
type QObjectListx struct {
	*qtrt.CObject
}

func (this *QObjectListx) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QObjectListx) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQObjectListxFromPointer(cthis unsafe.Pointer) *QObjectListx {
	return &QObjectListx{&qtrt.CObject{cthis}}
}
func NewQObjectListFromPointer(cthis unsafe.Pointer) *QObjectList {
	return &QObjectList{&qtrt.CObject{cthis}}
}
func (*QObjectListx) NewFromPointer(cthis unsafe.Pointer) *QObjectListx {
	return NewQObjectListxFromPointer(cthis)
}
func DeleteQObjectList(*QObjectList) {} // TODO

func (this *QObjectListx) At(i int) *QObject {
	rv, err := qtrt.InvokeQtFunc6("C_QObjectList_at_0", qtrt.FFI_TYPE_POINTER, this.GetCthis(), i)
	gopp.ErrPrint(err, rv)
	return NewQObjectFromPointer(unsafe.Pointer(uintptr(rv)))
}

func (this *QObjectListx) Count() int {
	rv, err := qtrt.InvokeQtFunc6("C_QObjectList_count_0", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int)
}

func (this *QObjectListx) ConvertToSlice() (lst []*QObject) {
	for i := 0; i < this.Count(); i++ {
		lst = append(lst, this.At(i))
	}
	return
}

func (*QObjectListx) ConvertFromSlice(lst []*QObject) *QObjectListx {
	return nil
}

/////
type QFileInfoListx struct {
	*qtrt.CObject
}

func (this *QFileInfoListx) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QFileInfoListx) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQFileInfoListxFromPointer(cthis unsafe.Pointer) *QFileInfoListx {
	return &QFileInfoListx{&qtrt.CObject{cthis}}
}
func NewQFileInfoListFromPointer(cthis unsafe.Pointer) *QFileInfoList {
	return &QFileInfoList{&qtrt.CObject{cthis}}
}
func (*QFileInfoListx) NewFromPointer(cthis unsafe.Pointer) *QFileInfoListx {
	return NewQFileInfoListxFromPointer(cthis)
}

func (this *QFileInfoListx) At(i int) *QFileInfo {
	rv, err := qtrt.InvokeQtFunc6("C_QFileInfoList_at_0", qtrt.FFI_TYPE_POINTER, this.GetCthis(), i)
	gopp.ErrPrint(err, rv)
	return NewQFileInfoFromPointer(unsafe.Pointer(uintptr(rv)))
}

func (this *QFileInfoListx) Count() int {
	rv, err := qtrt.InvokeQtFunc6("C_QFileInfoList_count_0", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int)
}

func (this *QFileInfoListx) ConvertToSlice() (lst []*QFileInfo) {
	for i := 0; i < this.Count(); i++ {
		lst = append(lst, this.At(i))
	}
	return
}

func (*QFileInfoListx) ConvertFromSlice(lst []*QFileInfo) *QFileInfoListx {
	return nil
}

/////
type QModelIndexListx struct {
	*qtrt.CObject
}

func (this *QModelIndexListx) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QModelIndexListx) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQModelIndexListxFromPointer(cthis unsafe.Pointer) *QModelIndexListx {
	return &QModelIndexListx{&qtrt.CObject{cthis}}
}
func NewQModelIndexListFromPointer(cthis unsafe.Pointer) *QModelIndexList {
	return &QModelIndexList{&qtrt.CObject{cthis}}
}
func (*QModelIndexListx) NewFromPointer(cthis unsafe.Pointer) *QModelIndexListx {
	return NewQModelIndexListxFromPointer(cthis)
}

func (this *QModelIndexListx) At(i int) *QModelIndex {
	rv, err := qtrt.InvokeQtFunc6("C_QModelIndexList_at_0", qtrt.FFI_TYPE_POINTER, this.GetCthis(), i)
	gopp.ErrPrint(err, rv)
	return NewQModelIndexFromPointer(unsafe.Pointer(uintptr(rv)))
}

func (this *QModelIndexListx) Count() int {
	rv, err := qtrt.InvokeQtFunc6("C_QModelIndexList_count_0", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int)
}

func (this *QModelIndexListx) ConvertToSlice() (lst []*QModelIndex) {
	for i := 0; i < this.Count(); i++ {
		lst = append(lst, this.At(i))
	}
	return
}

func (*QModelIndexListx) ConvertFromSlice(lst []*QModelIndex) *QModelIndexListx {
	return nil
}

/////
type QVariantListx struct {
	*qtrt.CObject
}

func (this *QVariantListx) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QVariantListx) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQVariantListxFromPointer(cthis unsafe.Pointer) *QVariantListx {
	return &QVariantListx{&qtrt.CObject{cthis}}
}
func NewQVariantListFromPointer(cthis unsafe.Pointer) *QVariantList {
	return &QVariantList{&qtrt.CObject{cthis}}
}
func (*QVariantListx) NewFromPointer(cthis unsafe.Pointer) *QVariantListx {
	return NewQVariantListxFromPointer(cthis)
}

func (this *QVariantListx) At(i int) *QVariant {
	rv, err := qtrt.InvokeQtFunc6("C_QVariantList_at_0", qtrt.FFI_TYPE_POINTER, this.GetCthis(), i)
	gopp.ErrPrint(err, rv)
	return NewQVariantFromPointer(unsafe.Pointer(uintptr(rv)))
}

func (this *QVariantListx) Count() int {
	rv, err := qtrt.InvokeQtFunc6("C_QVariantList_count_0", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int)
}

func (this *QVariantListx) ConvertToSlice() (lst []*QVariant) {
	for i := 0; i < this.Count(); i++ {
		lst = append(lst, this.At(i))
	}
	return
}

func (*QVariantListx) ConvertFromSlice(lst []*QVariant) *QVariantListx {
	return nil
}
