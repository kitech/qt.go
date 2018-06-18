package qtwidgets

// /usr/include/qt/QtWidgets/qdockwidget.h
// #include <qdockwidget.h>
// #include <QtWidgets>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 0
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "log"
import "github.com/kitech/qt.go/qtrt"
import "github.com/kitech/qt.go/qtcore"
import "github.com/kitech/qt.go/qtgui"

//  ext block end

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
	if false {
		qtcore.KeepMe()
	}
	if false {
		qtgui.KeepMe()
	}
}

//  keep block end

//  body block begin
type QDockWidgetList struct {
	*qtrt.CObject
}

// QList<T> & operator=(const QList<T> &)
func (this *QDockWidgetList) Operator_equal_0() *QDockWidgetList {
	// QDockWidgetList_operator_equal_0()
	rv, err := qtrt.InvokeQtFunc6("QDockWidgetList_operator_equal_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QList<T> & operator=(QList<T> &&)
func (this *QDockWidgetList) Operator_equal_1() *QDockWidgetList {
	// QDockWidgetList_operator_equal_1()
	rv, err := qtrt.InvokeQtFunc6("QDockWidgetList_operator_equal_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// void swap(QList<T> &)
func (this *QDockWidgetList) Swap_0() {
	// QDockWidgetList_swap_0()
	rv, err := qtrt.InvokeQtFunc6("QDockWidgetList_swap_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// bool operator==(const QList<T> &)
func (this *QDockWidgetList) Operator_equal_equal_0() bool {
	// QDockWidgetList_operator_equal_equal_0()
	rv, err := qtrt.InvokeQtFunc6("QDockWidgetList_operator_equal_equal_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// bool operator!=(const QList<T> &)
func (this *QDockWidgetList) Operator_not_equal_0() bool {
	// QDockWidgetList_operator_not_equal_0()
	rv, err := qtrt.InvokeQtFunc6("QDockWidgetList_operator_not_equal_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// int size()
func (this *QDockWidgetList) Size_0() int {
	// QDockWidgetList_size_0()
	rv, err := qtrt.InvokeQtFunc6("QDockWidgetList_size_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// void detach()
func (this *QDockWidgetList) Detach_0() {
	// QDockWidgetList_detach_0()
	rv, err := qtrt.InvokeQtFunc6("QDockWidgetList_detach_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void detachShared()
func (this *QDockWidgetList) DetachShared_0() {
	// QDockWidgetList_detachShared_0()
	rv, err := qtrt.InvokeQtFunc6("QDockWidgetList_detachShared_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// bool isDetached()
func (this *QDockWidgetList) IsDetached_0() bool {
	// QDockWidgetList_isDetached_0()
	rv, err := qtrt.InvokeQtFunc6("QDockWidgetList_isDetached_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// void setSharable(bool)
func (this *QDockWidgetList) SetSharable_0() {
	// QDockWidgetList_setSharable_0()
	rv, err := qtrt.InvokeQtFunc6("QDockWidgetList_setSharable_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// bool isSharedWith(const QList<T> &)
func (this *QDockWidgetList) IsSharedWith_0() bool {
	// QDockWidgetList_isSharedWith_0()
	rv, err := qtrt.InvokeQtFunc6("QDockWidgetList_isSharedWith_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// bool isEmpty()
func (this *QDockWidgetList) IsEmpty_0() bool {
	// QDockWidgetList_isEmpty_0()
	rv, err := qtrt.InvokeQtFunc6("QDockWidgetList_isEmpty_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// void clear()
func (this *QDockWidgetList) Clear_0() {
	// QDockWidgetList_clear_0()
	rv, err := qtrt.InvokeQtFunc6("QDockWidgetList_clear_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// const T & at(int)
func (this *QDockWidgetList) At_0() *QDockWidget {
	// QDockWidgetList_at_0()
	rv, err := qtrt.InvokeQtFunc6("QDockWidgetList_at_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QDockWidget{}
}

// const T & operator[](int)
func (this *QDockWidgetList) Operator_get_index_0() *QDockWidget {
	// QDockWidgetList_operator_get_index_0()
	rv, err := qtrt.InvokeQtFunc6("QDockWidgetList_operator_get_index_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QDockWidget{}
}

// T & operator[](int)
func (this *QDockWidgetList) Operator_get_index_1() *QDockWidget {
	// QDockWidgetList_operator_get_index_1()
	rv, err := qtrt.InvokeQtFunc6("QDockWidgetList_operator_get_index_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QDockWidget{}
}

// void reserve(int)
func (this *QDockWidgetList) Reserve_0() {
	// QDockWidgetList_reserve_0()
	rv, err := qtrt.InvokeQtFunc6("QDockWidgetList_reserve_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void append(const T &)
func (this *QDockWidgetList) Append_0() {
	// QDockWidgetList_append_0()
	rv, err := qtrt.InvokeQtFunc6("QDockWidgetList_append_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void append(const QList<T> &)
func (this *QDockWidgetList) Append_1() {
	// QDockWidgetList_append_1()
	rv, err := qtrt.InvokeQtFunc6("QDockWidgetList_append_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void prepend(const T &)
func (this *QDockWidgetList) Prepend_0() {
	// QDockWidgetList_prepend_0()
	rv, err := qtrt.InvokeQtFunc6("QDockWidgetList_prepend_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void insert(int, const T &)
func (this *QDockWidgetList) Insert_0() {
	// QDockWidgetList_insert_0()
	rv, err := qtrt.InvokeQtFunc6("QDockWidgetList_insert_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void replace(int, const T &)
func (this *QDockWidgetList) Replace_0() {
	// QDockWidgetList_replace_0()
	rv, err := qtrt.InvokeQtFunc6("QDockWidgetList_replace_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void removeAt(int)
func (this *QDockWidgetList) RemoveAt_0() {
	// QDockWidgetList_removeAt_0()
	rv, err := qtrt.InvokeQtFunc6("QDockWidgetList_removeAt_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// int removeAll(const T &)
func (this *QDockWidgetList) RemoveAll_0() int {
	// QDockWidgetList_removeAll_0()
	rv, err := qtrt.InvokeQtFunc6("QDockWidgetList_removeAll_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// bool removeOne(const T &)
func (this *QDockWidgetList) RemoveOne_0() bool {
	// QDockWidgetList_removeOne_0()
	rv, err := qtrt.InvokeQtFunc6("QDockWidgetList_removeOne_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// T takeAt(int)
func (this *QDockWidgetList) TakeAt_0() *QDockWidget {
	// QDockWidgetList_takeAt_0()
	rv, err := qtrt.InvokeQtFunc6("QDockWidgetList_takeAt_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QDockWidget{}
}

// T takeFirst()
func (this *QDockWidgetList) TakeFirst_0() *QDockWidget {
	// QDockWidgetList_takeFirst_0()
	rv, err := qtrt.InvokeQtFunc6("QDockWidgetList_takeFirst_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QDockWidget{}
}

// T takeLast()
func (this *QDockWidgetList) TakeLast_0() *QDockWidget {
	// QDockWidgetList_takeLast_0()
	rv, err := qtrt.InvokeQtFunc6("QDockWidgetList_takeLast_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QDockWidget{}
}

// void move(int, int)
func (this *QDockWidgetList) Move_0() {
	// QDockWidgetList_move_0()
	rv, err := qtrt.InvokeQtFunc6("QDockWidgetList_move_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void swap(int, int)
func (this *QDockWidgetList) Swap_1() {
	// QDockWidgetList_swap_1()
	rv, err := qtrt.InvokeQtFunc6("QDockWidgetList_swap_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// int indexOf(const T &, int)
func (this *QDockWidgetList) IndexOf_0() int {
	// QDockWidgetList_indexOf_0()
	rv, err := qtrt.InvokeQtFunc6("QDockWidgetList_indexOf_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// int lastIndexOf(const T &, int)
func (this *QDockWidgetList) LastIndexOf_0() int {
	// QDockWidgetList_lastIndexOf_0()
	rv, err := qtrt.InvokeQtFunc6("QDockWidgetList_lastIndexOf_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// bool contains(const T &)
func (this *QDockWidgetList) Contains_0() bool {
	// QDockWidgetList_contains_0()
	rv, err := qtrt.InvokeQtFunc6("QDockWidgetList_contains_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// int count(const T &)
func (this *QDockWidgetList) Count_0() int {
	// QDockWidgetList_count_0()
	rv, err := qtrt.InvokeQtFunc6("QDockWidgetList_count_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// QList::iterator begin()
func (this *QDockWidgetList) Begin_0() {
	// QDockWidgetList_begin_0()
	rv, err := qtrt.InvokeQtFunc6("QDockWidgetList_begin_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_iterator begin()
func (this *QDockWidgetList) Begin_1() {
	// QDockWidgetList_begin_1()
	rv, err := qtrt.InvokeQtFunc6("QDockWidgetList_begin_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_iterator cbegin()
func (this *QDockWidgetList) Cbegin_0() {
	// QDockWidgetList_cbegin_0()
	rv, err := qtrt.InvokeQtFunc6("QDockWidgetList_cbegin_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_iterator constBegin()
func (this *QDockWidgetList) ConstBegin_0() {
	// QDockWidgetList_constBegin_0()
	rv, err := qtrt.InvokeQtFunc6("QDockWidgetList_constBegin_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::iterator end()
func (this *QDockWidgetList) End_0() {
	// QDockWidgetList_end_0()
	rv, err := qtrt.InvokeQtFunc6("QDockWidgetList_end_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_iterator end()
func (this *QDockWidgetList) End_1() {
	// QDockWidgetList_end_1()
	rv, err := qtrt.InvokeQtFunc6("QDockWidgetList_end_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_iterator cend()
func (this *QDockWidgetList) Cend_0() {
	// QDockWidgetList_cend_0()
	rv, err := qtrt.InvokeQtFunc6("QDockWidgetList_cend_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_iterator constEnd()
func (this *QDockWidgetList) ConstEnd_0() {
	// QDockWidgetList_constEnd_0()
	rv, err := qtrt.InvokeQtFunc6("QDockWidgetList_constEnd_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::reverse_iterator rbegin()
func (this *QDockWidgetList) Rbegin_0() {
	// QDockWidgetList_rbegin_0()
	rv, err := qtrt.InvokeQtFunc6("QDockWidgetList_rbegin_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::reverse_iterator rend()
func (this *QDockWidgetList) Rend_0() {
	// QDockWidgetList_rend_0()
	rv, err := qtrt.InvokeQtFunc6("QDockWidgetList_rend_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_reverse_iterator rbegin()
func (this *QDockWidgetList) Rbegin_1() {
	// QDockWidgetList_rbegin_1()
	rv, err := qtrt.InvokeQtFunc6("QDockWidgetList_rbegin_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_reverse_iterator rend()
func (this *QDockWidgetList) Rend_1() {
	// QDockWidgetList_rend_1()
	rv, err := qtrt.InvokeQtFunc6("QDockWidgetList_rend_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_reverse_iterator crbegin()
func (this *QDockWidgetList) Crbegin_0() {
	// QDockWidgetList_crbegin_0()
	rv, err := qtrt.InvokeQtFunc6("QDockWidgetList_crbegin_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_reverse_iterator crend()
func (this *QDockWidgetList) Crend_0() {
	// QDockWidgetList_crend_0()
	rv, err := qtrt.InvokeQtFunc6("QDockWidgetList_crend_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::iterator insert(QList::iterator, const T &)
func (this *QDockWidgetList) Insert_1() {
	// QDockWidgetList_insert_1()
	rv, err := qtrt.InvokeQtFunc6("QDockWidgetList_insert_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::iterator erase(QList::iterator)
func (this *QDockWidgetList) Erase_0() {
	// QDockWidgetList_erase_0()
	rv, err := qtrt.InvokeQtFunc6("QDockWidgetList_erase_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::iterator erase(QList::iterator, QList::iterator)
func (this *QDockWidgetList) Erase_1() {
	// QDockWidgetList_erase_1()
	rv, err := qtrt.InvokeQtFunc6("QDockWidgetList_erase_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// int count()
func (this *QDockWidgetList) Count_1() int {
	// QDockWidgetList_count_1()
	rv, err := qtrt.InvokeQtFunc6("QDockWidgetList_count_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// int length()
func (this *QDockWidgetList) Length_0() int {
	// QDockWidgetList_length_0()
	rv, err := qtrt.InvokeQtFunc6("QDockWidgetList_length_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// T & first()
func (this *QDockWidgetList) First_0() *QDockWidget {
	// QDockWidgetList_first_0()
	rv, err := qtrt.InvokeQtFunc6("QDockWidgetList_first_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QDockWidget{}
}

// const T & constFirst()
func (this *QDockWidgetList) ConstFirst_0() *QDockWidget {
	// QDockWidgetList_constFirst_0()
	rv, err := qtrt.InvokeQtFunc6("QDockWidgetList_constFirst_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QDockWidget{}
}

// const T & first()
func (this *QDockWidgetList) First_1() *QDockWidget {
	// QDockWidgetList_first_1()
	rv, err := qtrt.InvokeQtFunc6("QDockWidgetList_first_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QDockWidget{}
}

// T & last()
func (this *QDockWidgetList) Last_0() *QDockWidget {
	// QDockWidgetList_last_0()
	rv, err := qtrt.InvokeQtFunc6("QDockWidgetList_last_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QDockWidget{}
}

// const T & last()
func (this *QDockWidgetList) Last_1() *QDockWidget {
	// QDockWidgetList_last_1()
	rv, err := qtrt.InvokeQtFunc6("QDockWidgetList_last_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QDockWidget{}
}

// const T & constLast()
func (this *QDockWidgetList) ConstLast_0() *QDockWidget {
	// QDockWidgetList_constLast_0()
	rv, err := qtrt.InvokeQtFunc6("QDockWidgetList_constLast_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QDockWidget{}
}

// void removeFirst()
func (this *QDockWidgetList) RemoveFirst_0() {
	// QDockWidgetList_removeFirst_0()
	rv, err := qtrt.InvokeQtFunc6("QDockWidgetList_removeFirst_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void removeLast()
func (this *QDockWidgetList) RemoveLast_0() {
	// QDockWidgetList_removeLast_0()
	rv, err := qtrt.InvokeQtFunc6("QDockWidgetList_removeLast_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// bool startsWith(const T &)
func (this *QDockWidgetList) StartsWith_0() bool {
	// QDockWidgetList_startsWith_0()
	rv, err := qtrt.InvokeQtFunc6("QDockWidgetList_startsWith_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// bool endsWith(const T &)
func (this *QDockWidgetList) EndsWith_0() bool {
	// QDockWidgetList_endsWith_0()
	rv, err := qtrt.InvokeQtFunc6("QDockWidgetList_endsWith_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// QList<T> mid(int, int)
func (this *QDockWidgetList) Mid_0() *QDockWidgetList {
	// QDockWidgetList_mid_0()
	rv, err := qtrt.InvokeQtFunc6("QDockWidgetList_mid_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// T value(int)
func (this *QDockWidgetList) Value_0() *QDockWidget {
	// QDockWidgetList_value_0()
	rv, err := qtrt.InvokeQtFunc6("QDockWidgetList_value_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QDockWidget{}
}

// T value(int, const T &)
func (this *QDockWidgetList) Value_1() *QDockWidget {
	// QDockWidgetList_value_1()
	rv, err := qtrt.InvokeQtFunc6("QDockWidgetList_value_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QDockWidget{}
}

// void push_back(const T &)
func (this *QDockWidgetList) Push_back_0() {
	// QDockWidgetList_push_back_0()
	rv, err := qtrt.InvokeQtFunc6("QDockWidgetList_push_back_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void push_front(const T &)
func (this *QDockWidgetList) Push_front_0() {
	// QDockWidgetList_push_front_0()
	rv, err := qtrt.InvokeQtFunc6("QDockWidgetList_push_front_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// T & front()
func (this *QDockWidgetList) Front_0() *QDockWidget {
	// QDockWidgetList_front_0()
	rv, err := qtrt.InvokeQtFunc6("QDockWidgetList_front_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QDockWidget{}
}

// const T & front()
func (this *QDockWidgetList) Front_1() *QDockWidget {
	// QDockWidgetList_front_1()
	rv, err := qtrt.InvokeQtFunc6("QDockWidgetList_front_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QDockWidget{}
}

// T & back()
func (this *QDockWidgetList) Back_0() *QDockWidget {
	// QDockWidgetList_back_0()
	rv, err := qtrt.InvokeQtFunc6("QDockWidgetList_back_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QDockWidget{}
}

// const T & back()
func (this *QDockWidgetList) Back_1() *QDockWidget {
	// QDockWidgetList_back_1()
	rv, err := qtrt.InvokeQtFunc6("QDockWidgetList_back_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QDockWidget{}
}

// void pop_front()
func (this *QDockWidgetList) Pop_front_0() {
	// QDockWidgetList_pop_front_0()
	rv, err := qtrt.InvokeQtFunc6("QDockWidgetList_pop_front_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void pop_back()
func (this *QDockWidgetList) Pop_back_0() {
	// QDockWidgetList_pop_back_0()
	rv, err := qtrt.InvokeQtFunc6("QDockWidgetList_pop_back_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// bool empty()
func (this *QDockWidgetList) Empty_0() bool {
	// QDockWidgetList_empty_0()
	rv, err := qtrt.InvokeQtFunc6("QDockWidgetList_empty_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// QList<T> & operator+=(const QList<T> &)
func (this *QDockWidgetList) Operator_add_equal_0() *QDockWidgetList {
	// QDockWidgetList_operator_add_equal_0()
	rv, err := qtrt.InvokeQtFunc6("QDockWidgetList_operator_add_equal_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QList<T> operator+(const QList<T> &)
func (this *QDockWidgetList) Operator_add_0() *QDockWidgetList {
	// QDockWidgetList_operator_add_0()
	rv, err := qtrt.InvokeQtFunc6("QDockWidgetList_operator_add_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QList<T> & operator+=(const T &)
func (this *QDockWidgetList) Operator_add_equal_1() *QDockWidgetList {
	// QDockWidgetList_operator_add_equal_1()
	rv, err := qtrt.InvokeQtFunc6("QDockWidgetList_operator_add_equal_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QList<T> & operator<<(const T &)
func (this *QDockWidgetList) Operator_left_shift_0() *QDockWidgetList {
	// QDockWidgetList_operator_left_shift_0()
	rv, err := qtrt.InvokeQtFunc6("QDockWidgetList_operator_left_shift_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QList<T> & operator<<(const QList<T> &)
func (this *QDockWidgetList) Operator_left_shift_1() *QDockWidgetList {
	// QDockWidgetList_operator_left_shift_1()
	rv, err := qtrt.InvokeQtFunc6("QDockWidgetList_operator_left_shift_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QVector<T> toVector()
func (this *QDockWidgetList) ToVector_0() {
	// QDockWidgetList_toVector_0()
	rv, err := qtrt.InvokeQtFunc6("QDockWidgetList_toVector_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QSet<T> toSet()
func (this *QDockWidgetList) ToSet_0() {
	// QDockWidgetList_toSet_0()
	rv, err := qtrt.InvokeQtFunc6("QDockWidgetList_toSet_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList<T> fromVector(const QVector<T> &)
func (this *QDockWidgetList) FromVector_0() *QDockWidgetList {
	// QDockWidgetList_fromVector_0()
	rv, err := qtrt.InvokeQtFunc6("QDockWidgetList_fromVector_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QList<T> fromSet(const QSet<T> &)
func (this *QDockWidgetList) FromSet_0() *QDockWidgetList {
	// QDockWidgetList_fromSet_0()
	rv, err := qtrt.InvokeQtFunc6("QDockWidgetList_fromSet_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QList<T> fromStdList(const std::list<T> &)
func (this *QDockWidgetList) FromStdList_0() *QDockWidgetList {
	// QDockWidgetList_fromStdList_0()
	rv, err := qtrt.InvokeQtFunc6("QDockWidgetList_fromStdList_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// std::list<T> toStdList()
func (this *QDockWidgetList) ToStdList_0() {
	// QDockWidgetList_toStdList_0()
	rv, err := qtrt.InvokeQtFunc6("QDockWidgetList_toStdList_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::Node * detach_helper_grow(int, int)
func (this *QDockWidgetList) Detach_helper_grow_0() {
	// QDockWidgetList_detach_helper_grow_0()
	rv, err := qtrt.InvokeQtFunc6("QDockWidgetList_detach_helper_grow_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void detach_helper(int)
func (this *QDockWidgetList) Detach_helper_0() {
	// QDockWidgetList_detach_helper_0()
	rv, err := qtrt.InvokeQtFunc6("QDockWidgetList_detach_helper_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void detach_helper()
func (this *QDockWidgetList) Detach_helper_1() {
	// QDockWidgetList_detach_helper_1()
	rv, err := qtrt.InvokeQtFunc6("QDockWidgetList_detach_helper_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void dealloc(QListData::Data *)
func (this *QDockWidgetList) Dealloc_0() {
	// QDockWidgetList_dealloc_0()
	rv, err := qtrt.InvokeQtFunc6("QDockWidgetList_dealloc_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void node_construct(QList::Node *, const T &)
func (this *QDockWidgetList) Node_construct_0() {
	// QDockWidgetList_node_construct_0()
	rv, err := qtrt.InvokeQtFunc6("QDockWidgetList_node_construct_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void node_destruct(QList::Node *)
func (this *QDockWidgetList) Node_destruct_0() {
	// QDockWidgetList_node_destruct_0()
	rv, err := qtrt.InvokeQtFunc6("QDockWidgetList_node_destruct_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void node_copy(QList::Node *, QList::Node *, QList::Node *)
func (this *QDockWidgetList) Node_copy_0() {
	// QDockWidgetList_node_copy_0()
	rv, err := qtrt.InvokeQtFunc6("QDockWidgetList_node_copy_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void node_destruct(QList::Node *, QList::Node *)
func (this *QDockWidgetList) Node_destruct_1() {
	// QDockWidgetList_node_destruct_1()
	rv, err := qtrt.InvokeQtFunc6("QDockWidgetList_node_destruct_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// bool isValidIterator(const QList::iterator &)
func (this *QDockWidgetList) IsValidIterator_0() bool {
	// QDockWidgetList_isValidIterator_0()
	rv, err := qtrt.InvokeQtFunc6("QDockWidgetList_isValidIterator_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// bool op_eq_impl(const QList<T> &, QListData::NotArrayCompatibleLayout)
func (this *QDockWidgetList) Op_eq_impl_0() bool {
	// QDockWidgetList_op_eq_impl_0()
	rv, err := qtrt.InvokeQtFunc6("QDockWidgetList_op_eq_impl_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// bool op_eq_impl(const QList<T> &, QListData::ArrayCompatibleLayout)
func (this *QDockWidgetList) Op_eq_impl_1() bool {
	// QDockWidgetList_op_eq_impl_1()
	rv, err := qtrt.InvokeQtFunc6("QDockWidgetList_op_eq_impl_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// bool contains_impl(const T &, QListData::NotArrayCompatibleLayout)
func (this *QDockWidgetList) Contains_impl_0() bool {
	// QDockWidgetList_contains_impl_0()
	rv, err := qtrt.InvokeQtFunc6("QDockWidgetList_contains_impl_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// bool contains_impl(const T &, QListData::ArrayCompatibleLayout)
func (this *QDockWidgetList) Contains_impl_1() bool {
	// QDockWidgetList_contains_impl_1()
	rv, err := qtrt.InvokeQtFunc6("QDockWidgetList_contains_impl_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// int count_impl(const T &, QListData::NotArrayCompatibleLayout)
func (this *QDockWidgetList) Count_impl_0() int {
	// QDockWidgetList_count_impl_0()
	rv, err := qtrt.InvokeQtFunc6("QDockWidgetList_count_impl_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// int count_impl(const T &, QListData::ArrayCompatibleLayout)
func (this *QDockWidgetList) Count_impl_1() int {
	// QDockWidgetList_count_impl_1()
	rv, err := qtrt.InvokeQtFunc6("QDockWidgetList_count_impl_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

//  body block end
