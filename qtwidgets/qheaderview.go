package qtwidgets

// /usr/include/qt/QtWidgets/qheaderview.h
// #include <qheaderview.h>
// #include <QtWidgets>

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

//  ext block end

//  body block begin

// void updateSection(int)
func (this *QHeaderView) InheritUpdateSection(f func(logicalIndex int) /*void*/) {
	qtrt.SetAllInheritCallback(this, "updateSection", f)
}

// void resizeSections()
func (this *QHeaderView) InheritResizeSections(f func() /*void*/) {
	qtrt.SetAllInheritCallback(this, "resizeSections", f)
}

// void sectionsInserted(const QModelIndex &, int, int)
func (this *QHeaderView) InheritSectionsInserted(f func(parent *qtcore.QModelIndex, logicalFirst int, logicalLast int) /*void*/) {
	qtrt.SetAllInheritCallback(this, "sectionsInserted", f)
}

// void sectionsAboutToBeRemoved(const QModelIndex &, int, int)
func (this *QHeaderView) InheritSectionsAboutToBeRemoved(f func(parent *qtcore.QModelIndex, logicalFirst int, logicalLast int) /*void*/) {
	qtrt.SetAllInheritCallback(this, "sectionsAboutToBeRemoved", f)
}

// void initialize()
func (this *QHeaderView) InheritInitialize(f func() /*void*/) {
	qtrt.SetAllInheritCallback(this, "initialize", f)
}

// void initializeSections()
func (this *QHeaderView) InheritInitializeSections(f func() /*void*/) {
	qtrt.SetAllInheritCallback(this, "initializeSections", f)
}

// void currentChanged(const QModelIndex &, const QModelIndex &)
func (this *QHeaderView) InheritCurrentChanged(f func(current *qtcore.QModelIndex, old *qtcore.QModelIndex) /*void*/) {
	qtrt.SetAllInheritCallback(this, "currentChanged", f)
}

// bool event(QEvent *)
func (this *QHeaderView) InheritEvent(f func(e *qtcore.QEvent /*777 QEvent **/) bool) {
	qtrt.SetAllInheritCallback(this, "event", f)
}

// void paintEvent(QPaintEvent *)
func (this *QHeaderView) InheritPaintEvent(f func(e *qtgui.QPaintEvent /*777 QPaintEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "paintEvent", f)
}

// void mousePressEvent(QMouseEvent *)
func (this *QHeaderView) InheritMousePressEvent(f func(e *qtgui.QMouseEvent /*777 QMouseEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "mousePressEvent", f)
}

// void mouseMoveEvent(QMouseEvent *)
func (this *QHeaderView) InheritMouseMoveEvent(f func(e *qtgui.QMouseEvent /*777 QMouseEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "mouseMoveEvent", f)
}

// void mouseReleaseEvent(QMouseEvent *)
func (this *QHeaderView) InheritMouseReleaseEvent(f func(e *qtgui.QMouseEvent /*777 QMouseEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "mouseReleaseEvent", f)
}

// void mouseDoubleClickEvent(QMouseEvent *)
func (this *QHeaderView) InheritMouseDoubleClickEvent(f func(e *qtgui.QMouseEvent /*777 QMouseEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "mouseDoubleClickEvent", f)
}

// bool viewportEvent(QEvent *)
func (this *QHeaderView) InheritViewportEvent(f func(e *qtcore.QEvent /*777 QEvent **/) bool) {
	qtrt.SetAllInheritCallback(this, "viewportEvent", f)
}

// void paintSection(QPainter *, const QRect &, int)
func (this *QHeaderView) InheritPaintSection(f func(painter *qtgui.QPainter /*777 QPainter **/, rect *qtcore.QRect, logicalIndex int) /*void*/) {
	qtrt.SetAllInheritCallback(this, "paintSection", f)
}

// QSize sectionSizeFromContents(int)
func (this *QHeaderView) InheritSectionSizeFromContents(f func(logicalIndex int) unsafe.Pointer) {
	qtrt.SetAllInheritCallback(this, "sectionSizeFromContents", f)
}

// int horizontalOffset()
func (this *QHeaderView) InheritHorizontalOffset(f func() int) {
	qtrt.SetAllInheritCallback(this, "horizontalOffset", f)
}

// int verticalOffset()
func (this *QHeaderView) InheritVerticalOffset(f func() int) {
	qtrt.SetAllInheritCallback(this, "verticalOffset", f)
}

// void updateGeometries()
func (this *QHeaderView) InheritUpdateGeometries(f func() /*void*/) {
	qtrt.SetAllInheritCallback(this, "updateGeometries", f)
}

// void scrollContentsBy(int, int)
func (this *QHeaderView) InheritScrollContentsBy(f func(dx int, dy int) /*void*/) {
	qtrt.SetAllInheritCallback(this, "scrollContentsBy", f)
}

// void rowsInserted(const QModelIndex &, int, int)
func (this *QHeaderView) InheritRowsInserted(f func(parent *qtcore.QModelIndex, start int, end_ int) /*void*/) {
	qtrt.SetAllInheritCallback(this, "rowsInserted", f)
}

// QRect visualRect(const QModelIndex &)
func (this *QHeaderView) InheritVisualRect(f func(index *qtcore.QModelIndex) unsafe.Pointer) {
	qtrt.SetAllInheritCallback(this, "visualRect", f)
}

// void scrollTo(const QModelIndex &, QAbstractItemView::ScrollHint)
func (this *QHeaderView) InheritScrollTo(f func(index *qtcore.QModelIndex, hint int) /*void*/) {
	qtrt.SetAllInheritCallback(this, "scrollTo", f)
}

// QModelIndex indexAt(const QPoint &)
func (this *QHeaderView) InheritIndexAt(f func(p *qtcore.QPoint) unsafe.Pointer) {
	qtrt.SetAllInheritCallback(this, "indexAt", f)
}

// bool isIndexHidden(const QModelIndex &)
func (this *QHeaderView) InheritIsIndexHidden(f func(index *qtcore.QModelIndex) bool) {
	qtrt.SetAllInheritCallback(this, "isIndexHidden", f)
}

// QModelIndex moveCursor(QAbstractItemView::CursorAction, Qt::KeyboardModifiers)
func (this *QHeaderView) InheritMoveCursor(f func(arg0 int, arg1 int) unsafe.Pointer) {
	qtrt.SetAllInheritCallback(this, "moveCursor", f)
}

// void setSelection(const QRect &, QItemSelectionModel::SelectionFlags)
func (this *QHeaderView) InheritSetSelection(f func(rect *qtcore.QRect, flags int) /*void*/) {
	qtrt.SetAllInheritCallback(this, "setSelection", f)
}

// QRegion visualRegionForSelection(const QItemSelection &)
func (this *QHeaderView) InheritVisualRegionForSelection(f func(selection *qtcore.QItemSelection) unsafe.Pointer) {
	qtrt.SetAllInheritCallback(this, "visualRegionForSelection", f)
}

// void initStyleOption(QStyleOptionHeader *)
func (this *QHeaderView) InheritInitStyleOption(f func(option *QStyleOptionHeader /*777 QStyleOptionHeader **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "initStyleOption", f)
}

/*

 */
type QHeaderView struct {
	*QAbstractItemView
}
type QHeaderView_ITF interface {
	QAbstractItemView_ITF
	QHeaderView_PTR() *QHeaderView
}

func (ptr *QHeaderView) QHeaderView_PTR() *QHeaderView { return ptr }

func (this *QHeaderView) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QAbstractItemView.GetCthis()
	}
}
func (this *QHeaderView) SetCthis(cthis unsafe.Pointer) {
	this.QAbstractItemView = NewQAbstractItemViewFromPointer(cthis)
}
func NewQHeaderViewFromPointer(cthis unsafe.Pointer) *QHeaderView {
	bcthis0 := NewQAbstractItemViewFromPointer(cthis)
	return &QHeaderView{bcthis0}
}
func (*QHeaderView) NewFromPointer(cthis unsafe.Pointer) *QHeaderView {
	return NewQHeaderViewFromPointer(cthis)
}

// /usr/include/qt/QtWidgets/qheaderview.h:55
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject() const

/*

 */
func (this *QHeaderView) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QHeaderView10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qheaderview.h:77
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QHeaderView(Qt::Orientation, QWidget *)

/*
Creates a new generic header with the given orientation and parent.
*/
func (*QHeaderView) NewForInherit(orientation int, parent QWidget_ITF /*777 QWidget **/) *QHeaderView {
	return NewQHeaderView(orientation, parent)
}
func NewQHeaderView(orientation int, parent QWidget_ITF /*777 QWidget **/) *QHeaderView {
	var convArg1 unsafe.Pointer
	if parent != nil && parent.QWidget_PTR() != nil {
		convArg1 = parent.QWidget_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QHeaderViewC2EN2Qt11OrientationEP7QWidget", qtrt.FFI_TYPE_POINTER, orientation, convArg1)
	qtrt.ErrPrint(err, rv)
	gothis := NewQHeaderViewFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QHeaderView")
	return gothis
}

// /usr/include/qt/QtWidgets/qheaderview.h:77
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QHeaderView(Qt::Orientation, QWidget *)

/*
Creates a new generic header with the given orientation and parent.
*/
func (*QHeaderView) NewForInheritp(orientation int) *QHeaderView {
	return NewQHeaderViewp(orientation)
}
func NewQHeaderViewp(orientation int) *QHeaderView {
	// arg: 1, QWidget *=Pointer, QWidget=Record, , Invalid
	var convArg1 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN11QHeaderViewC2EN2Qt11OrientationEP7QWidget", qtrt.FFI_TYPE_POINTER, orientation, convArg1)
	qtrt.ErrPrint(err, rv)
	gothis := NewQHeaderViewFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QHeaderView")
	return gothis
}

// /usr/include/qt/QtWidgets/qheaderview.h:78
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QHeaderView()

/*

 */
func DeleteQHeaderView(this *QHeaderView) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QHeaderViewD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 48)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtWidgets/qheaderview.h:80
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void setModel(QAbstractItemModel *)

/*
Reimplemented from QAbstractItemView::setModel().
*/
func (this *QHeaderView) SetModel(model qtcore.QAbstractItemModel_ITF /*777 QAbstractItemModel **/) {
	var convArg0 unsafe.Pointer
	if model != nil && model.QAbstractItemModel_PTR() != nil {
		convArg0 = model.QAbstractItemModel_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QHeaderView8setModelEP18QAbstractItemModel", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:82
// index:0
// Public Visibility=Default Availability=Available
// [4] Qt::Orientation orientation() const

/*
Returns the orientation of the header.

See also Qt::Orientation.
*/
func (this *QHeaderView) Orientation() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QHeaderView11orientationEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:83
// index:0
// Public Visibility=Default Availability=Available
// [4] int offset() const

/*
Returns the offset of the header: this is the header's left-most (or top-most for vertical headers) visible pixel.

See also setOffset().
*/
func (this *QHeaderView) Offset() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QHeaderView6offsetEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qheaderview.h:84
// index:0
// Public Visibility=Default Availability=Available
// [4] int length() const

/*
Returns the length along the orientation of the header.

See also sizeHint(), setSectionResizeMode(), and offset().
*/
func (this *QHeaderView) Length() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QHeaderView6lengthEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qheaderview.h:85
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QSize sizeHint() const

/*
Reimplemented from QWidget::sizeHint().

Returns a suitable size hint for this header.

See also sectionSizeHint().
*/
func (this *QHeaderView) SizeHint() *qtcore.QSize /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QHeaderView8sizeHintEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSize)
	return rv2
}

// /usr/include/qt/QtWidgets/qheaderview.h:86
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void setVisible(bool)

/*
Reimplemented from QWidget::setVisible().
*/
func (this *QHeaderView) SetVisible(v bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QHeaderView10setVisibleEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), v)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:87
// index:0
// Public Visibility=Default Availability=Available
// [4] int sectionSizeHint(int) const

/*
Returns a suitable size hint for the section specified by logicalIndex.

Qt::SizeHintRole

See also sizeHint(), defaultSectionSize(), minimumSectionSize(), and maximumSectionSize().
*/
func (this *QHeaderView) SectionSizeHint(logicalIndex int) int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QHeaderView15sectionSizeHintEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), logicalIndex)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qheaderview.h:89
// index:0
// Public Visibility=Default Availability=Available
// [4] int visualIndexAt(int) const

/*
Returns the visual index of the section that covers the given position in the viewport.

See also logicalIndexAt().
*/
func (this *QHeaderView) VisualIndexAt(position int) int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QHeaderView13visualIndexAtEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), position)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qheaderview.h:90
// index:0
// Public Visibility=Default Availability=Available
// [4] int logicalIndexAt(int) const

/*
Returns the section that covers the given position in the viewport.

See also visualIndexAt() and isSectionHidden().
*/
func (this *QHeaderView) LogicalIndexAt(position int) int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QHeaderView14logicalIndexAtEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), position)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qheaderview.h:92
// index:1
// Public inline Visibility=Default Availability=Available
// [4] int logicalIndexAt(int, int) const

/*
Returns the section that covers the given position in the viewport.

See also visualIndexAt() and isSectionHidden().
*/
func (this *QHeaderView) LogicalIndexAt1(x int, y int) int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QHeaderView14logicalIndexAtEii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), x, y)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qheaderview.h:93
// index:2
// Public inline Visibility=Default Availability=Available
// [4] int logicalIndexAt(const QPoint &) const

/*
Returns the section that covers the given position in the viewport.

See also visualIndexAt() and isSectionHidden().
*/
func (this *QHeaderView) LogicalIndexAt2(pos qtcore.QPoint_ITF) int {
	var convArg0 unsafe.Pointer
	if pos != nil && pos.QPoint_PTR() != nil {
		convArg0 = pos.QPoint_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QHeaderView14logicalIndexAtERK6QPoint", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qheaderview.h:95
// index:0
// Public Visibility=Default Availability=Available
// [4] int sectionSize(int) const

/*
Returns the width (or height for vertical headers) of the given logicalIndex.

See also length(), setSectionResizeMode(), and defaultSectionSize().
*/
func (this *QHeaderView) SectionSize(logicalIndex int) int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QHeaderView11sectionSizeEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), logicalIndex)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qheaderview.h:96
// index:0
// Public Visibility=Default Availability=Available
// [4] int sectionPosition(int) const

/*
Returns the section position of the given logicalIndex, or -1 if the section is hidden. The position is measured in pixels from the first visible item's top-left corner to the top-left corner of the item with logicalIndex. The measurement is along the x-axis for horizontal headers and along the y-axis for vertical headers.

See also sectionViewportPosition().
*/
func (this *QHeaderView) SectionPosition(logicalIndex int) int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QHeaderView15sectionPositionEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), logicalIndex)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qheaderview.h:97
// index:0
// Public Visibility=Default Availability=Available
// [4] int sectionViewportPosition(int) const

/*
Returns the section viewport position of the given logicalIndex.

If the section is hidden, the return value is undefined.

See also sectionPosition() and isSectionHidden().
*/
func (this *QHeaderView) SectionViewportPosition(logicalIndex int) int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QHeaderView23sectionViewportPositionEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), logicalIndex)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qheaderview.h:99
// index:0
// Public Visibility=Default Availability=Available
// [-2] void moveSection(int, int)

/*
Moves the section at visual index from to occupy visual index to.

See also sectionsMoved().
*/
func (this *QHeaderView) MoveSection(from int, to int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QHeaderView11moveSectionEii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), from, to)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:100
// index:0
// Public Visibility=Default Availability=Available
// [-2] void swapSections(int, int)

/*
Swaps the section at visual index first with the section at visual index second.

This function was introduced in  Qt 4.2.

See also moveSection().
*/
func (this *QHeaderView) SwapSections(first int, second int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QHeaderView12swapSectionsEii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), first, second)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:101
// index:0
// Public Visibility=Default Availability=Available
// [-2] void resizeSection(int, int)

/*
Resizes the section specified by logicalIndex to size measured in pixels. The size parameter must be a value larger or equal to zero. A size equal to zero is however not recommended. In that situation hideSection should be used instead.

See also sectionResized(), resizeMode(), sectionSize(), and hideSection().
*/
func (this *QHeaderView) ResizeSection(logicalIndex int, size int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QHeaderView13resizeSectionEii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), logicalIndex, size)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:102
// index:0
// Public Visibility=Default Availability=Available
// [-2] void resizeSections(QHeaderView::ResizeMode)

/*
Resizes the sections according to the given mode, ignoring the current resize mode.

See also resizeMode() and sectionResized().
*/
func (this *QHeaderView) ResizeSections(mode int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QHeaderView14resizeSectionsENS_10ResizeModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), mode)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:206
// index:1
// Protected Visibility=Default Availability=Available
// [-2] void resizeSections()

/*
Resizes the sections according to the given mode, ignoring the current resize mode.

See also resizeMode() and sectionResized().
*/
func (this *QHeaderView) ResizeSections1() {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QHeaderView14resizeSectionsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:104
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isSectionHidden(int) const

/*
Returns true if the section specified by logicalIndex is explicitly hidden from the user; otherwise returns false.

See also hideSection(), showSection(), setSectionHidden(), and hiddenSectionCount().
*/
func (this *QHeaderView) IsSectionHidden(logicalIndex int) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QHeaderView15isSectionHiddenEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), logicalIndex)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qheaderview.h:105
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setSectionHidden(int, bool)

/*
If hide is true the section specified by logicalIndex is hidden; otherwise the section is shown.

See also isSectionHidden() and hiddenSectionCount().
*/
func (this *QHeaderView) SetSectionHidden(logicalIndex int, hide bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QHeaderView16setSectionHiddenEib", qtrt.FFI_TYPE_POINTER, this.GetCthis(), logicalIndex, hide)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:106
// index:0
// Public Visibility=Default Availability=Available
// [4] int hiddenSectionCount() const

/*
Returns the number of sections in the header that has been hidden.

This function was introduced in  Qt 4.1.

See also setSectionHidden() and isSectionHidden().
*/
func (this *QHeaderView) HiddenSectionCount() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QHeaderView18hiddenSectionCountEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qheaderview.h:108
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void hideSection(int)

/*
Hides the section specified by logicalIndex.

See also showSection(), isSectionHidden(), hiddenSectionCount(), and setSectionHidden().
*/
func (this *QHeaderView) HideSection(logicalIndex int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QHeaderView11hideSectionEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), logicalIndex)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:109
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void showSection(int)

/*
Shows the section specified by logicalIndex.

See also hideSection(), isSectionHidden(), hiddenSectionCount(), and setSectionHidden().
*/
func (this *QHeaderView) ShowSection(logicalIndex int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QHeaderView11showSectionEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), logicalIndex)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:111
// index:0
// Public Visibility=Default Availability=Available
// [4] int count() const

/*
Returns the number of sections in the header.

See also sectionCountChanged() and length().
*/
func (this *QHeaderView) Count() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QHeaderView5countEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qheaderview.h:112
// index:0
// Public Visibility=Default Availability=Available
// [4] int visualIndex(int) const

/*
Returns the visual index position of the section specified by the given logicalIndex, or -1 otherwise.

Hidden sections still have valid visual indexes.

See also logicalIndex().
*/
func (this *QHeaderView) VisualIndex(logicalIndex int) int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QHeaderView11visualIndexEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), logicalIndex)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qheaderview.h:113
// index:0
// Public Visibility=Default Availability=Available
// [4] int logicalIndex(int) const

/*
Returns the logicalIndex for the section at the given visualIndex position, or -1 if visualIndex < 0 or visualIndex >= QHeaderView::count().

Note that the visualIndex is not affected by hidden sections.

See also visualIndex() and sectionPosition().
*/
func (this *QHeaderView) LogicalIndex(visualIndex int) int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QHeaderView12logicalIndexEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), visualIndex)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qheaderview.h:115
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setSectionsMovable(bool)

/*
If movable is true, the header may be moved by the user; otherwise it is fixed in place.

This function was introduced in  Qt 5.0.

See also sectionsMovable() and sectionMoved().
*/
func (this *QHeaderView) SetSectionsMovable(movable bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QHeaderView18setSectionsMovableEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), movable)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:116
// index:0
// Public Visibility=Default Availability=Available
// [1] bool sectionsMovable() const

/*
Returns true if the header can be moved by the user; otherwise returns false.

This function was introduced in  Qt 5.0.

See also setSectionsMovable().
*/
func (this *QHeaderView) SectionsMovable() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QHeaderView15sectionsMovableEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qheaderview.h:122
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setSectionsClickable(bool)

/*
If clickable is true, the header will respond to single clicks.

This function was introduced in  Qt 5.0.

See also sectionsClickable(), sectionClicked(), sectionPressed(), and setSortIndicatorShown().
*/
func (this *QHeaderView) SetSectionsClickable(clickable bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QHeaderView20setSectionsClickableEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), clickable)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:123
// index:0
// Public Visibility=Default Availability=Available
// [1] bool sectionsClickable() const

/*
Returns true if the header is clickable; otherwise returns false. A clickable header could be set up to allow the user to change the representation of the data in the view related to the header.

This function was introduced in  Qt 5.0.

See also setSectionsClickable().
*/
func (this *QHeaderView) SectionsClickable() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QHeaderView17sectionsClickableEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qheaderview.h:129
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setHighlightSections(bool)

/*

 */
func (this *QHeaderView) SetHighlightSections(highlight bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QHeaderView20setHighlightSectionsEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), highlight)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:130
// index:0
// Public Visibility=Default Availability=Available
// [1] bool highlightSections() const

/*

 */
func (this *QHeaderView) HighlightSections() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QHeaderView17highlightSectionsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qheaderview.h:132
// index:0
// Public Visibility=Default Availability=Available
// [4] QHeaderView::ResizeMode sectionResizeMode(int) const

/*
Returns the resize mode that applies to the section specified by the given logicalIndex.

This function was introduced in  Qt 5.0.

See also setSectionResizeMode().
*/
func (this *QHeaderView) SectionResizeMode(logicalIndex int) int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QHeaderView17sectionResizeModeEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), logicalIndex)
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:133
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setSectionResizeMode(QHeaderView::ResizeMode)

/*
Sets the constraints on how the header can be resized to those described by the given mode.

This function was introduced in  Qt 5.0.

See also sectionResizeMode(), resizeMode(), length(), and sectionResized().
*/
func (this *QHeaderView) SetSectionResizeMode(mode int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QHeaderView20setSectionResizeModeENS_10ResizeModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), mode)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:134
// index:1
// Public Visibility=Default Availability=Available
// [-2] void setSectionResizeMode(int, QHeaderView::ResizeMode)

/*
Sets the constraints on how the header can be resized to those described by the given mode.

This function was introduced in  Qt 5.0.

See also sectionResizeMode(), resizeMode(), length(), and sectionResized().
*/
func (this *QHeaderView) SetSectionResizeMode1(logicalIndex int, mode int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QHeaderView20setSectionResizeModeEiNS_10ResizeModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), logicalIndex, mode)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:136
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setResizeContentsPrecision(int)

/*
Sets how precise QHeaderView should calculate the size when ResizeToContents is used. A low value will provide a less accurate but fast auto resize while a higher value will provide a more accurate resize that however can be slow.

The number precision specifies how many sections that should be consider when calculating the preferred size.

The default value is 1000 meaning that a horizontal column with auto-resize will look at maximum 1000 rows on calculating when doing an auto resize.

Special value 0 means that it will look at only the visible area. Special value -1 will imply looking at all elements.

This value is used in QTableView::sizeHintForColumn(), QTableView::sizeHintForRow() and QTreeView::sizeHintForColumn(). Reimplementing these functions can make this function not having an effect.

This function was introduced in  Qt 5.2.

See also resizeContentsPrecision(), setSectionResizeMode(), resizeSections(), QTableView::sizeHintForColumn(), QTableView::sizeHintForRow(), and QTreeView::sizeHintForColumn().
*/
func (this *QHeaderView) SetResizeContentsPrecision(precision int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QHeaderView26setResizeContentsPrecisionEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), precision)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:137
// index:0
// Public Visibility=Default Availability=Available
// [4] int resizeContentsPrecision() const

/*
Returns how precise QHeaderView will calculate on ResizeToContents.

This function was introduced in  Qt 5.2.

See also setResizeContentsPrecision() and setSectionResizeMode().
*/
func (this *QHeaderView) ResizeContentsPrecision() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QHeaderView23resizeContentsPrecisionEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qheaderview.h:148
// index:0
// Public Visibility=Default Availability=Available
// [4] int stretchSectionCount() const

/*
Returns the number of sections that are set to resize mode stretch. In views, this can be used to see if the headerview needs to resize the sections when the view's geometry changes.

This function was introduced in  Qt 4.1.

See also stretchLastSection and resizeMode().
*/
func (this *QHeaderView) StretchSectionCount() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QHeaderView19stretchSectionCountEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qheaderview.h:150
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setSortIndicatorShown(bool)

/*

 */
func (this *QHeaderView) SetSortIndicatorShown(show bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QHeaderView21setSortIndicatorShownEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), show)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:151
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isSortIndicatorShown() const

/*

 */
func (this *QHeaderView) IsSortIndicatorShown() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QHeaderView20isSortIndicatorShownEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qheaderview.h:153
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setSortIndicator(int, Qt::SortOrder)

/*
Sets the sort indicator for the section specified by the given logicalIndex in the direction specified by order, and removes the sort indicator from any other section that was showing it.

logicalIndex may be -1, in which case no sort indicator will be shown and the model will return to its natural, unsorted order. Note that not all models support this and may even crash in this case.

See also sortIndicatorSection() and sortIndicatorOrder().
*/
func (this *QHeaderView) SetSortIndicator(logicalIndex int, order int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QHeaderView16setSortIndicatorEiN2Qt9SortOrderE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), logicalIndex, order)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:154
// index:0
// Public Visibility=Default Availability=Available
// [4] int sortIndicatorSection() const

/*
Returns the logical index of the section that has a sort indicator. By default this is section 0.

See also setSortIndicator(), sortIndicatorOrder(), and setSortIndicatorShown().
*/
func (this *QHeaderView) SortIndicatorSection() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QHeaderView20sortIndicatorSectionEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qheaderview.h:155
// index:0
// Public Visibility=Default Availability=Available
// [4] Qt::SortOrder sortIndicatorOrder() const

/*
Returns the order for the sort indicator. If no section has a sort indicator the return value of this function is undefined.

See also setSortIndicator() and sortIndicatorSection().
*/
func (this *QHeaderView) SortIndicatorOrder() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QHeaderView18sortIndicatorOrderEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:157
// index:0
// Public Visibility=Default Availability=Available
// [1] bool stretchLastSection() const

/*

 */
func (this *QHeaderView) StretchLastSection() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QHeaderView18stretchLastSectionEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qheaderview.h:158
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setStretchLastSection(bool)

/*

 */
func (this *QHeaderView) SetStretchLastSection(stretch bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QHeaderView21setStretchLastSectionEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), stretch)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:160
// index:0
// Public Visibility=Default Availability=Available
// [1] bool cascadingSectionResizes() const

/*

 */
func (this *QHeaderView) CascadingSectionResizes() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QHeaderView23cascadingSectionResizesEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qheaderview.h:161
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setCascadingSectionResizes(bool)

/*

 */
func (this *QHeaderView) SetCascadingSectionResizes(enable bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QHeaderView26setCascadingSectionResizesEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), enable)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:163
// index:0
// Public Visibility=Default Availability=Available
// [4] int defaultSectionSize() const

/*

 */
func (this *QHeaderView) DefaultSectionSize() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QHeaderView18defaultSectionSizeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qheaderview.h:164
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setDefaultSectionSize(int)

/*

 */
func (this *QHeaderView) SetDefaultSectionSize(size int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QHeaderView21setDefaultSectionSizeEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), size)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:165
// index:0
// Public Visibility=Default Availability=Available
// [-2] void resetDefaultSectionSize()

/*

 */
func (this *QHeaderView) ResetDefaultSectionSize() {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QHeaderView23resetDefaultSectionSizeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:167
// index:0
// Public Visibility=Default Availability=Available
// [4] int minimumSectionSize() const

/*

 */
func (this *QHeaderView) MinimumSectionSize() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QHeaderView18minimumSectionSizeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qheaderview.h:168
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setMinimumSectionSize(int)

/*

 */
func (this *QHeaderView) SetMinimumSectionSize(size int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QHeaderView21setMinimumSectionSizeEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), size)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:169
// index:0
// Public Visibility=Default Availability=Available
// [4] int maximumSectionSize() const

/*

 */
func (this *QHeaderView) MaximumSectionSize() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QHeaderView18maximumSectionSizeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qheaderview.h:170
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setMaximumSectionSize(int)

/*

 */
func (this *QHeaderView) SetMaximumSectionSize(size int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QHeaderView21setMaximumSectionSizeEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), size)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:172
// index:0
// Public Visibility=Default Availability=Available
// [4] Qt::Alignment defaultAlignment() const

/*

 */
func (this *QHeaderView) DefaultAlignment() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QHeaderView16defaultAlignmentEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:173
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setDefaultAlignment(Qt::Alignment)

/*

 */
func (this *QHeaderView) SetDefaultAlignment(alignment int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QHeaderView19setDefaultAlignmentE6QFlagsIN2Qt13AlignmentFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), alignment)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:175
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void doItemsLayout()

/*

 */
func (this *QHeaderView) DoItemsLayout() {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QHeaderView13doItemsLayoutEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:176
// index:0
// Public Visibility=Default Availability=Available
// [1] bool sectionsMoved() const

/*
Returns true if sections in the header has been moved; otherwise returns false;

See also moveSection().
*/
func (this *QHeaderView) SectionsMoved() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QHeaderView13sectionsMovedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qheaderview.h:177
// index:0
// Public Visibility=Default Availability=Available
// [1] bool sectionsHidden() const

/*
Returns true if sections in the header has been hidden; otherwise returns false;

This function was introduced in  Qt 4.1.

See also setSectionHidden().
*/
func (this *QHeaderView) SectionsHidden() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QHeaderView14sectionsHiddenEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qheaderview.h:180
// index:0
// Public Visibility=Default Availability=Available
// [8] QByteArray saveState() const

/*
Saves the current state of this header view.

To restore the saved state, pass the return value to restoreState().

This function was introduced in  Qt 4.3.

See also restoreState().
*/
func (this *QHeaderView) SaveState() *qtcore.QByteArray /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QHeaderView9saveStateEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQByteArray)
	return rv2
}

// /usr/include/qt/QtWidgets/qheaderview.h:181
// index:0
// Public Visibility=Default Availability=Available
// [1] bool restoreState(const QByteArray &)

/*
Restores the state of this header view. This function returns true if the state was restored; otherwise returns false.

This function was introduced in  Qt 4.3.

See also saveState().
*/
func (this *QHeaderView) RestoreState(state qtcore.QByteArray_ITF) bool {
	var convArg0 unsafe.Pointer
	if state != nil && state.QByteArray_PTR() != nil {
		convArg0 = state.QByteArray_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QHeaderView12restoreStateERK10QByteArray", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qheaderview.h:184
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void reset()

/*
Reimplemented from QAbstractItemView::reset().
*/
func (this *QHeaderView) Reset() {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QHeaderView5resetEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:187
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setOffset(int)

/*
Sets the header's offset to offset.

See also offset() and length().
*/
func (this *QHeaderView) SetOffset(offset int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QHeaderView9setOffsetEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), offset)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:188
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setOffsetToSectionPosition(int)

/*
Sets the offset to the start of the section at the given visualSectionNumber. visualSectionNumber is the actual visible section when hiddenSections are not considered. That is not always the same as visualIndex().

This function was introduced in  Qt 4.2.

See also setOffset() and sectionPosition().
*/
func (this *QHeaderView) SetOffsetToSectionPosition(visualIndex int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QHeaderView26setOffsetToSectionPositionEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), visualIndex)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:189
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setOffsetToLastSection()

/*
Sets the offset to make the last section visible.

This function was introduced in  Qt 4.2.

See also setOffset(), sectionPosition(), and setOffsetToSectionPosition().
*/
func (this *QHeaderView) SetOffsetToLastSection() {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QHeaderView22setOffsetToLastSectionEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:190
// index:0
// Public Visibility=Default Availability=Available
// [-2] void headerDataChanged(Qt::Orientation, int, int)

/*
Updates the changed header sections with the given orientation, from logicalFirst to logicalLast inclusive.
*/
func (this *QHeaderView) HeaderDataChanged(orientation int, logicalFirst int, logicalLast int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QHeaderView17headerDataChangedEN2Qt11OrientationEii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), orientation, logicalFirst, logicalLast)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:193
// index:0
// Public Visibility=Default Availability=Available
// [-2] void sectionMoved(int, int, int)

/*
This signal is emitted when a section is moved. The section's logical index is specified by logicalIndex, the old index by oldVisualIndex, and the new index position by newVisualIndex.

See also moveSection().
*/
func (this *QHeaderView) SectionMoved(logicalIndex int, oldVisualIndex int, newVisualIndex int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QHeaderView12sectionMovedEiii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), logicalIndex, oldVisualIndex, newVisualIndex)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:194
// index:0
// Public Visibility=Default Availability=Available
// [-2] void sectionResized(int, int, int)

/*
This signal is emitted when a section is resized. The section's logical number is specified by logicalIndex, the old size by oldSize, and the new size by newSize.

See also resizeSection().
*/
func (this *QHeaderView) SectionResized(logicalIndex int, oldSize int, newSize int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QHeaderView14sectionResizedEiii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), logicalIndex, oldSize, newSize)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:195
// index:0
// Public Visibility=Default Availability=Available
// [-2] void sectionPressed(int)

/*
This signal is emitted when a section is pressed. The section's logical index is specified by logicalIndex.

See also setSectionsClickable().
*/
func (this *QHeaderView) SectionPressed(logicalIndex int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QHeaderView14sectionPressedEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), logicalIndex)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:196
// index:0
// Public Visibility=Default Availability=Available
// [-2] void sectionClicked(int)

/*
This signal is emitted when a section is clicked. The section's logical index is specified by logicalIndex.

Note that the sectionPressed signal will also be emitted.

See also setSectionsClickable() and sectionPressed().
*/
func (this *QHeaderView) SectionClicked(logicalIndex int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QHeaderView14sectionClickedEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), logicalIndex)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:197
// index:0
// Public Visibility=Default Availability=Available
// [-2] void sectionEntered(int)

/*
This signal is emitted when the cursor moves over the section and the left mouse button is pressed. The section's logical index is specified by logicalIndex.

This function was introduced in  Qt 4.3.

See also setSectionsClickable() and sectionPressed().
*/
func (this *QHeaderView) SectionEntered(logicalIndex int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QHeaderView14sectionEnteredEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), logicalIndex)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:198
// index:0
// Public Visibility=Default Availability=Available
// [-2] void sectionDoubleClicked(int)

/*
This signal is emitted when a section is double-clicked. The section's logical index is specified by logicalIndex.

See also setSectionsClickable().
*/
func (this *QHeaderView) SectionDoubleClicked(logicalIndex int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QHeaderView20sectionDoubleClickedEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), logicalIndex)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:199
// index:0
// Public Visibility=Default Availability=Available
// [-2] void sectionCountChanged(int, int)

/*
This signal is emitted when the number of sections changes, i.e., when sections are added or deleted. The original count is specified by oldCount, and the new count by newCount.

See also count(), length(), and headerDataChanged().
*/
func (this *QHeaderView) SectionCountChanged(oldCount int, newCount int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QHeaderView19sectionCountChangedEii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), oldCount, newCount)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:200
// index:0
// Public Visibility=Default Availability=Available
// [-2] void sectionHandleDoubleClicked(int)

/*
This signal is emitted when a section is double-clicked. The section's logical index is specified by logicalIndex.

See also setSectionsClickable().
*/
func (this *QHeaderView) SectionHandleDoubleClicked(logicalIndex int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QHeaderView26sectionHandleDoubleClickedEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), logicalIndex)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:201
// index:0
// Public Visibility=Default Availability=Available
// [-2] void geometriesChanged()

/*
This signal is emitted when the header's geometries have changed.

This function was introduced in  Qt 4.2.
*/
func (this *QHeaderView) GeometriesChanged() {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QHeaderView17geometriesChangedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:202
// index:0
// Public Visibility=Default Availability=Available
// [-2] void sortIndicatorChanged(int, Qt::SortOrder)

/*
This signal is emitted when the section containing the sort indicator or the order indicated is changed. The section's logical index is specified by logicalIndex and the sort order is specified by order.

This function was introduced in  Qt 4.3.

See also setSortIndicator().
*/
func (this *QHeaderView) SortIndicatorChanged(logicalIndex int, order int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QHeaderView20sortIndicatorChangedEiN2Qt9SortOrderE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), logicalIndex, order)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:205
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void updateSection(int)

/*

 */
func (this *QHeaderView) UpdateSection(logicalIndex int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QHeaderView13updateSectionEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), logicalIndex)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:207
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void sectionsInserted(const QModelIndex &, int, int)

/*
This slot is called when sections are inserted into the parent. logicalFirst and logicalLast indices signify where the new sections were inserted.

If only one section is inserted, logicalFirst and logicalLast will be the same.
*/
func (this *QHeaderView) SectionsInserted(parent qtcore.QModelIndex_ITF, logicalFirst int, logicalLast int) {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QModelIndex_PTR() != nil {
		convArg0 = parent.QModelIndex_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QHeaderView16sectionsInsertedERK11QModelIndexii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, logicalFirst, logicalLast)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:208
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void sectionsAboutToBeRemoved(const QModelIndex &, int, int)

/*
This slot is called when sections are removed from the parent. logicalFirst and logicalLast signify where the sections were removed.

If only one section is removed, logicalFirst and logicalLast will be the same.
*/
func (this *QHeaderView) SectionsAboutToBeRemoved(parent qtcore.QModelIndex_ITF, logicalFirst int, logicalLast int) {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QModelIndex_PTR() != nil {
		convArg0 = parent.QModelIndex_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QHeaderView24sectionsAboutToBeRemovedERK11QModelIndexii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, logicalFirst, logicalLast)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:212
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void initialize()

/*

 */
func (this *QHeaderView) Initialize() {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QHeaderView10initializeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:214
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void initializeSections()

/*

 */
func (this *QHeaderView) InitializeSections() {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QHeaderView18initializeSectionsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:215
// index:1
// Protected Visibility=Default Availability=Available
// [-2] void initializeSections(int, int)

/*

 */
func (this *QHeaderView) InitializeSections1(start int, end_ int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QHeaderView18initializeSectionsEii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), start, end_)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:216
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void currentChanged(const QModelIndex &, const QModelIndex &)

/*
Reimplemented from QAbstractItemView::currentChanged().
*/
func (this *QHeaderView) CurrentChanged(current qtcore.QModelIndex_ITF, old qtcore.QModelIndex_ITF) {
	var convArg0 unsafe.Pointer
	if current != nil && current.QModelIndex_PTR() != nil {
		convArg0 = current.QModelIndex_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if old != nil && old.QModelIndex_PTR() != nil {
		convArg1 = old.QModelIndex_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QHeaderView14currentChangedERK11QModelIndexS2_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:218
// index:0
// Protected virtual Visibility=Default Availability=Available
// [1] bool event(QEvent *)

/*
Reimplemented from QObject::event().
*/
func (this *QHeaderView) Event(e qtcore.QEvent_ITF /*777 QEvent **/) bool {
	var convArg0 unsafe.Pointer
	if e != nil && e.QEvent_PTR() != nil {
		convArg0 = e.QEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QHeaderView5eventEP6QEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qheaderview.h:219
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void paintEvent(QPaintEvent *)

/*
Reimplemented from QWidget::paintEvent().
*/
func (this *QHeaderView) PaintEvent(e qtgui.QPaintEvent_ITF /*777 QPaintEvent **/) {
	var convArg0 unsafe.Pointer
	if e != nil && e.QPaintEvent_PTR() != nil {
		convArg0 = e.QPaintEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QHeaderView10paintEventEP11QPaintEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:220
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void mousePressEvent(QMouseEvent *)

/*
Reimplemented from QWidget::mousePressEvent().
*/
func (this *QHeaderView) MousePressEvent(e qtgui.QMouseEvent_ITF /*777 QMouseEvent **/) {
	var convArg0 unsafe.Pointer
	if e != nil && e.QMouseEvent_PTR() != nil {
		convArg0 = e.QMouseEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QHeaderView15mousePressEventEP11QMouseEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:221
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void mouseMoveEvent(QMouseEvent *)

/*
Reimplemented from QWidget::mouseMoveEvent().
*/
func (this *QHeaderView) MouseMoveEvent(e qtgui.QMouseEvent_ITF /*777 QMouseEvent **/) {
	var convArg0 unsafe.Pointer
	if e != nil && e.QMouseEvent_PTR() != nil {
		convArg0 = e.QMouseEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QHeaderView14mouseMoveEventEP11QMouseEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:222
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void mouseReleaseEvent(QMouseEvent *)

/*
Reimplemented from QWidget::mouseReleaseEvent().
*/
func (this *QHeaderView) MouseReleaseEvent(e qtgui.QMouseEvent_ITF /*777 QMouseEvent **/) {
	var convArg0 unsafe.Pointer
	if e != nil && e.QMouseEvent_PTR() != nil {
		convArg0 = e.QMouseEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QHeaderView17mouseReleaseEventEP11QMouseEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:223
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void mouseDoubleClickEvent(QMouseEvent *)

/*
Reimplemented from QWidget::mouseDoubleClickEvent().
*/
func (this *QHeaderView) MouseDoubleClickEvent(e qtgui.QMouseEvent_ITF /*777 QMouseEvent **/) {
	var convArg0 unsafe.Pointer
	if e != nil && e.QMouseEvent_PTR() != nil {
		convArg0 = e.QMouseEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QHeaderView21mouseDoubleClickEventEP11QMouseEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:224
// index:0
// Protected virtual Visibility=Default Availability=Available
// [1] bool viewportEvent(QEvent *)

/*
Reimplemented from QAbstractScrollArea::viewportEvent().
*/
func (this *QHeaderView) ViewportEvent(e qtcore.QEvent_ITF /*777 QEvent **/) bool {
	var convArg0 unsafe.Pointer
	if e != nil && e.QEvent_PTR() != nil {
		convArg0 = e.QEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QHeaderView13viewportEventEP6QEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qheaderview.h:226
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void paintSection(QPainter *, const QRect &, int) const

/*
Paints the section specified by the given logicalIndex, using the given painter and rect.

Normally, you do not have to call this function.
*/
func (this *QHeaderView) PaintSection(painter qtgui.QPainter_ITF /*777 QPainter **/, rect qtcore.QRect_ITF, logicalIndex int) {
	var convArg0 unsafe.Pointer
	if painter != nil && painter.QPainter_PTR() != nil {
		convArg0 = painter.QPainter_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if rect != nil && rect.QRect_PTR() != nil {
		convArg1 = rect.QRect_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QHeaderView12paintSectionEP8QPainterRK5QRecti", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, logicalIndex)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:227
// index:0
// Protected virtual Visibility=Default Availability=Available
// [8] QSize sectionSizeFromContents(int) const

/*
Returns the size of the contents of the section specified by the given logicalIndex.

See also defaultSectionSize().
*/
func (this *QHeaderView) SectionSizeFromContents(logicalIndex int) *qtcore.QSize /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QHeaderView23sectionSizeFromContentsEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), logicalIndex)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSize)
	return rv2
}

// /usr/include/qt/QtWidgets/qheaderview.h:229
// index:0
// Protected virtual Visibility=Default Availability=Available
// [4] int horizontalOffset() const

/*
Reimplemented from QAbstractItemView::horizontalOffset().

Returns the horizontal offset of the header. This is 0 for vertical headers.

See also offset().
*/
func (this *QHeaderView) HorizontalOffset() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QHeaderView16horizontalOffsetEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qheaderview.h:230
// index:0
// Protected virtual Visibility=Default Availability=Available
// [4] int verticalOffset() const

/*
Reimplemented from QAbstractItemView::verticalOffset().

Returns the vertical offset of the header. This is 0 for horizontal headers.

See also offset().
*/
func (this *QHeaderView) VerticalOffset() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QHeaderView14verticalOffsetEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qheaderview.h:231
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void updateGeometries()

/*

 */
func (this *QHeaderView) UpdateGeometries() {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QHeaderView16updateGeometriesEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:232
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void scrollContentsBy(int, int)

/*

 */
func (this *QHeaderView) ScrollContentsBy(dx int, dy int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QHeaderView16scrollContentsByEii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), dx, dy)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:235
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void rowsInserted(const QModelIndex &, int, int)

/*

 */
func (this *QHeaderView) RowsInserted(parent qtcore.QModelIndex_ITF, start int, end_ int) {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QModelIndex_PTR() != nil {
		convArg0 = parent.QModelIndex_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QHeaderView12rowsInsertedERK11QModelIndexii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, start, end_)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:237
// index:0
// Protected virtual Visibility=Default Availability=Available
// [16] QRect visualRect(const QModelIndex &) const

/*

 */
func (this *QHeaderView) VisualRect(index qtcore.QModelIndex_ITF) *qtcore.QRect /*123*/ {
	var convArg0 unsafe.Pointer
	if index != nil && index.QModelIndex_PTR() != nil {
		convArg0 = index.QModelIndex_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QHeaderView10visualRectERK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQRectFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQRect)
	return rv2
}

// /usr/include/qt/QtWidgets/qheaderview.h:238
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void scrollTo(const QModelIndex &, QAbstractItemView::ScrollHint)

/*

 */
func (this *QHeaderView) ScrollTo(index qtcore.QModelIndex_ITF, hint int) {
	var convArg0 unsafe.Pointer
	if index != nil && index.QModelIndex_PTR() != nil {
		convArg0 = index.QModelIndex_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QHeaderView8scrollToERK11QModelIndexN17QAbstractItemView10ScrollHintE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, hint)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:240
// index:0
// Protected virtual Visibility=Default Availability=Available
// [24] QModelIndex indexAt(const QPoint &) const

/*

 */
func (this *QHeaderView) IndexAt(p qtcore.QPoint_ITF) *qtcore.QModelIndex /*123*/ {
	var convArg0 unsafe.Pointer
	if p != nil && p.QPoint_PTR() != nil {
		convArg0 = p.QPoint_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QHeaderView7indexAtERK6QPoint", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQModelIndexFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQModelIndex)
	return rv2
}

// /usr/include/qt/QtWidgets/qheaderview.h:241
// index:0
// Protected virtual Visibility=Default Availability=Available
// [1] bool isIndexHidden(const QModelIndex &) const

/*

 */
func (this *QHeaderView) IsIndexHidden(index qtcore.QModelIndex_ITF) bool {
	var convArg0 unsafe.Pointer
	if index != nil && index.QModelIndex_PTR() != nil {
		convArg0 = index.QModelIndex_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QHeaderView13isIndexHiddenERK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qheaderview.h:243
// index:0
// Protected virtual Visibility=Default Availability=Available
// [24] QModelIndex moveCursor(QAbstractItemView::CursorAction, Qt::KeyboardModifiers)

/*

 */
func (this *QHeaderView) MoveCursor(arg0 int, arg1 int) *qtcore.QModelIndex /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QHeaderView10moveCursorEN17QAbstractItemView12CursorActionE6QFlagsIN2Qt16KeyboardModifierEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0, arg1)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQModelIndexFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQModelIndex)
	return rv2
}

// /usr/include/qt/QtWidgets/qheaderview.h:244
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void setSelection(const QRect &, QItemSelectionModel::SelectionFlags)

/*
Reimplemented from QAbstractItemView::setSelection().

Selects the items in the given rect according to the specified flags.

The base class implementation does nothing.
*/
func (this *QHeaderView) SetSelection(rect qtcore.QRect_ITF, flags int) {
	var convArg0 unsafe.Pointer
	if rect != nil && rect.QRect_PTR() != nil {
		convArg0 = rect.QRect_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QHeaderView12setSelectionERK5QRect6QFlagsIN19QItemSelectionModel13SelectionFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, flags)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:245
// index:0
// Protected virtual Visibility=Default Availability=Available
// [8] QRegion visualRegionForSelection(const QItemSelection &) const

/*

 */
func (this *QHeaderView) VisualRegionForSelection(selection qtcore.QItemSelection_ITF) *qtgui.QRegion /*123*/ {
	var convArg0 unsafe.Pointer
	if selection != nil && selection.QItemSelection_PTR() != nil {
		convArg0 = selection.QItemSelection_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QHeaderView24visualRegionForSelectionERK14QItemSelection", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := qtgui.NewQRegionFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtgui.DeleteQRegion)
	return rv2
}

// /usr/include/qt/QtWidgets/qheaderview.h:246
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void initStyleOption(QStyleOptionHeader *) const

/*
Initialize option with the values from this QHeaderView. This method is useful for subclasses when they need a QStyleOptionHeader, but do not want to fill in all the information themselves.

See also QStyleOption::initFrom().
*/
func (this *QHeaderView) InitStyleOption(option QStyleOptionHeader_ITF /*777 QStyleOptionHeader **/) {
	var convArg0 unsafe.Pointer
	if option != nil && option.QStyleOptionHeader_PTR() != nil {
		convArg0 = option.QStyleOptionHeader_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QHeaderView15initStyleOptionEP18QStyleOptionHeader", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

/*
The resize mode specifies the behavior of the header sections. It can be set on the entire header view or on individual sections using setSectionResizeMode().



The following values are obsolete:

QHeaderView::CustomFixedUse Fixed instead.


See also setResizeMode(), setSectionResizeMode(), stretchLastSection, and minimumSectionSize.

*/
type QHeaderView__ResizeMode = int

// The user can resize the section. The section can also be resized programmatically using resizeSection(). The section size defaults to defaultSectionSize. (See also cascadingSectionResizes.)
const QHeaderView__Interactive QHeaderView__ResizeMode = 0

// QHeaderView will automatically resize the section to fill the available space. The size cannot be changed by the user or programmatically.
const QHeaderView__Stretch QHeaderView__ResizeMode = 1

// The user cannot resize the section. The section can only be resized programmatically using resizeSection(). The section size defaults to defaultSectionSize.
const QHeaderView__Fixed QHeaderView__ResizeMode = 2

//
const QHeaderView__ResizeToContents QHeaderView__ResizeMode = 3

//
const QHeaderView__Custom QHeaderView__ResizeMode = 2

func (this *QHeaderView) ResizeModeItemName(val int) string {
	return qtrt.GetClassEnumItemName(this, val)
}
func QHeaderView_ResizeModeItemName(val int) string {
	var nilthis *QHeaderView
	return nilthis.ResizeModeItemName(val)
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
