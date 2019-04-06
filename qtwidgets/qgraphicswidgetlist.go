// +build !minimal

package qtwidgets

// /usr/include/qt/QtWidgets/qgraphicswidget.h
// #include <qgraphicswidget.h>
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

func init_unused_10123() {
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
type QGraphicsWidgetList struct {
	*qtrt.CObject
}

// QList<T> & operator=(const QList<T> &)
func (this *QGraphicsWidgetList) Operator_equal0() *QGraphicsWidgetList {
	// QGraphicsWidgetList_operator_equal_0()
	rv, err := qtrt.InvokeQtFunc6("C_QGraphicsWidgetList_operator_equal_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QList<T> & operator=(QList<T> &&)
func (this *QGraphicsWidgetList) Operator_equal1() *QGraphicsWidgetList {
	// QGraphicsWidgetList_operator_equal_1()
	rv, err := qtrt.InvokeQtFunc6("C_QGraphicsWidgetList_operator_equal_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// void swap(QList<T> &)
func (this *QGraphicsWidgetList) Swap0() {
	// QGraphicsWidgetList_swap_0()
	rv, err := qtrt.InvokeQtFunc6("C_QGraphicsWidgetList_swap_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// bool operator==(const QList<T> &)
func (this *QGraphicsWidgetList) Operator_equal_equal0() bool {
	// QGraphicsWidgetList_operator_equal_equal_0()
	rv, err := qtrt.InvokeQtFunc6("C_QGraphicsWidgetList_operator_equal_equal_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// bool operator!=(const QList<T> &)
func (this *QGraphicsWidgetList) Operator_not_equal0() bool {
	// QGraphicsWidgetList_operator_not_equal_0()
	rv, err := qtrt.InvokeQtFunc6("C_QGraphicsWidgetList_operator_not_equal_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// int size()
func (this *QGraphicsWidgetList) Size0() int {
	// QGraphicsWidgetList_size_0()
	rv, err := qtrt.InvokeQtFunc6("C_QGraphicsWidgetList_size_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// void detach()
func (this *QGraphicsWidgetList) Detach0() {
	// QGraphicsWidgetList_detach_0()
	rv, err := qtrt.InvokeQtFunc6("C_QGraphicsWidgetList_detach_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void detachShared()
func (this *QGraphicsWidgetList) DetachShared0() {
	// QGraphicsWidgetList_detachShared_0()
	rv, err := qtrt.InvokeQtFunc6("C_QGraphicsWidgetList_detachShared_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// bool isDetached()
func (this *QGraphicsWidgetList) IsDetached0() bool {
	// QGraphicsWidgetList_isDetached_0()
	rv, err := qtrt.InvokeQtFunc6("C_QGraphicsWidgetList_isDetached_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// void setSharable(bool)
func (this *QGraphicsWidgetList) SetSharable0() {
	// QGraphicsWidgetList_setSharable_0()
	rv, err := qtrt.InvokeQtFunc6("C_QGraphicsWidgetList_setSharable_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// bool isSharedWith(const QList<T> &)
func (this *QGraphicsWidgetList) IsSharedWith0() bool {
	// QGraphicsWidgetList_isSharedWith_0()
	rv, err := qtrt.InvokeQtFunc6("C_QGraphicsWidgetList_isSharedWith_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// bool isEmpty()
func (this *QGraphicsWidgetList) IsEmpty0() bool {
	// QGraphicsWidgetList_isEmpty_0()
	rv, err := qtrt.InvokeQtFunc6("C_QGraphicsWidgetList_isEmpty_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// void clear()
func (this *QGraphicsWidgetList) Clear0() {
	// QGraphicsWidgetList_clear_0()
	rv, err := qtrt.InvokeQtFunc6("C_QGraphicsWidgetList_clear_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// const T & at(int)
func (this *QGraphicsWidgetList) At0() *QGraphicsWidget {
	// QGraphicsWidgetList_at_0()
	rv, err := qtrt.InvokeQtFunc6("C_QGraphicsWidgetList_at_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QGraphicsWidget{}
}

// const T & operator[](int)
func (this *QGraphicsWidgetList) Operator_get_index0() *QGraphicsWidget {
	// QGraphicsWidgetList_operator_get_index_0()
	rv, err := qtrt.InvokeQtFunc6("C_QGraphicsWidgetList_operator_get_index_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QGraphicsWidget{}
}

// T & operator[](int)
func (this *QGraphicsWidgetList) Operator_get_index1() *QGraphicsWidget {
	// QGraphicsWidgetList_operator_get_index_1()
	rv, err := qtrt.InvokeQtFunc6("C_QGraphicsWidgetList_operator_get_index_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QGraphicsWidget{}
}

// void reserve(int)
func (this *QGraphicsWidgetList) Reserve0() {
	// QGraphicsWidgetList_reserve_0()
	rv, err := qtrt.InvokeQtFunc6("C_QGraphicsWidgetList_reserve_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void append(const T &)
func (this *QGraphicsWidgetList) Append0() {
	// QGraphicsWidgetList_append_0()
	rv, err := qtrt.InvokeQtFunc6("C_QGraphicsWidgetList_append_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void append(const QList<T> &)
func (this *QGraphicsWidgetList) Append1() {
	// QGraphicsWidgetList_append_1()
	rv, err := qtrt.InvokeQtFunc6("C_QGraphicsWidgetList_append_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void prepend(const T &)
func (this *QGraphicsWidgetList) Prepend0() {
	// QGraphicsWidgetList_prepend_0()
	rv, err := qtrt.InvokeQtFunc6("C_QGraphicsWidgetList_prepend_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void insert(int, const T &)
func (this *QGraphicsWidgetList) Insert0() {
	// QGraphicsWidgetList_insert_0()
	rv, err := qtrt.InvokeQtFunc6("C_QGraphicsWidgetList_insert_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void replace(int, const T &)
func (this *QGraphicsWidgetList) Replace0() {
	// QGraphicsWidgetList_replace_0()
	rv, err := qtrt.InvokeQtFunc6("C_QGraphicsWidgetList_replace_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void removeAt(int)
func (this *QGraphicsWidgetList) RemoveAt0() {
	// QGraphicsWidgetList_removeAt_0()
	rv, err := qtrt.InvokeQtFunc6("C_QGraphicsWidgetList_removeAt_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// int removeAll(const T &)
func (this *QGraphicsWidgetList) RemoveAll0() int {
	// QGraphicsWidgetList_removeAll_0()
	rv, err := qtrt.InvokeQtFunc6("C_QGraphicsWidgetList_removeAll_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// bool removeOne(const T &)
func (this *QGraphicsWidgetList) RemoveOne0() bool {
	// QGraphicsWidgetList_removeOne_0()
	rv, err := qtrt.InvokeQtFunc6("C_QGraphicsWidgetList_removeOne_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// T takeAt(int)
func (this *QGraphicsWidgetList) TakeAt0() *QGraphicsWidget {
	// QGraphicsWidgetList_takeAt_0()
	rv, err := qtrt.InvokeQtFunc6("C_QGraphicsWidgetList_takeAt_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QGraphicsWidget{}
}

// T takeFirst()
func (this *QGraphicsWidgetList) TakeFirst0() *QGraphicsWidget {
	// QGraphicsWidgetList_takeFirst_0()
	rv, err := qtrt.InvokeQtFunc6("C_QGraphicsWidgetList_takeFirst_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QGraphicsWidget{}
}

// T takeLast()
func (this *QGraphicsWidgetList) TakeLast0() *QGraphicsWidget {
	// QGraphicsWidgetList_takeLast_0()
	rv, err := qtrt.InvokeQtFunc6("C_QGraphicsWidgetList_takeLast_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QGraphicsWidget{}
}

// void move(int, int)
func (this *QGraphicsWidgetList) Move0() {
	// QGraphicsWidgetList_move_0()
	rv, err := qtrt.InvokeQtFunc6("C_QGraphicsWidgetList_move_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void swap(int, int)
func (this *QGraphicsWidgetList) Swap1() {
	// QGraphicsWidgetList_swap_1()
	rv, err := qtrt.InvokeQtFunc6("C_QGraphicsWidgetList_swap_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// int indexOf(const T &, int)
func (this *QGraphicsWidgetList) IndexOf0() int {
	// QGraphicsWidgetList_indexOf_0()
	rv, err := qtrt.InvokeQtFunc6("C_QGraphicsWidgetList_indexOf_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// int lastIndexOf(const T &, int)
func (this *QGraphicsWidgetList) LastIndexOf0() int {
	// QGraphicsWidgetList_lastIndexOf_0()
	rv, err := qtrt.InvokeQtFunc6("C_QGraphicsWidgetList_lastIndexOf_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// bool contains(const T &)
func (this *QGraphicsWidgetList) Contains0() bool {
	// QGraphicsWidgetList_contains_0()
	rv, err := qtrt.InvokeQtFunc6("C_QGraphicsWidgetList_contains_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// int count(const T &)
func (this *QGraphicsWidgetList) Count0() int {
	// QGraphicsWidgetList_count_0()
	rv, err := qtrt.InvokeQtFunc6("C_QGraphicsWidgetList_count_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// QList::iterator begin()
func (this *QGraphicsWidgetList) Begin0() {
	// QGraphicsWidgetList_begin_0()
	rv, err := qtrt.InvokeQtFunc6("C_QGraphicsWidgetList_begin_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_iterator begin()
func (this *QGraphicsWidgetList) Begin1() {
	// QGraphicsWidgetList_begin_1()
	rv, err := qtrt.InvokeQtFunc6("C_QGraphicsWidgetList_begin_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_iterator cbegin()
func (this *QGraphicsWidgetList) Cbegin0() {
	// QGraphicsWidgetList_cbegin_0()
	rv, err := qtrt.InvokeQtFunc6("C_QGraphicsWidgetList_cbegin_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_iterator constBegin()
func (this *QGraphicsWidgetList) ConstBegin0() {
	// QGraphicsWidgetList_constBegin_0()
	rv, err := qtrt.InvokeQtFunc6("C_QGraphicsWidgetList_constBegin_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::iterator end()
func (this *QGraphicsWidgetList) End0() {
	// QGraphicsWidgetList_end_0()
	rv, err := qtrt.InvokeQtFunc6("C_QGraphicsWidgetList_end_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_iterator end()
func (this *QGraphicsWidgetList) End1() {
	// QGraphicsWidgetList_end_1()
	rv, err := qtrt.InvokeQtFunc6("C_QGraphicsWidgetList_end_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_iterator cend()
func (this *QGraphicsWidgetList) Cend0() {
	// QGraphicsWidgetList_cend_0()
	rv, err := qtrt.InvokeQtFunc6("C_QGraphicsWidgetList_cend_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_iterator constEnd()
func (this *QGraphicsWidgetList) ConstEnd0() {
	// QGraphicsWidgetList_constEnd_0()
	rv, err := qtrt.InvokeQtFunc6("C_QGraphicsWidgetList_constEnd_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::reverse_iterator rbegin()
func (this *QGraphicsWidgetList) Rbegin0() {
	// QGraphicsWidgetList_rbegin_0()
	rv, err := qtrt.InvokeQtFunc6("C_QGraphicsWidgetList_rbegin_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::reverse_iterator rend()
func (this *QGraphicsWidgetList) Rend0() {
	// QGraphicsWidgetList_rend_0()
	rv, err := qtrt.InvokeQtFunc6("C_QGraphicsWidgetList_rend_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_reverse_iterator rbegin()
func (this *QGraphicsWidgetList) Rbegin1() {
	// QGraphicsWidgetList_rbegin_1()
	rv, err := qtrt.InvokeQtFunc6("C_QGraphicsWidgetList_rbegin_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_reverse_iterator rend()
func (this *QGraphicsWidgetList) Rend1() {
	// QGraphicsWidgetList_rend_1()
	rv, err := qtrt.InvokeQtFunc6("C_QGraphicsWidgetList_rend_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_reverse_iterator crbegin()
func (this *QGraphicsWidgetList) Crbegin0() {
	// QGraphicsWidgetList_crbegin_0()
	rv, err := qtrt.InvokeQtFunc6("C_QGraphicsWidgetList_crbegin_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_reverse_iterator crend()
func (this *QGraphicsWidgetList) Crend0() {
	// QGraphicsWidgetList_crend_0()
	rv, err := qtrt.InvokeQtFunc6("C_QGraphicsWidgetList_crend_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::iterator insert(QList::iterator, const T &)
func (this *QGraphicsWidgetList) Insert1() {
	// QGraphicsWidgetList_insert_1()
	rv, err := qtrt.InvokeQtFunc6("C_QGraphicsWidgetList_insert_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::iterator erase(QList::iterator)
func (this *QGraphicsWidgetList) Erase0() {
	// QGraphicsWidgetList_erase_0()
	rv, err := qtrt.InvokeQtFunc6("C_QGraphicsWidgetList_erase_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::iterator erase(QList::iterator, QList::iterator)
func (this *QGraphicsWidgetList) Erase1() {
	// QGraphicsWidgetList_erase_1()
	rv, err := qtrt.InvokeQtFunc6("C_QGraphicsWidgetList_erase_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// int count()
func (this *QGraphicsWidgetList) Count1() int {
	// QGraphicsWidgetList_count_1()
	rv, err := qtrt.InvokeQtFunc6("C_QGraphicsWidgetList_count_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// int length()
func (this *QGraphicsWidgetList) Length0() int {
	// QGraphicsWidgetList_length_0()
	rv, err := qtrt.InvokeQtFunc6("C_QGraphicsWidgetList_length_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// T & first()
func (this *QGraphicsWidgetList) First0() *QGraphicsWidget {
	// QGraphicsWidgetList_first_0()
	rv, err := qtrt.InvokeQtFunc6("C_QGraphicsWidgetList_first_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QGraphicsWidget{}
}

// const T & constFirst()
func (this *QGraphicsWidgetList) ConstFirst0() *QGraphicsWidget {
	// QGraphicsWidgetList_constFirst_0()
	rv, err := qtrt.InvokeQtFunc6("C_QGraphicsWidgetList_constFirst_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QGraphicsWidget{}
}

// const T & first()
func (this *QGraphicsWidgetList) First1() *QGraphicsWidget {
	// QGraphicsWidgetList_first_1()
	rv, err := qtrt.InvokeQtFunc6("C_QGraphicsWidgetList_first_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QGraphicsWidget{}
}

// T & last()
func (this *QGraphicsWidgetList) Last0() *QGraphicsWidget {
	// QGraphicsWidgetList_last_0()
	rv, err := qtrt.InvokeQtFunc6("C_QGraphicsWidgetList_last_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QGraphicsWidget{}
}

// const T & last()
func (this *QGraphicsWidgetList) Last1() *QGraphicsWidget {
	// QGraphicsWidgetList_last_1()
	rv, err := qtrt.InvokeQtFunc6("C_QGraphicsWidgetList_last_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QGraphicsWidget{}
}

// const T & constLast()
func (this *QGraphicsWidgetList) ConstLast0() *QGraphicsWidget {
	// QGraphicsWidgetList_constLast_0()
	rv, err := qtrt.InvokeQtFunc6("C_QGraphicsWidgetList_constLast_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QGraphicsWidget{}
}

// void removeFirst()
func (this *QGraphicsWidgetList) RemoveFirst0() {
	// QGraphicsWidgetList_removeFirst_0()
	rv, err := qtrt.InvokeQtFunc6("C_QGraphicsWidgetList_removeFirst_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void removeLast()
func (this *QGraphicsWidgetList) RemoveLast0() {
	// QGraphicsWidgetList_removeLast_0()
	rv, err := qtrt.InvokeQtFunc6("C_QGraphicsWidgetList_removeLast_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// bool startsWith(const T &)
func (this *QGraphicsWidgetList) StartsWith0() bool {
	// QGraphicsWidgetList_startsWith_0()
	rv, err := qtrt.InvokeQtFunc6("C_QGraphicsWidgetList_startsWith_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// bool endsWith(const T &)
func (this *QGraphicsWidgetList) EndsWith0() bool {
	// QGraphicsWidgetList_endsWith_0()
	rv, err := qtrt.InvokeQtFunc6("C_QGraphicsWidgetList_endsWith_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// QList<T> mid(int, int)
func (this *QGraphicsWidgetList) Mid0() *QGraphicsWidgetList {
	// QGraphicsWidgetList_mid_0()
	rv, err := qtrt.InvokeQtFunc6("C_QGraphicsWidgetList_mid_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// T value(int)
func (this *QGraphicsWidgetList) Value0() *QGraphicsWidget {
	// QGraphicsWidgetList_value_0()
	rv, err := qtrt.InvokeQtFunc6("C_QGraphicsWidgetList_value_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QGraphicsWidget{}
}

// T value(int, const T &)
func (this *QGraphicsWidgetList) Value1() *QGraphicsWidget {
	// QGraphicsWidgetList_value_1()
	rv, err := qtrt.InvokeQtFunc6("C_QGraphicsWidgetList_value_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QGraphicsWidget{}
}

// void push_back(const T &)
func (this *QGraphicsWidgetList) Push_back0() {
	// QGraphicsWidgetList_push_back_0()
	rv, err := qtrt.InvokeQtFunc6("C_QGraphicsWidgetList_push_back_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void push_front(const T &)
func (this *QGraphicsWidgetList) Push_front0() {
	// QGraphicsWidgetList_push_front_0()
	rv, err := qtrt.InvokeQtFunc6("C_QGraphicsWidgetList_push_front_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// T & front()
func (this *QGraphicsWidgetList) Front0() *QGraphicsWidget {
	// QGraphicsWidgetList_front_0()
	rv, err := qtrt.InvokeQtFunc6("C_QGraphicsWidgetList_front_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QGraphicsWidget{}
}

// const T & front()
func (this *QGraphicsWidgetList) Front1() *QGraphicsWidget {
	// QGraphicsWidgetList_front_1()
	rv, err := qtrt.InvokeQtFunc6("C_QGraphicsWidgetList_front_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QGraphicsWidget{}
}

// T & back()
func (this *QGraphicsWidgetList) Back0() *QGraphicsWidget {
	// QGraphicsWidgetList_back_0()
	rv, err := qtrt.InvokeQtFunc6("C_QGraphicsWidgetList_back_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QGraphicsWidget{}
}

// const T & back()
func (this *QGraphicsWidgetList) Back1() *QGraphicsWidget {
	// QGraphicsWidgetList_back_1()
	rv, err := qtrt.InvokeQtFunc6("C_QGraphicsWidgetList_back_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QGraphicsWidget{}
}

// void pop_front()
func (this *QGraphicsWidgetList) Pop_front0() {
	// QGraphicsWidgetList_pop_front_0()
	rv, err := qtrt.InvokeQtFunc6("C_QGraphicsWidgetList_pop_front_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void pop_back()
func (this *QGraphicsWidgetList) Pop_back0() {
	// QGraphicsWidgetList_pop_back_0()
	rv, err := qtrt.InvokeQtFunc6("C_QGraphicsWidgetList_pop_back_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// bool empty()
func (this *QGraphicsWidgetList) Empty0() bool {
	// QGraphicsWidgetList_empty_0()
	rv, err := qtrt.InvokeQtFunc6("C_QGraphicsWidgetList_empty_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// QList<T> & operator+=(const QList<T> &)
func (this *QGraphicsWidgetList) Operator_add_equal0() *QGraphicsWidgetList {
	// QGraphicsWidgetList_operator_add_equal_0()
	rv, err := qtrt.InvokeQtFunc6("C_QGraphicsWidgetList_operator_add_equal_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QList<T> operator+(const QList<T> &)
func (this *QGraphicsWidgetList) Operator_add0() *QGraphicsWidgetList {
	// QGraphicsWidgetList_operator_add_0()
	rv, err := qtrt.InvokeQtFunc6("C_QGraphicsWidgetList_operator_add_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QList<T> & operator+=(const T &)
func (this *QGraphicsWidgetList) Operator_add_equal1() *QGraphicsWidgetList {
	// QGraphicsWidgetList_operator_add_equal_1()
	rv, err := qtrt.InvokeQtFunc6("C_QGraphicsWidgetList_operator_add_equal_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QList<T> & operator<<(const T &)
func (this *QGraphicsWidgetList) Operator_left_shift0() *QGraphicsWidgetList {
	// QGraphicsWidgetList_operator_left_shift_0()
	rv, err := qtrt.InvokeQtFunc6("C_QGraphicsWidgetList_operator_left_shift_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QList<T> & operator<<(const QList<T> &)
func (this *QGraphicsWidgetList) Operator_left_shift1() *QGraphicsWidgetList {
	// QGraphicsWidgetList_operator_left_shift_1()
	rv, err := qtrt.InvokeQtFunc6("C_QGraphicsWidgetList_operator_left_shift_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QVector<T> toVector()
func (this *QGraphicsWidgetList) ToVector0() {
	// QGraphicsWidgetList_toVector_0()
	rv, err := qtrt.InvokeQtFunc6("C_QGraphicsWidgetList_toVector_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QSet<T> toSet()
func (this *QGraphicsWidgetList) ToSet0() {
	// QGraphicsWidgetList_toSet_0()
	rv, err := qtrt.InvokeQtFunc6("C_QGraphicsWidgetList_toSet_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList<T> fromVector(const QVector<T> &)
func (this *QGraphicsWidgetList) FromVector0() *QGraphicsWidgetList {
	// QGraphicsWidgetList_fromVector_0()
	rv, err := qtrt.InvokeQtFunc6("C_QGraphicsWidgetList_fromVector_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QList<T> fromSet(const QSet<T> &)
func (this *QGraphicsWidgetList) FromSet0() *QGraphicsWidgetList {
	// QGraphicsWidgetList_fromSet_0()
	rv, err := qtrt.InvokeQtFunc6("C_QGraphicsWidgetList_fromSet_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QList<T> fromStdList(const std::list<T> &)
func (this *QGraphicsWidgetList) FromStdList0() *QGraphicsWidgetList {
	// QGraphicsWidgetList_fromStdList_0()
	rv, err := qtrt.InvokeQtFunc6("C_QGraphicsWidgetList_fromStdList_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// std::list<T> toStdList()
func (this *QGraphicsWidgetList) ToStdList0() {
	// QGraphicsWidgetList_toStdList_0()
	rv, err := qtrt.InvokeQtFunc6("C_QGraphicsWidgetList_toStdList_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::Node * detach_helper_grow(int, int)
func (this *QGraphicsWidgetList) Detach_helper_grow0() {
	// QGraphicsWidgetList_detach_helper_grow_0()
	rv, err := qtrt.InvokeQtFunc6("C_QGraphicsWidgetList_detach_helper_grow_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void detach_helper(int)
func (this *QGraphicsWidgetList) Detach_helper0() {
	// QGraphicsWidgetList_detach_helper_0()
	rv, err := qtrt.InvokeQtFunc6("C_QGraphicsWidgetList_detach_helper_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void detach_helper()
func (this *QGraphicsWidgetList) Detach_helper1() {
	// QGraphicsWidgetList_detach_helper_1()
	rv, err := qtrt.InvokeQtFunc6("C_QGraphicsWidgetList_detach_helper_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void dealloc(QListData::Data *)
func (this *QGraphicsWidgetList) Dealloc0() {
	// QGraphicsWidgetList_dealloc_0()
	rv, err := qtrt.InvokeQtFunc6("C_QGraphicsWidgetList_dealloc_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void node_construct(QList::Node *, const T &)
func (this *QGraphicsWidgetList) Node_construct0() {
	// QGraphicsWidgetList_node_construct_0()
	rv, err := qtrt.InvokeQtFunc6("C_QGraphicsWidgetList_node_construct_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void node_destruct(QList::Node *)
func (this *QGraphicsWidgetList) Node_destruct0() {
	// QGraphicsWidgetList_node_destruct_0()
	rv, err := qtrt.InvokeQtFunc6("C_QGraphicsWidgetList_node_destruct_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void node_copy(QList::Node *, QList::Node *, QList::Node *)
func (this *QGraphicsWidgetList) Node_copy0() {
	// QGraphicsWidgetList_node_copy_0()
	rv, err := qtrt.InvokeQtFunc6("C_QGraphicsWidgetList_node_copy_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void node_destruct(QList::Node *, QList::Node *)
func (this *QGraphicsWidgetList) Node_destruct1() {
	// QGraphicsWidgetList_node_destruct_1()
	rv, err := qtrt.InvokeQtFunc6("C_QGraphicsWidgetList_node_destruct_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// bool isValidIterator(const QList::iterator &)
func (this *QGraphicsWidgetList) IsValidIterator0() bool {
	// QGraphicsWidgetList_isValidIterator_0()
	rv, err := qtrt.InvokeQtFunc6("C_QGraphicsWidgetList_isValidIterator_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// bool op_eq_impl(const QList<T> &, QListData::NotArrayCompatibleLayout)
func (this *QGraphicsWidgetList) Op_eq_impl0() bool {
	// QGraphicsWidgetList_op_eq_impl_0()
	rv, err := qtrt.InvokeQtFunc6("C_QGraphicsWidgetList_op_eq_impl_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// bool op_eq_impl(const QList<T> &, QListData::ArrayCompatibleLayout)
func (this *QGraphicsWidgetList) Op_eq_impl1() bool {
	// QGraphicsWidgetList_op_eq_impl_1()
	rv, err := qtrt.InvokeQtFunc6("C_QGraphicsWidgetList_op_eq_impl_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// bool contains_impl(const T &, QListData::NotArrayCompatibleLayout)
func (this *QGraphicsWidgetList) Contains_impl0() bool {
	// QGraphicsWidgetList_contains_impl_0()
	rv, err := qtrt.InvokeQtFunc6("C_QGraphicsWidgetList_contains_impl_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// bool contains_impl(const T &, QListData::ArrayCompatibleLayout)
func (this *QGraphicsWidgetList) Contains_impl1() bool {
	// QGraphicsWidgetList_contains_impl_1()
	rv, err := qtrt.InvokeQtFunc6("C_QGraphicsWidgetList_contains_impl_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// int count_impl(const T &, QListData::NotArrayCompatibleLayout)
func (this *QGraphicsWidgetList) Count_impl0() int {
	// QGraphicsWidgetList_count_impl_0()
	rv, err := qtrt.InvokeQtFunc6("C_QGraphicsWidgetList_count_impl_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// int count_impl(const T &, QListData::ArrayCompatibleLayout)
func (this *QGraphicsWidgetList) Count_impl1() int {
	// QGraphicsWidgetList_count_impl_1()
	rv, err := qtrt.InvokeQtFunc6("C_QGraphicsWidgetList_count_impl_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

//  body block end
