import axios from "axios";

export const getWeather = async () => {
  try {
    const res = await axios.get(`/api/weather`);
    console.log(res.data);
    return res.data;
  } catch (err) {
    console.error("Failed to fetch weather:", err);
    return null;
  }
};
