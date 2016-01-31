package qt5
// auto generated, do not modify.
// created: Sun Jan 31 12:22:27 2016
// src-file: /QtGui/qpdfwriter.h
// dst-file: /src/gui/qpdfwriter.go
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
  // proto:  bool QPdfWriter::newPage();
extern void C_ZN10QPdfWriter7newPageEv(void* qthis); // 4
  // proto:  QString QPdfWriter::creator();
extern void C_ZNK10QPdfWriter7creatorEv(void* qthis); // 4
  // proto:  void QPdfWriter::~QPdfWriter();
extern void C_ZN10QPdfWriterD2Ev(void* qthis); // 4
  // proto:  void QPdfWriter::setTitle(const QString & title);
extern void C_ZN10QPdfWriter8setTitleERK7QString(void* qthis, void* arg0); // 4
  // proto:  QString QPdfWriter::title();
extern void C_ZNK10QPdfWriter5titleEv(void* qthis); // 4
  // proto:  void QPdfWriter::QPdfWriter(const QString & filename);
extern void C_ZN10QPdfWriterC2ERK7QString(void* qthis, void* arg0); // 3
  // proto:  void QPdfWriter::QPdfWriter(QIODevice * device);
extern void C_ZN10QPdfWriterC2EP9QIODevice(void* qthis, void* arg0); // 3
  // proto:  void QPdfWriter::setPageSizeMM(const QSizeF & size);
extern void C_ZN10QPdfWriter13setPageSizeMMERK6QSizeF(void* qthis, void* arg0); // 4
  // proto:  void QPdfWriter::setResolution(int resolution);
extern void C_ZN10QPdfWriter13setResolutionEi(void* qthis, int32_t arg0); // 4
  // proto:  const QMetaObject * QPdfWriter::metaObject();
extern void C_ZNK10QPdfWriter10metaObjectEv(void* qthis); // 4
  // proto:  void QPdfWriter::setCreator(const QString & creator);
extern void C_ZN10QPdfWriter10setCreatorERK7QString(void* qthis, void* arg0); // 4
  // proto:  int QPdfWriter::resolution();
extern void C_ZNK10QPdfWriter10resolutionEv(void* qthis); // 4
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

// class sizeof(QPdfWriter)=1
type QPdfWriter struct {
  /*qbase*/ QObject;
  qclsinst unsafe.Pointer /* *C.void */;
}

// newPage()
func (this *QPdfWriter) newPage(args ...interface{}) () {
  // newPage()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QPdfWriter7newPageEv
    // invoke: bool newPage()
    C.C_ZN10QPdfWriter7newPageEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPdfWriter", "newPage", args)
  }

}

// creator()
func (this *QPdfWriter) creator(args ...interface{}) () {
  // creator()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QPdfWriter7creatorEv
    // invoke: QString creator()
    C.C_ZNK10QPdfWriter7creatorEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPdfWriter", "creator", args)
  }

}

// ~QPdfWriter()
func (this *QPdfWriter) FreeQPdfWriter(args ...interface{}) () {
  // ~QPdfWriter()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QPdfWriterD0Ev
    // invoke: void ~QPdfWriter()
    C.C_ZN10QPdfWriterD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPdfWriter", "~QPdfWriter", args)
  }

}

// setTitle(const class QString &)
func (this *QPdfWriter) setTitle(args ...interface{}) () {
  // setTitle(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QPdfWriter8setTitleERK7QString
    // invoke: void setTitle(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN10QPdfWriter8setTitleERK7QString(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPdfWriter", "setTitle", args)
  }

}

// title()
func (this *QPdfWriter) title(args ...interface{}) () {
  // title()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QPdfWriter5titleEv
    // invoke: QString title()
    C.C_ZNK10QPdfWriter5titleEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPdfWriter", "title", args)
  }

}

// QPdfWriter(const class QString &)
func NewQPdfWriter(args ...interface{}) QPdfWriter {
  // QPdfWriter(const class QString &)
  // QPdfWriter(class QIODevice *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QIODevice{}) // "QIODevice *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QPdfWriterC1ERK7QString
    // invoke: void QPdfWriter(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C.C_ZN10QPdfWriterC2ERK7QString(qthis, arg0)
  case 1:
    // invoke: _ZN10QPdfWriterC1EP9QIODevice
    // invoke: void QPdfWriter(class QIODevice *)
    var arg0 = args[0].(QIODevice).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C.C_ZN10QPdfWriterC2EP9QIODevice(qthis, arg0)
  default:
    qtrt.ErrorResolve("QPdfWriter", "QPdfWriter", args)
  }

  return QPdfWriter{}
}

// setPageSizeMM(const class QSizeF &)
func (this *QPdfWriter) setPageSizeMM(args ...interface{}) () {
  // setPageSizeMM(const class QSizeF &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QSizeF{}) // "const QSizeF &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QPdfWriter13setPageSizeMMERK6QSizeF
    // invoke: void setPageSizeMM(const class QSizeF &)
    var arg0 = args[0].(QSizeF).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN10QPdfWriter13setPageSizeMMERK6QSizeF(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPdfWriter", "setPageSizeMM", args)
  }

}

// setResolution(int)
func (this *QPdfWriter) setResolution(args ...interface{}) () {
  // setResolution(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QPdfWriter13setResolutionEi
    // invoke: void setResolution(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN10QPdfWriter13setResolutionEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPdfWriter", "setResolution", args)
  }

}

// metaObject()
func (this *QPdfWriter) metaObject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QPdfWriter10metaObjectEv
    // invoke: const QMetaObject * metaObject()
    C.C_ZNK10QPdfWriter10metaObjectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPdfWriter", "metaObject", args)
  }

}

// setCreator(const class QString &)
func (this *QPdfWriter) setCreator(args ...interface{}) () {
  // setCreator(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QPdfWriter10setCreatorERK7QString
    // invoke: void setCreator(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN10QPdfWriter10setCreatorERK7QString(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPdfWriter", "setCreator", args)
  }

}

// resolution()
func (this *QPdfWriter) resolution(args ...interface{}) () {
  // resolution()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QPdfWriter10resolutionEv
    // invoke: int resolution()
    C.C_ZNK10QPdfWriter10resolutionEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPdfWriter", "resolution", args)
  }

}

// <= body block end

