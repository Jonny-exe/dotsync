import React, { useContext } from 'react';
import { ConsoleLine } from '../helpers/types';
import { ConsoleContext } from '../contexts/Contexts';

const Console: React.FC = () => {
  const { commandConsole } = useContext(ConsoleContext);
  return (
    <div className="console-wrapper">
      {commandConsole.map((line: ConsoleLine, i: number) => (
        <p key={i} style={{ color: line.failed ? 'red' : 'black' }}>
          {line.text}
        </p>
      ))}
    </div>
  );
};

export default Console;
