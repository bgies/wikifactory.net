# [Wikifactory.net](http://wikifactory.net)

This is the organisational website for Wikifactory Ltd.

## Getting Started

1.) Download and install the Google App Engine SDK for [Go] (https://cloud.google.com/appengine/downloads).

2.) Assuming you have [Python] (https://www.python.org/downloads/) 2.7.9 or 3.4 plus, install [Assetgen] (https://github.com/tav/assetgen) by running: 

    pip install https://pypi.python.org/packages/source/a/assetgen/assetgen-0.3.7.tar.gz

3.) Install freetype-go with: `go get code.google.com/p/freetype-go/freetype`.

4.) From the root directory:

- run `assetgen` to generate the js and css assets.

- initialise the development web server with `path/to/appengine/dev_appserver .`

5.) View the site at `http://localhost:8080/`.

NB: During development, you can run `assetgen --profile dev --watch` to have your assets regenerated automatically everytime you update the sources.
