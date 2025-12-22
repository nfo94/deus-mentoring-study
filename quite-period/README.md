Multithreading with Go

- Process are heavy weight, have their own memory space and other resources, so concurrency with it
  tends to be slower and consume more resources;
- Threads are units of work of a process. They share resources between them inside a single process;
- Green threads are user level threads. So when we’re dealing with concurrency here we call it
  cooperative multitasking: the user defines when to context switch the work;
- There are kernel level threads. So when we’re dealing with concurrency here we call it preemptive
  multitasking: the CPU decides when to context switch work;
- In Go we can synchronize threads with mutexes or channels;
