#Crowdcoin Go!  

This repo contains the porting of the Crowdcoin codebase written in c++
its written compiled with go 1.10 and following

#Phases of Development

1. Complete porting 
    1.a refactor and well organize the code
2. Rebuild the communication system creating a virtual network between the masternodes
3. Improve the tps laveraging the new changes
4. Build the new gui with electrum that will be able to accomodate
   4.a -3 to 5 step creation of a new masternode
   4.b creation of a masternode of another blockchain directly from the gui
   4.c -creation and add onchain of a document direclty in the gui wallet
5. intial tests for 0 connection transactions   
   