package qt5
// auto generated, do not modify.
// created: Sat Jan  2 12:23:25 2016
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
// #[link(name = "Qt5Core")]
// #[link(name = "Qt5Gui")]
// #[link(name = "Qt5Widgets")]
// #[link(name = "QtInline")]

/*
#include <stdlib.h>
#include <stdbool.h>
#include <stdint.h>
#include <wchar.h>
#include <uchar.h>
  // proto:  void QPdfWriter::~QPdfWriter();
extern void _ZN10QPdfWriterD0Ev(void* qthis);
  // proto:  void QPdfWriter::setCreator(const QString & creator);
extern void _ZN10QPdfWriter10setCreatorERK7QString(void* qthis, void* arg0);
  // proto:  void QPdfWriter::setPageSizeMM(const QSizeF & size);
extern void _ZN10QPdfWriter13setPageSizeMMERK6QSizeF(void* qthis, void* arg0);
  // proto:  void QPdfWriter::setResolution(int resolution);
extern void _ZN10QPdfWriter13setResolutionEi(void* qthis, int arg0);
  // proto:  void QPdfWriter::QPdfWriter(const QPdfWriter & );
extern void* dector_ZN10QPdfWriterC1ERKS_(void* arg0);
extern void _ZN10QPdfWriterC1ERKS_(void* qthis, void* arg0);
  // proto:  bool QPdfWriter::newPage();
extern void _ZN10QPdfWriter7newPageEv(void* qthis);
  // proto:  QString QPdfWriter::title();
extern void _ZNK10QPdfWriter5titleEv(void* qthis);
  // proto:  QString QPdfWriter::creator();
extern void _ZNK10QPdfWriter7creatorEv(void* qthis);
  // proto:  int QPdfWriter::resolution();
extern void _ZNK10QPdfWriter10resolutionEv(void* qthis);
  // proto:  const QMetaObject * QPdfWriter::metaObject();
extern void _ZNK10QPdfWriter10metaObjectEv(void* qthis);
  // proto:  void QPdfWriter::QPdfWriter(const QString & filename);
extern void* dector_ZN10QPdfWriterC1ERK7QString(void* arg0);
extern void _ZN10QPdfWriterC1ERK7QString(void* qthis, void* arg0);
  // proto:  void QPdfWriter::QPdfWriter(QIODevice * device);
extern void* dector_ZN10QPdfWriterC1EP9QIODevice(void* arg0);
extern void _ZN10QPdfWriterC1EP9QIODevice(void* qthis, void* arg0);
  // proto:  void QPdfWriter::setTitle(const QString & title);
extern void _ZN10QPdfWriter8setTitleERK7QString(void* qthis, void* arg0);
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
  qclsinst uint64 /* *mut c_void*/;
}

  // proto:  void QPdfWriter::~QPdfWriter();
func (this *QPdfWriter) FreeQPdfWriter(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QPdfWriter", "~QPdfWriter", args)
  }

}

  // proto:  void QPdfWriter::setCreator(const QString & creator);
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
  default:
    qtrt.ErrorResolve("QPdfWriter", "setCreator", args)
  }

}

  // proto:  void QPdfWriter::setPageSizeMM(const QSizeF & size);
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
  default:
    qtrt.ErrorResolve("QPdfWriter", "setPageSizeMM", args)
  }

}

  // proto:  void QPdfWriter::setResolution(int resolution);
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
  default:
    qtrt.ErrorResolve("QPdfWriter", "setResolution", args)
  }

}

  // proto:  void QPdfWriter::QPdfWriter(const QPdfWriter & );
func NewQPdfWriter(args ...interface{}) QPdfWriter {
  return QPdfWriter{}
}

  // proto:  bool QPdfWriter::newPage();
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
  default:
    qtrt.ErrorResolve("QPdfWriter", "newPage", args)
  }

}

  // proto:  QString QPdfWriter::title();
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
  default:
    qtrt.ErrorResolve("QPdfWriter", "title", args)
  }

}

  // proto:  QString QPdfWriter::creator();
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
  default:
    qtrt.ErrorResolve("QPdfWriter", "creator", args)
  }

}

  // proto:  int QPdfWriter::resolution();
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
  default:
    qtrt.ErrorResolve("QPdfWriter", "resolution", args)
  }

}

  // proto:  const QMetaObject * QPdfWriter::metaObject();
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
  default:
    qtrt.ErrorResolve("QPdfWriter", "metaObject", args)
  }

}

  // proto:  void QPdfWriter::setTitle(const QString & title);
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
  default:
    qtrt.ErrorResolve("QPdfWriter", "setTitle", args)
  }

}

// <= body block end

