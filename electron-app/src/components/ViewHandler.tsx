import React, { useContext } from 'react';
import Sync from './Sync';
import Init from './Init';
import Config from './Config';
import { ViewContext } from '../contexts/Contexts';

const ViewHandler: React.FC = () => {
  const { view } = useContext(ViewContext);

  return (
    <div className="view-handler">
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
