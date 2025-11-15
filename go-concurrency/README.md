NOTES

- `sync.WaitGroup` is a struct with internal state so you should pass it by reference instead of value.
  If you pass by reference, for example, while handling concurrency with Go routines, each Go routine
  will have it's own copy of the `WaitGroup`, so calling `Done` on a copy will not affect the original.
  Main built-in structs with internal state that need pointers:
  - sync.WaitGroup
    - sync.Mutex / sync.RWMutex
    - sync.Cond
    - sync.Pool
    - sync.Map
    - sync.Once
    - context.Context
      These manage internal counters, locks, or state that must be shared.
- There's `Print` but there's also `Sprint`. Can you believe this? I though the AI was hallucinating.
  Here's [the doc](https://pkg.go.dev/fmt#example-Sprint). Apparently it formats and returns, which is
  different than printing to the console with `Print`. Nice to know.
- Atomics only work in specific types: int32, int64, uint32, uint64, pointers.

  Jon Bodner on atomics:

  > "If you need to squeeze
  > out every last bit of performance and are an expert on writing concurrent code, you’ll be glad that Go
  > includes atomic support. For everyone else, use goroutines and mutexes to manage your concurrency needs".

-
