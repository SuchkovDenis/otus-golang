package hw06pipelineexecution

type (
	In  = <-chan interface{}
	Out = In
	Bi  = chan interface{}
)

type Stage func(in In) (out Out)

func ExecutePipeline(in In, done In, stages ...Stage) Out {
	out := in
	for i := range stages {
		out = wrap(done, stages[i](out))
	}
	return out
}

func wrap(done, in In) Out {
	out := make(Bi)
	go func() {
		defer close(out)
		for {
			select {
			case el, ok := <-in:
				if !ok {
					return
				}
				out <- el
			case <-done:
				return
			}
		}
	}()
	return out
}
