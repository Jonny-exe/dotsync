import React from 'react';
import { execute } from '../helpers/cli';

const Init: React.FC = () => {
  const handleInit = () => {
    execute('dotsync-cli -init', (output, error, stderr) => {
      console.log(output, error, stderr);
    });
  };

  return (
    <div className="initWrapper baseComponent">
      <button onClick={handleInit} type="button">Init</button>
    </div>
  );
};

export default Init;
