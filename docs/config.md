# Zend Configuration File (Full Reference)

This file configures the **appearance, behavior, and widgets** of Zend. It includes everything from UI backgrounds, color themes, fonts, widgets (clock, weather, music), to animations and transitions.

> **Note:** All keys are optional unless marked **REQUIRED**. Any unspecified option will fall back to its default value.

---

# Common Properties Reference

These properties are reused in multiple sections below.
Refer to them here to avoid repetitive explanations.

---

## Positioning of Elements on Screen

- Used by: clock, weather, music, visualizer, etc.
- Determines where the element will appear on the screen.
- Supported values:

  - Predefined keywords (anchors to screen edges/corners):

    - `"top"`, `"bottom"`, `"left"`, `"right"`
    - `"top-left"`, `"top-right"`, `"bottom-left"`, `"bottom-right"`
    - `"center"`

  - Custom coordinates as `"(x%, y%)"`:

    - `"(50%, 50%)"` = exact center
    - `"(10%, 90%)"` = near bottom-left

- Tips:

  - Use keywords for quick alignment.
  - Use percentages for precise positioning in complex layouts.

- Example:

```yaml
position: "top-right"
position: "(30%, 75%)"
```

- Usage:

```yaml
position: "<keyword or (x%, y%)>"
```

---

## File and Directory Paths

- Used by: image backgrounds, local music, icons
- Types:

  - Absolute: starts with `/` → full system path

- Examples:

```yaml
path: "/home/user/images/bg.png"
```

- Usage:

```yaml
path: "/home/user/music/"
path_list:
  - "/home/user/music/song1.mp3"
  - "/home/user/music/song2.mp3"
```

---

## Colors

- Used by: solid backgrounds, palettes, UI accenting
- Formats:

  - HEX: `"#rrggbb"` (e.g., `"#ff0000"` = red)
  - HEX with alpha: `"#rrggbbaa"` (e.g., `"#ff000080"` = semi-transparent red)

- Usage:

```yaml
color: "<hex|named|gradient>"
```

---

## Time Formatting (Clock & Time-Based Features)

- Uses Day.js format tokens.
- Examples:

  - `"HH:mm:ss"` → 24-hour format with seconds (default)
  - `"hh:mm A"` → 12-hour format with AM/PM
  - `"dddd, MMMM D"` → Full weekday and month (e.g., Wednesday, August 27)

- Usage:

```yaml
time_format: "<Day.js format>"
```

---

# Frontend Build Configuration

- Location of the compiled frontend (HTML/CSS/JS)
- Useful if you are using a custom-built frontend or deploying on a specific environment
- Examples:

```yaml
dist: "/home/user/myproject/frontend/dist"  # Absolute path
dist: "./dist"                              # Relative path
```

- Optional. If omitted, Zend will serve its built-in frontend.

```yaml
dist: "/home/hoppxi/devspace/zend/frontend/dist"
```

---

# Random Background Settings

```yaml
random:
  enabled: true
  use: "color"
  interval: 10
  max: 100
```

- **enabled**: Enable random background cycling (color or image).
  When enabled, Zend automatically picks a random color or fetches a random image from Lorem Picsum.
  Optional. Default: `true`.
- **use**: Type of random content: `"color"` → solid colors, `"image"` → random photos.
  Optional. Default: `"color"`.
- **interval**: How often to change the background in minutes. Default: `10`.
- **max**: Number of random options to keep in the pool:

  - Colors → up to 100 distinct colors
  - Images → up to 30 images
    Optional. Defaults: 100 (colors), 30 (images).

---

# Custom Image Background

```yaml
image:
  path: "images/background.png"
  # path_list:
  #   - "images/image1.png"
  #   - "images/image2.jpg"
```

- Use a fixed image instead of random background.
- Paths can be absolute or relative.
- `path_list` allows a folder of images for slideshow-like backgrounds.
- If both `random` and `image` are set, `image` takes priority.

---

# Image Rendering and Fitting

```yaml
resize:
  mode: "cover"
  position: "center"
  repeat: "no-repeat"
```

- **mode**: `"cover"` (fill), `"contain"` (fit without cropping), or any CSS background-size value.
- **position**: CSS background-position (keywords or percentages)
- **repeat**: `"no-repeat"`, `"repeat"`, `"repeat-x"`, `"repeat-y"`

---

# Background Refresh Control

```yaml
refresh:
  enabled: false
  interval: 60
  every_restart: false
```

- **enabled**: Auto-refresh background images from folder.
- **interval**: Time in seconds (default 60)
- **every_restart**: Pick a fresh image each time Zend restarts.

---

# Solid Color Background

```yaml
solid:
  enabled: true
  color: "#ff0000"
  use_as_accent: true
```

- Fixed color background, overrides random/image if enabled.
- `use_as_accent`: generate UI palette based on color.

---

# General UI Configuration

```yaml
general:
  blur: 5
  opacity: 1
  transition:
    type: "fade"
    duration: 1
    position: "center"
```

- **blur**: Background blur in pixels
- **opacity**: 0–1
- **transition**:

  - **type**: `"fade"`, `"slide"`, `"zoom"`, `"grow"`
  - **duration**: seconds
  - **position**: origin for `"slide"` or `"grow"`

---

# Font Configuration

```yaml
font:
  family: "Arial, sans-serif"
  size: 14
  weight: "normal"
  style: "normal"
```

- **family**: Font family or stack
- **size**: Pixels
- **weight**: `"normal"`, `"bold"`, or 100–900
- **style**: `"normal"`, `"italic"`, `"oblique"`

---

# Color Palette

```yaml
palette:
  primary: "#95ccff"
  on-primary: "#003352"
  primary-container: "#004a75"
  on-primary-container: "#cde5ff"
  secondary: "#b9c8da"
  on-secondary: "#233240"
  secondary-container: "#3a4857"
  on-secondary-container: "#d5e4f6"
  tertiary: "#d2bfe6"
  on-tertiary: "#382a49"
  tertiary-container: "#4f4061"
  on-tertiary-container: "#eedbff"
  error: "#ffb4ab"
  on-error: "#690005"
  error-container: "#93000a"
  on-error-container: "#ffb4ab"
  background: "#1a1c1e"
  on-background: "#e2e2e5"
  surface: "#1a1c1e"
  on-surface: "#e2e2e5"
  surface-variant: "#42474e"
  on-surface-variant: "#c2c7cf"
  outline: "#8c9198"
  outline-variant: "#42474e"
  shadow: "#000000"
  scrim: "#000000"
  inverse-surface: "#e2e2e5"
  inverse-on-surface: "#2f3033"
  inverse-primary: "#00639a"
  surface-dim: "#121416"
  surface-bright: "#37393c"
  surface-container-lowest: "#0c0e11"
  surface-container-low: "#1a1c1e"
  surface-container: "#1e2022"
  surface-container-high: "#282a2d"
  surface-container-highest: "#333538"
```

- Full customization of UI colors.
- If omitted, Zend generates colors automatically.

---

# Widget Customization

Widgets include **search bar, clock, weather, music player**.
Each widget can be enabled/disabled, positioned anywhere, and customized in appearance/behavior.

---

## Search Bar

```yaml
search_bar:
  engine: "google"
  suggestions: true
  icon: "icons/search.png"
  placeholder: "Search here..."
  open_in_new_tab: false
```

- **engine**: `"google"`, `"bing"`, `"brave"`, `"duckduckgo"`
- **suggestions**: live search suggestions
- **icon**: path or URL
- **placeholder**: input placeholder text
- **open_in_new_tab**: true/false

---

## Clock

```yaml
clock:
  enabled: true
  format:
    time: "HH:mm:ss"
    date: "YYYY-MM-DD"
  timezone: "UTC"
  position: "top-right"
```

- Uses Day.js formatting
- Position supports keywords or coordinates
- Optional, default enabled

---

## Weather

```yaml
weather:
  enabled: true
  location: "New York,US"
  units: "metric"
  api_key: "your_openweathermap_api_key"
  position: "top-left"
```

- Requires **OpenWeatherMap API key**
- Units: `"metric"`, `"imperial"`, `"standard"`

---

## Music Player

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

- **local.enabled**: use local files
- **path/path_list**: files or folders
- **shuffle**: play order
- **visualizer**: animated visual display during music playback
- **type**: `"bars"`, `"wave"`, `"circle"`, `"dots"`
- **size**: intensity/scale of visualizer
