package buffconn

import (
	"errors"
	"io"
	"net"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/iotaledger/hive.go/core/generics/event"
)

const graceTime = 10 * time.Millisecond
const maxMessageSize = 65 * 1024

var testMsg = []byte("test")

func TestBufferedConnection(t *testing.T) {
	t.Run("Close", func(t *testing.T) {
		conn1, conn2 := net.Pipe()
		buffConn1 := NewBufferedConnection(conn1, maxMessageSize)
		defer buffConn1.Close()
		buffConn2 := NewBufferedConnection(conn2, maxMessageSize)
		defer buffConn2.Close()

		var wg sync.WaitGroup
		wg.Add(2)
		buffConn2.Events.Close.Hook(event.NewClosure(func(_ *CloseEvent) { wg.Done() }))
		go func() {
			err := buffConn2.Read()
			assert.True(t, errors.Is(err, io.ErrClosedPipe), "unexpected error: %s", err)
			require.NoError(t, buffConn2.Close())
			wg.Done()
		}()

		err := buffConn1.Close()
		require.NoError(t, err)
		wg.Wait()
	})

	t.Run("Write", func(t *testing.T) {
		conn1, conn2 := net.Pipe()
		buffConn1 := NewBufferedConnection(conn1, maxMessageSize)
		defer buffConn1.Close()
		buffConn2 := NewBufferedConnection(conn2, maxMessageSize)
		defer buffConn2.Close()

		go func() {
			_ = buffConn2.Read()
		}()

		n, err := buffConn1.Write(testMsg)
		require.NoError(t, err)
		assert.EqualValues(t, len(testMsg), n)
	})

	t.Run("ReceiveMessage", func(t *testing.T) {
		conn1, conn2 := net.Pipe()
		buffConn1 := NewBufferedConnection(conn1, maxMessageSize)
		defer buffConn1.Close()
		buffConn2 := NewBufferedConnection(conn2, maxMessageSize)
		defer buffConn2.Close()

		var wg sync.WaitGroup
		wg.Add(2)
		buffConn2.Events.ReceiveMessage.Hook(event.NewClosure(func(event *ReceiveMessageEvent) {
			assert.EqualValues(t, testMsg, event.Data)
			wg.Done()
		}))
		go func() {
			err := buffConn2.Read()
			assert.True(t, errors.Is(err, io.EOF), "unexpected error: %s", err)
			wg.Done()
		}()

		n, err := buffConn1.Write(testMsg)
		require.NoError(t, err)
		assert.EqualValues(t, len(testMsg), n)

		time.Sleep(graceTime)

		err = buffConn1.Close()
		require.NoError(t, err)
		wg.Wait()
	})

	t.Run("ReceiveMany", func(t *testing.T) {
		conn1, conn2 := net.Pipe()
		buffConn1 := NewBufferedConnection(conn1, maxMessageSize)
		defer buffConn1.Close()
		buffConn2 := NewBufferedConnection(conn2, maxMessageSize)
		defer buffConn2.Close()

		const numWrites = 3

		var wg sync.WaitGroup
		wg.Add(numWrites)
		buffConn2.Events.ReceiveMessage.Hook(event.NewClosure(func(event *ReceiveMessageEvent) {
			assert.Equal(t, testMsg, event.Data)
			wg.Done()
		}))
		go func() {
			_ = buffConn2.Read()
		}()

		for i := 1; i <= numWrites; i++ {
			_, err := buffConn1.Write(testMsg)
			require.NoError(t, err)
			if i < numWrites {
				time.Sleep(IOTimeout + graceTime)
			}
		}
		wg.Wait()
	})

	t.Run("InvalidHeader", func(t *testing.T) {
		conn1, conn2 := net.Pipe()
		buffConn1 := NewBufferedConnection(conn1, maxMessageSize)
		defer buffConn1.Close()
		defer conn2.Close()

		var wg sync.WaitGroup
		wg.Add(1)
		go func() {
			err := buffConn1.Read()
			assert.True(t, errors.Is(err, ErrInvalidHeader), "unexpected error: %s", err)
			wg.Done()
		}()

		_, err := conn2.Write([]byte{0xff, 0xff, 0xff, 0xff})
		require.NoError(t, err)
		wg.Wait()
	})
}
