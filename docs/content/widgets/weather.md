---
title: "Weather Widget"
date: 2025-08-29
draft: false
---

# Weather Widget

Displays current weather information for a specified location. Requires an OpenWeatherMap API key.

---

## Example Configuration

```yaml
weather:
  enabled: true
  location: "New York,US"
  units: "metric"
  api_key: "your_openweathermap_api_key"
  position: "top-left"
```

### **Properties**

- **enabled**: true/false
- **location**: City and country code (e.g., `"London,GB"`)
- **units**: `"metric"`, `"imperial"`, `"standard"`
- **api_key**: OpenWeatherMap API key
- **position**: Widget location on screen

---

## Usage Tips

- Ensure API key is valid and has quota.
- Use units that match your audience (metric vs imperial).
- Combine with CSS transitions for smooth updates.
