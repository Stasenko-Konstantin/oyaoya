hello: main.go
	fyne-cross windows -arch=386
	cp fyne-cross/bin/windows-386/oyaoya.exe oyaoya.exe
	rm -rf fyne-cross
	rm Icon.png
