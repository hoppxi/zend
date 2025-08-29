---
title: "Color Palette"
date: 2025-08-29
draft: false
---

# Color Palette

Zend allows full customization of UI colors via a **palette configuration**. This controls everything from primary and secondary colors to surfaces, outlines, and error states.

---

## Example Palette

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
```

- **primary / secondary / tertiary**: Main branding colors
- **on-**\*: Text/icon color for the corresponding background
- **containers**: Color for surfaces containing content (cards, sections)
- **background / surface**: Default page and widget backgrounds
- **error / on-error**: Colors used for error states and messages

---

## **Generating a Palette**

You have two main options for generating a cohesive theme:

1. **Material Color Utilities (MCU)**

   - Use the [Material Color Utilities](https://material.io/resources/color) to generate consistent palettes for primary, secondary, and tertiary colors.
   - Ensures accessible contrast and coherent design across your UI.

2. **Using `mcu-cli`**

   - If you prefer a CLI workflow, you can use my module:

     ```bash
     git clone https://github.com/hoppxi/mcu-cli.git
     cd mcu-cli
     ./mcuc generate <base-color>
     ```

   - This generates a complete Zend palette from a single base color.

> You can choose either method; the resulting colors can be pasted directly into your `palette` YAML in the Zend configuration.

---

## **Tips**

- Keep **text contrast** in mind — always check `on-primary` and `on-surface` readability.
- Use **secondary and tertiary colors** for accents, highlights, and non-critical UI elements.
- Error colors should be **clearly visible** but not overly jarring.
