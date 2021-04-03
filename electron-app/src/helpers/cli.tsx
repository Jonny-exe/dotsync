const { exec } = require('child_process');

export const execute = (
  command: string,
  callback: (output: string, error: any, stderr: string) => any
) => {
  exec(command, (error: any, stdout: any, stderr: any) => {
    callback(stdout, error, stderr);
  });
};
