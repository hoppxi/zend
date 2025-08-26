export interface ZendConfig {
  image?: string; // Single background image path
  images?: string; // Folder of background images
  color?: string; // Hex color
  random?: RandomMode; // Random selection mode
  palette?: Palette; // Material color palette
  engine?: string; // Search engine (google, bing, brave, etc.)
  clock?: boolean; // Show clock
  suggestions?: boolean; // Enable live search suggestions
  music?: string; // Single music file
  musics?: string; // Folder of music files
  visualizer?: boolean; // Enable music visualizer
}

export type RandomMode = "image" | "color" | "null" | undefined;
