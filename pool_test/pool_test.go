package pool_test

import (
	"errors"
	"testing"
	"time"
)

type ReusableObj struct {
}

type ObjPool struct {
	bufChan chan *ReusableObj
}

func NewObjPool(num int) *ObjPool {
	objPool := &ObjPool{}
	objPool.bufChan = make(chan *ReusableObj, num)
	for i := 0; i < num; i++ {
		objPool.bufChan <- &ReusableObj{}
	}
	return objPool
}

func (p *ObjPool) GetObj(timeout time.Duration) (*ReusableObj, error) {
	select {
	case ret := <-p.bufChan:
		return ret, nil
	case <-time.After(timeout):
		return nil, errors.New("timeout")
	}
}

func (p *ObjPool) ReleaseObj(obj *ReusableObj) error {
	select {
	case p.bufChan <- obj:
		return nil
	default:
		return errors.New("overflow....")
	}
}

func TestPool(t *testing.T) {
	pool := NewObjPool(10)
	// 由于线程池是满的，不能往里面增加资源
	// if err := pool.ReleaseObj(&ReusableObj{}); err != nil {
	// 	t.Error(err)
	// }
	for i := 0; i < 10; i++ {
		if v, err := pool.GetObj(1 * time.Second); err != nil {
			t.Error(err)
		} else {
			t.Logf("%T", v)
			// if err := pool.ReleaseObj(v); err != nil {
			// 	t.Error(err)
			// }
		}
	}
}
