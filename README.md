Temporary experiment for getting vendoring to work together with Appengine SDK.
The vendored files were populated with `godep save .`.

The following examples fails in different ways:

	gcloud app deploy app.yaml
	goapp deploy -application foobar -version 1 app.yaml
	appcfg.py update -A foobar -V 1 app.yaml
	GOPATH= appcfg.py update -A foobar -V 1 app.yaml
	GOPATH= goapp deploy -application foobar -version 1 app.yaml

https://code.google.com/p/googleappengine/issues/detail?id=13353

Versions used at the time of testing:

Google Cloud SDK 130.0.0
app-engine-go-darwin-x86_64 20160927
app-engine-python 1.9.40