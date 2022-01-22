# fivem-audit

## About
Logs certain events to another webserver for display outside of FiveM. Can be used for monitoring/graphs/statistics

## Setup
#### FiveM
1. Add the files in [audit](https://github.com/Jepzter/fivem-audit/tree/main/audit) to /txData/CFXDefault_DAFCFA.base/resources/[insert_folder_here]/audit
2. Modify the host:port if changed
3. Open `server.cfg` and add the line `ensure audit` under the resources section
4. Start your FiveM server

#### Webservice
1. Build the go application by using `go build *.go` in the root directory of the webservice (skip this step if you downloaded the release)
2. Start the webservice by executing the binary `./server.go -port=8090`

#### UI 
1. Run `npm install`
2. Run `npm run serve` or `npm run build` if you have a webserver
3. Access the ui by going to http://localhost:8080
