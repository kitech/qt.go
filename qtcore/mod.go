package qtcore

import (
	"unsafe"

	"github.com/kitech/qt.go/qtrt"
)

/*
#cgo CFLAGS: -std=c11
// #cgo LDFLAGS: -lQt5Inline
//  -lQt5Core -lQt5Gui -lQt5Widgets
*/
// import "C"

func KeepMe() {}
func init()   { qtrt.RegisterSubPackage("Core") }

type Voidptr = unsafe.Pointer

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

// func NewQObjectListFromPointer(cthis unsafe.Pointer) *QObjectList {
// 	return &QObjectList{&qtrt.CObject{cthis}}
// }
func (*QObjectListx) NewFromPointer(cthis unsafe.Pointer) *QObjectListx {
	return NewQObjectListxFromPointer(cthis)
}

// func DeleteQObjectList(*QObjectList) {} // TODO

func (this *QObjectListx) At(i int) *QObject {
	rv, err := qtrt.InvokeQtFunc6("C_QObjectList_at_0", qtrt.FFI_TYPE_POINTER, this.GetCthis(), i)
	qtrt.ErrPrint(err, rv)
	return QObjectFromptr(qtrt.DerefPtr2(unsafe.Pointer(uintptr(rv))))
}

func (this *QObjectListx) Count(that *QObject) int {
	rv, err := qtrt.InvokeQtFunc6("C_QObjectList_count_0", qtrt.FFI_TYPE_POINTER, this.GetCthis(), that.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int)
}
func (this *QObjectListx) Count1() int {
	rv, err := qtrt.InvokeQtFunc6("C_QObjectList_count_1", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int)
}

func (this *QObjectListx) ConvertToSlice() (lst []*QObject) {
	for i := 0; i < this.Count1(); i++ {
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

// func NewQFileInfoListFromPointer(cthis unsafe.Pointer) *QFileInfoList {
// 	return &QFileInfoList{&qtrt.CObject{cthis}}
// }
func (*QFileInfoListx) NewFromPointer(cthis unsafe.Pointer) *QFileInfoListx {
	return NewQFileInfoListxFromPointer(cthis)
}

// func (this *QFileInfoListx) At(i int) *QFileInfo {
// 	rv, err := qtrt.InvokeQtFunc6("C_QFileInfoList_at_0", qtrt.FFI_TYPE_POINTER, this.GetCthis(), i)
// 	qtrt.ErrPrint(err, rv)
// 	return NewQFileInfoFromPointer(unsafe.Pointer(uintptr(rv)))
// }

// func (this *QFileInfoListx) Count(that *QFileInfo) int {
// 	rv, err := qtrt.InvokeQtFunc6("C_QFileInfoList_count_0", qtrt.FFI_TYPE_POINTER, this.GetCthis(), that.GetCthis())
// 	qtrt.ErrPrint(err, rv)
// 	return qtrt.Cretval2go("int", rv).(int)
// }
// func (this *QFileInfoListx) Count1() int {
// 	rv, err := qtrt.InvokeQtFunc6("C_QFileInfoList_count_1", qtrt.FFI_TYPE_POINTER, this.GetCthis())
// 	qtrt.ErrPrint(err, rv)
// 	return qtrt.Cretval2go("int", rv).(int)
// }

// func (this *QFileInfoListx) ConvertToSlice() (lst []*QFileInfo) {
// 	for i := 0; i < this.Count1(); i++ {
// 		lst = append(lst, this.At(i))
// 	}
// 	return
// }

// func (*QFileInfoListx) ConvertFromSlice(lst []*QFileInfo) *QFileInfoListx {
// 	return nil
// }

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
func (*QModelIndexListx) NewFromPointer(cthis unsafe.Pointer) *QModelIndexListx {
	return NewQModelIndexListxFromPointer(cthis)
}

func (this *QModelIndexListx) Count1() int {
	rv, err := qtrt.InvokeQtFunc6("C_QModelIndexList_count_1", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int)
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

// func NewQVariantListFromPointer(cthis unsafe.Pointer) *QVariantList {
// 	return &QVariantList{&qtrt.CObject{cthis}}
// }
func (*QVariantListx) NewFromPointer(cthis unsafe.Pointer) *QVariantListx {
	return NewQVariantListxFromPointer(cthis)
}

func (this *QVariantListx) At(i int) *QVariant {
	rv, err := qtrt.InvokeQtFunc6("C_QVariantList_at_0", qtrt.FFI_TYPE_POINTER, this.GetCthis(), i)
	qtrt.ErrPrint(err, rv)
	return QVariantFromptr(unsafe.Pointer(uintptr(rv)))
}

func (this *QVariantListx) Count1() int {
	rv, err := qtrt.InvokeQtFunc6("C_QVariantList_count_1", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int)
}

func (this *QVariantListx) Size() int {
	rv, err := qtrt.InvokeQtFunc6("C_QVariantList_size_0", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int)
}

func (this *QVariantListx) ConvertToSlice() (lst []*QVariant) {
	for i := 0; i < this.Size(); i++ {
		lst = append(lst, this.At(i))
	}
	return
}

func (*QVariantListx) ConvertFromSlice(lst []*QVariant) *QVariantListx {
	return nil
}

/////
type QStringListx struct {
	*qtrt.CObject
}

func (this *QStringListx) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QStringListx) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQStringListxFromPointer(cthis unsafe.Pointer) *QStringListx {
	return &QStringListx{&qtrt.CObject{cthis}}
}

// func _NewQStringListFromPointer(cthis unsafe.Pointer) *QStringList {
// 	return &QStringList{&qtrt.CObject{cthis}}
// }
func (*QStringListx) NewFromPointer(cthis unsafe.Pointer) *QStringListx {
	return NewQStringListxFromPointer(cthis)
}
func DeleteQStringListx(this *QStringListx) {
	qtrt.InvokeQtFunc6("C_QStringList_dtor", qtrt.FFI_TYPE_POINTER, this.GetCthis())
}
func (this *QStringListx) Release() { this.CObject = nil }

func (this *QStringListx) At(i int) string {
	rv, err := qtrt.InvokeQtFunc6("C_QStringList_at_0", qtrt.FFI_TYPE_POINTER, this.GetCthis(), i)
	qtrt.ErrPrint(err, rv)
	qstr := QStringFromptr(unsafe.Pointer(uintptr(rv)))
	ret := qstr.ToUtf8().Data()
	DeleteQString(qstr)
	return ret
}

func (this *QStringListx) Count(s string) int {
	var convArg0 = qtrt.CString(s)
	defer qtrt.FreeMem(convArg0)
	rv, err := qtrt.InvokeQtFunc6("C_QStringList_count_0", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int)
}

func (this *QStringListx) Count1() int {
	rv, err := qtrt.InvokeQtFunc6("C_QStringList_count_1", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int)
}

func (this *QStringListx) Size() int {
	rv, err := qtrt.InvokeQtFunc6("C_QStringList_size_0", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int)
}

func (this *QStringListx) ConvertToSlice() (lst []string) {
	for i := 0; i < this.Size(); i++ {
		lst = append(lst, this.At(i))
	}
	return
}

func (*QStringListx) ConvertFromSlice(lst []string) *QStringListx {
	return nil
}

/////
type QUrlListx struct {
	*qtrt.CObject
}

func (this *QUrlListx) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QUrlListx) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQUrlListxFromPointer(cthis unsafe.Pointer) *QUrlListx {
	return &QUrlListx{&qtrt.CObject{cthis}}
}

// func NewQUrlListFromPointer(cthis unsafe.Pointer) *QUrlList {
// 	return &QUrlList{&qtrt.CObject{cthis}}
// }
func (*QUrlListx) NewFromPointer(cthis unsafe.Pointer) *QUrlListx {
	return NewQUrlListxFromPointer(cthis)
}
func DeleteQUrlListx(this *QUrlListx) {
	qtrt.InvokeQtFunc6("C_QUrlList_dtor", qtrt.FFI_TYPE_POINTER, this.GetCthis())
}
func (this *QUrlListx) Release() { this.CObject = nil }

// func (this *QUrlListx) At(i int) *QUrl {
// 	rv, err := qtrt.InvokeQtFunc6("C_QUrlList_at_0", qtrt.FFI_TYPE_POINTER, this.GetCthis(), i)
// 	qtrt.ErrPrint(err, rv)
// 	urlo := NewQUrlFromPointer(unsafe.Pointer(uintptr(rv)))
// 	return urlo
// }

func (this *QUrlListx) Count(s string) int {
	var convArg0 = qtrt.CString(s)
	defer qtrt.FreeMem(convArg0)
	rv, err := qtrt.InvokeQtFunc6("C_QUrlList_count_0", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int)
}

func (this *QUrlListx) Count1() int {
	rv, err := qtrt.InvokeQtFunc6("C_QUrlList_count_1", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int)
}

func (this *QUrlListx) Size() int {
	rv, err := qtrt.InvokeQtFunc6("C_QUrlList_size_0", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int)
}

// func (this *QUrlListx) ConvertToSlice() (lst []*QUrl) {
// 	for i := 0; i < this.Size(); i++ {
// 		lst = append(lst, this.At(i))
// 	}
// 	return
// }

// func (*QUrlListx) ConvertFromSlice(lst []*QUrl) *QUrlListx {
// 	return nil
// }

/////
type QSizeListx struct {
	*qtrt.CObject
}

func (this *QSizeListx) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QSizeListx) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQSizeListxFromPointer(cthis unsafe.Pointer) *QSizeListx {
	return &QSizeListx{&qtrt.CObject{cthis}}
}

// func NewQSizeListFromPointer(cthis unsafe.Pointer) *QSizeList {
// 	return &QSizeList{&qtrt.CObject{cthis}}
// }
func (*QSizeListx) NewFromPointer(cthis unsafe.Pointer) *QSizeListx {
	return NewQSizeListxFromPointer(cthis)
}
func DeleteQSizeListx(this *QSizeListx) {
	qtrt.InvokeQtFunc6("C_QSizeList_dtor", qtrt.FFI_TYPE_POINTER, this.GetCthis())
}
func (this *QSizeListx) Release() { this.CObject = nil }

func (this *QSizeListx) At(i int) *QSize {
	rv, err := qtrt.InvokeQtFunc6("C_QSizeList_at_0", qtrt.FFI_TYPE_POINTER, this.GetCthis(), i)
	qtrt.ErrPrint(err, rv)
	urlo := QSizeFromptr(unsafe.Pointer(uintptr(rv)))
	return urlo
}

func (this *QSizeListx) Count(s string) int {
	var convArg0 = qtrt.CString(s)
	defer qtrt.FreeMem(convArg0)
	rv, err := qtrt.InvokeQtFunc6("C_QSizeList_count_0", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int)
}

func (this *QSizeListx) Count1() int {
	rv, err := qtrt.InvokeQtFunc6("C_QSizeList_count_1", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int)
}

func (this *QSizeListx) Size() int {
	rv, err := qtrt.InvokeQtFunc6("C_QSizeList_size_0", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int)
}

func (this *QSizeListx) ConvertToSlice() (lst []*QSize) {
	for i := 0; i < this.Size(); i++ {
		lst = append(lst, this.At(i))
	}
	return
}

func (*QSizeListx) ConvertFromSlice(lst []*QSize) *QSizeListx {
	return nil
}
