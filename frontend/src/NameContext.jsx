import React, { createContext, useState } from 'react';

const NameContext = createContext();

export function NameProvider({ children }) {
  const [name, setName] = useState('');
  const [poll, setPoll] = useState();

  return (
    <NameContext.Provider value={{ name, setName, poll, setPoll }}>
      {children}
    </NameContext.Provider>
  );
}

export default NameContext;
