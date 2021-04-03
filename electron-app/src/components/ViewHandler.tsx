import React, { useContext } from 'react';
import Sync from './Sync';
import Init from './Init';
import Config from './Config';
import { ViewContext } from '../contexts/Contexts';

const ViewHandler: React.FC = () => {
  const { view } = useContext(ViewContext);

  const style = {
    color: 'white',
  };

  return (
    <div className="app" style={style}>
      {view !== 'config' ? (
        <>
          <Init />
          <Sync />
          <Config />
        </>
      ) : (
        <Config />
      )}
    </div>
  );
};

export default ViewHandler;
