---
title: "Music Player"
date: 2025-08-29
draft: false
---

# Music Player Widget

Plays local audio files with optional visualizer effects.

---

## Example Configuration

```yaml
music:
  enabled: true
  position: "bottom-left"
  local:
    enabled: true
    path: "music/"
    path_list:
      - "music/song1.mp3"
      - "music/song2.mp3"
      - "music/song3.mp3"
    shuffle: true
  visualizer:
    enabled: true
    type: "bars"
    size: 50
    position: "bottom"
```

### **Properties**

- **local.enabled**: Use local files (true/false)
- **path / path_list**: Files or folders to play
- **shuffle**: Play order (true/false)
- **visualizer.enabled**: Show animated visualizer
- **visualizer.type**: `"bars"`, `"wave"`, `"circle"`, `"dots"`
- **visualizer.size**: Scale/intensity of visualizer
- **visualizer.position**: Where visualizer appears

---

## Usage Tips

- Ensure music files exist at specified paths.
- Shuffle and visualizer options enhance user experience.
- Keep the player position consistent with other UI elements.
