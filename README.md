# CHASE  
# Concurrent Harvester for Adversary Studies and Examination  

Capability
---------------
CHASE is a plugin for the TRAM tool developed by MITRE. TRAM leverages machine learning to detect
TTPs in threat reports and map them to ATT&CK.  

The main goal of CHASE is to quickly gather large amounts of relevant threat report URLs and feed them into TRAM. In doing so, we can quickly help
analysts strengthen TRAM's model and improve its capability.

How it Works
---------------

CHASE integrates the Feedly Cloud API. Feedly is a service that aggregates blogs based on your interest and provides an API that easily lets you access that data programmatically. We leverage this API to help us fetch relevant security reports and pass those titles and urls to TRAM.

### Feeds for PoC

These are the preloaded feeds in config/feeds.json. Of course, this JSON file is completely customizable based on what your needs are. Just keep in mind that the goal of CHASE is to feed TRAM with as many relevant reports as possible to help train the model.

- FireEye Threat Research
- Darknet
- Schneier on Security
- Sophos Labs

Usage
---------------

1. Get a developer access token to authenticate to the Feedly Cloud API (it's free).
2. Start TRAM.
3. Configure config/feeds.json with more relevant Feedly feeds.
4. Build CHASE. 
``` 
cd cmd\chase  
go build  
./chase.exe
```  
5. 
Send a POST to http://localhost:7000/ with the following body:  

```
{
    "token": "your_developer_token"
}
```

That's it! Now let TRAM do its thing!

Impact
----------------

There is an immense backlog of threat reports that the ATT&CK team (and others) still need to analyze and map to ATT&CK to strengthen their detection mechanisms. CHASE will help alleviate the still manual process of finding relevant reports to provide to TRAM by searching the Internet for them.  

CHASE also segments this process from TRAM, that way, TRAM can focus directly on effectively detecting TTPs rather than worry about where it gets these reports from as well. This plugin then makes it more practical to incorporate TRAM in a microservice architecture. It would then become trivial to automate interaction with CHASE so it consistently provides reports to TRAM.

TODO
--------------
- Dockerize CHASE.