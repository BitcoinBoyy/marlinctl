***Building***
go build -o build/marlinctl;
cd build;
./marlinctl help

***Usage***
First kill old supervisord process [ps -ef | grep supervisord; sudo kill -9 <pid>]
sudo supervisord -c ../supervisord.conf
sudo ./marlinctl help

sudo ./marlinctl beacon start --param1 n
