# mylittleethorse_server
+++++INSTALL

+++METHOD1 :
+ you need golang 1.7+

sudo apt-get install golang

+ you need to set GOPATH

export GOPATH=$HOME/go

+ get golang deps 

go get "github.com/ethereum/go-ethereum"
go get "github.com/gorilla/mux"

+++METHOD2 :
chmod u+x install.sh
sudo install.sh

+++++BUILD

go build -o mle --tags linux


+++++RUN

./mle


+++++BUILD + RUN

./run.sh
