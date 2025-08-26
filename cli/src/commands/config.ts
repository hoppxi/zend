import { setConfig, getConfig, loadConfig } from "../utils/config";
import { buildUrl } from "../utils/url";
import { openBrowser } from "../utils/browser";
import { Logger } from "../utils/logger";

export async function configCommand(file?: string) {
  // @ts-expect-error extract opts from commander
  const opts = this.opts();

  if (opts.set) {
    const [key, value] = opts.set.split("=");
    if (!key || value === undefined) {
      Logger.error("Invalid format: --set key=value");
      return;
    }
    await setConfig(key, value, file);
    Logger.success(`Set ${key} = ${value}`);
    return;
  }

  if (opts.get) {
    const val = await getConfig(opts.get, file);
    Logger.info(`${opts.get} = ${val}`);
    return;
  }

  if (file) {
    const config = await loadConfig(file);
    const url = buildUrl(config);
    Logger.info(`Opening Zend with config: ${file}`);
    Logger.info(`URL → ${url}`);
    await openBrowser(url);
    return;
  }

  Logger.error(
    "Use --set <key=value>, --get <key>, or provide a config file to launch Zend"
  );
}
