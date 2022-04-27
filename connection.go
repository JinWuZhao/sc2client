package sc2client

import (
	"context"
	"fmt"
	"io"

	"google.golang.org/protobuf/proto"
	"nhooyr.io/websocket"

	"github.com/jinwuzhao/sc2client/sc2proto"
)

type Connection struct {
	conn *websocket.Conn
}

func DialSC2(ctx context.Context, host string, port int) (*Connection, error) {
	wsURL := fmt.Sprintf("ws://%s:%d/sc2api", host, port)
	conn, _, err := websocket.Dial(ctx, wsURL, nil)
	if err != nil {
		return nil, fmt.Errorf("websocket.Dial() error: %w", err)
	}
	return &Connection{
		conn: conn,
	}, nil
}

func (c *Connection) Read(ctx context.Context, rsp *sc2proto.Response) error {
	typ, r, err := c.conn.Reader(ctx)
	if err != nil {
		return fmt.Errorf("c.conn.Reader() error: %w", err)
	}

	if typ != websocket.MessageBinary {
		_ = c.conn.Close(websocket.StatusUnsupportedData, "expected binary message")
		return fmt.Errorf("expected binary message for protobuf but got: %v", typ)
	}

	b, err := io.ReadAll(r)
	if err != nil {
		return fmt.Errorf("io.ReadAll() error: %w", err)
	}

	err = proto.Unmarshal(b, rsp)
	if err != nil {
		_ = c.conn.Close(websocket.StatusInvalidFramePayloadData, "failed to unmarshal protobuf")
		return fmt.Errorf("failed to unmarshal protobuf: %w", err)
	}

	return nil
}

func (c *Connection) Write(ctx context.Context, req *sc2proto.Request) error {
	b, err := proto.Marshal(req)
	if err != nil {
		return fmt.Errorf("proto.Marshal() error: %w", err)
	}

	err = c.conn.Write(ctx, websocket.MessageBinary, b)
	if err != nil {
		return fmt.Errorf("c.conn.Write() error: %w", err)
	}

	return nil
}

func (c *Connection) Close() {
	_ = c.conn.Close(websocket.StatusNoStatusRcvd, "")
}
