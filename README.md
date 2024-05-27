# gRPC2

1) greetGRPC : unary RPC ; client sends name ; server responds with "Hello" + name sent by client

2) sumURPC : unray RPC ; client takes 2 number from user & sends them to server ; server adds them up & returns result

3) serverStream : server streaming RPC ; modified version of greetGRPC ; here client sends name ; server responds with `STREAM` of greets
