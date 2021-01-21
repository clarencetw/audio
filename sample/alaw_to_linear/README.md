# alaw to linear sample

1. download sample:
   `wget http://www.music.helsinki.fi/tmt/opetus/uusmedia/esim/a2002011001-e02-8kHz.wav`

2. wav to pcm:
   `ffmpeg -i a2002011001-e02-8kHz.wav -ar 8000 -ac 1 -ab 64 -f alaw output.g711`
3. run sample:
   `go run main.go`
