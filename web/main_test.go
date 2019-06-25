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
	ws.WriteMessage(1, []byte("{ \"type\": \"new\", \"width\": 80, \"height\": 80, \"tick\": 10 }"))
	err = ws.ReadJSON(&ack)
	assert.NoError(err)
	assert.Equal(newAckOk(), ack)

	var gs engine.GameState
	ws.WriteMessage(1, []byte("{ \"type\": \"new\", \"width\": 80, \"height\": 80, \"tick\": 10 }"))

	// Read first state back from engine
	err = ws.ReadJSON(&gs)
	assert.NoError(err)
	assert.IsType(engine.GameState{}, gs)

	err = ws.ReadJSON(&ack)
	assert.NoError(err)
	assert.Equal(newAckError("game already exists"), ack)
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
	ws.WriteMessage(1, []byte("{ \"type\": \"new\", \"width\": 80, \"height\": 80, \"tick\": 10 }"))
	err = ws.ReadJSON(&ack)
	assert.NoError(err)
	assert.Equal(ack, newAckOk())

	var gs engine.GameState
	// Read first state back from engine
	err = ws.ReadJSON(&gs)
	assert.NoError(err)
	assert.IsType(engine.GameState{}, gs)

	ws.WriteMessage(1, []byte("{ \"type\": \"destroy\" }"))
	err = ws.ReadJSON(&ack)
	assert.NoError(err)
	assert.Equal(ack, newAckOk())
}
