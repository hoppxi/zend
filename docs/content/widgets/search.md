---
title: "Search Bar"
date: 2025-08-29
draft: false
---

# Search Bar

The search bar widget allows users to perform searches directly from Zend. You can customize the engine, suggestions, icon, placeholder, and behavior.

---

## Example Configuration

```yaml
search_bar:
  engine: "google"
  suggestions: true
  icon: "icons/search.png"
  placeholder: "Search here..."
  open_in_new_tab: false
```

### **Properties**

- **engine**: `"google"`, `"bing"`, `"brave"`, `"duckduckgo"`
- **suggestions**: Enable live search suggestions
- **icon**: Path or URL for the search icon
- **placeholder**: Input placeholder text
- **open_in_new_tab**: Open results in a new tab (true/false)

---

## Usage Tips

- Choose a default search engine appropriate for your users.
- Use icons that match your theme.
- Live suggestions improve UX but require internet connectivity.
