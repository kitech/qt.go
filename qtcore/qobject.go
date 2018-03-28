package qtcore

// /usr/include/qt/QtCore/qobject.h
// #include <qobject.h>
// #include <QtCore>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 2
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "log"
import "github.com/kitech/qt.go/qtrt"

//  ext block end

//  body block begin

// QObject * sender()
func (this *QObject) InheritSender(f func() unsafe.Pointer /*666*/) {
	qtrt.SetAllInheritCallback(this, "sender", f)
}

// int senderSignalIndex()
func (this *QObject) InheritSenderSignalIndex(f func() int) {
	qtrt.SetAllInheritCallback(this, "senderSignalIndex", f)
}

// int receivers(const char *)
func (this *QObject) InheritReceivers(f func(signal string) int) {
	qtrt.SetAllInheritCallback(this, "receivers", f)
}

// bool isSignalConnected(const QMetaMethod &)
func (this *QObject) InheritIsSignalConnected(f func(signal *QMetaMethod) bool) {
	qtrt.SetAllInheritCallback(this, "isSignalConnected", f)
}

// void timerEvent(QTimerEvent *)
func (this *QObject) InheritTimerEvent(f func(event *QTimerEvent /*777 QTimerEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "timerEvent", f)
}

// void childEvent(QChildEvent *)
func (this *QObject) InheritChildEvent(f func(event *QChildEvent /*777 QChildEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "childEvent", f)
}

// void customEvent(QEvent *)
func (this *QObject) InheritCustomEvent(f func(event *QEvent /*777 QEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "customEvent", f)
}

// void connectNotify(const QMetaMethod &)
func (this *QObject) InheritConnectNotify(f func(signal *QMetaMethod) /*void*/) {
	qtrt.SetAllInheritCallback(this, "connectNotify", f)
}

// void disconnectNotify(const QMetaMethod &)
func (this *QObject) InheritDisconnectNotify(f func(signal *QMetaMethod) /*void*/) {
	qtrt.SetAllInheritCallback(this, "disconnectNotify", f)
}

/*

 */
type QObject struct {
	*qtrt.CObject
}
type QObject_ITF interface {
	QObject_PTR() *QObject
}

func (ptr *QObject) QObject_PTR() *QObject { return ptr }

func (this *QObject) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QObject) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQObjectFromPointer(cthis unsafe.Pointer) *QObject {
	return &QObject{&qtrt.CObject{cthis}}
}
func (*QObject) NewFromPointer(cthis unsafe.Pointer) *QObject {
	return NewQObjectFromPointer(cthis)
}

// /usr/include/qt/QtCore/qobject.h:119
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject() const

/*
Returns a pointer to the meta-object of this object.

A meta-object contains information about a class that inherits QObject, e.g. class name, superclass name, properties, signals and slots. Every QObject subclass that contains the Q_OBJECT macro will have a meta-object.

The meta-object information is required by the signal/slot connection mechanism and the property system. The inherits() function also makes use of the meta-object.

If you have no pointer to an actual object instance but still want to access the meta-object of a class, you can use staticMetaObject.

Example:


  QObject *obj = new QPushButton;
  obj->metaObject()->className();             // returns "QPushButton"

  QPushButton::staticMetaObject.className();  // returns "QPushButton"



See also staticMetaObject.
*/
func (this *QObject) MetaObject() *QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QObject10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtCore/qobject.h:124
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QObject(QObject *)

/*
Constructs an object with parent object parent.

The parent of an object may be viewed as the object's owner. For instance, a dialog box is the parent of the OK and Cancel buttons it contains.

The destructor of a parent object destroys all child objects.

Setting parent to 0 constructs an object with no parent. If the object is a widget, it will become a top-level window.

See also parent(), findChild(), and findChildren().
*/
func NewQObject(parent QObject_ITF /*777 QObject **/) *QObject {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QObject_PTR() != nil {
		convArg0 = parent.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN7QObjectC2EPS_", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQObjectFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQObject)
	return gothis
}

// /usr/include/qt/QtCore/qobject.h:124
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QObject(QObject *)

/*
Constructs an object with parent object parent.

The parent of an object may be viewed as the object's owner. For instance, a dialog box is the parent of the OK and Cancel buttons it contains.

The destructor of a parent object destroys all child objects.

Setting parent to 0 constructs an object with no parent. If the object is a widget, it will become a top-level window.

See also parent(), findChild(), and findChildren().
*/
func NewQObject__() *QObject {
	// arg: 0, QObject *=Pointer, QObject=Record,
	var convArg0 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN7QObjectC2EPS_", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQObjectFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQObject)
	return gothis
}

// /usr/include/qt/QtCore/qobject.h:125
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QObject()

/*

 */
func DeleteQObject(this *QObject) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QObjectD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 16)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtCore/qobject.h:127
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool event(QEvent *)

/*
This virtual function receives events to an object and should return true if the event e was recognized and processed.

The event() function can be reimplemented to customize the behavior of an object.

Make sure you call the parent event class implementation for all the events you did not handle.

Example:


  class MyClass : public QWidget
  {
      Q_OBJECT

  public:
      MyClass(QWidget *parent = 0);
      ~MyClass();

      bool event(QEvent* ev)
      {
          if (ev->type() == QEvent::PolishRequest) {
              // overwrite handling of PolishRequest if any
              doThings();
              return true;
          } else  if (ev->type() == QEvent::Show) {
              // complement handling of Show if any
              doThings2();
              QWidget::event(ev);
              return true;
          }
          // Make sure the rest of events are handled
          return QWidget::event(ev);
      }
  };



See also installEventFilter(), timerEvent(), QCoreApplication::sendEvent(), and QCoreApplication::postEvent().
*/
func (this *QObject) Event(event QEvent_ITF /*777 QEvent **/) bool {
	var convArg0 unsafe.Pointer
	if event != nil && event.QEvent_PTR() != nil {
		convArg0 = event.QEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN7QObject5eventEP6QEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qobject.h:128
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool eventFilter(QObject *, QEvent *)

/*
Filters events if this object has been installed as an event filter for the watched object.

In your reimplementation of this function, if you want to filter the event out, i.e. stop it being handled further, return true; otherwise return false.

Example:


  class MainWindow : public QMainWindow
  {
  public:
      MainWindow();

  protected:
      bool eventFilter(QObject *obj, QEvent *ev);

  private:
      QTextEdit *textEdit;
  };

  MainWindow::MainWindow()
  {
      textEdit = new QTextEdit;
      setCentralWidget(textEdit);

      textEdit->installEventFilter(this);
  }

  bool MainWindow::eventFilter(QObject *obj, QEvent *event)
  {
      if (obj == textEdit) {
          if (event->type() == QEvent::KeyPress) {
              QKeyEvent *keyEvent = static_cast<QKeyEvent*>(event);
              qDebug() << "Ate key press" << keyEvent->key();
              return true;
          } else {
              return false;
          }
      } else {
          // pass the event on to the parent class
          return QMainWindow::eventFilter(obj, event);
      }
  }



Notice in the example above that unhandled events are passed to the base class's eventFilter() function, since the base class might have reimplemented eventFilter() for its own internal purposes.

Warning: If you delete the receiver object in this function, be sure to return true. Otherwise, Qt will forward the event to the deleted object and the program might crash.

See also installEventFilter().
*/
func (this *QObject) EventFilter(watched QObject_ITF /*777 QObject **/, event QEvent_ITF /*777 QEvent **/) bool {
	var convArg0 unsafe.Pointer
	if watched != nil && watched.QObject_PTR() != nil {
		convArg0 = watched.QObject_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if event != nil && event.QEvent_PTR() != nil {
		convArg1 = event.QEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN7QObject11eventFilterEPS_P6QEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qobject.h:145
// index:0
// Public Visibility=Default Availability=Available
// [8] QString objectName() const

/*

 */
func (this *QObject) ObjectName() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QObject10objectNameEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qobject.h:146
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setObjectName(const QString &)

/*

 */
func (this *QObject) SetObjectName(name string) {
	var tmpArg0 = NewQString_5(name)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN7QObject13setObjectNameERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qobject.h:148
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isWidgetType() const

/*
Returns true if the object is a widget; otherwise returns false.

Calling this function is equivalent to calling inherits("QWidget"), except that it is much faster.
*/
func (this *QObject) IsWidgetType() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QObject12isWidgetTypeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qobject.h:149
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isWindowType() const

/*
Returns true if the object is a window; otherwise returns false.

Calling this function is equivalent to calling inherits("QWindow"), except that it is much faster.
*/
func (this *QObject) IsWindowType() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QObject12isWindowTypeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qobject.h:151
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool signalsBlocked() const

/*
Returns true if signals are blocked; otherwise returns false.

Signals are not blocked by default.

See also blockSignals() and QSignalBlocker.
*/
func (this *QObject) SignalsBlocked() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QObject14signalsBlockedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qobject.h:152
// index:0
// Public Visibility=Default Availability=Available
// [1] bool blockSignals(bool)

/*
If block is true, signals emitted by this object are blocked (i.e., emitting a signal will not invoke anything connected to it). If block is false, no such blocking will occur.

The return value is the previous value of signalsBlocked().

Note that the destroyed() signal will be emitted even if the signals for this object have been blocked.

Signals emitted while being blocked are not buffered.

See also signalsBlocked() and QSignalBlocker.
*/
func (this *QObject) BlockSignals(b bool) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QObject12blockSignalsEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), b)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qobject.h:154
// index:0
// Public Visibility=Default Availability=Available
// [8] QThread * thread() const

/*
Returns the thread in which the object lives.

See also moveToThread().
*/
func (this *QObject) Thread() *QThread /*777 QThread **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QObject6threadEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQThreadFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtCore/qobject.h:155
// index:0
// Public Visibility=Default Availability=Available
// [-2] void moveToThread(QThread *)

/*
Changes the thread affinity for this object and its children. The object cannot be moved if it has a parent. Event processing will continue in the targetThread.

To move an object to the main thread, use QApplication::instance() to retrieve a pointer to the current application, and then use QApplication::thread() to retrieve the thread in which the application lives. For example:


  myObject->moveToThread(QApplication::instance()->thread());



If targetThread is zero, all event processing for this object and its children stops.

Note that all active timers for the object will be reset. The timers are first stopped in the current thread and restarted (with the same interval) in the targetThread. As a result, constantly moving an object between threads can postpone timer events indefinitely.

A QEvent::ThreadChange event is sent to this object just before the thread affinity is changed. You can handle this event to perform any special processing. Note that any new events that are posted to this object will be handled in the targetThread.

Warning: This function is not thread-safe; the current thread must be same as the current thread affinity. In other words, this function can only "push" an object from the current thread to another thread, it cannot "pull" an object from any arbitrary thread to the current thread.

See also thread().
*/
func (this *QObject) MoveToThread(thread QThread_ITF /*777 QThread **/) {
	var convArg0 unsafe.Pointer
	if thread != nil && thread.QThread_PTR() != nil {
		convArg0 = thread.QThread_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN7QObject12moveToThreadEP7QThread", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qobject.h:157
// index:0
// Public Visibility=Default Availability=Available
// [4] int startTimer(int, Qt::TimerType)

/*
Starts a timer and returns a timer identifier, or returns zero if it could not start a timer.

A timer event will occur every interval milliseconds until killTimer() is called. If interval is 0, then the timer event occurs once every time there are no more window system events to process.

The virtual timerEvent() function is called with the QTimerEvent event parameter class when a timer event occurs. Reimplement this function to get timer events.

If multiple timers are running, the QTimerEvent::timerId() can be used to find out which timer was activated.

Example:


  class MyObject : public QObject
  {
      Q_OBJECT

  public:
      MyObject(QObject *parent = 0);

  protected:
      void timerEvent(QTimerEvent *event);
  };

  MyObject::MyObject(QObject *parent)
      : QObject(parent)
  {
      startTimer(50);     // 50-millisecond timer
      startTimer(1000);   // 1-second timer
      startTimer(60000);  // 1-minute timer

      using namespace std::chrono;
      startTimer(milliseconds(50));
      startTimer(seconds(1));
      startTimer(minutes(1));

      // since C++14 we can use std::chrono::duration literals, e.g.:
      startTimer(100ms);
      startTimer(5s);
      startTimer(2min);
      startTimer(1h);
  }

  void MyObject::timerEvent(QTimerEvent *event)
  {
      qDebug() << "Timer ID:" << event->timerId();
  }



Note that QTimer's accuracy depends on the underlying operating system and hardware. The timerType argument allows you to customize the accuracy of the timer. See Qt::TimerType for information on the different timer types. Most platforms support an accuracy of 20 milliseconds; some provide more. If Qt is unable to deliver the requested number of timer events, it will silently discard some.

The QTimer class provides a high-level programming interface with single-shot timers and timer signals instead of events. There is also a QBasicTimer class that is more lightweight than QTimer and less clumsy than using timer IDs directly.

See also timerEvent(), killTimer(), and QTimer::singleShot().
*/
func (this *QObject) StartTimer(interval int, timerType int) int {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QObject10startTimerEiN2Qt9TimerTypeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), interval, timerType)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qobject.h:157
// index:0
// Public Visibility=Default Availability=Available
// [4] int startTimer(int, Qt::TimerType)

/*
Starts a timer and returns a timer identifier, or returns zero if it could not start a timer.

A timer event will occur every interval milliseconds until killTimer() is called. If interval is 0, then the timer event occurs once every time there are no more window system events to process.

The virtual timerEvent() function is called with the QTimerEvent event parameter class when a timer event occurs. Reimplement this function to get timer events.

If multiple timers are running, the QTimerEvent::timerId() can be used to find out which timer was activated.

Example:


  class MyObject : public QObject
  {
      Q_OBJECT

  public:
      MyObject(QObject *parent = 0);

  protected:
      void timerEvent(QTimerEvent *event);
  };

  MyObject::MyObject(QObject *parent)
      : QObject(parent)
  {
      startTimer(50);     // 50-millisecond timer
      startTimer(1000);   // 1-second timer
      startTimer(60000);  // 1-minute timer

      using namespace std::chrono;
      startTimer(milliseconds(50));
      startTimer(seconds(1));
      startTimer(minutes(1));

      // since C++14 we can use std::chrono::duration literals, e.g.:
      startTimer(100ms);
      startTimer(5s);
      startTimer(2min);
      startTimer(1h);
  }

  void MyObject::timerEvent(QTimerEvent *event)
  {
      qDebug() << "Timer ID:" << event->timerId();
  }



Note that QTimer's accuracy depends on the underlying operating system and hardware. The timerType argument allows you to customize the accuracy of the timer. See Qt::TimerType for information on the different timer types. Most platforms support an accuracy of 20 milliseconds; some provide more. If Qt is unable to deliver the requested number of timer events, it will silently discard some.

The QTimer class provides a high-level programming interface with single-shot timers and timer signals instead of events. There is also a QBasicTimer class that is more lightweight than QTimer and less clumsy than using timer IDs directly.

See also timerEvent(), killTimer(), and QTimer::singleShot().
*/
func (this *QObject) StartTimer__(interval int) int {
	// arg: 1, Qt::TimerType=Elaborated, Qt::TimerType=Enum,
	timerType := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN7QObject10startTimerEiN2Qt9TimerTypeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), interval, timerType)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qobject.h:165
// index:0
// Public Visibility=Default Availability=Available
// [-2] void killTimer(int)

/*
Kills the timer with timer identifier, id.

The timer identifier is returned by startTimer() when a timer event is started.

See also timerEvent() and startTimer().
*/
func (this *QObject) KillTimer(id int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QObject9killTimerEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), id)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qobject.h:208
// index:0
// Public inline Visibility=Default Availability=Available
// [8] const QObjectList & children() const

/*
Returns a list of child objects. The QObjectList class is defined in the <QObject> header file as the following:


  typedef QList<QObject*> QObjectList;



The first child added is the first object in the list and the last child added is the last object in the list, i.e. new children are appended at the end.

Note that the list order changes when QWidget children are raised or lowered. A widget that is raised becomes the last object in the list, and a widget that is lowered becomes the first object in the list.

See also findChild(), findChildren(), parent(), and setParent().
*/
func (this *QObject) Children() *QObjectList {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QObject8childrenEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQObjectListFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQObjectList)
	return rv2
}

// /usr/include/qt/QtCore/qobject.h:210
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setParent(QObject *)

/*
Makes the object a child of parent.

See also parent() and children().
*/
func (this *QObject) SetParent(parent QObject_ITF /*777 QObject **/) {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QObject_PTR() != nil {
		convArg0 = parent.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN7QObject9setParentEPS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qobject.h:211
// index:0
// Public Visibility=Default Availability=Available
// [-2] void installEventFilter(QObject *)

/*
Installs an event filter filterObj on this object. For example:


  monitoredObj->installEventFilter(filterObj);



An event filter is an object that receives all events that are sent to this object. The filter can either stop the event or forward it to this object. The event filter filterObj receives events via its eventFilter() function. The eventFilter() function must return true if the event should be filtered, (i.e. stopped); otherwise it must return false.

If multiple event filters are installed on a single object, the filter that was installed last is activated first.

Here's a KeyPressEater class that eats the key presses of its monitored objects:


  class KeyPressEater : public QObject
  {
      Q_OBJECT
      ...

  protected:
      bool eventFilter(QObject *obj, QEvent *event);
  };

  bool KeyPressEater::eventFilter(QObject *obj, QEvent *event)
  {
      if (event->type() == QEvent::KeyPress) {
          QKeyEvent *keyEvent = static_cast<QKeyEvent *>(event);
          qDebug("Ate key press %d", keyEvent->key());
          return true;
      } else {
          // standard event processing
          return QObject::eventFilter(obj, event);
      }
  }



And here's how to install it on two widgets:


  KeyPressEater *keyPressEater = new KeyPressEater(this);
  QPushButton *pushButton = new QPushButton(this);
  QListView *listView = new QListView(this);

  pushButton->installEventFilter(keyPressEater);
  listView->installEventFilter(keyPressEater);



The QShortcut class, for example, uses this technique to intercept shortcut key presses.

Warning: If you delete the receiver object in your eventFilter() function, be sure to return true. If you return false, Qt sends the event to the deleted object and the program will crash.

Note that the filtering object must be in the same thread as this object. If filterObj is in a different thread, this function does nothing. If either filterObj or this object are moved to a different thread after calling this function, the event filter will not be called until both objects have the same thread affinity again (it is not removed).

See also removeEventFilter(), eventFilter(), and event().
*/
func (this *QObject) InstallEventFilter(filterObj QObject_ITF /*777 QObject **/) {
	var convArg0 unsafe.Pointer
	if filterObj != nil && filterObj.QObject_PTR() != nil {
		convArg0 = filterObj.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN7QObject18installEventFilterEPS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qobject.h:212
// index:0
// Public Visibility=Default Availability=Available
// [-2] void removeEventFilter(QObject *)

/*
Removes an event filter object obj from this object. The request is ignored if such an event filter has not been installed.

All event filters for this object are automatically removed when this object is destroyed.

It is always safe to remove an event filter, even during event filter activation (i.e. from the eventFilter() function).

See also installEventFilter(), eventFilter(), and event().
*/
func (this *QObject) RemoveEventFilter(obj QObject_ITF /*777 QObject **/) {
	var convArg0 unsafe.Pointer
	if obj != nil && obj.QObject_PTR() != nil {
		convArg0 = obj.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN7QObject17removeEventFilterEPS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qobject.h:214
// index:0
// Public static Visibility=Default Availability=Available
// [8] QMetaObject::Connection connect(const QObject *, const char *, const QObject *, const char *, Qt::ConnectionType)

/*
Creates a connection of the given type from the signal in the sender object to the method in the receiver object. Returns a handle to the connection that can be used to disconnect it later.

You must use the SIGNAL() and SLOT() macros when specifying the signal and the method, for example:


  QLabel *label = new QLabel;
  QScrollBar *scrollBar = new QScrollBar;
  QObject::connect(scrollBar, SIGNAL(valueChanged(int)),
                   label,  SLOT(setNum(int)));



This example ensures that the label always displays the current scroll bar value. Note that the signal and slots parameters must not contain any variable names, only the type. E.g. the following would not work and return false:


  // WRONG
  QObject::connect(scrollBar, SIGNAL(valueChanged(int value)),
                   label, SLOT(setNum(int value)));



A signal can also be connected to another signal:


  class MyWidget : public QWidget
  {
      Q_OBJECT

  public:
      MyWidget();

  signals:
      void buttonClicked();

  private:
      QPushButton *myButton;
  };

  MyWidget::MyWidget()
  {
      myButton = new QPushButton(this);
      connect(myButton, SIGNAL(clicked()),
              this, SIGNAL(buttonClicked()));
  }



In this example, the MyWidget constructor relays a signal from a private member variable, and makes it available under a name that relates to MyWidget.

A signal can be connected to many slots and signals. Many signals can be connected to one slot.

If a signal is connected to several slots, the slots are activated in the same order in which the connections were made, when the signal is emitted.

The function returns a QMetaObject::Connection that represents a handle to a connection if it successfully connects the signal to the slot. The connection handle will be invalid if it cannot create the connection, for example, if QObject is unable to verify the existence of either signal or method, or if their signatures aren't compatible. You can check if the handle is valid by casting it to a bool.

By default, a signal is emitted for every connection you make; two signals are emitted for duplicate connections. You can break all of these connections with a single disconnect() call. If you pass the Qt::UniqueConnection type, the connection will only be made if it is not a duplicate. If there is already a duplicate (exact same signal to the exact same slot on the same objects), the connection will fail and connect will return an invalid QMetaObject::Connection.

Note: Qt::UniqueConnections do not work for lambdas, non-member functions and functors; they only apply to connecting to member functions.

The optional type parameter describes the type of connection to establish. In particular, it determines whether a particular signal is delivered to a slot immediately or queued for delivery at a later time. If the signal is queued, the parameters must be of types that are known to Qt's meta-object system, because Qt needs to copy the arguments to store them in an event behind the scenes. If you try to use a queued connection and get the error message


  QObject::connect: Cannot queue arguments of type 'MyType'
  (Make sure 'MyType' is registered using qRegisterMetaType().)



call qRegisterMetaType() to register the data type before you establish the connection.

Note: This function is thread-safe.

See also disconnect(), sender(), qRegisterMetaType(), Q_DECLARE_METATYPE(), and Differences between String-Based and Functor-Based Connections.
*/
func (this *QObject) Connect(sender QObject_ITF /*777 const QObject **/, signal string, receiver QObject_ITF /*777 const QObject **/, member string, arg4 int) int {
	var convArg0 unsafe.Pointer
	if sender != nil && sender.QObject_PTR() != nil {
		convArg0 = sender.QObject_PTR().GetCthis()
	}
	var convArg1 = qtrt.CString(signal)
	defer qtrt.FreeMem(convArg1)
	var convArg2 unsafe.Pointer
	if receiver != nil && receiver.QObject_PTR() != nil {
		convArg2 = receiver.QObject_PTR().GetCthis()
	}
	var convArg3 = qtrt.CString(member)
	defer qtrt.FreeMem(convArg3)
	rv, err := qtrt.InvokeQtFunc6("_ZN7QObject7connectEPKS_PKcS1_S3_N2Qt14ConnectionTypeE", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, convArg2, convArg3, arg4)
	qtrt.ErrPrint(err, rv)
	return int(rv)
}
func QObject_Connect(sender QObject_ITF /*777 const QObject **/, signal string, receiver QObject_ITF /*777 const QObject **/, member string, arg4 int) int {
	var nilthis *QObject
	rv := nilthis.Connect(sender, signal, receiver, member, arg4)
	return rv
}

// /usr/include/qt/QtCore/qobject.h:214
// index:0
// Public static Visibility=Default Availability=Available
// [8] QMetaObject::Connection connect(const QObject *, const char *, const QObject *, const char *, Qt::ConnectionType)

/*
Creates a connection of the given type from the signal in the sender object to the method in the receiver object. Returns a handle to the connection that can be used to disconnect it later.

You must use the SIGNAL() and SLOT() macros when specifying the signal and the method, for example:


  QLabel *label = new QLabel;
  QScrollBar *scrollBar = new QScrollBar;
  QObject::connect(scrollBar, SIGNAL(valueChanged(int)),
                   label,  SLOT(setNum(int)));



This example ensures that the label always displays the current scroll bar value. Note that the signal and slots parameters must not contain any variable names, only the type. E.g. the following would not work and return false:


  // WRONG
  QObject::connect(scrollBar, SIGNAL(valueChanged(int value)),
                   label, SLOT(setNum(int value)));



A signal can also be connected to another signal:


  class MyWidget : public QWidget
  {
      Q_OBJECT

  public:
      MyWidget();

  signals:
      void buttonClicked();

  private:
      QPushButton *myButton;
  };

  MyWidget::MyWidget()
  {
      myButton = new QPushButton(this);
      connect(myButton, SIGNAL(clicked()),
              this, SIGNAL(buttonClicked()));
  }



In this example, the MyWidget constructor relays a signal from a private member variable, and makes it available under a name that relates to MyWidget.

A signal can be connected to many slots and signals. Many signals can be connected to one slot.

If a signal is connected to several slots, the slots are activated in the same order in which the connections were made, when the signal is emitted.

The function returns a QMetaObject::Connection that represents a handle to a connection if it successfully connects the signal to the slot. The connection handle will be invalid if it cannot create the connection, for example, if QObject is unable to verify the existence of either signal or method, or if their signatures aren't compatible. You can check if the handle is valid by casting it to a bool.

By default, a signal is emitted for every connection you make; two signals are emitted for duplicate connections. You can break all of these connections with a single disconnect() call. If you pass the Qt::UniqueConnection type, the connection will only be made if it is not a duplicate. If there is already a duplicate (exact same signal to the exact same slot on the same objects), the connection will fail and connect will return an invalid QMetaObject::Connection.

Note: Qt::UniqueConnections do not work for lambdas, non-member functions and functors; they only apply to connecting to member functions.

The optional type parameter describes the type of connection to establish. In particular, it determines whether a particular signal is delivered to a slot immediately or queued for delivery at a later time. If the signal is queued, the parameters must be of types that are known to Qt's meta-object system, because Qt needs to copy the arguments to store them in an event behind the scenes. If you try to use a queued connection and get the error message


  QObject::connect: Cannot queue arguments of type 'MyType'
  (Make sure 'MyType' is registered using qRegisterMetaType().)



call qRegisterMetaType() to register the data type before you establish the connection.

Note: This function is thread-safe.

See also disconnect(), sender(), qRegisterMetaType(), Q_DECLARE_METATYPE(), and Differences between String-Based and Functor-Based Connections.
*/
func (this *QObject) Connect__(sender QObject_ITF /*777 const QObject **/, signal string, receiver QObject_ITF /*777 const QObject **/, member string) int {
	var convArg0 unsafe.Pointer
	if sender != nil && sender.QObject_PTR() != nil {
		convArg0 = sender.QObject_PTR().GetCthis()
	}
	var convArg1 = qtrt.CString(signal)
	defer qtrt.FreeMem(convArg1)
	var convArg2 unsafe.Pointer
	if receiver != nil && receiver.QObject_PTR() != nil {
		convArg2 = receiver.QObject_PTR().GetCthis()
	}
	var convArg3 = qtrt.CString(member)
	defer qtrt.FreeMem(convArg3)
	// arg: 4, Qt::ConnectionType=Elaborated, Qt::ConnectionType=Enum,
	arg4 := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN7QObject7connectEPKS_PKcS1_S3_N2Qt14ConnectionTypeE", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, convArg2, convArg3, arg4)
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtCore/qobject.h:217
// index:1
// Public static Visibility=Default Availability=Available
// [8] QMetaObject::Connection connect(const QObject *, const QMetaMethod &, const QObject *, const QMetaMethod &, Qt::ConnectionType)

/*
Creates a connection of the given type from the signal in the sender object to the method in the receiver object. Returns a handle to the connection that can be used to disconnect it later.

You must use the SIGNAL() and SLOT() macros when specifying the signal and the method, for example:


  QLabel *label = new QLabel;
  QScrollBar *scrollBar = new QScrollBar;
  QObject::connect(scrollBar, SIGNAL(valueChanged(int)),
                   label,  SLOT(setNum(int)));



This example ensures that the label always displays the current scroll bar value. Note that the signal and slots parameters must not contain any variable names, only the type. E.g. the following would not work and return false:


  // WRONG
  QObject::connect(scrollBar, SIGNAL(valueChanged(int value)),
                   label, SLOT(setNum(int value)));



A signal can also be connected to another signal:


  class MyWidget : public QWidget
  {
      Q_OBJECT

  public:
      MyWidget();

  signals:
      void buttonClicked();

  private:
      QPushButton *myButton;
  };

  MyWidget::MyWidget()
  {
      myButton = new QPushButton(this);
      connect(myButton, SIGNAL(clicked()),
              this, SIGNAL(buttonClicked()));
  }



In this example, the MyWidget constructor relays a signal from a private member variable, and makes it available under a name that relates to MyWidget.

A signal can be connected to many slots and signals. Many signals can be connected to one slot.

If a signal is connected to several slots, the slots are activated in the same order in which the connections were made, when the signal is emitted.

The function returns a QMetaObject::Connection that represents a handle to a connection if it successfully connects the signal to the slot. The connection handle will be invalid if it cannot create the connection, for example, if QObject is unable to verify the existence of either signal or method, or if their signatures aren't compatible. You can check if the handle is valid by casting it to a bool.

By default, a signal is emitted for every connection you make; two signals are emitted for duplicate connections. You can break all of these connections with a single disconnect() call. If you pass the Qt::UniqueConnection type, the connection will only be made if it is not a duplicate. If there is already a duplicate (exact same signal to the exact same slot on the same objects), the connection will fail and connect will return an invalid QMetaObject::Connection.

Note: Qt::UniqueConnections do not work for lambdas, non-member functions and functors; they only apply to connecting to member functions.

The optional type parameter describes the type of connection to establish. In particular, it determines whether a particular signal is delivered to a slot immediately or queued for delivery at a later time. If the signal is queued, the parameters must be of types that are known to Qt's meta-object system, because Qt needs to copy the arguments to store them in an event behind the scenes. If you try to use a queued connection and get the error message


  QObject::connect: Cannot queue arguments of type 'MyType'
  (Make sure 'MyType' is registered using qRegisterMetaType().)



call qRegisterMetaType() to register the data type before you establish the connection.

Note: This function is thread-safe.

See also disconnect(), sender(), qRegisterMetaType(), Q_DECLARE_METATYPE(), and Differences between String-Based and Functor-Based Connections.
*/
func (this *QObject) Connect_1(sender QObject_ITF /*777 const QObject **/, signal QMetaMethod_ITF, receiver QObject_ITF /*777 const QObject **/, method QMetaMethod_ITF, type_ int) int {
	var convArg0 unsafe.Pointer
	if sender != nil && sender.QObject_PTR() != nil {
		convArg0 = sender.QObject_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if signal != nil && signal.QMetaMethod_PTR() != nil {
		convArg1 = signal.QMetaMethod_PTR().GetCthis()
	}
	var convArg2 unsafe.Pointer
	if receiver != nil && receiver.QObject_PTR() != nil {
		convArg2 = receiver.QObject_PTR().GetCthis()
	}
	var convArg3 unsafe.Pointer
	if method != nil && method.QMetaMethod_PTR() != nil {
		convArg3 = method.QMetaMethod_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN7QObject7connectEPKS_RK11QMetaMethodS1_S4_N2Qt14ConnectionTypeE", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, convArg2, convArg3, type_)
	qtrt.ErrPrint(err, rv)
	return int(rv)
}
func QObject_Connect_1(sender QObject_ITF /*777 const QObject **/, signal QMetaMethod_ITF, receiver QObject_ITF /*777 const QObject **/, method QMetaMethod_ITF, type_ int) int {
	var nilthis *QObject
	rv := nilthis.Connect_1(sender, signal, receiver, method, type_)
	return rv
}

// /usr/include/qt/QtCore/qobject.h:217
// index:1
// Public static Visibility=Default Availability=Available
// [8] QMetaObject::Connection connect(const QObject *, const QMetaMethod &, const QObject *, const QMetaMethod &, Qt::ConnectionType)

/*
Creates a connection of the given type from the signal in the sender object to the method in the receiver object. Returns a handle to the connection that can be used to disconnect it later.

You must use the SIGNAL() and SLOT() macros when specifying the signal and the method, for example:


  QLabel *label = new QLabel;
  QScrollBar *scrollBar = new QScrollBar;
  QObject::connect(scrollBar, SIGNAL(valueChanged(int)),
                   label,  SLOT(setNum(int)));



This example ensures that the label always displays the current scroll bar value. Note that the signal and slots parameters must not contain any variable names, only the type. E.g. the following would not work and return false:


  // WRONG
  QObject::connect(scrollBar, SIGNAL(valueChanged(int value)),
                   label, SLOT(setNum(int value)));



A signal can also be connected to another signal:


  class MyWidget : public QWidget
  {
      Q_OBJECT

  public:
      MyWidget();

  signals:
      void buttonClicked();

  private:
      QPushButton *myButton;
  };

  MyWidget::MyWidget()
  {
      myButton = new QPushButton(this);
      connect(myButton, SIGNAL(clicked()),
              this, SIGNAL(buttonClicked()));
  }



In this example, the MyWidget constructor relays a signal from a private member variable, and makes it available under a name that relates to MyWidget.

A signal can be connected to many slots and signals. Many signals can be connected to one slot.

If a signal is connected to several slots, the slots are activated in the same order in which the connections were made, when the signal is emitted.

The function returns a QMetaObject::Connection that represents a handle to a connection if it successfully connects the signal to the slot. The connection handle will be invalid if it cannot create the connection, for example, if QObject is unable to verify the existence of either signal or method, or if their signatures aren't compatible. You can check if the handle is valid by casting it to a bool.

By default, a signal is emitted for every connection you make; two signals are emitted for duplicate connections. You can break all of these connections with a single disconnect() call. If you pass the Qt::UniqueConnection type, the connection will only be made if it is not a duplicate. If there is already a duplicate (exact same signal to the exact same slot on the same objects), the connection will fail and connect will return an invalid QMetaObject::Connection.

Note: Qt::UniqueConnections do not work for lambdas, non-member functions and functors; they only apply to connecting to member functions.

The optional type parameter describes the type of connection to establish. In particular, it determines whether a particular signal is delivered to a slot immediately or queued for delivery at a later time. If the signal is queued, the parameters must be of types that are known to Qt's meta-object system, because Qt needs to copy the arguments to store them in an event behind the scenes. If you try to use a queued connection and get the error message


  QObject::connect: Cannot queue arguments of type 'MyType'
  (Make sure 'MyType' is registered using qRegisterMetaType().)



call qRegisterMetaType() to register the data type before you establish the connection.

Note: This function is thread-safe.

See also disconnect(), sender(), qRegisterMetaType(), Q_DECLARE_METATYPE(), and Differences between String-Based and Functor-Based Connections.
*/
func (this *QObject) Connect_1_(sender QObject_ITF /*777 const QObject **/, signal QMetaMethod_ITF, receiver QObject_ITF /*777 const QObject **/, method QMetaMethod_ITF) int {
	var convArg0 unsafe.Pointer
	if sender != nil && sender.QObject_PTR() != nil {
		convArg0 = sender.QObject_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if signal != nil && signal.QMetaMethod_PTR() != nil {
		convArg1 = signal.QMetaMethod_PTR().GetCthis()
	}
	var convArg2 unsafe.Pointer
	if receiver != nil && receiver.QObject_PTR() != nil {
		convArg2 = receiver.QObject_PTR().GetCthis()
	}
	var convArg3 unsafe.Pointer
	if method != nil && method.QMetaMethod_PTR() != nil {
		convArg3 = method.QMetaMethod_PTR().GetCthis()
	}
	// arg: 4, Qt::ConnectionType=Elaborated, Qt::ConnectionType=Enum,
	type_ := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN7QObject7connectEPKS_RK11QMetaMethodS1_S4_N2Qt14ConnectionTypeE", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, convArg2, convArg3, type_)
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtCore/qobject.h:221
// index:2
// Public inline Visibility=Default Availability=Available
// [8] QMetaObject::Connection connect(const QObject *, const char *, const char *, Qt::ConnectionType) const

/*
Creates a connection of the given type from the signal in the sender object to the method in the receiver object. Returns a handle to the connection that can be used to disconnect it later.

You must use the SIGNAL() and SLOT() macros when specifying the signal and the method, for example:


  QLabel *label = new QLabel;
  QScrollBar *scrollBar = new QScrollBar;
  QObject::connect(scrollBar, SIGNAL(valueChanged(int)),
                   label,  SLOT(setNum(int)));



This example ensures that the label always displays the current scroll bar value. Note that the signal and slots parameters must not contain any variable names, only the type. E.g. the following would not work and return false:


  // WRONG
  QObject::connect(scrollBar, SIGNAL(valueChanged(int value)),
                   label, SLOT(setNum(int value)));



A signal can also be connected to another signal:


  class MyWidget : public QWidget
  {
      Q_OBJECT

  public:
      MyWidget();

  signals:
      void buttonClicked();

  private:
      QPushButton *myButton;
  };

  MyWidget::MyWidget()
  {
      myButton = new QPushButton(this);
      connect(myButton, SIGNAL(clicked()),
              this, SIGNAL(buttonClicked()));
  }



In this example, the MyWidget constructor relays a signal from a private member variable, and makes it available under a name that relates to MyWidget.

A signal can be connected to many slots and signals. Many signals can be connected to one slot.

If a signal is connected to several slots, the slots are activated in the same order in which the connections were made, when the signal is emitted.

The function returns a QMetaObject::Connection that represents a handle to a connection if it successfully connects the signal to the slot. The connection handle will be invalid if it cannot create the connection, for example, if QObject is unable to verify the existence of either signal or method, or if their signatures aren't compatible. You can check if the handle is valid by casting it to a bool.

By default, a signal is emitted for every connection you make; two signals are emitted for duplicate connections. You can break all of these connections with a single disconnect() call. If you pass the Qt::UniqueConnection type, the connection will only be made if it is not a duplicate. If there is already a duplicate (exact same signal to the exact same slot on the same objects), the connection will fail and connect will return an invalid QMetaObject::Connection.

Note: Qt::UniqueConnections do not work for lambdas, non-member functions and functors; they only apply to connecting to member functions.

The optional type parameter describes the type of connection to establish. In particular, it determines whether a particular signal is delivered to a slot immediately or queued for delivery at a later time. If the signal is queued, the parameters must be of types that are known to Qt's meta-object system, because Qt needs to copy the arguments to store them in an event behind the scenes. If you try to use a queued connection and get the error message


  QObject::connect: Cannot queue arguments of type 'MyType'
  (Make sure 'MyType' is registered using qRegisterMetaType().)



call qRegisterMetaType() to register the data type before you establish the connection.

Note: This function is thread-safe.

See also disconnect(), sender(), qRegisterMetaType(), Q_DECLARE_METATYPE(), and Differences between String-Based and Functor-Based Connections.
*/
func (this *QObject) Connect_2(sender QObject_ITF /*777 const QObject **/, signal string, member string, type_ int) int {
	var convArg0 unsafe.Pointer
	if sender != nil && sender.QObject_PTR() != nil {
		convArg0 = sender.QObject_PTR().GetCthis()
	}
	var convArg1 = qtrt.CString(signal)
	defer qtrt.FreeMem(convArg1)
	var convArg2 = qtrt.CString(member)
	defer qtrt.FreeMem(convArg2)
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QObject7connectEPKS_PKcS3_N2Qt14ConnectionTypeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2, type_)
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtCore/qobject.h:221
// index:2
// Public inline Visibility=Default Availability=Available
// [8] QMetaObject::Connection connect(const QObject *, const char *, const char *, Qt::ConnectionType) const

/*
Creates a connection of the given type from the signal in the sender object to the method in the receiver object. Returns a handle to the connection that can be used to disconnect it later.

You must use the SIGNAL() and SLOT() macros when specifying the signal and the method, for example:


  QLabel *label = new QLabel;
  QScrollBar *scrollBar = new QScrollBar;
  QObject::connect(scrollBar, SIGNAL(valueChanged(int)),
                   label,  SLOT(setNum(int)));



This example ensures that the label always displays the current scroll bar value. Note that the signal and slots parameters must not contain any variable names, only the type. E.g. the following would not work and return false:


  // WRONG
  QObject::connect(scrollBar, SIGNAL(valueChanged(int value)),
                   label, SLOT(setNum(int value)));



A signal can also be connected to another signal:


  class MyWidget : public QWidget
  {
      Q_OBJECT

  public:
      MyWidget();

  signals:
      void buttonClicked();

  private:
      QPushButton *myButton;
  };

  MyWidget::MyWidget()
  {
      myButton = new QPushButton(this);
      connect(myButton, SIGNAL(clicked()),
              this, SIGNAL(buttonClicked()));
  }



In this example, the MyWidget constructor relays a signal from a private member variable, and makes it available under a name that relates to MyWidget.

A signal can be connected to many slots and signals. Many signals can be connected to one slot.

If a signal is connected to several slots, the slots are activated in the same order in which the connections were made, when the signal is emitted.

The function returns a QMetaObject::Connection that represents a handle to a connection if it successfully connects the signal to the slot. The connection handle will be invalid if it cannot create the connection, for example, if QObject is unable to verify the existence of either signal or method, or if their signatures aren't compatible. You can check if the handle is valid by casting it to a bool.

By default, a signal is emitted for every connection you make; two signals are emitted for duplicate connections. You can break all of these connections with a single disconnect() call. If you pass the Qt::UniqueConnection type, the connection will only be made if it is not a duplicate. If there is already a duplicate (exact same signal to the exact same slot on the same objects), the connection will fail and connect will return an invalid QMetaObject::Connection.

Note: Qt::UniqueConnections do not work for lambdas, non-member functions and functors; they only apply to connecting to member functions.

The optional type parameter describes the type of connection to establish. In particular, it determines whether a particular signal is delivered to a slot immediately or queued for delivery at a later time. If the signal is queued, the parameters must be of types that are known to Qt's meta-object system, because Qt needs to copy the arguments to store them in an event behind the scenes. If you try to use a queued connection and get the error message


  QObject::connect: Cannot queue arguments of type 'MyType'
  (Make sure 'MyType' is registered using qRegisterMetaType().)



call qRegisterMetaType() to register the data type before you establish the connection.

Note: This function is thread-safe.

See also disconnect(), sender(), qRegisterMetaType(), Q_DECLARE_METATYPE(), and Differences between String-Based and Functor-Based Connections.
*/
func (this *QObject) Connect_2_(sender QObject_ITF /*777 const QObject **/, signal string, member string) int {
	var convArg0 unsafe.Pointer
	if sender != nil && sender.QObject_PTR() != nil {
		convArg0 = sender.QObject_PTR().GetCthis()
	}
	var convArg1 = qtrt.CString(signal)
	defer qtrt.FreeMem(convArg1)
	var convArg2 = qtrt.CString(member)
	defer qtrt.FreeMem(convArg2)
	// arg: 3, Qt::ConnectionType=Elaborated, Qt::ConnectionType=Enum,
	type_ := 0
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QObject7connectEPKS_PKcS3_N2Qt14ConnectionTypeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2, type_)
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtCore/qobject.h:343
// index:0
// Public static Visibility=Default Availability=Available
// [1] bool disconnect(const QObject *, const char *, const QObject *, const char *)

/*
Disconnects signal in object sender from method in object receiver. Returns true if the connection is successfully broken; otherwise returns false.

A signal-slot connection is removed when either of the objects involved are destroyed.

disconnect() is typically used in three ways, as the following examples demonstrate.
*/
func (this *QObject) Disconnect(sender QObject_ITF /*777 const QObject **/, signal string, receiver QObject_ITF /*777 const QObject **/, member string) bool {
	var convArg0 unsafe.Pointer
	if sender != nil && sender.QObject_PTR() != nil {
		convArg0 = sender.QObject_PTR().GetCthis()
	}
	var convArg1 = qtrt.CString(signal)
	defer qtrt.FreeMem(convArg1)
	var convArg2 unsafe.Pointer
	if receiver != nil && receiver.QObject_PTR() != nil {
		convArg2 = receiver.QObject_PTR().GetCthis()
	}
	var convArg3 = qtrt.CString(member)
	defer qtrt.FreeMem(convArg3)
	rv, err := qtrt.InvokeQtFunc6("_ZN7QObject10disconnectEPKS_PKcS1_S3_", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, convArg2, convArg3)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}
func QObject_Disconnect(sender QObject_ITF /*777 const QObject **/, signal string, receiver QObject_ITF /*777 const QObject **/, member string) bool {
	var nilthis *QObject
	rv := nilthis.Disconnect(sender, signal, receiver, member)
	return rv
}

// /usr/include/qt/QtCore/qobject.h:345
// index:1
// Public static Visibility=Default Availability=Available
// [1] bool disconnect(const QObject *, const QMetaMethod &, const QObject *, const QMetaMethod &)

/*
Disconnects signal in object sender from method in object receiver. Returns true if the connection is successfully broken; otherwise returns false.

A signal-slot connection is removed when either of the objects involved are destroyed.

disconnect() is typically used in three ways, as the following examples demonstrate.
*/
func (this *QObject) Disconnect_1(sender QObject_ITF /*777 const QObject **/, signal QMetaMethod_ITF, receiver QObject_ITF /*777 const QObject **/, member QMetaMethod_ITF) bool {
	var convArg0 unsafe.Pointer
	if sender != nil && sender.QObject_PTR() != nil {
		convArg0 = sender.QObject_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if signal != nil && signal.QMetaMethod_PTR() != nil {
		convArg1 = signal.QMetaMethod_PTR().GetCthis()
	}
	var convArg2 unsafe.Pointer
	if receiver != nil && receiver.QObject_PTR() != nil {
		convArg2 = receiver.QObject_PTR().GetCthis()
	}
	var convArg3 unsafe.Pointer
	if member != nil && member.QMetaMethod_PTR() != nil {
		convArg3 = member.QMetaMethod_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN7QObject10disconnectEPKS_RK11QMetaMethodS1_S4_", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, convArg2, convArg3)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}
func QObject_Disconnect_1(sender QObject_ITF /*777 const QObject **/, signal QMetaMethod_ITF, receiver QObject_ITF /*777 const QObject **/, member QMetaMethod_ITF) bool {
	var nilthis *QObject
	rv := nilthis.Disconnect_1(sender, signal, receiver, member)
	return rv
}

// /usr/include/qt/QtCore/qobject.h:347
// index:2
// Public inline Visibility=Default Availability=Available
// [1] bool disconnect(const char *, const QObject *, const char *) const

/*
Disconnects signal in object sender from method in object receiver. Returns true if the connection is successfully broken; otherwise returns false.

A signal-slot connection is removed when either of the objects involved are destroyed.

disconnect() is typically used in three ways, as the following examples demonstrate.
*/
func (this *QObject) Disconnect_2(signal string, receiver QObject_ITF /*777 const QObject **/, member string) bool {
	var convArg0 = qtrt.CString(signal)
	defer qtrt.FreeMem(convArg0)
	var convArg1 unsafe.Pointer
	if receiver != nil && receiver.QObject_PTR() != nil {
		convArg1 = receiver.QObject_PTR().GetCthis()
	}
	var convArg2 = qtrt.CString(member)
	defer qtrt.FreeMem(convArg2)
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QObject10disconnectEPKcPKS_S1_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qobject.h:347
// index:2
// Public inline Visibility=Default Availability=Available
// [1] bool disconnect(const char *, const QObject *, const char *) const

/*
Disconnects signal in object sender from method in object receiver. Returns true if the connection is successfully broken; otherwise returns false.

A signal-slot connection is removed when either of the objects involved are destroyed.

disconnect() is typically used in three ways, as the following examples demonstrate.
*/
func (this *QObject) Disconnect_2_() bool {
	// arg: 0, const char *=Pointer, =Invalid,
	var convArg0 unsafe.Pointer
	// arg: 1, const QObject *=Pointer, QObject=Record,
	var convArg1 unsafe.Pointer
	// arg: 2, const char *=Pointer, =Invalid,
	var convArg2 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QObject10disconnectEPKcPKS_S1_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qobject.h:347
// index:2
// Public inline Visibility=Default Availability=Available
// [1] bool disconnect(const char *, const QObject *, const char *) const

/*
Disconnects signal in object sender from method in object receiver. Returns true if the connection is successfully broken; otherwise returns false.

A signal-slot connection is removed when either of the objects involved are destroyed.

disconnect() is typically used in three ways, as the following examples demonstrate.
*/
func (this *QObject) Disconnect_2_1(signal string) bool {
	var convArg0 = qtrt.CString(signal)
	defer qtrt.FreeMem(convArg0)
	// arg: 1, const QObject *=Pointer, QObject=Record,
	var convArg1 unsafe.Pointer
	// arg: 2, const char *=Pointer, =Invalid,
	var convArg2 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QObject10disconnectEPKcPKS_S1_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qobject.h:347
// index:2
// Public inline Visibility=Default Availability=Available
// [1] bool disconnect(const char *, const QObject *, const char *) const

/*
Disconnects signal in object sender from method in object receiver. Returns true if the connection is successfully broken; otherwise returns false.

A signal-slot connection is removed when either of the objects involved are destroyed.

disconnect() is typically used in three ways, as the following examples demonstrate.
*/
func (this *QObject) Disconnect_2_2(signal string, receiver QObject_ITF /*777 const QObject **/) bool {
	var convArg0 = qtrt.CString(signal)
	defer qtrt.FreeMem(convArg0)
	var convArg1 unsafe.Pointer
	if receiver != nil && receiver.QObject_PTR() != nil {
		convArg1 = receiver.QObject_PTR().GetCthis()
	}
	// arg: 2, const char *=Pointer, =Invalid,
	var convArg2 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QObject10disconnectEPKcPKS_S1_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qobject.h:350
// index:3
// Public inline Visibility=Default Availability=Available
// [1] bool disconnect(const QObject *, const char *) const

/*
Disconnects signal in object sender from method in object receiver. Returns true if the connection is successfully broken; otherwise returns false.

A signal-slot connection is removed when either of the objects involved are destroyed.

disconnect() is typically used in three ways, as the following examples demonstrate.
*/
func (this *QObject) Disconnect_3(receiver QObject_ITF /*777 const QObject **/, member string) bool {
	var convArg0 unsafe.Pointer
	if receiver != nil && receiver.QObject_PTR() != nil {
		convArg0 = receiver.QObject_PTR().GetCthis()
	}
	var convArg1 = qtrt.CString(member)
	defer qtrt.FreeMem(convArg1)
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QObject10disconnectEPKS_PKc", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qobject.h:350
// index:3
// Public inline Visibility=Default Availability=Available
// [1] bool disconnect(const QObject *, const char *) const

/*
Disconnects signal in object sender from method in object receiver. Returns true if the connection is successfully broken; otherwise returns false.

A signal-slot connection is removed when either of the objects involved are destroyed.

disconnect() is typically used in three ways, as the following examples demonstrate.
*/
func (this *QObject) Disconnect_3_(receiver QObject_ITF /*777 const QObject **/) bool {
	var convArg0 unsafe.Pointer
	if receiver != nil && receiver.QObject_PTR() != nil {
		convArg0 = receiver.QObject_PTR().GetCthis()
	}
	// arg: 1, const char *=Pointer, =Invalid,
	var convArg1 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QObject10disconnectEPKS_PKc", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qobject.h:391
// index:0
// Public Visibility=Default Availability=Available
// [-2] void dumpObjectTree()

/*
Dumps a tree of children to the debug output.

Note: before Qt 5.9, this function was not const.

See also dumpObjectInfo().
*/
func (this *QObject) DumpObjectTree() {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QObject14dumpObjectTreeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qobject.h:394
// index:1
// Public Visibility=Default Availability=Available
// [-2] void dumpObjectTree() const

/*
Dumps a tree of children to the debug output.

Note: before Qt 5.9, this function was not const.

See also dumpObjectInfo().
*/
func (this *QObject) DumpObjectTree_1() {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QObject14dumpObjectTreeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qobject.h:392
// index:0
// Public Visibility=Default Availability=Available
// [-2] void dumpObjectInfo()

/*
Dumps information about signal connections, etc. for this object to the debug output.

Note: before Qt 5.9, this function was not const.

See also dumpObjectTree().
*/
func (this *QObject) DumpObjectInfo() {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QObject14dumpObjectInfoEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qobject.h:395
// index:1
// Public Visibility=Default Availability=Available
// [-2] void dumpObjectInfo() const

/*
Dumps information about signal connections, etc. for this object to the debug output.

Note: before Qt 5.9, this function was not const.

See also dumpObjectTree().
*/
func (this *QObject) DumpObjectInfo_1() {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QObject14dumpObjectInfoEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qobject.h:398
// index:0
// Public Visibility=Default Availability=Available
// [1] bool setProperty(const char *, const QVariant &)

/*
Sets the value of the object's name property to value.

If the property is defined in the class using Q_PROPERTY then true is returned on success and false otherwise. If the property is not defined using Q_PROPERTY, and therefore not listed in the meta-object, it is added as a dynamic property and false is returned.

Information about all available properties is provided through the metaObject() and dynamicPropertyNames().

Dynamic properties can be queried again using property() and can be removed by setting the property value to an invalid QVariant. Changing the value of a dynamic property causes a QDynamicPropertyChangeEvent to be sent to the object.

Note: Dynamic properties starting with "_q_" are reserved for internal purposes.

See also property(), metaObject(), dynamicPropertyNames(), and QMetaProperty::write().
*/
func (this *QObject) SetProperty(name string, value QVariant_ITF) bool {
	var convArg0 = qtrt.CString(name)
	defer qtrt.FreeMem(convArg0)
	var convArg1 unsafe.Pointer
	if value != nil && value.QVariant_PTR() != nil {
		convArg1 = value.QVariant_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN7QObject11setPropertyEPKcRK8QVariant", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qobject.h:399
// index:0
// Public Visibility=Default Availability=Available
// [16] QVariant property(const char *) const

/*
Returns the value of the object's name property.

If no such property exists, the returned variant is invalid.

Information about all available properties is provided through the metaObject() and dynamicPropertyNames().

See also setProperty(), QVariant::isValid(), metaObject(), and dynamicPropertyNames().
*/
func (this *QObject) Property(name string) *QVariant /*123*/ {
	var convArg0 = qtrt.CString(name)
	defer qtrt.FreeMem(convArg0)
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QObject8propertyEPKc", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQVariantFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQVariant)
	return rv2
}

// /usr/include/qt/QtCore/qobject.h:404
// index:0
// Public static Visibility=Default Availability=Available
// [4] uint registerUserData()

/*

 */
func (this *QObject) RegisterUserData() uint {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QObject16registerUserDataEv", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	return uint(rv) // 222
}
func QObject_RegisterUserData() uint {
	var nilthis *QObject
	rv := nilthis.RegisterUserData()
	return rv
}

// /usr/include/qt/QtCore/qobject.h:405
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setUserData(uint, QObjectUserData *)

/*

 */
func (this *QObject) SetUserData(id uint, data QObjectUserData_ITF /*777 QObjectUserData **/) {
	var convArg1 unsafe.Pointer
	if data != nil && data.QObjectUserData_PTR() != nil {
		convArg1 = data.QObjectUserData_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN7QObject11setUserDataEjP15QObjectUserData", qtrt.FFI_TYPE_POINTER, this.GetCthis(), id, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qobject.h:406
// index:0
// Public Visibility=Default Availability=Available
// [8] QObjectUserData * userData(uint) const

/*

 */
func (this *QObject) UserData(id uint) *QObjectUserData /*777 QObjectUserData **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QObject8userDataEj", qtrt.FFI_TYPE_POINTER, this.GetCthis(), id)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQObjectUserDataFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtCore/qobject.h:410
// index:0
// Public Visibility=Default Availability=Available
// [-2] void destroyed(QObject *)

/*
This signal is emitted immediately before the object obj is destroyed, and can not be blocked.

All the objects's children are destroyed immediately after this signal is emitted.

See also deleteLater() and QPointer.
*/
func (this *QObject) Destroyed(arg0 QObject_ITF /*777 QObject **/) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QObject_PTR() != nil {
		convArg0 = arg0.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN7QObject9destroyedEPS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qobject.h:410
// index:0
// Public Visibility=Default Availability=Available
// [-2] void destroyed(QObject *)

/*
This signal is emitted immediately before the object obj is destroyed, and can not be blocked.

All the objects's children are destroyed immediately after this signal is emitted.

See also deleteLater() and QPointer.
*/
func (this *QObject) Destroyed__() {
	// arg: 0, QObject *=Pointer, QObject=Record,
	var convArg0 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN7QObject9destroyedEPS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qobject.h:414
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QObject * parent() const

/*
Returns a pointer to the parent object.

See also setParent() and children().
*/
func (this *QObject) Parent() *QObject /*777 QObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QObject6parentEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtCore/qobject.h:416
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool inherits(const char *) const

/*
Returns true if this object is an instance of a class that inherits className or a QObject subclass that inherits className; otherwise returns false.

A class is considered to inherit itself.

Example:


  QTimer *timer = new QTimer;         // QTimer inherits QObject
  timer->inherits("QTimer");          // returns true
  timer->inherits("QObject");         // returns true
  timer->inherits("QAbstractButton"); // returns false

  // QVBoxLayout inherits QObject and QLayoutItem
  QVBoxLayout *layout = new QVBoxLayout;
  layout->inherits("QObject");        // returns true
  layout->inherits("QLayoutItem");    // returns true (even though QLayoutItem is not a QObject)



If you need to determine whether an object is an instance of a particular class for the purpose of casting it, consider using qobject_cast<Type *>(object) instead.

See also metaObject() and qobject_cast().
*/
func (this *QObject) Inherits(classname string) bool {
	var convArg0 = qtrt.CString(classname)
	defer qtrt.FreeMem(convArg0)
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QObject8inheritsEPKc", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qobject.h:420
// index:0
// Public Visibility=Default Availability=Available
// [-2] void deleteLater()

/*
Schedules this object for deletion.

The object will be deleted when control returns to the event loop. If the event loop is not running when this function is called (e.g. deleteLater() is called on an object before QCoreApplication::exec()), the object will be deleted once the event loop is started. If deleteLater() is called after the main event loop has stopped, the object will not be deleted. Since Qt 4.8, if deleteLater() is called on an object that lives in a thread with no running event loop, the object will be destroyed when the thread finishes.

Note that entering and leaving a new event loop (e.g., by opening a modal dialog) will not perform the deferred deletion; for the object to be deleted, the control must return to the event loop from which deleteLater() was called.

Note: It is safe to call this function more than once; when the first deferred deletion event is delivered, any pending events for the object are removed from the event queue.

See also destroyed() and QPointer.
*/
func (this *QObject) DeleteLater() {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QObject11deleteLaterEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qobject.h:423
// index:0
// Protected Visibility=Default Availability=Available
// [8] QObject * sender() const

/*
Returns a pointer to the object that sent the signal, if called in a slot activated by a signal; otherwise it returns 0. The pointer is valid only during the execution of the slot that calls this function from this object's thread context.

The pointer returned by this function becomes invalid if the sender is destroyed, or if the slot is disconnected from the sender's signal.

Warning: This function violates the object-oriented principle of modularity. However, getting access to the sender might be useful when many signals are connected to a single slot.

Warning: As mentioned above, the return value of this function is not valid when the slot is called via a Qt::DirectConnection from a thread different from this object's thread. Do not use this function in this type of scenario.

See also senderSignalIndex().
*/
func (this *QObject) Sender() *QObject /*777 QObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QObject6senderEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtCore/qobject.h:424
// index:0
// Protected Visibility=Default Availability=Available
// [4] int senderSignalIndex() const

/*
Returns the meta-method index of the signal that called the currently executing slot, which is a member of the class returned by sender(). If called outside of a slot activated by a signal, -1 is returned.

For signals with default parameters, this function will always return the index with all parameters, regardless of which was used with connect(). For example, the signal destroyed(QObject *obj = 0) will have two different indexes (with and without the parameter), but this function will always return the index with a parameter. This does not apply when overloading signals with different parameters.

Warning: This function violates the object-oriented principle of modularity. However, getting access to the signal index might be useful when many signals are connected to a single slot.

Warning: The return value of this function is not valid when the slot is called via a Qt::DirectConnection from a thread different from this object's thread. Do not use this function in this type of scenario.

This function was introduced in  Qt 4.8.

See also sender(), QMetaObject::indexOfSignal(), and QMetaObject::method().
*/
func (this *QObject) SenderSignalIndex() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QObject17senderSignalIndexEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qobject.h:425
// index:0
// Protected Visibility=Default Availability=Available
// [4] int receivers(const char *) const

/*
Returns the number of receivers connected to the signal.

Since both slots and signals can be used as receivers for signals, and the same connections can be made many times, the number of receivers is the same as the number of connections made from this signal.

When calling this function, you can use the SIGNAL() macro to pass a specific signal:


  if (receivers(SIGNAL(valueChanged(QByteArray))) > 0) {
      QByteArray data;
      get_the_value(&data);       // expensive operation
      emit valueChanged(data);
  }



Warning: This function violates the object-oriented principle of modularity. However, it might be useful when you need to perform expensive initialization only if something is connected to a signal.

See also isSignalConnected().
*/
func (this *QObject) Receivers(signal string) int {
	var convArg0 = qtrt.CString(signal)
	defer qtrt.FreeMem(convArg0)
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QObject9receiversEPKc", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qobject.h:426
// index:0
// Protected Visibility=Default Availability=Available
// [1] bool isSignalConnected(const QMetaMethod &) const

/*
Returns true if the signal is connected to at least one receiver, otherwise returns false.

signal must be a signal member of this object, otherwise the behaviour is undefined.


  static const QMetaMethod valueChangedSignal = QMetaMethod::fromSignal(&MyObject::valueChanged);
  if (isSignalConnected(valueChangedSignal)) {
      QByteArray data;
      data = get_the_value();       // expensive operation
      emit valueChanged(data);
  }



As the code snippet above illustrates, you can use this function to avoid emitting a signal that nobody listens to.

Warning: This function violates the object-oriented principle of modularity. However, it might be useful when you need to perform expensive initialization only if something is connected to a signal.

This function was introduced in  Qt 5.0.
*/
func (this *QObject) IsSignalConnected(signal QMetaMethod_ITF) bool {
	var convArg0 unsafe.Pointer
	if signal != nil && signal.QMetaMethod_PTR() != nil {
		convArg0 = signal.QMetaMethod_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QObject17isSignalConnectedERK11QMetaMethod", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qobject.h:428
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void timerEvent(QTimerEvent *)

/*
This event handler can be reimplemented in a subclass to receive timer events for the object.

QTimer provides a higher-level interface to the timer functionality, and also more general information about timers. The timer event is passed in the event parameter.

See also startTimer(), killTimer(), and event().
*/
func (this *QObject) TimerEvent(event QTimerEvent_ITF /*777 QTimerEvent **/) {
	var convArg0 unsafe.Pointer
	if event != nil && event.QTimerEvent_PTR() != nil {
		convArg0 = event.QTimerEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN7QObject10timerEventEP11QTimerEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qobject.h:429
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void childEvent(QChildEvent *)

/*
This event handler can be reimplemented in a subclass to receive child events. The event is passed in the event parameter.

QEvent::ChildAdded and QEvent::ChildRemoved events are sent to objects when children are added or removed. In both cases you can only rely on the child being a QObject, or if isWidgetType() returns true, a QWidget. (This is because, in the ChildAdded case, the child is not yet fully constructed, and in the ChildRemoved case it might have been destructed already).

QEvent::ChildPolished events are sent to widgets when children are polished, or when polished children are added. If you receive a child polished event, the child's construction is usually completed. However, this is not guaranteed, and multiple polish events may be delivered during the execution of a widget's constructor.

For every child widget, you receive one ChildAdded event, zero or more ChildPolished events, and one ChildRemoved event.

The ChildPolished event is omitted if a child is removed immediately after it is added. If a child is polished several times during construction and destruction, you may receive several child polished events for the same child, each time with a different virtual table.

See also event().
*/
func (this *QObject) ChildEvent(event QChildEvent_ITF /*777 QChildEvent **/) {
	var convArg0 unsafe.Pointer
	if event != nil && event.QChildEvent_PTR() != nil {
		convArg0 = event.QChildEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN7QObject10childEventEP11QChildEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qobject.h:430
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void customEvent(QEvent *)

/*
This event handler can be reimplemented in a subclass to receive custom events. Custom events are user-defined events with a type value at least as large as the QEvent::User item of the QEvent::Type enum, and is typically a QEvent subclass. The event is passed in the event parameter.

See also event() and QEvent.
*/
func (this *QObject) CustomEvent(event QEvent_ITF /*777 QEvent **/) {
	var convArg0 unsafe.Pointer
	if event != nil && event.QEvent_PTR() != nil {
		convArg0 = event.QEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN7QObject11customEventEP6QEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qobject.h:432
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void connectNotify(const QMetaMethod &)

/*
This virtual function is called when something has been connected to signal in this object.

If you want to compare signal with a specific signal, you can use QMetaMethod::fromSignal() as follows:


  if (signal == QMetaMethod::fromSignal(&MyObject::valueChanged)) {
      // signal is valueChanged
  }



Warning: This function violates the object-oriented principle of modularity. However, it might be useful when you need to perform expensive initialization only if something is connected to a signal.

Warning: This function is called from the thread which performs the connection, which may be a different thread from the thread in which this object lives.

This function was introduced in  Qt 5.0.

See also connect() and disconnectNotify().
*/
func (this *QObject) ConnectNotify(signal QMetaMethod_ITF) {
	var convArg0 unsafe.Pointer
	if signal != nil && signal.QMetaMethod_PTR() != nil {
		convArg0 = signal.QMetaMethod_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN7QObject13connectNotifyERK11QMetaMethod", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qobject.h:433
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void disconnectNotify(const QMetaMethod &)

/*
This virtual function is called when something has been disconnected from signal in this object.

See connectNotify() for an example of how to compare signal with a specific signal.

If all signals were disconnected from this object (e.g., the signal argument to disconnect() was 0), disconnectNotify() is only called once, and the signal will be an invalid QMetaMethod (QMetaMethod::isValid() returns false).

Warning: This function violates the object-oriented principle of modularity. However, it might be useful for optimizing access to expensive resources.

Warning: This function is called from the thread which performs the disconnection, which may be a different thread from the thread in which this object lives. This function may also be called with a QObject internal mutex locked. It is therefore not allowed to re-enter any of any QObject functions from your reimplementation and if you lock a mutex in your reimplementation, make sure that you don't call QObject functions with that mutex held in other places or it will result in a deadlock.

This function was introduced in  Qt 5.0.

See also disconnect() and connectNotify().
*/
func (this *QObject) DisconnectNotify(signal QMetaMethod_ITF) {
	var convArg0 unsafe.Pointer
	if signal != nil && signal.QMetaMethod_PTR() != nil {
		convArg0 = signal.QMetaMethod_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN7QObject16disconnectNotifyERK11QMetaMethod", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
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
