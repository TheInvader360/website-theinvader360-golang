#ABOUT

[http://www.theinvader360.com](http://www.theinvader360.com)

A simple website to advertise my apps. Also somewhere to host auxiliary pages (about, privacy, ads.txt, etc).

Must be easy to maintain. Does not need a dynamic DB. Data could be stored in flat files for ease of management and versioning.

Should offer some of the benefits of a static site (content in version control, easy management, some security) and some of the benefits of a dynamic site (templating, server side logic).

Most importantly - free to host on GAE and a good excuse to dabble in a bit of golang :)

#INITIAL SETUP

[Install cloud sdk](https://cloud.google.com/sdk/install)

[Install go 1.12.9](https://golang.org/doc/install)

    gcloud init

Select "*create a new project...*"

    website-theinvader360-golang
    gcloud app create

Select "*europe-west2*"

[Enable cloudbuild](https://console.developers.google.com/apis/api/cloudbuild.googleapis.com/overview?project=website-theinvader360-golang)

    mkdir -p ~/go/src/website-theinvader360-golang/
    cd ~/go/src/website-theinvader360-golang/

* Create app.yaml
* Create main.go

Deploy to cloud:

    cd ~/go/src/website-theinvader360-golang
    gcloud app deploy
    gcloud app browse

#UPDATES

Update app as required...

Run locally:

    go build && ./website-theinvader360-golang

Test locally: [http://localhost:8080/](http://localhost:8080/)

Deploy to cloud:

    gcloud app deploy
    gcloud app browse
