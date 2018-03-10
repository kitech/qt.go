package qtcore

// /usr/include/qt/QtCore/qabstractnativeeventfilter.h
// #include <qabstractnativeeventfilter.h>
// #include <QtCore>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 7
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "log"
import "github.com/kitech/qt.go/qtrt"

//  ext block end

//  body block begin

/*

 */
type QAbstractNativeEventFilter struct {
	*qtrt.CObject
}
type QAbstractNativeEventFilter_ITF interface {
	QAbstractNativeEventFilter_PTR() *QAbstractNativeEventFilter
}

func (ptr *QAbstractNativeEventFilter) QAbstractNativeEventFilter_PTR() *QAbstractNativeEventFilter {
	return ptr
}

func (this *QAbstractNativeEventFilter) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QAbstractNativeEventFilter) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQAbstractNativeEventFilterFromPointer(cthis unsafe.Pointer) *QAbstractNativeEventFilter {
	return &QAbstractNativeEventFilter{&qtrt.CObject{cthis}}
}
func (*QAbstractNativeEventFilter) NewFromPointer(cthis unsafe.Pointer) *QAbstractNativeEventFilter {
	return NewQAbstractNativeEventFilterFromPointer(cthis)
}

// /usr/include/qt/QtCore/qabstractnativeeventfilter.h:52
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QAbstractNativeEventFilter()

/*
Creates a native event filter.

By default this doesn't do anything. Remember to install it on the application object.
*/
func NewQAbstractNativeEventFilter() *QAbstractNativeEventFilter {
	rv, err := qtrt.InvokeQtFunc6("_ZN26QAbstractNativeEventFilterC2Ev", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	gothis := NewQAbstractNativeEventFilterFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQAbstractNativeEventFilter)
	return gothis
}

// /usr/include/qt/QtCore/qabstractnativeeventfilter.h:53
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QAbstractNativeEventFilter()

/*

 */
func DeleteQAbstractNativeEventFilter(this *QAbstractNativeEventFilter) {
	rv, err := qtrt.InvokeQtFunc6("_ZN26QAbstractNativeEventFilterD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 16)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtCore/qabstractnativeeventfilter.h:55
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [1] bool nativeEventFilter(const QByteArray &, void *, long *)

/*
This method is called for every native event.

Note: The filter function here receives native messages, for example, MSG or XCB event structs.

It is called by the QPA platform plugin. On Windows, it is called by the event dispatcher.

The type of event eventType is specific to the platform plugin chosen at run-time, and can be used to cast message to the right type.

On X11, eventType is set to "xcb_generic_event_t", and the message can be casted to a xcb_generic_event_t pointer.

On Windows, eventType is set to "windows_generic_MSG" for messages sent to toplevel windows, and "windows_dispatcher_MSG" for system-wide messages such as messages from a registered hot key. In both cases, the message can be casted to a MSG pointer. The result pointer is only used on Windows, and corresponds to the LRESULT pointer.

On macOS, eventType is set to "mac_generic_NSEvent", and the message can be casted to an NSEvent pointer.

In your reimplementation of this function, if you want to filter the message out, i.e. stop it being handled further, return true; otherwise return false.

Linux example


  class MyXcbEventFilter : public QAbstractNativeEventFilter
  {
  public:
      bool nativeEventFilter(const QByteArray &eventType, void *message, long *) override
      {
          if (eventType == "xcb_generic_event_t") {
              xcb_generic_event_t* ev = static_cast<xcb_generic_event_t *>(message);
              // ...
          }
          return false;
      }
  };



macOS example

mycocoaeventfilter.h:


  #include <QAbstractNativeEventFilter>

  class MyCocoaEventFilter : public QAbstractNativeEventFilter
  {
  public:
      bool nativeEventFilter(const QByteArray &eventType, void *message, long *) override;
  };



mycocoaeventfilter.mm:


  #include "mycocoaeventfilter.h"

  #import <AppKit/AppKit.h>

  bool CocoaNativeEventFilter::nativeEventFilter(const QByteArray &eventType, void *message, long *)
  {
      if (eventType == "mac_generic_NSEvent") {
          NSEvent *event = static_cast<NSEvent *>(message);
          if ([event type] == NSKeyDown) {
              // Handle key event
              qDebug() << QString::fromNSString([event characters]);
          }
      }
      return false;
  }



myapp.pro:


  HEADERS += mycocoaeventfilter.h
  OBJECTIVE_SOURCES += mycocoaeventfilter.mm
  LIBS += -framework AppKit
*/
func (this *QAbstractNativeEventFilter) NativeEventFilter(eventType QByteArray_ITF, message unsafe.Pointer /*666*/, result unsafe.Pointer /*666*/) bool {
	var convArg0 unsafe.Pointer
	if eventType != nil && eventType.QByteArray_PTR() != nil {
		convArg0 = eventType.QByteArray_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN26QAbstractNativeEventFilter17nativeEventFilterERK10QByteArrayPvPl", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, message, result)
	qtrt.ErrPrint(err, rv)
	return rv != 0
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
}

//  keep block end
