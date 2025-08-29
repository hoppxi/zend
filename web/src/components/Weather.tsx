import React, { useEffect, useState } from "react";

type WeatherData = {
  temperature: number;
  windspeed: number;
  weathercode: number;
};

const weatherIcon = (code: number) => {
  // Minimal mapping for demo purposes
  if ([0].includes(code)) return "☀️";
  if ([1, 2, 3].includes(code)) return "⛅";
  if ([45, 48].includes(code)) return "🌫️";
  if ([51, 53, 55, 56, 57, 61, 63, 65, 66, 67, 80, 81, 82].includes(code))
    return "🌧️";
  if ([71, 73, 75, 77, 85, 86].includes(code)) return "❄️";
  if ([95, 96, 99].includes(code)) return "⛈️";
  return "🌤️";
};

export const Weather: React.FC = () => {
  const [coords, setCoords] = useState<{ lat: number; lon: number } | null>(
    null
  );
  const [data, setData] = useState<WeatherData | null>(null);
  const [err, setErr] = useState<string | null>(null);

  useEffect(() => {
    if (!("geolocation" in navigator)) {
      setErr("Geolocation not available");
      return;
    }

    navigator.geolocation.getCurrentPosition(
      (pos) =>
        setCoords({ lat: pos.coords.latitude, lon: pos.coords.longitude }),
      () => setErr("Location blocked")
    );
  }, []);

  useEffect(() => {
    const fetchWeather = async () => {
      if (!coords) return;

      try {
        const url = `https://api.open-meteo.com/v1/forecast?latitude=${coords.lat}&longitude=${coords.lon}&current_weather=true`;
        const res = await fetch(url);
        const json = await res.json();

        if (!json.current_weather) {
          setErr("Weather data unavailable");
          return;
        }

        const cw = json.current_weather;
        setData({
          temperature: cw.temperature,
          windspeed: cw.windspeed,
          weathercode: cw.weathercode,
        });
      } catch {
        setErr("Weather unavailable");
      }
    };

    fetchWeather();
  }, [coords]);

  return (
    <div className="weather">
      {data ? (
        <>
          <span className="wx-emoji" aria-hidden={true}>
            {weatherIcon(data.weathercode)}
          </span>
          <span className="wx-temp">{Math.round(data.temperature)}°</span>
          <span className="wx-wind">{Math.round(data.windspeed)} km/h</span>
        </>
      ) : (
        <span className="wx-loading">{err ?? "Loading weather…"}</span>
      )}
    </div>
  );
};
