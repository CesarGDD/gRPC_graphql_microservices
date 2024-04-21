import React from 'react';
import logo from './logo.svg';
import './App.css';
import { useGetUsersQuery } from './graphql/generated/graphql';

function App() {
  const [{ data, fetching, error }] = useGetUsersQuery();

  if (fetching) return <p>Loading...</p>;
  if (error) return <p>Oh no... {error.message}</p>;

  return (
    <div className="App">
      <header className="App-header">
        <img src={logo} className="App-logo" alt="logo" />
        <p>
          Edit <code>src/App.tsx</code> and save to reload.
        </p>
      <div>
        {data?.getUsers.map(user => (
          <p key={user.userId}>{user.username}</p>
        ))}
      </div>
      </header>
    </div>
  );
}

export default App;
