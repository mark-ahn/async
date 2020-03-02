package async

// import "context"

// var pool = NewPoolOfSomeThenOtherImpl()

// func ChainOfSomeThenOtherUsingAnother(ctx context.Context, req *WorkOfSomeThenOther, start WorkerOfPushSomeThenOther, next WorkerOfPushSomeThenOther) {
// 	go func() {
// 		ch := pool.GetChReturnOfOther()
// 		defer pool.PutChReturnOfOther(ch)

// 		work := pool.GetWorkOfSomeThenOther()
// 		work.Argument = req.Argument
// 		work.ReturnCh = ch

// 		start.Push(ctx, work)
// 		defer func() {
// 			work.ReturnCh = nil
// 			pool.PutWorkOfSomeThenOther(work)
// 		}()

// 		next_arg := <-ch

// 		work.Argument = next_arg
// 		work.ReturnCh = req.ReturnCh
// 		next.Push(ctx, work)
// 	}()
// }
