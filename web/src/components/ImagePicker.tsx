import React, { useEffect, useState } from "react";
import { getImage } from "@/api/image";

export const ImagePicker: React.FC<{ onPick: (url: string) => void }> = ({
  onPick,
}) => {
  const [images, setImages] = useState<string[]>([]);
  const [open, setOpen] = useState(false);

  useEffect(() => {
    (async () => {
      const images = await getImage();
      setImages(images);
    })();
  }, []);

  return (
    <div className="picker">
      <button
        className="btn"
        onClick={() => setOpen((v) => !v)}
        aria-expanded={open}
      >
        🖼️ Background
      </button>

      {open && (
        <div className="picker-grid">
          {images.length === 0 && <div className="muted">No images</div>}
          {images.map((src) => (
            <button
              key={src}
              className="thumb"
              onClick={() => onPick(src)}
              title={src}
            >
              <img src={src} loading="lazy" alt="background option" />
            </button>
          ))}
        </div>
      )}
    </div>
  );
};
