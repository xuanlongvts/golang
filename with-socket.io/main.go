package main

import (
	socketio "github.com/googollee/go-socket.io"
	"github.com/labstack/echo"
	"log"
)

func main() {
	server := socketio.NewServer(nil)
	server.OnConnect("/", func(conn socketio.Conn) error {
		conn.SetContext("")
		log.Println("connected: ", conn.ID())
		return nil
	})

	server.OnEvent("/", "notice", func(s socketio.Conn, msg string) {
		log.Println("notice: ", msg)
		s.Emit("reply", " have "+msg)
	})

	server.OnEvent("/chat", "msg", func(s socketio.Conn, msg string) string {
		s.SetContext(msg)
		return "recv " + msg
	})

	server.OnEvent("/", "echo", func(s socketio.Conn, msg interface{}) {
		s.Emit("echo: ", msg)
	})

	server.OnEvent("/", "bye", func(s socketio.Conn) string {
		last := s.Context().(string)
		s.Emit("bye", last)
		s.Close()
		return last
	})

	server.OnError("/", func(conn socketio.Conn, err error) {
		log.Println("meet error: ", err.Error())
	})

	server.OnDisconnect("/", func(conn socketio.Conn, s string) {
		log.Println("closed: ", s)
	})

	go server.Serve()
	defer server.Close()

	e := echo.New()
	e.HideBanner = true

	e.Static("/", "./asset")
	e.Any("socket.io/", func(context echo.Context) error {
		server.ServeHTTP(context.Response(), context.Request())
		return nil
	})
	e.Logger.Fatal(e.Start(":1323"))
}
