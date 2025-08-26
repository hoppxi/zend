import chalk from "chalk";

export const Logger = {
  info: (msg: string) => console.log(chalk.blue("INFO"), msg),
  success: (msg: string) => console.log(chalk.green("SUCCESS"), msg),
  warn: (msg: string) => console.log(chalk.yellow("WARN"), msg),
  error: (msg: string) => console.error(chalk.red("ERROR"), msg),
};
