# Zend

**Zend** is a minimal home page for browsers, offering beautiful backgrounds, live suggestions, a clock, music playback, and more. Packed wit cli tool to config from terminal.

## Usage

```bash
zend [options] [command]
```

Launch the Zend home page with optional customization.

### Examples

- Launch Zend with a specific background image:

```bash
zend -i ~/Pictures/wallpapers
```

- Launch Zend with a random color:

```bash
zend -r color
```

- Enable clock and live suggestions:

```bash
zend -t -s
```

- Set a custom search engine and main color:

```bash
zend -e duckduckgo -c "#3498db"
```

- Play music with visualizer:

```bash
zend -m ~/Music -v
```
