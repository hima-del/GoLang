**Create an instance**

* Create an AWS Account
* enter credit card info
* Sign into console

**Create a new EC2 instance**

* services / EC2
* launch instance
* choose your instance
* add storage / 30GB free
* add tags / webserver
* security / ssh / http
* launch
* create new key pair / download

**Deploy your binary**

* mv [src] [dst] / sudo chmod 400 your.pem

**Build hello world**

* GOOS=linux GOARCH=amd64 go build

**Copy your binary to the sever**

* scp -i /path/to/[your].pem ./main ec2-user@[public-DNS]:
* "ec2-user" might be "ubuntu" depending upon your machine
* say "yes" to The authenticity of host ... can't be established.

**SSH into your server**

* ssh -i /path/to/[your].pem ec2-user@[public-DNS]

**Run your code**

* sudo chmod 700 mybinary
* sudo ./mybinary
* check it in a browser at [public-IP]

**Exit**

* ctrl + c
* exit


**Persisting your application**

* To run our application after the terminal session has ended, we must do one of the following:

**Possible options**

* screen
* init.d
* upstart
* system.d

**System.d**

**Create a configuration file**

* cd /etc/systemd/system/
* sudo nano <filename>.service

```
[Unit]
Description=Go Server

[Service]
ExecStart=/home/<username>/<exepath>
User=root
Group=root
Restart=always

[Install]
WantedBy=multi-user.target
```

**Add the service to systemd**

* sudo systemctl enable <filename>.service

**Activate the service**

* sudo systemctl start <filename>.service

**Check if systemd started it**

* sudo systemctl status <filename>.service

**Stop systemd if so desired**

* sudo systemctl stop <filename>.service
