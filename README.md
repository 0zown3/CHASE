# CHASE  
# Concurrent Harvester for Adversary Studies and Examination  

CHASE aims to be a plugin for TRAM that continuously fetches threat actor reports from known and trusted security research orgs by identifying APTX and aliases in article titles.

CHASE runs, by default, on port 7000 exposing a single endpoint / which only accepts a POST request with the body:
{"dtg": str, "apt": str}