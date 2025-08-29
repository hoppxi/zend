import axios from "axios";

export type Track = {
  title: string;
  url: string;
  artist?: string;
  cover?: string;
};

export const getMusic = async () => {
  try {
    const res = await axios.get("/api/music/");
    const html = res.data;

    const parser = new DOMParser();
    const doc = parser.parseFromString(html, "text/html");

    const links = Array.from(doc.querySelectorAll("a"))
      .map((a) => (a as HTMLAnchorElement).getAttribute("href") || "")
      .filter((href) => href.match(/\.(mp3|wav|ogg|flac|m4a)$/i));

    const parsed: Track[] = links.map((filename) => ({
      title: decodeURIComponent(filename.split("/").pop() || "Track"),
      url: `/api/music/${filename.replace(/^\/+/, "")}`,
    }));

    return parsed;
  } catch (error) {
    console.error("Error fetching music:", error);
    return [];
  }
};
