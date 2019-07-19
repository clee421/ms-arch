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
  const username = useFormInput('', {
    name: 'username',
  });
  const password = useFormInput('', {
    type: 'password',
    name: 'password',
  });

  function handleSubmit(): void {
    // eslint-disable-next-line
    console.log(`Username: ${username.value}; Password: ${password.value}`);
  }

  return (
    <div className="app-container">
      <div className="input-boxes">
        <label>
          <span>username</span>
          <input {...username} />
        </label>
        <label>
          <span>password</span>
          <input {...password} />
        </label>
      </div>

      <div className="submit-container">
        <button type="button" onClick={handleSubmit}>
          Submit
        </button>
      </div>

      <div className="token-container">{`Token: ${token}`}</div>
    </div>
  );
};

export default App;
