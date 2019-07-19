import React, { ReactElement } from 'react';
import { useSelector } from 'react-redux';

// Types
import { AppState } from '../store';

// Hooks
import { useFormInput } from '../hooks/form-input';

// Style
import './app.scss';

const App: React.FC = (): ReactElement => {
  const token = useSelector((state: AppState): string => state.system.token);
  const username = useFormInput('', { type: 'text', name: 'username' });
  const password = useFormInput('', { type: 'password', name: 'password' });

  return (
    <div className="app-container">
      <div className="input-boxes">
        <label>
          username
          <input {...username} />
        </label>
        <label>
          password
          <input {...password} />
        </label>
      </div>

      <div className="token-container">{`Token: ${token}`}</div>
    </div>
  );
};

export default App;
