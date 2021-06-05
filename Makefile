debug-tangram:
	go run -mod vendor cmd/server/main.go -map-provider tangramjs -nextzen-apikey $(APIKEY)

debug-protomaps:
	go run -mod vendor cmd/server/main.go -map-provider protomaps -protomaps-tile-url https://static.sfomuseum.org/pmtiles/sfo.pmtiles

debug-tilepack:
	go run -mod vendor cmd/server/main.go -map-provider tangramjs -tilezen-enable-tilepack -tilezen-tilepack-path /usr/local/data/sf.db -nextzen-tile-url http://localhost:8080/tilezen/vector/v1/512/all/{z}/{x}/{y}.mvt
