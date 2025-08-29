import axios from "axios";

export const getImage = async () => {
  try {
    const res = await axios.get("/api/image/");
    const html = res.data;

    const parser = new DOMParser();
    const doc = parser.parseFromString(html, "text/html");

    const links = Array.from(doc.querySelectorAll("a"));

    const imgs = links
      .map((a) => a.getAttribute("href") || "")
      .filter((href) =>
        href.match(/\.(jpg|jpeg|png|gif|webp|bmp|tiff|svg|ico|avif)$/i)
      )
      .map((filename) => `/api/image/${filename.replace(/^\/+/, "")}`); // normalize

    return imgs;
  } catch (error) {
    console.error("Error fetching images:", error);
    return [];
  }
};
