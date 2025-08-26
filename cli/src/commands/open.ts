import { resolveConfig } from "../utils/config";
import { buildUrl } from "../utils/url";
import { openBrowser } from "../utils/browser";
import { Logger } from "../utils/logger";

export async function openCommand(
  cliOpts: Record<string, string | boolean | string[] | Record<string, string>>
) {
  const config = await resolveConfig(cliOpts);
  const url = buildUrl(config);

  Logger.info(`Launching Zend → ${url}`);
  await openBrowser(url);
}
