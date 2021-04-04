import React, { useContext } from 'react';
import { execute } from '../helpers/cli';
import { ConsoleLine } from '../helpers/types';
import { ConsoleContext } from '../contexts/Contexts';

interface Props {
  type: string;
}

const OptionButton: React.FC<Props> = ({ type }) => {
  const { commandConsole, setCommandConsole } = useContext(ConsoleContext);
  const handleCli = () => {
    execute(`dotsync-cli -${type}`, (output, error, stderr) => {
      const currentDate = new Date();
      const time = `${currentDate.getHours()}:${currentDate.getMinutes()}:${currentDate.getSeconds()}`;
      const consoleLine: ConsoleLine = {
        text: `${time}: ${type}`,
        failed: error !== null,
        debugInfo: stderr,
      };
      setCommandConsole([...commandConsole, consoleLine]);
    });
  };
  return (
    <>
      <button onClick={() => handleCli()} type="button">
        {type}
      </button>
    </>
  );
};
export default OptionButton;
