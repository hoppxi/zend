---
title: "Background Settings"
date: 2025-08-29
draft: false
---

# Background Settings

Zend supports multiple ways to configure the homepage background, including **random backgrounds**, **solid colors**, and **custom images**.

---

## Random Background

```yaml
random:
  enabled: true
  use: "color" # "color" for solid colors, "image" for random photos
  interval: 10 # minutes between changes
  max: 100 # number of options (colors or images)
```

- **enabled**: Toggle random background cycling on/off (default: true)
- **use**: Type of random content (`color` or `image`)
- **interval**: Time interval in minutes for background change
- **max**: Maximum number of random options

> If both `random` and `image` are configured, `image` takes priority.

---

## Solid Color Background

```yaml
solid:
  enabled: true
  color: "#ff0000" # HEX code
  use_as_accent: true # Generate UI palette based on this color
```

- **enabled**: Enable fixed color background
- **color**: HEX color value
- **use_as_accent**: Whether to use this color for accenting the UI

---

## Custom Image Background

```yaml
image:
  path: "images/background.png"
  # path_list:
  #   - "images/image1.png"
  #   - "images/image2.jpg"
```

- **path**: Single fixed image
- **path_list**: Optional list of images for a slideshow
- Supports **absolute or relative paths**

---

## Image Rendering and Fitting

```yaml
resize:
  mode: "cover" # "cover", "contain", or any CSS background-size value
  position: "center" # CSS background-position
  repeat: "no-repeat" # "no-repeat", "repeat", "repeat-x", "repeat-y"
```

- Controls how background images are displayed.

---

## Background Refresh Control

```yaml
refresh:
  enabled: false
  interval: 60 # seconds
  every_restart: false
```

- **enabled**: Auto-refresh backgrounds
- **interval**: Time in seconds between refreshes
- **every_restart**: Pick a fresh image each restart
