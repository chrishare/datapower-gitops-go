# datapower-gitops-go

Go rewrite of datapower tooling I'd formerly written in python.

This utility performs datapower maintenance activities (deployments, validations, upgrades) in a gitops-styled fashion, with some standalone CLI functionality too.

### Version history

0.01 - Create a log entry
0.02 - 
### TODO

* Add tocfg method for all times, to convert from the standard yaml/json/soap format to text cfg
* Add reboot script

### Commands

* Start a docker server for testing

docker run -it  -v $PWD/config:/opt/ibm/datapower/drouter/config  -v $PWD/local:/opt/ibm/datapower/drouter/local  -e DATAPOWER_ACCEPT_LICENSE=true  -e DATAPOWER_INTERACTIVE=true  -p 9090:9090  -p 9022:22  -p 5554:5554  -p 8000-8010:8000-8010  --name idg  icr.io/cpopen/datapower/datapower-limited:10.0.1.5

