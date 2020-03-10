# 伪造
.PHONY: yggdrasill master slave

yggdrasill: master slave 

master:
		cd master && protoc *.proto --go_out=plugins=grpc:.

slave:
		cd slave && protoc *.proto --go_out=plugins=grpc:.