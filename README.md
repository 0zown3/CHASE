# CHASE  
# Concurrent Harvester for Adversary Studies and Examination  

Capability
---------------
CHASE is a plugin for the TRAM tool developed by MITRE. TRAM leverages machine learning to detect
TTPs in threat reports and map them to adversaries.  

The main goal of CHASE is to quickly gather large amounts of relevant report URLs for a specificed adversary to feed that array into TRAM. TRAM will then queue up those urls and begin its work. CHASE
will aid in training TRAMs model and also help analysts gather relevant reports on specific groups.

Usage
---------------
`cd tests\`  
`go test` # verify app passes all tests  
`cd cmd\chase`  
`go build`  
`./chase.exe`  

Send a POST to http://localhost:7000/ with the following body:  
```
{
    "APT": string
}
```
The string in the APT field MUST correspond to an ATT&CK mapping (i.e. APT28).  
See: https://attack.mitre.org/groups/  


TODO
--------------
- Gather relevant URLs per specified APT and their aliases.
- Communicate with TRAM instance and send array to endpoint.
- Dockerize CHASE.