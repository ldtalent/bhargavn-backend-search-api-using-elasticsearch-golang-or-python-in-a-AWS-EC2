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
