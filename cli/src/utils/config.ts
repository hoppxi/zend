import fs from "fs/promises";
import path from "path";
import os from "os";
import { ZendConfig } from "../types";

const CONFIG_DIR =
  process.env.XDG_CONFIG_HOME || path.join(os.homedir(), ".config");
const CONFIG_PATH = path.join(CONFIG_DIR, "zend", "config.json");

export async function loadConfig(customPath?: string): Promise<ZendConfig> {
  const file = customPath || CONFIG_PATH;
  try {
    const data = await fs.readFile(file, "utf8");
    return JSON.parse(data);
  } catch {
    return {};
  }
}

export async function saveConfig(config: ZendConfig, customPath?: string) {
  const file = customPath || CONFIG_PATH;
  await fs.mkdir(path.dirname(file), { recursive: true });
  await fs.writeFile(file, JSON.stringify(config, null, 2), "utf8");
}

export async function resolveConfig(cliOpts: ZendConfig, customPath?: string) {
  const fileConfig = await loadConfig(customPath);
  return { ...fileConfig, ...cliOpts };
}

export async function setConfig(
  key: string,
  value:
    | string
    | boolean
    | undefined
    | null
    | string[]
    | Record<string, string>,
  customPath?: string
) {
  const cfg = await loadConfig(customPath);
  cfg[key as keyof ZendConfig] = value;
  await saveConfig(cfg, customPath);
}

export async function getConfig(key: string, customPath?: string) {
  const cfg = await loadConfig(customPath);
  return cfg[key as keyof ZendConfig] ?? null;
}
