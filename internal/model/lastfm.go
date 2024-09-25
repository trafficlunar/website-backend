package model

type LastFMAPI struct {
	RecentTracks struct {
		TrackList []struct {
			Artist struct {
				Text string `json:"#text"`
			} `json:"artist"`
			Name  string `json:"name"`
			Image []struct {
				Text string `json:"#text"`
			} `json:"image"`
			Url        string `json:"url"`
			Attributes *struct {
				NowPlaying string `json:"nowplaying"`
			} `json:"@attr,omitempty"`
		} `json:"track"`
	} `json:"recenttracks"`
}

type LastFMData struct {
	Song    string `json:"song"`
	Artist  string `json:"artist"`
	Image   string `json:"image"`
	Url     string `json:"url"`
	Playing bool   `json:"playing"`
}
