import axios from "axios";

export const getConfig = async () => {
  try {
    const res = await axios.get(`/api/config`);
    console.log(res.data);
    return res.data;
  } catch (err) {
    console.error("Failed to fetch Zend config:", err);
    return null;
  }
};
