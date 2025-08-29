---
title: "General UI Configuration"
date: 2025-08-29
draft: false
---

# General UI Configuration

Zend allows fine-tuning of general UI properties such as **blur, opacity, and transition effects** to control the look and feel of your homepage.

---

## Example UI Configuration

```yaml
general:
  blur: 5
  opacity: 1
  transition:
    type: "fade"
    duration: 1
    position: "center"
```

---

## **Properties**

- **blur**: Background blur in pixels

  - Example: `5` → slight blur, `0` → no blur

- **opacity**: Overall transparency of UI elements (0–1)

  - `0` = fully transparent, `1` = fully opaque

- **transition**: Animation settings for UI elements

  - **type**: `"fade"`, `"slide"`, `"zoom"`, `"grow"`
  - **duration**: Duration in seconds
  - **position**: Origin for `"slide"` or `"grow"` transitions

---

## **Usage Example**

```yaml
general:
  blur: 10
  opacity: 0.9
  transition:
    type: "slide"
    duration: 0.5
    position: "top-left"
```

- This configuration creates a **slightly blurred UI** with semi-transparent elements and smooth slide animations starting from the top-left corner.

---

## **Tips**

- Adjust **blur** and **opacity** carefully — too much blur or transparency can reduce readability.
- Use **transition effects** sparingly to maintain a smooth user experience.
- Combine UI settings with your **color palette** for cohesive visuals.
