flash: flash-tinygo

prepare-gopherconeu:
	go run cmd/main.go -conf=gopherconeu

flash-gopherconeu: prepare-gopherconeu perform-flash

prepare-gopherconie:
	go run cmd/main.go -conf=gopherconie

flash-gopherconie: prepare-gopherconie perform-flash

prepare-gopherconuk:
	go run cmd/main.go -conf=gopherconuk

flash-gopherconuk: prepare-gopherconuk perform-flash

prepare-gopherconus:
	go run cmd/main.go -conf=gopherconus

flash-gopherconus: prepare-gopherconus perform-flash

prepare-tinygo:
	go run cmd/main.go -conf=tinygo

flash-tinygo: prepare-tinygo perform-flash

prepare-golab:
	go run cmd/main.go -conf=golab

flash-golab: prepare-golab perform-flash

perform-flash:
	tinygo flash -size short -target gopher-badge -ldflags="-X main.YourName='$(NAME)' \
	-X main.YourTitleA1='$(TITLEA1)' -X main.YourTitleA2='$(TITLEA2)'  -X main.YourTitleB1='$(TITLEB1)' \
	 -X main.YourTitleB2='$(TITLEB2)' -X main.YourMarqueeTop='$(MARQUEETOP)' -X main.YourMarqueeMiddle='$(MARQUEEMIDDLE)' \
	 -X main.YourMarqueeBottom='$(MARQUEEBOTTOM)' -X main.YourQRText='$(QRTEXT)'" .
