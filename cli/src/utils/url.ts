import { ZendConfig } from "../types/index";

function randomHexColor(): string {
  return (
    "#" +
    Math.floor(Math.random() * 0xffffff)
      .toString(16)
      .padStart(6, "0")
  );
}

export function buildUrl(config: ZendConfig): string {
  const base = "http://localhost:5000"; // test

  let random: string;
  if (config.random === "color") {
    random = randomHexColor();
  } else if (config.random === "image") {
    random = "image";
  } else {
    random = "n";
  }

  // Encode potentially problematic paths (image/music) properly
  const encodePath = (p?: string) => (p ? encodeURIComponent(p) : "n");

  const segments = [
    encodePath(config.image || config.images),
    encodePath(config.color),
    encodePath(config.music || config.musics),
    encodePath(config.engine),
    config.palette ? encodeURIComponent(JSON.stringify(config.palette)) : "n",
    encodePath(random),
    config.clock ? "t" : "f",
    config.visualizer ? "t" : "f",
    config.suggestions ? "t" : "f",
  ];

  return `${base}/${segments.join("/")}`;
}
