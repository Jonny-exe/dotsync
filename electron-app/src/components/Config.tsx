import React, { useContext } from 'react';
import { ViewContext } from '../contexts/Contexts';
import { execute } from '../helpers/cli';

const Config: React.FC = () => {
  const { view, setView } = useContext(ViewContext);

  return (
    <>
      {view === 'config' ? (
        <div>
          <h1> These are your configs </h1>
          <button
            type="button"
            onClick={() => {
              execute('dotsync-cli -config', (output, error, stderr) => {
                console.log(output, error, stderr);
              });
            }}
          >
            Execute
          </button>
          <button type="button">See output</button>
          <button type="button" onClick={() => setView('home')}>
            Back
          </button>
        </div>
      ) : (
        <div className="configWrapper base-component">
          <button type="button" onClick={() => setView('config')}>
            Config
          </button>
        </div>
      )}
    </>
  );
};

export default Config;
