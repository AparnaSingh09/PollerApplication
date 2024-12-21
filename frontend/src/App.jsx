import React from 'react';
import { BrowserRouter as Router, Route, Routes } from 'react-router-dom';
import { NameProvider } from './NameContext';
import FirstComponent from './FirstComponent';
import SecondComponent from './SecondComponent';
import ThirdComponent from './ThirdComponent';

function App() {
  return (
    <Router>
      <NameProvider>
        <Routes>
          <Route path="*" element={<FirstComponent />} />
          <Route path="/app/second" element={<SecondComponent />} />
          <Route path="/app/third" element={<ThirdComponent />} />
          </Routes>
      </NameProvider>
    </Router>
  );
}

export default App;

































// import { useState } from 'react'
// import reactLogo from './assets/react.svg'
// import viteLogo from '/vite.svg'
// import './App.css'

// function App() {
//   const [count, setCount] = useState(0)

//   return (
//     <>
//       <div>
//         <a href="https://vitejs.dev" target="_blank">
//           <img src={viteLogo} className="logo" alt="Vite logo" />
//         </a>
//         <a href="https://react.dev" target="_blank">
//           <img src={reactLogo} className="logo react" alt="React logo" />
//         </a>
//       </div>
//       <h1>Vite + React</h1>
//       <div className="card">
//         <button onClick={() => setCount((count) => count + 1)}>
//           count is {count}
//         </button>
//         <p>
//           Edit <code>src/App.jsx</code> and save to test HMR
//         </p>
//       </div>
//       <p className="read-the-docs">
//         Click on the Vite and React logos to learn more
//         second line
//         third line
//         in app.jsx
//       </p>
//     </>
//   )
// }

// export default App
