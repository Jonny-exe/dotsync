import React from 'react';
import { ConsoleLine } from '../helpers/types';

const setViewFunc = (hello: string) => hello;
const setConsoleFunc = (newConsole: ConsoleLine[]) => newConsole;

export const ViewContext = React.createContext<{
  view: string;
  setView: (newView: string) => any;
}>({ view: 'home', setView: setViewFunc });
export const ConsoleContext = React.createContext<{
  commandConsole: ConsoleLine[];
  setCommandConsole: (newView: ConsoleLine[]) => any;
}>({ commandConsole: [], setCommandConsole: setConsoleFunc });
