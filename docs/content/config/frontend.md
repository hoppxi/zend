---
title: "Frontend Build Configuration"
date: 2025-08-29
draft: false
---

# Frontend Build Configuration

Zend allows you to specify the location of your **compiled frontend files** (HTML/CSS/JS). This is useful if you are deploying a custom-built frontend or using a specific environment.

---

## Example Configuration

```yaml
dist: "/home/user/myproject/frontend/dist" # Absolute path
# dist: "./dist"                             # Relative path
```

- **dist**: Location of the compiled frontend.

  - Can be **absolute** or **relative** to your Zend project.
  - Optional: if omitted, Zend will serve its **built-in frontend**.

---

## Usage Tips

- When using a custom frontend, make sure the `dist` folder contains **index.html** and all required assets.
- For development, relative paths (`./dist`) are often easier to manage.
- For production or deployment, use absolute paths for consistency.
