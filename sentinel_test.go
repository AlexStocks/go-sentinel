package sentinel

import (
	"testing"
	"fmt"
	"sync"
	"time"
)

import (
	"github.com/garyburd/redigo/redis"
)

func TestSentinel(t *testing.T) {
	st := NewSentinel(
		[]string{"192.168.10.100:26380"},
	)
	defer st.Close()

	instances, err := st.GetInstances()
	if err != nil {
		t.Errorf("st.GetInstances, error:%#v\n", err)
		t.FailNow()
	}

	for idx, inst := range instances {
		t.Logf("idx:%d, instance:%#v\n", idx, inst)
		err = st.Discover(inst.Name)
		if err != nil {
			t.Log(err)
			t.FailNow()
		}
	}

	addrs := st.GetSentinels()
	t.Logf("sentinel instances:%#v\n", addrs)

	st.Close()
}

func TestSentinelWatcher(t *testing.T) {
	st := NewSentinel(
		[]string{"192.168.10.100:26380"},
	)
	defer st.Close()

	instances, err := st.GetInstances()
	if err != nil {
		t.Errorf("st.GetInstances, error:%#v\n", err)
		t.FailNow()
	}

	for _, inst := range instances {
		err = st.Discover(inst.Name)
		if err != nil {
			t.Log(err)
			t.FailNow()
		}
	}

	wg := &sync.WaitGroup{}
	watcher, err := st.MakeSentinelWatcher()
	w, err := watcher.Watch()
	_ = w
	go func() {
		defer wg.Done()
		wg.Add(1)
		for addr := range w {
			t.Logf("redis instance switch: %#v\n", addr)
		}
		fmt.Println("watch exit")
	}()
	time.Sleep(20 * time.Second)
	fmt.Println("close")
	watcher.Close()
	wg.Wait()
}

func TestSentinelPool(t *testing.T) {
	sp := NewSentinel(
		[]string{"192.168.10.100:26380"},
	)
	defer sp.Close()

	instances, err := sp.GetInstances()
	if err != nil {
		t.Errorf("st.GetInstances, error:%#v\n", err)
		t.FailNow()
	}

	for i, inst := range instances {
		err = sp.Discover(inst.Name)
		if err != nil {
			t.Log(err)
			t.FailNow()
		}

		conn := sp.GetConn(fmt.Sprintf("%s:%d", inst.Master.IP, inst.Master.Port))
		if conn == nil {
			fmt.Println("get conn fail, ", i)
			continue
		}
		s, err := redis.String(conn.Do("INFO"))
		if err != nil {
			fmt.Println("do command error:", err)
			fmt.Printf("do command error for master addr{idx:%s, addr:%#v}", i, inst.Master)
			continue
		}
		fmt.Printf("idx:%s, addr:%#v, info:%#v", i, inst.Master, s)
		time.Sleep(1 * time.Second)
	}
}