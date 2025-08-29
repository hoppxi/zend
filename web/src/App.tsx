import React, { useEffect, useMemo, useState } from "react";
import { Clock } from "@/components/Clock";
import { Weather } from "@/components/Weather";
import { ImagePicker } from "@/components/ImagePicker";
import { MusicPlayer } from "@/components/MusicPlayer";
import { SearchBar } from "@/components/SearchBar";

import "@/styles/main.scss";

export default function App() {
  const [bgUrl, setBgUrl] = useState<string | null>(null);

  useEffect(() => {
    if (!bgUrl) return;
    const el = document.documentElement;
    el.style.setProperty("--bg-image", `url("${bgUrl}")`);
  }, [bgUrl]);

  // Fallback gradient when no background selected
  const backgroundStyle = useMemo(
    () => ({
      backgroundImage: bgUrl
        ? undefined
        : "radial-gradient(1200px 800px at 20% 10%, rgba(255,255,255,0.12), transparent), radial-gradient(700px 500px at 80% 80%, rgba(255,255,255,0.08), transparent)",
    }),
    [bgUrl]
  );

  return (
    <div className="app" style={backgroundStyle}>
      <div className="glass">
        <div className="top-row">
          <SearchBar />
          <div className="top-right">
            <Weather />
          </div>
        </div>

        <div className="center">
          <Clock />
        </div>

        <div className="bottom">
          <ImagePicker onPick={setBgUrl} />
          <MusicPlayer />
        </div>
      </div>
    </div>
  );
}
