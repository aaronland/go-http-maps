debug-tagram:
	go run -mod vendor cmd/server/main.go -map-renderer tangramjs -nextzen-apikey $(APIKEY)

debug-protomaps:
	go run -mod vendor cmd/server/main.go -map-renderer protomaps -protomaps-tile-url https://static.sfomuseum.org/pmtiles/sfo.pmtiles
