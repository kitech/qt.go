package qtgui

// /usr/include/qt/QtGui/qaccessible.h
// #include <qaccessible.h>
// #include <QtGui>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 25
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "gopp"
import "qt.go/qtrt"
import "qt.go/qtcore"

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
		qtrt.KeepMe()
	}
	if false {
		gopp.KeepMe()
	}
	if false {
		qtcore.KeepMe()
	}
}

//  ext block end

//  body block begin

type QAccessibleTextInterface struct {
	*qtrt.CObject
}

func (this *QAccessibleTextInterface) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QAccessibleTextInterface) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQAccessibleTextInterfaceFromPointer(cthis unsafe.Pointer) *QAccessibleTextInterface {
	return &QAccessibleTextInterface{&qtrt.CObject{cthis}}
}
func (*QAccessibleTextInterface) NewFromPointer(cthis unsafe.Pointer) *QAccessibleTextInterface {
	return NewQAccessibleTextInterfaceFromPointer(cthis)
}

// /usr/include/qt/QtGui/qaccessible.h:523
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QAccessibleTextInterface()
func DeleteQAccessibleTextInterface(this *QAccessibleTextInterface) {
	rv, err := qtrt.InvokeQtFunc6("_ZN24QAccessibleTextInterfaceD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 8)
	gopp.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtGui/qaccessible.h:525
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [-2] void selection(int, int *, int *)
func (this *QAccessibleTextInterface) Selection(selectionIndex int, startOffset unsafe.Pointer /*666*/, endOffset unsafe.Pointer /*666*/) {
	rv, err := qtrt.InvokeQtFunc6("_ZNK24QAccessibleTextInterface9selectionEiPiS0_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), selectionIndex, &startOffset, &endOffset)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qaccessible.h:526
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [4] int selectionCount()
func (this *QAccessibleTextInterface) SelectionCount() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK24QAccessibleTextInterface14selectionCountEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qaccessible.h:527
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [-2] void addSelection(int, int)
func (this *QAccessibleTextInterface) AddSelection(startOffset int, endOffset int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN24QAccessibleTextInterface12addSelectionEii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), startOffset, endOffset)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qaccessible.h:528
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [-2] void removeSelection(int)
func (this *QAccessibleTextInterface) RemoveSelection(selectionIndex int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN24QAccessibleTextInterface15removeSelectionEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), selectionIndex)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qaccessible.h:529
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [-2] void setSelection(int, int, int)
func (this *QAccessibleTextInterface) SetSelection(selectionIndex int, startOffset int, endOffset int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN24QAccessibleTextInterface12setSelectionEiii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), selectionIndex, startOffset, endOffset)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qaccessible.h:532
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [4] int cursorPosition()
func (this *QAccessibleTextInterface) CursorPosition() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK24QAccessibleTextInterface14cursorPositionEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qaccessible.h:533
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [-2] void setCursorPosition(int)
func (this *QAccessibleTextInterface) SetCursorPosition(position int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN24QAccessibleTextInterface17setCursorPositionEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), position)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qaccessible.h:536
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [8] QString text(int, int)
func (this *QAccessibleTextInterface) Text(startOffset int, endOffset int) string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK24QAccessibleTextInterface4textEii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), startOffset, endOffset)
	gopp.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtGui/qaccessible.h:537
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QString textBeforeOffset(int, QAccessible::TextBoundaryType, int *, int *)
func (this *QAccessibleTextInterface) TextBeforeOffset(offset int, boundaryType int, startOffset unsafe.Pointer /*666*/, endOffset unsafe.Pointer /*666*/) string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK24QAccessibleTextInterface16textBeforeOffsetEiN11QAccessible16TextBoundaryTypeEPiS2_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), offset, boundaryType, &startOffset, &endOffset)
	gopp.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtGui/qaccessible.h:539
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QString textAfterOffset(int, QAccessible::TextBoundaryType, int *, int *)
func (this *QAccessibleTextInterface) TextAfterOffset(offset int, boundaryType int, startOffset unsafe.Pointer /*666*/, endOffset unsafe.Pointer /*666*/) string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK24QAccessibleTextInterface15textAfterOffsetEiN11QAccessible16TextBoundaryTypeEPiS2_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), offset, boundaryType, &startOffset, &endOffset)
	gopp.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtGui/qaccessible.h:541
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QString textAtOffset(int, QAccessible::TextBoundaryType, int *, int *)
func (this *QAccessibleTextInterface) TextAtOffset(offset int, boundaryType int, startOffset unsafe.Pointer /*666*/, endOffset unsafe.Pointer /*666*/) string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK24QAccessibleTextInterface12textAtOffsetEiN11QAccessible16TextBoundaryTypeEPiS2_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), offset, boundaryType, &startOffset, &endOffset)
	gopp.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtGui/qaccessible.h:543
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [4] int characterCount()
func (this *QAccessibleTextInterface) CharacterCount() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK24QAccessibleTextInterface14characterCountEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qaccessible.h:546
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [16] QRect characterRect(int)
func (this *QAccessibleTextInterface) CharacterRect(offset int) *qtcore.QRect /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK24QAccessibleTextInterface13characterRectEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), offset)
	gopp.ErrPrint(err, rv)
	rv2 := qtcore.NewQRectFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQRect)
	return rv2
}

// /usr/include/qt/QtGui/qaccessible.h:547
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [4] int offsetAtPoint(const QPoint &)
func (this *QAccessibleTextInterface) OffsetAtPoint(point *qtcore.QPoint) int {
	var convArg0 = point.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK24QAccessibleTextInterface13offsetAtPointERK6QPoint", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qaccessible.h:549
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [-2] void scrollToSubstring(int, int)
func (this *QAccessibleTextInterface) ScrollToSubstring(startIndex int, endIndex int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN24QAccessibleTextInterface17scrollToSubstringEii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), startIndex, endIndex)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qaccessible.h:550
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [8] QString attributes(int, int *, int *)
func (this *QAccessibleTextInterface) Attributes(offset int, startOffset unsafe.Pointer /*666*/, endOffset unsafe.Pointer /*666*/) string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK24QAccessibleTextInterface10attributesEiPiS0_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), offset, &startOffset, &endOffset)
	gopp.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

//  body block end
