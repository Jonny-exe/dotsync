import React, { useContext } from 'react';
import Sync from './Sync';
import Init from './Init';
import Config from './Config';
import Console from './Console';
import { ViewContext } from '../contexts/Contexts';

const ViewHandler: React.FC = () => {
  const { view } = useContext(ViewContext);

  return (
    <>
      {view !== 'config' ? (
        <>
          <div className="options-wrapper">
            <Init />
            <Sync />
            <Config />
          </div>
          <Console />
        </>
      ) : (
        <Config />
      )}
    </>
  );
};

export default ViewHandler;
