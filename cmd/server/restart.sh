git pull
make build
killall quick_srv
cd ../..
cat ./output/nohup.out >> ./output/nohup.out1
nohup ./output/quick_srv > ./output/nohup.out &