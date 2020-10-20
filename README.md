<h2>Basic Architecture/Flow of x-tracer:</h2>

![alt text](https://sheenampathak.com/wp-content/uploads/2020/06/Screenshot-from-2020-06-10-13-48-07.png)

<b>x-tracer flow:</b>
1. x-tracer server is deployed on the master node
2. x-agent client deploys on the worker node(in which our target pod is running)
3. x-agent creation triggers a go module named ```probeparser```, which executes 7 different probes(ebpf tools)
4. 7 probes traces the logs of the target_pod's processes using namespace ID(as every process PID in container belongs to the same namespace ID) 
5. These generated logs from the probes are channelized to the x-tracer server in real time


<h2> Installation Steps: </h2>

Linux Kernel: 4.15.0-64-generic

<h3>1. Install Go Language</h3>

Currently tested with go1.14.2
 
<h3>2. Install BCC tools </h3>

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
