package clock

import "time"

type Cmos struct {
	Second     int
	Minute     int
	Hour       int
	DayOfMonth int
	Month      int
	Year       int
	Century    int
}

func ReadCmos() Cmos {
	var t Cmos
	for {
		readCmos(&t)
		if bcdDecode(readCmosSecond()) == t.Second {
			break
		}
	}
	return t
}

func (c *Cmos) Time() time.Time {
	return time.Date(c.Year, time.Month(c.Month), c.DayOfMonth, c.Hour, c.Minute, c.Second, 0, time.UTC)
}

// https://wiki.osdev.org/CMOS
func readCmos(t *Cmos) {
	t.Century = bcdDecode(readCmosReg(0x32))
	t.Year = bcdDecode(readCmosReg(0x09)) + bcdDecode(readCmosReg(0x32))*100
	t.Month = bcdDecode(readCmosReg(0x08))
	t.DayOfMonth = bcdDecode(readCmosReg(0x07))
	t.Hour = bcdDecode(readCmosReg(0x06))
	t.Minute = bcdDecode(readCmosReg(0x2))
	t.Second = bcdDecode(readCmosReg(0x00))

}

func readCmosSecond() int {
	return readCmosReg(0x00)
}

func bcdDecode(v int) int {
	return v&0x0F + v/16*10
}

//go:NOSPLIT
func readCmosReg(reg int16) int {
	sys.Outb(0x70, 0x80|byte(reg))
	return int(sys.Inb(0x71))
}
