//  header block begin
// /usr/include/qt/QtWidgets/qheaderview.h
// #include <qheaderview.h>
// #include <QtWidgets>
package qtwidgets

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 17
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
	cthis unsafe.Pointer
}

// /usr/include/qt/QtWidgets/qheaderview.h:55
// index:0
// virtual
// const QMetaObject * metaObject()
func (this *QHeaderView) MetaObject() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QHeaderView10metaObjectEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:77
// index:0
// void QHeaderView(Qt::Orientation, class QWidget *)
func NewQHeaderView(orientation int, parent unsafe.Pointer) *QHeaderView {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QHeaderViewC2EN2Qt11OrientationEP7QWidget", ffiqt.FFI_TYPE_VOID, cthis, &orientation, parent)
	gopp.ErrPrint(err, rv)
	return &QHeaderView{cthis}
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
	// 0: (, QAbstractItemModel * model), (model)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QHeaderView8setModelEP18QAbstractItemModel", ffiqt.FFI_TYPE_VOID, this.cthis, model)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:82
// index:0
// Qt::Orientation orientation()
func (this *QHeaderView) Orientation() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QHeaderView11orientationEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:83
// index:0
// int offset()
func (this *QHeaderView) Offset() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QHeaderView6offsetEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:84
// index:0
// int length()
func (this *QHeaderView) Length() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QHeaderView6lengthEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:85
// index:0
// virtual
// QSize sizeHint()
func (this *QHeaderView) SizeHint() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QHeaderView8sizeHintEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:86
// index:0
// virtual
// void setVisible(_Bool)
func (this *QHeaderView) SetVisible(v bool) {
	// 0: (, bool v), (&v)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QHeaderView10setVisibleEb", ffiqt.FFI_TYPE_VOID, this.cthis, &v)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:87
// index:0
// int sectionSizeHint(int)
func (this *QHeaderView) SectionSizeHint(logicalIndex int) {
	// 0: (, int logicalIndex), (&logicalIndex)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QHeaderView15sectionSizeHintEi", ffiqt.FFI_TYPE_VOID, this.cthis, &logicalIndex)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:89
// index:0
// int visualIndexAt(int)
func (this *QHeaderView) VisualIndexAt(position int) {
	// 0: (, int position), (&position)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QHeaderView13visualIndexAtEi", ffiqt.FFI_TYPE_VOID, this.cthis, &position)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:90
// index:0
// int logicalIndexAt(int)
func (this *QHeaderView) LogicalIndexAt(position int) {
	// 0: (, int position), (&position)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QHeaderView14logicalIndexAtEi", ffiqt.FFI_TYPE_VOID, this.cthis, &position)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:92
// index:1
// inline
// int logicalIndexAt(int, int)
func (this *QHeaderView) LogicalIndexAt_1(x int, y int) {
	// 1: (, int x, int y), (&x, &y)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QHeaderView14logicalIndexAtEii", ffiqt.FFI_TYPE_VOID, this.cthis, &x, &y)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:93
// index:2
// inline
// int logicalIndexAt(const class QPoint &)
func (this *QHeaderView) LogicalIndexAt_2(pos unsafe.Pointer) {
	// 2: (, const QPoint & pos), (pos)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QHeaderView14logicalIndexAtERK6QPoint", ffiqt.FFI_TYPE_VOID, this.cthis, pos)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:95
// index:0
// int sectionSize(int)
func (this *QHeaderView) SectionSize(logicalIndex int) {
	// 0: (, int logicalIndex), (&logicalIndex)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QHeaderView11sectionSizeEi", ffiqt.FFI_TYPE_VOID, this.cthis, &logicalIndex)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:96
// index:0
// int sectionPosition(int)
func (this *QHeaderView) SectionPosition(logicalIndex int) {
	// 0: (, int logicalIndex), (&logicalIndex)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QHeaderView15sectionPositionEi", ffiqt.FFI_TYPE_VOID, this.cthis, &logicalIndex)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:97
// index:0
// int sectionViewportPosition(int)
func (this *QHeaderView) SectionViewportPosition(logicalIndex int) {
	// 0: (, int logicalIndex), (&logicalIndex)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QHeaderView23sectionViewportPositionEi", ffiqt.FFI_TYPE_VOID, this.cthis, &logicalIndex)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:99
// index:0
// void moveSection(int, int)
func (this *QHeaderView) MoveSection(from int, to int) {
	// 0: (, int from, int to), (&from, &to)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QHeaderView11moveSectionEii", ffiqt.FFI_TYPE_VOID, this.cthis, &from, &to)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:100
// index:0
// void swapSections(int, int)
func (this *QHeaderView) SwapSections(first int, second int) {
	// 0: (, int first, int second), (&first, &second)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QHeaderView12swapSectionsEii", ffiqt.FFI_TYPE_VOID, this.cthis, &first, &second)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:101
// index:0
// void resizeSection(int, int)
func (this *QHeaderView) ResizeSection(logicalIndex int, size int) {
	// 0: (, int logicalIndex, int size), (&logicalIndex, &size)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QHeaderView13resizeSectionEii", ffiqt.FFI_TYPE_VOID, this.cthis, &logicalIndex, &size)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:102
// index:0
// void resizeSections(class QHeaderView::ResizeMode)
func (this *QHeaderView) ResizeSections(mode int) {
	// 0: (, QHeaderView::ResizeMode mode), (&mode)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QHeaderView14resizeSectionsENS_10ResizeModeE", ffiqt.FFI_TYPE_VOID, this.cthis, &mode)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:104
// index:0
// bool isSectionHidden(int)
func (this *QHeaderView) IsSectionHidden(logicalIndex int) {
	// 0: (, int logicalIndex), (&logicalIndex)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QHeaderView15isSectionHiddenEi", ffiqt.FFI_TYPE_VOID, this.cthis, &logicalIndex)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:105
// index:0
// void setSectionHidden(int, _Bool)
func (this *QHeaderView) SetSectionHidden(logicalIndex int, hide bool) {
	// 0: (, int logicalIndex, bool hide), (&logicalIndex, &hide)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QHeaderView16setSectionHiddenEib", ffiqt.FFI_TYPE_VOID, this.cthis, &logicalIndex, &hide)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:106
// index:0
// int hiddenSectionCount()
func (this *QHeaderView) HiddenSectionCount() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QHeaderView18hiddenSectionCountEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:108
// index:0
// inline
// void hideSection(int)
func (this *QHeaderView) HideSection(logicalIndex int) {
	// 0: (, int logicalIndex), (&logicalIndex)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QHeaderView11hideSectionEi", ffiqt.FFI_TYPE_VOID, this.cthis, &logicalIndex)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:109
// index:0
// inline
// void showSection(int)
func (this *QHeaderView) ShowSection(logicalIndex int) {
	// 0: (, int logicalIndex), (&logicalIndex)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QHeaderView11showSectionEi", ffiqt.FFI_TYPE_VOID, this.cthis, &logicalIndex)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:111
// index:0
// int count()
func (this *QHeaderView) Count() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QHeaderView5countEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:112
// index:0
// int visualIndex(int)
func (this *QHeaderView) VisualIndex(logicalIndex int) {
	// 0: (, int logicalIndex), (&logicalIndex)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QHeaderView11visualIndexEi", ffiqt.FFI_TYPE_VOID, this.cthis, &logicalIndex)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:113
// index:0
// int logicalIndex(int)
func (this *QHeaderView) LogicalIndex(visualIndex int) {
	// 0: (, int visualIndex), (&visualIndex)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QHeaderView12logicalIndexEi", ffiqt.FFI_TYPE_VOID, this.cthis, &visualIndex)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:115
// index:0
// void setSectionsMovable(_Bool)
func (this *QHeaderView) SetSectionsMovable(movable bool) {
	// 0: (, bool movable), (&movable)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QHeaderView18setSectionsMovableEb", ffiqt.FFI_TYPE_VOID, this.cthis, &movable)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:116
// index:0
// bool sectionsMovable()
func (this *QHeaderView) SectionsMovable() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QHeaderView15sectionsMovableEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:122
// index:0
// void setSectionsClickable(_Bool)
func (this *QHeaderView) SetSectionsClickable(clickable bool) {
	// 0: (, bool clickable), (&clickable)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QHeaderView20setSectionsClickableEb", ffiqt.FFI_TYPE_VOID, this.cthis, &clickable)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:123
// index:0
// bool sectionsClickable()
func (this *QHeaderView) SectionsClickable() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QHeaderView17sectionsClickableEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:129
// index:0
// void setHighlightSections(_Bool)
func (this *QHeaderView) SetHighlightSections(highlight bool) {
	// 0: (, bool highlight), (&highlight)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QHeaderView20setHighlightSectionsEb", ffiqt.FFI_TYPE_VOID, this.cthis, &highlight)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:130
// index:0
// bool highlightSections()
func (this *QHeaderView) HighlightSections() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QHeaderView17highlightSectionsEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:132
// index:0
// QHeaderView::ResizeMode sectionResizeMode(int)
func (this *QHeaderView) SectionResizeMode(logicalIndex int) {
	// 0: (, int logicalIndex), (&logicalIndex)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QHeaderView17sectionResizeModeEi", ffiqt.FFI_TYPE_VOID, this.cthis, &logicalIndex)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:133
// index:0
// void setSectionResizeMode(enum QHeaderView::ResizeMode)
func (this *QHeaderView) SetSectionResizeMode(mode int) {
	// 0: (, QHeaderView::ResizeMode mode), (&mode)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QHeaderView20setSectionResizeModeENS_10ResizeModeE", ffiqt.FFI_TYPE_VOID, this.cthis, &mode)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:134
// index:1
// void setSectionResizeMode(int, enum QHeaderView::ResizeMode)
func (this *QHeaderView) SetSectionResizeMode_1(logicalIndex int, mode int) {
	// 1: (, int logicalIndex, QHeaderView::ResizeMode mode), (&logicalIndex, &mode)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QHeaderView20setSectionResizeModeEiNS_10ResizeModeE", ffiqt.FFI_TYPE_VOID, this.cthis, &logicalIndex, &mode)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:136
// index:0
// void setResizeContentsPrecision(int)
func (this *QHeaderView) SetResizeContentsPrecision(precision int) {
	// 0: (, int precision), (&precision)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QHeaderView26setResizeContentsPrecisionEi", ffiqt.FFI_TYPE_VOID, this.cthis, &precision)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:137
// index:0
// int resizeContentsPrecision()
func (this *QHeaderView) ResizeContentsPrecision() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QHeaderView23resizeContentsPrecisionEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:148
// index:0
// int stretchSectionCount()
func (this *QHeaderView) StretchSectionCount() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QHeaderView19stretchSectionCountEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:150
// index:0
// void setSortIndicatorShown(_Bool)
func (this *QHeaderView) SetSortIndicatorShown(show bool) {
	// 0: (, bool show), (&show)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QHeaderView21setSortIndicatorShownEb", ffiqt.FFI_TYPE_VOID, this.cthis, &show)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:151
// index:0
// bool isSortIndicatorShown()
func (this *QHeaderView) IsSortIndicatorShown() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QHeaderView20isSortIndicatorShownEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:153
// index:0
// void setSortIndicator(int, Qt::SortOrder)
func (this *QHeaderView) SetSortIndicator(logicalIndex int, order int) {
	// 0: (, int logicalIndex, Qt::SortOrder order), (&logicalIndex, &order)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QHeaderView16setSortIndicatorEiN2Qt9SortOrderE", ffiqt.FFI_TYPE_VOID, this.cthis, &logicalIndex, &order)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:154
// index:0
// int sortIndicatorSection()
func (this *QHeaderView) SortIndicatorSection() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QHeaderView20sortIndicatorSectionEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:155
// index:0
// Qt::SortOrder sortIndicatorOrder()
func (this *QHeaderView) SortIndicatorOrder() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QHeaderView18sortIndicatorOrderEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:157
// index:0
// bool stretchLastSection()
func (this *QHeaderView) StretchLastSection() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QHeaderView18stretchLastSectionEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:158
// index:0
// void setStretchLastSection(_Bool)
func (this *QHeaderView) SetStretchLastSection(stretch bool) {
	// 0: (, bool stretch), (&stretch)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QHeaderView21setStretchLastSectionEb", ffiqt.FFI_TYPE_VOID, this.cthis, &stretch)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:160
// index:0
// bool cascadingSectionResizes()
func (this *QHeaderView) CascadingSectionResizes() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QHeaderView23cascadingSectionResizesEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:161
// index:0
// void setCascadingSectionResizes(_Bool)
func (this *QHeaderView) SetCascadingSectionResizes(enable bool) {
	// 0: (, bool enable), (&enable)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QHeaderView26setCascadingSectionResizesEb", ffiqt.FFI_TYPE_VOID, this.cthis, &enable)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:163
// index:0
// int defaultSectionSize()
func (this *QHeaderView) DefaultSectionSize() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QHeaderView18defaultSectionSizeEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:164
// index:0
// void setDefaultSectionSize(int)
func (this *QHeaderView) SetDefaultSectionSize(size int) {
	// 0: (, int size), (&size)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QHeaderView21setDefaultSectionSizeEi", ffiqt.FFI_TYPE_VOID, this.cthis, &size)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:165
// index:0
// void resetDefaultSectionSize()
func (this *QHeaderView) ResetDefaultSectionSize() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QHeaderView23resetDefaultSectionSizeEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:167
// index:0
// int minimumSectionSize()
func (this *QHeaderView) MinimumSectionSize() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QHeaderView18minimumSectionSizeEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:168
// index:0
// void setMinimumSectionSize(int)
func (this *QHeaderView) SetMinimumSectionSize(size int) {
	// 0: (, int size), (&size)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QHeaderView21setMinimumSectionSizeEi", ffiqt.FFI_TYPE_VOID, this.cthis, &size)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:169
// index:0
// int maximumSectionSize()
func (this *QHeaderView) MaximumSectionSize() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QHeaderView18maximumSectionSizeEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:170
// index:0
// void setMaximumSectionSize(int)
func (this *QHeaderView) SetMaximumSectionSize(size int) {
	// 0: (, int size), (&size)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QHeaderView21setMaximumSectionSizeEi", ffiqt.FFI_TYPE_VOID, this.cthis, &size)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:172
// index:0
// Qt::Alignment defaultAlignment()
func (this *QHeaderView) DefaultAlignment() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QHeaderView16defaultAlignmentEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:175
// index:0
// virtual
// void doItemsLayout()
func (this *QHeaderView) DoItemsLayout() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QHeaderView13doItemsLayoutEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:176
// index:0
// bool sectionsMoved()
func (this *QHeaderView) SectionsMoved() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QHeaderView13sectionsMovedEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:177
// index:0
// bool sectionsHidden()
func (this *QHeaderView) SectionsHidden() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QHeaderView14sectionsHiddenEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:180
// index:0
// QByteArray saveState()
func (this *QHeaderView) SaveState() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QHeaderView9saveStateEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:181
// index:0
// bool restoreState(const class QByteArray &)
func (this *QHeaderView) RestoreState(state unsafe.Pointer) {
	// 0: (, const QByteArray & state), (state)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QHeaderView12restoreStateERK10QByteArray", ffiqt.FFI_TYPE_VOID, this.cthis, state)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:184
// index:0
// virtual
// void reset()
func (this *QHeaderView) Reset() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QHeaderView5resetEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:187
// index:0
// void setOffset(int)
func (this *QHeaderView) SetOffset(offset int) {
	// 0: (, int offset), (&offset)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QHeaderView9setOffsetEi", ffiqt.FFI_TYPE_VOID, this.cthis, &offset)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:188
// index:0
// void setOffsetToSectionPosition(int)
func (this *QHeaderView) SetOffsetToSectionPosition(visualIndex int) {
	// 0: (, int visualIndex), (&visualIndex)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QHeaderView26setOffsetToSectionPositionEi", ffiqt.FFI_TYPE_VOID, this.cthis, &visualIndex)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:189
// index:0
// void setOffsetToLastSection()
func (this *QHeaderView) SetOffsetToLastSection() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QHeaderView22setOffsetToLastSectionEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:190
// index:0
// void headerDataChanged(Qt::Orientation, int, int)
func (this *QHeaderView) HeaderDataChanged(orientation int, logicalFirst int, logicalLast int) {
	// 0: (, Qt::Orientation orientation, int logicalFirst, int logicalLast), (&orientation, &logicalFirst, &logicalLast)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QHeaderView17headerDataChangedEN2Qt11OrientationEii", ffiqt.FFI_TYPE_VOID, this.cthis, &orientation, &logicalFirst, &logicalLast)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:193
// index:0
// void sectionMoved(int, int, int)
func (this *QHeaderView) SectionMoved(logicalIndex int, oldVisualIndex int, newVisualIndex int) {
	// 0: (, int logicalIndex, int oldVisualIndex, int newVisualIndex), (&logicalIndex, &oldVisualIndex, &newVisualIndex)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QHeaderView12sectionMovedEiii", ffiqt.FFI_TYPE_VOID, this.cthis, &logicalIndex, &oldVisualIndex, &newVisualIndex)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:194
// index:0
// void sectionResized(int, int, int)
func (this *QHeaderView) SectionResized(logicalIndex int, oldSize int, newSize int) {
	// 0: (, int logicalIndex, int oldSize, int newSize), (&logicalIndex, &oldSize, &newSize)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QHeaderView14sectionResizedEiii", ffiqt.FFI_TYPE_VOID, this.cthis, &logicalIndex, &oldSize, &newSize)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:195
// index:0
// void sectionPressed(int)
func (this *QHeaderView) SectionPressed(logicalIndex int) {
	// 0: (, int logicalIndex), (&logicalIndex)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QHeaderView14sectionPressedEi", ffiqt.FFI_TYPE_VOID, this.cthis, &logicalIndex)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:196
// index:0
// void sectionClicked(int)
func (this *QHeaderView) SectionClicked(logicalIndex int) {
	// 0: (, int logicalIndex), (&logicalIndex)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QHeaderView14sectionClickedEi", ffiqt.FFI_TYPE_VOID, this.cthis, &logicalIndex)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:197
// index:0
// void sectionEntered(int)
func (this *QHeaderView) SectionEntered(logicalIndex int) {
	// 0: (, int logicalIndex), (&logicalIndex)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QHeaderView14sectionEnteredEi", ffiqt.FFI_TYPE_VOID, this.cthis, &logicalIndex)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:198
// index:0
// void sectionDoubleClicked(int)
func (this *QHeaderView) SectionDoubleClicked(logicalIndex int) {
	// 0: (, int logicalIndex), (&logicalIndex)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QHeaderView20sectionDoubleClickedEi", ffiqt.FFI_TYPE_VOID, this.cthis, &logicalIndex)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:199
// index:0
// void sectionCountChanged(int, int)
func (this *QHeaderView) SectionCountChanged(oldCount int, newCount int) {
	// 0: (, int oldCount, int newCount), (&oldCount, &newCount)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QHeaderView19sectionCountChangedEii", ffiqt.FFI_TYPE_VOID, this.cthis, &oldCount, &newCount)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:200
// index:0
// void sectionHandleDoubleClicked(int)
func (this *QHeaderView) SectionHandleDoubleClicked(logicalIndex int) {
	// 0: (, int logicalIndex), (&logicalIndex)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QHeaderView26sectionHandleDoubleClickedEi", ffiqt.FFI_TYPE_VOID, this.cthis, &logicalIndex)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:201
// index:0
// void geometriesChanged()
func (this *QHeaderView) GeometriesChanged() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QHeaderView17geometriesChangedEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:202
// index:0
// void sortIndicatorChanged(int, Qt::SortOrder)
func (this *QHeaderView) SortIndicatorChanged(logicalIndex int, order int) {
	// 0: (, int logicalIndex, Qt::SortOrder order), (&logicalIndex, &order)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QHeaderView20sortIndicatorChangedEiN2Qt9SortOrderE", ffiqt.FFI_TYPE_VOID, this.cthis, &logicalIndex, &order)
	gopp.ErrPrint(err, rv)
}

//  body block end
