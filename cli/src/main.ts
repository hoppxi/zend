import { Command } from "commander";
import { openCommand } from "./commands/open";
import { configCommand } from "./commands/config";
import pkg from "../package.json";

const program = new Command();

program
  .name("zend")
  .description("Minimal browser landing page launcher")
  .version(pkg.version, "-V, --version", "Output the version number")
  .showHelpAfterError()
  .showSuggestionAfterError()
  .helpOption("-h, --help", "Display help for command")
  .option("-i, --image <path>", "File/folder path for background image(s)")
  .option("-c, --color <hex>", "Main color (hex)")
  .option("-r, --random <type>", "Random mode: image|color|null")
  .option("-p, --palette <json>", "Material palette JSON")
  .option("-e, --engine <engine>", "Search engine")
  .option("-t, --clock", "Enable clock")
  .option("-s, --suggestions", "Enable live suggestions")
  .option("-m, --music <path>", "File/folder path for music")
  .option("-v, --visualizer", "Enable music visualizer")
  .action(openCommand);

program
  .command("config")
  .description("Manage zend config")
  .option("--get <key>", "Get config value")
  .option("--set <key=value>", "Set config value")
  .argument("[file]", "Use custom config file")
  .action(configCommand);

program.parse(process.argv);
