package qt5
// auto generated, do not modify.
// created: Sun Jan 31 12:22:27 2016
// src-file: /QtWidgets/qheaderview.h
// dst-file: /src/widgets/qheaderview.go
//

// header block begin =>


// <= header block end

// main block begin =>
// <= main block end

// use block begin =>
import "fmt"
import "reflect"
import "unsafe"
import "qtrt"
// <= use block end

// ext block begin =>

/*
#include <stdlib.h>
#include <stdbool.h>
#include <stdint.h>
#include <wchar.h>
#include <uchar.h>
  // proto:  bool QHeaderView::cascadingSectionResizes();
extern void C_ZNK11QHeaderView23cascadingSectionResizesEv(void* qthis); // 4
  // proto:  void QHeaderView::setResizeContentsPrecision(int precision);
extern void C_ZN11QHeaderView26setResizeContentsPrecisionEi(void* qthis, int32_t arg0); // 4
  // proto:  bool QHeaderView::sectionsHidden();
extern void C_ZNK11QHeaderView14sectionsHiddenEv(void* qthis); // 4
  // proto:  int QHeaderView::sectionSizeHint(int logicalIndex);
extern void C_ZNK11QHeaderView15sectionSizeHintEi(void* qthis, int32_t arg0); // 4
  // proto:  Qt::SortOrder QHeaderView::sortIndicatorOrder();
extern void C_ZNK11QHeaderView18sortIndicatorOrderEv(void* qthis); // 4
  // proto:  void QHeaderView::resetDefaultSectionSize();
extern void C_ZN11QHeaderView23resetDefaultSectionSizeEv(void* qthis); // 4
  // proto:  void QHeaderView::setOffset(int offset);
extern void C_ZN11QHeaderView9setOffsetEi(void* qthis, int32_t arg0); // 4
  // proto:  void QHeaderView::reset();
extern void C_ZN11QHeaderView5resetEv(void* qthis); // 4
  // proto:  int QHeaderView::sectionPosition(int logicalIndex);
extern void C_ZNK11QHeaderView15sectionPositionEi(void* qthis, int32_t arg0); // 4
  // proto:  int QHeaderView::logicalIndex(int visualIndex);
extern void C_ZNK11QHeaderView12logicalIndexEi(void* qthis, int32_t arg0); // 4
  // proto:  int QHeaderView::sectionSize(int logicalIndex);
extern void C_ZNK11QHeaderView11sectionSizeEi(void* qthis, int32_t arg0); // 4
  // proto:  const QMetaObject * QHeaderView::metaObject();
extern void C_ZNK11QHeaderView10metaObjectEv(void* qthis); // 4
  // proto:  void QHeaderView::hideSection(int logicalIndex);
extern void C_ZN11QHeaderView11hideSectionEi(void* qthis, int32_t arg0); // 2
  // proto:  void QHeaderView::setMaximumSectionSize(int size);
extern void C_ZN11QHeaderView21setMaximumSectionSizeEi(void* qthis, int32_t arg0); // 4
  // proto:  QHeaderView::ResizeMode QHeaderView::sectionResizeMode(int logicalIndex);
extern void C_ZNK11QHeaderView17sectionResizeModeEi(void* qthis, int32_t arg0); // 4
  // proto:  void QHeaderView::~QHeaderView();
extern void C_ZN11QHeaderViewD2Ev(void* qthis); // 4
  // proto:  bool QHeaderView::restoreState(const QByteArray & state);
extern void C_ZN11QHeaderView12restoreStateERK10QByteArray(void* qthis, void* arg0); // 4
  // proto:  void QHeaderView::setOffsetToSectionPosition(int visualIndex);
extern void C_ZN11QHeaderView26setOffsetToSectionPositionEi(void* qthis, int32_t arg0); // 4
  // proto:  void QHeaderView::doItemsLayout();
extern void C_ZN11QHeaderView13doItemsLayoutEv(void* qthis); // 4
  // proto:  void QHeaderView::setHighlightSections(bool highlight);
extern void C_ZN11QHeaderView20setHighlightSectionsEb(void* qthis, bool arg0); // 4
  // proto:  int QHeaderView::resizeContentsPrecision();
extern void C_ZNK11QHeaderView23resizeContentsPrecisionEv(void* qthis); // 4
  // proto:  bool QHeaderView::sectionsClickable();
extern void C_ZNK11QHeaderView17sectionsClickableEv(void* qthis); // 4
  // proto:  int QHeaderView::logicalIndexAt(const QPoint & pos);
extern void C_ZNK11QHeaderView14logicalIndexAtERK6QPoint(void* qthis, void* arg0); // 2
  // proto:  int QHeaderView::logicalIndexAt(int position);
extern void C_ZNK11QHeaderView14logicalIndexAtEi(void* qthis, int32_t arg0); // 4
  // proto:  int QHeaderView::logicalIndexAt(int x, int y);
extern void C_ZNK11QHeaderView14logicalIndexAtEii(void* qthis, int32_t arg0, int32_t arg1); // 2
  // proto:  int QHeaderView::offset();
extern void C_ZNK11QHeaderView6offsetEv(void* qthis); // 4
  // proto:  void QHeaderView::setMinimumSectionSize(int size);
extern void C_ZN11QHeaderView21setMinimumSectionSizeEi(void* qthis, int32_t arg0); // 4
  // proto:  QSize QHeaderView::sizeHint();
extern void C_ZNK11QHeaderView8sizeHintEv(void* qthis); // 4
  // proto:  void QHeaderView::setVisible(bool v);
extern void C_ZN11QHeaderView10setVisibleEb(void* qthis, bool arg0); // 4
  // proto:  void QHeaderView::setCascadingSectionResizes(bool enable);
extern void C_ZN11QHeaderView26setCascadingSectionResizesEb(void* qthis, bool arg0); // 4
  // proto:  bool QHeaderView::sectionsMovable();
extern void C_ZNK11QHeaderView15sectionsMovableEv(void* qthis); // 4
  // proto:  int QHeaderView::length();
extern void C_ZNK11QHeaderView6lengthEv(void* qthis); // 4
  // proto:  int QHeaderView::sortIndicatorSection();
extern void C_ZNK11QHeaderView20sortIndicatorSectionEv(void* qthis); // 4
  // proto:  void QHeaderView::resizeSection(int logicalIndex, int size);
extern void C_ZN11QHeaderView13resizeSectionEii(void* qthis, int32_t arg0, int32_t arg1); // 4
  // proto:  void QHeaderView::setSectionsMovable(bool movable);
extern void C_ZN11QHeaderView18setSectionsMovableEb(void* qthis, bool arg0); // 4
  // proto:  int QHeaderView::maximumSectionSize();
extern void C_ZNK11QHeaderView18maximumSectionSizeEv(void* qthis); // 4
  // proto:  void QHeaderView::moveSection(int from, int to);
extern void C_ZN11QHeaderView11moveSectionEii(void* qthis, int32_t arg0, int32_t arg1); // 4
  // proto:  bool QHeaderView::isSectionHidden(int logicalIndex);
extern void C_ZNK11QHeaderView15isSectionHiddenEi(void* qthis, int32_t arg0); // 4
  // proto:  int QHeaderView::visualIndexAt(int position);
extern void C_ZNK11QHeaderView13visualIndexAtEi(void* qthis, int32_t arg0); // 4
  // proto:  void QHeaderView::setSortIndicatorShown(bool show);
extern void C_ZN11QHeaderView21setSortIndicatorShownEb(void* qthis, bool arg0); // 4
  // proto:  Qt::Alignment QHeaderView::defaultAlignment();
extern void C_ZNK11QHeaderView16defaultAlignmentEv(void* qthis); // 4
  // proto:  int QHeaderView::hiddenSectionCount();
extern void C_ZNK11QHeaderView18hiddenSectionCountEv(void* qthis); // 4
  // proto:  void QHeaderView::setStretchLastSection(bool stretch);
extern void C_ZN11QHeaderView21setStretchLastSectionEb(void* qthis, bool arg0); // 4
  // proto:  int QHeaderView::count();
extern void C_ZNK11QHeaderView5countEv(void* qthis); // 4
  // proto:  int QHeaderView::minimumSectionSize();
extern void C_ZNK11QHeaderView18minimumSectionSizeEv(void* qthis); // 4
  // proto:  QByteArray QHeaderView::saveState();
extern void C_ZNK11QHeaderView9saveStateEv(void* qthis); // 4
  // proto:  Qt::Orientation QHeaderView::orientation();
extern void C_ZNK11QHeaderView11orientationEv(void* qthis); // 4
  // proto:  void QHeaderView::setSectionHidden(int logicalIndex, bool hide);
extern void C_ZN11QHeaderView16setSectionHiddenEib(void* qthis, int32_t arg0, bool arg1); // 4
  // proto:  void QHeaderView::setOffsetToLastSection();
extern void C_ZN11QHeaderView22setOffsetToLastSectionEv(void* qthis); // 4
  // proto:  bool QHeaderView::stretchLastSection();
extern void C_ZNK11QHeaderView18stretchLastSectionEv(void* qthis); // 4
  // proto:  bool QHeaderView::isSortIndicatorShown();
extern void C_ZNK11QHeaderView20isSortIndicatorShownEv(void* qthis); // 4
  // proto:  void QHeaderView::showSection(int logicalIndex);
extern void C_ZN11QHeaderView11showSectionEi(void* qthis, int32_t arg0); // 2
  // proto:  void QHeaderView::swapSections(int first, int second);
extern void C_ZN11QHeaderView12swapSectionsEii(void* qthis, int32_t arg0, int32_t arg1); // 4
  // proto:  void QHeaderView::setSectionsClickable(bool clickable);
extern void C_ZN11QHeaderView20setSectionsClickableEb(void* qthis, bool arg0); // 4
  // proto:  void QHeaderView::setModel(QAbstractItemModel * model);
extern void C_ZN11QHeaderView8setModelEP18QAbstractItemModel(void* qthis, void* arg0); // 4
  // proto:  void QHeaderView::setDefaultSectionSize(int size);
extern void C_ZN11QHeaderView21setDefaultSectionSizeEi(void* qthis, int32_t arg0); // 4
  // proto:  bool QHeaderView::sectionsMoved();
extern void C_ZNK11QHeaderView13sectionsMovedEv(void* qthis); // 4
  // proto:  int QHeaderView::sectionViewportPosition(int logicalIndex);
extern void C_ZNK11QHeaderView23sectionViewportPositionEi(void* qthis, int32_t arg0); // 4
  // proto:  int QHeaderView::stretchSectionCount();
extern void C_ZNK11QHeaderView19stretchSectionCountEv(void* qthis); // 4
  // proto:  int QHeaderView::visualIndex(int logicalIndex);
extern void C_ZNK11QHeaderView11visualIndexEi(void* qthis, int32_t arg0); // 4
  // proto:  bool QHeaderView::highlightSections();
extern void C_ZNK11QHeaderView17highlightSectionsEv(void* qthis); // 4
  // proto:  int QHeaderView::defaultSectionSize();
extern void C_ZNK11QHeaderView18defaultSectionSizeEv(void* qthis); // 4
*/
import "C"
// } // <= ext block end

// body block begin =>
func init() {
  if false {qtrt.KeepMe()}
  if false {fmt.Println(123)}
  if false {reflect.TypeOf(123)}
  if false {reflect.TypeOf(unsafe.Sizeof(0))}
}

// class sizeof(QHeaderView)=1
type QHeaderView struct {
  /*qbase*/ QAbstractItemView;
  qclsinst unsafe.Pointer /* *C.void */;
//  _sectionHandleDoubleClicked QHeaderView_sectionHandleDoubleClicked_signal;
//  _sectionEntered QHeaderView_sectionEntered_signal;
//  _sortIndicatorChanged QHeaderView_sortIndicatorChanged_signal;
//  _sectionClicked QHeaderView_sectionClicked_signal;
//  _sectionCountChanged QHeaderView_sectionCountChanged_signal;
//  _geometriesChanged QHeaderView_geometriesChanged_signal;
//  _sectionMoved QHeaderView_sectionMoved_signal;
//  _sectionPressed QHeaderView_sectionPressed_signal;
//  _sectionDoubleClicked QHeaderView_sectionDoubleClicked_signal;
//  _sectionResized QHeaderView_sectionResized_signal;
}

// cascadingSectionResizes()
func (this *QHeaderView) cascadingSectionResizes(args ...interface{}) () {
  // cascadingSectionResizes()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QHeaderView23cascadingSectionResizesEv
    // invoke: bool cascadingSectionResizes()
    C.C_ZNK11QHeaderView23cascadingSectionResizesEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QHeaderView", "cascadingSectionResizes", args)
  }

}

// setResizeContentsPrecision(int)
func (this *QHeaderView) setResizeContentsPrecision(args ...interface{}) () {
  // setResizeContentsPrecision(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QHeaderView26setResizeContentsPrecisionEi
    // invoke: void setResizeContentsPrecision(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN11QHeaderView26setResizeContentsPrecisionEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QHeaderView", "setResizeContentsPrecision", args)
  }

}

// sectionsHidden()
func (this *QHeaderView) sectionsHidden(args ...interface{}) () {
  // sectionsHidden()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QHeaderView14sectionsHiddenEv
    // invoke: bool sectionsHidden()
    C.C_ZNK11QHeaderView14sectionsHiddenEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QHeaderView", "sectionsHidden", args)
  }

}

// sectionSizeHint(int)
func (this *QHeaderView) sectionSizeHint(args ...interface{}) () {
  // sectionSizeHint(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QHeaderView15sectionSizeHintEi
    // invoke: int sectionSizeHint(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.C_ZNK11QHeaderView15sectionSizeHintEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QHeaderView", "sectionSizeHint", args)
  }

}

// sortIndicatorOrder()
func (this *QHeaderView) sortIndicatorOrder(args ...interface{}) () {
  // sortIndicatorOrder()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QHeaderView18sortIndicatorOrderEv
    // invoke: Qt::SortOrder sortIndicatorOrder()
    C.C_ZNK11QHeaderView18sortIndicatorOrderEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QHeaderView", "sortIndicatorOrder", args)
  }

}

// resetDefaultSectionSize()
func (this *QHeaderView) resetDefaultSectionSize(args ...interface{}) () {
  // resetDefaultSectionSize()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QHeaderView23resetDefaultSectionSizeEv
    // invoke: void resetDefaultSectionSize()
    C.C_ZN11QHeaderView23resetDefaultSectionSizeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QHeaderView", "resetDefaultSectionSize", args)
  }

}

// setOffset(int)
func (this *QHeaderView) setOffset(args ...interface{}) () {
  // setOffset(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QHeaderView9setOffsetEi
    // invoke: void setOffset(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN11QHeaderView9setOffsetEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QHeaderView", "setOffset", args)
  }

}

// reset()
func (this *QHeaderView) reset(args ...interface{}) () {
  // reset()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QHeaderView5resetEv
    // invoke: void reset()
    C.C_ZN11QHeaderView5resetEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QHeaderView", "reset", args)
  }

}

// sectionPosition(int)
func (this *QHeaderView) sectionPosition(args ...interface{}) () {
  // sectionPosition(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QHeaderView15sectionPositionEi
    // invoke: int sectionPosition(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.C_ZNK11QHeaderView15sectionPositionEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QHeaderView", "sectionPosition", args)
  }

}

// logicalIndex(int)
func (this *QHeaderView) logicalIndex(args ...interface{}) () {
  // logicalIndex(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QHeaderView12logicalIndexEi
    // invoke: int logicalIndex(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.C_ZNK11QHeaderView12logicalIndexEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QHeaderView", "logicalIndex", args)
  }

}

// sectionSize(int)
func (this *QHeaderView) sectionSize(args ...interface{}) () {
  // sectionSize(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QHeaderView11sectionSizeEi
    // invoke: int sectionSize(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.C_ZNK11QHeaderView11sectionSizeEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QHeaderView", "sectionSize", args)
  }

}

// metaObject()
func (this *QHeaderView) metaObject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QHeaderView10metaObjectEv
    // invoke: const QMetaObject * metaObject()
    C.C_ZNK11QHeaderView10metaObjectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QHeaderView", "metaObject", args)
  }

}

// hideSection(int)
func (this *QHeaderView) hideSection(args ...interface{}) () {
  // hideSection(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QHeaderView11hideSectionEi
    // invoke: void hideSection(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN11QHeaderView11hideSectionEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QHeaderView", "hideSection", args)
  }

}

// setMaximumSectionSize(int)
func (this *QHeaderView) setMaximumSectionSize(args ...interface{}) () {
  // setMaximumSectionSize(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QHeaderView21setMaximumSectionSizeEi
    // invoke: void setMaximumSectionSize(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN11QHeaderView21setMaximumSectionSizeEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QHeaderView", "setMaximumSectionSize", args)
  }

}

// sectionResizeMode(int)
func (this *QHeaderView) sectionResizeMode(args ...interface{}) () {
  // sectionResizeMode(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QHeaderView17sectionResizeModeEi
    // invoke: QHeaderView::ResizeMode sectionResizeMode(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.C_ZNK11QHeaderView17sectionResizeModeEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QHeaderView", "sectionResizeMode", args)
  }

}

// ~QHeaderView()
func (this *QHeaderView) FreeQHeaderView(args ...interface{}) () {
  // ~QHeaderView()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QHeaderViewD0Ev
    // invoke: void ~QHeaderView()
    C.C_ZN11QHeaderViewD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QHeaderView", "~QHeaderView", args)
  }

}

// restoreState(const class QByteArray &)
func (this *QHeaderView) restoreState(args ...interface{}) () {
  // restoreState(const class QByteArray &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QByteArray{}) // "const QByteArray &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QHeaderView12restoreStateERK10QByteArray
    // invoke: bool restoreState(const class QByteArray &)
    var arg0 = args[0].(QByteArray).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN11QHeaderView12restoreStateERK10QByteArray(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QHeaderView", "restoreState", args)
  }

}

// setOffsetToSectionPosition(int)
func (this *QHeaderView) setOffsetToSectionPosition(args ...interface{}) () {
  // setOffsetToSectionPosition(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QHeaderView26setOffsetToSectionPositionEi
    // invoke: void setOffsetToSectionPosition(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN11QHeaderView26setOffsetToSectionPositionEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QHeaderView", "setOffsetToSectionPosition", args)
  }

}

// doItemsLayout()
func (this *QHeaderView) doItemsLayout(args ...interface{}) () {
  // doItemsLayout()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QHeaderView13doItemsLayoutEv
    // invoke: void doItemsLayout()
    C.C_ZN11QHeaderView13doItemsLayoutEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QHeaderView", "doItemsLayout", args)
  }

}

// setHighlightSections(_Bool)
func (this *QHeaderView) setHighlightSections(args ...interface{}) () {
  // setHighlightSections(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QHeaderView20setHighlightSectionsEb
    // invoke: void setHighlightSections(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C.C_ZN11QHeaderView20setHighlightSectionsEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QHeaderView", "setHighlightSections", args)
  }

}

// resizeContentsPrecision()
func (this *QHeaderView) resizeContentsPrecision(args ...interface{}) () {
  // resizeContentsPrecision()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QHeaderView23resizeContentsPrecisionEv
    // invoke: int resizeContentsPrecision()
    C.C_ZNK11QHeaderView23resizeContentsPrecisionEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QHeaderView", "resizeContentsPrecision", args)
  }

}

// sectionsClickable()
func (this *QHeaderView) sectionsClickable(args ...interface{}) () {
  // sectionsClickable()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QHeaderView17sectionsClickableEv
    // invoke: bool sectionsClickable()
    C.C_ZNK11QHeaderView17sectionsClickableEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QHeaderView", "sectionsClickable", args)
  }

}

// logicalIndexAt(const class QPoint &)
func (this *QHeaderView) logicalIndexAt(args ...interface{}) () {
  // logicalIndexAt(const class QPoint &)
  // logicalIndexAt(int)
  // logicalIndexAt(int, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPoint{}) // "const QPoint &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.Int32Ty(false) // "int"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = qtrt.Int32Ty(false) // "int"
  vtys[2][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QHeaderView14logicalIndexAtERK6QPoint
    // invoke: int logicalIndexAt(const class QPoint &)
    var arg0 = args[0].(QPoint).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZNK11QHeaderView14logicalIndexAtERK6QPoint(this.qclsinst, arg0)
  case 1:
    // invoke: _ZNK11QHeaderView14logicalIndexAtEi
    // invoke: int logicalIndexAt(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.C_ZNK11QHeaderView14logicalIndexAtEi(this.qclsinst, arg0)
  case 2:
    // invoke: _ZNK11QHeaderView14logicalIndexAtEii
    // invoke: int logicalIndexAt(int, int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    C.C_ZNK11QHeaderView14logicalIndexAtEii(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QHeaderView", "logicalIndexAt", args)
  }

}

// offset()
func (this *QHeaderView) offset(args ...interface{}) () {
  // offset()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QHeaderView6offsetEv
    // invoke: int offset()
    C.C_ZNK11QHeaderView6offsetEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QHeaderView", "offset", args)
  }

}

// setMinimumSectionSize(int)
func (this *QHeaderView) setMinimumSectionSize(args ...interface{}) () {
  // setMinimumSectionSize(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QHeaderView21setMinimumSectionSizeEi
    // invoke: void setMinimumSectionSize(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN11QHeaderView21setMinimumSectionSizeEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QHeaderView", "setMinimumSectionSize", args)
  }

}

// sizeHint()
func (this *QHeaderView) sizeHint(args ...interface{}) () {
  // sizeHint()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QHeaderView8sizeHintEv
    // invoke: QSize sizeHint()
    C.C_ZNK11QHeaderView8sizeHintEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QHeaderView", "sizeHint", args)
  }

}

// setVisible(_Bool)
func (this *QHeaderView) setVisible(args ...interface{}) () {
  // setVisible(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QHeaderView10setVisibleEb
    // invoke: void setVisible(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C.C_ZN11QHeaderView10setVisibleEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QHeaderView", "setVisible", args)
  }

}

// setCascadingSectionResizes(_Bool)
func (this *QHeaderView) setCascadingSectionResizes(args ...interface{}) () {
  // setCascadingSectionResizes(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QHeaderView26setCascadingSectionResizesEb
    // invoke: void setCascadingSectionResizes(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C.C_ZN11QHeaderView26setCascadingSectionResizesEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QHeaderView", "setCascadingSectionResizes", args)
  }

}

// sectionsMovable()
func (this *QHeaderView) sectionsMovable(args ...interface{}) () {
  // sectionsMovable()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QHeaderView15sectionsMovableEv
    // invoke: bool sectionsMovable()
    C.C_ZNK11QHeaderView15sectionsMovableEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QHeaderView", "sectionsMovable", args)
  }

}

// length()
func (this *QHeaderView) length(args ...interface{}) () {
  // length()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QHeaderView6lengthEv
    // invoke: int length()
    C.C_ZNK11QHeaderView6lengthEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QHeaderView", "length", args)
  }

}

// sortIndicatorSection()
func (this *QHeaderView) sortIndicatorSection(args ...interface{}) () {
  // sortIndicatorSection()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QHeaderView20sortIndicatorSectionEv
    // invoke: int sortIndicatorSection()
    C.C_ZNK11QHeaderView20sortIndicatorSectionEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QHeaderView", "sortIndicatorSection", args)
  }

}

// resizeSection(int, int)
func (this *QHeaderView) resizeSection(args ...interface{}) () {
  // resizeSection(int, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QHeaderView13resizeSectionEii
    // invoke: void resizeSection(int, int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    C.C_ZN11QHeaderView13resizeSectionEii(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QHeaderView", "resizeSection", args)
  }

}

// setSectionsMovable(_Bool)
func (this *QHeaderView) setSectionsMovable(args ...interface{}) () {
  // setSectionsMovable(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QHeaderView18setSectionsMovableEb
    // invoke: void setSectionsMovable(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C.C_ZN11QHeaderView18setSectionsMovableEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QHeaderView", "setSectionsMovable", args)
  }

}

// maximumSectionSize()
func (this *QHeaderView) maximumSectionSize(args ...interface{}) () {
  // maximumSectionSize()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QHeaderView18maximumSectionSizeEv
    // invoke: int maximumSectionSize()
    C.C_ZNK11QHeaderView18maximumSectionSizeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QHeaderView", "maximumSectionSize", args)
  }

}

// moveSection(int, int)
func (this *QHeaderView) moveSection(args ...interface{}) () {
  // moveSection(int, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QHeaderView11moveSectionEii
    // invoke: void moveSection(int, int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    C.C_ZN11QHeaderView11moveSectionEii(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QHeaderView", "moveSection", args)
  }

}

// isSectionHidden(int)
func (this *QHeaderView) isSectionHidden(args ...interface{}) () {
  // isSectionHidden(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QHeaderView15isSectionHiddenEi
    // invoke: bool isSectionHidden(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.C_ZNK11QHeaderView15isSectionHiddenEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QHeaderView", "isSectionHidden", args)
  }

}

// visualIndexAt(int)
func (this *QHeaderView) visualIndexAt(args ...interface{}) () {
  // visualIndexAt(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QHeaderView13visualIndexAtEi
    // invoke: int visualIndexAt(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.C_ZNK11QHeaderView13visualIndexAtEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QHeaderView", "visualIndexAt", args)
  }

}

// setSortIndicatorShown(_Bool)
func (this *QHeaderView) setSortIndicatorShown(args ...interface{}) () {
  // setSortIndicatorShown(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QHeaderView21setSortIndicatorShownEb
    // invoke: void setSortIndicatorShown(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C.C_ZN11QHeaderView21setSortIndicatorShownEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QHeaderView", "setSortIndicatorShown", args)
  }

}

// defaultAlignment()
func (this *QHeaderView) defaultAlignment(args ...interface{}) () {
  // defaultAlignment()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QHeaderView16defaultAlignmentEv
    // invoke: Qt::Alignment defaultAlignment()
    C.C_ZNK11QHeaderView16defaultAlignmentEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QHeaderView", "defaultAlignment", args)
  }

}

// hiddenSectionCount()
func (this *QHeaderView) hiddenSectionCount(args ...interface{}) () {
  // hiddenSectionCount()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QHeaderView18hiddenSectionCountEv
    // invoke: int hiddenSectionCount()
    C.C_ZNK11QHeaderView18hiddenSectionCountEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QHeaderView", "hiddenSectionCount", args)
  }

}

// setStretchLastSection(_Bool)
func (this *QHeaderView) setStretchLastSection(args ...interface{}) () {
  // setStretchLastSection(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QHeaderView21setStretchLastSectionEb
    // invoke: void setStretchLastSection(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C.C_ZN11QHeaderView21setStretchLastSectionEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QHeaderView", "setStretchLastSection", args)
  }

}

// count()
func (this *QHeaderView) count(args ...interface{}) () {
  // count()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QHeaderView5countEv
    // invoke: int count()
    C.C_ZNK11QHeaderView5countEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QHeaderView", "count", args)
  }

}

// minimumSectionSize()
func (this *QHeaderView) minimumSectionSize(args ...interface{}) () {
  // minimumSectionSize()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QHeaderView18minimumSectionSizeEv
    // invoke: int minimumSectionSize()
    C.C_ZNK11QHeaderView18minimumSectionSizeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QHeaderView", "minimumSectionSize", args)
  }

}

// saveState()
func (this *QHeaderView) saveState(args ...interface{}) () {
  // saveState()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QHeaderView9saveStateEv
    // invoke: QByteArray saveState()
    C.C_ZNK11QHeaderView9saveStateEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QHeaderView", "saveState", args)
  }

}

// orientation()
func (this *QHeaderView) orientation(args ...interface{}) () {
  // orientation()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QHeaderView11orientationEv
    // invoke: Qt::Orientation orientation()
    C.C_ZNK11QHeaderView11orientationEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QHeaderView", "orientation", args)
  }

}

// setSectionHidden(int, _Bool)
func (this *QHeaderView) setSectionHidden(args ...interface{}) () {
  // setSectionHidden(int, _Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QHeaderView16setSectionHiddenEib
    // invoke: void setSectionHidden(int, _Bool)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.bool(args[1].(bool))
    if false {fmt.Println(arg1)}
    C.C_ZN11QHeaderView16setSectionHiddenEib(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QHeaderView", "setSectionHidden", args)
  }

}

// setOffsetToLastSection()
func (this *QHeaderView) setOffsetToLastSection(args ...interface{}) () {
  // setOffsetToLastSection()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QHeaderView22setOffsetToLastSectionEv
    // invoke: void setOffsetToLastSection()
    C.C_ZN11QHeaderView22setOffsetToLastSectionEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QHeaderView", "setOffsetToLastSection", args)
  }

}

// stretchLastSection()
func (this *QHeaderView) stretchLastSection(args ...interface{}) () {
  // stretchLastSection()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QHeaderView18stretchLastSectionEv
    // invoke: bool stretchLastSection()
    C.C_ZNK11QHeaderView18stretchLastSectionEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QHeaderView", "stretchLastSection", args)
  }

}

// isSortIndicatorShown()
func (this *QHeaderView) isSortIndicatorShown(args ...interface{}) () {
  // isSortIndicatorShown()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QHeaderView20isSortIndicatorShownEv
    // invoke: bool isSortIndicatorShown()
    C.C_ZNK11QHeaderView20isSortIndicatorShownEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QHeaderView", "isSortIndicatorShown", args)
  }

}

// showSection(int)
func (this *QHeaderView) showSection(args ...interface{}) () {
  // showSection(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QHeaderView11showSectionEi
    // invoke: void showSection(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN11QHeaderView11showSectionEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QHeaderView", "showSection", args)
  }

}

// swapSections(int, int)
func (this *QHeaderView) swapSections(args ...interface{}) () {
  // swapSections(int, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QHeaderView12swapSectionsEii
    // invoke: void swapSections(int, int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    C.C_ZN11QHeaderView12swapSectionsEii(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QHeaderView", "swapSections", args)
  }

}

// setSectionsClickable(_Bool)
func (this *QHeaderView) setSectionsClickable(args ...interface{}) () {
  // setSectionsClickable(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QHeaderView20setSectionsClickableEb
    // invoke: void setSectionsClickable(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C.C_ZN11QHeaderView20setSectionsClickableEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QHeaderView", "setSectionsClickable", args)
  }

}

// setModel(class QAbstractItemModel *)
func (this *QHeaderView) setModel(args ...interface{}) () {
  // setModel(class QAbstractItemModel *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QAbstractItemModel{}) // "QAbstractItemModel *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QHeaderView8setModelEP18QAbstractItemModel
    // invoke: void setModel(class QAbstractItemModel *)
    var arg0 = args[0].(QAbstractItemModel).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN11QHeaderView8setModelEP18QAbstractItemModel(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QHeaderView", "setModel", args)
  }

}

// setDefaultSectionSize(int)
func (this *QHeaderView) setDefaultSectionSize(args ...interface{}) () {
  // setDefaultSectionSize(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QHeaderView21setDefaultSectionSizeEi
    // invoke: void setDefaultSectionSize(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN11QHeaderView21setDefaultSectionSizeEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QHeaderView", "setDefaultSectionSize", args)
  }

}

// sectionsMoved()
func (this *QHeaderView) sectionsMoved(args ...interface{}) () {
  // sectionsMoved()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QHeaderView13sectionsMovedEv
    // invoke: bool sectionsMoved()
    C.C_ZNK11QHeaderView13sectionsMovedEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QHeaderView", "sectionsMoved", args)
  }

}

// sectionViewportPosition(int)
func (this *QHeaderView) sectionViewportPosition(args ...interface{}) () {
  // sectionViewportPosition(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QHeaderView23sectionViewportPositionEi
    // invoke: int sectionViewportPosition(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.C_ZNK11QHeaderView23sectionViewportPositionEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QHeaderView", "sectionViewportPosition", args)
  }

}

// stretchSectionCount()
func (this *QHeaderView) stretchSectionCount(args ...interface{}) () {
  // stretchSectionCount()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QHeaderView19stretchSectionCountEv
    // invoke: int stretchSectionCount()
    C.C_ZNK11QHeaderView19stretchSectionCountEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QHeaderView", "stretchSectionCount", args)
  }

}

// visualIndex(int)
func (this *QHeaderView) visualIndex(args ...interface{}) () {
  // visualIndex(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QHeaderView11visualIndexEi
    // invoke: int visualIndex(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.C_ZNK11QHeaderView11visualIndexEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QHeaderView", "visualIndex", args)
  }

}

// highlightSections()
func (this *QHeaderView) highlightSections(args ...interface{}) () {
  // highlightSections()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QHeaderView17highlightSectionsEv
    // invoke: bool highlightSections()
    C.C_ZNK11QHeaderView17highlightSectionsEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QHeaderView", "highlightSections", args)
  }

}

// defaultSectionSize()
func (this *QHeaderView) defaultSectionSize(args ...interface{}) () {
  // defaultSectionSize()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QHeaderView18defaultSectionSizeEv
    // invoke: int defaultSectionSize()
    C.C_ZNK11QHeaderView18defaultSectionSizeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QHeaderView", "defaultSectionSize", args)
  }

}

// <= body block end

