import React from 'react';
// import { ConsoleContext } from '../contexts/Contexts';
import { execute } from '../helpers/cli';
import { ConsoleLine } from '../helpers/types';

const Sync: React.FC = () => {
  const handleSync = () => {
    execute('dotsync-cli -sync', (output, error, stderr) => {
      const consoleLine: ConsoleLine = {
        text: 'test',
        failed: false,
      };
      console.log(consoleLine, output);
      console.log(stderr);
    });
  };

  return (
    <div className="syncWrapper base-component">
      <button onClick={handleSync} type="button"> Sync </button>
    </div>
  );
};

export default Sync;
