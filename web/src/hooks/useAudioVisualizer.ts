import { useCallback, useEffect, useRef, useState } from "react";

export const useAudioVisualizer = () => {
  const analyserRef = useRef<AnalyserNode | null>(null);
  const srcRef = useRef<MediaElementAudioSourceNode | null>(null);
  const ctxRef = useRef<AudioContext | null>(null);
  const [dataArray, setDataArray] = useState<Uint8Array>(new Uint8Array(0));

  const attach = useCallback((audio: HTMLAudioElement) => {
    if (srcRef.current) return;
    const ctx = new (window.AudioContext ||
      (window as any).webkitAudioContext)();
    const source = ctx.createMediaElementSource(audio);
    const analyser = ctx.createAnalyser();
    analyser.fftSize = 256;
    source.connect(analyser);
    analyser.connect(ctx.destination);
    analyserRef.current = analyser;
    srcRef.current = source;
    ctxRef.current = ctx;

    const arr = new Uint8Array(analyser.frequencyBinCount);
    setDataArray(arr);
  }, []);

  useEffect(() => {
    let id: number;
    const tick = () => {
      const analyser = analyserRef.current;
      if (analyser && dataArray.length) {
        analyser.getByteFrequencyData(dataArray as Uint8Array<ArrayBuffer>);
      }
      id = requestAnimationFrame(tick);
    };
    id = requestAnimationFrame(tick);
    return () => cancelAnimationFrame(id);
  }, [dataArray]);

  return { attach, dataArray };
};
