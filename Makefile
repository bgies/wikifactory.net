# Public Domain (-) 2015 The Wikifactory.net Website Authors.
# See the Wikifactory.net UNLICENSE file for details.

site: sync
	@assetgen
	@go run gensite.go

clean:
	@assetgen --clean
	@rm -rf www
