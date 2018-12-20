
C++ and Go have different name specific. 

Although make effort similar to origin C++ name, but still some converts:

### Qt components

All Qt component names are converted to go package names with lowercase:

* QtCore => qtcore.
* QtGui => qtgui.
* QtWidgets => qtwidgets.

### class name

class name is the same.

### method name:
* title the first charactor: `show() => Show()`

* overloaded methods

  According to overload index, add a number suffix. The first `0` is omitted.
 
  `arg() => Arg1()`

* has default arguments methods
  
  First, qt.go use `pn` where `n` is a number suffix.
  And the `n` is omitted default arguments count. The first `0` is omitted.
  
  For C++:
  
  `void setShortcutEnabled(int id, bool enable = true);` 
  
  For Go, there will be two methods with different `pn` suffix:
  
  `func (this *QWidget) SetShortcutEnabled(id int, enable bool)`
  `func (this *QWidget) SetShortcutEnabledp(id int)`

### static methods

qt.go use `class_method` as the full method name.

For convenient, the class instance has a `method` method. Such as:

```
func (this *QCoreApplication) Flush()
func QCoreApplication_Flush()
```

### class constructor

Add `New` prefix to class name: `NewQPushButton(...)` 

### protected methods

Add `Inherit` prefix to method name: `InheritMouseMoveEvent(...)`

### class internal enums

qt.go use `class__enumname` as the full name.

For C++:

`QWidget::DrawWindowBackground`
`QWidget::DrawChildren`

For Go:

`const QWidget__DrawWindowBackground = 1`
`const QWidget__DrawChildren         = 2`

### qt namespace enums

qt.go use `namespace__enumname` as the full name:

For C++:

`Qt::black`
`Qt::white`

For Go:

`const Qt__black = 2`
`const Qt__white = 3`

### QString

Except `QString` class wrapper itself, all other `QString` are replaced with go `string`.

### QObject and subclass type arguments

Since it can passby all subclass type, so qt.go use interface type to emulate this feature.

For every Qt class, there is a corresponding interface: `QWidget_ITF`

Such one can passby any `QWidget` subclass instance just like in C++.

