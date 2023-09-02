package initialization

import (
	"context"
	"douyin/global"
)

func InitializeContext() {
	ctx, _ := context.WithCancel(context.Background())
	global.SERVER_CONTEXT = &ctx
	// 	defer cancel()

	// // Cleanly shutdown and flush telemetry when the application exits.
	//
	//	defer func(ctx context.Context) {
	//		// Do not make the application hang when it is shutdown.
	//		ctx, cancel = context.WithTimeout(ctx, time.Second*5)
	//		defer cancel()
	//		if err := global.SERVER_TRACE_PROVIDER.Shutdown(ctx); err != nil {
	//			log.Fatal(err)
	//		}
	//	}(ctx)
}
