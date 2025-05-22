Projects deals with capturing data/logs from vechicles/truck over LTE/LoRa protocol to datastrem and eventually to datawarehouse. This data eventually feeds into AI pipeline which predicts Remaining Useful Life(RUL) of vehicles, which help in making decisions to repair/replace/modify parts. 

Main components of the projects: 
1. Edge-agent - To collect CAN/1553 bus logs from vehicles. 
2. Data-streaming - Collected logs are feed into data-stream(kafka). 
3. Cloud/Server - To process and save log feed coming from agent. Also helps in creating dataset to train ML models.
4. ML Model - To train the AI/ML model to predict Remaining Useful Life(RUL) for vehicle.
5. UI/Dashboard - To show which parts/components required to replaced on basis of inference from ML model.