# gRPC2

1) greetGRPC : unary RPC ; client sends name ; server responds with "Hello" + name sent by client

2) sumURPC : unray RPC ; client takes 2 number from user & sends them to server ; server adds them up & returns result

3) serverStream : server streaming RPC ; modified version of greetGRPC ; here client sends name ; server responds with `STREAM` of greets

4) factors : server streaming RPC ; Client sends an integer , server receives it & returns `stream` of `factors` of that number

5) clientStream : clinet streaming RPC ; client sends `stream` of names ; server receives all & respond/greets at once

6) average : client streaming RPC ; client sends `stream` of numbers ; server waits for all requests , calculates average of numbers & returns result.

7) biDstream : bi-directional streaming RPC ; client sends `stream` of names ; server responds with greeting + names ; Async streaming

8) maxNum : bi-directional streaming RPC ; client sends `stream` of numbers ; server asserts maximum among current & privous max number ; returns `stream` of `set of max numbers`
