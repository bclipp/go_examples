
go mod init data_ingestion

sudo apt-get remove -y gccgo-go && wget http://golang.org/dl/go1.13.linux-amd64.tar.gz && sudo apt-get -y install gcc && sudo tar -C /usr/local -xzf go1.8.2.linux-amd64.tar.gz && echo "export PATH=\$PATH:/usr/local/go/bin" >> ~/.bashrc

### building library
go install


