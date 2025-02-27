package main

import (
	"fmt"
	"sync"
	"time"
)

type Meeting struct {
	id           int
	meetingPIC   string
	meetingTitle string
	meetingTime  time.Duration
}

type MeetingSession struct {
	meetQueue chan Meeting
	wg        sync.WaitGroup
	mutex     sync.Mutex
}

func NewMeetingSession() *MeetingSession {
	return &MeetingSession{
		meetQueue: make(chan Meeting, 5),
	}
}

func (meetSession *MeetingSession) onMeet(meetingRoomID int) {
	defer meetSession.wg.Done()
	for meet := range meetSession.meetQueue {
		fmt.Printf("[Concurrency] Karyawan %s membuka meeting dengan pembahasan #%s pada ruangan %d\n",
			meet.meetingPIC, meet.meetingTitle, meetingRoomID)
		time.Sleep(meet.meetingTime * time.Second) // sebuah jeda
		fmt.Printf("[Concurrency] Karyawan %s menutup meeting dengan pembahasan #%s pada ruangan %d\n",
			meet.meetingPIC, meet.meetingTitle, meetingRoomID)
	}
}

func RunConcurrent(meetings []Meeting) time.Duration {
	start := time.Now()

	meetingSession := NewMeetingSession()

	meetingRoom := 3
	meetingSession.wg.Add(meetingRoom)

	for i := 1; i <= meetingRoom; i++ {
		go meetingSession.onMeet(i)
	}

	for _, meeting := range meetings {
		meetingSession.meetQueue <- meeting
	}

	close(meetingSession.meetQueue)
	meetingSession.wg.Wait()

	return time.Since(start)
}

func ConcurrencyCase() {
	meetings := []Meeting{
		{1, "Ichsan", "Pemasangan GPS Tracker pada Kendaraaan Surveyor", 2},
		{2, "Tasya", "Peningkatan penjualan untuk Div. Sales", 2},
		{3, "Fauzan", "Presentasi KPI Tim IT", 2},
		{4, "Andra", "Meeting All Branch Tim Purchasing", 2},
		{5, "Kinan", "Presentasi Vendor RFID untuk Warehouse", 2},
	}

	fmt.Println("Concurrency Process")
	concTime := RunConcurrent(meetings)
	fmt.Printf("Waktu concurrency: %v\n", concTime)
}

func main() {
	ConcurrencyCase()
}
