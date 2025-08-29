import React, { useEffect, useState } from "react";

interface BackgroundProps {
  config: any;
  children: React.ReactNode;
}

const Background: React.FC<BackgroundProps> = ({ config, children }) => {
  const [bgStyle, setBgStyle] = useState<React.CSSProperties>({});

  // Random colors or images if enabled
  useEffect(() => {
    const applyRandom = () => {
      let style: React.CSSProperties = {};

      // Random color
      if (config.random?.enabled && config.random.use === "color") {
        const min = config.random.min || 0;
        const max = config.random.max || 100;
        const r = Math.floor(Math.random() * (max - min) + min);
        const g = Math.floor(Math.random() * (max - min) + min);
        const b = Math.floor(Math.random() * (max - min) + min);
        style.backgroundColor = `rgb(${r},${g},${b})`;
      }

      // Random image
      if (
        config.random?.enabled &&
        config.random.use === "image" &&
        Array.isArray(config.image?.path_list)
      ) {
        const images = config.image.path_list;
        const index = Math.floor(Math.random() * images.length);
        style.backgroundImage = `url(${images[index]})`;
      }

      // Fallback to static image
      if (!style.backgroundImage && config.image?.path) {
        style.backgroundImage = `url(${config.image.path})`;
      }

      // Fallback to solid color
      if (
        !style.backgroundImage &&
        !style.backgroundColor &&
        config.solid?.enabled
      ) {
        style.backgroundColor = config.solid.color;
      }

      // Resize, position, repeat
      style.backgroundSize = config.resize?.mode || "cover";
      style.backgroundPosition = config.resize?.position || "center";
      style.backgroundRepeat = config.resize?.repeat || "no-repeat";

      // Blur and opacity
      style.filter = config.blur ? `blur(${config.blur}px)` : undefined;
      style.opacity = config.opacity;

      setBgStyle(style);
    };

    applyRandom();

    if (config.random?.enabled && config.random.interval) {
      const interval = setInterval(applyRandom, config.random.interval * 1000);
      return () => clearInterval(interval);
    }
  }, [config]);

  // Handle CSS transitions
  const transitionStyle: React.CSSProperties = {};
  if (config.transition?.type) {
    const duration = config.transition.duration || 1;
    switch (config.transition.type) {
      case "fade":
        transitionStyle.transition = `all ${duration}s ease-in-out`;
        break;
      case "slide":
        transitionStyle.transition = `transform ${duration}s ease-in-out`;
        break;
      case "zoom":
        transitionStyle.transition = `transform ${duration}s ease-in-out`;
        break;
      case "grow":
        transitionStyle.transition = `transform ${duration}s ease-in-out`;
        break;
      default:
        break;
    }
  }

  return (
    <>
      <div
        className="zend-background"
        style={{ ...bgStyle, ...transitionStyle }}
      ></div>
      {children}
    </>
  );
};

export default Background;
