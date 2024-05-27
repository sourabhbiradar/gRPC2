# server streaming

stream of response from server

`NEW` : 
1) `stream` keyword in proto service block

2) GreetMultiple (req *pb.GreetReq, strm pb.GreetService_GreetMultipleServer)

3) `strm.Send()` & `strm.Recv()`



