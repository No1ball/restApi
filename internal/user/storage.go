package user

import (
	"errors"
	"github.com/mowshon/moviego"
	"os"
)

// videos are stored in map, where key is a unique session id;
// one session may have only one instance of videos,
// it means that no change can be undone
var videos map[int]*moviego.Video

func init() {
	videos = make(map[int]*moviego.Video)
}

// GetVideo returns videos stored during session with specified sessionId
// or error in case of its absence
func getVideo(videoInfo VideoInfo) (*moviego.Video, error) {
	res, ok := videos[videoInfo.id]

	if ok {
		return res, nil
	} else {
		return nil, errors.New("no videos for given session")
	}
}

// set videos with specified id;
// call it only to replace existing videos with changed;
// if you need to create new videos, call loadVideo()
func setVideo(videoInfo VideoInfo, newVideo *moviego.Video) error {
	_, ok := videos[videoInfo.id]

	if ok {
		videos[videoInfo.id].Output(videoInfo.path)
		videos[videoInfo.id] = newVideo
		return nil
	} else {
		return errors.New("attempt to modify unknown videos")
	}
}

// uploads existing video to RAM
// this function DOES NOT upload video from client
func loadVideo(videoInfo VideoInfo) error {
	_, ok := videos[videoInfo.id]

	if !ok {
		video, error := moviego.Load(videoInfo.path)

		if error != nil {
			return errors.New("lack of memory")
		}

		videos[videoInfo.id] = &video
		return nil

	} else {
		return errors.New("video is in RAM already")
	}
}

// removes video both from RAM and server
func deleteVideo(videoInfo VideoInfo) error {
	_, ok := videos[videoInfo.id]

	if ok {
		delete(videos, videoInfo.id)
		os.Remove(videoInfo.path)
		return nil
	} else {
		return errors.New("attempt to remove absent video")
	}
}
