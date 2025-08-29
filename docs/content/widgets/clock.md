---
title: "Clock"
date: 2025-08-29
draft: false
---

# Clock Widget

Displays the current time and optionally the date. Fully customizable in format, timezone, and position.

---

## Example Configuration

```yaml
clock:
  enabled: true
  format:
    time: "HH:mm:ss"
    date: "YYYY-MM-DD"
  timezone: "UTC"
  position: "top-right"
```

### **Properties**

- **enabled**: true/false
- **format.time**: Day.js format string for time
- **format.date**: Day.js format string for date
- **timezone**: Timezone identifier (e.g., `"UTC"`, `"America/New_York"`)
- **position**: Predefined keyword or `(x%, y%)` coordinates

---

## Usage Tips

- Use `"HH:mm:ss"` for 24-hour format or `"hh:mm A"` for 12-hour format.
- Position the clock where it doesn’t overlap other widgets.
- Optional: combine with background opacity for better readability.
