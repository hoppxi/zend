---
title: "Font Configuration"
date: 2025-08-29
draft: false
---

# Font Configuration

Zend allows full customization of fonts used throughout the homepage, including **family, size, weight, and style**. This ensures a consistent look across all widgets and UI elements.

---

## Font Settings

```yaml
font:
  family: "Arial, sans-serif"
  size: 14
  weight: "normal"
  style: "normal"
```

### **Properties**

- **family**: Font family or fallback stack

  - Example: `"Arial, sans-serif"` or `"Roboto, sans-serif"`

- **size**: Font size in **pixels**

  - Example: `14`, `16`, `18`

- **weight**: Font weight

  - `"normal"`, `"bold"`, or numeric values `100–900`

- **style**: Font style

  - `"normal"`, `"italic"`, or `"oblique"`

---

## **Usage Example**

```yaml
font:
  family: "Roboto, sans-serif"
  size: 16
  weight: "bold"
  style: "italic"
```

- This configuration will render text in **Roboto**, size **16px**, bold and italicized.

---

## **Tips**

- Always provide a fallback font family to ensure proper rendering if a custom font is unavailable.
- Combine font size and weight with your UI layout to maintain readability.
- Use `"normal"` weight for body text and `"bold"` for headings or emphasis.
- Apply consistent font settings across all widgets for a cohesive design.
