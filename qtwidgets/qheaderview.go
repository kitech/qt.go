package qtwidgets

// /usr/include/qt/QtWidgets/qheaderview.h
// #include <qheaderview.h>
// #include <QtWidgets>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 27
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "gopp"
import "qt.go/cffiqt"
import "qt.go/qtrt"
import "qt.go/qtcore"
import "qt.go/qtgui"

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
		ffiqt.KeepMe()
	}
	if false {
		gopp.KeepMe()
	}
	if false {
		qtcore.KeepMe()
	}
	if false {
		qtgui.KeepMe()
	}
}

//  ext block end

//  body block begin
// void updateSection(int)
func (this *QHeaderView) InheritUpdateSection(f func(logicalIndex int)) {
	ffiqt.SetAllInheritCallback(this, "updateSection", f)
}

// void resizeSections()
func (this *QHeaderView) InheritResizeSections(f func()) {
	ffiqt.SetAllInheritCallback(this, "resizeSections", f)
}

// void sectionsInserted(const class QModelIndex &, int, int)
func (this *QHeaderView) InheritSectionsInserted(f func(parent *qtcore.QModelIndex, logicalFirst int, logicalLast int)) {
	ffiqt.SetAllInheritCallback(this, "sectionsInserted", f)
}

// void sectionsAboutToBeRemoved(const class QModelIndex &, int, int)
func (this *QHeaderView) InheritSectionsAboutToBeRemoved(f func(parent *qtcore.QModelIndex, logicalFirst int, logicalLast int)) {
	ffiqt.SetAllInheritCallback(this, "sectionsAboutToBeRemoved", f)
}

// void initialize()
func (this *QHeaderView) InheritInitialize(f func()) {
	ffiqt.SetAllInheritCallback(this, "initialize", f)
}

// void initializeSections()
func (this *QHeaderView) InheritInitializeSections(f func()) {
	ffiqt.SetAllInheritCallback(this, "initializeSections", f)
}

// void currentChanged(const class QModelIndex &, const class QModelIndex &)
func (this *QHeaderView) InheritCurrentChanged(f func(current *qtcore.QModelIndex, old *qtcore.QModelIndex)) {
	ffiqt.SetAllInheritCallback(this, "currentChanged", f)
}

// bool event(class QEvent *)
func (this *QHeaderView) InheritEvent(f func(e *qtcore.QEvent /*777 QEvent **/) bool) {
	ffiqt.SetAllInheritCallback(this, "event", f)
}

// void paintEvent(class QPaintEvent *)
func (this *QHeaderView) InheritPaintEvent(f func(e *qtgui.QPaintEvent /*777 QPaintEvent **/)) {
	ffiqt.SetAllInheritCallback(this, "paintEvent", f)
}

// void mousePressEvent(class QMouseEvent *)
func (this *QHeaderView) InheritMousePressEvent(f func(e *qtgui.QMouseEvent /*777 QMouseEvent **/)) {
	ffiqt.SetAllInheritCallback(this, "mousePressEvent", f)
}

// void mouseMoveEvent(class QMouseEvent *)
func (this *QHeaderView) InheritMouseMoveEvent(f func(e *qtgui.QMouseEvent /*777 QMouseEvent **/)) {
	ffiqt.SetAllInheritCallback(this, "mouseMoveEvent", f)
}

// void mouseReleaseEvent(class QMouseEvent *)
func (this *QHeaderView) InheritMouseReleaseEvent(f func(e *qtgui.QMouseEvent /*777 QMouseEvent **/)) {
	ffiqt.SetAllInheritCallback(this, "mouseReleaseEvent", f)
}

// void mouseDoubleClickEvent(class QMouseEvent *)
func (this *QHeaderView) InheritMouseDoubleClickEvent(f func(e *qtgui.QMouseEvent /*777 QMouseEvent **/)) {
	ffiqt.SetAllInheritCallback(this, "mouseDoubleClickEvent", f)
}

// bool viewportEvent(class QEvent *)
func (this *QHeaderView) InheritViewportEvent(f func(e *qtcore.QEvent /*777 QEvent **/) bool) {
	ffiqt.SetAllInheritCallback(this, "viewportEvent", f)
}

// void paintSection(class QPainter *, const class QRect &, int)
func (this *QHeaderView) InheritPaintSection(f func(painter *qtgui.QPainter /*777 QPainter **/, rect *qtcore.QRect, logicalIndex int)) {
	ffiqt.SetAllInheritCallback(this, "paintSection", f)
}

// QSize sectionSizeFromContents(int)
func (this *QHeaderView) InheritSectionSizeFromContents(f func(logicalIndex int) unsafe.Pointer) {
	ffiqt.SetAllInheritCallback(this, "sectionSizeFromContents", f)
}

// int horizontalOffset()
func (this *QHeaderView) InheritHorizontalOffset(f func() int) {
	ffiqt.SetAllInheritCallback(this, "horizontalOffset", f)
}

// int verticalOffset()
func (this *QHeaderView) InheritVerticalOffset(f func() int) {
	ffiqt.SetAllInheritCallback(this, "verticalOffset", f)
}

// void updateGeometries()
func (this *QHeaderView) InheritUpdateGeometries(f func()) {
	ffiqt.SetAllInheritCallback(this, "updateGeometries", f)
}

// void scrollContentsBy(int, int)
func (this *QHeaderView) InheritScrollContentsBy(f func(dx int, dy int)) {
	ffiqt.SetAllInheritCallback(this, "scrollContentsBy", f)
}

// void rowsInserted(const class QModelIndex &, int, int)
func (this *QHeaderView) InheritRowsInserted(f func(parent *qtcore.QModelIndex, start int, end int)) {
	ffiqt.SetAllInheritCallback(this, "rowsInserted", f)
}

// QRect visualRect(const class QModelIndex &)
func (this *QHeaderView) InheritVisualRect(f func(index *qtcore.QModelIndex) unsafe.Pointer) {
	ffiqt.SetAllInheritCallback(this, "visualRect", f)
}

// void scrollTo(const class QModelIndex &, enum QAbstractItemView::ScrollHint)
func (this *QHeaderView) InheritScrollTo(f func(index *qtcore.QModelIndex, hint int)) {
	ffiqt.SetAllInheritCallback(this, "scrollTo", f)
}

// QModelIndex indexAt(const class QPoint &)
func (this *QHeaderView) InheritIndexAt(f func(p *qtcore.QPoint) unsafe.Pointer) {
	ffiqt.SetAllInheritCallback(this, "indexAt", f)
}

// bool isIndexHidden(const class QModelIndex &)
func (this *QHeaderView) InheritIsIndexHidden(f func(index *qtcore.QModelIndex) bool) {
	ffiqt.SetAllInheritCallback(this, "isIndexHidden", f)
}

// QModelIndex moveCursor(enum QAbstractItemView::CursorAction, Qt::KeyboardModifiers)
func (this *QHeaderView) InheritMoveCursor(f func(arg0 int, arg1 int) unsafe.Pointer) {
	ffiqt.SetAllInheritCallback(this, "moveCursor", f)
}

// void setSelection(const class QRect &, class QItemSelectionModel::SelectionFlags)
func (this *QHeaderView) InheritSetSelection(f func(rect *qtcore.QRect, flags int)) {
	ffiqt.SetAllInheritCallback(this, "setSelection", f)
}

// QRegion visualRegionForSelection(const class QItemSelection &)
func (this *QHeaderView) InheritVisualRegionForSelection(f func(selection *qtcore.QItemSelection) unsafe.Pointer) {
	ffiqt.SetAllInheritCallback(this, "visualRegionForSelection", f)
}

// void initStyleOption(class QStyleOptionHeader *)
func (this *QHeaderView) InheritInitStyleOption(f func(option *QStyleOptionHeader /*777 QStyleOptionHeader **/)) {
	ffiqt.SetAllInheritCallback(this, "initStyleOption", f)
}

type QHeaderView struct {
	*QAbstractItemView
}

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
// [8] const QMetaObject * metaObject()
func (this *QHeaderView) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QHeaderView10metaObjectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qheaderview.h:77
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QHeaderView(Qt::Orientation, QWidget *)
func NewQHeaderView(orientation int, parent *QWidget /*777 QWidget **/) *QHeaderView {
	var convArg1 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QHeaderViewC2EN2Qt11OrientationEP7QWidget", ffiqt.FFI_TYPE_POINTER, orientation, convArg1)
	gopp.ErrPrint(err, rv)
	gothis := NewQHeaderViewFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtWidgets/qheaderview.h:78
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QHeaderView()
func DeleteQHeaderView(this *QHeaderView) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QHeaderViewD2Ev", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 48)
	gopp.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtWidgets/qheaderview.h:80
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void setModel(QAbstractItemModel *)
func (this *QHeaderView) SetModel(model *qtcore.QAbstractItemModel /*777 QAbstractItemModel **/) {
	var convArg0 = model.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QHeaderView8setModelEP18QAbstractItemModel", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:82
// index:0
// Public Visibility=Default Availability=Available
// [4] Qt::Orientation orientation()
func (this *QHeaderView) Orientation() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QHeaderView11orientationEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:83
// index:0
// Public Visibility=Default Availability=Available
// [4] int offset()
func (this *QHeaderView) Offset() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QHeaderView6offsetEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qheaderview.h:84
// index:0
// Public Visibility=Default Availability=Available
// [4] int length()
func (this *QHeaderView) Length() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QHeaderView6lengthEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qheaderview.h:85
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QSize sizeHint()
func (this *QHeaderView) SizeHint() *qtcore.QSize /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QHeaderView8sizeHintEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSize)
	return rv2
}

// /usr/include/qt/QtWidgets/qheaderview.h:86
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void setVisible(_Bool)
func (this *QHeaderView) SetVisible(v bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QHeaderView10setVisibleEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), v)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:87
// index:0
// Public Visibility=Default Availability=Available
// [4] int sectionSizeHint(int)
func (this *QHeaderView) SectionSizeHint(logicalIndex int) int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QHeaderView15sectionSizeHintEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), logicalIndex)
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qheaderview.h:89
// index:0
// Public Visibility=Default Availability=Available
// [4] int visualIndexAt(int)
func (this *QHeaderView) VisualIndexAt(position int) int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QHeaderView13visualIndexAtEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), position)
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qheaderview.h:90
// index:0
// Public Visibility=Default Availability=Available
// [4] int logicalIndexAt(int)
func (this *QHeaderView) LogicalIndexAt(position int) int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QHeaderView14logicalIndexAtEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), position)
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qheaderview.h:92
// index:1
// Public inline Visibility=Default Availability=Available
// [4] int logicalIndexAt(int, int)
func (this *QHeaderView) LogicalIndexAt_1(x int, y int) int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QHeaderView14logicalIndexAtEii", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), x, y)
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qheaderview.h:93
// index:2
// Public inline Visibility=Default Availability=Available
// [4] int logicalIndexAt(const QPoint &)
func (this *QHeaderView) LogicalIndexAt_2(pos *qtcore.QPoint) int {
	var convArg0 = pos.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QHeaderView14logicalIndexAtERK6QPoint", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qheaderview.h:95
// index:0
// Public Visibility=Default Availability=Available
// [4] int sectionSize(int)
func (this *QHeaderView) SectionSize(logicalIndex int) int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QHeaderView11sectionSizeEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), logicalIndex)
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qheaderview.h:96
// index:0
// Public Visibility=Default Availability=Available
// [4] int sectionPosition(int)
func (this *QHeaderView) SectionPosition(logicalIndex int) int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QHeaderView15sectionPositionEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), logicalIndex)
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qheaderview.h:97
// index:0
// Public Visibility=Default Availability=Available
// [4] int sectionViewportPosition(int)
func (this *QHeaderView) SectionViewportPosition(logicalIndex int) int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QHeaderView23sectionViewportPositionEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), logicalIndex)
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qheaderview.h:99
// index:0
// Public Visibility=Default Availability=Available
// [-2] void moveSection(int, int)
func (this *QHeaderView) MoveSection(from int, to int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QHeaderView11moveSectionEii", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), from, to)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:100
// index:0
// Public Visibility=Default Availability=Available
// [-2] void swapSections(int, int)
func (this *QHeaderView) SwapSections(first int, second int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QHeaderView12swapSectionsEii", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), first, second)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:101
// index:0
// Public Visibility=Default Availability=Available
// [-2] void resizeSection(int, int)
func (this *QHeaderView) ResizeSection(logicalIndex int, size int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QHeaderView13resizeSectionEii", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), logicalIndex, size)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:102
// index:0
// Public Visibility=Default Availability=Available
// [-2] void resizeSections(QHeaderView::ResizeMode)
func (this *QHeaderView) ResizeSections(mode int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QHeaderView14resizeSectionsENS_10ResizeModeE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), mode)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:206
// index:1
// Protected Visibility=Default Availability=Available
// [-2] void resizeSections()
func (this *QHeaderView) ResizeSections_1() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QHeaderView14resizeSectionsEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:104
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isSectionHidden(int)
func (this *QHeaderView) IsSectionHidden(logicalIndex int) bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QHeaderView15isSectionHiddenEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), logicalIndex)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qheaderview.h:105
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setSectionHidden(int, _Bool)
func (this *QHeaderView) SetSectionHidden(logicalIndex int, hide bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QHeaderView16setSectionHiddenEib", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), logicalIndex, hide)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:106
// index:0
// Public Visibility=Default Availability=Available
// [4] int hiddenSectionCount()
func (this *QHeaderView) HiddenSectionCount() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QHeaderView18hiddenSectionCountEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qheaderview.h:108
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void hideSection(int)
func (this *QHeaderView) HideSection(logicalIndex int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QHeaderView11hideSectionEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), logicalIndex)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:109
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void showSection(int)
func (this *QHeaderView) ShowSection(logicalIndex int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QHeaderView11showSectionEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), logicalIndex)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:111
// index:0
// Public Visibility=Default Availability=Available
// [4] int count()
func (this *QHeaderView) Count() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QHeaderView5countEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qheaderview.h:112
// index:0
// Public Visibility=Default Availability=Available
// [4] int visualIndex(int)
func (this *QHeaderView) VisualIndex(logicalIndex int) int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QHeaderView11visualIndexEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), logicalIndex)
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qheaderview.h:113
// index:0
// Public Visibility=Default Availability=Available
// [4] int logicalIndex(int)
func (this *QHeaderView) LogicalIndex(visualIndex int) int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QHeaderView12logicalIndexEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), visualIndex)
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qheaderview.h:115
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setSectionsMovable(_Bool)
func (this *QHeaderView) SetSectionsMovable(movable bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QHeaderView18setSectionsMovableEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), movable)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:116
// index:0
// Public Visibility=Default Availability=Available
// [1] bool sectionsMovable()
func (this *QHeaderView) SectionsMovable() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QHeaderView15sectionsMovableEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qheaderview.h:122
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setSectionsClickable(_Bool)
func (this *QHeaderView) SetSectionsClickable(clickable bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QHeaderView20setSectionsClickableEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), clickable)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:123
// index:0
// Public Visibility=Default Availability=Available
// [1] bool sectionsClickable()
func (this *QHeaderView) SectionsClickable() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QHeaderView17sectionsClickableEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qheaderview.h:129
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setHighlightSections(_Bool)
func (this *QHeaderView) SetHighlightSections(highlight bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QHeaderView20setHighlightSectionsEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), highlight)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:130
// index:0
// Public Visibility=Default Availability=Available
// [1] bool highlightSections()
func (this *QHeaderView) HighlightSections() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QHeaderView17highlightSectionsEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qheaderview.h:132
// index:0
// Public Visibility=Default Availability=Available
// [4] QHeaderView::ResizeMode sectionResizeMode(int)
func (this *QHeaderView) SectionResizeMode(logicalIndex int) int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QHeaderView17sectionResizeModeEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), logicalIndex)
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:133
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setSectionResizeMode(enum QHeaderView::ResizeMode)
func (this *QHeaderView) SetSectionResizeMode(mode int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QHeaderView20setSectionResizeModeENS_10ResizeModeE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), mode)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:134
// index:1
// Public Visibility=Default Availability=Available
// [-2] void setSectionResizeMode(int, enum QHeaderView::ResizeMode)
func (this *QHeaderView) SetSectionResizeMode_1(logicalIndex int, mode int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QHeaderView20setSectionResizeModeEiNS_10ResizeModeE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), logicalIndex, mode)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:136
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setResizeContentsPrecision(int)
func (this *QHeaderView) SetResizeContentsPrecision(precision int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QHeaderView26setResizeContentsPrecisionEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), precision)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:137
// index:0
// Public Visibility=Default Availability=Available
// [4] int resizeContentsPrecision()
func (this *QHeaderView) ResizeContentsPrecision() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QHeaderView23resizeContentsPrecisionEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qheaderview.h:148
// index:0
// Public Visibility=Default Availability=Available
// [4] int stretchSectionCount()
func (this *QHeaderView) StretchSectionCount() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QHeaderView19stretchSectionCountEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qheaderview.h:150
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setSortIndicatorShown(_Bool)
func (this *QHeaderView) SetSortIndicatorShown(show bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QHeaderView21setSortIndicatorShownEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), show)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:151
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isSortIndicatorShown()
func (this *QHeaderView) IsSortIndicatorShown() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QHeaderView20isSortIndicatorShownEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qheaderview.h:153
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setSortIndicator(int, Qt::SortOrder)
func (this *QHeaderView) SetSortIndicator(logicalIndex int, order int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QHeaderView16setSortIndicatorEiN2Qt9SortOrderE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), logicalIndex, order)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:154
// index:0
// Public Visibility=Default Availability=Available
// [4] int sortIndicatorSection()
func (this *QHeaderView) SortIndicatorSection() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QHeaderView20sortIndicatorSectionEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qheaderview.h:155
// index:0
// Public Visibility=Default Availability=Available
// [4] Qt::SortOrder sortIndicatorOrder()
func (this *QHeaderView) SortIndicatorOrder() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QHeaderView18sortIndicatorOrderEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:157
// index:0
// Public Visibility=Default Availability=Available
// [1] bool stretchLastSection()
func (this *QHeaderView) StretchLastSection() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QHeaderView18stretchLastSectionEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qheaderview.h:158
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setStretchLastSection(_Bool)
func (this *QHeaderView) SetStretchLastSection(stretch bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QHeaderView21setStretchLastSectionEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), stretch)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:160
// index:0
// Public Visibility=Default Availability=Available
// [1] bool cascadingSectionResizes()
func (this *QHeaderView) CascadingSectionResizes() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QHeaderView23cascadingSectionResizesEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qheaderview.h:161
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setCascadingSectionResizes(_Bool)
func (this *QHeaderView) SetCascadingSectionResizes(enable bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QHeaderView26setCascadingSectionResizesEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), enable)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:163
// index:0
// Public Visibility=Default Availability=Available
// [4] int defaultSectionSize()
func (this *QHeaderView) DefaultSectionSize() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QHeaderView18defaultSectionSizeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qheaderview.h:164
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setDefaultSectionSize(int)
func (this *QHeaderView) SetDefaultSectionSize(size int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QHeaderView21setDefaultSectionSizeEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), size)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:165
// index:0
// Public Visibility=Default Availability=Available
// [-2] void resetDefaultSectionSize()
func (this *QHeaderView) ResetDefaultSectionSize() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QHeaderView23resetDefaultSectionSizeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:167
// index:0
// Public Visibility=Default Availability=Available
// [4] int minimumSectionSize()
func (this *QHeaderView) MinimumSectionSize() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QHeaderView18minimumSectionSizeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qheaderview.h:168
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setMinimumSectionSize(int)
func (this *QHeaderView) SetMinimumSectionSize(size int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QHeaderView21setMinimumSectionSizeEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), size)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:169
// index:0
// Public Visibility=Default Availability=Available
// [4] int maximumSectionSize()
func (this *QHeaderView) MaximumSectionSize() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QHeaderView18maximumSectionSizeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qheaderview.h:170
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setMaximumSectionSize(int)
func (this *QHeaderView) SetMaximumSectionSize(size int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QHeaderView21setMaximumSectionSizeEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), size)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:172
// index:0
// Public Visibility=Default Availability=Available
// [4] Qt::Alignment defaultAlignment()
func (this *QHeaderView) DefaultAlignment() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QHeaderView16defaultAlignmentEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:173
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setDefaultAlignment(Qt::Alignment)
func (this *QHeaderView) SetDefaultAlignment(alignment int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QHeaderView19setDefaultAlignmentE6QFlagsIN2Qt13AlignmentFlagEE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), alignment)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:175
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void doItemsLayout()
func (this *QHeaderView) DoItemsLayout() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QHeaderView13doItemsLayoutEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:176
// index:0
// Public Visibility=Default Availability=Available
// [1] bool sectionsMoved()
func (this *QHeaderView) SectionsMoved() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QHeaderView13sectionsMovedEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qheaderview.h:177
// index:0
// Public Visibility=Default Availability=Available
// [1] bool sectionsHidden()
func (this *QHeaderView) SectionsHidden() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QHeaderView14sectionsHiddenEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qheaderview.h:180
// index:0
// Public Visibility=Default Availability=Available
// [8] QByteArray saveState()
func (this *QHeaderView) SaveState() *qtcore.QByteArray /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QHeaderView9saveStateEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQByteArray)
	return rv2
}

// /usr/include/qt/QtWidgets/qheaderview.h:181
// index:0
// Public Visibility=Default Availability=Available
// [1] bool restoreState(const QByteArray &)
func (this *QHeaderView) RestoreState(state *qtcore.QByteArray) bool {
	var convArg0 = state.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QHeaderView12restoreStateERK10QByteArray", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qheaderview.h:184
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void reset()
func (this *QHeaderView) Reset() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QHeaderView5resetEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:187
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setOffset(int)
func (this *QHeaderView) SetOffset(offset int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QHeaderView9setOffsetEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), offset)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:188
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setOffsetToSectionPosition(int)
func (this *QHeaderView) SetOffsetToSectionPosition(visualIndex int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QHeaderView26setOffsetToSectionPositionEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), visualIndex)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:189
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setOffsetToLastSection()
func (this *QHeaderView) SetOffsetToLastSection() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QHeaderView22setOffsetToLastSectionEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:190
// index:0
// Public Visibility=Default Availability=Available
// [-2] void headerDataChanged(Qt::Orientation, int, int)
func (this *QHeaderView) HeaderDataChanged(orientation int, logicalFirst int, logicalLast int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QHeaderView17headerDataChangedEN2Qt11OrientationEii", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), orientation, logicalFirst, logicalLast)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:193
// index:0
// Public Visibility=Default Availability=Available
// [-2] void sectionMoved(int, int, int)
func (this *QHeaderView) SectionMoved(logicalIndex int, oldVisualIndex int, newVisualIndex int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QHeaderView12sectionMovedEiii", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), logicalIndex, oldVisualIndex, newVisualIndex)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:194
// index:0
// Public Visibility=Default Availability=Available
// [-2] void sectionResized(int, int, int)
func (this *QHeaderView) SectionResized(logicalIndex int, oldSize int, newSize int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QHeaderView14sectionResizedEiii", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), logicalIndex, oldSize, newSize)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:195
// index:0
// Public Visibility=Default Availability=Available
// [-2] void sectionPressed(int)
func (this *QHeaderView) SectionPressed(logicalIndex int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QHeaderView14sectionPressedEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), logicalIndex)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:196
// index:0
// Public Visibility=Default Availability=Available
// [-2] void sectionClicked(int)
func (this *QHeaderView) SectionClicked(logicalIndex int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QHeaderView14sectionClickedEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), logicalIndex)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:197
// index:0
// Public Visibility=Default Availability=Available
// [-2] void sectionEntered(int)
func (this *QHeaderView) SectionEntered(logicalIndex int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QHeaderView14sectionEnteredEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), logicalIndex)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:198
// index:0
// Public Visibility=Default Availability=Available
// [-2] void sectionDoubleClicked(int)
func (this *QHeaderView) SectionDoubleClicked(logicalIndex int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QHeaderView20sectionDoubleClickedEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), logicalIndex)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:199
// index:0
// Public Visibility=Default Availability=Available
// [-2] void sectionCountChanged(int, int)
func (this *QHeaderView) SectionCountChanged(oldCount int, newCount int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QHeaderView19sectionCountChangedEii", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), oldCount, newCount)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:200
// index:0
// Public Visibility=Default Availability=Available
// [-2] void sectionHandleDoubleClicked(int)
func (this *QHeaderView) SectionHandleDoubleClicked(logicalIndex int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QHeaderView26sectionHandleDoubleClickedEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), logicalIndex)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:201
// index:0
// Public Visibility=Default Availability=Available
// [-2] void geometriesChanged()
func (this *QHeaderView) GeometriesChanged() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QHeaderView17geometriesChangedEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:202
// index:0
// Public Visibility=Default Availability=Available
// [-2] void sortIndicatorChanged(int, Qt::SortOrder)
func (this *QHeaderView) SortIndicatorChanged(logicalIndex int, order int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QHeaderView20sortIndicatorChangedEiN2Qt9SortOrderE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), logicalIndex, order)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:205
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void updateSection(int)
func (this *QHeaderView) UpdateSection(logicalIndex int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QHeaderView13updateSectionEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), logicalIndex)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:207
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void sectionsInserted(const QModelIndex &, int, int)
func (this *QHeaderView) SectionsInserted(parent *qtcore.QModelIndex, logicalFirst int, logicalLast int) {
	var convArg0 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QHeaderView16sectionsInsertedERK11QModelIndexii", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, logicalFirst, logicalLast)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:208
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void sectionsAboutToBeRemoved(const QModelIndex &, int, int)
func (this *QHeaderView) SectionsAboutToBeRemoved(parent *qtcore.QModelIndex, logicalFirst int, logicalLast int) {
	var convArg0 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QHeaderView24sectionsAboutToBeRemovedERK11QModelIndexii", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, logicalFirst, logicalLast)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:212
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void initialize()
func (this *QHeaderView) Initialize() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QHeaderView10initializeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:214
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void initializeSections()
func (this *QHeaderView) InitializeSections() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QHeaderView18initializeSectionsEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:215
// index:1
// Protected Visibility=Default Availability=Available
// [-2] void initializeSections(int, int)
func (this *QHeaderView) InitializeSections_1(start int, end int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QHeaderView18initializeSectionsEii", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), start, end)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:216
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void currentChanged(const QModelIndex &, const QModelIndex &)
func (this *QHeaderView) CurrentChanged(current *qtcore.QModelIndex, old *qtcore.QModelIndex) {
	var convArg0 = current.GetCthis()
	var convArg1 = old.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QHeaderView14currentChangedERK11QModelIndexS2_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:218
// index:0
// Protected virtual Visibility=Default Availability=Available
// [1] bool event(QEvent *)
func (this *QHeaderView) Event(e *qtcore.QEvent /*777 QEvent **/) bool {
	var convArg0 = e.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QHeaderView5eventEP6QEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qheaderview.h:219
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void paintEvent(QPaintEvent *)
func (this *QHeaderView) PaintEvent(e *qtgui.QPaintEvent /*777 QPaintEvent **/) {
	var convArg0 = e.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QHeaderView10paintEventEP11QPaintEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:220
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void mousePressEvent(QMouseEvent *)
func (this *QHeaderView) MousePressEvent(e *qtgui.QMouseEvent /*777 QMouseEvent **/) {
	var convArg0 = e.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QHeaderView15mousePressEventEP11QMouseEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:221
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void mouseMoveEvent(QMouseEvent *)
func (this *QHeaderView) MouseMoveEvent(e *qtgui.QMouseEvent /*777 QMouseEvent **/) {
	var convArg0 = e.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QHeaderView14mouseMoveEventEP11QMouseEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:222
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void mouseReleaseEvent(QMouseEvent *)
func (this *QHeaderView) MouseReleaseEvent(e *qtgui.QMouseEvent /*777 QMouseEvent **/) {
	var convArg0 = e.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QHeaderView17mouseReleaseEventEP11QMouseEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:223
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void mouseDoubleClickEvent(QMouseEvent *)
func (this *QHeaderView) MouseDoubleClickEvent(e *qtgui.QMouseEvent /*777 QMouseEvent **/) {
	var convArg0 = e.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QHeaderView21mouseDoubleClickEventEP11QMouseEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:224
// index:0
// Protected virtual Visibility=Default Availability=Available
// [1] bool viewportEvent(QEvent *)
func (this *QHeaderView) ViewportEvent(e *qtcore.QEvent /*777 QEvent **/) bool {
	var convArg0 = e.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QHeaderView13viewportEventEP6QEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qheaderview.h:226
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void paintSection(QPainter *, const QRect &, int)
func (this *QHeaderView) PaintSection(painter *qtgui.QPainter /*777 QPainter **/, rect *qtcore.QRect, logicalIndex int) {
	var convArg0 = painter.GetCthis()
	var convArg1 = rect.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QHeaderView12paintSectionEP8QPainterRK5QRecti", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, logicalIndex)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:227
// index:0
// Protected virtual Visibility=Default Availability=Available
// [8] QSize sectionSizeFromContents(int)
func (this *QHeaderView) SectionSizeFromContents(logicalIndex int) *qtcore.QSize /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QHeaderView23sectionSizeFromContentsEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), logicalIndex)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSize)
	return rv2
}

// /usr/include/qt/QtWidgets/qheaderview.h:229
// index:0
// Protected virtual Visibility=Default Availability=Available
// [4] int horizontalOffset()
func (this *QHeaderView) HorizontalOffset() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QHeaderView16horizontalOffsetEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qheaderview.h:230
// index:0
// Protected virtual Visibility=Default Availability=Available
// [4] int verticalOffset()
func (this *QHeaderView) VerticalOffset() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QHeaderView14verticalOffsetEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qheaderview.h:231
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void updateGeometries()
func (this *QHeaderView) UpdateGeometries() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QHeaderView16updateGeometriesEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:232
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void scrollContentsBy(int, int)
func (this *QHeaderView) ScrollContentsBy(dx int, dy int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QHeaderView16scrollContentsByEii", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), dx, dy)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:235
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void rowsInserted(const QModelIndex &, int, int)
func (this *QHeaderView) RowsInserted(parent *qtcore.QModelIndex, start int, end int) {
	var convArg0 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QHeaderView12rowsInsertedERK11QModelIndexii", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, start, end)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:237
// index:0
// Protected virtual Visibility=Default Availability=Available
// [16] QRect visualRect(const QModelIndex &)
func (this *QHeaderView) VisualRect(index *qtcore.QModelIndex) *qtcore.QRect /*123*/ {
	var convArg0 = index.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QHeaderView10visualRectERK11QModelIndex", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQRectFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQRect)
	return rv2
}

// /usr/include/qt/QtWidgets/qheaderview.h:238
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void scrollTo(const QModelIndex &, enum QAbstractItemView::ScrollHint)
func (this *QHeaderView) ScrollTo(index *qtcore.QModelIndex, hint int) {
	var convArg0 = index.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QHeaderView8scrollToERK11QModelIndexN17QAbstractItemView10ScrollHintE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, hint)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:240
// index:0
// Protected virtual Visibility=Default Availability=Available
// [24] QModelIndex indexAt(const QPoint &)
func (this *QHeaderView) IndexAt(p *qtcore.QPoint) *qtcore.QModelIndex /*123*/ {
	var convArg0 = p.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QHeaderView7indexAtERK6QPoint", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQModelIndexFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQModelIndex)
	return rv2
}

// /usr/include/qt/QtWidgets/qheaderview.h:241
// index:0
// Protected virtual Visibility=Default Availability=Available
// [1] bool isIndexHidden(const QModelIndex &)
func (this *QHeaderView) IsIndexHidden(index *qtcore.QModelIndex) bool {
	var convArg0 = index.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QHeaderView13isIndexHiddenERK11QModelIndex", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qheaderview.h:243
// index:0
// Protected virtual Visibility=Default Availability=Available
// [24] QModelIndex moveCursor(enum QAbstractItemView::CursorAction, Qt::KeyboardModifiers)
func (this *QHeaderView) MoveCursor(arg0 int, arg1 int) *qtcore.QModelIndex /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QHeaderView10moveCursorEN17QAbstractItemView12CursorActionE6QFlagsIN2Qt16KeyboardModifierEE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), arg0, arg1)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQModelIndexFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQModelIndex)
	return rv2
}

// /usr/include/qt/QtWidgets/qheaderview.h:244
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void setSelection(const QRect &, QItemSelectionModel::SelectionFlags)
func (this *QHeaderView) SetSelection(rect *qtcore.QRect, flags int) {
	var convArg0 = rect.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QHeaderView12setSelectionERK5QRect6QFlagsIN19QItemSelectionModel13SelectionFlagEE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, flags)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:245
// index:0
// Protected virtual Visibility=Default Availability=Available
// [8] QRegion visualRegionForSelection(const QItemSelection &)
func (this *QHeaderView) VisualRegionForSelection(selection *qtcore.QItemSelection) *qtgui.QRegion /*123*/ {
	var convArg0 = selection.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QHeaderView24visualRegionForSelectionERK14QItemSelection", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtgui.NewQRegionFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtgui.DeleteQRegion)
	return rv2
}

// /usr/include/qt/QtWidgets/qheaderview.h:246
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void initStyleOption(QStyleOptionHeader *)
func (this *QHeaderView) InitStyleOption(option *QStyleOptionHeader /*777 QStyleOptionHeader **/) {
	var convArg0 = option.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QHeaderView15initStyleOptionEP18QStyleOptionHeader", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

type QHeaderView__ResizeMode = int

const QHeaderView__Interactive QHeaderView__ResizeMode = 0
const QHeaderView__Stretch QHeaderView__ResizeMode = 1
const QHeaderView__Fixed QHeaderView__ResizeMode = 2
const QHeaderView__ResizeToContents QHeaderView__ResizeMode = 3
const QHeaderView__Custom QHeaderView__ResizeMode = 2

//  body block end
