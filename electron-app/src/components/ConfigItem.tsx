import React from 'react';

interface IObjectKeys {
  [key: string]: string | number;
}

interface Config extends IObjectKeys {
  'gh-username': string;
  'gh-access-token': string;
}

interface Props {
  name: string;
  value: Config;
  setValue: (newConfig: Config) => void;
  objectKey: string;
}

const ConfigItem: React.FC<Props> = ({ objectKey, name, value, setValue }) => {
  return (
    <div className="username-container">
      <div className="username">{name}</div>
      <div className="username-input-wrapper">
        <input
          className="username-input"
          placeholder="username"
          type="text"
          onChange={(e: any) =>
            setValue({ ...value, [objectKey]: e.target.value })
          }
          value={value[objectKey]}
        />
      </div>
    </div>
  );
};

export default ConfigItem;
