***Building***
go build -o build/marlin-cli; 
cd build; 
./marlin-cli help

***Usage***
First kill old supervisord process [ps -ef | grep supervisord; sudo kill -9 <pid>]
sudo supervisord -c ../supervisord.conf
sudo ./marlin-cli help

sudo ./marlin-cli beacon start --param1 n