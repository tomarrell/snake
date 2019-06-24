package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gorilla/websocket"
	"github.com/stretchr/testify/assert"
	"github.com/tomarrell/snake/engine"
)

func TestInvalidJSON(t *testing.T) {
	assert := assert.New(t)

	e := engine.NewEngine()
	s := httptest.NewServer(http.HandlerFunc(websocketHandler(e)))
	u := "ws" + strings.TrimPrefix(s.URL, "http")

	ws, _, err := websocket.DefaultDialer.Dial(u, nil)
	if err != nil {
		t.Fatalf("%v", err)
	}
	defer ws.Close()

	var ack mAck
	ws.WriteMessage(1, []byte("{ \"type\": "))
	err = ws.ReadJSON(&ack)
	assert.NoError(err)
	assert.Equal(ack, newAckError("message is not valid json"))
}

func TestCreate(t *testing.T) {
	assert := assert.New(t)

	e := engine.NewEngine()
	s := httptest.NewServer(http.HandlerFunc(websocketHandler(e)))

	// Convert http://127.0.0.1 to ws://127.0.0.
	u := "ws" + strings.TrimPrefix(s.URL, "http")

	// Connect to the server
	ws, _, err := websocket.DefaultDialer.Dial(u, nil)
	if err != nil {
		t.Fatalf("%v", err)
	}
	defer ws.Close()

	var ack mAck
	ws.WriteMessage(1, []byte("{ \"type\": \"new\" }"))
	err = ws.ReadJSON(&ack)
	assert.NoError(err)
	assert.Equal(ack, newAckOk())

	ws.WriteMessage(1, []byte("{ \"type\": \"new\" }"))
	err = ws.ReadJSON(&ack)
	assert.NoError(err)
	assert.Equal(ack, newAckError("game already exists"))
}

func TestDestroy(t *testing.T) {
	assert := assert.New(t)

	e := engine.NewEngine()
	s := httptest.NewServer(http.HandlerFunc(websocketHandler(e)))
	u := "ws" + strings.TrimPrefix(s.URL, "http")

	ws, _, err := websocket.DefaultDialer.Dial(u, nil)
	if err != nil {
		t.Fatalf("%v", err)
	}
	defer ws.Close()

	var ack mAck
	ws.WriteMessage(1, []byte("{ \"type\": \"new\" }"))
	err = ws.ReadJSON(&ack)
	assert.NoError(err)
	assert.Equal(ack, newAckOk())

	ws.WriteMessage(1, []byte("{ \"type\": \"destroy\" }"))
	err = ws.ReadJSON(&ack)
	assert.NoError(err)
	assert.Equal(ack, newAckOk())
}
