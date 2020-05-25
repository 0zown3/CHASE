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

How it Works
---------------

This POST request will kick off operations. CHASE will use that APT string as a key to fetch known aliases for that adversary from apt_mappings.json in /config.

From there, concurrent calls will be made to the main crawling function that will continuously feed TRAM's REST API with report URLs until it's finished. 

Impact
----------------

There is an immense backlog of threat reports that the ATT&CK team (and others) still need to analyze and map to ATT&CK to strengthen their detection mechanisms. CHASE will help alleviate the still manual process of finding relevant reports to
provide to TRAM by searching the Internet for them. 

The analyst can then predominantly focus on verifying TRAM's results and continuing to train its model. 


TODO
--------------
- Gather relevant URLs per specified APT and their aliases.
- Communicate with TRAM's REST API.
- Dockerize CHASE.