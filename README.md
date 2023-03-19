# How to build a backend API with elasticsearch and deploy it into the AWS EC2 using Golang
## Overview:
In this tutorial you will get an idea of how we can use elasticsearch for building API and might give you an idea that how you could integrate it with your other backend Apis (other backend apis means apis that you build to serve your other task based on your problem statement). The pre-requisites for this tutorial is a little programming knowledge on Golang with a basics of API. 
You will get to learn the following topics
1. How to configure AWS EC2 instance.
2. Basics of elasticsearch.
3. How to install and configure elasticsearch.
4. Installing Golang and building a simple search backend API in Golang.

## How to configure AWS EC2 instance:
1. Sign Up/Log in to AWS.
2. Go to the EC2 from the navigation panel and click launch instance.
3. Select region 
4. Give server name
5. create key and download the key file. You can add custom security key with port number according to your need in my case that is 1900. Also default tick SSH client support with port 22.
6.Select storage size
7.Locate private key file which you have stored.In my case filename.pem , give necessary permission and run chmod 400 bnathkey.pem

## Connect your remote server from your pc using ssh command
Move to the folder where you download the pem key file.
Then open terminal on that folder and run the following command
```
ssh -i path/directory/to/the/filename.pem username@ipaddress
```
for example in my case ssh -i e:\bnathkey.pem ubuntu@18.223.29.3

## Install Golang
To install Go in EC2 run the following command in AWS EC2 linux terminal.
```
sudo apt update
sudo apt upgrade
sudo apt search golang-go
sudo apt search gccgo-go
sudo apt install golang-go
```
## Install Elasticsearch
Run the following command in ec2 terminal
```
sudo apt-get install default-jre
https://artifacts.elastic.co/GPG-KEY-elasticsearch | sudo apt-key add -
wget -qO - https://artifacts.elastic.co/GPG-KEY-elasticsearch | sudo apt-key add -

wget https://artifacts.elastic.co/downloads/elasticsearch/elasticsearch-5.3.0.deb
sudo update-rc.d elasticsearch defaults 95 10
sudo /etc/init.d/elasticsearch start
```
Since elasticsearch takes huge amount of RAM you can reduce the memory allocation size.
As for the demonstration purpose I have use EC2 free tier instance so I have to reduce that using following command
```
sudo nano /etc/elasticsearch/jvm.options
```
Find the line
```
-Xms1g
-Xmx1g
```
And reduce it according to your available resources. In my case that would be
```
-Xms312m
-Xmx312m
```
Then run 
``` curl -XGET 'http://localhost:9200'```
That will start the elasticsearch server on port 9200
Now you have to put some data on elasticsearch.

#### Storing csv data to elasticsearch
Since I couldnot use Logstash due limitation of resources, I used curl to store data in Elasticsearch.
So, There is a tricky way I found that will do my work flawlessly.
1. At first convert csv data to json format. (using your coding skill or any online converter tools from internet ;) )
2. Push it to remote server (i.e, EC2 server) using git or filezilla like application.
3. Though your data is on json format, yet you can not push your data to elasticsearch. So first you need to convert json format to elasticsearch acceptable format. This can be done by running the following command
```sudo jq -c -r ".[]" input.json | while read line; do echo '{"index":{}}'; echo $line; done > output.json``` Where input.json is my json file name and  output.json is my elasicsearch compatible json.
The above is the jq command and work on linux (As in linux Jq is installed by default but if your ststem doesn't have jq command than just install it and then run the command).
4. Then upload the data using following command 
```curl -XPOST localhost:9200/students/your_type/_bulk -H "Content-Type: application/x-ndjson" --data-binary @output.json```
here students is my index name.

## Building a simple search API 
Run the main.go file with the following command
```
sudo go mod init foldername
sudo go mod tidy
sudo go run main.go // where main is my file name 
```
### So what have I done in main.go file?
I have students -sheets(2).csv which is a small datasets of student and that needs to be uploaded in elasticsearch. Now I have to make an API that serve data based on the search request, say I need the student data whose firstname is John so my search api return the result of all the students whose firstname is John. See the API respose of my search request as shown bellow.
<img title="" alt="Alt text" src="/images/boo.svg">
