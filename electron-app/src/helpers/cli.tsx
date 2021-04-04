const { exec } = require('child_process');
const fs = require('fs');
const yaml = require('js-yaml');

export const execute = (
  command: string,
  callback: (output: string, error: any, stderr: string) => any
) => {
  exec(command, (error: any, stdout: any, stderr: any) => {
    callback(stdout, error, stderr);
  });
};

const path = `${process.env.HOME}/.dotsync.yaml`;
console.log(path);

export const readConfig = () => {
  try {
    const fileContents = fs.readFileSync(path, 'utf8');
    const data = yaml.load(fileContents);
    return data;
  } catch {
    return {};
  }
};

export const writeConfig = (config: Record<string, unknown>) => {
  try {
    const yamlStr = yaml.dump(config);
    fs.writeFileSync(path, yamlStr, 'utf8');
    return null;
  } catch {
    return true;
  }
};
readConfig();
