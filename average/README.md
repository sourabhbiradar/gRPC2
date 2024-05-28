# Average

client sends `stream` of numbers ; server waits for all requests , calculates average of numbers & returns responds with result.

`NEW` : Client sends `int32` ; server responds with `double`

`logic` : server >> avg.go

    var sum int32= 0
	count := 0

	for {
		req, err := strm.Recv()

		if err == io.EOF {
           return strm.SendAndClose(&pb.AvgRes{
			Result: float64(sum) / float64(count), // calulates avg
		   })
		}

		if err != nil {
			log.Fatalf("Error while reading inputs : %v\n", err)
		}

		sum += req.Number  // sum = sum + req.Number
		count++  // count = count + 1
	}

 `cmd` : make protoc 

 `cmd` : make serverrun

 `cmd` : make clientrun
