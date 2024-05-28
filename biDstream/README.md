# Bi-Directional streaming RPC

client sends `stream` of names ; server responds with greeting + names ; Async streaming

`NEW` : used `Go routines` to send & receive in client/greetE1.go

ch := make(chan struct{})  // Create a channel of type struct{}
close(ch)                 // Close the channel
<-ch                      // Receive from the channel

`cmd` : make protoc

`cmd` : make runserver

`cmd` : make runclient