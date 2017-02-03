package himago

import "testing"

// Download a single tile that exists
func TestDownloadSingleTile(t *testing.T) {
	_, err := downloadTile("http://himawari8-dl.nict.go.jp/himawari8/img/D531106/2d/550/2017/02/03/191000_1_0.png")
	if err != nil {
		t.Error(err)
	}
}
