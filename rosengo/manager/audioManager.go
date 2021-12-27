package manager

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"path"
	"path/filepath"
	"strings"

	"github.com/MrWormHole/rosengo/rosengo/assets"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/mp3"
	"github.com/hajimehoshi/ebiten/v2/audio/vorbis"
	"github.com/hajimehoshi/ebiten/v2/audio/wav"
)

type AudioManager struct {
	audioPlayers map[string]*audio.Player
	audioContext *audio.Context
	sampleRate   int
	isMuted      bool
}

func NewAudioManager(sampleRate int) (*AudioManager, error) {
	return &AudioManager{
		audioPlayers: make(map[string]*audio.Player),
		audioContext: audio.NewContext(sampleRate),
		sampleRate:   sampleRate,
		isMuted:      false,
	}, nil
}

func (m *AudioManager) LoadAll(dir string) error {
	if strings.TrimSpace(dir) == "" {
		return errors.New("AudioManager.Load: directory name can not be blank")
	}

	sounds, err := assets.Bundle.ReadDir(dir)
	if err != nil {
		return fmt.Errorf("AudioManager.LoadAll: failed to read sounds directory: %v", err)
	}

	for i := range sounds {
		name := sounds[i].Name()
		extension := filepath.Ext(name)
		if extension != ".ogg" && extension != ".wav" && extension != ".mp3" {
			continue
		}

		rawFile, err := assets.Bundle.ReadFile(path.Join(dir, name))
		if err != nil {
			return fmt.Errorf("AudioManager.LoadAll: failed to read sound file: %v", err)
		}

		var src io.ReadSeeker
		switch extension {
		case ".ogg":
			stream, err := vorbis.DecodeWithSampleRate(m.sampleRate, bytes.NewReader(rawFile))
			if err != nil {
				return fmt.Errorf("AudioManager.LoadAll: failed to decode ogg: %v", err)
			}
			src = stream
		case ".wav":
			stream, err := wav.DecodeWithSampleRate(m.sampleRate, bytes.NewReader(rawFile))
			if err != nil {
				return fmt.Errorf("AudioManager.LoadAll: failed to decode wav: %v", err)
			}
			src = stream
		case ".mp3":
			stream, err := mp3.DecodeWithSampleRate(m.sampleRate, bytes.NewReader(rawFile))
			if err != nil {
				return fmt.Errorf("AudioManager.LoadAll: failed to decode mp3: %v", err)
			}
			src = stream
		}

		audioPlayer, err := m.audioContext.NewPlayer(src)
		if err != nil {
			return fmt.Errorf("AudioManager.LoadAll: failed to create a new audio player: %v", err)
		}

		m.audioPlayers[name] = audioPlayer
	}
	return nil
}

func (m *AudioManager) Close() error {
	for key, p := range m.audioPlayers {
		if err := p.Close(); err != nil {
			return fmt.Errorf("AudioManager.CloseAll: failed to close an audio player: %v", err)
		}
		delete(m.audioPlayers, key)
	}
	return nil
}

func (m *AudioManager) checkAudioPlayer(audio string) error {
	_, ok := m.audioPlayers[audio]
	if !ok {
		return fmt.Errorf("AudioManager.checkAudioPlayer: audio player is missing for the given audio name %s", audio)
	}
	return nil
}

func (m *AudioManager) IsPlaying(audio string) bool {
	if err := m.checkAudioPlayer(audio); err != nil {
		return false
	}
	return m.audioPlayers[audio].IsPlaying()
}

func (m *AudioManager) SetVolume(audio string, volume float64) error {
	if m.isMuted {
		return nil
	}

	if err := m.checkAudioPlayer(audio); err != nil {
		return fmt.Errorf("AudioManager.SetVolume: %v", err)
	}
	m.audioPlayers[audio].SetVolume(volume)
	return nil
}

func (m *AudioManager) GetVolume(audio string) float64 {
	if m.isMuted {
		return 0
	}

	if err := m.checkAudioPlayer(audio); err != nil {
		return 0
	}
	return m.audioPlayers[audio].Volume()
}

func (m *AudioManager) Pause(audio string) error {
	if m.isMuted {
		return nil
	}

	if err := m.checkAudioPlayer(audio); err != nil {
		return fmt.Errorf("AudioManager.Pause: %v", err)
	}
	m.audioPlayers[audio].Pause()
	return nil
}

func (m *AudioManager) Play(audio string) error {
	if m.isMuted {
		return nil
	}

	if err := m.checkAudioPlayer(audio); err != nil {
		return fmt.Errorf("AudioManager.Play: %v", err)
	}
	m.audioPlayers[audio].Play()
	return nil
}

func (m *AudioManager) Replay(audio string) error {
	if m.isMuted {
		return nil
	}

	if err := m.checkAudioPlayer(audio); err != nil {
		return fmt.Errorf("AudioManager.Replay: %v", err)
	}
	if err := m.audioPlayers[audio].Rewind(); err != nil {
		return fmt.Errorf("AudioManager.Replay: failed to rewind audio player: %v", err)
	}
	m.audioPlayers[audio].Play()
	return nil
}
