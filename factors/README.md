# factors

# server streaming RPC

Client sends an integer , server receives it & returns `stream` of `factors` of that number

server >> factors.go

`logic` :

    number := num.Number
	divisor := int64(2)

	for number > 1 {
		if number%divisor == 0 {
			strm.Send(&pb.FactorResp{
				Result: divisor,
			})
			number /= divisor
		} else {
			divisor++
		}
	}

`cmd` : make run_server

`cmd` : make run_client

