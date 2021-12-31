package manager

import (
	"bytes"
	"errors"
	"fmt"
	"image"
	_ "image/png"
	"path"
	"path/filepath"
	"strings"

	"github.com/MrWormHole/rosengo/rosengo/assets"
	"github.com/hajimehoshi/ebiten/v2"
)

type SpriteManager struct {
	images map[string]*ebiten.Image
}

func NewSpriteManager() (*SpriteManager, error) {
	images := make(map[string]*ebiten.Image)
	return &SpriteManager{
		images: images,
	}, nil
}

func (m *SpriteManager) LoadAll(dir string) error {
	if strings.TrimSpace(dir) == "" {
		return errors.New("SpriteManager.Load: directory name can not be blank")
	}

	images, err := assets.Bundle.ReadDir(dir)
	if err != nil {
		return fmt.Errorf("SpriteManager.LoadAll: failed to read images directory: %v", err)
	}

	for i := range images {
		name := images[i].Name()
		extension := filepath.Ext(name)
		if extension != ".png" {
			continue
		}

		rawFile, err := assets.Bundle.ReadFile(path.Join(dir, name))
		if err != nil {
			return fmt.Errorf("SpriteManager.LoadAll: failed to read image file: %v", err)
		}

		img, _, err := image.Decode(bytes.NewReader(rawFile))
		if err != nil {
			return fmt.Errorf("SpriteManager.LoadAll: failed to decode image: %v", err)
		}

		key := name[:len(name)-len(extension)]
		m.images[key] = ebiten.NewImageFromImage(img)
	}
	return nil
}

func (m *SpriteManager) Dispose() {
	for key, i := range m.images {
		i.Dispose()
		delete(m.images, key)
	}
}

func (m *SpriteManager) GetImage(name string) (*ebiten.Image, error) {
	if strings.TrimSpace(name) == "" {
		return nil, errors.New("SpriteManager.GetImage: name can not be blank")
	}

	if img, ok := m.images[name]; ok {
		return img, nil
	}

	return nil, fmt.Errorf("SpriteManager.GetImage: image not found: %v", name)
}
