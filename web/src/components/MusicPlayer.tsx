import React, { useEffect, useState } from "react";
import AudioPlayer from "react-h5-audio-player";
import "react-h5-audio-player/lib/styles.css";
import { useAudioVisualizer } from "@/hooks/useAudioVisualizer";
import { Track, getMusic } from "@/api/music";

export const MusicPlayer: React.FC = () => {
  const [tracks, setTracks] = useState<Track[]>([]);
  const [current, setCurrent] = useState<number>(0);
  const { attach, dataArray } = useAudioVisualizer();

  useEffect(() => {
    (async () => {
      const musics = await getMusic();
      setTracks(musics);
    })();
  }, []);

  const handleClickNext = () => {
    setCurrent((prev) => (prev + 1) % tracks.length);
  };

  const handleClickPrevious = () => {
    setCurrent((prev) => (prev - 1 + tracks.length) % tracks.length);
  };

  const currentTrack = tracks[current];

  return (
    <div className="music">
      {/* Now Playing */}
      {currentTrack && (
        <div className="now-playing">
          {currentTrack.cover ? (
            <img src={currentTrack.cover} alt="cover" className="cover" />
          ) : (
            <span className="disc">💿</span>
          )}
          <div>
            <div className="t-title">{currentTrack.title}</div>
            {currentTrack.artist && (
              <div className="t-artist">{currentTrack.artist}</div>
            )}
          </div>
        </div>
      )}

      {/* Custom Audio Player */}
      {currentTrack && (
        <AudioPlayer
          autoPlay
          src={currentTrack.url}
          showJumpControls={true}
          onClickPrevious={handleClickPrevious}
          onClickNext={handleClickNext}
          onEnded={handleClickNext}
        />
      )}

      {/* Playlist */}
      <div className="picker-list">
        {tracks.length === 0 && <div className="muted">No tracks</div>}
        {tracks.map((t, i) => (
          <button
            key={t.url}
            className={`track ${i === current ? "active" : ""}`}
            onClick={() => setCurrent(i)}
            title={`${t.title}${t.artist ? " — " + t.artist : ""}`}
          >
            {t.cover ? (
              <img src={t.cover} alt="cover" />
            ) : (
              <span className="disc">💿</span>
            )}
            <span className="track-text">
              <span className="t-title">{t.title}</span>
              {t.artist && <span className="t-artist">{t.artist}</span>}
            </span>
          </button>
        ))}
      </div>

      {/* Visualizer */}
      <canvas
        id="visualizer"
        className="visualizer"
        width={600}
        height={60}
        ref={(el) => {
          if (!el) return;
          const ctx = el.getContext("2d");
          let raf: number;
          const render = () => {
            if (!ctx) return;
            const w = el.width;
            const h = el.height;
            ctx.clearRect(0, 0, w, h);
            if (dataArray.length) {
              const barWidth = Math.max(2, Math.floor(w / dataArray.length));
              for (let i = 0; i < dataArray.length; i++) {
                const v = dataArray[i] / 255;
                const barHeight = Math.max(2, Math.floor(v * h));
                const x = i * barWidth;
                const y = h - barHeight;
                ctx.fillRect(x, y, barWidth - 1, barHeight);
              }
            }
            raf = requestAnimationFrame(render);
          };
          raf = requestAnimationFrame(render);
          return () => cancelAnimationFrame(raf);
        }}
      />
    </div>
  );
};
