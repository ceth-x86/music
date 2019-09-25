package engine

type DownloadResult struct {
	Downloaded bool
	Album      uint
	Single     uint
}

func DownloadTrackError() *DownloadResult {
	return &DownloadResult{
		Downloaded: false,
		Album:      0,
		Single:     0,
	}
}
