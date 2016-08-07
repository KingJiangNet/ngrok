
##Build if you prefer
If you want to build binaries from scratch, please follow the following steps.

*yum, apt_get or others*
- Install prerequisites

```
sudo yum update
sudo yum install build-essential golang mercurial git
```

- Downlaod source
```
git clone https://github.com/inconshreveable/ngrok.git ngrok
cd ngrok
```

- Update git
```
git requirement: >= 1.7.9.5
yum --enablerepo=rpmforge-extras install git
```

- Insall golang
```
wget https://storage.googleapis.com/golang/go1.6.3.linux-386.tar.gz
tar -xzvf go1.6.3.linux-386.tar.gz
sudo mv go /usr/local/
```

- Add golang to path
```
nano /etc/profile
export PATH=$PATH:/usr/local/go/bin
source /etc/profile
```

- Build
*Build any platform you like*
```
GOOS=linux GOARCH=386 make release-server 
GOOS=linux GOARCH=386 make release-client 

GOOS=linux GOARCH=amd64 make release-server 
GOOS=linux GOARCH=amd64 make release-client

#Raspberry Pi and its sisters
GOOS=linux GOARCH=arm make release-server
GOOS=linux GOARCH=arm make release-client

GOOS=windows GOARCH=386 make release-server
GOOS=windows GOARCH=386 make release-client

GOOS=windows GOARCH=amd64 make release-server
GOOS=windows GOARCH=amd64 make release-client

#Mac OS
GOOS=darwin GOARCH=386 make release-server
GOOS=darwin GOARCH=386 make release-client

GOOS=darwin GOARCH=amd64 make release-server
GOOS=darwin GOARCH=amd64 make release-client
```

Now the build is completed.

## Start server

- SSL Certificate
If tou don't want to get a SSL certificate from a public CA, use these commands to create a self-signed certificate
Actually there are some free SSL certificate vendors which are much better than singed by self

```
openssl genrsa -out ca.key 2048
openssl req -new -x509 -nodes -key ca.key -days 10000 -subj "/CN=yourserver.com" -out cacert.pem
openssl genrsa -out server.key 2048
openssl req -new -key server.key -subj "/CN=yourserver.com" -out server.csr
openssl x509 -req -in server.csr -CA cacert.pem -CAkey ca.key -CAcreateserial -days 10000 -out server.crt
```

- Launch parameters:
```
./bin/ngrokd -tlsKey=server.key -tlsCrt=server.crt -domain="yourserver.com" -httpAddr=":83" -httpsAddr=":446" -tunnelAddr=":930"
```

## Start client
- Config
```
sudo nano ~/.ngrok
```

```yaml
server_addr: yourserver.com:930
trust_host_root_certs: true
useInsecureSkipVerify: true
#CACrtPath: "/cacert.pem" # Set When useInsecureSkipVerify is false for higher security
tunnels:
web:
    #auth: "AuthUser:AuthPassWord"
    proto:
	http: 80
ssh:
    remote_port: 2222
    proto:
	tcp: 22
```

- Add DNS record
e.g.
abc.yourserver.com -> your server's public IP address
yourserver.com -> your server's public IP address

- Test
While DNS record takes effect. open "http://abc.yourserver.com:83" in your browser.
