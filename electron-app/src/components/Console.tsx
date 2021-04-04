import React, { useContext, useState } from 'react';
import { ConsoleLine } from '../helpers/types';
import { ConsoleContext } from '../contexts/Contexts';

const Console: React.FC = () => {
  const { commandConsole } = useContext(ConsoleContext);
  const [isDevMode, setIsDevMode] = useState(false);

  return (
    <div className="console-wrapper">
      <div>
        <button
          type="button"
          className="dev-mode-console-button"
          onClick={() => setIsDevMode(!isDevMode)}
        >
          {'<>'}
        </button>
      </div>
      <div className="console-info">
        {commandConsole.map((line: ConsoleLine, i: number) => (
          <p key={i} style={{ color: line.failed ? 'red' : 'black' }}>
            {isDevMode ? line.debugInfo : line.text}
          </p>
        ))}
      </div>
    </div>
  );
};

export default Console;
