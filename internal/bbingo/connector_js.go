//go:build wasm && js

package bbingo

import (
	"github.com/nats-io/nats.go"
	"sync"
	"syscall/js"
)

/* TODO we need
- Publish(subject,data)
- Test si stream a des messages
- ET lastmessage ()
- Chansubscribe
*/

var (
	streamMsgC         = make(chan bool)
	streamLastMsg      = make(chan string)
	initOnce           sync.Once
	StreamHasMessageCB = js.FuncOf(func(this js.Value, args []js.Value) any {
		if len(args) == 0 || args[0].Type() != js.TypeBoolean {
			streamMsgC <- false
			return nil
		}
		streamMsgC <- args[0].Bool()
		return nil
	})
	StreamLastMsgCB = js.FuncOf(func(this js.Value, args []js.Value) any {
		if len(args) == 0 || args[0].Type() != js.TypeString {
			streamLastMsg <- ""
			return nil
		}
		streamLastMsg <- args[0].String()
		return nil
	})
	SubCB = js.FuncOf(func(this js.Value, args []js.Value) any {
		if len(args) != 3 || args[0].Type() != js.TypeString || args[1].Type() != js.TypeString || args[2].Type() != js.TypeNumber {
			return nil
		}
		chans[args[1].Int()] <- msgData{args[0].String(), args[1].String()}
		return nil
	})
)

func init() {
	// set up some callbacks for JS
	window := js.Global()
	doc := window.Get("document")
	window.Set("StreamHasMessageCB", StreamHasMessageCB)
	doc.Set("StreamLastMsgCB", StreamLastMsgCB)
	window.Get("console").Call("log", "go harnessed")
}

func Publish(subject, data string) { js.Global().Call("publish", subject, data) }

func StreamHasMessage(streamname string) bool {
	js.Global().Call("streamHasMsgs", streamname)
	return <-streamMsgC
}

func StreamLastMessage(streamname string) string {
	js.Global().Call("streamLastMsg", streamname)
	return <-streamLastMsg
}

type msgData struct{ Subject, Data string }

var chans = make([]chan msgData, 0)
var chansM sync.Mutex

func ChanSubscribe(subject string, ch chan *nats.Msg) (*nats.Subscription, error) { // TODO
	js.Global().Call("subscribe", subject, len(chans))
	chansM.Lock()
	c := make(chan msgData)
	chans = append(chans, c)
	go func(c chan msgData) {
		for m := range c {
			ch <- &nats.Msg{Subject: m.Subject, Data: []byte(m.Data)} // FIXME you can't use NATS here... can you ?
		}
	}(c)
	return nil, nil
}
