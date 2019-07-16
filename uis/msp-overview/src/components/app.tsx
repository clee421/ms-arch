import React, { ReactElement } from 'react';
import { useSelector } from 'react-redux';

// Types
import { AppState } from '../store';

import './app.scss';

const App: React.FC = (): ReactElement => {
  const token = useSelector((state: AppState): string => state.system.token);
  return (
    <div className="app-container">
      <div className="input-boxes">
        <label>
          username
          <input type="text" name="username" />
        </label>
        <label>
          password
          <input type="text" name="password" />
        </label>
        <label>
          token
          <input type="text" name="token" value={token} />
        </label>
      </div>
    </div>
  );
};

export default App;
