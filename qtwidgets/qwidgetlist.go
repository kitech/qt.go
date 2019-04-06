package qtwidgets

// /usr/include/qt/QtWidgets/qwidget.h
// #include <qwidget.h>
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

func init_unused_10055() {
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
type QWidgetList struct {
	*qtrt.CObject
}

// QList<T> & operator=(const QList<T> &)
func (this *QWidgetList) Operator_equal0() *QWidgetList {
	// QWidgetList_operator_equal_0()
	rv, err := qtrt.InvokeQtFunc6("C_QWidgetList_operator_equal_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QList<T> & operator=(QList<T> &&)
func (this *QWidgetList) Operator_equal1() *QWidgetList {
	// QWidgetList_operator_equal_1()
	rv, err := qtrt.InvokeQtFunc6("C_QWidgetList_operator_equal_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// void swap(QList<T> &)
func (this *QWidgetList) Swap0() {
	// QWidgetList_swap_0()
	rv, err := qtrt.InvokeQtFunc6("C_QWidgetList_swap_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// bool operator==(const QList<T> &)
func (this *QWidgetList) Operator_equal_equal0() bool {
	// QWidgetList_operator_equal_equal_0()
	rv, err := qtrt.InvokeQtFunc6("C_QWidgetList_operator_equal_equal_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// bool operator!=(const QList<T> &)
func (this *QWidgetList) Operator_not_equal0() bool {
	// QWidgetList_operator_not_equal_0()
	rv, err := qtrt.InvokeQtFunc6("C_QWidgetList_operator_not_equal_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// int size()
func (this *QWidgetList) Size0() int {
	// QWidgetList_size_0()
	rv, err := qtrt.InvokeQtFunc6("C_QWidgetList_size_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// void detach()
func (this *QWidgetList) Detach0() {
	// QWidgetList_detach_0()
	rv, err := qtrt.InvokeQtFunc6("C_QWidgetList_detach_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void detachShared()
func (this *QWidgetList) DetachShared0() {
	// QWidgetList_detachShared_0()
	rv, err := qtrt.InvokeQtFunc6("C_QWidgetList_detachShared_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// bool isDetached()
func (this *QWidgetList) IsDetached0() bool {
	// QWidgetList_isDetached_0()
	rv, err := qtrt.InvokeQtFunc6("C_QWidgetList_isDetached_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// void setSharable(bool)
func (this *QWidgetList) SetSharable0() {
	// QWidgetList_setSharable_0()
	rv, err := qtrt.InvokeQtFunc6("C_QWidgetList_setSharable_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// bool isSharedWith(const QList<T> &)
func (this *QWidgetList) IsSharedWith0() bool {
	// QWidgetList_isSharedWith_0()
	rv, err := qtrt.InvokeQtFunc6("C_QWidgetList_isSharedWith_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// bool isEmpty()
func (this *QWidgetList) IsEmpty0() bool {
	// QWidgetList_isEmpty_0()
	rv, err := qtrt.InvokeQtFunc6("C_QWidgetList_isEmpty_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// void clear()
func (this *QWidgetList) Clear0() {
	// QWidgetList_clear_0()
	rv, err := qtrt.InvokeQtFunc6("C_QWidgetList_clear_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// const T & at(int)
func (this *QWidgetList) At0() *QWidget {
	// QWidgetList_at_0()
	rv, err := qtrt.InvokeQtFunc6("C_QWidgetList_at_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QWidget{}
}

// const T & operator[](int)
func (this *QWidgetList) Operator_get_index0() *QWidget {
	// QWidgetList_operator_get_index_0()
	rv, err := qtrt.InvokeQtFunc6("C_QWidgetList_operator_get_index_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QWidget{}
}

// T & operator[](int)
func (this *QWidgetList) Operator_get_index1() *QWidget {
	// QWidgetList_operator_get_index_1()
	rv, err := qtrt.InvokeQtFunc6("C_QWidgetList_operator_get_index_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QWidget{}
}

// void reserve(int)
func (this *QWidgetList) Reserve0() {
	// QWidgetList_reserve_0()
	rv, err := qtrt.InvokeQtFunc6("C_QWidgetList_reserve_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void append(const T &)
func (this *QWidgetList) Append0() {
	// QWidgetList_append_0()
	rv, err := qtrt.InvokeQtFunc6("C_QWidgetList_append_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void append(const QList<T> &)
func (this *QWidgetList) Append1() {
	// QWidgetList_append_1()
	rv, err := qtrt.InvokeQtFunc6("C_QWidgetList_append_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void prepend(const T &)
func (this *QWidgetList) Prepend0() {
	// QWidgetList_prepend_0()
	rv, err := qtrt.InvokeQtFunc6("C_QWidgetList_prepend_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void insert(int, const T &)
func (this *QWidgetList) Insert0() {
	// QWidgetList_insert_0()
	rv, err := qtrt.InvokeQtFunc6("C_QWidgetList_insert_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void replace(int, const T &)
func (this *QWidgetList) Replace0() {
	// QWidgetList_replace_0()
	rv, err := qtrt.InvokeQtFunc6("C_QWidgetList_replace_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void removeAt(int)
func (this *QWidgetList) RemoveAt0() {
	// QWidgetList_removeAt_0()
	rv, err := qtrt.InvokeQtFunc6("C_QWidgetList_removeAt_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// int removeAll(const T &)
func (this *QWidgetList) RemoveAll0() int {
	// QWidgetList_removeAll_0()
	rv, err := qtrt.InvokeQtFunc6("C_QWidgetList_removeAll_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// bool removeOne(const T &)
func (this *QWidgetList) RemoveOne0() bool {
	// QWidgetList_removeOne_0()
	rv, err := qtrt.InvokeQtFunc6("C_QWidgetList_removeOne_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// T takeAt(int)
func (this *QWidgetList) TakeAt0() *QWidget {
	// QWidgetList_takeAt_0()
	rv, err := qtrt.InvokeQtFunc6("C_QWidgetList_takeAt_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QWidget{}
}

// T takeFirst()
func (this *QWidgetList) TakeFirst0() *QWidget {
	// QWidgetList_takeFirst_0()
	rv, err := qtrt.InvokeQtFunc6("C_QWidgetList_takeFirst_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QWidget{}
}

// T takeLast()
func (this *QWidgetList) TakeLast0() *QWidget {
	// QWidgetList_takeLast_0()
	rv, err := qtrt.InvokeQtFunc6("C_QWidgetList_takeLast_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QWidget{}
}

// void move(int, int)
func (this *QWidgetList) Move0() {
	// QWidgetList_move_0()
	rv, err := qtrt.InvokeQtFunc6("C_QWidgetList_move_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void swap(int, int)
func (this *QWidgetList) Swap1() {
	// QWidgetList_swap_1()
	rv, err := qtrt.InvokeQtFunc6("C_QWidgetList_swap_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// int indexOf(const T &, int)
func (this *QWidgetList) IndexOf0() int {
	// QWidgetList_indexOf_0()
	rv, err := qtrt.InvokeQtFunc6("C_QWidgetList_indexOf_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// int lastIndexOf(const T &, int)
func (this *QWidgetList) LastIndexOf0() int {
	// QWidgetList_lastIndexOf_0()
	rv, err := qtrt.InvokeQtFunc6("C_QWidgetList_lastIndexOf_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// bool contains(const T &)
func (this *QWidgetList) Contains0() bool {
	// QWidgetList_contains_0()
	rv, err := qtrt.InvokeQtFunc6("C_QWidgetList_contains_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// int count(const T &)
func (this *QWidgetList) Count0() int {
	// QWidgetList_count_0()
	rv, err := qtrt.InvokeQtFunc6("C_QWidgetList_count_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// QList::iterator begin()
func (this *QWidgetList) Begin0() {
	// QWidgetList_begin_0()
	rv, err := qtrt.InvokeQtFunc6("C_QWidgetList_begin_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_iterator begin()
func (this *QWidgetList) Begin1() {
	// QWidgetList_begin_1()
	rv, err := qtrt.InvokeQtFunc6("C_QWidgetList_begin_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_iterator cbegin()
func (this *QWidgetList) Cbegin0() {
	// QWidgetList_cbegin_0()
	rv, err := qtrt.InvokeQtFunc6("C_QWidgetList_cbegin_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_iterator constBegin()
func (this *QWidgetList) ConstBegin0() {
	// QWidgetList_constBegin_0()
	rv, err := qtrt.InvokeQtFunc6("C_QWidgetList_constBegin_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::iterator end()
func (this *QWidgetList) End0() {
	// QWidgetList_end_0()
	rv, err := qtrt.InvokeQtFunc6("C_QWidgetList_end_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_iterator end()
func (this *QWidgetList) End1() {
	// QWidgetList_end_1()
	rv, err := qtrt.InvokeQtFunc6("C_QWidgetList_end_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_iterator cend()
func (this *QWidgetList) Cend0() {
	// QWidgetList_cend_0()
	rv, err := qtrt.InvokeQtFunc6("C_QWidgetList_cend_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_iterator constEnd()
func (this *QWidgetList) ConstEnd0() {
	// QWidgetList_constEnd_0()
	rv, err := qtrt.InvokeQtFunc6("C_QWidgetList_constEnd_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::reverse_iterator rbegin()
func (this *QWidgetList) Rbegin0() {
	// QWidgetList_rbegin_0()
	rv, err := qtrt.InvokeQtFunc6("C_QWidgetList_rbegin_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::reverse_iterator rend()
func (this *QWidgetList) Rend0() {
	// QWidgetList_rend_0()
	rv, err := qtrt.InvokeQtFunc6("C_QWidgetList_rend_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_reverse_iterator rbegin()
func (this *QWidgetList) Rbegin1() {
	// QWidgetList_rbegin_1()
	rv, err := qtrt.InvokeQtFunc6("C_QWidgetList_rbegin_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_reverse_iterator rend()
func (this *QWidgetList) Rend1() {
	// QWidgetList_rend_1()
	rv, err := qtrt.InvokeQtFunc6("C_QWidgetList_rend_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_reverse_iterator crbegin()
func (this *QWidgetList) Crbegin0() {
	// QWidgetList_crbegin_0()
	rv, err := qtrt.InvokeQtFunc6("C_QWidgetList_crbegin_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_reverse_iterator crend()
func (this *QWidgetList) Crend0() {
	// QWidgetList_crend_0()
	rv, err := qtrt.InvokeQtFunc6("C_QWidgetList_crend_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::iterator insert(QList::iterator, const T &)
func (this *QWidgetList) Insert1() {
	// QWidgetList_insert_1()
	rv, err := qtrt.InvokeQtFunc6("C_QWidgetList_insert_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::iterator erase(QList::iterator)
func (this *QWidgetList) Erase0() {
	// QWidgetList_erase_0()
	rv, err := qtrt.InvokeQtFunc6("C_QWidgetList_erase_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::iterator erase(QList::iterator, QList::iterator)
func (this *QWidgetList) Erase1() {
	// QWidgetList_erase_1()
	rv, err := qtrt.InvokeQtFunc6("C_QWidgetList_erase_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// int count()
func (this *QWidgetList) Count1() int {
	// QWidgetList_count_1()
	rv, err := qtrt.InvokeQtFunc6("C_QWidgetList_count_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// int length()
func (this *QWidgetList) Length0() int {
	// QWidgetList_length_0()
	rv, err := qtrt.InvokeQtFunc6("C_QWidgetList_length_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// T & first()
func (this *QWidgetList) First0() *QWidget {
	// QWidgetList_first_0()
	rv, err := qtrt.InvokeQtFunc6("C_QWidgetList_first_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QWidget{}
}

// const T & constFirst()
func (this *QWidgetList) ConstFirst0() *QWidget {
	// QWidgetList_constFirst_0()
	rv, err := qtrt.InvokeQtFunc6("C_QWidgetList_constFirst_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QWidget{}
}

// const T & first()
func (this *QWidgetList) First1() *QWidget {
	// QWidgetList_first_1()
	rv, err := qtrt.InvokeQtFunc6("C_QWidgetList_first_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QWidget{}
}

// T & last()
func (this *QWidgetList) Last0() *QWidget {
	// QWidgetList_last_0()
	rv, err := qtrt.InvokeQtFunc6("C_QWidgetList_last_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QWidget{}
}

// const T & last()
func (this *QWidgetList) Last1() *QWidget {
	// QWidgetList_last_1()
	rv, err := qtrt.InvokeQtFunc6("C_QWidgetList_last_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QWidget{}
}

// const T & constLast()
func (this *QWidgetList) ConstLast0() *QWidget {
	// QWidgetList_constLast_0()
	rv, err := qtrt.InvokeQtFunc6("C_QWidgetList_constLast_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QWidget{}
}

// void removeFirst()
func (this *QWidgetList) RemoveFirst0() {
	// QWidgetList_removeFirst_0()
	rv, err := qtrt.InvokeQtFunc6("C_QWidgetList_removeFirst_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void removeLast()
func (this *QWidgetList) RemoveLast0() {
	// QWidgetList_removeLast_0()
	rv, err := qtrt.InvokeQtFunc6("C_QWidgetList_removeLast_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// bool startsWith(const T &)
func (this *QWidgetList) StartsWith0() bool {
	// QWidgetList_startsWith_0()
	rv, err := qtrt.InvokeQtFunc6("C_QWidgetList_startsWith_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// bool endsWith(const T &)
func (this *QWidgetList) EndsWith0() bool {
	// QWidgetList_endsWith_0()
	rv, err := qtrt.InvokeQtFunc6("C_QWidgetList_endsWith_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// QList<T> mid(int, int)
func (this *QWidgetList) Mid0() *QWidgetList {
	// QWidgetList_mid_0()
	rv, err := qtrt.InvokeQtFunc6("C_QWidgetList_mid_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// T value(int)
func (this *QWidgetList) Value0() *QWidget {
	// QWidgetList_value_0()
	rv, err := qtrt.InvokeQtFunc6("C_QWidgetList_value_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QWidget{}
}

// T value(int, const T &)
func (this *QWidgetList) Value1() *QWidget {
	// QWidgetList_value_1()
	rv, err := qtrt.InvokeQtFunc6("C_QWidgetList_value_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QWidget{}
}

// void push_back(const T &)
func (this *QWidgetList) Push_back0() {
	// QWidgetList_push_back_0()
	rv, err := qtrt.InvokeQtFunc6("C_QWidgetList_push_back_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void push_front(const T &)
func (this *QWidgetList) Push_front0() {
	// QWidgetList_push_front_0()
	rv, err := qtrt.InvokeQtFunc6("C_QWidgetList_push_front_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// T & front()
func (this *QWidgetList) Front0() *QWidget {
	// QWidgetList_front_0()
	rv, err := qtrt.InvokeQtFunc6("C_QWidgetList_front_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QWidget{}
}

// const T & front()
func (this *QWidgetList) Front1() *QWidget {
	// QWidgetList_front_1()
	rv, err := qtrt.InvokeQtFunc6("C_QWidgetList_front_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QWidget{}
}

// T & back()
func (this *QWidgetList) Back0() *QWidget {
	// QWidgetList_back_0()
	rv, err := qtrt.InvokeQtFunc6("C_QWidgetList_back_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QWidget{}
}

// const T & back()
func (this *QWidgetList) Back1() *QWidget {
	// QWidgetList_back_1()
	rv, err := qtrt.InvokeQtFunc6("C_QWidgetList_back_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QWidget{}
}

// void pop_front()
func (this *QWidgetList) Pop_front0() {
	// QWidgetList_pop_front_0()
	rv, err := qtrt.InvokeQtFunc6("C_QWidgetList_pop_front_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void pop_back()
func (this *QWidgetList) Pop_back0() {
	// QWidgetList_pop_back_0()
	rv, err := qtrt.InvokeQtFunc6("C_QWidgetList_pop_back_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// bool empty()
func (this *QWidgetList) Empty0() bool {
	// QWidgetList_empty_0()
	rv, err := qtrt.InvokeQtFunc6("C_QWidgetList_empty_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// QList<T> & operator+=(const QList<T> &)
func (this *QWidgetList) Operator_add_equal0() *QWidgetList {
	// QWidgetList_operator_add_equal_0()
	rv, err := qtrt.InvokeQtFunc6("C_QWidgetList_operator_add_equal_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QList<T> operator+(const QList<T> &)
func (this *QWidgetList) Operator_add0() *QWidgetList {
	// QWidgetList_operator_add_0()
	rv, err := qtrt.InvokeQtFunc6("C_QWidgetList_operator_add_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QList<T> & operator+=(const T &)
func (this *QWidgetList) Operator_add_equal1() *QWidgetList {
	// QWidgetList_operator_add_equal_1()
	rv, err := qtrt.InvokeQtFunc6("C_QWidgetList_operator_add_equal_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QList<T> & operator<<(const T &)
func (this *QWidgetList) Operator_left_shift0() *QWidgetList {
	// QWidgetList_operator_left_shift_0()
	rv, err := qtrt.InvokeQtFunc6("C_QWidgetList_operator_left_shift_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QList<T> & operator<<(const QList<T> &)
func (this *QWidgetList) Operator_left_shift1() *QWidgetList {
	// QWidgetList_operator_left_shift_1()
	rv, err := qtrt.InvokeQtFunc6("C_QWidgetList_operator_left_shift_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QVector<T> toVector()
func (this *QWidgetList) ToVector0() {
	// QWidgetList_toVector_0()
	rv, err := qtrt.InvokeQtFunc6("C_QWidgetList_toVector_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QSet<T> toSet()
func (this *QWidgetList) ToSet0() {
	// QWidgetList_toSet_0()
	rv, err := qtrt.InvokeQtFunc6("C_QWidgetList_toSet_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList<T> fromVector(const QVector<T> &)
func (this *QWidgetList) FromVector0() *QWidgetList {
	// QWidgetList_fromVector_0()
	rv, err := qtrt.InvokeQtFunc6("C_QWidgetList_fromVector_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QList<T> fromSet(const QSet<T> &)
func (this *QWidgetList) FromSet0() *QWidgetList {
	// QWidgetList_fromSet_0()
	rv, err := qtrt.InvokeQtFunc6("C_QWidgetList_fromSet_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QList<T> fromStdList(const std::list<T> &)
func (this *QWidgetList) FromStdList0() *QWidgetList {
	// QWidgetList_fromStdList_0()
	rv, err := qtrt.InvokeQtFunc6("C_QWidgetList_fromStdList_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// std::list<T> toStdList()
func (this *QWidgetList) ToStdList0() {
	// QWidgetList_toStdList_0()
	rv, err := qtrt.InvokeQtFunc6("C_QWidgetList_toStdList_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::Node * detach_helper_grow(int, int)
func (this *QWidgetList) Detach_helper_grow0() {
	// QWidgetList_detach_helper_grow_0()
	rv, err := qtrt.InvokeQtFunc6("C_QWidgetList_detach_helper_grow_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void detach_helper(int)
func (this *QWidgetList) Detach_helper0() {
	// QWidgetList_detach_helper_0()
	rv, err := qtrt.InvokeQtFunc6("C_QWidgetList_detach_helper_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void detach_helper()
func (this *QWidgetList) Detach_helper1() {
	// QWidgetList_detach_helper_1()
	rv, err := qtrt.InvokeQtFunc6("C_QWidgetList_detach_helper_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void dealloc(QListData::Data *)
func (this *QWidgetList) Dealloc0() {
	// QWidgetList_dealloc_0()
	rv, err := qtrt.InvokeQtFunc6("C_QWidgetList_dealloc_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void node_construct(QList::Node *, const T &)
func (this *QWidgetList) Node_construct0() {
	// QWidgetList_node_construct_0()
	rv, err := qtrt.InvokeQtFunc6("C_QWidgetList_node_construct_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void node_destruct(QList::Node *)
func (this *QWidgetList) Node_destruct0() {
	// QWidgetList_node_destruct_0()
	rv, err := qtrt.InvokeQtFunc6("C_QWidgetList_node_destruct_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void node_copy(QList::Node *, QList::Node *, QList::Node *)
func (this *QWidgetList) Node_copy0() {
	// QWidgetList_node_copy_0()
	rv, err := qtrt.InvokeQtFunc6("C_QWidgetList_node_copy_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void node_destruct(QList::Node *, QList::Node *)
func (this *QWidgetList) Node_destruct1() {
	// QWidgetList_node_destruct_1()
	rv, err := qtrt.InvokeQtFunc6("C_QWidgetList_node_destruct_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// bool isValidIterator(const QList::iterator &)
func (this *QWidgetList) IsValidIterator0() bool {
	// QWidgetList_isValidIterator_0()
	rv, err := qtrt.InvokeQtFunc6("C_QWidgetList_isValidIterator_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// bool op_eq_impl(const QList<T> &, QListData::NotArrayCompatibleLayout)
func (this *QWidgetList) Op_eq_impl0() bool {
	// QWidgetList_op_eq_impl_0()
	rv, err := qtrt.InvokeQtFunc6("C_QWidgetList_op_eq_impl_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// bool op_eq_impl(const QList<T> &, QListData::ArrayCompatibleLayout)
func (this *QWidgetList) Op_eq_impl1() bool {
	// QWidgetList_op_eq_impl_1()
	rv, err := qtrt.InvokeQtFunc6("C_QWidgetList_op_eq_impl_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// bool contains_impl(const T &, QListData::NotArrayCompatibleLayout)
func (this *QWidgetList) Contains_impl0() bool {
	// QWidgetList_contains_impl_0()
	rv, err := qtrt.InvokeQtFunc6("C_QWidgetList_contains_impl_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// bool contains_impl(const T &, QListData::ArrayCompatibleLayout)
func (this *QWidgetList) Contains_impl1() bool {
	// QWidgetList_contains_impl_1()
	rv, err := qtrt.InvokeQtFunc6("C_QWidgetList_contains_impl_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// int count_impl(const T &, QListData::NotArrayCompatibleLayout)
func (this *QWidgetList) Count_impl0() int {
	// QWidgetList_count_impl_0()
	rv, err := qtrt.InvokeQtFunc6("C_QWidgetList_count_impl_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// int count_impl(const T &, QListData::ArrayCompatibleLayout)
func (this *QWidgetList) Count_impl1() int {
	// QWidgetList_count_impl_1()
	rv, err := qtrt.InvokeQtFunc6("C_QWidgetList_count_impl_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

//  body block end
