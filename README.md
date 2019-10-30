# stream_learn
rpc stream


protoc --go_out=output_directory input_directory/file.proto

pprof 文件分析：
go tool pprof -svg cpu.pprof > cpu.svg
依赖于 graphviz 工具