package main

import (
	"github.com/radicalmind/xeon"
	"github.com/radicalmind/xeon/context"
	"github.com/radicalmind/xeon/sessions"
)

func main() {
	app := xeon.New()
	// output startup banner and error logs on os.Stdout

	sess := sessions.New(sessions.Config{Cookie: "myappsessionid"})
	app.AttachSessionManager(sess)

	app.Get("/set", func(ctx context.Context) {
		ctx.Session().SetFlash("name", "xeon")
		ctx.Writef("Message setted, is available for the next request")
	})

	app.Get("/get", func(ctx context.Context) {
		name := ctx.Session().GetFlashString("name")
		if name == "" {
			ctx.Writef("Empty name!!")
			return
		}
		ctx.Writef("Hello %s", name)
	})

	app.Get("/test", func(ctx context.Context) {
		name := ctx.Session().GetFlashString("name")
		if name == "" {
			ctx.Writef("Empty name!!")
			return
		}

		ctx.Writef("Ok you are comming from /set ,the value of the name is %s", name)
		ctx.Writef(", and again from the same context: %s", name)

	})

	app.Run(xeon.Addr(":8080"))
}
