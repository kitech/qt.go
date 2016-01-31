package qt5
// auto generated, do not modify.
// created: Sun Jan 31 23:40:52 2016
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
extern bool C_ZNK11QHeaderView23cascadingSectionResizesEv(void* qthis); // 4
  // proto:  void QHeaderView::setResizeContentsPrecision(int precision);
extern void C_ZN11QHeaderView26setResizeContentsPrecisionEi(void* qthis, int32_t arg0); // 4
  // proto:  bool QHeaderView::sectionsHidden();
extern bool C_ZNK11QHeaderView14sectionsHiddenEv(void* qthis); // 4
  // proto:  int QHeaderView::sectionSizeHint(int logicalIndex);
extern int32_t C_ZNK11QHeaderView15sectionSizeHintEi(void* qthis, int32_t arg0); // 4
  // proto:  Qt::SortOrder QHeaderView::sortIndicatorOrder();
extern void C_ZNK11QHeaderView18sortIndicatorOrderEv(void* qthis); // 4
  // proto:  void QHeaderView::resetDefaultSectionSize();
extern void C_ZN11QHeaderView23resetDefaultSectionSizeEv(void* qthis); // 4
  // proto:  void QHeaderView::setOffset(int offset);
extern void C_ZN11QHeaderView9setOffsetEi(void* qthis, int32_t arg0); // 4
  // proto:  void QHeaderView::reset();
extern void C_ZN11QHeaderView5resetEv(void* qthis); // 4
  // proto:  int QHeaderView::sectionPosition(int logicalIndex);
extern int32_t C_ZNK11QHeaderView15sectionPositionEi(void* qthis, int32_t arg0); // 4
  // proto:  int QHeaderView::logicalIndex(int visualIndex);
extern int32_t C_ZNK11QHeaderView12logicalIndexEi(void* qthis, int32_t arg0); // 4
  // proto:  int QHeaderView::sectionSize(int logicalIndex);
extern int32_t C_ZNK11QHeaderView11sectionSizeEi(void* qthis, int32_t arg0); // 4
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
extern bool C_ZN11QHeaderView12restoreStateERK10QByteArray(void* qthis, void* arg0); // 4
  // proto:  void QHeaderView::setOffsetToSectionPosition(int visualIndex);
extern void C_ZN11QHeaderView26setOffsetToSectionPositionEi(void* qthis, int32_t arg0); // 4
  // proto:  void QHeaderView::doItemsLayout();
extern void C_ZN11QHeaderView13doItemsLayoutEv(void* qthis); // 4
  // proto:  void QHeaderView::setHighlightSections(bool highlight);
extern void C_ZN11QHeaderView20setHighlightSectionsEb(void* qthis, bool arg0); // 4
  // proto:  int QHeaderView::resizeContentsPrecision();
extern int32_t C_ZNK11QHeaderView23resizeContentsPrecisionEv(void* qthis); // 4
  // proto:  bool QHeaderView::sectionsClickable();
extern bool C_ZNK11QHeaderView17sectionsClickableEv(void* qthis); // 4
  // proto:  int QHeaderView::logicalIndexAt(const QPoint & pos);
extern int32_t C_ZNK11QHeaderView14logicalIndexAtERK6QPoint(void* qthis, void* arg0); // 2
  // proto:  int QHeaderView::logicalIndexAt(int position);
extern int32_t C_ZNK11QHeaderView14logicalIndexAtEi(void* qthis, int32_t arg0); // 4
  // proto:  int QHeaderView::logicalIndexAt(int x, int y);
extern int32_t C_ZNK11QHeaderView14logicalIndexAtEii(void* qthis, int32_t arg0, int32_t arg1); // 2
  // proto:  int QHeaderView::offset();
extern int32_t C_ZNK11QHeaderView6offsetEv(void* qthis); // 4
  // proto:  void QHeaderView::setMinimumSectionSize(int size);
extern void C_ZN11QHeaderView21setMinimumSectionSizeEi(void* qthis, int32_t arg0); // 4
  // proto:  QSize QHeaderView::sizeHint();
extern void* C_ZNK11QHeaderView8sizeHintEv(void* qthis); // 4
  // proto:  void QHeaderView::setVisible(bool v);
extern void C_ZN11QHeaderView10setVisibleEb(void* qthis, bool arg0); // 4
  // proto:  void QHeaderView::setCascadingSectionResizes(bool enable);
extern void C_ZN11QHeaderView26setCascadingSectionResizesEb(void* qthis, bool arg0); // 4
  // proto:  bool QHeaderView::sectionsMovable();
extern bool C_ZNK11QHeaderView15sectionsMovableEv(void* qthis); // 4
  // proto:  int QHeaderView::length();
extern int32_t C_ZNK11QHeaderView6lengthEv(void* qthis); // 4
  // proto:  int QHeaderView::sortIndicatorSection();
extern int32_t C_ZNK11QHeaderView20sortIndicatorSectionEv(void* qthis); // 4
  // proto:  void QHeaderView::resizeSection(int logicalIndex, int size);
extern void C_ZN11QHeaderView13resizeSectionEii(void* qthis, int32_t arg0, int32_t arg1); // 4
  // proto:  void QHeaderView::setSectionsMovable(bool movable);
extern void C_ZN11QHeaderView18setSectionsMovableEb(void* qthis, bool arg0); // 4
  // proto:  int QHeaderView::maximumSectionSize();
extern int32_t C_ZNK11QHeaderView18maximumSectionSizeEv(void* qthis); // 4
  // proto:  void QHeaderView::moveSection(int from, int to);
extern void C_ZN11QHeaderView11moveSectionEii(void* qthis, int32_t arg0, int32_t arg1); // 4
  // proto:  bool QHeaderView::isSectionHidden(int logicalIndex);
extern bool C_ZNK11QHeaderView15isSectionHiddenEi(void* qthis, int32_t arg0); // 4
  // proto:  int QHeaderView::visualIndexAt(int position);
extern int32_t C_ZNK11QHeaderView13visualIndexAtEi(void* qthis, int32_t arg0); // 4
  // proto:  void QHeaderView::setSortIndicatorShown(bool show);
extern void C_ZN11QHeaderView21setSortIndicatorShownEb(void* qthis, bool arg0); // 4
  // proto:  Qt::Alignment QHeaderView::defaultAlignment();
extern void C_ZNK11QHeaderView16defaultAlignmentEv(void* qthis); // 4
  // proto:  int QHeaderView::hiddenSectionCount();
extern int32_t C_ZNK11QHeaderView18hiddenSectionCountEv(void* qthis); // 4
  // proto:  void QHeaderView::setStretchLastSection(bool stretch);
extern void C_ZN11QHeaderView21setStretchLastSectionEb(void* qthis, bool arg0); // 4
  // proto:  int QHeaderView::count();
extern int32_t C_ZNK11QHeaderView5countEv(void* qthis); // 4
  // proto:  int QHeaderView::minimumSectionSize();
extern int32_t C_ZNK11QHeaderView18minimumSectionSizeEv(void* qthis); // 4
  // proto:  QByteArray QHeaderView::saveState();
extern void* C_ZNK11QHeaderView9saveStateEv(void* qthis); // 4
  // proto:  Qt::Orientation QHeaderView::orientation();
extern void C_ZNK11QHeaderView11orientationEv(void* qthis); // 4
  // proto:  void QHeaderView::setSectionHidden(int logicalIndex, bool hide);
extern void C_ZN11QHeaderView16setSectionHiddenEib(void* qthis, int32_t arg0, bool arg1); // 4
  // proto:  void QHeaderView::setOffsetToLastSection();
extern void C_ZN11QHeaderView22setOffsetToLastSectionEv(void* qthis); // 4
  // proto:  bool QHeaderView::stretchLastSection();
extern bool C_ZNK11QHeaderView18stretchLastSectionEv(void* qthis); // 4
  // proto:  bool QHeaderView::isSortIndicatorShown();
extern bool C_ZNK11QHeaderView20isSortIndicatorShownEv(void* qthis); // 4
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
extern bool C_ZNK11QHeaderView13sectionsMovedEv(void* qthis); // 4
  // proto:  int QHeaderView::sectionViewportPosition(int logicalIndex);
extern int32_t C_ZNK11QHeaderView23sectionViewportPositionEi(void* qthis, int32_t arg0); // 4
  // proto:  int QHeaderView::stretchSectionCount();
extern int32_t C_ZNK11QHeaderView19stretchSectionCountEv(void* qthis); // 4
  // proto:  int QHeaderView::visualIndex(int logicalIndex);
extern int32_t C_ZNK11QHeaderView11visualIndexEi(void* qthis, int32_t arg0); // 4
  // proto:  bool QHeaderView::highlightSections();
extern bool C_ZNK11QHeaderView17highlightSectionsEv(void* qthis); // 4
  // proto:  int QHeaderView::defaultSectionSize();
extern int32_t C_ZNK11QHeaderView18defaultSectionSizeEv(void* qthis); // 4
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
func (this *QHeaderView) Cascadingsectionresizes(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK11QHeaderView23cascadingSectionResizesEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QHeaderView", "cascadingSectionResizes", args)
  }

  return
}

// setResizeContentsPrecision(int)
func (this *QHeaderView) Setresizecontentsprecision(args ...interface{}) () {
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

  return
}

// sectionsHidden()
func (this *QHeaderView) Sectionshidden(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK11QHeaderView14sectionsHiddenEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QHeaderView", "sectionsHidden", args)
  }

  return
}

// sectionSizeHint(int)
func (this *QHeaderView) Sectionsizehint(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK11QHeaderView15sectionSizeHintEi(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QHeaderView", "sectionSizeHint", args)
  }

  return
}

// sortIndicatorOrder()
func (this *QHeaderView) Sortindicatororder(args ...interface{}) () {
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

  return
}

// resetDefaultSectionSize()
func (this *QHeaderView) Resetdefaultsectionsize(args ...interface{}) () {
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

  return
}

// setOffset(int)
func (this *QHeaderView) Setoffset(args ...interface{}) () {
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

  return
}

// reset()
func (this *QHeaderView) Reset(args ...interface{}) () {
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

  return
}

// sectionPosition(int)
func (this *QHeaderView) Sectionposition(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK11QHeaderView15sectionPositionEi(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QHeaderView", "sectionPosition", args)
  }

  return
}

// logicalIndex(int)
func (this *QHeaderView) Logicalindex(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK11QHeaderView12logicalIndexEi(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QHeaderView", "logicalIndex", args)
  }

  return
}

// sectionSize(int)
func (this *QHeaderView) Sectionsize(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK11QHeaderView11sectionSizeEi(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QHeaderView", "sectionSize", args)
  }

  return
}

// metaObject()
func (this *QHeaderView) Metaobject(args ...interface{}) () {
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

  return
}

// hideSection(int)
func (this *QHeaderView) Hidesection(args ...interface{}) () {
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

  return
}

// setMaximumSectionSize(int)
func (this *QHeaderView) Setmaximumsectionsize(args ...interface{}) () {
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

  return
}

// sectionResizeMode(int)
func (this *QHeaderView) Sectionresizemode(args ...interface{}) () {
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

  return
}

// ~QHeaderView()
func (this *QHeaderView) Freeqheaderview(args ...interface{}) () {
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

  return
}

// restoreState(const class QByteArray &)
func (this *QHeaderView) Restorestate(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZN11QHeaderView12restoreStateERK10QByteArray(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QHeaderView", "restoreState", args)
  }

  return
}

// setOffsetToSectionPosition(int)
func (this *QHeaderView) Setoffsettosectionposition(args ...interface{}) () {
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

  return
}

// doItemsLayout()
func (this *QHeaderView) Doitemslayout(args ...interface{}) () {
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

  return
}

// setHighlightSections(_Bool)
func (this *QHeaderView) Sethighlightsections(args ...interface{}) () {
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

  return
}

// resizeContentsPrecision()
func (this *QHeaderView) Resizecontentsprecision(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK11QHeaderView23resizeContentsPrecisionEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QHeaderView", "resizeContentsPrecision", args)
  }

  return
}

// sectionsClickable()
func (this *QHeaderView) Sectionsclickable(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK11QHeaderView17sectionsClickableEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QHeaderView", "sectionsClickable", args)
  }

  return
}

// logicalIndexAt(const class QPoint &)
func (this *QHeaderView) Logicalindexat(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK11QHeaderView14logicalIndexAtERK6QPoint(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  case 1:
    // invoke: _ZNK11QHeaderView14logicalIndexAtEi
    // invoke: int logicalIndexAt(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK11QHeaderView14logicalIndexAtEi(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  case 2:
    // invoke: _ZNK11QHeaderView14logicalIndexAtEii
    // invoke: int logicalIndexAt(int, int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var ret0 = C.C_ZNK11QHeaderView14logicalIndexAtEii(this.qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QHeaderView", "logicalIndexAt", args)
  }

  return
}

// offset()
func (this *QHeaderView) Offset(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK11QHeaderView6offsetEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QHeaderView", "offset", args)
  }

  return
}

// setMinimumSectionSize(int)
func (this *QHeaderView) Setminimumsectionsize(args ...interface{}) () {
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

  return
}

// sizeHint()
func (this *QHeaderView) Sizehint(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK11QHeaderView8sizeHintEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QSize{}) // "QSize"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QHeaderView", "sizeHint", args)
  }

  return
}

// setVisible(_Bool)
func (this *QHeaderView) Setvisible(args ...interface{}) () {
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

  return
}

// setCascadingSectionResizes(_Bool)
func (this *QHeaderView) Setcascadingsectionresizes(args ...interface{}) () {
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

  return
}

// sectionsMovable()
func (this *QHeaderView) Sectionsmovable(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK11QHeaderView15sectionsMovableEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QHeaderView", "sectionsMovable", args)
  }

  return
}

// length()
func (this *QHeaderView) Length(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK11QHeaderView6lengthEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QHeaderView", "length", args)
  }

  return
}

// sortIndicatorSection()
func (this *QHeaderView) Sortindicatorsection(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK11QHeaderView20sortIndicatorSectionEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QHeaderView", "sortIndicatorSection", args)
  }

  return
}

// resizeSection(int, int)
func (this *QHeaderView) Resizesection(args ...interface{}) () {
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

  return
}

// setSectionsMovable(_Bool)
func (this *QHeaderView) Setsectionsmovable(args ...interface{}) () {
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

  return
}

// maximumSectionSize()
func (this *QHeaderView) Maximumsectionsize(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK11QHeaderView18maximumSectionSizeEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QHeaderView", "maximumSectionSize", args)
  }

  return
}

// moveSection(int, int)
func (this *QHeaderView) Movesection(args ...interface{}) () {
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

  return
}

// isSectionHidden(int)
func (this *QHeaderView) Issectionhidden(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK11QHeaderView15isSectionHiddenEi(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QHeaderView", "isSectionHidden", args)
  }

  return
}

// visualIndexAt(int)
func (this *QHeaderView) Visualindexat(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK11QHeaderView13visualIndexAtEi(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QHeaderView", "visualIndexAt", args)
  }

  return
}

// setSortIndicatorShown(_Bool)
func (this *QHeaderView) Setsortindicatorshown(args ...interface{}) () {
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

  return
}

// defaultAlignment()
func (this *QHeaderView) Defaultalignment(args ...interface{}) () {
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

  return
}

// hiddenSectionCount()
func (this *QHeaderView) Hiddensectioncount(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK11QHeaderView18hiddenSectionCountEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QHeaderView", "hiddenSectionCount", args)
  }

  return
}

// setStretchLastSection(_Bool)
func (this *QHeaderView) Setstretchlastsection(args ...interface{}) () {
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

  return
}

// count()
func (this *QHeaderView) Count(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK11QHeaderView5countEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QHeaderView", "count", args)
  }

  return
}

// minimumSectionSize()
func (this *QHeaderView) Minimumsectionsize(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK11QHeaderView18minimumSectionSizeEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QHeaderView", "minimumSectionSize", args)
  }

  return
}

// saveState()
func (this *QHeaderView) Savestate(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK11QHeaderView9saveStateEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QByteArray{}) // "QByteArray"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QHeaderView", "saveState", args)
  }

  return
}

// orientation()
func (this *QHeaderView) Orientation(args ...interface{}) () {
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

  return
}

// setSectionHidden(int, _Bool)
func (this *QHeaderView) Setsectionhidden(args ...interface{}) () {
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

  return
}

// setOffsetToLastSection()
func (this *QHeaderView) Setoffsettolastsection(args ...interface{}) () {
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

  return
}

// stretchLastSection()
func (this *QHeaderView) Stretchlastsection(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK11QHeaderView18stretchLastSectionEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QHeaderView", "stretchLastSection", args)
  }

  return
}

// isSortIndicatorShown()
func (this *QHeaderView) Issortindicatorshown(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK11QHeaderView20isSortIndicatorShownEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QHeaderView", "isSortIndicatorShown", args)
  }

  return
}

// showSection(int)
func (this *QHeaderView) Showsection(args ...interface{}) () {
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

  return
}

// swapSections(int, int)
func (this *QHeaderView) Swapsections(args ...interface{}) () {
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

  return
}

// setSectionsClickable(_Bool)
func (this *QHeaderView) Setsectionsclickable(args ...interface{}) () {
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

  return
}

// setModel(class QAbstractItemModel *)
func (this *QHeaderView) Setmodel(args ...interface{}) () {
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

  return
}

// setDefaultSectionSize(int)
func (this *QHeaderView) Setdefaultsectionsize(args ...interface{}) () {
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

  return
}

// sectionsMoved()
func (this *QHeaderView) Sectionsmoved(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK11QHeaderView13sectionsMovedEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QHeaderView", "sectionsMoved", args)
  }

  return
}

// sectionViewportPosition(int)
func (this *QHeaderView) Sectionviewportposition(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK11QHeaderView23sectionViewportPositionEi(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QHeaderView", "sectionViewportPosition", args)
  }

  return
}

// stretchSectionCount()
func (this *QHeaderView) Stretchsectioncount(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK11QHeaderView19stretchSectionCountEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QHeaderView", "stretchSectionCount", args)
  }

  return
}

// visualIndex(int)
func (this *QHeaderView) Visualindex(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK11QHeaderView11visualIndexEi(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QHeaderView", "visualIndex", args)
  }

  return
}

// highlightSections()
func (this *QHeaderView) Highlightsections(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK11QHeaderView17highlightSectionsEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QHeaderView", "highlightSections", args)
  }

  return
}

// defaultSectionSize()
func (this *QHeaderView) Defaultsectionsize(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK11QHeaderView18defaultSectionSizeEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QHeaderView", "defaultSectionSize", args)
  }

  return
}

// <= body block end

