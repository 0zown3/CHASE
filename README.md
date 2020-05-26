# CHASE  
# Concurrent Harvester for Adversary Studies and Examination  

Capability
---------------
CHASE is a plugin for the TRAM tool developed by MITRE. TRAM leverages machine learning to detect
TTPs in threat reports and map them to adversaries.  

The main goal of CHASE is to quickly gather large amounts of relevant threat report URLs and feed them into TRAM. TRAM will then queue up those urls and begin its work. 

How it Works
---------------

To start things off, you simply need to send a POST request to localhost:7000 (when CHASE is running, or whichever port you choose to expose in a Docker/k8s environment).  

CHASE integrates the Feedly API. This allows for the aggregation of blogs from well known and reliable security blogs such as FireEye Threat Research and more. You'll need to get free
access token from Feedly by signing up (or if you really want to beef up your TRAM instance, opt for an enterprise token). However, the free token should be more than enough for daily
requests.  

You simply send a POST request with the following body:

{
    "token": "your_token"
}

From there, CHASE will incorporate your token in subsequent requests to the supported feeds (over HTTPS) in config/feeds.json. 

Usage
---------------
`cd tests\`  
`go test` # verify app passes all tests  
`cd cmd\chase`  
`go build`  
`./chase.exe`  

Send a POST to http://localhost:7000/ 

Impact
----------------

There is an immense backlog of threat reports that the ATT&CK team (and others) still need to analyze and map to ATT&CK to strengthen their detection mechanisms. CHASE will help alleviate the still manual process of finding relevant reports to
provide to TRAM by searching the Internet for them. 

The analyst can then predominantly focus on verifying TRAM's results and continuing to train its model. 


TODO
--------------
- Gather relevant URLs.
- Communicate with TRAM's REST API.
- Dockerize CHASE.