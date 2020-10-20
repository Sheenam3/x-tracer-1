x-tracer traces the process log inside a pod and stream the logs to the x-tracer server in real time.

x-tracer includes 7 ebp tools(BCC),probes to trace the process events:

1. Tcp connections: closed, active, established, life
2. Block device I/O 
3. New executed processes
4. Cache kernel function calls


<h2> Installation Steps: </h2>

Linux Kernel: 4.15.0-64-generic

<h3>1. Install Go Language</h3>

Currently tested with go1.14.2
 
<h3>2. Install BCC tools(on master node) </h3>

For Bionic(Ubuntu 18.04 LTS)

Install build dependencies


 ```
 sudo apt-get -y install bison build-essential cmake flex git libedit-dev\
 libllvm6.0 llvm-6.0-dev libclang-6.0-dev python zlib1g-dev libelf-dev
 ```

Fetch and compile BCC 

```
git clone https://github.com/iovisor/bcc.git
mkdir bcc/build 
cd bcc
git checkout v0.16.0
cd build
cmake ..
make
sudo make install
cmake -DPYTHON_CMD=python3 .. # build python3 binding
pushd src/python/
make
sudo make install
popd

``` 
<h3>3. Get and run x-tracer</h3>

```
go get github.com/ITRI-ICL-Peregrine/x-tracer
cd $GOPATH/src/github.com/ITRI-ICL-Peregrine/x-tracer
go build x-tracer.go
./x-tracer

```
