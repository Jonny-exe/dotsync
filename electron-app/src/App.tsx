import React, { useState } from 'react';
import { BrowserRouter as Router, Switch, Route } from 'react-router-dom';
import './App.global.css';
import ViewHandler from './components/ViewHandler';
import { ConsoleContext, ViewContext } from './contexts/Contexts';
import { ConsoleLine } from './helpers/types';

const Body = () => {
  const [view, setView] = useState('home');
  const [commandConsole, setCommandConsole] = useState<ConsoleLine[]>([]);
  return (
    <div className="app">
      <h1> Choose one </h1>

      <ConsoleContext.Provider value={{ commandConsole, setCommandConsole }}>
        <ViewContext.Provider value={{ view, setView }}>
          <ViewHandler />
        </ViewContext.Provider>
      </ConsoleContext.Provider>
    </div>
  );
};

export default function App() {
  return (
    <Router>
      <Switch>
        <Route path="/" component={Body} />
      </Switch>
    </Router>
  );
}
