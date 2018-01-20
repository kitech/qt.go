//  header block begin
// /usr/include/qt/QtWidgets/qheaderview.h
// #include <qheaderview.h>
// #include <QtWidgets>
package qtwidgets

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
import "qtrt"
import "mkuse/cffiqt"
import "gopp"
import "qtcore"
import "qtgui"

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
type QHeaderView struct {
	*QAbstractItemView
}

func (this *QHeaderView) GetCthis() unsafe.Pointer {
	return this.QAbstractItemView.GetCthis()
}

// /usr/include/qt/QtWidgets/qheaderview.h:55
// index:0
// virtual
// const QMetaObject * metaObject()
func (this *QHeaderView) MetaObject() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QHeaderView10metaObjectEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:77
// index:0
// void QHeaderView(Qt::Orientation, class QWidget *)
func NewQHeaderView(orientation int, parent unsafe.Pointer) *QHeaderView {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QHeaderViewC2EN2Qt11OrientationEP7QWidget", ffiqt.FFI_TYPE_VOID, cthis, &orientation, parent)
	gopp.ErrPrint(err, rv)
	gothis := NewQHeaderViewFromPointer(cthis)
	return gothis
}
func NewQHeaderViewFromPointer(cthis unsafe.Pointer) *QHeaderView {
	bcthis0 := NewQAbstractItemViewFromPointer(cthis)
	return &QHeaderView{bcthis0}
}

// /usr/include/qt/QtWidgets/qheaderview.h:211
// index:1
// void QHeaderView(class QHeaderViewPrivate &, Qt::Orientation, class QWidget *)
func NewQHeaderView_1(dd unsafe.Pointer, orientation int, parent unsafe.Pointer) *QHeaderView {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QHeaderViewC2ER18QHeaderViewPrivateN2Qt11OrientationEP7QWidget", ffiqt.FFI_TYPE_VOID, cthis, dd, &orientation, parent)
	gopp.ErrPrint(err, rv)
	gothis := NewQHeaderViewFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtWidgets/qheaderview.h:78
// index:0
// virtual
// void ~QHeaderView()
func DeleteQHeaderView(*QHeaderView) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QHeaderViewD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:80
// index:0
// virtual
// void setModel(class QAbstractItemModel *)
func (this *QHeaderView) SetModel(model unsafe.Pointer) {
	// 0: (, model QAbstractItemModel *), (model)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QHeaderView8setModelEP18QAbstractItemModel", ffiqt.FFI_TYPE_VOID, this.GetCthis(), model)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:82
// index:0
// Qt::Orientation orientation()
func (this *QHeaderView) Orientation() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QHeaderView11orientationEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:83
// index:0
// int offset()
func (this *QHeaderView) Offset() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QHeaderView6offsetEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:84
// index:0
// int length()
func (this *QHeaderView) Length() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QHeaderView6lengthEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:85
// index:0
// virtual
// QSize sizeHint()
func (this *QHeaderView) SizeHint() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QHeaderView8sizeHintEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:86
// index:0
// virtual
// void setVisible(_Bool)
func (this *QHeaderView) SetVisible(v bool) {
	// 0: (, v bool), (&v)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QHeaderView10setVisibleEb", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &v)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:87
// index:0
// int sectionSizeHint(int)
func (this *QHeaderView) SectionSizeHint(logicalIndex int) {
	// 0: (, logicalIndex int), (&logicalIndex)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QHeaderView15sectionSizeHintEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &logicalIndex)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:89
// index:0
// int visualIndexAt(int)
func (this *QHeaderView) VisualIndexAt(position int) {
	// 0: (, position int), (&position)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QHeaderView13visualIndexAtEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &position)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:90
// index:0
// int logicalIndexAt(int)
func (this *QHeaderView) LogicalIndexAt(position int) {
	// 0: (, position int), (&position)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QHeaderView14logicalIndexAtEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &position)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:92
// index:1
// inline
// int logicalIndexAt(int, int)
func (this *QHeaderView) LogicalIndexAt_1(x int, y int) {
	// 1: (, x int, y int), (&x, &y)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QHeaderView14logicalIndexAtEii", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &x, &y)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:93
// index:2
// inline
// int logicalIndexAt(const class QPoint &)
func (this *QHeaderView) LogicalIndexAt_2(pos unsafe.Pointer) {
	// 2: (, pos const QPoint &), (pos)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QHeaderView14logicalIndexAtERK6QPoint", ffiqt.FFI_TYPE_VOID, this.GetCthis(), pos)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:95
// index:0
// int sectionSize(int)
func (this *QHeaderView) SectionSize(logicalIndex int) {
	// 0: (, logicalIndex int), (&logicalIndex)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QHeaderView11sectionSizeEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &logicalIndex)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:96
// index:0
// int sectionPosition(int)
func (this *QHeaderView) SectionPosition(logicalIndex int) {
	// 0: (, logicalIndex int), (&logicalIndex)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QHeaderView15sectionPositionEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &logicalIndex)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:97
// index:0
// int sectionViewportPosition(int)
func (this *QHeaderView) SectionViewportPosition(logicalIndex int) {
	// 0: (, logicalIndex int), (&logicalIndex)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QHeaderView23sectionViewportPositionEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &logicalIndex)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:99
// index:0
// void moveSection(int, int)
func (this *QHeaderView) MoveSection(from int, to int) {
	// 0: (, from int, to int), (&from, &to)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QHeaderView11moveSectionEii", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &from, &to)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:100
// index:0
// void swapSections(int, int)
func (this *QHeaderView) SwapSections(first int, second int) {
	// 0: (, first int, second int), (&first, &second)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QHeaderView12swapSectionsEii", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &first, &second)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:101
// index:0
// void resizeSection(int, int)
func (this *QHeaderView) ResizeSection(logicalIndex int, size int) {
	// 0: (, logicalIndex int, size int), (&logicalIndex, &size)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QHeaderView13resizeSectionEii", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &logicalIndex, &size)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:102
// index:0
// void resizeSections(class QHeaderView::ResizeMode)
func (this *QHeaderView) ResizeSections(mode int) {
	// 0: (, mode QHeaderView::ResizeMode), (&mode)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QHeaderView14resizeSectionsENS_10ResizeModeE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &mode)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:206
// index:1
// void resizeSections()
func (this *QHeaderView) ResizeSections_1() {
	// 1: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QHeaderView14resizeSectionsEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:104
// index:0
// bool isSectionHidden(int)
func (this *QHeaderView) IsSectionHidden(logicalIndex int) {
	// 0: (, logicalIndex int), (&logicalIndex)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QHeaderView15isSectionHiddenEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &logicalIndex)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:105
// index:0
// void setSectionHidden(int, _Bool)
func (this *QHeaderView) SetSectionHidden(logicalIndex int, hide bool) {
	// 0: (, logicalIndex int, hide bool), (&logicalIndex, &hide)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QHeaderView16setSectionHiddenEib", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &logicalIndex, &hide)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:106
// index:0
// int hiddenSectionCount()
func (this *QHeaderView) HiddenSectionCount() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QHeaderView18hiddenSectionCountEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:108
// index:0
// inline
// void hideSection(int)
func (this *QHeaderView) HideSection(logicalIndex int) {
	// 0: (, logicalIndex int), (&logicalIndex)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QHeaderView11hideSectionEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &logicalIndex)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:109
// index:0
// inline
// void showSection(int)
func (this *QHeaderView) ShowSection(logicalIndex int) {
	// 0: (, logicalIndex int), (&logicalIndex)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QHeaderView11showSectionEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &logicalIndex)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:111
// index:0
// int count()
func (this *QHeaderView) Count() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QHeaderView5countEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:112
// index:0
// int visualIndex(int)
func (this *QHeaderView) VisualIndex(logicalIndex int) {
	// 0: (, logicalIndex int), (&logicalIndex)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QHeaderView11visualIndexEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &logicalIndex)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:113
// index:0
// int logicalIndex(int)
func (this *QHeaderView) LogicalIndex(visualIndex int) {
	// 0: (, visualIndex int), (&visualIndex)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QHeaderView12logicalIndexEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &visualIndex)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:115
// index:0
// void setSectionsMovable(_Bool)
func (this *QHeaderView) SetSectionsMovable(movable bool) {
	// 0: (, movable bool), (&movable)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QHeaderView18setSectionsMovableEb", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &movable)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:116
// index:0
// bool sectionsMovable()
func (this *QHeaderView) SectionsMovable() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QHeaderView15sectionsMovableEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:122
// index:0
// void setSectionsClickable(_Bool)
func (this *QHeaderView) SetSectionsClickable(clickable bool) {
	// 0: (, clickable bool), (&clickable)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QHeaderView20setSectionsClickableEb", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &clickable)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:123
// index:0
// bool sectionsClickable()
func (this *QHeaderView) SectionsClickable() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QHeaderView17sectionsClickableEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:129
// index:0
// void setHighlightSections(_Bool)
func (this *QHeaderView) SetHighlightSections(highlight bool) {
	// 0: (, highlight bool), (&highlight)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QHeaderView20setHighlightSectionsEb", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &highlight)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:130
// index:0
// bool highlightSections()
func (this *QHeaderView) HighlightSections() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QHeaderView17highlightSectionsEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:132
// index:0
// QHeaderView::ResizeMode sectionResizeMode(int)
func (this *QHeaderView) SectionResizeMode(logicalIndex int) {
	// 0: (, logicalIndex int), (&logicalIndex)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QHeaderView17sectionResizeModeEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &logicalIndex)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:133
// index:0
// void setSectionResizeMode(enum QHeaderView::ResizeMode)
func (this *QHeaderView) SetSectionResizeMode(mode int) {
	// 0: (, mode QHeaderView::ResizeMode), (&mode)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QHeaderView20setSectionResizeModeENS_10ResizeModeE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &mode)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:134
// index:1
// void setSectionResizeMode(int, enum QHeaderView::ResizeMode)
func (this *QHeaderView) SetSectionResizeMode_1(logicalIndex int, mode int) {
	// 1: (, logicalIndex int, mode QHeaderView::ResizeMode), (&logicalIndex, &mode)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QHeaderView20setSectionResizeModeEiNS_10ResizeModeE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &logicalIndex, &mode)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:136
// index:0
// void setResizeContentsPrecision(int)
func (this *QHeaderView) SetResizeContentsPrecision(precision int) {
	// 0: (, precision int), (&precision)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QHeaderView26setResizeContentsPrecisionEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &precision)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:137
// index:0
// int resizeContentsPrecision()
func (this *QHeaderView) ResizeContentsPrecision() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QHeaderView23resizeContentsPrecisionEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:148
// index:0
// int stretchSectionCount()
func (this *QHeaderView) StretchSectionCount() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QHeaderView19stretchSectionCountEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:150
// index:0
// void setSortIndicatorShown(_Bool)
func (this *QHeaderView) SetSortIndicatorShown(show bool) {
	// 0: (, show bool), (&show)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QHeaderView21setSortIndicatorShownEb", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &show)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:151
// index:0
// bool isSortIndicatorShown()
func (this *QHeaderView) IsSortIndicatorShown() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QHeaderView20isSortIndicatorShownEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:153
// index:0
// void setSortIndicator(int, Qt::SortOrder)
func (this *QHeaderView) SetSortIndicator(logicalIndex int, order int) {
	// 0: (, logicalIndex int, order Qt::SortOrder), (&logicalIndex, &order)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QHeaderView16setSortIndicatorEiN2Qt9SortOrderE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &logicalIndex, &order)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:154
// index:0
// int sortIndicatorSection()
func (this *QHeaderView) SortIndicatorSection() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QHeaderView20sortIndicatorSectionEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:155
// index:0
// Qt::SortOrder sortIndicatorOrder()
func (this *QHeaderView) SortIndicatorOrder() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QHeaderView18sortIndicatorOrderEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:157
// index:0
// bool stretchLastSection()
func (this *QHeaderView) StretchLastSection() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QHeaderView18stretchLastSectionEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:158
// index:0
// void setStretchLastSection(_Bool)
func (this *QHeaderView) SetStretchLastSection(stretch bool) {
	// 0: (, stretch bool), (&stretch)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QHeaderView21setStretchLastSectionEb", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &stretch)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:160
// index:0
// bool cascadingSectionResizes()
func (this *QHeaderView) CascadingSectionResizes() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QHeaderView23cascadingSectionResizesEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:161
// index:0
// void setCascadingSectionResizes(_Bool)
func (this *QHeaderView) SetCascadingSectionResizes(enable bool) {
	// 0: (, enable bool), (&enable)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QHeaderView26setCascadingSectionResizesEb", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &enable)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:163
// index:0
// int defaultSectionSize()
func (this *QHeaderView) DefaultSectionSize() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QHeaderView18defaultSectionSizeEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:164
// index:0
// void setDefaultSectionSize(int)
func (this *QHeaderView) SetDefaultSectionSize(size int) {
	// 0: (, size int), (&size)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QHeaderView21setDefaultSectionSizeEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &size)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:165
// index:0
// void resetDefaultSectionSize()
func (this *QHeaderView) ResetDefaultSectionSize() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QHeaderView23resetDefaultSectionSizeEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:167
// index:0
// int minimumSectionSize()
func (this *QHeaderView) MinimumSectionSize() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QHeaderView18minimumSectionSizeEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:168
// index:0
// void setMinimumSectionSize(int)
func (this *QHeaderView) SetMinimumSectionSize(size int) {
	// 0: (, size int), (&size)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QHeaderView21setMinimumSectionSizeEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &size)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:169
// index:0
// int maximumSectionSize()
func (this *QHeaderView) MaximumSectionSize() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QHeaderView18maximumSectionSizeEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:170
// index:0
// void setMaximumSectionSize(int)
func (this *QHeaderView) SetMaximumSectionSize(size int) {
	// 0: (, size int), (&size)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QHeaderView21setMaximumSectionSizeEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &size)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:172
// index:0
// Qt::Alignment defaultAlignment()
func (this *QHeaderView) DefaultAlignment() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QHeaderView16defaultAlignmentEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:173
// index:0
// void setDefaultAlignment(Qt::Alignment)
func (this *QHeaderView) SetDefaultAlignment(alignment int) {
	// 0: (, QFlags<Qt::AlignmentFlag> alignment), (&alignment)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QHeaderView19setDefaultAlignmentE6QFlagsIN2Qt13AlignmentFlagEE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &alignment)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:175
// index:0
// virtual
// void doItemsLayout()
func (this *QHeaderView) DoItemsLayout() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QHeaderView13doItemsLayoutEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:176
// index:0
// bool sectionsMoved()
func (this *QHeaderView) SectionsMoved() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QHeaderView13sectionsMovedEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:177
// index:0
// bool sectionsHidden()
func (this *QHeaderView) SectionsHidden() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QHeaderView14sectionsHiddenEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:180
// index:0
// QByteArray saveState()
func (this *QHeaderView) SaveState() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QHeaderView9saveStateEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:181
// index:0
// bool restoreState(const class QByteArray &)
func (this *QHeaderView) RestoreState(state unsafe.Pointer) {
	// 0: (, state const QByteArray &), (state)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QHeaderView12restoreStateERK10QByteArray", ffiqt.FFI_TYPE_VOID, this.GetCthis(), state)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:184
// index:0
// virtual
// void reset()
func (this *QHeaderView) Reset() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QHeaderView5resetEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:187
// index:0
// void setOffset(int)
func (this *QHeaderView) SetOffset(offset int) {
	// 0: (, offset int), (&offset)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QHeaderView9setOffsetEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &offset)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:188
// index:0
// void setOffsetToSectionPosition(int)
func (this *QHeaderView) SetOffsetToSectionPosition(visualIndex int) {
	// 0: (, visualIndex int), (&visualIndex)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QHeaderView26setOffsetToSectionPositionEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &visualIndex)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:189
// index:0
// void setOffsetToLastSection()
func (this *QHeaderView) SetOffsetToLastSection() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QHeaderView22setOffsetToLastSectionEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:190
// index:0
// void headerDataChanged(Qt::Orientation, int, int)
func (this *QHeaderView) HeaderDataChanged(orientation int, logicalFirst int, logicalLast int) {
	// 0: (, orientation Qt::Orientation, logicalFirst int, logicalLast int), (&orientation, &logicalFirst, &logicalLast)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QHeaderView17headerDataChangedEN2Qt11OrientationEii", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &orientation, &logicalFirst, &logicalLast)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:193
// index:0
// void sectionMoved(int, int, int)
func (this *QHeaderView) SectionMoved(logicalIndex int, oldVisualIndex int, newVisualIndex int) {
	// 0: (, logicalIndex int, oldVisualIndex int, newVisualIndex int), (&logicalIndex, &oldVisualIndex, &newVisualIndex)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QHeaderView12sectionMovedEiii", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &logicalIndex, &oldVisualIndex, &newVisualIndex)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:194
// index:0
// void sectionResized(int, int, int)
func (this *QHeaderView) SectionResized(logicalIndex int, oldSize int, newSize int) {
	// 0: (, logicalIndex int, oldSize int, newSize int), (&logicalIndex, &oldSize, &newSize)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QHeaderView14sectionResizedEiii", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &logicalIndex, &oldSize, &newSize)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:195
// index:0
// void sectionPressed(int)
func (this *QHeaderView) SectionPressed(logicalIndex int) {
	// 0: (, logicalIndex int), (&logicalIndex)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QHeaderView14sectionPressedEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &logicalIndex)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:196
// index:0
// void sectionClicked(int)
func (this *QHeaderView) SectionClicked(logicalIndex int) {
	// 0: (, logicalIndex int), (&logicalIndex)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QHeaderView14sectionClickedEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &logicalIndex)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:197
// index:0
// void sectionEntered(int)
func (this *QHeaderView) SectionEntered(logicalIndex int) {
	// 0: (, logicalIndex int), (&logicalIndex)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QHeaderView14sectionEnteredEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &logicalIndex)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:198
// index:0
// void sectionDoubleClicked(int)
func (this *QHeaderView) SectionDoubleClicked(logicalIndex int) {
	// 0: (, logicalIndex int), (&logicalIndex)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QHeaderView20sectionDoubleClickedEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &logicalIndex)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:199
// index:0
// void sectionCountChanged(int, int)
func (this *QHeaderView) SectionCountChanged(oldCount int, newCount int) {
	// 0: (, oldCount int, newCount int), (&oldCount, &newCount)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QHeaderView19sectionCountChangedEii", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &oldCount, &newCount)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:200
// index:0
// void sectionHandleDoubleClicked(int)
func (this *QHeaderView) SectionHandleDoubleClicked(logicalIndex int) {
	// 0: (, logicalIndex int), (&logicalIndex)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QHeaderView26sectionHandleDoubleClickedEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &logicalIndex)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:201
// index:0
// void geometriesChanged()
func (this *QHeaderView) GeometriesChanged() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QHeaderView17geometriesChangedEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:202
// index:0
// void sortIndicatorChanged(int, Qt::SortOrder)
func (this *QHeaderView) SortIndicatorChanged(logicalIndex int, order int) {
	// 0: (, logicalIndex int, order Qt::SortOrder), (&logicalIndex, &order)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QHeaderView20sortIndicatorChangedEiN2Qt9SortOrderE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &logicalIndex, &order)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:205
// index:0
// void updateSection(int)
func (this *QHeaderView) UpdateSection(logicalIndex int) {
	// 0: (, logicalIndex int), (&logicalIndex)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QHeaderView13updateSectionEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &logicalIndex)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:207
// index:0
// void sectionsInserted(const class QModelIndex &, int, int)
func (this *QHeaderView) SectionsInserted(parent unsafe.Pointer, logicalFirst int, logicalLast int) {
	// 0: (, parent const QModelIndex &, logicalFirst int, logicalLast int), (parent, &logicalFirst, &logicalLast)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QHeaderView16sectionsInsertedERK11QModelIndexii", ffiqt.FFI_TYPE_VOID, this.GetCthis(), parent, &logicalFirst, &logicalLast)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:208
// index:0
// void sectionsAboutToBeRemoved(const class QModelIndex &, int, int)
func (this *QHeaderView) SectionsAboutToBeRemoved(parent unsafe.Pointer, logicalFirst int, logicalLast int) {
	// 0: (, parent const QModelIndex &, logicalFirst int, logicalLast int), (parent, &logicalFirst, &logicalLast)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QHeaderView24sectionsAboutToBeRemovedERK11QModelIndexii", ffiqt.FFI_TYPE_VOID, this.GetCthis(), parent, &logicalFirst, &logicalLast)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:212
// index:0
// void initialize()
func (this *QHeaderView) Initialize() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QHeaderView10initializeEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:214
// index:0
// void initializeSections()
func (this *QHeaderView) InitializeSections() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QHeaderView18initializeSectionsEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:215
// index:1
// void initializeSections(int, int)
func (this *QHeaderView) InitializeSections_1(start int, end int) {
	// 1: (, start int, end int), (&start, &end)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QHeaderView18initializeSectionsEii", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &start, &end)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:216
// index:0
// virtual
// void currentChanged(const class QModelIndex &, const class QModelIndex &)
func (this *QHeaderView) CurrentChanged(current unsafe.Pointer, old unsafe.Pointer) {
	// 0: (, current const QModelIndex &, old const QModelIndex &), (current, old)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QHeaderView14currentChangedERK11QModelIndexS2_", ffiqt.FFI_TYPE_VOID, this.GetCthis(), current, old)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:218
// index:0
// virtual
// bool event(class QEvent *)
func (this *QHeaderView) Event(e unsafe.Pointer) {
	// 0: (, e QEvent *), (e)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QHeaderView5eventEP6QEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), e)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:219
// index:0
// virtual
// void paintEvent(class QPaintEvent *)
func (this *QHeaderView) PaintEvent(e unsafe.Pointer) {
	// 0: (, e QPaintEvent *), (e)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QHeaderView10paintEventEP11QPaintEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), e)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:220
// index:0
// virtual
// void mousePressEvent(class QMouseEvent *)
func (this *QHeaderView) MousePressEvent(e unsafe.Pointer) {
	// 0: (, e QMouseEvent *), (e)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QHeaderView15mousePressEventEP11QMouseEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), e)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:221
// index:0
// virtual
// void mouseMoveEvent(class QMouseEvent *)
func (this *QHeaderView) MouseMoveEvent(e unsafe.Pointer) {
	// 0: (, e QMouseEvent *), (e)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QHeaderView14mouseMoveEventEP11QMouseEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), e)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:222
// index:0
// virtual
// void mouseReleaseEvent(class QMouseEvent *)
func (this *QHeaderView) MouseReleaseEvent(e unsafe.Pointer) {
	// 0: (, e QMouseEvent *), (e)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QHeaderView17mouseReleaseEventEP11QMouseEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), e)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:223
// index:0
// virtual
// void mouseDoubleClickEvent(class QMouseEvent *)
func (this *QHeaderView) MouseDoubleClickEvent(e unsafe.Pointer) {
	// 0: (, e QMouseEvent *), (e)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QHeaderView21mouseDoubleClickEventEP11QMouseEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), e)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:224
// index:0
// virtual
// bool viewportEvent(class QEvent *)
func (this *QHeaderView) ViewportEvent(e unsafe.Pointer) {
	// 0: (, e QEvent *), (e)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QHeaderView13viewportEventEP6QEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), e)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:226
// index:0
// virtual
// void paintSection(class QPainter *, const class QRect &, int)
func (this *QHeaderView) PaintSection(painter unsafe.Pointer, rect unsafe.Pointer, logicalIndex int) {
	// 0: (, painter QPainter *, rect const QRect &, logicalIndex int), (painter, rect, &logicalIndex)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QHeaderView12paintSectionEP8QPainterRK5QRecti", ffiqt.FFI_TYPE_VOID, this.GetCthis(), painter, rect, &logicalIndex)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:227
// index:0
// virtual
// QSize sectionSizeFromContents(int)
func (this *QHeaderView) SectionSizeFromContents(logicalIndex int) {
	// 0: (, logicalIndex int), (&logicalIndex)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QHeaderView23sectionSizeFromContentsEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &logicalIndex)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:229
// index:0
// virtual
// int horizontalOffset()
func (this *QHeaderView) HorizontalOffset() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QHeaderView16horizontalOffsetEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:230
// index:0
// virtual
// int verticalOffset()
func (this *QHeaderView) VerticalOffset() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QHeaderView14verticalOffsetEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:231
// index:0
// virtual
// void updateGeometries()
func (this *QHeaderView) UpdateGeometries() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QHeaderView16updateGeometriesEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:232
// index:0
// virtual
// void scrollContentsBy(int, int)
func (this *QHeaderView) ScrollContentsBy(dx int, dy int) {
	// 0: (, dx int, dy int), (&dx, &dy)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QHeaderView16scrollContentsByEii", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &dx, &dy)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:235
// index:0
// virtual
// void rowsInserted(const class QModelIndex &, int, int)
func (this *QHeaderView) RowsInserted(parent unsafe.Pointer, start int, end int) {
	// 0: (, parent const QModelIndex &, start int, end int), (parent, &start, &end)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QHeaderView12rowsInsertedERK11QModelIndexii", ffiqt.FFI_TYPE_VOID, this.GetCthis(), parent, &start, &end)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:237
// index:0
// virtual
// QRect visualRect(const class QModelIndex &)
func (this *QHeaderView) VisualRect(index unsafe.Pointer) {
	// 0: (, index const QModelIndex &), (index)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QHeaderView10visualRectERK11QModelIndex", ffiqt.FFI_TYPE_VOID, this.GetCthis(), index)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:238
// index:0
// virtual
// void scrollTo(const class QModelIndex &, enum QAbstractItemView::ScrollHint)
func (this *QHeaderView) ScrollTo(index unsafe.Pointer, hint int) {
	// 0: (, index const QModelIndex &, hint QAbstractItemView::ScrollHint), (index, &hint)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QHeaderView8scrollToERK11QModelIndexN17QAbstractItemView10ScrollHintE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), index, &hint)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:240
// index:0
// virtual
// QModelIndex indexAt(const class QPoint &)
func (this *QHeaderView) IndexAt(p unsafe.Pointer) {
	// 0: (, p const QPoint &), (p)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QHeaderView7indexAtERK6QPoint", ffiqt.FFI_TYPE_VOID, this.GetCthis(), p)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:241
// index:0
// virtual
// bool isIndexHidden(const class QModelIndex &)
func (this *QHeaderView) IsIndexHidden(index unsafe.Pointer) {
	// 0: (, index const QModelIndex &), (index)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QHeaderView13isIndexHiddenERK11QModelIndex", ffiqt.FFI_TYPE_VOID, this.GetCthis(), index)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:243
// index:0
// virtual
// QModelIndex moveCursor(enum QAbstractItemView::CursorAction, Qt::KeyboardModifiers)
func (this *QHeaderView) MoveCursor(arg0 int, arg1 int) {
	// 0: (, QAbstractItemView::CursorAction arg0, Qt::KeyboardModifiers arg1), (&arg0, &arg1)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QHeaderView10moveCursorEN17QAbstractItemView12CursorActionE6QFlagsIN2Qt16KeyboardModifierEE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &arg0, &arg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:244
// index:0
// virtual
// void setSelection(const class QRect &, class QItemSelectionModel::SelectionFlags)
func (this *QHeaderView) SetSelection(rect unsafe.Pointer, flags int) {
	// 0: (, rect const QRect &, QFlags<QItemSelectionModel::SelectionFlag> flags), (rect, &flags)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QHeaderView12setSelectionERK5QRect6QFlagsIN19QItemSelectionModel13SelectionFlagEE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), rect, &flags)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:245
// index:0
// virtual
// QRegion visualRegionForSelection(const class QItemSelection &)
func (this *QHeaderView) VisualRegionForSelection(selection unsafe.Pointer) {
	// 0: (, selection const QItemSelection &), (selection)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QHeaderView24visualRegionForSelectionERK14QItemSelection", ffiqt.FFI_TYPE_VOID, this.GetCthis(), selection)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:246
// index:0
// void initStyleOption(class QStyleOptionHeader *)
func (this *QHeaderView) InitStyleOption(option unsafe.Pointer) {
	// 0: (, option QStyleOptionHeader *), (option)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QHeaderView15initStyleOptionEP18QStyleOptionHeader", ffiqt.FFI_TYPE_VOID, this.GetCthis(), option)
	gopp.ErrPrint(err, rv)
}

//  body block end
