package qtqt

import "unsafe"

func DisconnectObject(cPtr unsafe.Pointer) {
	signalsMutex.Lock()
	if _, ok := signals[cPtr]; ok {
		delete(signals, cPtr)
	}
	signalsMutex.Unlock()
	Unregister(cPtr)
}

func DumpSignalsByObj(cPtr unsafe.Pointer) {
	Debug("##############################\tSIGNALSTABLE_START\t##############################")
	signalsMutex.Lock()
	if entry, ok := signals[cPtr]; ok {
		Debug(cPtr, entry)
	}
	signalsMutex.Unlock()
	Debug("##############################\tSIGNALSTABLE_END\t##############################")
}

// get connected signals
func GetActiveSignals(cPtr unsafe.Pointer) (names []string) {
	signalsMutex.Lock()
	if _, ok := signals[cPtr]; ok {
		for signame, _ := range signals[cPtr] {
			names = append(names, signame)
		}
	}
	signalsMutex.Unlock()
	return
}

func CountSignalsByObj(cPtr unsafe.Pointer) (c int) {
	signalsMutex.Lock()
	if entry, ok := signals[cPtr]; ok {
		c = len(entry)
	}
	signalsMutex.Unlock()
	return
}
