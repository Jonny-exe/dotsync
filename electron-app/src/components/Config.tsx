import React, { useContext, useState } from 'react';
import ConfigItem from './ConfigItem';
import { ViewContext, ConsoleContext } from '../contexts/Contexts';
import { readConfig, writeConfig } from '../helpers/cli';
import { ConsoleLine } from '../helpers/types';

const Config: React.FC = () => {
  const { view, setView } = useContext(ViewContext);
  const { commandConsole, setCommandConsole } = useContext(ConsoleContext);
  const [userConfig, setUserConfig] = useState(readConfig());
  console.log(userConfig);

  const saveConfig = () => {
    const error = writeConfig(userConfig);
    const currentDate = new Date();
    const time = `${currentDate.getHours()}:${currentDate.getMinutes()}:${currentDate.getSeconds()}`;
    const type = 'Config updated';

    const consoleLine: ConsoleLine = {
      text: `${time}: ${type}`,
      failed: error !== null,
      debugInfo: `${time}: ${type}`,
    };
    setCommandConsole([...commandConsole, consoleLine]);
  };

  return (
    <>
      {view === 'config' ? (
        <>
          <h1 style={{ fontWeight: 'normal', color: 'white' }}> Config </h1>
          <div className="config-wrapper">
            {Object.entries(userConfig).length !== 0 ? (
              <>
                <ConfigItem
                  name='Username:'
                  value={userConfig}
                  setValue={setUserConfig}
                  objectKey='gh-access-token'
                />
                <ConfigItem
                  name='Access token:'
                  value={userConfig}
                  setValue={setUserConfig}
                  objectKey='gh-username'
                />
                <div>
                  <button
                    className="config-button"
                    type="button"
                    onClick={saveConfig}
                  >
                    Save
                  </button>
                  <button
                    className="config-button"
                    type="button"
                    onClick={() => setView('home')}
                  >
                    Back
                  </button>
                </div>
              </>
            ) : (
              <>
                <h2 style={{ fontWeight: 'normal' }}>
                  {' '}
                  No config file was found or config file was empty
                </h2>
                <button
                  className="config-button"
                  type="button"
                  onClick={() => setView('home')}
                >
                  {' '}
                  Back{' '}
                </button>
              </>
            )}
          </div>
        </>
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
