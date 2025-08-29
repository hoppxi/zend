---
title: "Zend Configuration File"
date: 2025-08-29
draft: false
---

# Zend Configuration File

This file configures the **appearance, behavior, and widgets** of Zend. It includes settings for UI, colors, fonts, backgrounds, widgets (clock, weather, music), and transitions.

> **Note:** All keys are optional unless marked **REQUIRED**. Unspecified options will fall back to default values.

---

## Common Properties Reference

Some properties are reused across multiple sections:

### Positioning of Elements

- Used by: clock, weather, music, visualizer
- Determines element location on screen
- Supported values:
  - Predefined keywords: `"top"`, `"bottom"`, `"left"`, `"right"`, `"top-left"`, `"top-right"`, `"bottom-left"`, `"bottom-right"`, `"center"`
  - Custom coordinates: `"(x%, y%)"`
    - Example: `"(50%, 50%)"` = center

```yaml
position: "top-right"
position: "(30%, 75%)"
```

---

### File and Directory Paths

- Used by: image backgrounds, local music, icons
- Supports absolute paths or lists:

```yaml
path: "/home/user/images/bg.png"
path_list:
  - "/home/user/music/song1.mp3"
  - "/home/user/music/song2.mp3"
```

---

### Colors

- Formats: HEX `"#rrggbb"` or `"#rrggbbaa"`

```yaml
color: "#ff0000"
```

---

### Time Formatting (Clock & Time-Based Features)

- Uses Day.js tokens:

```yaml
time_format: "HH:mm:ss"  # 24-hour with seconds
time_format: "hh:mm A"   # 12-hour AM/PM
time_format: "dddd, MMMM D"  # Full weekday & month
```

---

## Full Example

```yaml
general:
  blur: 5
  opacity: 1
  transition:
    type: "fade"
    duration: 1
    position: "center"

font:
  family: "Arial, sans-serif"
  size: 14
  weight: "normal"
  style: "normal"

palette:
  primary: "#95ccff"
  on-primary: "#003352"
  secondary: "#b9c8da"
  on-secondary: "#233240"

background:
  solid:
    enabled: true
    color: "#ff0000"
    use_as_accent: true
  random:
    enabled: true
    use: "color"
    interval: 10
    max: 100

clock:
  enabled: true
  format:
    time: "HH:mm:ss"
    date: "YYYY-MM-DD"
  timezone: "UTC"
  position: "top-right"

weather:
  enabled: true
  location: "New York,US"
  units: "metric"
  api_key: "your_openweathermap_api_key"
  position: "top-left"
```
